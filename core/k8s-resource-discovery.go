package core

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/build/kubernetes/api"
	"google.golang.org/grpc"
	"istio-service-mesh/constants"
	pb "istio-service-mesh/core/proto"
	"istio-service-mesh/types"
	"istio-service-mesh/utils"
	v1 "k8s.io/api/apps/v1"
	autoscale "k8s.io/api/autoscaling/v2beta2"
	batch "k8s.io/api/batch/v1"
	"k8s.io/api/batch/v1beta1"
	batchv1 "k8s.io/api/batch/v1beta1"
	v2 "k8s.io/api/core/v1"
	rbac "k8s.io/api/rbac/v1"
	storage "k8s.io/api/storage/v1"
	"math/rand"
	"reflect"
	"sigs.k8s.io/yaml"
	"strconv"
	"sync"
)

type GrpcConn struct {
	Connection *grpc.ClientConn
	ProjectId  string
	CompanyId  string
	token      string
}

var serviceTemplates []*types.ServiceTemplate

func (s *Server) GetK8SResource(ctx context.Context, request *pb.K8SResourceRequest) (response *pb.K8SResourceResponse, err error) {
	response = new(pb.K8SResourceResponse)
	utils.Info.Println(reflect.TypeOf(ctx))

	if request.CompanyId == "" || request.ProjectId == "" {
		return &pb.K8SResourceResponse{}, errors.New("projectId or companyId must not be empty")
	}

	conn, err := grpc.DialContext(ctx, constants.K8sEngineGRPCURL, grpc.WithInsecure())
	if err != nil {
		utils.Error.Println(err)
		return &pb.K8SResourceResponse{}, err
	}
	defer conn.Close()

	grpcConn := &GrpcConn{
		Connection: conn,
		ProjectId:  request.ProjectId,
		CompanyId:  request.CompanyId,
		token:      request.Token,
	}

	namespaces := request.Namespaces

	var wg sync.WaitGroup
	for _, namespace := range namespaces {
		//deployments
		deploymentList, err := grpcConn.getAllDeployments(ctx, namespace)
		if err != nil {
			return &pb.K8SResourceResponse{}, err
		}
		err = grpcConn.deploymentk8sToCp(ctx, deploymentList.Items, &wg)
		if err != nil {
			return &pb.K8SResourceResponse{}, err
		}
		//deployments

		//statefulsets
		//statefulSetList, err := grpcConn.getAllStatefulsets(ctx, namespace)
		//if err != nil {
		//	return &pb.K8SResourceResponse{}, err
		//}
		//err = grpcConn.statefulsetsK8sToCp(ctx, statefulSetList.Items)
		//if err != nil {
		//	return &pb.K8SResourceResponse{}, err
		//}
		////statefulsets
		//
		////daemonsets
		//daemonSetList, err := grpcConn.getAllDaemonsets(ctx, namespace)
		//if err != nil {
		//	return &pb.K8SResourceResponse{}, err
		//}
		//err = grpcConn.daemonsetK8sToCp(ctx, daemonSetList.Items)
		//if err != nil {
		//	return &pb.K8SResourceResponse{}, err
		//}
		////daemonsets
		//
		////cronjobs
		//cronJobList, err := grpcConn.getAllCronjobs(ctx, namespace)
		//if err != nil {
		//	return &pb.K8SResourceResponse{}, err
		//}
		//err = grpcConn.cronjobK8sToCp(ctx, cronJobList.Items)
		//if err != nil {
		//	return &pb.K8SResourceResponse{}, err
		//}
		////cronjobs
		//
		////jobs
		//JobList, err := grpcConn.getAllJobs(ctx, namespace)
		//if err != nil {
		//	return &pb.K8SResourceResponse{}, err
		//}
		//err = grpcConn.jobK8sToCp(ctx, JobList.Items)
		//if err != nil {
		//	return &pb.K8SResourceResponse{}, err
		//}
		//jobs
	}

	wg.Wait()
	bytes, err := json.Marshal(serviceTemplates)
	if err != nil {
		return &pb.K8SResourceResponse{}, err
	}

	response.Resource = bytes
	return response, err
}

func (conn *GrpcConn) jobK8sToCp(ctx context.Context, jobs []batch.Job) error {
	for _, job := range jobs {
		utils.Info.Printf("Resolving dependencies of job :%v within namespace : %v", job.Name, job.Namespace)
		jobTemp, err := getCpConvertedTemplate(job, job.Kind)
		if err != nil {
			utils.Error.Println(err)
			return err
		}
		namespace := job.Namespace
		if job.Spec.Template.Spec.ServiceAccountName != "" {
			svcname := job.Spec.Template.Spec.ServiceAccountName
			svcaccount, err := conn.getSvcAccount(ctx, svcname, namespace)
			if err != nil {
				return err
			}
			svcAccTemp, err := getCpConvertedTemplate(svcaccount, svcaccount.Kind)
			if err != nil {
				utils.Error.Println(err)
				return err
			}
			if !isAlreadyExist(svcAccTemp.NameSpace, svcAccTemp.ServiceSubType, svcAccTemp.Name) {
				rbacDependencies, err := conn.getK8sRbacResources(ctx, namespace, svcaccount, svcAccTemp)
				if err != nil {
					utils.Error.Println(err)
					return err
				}

				var strpointer = new(string)
				*strpointer = "serviceaccount"
				for _, rbacTemp := range rbacDependencies {
					if *rbacTemp.ServiceSubType == *strpointer {
						jobTemp.BeforeServices = append(jobTemp.BeforeServices, rbacTemp.ServiceId)
						rbacTemp.AfterServices = append(rbacTemp.AfterServices, jobTemp.ServiceId)
					}
					serviceTemplates = append(serviceTemplates, rbacTemp)
				}
			} else {
				for _, service := range serviceTemplates {
					if *service.NameSpace == *svcAccTemp.NameSpace && *service.ServiceSubType == *svcAccTemp.ServiceSubType && *service.Name == *svcAccTemp.Name {
						jobTemp.BeforeServices = append(jobTemp.BeforeServices, service.ServiceId)
						service.AfterServices = append(service.AfterServices, jobTemp.ServiceId)
					}
				}
			}
		}

		//image pull secrets
		for _, objRef := range job.Spec.Template.Spec.ImagePullSecrets {
			secretname := objRef.Name
			secret, err := conn.getSecret(ctx, secretname, namespace)
			if err != nil {
				return err
			}
			secretTemp, err := getCpConvertedTemplate(secret, secret.Kind)
			if err != nil {
				utils.Error.Println(err)
				return err
			}
			if !isAlreadyExist(secretTemp.NameSpace, secretTemp.ServiceSubType, secretTemp.Name) {
				jobTemp.BeforeServices = append(jobTemp.BeforeServices, secretTemp.ServiceId)
				secretTemp.AfterServices = append(secretTemp.AfterServices, jobTemp.ServiceId)
				serviceTemplates = append(serviceTemplates, secretTemp)
			} else {
				isSameJob := false
				for _, serviceId := range secretTemp.AfterServices {
					if *serviceId == *jobTemp.ServiceId {
						isSameJob = true
						break
					}
				}
				if !isSameJob {
					jobTemp.BeforeServices = append(jobTemp.BeforeServices, secretTemp.ServiceId)
					secretTemp.AfterServices = append(secretTemp.AfterServices, jobTemp.ServiceId)
				}
			}
		}

		//hpa finding
		hpaList, err := conn.getAllHpas(ctx, namespace)
		if err != nil {
			return err
		}
		if len(hpaList.Items) > 0 {
			for _, hpa := range hpaList.Items {
				if hpa.Spec.ScaleTargetRef.APIVersion == job.APIVersion && hpa.Spec.ScaleTargetRef.Kind == job.Kind && hpa.Spec.ScaleTargetRef.Name == job.Name {
					hpaTemplate, err := getCpConvertedTemplate(hpa, hpa.Kind)
					if err != nil {
						utils.Error.Println(err)
						return err
					}
					hpaTemplate.BeforeServices = append(hpaTemplate.BeforeServices, jobTemp.ServiceId)
					jobTemp.AfterServices = append(jobTemp.AfterServices, hpaTemplate.ServiceId)
					serviceTemplates = append(serviceTemplates, hpaTemplate)
				}
			}
		}

		//kubernetes service depecndency findings
		var kubeSvcList *v2.ServiceList
		for key, value := range job.Spec.Template.Labels {
			kubeSvcList, err = conn.getKubernetesServices(ctx, key, value, namespace)
			if err != nil {
				return err
			}
		}

		//container dependency finding
		for _, container := range job.Spec.Template.Spec.Containers {

			//resolving dependencies of kubernetes service
			if kubeSvcList != nil && len(kubeSvcList.Items) > 0 {
				for _, kubeSvc := range kubeSvcList.Items {
					if isPortMatched(&kubeSvc, &container) {
						k8serviceTemp, err := getCpConvertedTemplate(kubeSvc, kubeSvc.Kind)
						if err != nil {
							utils.Error.Println(err)
							return err
						}
						if !isAlreadyExist(k8serviceTemp.NameSpace, k8serviceTemp.ServiceSubType, k8serviceTemp.Name) {
							k8serviceTemp.BeforeServices = append(k8serviceTemp.BeforeServices, jobTemp.ServiceId)
							jobTemp.AfterServices = append(jobTemp.AfterServices, k8serviceTemp.ServiceId)
							serviceTemplates = append(serviceTemplates, k8serviceTemp)
						} else {
							isSameJob := false
							for _, serviceId := range k8serviceTemp.BeforeServices {
								if *serviceId == *jobTemp.ServiceId {
									isSameJob = true
									break
								}
							}
							if !isSameJob {
								k8serviceTemp.BeforeServices = append(k8serviceTemp.BeforeServices, jobTemp.ServiceId)
								jobTemp.AfterServices = append(jobTemp.AfterServices, k8serviceTemp.ServiceId)
							}
						}
					}
				}
			}

			//discovering secret and config maps in deployment containers
			for _, env := range container.Env {
				if env.ValueFrom != nil && env.ValueFrom.SecretKeyRef != nil {
					secretname := env.ValueFrom.SecretKeyRef.Name
					secret, err := conn.getSecret(ctx, secretname, namespace)
					if err != nil {
						return err
					}
					secretTemp, err := getCpConvertedTemplate(secret, secret.Kind)
					if err != nil {
						utils.Error.Println(err)
						return err
					}
					if !isAlreadyExist(secretTemp.NameSpace, secretTemp.ServiceSubType, secretTemp.Name) {
						jobTemp.BeforeServices = append(jobTemp.BeforeServices, secretTemp.ServiceId)
						secretTemp.AfterServices = append(secretTemp.AfterServices, jobTemp.ServiceId)
						serviceTemplates = append(serviceTemplates, secretTemp)
					} else {
						isSameJob := false
						for _, serviceId := range secretTemp.AfterServices {
							if *serviceId == *jobTemp.ServiceId {
								isSameJob = true
								break
							}
						}
						if !isSameJob {
							jobTemp.BeforeServices = append(jobTemp.BeforeServices, secretTemp.ServiceId)
							secretTemp.AfterServices = append(secretTemp.AfterServices, jobTemp.ServiceId)
						}
					}
				} else if env.ValueFrom != nil && env.ValueFrom.ConfigMapKeyRef != nil {
					configmapname := env.ValueFrom.ConfigMapKeyRef.Name
					configmap, err := conn.getConfigMap(ctx, configmapname, namespace)
					if err != nil {
						return err
					}
					configmapTemp, err := getCpConvertedTemplate(configmap, configmap.Kind)
					if err != nil {
						utils.Error.Println(err)
						return err
					}
					if !isAlreadyExist(configmapTemp.NameSpace, configmapTemp.ServiceSubType, configmapTemp.Name) {
						jobTemp.BeforeServices = append(jobTemp.BeforeServices, configmapTemp.ServiceId)
						configmapTemp.AfterServices = append(configmapTemp.AfterServices, jobTemp.ServiceId)
						serviceTemplates = append(serviceTemplates, configmapTemp)
					} else {
						isSameJob := false
						for _, serviceId := range configmapTemp.AfterServices {
							if *serviceId == *jobTemp.ServiceId {
								isSameJob = true
								break
							}
						}
						if !isSameJob {
							jobTemp.BeforeServices = append(jobTemp.BeforeServices, configmapTemp.ServiceId)
							configmapTemp.AfterServices = append(configmapTemp.AfterServices, jobTemp.ServiceId)
						}
					}
				}
			}
		}

		//volume dependency finding
		for _, vol := range job.Spec.Template.Spec.Volumes {
			if vol.Secret != nil {
				secretname := vol.Secret.SecretName
				secret, err := conn.getSecret(ctx, secretname, namespace)
				if err != nil {
					return err
				}
				secretTemp, err := getCpConvertedTemplate(secret, secret.Kind)
				if err != nil {
					utils.Error.Println(err)
					return err
				}
				if !isAlreadyExist(secretTemp.NameSpace, secretTemp.ServiceSubType, secretTemp.Name) {
					jobTemp.BeforeServices = append(jobTemp.BeforeServices, secretTemp.ServiceId)
					secretTemp.AfterServices = append(secretTemp.AfterServices, jobTemp.ServiceId)
					serviceTemplates = append(serviceTemplates, secretTemp)
				} else {
					isSameJob := false
					for _, serviceId := range secretTemp.AfterServices {
						if *serviceId == *jobTemp.ServiceId {
							isSameJob = true
							break
						}
					}
					if !isSameJob {
						jobTemp.BeforeServices = append(jobTemp.BeforeServices, secretTemp.ServiceId)
						secretTemp.AfterServices = append(secretTemp.AfterServices, jobTemp.ServiceId)
					}
				}
			} else if vol.ConfigMap != nil {
				configmapname := vol.ConfigMap.Name
				configmap, err := conn.getConfigMap(ctx, configmapname, namespace)
				if err != nil {
					return err
				}
				configmapTemp, err := getCpConvertedTemplate(configmap, configmap.Kind)
				if err != nil {
					utils.Error.Println(err)
					return err
				}
				if !isAlreadyExist(configmapTemp.NameSpace, configmapTemp.ServiceSubType, configmapTemp.Name) {
					jobTemp.BeforeServices = append(jobTemp.BeforeServices, configmapTemp.ServiceId)
					configmapTemp.AfterServices = append(configmapTemp.AfterServices, jobTemp.ServiceId)
					serviceTemplates = append(serviceTemplates, configmapTemp)
				} else {
					isSameJob := false
					for _, serviceId := range configmapTemp.AfterServices {
						if *serviceId == *jobTemp.ServiceId {
							isSameJob = true
							break
						}
					}
					if !isSameJob {
						jobTemp.BeforeServices = append(jobTemp.BeforeServices, configmapTemp.ServiceId)
						configmapTemp.AfterServices = append(configmapTemp.AfterServices, jobTemp.ServiceId)
					}
				}
			} else if vol.PersistentVolumeClaim != nil {
				pvcname := vol.PersistentVolumeClaim.ClaimName
				pvc, err := conn.getPvc(ctx, pvcname, namespace)
				if err != nil {
					return err
				}
				pvcTemp, err := getCpConvertedTemplate(pvc, pvc.Kind)
				if err != nil {
					utils.Error.Println(err)
					return err
				}
				if !isAlreadyExist(pvcTemp.NameSpace, pvcTemp.ServiceSubType, pvcTemp.Name) {
					jobTemp.BeforeServices = append(jobTemp.BeforeServices, pvcTemp.ServiceId)
					pvcTemp.AfterServices = append(pvcTemp.AfterServices, jobTemp.ServiceId)
					serviceTemplates = append(serviceTemplates, pvcTemp)
				} else {
					for _, service := range serviceTemplates {
						if *service.NameSpace == *pvcTemp.NameSpace && *service.ServiceSubType == *pvcTemp.ServiceSubType && *service.Name == *pvcTemp.Name {
							jobTemp.BeforeServices = append(jobTemp.BeforeServices, service.ServiceId)
							service.AfterServices = append(service.AfterServices, jobTemp.ServiceId)
						}
					}
				}
				if pvc.Spec.StorageClassName != nil {
					storageClassName := *pvc.Spec.StorageClassName
					storageClass, err := conn.getStorageClass(ctx, storageClassName, namespace)
					if err != nil {
						utils.Error.Println(err)
						return err
					}
					storageClassTemp, err := getCpConvertedTemplate(storageClass, storageClass.Kind)
					if err != nil {
						utils.Error.Println(err)
						return err
					}
					if !isAlreadyExist(storageClassTemp.NameSpace, storageClassTemp.ServiceSubType, storageClassTemp.Name) {
						pvcTemp.BeforeServices = append(pvcTemp.BeforeServices, storageClassTemp.ServiceId)
						storageClassTemp.AfterServices = append(storageClassTemp.AfterServices, pvcTemp.ServiceId)
						serviceTemplates = append(serviceTemplates, storageClassTemp)
					} else {
						isPVCexist := false
						for _, serviceId := range storageClassTemp.AfterServices {
							if *serviceId == *pvcTemp.ServiceId {
								isPVCexist = true
								break
							}
						}
						if !isPVCexist {
							pvcTemp.BeforeServices = append(pvcTemp.BeforeServices, storageClassTemp.ServiceId)
							storageClassTemp.AfterServices = append(storageClassTemp.AfterServices, pvcTemp.ServiceId)
						}
					}
				}
			}

			if vol.Projected != nil {
				for _, source := range vol.Projected.Sources {
					if source.ConfigMap != nil {
						configmapname := vol.ConfigMap.Name
						configmap, err := conn.getConfigMap(ctx, configmapname, namespace)
						if err != nil {
							return err
						}
						configmapTemp, err := getCpConvertedTemplate(configmap, configmap.Kind)
						if err != nil {
							utils.Error.Println(err)
							return err
						}
						if !isAlreadyExist(configmapTemp.NameSpace, configmapTemp.ServiceSubType, configmapTemp.Name) {
							jobTemp.BeforeServices = append(jobTemp.BeforeServices, configmapTemp.ServiceId)
							configmapTemp.AfterServices = append(configmapTemp.AfterServices, jobTemp.ServiceId)
							serviceTemplates = append(serviceTemplates, configmapTemp)
						} else {
							isSameJob := false
							for _, serviceId := range configmapTemp.AfterServices {
								if *serviceId == *jobTemp.ServiceId {
									isSameJob = true
									break
								}
							}
							if !isSameJob {
								jobTemp.BeforeServices = append(jobTemp.BeforeServices, configmapTemp.ServiceId)
								configmapTemp.AfterServices = append(configmapTemp.AfterServices, jobTemp.ServiceId)
							}
						}
					} else if source.Secret != nil {
						secretname := vol.Secret.SecretName
						secret, err := conn.getSecret(ctx, secretname, namespace)
						if err != nil {
							return err
						}
						secretTemp, err := getCpConvertedTemplate(secret, secret.Kind)
						if err != nil {
							utils.Error.Println(err)
							return err
						}
						if !isAlreadyExist(secretTemp.NameSpace, secretTemp.ServiceSubType, secretTemp.Name) {
							jobTemp.BeforeServices = append(jobTemp.BeforeServices, secretTemp.ServiceId)
							secretTemp.AfterServices = append(secretTemp.AfterServices, jobTemp.ServiceId)
							serviceTemplates = append(serviceTemplates, secretTemp)
						} else {
							isSameJob := false
							for _, serviceId := range secretTemp.AfterServices {
								if *serviceId == *jobTemp.ServiceId {
									isSameJob = true
									break
								}
							}
							if !isSameJob {
								jobTemp.BeforeServices = append(jobTemp.BeforeServices, secretTemp.ServiceId)
								secretTemp.AfterServices = append(secretTemp.AfterServices, jobTemp.ServiceId)
							}
						}
					}
				}
			}
		}

		serviceTemplates = append(serviceTemplates, jobTemp)

	}
	return nil
}

