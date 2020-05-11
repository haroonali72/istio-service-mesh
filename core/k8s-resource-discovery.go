package core

import (
	"bitbucket.org/cloudplex-devs/istio-service-mesh/constants"
	"bitbucket.org/cloudplex-devs/istio-service-mesh/utils"
	meshConstants "bitbucket.org/cloudplex-devs/microservices-mesh-engine/constants"
	pb "bitbucket.org/cloudplex-devs/microservices-mesh-engine/core/services/proto"
	svcTypes "bitbucket.org/cloudplex-devs/microservices-mesh-engine/types"
	meshTypes "bitbucket.org/cloudplex-devs/microservices-mesh-engine/types/services"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/build/kubernetes/api"
	"google.golang.org/grpc"
	istioClient "istio.io/client-go/pkg/apis/networking/v1alpha3"
	v1 "k8s.io/api/apps/v1"
	autoscale "k8s.io/api/autoscaling/v1"
	batch "k8s.io/api/batch/v1"
	"k8s.io/api/batch/v1beta1"
	v2 "k8s.io/api/core/v1"
	rbac "k8s.io/api/rbac/v1"
	storage "k8s.io/api/storage/v1"
	"math/rand"
	"reflect"
	"sigs.k8s.io/yaml"
	"strconv"
	"strings"
	"sync"
)

type GrpcConn struct {
	Connection *grpc.ClientConn
	ProjectId  string
	CompanyId  string
	token      string
}

var serviceTemplates []*svcTypes.ServiceTemplate

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
		statefulSetList, err := grpcConn.getAllStatefulsets(ctx, namespace)
		if err != nil {
			return &pb.K8SResourceResponse{}, err
		}
		err = grpcConn.statefulsetsK8sToCp(ctx, statefulSetList.Items, &wg)
		if err != nil {
			return &pb.K8SResourceResponse{}, err
		}
		//statefulsets

		//daemonsets
		daemonSetList, err := grpcConn.getAllDaemonsets(ctx, namespace)
		if err != nil {
			return &pb.K8SResourceResponse{}, err
		}
		err = grpcConn.daemonsetK8sToCp(ctx, daemonSetList.Items, &wg)
		if err != nil {
			return &pb.K8SResourceResponse{}, err
		}
		//daemonsets

		//cronjobs
		cronJobList, err := grpcConn.getAllCronjobs(ctx, namespace)
		if err != nil {
			return &pb.K8SResourceResponse{}, err
		}
		err = grpcConn.cronjobK8sToCp(ctx, cronJobList.Items, &wg)
		if err != nil {
			return &pb.K8SResourceResponse{}, err
		}
		//cronjobs

		//jobs
		JobList, err := grpcConn.getAllJobs(ctx, namespace)
		if err != nil {
			return &pb.K8SResourceResponse{}, err
		}
		err = grpcConn.jobK8sToCp(ctx, JobList.Items, &wg)
		if err != nil {
			return &pb.K8SResourceResponse{}, err
		}
		//jobs
	}

	wg.Wait()
	bytes, err := json.Marshal(serviceTemplates)
	if err != nil {
		return &pb.K8SResourceResponse{}, err
	}

	serviceTemplates = nil
	response.Resource = bytes
	return response, err
}

func (conn *GrpcConn) jobK8sToCp(ctx context.Context, jobs []batch.Job, wg *sync.WaitGroup) error {
	for _, job := range jobs {
		wg.Add(1)
		go conn.ResolveJobDependencies(job, wg, ctx)
	}
	return nil
}