func (conn *GrpcConn) cronjobK8sToCp(ctx context.Context, cronjobs []v1beta1.CronJob) error {
	for _, cronjob := range cronjobs {
		utils.Info.Printf("Resolving dependencies of cronjob :%v within namespace : %v", cronjob.Name, cronjob.Namespace)
		cronjobTemp, err := getCpConvertedTemplate(cronjob, cronjob.Kind)
		if err != nil {
			utils.Error.Println(err)
			return err
		}
		namespace := cronjob.Namespace
		if cronjob.Spec.JobTemplate.Spec.Template.Spec.ServiceAccountName != "" {
			svcname := cronjob.Spec.JobTemplate.Spec.Template.Spec.ServiceAccountName
			svcaccount, err := conn.getSvcAccount(ctx, svcname, namespace)
			if err != nil {
				return err
			}
			svcAccTemp, err := getCpConvertedTemplate(svcaccount, svcaccount.Kind)
			if err != nil {
				utils.Error.Println(err)
				return err
			}
			if !isAlreadyExist(svcAccTemp.NameSpace, svcAccTemp.ServiceSubType, svcAccTemp.Name) {
				rbacDependencies, err := conn.getK8sRbacResources(ctx, namespace, svcaccount, svcAccTemp)
				if err != nil {
					utils.Error.Println(err)
					return err
				}

				var strpointer = new(string)
				*strpointer = "serviceaccount"
				for _, rbacTemp := range rbacDependencies {
					if *rbacTemp.ServiceSubType == *strpointer {
						cronjobTemp.BeforeServices = append(cronjobTemp.BeforeServices, rbacTemp.ServiceId)
						rbacTemp.AfterServices = append(rbacTemp.AfterServices, cronjobTemp.ServiceId)
					}
					serviceTemplates = append(serviceTemplates, rbacTemp)
				}
			} else {
				for _, service := range serviceTemplates {
					if *service.NameSpace == *svcAccTemp.NameSpace && *service.ServiceSubType == *svcAccTemp.ServiceSubType && *service.Name == *svcAccTemp.Name {
						cronjobTemp.BeforeServices = append(cronjobTemp.BeforeServices, service.ServiceId)
						service.AfterServices = append(service.AfterServices, cronjobTemp.ServiceId)
					}
				}
			}
		}

		//image pull secrets
		for _, objRef := range cronjob.Spec.JobTemplate.Spec.Template.Spec.ImagePullSecrets {
			secretname := objRef.Name
			secret, err := conn.getSecret(ctx, secretname, namespace)
			if err != nil {
				return err
			}
			secretTemp, err := getCpConvertedTemplate(secret, secret.Kind)
			if err != nil {
				utils.Error.Println(err)
				return err
			}
			if !isAlreadyExist(secretTemp.NameSpace, secretTemp.ServiceSubType, secretTemp.Name) {
				cronjobTemp.BeforeServices = append(cronjobTemp.BeforeServices, secretTemp.ServiceId)
				secretTemp.AfterServices = append(secretTemp.AfterServices, cronjobTemp.ServiceId)
				serviceTemplates = append(serviceTemplates, secretTemp)
			} else {
				isSameCronJob := false
				for _, serviceId := range secretTemp.AfterServices {
					if *serviceId == *cronjobTemp.ServiceId {
						isSameCronJob = true
						break
					}
				}
				if !isSameCronJob {
					cronjobTemp.BeforeServices = append(cronjobTemp.BeforeServices, secretTemp.ServiceId)
					secretTemp.AfterServices = append(secretTemp.AfterServices, cronjobTemp.ServiceId)
				}
			}
		}

		hpaList, err := conn.getAllHpas(ctx, namespace)
		if err != nil {
			return err
		}
		if len(hpaList.Items) > 0 {
			for _, hpa := range hpaList.Items {
				if hpa.Spec.ScaleTargetRef.APIVersion == cronjob.APIVersion && hpa.Spec.ScaleTargetRef.Kind == cronjob.Kind && hpa.Spec.ScaleTargetRef.Name == cronjob.Name {
					hpaTemplate, err := getCpConvertedTemplate(hpa, hpa.Kind)
					if err != nil {
						utils.Error.Println(err)
						return err
					}
					hpaTemplate.BeforeServices = append(hpaTemplate.BeforeServices, cronjobTemp.ServiceId)
					cronjobTemp.AfterServices = append(cronjobTemp.AfterServices, hpaTemplate.ServiceId)
					serviceTemplates = append(serviceTemplates, hpaTemplate)
				}
			}
		}

		//kubernetes service depecndency findings
		var kubeSvcList *v2.ServiceList
		for key, value := range cronjob.Spec.JobTemplate.Spec.Template.Labels {
			kubeSvcList, err = conn.getKubernetesServices(ctx, key, value, namespace)
			if err != nil {
				return err
			}
		}

		//container dependency finding
		for _, container := range cronjob.Spec.JobTemplate.Spec.Template.Spec.Containers {

			//resolving dependencies of kubernetes service
			if kubeSvcList != nil && len(kubeSvcList.Items) > 0 {
				for _, kubeSvc := range kubeSvcList.Items {
					if isPortMatched(&kubeSvc, &container) {
						k8serviceTemp, err := getCpConvertedTemplate(kubeSvc, kubeSvc.Kind)
						if err != nil {
							utils.Error.Println(err)
							return err
						}
						if !isAlreadyExist(k8serviceTemp.NameSpace, k8serviceTemp.ServiceSubType, k8serviceTemp.Name) {
							k8serviceTemp.BeforeServices = append(k8serviceTemp.BeforeServices, cronjobTemp.ServiceId)
							cronjobTemp.AfterServices = append(cronjobTemp.AfterServices, k8serviceTemp.ServiceId)
							serviceTemplates = append(serviceTemplates, k8serviceTemp)
						} else {
							isSameCronJob := false
							for _, serviceId := range k8serviceTemp.BeforeServices {
								if *serviceId == *cronjobTemp.ServiceId {
									isSameCronJob = true
									break
								}
							}
							if !isSameCronJob {
								k8serviceTemp.BeforeServices = append(k8serviceTemp.BeforeServices, cronjobTemp.ServiceId)
								cronjobTemp.AfterServices = append(cronjobTemp.AfterServices, k8serviceTemp.ServiceId)
							}
						}
					}
				}
			}

			//discovering secret and config maps in deployment containers
			for _, env := range container.Env {
				if env.ValueFrom != nil && env.ValueFrom.SecretKeyRef != nil {
					secretname := env.ValueFrom.SecretKeyRef.Name
					secret, err := conn.getSecret(ctx, secretname, namespace)
					if err != nil {
						return err
					}
					secretTemp, err := getCpConvertedTemplate(secret, secret.Kind)
					if err != nil {
						utils.Error.Println(err)
						return err
					}
					if !isAlreadyExist(secretTemp.NameSpace, secretTemp.ServiceSubType, secretTemp.Name) {
						cronjobTemp.BeforeServices = append(cronjobTemp.BeforeServices, secretTemp.ServiceId)
						secretTemp.AfterServices = append(secretTemp.AfterServices, cronjobTemp.ServiceId)
						serviceTemplates = append(serviceTemplates, secretTemp)
					} else {
						isSameCronJob := false
						for _, serviceId := range secretTemp.AfterServices {
							if *serviceId == *cronjobTemp.ServiceId {
								isSameCronJob = true
								break
							}
						}
						if !isSameCronJob {
							cronjobTemp.BeforeServices = append(cronjobTemp.BeforeServices, secretTemp.ServiceId)
							secretTemp.AfterServices = append(secretTemp.AfterServices, cronjobTemp.ServiceId)
						}
					}
				} else if env.ValueFrom != nil && env.ValueFrom.ConfigMapKeyRef != nil {
					configmapname := env.ValueFrom.ConfigMapKeyRef.Name
					configmap, err := conn.getConfigMap(ctx, configmapname, namespace)
					if err != nil {
						return err
					}
					configmapTemp, err := getCpConvertedTemplate(configmap, configmap.Kind)
					if err != nil {
						utils.Error.Println(err)
						return err
					}
					if !isAlreadyExist(configmapTemp.NameSpace, configmapTemp.ServiceSubType, configmapTemp.Name) {
						cronjobTemp.BeforeServices = append(cronjobTemp.BeforeServices, configmapTemp.ServiceId)
						configmapTemp.AfterServices = append(configmapTemp.AfterServices, cronjobTemp.ServiceId)
						serviceTemplates = append(serviceTemplates, configmapTemp)
					} else {
						isSameCronJob := false
						for _, serviceId := range configmapTemp.AfterServices {
							if *serviceId == *cronjobTemp.ServiceId {
								isSameCronJob = true
								break
							}
						}
						if !isSameCronJob {
							cronjobTemp.BeforeServices = append(cronjobTemp.BeforeServices, configmapTemp.ServiceId)
							configmapTemp.AfterServices = append(configmapTemp.AfterServices, cronjobTemp.ServiceId)
						}
					}
				}
			}
		}

		//volume dependency finding
		for _, vol := range cronjob.Spec.JobTemplate.Spec.Template.Spec.Volumes {
			if vol.Secret != nil {
				secretname := vol.Secret.SecretName
				secret, err := conn.getSecret(ctx, secretname, namespace)
				if err != nil {
					return err
				}
				secretTemp, err := getCpConvertedTemplate(secret, secret.Kind)
				if err != nil {
					utils.Error.Println(err)
					return err
				}
				if !isAlreadyExist(secretTemp.NameSpace, secretTemp.ServiceSubType, secretTemp.Name) {
					cronjobTemp.BeforeServices = append(cronjobTemp.BeforeServices, secretTemp.ServiceId)
					secretTemp.AfterServices = append(secretTemp.AfterServices, cronjobTemp.ServiceId)
					serviceTemplates = append(serviceTemplates, secretTemp)
				} else {
					isSameCronJob := false
					for _, serviceId := range secretTemp.AfterServices {
						if *serviceId == *cronjobTemp.ServiceId {
							isSameCronJob = true
							break
						}
					}
					if !isSameCronJob {
						cronjobTemp.BeforeServices = append(cronjobTemp.BeforeServices, secretTemp.ServiceId)
						secretTemp.AfterServices = append(secretTemp.AfterServices, cronjobTemp.ServiceId)
					}
				}
			} else if vol.ConfigMap != nil {
				configmapname := vol.ConfigMap.Name
				configmap, err := conn.getConfigMap(ctx, configmapname, namespace)
				if err != nil {
					return err
				}
				configmapTemp, err := getCpConvertedTemplate(configmap, configmap.Kind)
				if err != nil {
					utils.Error.Println(err)
					return err
				}
				if !isAlreadyExist(configmapTemp.NameSpace, configmapTemp.ServiceSubType, configmapTemp.Name) {
					cronjobTemp.BeforeServices = append(cronjobTemp.BeforeServices, configmapTemp.ServiceId)
					configmapTemp.AfterServices = append(configmapTemp.AfterServices, cronjobTemp.ServiceId)
					serviceTemplates = append(serviceTemplates, configmapTemp)
				} else {
					isSameCronJob := false
					for _, serviceId := range configmapTemp.AfterServices {
						if *serviceId == *cronjobTemp.ServiceId {
							isSameCronJob = true
							break
						}
					}
					if !isSameCronJob {
						cronjobTemp.BeforeServices = append(cronjobTemp.BeforeServices, configmapTemp.ServiceId)
						configmapTemp.AfterServices = append(configmapTemp.AfterServices, cronjobTemp.ServiceId)
					}
				}
			} else if vol.PersistentVolumeClaim != nil {
				pvcname := vol.PersistentVolumeClaim.ClaimName
				pvc, err := conn.getPvc(ctx, pvcname, namespace)
				if err != nil {
					return err
				}
				pvcTemp, err := getCpConvertedTemplate(pvc, pvc.Kind)
				if err != nil {
					utils.Error.Println(err)
					return err
				}
				if !isAlreadyExist(pvcTemp.NameSpace, pvcTemp.ServiceSubType, pvcTemp.Name) {
					cronjobTemp.BeforeServices = append(cronjobTemp.BeforeServices, pvcTemp.ServiceId)
					pvcTemp.AfterServices = append(pvcTemp.AfterServices, cronjobTemp.ServiceId)
					serviceTemplates = append(serviceTemplates, pvcTemp)
				} else {
					for _, service := range serviceTemplates {
						if *service.NameSpace == *pvcTemp.NameSpace && *service.ServiceSubType == *pvcTemp.ServiceSubType && *service.Name == *pvcTemp.Name {
							cronjobTemp.BeforeServices = append(cronjobTemp.BeforeServices, service.ServiceId)
							service.AfterServices = append(service.AfterServices, cronjobTemp.ServiceId)
						}
					}
				}
				if pvc.Spec.StorageClassName != nil {
					storageClassName := *pvc.Spec.StorageClassName
					storageClass, err := conn.getStorageClass(ctx, storageClassName, namespace)
					if err != nil {
						utils.Error.Println(err)
						return err
					}
					storageClassTemp, err := getCpConvertedTemplate(storageClass, storageClass.Kind)
					if err != nil {
						utils.Error.Println(err)
						return err
					}
					if !isAlreadyExist(storageClassTemp.NameSpace, storageClassTemp.ServiceSubType, storageClassTemp.Name) {
						pvcTemp.BeforeServices = append(pvcTemp.BeforeServices, storageClassTemp.ServiceId)
						storageClassTemp.AfterServices = append(storageClassTemp.AfterServices, pvcTemp.ServiceId)
						serviceTemplates = append(serviceTemplates, storageClassTemp)
					} else {
						isPVCexist := false
						for _, serviceId := range storageClassTemp.AfterServices {
							if *serviceId == *pvcTemp.ServiceId {
								isPVCexist = true
								break
							}
						}
						if !isPVCexist {
							pvcTemp.BeforeServices = append(pvcTemp.BeforeServices, storageClassTemp.ServiceId)
							storageClassTemp.AfterServices = append(storageClassTemp.AfterServices, pvcTemp.ServiceId)
						}
					}
				}
			}

			if vol.Projected != nil {
				for _, source := range vol.Projected.Sources {
					if source.ConfigMap != nil {
						configmapname := vol.ConfigMap.Name
						configmap, err := conn.getConfigMap(ctx, configmapname, namespace)
						if err != nil {
							return err
						}
						configmapTemp, err := getCpConvertedTemplate(configmap, configmap.Kind)
						if err != nil {
							utils.Error.Println(err)
							return err
						}
						if !isAlreadyExist(configmapTemp.NameSpace, configmapTemp.ServiceSubType, configmapTemp.Name) {
							cronjobTemp.BeforeServices = append(cronjobTemp.BeforeServices, configmapTemp.ServiceId)
							configmapTemp.AfterServices = append(configmapTemp.AfterServices, cronjobTemp.ServiceId)
							serviceTemplates = append(serviceTemplates, configmapTemp)
						} else {
							isSameCronJob := false
							for _, serviceId := range configmapTemp.AfterServices {
								if *serviceId == *cronjobTemp.ServiceId {
									isSameCronJob = true
									break
								}
							}
							if !isSameCronJob {
								cronjobTemp.BeforeServices = append(cronjobTemp.BeforeServices, configmapTemp.ServiceId)
								configmapTemp.AfterServices = append(configmapTemp.AfterServices, cronjobTemp.ServiceId)
							}
						}
					} else if source.Secret != nil {
						secretname := vol.Secret.SecretName
						secret, err := conn.getSecret(ctx, secretname, namespace)
						if err != nil {
							return err
						}
						secretTemp, err := getCpConvertedTemplate(secret, secret.Kind)
						if err != nil {
							utils.Error.Println(err)
							return err
						}
						if !isAlreadyExist(secretTemp.NameSpace, secretTemp.ServiceSubType, secretTemp.Name) {
							cronjobTemp.BeforeServices = append(cronjobTemp.BeforeServices, secretTemp.ServiceId)
							secretTemp.AfterServices = append(secretTemp.AfterServices, cronjobTemp.ServiceId)
							serviceTemplates = append(serviceTemplates, secretTemp)
						} else {
							isSameCronJob := false
							for _, serviceId := range secretTemp.AfterServices {
								if *serviceId == *cronjobTemp.ServiceId {
									isSameCronJob = true
									break
								}
							}
							if !isSameCronJob {
								cronjobTemp.BeforeServices = append(cronjobTemp.BeforeServices, secretTemp.ServiceId)
								secretTemp.AfterServices = append(secretTemp.AfterServices, cronjobTemp.ServiceId)
							}
						}
					}
				}
			}
		}

		serviceTemplates = append(serviceTemplates, cronjobTemp)
	}
	return nil
}

func (conn *GrpcConn) daemonsetK8sToCp(ctx context.Context, daemonsets []v1.DaemonSet) error {
	for _, daemonset := range daemonsets {
		utils.Info.Printf("Resolving dependencies of daemonset :%v within namespace : %v", daemonset.Name, daemonset.Namespace)
		daemonsetTemp, err := getCpConvertedTemplate(daemonset, daemonset.Kind)
		if err != nil {
			utils.Error.Println(err)
			return err
		}
		namespace := daemonset.Namespace
		if daemonset.Spec.Template.Spec.ServiceAccountName != "" {
			svcname := daemonset.Spec.Template.Spec.ServiceAccountName
			svcaccount, err := conn.getSvcAccount(ctx, svcname, namespace)
			if err != nil {
				return err
			}
			svcAccTemp, err := getCpConvertedTemplate(svcaccount, svcaccount.Kind)
			if err != nil {
				utils.Error.Println(err)
				return err
			}
			if !isAlreadyExist(svcAccTemp.NameSpace, svcAccTemp.ServiceSubType, svcAccTemp.Name) {
				rbacDependencies, err := conn.getK8sRbacResources(ctx, namespace, svcaccount, svcAccTemp)
				if err != nil {
					utils.Error.Println(err)
					return err
				}

				var strpointer = new(string)
				*strpointer = "serviceaccount"
				for _, rbacTemp := range rbacDependencies {
					if *rbacTemp.ServiceSubType == *strpointer {
						daemonsetTemp.BeforeServices = append(daemonsetTemp.BeforeServices, rbacTemp.ServiceId)
						rbacTemp.AfterServices = append(rbacTemp.AfterServices, daemonsetTemp.ServiceId)
					}
					serviceTemplates = append(serviceTemplates, rbacTemp)
				}
			} else {
				for _, service := range serviceTemplates {
					if *service.NameSpace == *svcAccTemp.NameSpace && *service.ServiceSubType == *svcAccTemp.ServiceSubType && *service.Name == *svcAccTemp.Name {
						daemonsetTemp.BeforeServices = append(daemonsetTemp.BeforeServices, service.ServiceId)
						service.AfterServices = append(service.AfterServices, daemonsetTemp.ServiceId)
					}
				}
			}
		}

		//image pull secrets
		for _, objRef := range daemonset.Spec.Template.Spec.ImagePullSecrets {
			secretname := objRef.Name
			secret, err := conn.getSecret(ctx, secretname, namespace)
			if err != nil {
				return err
			}
			secretTemp, err := getCpConvertedTemplate(secret, secret.Kind)
			if err != nil {
				utils.Error.Println(err)
				return err
			}
			if !isAlreadyExist(secretTemp.NameSpace, secretTemp.ServiceSubType, secretTemp.Name) {
				daemonsetTemp.BeforeServices = append(daemonsetTemp.BeforeServices, secretTemp.ServiceId)
				secretTemp.AfterServices = append(secretTemp.AfterServices, daemonsetTemp.ServiceId)
				serviceTemplates = append(serviceTemplates, secretTemp)
			} else {
				isSameDaemonSet := false
				for _, serviceId := range secretTemp.AfterServices {
					if *serviceId == *daemonsetTemp.ServiceId {
						isSameDaemonSet = true
						break
					}
				}
				if !isSameDaemonSet {
					daemonsetTemp.BeforeServices = append(daemonsetTemp.BeforeServices, secretTemp.ServiceId)
					secretTemp.AfterServices = append(secretTemp.AfterServices, daemonsetTemp.ServiceId)
				}
			}
		}

		//kubernetes service depecndency findings
		var kubeSvcList *v2.ServiceList
		for key, value := range daemonset.Spec.Template.Labels {
			kubeSvcList, err = conn.getKubernetesServices(ctx, key, value, namespace)
			if err != nil {
				return err
			}
		}

		//container dependency finding
		for _, container := range daemonset.Spec.Template.Spec.Containers {

			//resolving dependencies of kubernetes service
			if kubeSvcList != nil && len(kubeSvcList.Items) > 0 {
				for _, kubeSvc := range kubeSvcList.Items {
					if isPortMatched(&kubeSvc, &container) {
						k8serviceTemp, err := getCpConvertedTemplate(kubeSvc, kubeSvc.Kind)
						if err != nil {
							utils.Error.Println(err)
							return err
						}
						if !isAlreadyExist(k8serviceTemp.NameSpace, k8serviceTemp.ServiceSubType, k8serviceTemp.Name) {
							k8serviceTemp.BeforeServices = append(k8serviceTemp.BeforeServices, daemonsetTemp.ServiceId)
							daemonsetTemp.AfterServices = append(daemonsetTemp.AfterServices, k8serviceTemp.ServiceId)
							serviceTemplates = append(serviceTemplates, k8serviceTemp)
						} else {
							isSameDaemonSet := false
							for _, serviceId := range k8serviceTemp.BeforeServices {
								if *serviceId == *daemonsetTemp.ServiceId {
									isSameDaemonSet = true
									break
								}
							}
							if !isSameDaemonSet {
								k8serviceTemp.BeforeServices = append(k8serviceTemp.BeforeServices, daemonsetTemp.ServiceId)
								daemonsetTemp.AfterServices = append(daemonsetTemp.AfterServices, k8serviceTemp.ServiceId)
							}
						}
					}
				}
			}

			//discovering secret and config maps in deployment containers
			for _, env := range container.Env {
				if env.ValueFrom != nil && env.ValueFrom.SecretKeyRef != nil {
					secretname := env.ValueFrom.SecretKeyRef.Name
					secret, err := conn.getSecret(ctx, secretname, namespace)
					if err != nil {
						return err
					}
					secretTemp, err := getCpConvertedTemplate(secret, secret.Kind)
					if err != nil {
						utils.Error.Println(err)
						return err
					}
					if !isAlreadyExist(secretTemp.NameSpace, secretTemp.ServiceSubType, secretTemp.Name) {
						daemonsetTemp.BeforeServices = append(daemonsetTemp.BeforeServices, secretTemp.ServiceId)
						secretTemp.AfterServices = append(secretTemp.AfterServices, daemonsetTemp.ServiceId)
						serviceTemplates = append(serviceTemplates, secretTemp)
					} else {
						isSameDaemonSet := false
						for _, serviceId := range secretTemp.AfterServices {
							if *serviceId == *daemonsetTemp.ServiceId {
								isSameDaemonSet = true
								break
							}
						}
						if !isSameDaemonSet {
							daemonsetTemp.BeforeServices = append(daemonsetTemp.BeforeServices, secretTemp.ServiceId)
							secretTemp.AfterServices = append(secretTemp.AfterServices, daemonsetTemp.ServiceId)
						}
					}
				} else if env.ValueFrom != nil && env.ValueFrom.ConfigMapKeyRef != nil {
					configmapname := env.ValueFrom.ConfigMapKeyRef.Name
					configmap, err := conn.getConfigMap(ctx, configmapname, namespace)
					if err != nil {
						return err
					}
					configmapTemp, err := getCpConvertedTemplate(configmap, configmap.Kind)
					if err != nil {
						utils.Error.Println(err)
						return err
					}
					if !isAlreadyExist(configmapTemp.NameSpace, configmapTemp.ServiceSubType, configmapTemp.Name) {
						daemonsetTemp.BeforeServices = append(daemonsetTemp.BeforeServices, configmapTemp.ServiceId)
						configmapTemp.AfterServices = append(configmapTemp.AfterServices, daemonsetTemp.ServiceId)
						serviceTemplates = append(serviceTemplates, configmapTemp)
					} else {
						isSameDaemonSet := false
						for _, serviceId := range configmapTemp.AfterServices {
							if *serviceId == *daemonsetTemp.ServiceId {
								isSameDaemonSet = true
								break
							}
						}
						if !isSameDaemonSet {
							daemonsetTemp.BeforeServices = append(daemonsetTemp.BeforeServices, configmapTemp.ServiceId)
							configmapTemp.AfterServices = append(configmapTemp.AfterServices, daemonsetTemp.ServiceId)
						}
					}
				}
			}
		}

		//volume dependency finding
		for _, vol := range daemonset.Spec.Template.Spec.Volumes {
			if vol.Secret != nil {
				secretname := vol.Secret.SecretName
				secret, err := conn.getSecret(ctx, secretname, namespace)
				if err != nil {
					return err
				}
				secretTemp, err := getCpConvertedTemplate(secret, secret.Kind)
				if err != nil {
					utils.Error.Println(err)
					return err
				}
				if !isAlreadyExist(secretTemp.NameSpace, secretTemp.ServiceSubType, secretTemp.Name) {
					daemonsetTemp.BeforeServices = append(daemonsetTemp.BeforeServices, secretTemp.ServiceId)
					secretTemp.AfterServices = append(secretTemp.AfterServices, daemonsetTemp.ServiceId)
					serviceTemplates = append(serviceTemplates, secretTemp)
				} else {
					isSameDaemonSet := false
					for _, serviceId := range secretTemp.AfterServices {
						if *serviceId == *daemonsetTemp.ServiceId {
							isSameDaemonSet = true
							break
						}
					}
					if !isSameDaemonSet {
						daemonsetTemp.BeforeServices = append(daemonsetTemp.BeforeServices, secretTemp.ServiceId)
						secretTemp.AfterServices = append(secretTemp.AfterServices, daemonsetTemp.ServiceId)
					}
				}
			} else if vol.ConfigMap != nil {
				configmapname := vol.ConfigMap.Name
				configmap, err := conn.getConfigMap(ctx, configmapname, namespace)
				if err != nil {
					return err
				}
				configmapTemp, err := getCpConvertedTemplate(configmap, configmap.Kind)
				if err != nil {
					utils.Error.Println(err)
					return err
				}
				if !isAlreadyExist(configmapTemp.NameSpace, configmapTemp.ServiceSubType, configmapTemp.Name) {
					daemonsetTemp.BeforeServices = append(daemonsetTemp.BeforeServices, configmapTemp.ServiceId)
					configmapTemp.AfterServices = append(configmapTemp.AfterServices, daemonsetTemp.ServiceId)
					serviceTemplates = append(serviceTemplates, configmapTemp)
				} else {
					isSameDaemonSet := false
					for _, serviceId := range configmapTemp.AfterServices {
						if *serviceId == *daemonsetTemp.ServiceId {
							isSameDaemonSet = true
							break
						}
					}
					if !isSameDaemonSet {
						daemonsetTemp.BeforeServices = append(daemonsetTemp.BeforeServices, configmapTemp.ServiceId)
						configmapTemp.AfterServices = append(configmapTemp.AfterServices, daemonsetTemp.ServiceId)
					}
				}
			} else if vol.PersistentVolumeClaim != nil {
				pvcname := vol.PersistentVolumeClaim.ClaimName
				pvc, err := conn.getPvc(ctx, pvcname, namespace)
				if err != nil {
					return err
				}
				pvcTemp, err := getCpConvertedTemplate(pvc, pvc.Kind)
				if err != nil {
					utils.Error.Println(err)
					return err
				}
				if !isAlreadyExist(pvcTemp.NameSpace, pvcTemp.ServiceSubType, pvcTemp.Name) {
					daemonsetTemp.BeforeServices = append(daemonsetTemp.BeforeServices, pvcTemp.ServiceId)
					pvcTemp.AfterServices = append(pvcTemp.AfterServices, daemonsetTemp.ServiceId)
					serviceTemplates = append(serviceTemplates, pvcTemp)
				} else {
					for _, service := range serviceTemplates {
						if *service.NameSpace == *pvcTemp.NameSpace && *service.ServiceSubType == *pvcTemp.ServiceSubType && *service.Name == *pvcTemp.Name {
							daemonsetTemp.BeforeServices = append(daemonsetTemp.BeforeServices, service.ServiceId)
							service.AfterServices = append(service.AfterServices, daemonsetTemp.ServiceId)
						}
					}
				}
				if pvc.Spec.StorageClassName != nil {
					storageClassName := *pvc.Spec.StorageClassName
					storageClass, err := conn.getStorageClass(ctx, storageClassName, namespace)
					if err != nil {
						utils.Error.Println(err)
						return err
					}
					storageClassTemp, err := getCpConvertedTemplate(storageClass, storageClass.Kind)
					if err != nil {
						utils.Error.Println(err)
						return err
					}
					if !isAlreadyExist(storageClassTemp.NameSpace, storageClassTemp.ServiceSubType, storageClassTemp.Name) {
						pvcTemp.BeforeServices = append(pvcTemp.BeforeServices, storageClassTemp.ServiceId)
						storageClassTemp.AfterServices = append(storageClassTemp.AfterServices, pvcTemp.ServiceId)
						serviceTemplates = append(serviceTemplates, storageClassTemp)
					} else {
						isPVCexist := false
						for _, serviceId := range storageClassTemp.AfterServices {
							if *serviceId == *pvcTemp.ServiceId {
								isPVCexist = true
								break
							}
						}
						if !isPVCexist {
							pvcTemp.BeforeServices = append(pvcTemp.BeforeServices, storageClassTemp.ServiceId)
							storageClassTemp.AfterServices = append(storageClassTemp.AfterServices, pvcTemp.ServiceId)
						}
					}
				}
			}

			if vol.Projected != nil {
				for _, source := range vol.Projected.Sources {
					if source.ConfigMap != nil {
						configmapname := vol.ConfigMap.Name
						configmap, err := conn.getConfigMap(ctx, configmapname, namespace)
						if err != nil {
							return err
						}
						configmapTemp, err := getCpConvertedTemplate(configmap, configmap.Kind)
						if err != nil {
							utils.Error.Println(err)
							return err
						}
						if !isAlreadyExist(configmapTemp.NameSpace, configmapTemp.ServiceSubType, configmapTemp.Name) {
							daemonsetTemp.BeforeServices = append(daemonsetTemp.BeforeServices, configmapTemp.ServiceId)
							configmapTemp.AfterServices = append(configmapTemp.AfterServices, daemonsetTemp.ServiceId)
							serviceTemplates = append(serviceTemplates, configmapTemp)
						} else {
							isSameDaemonSet := false
							for _, serviceId := range configmapTemp.AfterServices {
								if *serviceId == *daemonsetTemp.ServiceId {
									isSameDaemonSet = true
									break
								}
							}
							if !isSameDaemonSet {
								daemonsetTemp.BeforeServices = append(daemonsetTemp.BeforeServices, configmapTemp.ServiceId)
								configmapTemp.AfterServices = append(configmapTemp.AfterServices, daemonsetTemp.ServiceId)
							}
						}
					} else if source.Secret != nil {
						secretname := vol.Secret.SecretName
						secret, err := conn.getSecret(ctx, secretname, namespace)
						if err != nil {
							return err
						}
						secretTemp, err := getCpConvertedTemplate(secret, secret.Kind)
						if err != nil {
							utils.Error.Println(err)
							return err
						}
						if !isAlreadyExist(secretTemp.NameSpace, secretTemp.ServiceSubType, secretTemp.Name) {
							daemonsetTemp.BeforeServices = append(daemonsetTemp.BeforeServices, secretTemp.ServiceId)
							secretTemp.AfterServices = append(secretTemp.AfterServices, daemonsetTemp.ServiceId)
							serviceTemplates = append(serviceTemplates, secretTemp)
						} else {
							isSameDaemonSet := false
							for _, serviceId := range secretTemp.AfterServices {
								if *serviceId == *daemonsetTemp.ServiceId {
									isSameDaemonSet = true
									break
								}
							}
							if !isSameDaemonSet {
								daemonsetTemp.BeforeServices = append(daemonsetTemp.BeforeServices, secretTemp.ServiceId)
								secretTemp.AfterServices = append(secretTemp.AfterServices, daemonsetTemp.ServiceId)
							}
						}
					}
				}
			}
		}

		serviceTemplates = append(serviceTemplates, daemonsetTemp)
	}

	return nil
}