func (conn *GrpcConn) ResolveJobDependencies(job batch.Job, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	utils.Info.Printf("Resolving dependencies of job :%v within namespace : %v", job.Name, job.Namespace)
	jobTemp, err := getCpConvertedTemplate(job, job.Kind)
	if err != nil {
		utils.Error.Println(err)
		return
	}
	namespace := job.Namespace
	if job.Spec.Template.Spec.ServiceAccountName != "" {
		svcname := job.Spec.Template.Spec.ServiceAccountName
		svcaccount, err := conn.getSvcAccount(ctx, svcname, namespace)
		if err != nil {
			return
		}
		svcAccTemp, err := getCpConvertedTemplate(svcaccount, svcaccount.Kind)
		if err != nil {
			utils.Error.Println(err)
			return
		}
		if !isAlreadyExist(svcAccTemp.Namespace, svcAccTemp.ServiceSubType, svcAccTemp.Name) {
			rbacDependencies, err := conn.getK8sRbacResources(ctx, namespace, svcaccount, svcAccTemp)
			if err != nil {
				utils.Error.Println(err)
				return
			}

			for _, rbacTemp := range rbacDependencies {
				if rbacTemp.ServiceSubType == meshConstants.ServiceAccount {
					jobTemp.BeforeServices = append(jobTemp.BeforeServices, &rbacTemp.ServiceId)
					rbacTemp.AfterServices = append(rbacTemp.AfterServices, &jobTemp.ServiceId)
				}
				serviceTemplates = append(serviceTemplates, rbacTemp)
			}
		} else {
			service := GetExistingService(svcAccTemp.Namespace, svcAccTemp.ServiceSubType, svcAccTemp.Name)
			jobTemp.BeforeServices = append(jobTemp.BeforeServices, &service.ServiceId)
			service.AfterServices = append(service.AfterServices, &jobTemp.ServiceId)
		}
	}

	//image pull secrets
	for _, objRef := range job.Spec.Template.Spec.ImagePullSecrets {
		secretname := objRef.Name
		secret, err := conn.getSecret(ctx, secretname, namespace)
		if err != nil {
			return
		}
		secretTemp, err := getCpConvertedTemplate(secret, secret.Kind)
		if err != nil {
			utils.Error.Println(err)
			return
		}
		if !isAlreadyExist(secretTemp.Namespace, secretTemp.ServiceSubType, secretTemp.Name) {
			jobTemp.BeforeServices = append(jobTemp.BeforeServices, &secretTemp.ServiceId)
			secretTemp.AfterServices = append(secretTemp.AfterServices, &jobTemp.ServiceId)
			serviceTemplates = append(serviceTemplates, secretTemp)
		} else {
			isSameJob := false
			for _, serviceId := range secretTemp.AfterServices {
				if *serviceId == jobTemp.ServiceId {
					isSameJob = true
					break
				}
			}
			if !isSameJob {
				jobTemp.BeforeServices = append(jobTemp.BeforeServices, &secretTemp.ServiceId)
				secretTemp.AfterServices = append(secretTemp.AfterServices, &jobTemp.ServiceId)
			}
		}
	}

	//hpa finding
	hpaList, err := conn.getAllHpas(ctx, namespace)
	if err != nil {
		return
	}
	if len(hpaList.Items) > 0 {
		for _, hpa := range hpaList.Items {
			if hpa.Spec.ScaleTargetRef.APIVersion == job.APIVersion && strings.ToLower(hpa.Spec.ScaleTargetRef.Kind) == strings.ToLower(job.Kind) && hpa.Spec.ScaleTargetRef.Name == job.Name {
				hpaTemplate, err := getCpConvertedTemplate(hpa, hpa.Kind)
				if err != nil {
					utils.Error.Println(err)
					return
				}
				hpaTemplate.BeforeServices = append(hpaTemplate.BeforeServices, &jobTemp.ServiceId)
				jobTemp.AfterServices = append(jobTemp.AfterServices, &hpaTemplate.ServiceId)
				serviceTemplates = append(serviceTemplates, hpaTemplate)
			}
		}
	}

	//kubernetes service depecndency findings
	var kubeSvcList []*v2.Service
	labels := make(map[string]string)
	for key, value := range job.Spec.Template.Labels {
		labels[key] = value
		kubesvclist, err := conn.getKubernetesServices(ctx, key, value, namespace)
		if err != nil {
			return
		}
		for _, kubesvc := range kubesvclist.Items {
			kubeSvcList = append(kubeSvcList, &kubesvc)
		}
	}

	//container dependency finding
	for _, container := range job.Spec.Template.Spec.Containers {

		//resolving dependencies of kubernetes service
		if kubeSvcList != nil && len(kubeSvcList) > 0 {
			for _, kubeSvc := range kubeSvcList {
				if isPortMatched(kubeSvc, &container) {
					k8serviceTemp, err := getCpConvertedTemplate(kubeSvc, kubeSvc.Kind)
					if err != nil {
						utils.Error.Println(err)
						return
					}
					//istio components creation
					istioSvcTemps, err := CreateIstioComponents(k8serviceTemp, labels)
					if err != nil {
						utils.Error.Println(err)
						return
					}
					for _, istioSvc := range istioSvcTemps {
						serviceTemplates = append(serviceTemplates, istioSvc)
					}
					//istio components creation
					if !isAlreadyExist(k8serviceTemp.Namespace, k8serviceTemp.ServiceSubType, k8serviceTemp.Name) {
						k8serviceTemp.BeforeServices = append(k8serviceTemp.BeforeServices, &jobTemp.ServiceId)
						jobTemp.AfterServices = append(jobTemp.AfterServices, &k8serviceTemp.ServiceId)
						serviceTemplates = append(serviceTemplates, k8serviceTemp)
					} else {
						isSameJob := false
						for _, serviceId := range k8serviceTemp.BeforeServices {
							if *serviceId == jobTemp.ServiceId {
								isSameJob = true
								break
							}
						}
						if !isSameJob {
							k8serviceTemp.BeforeServices = append(k8serviceTemp.BeforeServices, &jobTemp.ServiceId)
							jobTemp.AfterServices = append(jobTemp.AfterServices, &k8serviceTemp.ServiceId)
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
					return
				}
				secretTemp, err := getCpConvertedTemplate(secret, secret.Kind)
				if err != nil {
					utils.Error.Println(err)
					return
				}
				if !isAlreadyExist(secretTemp.Namespace, secretTemp.ServiceSubType, secretTemp.Name) {
					jobTemp.BeforeServices = append(jobTemp.BeforeServices, &secretTemp.ServiceId)
					secretTemp.AfterServices = append(secretTemp.AfterServices, &jobTemp.ServiceId)
					serviceTemplates = append(serviceTemplates, secretTemp)
				} else {
					isSameJob := false
					for _, serviceId := range secretTemp.AfterServices {
						if *serviceId == jobTemp.ServiceId {
							isSameJob = true
							break
						}
					}
					if !isSameJob {
						jobTemp.BeforeServices = append(jobTemp.BeforeServices, &secretTemp.ServiceId)
						secretTemp.AfterServices = append(secretTemp.AfterServices, &jobTemp.ServiceId)
					}
				}
			} else if env.ValueFrom != nil && env.ValueFrom.ConfigMapKeyRef != nil {
				configmapname := env.ValueFrom.ConfigMapKeyRef.Name
				configmap, err := conn.getConfigMap(ctx, configmapname, namespace)
				if err != nil {
					return
				}
				configmapTemp, err := getCpConvertedTemplate(configmap, configmap.Kind)
				if err != nil {
					utils.Error.Println(err)
					return
				}
				if !isAlreadyExist(configmapTemp.Namespace, configmapTemp.ServiceSubType, configmapTemp.Name) {
					jobTemp.BeforeServices = append(jobTemp.BeforeServices, &configmapTemp.ServiceId)
					configmapTemp.AfterServices = append(configmapTemp.AfterServices, &jobTemp.ServiceId)
					serviceTemplates = append(serviceTemplates, configmapTemp)
				} else {
					isSameJob := false
					for _, serviceId := range configmapTemp.AfterServices {
						if *serviceId == jobTemp.ServiceId {
							isSameJob = true
							break
						}
					}
					if !isSameJob {
						jobTemp.BeforeServices = append(jobTemp.BeforeServices, &configmapTemp.ServiceId)
						configmapTemp.AfterServices = append(configmapTemp.AfterServices, &jobTemp.ServiceId)
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
				return
			}
			secretTemp, err := getCpConvertedTemplate(secret, secret.Kind)
			if err != nil {
				utils.Error.Println(err)
				return
			}
			if !isAlreadyExist(secretTemp.Namespace, secretTemp.ServiceSubType, secretTemp.Name) {
				jobTemp.BeforeServices = append(jobTemp.BeforeServices, &secretTemp.ServiceId)
				secretTemp.AfterServices = append(secretTemp.AfterServices, &jobTemp.ServiceId)
				serviceTemplates = append(serviceTemplates, secretTemp)
			} else {
				isSameJob := false
				for _, serviceId := range secretTemp.AfterServices {
					if *serviceId == jobTemp.ServiceId {
						isSameJob = true
						break
					}
				}
				if !isSameJob {
					jobTemp.BeforeServices = append(jobTemp.BeforeServices, &secretTemp.ServiceId)
					secretTemp.AfterServices = append(secretTemp.AfterServices, &jobTemp.ServiceId)
				}
			}
		} else if vol.ConfigMap != nil {
			configmapname := vol.ConfigMap.Name
			configmap, err := conn.getConfigMap(ctx, configmapname, namespace)
			if err != nil {
				return
			}
			configmapTemp, err := getCpConvertedTemplate(configmap, configmap.Kind)
			if err != nil {
				utils.Error.Println(err)
				return
			}
			if !isAlreadyExist(configmapTemp.Namespace, configmapTemp.ServiceSubType, configmapTemp.Name) {
				jobTemp.BeforeServices = append(jobTemp.BeforeServices, &configmapTemp.ServiceId)
				configmapTemp.AfterServices = append(configmapTemp.AfterServices, &jobTemp.ServiceId)
				serviceTemplates = append(serviceTemplates, configmapTemp)
			} else {
				isSameJob := false
				for _, serviceId := range configmapTemp.AfterServices {
					if *serviceId == jobTemp.ServiceId {
						isSameJob = true
						break
					}
				}
				if !isSameJob {
					jobTemp.BeforeServices = append(jobTemp.BeforeServices, &configmapTemp.ServiceId)
					configmapTemp.AfterServices = append(configmapTemp.AfterServices, &jobTemp.ServiceId)
				}
			}
		} else if vol.PersistentVolumeClaim != nil {
			pvcname := vol.PersistentVolumeClaim.ClaimName
			pvc, err := conn.getPvc(ctx, pvcname, namespace)
			if err != nil {
				return
			}
			pvcTemp, err := getCpConvertedTemplate(pvc, pvc.Kind)
			if err != nil {
				utils.Error.Println(err)
				return
			}
			if !isAlreadyExist(pvcTemp.Namespace, pvcTemp.ServiceSubType, pvcTemp.Name) {
				jobTemp.BeforeServices = append(jobTemp.BeforeServices, &pvcTemp.ServiceId)
				pvcTemp.AfterServices = append(pvcTemp.AfterServices, &jobTemp.ServiceId)
				serviceTemplates = append(serviceTemplates, pvcTemp)
			} else {
				service := GetExistingService(pvcTemp.Namespace, pvcTemp.ServiceSubType, pvcTemp.Name)
				jobTemp.BeforeServices = append(jobTemp.BeforeServices, &service.ServiceId)
				service.AfterServices = append(service.AfterServices, &jobTemp.ServiceId)
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
				if !isAlreadyExist(storageClassTemp.Namespace, storageClassTemp.ServiceSubType, storageClassTemp.Name) {
					pvcTemp.BeforeServices = append(pvcTemp.BeforeServices, &storageClassTemp.ServiceId)
					storageClassTemp.AfterServices = append(storageClassTemp.AfterServices, &pvcTemp.ServiceId)
					serviceTemplates = append(serviceTemplates, storageClassTemp)
				} else {
					isPVCexist := false
					for _, serviceId := range storageClassTemp.AfterServices {
						if *serviceId == pvcTemp.ServiceId {
							isPVCexist = true
							break
						}
					}
					if !isPVCexist {
						pvcTemp.BeforeServices = append(pvcTemp.BeforeServices, &storageClassTemp.ServiceId)
						storageClassTemp.AfterServices = append(storageClassTemp.AfterServices, &pvcTemp.ServiceId)
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
						return
					}
					configmapTemp, err := getCpConvertedTemplate(configmap, configmap.Kind)
					if err != nil {
						utils.Error.Println(err)
						return
					}
					if !isAlreadyExist(configmapTemp.Namespace, configmapTemp.ServiceSubType, configmapTemp.Name) {
						jobTemp.BeforeServices = append(jobTemp.BeforeServices, &configmapTemp.ServiceId)
						configmapTemp.AfterServices = append(configmapTemp.AfterServices, &jobTemp.ServiceId)
						serviceTemplates = append(serviceTemplates, configmapTemp)
					} else {
						isSameJob := false
						for _, serviceId := range configmapTemp.AfterServices {
							if *serviceId == jobTemp.ServiceId {
								isSameJob = true
								break
							}
						}
						if !isSameJob {
							jobTemp.BeforeServices = append(jobTemp.BeforeServices, &configmapTemp.ServiceId)
							configmapTemp.AfterServices = append(configmapTemp.AfterServices, &jobTemp.ServiceId)
						}
					}
				} else if source.Secret != nil {
					secretname := vol.Secret.SecretName
					secret, err := conn.getSecret(ctx, secretname, namespace)
					if err != nil {
						return
					}
					secretTemp, err := getCpConvertedTemplate(secret, secret.Kind)
					if err != nil {
						utils.Error.Println(err)
						return
					}
					if !isAlreadyExist(secretTemp.Namespace, secretTemp.ServiceSubType, secretTemp.Name) {
						jobTemp.BeforeServices = append(jobTemp.BeforeServices, &secretTemp.ServiceId)
						secretTemp.AfterServices = append(secretTemp.AfterServices, &jobTemp.ServiceId)
						serviceTemplates = append(serviceTemplates, secretTemp)
					} else {
						isSameJob := false
						for _, serviceId := range secretTemp.AfterServices {
							if *serviceId == jobTemp.ServiceId {
								isSameJob = true
								break
							}
						}
						if !isSameJob {
							jobTemp.BeforeServices = append(jobTemp.BeforeServices, &secretTemp.ServiceId)
							secretTemp.AfterServices = append(secretTemp.AfterServices, &jobTemp.ServiceId)
						}
					}
				}
			}
		}
	}

	serviceTemplates = append(serviceTemplates, jobTemp)
}

func (conn *GrpcConn) ResolveCronJobDependencies(cronjob v1beta1.CronJob, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	utils.Info.Printf("Resolving dependencies of cronjob :%v within namespace : %v", cronjob.Name, cronjob.Namespace)
	cronjobTemp, err := getCpConvertedTemplate(cronjob, cronjob.Kind)
	if err != nil {
		utils.Error.Println(err)
		return
	}
	namespace := cronjob.Namespace
	if cronjob.Spec.JobTemplate.Spec.Template.Spec.ServiceAccountName != "" {
		svcname := cronjob.Spec.JobTemplate.Spec.Template.Spec.ServiceAccountName
		svcaccount, err := conn.getSvcAccount(ctx, svcname, namespace)
		if err != nil {
			return
		}
		svcAccTemp, err := getCpConvertedTemplate(svcaccount, svcaccount.Kind)
		if err != nil {
			utils.Error.Println(err)
			return
		}
		if !isAlreadyExist(svcAccTemp.Namespace, svcAccTemp.ServiceSubType, svcAccTemp.Name) {
			rbacDependencies, err := conn.getK8sRbacResources(ctx, namespace, svcaccount, svcAccTemp)
			if err != nil {
				utils.Error.Println(err)
				return
			}

			for _, rbacTemp := range rbacDependencies {
				if rbacTemp.ServiceSubType == meshConstants.ServiceAccount {
					cronjobTemp.BeforeServices = append(cronjobTemp.BeforeServices, &rbacTemp.ServiceId)
					rbacTemp.AfterServices = append(rbacTemp.AfterServices, &cronjobTemp.ServiceId)
				}
				serviceTemplates = append(serviceTemplates, rbacTemp)
			}
		} else {
			service := GetExistingService(svcAccTemp.Namespace, svcAccTemp.ServiceSubType, svcAccTemp.Name)
			cronjobTemp.BeforeServices = append(cronjobTemp.BeforeServices, &service.ServiceId)
			service.AfterServices = append(service.AfterServices, &cronjobTemp.ServiceId)
		}
	}

	//image pull secrets
	for _, objRef := range cronjob.Spec.JobTemplate.Spec.Template.Spec.ImagePullSecrets {
		secretname := objRef.Name
		secret, err := conn.getSecret(ctx, secretname, namespace)
		if err != nil {
			return
		}
		secretTemp, err := getCpConvertedTemplate(secret, secret.Kind)
		if err != nil {
			utils.Error.Println(err)
			return
		}
		if !isAlreadyExist(secretTemp.Namespace, secretTemp.ServiceSubType, secretTemp.Name) {
			cronjobTemp.BeforeServices = append(cronjobTemp.BeforeServices, &secretTemp.ServiceId)
			secretTemp.AfterServices = append(secretTemp.AfterServices, &cronjobTemp.ServiceId)
			serviceTemplates = append(serviceTemplates, secretTemp)
		} else {
			isSameCronJob := false
			for _, serviceId := range secretTemp.AfterServices {
				if *serviceId == cronjobTemp.ServiceId {
					isSameCronJob = true
					break
				}
			}
			if !isSameCronJob {
				cronjobTemp.BeforeServices = append(cronjobTemp.BeforeServices, &secretTemp.ServiceId)
				secretTemp.AfterServices = append(secretTemp.AfterServices, &cronjobTemp.ServiceId)
			}
		}
	}

	hpaList, err := conn.getAllHpas(ctx, namespace)
	if err != nil {
		return
	}
	if len(hpaList.Items) > 0 {
		for _, hpa := range hpaList.Items {
			if hpa.Spec.ScaleTargetRef.APIVersion == cronjob.APIVersion && strings.ToLower(hpa.Spec.ScaleTargetRef.Kind) == strings.ToLower(cronjob.Kind) && hpa.Spec.ScaleTargetRef.Name == cronjob.Name {
				hpaTemplate, err := getCpConvertedTemplate(hpa, hpa.Kind)
				if err != nil {
					utils.Error.Println(err)
					return
				}
				hpaTemplate.BeforeServices = append(hpaTemplate.BeforeServices, &cronjobTemp.ServiceId)
				cronjobTemp.AfterServices = append(cronjobTemp.AfterServices, &hpaTemplate.ServiceId)
				serviceTemplates = append(serviceTemplates, hpaTemplate)
			}
		}
	}

	//kubernetes service depecndency findings
	var kubeSvcList []*v2.Service
	labels := make(map[string]string)
	for key, value := range cronjob.Spec.JobTemplate.Spec.Template.Labels {
		labels[key] = value
		kubesvclist, err := conn.getKubernetesServices(ctx, key, value, namespace)
		if err != nil {
			return
		}
		for _, kubesvc := range kubesvclist.Items {
			kubeSvcList = append(kubeSvcList, &kubesvc)
		}
	}

	//container dependency finding
	for _, container := range cronjob.Spec.JobTemplate.Spec.Template.Spec.Containers {

		//resolving dependencies of kubernetes service
		if kubeSvcList != nil && len(kubeSvcList) > 0 {
			for _, kubeSvc := range kubeSvcList {
				if isPortMatched(kubeSvc, &container) {
					k8serviceTemp, err := getCpConvertedTemplate(kubeSvc, kubeSvc.Kind)
					if err != nil {
						utils.Error.Println(err)
						return
					}
					//istio components creation
					istioSvcTemps, err := CreateIstioComponents(k8serviceTemp, labels)
					if err != nil {
						utils.Error.Println(err)
						return
					}
					for _, istioSvc := range istioSvcTemps {
						serviceTemplates = append(serviceTemplates, istioSvc)
					}
					//istio components creation
					if !isAlreadyExist(k8serviceTemp.Namespace, k8serviceTemp.ServiceSubType, k8serviceTemp.Name) {
						k8serviceTemp.BeforeServices = append(k8serviceTemp.BeforeServices, &cronjobTemp.ServiceId)
						cronjobTemp.AfterServices = append(cronjobTemp.AfterServices, &k8serviceTemp.ServiceId)
						serviceTemplates = append(serviceTemplates, k8serviceTemp)
					} else {
						isSameCronJob := false
						for _, serviceId := range k8serviceTemp.BeforeServices {
							if *serviceId == cronjobTemp.ServiceId {
								isSameCronJob = true
								break
							}
						}
						if !isSameCronJob {
							k8serviceTemp.BeforeServices = append(k8serviceTemp.BeforeServices, &cronjobTemp.ServiceId)
							cronjobTemp.AfterServices = append(cronjobTemp.AfterServices, &k8serviceTemp.ServiceId)
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
					return
				}
				secretTemp, err := getCpConvertedTemplate(secret, secret.Kind)
				if err != nil {
					utils.Error.Println(err)
					return
				}
				if !isAlreadyExist(secretTemp.Namespace, secretTemp.ServiceSubType, secretTemp.Name) {
					cronjobTemp.BeforeServices = append(cronjobTemp.BeforeServices, &secretTemp.ServiceId)
					secretTemp.AfterServices = append(secretTemp.AfterServices, &cronjobTemp.ServiceId)
					serviceTemplates = append(serviceTemplates, secretTemp)
				} else {
					isSameCronJob := false
					for _, serviceId := range secretTemp.AfterServices {
						if *serviceId == cronjobTemp.ServiceId {
							isSameCronJob = true
							break
						}
					}
					if !isSameCronJob {
						cronjobTemp.BeforeServices = append(cronjobTemp.BeforeServices, &secretTemp.ServiceId)
						secretTemp.AfterServices = append(secretTemp.AfterServices, &cronjobTemp.ServiceId)
					}
				}
			} else if env.ValueFrom != nil && env.ValueFrom.ConfigMapKeyRef != nil {
				configmapname := env.ValueFrom.ConfigMapKeyRef.Name
				configmap, err := conn.getConfigMap(ctx, configmapname, namespace)
				if err != nil {
					return
				}
				configmapTemp, err := getCpConvertedTemplate(configmap, configmap.Kind)
				if err != nil {
					utils.Error.Println(err)
					return
				}
				if !isAlreadyExist(configmapTemp.Namespace, configmapTemp.ServiceSubType, configmapTemp.Name) {
					cronjobTemp.BeforeServices = append(cronjobTemp.BeforeServices, &configmapTemp.ServiceId)
					configmapTemp.AfterServices = append(configmapTemp.AfterServices, &cronjobTemp.ServiceId)
					serviceTemplates = append(serviceTemplates, configmapTemp)
				} else {
					isSameCronJob := false
					for _, serviceId := range configmapTemp.AfterServices {
						if *serviceId == cronjobTemp.ServiceId {
							isSameCronJob = true
							break
						}
					}
					if !isSameCronJob {
						cronjobTemp.BeforeServices = append(cronjobTemp.BeforeServices, &configmapTemp.ServiceId)
						configmapTemp.AfterServices = append(configmapTemp.AfterServices, &cronjobTemp.ServiceId)
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
				return
			}
			secretTemp, err := getCpConvertedTemplate(secret, secret.Kind)
			if err != nil {
				utils.Error.Println(err)
				return
			}
			if !isAlreadyExist(secretTemp.Namespace, secretTemp.ServiceSubType, secretTemp.Name) {
				cronjobTemp.BeforeServices = append(cronjobTemp.BeforeServices, &secretTemp.ServiceId)
				secretTemp.AfterServices = append(secretTemp.AfterServices, &cronjobTemp.ServiceId)
				serviceTemplates = append(serviceTemplates, secretTemp)
			} else {
				isSameCronJob := false
				for _, serviceId := range secretTemp.AfterServices {
					if *serviceId == cronjobTemp.ServiceId {
						isSameCronJob = true
						break
					}
				}
				if !isSameCronJob {
					cronjobTemp.BeforeServices = append(cronjobTemp.BeforeServices, &secretTemp.ServiceId)
					secretTemp.AfterServices = append(secretTemp.AfterServices, &cronjobTemp.ServiceId)
				}
			}
		} else if vol.ConfigMap != nil {
			configmapname := vol.ConfigMap.Name
			configmap, err := conn.getConfigMap(ctx, configmapname, namespace)
			if err != nil {
				return
			}
			configmapTemp, err := getCpConvertedTemplate(configmap, configmap.Kind)
			if err != nil {
				utils.Error.Println(err)
				return
			}
			if !isAlreadyExist(configmapTemp.Namespace, configmapTemp.ServiceSubType, configmapTemp.Name) {
				cronjobTemp.BeforeServices = append(cronjobTemp.BeforeServices, &configmapTemp.ServiceId)
				configmapTemp.AfterServices = append(configmapTemp.AfterServices, &cronjobTemp.ServiceId)
				serviceTemplates = append(serviceTemplates, configmapTemp)
			} else {
				isSameCronJob := false
				for _, serviceId := range configmapTemp.AfterServices {
					if *serviceId == cronjobTemp.ServiceId {
						isSameCronJob = true
						break
					}
				}
				if !isSameCronJob {
					cronjobTemp.BeforeServices = append(cronjobTemp.BeforeServices, &configmapTemp.ServiceId)
					configmapTemp.AfterServices = append(configmapTemp.AfterServices, &cronjobTemp.ServiceId)
				}
			}
		} else if vol.PersistentVolumeClaim != nil {
			pvcname := vol.PersistentVolumeClaim.ClaimName
			pvc, err := conn.getPvc(ctx, pvcname, namespace)
			if err != nil {
				return
			}
			pvcTemp, err := getCpConvertedTemplate(pvc, pvc.Kind)
			if err != nil {
				utils.Error.Println(err)
				return
			}
			if !isAlreadyExist(pvcTemp.Namespace, pvcTemp.ServiceSubType, pvcTemp.Name) {
				cronjobTemp.BeforeServices = append(cronjobTemp.BeforeServices, &pvcTemp.ServiceId)
				pvcTemp.AfterServices = append(pvcTemp.AfterServices, &cronjobTemp.ServiceId)
				serviceTemplates = append(serviceTemplates, pvcTemp)
			} else {
				service := GetExistingService(pvcTemp.Namespace, pvcTemp.ServiceSubType, pvcTemp.Name)
				cronjobTemp.BeforeServices = append(cronjobTemp.BeforeServices, &service.ServiceId)
				service.AfterServices = append(service.AfterServices, &cronjobTemp.ServiceId)
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
				if !isAlreadyExist(storageClassTemp.Namespace, storageClassTemp.ServiceSubType, storageClassTemp.Name) {
					pvcTemp.BeforeServices = append(pvcTemp.BeforeServices, &storageClassTemp.ServiceId)
					storageClassTemp.AfterServices = append(storageClassTemp.AfterServices, &pvcTemp.ServiceId)
					serviceTemplates = append(serviceTemplates, storageClassTemp)
				} else {
					isPVCexist := false
					for _, serviceId := range storageClassTemp.AfterServices {
						if *serviceId == pvcTemp.ServiceId {
							isPVCexist = true
							break
						}
					}
					if !isPVCexist {
						pvcTemp.BeforeServices = append(pvcTemp.BeforeServices, &storageClassTemp.ServiceId)
						storageClassTemp.AfterServices = append(storageClassTemp.AfterServices, &pvcTemp.ServiceId)
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
						return
					}
					configmapTemp, err := getCpConvertedTemplate(configmap, configmap.Kind)
					if err != nil {
						utils.Error.Println(err)
						return
					}
					if !isAlreadyExist(configmapTemp.Namespace, configmapTemp.ServiceSubType, configmapTemp.Name) {
						cronjobTemp.BeforeServices = append(cronjobTemp.BeforeServices, &configmapTemp.ServiceId)
						configmapTemp.AfterServices = append(configmapTemp.AfterServices, &cronjobTemp.ServiceId)
						serviceTemplates = append(serviceTemplates, configmapTemp)
					} else {
						isSameCronJob := false
						for _, serviceId := range configmapTemp.AfterServices {
							if *serviceId == cronjobTemp.ServiceId {
								isSameCronJob = true
								break
							}
						}
						if !isSameCronJob {
							cronjobTemp.BeforeServices = append(cronjobTemp.BeforeServices, &configmapTemp.ServiceId)
							configmapTemp.AfterServices = append(configmapTemp.AfterServices, &cronjobTemp.ServiceId)
						}
					}
				} else if source.Secret != nil {
					secretname := vol.Secret.SecretName
					secret, err := conn.getSecret(ctx, secretname, namespace)
					if err != nil {
						return
					}
					secretTemp, err := getCpConvertedTemplate(secret, secret.Kind)
					if err != nil {
						utils.Error.Println(err)
						return
					}
					if !isAlreadyExist(secretTemp.Namespace, secretTemp.ServiceSubType, secretTemp.Name) {
						cronjobTemp.BeforeServices = append(cronjobTemp.BeforeServices, &secretTemp.ServiceId)
						secretTemp.AfterServices = append(secretTemp.AfterServices, &cronjobTemp.ServiceId)
						serviceTemplates = append(serviceTemplates, secretTemp)
					} else {
						isSameCronJob := false
						for _, serviceId := range secretTemp.AfterServices {
							if *serviceId == cronjobTemp.ServiceId {
								isSameCronJob = true
								break
							}
						}
						if !isSameCronJob {
							cronjobTemp.BeforeServices = append(cronjobTemp.BeforeServices, &secretTemp.ServiceId)
							secretTemp.AfterServices = append(secretTemp.AfterServices, &cronjobTemp.ServiceId)
						}
					}
				}
			}
		}
	}

	serviceTemplates = append(serviceTemplates, cronjobTemp)
}

func (conn *GrpcConn) cronjobK8sToCp(ctx context.Context, cronjobs []v1beta1.CronJob, wg *sync.WaitGroup) error {
	for _, cronjob := range cronjobs {
		wg.Add(1)
		go conn.ResolveCronJobDependencies(cronjob, wg, ctx)
	}
	return nil
}

func (conn *GrpcConn) daemonsetK8sToCp(ctx context.Context, daemonsets []v1.DaemonSet, wg *sync.WaitGroup) error {
	for _, daemonset := range daemonsets {
		wg.Add(1)
		go conn.ResolveDaemonSetDependencies(daemonset, wg, ctx)
	}

	return nil
}

func (conn *GrpcConn) ResolveDaemonSetDependencies(daemonset v1.DaemonSet, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	utils.Info.Printf("Resolving dependencies of daemonset :%v within namespace : %v", daemonset.Name, daemonset.Namespace)
	daemonsetTemp, err := getCpConvertedTemplate(daemonset, daemonset.Kind)
	if err != nil {
		utils.Error.Println(err)
		return
	}
	namespace := daemonset.Namespace
	if daemonset.Spec.Template.Spec.ServiceAccountName != "" {
		svcname := daemonset.Spec.Template.Spec.ServiceAccountName
		svcaccount, err := conn.getSvcAccount(ctx, svcname, namespace)
		if err != nil {
			return
		}
		svcAccTemp, err := getCpConvertedTemplate(svcaccount, svcaccount.Kind)
		if err != nil {
			utils.Error.Println(err)
			return
		}
		if !isAlreadyExist(svcAccTemp.Namespace, svcAccTemp.ServiceSubType, svcAccTemp.Name) {
			rbacDependencies, err := conn.getK8sRbacResources(ctx, namespace, svcaccount, svcAccTemp)
			if err != nil {
				utils.Error.Println(err)
				return
			}

			var strpointer = new(string)
			*strpointer = "serviceaccount"

			for _, rbacTemp := range rbacDependencies {
				if rbacTemp.ServiceSubType == meshConstants.ServiceAccount {
					daemonsetTemp.BeforeServices = append(daemonsetTemp.BeforeServices, &rbacTemp.ServiceId)
					rbacTemp.AfterServices = append(rbacTemp.AfterServices, &daemonsetTemp.ServiceId)
				}
				serviceTemplates = append(serviceTemplates, rbacTemp)
			}
		} else {
			service := GetExistingService(svcAccTemp.Namespace, svcAccTemp.ServiceSubType, svcAccTemp.Name)
			daemonsetTemp.BeforeServices = append(daemonsetTemp.BeforeServices, &service.ServiceId)
			service.AfterServices = append(service.AfterServices, &daemonsetTemp.ServiceId)
		}
	}

	//image pull secrets
	for _, objRef := range daemonset.Spec.Template.Spec.ImagePullSecrets {
		secretname := objRef.Name
		secret, err := conn.getSecret(ctx, secretname, namespace)
		if err != nil {
			return
		}
		secretTemp, err := getCpConvertedTemplate(secret, secret.Kind)
		if err != nil {
			utils.Error.Println(err)
			return
		}
		if !isAlreadyExist(secretTemp.Namespace, secretTemp.ServiceSubType, secretTemp.Name) {
			daemonsetTemp.BeforeServices = append(daemonsetTemp.BeforeServices, &secretTemp.ServiceId)
			secretTemp.AfterServices = append(secretTemp.AfterServices, &daemonsetTemp.ServiceId)
			serviceTemplates = append(serviceTemplates, secretTemp)
		} else {
			isSameDaemonSet := false
			for _, serviceId := range secretTemp.AfterServices {
				if *serviceId == daemonsetTemp.ServiceId {
					isSameDaemonSet = true
					break
				}
			}
			if !isSameDaemonSet {
				daemonsetTemp.BeforeServices = append(daemonsetTemp.BeforeServices, &secretTemp.ServiceId)
				secretTemp.AfterServices = append(secretTemp.AfterServices, &daemonsetTemp.ServiceId)
			}
		}
	}

	//kubernetes service depecndency findings
	var kubeSvcList []*v2.Service
	labels := make(map[string]string)
	for key, value := range daemonset.Spec.Template.Labels {
		labels[key] = value
		kubesvclist, err := conn.getKubernetesServices(ctx, key, value, namespace)
		if err != nil {
			return
		}
		for _, kubesvc := range kubesvclist.Items {
			kubeSvcList = append(kubeSvcList, &kubesvc)
		}
	}

	//container dependency finding
	for _, container := range daemonset.Spec.Template.Spec.Containers {

		//resolving dependencies of kubernetes service
		if kubeSvcList != nil && len(kubeSvcList) > 0 {
			for _, kubeSvc := range kubeSvcList {
				if isPortMatched(kubeSvc, &container) {
					k8serviceTemp, err := getCpConvertedTemplate(kubeSvc, kubeSvc.Kind)
					if err != nil {
						utils.Error.Println(err)
						return
					}
					//istio components creation
					istioSvcTemps, err := CreateIstioComponents(k8serviceTemp, labels)
					if err != nil {
						utils.Error.Println(err)
						return
					}
					for _, istioSvc := range istioSvcTemps {
						serviceTemplates = append(serviceTemplates, istioSvc)
					}
					//istio components creation
					if !isAlreadyExist(k8serviceTemp.Namespace, k8serviceTemp.ServiceSubType, k8serviceTemp.Name) {
						k8serviceTemp.BeforeServices = append(k8serviceTemp.BeforeServices, &daemonsetTemp.ServiceId)
						daemonsetTemp.AfterServices = append(daemonsetTemp.AfterServices, &k8serviceTemp.ServiceId)
						serviceTemplates = append(serviceTemplates, k8serviceTemp)
					} else {
						isSameDaemonSet := false
						for _, serviceId := range k8serviceTemp.BeforeServices {
							if *serviceId == daemonsetTemp.ServiceId {
								isSameDaemonSet = true
								break
							}
						}
						if !isSameDaemonSet {
							k8serviceTemp.BeforeServices = append(k8serviceTemp.BeforeServices, &daemonsetTemp.ServiceId)
							daemonsetTemp.AfterServices = append(daemonsetTemp.AfterServices, &k8serviceTemp.ServiceId)
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
					return
				}
				secretTemp, err := getCpConvertedTemplate(secret, secret.Kind)
				if err != nil {
					utils.Error.Println(err)
					return
				}
				if !isAlreadyExist(secretTemp.Namespace, secretTemp.ServiceSubType, secretTemp.Name) {
					daemonsetTemp.BeforeServices = append(daemonsetTemp.BeforeServices, &secretTemp.ServiceId)
					secretTemp.AfterServices = append(secretTemp.AfterServices, &daemonsetTemp.ServiceId)
					serviceTemplates = append(serviceTemplates, secretTemp)
				} else {
					isSameDaemonSet := false
					for _, serviceId := range secretTemp.AfterServices {
						if *serviceId == daemonsetTemp.ServiceId {
							isSameDaemonSet = true
							break
						}
					}
					if !isSameDaemonSet {
						daemonsetTemp.BeforeServices = append(daemonsetTemp.BeforeServices, &secretTemp.ServiceId)
						secretTemp.AfterServices = append(secretTemp.AfterServices, &daemonsetTemp.ServiceId)
					}
				}
			} else if env.ValueFrom != nil && env.ValueFrom.ConfigMapKeyRef != nil {
				configmapname := env.ValueFrom.ConfigMapKeyRef.Name
				configmap, err := conn.getConfigMap(ctx, configmapname, namespace)
				if err != nil {
					return
				}
				configmapTemp, err := getCpConvertedTemplate(configmap, configmap.Kind)
				if err != nil {
					utils.Error.Println(err)
					return
				}
				if !isAlreadyExist(configmapTemp.Namespace, configmapTemp.ServiceSubType, configmapTemp.Name) {
					daemonsetTemp.BeforeServices = append(daemonsetTemp.BeforeServices, &configmapTemp.ServiceId)
					configmapTemp.AfterServices = append(configmapTemp.AfterServices, &daemonsetTemp.ServiceId)
					serviceTemplates = append(serviceTemplates, configmapTemp)
				} else {
					isSameDaemonSet := false
					for _, serviceId := range configmapTemp.AfterServices {
						if *serviceId == daemonsetTemp.ServiceId {
							isSameDaemonSet = true
							break
						}
					}
					if !isSameDaemonSet {
						daemonsetTemp.BeforeServices = append(daemonsetTemp.BeforeServices, &configmapTemp.ServiceId)
						configmapTemp.AfterServices = append(configmapTemp.AfterServices, &daemonsetTemp.ServiceId)
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
				return
			}
			secretTemp, err := getCpConvertedTemplate(secret, secret.Kind)
			if err != nil {
				utils.Error.Println(err)
				return
			}
			if !isAlreadyExist(secretTemp.Namespace, secretTemp.ServiceSubType, secretTemp.Name) {
				daemonsetTemp.BeforeServices = append(daemonsetTemp.BeforeServices, &secretTemp.ServiceId)
				secretTemp.AfterServices = append(secretTemp.AfterServices, &daemonsetTemp.ServiceId)
				serviceTemplates = append(serviceTemplates, secretTemp)
			} else {
				isSameDaemonSet := false
				for _, serviceId := range secretTemp.AfterServices {
					if *serviceId == daemonsetTemp.ServiceId {
						isSameDaemonSet = true
						break
					}
				}
				if !isSameDaemonSet {
					daemonsetTemp.BeforeServices = append(daemonsetTemp.BeforeServices, &secretTemp.ServiceId)
					secretTemp.AfterServices = append(secretTemp.AfterServices, &daemonsetTemp.ServiceId)
				}
			}
		} else if vol.ConfigMap != nil {
			configmapname := vol.ConfigMap.Name
			configmap, err := conn.getConfigMap(ctx, configmapname, namespace)
			if err != nil {
				return
			}
			configmapTemp, err := getCpConvertedTemplate(configmap, configmap.Kind)
			if err != nil {
				utils.Error.Println(err)
				return
			}
			if !isAlreadyExist(configmapTemp.Namespace, configmapTemp.ServiceSubType, configmapTemp.Name) {
				daemonsetTemp.BeforeServices = append(daemonsetTemp.BeforeServices, &configmapTemp.ServiceId)
				configmapTemp.AfterServices = append(configmapTemp.AfterServices, &daemonsetTemp.ServiceId)
				serviceTemplates = append(serviceTemplates, configmapTemp)
			} else {
				isSameDaemonSet := false
				for _, serviceId := range configmapTemp.AfterServices {
					if *serviceId == daemonsetTemp.ServiceId {
						isSameDaemonSet = true
						break
					}
				}
				if !isSameDaemonSet {
					daemonsetTemp.BeforeServices = append(daemonsetTemp.BeforeServices, &configmapTemp.ServiceId)
					configmapTemp.AfterServices = append(configmapTemp.AfterServices, &daemonsetTemp.ServiceId)
				}
			}
		} else if vol.PersistentVolumeClaim != nil {
			pvcname := vol.PersistentVolumeClaim.ClaimName
			pvc, err := conn.getPvc(ctx, pvcname, namespace)
			if err != nil {
				return
			}
			pvcTemp, err := getCpConvertedTemplate(pvc, pvc.Kind)
			if err != nil {
				utils.Error.Println(err)
				return
			}
			if !isAlreadyExist(pvcTemp.Namespace, pvcTemp.ServiceSubType, pvcTemp.Name) {
				daemonsetTemp.BeforeServices = append(daemonsetTemp.BeforeServices, &pvcTemp.ServiceId)
				pvcTemp.AfterServices = append(pvcTemp.AfterServices, &daemonsetTemp.ServiceId)
				serviceTemplates = append(serviceTemplates, pvcTemp)
			} else {
				service := GetExistingService(pvcTemp.Namespace, pvcTemp.ServiceSubType, pvcTemp.Name)
				daemonsetTemp.BeforeServices = append(daemonsetTemp.BeforeServices, &service.ServiceId)
				service.AfterServices = append(service.AfterServices, &daemonsetTemp.ServiceId)
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
				if !isAlreadyExist(storageClassTemp.Namespace, storageClassTemp.ServiceSubType, storageClassTemp.Name) {
					pvcTemp.BeforeServices = append(pvcTemp.BeforeServices, &storageClassTemp.ServiceId)
					storageClassTemp.AfterServices = append(storageClassTemp.AfterServices, &pvcTemp.ServiceId)
					serviceTemplates = append(serviceTemplates, storageClassTemp)
				} else {
					isPVCexist := false
					for _, serviceId := range storageClassTemp.AfterServices {
						if *serviceId == pvcTemp.ServiceId {
							isPVCexist = true
							break
						}
					}
					if !isPVCexist {
						pvcTemp.BeforeServices = append(pvcTemp.BeforeServices, &storageClassTemp.ServiceId)
						storageClassTemp.AfterServices = append(storageClassTemp.AfterServices, &pvcTemp.ServiceId)
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
						return
					}
					configmapTemp, err := getCpConvertedTemplate(configmap, configmap.Kind)
					if err != nil {
						utils.Error.Println(err)
						return
					}
					if !isAlreadyExist(configmapTemp.Namespace, configmapTemp.ServiceSubType, configmapTemp.Name) {
						daemonsetTemp.BeforeServices = append(daemonsetTemp.BeforeServices, &configmapTemp.ServiceId)
						configmapTemp.AfterServices = append(configmapTemp.AfterServices, &daemonsetTemp.ServiceId)
						serviceTemplates = append(serviceTemplates, configmapTemp)
					} else {
						isSameDaemonSet := false
						for _, serviceId := range configmapTemp.AfterServices {
							if *serviceId == daemonsetTemp.ServiceId {
								isSameDaemonSet = true
								break
							}
						}
						if !isSameDaemonSet {
							daemonsetTemp.BeforeServices = append(daemonsetTemp.BeforeServices, &configmapTemp.ServiceId)
							configmapTemp.AfterServices = append(configmapTemp.AfterServices, &daemonsetTemp.ServiceId)
						}
					}
				} else if source.Secret != nil {
					secretname := vol.Secret.SecretName
					secret, err := conn.getSecret(ctx, secretname, namespace)
					if err != nil {
						return
					}
					secretTemp, err := getCpConvertedTemplate(secret, secret.Kind)
					if err != nil {
						utils.Error.Println(err)
						return
					}
					if !isAlreadyExist(secretTemp.Namespace, secretTemp.ServiceSubType, secretTemp.Name) {
						daemonsetTemp.BeforeServices = append(daemonsetTemp.BeforeServices, &secretTemp.ServiceId)
						secretTemp.AfterServices = append(secretTemp.AfterServices, &daemonsetTemp.ServiceId)
						serviceTemplates = append(serviceTemplates, secretTemp)
					} else {
						isSameDaemonSet := false
						for _, serviceId := range secretTemp.AfterServices {
							if *serviceId == daemonsetTemp.ServiceId {
								isSameDaemonSet = true
								break
							}
						}
						if !isSameDaemonSet {
							daemonsetTemp.BeforeServices = append(daemonsetTemp.BeforeServices, &secretTemp.ServiceId)
							secretTemp.AfterServices = append(secretTemp.AfterServices, &daemonsetTemp.ServiceId)
						}
					}
				}
			}
		}
	}

	serviceTemplates = append(serviceTemplates, daemonsetTemp)
}

func (conn *GrpcConn) ResolveStatefulSetDependencies(statefulset v1.StatefulSet, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	utils.Info.Printf("Resolving dependencies of statefulset :%v within namespace : %v", statefulset.Name, statefulset.Namespace)
	stsTemp, err := getCpConvertedTemplate(statefulset, statefulset.Kind)
	if err != nil {
		utils.Error.Println(err)
		return
	}
	namespace := statefulset.Namespace
	if statefulset.Spec.Template.Spec.ServiceAccountName != "" {
		svcname := statefulset.Spec.Template.Spec.ServiceAccountName
		svcaccount, err := conn.getSvcAccount(ctx, svcname, namespace)
		if err != nil {
			return
		}
		svcAccTemp, err := getCpConvertedTemplate(svcaccount, svcaccount.Kind)
		if err != nil {
			utils.Error.Println(err)
			return
		}
		if !isAlreadyExist(svcAccTemp.Namespace, svcAccTemp.ServiceSubType, svcAccTemp.Name) {
			rbacDependencies, err := conn.getK8sRbacResources(ctx, namespace, svcaccount, svcAccTemp)
			if err != nil {
				utils.Error.Println(err)
				return
			}

			for _, rbacTemp := range rbacDependencies {
				if rbacTemp.ServiceSubType == meshConstants.ServiceAccount {
					stsTemp.BeforeServices = append(stsTemp.BeforeServices, &rbacTemp.ServiceId)
					rbacTemp.AfterServices = append(rbacTemp.AfterServices, &stsTemp.ServiceId)
				}
				serviceTemplates = append(serviceTemplates, rbacTemp)
			}
		} else {
			service := GetExistingService(svcAccTemp.Namespace, svcAccTemp.ServiceSubType, svcAccTemp.Name)
			stsTemp.BeforeServices = append(stsTemp.BeforeServices, &service.ServiceId)
			service.AfterServices = append(service.AfterServices, &stsTemp.ServiceId)
		}
	}

	//image pull secrets
	for _, objRef := range statefulset.Spec.Template.Spec.ImagePullSecrets {
		secretname := objRef.Name
		secret, err := conn.getSecret(ctx, secretname, namespace)
		if err != nil {
			return
		}
		secretTemp, err := getCpConvertedTemplate(secret, secret.Kind)
		if err != nil {
			utils.Error.Println(err)
			return
		}
		if !isAlreadyExist(secretTemp.Namespace, secretTemp.ServiceSubType, secretTemp.Name) {
			stsTemp.BeforeServices = append(stsTemp.BeforeServices, &secretTemp.ServiceId)
			secretTemp.AfterServices = append(secretTemp.AfterServices, &stsTemp.ServiceId)
			serviceTemplates = append(serviceTemplates, secretTemp)
		} else {
			isSameStatefulSet := false
			for _, serviceId := range secretTemp.AfterServices {
				if *serviceId == stsTemp.ServiceId {
					isSameStatefulSet = true
					break
				}
			}
			if !isSameStatefulSet {
				stsTemp.BeforeServices = append(stsTemp.BeforeServices, &secretTemp.ServiceId)
				secretTemp.AfterServices = append(secretTemp.AfterServices, &stsTemp.ServiceId)
			}
		}
	}

	hpaList, err := conn.getAllHpas(ctx, namespace)
	if err != nil {
		return
	}

	if len(hpaList.Items) > 0 {
		for _, hpa := range hpaList.Items {
			if hpa.Spec.ScaleTargetRef.APIVersion == statefulset.APIVersion && strings.ToLower(hpa.Spec.ScaleTargetRef.Kind) == strings.ToLower(statefulset.Kind) && hpa.Spec.ScaleTargetRef.Name == statefulset.Name {
				hpaTemplate, err := getCpConvertedTemplate(hpa, hpa.Kind)
				if err != nil {
					utils.Error.Println(err)
					return
				}
				hpaTemplate.BeforeServices = append(hpaTemplate.BeforeServices, &stsTemp.ServiceId)
				stsTemp.AfterServices = append(stsTemp.AfterServices, &hpaTemplate.ServiceId)
				serviceTemplates = append(serviceTemplates, hpaTemplate)
			}
		}
	}

	//kubernetes service depecndency findings
	var kubeSvcList []*v2.Service
	labels := make(map[string]string)
	for key, value := range statefulset.Spec.Template.Labels {
		labels[key] = value
		kubesvclist, err := conn.getKubernetesServices(ctx, key, value, namespace)
		if err != nil {
			return
		}
		for _, kubesvc := range kubesvclist.Items {
			kubeSvcList = append(kubeSvcList, &kubesvc)
		}
	}

	//container dependency finding
	for _, container := range statefulset.Spec.Template.Spec.Containers {

		//resolving dependencies of kubernetes service
		if kubeSvcList != nil && len(kubeSvcList) > 0 {
			for _, kubeSvc := range kubeSvcList {
				if isPortMatched(kubeSvc, &container) {
					k8serviceTemp, err := getCpConvertedTemplate(kubeSvc, kubeSvc.Kind)
					if err != nil {
						utils.Error.Println(err)
						return
					}
					//istio components creation
					istioSvcTemps, err := CreateIstioComponents(k8serviceTemp, labels)
					if err != nil {
						utils.Error.Println(err)
						return
					}
					for _, istioSvc := range istioSvcTemps {
						serviceTemplates = append(serviceTemplates, istioSvc)
					}
					//istio components creation
					if !isAlreadyExist(k8serviceTemp.Namespace, k8serviceTemp.ServiceSubType, k8serviceTemp.Name) {
						k8serviceTemp.BeforeServices = append(k8serviceTemp.BeforeServices, &stsTemp.ServiceId)
						stsTemp.AfterServices = append(stsTemp.AfterServices, &k8serviceTemp.ServiceId)
						serviceTemplates = append(serviceTemplates, k8serviceTemp)
					} else {
						isSameStatefulSet := false
						for _, serviceId := range k8serviceTemp.BeforeServices {
							if *serviceId == stsTemp.ServiceId {
								isSameStatefulSet = true
								break
							}
						}
						if !isSameStatefulSet {
							k8serviceTemp.BeforeServices = append(k8serviceTemp.BeforeServices, &stsTemp.ServiceId)
							stsTemp.AfterServices = append(stsTemp.AfterServices, &k8serviceTemp.ServiceId)
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
					return
				}
				secretTemp, err := getCpConvertedTemplate(secret, secret.Kind)
				if err != nil {
					utils.Error.Println(err)
					return
				}
				if !isAlreadyExist(secretTemp.Namespace, secretTemp.ServiceSubType, secretTemp.Name) {
					stsTemp.BeforeServices = append(stsTemp.BeforeServices, &secretTemp.ServiceId)
					secretTemp.AfterServices = append(secretTemp.AfterServices, &stsTemp.ServiceId)
					serviceTemplates = append(serviceTemplates, secretTemp)
				} else {
					isSameStatefulSet := false
					for _, serviceId := range secretTemp.AfterServices {
						if *serviceId == stsTemp.ServiceId {
							isSameStatefulSet = true
							break
						}
					}
					if !isSameStatefulSet {
						stsTemp.BeforeServices = append(stsTemp.BeforeServices, &secretTemp.ServiceId)
						secretTemp.AfterServices = append(secretTemp.AfterServices, &stsTemp.ServiceId)
					}
				}
			} else if env.ValueFrom != nil && env.ValueFrom.ConfigMapKeyRef != nil {
				configmapname := env.ValueFrom.ConfigMapKeyRef.Name
				configmap, err := conn.getConfigMap(ctx, configmapname, namespace)
				if err != nil {
					return
				}
				configmapTemp, err := getCpConvertedTemplate(configmap, configmap.Kind)
				if err != nil {
					utils.Error.Println(err)
					return
				}
				if !isAlreadyExist(configmapTemp.Namespace, configmapTemp.ServiceSubType, configmapTemp.Name) {
					stsTemp.BeforeServices = append(stsTemp.BeforeServices, &configmapTemp.ServiceId)
					configmapTemp.AfterServices = append(configmapTemp.AfterServices, &stsTemp.ServiceId)
					serviceTemplates = append(serviceTemplates, configmapTemp)
				} else {
					isSameStatefulSet := false
					for _, serviceId := range configmapTemp.AfterServices {
						if *serviceId == stsTemp.ServiceId {
							isSameStatefulSet = true
							break
						}
					}
					if !isSameStatefulSet {
						stsTemp.BeforeServices = append(stsTemp.BeforeServices, &configmapTemp.ServiceId)
						configmapTemp.AfterServices = append(configmapTemp.AfterServices, &stsTemp.ServiceId)
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
				return
			}
			secretTemp, err := getCpConvertedTemplate(secret, secret.Kind)
			if err != nil {
				utils.Error.Println(err)
				return
			}
			if !isAlreadyExist(secretTemp.Namespace, secretTemp.ServiceSubType, secretTemp.Name) {
				stsTemp.BeforeServices = append(stsTemp.BeforeServices, &secretTemp.ServiceId)
				secretTemp.AfterServices = append(secretTemp.AfterServices, &stsTemp.ServiceId)
				serviceTemplates = append(serviceTemplates, secretTemp)
			} else {
				isSameStatefulSet := false
				for _, serviceId := range secretTemp.AfterServices {
					if *serviceId == stsTemp.ServiceId {
						isSameStatefulSet = true
						break
					}
				}
				if !isSameStatefulSet {
					stsTemp.BeforeServices = append(stsTemp.BeforeServices, &secretTemp.ServiceId)
					secretTemp.AfterServices = append(secretTemp.AfterServices, &stsTemp.ServiceId)
				}
			}
		} else if vol.ConfigMap != nil {
			configmapname := vol.ConfigMap.Name
			configmap, err := conn.getConfigMap(ctx, configmapname, namespace)
			if err != nil {
				return
			}
			configmapTemp, err := getCpConvertedTemplate(configmap, configmap.Kind)
			if err != nil {
				utils.Error.Println(err)
				return
			}
			if !isAlreadyExist(configmapTemp.Namespace, configmapTemp.ServiceSubType, configmapTemp.Name) {
				stsTemp.BeforeServices = append(stsTemp.BeforeServices, &configmapTemp.ServiceId)
				configmapTemp.AfterServices = append(configmapTemp.AfterServices, &stsTemp.ServiceId)
				serviceTemplates = append(serviceTemplates, configmapTemp)
			} else {
				isSameStatefulSet := false
				for _, serviceId := range configmapTemp.AfterServices {
					if *serviceId == stsTemp.ServiceId {
						isSameStatefulSet = true
						break
					}
				}
				if !isSameStatefulSet {
					stsTemp.BeforeServices = append(stsTemp.BeforeServices, &configmapTemp.ServiceId)
					configmapTemp.AfterServices = append(configmapTemp.AfterServices, &stsTemp.ServiceId)
				}
			}
		} else if vol.PersistentVolumeClaim != nil {
			pvcname := vol.PersistentVolumeClaim.ClaimName
			pvc, err := conn.getPvc(ctx, pvcname, namespace)
			if err != nil {
				return
			}
			pvcTemp, err := getCpConvertedTemplate(pvc, pvc.Kind)
			if err != nil {
				utils.Error.Println(err)
				return
			}
			if !isAlreadyExist(pvcTemp.Namespace, pvcTemp.ServiceSubType, pvcTemp.Name) {
				stsTemp.BeforeServices = append(stsTemp.BeforeServices, &pvcTemp.ServiceId)
				pvcTemp.AfterServices = append(pvcTemp.AfterServices, &stsTemp.ServiceId)
				serviceTemplates = append(serviceTemplates, pvcTemp)
			} else {
				service := GetExistingService(pvcTemp.Namespace, pvcTemp.ServiceSubType, pvcTemp.Name)
				stsTemp.BeforeServices = append(stsTemp.BeforeServices, &service.ServiceId)
				service.AfterServices = append(service.AfterServices, &stsTemp.ServiceId)
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
				if !isAlreadyExist(storageClassTemp.Namespace, storageClassTemp.ServiceSubType, storageClassTemp.Name) {
					pvcTemp.BeforeServices = append(pvcTemp.BeforeServices, &storageClassTemp.ServiceId)
					storageClassTemp.AfterServices = append(storageClassTemp.AfterServices, &pvcTemp.ServiceId)
					serviceTemplates = append(serviceTemplates, storageClassTemp)
				} else {
					isPVCexist := false
					for _, serviceId := range storageClassTemp.AfterServices {
						if *serviceId == pvcTemp.ServiceId {
							isPVCexist = true
							break
						}
					}
					if !isPVCexist {
						pvcTemp.BeforeServices = append(pvcTemp.BeforeServices, &storageClassTemp.ServiceId)
						storageClassTemp.AfterServices = append(storageClassTemp.AfterServices, &pvcTemp.ServiceId)
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
						return
					}
					configmapTemp, err := getCpConvertedTemplate(configmap, configmap.Kind)
					if err != nil {
						utils.Error.Println(err)
						return
					}
					if !isAlreadyExist(configmapTemp.Namespace, configmapTemp.ServiceSubType, configmapTemp.Name) {
						stsTemp.BeforeServices = append(stsTemp.BeforeServices, &configmapTemp.ServiceId)
						configmapTemp.AfterServices = append(configmapTemp.AfterServices, &stsTemp.ServiceId)
						serviceTemplates = append(serviceTemplates, configmapTemp)
					} else {
						isSameStatefulSet := false
						for _, serviceId := range configmapTemp.AfterServices {
							if *serviceId == stsTemp.ServiceId {
								isSameStatefulSet = true
								break
							}
						}
						if !isSameStatefulSet {
							stsTemp.BeforeServices = append(stsTemp.BeforeServices, &configmapTemp.ServiceId)
							configmapTemp.AfterServices = append(configmapTemp.AfterServices, &stsTemp.ServiceId)
						}
					}
				} else if source.Secret != nil {
					secretname := vol.Secret.SecretName
					secret, err := conn.getSecret(ctx, secretname, namespace)
					if err != nil {
						return
					}
					secretTemp, err := getCpConvertedTemplate(secret, secret.Kind)
					if err != nil {
						utils.Error.Println(err)
						return
					}
					if !isAlreadyExist(secretTemp.Namespace, secretTemp.ServiceSubType, secretTemp.Name) {
						stsTemp.BeforeServices = append(stsTemp.BeforeServices, &secretTemp.ServiceId)
						secretTemp.AfterServices = append(secretTemp.AfterServices, &stsTemp.ServiceId)
						serviceTemplates = append(serviceTemplates, secretTemp)
					} else {
						isSameStatefulSet := false
						for _, serviceId := range secretTemp.AfterServices {
							if *serviceId == stsTemp.ServiceId {
								isSameStatefulSet = true
								break
							}
						}
						if !isSameStatefulSet {
							stsTemp.BeforeServices = append(stsTemp.BeforeServices, &secretTemp.ServiceId)
							secretTemp.AfterServices = append(secretTemp.AfterServices, &stsTemp.ServiceId)
						}
					}
				}
			}
		}
	}

	serviceTemplates = append(serviceTemplates, stsTemp)
}

func (conn *GrpcConn) statefulsetsK8sToCp(ctx context.Context, statefulsets []v1.StatefulSet, wg *sync.WaitGroup) error {

	for _, statefulset := range statefulsets {
		wg.Add(1)
		go conn.ResolveStatefulSetDependencies(statefulset, wg, ctx)
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
		if !isAlreadyExist(svcAccTemp.Namespace, svcAccTemp.ServiceSubType, svcAccTemp.Name) {
			rbacDependencies, err := conn.getK8sRbacResources(ctx, namespace, svcaccount, svcAccTemp)
			if err != nil {
				utils.Error.Println(err)
				return
			}

			var strpointer = new(string)
			*strpointer = "serviceaccount"
			for _, rbacTemp := range rbacDependencies {
				if rbacTemp.ServiceSubType == meshConstants.ServiceAccount {
					depTemp.BeforeServices = append(depTemp.BeforeServices, &rbacTemp.ServiceId)
					rbacTemp.AfterServices = append(rbacTemp.AfterServices, &depTemp.ServiceId)
				}
				serviceTemplates = append(serviceTemplates, rbacTemp)
			}
		} else {
			service := GetExistingService(svcAccTemp.Namespace, svcAccTemp.ServiceSubType, svcAccTemp.Name)
			depTemp.BeforeServices = append(depTemp.BeforeServices, &service.ServiceId)
			service.AfterServices = append(service.AfterServices, &depTemp.ServiceId)
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
		if !isAlreadyExist(secretTemp.Namespace, secretTemp.ServiceSubType, secretTemp.Name) {
			depTemp.BeforeServices = append(depTemp.BeforeServices, &secretTemp.ServiceId)
			secretTemp.AfterServices = append(secretTemp.AfterServices, &depTemp.ServiceId)
			serviceTemplates = append(serviceTemplates, secretTemp)
		} else {
			isSameDeployment := false
			for _, serviceId := range secretTemp.AfterServices {
				if *serviceId == depTemp.ServiceId {
					isSameDeployment = true
					break
				}
			}
			if !isSameDeployment {
				depTemp.BeforeServices = append(depTemp.BeforeServices, &secretTemp.ServiceId)
				secretTemp.AfterServices = append(secretTemp.AfterServices, &depTemp.ServiceId)
			}
		}
	}

	hpaList, err := conn.getAllHpas(ctx, namespace)
	if err != nil {
		return
	}
	if len(hpaList.Items) > 0 {
		for _, hpa := range hpaList.Items {
			if hpa.Spec.ScaleTargetRef.APIVersion == dep.APIVersion && strings.ToLower(hpa.Spec.ScaleTargetRef.Kind) == strings.ToLower(dep.Kind) && hpa.Spec.ScaleTargetRef.Name == dep.Name {
				hpaTemplate, err := getCpConvertedTemplate(hpa, hpa.Kind)
				if err != nil {
					utils.Error.Println(err)
					return
				}
				hpaTemplate.BeforeServices = append(hpaTemplate.BeforeServices, &depTemp.ServiceId)
				depTemp.AfterServices = append(depTemp.AfterServices, &hpaTemplate.ServiceId)
				serviceTemplates = append(serviceTemplates, hpaTemplate)
			}
		}
	}

	//finding kubernetes service
	var kubeSvcList []*v2.Service
	labels := make(map[string]string)
	for key, value := range dep.Spec.Template.Labels {
		labels[key] = value
		kubesvclist, err := conn.getKubernetesServices(ctx, key, value, namespace)
		if err != nil {
			utils.Error.Println(err)
			return
		}
		for _, kubesvc := range kubesvclist.Items {
			kubeSvcList = append(kubeSvcList, &kubesvc)
		}
	}

	//container dependency finding
	for _, container := range dep.Spec.Template.Spec.Containers {

		//resolving dependencies of kubernetes service
		if kubeSvcList != nil && len(kubeSvcList) > 0 {
			for _, kubeSvc := range kubeSvcList {
				if isPortMatched(kubeSvc, &container) {
					k8serviceTemp, err := getCpConvertedTemplate(kubeSvc, kubeSvc.Kind)
					if err != nil {
						utils.Error.Println(err)
						return
					}

					if conn.isIstioEnabled(ctx) {
						//Istio components discovery
						err = conn.discoverIstioComponents(ctx, k8serviceTemp, namespace)
						if err != nil {
							utils.Error.Println(err)
							return
						}
					} else {
						//istio components creation
						istioSvcTemps, err := CreateIstioComponents(k8serviceTemp, labels)
						if err != nil {
							utils.Error.Println(err)
							return
						}
						for _, istioSvc := range istioSvcTemps {
							serviceTemplates = append(serviceTemplates, istioSvc)
						}
						//istio components creation
					}

					if !isAlreadyExist(k8serviceTemp.Namespace, k8serviceTemp.ServiceSubType, k8serviceTemp.Name) {
						k8serviceTemp.BeforeServices = append(k8serviceTemp.BeforeServices, &depTemp.ServiceId)
						depTemp.AfterServices = append(depTemp.AfterServices, &k8serviceTemp.ServiceId)
						serviceTemplates = append(serviceTemplates, k8serviceTemp)
					} else {
						isSameDeployment := false
						for _, serviceId := range k8serviceTemp.BeforeServices {
							if *serviceId == depTemp.ServiceId {
								isSameDeployment = true
								break
							}
						}
						if !isSameDeployment {
							k8serviceTemp.BeforeServices = append(k8serviceTemp.BeforeServices, &depTemp.ServiceId)
							depTemp.AfterServices = append(depTemp.AfterServices, &k8serviceTemp.ServiceId)
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
				if !isAlreadyExist(secretTemp.Namespace, secretTemp.ServiceSubType, secretTemp.Name) {
					depTemp.BeforeServices = append(depTemp.BeforeServices, &secretTemp.ServiceId)
					secretTemp.AfterServices = append(secretTemp.AfterServices, &depTemp.ServiceId)
					serviceTemplates = append(serviceTemplates, secretTemp)
				} else {
					isSameDeployment := false
					for _, serviceId := range secretTemp.AfterServices {
						if *serviceId == depTemp.ServiceId {
							isSameDeployment = true
							break
						}
					}
					if !isSameDeployment {
						depTemp.BeforeServices = append(depTemp.BeforeServices, &secretTemp.ServiceId)
						secretTemp.AfterServices = append(secretTemp.AfterServices, &depTemp.ServiceId)
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
				if !isAlreadyExist(configmapTemp.Namespace, configmapTemp.ServiceSubType, configmapTemp.Name) {
					depTemp.BeforeServices = append(depTemp.BeforeServices, &configmapTemp.ServiceId)
					configmapTemp.AfterServices = append(configmapTemp.AfterServices, &depTemp.ServiceId)
					serviceTemplates = append(serviceTemplates, configmapTemp)
				} else {
					isSameDeployment := false
					for _, serviceId := range configmapTemp.AfterServices {
						if *serviceId == depTemp.ServiceId {
							isSameDeployment = true
							break
						}
					}
					if !isSameDeployment {
						depTemp.BeforeServices = append(depTemp.BeforeServices, &configmapTemp.ServiceId)
						configmapTemp.AfterServices = append(configmapTemp.AfterServices, &depTemp.ServiceId)
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
			if !isAlreadyExist(secretTemp.Namespace, secretTemp.ServiceSubType, secretTemp.Name) {
				depTemp.BeforeServices = append(depTemp.BeforeServices, &secretTemp.ServiceId)
				secretTemp.AfterServices = append(secretTemp.AfterServices, &depTemp.ServiceId)
				serviceTemplates = append(serviceTemplates, secretTemp)
			} else {
				isSameDeployment := false
				for _, serviceId := range secretTemp.AfterServices {
					if *serviceId == depTemp.ServiceId {
						isSameDeployment = true
						break
					}
				}
				if !isSameDeployment {
					depTemp.BeforeServices = append(depTemp.BeforeServices, &secretTemp.ServiceId)
					secretTemp.AfterServices = append(secretTemp.AfterServices, &depTemp.ServiceId)
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
			if !isAlreadyExist(configmapTemp.Namespace, configmapTemp.ServiceSubType, configmapTemp.Name) {
				depTemp.BeforeServices = append(depTemp.BeforeServices, &configmapTemp.ServiceId)
				configmapTemp.AfterServices = append(configmapTemp.AfterServices, &depTemp.ServiceId)
				serviceTemplates = append(serviceTemplates, configmapTemp)
			} else {
				isSameDeployment := false
				for _, serviceId := range configmapTemp.AfterServices {
					if *serviceId == depTemp.ServiceId {
						isSameDeployment = true
						break
					}
				}
				if !isSameDeployment {
					depTemp.BeforeServices = append(depTemp.BeforeServices, &configmapTemp.ServiceId)
					configmapTemp.AfterServices = append(configmapTemp.AfterServices, &depTemp.ServiceId)
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
			if !isAlreadyExist(pvcTemp.Namespace, pvcTemp.ServiceSubType, pvcTemp.Name) {
				depTemp.BeforeServices = append(depTemp.BeforeServices, &pvcTemp.ServiceId)
				pvcTemp.AfterServices = append(pvcTemp.AfterServices, &depTemp.ServiceId)
				serviceTemplates = append(serviceTemplates, pvcTemp)
			} else {
				service := GetExistingService(pvcTemp.Namespace, pvcTemp.ServiceSubType, pvcTemp.Name)
				depTemp.BeforeServices = append(depTemp.BeforeServices, &service.ServiceId)
				service.AfterServices = append(service.AfterServices, &depTemp.ServiceId)
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
				if !isAlreadyExist(storageClassTemp.Namespace, storageClassTemp.ServiceSubType, storageClassTemp.Name) {
					pvcTemp.BeforeServices = append(pvcTemp.BeforeServices, &storageClassTemp.ServiceId)
					storageClassTemp.AfterServices = append(storageClassTemp.AfterServices, &pvcTemp.ServiceId)
					serviceTemplates = append(serviceTemplates, storageClassTemp)
				} else {
					isPVCexist := false
					for _, serviceId := range storageClassTemp.AfterServices {
						if *serviceId == pvcTemp.ServiceId {
							isPVCexist = true
							break
						}
					}
					if !isPVCexist {
						pvcTemp.BeforeServices = append(pvcTemp.BeforeServices, &storageClassTemp.ServiceId)
						storageClassTemp.AfterServices = append(storageClassTemp.AfterServices, &pvcTemp.ServiceId)
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
					if !isAlreadyExist(configmapTemp.Namespace, configmapTemp.ServiceSubType, configmapTemp.Name) {
						depTemp.BeforeServices = append(depTemp.BeforeServices, &configmapTemp.ServiceId)
						configmapTemp.AfterServices = append(configmapTemp.AfterServices, &depTemp.ServiceId)
						serviceTemplates = append(serviceTemplates, configmapTemp)
					} else {
						isSameDeployment := false
						for _, serviceId := range configmapTemp.AfterServices {
							if *serviceId == depTemp.ServiceId {
								isSameDeployment = true
								break
							}
						}
						if !isSameDeployment {
							depTemp.BeforeServices = append(depTemp.BeforeServices, &configmapTemp.ServiceId)
							configmapTemp.AfterServices = append(configmapTemp.AfterServices, &depTemp.ServiceId)
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
					if !isAlreadyExist(secretTemp.Namespace, secretTemp.ServiceSubType, secretTemp.Name) {
						depTemp.BeforeServices = append(depTemp.BeforeServices, &secretTemp.ServiceId)
						secretTemp.AfterServices = append(secretTemp.AfterServices, &depTemp.ServiceId)
						serviceTemplates = append(serviceTemplates, secretTemp)
					} else {
						isSameDeployment := false
						for _, serviceId := range secretTemp.AfterServices {
							if *serviceId == depTemp.ServiceId {
								isSameDeployment = true
								break
							}
						}
						if !isSameDeployment {
							depTemp.BeforeServices = append(depTemp.BeforeServices, &secretTemp.ServiceId)
							secretTemp.AfterServices = append(secretTemp.AfterServices, &depTemp.ServiceId)
						}
					}
				}
			}
		}
	}

	serviceTemplates = append(serviceTemplates, depTemp)
}

func (conn *GrpcConn) discoverIstioComponents(ctx context.Context, svcTemp *svcTypes.ServiceTemplate, namespace string) error {
	err := conn.discoverIstioDestinationRules(ctx, svcTemp, namespace)
	if err != nil {
		return err
	}

	err = conn.discoverIstioVirtualServices(ctx, svcTemp, namespace)
	if err != nil {
		return err
	}

	err = conn.discoverIstioServiceEntries(ctx, namespace)
	if err != nil {
		return err
	}

	return nil
}

func (conn *GrpcConn) discoverIstioServiceEntries(ctx context.Context, namespace string) error {
	svcEntryList, err := conn.getAllServiceEntries(ctx, namespace)
	if err != nil {
		return err
	}
	for _, svcEntry := range svcEntryList.Items {
		svcEntryTemp, err := getCpConvertedTemplate(svcEntry, svcEntry.Kind)
		if err != nil {
			return err
		}
		serviceTemplates = append(serviceTemplates, svcEntryTemp)
	}
}

func (conn *GrpcConn) discoverIstioDestinationRules(ctx context.Context, svcTemp *svcTypes.ServiceTemplate, namespace string) error {

	drList, err := conn.getAllDestinationRules(ctx, namespace)
	if err != nil {
		return err
	}
	for _, dr := range drList.Items {
		if dr.Spec.Host == svcTemp.Name {
			drTemp, err := getCpConvertedTemplate(dr, dr.Kind)
			if err != nil {
				return err
			}
			drTemp.BeforeServices = append(drTemp.BeforeServices, &svcTemp.ServiceId)
			svcTemp.AfterServices = append(svcTemp.AfterServices, &drTemp.ServiceId)
			serviceTemplates = append(serviceTemplates, drTemp)
			break
		}
	}

	return nil
}

func (conn *GrpcConn) discoverIstioVirtualServices(ctx context.Context, svcTemp *svcTypes.ServiceTemplate, namespace string) error {

	vsList, err := conn.getAllVirtualServices(ctx, namespace)
	if err != nil {
		return err
	}
	for _, vs := range vsList.Items {
		vsTemp, err := getCpConvertedTemplate(vs, vs.Kind)
		if err != nil {
			return err
		}
		for _, http := range vs.Spec.Http {
			for _, route := range http.Route {
				if !isAlreadyExist(vsTemp.Namespace, vsTemp.ServiceSubType, vsTemp.Name) && route.Destination.Host == svcTemp.Name {
					vsTemp.BeforeServices = append(vsTemp.BeforeServices, &svcTemp.ServiceId)
					svcTemp.AfterServices = append(svcTemp.AfterServices, &vsTemp.ServiceId)
					serviceTemplates = append(serviceTemplates, vsTemp)
					break
				}
			}
		}

		//istio gateway discovery
		for _, gateway := range vs.Spec.Gateways {
			istioGateway, err := conn.getIstioGateway(ctx, gateway, namespace)
			if err != nil {
				return err
			}

			gatewayTemp, err := getCpConvertedTemplate(istioGateway, istioGateway.Kind)
			if err != nil {
				return err
			}

			gatewayTemp.BeforeServices = append(gatewayTemp.BeforeServices, &vsTemp.ServiceId)
			vsTemp.AfterServices = append(vsTemp.AfterServices, &gatewayTemp.ServiceId)
			serviceTemplates = append(serviceTemplates, gatewayTemp)
		}
	}

	return nil
}

func (conn *GrpcConn) getAllServiceEntries(ctx context.Context, namespace string) (*istioClient.ServiceEntryList, error) {
	response, err := pb.NewK8SResourceClient(conn.Connection).GetK8SResource(ctx, &pb.K8SResourceRequest{
		ProjectId: conn.ProjectId,
		CompanyId: conn.CompanyId,
		Token:     conn.token,
		Command:   "kubectl",
		Args:      []string{"get", "serviceentry", "-n", namespace, "-o", "json"},
	})
	if err != nil {
		utils.Error.Println(err)
		return nil, err
	}

	var serviceEntryList *istioClient.ServiceEntryList
	err = json.Unmarshal(response.Resource, &serviceEntryList)
	if err != nil {
		utils.Error.Println(err)
		return nil, err
	}

	return serviceEntryList, nil
}

func (conn *GrpcConn) getIstioGateway(ctx context.Context, gatewayName, namespace string) (*istioClient.Gateway, error) {
	response, err := pb.NewK8SResourceClient(conn.Connection).GetK8SResource(ctx, &pb.K8SResourceRequest{
		ProjectId: conn.ProjectId,
		CompanyId: conn.CompanyId,
		Token:     conn.token,
		Command:   "kubectl",
		Args:      []string{"get", "gateway", gatewayName, "-n", namespace, "-o", "json"},
	})
	if err != nil {
		utils.Error.Println(err)
		return nil, err
	}

	var istioGateway *istioClient.Gateway
	err = json.Unmarshal(response.Resource, &istioGateway)
	if err != nil {
		utils.Error.Println(err)
		return nil, err
	}

	return istioGateway, nil
}

func (conn *GrpcConn) getAllDestinationRules(ctx context.Context, namespace string) (*istioClient.DestinationRuleList, error) {
	response, err := pb.NewK8SResourceClient(conn.Connection).GetK8SResource(ctx, &pb.K8SResourceRequest{
		ProjectId: conn.ProjectId,
		CompanyId: conn.CompanyId,
		Token:     conn.token,
		Command:   "kubectl",
		Args:      []string{"get", "dr", "-n", namespace, "-o", "json"},
	})
	if err != nil {
		utils.Error.Println(err)
		return nil, err
	}

	var destinationRuleList *istioClient.DestinationRuleList
	err = json.Unmarshal(response.Resource, &destinationRuleList)
	if err != nil {
		utils.Error.Println(err)
		return nil, err
	}

	return destinationRuleList, nil
}

func (conn *GrpcConn) getAllVirtualServices(ctx context.Context, namespace string) (*istioClient.VirtualServiceList, error) {
	response, err := pb.NewK8SResourceClient(conn.Connection).GetK8SResource(ctx, &pb.K8SResourceRequest{
		ProjectId: conn.ProjectId,
		CompanyId: conn.CompanyId,
		Token:     conn.token,
		Command:   "kubectl",
		Args:      []string{"get", "vs", "-n", namespace, "-o", "json"},
	})
	if err != nil {
		utils.Error.Println(err)
		return nil, err
	}

	var virtualServiceList *istioClient.VirtualServiceList
	err = json.Unmarshal(response.Resource, &virtualServiceList)
	if err != nil {
		utils.Error.Println(err)
		return nil, err
	}

	return virtualServiceList, nil
}

func (conn *GrpcConn) isIstioEnabled(ctx context.Context) bool {
	_, err := pb.NewK8SResourceClient(conn.Connection).GetK8SResource(ctx, &pb.K8SResourceRequest{
		ProjectId: conn.ProjectId,
		CompanyId: conn.CompanyId,
		Token:     conn.token,
		Command:   "kubectl",
		Args:      []string{"get", "ns", "istio-system"},
	})
	if err != nil {
		utils.Error.Println(err)
		return false
	}

	return true
}

func GetExistingService(namespace string, svcsubtype meshConstants.ServiceSubType, name string) *svcTypes.ServiceTemplate {
	for _, service := range serviceTemplates {
		if service.ServiceSubType == svcsubtype && service.Name == name {
			return service
		} else if service.Namespace == namespace && service.ServiceSubType == svcsubtype && service.Name == name {
			return service
		}
	}
	return nil
}

func (conn *GrpcConn) getK8sRbacResources(ctx context.Context, namespace string, k8svcAcc *api.ServiceAccount, svcAccTemp *svcTypes.ServiceTemplate) ([]*svcTypes.ServiceTemplate, error) {

	var rbacServiceTemplates []*svcTypes.ServiceTemplate
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
	//			svcAccTemp.AfterServices = append(svcAccTemp.AfterServices, &secretTemp.ServiceId)
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
			if sub.Kind == constants.ServiceAccount.String() && sub.Name == k8svcAcc.Name {
				if clstrrolebind.RoleRef.Kind == constants.ClusterRole.String() {
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
							if !isAlreadyExist(clstrrolebindTemp.Namespace, clstrrolebindTemp.ServiceSubType, clstrrolebindTemp.Name) {
								clstrrolebindTemp.BeforeServices = append(clstrrolebindTemp.BeforeServices, &clstrroleTemp.ServiceId)
								clstrroleTemp.AfterServices = append(clstrroleTemp.AfterServices, &clstrrolebindTemp.ServiceId)

								clstrrolebindTemp.BeforeServices = append(clstrrolebindTemp.BeforeServices, &svcAccTemp.ServiceId)
								svcAccTemp.AfterServices = append(svcAccTemp.AfterServices, &clstrrolebindTemp.ServiceId)

								rbacServiceTemplates = append(rbacServiceTemplates, clstrrolebindTemp)
								rbacServiceTemplates = append(rbacServiceTemplates, clstrroleTemp)
							} else {
								clstrrolebindTemp.BeforeServices = append(clstrrolebindTemp.BeforeServices, &svcAccTemp.ServiceId)
								svcAccTemp.AfterServices = append(svcAccTemp.AfterServices, &clstrrolebindTemp.ServiceId)
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
			if sub.Kind == constants.ServiceAccount.String() && sub.Name == k8svcAcc.Name {
				if rolebinding.RoleRef.Kind == constants.Role.String() {
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
							if !isAlreadyExist(rolebindTemp.Namespace, rolebindTemp.ServiceSubType, rolebindTemp.Name) {
								rolebindTemp.BeforeServices = append(rolebindTemp.BeforeServices, &roleTemp.ServiceId)
								roleTemp.AfterServices = append(roleTemp.AfterServices, &rolebindTemp.ServiceId)

								rolebindTemp.BeforeServices = append(rolebindTemp.BeforeServices, &svcAccTemp.ServiceId)
								svcAccTemp.AfterServices = append(svcAccTemp.AfterServices, &rolebindTemp.ServiceId)

								rbacServiceTemplates = append(rbacServiceTemplates, rolebindTemp)
								rbacServiceTemplates = append(rbacServiceTemplates, roleTemp)
							} else {
								rolebindTemp.BeforeServices = append(rolebindTemp.BeforeServices, &svcAccTemp.ServiceId)
								svcAccTemp.AfterServices = append(svcAccTemp.AfterServices, &rolebindTemp.ServiceId)
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
		Args:      []string{"get", "daemonsets", "-n", namespace, "-o", "json"},
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

func isAlreadyExist(namespace string, svcsubtype meshConstants.ServiceSubType, name string) bool {
	//utils.Info.Printf("Checking existence of type :%v name :%v within namespace :%v", *svcsubtype, *name, *namespace)
	for _, val := range serviceTemplates {

		if val.ServiceSubType == svcsubtype && val.Name == name {
			return true
		} else if val.Namespace == namespace && val.ServiceSubType == svcsubtype && val.Name == name {
			return true
		}

	}
	return false
}

func getCpConvertedTemplate(data interface{}, kind string) (*svcTypes.ServiceTemplate, error) {

	var template *svcTypes.ServiceTemplate
	switch constants.K8sKind(kind) {
	case constants.Deployment:
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
		template.ServiceId = id
	//case constants.CronJob:
	//	bytes, err := json.Marshal(data)
	//	if err != nil {
	//		utils.Error.Println(err)
	//		return nil, err
	//	}
	//	var cronjob batchv1.CronJob
	//	err = json.Unmarshal(bytes, &cronjob)
	//	if err != nil {
	//		utils.Error.Println(err)
	//		return nil, err
	//	}
	//	CpCronJob, err := convertToCPCronJob(&cronjob)
	//	if err != nil {
	//		utils.Error.Println(err)
	//		return nil, err
	//	}
	//
	//	bytes, err = json.Marshal(CpCronJob)
	//	if err != nil {
	//		utils.Error.Println(err)
	//		return nil, err
	//	}
	//	err = json.Unmarshal(bytes, &template)
	//	if err != nil {
	//		utils.Error.Println(err)
	//		return nil, err
	//	}
	//	id := strconv.Itoa(rand.Int())
	//	template.ServiceId = id
	case constants.Job:
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
		template.ServiceId = id
	case constants.DaemonSet:
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
		template.ServiceId = id
	case constants.StatefulSet:
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
		template.ServiceId = id
	case constants.Service:
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
		template.ServiceId = id
		if isAlreadyExist(template.Namespace, template.ServiceSubType, template.Name) {
			template = GetExistingService(template.Namespace, template.ServiceSubType, template.Name)
		}
	//case constants.HPA:
	//	bytes, err := json.Marshal(data)
	//	if err != nil {
	//		utils.Error.Println(err)
	//		return nil, err
	//	}
	//	var hpa autoscale.HorizontalPodAutoscaler
	//	err = json.Unmarshal(bytes, &hpa)
	//	if err != nil {
	//		utils.Error.Println(err)
	//		return nil, err
	//	}
	//	CpHpa, err := ConvertToCPHPA(&hpa)
	//	if err != nil {
	//		utils.Error.Println(err)
	//		return nil, err
	//	}
	//	bytes, err = json.Marshal(CpHpa)
	//	if err != nil {
	//		utils.Error.Println(err)
	//		return nil, err
	//	}
	//	err = json.Unmarshal(bytes, &template)
	//	if err != nil {
	//		utils.Error.Println(err)
	//		return nil, err
	//	}
	//	id := strconv.Itoa(rand.Int())
	//	template.ServiceId = id
	case constants.ConfigMap:
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
		template.ServiceId = id
		if isAlreadyExist(template.Namespace, template.ServiceSubType, template.Name) {
			template = GetExistingService(template.Namespace, template.ServiceSubType, template.Name)
		}
	case constants.Secret:
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
		template.ServiceId = id
		if isAlreadyExist(template.Namespace, template.ServiceSubType, template.Name) {
			template = GetExistingService(template.Namespace, template.ServiceSubType, template.Name)
		}
	case constants.ServiceAccount:
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
		template.ServiceId = id
		if isAlreadyExist(template.Namespace, template.ServiceSubType, template.Name) {
			template = GetExistingService(template.Namespace, template.ServiceSubType, template.Name)
		}
	case constants.Role:
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
		template.ServiceId = id
		if isAlreadyExist(template.Namespace, template.ServiceSubType, template.Name) {
			template = GetExistingService(template.Namespace, template.ServiceSubType, template.Name)
		}
	case constants.RoleBinding:
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
		template.ServiceId = id
		if isAlreadyExist(template.Namespace, template.ServiceSubType, template.Name) {
			template = GetExistingService(template.Namespace, template.ServiceSubType, template.Name)
		}
	case constants.ClusterRole:
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
		template.ServiceId = id
		if isAlreadyExist(template.Namespace, template.ServiceSubType, template.Name) {
			template = GetExistingService(template.Namespace, template.ServiceSubType, template.Name)
		}
	case constants.ClusterRoleBinding:
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
		template.ServiceId = id
		if isAlreadyExist(template.Namespace, template.ServiceSubType, template.Name) {
			template = GetExistingService(template.Namespace, template.ServiceSubType, template.Name)
		}
	case constants.PersistentVolume:
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
		template.ServiceId = id
	case constants.PersistentVolumeClaim:
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
		template.ServiceId = id
		if isAlreadyExist(template.Namespace, template.ServiceSubType, template.Name) {
			template = GetExistingService(template.Namespace, template.ServiceSubType, template.Name)
		}
	case constants.StorageClass:
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
		template.ServiceId = id
		if isAlreadyExist(template.Namespace, template.ServiceSubType, template.Name) {
			template = GetExistingService(template.Namespace, template.ServiceSubType, template.Name)
		}
	default:
		utils.Info.Println("Kind does not exist in defined switch cases")
		return nil, errors.New("type does not exit")
	}

	switch constants.MeshKind(kind) {
	case constants.DestinationRule:
		bytes, err := json.Marshal(data)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		var dr istioClient.DestinationRule
		err = json.Unmarshal(bytes, &dr)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		CpDr, err := convertToCPDestinationRule(&dr)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		bytes, err = json.Marshal(CpDr)
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
		template.ServiceId = id
		if isAlreadyExist(template.Namespace, template.ServiceSubType, template.Name) {
			template = GetExistingService(template.Namespace, template.ServiceSubType, template.Name)
		}
	case constants.VirtualService:
		bytes, err := json.Marshal(data)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		var vs istioClient.VirtualService
		err = json.Unmarshal(bytes, &vs)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		CpVs, err := convertToCPVirtualService(&vs)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		bytes, err = json.Marshal(CpVs)
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
		template.ServiceId = id
		if isAlreadyExist(template.Namespace, template.ServiceSubType, template.Name) {
			template = GetExistingService(template.Namespace, template.ServiceSubType, template.Name)
		}
	case constants.Gateway:
		bytes, err := json.Marshal(data)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		var gateway istioClient.Gateway
		err = json.Unmarshal(bytes, &gateway)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		CpGateway, err := convertToCPGateway(&gateway)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		bytes, err = json.Marshal(CpGateway)
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
		template.ServiceId = id
		if isAlreadyExist(template.Namespace, template.ServiceSubType, template.Name) {
			template = GetExistingService(template.Namespace, template.ServiceSubType, template.Name)
		}
	case constants.ServiceEntry:
		bytes, err := json.Marshal(data)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		var svcEntry istioClient.ServiceEntry
		err = json.Unmarshal(bytes, &svcEntry)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		CpSvcEntry, err := convertToCPServiceEntry(&svcEntry)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		bytes, err = json.Marshal(CpSvcEntry)
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
		template.ServiceId = id
		if isAlreadyExist(template.Namespace, template.ServiceSubType, template.Name) {
			template = GetExistingService(template.Namespace, template.ServiceSubType, template.Name)
		}
	default:
		utils.Info.Println("Kind does not exist in defined switch cases")
		return nil, errors.New("type does not exit")
	}

	return template, nil

}

func CreateIstioComponents(svcTemp *svcTypes.ServiceTemplate, labels map[string]string) ([]*svcTypes.ServiceTemplate, error) {

	var svcComponents []*svcTypes.ServiceTemplate

	cpKubeService := new(meshTypes.Service)
	cpKubeService.Name = svcTemp.Name
	cpKubeService.ServiceId = svcTemp.ServiceId
	cpKubeService.ServiceType = svcTemp.ServiceType
	cpKubeService.ServiceSubType = svcTemp.ServiceSubType
	cpKubeService.Version = svcTemp.Version
	if svcTemp.Namespace != "" {
		cpKubeService.Namespace = svcTemp.Namespace
	} else {
		cpKubeService.Namespace = "default"
	}
	bytes, err := json.Marshal(svcTemp.ServiceAttributes)
	if err != nil {
		utils.Error.Println(err)
		return nil, err
	}
	err = json.Unmarshal(bytes, &cpKubeService.ServiceAttributes)
	if err != nil {
		utils.Error.Println(err)
		return nil, err
	}

	if !isAlreadyExist(svcTemp.Namespace, meshConstants.VirtualService, svcTemp.Name) {
		destRule := new(meshTypes.DestinationRules)
		istioVS := new(meshTypes.VirtualService)
		istioVS.Name = cpKubeService.Name
		istioVS.Namespace = cpKubeService.Namespace
		istioVS.Version = cpKubeService.Version
		istioVS.ServiceType = meshConstants.MeshType
		istioVS.ServiceSubType = meshConstants.VirtualService
		istioVS.ServiceAttributes = new(meshTypes.VSServiceAttribute)
		for _, value := range cpKubeService.ServiceAttributes.Selector {
			istioVS.ServiceAttributes.Hosts = []string{value}
			for _, label := range labels {
				if label == value {
					continue
				}
				http := new(meshTypes.Http)
				httpRoute := new(meshTypes.HttpRoute)
				routeRule := new(meshTypes.RouteDestination)
				routeRule.Host = value
				routeRule.Subset = label
				routeRule.Port = cpKubeService.ServiceAttributes.Ports[0].Port
				httpRoute.Routes = append(httpRoute.Routes, routeRule)
				http.HttpRoute = append(http.HttpRoute, httpRoute)
				istioVS.ServiceAttributes.Http = append(istioVS.ServiceAttributes.Http, http)

			}
		}

		destRule.ServiceType = meshConstants.MeshType
		destRule.ServiceSubType = meshConstants.DestinationRule
		destRule.Name = cpKubeService.Name
		destRule.Namespace = cpKubeService.Namespace
		for _, value := range cpKubeService.ServiceAttributes.Selector {
			destRule.ServiceAttributes.Host = value
			for key, label := range labels {
				subset := new(meshTypes.Subset)
				if label == value {
					continue
				} else {
					subset.Name = label
					lab := make(map[string]string)
					lab[key] = label
					subset.Labels = &lab
					destRule.ServiceAttributes.Subsets = append(destRule.ServiceAttributes.Subsets, subset)
				}
			}
		}

		var VStemplate *svcTypes.ServiceTemplate
		bytes, err = json.Marshal(istioVS)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		err = json.Unmarshal(bytes, &VStemplate)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		id := strconv.Itoa(rand.Int())
		VStemplate.ServiceId = id
		svcTemp.AfterServices = append(svcTemp.AfterServices, &VStemplate.ServiceId)
		VStemplate.BeforeServices = append(VStemplate.BeforeServices, &svcTemp.ServiceId)

		var DStemplate *svcTypes.ServiceTemplate
		bytes, err = json.Marshal(destRule)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		err = json.Unmarshal(bytes, &DStemplate)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		id = strconv.Itoa(rand.Int())
		DStemplate.ServiceId = id
		svcTemp.AfterServices = append(svcTemp.AfterServices, &DStemplate.ServiceId)
		DStemplate.BeforeServices = append(DStemplate.BeforeServices, &svcTemp.ServiceId)

		//svcComponents = append(svcComponents, svcTemp)
		svcComponents = append(svcComponents, VStemplate)
		svcComponents = append(svcComponents, DStemplate)
		return svcComponents, nil

	} else {

		VStemplate := GetExistingService(svcTemp.Namespace, meshConstants.VirtualService, svcTemp.Name)
		DRtemplate := GetExistingService(svcTemp.Namespace, meshConstants.DestinationRule, svcTemp.Name)
		bytes, err := json.Marshal(VStemplate.ServiceAttributes)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		virtualServiceAttr := new(meshTypes.VSServiceAttribute)
		err = json.Unmarshal(bytes, &virtualServiceAttr)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}

		for _, value := range cpKubeService.ServiceAttributes.Selector {
			for _, label := range labels {
				if label == value {
					continue
				}
				routeRule := new(meshTypes.RouteDestination)
				routeRule.Host = value
				routeRule.Subset = label
				routeRule.Port = cpKubeService.ServiceAttributes.Ports[0].Port
				virtualServiceAttr.Http[0].HttpRoute[0].Routes = append(virtualServiceAttr.Http[0].HttpRoute[0].Routes, routeRule)

			}
		}

		bytes, err = json.Marshal(virtualServiceAttr)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		err = json.Unmarshal(bytes, &VStemplate.ServiceAttributes)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}

		bytes, err = json.Marshal(DRtemplate.ServiceAttributes)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		var destRuleAttr meshTypes.DRServiceAttribute
		err = json.Unmarshal(bytes, &destRuleAttr)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}

		for _, value := range cpKubeService.ServiceAttributes.Selector {
			destRuleAttr.Host = value
			for key, label := range labels {
				subset := new(meshTypes.Subset)
				if label == value {
					continue
				} else {
					subset.Name = label
					lab := make(map[string]string)
					lab[key] = label
					subset.Labels = &lab
					destRuleAttr.Subsets = append(destRuleAttr.Subsets, subset)
				}
			}
		}
		bytes, err = json.Marshal(destRuleAttr)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		err = json.Unmarshal(bytes, &DRtemplate.ServiceAttributes)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		return nil, nil
	}
}

//func CreateIstioVirtualService(svc v2.Service, labels map[string]string) (*istioClient.VirtualService, error) {
//	istioVS := new(istioClient.VirtualService)
//	istioVS.Kind = "VirtualService"
//	istioVS.APIVersion = "networking.istio.io/v1alpha3"
//	istioVS.Name = svc.Name
//	istioVS.Namespace = svc.Namespace
//	istioVS.Spec.Hosts = []string{svc.Name}
//
//	for key, value := range labels {
//		if key == "app" {
//			continue
//		}
//		httpRoute := new(v1alpha3.HTTPRoute)
//		routeRule := new(v1alpha3.HTTPRouteDestination)
//		routeRule.Destination = &v1alpha3.Destination{Host: svc.Name, Subset: value}
//		httpRoute.Route = append(httpRoute.Route, routeRule)
//		istioVS.Spec.Http = append(istioVS.Spec.Http, httpRoute)
//	}
//
//	return istioVS, nil
//
//}
//
//func CreateIstioDestinationRule(svc v2.Service, labels map[string]string) (*istioClient.DestinationRule, error) {
//	destRule := new(istioClient.DestinationRule)
//	destRule.Kind = "DestinationRule"
//	destRule.APIVersion = "networking.istio.io/v1alpha3"
//	destRule.Name = svc.Name
//	destRule.Namespace = svc.Namespace
//	destRule.Spec.Host = svc.Name
//	destRule.Spec.TrafficPolicy = new(v1alpha3.TrafficPolicy)
//	destRule.Spec.TrafficPolicy.Tls = new(v1alpha3.TLSSettings)
//	destRule.Spec.TrafficPolicy.Tls.Mode = v1alpha3.TLSSettings_ISTIO_MUTUAL
//	destRule.Spec.TrafficPolicy.LoadBalancer = new(v1alpha3.LoadBalancerSettings)
//	destRule.Spec.TrafficPolicy.LoadBalancer.LbPolicy = &v1alpha3.LoadBalancerSettings_Simple{
//		Simple: v1alpha3.LoadBalancerSettings_ROUND_ROBIN,
//	}
//
//	var subsets []*v1alpha3.Subset
//	for key, value := range labels {
//		subset := new(v1alpha3.Subset)
//		if value == svc.Name {
//			continue
//		} else {
//			subset.Name = value
//			subset.Labels = make(map[string]string)
//			subset.Labels[key] = value
//			subsets = append(subsets, subset)
//		}
//	}
//	return destRule, nil
//}