func (conn *GrpcConn) statefulsetsK8sToCp(ctx context.Context, statefulsets []v1.StatefulSet) error {

	for _, statefulset := range statefulsets {
		utils.Info.Printf("Resolving dependencies of statefulset :%v within namespace : %v", statefulset.Name, statefulset.Namespace)
		stsTemp, err := getCpConvertedTemplate(statefulset, statefulset.Kind)
		if err != nil {
			utils.Error.Println(err)
			return err
		}
		namespace := statefulset.Namespace
		if statefulset.Spec.Template.Spec.ServiceAccountName != "" {
			svcname := statefulset.Spec.Template.Spec.ServiceAccountName
			svcaccount, err := conn.getSvcAccount(ctx, svcname, namespace)
			if err != nil {
				return err
			}
			svcAccTemp, err := getCpConvertedTemplate(svcaccount, svcaccount.Kind)
			if err != nil {
				utils.Error.Println(err)
				return err
			}
			if !isAlreadyExist(svcAccTemp.NameSpace, svcAccTemp.ServiceSubType, svcAccTemp.Name) {
				rbacDependencies, err := conn.getK8sRbacResources(ctx, namespace, svcaccount, svcAccTemp)
				if err != nil {
					utils.Error.Println(err)
					return err
				}

				var strpointer = new(string)
				*strpointer = "serviceaccount"
				for _, rbacTemp := range rbacDependencies {
					if *rbacTemp.ServiceSubType == *strpointer {
						stsTemp.BeforeServices = append(stsTemp.BeforeServices, rbacTemp.ServiceId)
						rbacTemp.AfterServices = append(rbacTemp.AfterServices, stsTemp.ServiceId)
					}
					serviceTemplates = append(serviceTemplates, rbacTemp)
				}
			} else {
				for _, service := range serviceTemplates {
					if *service.NameSpace == *svcAccTemp.NameSpace && *service.ServiceSubType == *svcAccTemp.ServiceSubType && *service.Name == *svcAccTemp.Name {
						stsTemp.BeforeServices = append(stsTemp.BeforeServices, service.ServiceId)
						service.AfterServices = append(service.AfterServices, stsTemp.ServiceId)
					}
				}
			}
		}

		//image pull secrets
		for _, objRef := range statefulset.Spec.Template.Spec.ImagePullSecrets {
			secretname := objRef.Name
			secret, err := conn.getSecret(ctx, secretname, namespace)
			if err != nil {
				return err
			}
			secretTemp, err := getCpConvertedTemplate(secret, secret.Kind)
			if err != nil {
				utils.Error.Println(err)
				return err
			}
			if !isAlreadyExist(secretTemp.NameSpace, secretTemp.ServiceSubType, secretTemp.Name) {
				stsTemp.BeforeServices = append(stsTemp.BeforeServices, secretTemp.ServiceId)
				secretTemp.AfterServices = append(secretTemp.AfterServices, stsTemp.ServiceId)
				serviceTemplates = append(serviceTemplates, secretTemp)
			} else {
				isSameStatefulSet := false
				for _, serviceId := range secretTemp.AfterServices {
					if *serviceId == *stsTemp.ServiceId {
						isSameStatefulSet = true
						break
					}
				}
				if !isSameStatefulSet {
					stsTemp.BeforeServices = append(stsTemp.BeforeServices, secretTemp.ServiceId)
					secretTemp.AfterServices = append(secretTemp.AfterServices, stsTemp.ServiceId)
				}
			}
		}

		hpaList, err := conn.getAllHpas(ctx, namespace)
		if err != nil {
			return err
		}

		if len(hpaList.Items) > 0 {
			for _, hpa := range hpaList.Items {
				if hpa.Spec.ScaleTargetRef.APIVersion == statefulset.APIVersion && hpa.Spec.ScaleTargetRef.Kind == statefulset.Kind && hpa.Spec.ScaleTargetRef.Name == statefulset.Name {
					hpaTemplate, err := getCpConvertedTemplate(hpa, hpa.Kind)
					if err != nil {
						utils.Error.Println(err)
						return err
					}
					hpaTemplate.BeforeServices = append(hpaTemplate.BeforeServices, stsTemp.ServiceId)
					stsTemp.AfterServices = append(stsTemp.AfterServices, hpaTemplate.ServiceId)
					serviceTemplates = append(serviceTemplates, hpaTemplate)
				}
			}
		}

		//kubernetes service depecndency findings
		var kubeSvcList *v2.ServiceList
		for key, value := range statefulset.Spec.Template.Labels {
			kubeSvcList, err = conn.getKubernetesServices(ctx, key, value, namespace)
			if err != nil {
				return err
			}
		}

		//container dependency finding
		for _, container := range statefulset.Spec.Template.Spec.Containers {

			//resolving dependencies of kubernetes service
			if kubeSvcList != nil && len(kubeSvcList.Items) > 0 {
				for _, kubeSvc := range kubeSvcList.Items {
					if isPortMatched(&kubeSvc, &container) {
						k8serviceTemp, err := getCpConvertedTemplate(kubeSvc, kubeSvc.Kind)
						if err != nil {
							utils.Error.Println(err)
							return err
						}
						if !isAlreadyExist(k8serviceTemp.NameSpace, k8serviceTemp.ServiceSubType, k8serviceTemp.Name) {
							k8serviceTemp.BeforeServices = append(k8serviceTemp.BeforeServices, stsTemp.ServiceId)
							stsTemp.AfterServices = append(stsTemp.AfterServices, k8serviceTemp.ServiceId)
							serviceTemplates = append(serviceTemplates, k8serviceTemp)
						} else {
							isSameStatefulSet := false
							for _, serviceId := range k8serviceTemp.BeforeServices {
								if *serviceId == *stsTemp.ServiceId {
									isSameStatefulSet = true
									break
								}
							}
							if !isSameStatefulSet {
								k8serviceTemp.BeforeServices = append(k8serviceTemp.BeforeServices, stsTemp.ServiceId)
								stsTemp.AfterServices = append(stsTemp.AfterServices, k8serviceTemp.ServiceId)
							}
						}
					}
				}
			}

			//discovering secret and config maps in deployment containers
			for _, env := range container.Env {
				if env.ValueFrom != nil && env.ValueFrom.SecretKeyRef != nil {
					secretname := env.ValueFrom.SecretKeyRef.Name
					secret, err := conn.getSecret(ctx, secretname, namespace)
					if err != nil {
						return err
					}
					secretTemp, err := getCpConvertedTemplate(secret, secret.Kind)
					if err != nil {
						utils.Error.Println(err)
						return err
					}
					if !isAlreadyExist(secretTemp.NameSpace, secretTemp.ServiceSubType, secretTemp.Name) {
						stsTemp.BeforeServices = append(stsTemp.BeforeServices, secretTemp.ServiceId)
						secretTemp.AfterServices = append(secretTemp.AfterServices, stsTemp.ServiceId)
						serviceTemplates = append(serviceTemplates, secretTemp)
					} else {
						isSameStatefulSet := false
						for _, serviceId := range secretTemp.AfterServices {
							if *serviceId == *stsTemp.ServiceId {
								isSameStatefulSet = true
								break
							}
						}
						if !isSameStatefulSet {
							stsTemp.BeforeServices = append(stsTemp.BeforeServices, secretTemp.ServiceId)
							secretTemp.AfterServices = append(secretTemp.AfterServices, stsTemp.ServiceId)
						}
					}
				} else if env.ValueFrom != nil && env.ValueFrom.ConfigMapKeyRef != nil {
					configmapname := env.ValueFrom.ConfigMapKeyRef.Name
					configmap, err := conn.getConfigMap(ctx, configmapname, namespace)
					if err != nil {
						return err
					}
					configmapTemp, err := getCpConvertedTemplate(configmap, configmap.Kind)
					if err != nil {
						utils.Error.Println(err)
						return err
					}
					if !isAlreadyExist(configmapTemp.NameSpace, configmapTemp.ServiceSubType, configmapTemp.Name) {
						stsTemp.BeforeServices = append(stsTemp.BeforeServices, configmapTemp.ServiceId)
						configmapTemp.AfterServices = append(configmapTemp.AfterServices, stsTemp.ServiceId)
						serviceTemplates = append(serviceTemplates, configmapTemp)
					} else {
						isSameStatefulSet := false
						for _, serviceId := range configmapTemp.AfterServices {
							if *serviceId == *stsTemp.ServiceId {
								isSameStatefulSet = true
								break
							}
						}
						if !isSameStatefulSet {
							stsTemp.BeforeServices = append(stsTemp.BeforeServices, configmapTemp.ServiceId)
							configmapTemp.AfterServices = append(configmapTemp.AfterServices, stsTemp.ServiceId)
						}
					}
				}
			}
		}

		//volume dependency finding
		for _, vol := range statefulset.Spec.Template.Spec.Volumes {
			if vol.Secret != nil {
				secretname := vol.Secret.SecretName
				secret, err := conn.getSecret(ctx, secretname, namespace)
				if err != nil {
					return err
				}
				secretTemp, err := getCpConvertedTemplate(secret, secret.Kind)
				if err != nil {
					utils.Error.Println(err)
					return err
				}
				if !isAlreadyExist(secretTemp.NameSpace, secretTemp.ServiceSubType, secretTemp.Name) {
					stsTemp.BeforeServices = append(stsTemp.BeforeServices, secretTemp.ServiceId)
					secretTemp.AfterServices = append(secretTemp.AfterServices, stsTemp.ServiceId)
					serviceTemplates = append(serviceTemplates, secretTemp)
				} else {
					isSameStatefulSet := false
					for _, serviceId := range secretTemp.AfterServices {
						if *serviceId == *stsTemp.ServiceId {
							isSameStatefulSet = true
							break
						}
					}
					if !isSameStatefulSet {
						stsTemp.BeforeServices = append(stsTemp.BeforeServices, secretTemp.ServiceId)
						secretTemp.AfterServices = append(secretTemp.AfterServices, stsTemp.ServiceId)
					}
				}
			} else if vol.ConfigMap != nil {
				configmapname := vol.ConfigMap.Name
				configmap, err := conn.getConfigMap(ctx, configmapname, namespace)
				if err != nil {
					return err
				}
				configmapTemp, err := getCpConvertedTemplate(configmap, configmap.Kind)
				if err != nil {
					utils.Error.Println(err)
					return err
				}
				if !isAlreadyExist(configmapTemp.NameSpace, configmapTemp.ServiceSubType, configmapTemp.Name) {
					stsTemp.BeforeServices = append(stsTemp.BeforeServices, configmapTemp.ServiceId)
					configmapTemp.AfterServices = append(configmapTemp.AfterServices, stsTemp.ServiceId)
					serviceTemplates = append(serviceTemplates, configmapTemp)
				} else {
					isSameStatefulSet := false
					for _, serviceId := range configmapTemp.AfterServices {
						if *serviceId == *stsTemp.ServiceId {
							isSameStatefulSet = true
							break
						}
					}
					if !isSameStatefulSet {
						stsTemp.BeforeServices = append(stsTemp.BeforeServices, configmapTemp.ServiceId)
						configmapTemp.AfterServices = append(configmapTemp.AfterServices, stsTemp.ServiceId)
					}
				}
			} else if vol.PersistentVolumeClaim != nil {
				pvcname := vol.PersistentVolumeClaim.ClaimName
				pvc, err := conn.getPvc(ctx, pvcname, namespace)
				if err != nil {
					return err
				}
				pvcTemp, err := getCpConvertedTemplate(pvc, pvc.Kind)
				if err != nil {
					utils.Error.Println(err)
					return err
				}
				if !isAlreadyExist(pvcTemp.NameSpace, pvcTemp.ServiceSubType, pvcTemp.Name) {
					stsTemp.BeforeServices = append(stsTemp.BeforeServices, pvcTemp.ServiceId)
					pvcTemp.AfterServices = append(pvcTemp.AfterServices, stsTemp.ServiceId)
					serviceTemplates = append(serviceTemplates, pvcTemp)
				} else {
					for _, service := range serviceTemplates {
						if *service.NameSpace == *pvcTemp.NameSpace && *service.ServiceSubType == *pvcTemp.ServiceSubType && *service.Name == *pvcTemp.Name {
							stsTemp.BeforeServices = append(stsTemp.BeforeServices, service.ServiceId)
							service.AfterServices = append(service.AfterServices, stsTemp.ServiceId)
						}
					}
				}
				if pvc.Spec.StorageClassName != nil {
					storageClassName := *pvc.Spec.StorageClassName
					storageClass, err := conn.getStorageClass(ctx, storageClassName, namespace)
					if err != nil {
						utils.Error.Println(err)
						return err
					}
					storageClassTemp, err := getCpConvertedTemplate(storageClass, storageClass.Kind)
					if err != nil {
						utils.Error.Println(err)
						return err
					}
					if !isAlreadyExist(storageClassTemp.NameSpace, storageClassTemp.ServiceSubType, storageClassTemp.Name) {
						pvcTemp.BeforeServices = append(pvcTemp.BeforeServices, storageClassTemp.ServiceId)
						storageClassTemp.AfterServices = append(storageClassTemp.AfterServices, pvcTemp.ServiceId)
						serviceTemplates = append(serviceTemplates, storageClassTemp)
					} else {
						isPVCexist := false
						for _, serviceId := range storageClassTemp.AfterServices {
							if *serviceId == *pvcTemp.ServiceId {
								isPVCexist = true
								break
							}
						}
						if !isPVCexist {
							pvcTemp.BeforeServices = append(pvcTemp.BeforeServices, storageClassTemp.ServiceId)
							storageClassTemp.AfterServices = append(storageClassTemp.AfterServices, pvcTemp.ServiceId)
						}
					}
				}
			}

			if vol.Projected != nil {
				for _, source := range vol.Projected.Sources {
					if source.ConfigMap != nil {
						configmapname := vol.ConfigMap.Name
						configmap, err := conn.getConfigMap(ctx, configmapname, namespace)
						if err != nil {
							return err
						}
						configmapTemp, err := getCpConvertedTemplate(configmap, configmap.Kind)
						if err != nil {
							utils.Error.Println(err)
							return err
						}
						if !isAlreadyExist(configmapTemp.NameSpace, configmapTemp.ServiceSubType, configmapTemp.Name) {
							stsTemp.BeforeServices = append(stsTemp.BeforeServices, configmapTemp.ServiceId)
							configmapTemp.AfterServices = append(configmapTemp.AfterServices, stsTemp.ServiceId)
							serviceTemplates = append(serviceTemplates, configmapTemp)
						} else {
							isSameStatefulSet := false
							for _, serviceId := range configmapTemp.AfterServices {
								if *serviceId == *stsTemp.ServiceId {
									isSameStatefulSet = true
									break
								}
							}
							if !isSameStatefulSet {
								stsTemp.BeforeServices = append(stsTemp.BeforeServices, configmapTemp.ServiceId)
								configmapTemp.AfterServices = append(configmapTemp.AfterServices, stsTemp.ServiceId)
							}
						}
					} else if source.Secret != nil {
						secretname := vol.Secret.SecretName
						secret, err := conn.getSecret(ctx, secretname, namespace)
						if err != nil {
							return err
						}
						secretTemp, err := getCpConvertedTemplate(secret, secret.Kind)
						if err != nil {
							utils.Error.Println(err)
							return err
						}
						if !isAlreadyExist(secretTemp.NameSpace, secretTemp.ServiceSubType, secretTemp.Name) {
							stsTemp.BeforeServices = append(stsTemp.BeforeServices, secretTemp.ServiceId)
							secretTemp.AfterServices = append(secretTemp.AfterServices, stsTemp.ServiceId)
							serviceTemplates = append(serviceTemplates, secretTemp)
						} else {
							isSameStatefulSet := false
							for _, serviceId := range secretTemp.AfterServices {
								if *serviceId == *stsTemp.ServiceId {
									isSameStatefulSet = true
									break
								}
							}
							if !isSameStatefulSet {
								stsTemp.BeforeServices = append(stsTemp.BeforeServices, secretTemp.ServiceId)
								secretTemp.AfterServices = append(secretTemp.AfterServices, stsTemp.ServiceId)
							}
						}
					}
				}
			}
		}

		serviceTemplates = append(serviceTemplates, stsTemp)
	}
	return nil

}

func (conn *GrpcConn) deploymentk8sToCp(ctx context.Context, deployments []v1.Deployment, wg *sync.WaitGroup) error {

	for _, dep := range deployments {
		wg.Add(1)
		go conn.ResolveDeploymentDependencies(dep, wg, ctx)
	}
	return nil
}

func (conn *GrpcConn) ResolveDeploymentDependencies(dep v1.Deployment, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	if dep.Name == "antelope" {
		fmt.Println("hello from antelope")
	}
	utils.Info.Printf("Resolving dependencies of deployment :%v within namespace : %v", dep.Name, dep.Namespace)
	depTemp, err := getCpConvertedTemplate(dep, dep.Kind)
	if err != nil {
		utils.Error.Println(err)
		return
	}

	namespace := dep.Namespace
	//checking for the service account if name not empty then getting cluster role and cluster role  binding against that service account
	if dep.Spec.Template.Spec.ServiceAccountName != "" {
		svcname := dep.Spec.Template.Spec.ServiceAccountName
		svcaccount, err := conn.getSvcAccount(ctx, svcname, namespace)
		if err != nil {
			return
		}
		svcAccTemp, err := getCpConvertedTemplate(svcaccount, svcaccount.Kind)
		if err != nil {
			utils.Error.Println(err)
			return
		}
		if !isAlreadyExist(svcAccTemp.NameSpace, svcAccTemp.ServiceSubType, svcAccTemp.Name) {
			rbacDependencies, err := conn.getK8sRbacResources(ctx, namespace, svcaccount, svcAccTemp)
			if err != nil {
				utils.Error.Println(err)
				return
			}

			var strpointer = new(string)
			*strpointer = "serviceaccount"
			for _, rbacTemp := range rbacDependencies {
				if *rbacTemp.ServiceSubType == *strpointer {
					depTemp.BeforeServices = append(depTemp.BeforeServices, rbacTemp.ServiceId)
					rbacTemp.AfterServices = append(rbacTemp.AfterServices, depTemp.ServiceId)
				}
				serviceTemplates = append(serviceTemplates, rbacTemp)
			}
		} else {
			for _, service := range serviceTemplates {
				if *service.NameSpace == *svcAccTemp.NameSpace && *service.ServiceSubType == *svcAccTemp.ServiceSubType && *service.Name == *svcAccTemp.Name {
					depTemp.BeforeServices = append(depTemp.BeforeServices, service.ServiceId)
					service.AfterServices = append(service.AfterServices, depTemp.ServiceId)
				}
			}
		}

	}

	//image pull secrets
	for _, objRef := range dep.Spec.Template.Spec.ImagePullSecrets {
		secretname := objRef.Name
		secret, err := conn.getSecret(ctx, secretname, namespace)
		if err != nil {
			utils.Error.Println(err)
			return
		}
		secretTemp, err := getCpConvertedTemplate(secret, secret.Kind)
		if err != nil {
			utils.Error.Println(err)
			return
		}
		if !isAlreadyExist(secretTemp.NameSpace, secretTemp.ServiceSubType, secretTemp.Name) {
			depTemp.BeforeServices = append(depTemp.BeforeServices, secretTemp.ServiceId)
			secretTemp.AfterServices = append(secretTemp.AfterServices, depTemp.ServiceId)
			serviceTemplates = append(serviceTemplates, secretTemp)
		} else {
			isSameDeployment := false
			for _, serviceId := range secretTemp.AfterServices {
				if *serviceId == *depTemp.ServiceId {
					isSameDeployment = true
					break
				}
			}
			if !isSameDeployment {
				depTemp.BeforeServices = append(depTemp.BeforeServices, secretTemp.ServiceId)
				secretTemp.AfterServices = append(secretTemp.AfterServices, depTemp.ServiceId)
			}
		}
	}

	hpaList, err := conn.getAllHpas(ctx, namespace)
	if err != nil {
		return
	}
	if len(hpaList.Items) > 0 {
		for _, hpa := range hpaList.Items {
			if hpa.Spec.ScaleTargetRef.APIVersion == dep.APIVersion && hpa.Spec.ScaleTargetRef.Kind == dep.Kind && hpa.Spec.ScaleTargetRef.Name == dep.Name {
				hpaTemplate, err := getCpConvertedTemplate(hpa, hpa.Kind)
				if err != nil {
					utils.Error.Println(err)
					return
				}
				hpaTemplate.BeforeServices = append(hpaTemplate.BeforeServices, depTemp.ServiceId)
				depTemp.AfterServices = append(depTemp.AfterServices, hpaTemplate.ServiceId)
				serviceTemplates = append(serviceTemplates, hpaTemplate)
			}
		}
	}

	//finding kubernetes service
	var kubeSvcList *v2.ServiceList
	for key, value := range dep.Spec.Template.Labels {
		kubeSvcList, err = conn.getKubernetesServices(ctx, key, value, namespace)
		if err != nil {
			utils.Error.Println(err)
			return
		}
	}

	//container dependency finding
	for _, container := range dep.Spec.Template.Spec.Containers {

		//resolving dependencies of kubernetes service
		if kubeSvcList != nil && len(kubeSvcList.Items) > 0 {
			for _, kubeSvc := range kubeSvcList.Items {
				if isPortMatched(&kubeSvc, &container) {
					k8serviceTemp, err := getCpConvertedTemplate(kubeSvc, kubeSvc.Kind)
					if err != nil {
						utils.Error.Println(err)
						return
					}
					if !isAlreadyExist(k8serviceTemp.NameSpace, k8serviceTemp.ServiceSubType, k8serviceTemp.Name) {
						k8serviceTemp.BeforeServices = append(k8serviceTemp.BeforeServices, depTemp.ServiceId)
						depTemp.AfterServices = append(depTemp.AfterServices, k8serviceTemp.ServiceId)
						serviceTemplates = append(serviceTemplates, k8serviceTemp)
					} else {
						isSameDeployment := false
						for _, serviceId := range k8serviceTemp.BeforeServices {
							if *serviceId == *depTemp.ServiceId {
								isSameDeployment = true
								break
							}
						}
						if !isSameDeployment {
							k8serviceTemp.BeforeServices = append(k8serviceTemp.BeforeServices, depTemp.ServiceId)
							depTemp.AfterServices = append(depTemp.AfterServices, k8serviceTemp.ServiceId)
						}
					}
				}
			}
		}

		//discovering secret and config maps in deployment containers
		for _, env := range container.Env {
			if env.ValueFrom != nil && env.ValueFrom.SecretKeyRef != nil {
				secretname := env.ValueFrom.SecretKeyRef.Name
				secret, err := conn.getSecret(ctx, secretname, namespace)
				if err != nil {
					utils.Error.Println(err)
					return
				}
				secretTemp, err := getCpConvertedTemplate(secret, secret.Kind)
				if err != nil {
					utils.Error.Println(err)
					return
				}
				if !isAlreadyExist(secretTemp.NameSpace, secretTemp.ServiceSubType, secretTemp.Name) {
					depTemp.BeforeServices = append(depTemp.BeforeServices, secretTemp.ServiceId)
					secretTemp.AfterServices = append(secretTemp.AfterServices, depTemp.ServiceId)
					serviceTemplates = append(serviceTemplates, secretTemp)
				} else {
					isSameDeployment := false
					for _, serviceId := range secretTemp.AfterServices {
						if *serviceId == *depTemp.ServiceId {
							isSameDeployment = true
							break
						}
					}
					if !isSameDeployment {
						depTemp.BeforeServices = append(depTemp.BeforeServices, secretTemp.ServiceId)
						secretTemp.AfterServices = append(secretTemp.AfterServices, depTemp.ServiceId)
					}
				}
			} else if env.ValueFrom != nil && env.ValueFrom.ConfigMapKeyRef != nil {
				configmapname := env.ValueFrom.ConfigMapKeyRef.Name
				configmap, err := conn.getConfigMap(ctx, configmapname, namespace)
				if err != nil {
					utils.Error.Println(err)
					return
				}
				configmapTemp, err := getCpConvertedTemplate(configmap, configmap.Kind)
				if err != nil {
					utils.Error.Println(err)
					return
				}
				if !isAlreadyExist(configmapTemp.NameSpace, configmapTemp.ServiceSubType, configmapTemp.Name) {
					depTemp.BeforeServices = append(depTemp.BeforeServices, configmapTemp.ServiceId)
					configmapTemp.AfterServices = append(configmapTemp.AfterServices, depTemp.ServiceId)
					serviceTemplates = append(serviceTemplates, configmapTemp)
				} else {
					isSameDeployment := false
					for _, serviceId := range configmapTemp.AfterServices {
						if *serviceId == *depTemp.ServiceId {
							isSameDeployment = true
							break
						}
					}
					if !isSameDeployment {
						depTemp.BeforeServices = append(depTemp.BeforeServices, configmapTemp.ServiceId)
						configmapTemp.AfterServices = append(configmapTemp.AfterServices, depTemp.ServiceId)
					}
				}
			}
		}
	}

	//volume dependency finding
	for _, vol := range dep.Spec.Template.Spec.Volumes {
		if vol.Secret != nil {
			secretname := vol.Secret.SecretName
			secret, err := conn.getSecret(ctx, secretname, namespace)
			if err != nil {
				utils.Error.Println(err)
				return
			}
			secretTemp, err := getCpConvertedTemplate(secret, secret.Kind)
			if err != nil {
				utils.Error.Println(err)
				return
			}
			if !isAlreadyExist(secretTemp.NameSpace, secretTemp.ServiceSubType, secretTemp.Name) {
				depTemp.BeforeServices = append(depTemp.BeforeServices, secretTemp.ServiceId)
				secretTemp.AfterServices = append(secretTemp.AfterServices, depTemp.ServiceId)
				serviceTemplates = append(serviceTemplates, secretTemp)
			} else {
				isSameDeployment := false
				for _, serviceId := range secretTemp.AfterServices {
					if *serviceId == *depTemp.ServiceId {
						isSameDeployment = true
						break
					}
				}
				if !isSameDeployment {
					depTemp.BeforeServices = append(depTemp.BeforeServices, secretTemp.ServiceId)
					secretTemp.AfterServices = append(secretTemp.AfterServices, depTemp.ServiceId)
				}
			}
		} else if vol.ConfigMap != nil {
			configmapname := vol.ConfigMap.Name
			configmap, err := conn.getConfigMap(ctx, configmapname, namespace)
			if err != nil {
				utils.Error.Println(err)
				return
			}
			configmapTemp, err := getCpConvertedTemplate(configmap, configmap.Kind)
			if err != nil {
				utils.Error.Println(err)
				return
			}
			if !isAlreadyExist(configmapTemp.NameSpace, configmapTemp.ServiceSubType, configmapTemp.Name) {
				depTemp.BeforeServices = append(depTemp.BeforeServices, configmapTemp.ServiceId)
				configmapTemp.AfterServices = append(configmapTemp.AfterServices, depTemp.ServiceId)
				serviceTemplates = append(serviceTemplates, configmapTemp)
			} else {
				isSameDeployment := false
				for _, serviceId := range configmapTemp.AfterServices {
					if *serviceId == *depTemp.ServiceId {
						isSameDeployment = true
						break
					}
				}
				if !isSameDeployment {
					depTemp.BeforeServices = append(depTemp.BeforeServices, configmapTemp.ServiceId)
					configmapTemp.AfterServices = append(configmapTemp.AfterServices, depTemp.ServiceId)
				}
			}
		} else if vol.PersistentVolumeClaim != nil {
			pvcname := vol.PersistentVolumeClaim.ClaimName
			pvc, err := conn.getPvc(ctx, pvcname, namespace)
			if err != nil {
				utils.Error.Println(err)
				return
			}
			pvcTemp, err := getCpConvertedTemplate(pvc, pvc.Kind)
			if err != nil {
				utils.Error.Println(err)
				return
			}
			if !isAlreadyExist(pvcTemp.NameSpace, pvcTemp.ServiceSubType, pvcTemp.Name) {
				depTemp.BeforeServices = append(depTemp.BeforeServices, pvcTemp.ServiceId)
				pvcTemp.AfterServices = append(pvcTemp.AfterServices, depTemp.ServiceId)
				serviceTemplates = append(serviceTemplates, pvcTemp)
			} else {
				for _, service := range serviceTemplates {
					if *service.NameSpace == *pvcTemp.NameSpace && *service.ServiceSubType == *pvcTemp.ServiceSubType && *service.Name == *pvcTemp.Name {
						depTemp.BeforeServices = append(depTemp.BeforeServices, service.ServiceId)
						service.AfterServices = append(service.AfterServices, depTemp.ServiceId)
					}
				}
			}
			if pvc.Spec.StorageClassName != nil {
				storageClassName := *pvc.Spec.StorageClassName
				storageClass, err := conn.getStorageClass(ctx, storageClassName, namespace)
				if err != nil {
					utils.Error.Println(err)
					return
				}
				storageClassTemp, err := getCpConvertedTemplate(storageClass, storageClass.Kind)
				if err != nil {
					utils.Error.Println(err)
					return
				}
				if !isAlreadyExist(storageClassTemp.NameSpace, storageClassTemp.ServiceSubType, storageClassTemp.Name) {
					pvcTemp.BeforeServices = append(pvcTemp.BeforeServices, storageClassTemp.ServiceId)
					storageClassTemp.AfterServices = append(storageClassTemp.AfterServices, pvcTemp.ServiceId)
					serviceTemplates = append(serviceTemplates, storageClassTemp)
				} else {
					isPVCexist := false
					for _, serviceId := range storageClassTemp.AfterServices {
						if *serviceId == *pvcTemp.ServiceId {
							isPVCexist = true
							break
						}
					}
					if !isPVCexist {
						pvcTemp.BeforeServices = append(pvcTemp.BeforeServices, storageClassTemp.ServiceId)
						storageClassTemp.AfterServices = append(storageClassTemp.AfterServices, pvcTemp.ServiceId)
					}
				}
			}
		}

		if vol.Projected != nil {
			for _, source := range vol.Projected.Sources {
				if source.ConfigMap != nil {
					configmapname := vol.ConfigMap.Name
					configmap, err := conn.getConfigMap(ctx, configmapname, namespace)
					if err != nil {
						utils.Error.Println(err)
						return
					}
					configmapTemp, err := getCpConvertedTemplate(configmap, configmap.Kind)
					if err != nil {
						utils.Error.Println(err)
						return
					}
					if !isAlreadyExist(configmapTemp.NameSpace, configmapTemp.ServiceSubType, configmapTemp.Name) {
						depTemp.BeforeServices = append(depTemp.BeforeServices, configmapTemp.ServiceId)
						configmapTemp.AfterServices = append(configmapTemp.AfterServices, depTemp.ServiceId)
						serviceTemplates = append(serviceTemplates, configmapTemp)
					} else {
						isSameDeployment := false
						for _, serviceId := range configmapTemp.AfterServices {
							if *serviceId == *depTemp.ServiceId {
								isSameDeployment = true
								break
							}
						}
						if !isSameDeployment {
							depTemp.BeforeServices = append(depTemp.BeforeServices, configmapTemp.ServiceId)
							configmapTemp.AfterServices = append(configmapTemp.AfterServices, depTemp.ServiceId)
						}
					}
				} else if source.Secret != nil {
					secretname := vol.Secret.SecretName
					secret, err := conn.getSecret(ctx, secretname, namespace)
					if err != nil {
						utils.Error.Println(err)
						return
					}
					secretTemp, err := getCpConvertedTemplate(secret, secret.Kind)
					if err != nil {
						utils.Error.Println(err)
						return
					}
					if !isAlreadyExist(secretTemp.NameSpace, secretTemp.ServiceSubType, secretTemp.Name) {
						depTemp.BeforeServices = append(depTemp.BeforeServices, secretTemp.ServiceId)
						secretTemp.AfterServices = append(secretTemp.AfterServices, depTemp.ServiceId)
						serviceTemplates = append(serviceTemplates, secretTemp)
					} else {
						isSameDeployment := false
						for _, serviceId := range secretTemp.AfterServices {
							if *serviceId == *depTemp.ServiceId {
								isSameDeployment = true
								break
							}
						}
						if !isSameDeployment {
							depTemp.BeforeServices = append(depTemp.BeforeServices, secretTemp.ServiceId)
							secretTemp.AfterServices = append(secretTemp.AfterServices, depTemp.ServiceId)
						}
					}
				}
			}
		}
	}

	serviceTemplates = append(serviceTemplates, depTemp)
}

func GetExistingService(namespace, svcsubtype, name *string) *types.ServiceTemplate {
	for _, service := range serviceTemplates {
		if *service.NameSpace == *namespace && *service.ServiceSubType == *svcsubtype && *service.Name == *name {
			return service
		}
	}
	return nil
}

func (conn *GrpcConn) getK8sRbacResources(ctx context.Context, namespace string, k8svcAcc *api.ServiceAccount, svcAccTemp *types.ServiceTemplate) ([]*types.ServiceTemplate, error) {

	var rbacServiceTemplates []*types.ServiceTemplate
	//creating secrets for service account
	//for _, secret := range k8svcAcc.Secrets {
	//	if secret.Name != "" {
	//		secretname := secret.Name
	//		if secret.Namespace != "" {
	//			namespace = secret.Namespace
	//		}
	//		scrt, err := conn.getSecret(ctx, secretname, namespace)
	//		if err != nil {
	//			return nil, err
	//		} else {
	//			secretTemp, err := getCpConvertedTemplate(scrt, scrt.Kind)
	//			if err != nil {
	//				utils.Error.Println(err)
	//				return nil, err
	//			}
	//			//this is doubtful
	//			secretTemp.BeforeServices = append(secretTemp.BeforeServices, svcAccTemp.ServiceId)
	//			rbacServiceTemplates = append(rbacServiceTemplates, secretTemp)
	//			svcAccTemp.AfterServices = append(svcAccTemp.AfterServices, secretTemp.ServiceId)
	//			//this is doubtful
	//		}
	//	}
	//}

	clusterrolebindings, err := conn.getAllClusterRoleBindings(ctx)
	if err != nil {
		return nil, err
	}

	for _, clstrrolebind := range clusterrolebindings.Items {
		for _, sub := range clstrrolebind.Subjects {
			if sub.Kind == "ServiceAccount" && sub.Name == k8svcAcc.Name {
				if clstrrolebind.RoleRef.Kind == "ClusterRole" {
					clusterrolename := clstrrolebind.RoleRef.Name
					clstrrole, err := conn.getClusterRole(ctx, clusterrolename)
					if err != nil {
						return nil, err
					} else {
						clstrroleTemp, err := getCpConvertedTemplate(clstrrole, clstrrole.Kind)
						if err != nil {
							utils.Error.Println(err)
							return nil, err
						} else {

							clstrrolebindTemp, err := getCpConvertedTemplate(clstrrolebind, clstrrolebind.Kind)
							if err != nil {
								utils.Error.Println(err)
								return nil, err
							}
							if !isAlreadyExist(clstrrolebindTemp.NameSpace, clstrrolebindTemp.ServiceSubType, clstrrolebindTemp.Name) {
								clstrrolebindTemp.BeforeServices = append(clstrrolebindTemp.BeforeServices, clstrroleTemp.ServiceId)
								clstrroleTemp.AfterServices = append(clstrroleTemp.AfterServices, clstrrolebindTemp.ServiceId)

								clstrrolebindTemp.BeforeServices = append(clstrrolebindTemp.BeforeServices, svcAccTemp.ServiceId)
								svcAccTemp.AfterServices = append(svcAccTemp.AfterServices, clstrrolebindTemp.ServiceId)

								rbacServiceTemplates = append(rbacServiceTemplates, clstrrolebindTemp)
								rbacServiceTemplates = append(rbacServiceTemplates, clstrroleTemp)
							} else {
								clstrrolebindTemp.BeforeServices = append(clstrrolebindTemp.BeforeServices, svcAccTemp.ServiceId)
								svcAccTemp.AfterServices = append(svcAccTemp.AfterServices, clstrrolebindTemp.ServiceId)
							}
						}
					}
				}
				break
			}
		}
	}

	rolebindings, err := conn.getAllRoleBindings(ctx, namespace)
	if err != nil {
		return nil, err
	}

	for _, rolebinding := range rolebindings.Items {
		for _, sub := range rolebinding.Subjects {
			if sub.Kind == "ServiceAccount" && sub.Name == k8svcAcc.Name {
				if rolebinding.RoleRef.Kind == "Role" {
					rolename := rolebinding.RoleRef.Name
					role, err := conn.getRole(ctx, rolename, namespace)
					if err != nil {
						return nil, err
					} else {
						roleTemp, err := getCpConvertedTemplate(role, role.Kind)
						if err != nil {
							utils.Error.Println(err)
							return nil, err
						} else {
							rolebindTemp, err := getCpConvertedTemplate(rolebinding, rolebinding.Kind)
							if err != nil {
								utils.Error.Println(err)
								return nil, err
							}
							if !isAlreadyExist(rolebindTemp.NameSpace, rolebindTemp.ServiceSubType, rolebindTemp.Name) {
								rolebindTemp.BeforeServices = append(rolebindTemp.BeforeServices, roleTemp.ServiceId)
								roleTemp.AfterServices = append(roleTemp.AfterServices, rolebindTemp.ServiceId)

								rolebindTemp.BeforeServices = append(rolebindTemp.BeforeServices, svcAccTemp.ServiceId)
								svcAccTemp.AfterServices = append(svcAccTemp.AfterServices, rolebindTemp.ServiceId)

								rbacServiceTemplates = append(rbacServiceTemplates, rolebindTemp)
								rbacServiceTemplates = append(rbacServiceTemplates, roleTemp)
							} else {
								rolebindTemp.BeforeServices = append(rolebindTemp.BeforeServices, svcAccTemp.ServiceId)
								svcAccTemp.AfterServices = append(svcAccTemp.AfterServices, rolebindTemp.ServiceId)
							}
						}
					}
				}
				break
			}
		}
	}

	rbacServiceTemplates = append(rbacServiceTemplates, svcAccTemp)
	return rbacServiceTemplates, nil
}

func (conn *GrpcConn) getAllCronjobs(ctx context.Context, namespace string) (*v1beta1.CronJobList, error) {
	response, err := pb.NewK8SResourceClient(conn.Connection).GetK8SResource(ctx, &pb.K8SResourceRequest{
		ProjectId: conn.ProjectId,
		CompanyId: conn.CompanyId,
		Token:     conn.token,
		Command:   "kubectl",
		Args:      []string{"get", "cronjobs", "-n", namespace, "-o", "json"},
	})
	if err != nil {
		utils.Error.Println(err)
		return nil, errors.New("error from grpc server :" + err.Error())
	}

	var cronjobList *v1beta1.CronJobList
	err = json.Unmarshal(response.Resource, &cronjobList)
	if err != nil {
		utils.Error.Println(err)
		return nil, err
	}

	return cronjobList, nil
}

func (conn *GrpcConn) getAllDaemonsets(ctx context.Context, namespace string) (*v1.DaemonSetList, error) {
	response, err := pb.NewK8SResourceClient(conn.Connection).GetK8SResource(ctx, &pb.K8SResourceRequest{
		ProjectId: conn.ProjectId,
		CompanyId: conn.CompanyId,
		Token:     conn.token,
		Command:   "kubectl",
		Args:      []string{"get", "statefulsets", "-n", namespace, "-o", "json"},
	})
	if err != nil {
		utils.Error.Println(err)
		return nil, errors.New("error from grpc server :" + err.Error())
	}

	var daemonsetList *v1.DaemonSetList
	err = json.Unmarshal(response.Resource, &daemonsetList)
	if err != nil {
		utils.Error.Println(err)
		return nil, err
	}

	return daemonsetList, nil
}

func (conn *GrpcConn) getAllStatefulsets(ctx context.Context, namespace string) (*v1.StatefulSetList, error) {
	response, err := pb.NewK8SResourceClient(conn.Connection).GetK8SResource(ctx, &pb.K8SResourceRequest{
		ProjectId: conn.ProjectId,
		CompanyId: conn.CompanyId,
		Token:     conn.token,
		Command:   "kubectl",
		Args:      []string{"get", "statefulsets", "-n", namespace, "-o", "json"},
	})
	if err != nil {
		utils.Error.Println(err)
		return nil, errors.New("error from grpc server :" + err.Error())
	}

	var statefulsetList *v1.StatefulSetList
	err = json.Unmarshal(response.Resource, &statefulsetList)
	if err != nil {
		utils.Error.Println(err)
		return nil, err
	}

	return statefulsetList, nil
}

func (conn *GrpcConn) getAllDeployments(ctx context.Context, namespace string) (*v1.DeploymentList, error) {
	response, err := pb.NewK8SResourceClient(conn.Connection).GetK8SResource(ctx, &pb.K8SResourceRequest{
		ProjectId: conn.ProjectId,
		CompanyId: conn.CompanyId,
		Token:     conn.token,
		Command:   "kubectl",
		Args:      []string{"get", "deployments", "-n", namespace, "-o", "yaml"},
	})
	if err != nil {
		utils.Error.Println(err)
		return nil, errors.New("error from grpc server :" + err.Error())
	}

	var deploymentList *v1.DeploymentList
	err = yaml.Unmarshal(response.Resource, &deploymentList)
	if err != nil {
		utils.Error.Println(err)
		return nil, err
	}

	return deploymentList, nil
}

func (conn *GrpcConn) getAllJobs(ctx context.Context, namespace string) (*batch.JobList, error) {
	response, err := pb.NewK8SResourceClient(conn.Connection).GetK8SResource(ctx, &pb.K8SResourceRequest{
		ProjectId: conn.ProjectId,
		CompanyId: conn.CompanyId,
		Token:     conn.token,
		Command:   "kubectl",
		Args:      []string{"get", "job", "-n", namespace, "-o", "json"},
	})
	if err != nil {
		utils.Error.Println(err)
		return nil, errors.New("error from grpc server :" + err.Error())
	}

	var jobList *batch.JobList
	err = json.Unmarshal(response.Resource, &jobList)
	if err != nil {
		utils.Error.Println(err)
		return nil, err
	}

	return jobList, nil
}

func (conn *GrpcConn) getAllHpas(ctx context.Context, namespace string) (*autoscale.HorizontalPodAutoscalerList, error) {
	response, err := pb.NewK8SResourceClient(conn.Connection).GetK8SResource(ctx, &pb.K8SResourceRequest{
		ProjectId: conn.ProjectId,
		CompanyId: conn.CompanyId,
		Token:     conn.token,
		Command:   "kubectl",
		Args:      []string{"get", "hpa", "-n", namespace, "-o", "json"},
	})
	if err != nil {
		utils.Error.Println(err)
		return nil, errors.New("error from grpc server :" + err.Error())
	}

	var hpaList *autoscale.HorizontalPodAutoscalerList
	err = json.Unmarshal(response.Resource, &hpaList)
	if err != nil {
		utils.Error.Println(err)
		return nil, err
	}

	return hpaList, nil
}

func (conn *GrpcConn) getStorageClass(ctx context.Context, storgaClassName, namespace string) (*storage.StorageClass, error) {
	response, err := pb.NewK8SResourceClient(conn.Connection).GetK8SResource(ctx, &pb.K8SResourceRequest{
		ProjectId: conn.ProjectId,
		CompanyId: conn.CompanyId,
		Token:     conn.token,
		Command:   "kubectl",
		Args:      []string{"get", "storageclass", storgaClassName, "-n", namespace, "-o", "json"},
	})
	if err != nil {
		utils.Error.Println(err)
		return nil, errors.New("error from grpc server :" + err.Error())
	}

	var storageClass *storage.StorageClass
	err = json.Unmarshal(response.Resource, &storageClass)
	if err != nil {
		utils.Error.Println(err)
		return nil, err
	}

	return storageClass, nil
}

func (conn *GrpcConn) getKubernetesServices(ctx context.Context, key, value, namespace string) (*v2.ServiceList, error) {

	response, err := pb.NewK8SResourceClient(conn.Connection).GetK8SResource(ctx, &pb.K8SResourceRequest{
		ProjectId: conn.ProjectId,
		CompanyId: conn.CompanyId,
		Token:     conn.token,
		Command:   "kubectl",
		Args:      []string{"get", "svc", "-n", namespace, "-o", "json"},
	})
	if err != nil {
		utils.Error.Println(err)
		return nil, errors.New("error from grpc server :" + err.Error())
	}

	var kubeServiceList *v2.ServiceList
	err = json.Unmarshal(response.Resource, &kubeServiceList)
	if err != nil {
		utils.Error.Println(err)
		return nil, err
	}

	var serviceList = new(v2.ServiceList)
	for _, item := range kubeServiceList.Items {
		for kubeKey, kubeLabel := range item.Spec.Selector {
			if kubeKey == key && kubeLabel == value {
				serviceList.Items = append(serviceList.Items, item)
			}
		}
	}

	return serviceList, nil
}

func (conn *GrpcConn) getRole(ctx context.Context, rolename, namespace string) (*rbac.Role, error) {
	response, err := pb.NewK8SResourceClient(conn.Connection).GetK8SResource(ctx, &pb.K8SResourceRequest{
		ProjectId: conn.ProjectId,
		CompanyId: conn.CompanyId,
		Token:     conn.token,
		Command:   "kubectl",
		Args:      []string{"get", "roles", rolename, "-n", namespace, "-o", "yaml"},
	})
	if err != nil {
		utils.Error.Println(err)
		return nil, errors.New("error from grpc server :" + err.Error())
	}

	var role *rbac.Role
	err = yaml.Unmarshal(response.Resource, &role)
	if err != nil {
		utils.Error.Println(err)
		return nil, err
	}

	return role, nil
}

func (conn *GrpcConn) getAllRoleBindings(ctx context.Context, namespace string) (*rbac.RoleBindingList, error) {
	response, err := pb.NewK8SResourceClient(conn.Connection).GetK8SResource(ctx, &pb.K8SResourceRequest{
		ProjectId: conn.ProjectId,
		CompanyId: conn.CompanyId,
		Token:     conn.token,
		Command:   "kubectl",
		Args:      []string{"get", "rolebindings", "-n", namespace, "-o", "json"},
	})
	if err != nil {
		utils.Error.Println(err)
		return nil, errors.New("error from grpc server :" + err.Error())
	}

	var rolebindings *rbac.RoleBindingList
	err = json.Unmarshal(response.Resource, &rolebindings)
	if err != nil {
		utils.Error.Println(err)
		return nil, err
	}

	return rolebindings, nil
}

func (conn *GrpcConn) getPvc(ctx context.Context, pvcname, namespace string) (*v2.PersistentVolumeClaim, error) {
	response, err := pb.NewK8SResourceClient(conn.Connection).GetK8SResource(ctx, &pb.K8SResourceRequest{
		ProjectId: conn.ProjectId,
		CompanyId: conn.CompanyId,
		Token:     conn.token,
		Command:   "kubectl",
		Args:      []string{"get", "pvc", pvcname, "-n", namespace, "-o", "yaml"},
	})
	if err != nil {
		utils.Error.Println(err)
		return nil, errors.New("error from grpc server :" + err.Error())
	}

	var pvc *v2.PersistentVolumeClaim
	err = yaml.Unmarshal(response.Resource, &pvc)
	if err != nil {
		utils.Error.Println(err)
		return nil, err
	}

	return pvc, nil
}

func (conn *GrpcConn) getConfigMap(ctx context.Context, configmapname, namespace string) (*v2.ConfigMap, error) {
	response, err := pb.NewK8SResourceClient(conn.Connection).GetK8SResource(ctx, &pb.K8SResourceRequest{
		ProjectId: conn.ProjectId,
		CompanyId: conn.CompanyId,
		Token:     conn.token,
		Command:   "kubectl",
		Args:      []string{"get", "configmaps", configmapname, "-n", namespace, "-o", "json"},
	})
	if err != nil {
		utils.Error.Println(err)
		return nil, errors.New("error from grpc server :" + err.Error())
	}

	var configmap *v2.ConfigMap
	err = json.Unmarshal(response.Resource, &configmap)
	if err != nil {
		utils.Error.Println(err)
		return nil, err
	}

	return configmap, nil
}

func (conn *GrpcConn) getSvcAccount(ctx context.Context, svcname, namespace string) (*api.ServiceAccount, error) {
	response, err := pb.NewK8SResourceClient(conn.Connection).GetK8SResource(ctx, &pb.K8SResourceRequest{
		ProjectId: conn.ProjectId,
		CompanyId: conn.CompanyId,
		Token:     conn.token,
		Command:   "kubectl",
		Args:      []string{"get", "sa", svcname, "-n", namespace, "-o", "yaml"},
	})
	if err != nil {
		utils.Error.Println(err)
		return nil, errors.New("error from grpc server :" + err.Error())
	}

	var svcAcc *api.ServiceAccount
	err = yaml.Unmarshal(response.Resource, &svcAcc)
	if err != nil {
		utils.Error.Println(err)
		return nil, err
	}

	return svcAcc, nil
}

func (conn *GrpcConn) getSecret(ctx context.Context, secretname, namespace string) (*v2.Secret, error) {
	response, err := pb.NewK8SResourceClient(conn.Connection).GetK8SResource(ctx, &pb.K8SResourceRequest{
		ProjectId: conn.ProjectId,
		CompanyId: conn.CompanyId,
		Token:     conn.token,
		Command:   "kubectl",
		Args:      []string{"get", "secrets", secretname, "-n", namespace, "-o", "yaml"},
	})
	if err != nil {
		utils.Error.Println(err)
		return nil, errors.New("error from grpc server :" + err.Error())
	}

	var scrt *v2.Secret
	err = yaml.Unmarshal(response.Resource, &scrt)
	if err != nil {
		utils.Error.Println(err)
		return nil, err
	}

	return scrt, nil
}

func (conn *GrpcConn) getAllClusterRoleBindings(ctx context.Context) (*rbac.ClusterRoleBindingList, error) {
	response, err := pb.NewK8SResourceClient(conn.Connection).GetK8SResource(ctx, &pb.K8SResourceRequest{
		ProjectId: conn.ProjectId,
		CompanyId: conn.CompanyId,
		Token:     conn.token,
		Command:   "kubectl",
		Args:      []string{"get", "clusterrolebindings", "-o", "json"},
	})
	if err != nil {
		utils.Error.Println(err)
		return nil, errors.New("error from grpc server :" + err.Error())
	}

	var clusterrolebindings *rbac.ClusterRoleBindingList
	err = json.Unmarshal(response.Resource, &clusterrolebindings)
	if err != nil {
		utils.Error.Println(err)
		return nil, err
	}

	return clusterrolebindings, nil
}

func (conn *GrpcConn) getClusterRole(ctx context.Context, clusterrolename string) (*rbac.ClusterRole, error) {
	response, err := pb.NewK8SResourceClient(conn.Connection).GetK8SResource(ctx, &pb.K8SResourceRequest{
		ProjectId: conn.ProjectId,
		CompanyId: conn.CompanyId,
		Token:     conn.token,
		Command:   "kubectl",
		Args:      []string{"get", "clusterrole", clusterrolename, "-o", "yaml"},
	})
	if err != nil {
		utils.Error.Println(err)
		return nil, errors.New("error from grpc server :" + err.Error())
	}

	var clusterrole *rbac.ClusterRole
	err = yaml.Unmarshal(response.Resource, &clusterrole)
	if err != nil {
		utils.Error.Println(err)
		return nil, err
	}

	return clusterrole, nil
}

func isPortMatched(kubeSvc *v2.Service, Container *v2.Container) bool {
	for _, kubeSvcPort := range kubeSvc.Spec.Ports {
		for _, containerPort := range Container.Ports {
			if kubeSvcPort.Protocol == containerPort.Protocol && kubeSvcPort.TargetPort.IntVal == containerPort.ContainerPort {
				return true
			}
		}
	}
	return false
}

func isAlreadyExist(namespace, svcsubtype, name *string) bool {
	utils.Info.Printf("Checking existence of type :%v name :%v within namespace :%v", *svcsubtype, *name, *namespace)
	if *svcsubtype == "cluster_role_binding" {
		fmt.Println("hellow pakistan")
	}
	for _, val := range serviceTemplates {
		if (*val.NameSpace == *namespace || (val.NameSpace == nil && namespace == nil)) && *val.ServiceSubType == *svcsubtype && *val.Name == *name {
			return true
		}
	}
	return false
}

func getCpConvertedTemplate(data interface{}, kind string) (*types.ServiceTemplate, error) {

	var template *types.ServiceTemplate
	switch kind {
	case "Deployment":
		CpDeployment, err := convertToCPDeployment(data)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		bytes, err := json.Marshal(CpDeployment)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}

		err = json.Unmarshal(bytes, &template)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		id := strconv.Itoa(rand.Int())
		template.ServiceId = &id
	case "CronJob":
		bytes, err := json.Marshal(data)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		var cronjob batchv1.CronJob
		err = json.Unmarshal(bytes, &cronjob)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		CpCronJob, err := convertToCPCronJob(&cronjob)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}

		bytes, err = json.Marshal(CpCronJob)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		err = json.Unmarshal(bytes, &template)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		id := strconv.Itoa(rand.Int())
		template.ServiceId = &id
	case "Job":
		bytes, err := json.Marshal(data)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		var job batch.Job
		err = json.Unmarshal(bytes, &job)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		CpJob, err := convertToCPJob(&job)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		bytes, err = json.Marshal(CpJob)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		err = json.Unmarshal(bytes, &template)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		id := strconv.Itoa(rand.Int())
		template.ServiceId = &id
	case "DaemonSet":
		CpDaemonset, err := convertToCPDaemonSet(data)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		bytes, err := json.Marshal(CpDaemonset)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		err = json.Unmarshal(bytes, &template)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		id := strconv.Itoa(rand.Int())
		template.ServiceId = &id
	case "StatefulSet":
		CpStatefuleSet, err := convertToCPStatefulSet(data)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		bytes, err := json.Marshal(CpStatefuleSet)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		err = json.Unmarshal(bytes, &template)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		id := strconv.Itoa(rand.Int())
		template.ServiceId = &id
	case "Service":
		bytes, err := json.Marshal(data)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		var k8Service v2.Service
		err = json.Unmarshal(bytes, &k8Service)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		CpKubeService, err := convertToCPKubernetesService(&k8Service)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		bytes, err = json.Marshal(CpKubeService)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		err = json.Unmarshal(bytes, &template)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		id := strconv.Itoa(rand.Int())
		template.ServiceId = &id
		if isAlreadyExist(template.NameSpace, template.ServiceSubType, template.Name) {
			template = GetExistingService(template.NameSpace, template.ServiceSubType, template.Name)
		}
	case "HorizontalPodAutoscaler":
		bytes, err := json.Marshal(data)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		var hpa autoscale.HorizontalPodAutoscaler
		err = json.Unmarshal(bytes, &hpa)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		CpHpa, err := ConvertToCPHPA(&hpa)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		bytes, err = json.Marshal(CpHpa)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		err = json.Unmarshal(bytes, &template)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		id := strconv.Itoa(rand.Int())
		template.ServiceId = &id
	case "ConfigMap":
		bytes, err := json.Marshal(data)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		var configmap v2.ConfigMap
		err = json.Unmarshal(bytes, &configmap)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		CpConfigMap, err := ConvertToCPConfigMap(&configmap)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		bytes, err = json.Marshal(CpConfigMap)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		err = json.Unmarshal(bytes, &template)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		id := strconv.Itoa(rand.Int())
		template.ServiceId = &id
		if isAlreadyExist(template.NameSpace, template.ServiceSubType, template.Name) {
			template = GetExistingService(template.NameSpace, template.ServiceSubType, template.Name)
		}
	case "Secret":
		bytes, err := json.Marshal(data)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		var secret v2.Secret
		err = json.Unmarshal(bytes, &secret)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		CpSecret, err := ConvertToCPSecret(&secret)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		bytes, err = json.Marshal(CpSecret)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		err = json.Unmarshal(bytes, &template)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		id := strconv.Itoa(rand.Int())
		template.ServiceId = &id
		if isAlreadyExist(template.NameSpace, template.ServiceSubType, template.Name) {
			template = GetExistingService(template.NameSpace, template.ServiceSubType, template.Name)
		}
	case "ServiceAccount":
		bytes, err := json.Marshal(data)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		var serviceaccount v2.ServiceAccount
		err = json.Unmarshal(bytes, &serviceaccount)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		CpServiceAccount, err := convertToCPServiceAccount(&serviceaccount)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		bytes, err = json.Marshal(CpServiceAccount)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		err = json.Unmarshal(bytes, &template)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		id := strconv.Itoa(rand.Int())
		template.ServiceId = &id
		if isAlreadyExist(template.NameSpace, template.ServiceSubType, template.Name) {
			template = GetExistingService(template.NameSpace, template.ServiceSubType, template.Name)
		}
	case "Role":
		bytes, err := json.Marshal(data)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		var role rbac.Role
		err = json.Unmarshal(bytes, &role)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		CpRole, err := ConvertToCPRole(&role)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		bytes, err = json.Marshal(CpRole)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		err = json.Unmarshal(bytes, &template)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		id := strconv.Itoa(rand.Int())
		template.ServiceId = &id
		if isAlreadyExist(template.NameSpace, template.ServiceSubType, template.Name) {
			template = GetExistingService(template.NameSpace, template.ServiceSubType, template.Name)
		}
	case "RoleBinding":
		bytes, err := json.Marshal(data)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		var roleBinding rbac.RoleBinding
		err = json.Unmarshal(bytes, &roleBinding)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		CpRoleBinding, err := ConvertToCPRoleBinding(&roleBinding)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		bytes, err = json.Marshal(CpRoleBinding)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		err = json.Unmarshal(bytes, &template)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		id := strconv.Itoa(rand.Int())
		template.ServiceId = &id
		if isAlreadyExist(template.NameSpace, template.ServiceSubType, template.Name) {
			template = GetExistingService(template.NameSpace, template.ServiceSubType, template.Name)
		}
	case "ClusterRole":
		bytes, err := json.Marshal(data)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		var clusterRole rbac.ClusterRole
		err = json.Unmarshal(bytes, &clusterRole)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		CpClusterRole, err := ConvertToCPClusterRole(&clusterRole)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		bytes, err = json.Marshal(CpClusterRole)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		err = json.Unmarshal(bytes, &template)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		id := strconv.Itoa(rand.Int())
		template.ServiceId = &id
		if isAlreadyExist(template.NameSpace, template.ServiceSubType, template.Name) {
			template = GetExistingService(template.NameSpace, template.ServiceSubType, template.Name)
		}
	case "ClusterRoleBinding":
		bytes, err := json.Marshal(data)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		var clusterRoleBinding rbac.ClusterRoleBinding
		err = json.Unmarshal(bytes, &clusterRoleBinding)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		CpClusterRoleBinding, err := ConvertToCPClusterRoleBinding(&clusterRoleBinding)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		bytes, err = json.Marshal(CpClusterRoleBinding)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		err = json.Unmarshal(bytes, &template)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		id := strconv.Itoa(rand.Int())
		template.ServiceId = &id
		if isAlreadyExist(template.NameSpace, template.ServiceSubType, template.Name) {
			template = GetExistingService(template.NameSpace, template.ServiceSubType, template.Name)
		}
	case "PersistentVolume":
		bytes, err := json.Marshal(data)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		var persistenVolume v2.PersistentVolume
		err = json.Unmarshal(bytes, &persistenVolume)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		CpPersistentVolume, err := convertToCPPersistentVolume(&persistenVolume)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		bytes, err = json.Marshal(CpPersistentVolume)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		err = json.Unmarshal(bytes, &template)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		id := strconv.Itoa(rand.Int())
		template.ServiceId = &id
	case "PersistentVolumeClaim":
		bytes, err := json.Marshal(data)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		var persistenVolumeClaim v2.PersistentVolumeClaim
		err = json.Unmarshal(bytes, &persistenVolumeClaim)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		CpPVC, err := convertToCPPersistentVolumeClaim(&persistenVolumeClaim)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		bytes, err = json.Marshal(CpPVC)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		err = json.Unmarshal(bytes, &template)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		id := strconv.Itoa(rand.Int())
		template.ServiceId = &id
		if isAlreadyExist(template.NameSpace, template.ServiceSubType, template.Name) {
			template = GetExistingService(template.NameSpace, template.ServiceSubType, template.Name)
		}
	case "StorageClass":
		bytes, err := json.Marshal(data)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		var storageClass storage.StorageClass
		err = json.Unmarshal(bytes, &storageClass)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		CpStorageClass, err := convertToCPStorageClass(&storageClass)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		bytes, err = json.Marshal(CpStorageClass)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		err = json.Unmarshal(bytes, &template)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		id := strconv.Itoa(rand.Int())
		template.ServiceId = &id
		if isAlreadyExist(template.NameSpace, template.ServiceSubType, template.Name) {
			template = GetExistingService(template.NameSpace, template.ServiceSubType, template.Name)
		}
	default:
		utils.Info.Println("Kind does not exist in defined switch cases")
		return nil, errors.New("type does not exit")
	}

	return template, nil

}
