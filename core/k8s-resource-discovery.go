package core

import (
	"bitbucket.org/cloudplex-devs/istio-service-mesh/constants"
	"bitbucket.org/cloudplex-devs/istio-service-mesh/types"
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

	utils.Info.Printf("Resolving dependency :%v", getLogData(types.AppDiscoveryLog{
		ProjectId:   conn.ProjectId,
		ServiceType: string(constants.Job),
		ServiceName: job.Name,
		Namespace:   job.Namespace,
	}))

	jobTemp, err := conn.getCpConvertedTemplate(job, job.Kind)
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
		svcAccTemp, err := conn.getCpConvertedTemplate(svcaccount, svcaccount.Kind)
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
					jobTemp.Embeds = append(jobTemp.Embeds, rbacTemp.ServiceId)
					jobTemp.Embeds = append(jobTemp.Embeds, rbacTemp.ServiceId)
					rbacTemp.Deleted = true
					if len(rbacTemp.Embeds) > 0 {
						rbacTemp.IsEmbedded = true
					}

				}
				serviceTemplates = append(serviceTemplates, rbacTemp)
			}
		} else {
			service := GetExistingService(svcAccTemp.Namespace, svcAccTemp.ServiceSubType, svcAccTemp.Name)
			jobTemp.BeforeServices = append(jobTemp.BeforeServices, &service.ServiceId)
			service.AfterServices = append(service.AfterServices, &jobTemp.ServiceId)
			jobTemp.Embeds = append(jobTemp.Embeds, service.ServiceId)

		}
	}

	//image pull secrets
	for _, objRef := range job.Spec.Template.Spec.ImagePullSecrets {
		secretname := objRef.Name
		secret, err := conn.getSecret(ctx, secretname, namespace)
		if err != nil {
			return
		}
		secretTemp, err := conn.getCpConvertedTemplate(secret, secret.Kind)
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
				hpaTemplate, err := conn.getCpConvertedTemplate(hpa, hpa.Kind)
				if err != nil {
					utils.Error.Println(err)
					return
				}
				hpaTemplate.BeforeServices = append(hpaTemplate.BeforeServices, &jobTemp.ServiceId)
				jobTemp.AfterServices = append(jobTemp.AfterServices, &hpaTemplate.ServiceId)
				jobTemp.Embeds = append(jobTemp.Embeds, hpaTemplate.ServiceId)
				serviceTemplates = append(serviceTemplates, hpaTemplate)
			}
		}
	}

	//kubernetes service depecndency findings
	labels := getLabels(job, job.Kind) //it is better to write get label function
	kubeSvcList, err := conn.getKubernetesServices(ctx, labels, namespace)
	if err != nil {
		return
	}

	//container dependency finding
	for _, container := range job.Spec.Template.Spec.Containers {

		//resolving dependencies of kubernetes service
		if kubeSvcList != nil && len(kubeSvcList.Items) > 0 {
			for _, kubeSvc := range kubeSvcList.Items {
				if isPortMatched(&kubeSvc, &container) {
					k8serviceTemp, err := conn.getCpConvertedTemplate(kubeSvc, kubeSvc.Kind)
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
						k8serviceTemp.BeforeServices = append(k8serviceTemp.BeforeServices, &jobTemp.ServiceId)
						jobTemp.AfterServices = append(jobTemp.AfterServices, &k8serviceTemp.ServiceId)
						for _, key := range k8serviceTemp.AfterServices {
							jobTemp.Embeds = append(jobTemp.Embeds, *key)
						}
						jobTemp.Embeds = append(jobTemp.Embeds, k8serviceTemp.ServiceId)
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
							for _, key := range k8serviceTemp.AfterServices {
								jobTemp.Embeds = append(jobTemp.Embeds, *key)
							}
							jobTemp.Embeds = append(jobTemp.Embeds, k8serviceTemp.ServiceId)
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
				secretTemp, err := conn.getCpConvertedTemplate(secret, secret.Kind)
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
				configmapTemp, err := conn.getCpConvertedTemplate(configmap, configmap.Kind)
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
			secretTemp, err := conn.getCpConvertedTemplate(secret, secret.Kind)
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
			configmapTemp, err := conn.getCpConvertedTemplate(configmap, configmap.Kind)
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
			pvcTemp, err := conn.getCpConvertedTemplate(pvc, pvc.Kind)
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
				storageClassTemp, err := conn.getCpConvertedTemplate(storageClass, storageClass.Kind)
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
					configmapTemp, err := conn.getCpConvertedTemplate(configmap, configmap.Kind)
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
					secretTemp, err := conn.getCpConvertedTemplate(secret, secret.Kind)
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

	utils.Info.Printf("Resolved dependency :%v", getLogData(types.AppDiscoveryLog{
		ProjectId:   conn.ProjectId,
		ServiceType: string(constants.Job),
		ServiceName: job.Name,
		Namespace:   job.Namespace,
	}))
}

func (conn *GrpcConn) ResolveCronJobDependencies(cronjob v1beta1.CronJob, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()

	utils.Info.Printf("Resolving dependency :%v", getLogData(types.AppDiscoveryLog{
		ProjectId:   conn.ProjectId,
		ServiceType: string(constants.CronJob),
		ServiceName: cronjob.Name,
		Namespace:   cronjob.Namespace,
	}))

	cronjobTemp, err := conn.getCpConvertedTemplate(cronjob, cronjob.Kind)
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
		svcAccTemp, err := conn.getCpConvertedTemplate(svcaccount, svcaccount.Kind)
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
					cronjobTemp.Embeds = append(cronjobTemp.Embeds, rbacTemp.ServiceId)
				}
				serviceTemplates = append(serviceTemplates, rbacTemp)
			}
		} else {
			service := GetExistingService(svcAccTemp.Namespace, svcAccTemp.ServiceSubType, svcAccTemp.Name)
			cronjobTemp.BeforeServices = append(cronjobTemp.BeforeServices, &service.ServiceId)
			service.AfterServices = append(service.AfterServices, &cronjobTemp.ServiceId)
			cronjobTemp.Embeds = append(cronjobTemp.Embeds, service.ServiceId)

		}
	}

	//image pull secrets
	for _, objRef := range cronjob.Spec.JobTemplate.Spec.Template.Spec.ImagePullSecrets {
		secretname := objRef.Name
		secret, err := conn.getSecret(ctx, secretname, namespace)
		if err != nil {
			return
		}
		secretTemp, err := conn.getCpConvertedTemplate(secret, secret.Kind)
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
				hpaTemplate, err := conn.getCpConvertedTemplate(hpa, hpa.Kind)
				if err != nil {
					utils.Error.Println(err)
					return
				}
				hpaTemplate.BeforeServices = append(hpaTemplate.BeforeServices, &cronjobTemp.ServiceId)
				cronjobTemp.AfterServices = append(cronjobTemp.AfterServices, &hpaTemplate.ServiceId)
				cronjobTemp.Embeds = append(cronjobTemp.Embeds, hpaTemplate.ServiceId)
				serviceTemplates = append(serviceTemplates, hpaTemplate)
			}
		}
	}

	//kubernetes service depecndency findings
	labels := getLabels(cronjob, cronjob.Kind) //it is better to write get label function
	kubeSvcList, err := conn.getKubernetesServices(ctx, labels, namespace)
	if err != nil {
		return
	}

	//container dependency finding
	for _, container := range cronjob.Spec.JobTemplate.Spec.Template.Spec.Containers {

		//resolving dependencies of kubernetes service
		if kubeSvcList != nil && len(kubeSvcList.Items) > 0 {
			for _, kubeSvc := range kubeSvcList.Items {
				if isPortMatched(&kubeSvc, &container) {
					k8serviceTemp, err := conn.getCpConvertedTemplate(kubeSvc, kubeSvc.Kind)
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
						k8serviceTemp.BeforeServices = append(k8serviceTemp.BeforeServices, &cronjobTemp.ServiceId)
						cronjobTemp.AfterServices = append(cronjobTemp.AfterServices, &k8serviceTemp.ServiceId)
						for _, key := range k8serviceTemp.AfterServices {
							cronjobTemp.Embeds = append(cronjobTemp.Embeds, *key)
						}
						cronjobTemp.Embeds = append(cronjobTemp.Embeds, k8serviceTemp.ServiceId)
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
							for _, key := range k8serviceTemp.AfterServices {
								cronjobTemp.Embeds = append(cronjobTemp.Embeds, *key)
							}
							cronjobTemp.Embeds = append(cronjobTemp.Embeds, k8serviceTemp.ServiceId)

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
				secretTemp, err := conn.getCpConvertedTemplate(secret, secret.Kind)
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
				configmapTemp, err := conn.getCpConvertedTemplate(configmap, configmap.Kind)
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
			secretTemp, err := conn.getCpConvertedTemplate(secret, secret.Kind)
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
			configmapTemp, err := conn.getCpConvertedTemplate(configmap, configmap.Kind)
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
			pvcTemp, err := conn.getCpConvertedTemplate(pvc, pvc.Kind)
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
				storageClassTemp, err := conn.getCpConvertedTemplate(storageClass, storageClass.Kind)
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
					configmapTemp, err := conn.getCpConvertedTemplate(configmap, configmap.Kind)
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
					secretTemp, err := conn.getCpConvertedTemplate(secret, secret.Kind)
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

	utils.Info.Printf("Resolved dependency :%v", getLogData(types.AppDiscoveryLog{
		ProjectId:   conn.ProjectId,
		ServiceType: string(constants.CronJob),
		ServiceName: cronjob.Name,
		Namespace:   cronjob.Namespace,
	}))
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

	utils.Info.Printf("Resolving dependency :%v", getLogData(types.AppDiscoveryLog{
		ProjectId:   conn.ProjectId,
		ServiceType: string(constants.DaemonSet),
		ServiceName: daemonset.Name,
		Namespace:   daemonset.Namespace,
	}))

	daemonsetTemp, err := conn.getCpConvertedTemplate(daemonset, daemonset.Kind)
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
		svcAccTemp, err := conn.getCpConvertedTemplate(svcaccount, svcaccount.Kind)
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
					daemonsetTemp.Embeds = append(daemonsetTemp.Embeds, rbacTemp.ServiceId)

				}
				serviceTemplates = append(serviceTemplates, rbacTemp)
			}
		} else {
			service := GetExistingService(svcAccTemp.Namespace, svcAccTemp.ServiceSubType, svcAccTemp.Name)
			daemonsetTemp.BeforeServices = append(daemonsetTemp.BeforeServices, &service.ServiceId)
			service.AfterServices = append(service.AfterServices, &daemonsetTemp.ServiceId)
			daemonsetTemp.Embeds = append(daemonsetTemp.Embeds, service.ServiceId)
		}
	}

	//image pull secrets
	for _, objRef := range daemonset.Spec.Template.Spec.ImagePullSecrets {
		secretname := objRef.Name
		secret, err := conn.getSecret(ctx, secretname, namespace)
		if err != nil {
			return
		}
		secretTemp, err := conn.getCpConvertedTemplate(secret, secret.Kind)
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
	labels := getLabels(daemonset, daemonset.Kind) //it is better to write get label function
	kubeSvcList, err := conn.getKubernetesServices(ctx, labels, namespace)
	if err != nil {
		return
	}

	//container dependency finding
	for _, container := range daemonset.Spec.Template.Spec.Containers {

		//resolving dependencies of kubernetes service
		if kubeSvcList != nil && len(kubeSvcList.Items) > 0 {
			for _, kubeSvc := range kubeSvcList.Items {
				if isPortMatched(&kubeSvc, &container) {
					k8serviceTemp, err := conn.getCpConvertedTemplate(kubeSvc, kubeSvc.Kind)
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
						k8serviceTemp.BeforeServices = append(k8serviceTemp.BeforeServices, &daemonsetTemp.ServiceId)
						daemonsetTemp.AfterServices = append(daemonsetTemp.AfterServices, &k8serviceTemp.ServiceId)
						for _, key := range k8serviceTemp.AfterServices {
							daemonsetTemp.Embeds = append(daemonsetTemp.Embeds, *key)
						}
						daemonsetTemp.Embeds = append(daemonsetTemp.Embeds, k8serviceTemp.ServiceId)
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
							for _, key := range k8serviceTemp.AfterServices {
								daemonsetTemp.Embeds = append(daemonsetTemp.Embeds, *key)
							}
							daemonsetTemp.Embeds = append(daemonsetTemp.Embeds, k8serviceTemp.ServiceId)
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
				secretTemp, err := conn.getCpConvertedTemplate(secret, secret.Kind)
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
				configmapTemp, err := conn.getCpConvertedTemplate(configmap, configmap.Kind)
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
			secretTemp, err := conn.getCpConvertedTemplate(secret, secret.Kind)
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
			configmapTemp, err := conn.getCpConvertedTemplate(configmap, configmap.Kind)
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
			pvcTemp, err := conn.getCpConvertedTemplate(pvc, pvc.Kind)
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
				storageClassTemp, err := conn.getCpConvertedTemplate(storageClass, storageClass.Kind)
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
					configmapTemp, err := conn.getCpConvertedTemplate(configmap, configmap.Kind)
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
					secretTemp, err := conn.getCpConvertedTemplate(secret, secret.Kind)
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

	utils.Info.Printf("Resolved dependency :%v", getLogData(types.AppDiscoveryLog{
		ProjectId:   conn.ProjectId,
		ServiceType: string(constants.DaemonSet),
		ServiceName: daemonset.Name,
		Namespace:   daemonset.Namespace,
	}))
}

func (conn *GrpcConn) ResolveStatefulSetDependencies(statefulset v1.StatefulSet, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()

	utils.Info.Printf("Resolving dependency :%v", getLogData(types.AppDiscoveryLog{
		ProjectId:   conn.ProjectId,
		ServiceType: string(constants.StatefulSet),
		ServiceName: statefulset.Name,
		Namespace:   statefulset.Namespace,
	}))

	stsTemp, err := conn.getCpConvertedTemplate(statefulset, statefulset.Kind)
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
		svcAccTemp, err := conn.getCpConvertedTemplate(svcaccount, svcaccount.Kind)
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
					stsTemp.Embeds = append(stsTemp.Embeds, rbacTemp.ServiceId)

				}
				serviceTemplates = append(serviceTemplates, rbacTemp)
			}
		} else {
			service := GetExistingService(svcAccTemp.Namespace, svcAccTemp.ServiceSubType, svcAccTemp.Name)
			stsTemp.BeforeServices = append(stsTemp.BeforeServices, &service.ServiceId)
			service.AfterServices = append(service.AfterServices, &stsTemp.ServiceId)
			stsTemp.Embeds = append(stsTemp.Embeds, service.ServiceId)

		}
	}

	//image pull secrets
	for _, objRef := range statefulset.Spec.Template.Spec.ImagePullSecrets {
		secretname := objRef.Name
		secret, err := conn.getSecret(ctx, secretname, namespace)
		if err != nil {
			return
		}
		secretTemp, err := conn.getCpConvertedTemplate(secret, secret.Kind)
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
				hpaTemplate, err := conn.getCpConvertedTemplate(hpa, hpa.Kind)
				if err != nil {
					utils.Error.Println(err)
					return
				}
				hpaTemplate.BeforeServices = append(hpaTemplate.BeforeServices, &stsTemp.ServiceId)
				stsTemp.AfterServices = append(stsTemp.AfterServices, &hpaTemplate.ServiceId)
				stsTemp.Embeds = append(stsTemp.Embeds, hpaTemplate.ServiceId)
				serviceTemplates = append(serviceTemplates, hpaTemplate)
			}
		}
	}

	//kubernetes service depecndency findings
	labels := getLabels(statefulset, statefulset.Kind) //it is better to write get label function
	kubeSvcList, err := conn.getKubernetesServices(ctx, labels, namespace)
	if err != nil {
		return
	}

	//container dependency finding
	for _, container := range statefulset.Spec.Template.Spec.Containers {

		//resolving dependencies of kubernetes service
		if kubeSvcList != nil && len(kubeSvcList.Items) > 0 {
			for _, kubeSvc := range kubeSvcList.Items {
				if isPortMatched(&kubeSvc, &container) {
					k8serviceTemp, err := conn.getCpConvertedTemplate(kubeSvc, kubeSvc.Kind)
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
						k8serviceTemp.BeforeServices = append(k8serviceTemp.BeforeServices, &stsTemp.ServiceId)
						stsTemp.AfterServices = append(stsTemp.AfterServices, &k8serviceTemp.ServiceId)
						for _, key := range k8serviceTemp.AfterServices {
							stsTemp.Embeds = append(stsTemp.Embeds, *key)
						}
						stsTemp.Embeds = append(stsTemp.Embeds, k8serviceTemp.ServiceId)
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
							for _, key := range k8serviceTemp.AfterServices {
								stsTemp.Embeds = append(stsTemp.Embeds, *key)
							}
							stsTemp.Embeds = append(stsTemp.Embeds, k8serviceTemp.ServiceId)
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
				secretTemp, err := conn.getCpConvertedTemplate(secret, secret.Kind)
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
				configmapTemp, err := conn.getCpConvertedTemplate(configmap, configmap.Kind)
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
			secretTemp, err := conn.getCpConvertedTemplate(secret, secret.Kind)
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
			configmapTemp, err := conn.getCpConvertedTemplate(configmap, configmap.Kind)
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
			pvcTemp, err := conn.getCpConvertedTemplate(pvc, pvc.Kind)
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
				storageClassTemp, err := conn.getCpConvertedTemplate(storageClass, storageClass.Kind)
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
					configmapTemp, err := conn.getCpConvertedTemplate(configmap, configmap.Kind)
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
					secretTemp, err := conn.getCpConvertedTemplate(secret, secret.Kind)
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

	utils.Info.Printf("Resolved dependency :%v", getLogData(types.AppDiscoveryLog{
		ProjectId:   conn.ProjectId,
		ServiceType: string(constants.StatefulSet),
		ServiceName: statefulset.Name,
		Namespace:   statefulset.Namespace,
	}))
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
	utils.Info.Printf("Resolving dependency :%v", getLogData(types.AppDiscoveryLog{
		ProjectId:   conn.ProjectId,
		ServiceType: string(constants.Deployment),
		ServiceName: dep.Name,
		Namespace:   dep.Namespace,
	}))

	depTemp, err := conn.getCpConvertedTemplate(dep, dep.Kind)
	if err != nil {
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
		svcAccTemp, err := conn.getCpConvertedTemplate(svcaccount, svcaccount.Kind)
		if err != nil {
			return
		}
		if !isAlreadyExist(svcAccTemp.Namespace, svcAccTemp.ServiceSubType, svcAccTemp.Name) {
			rbacDependencies, err := conn.getK8sRbacResources(ctx, namespace, svcaccount, svcAccTemp)
			if err != nil {
				return
			}

			for _, rbacTemp := range rbacDependencies {
				if rbacTemp.ServiceSubType == meshConstants.ServiceAccount {
					depTemp.BeforeServices = append(depTemp.BeforeServices, &rbacTemp.ServiceId)
					rbacTemp.AfterServices = append(rbacTemp.AfterServices, &depTemp.ServiceId)
					depTemp.Embeds = append(depTemp.Embeds, rbacTemp.ServiceId)
					rbacTemp.Deleted = true
					if len(rbacTemp.Embeds) > 0 {
						rbacTemp.IsEmbedded = true
					}

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
			return
		}
		secretTemp, err := conn.getCpConvertedTemplate(secret, secret.Kind)
		if err != nil {
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
				hpaTemplate, err := conn.getCpConvertedTemplate(hpa, hpa.Kind)
				if err != nil {
					return
				}
				hpaTemplate.BeforeServices = append(hpaTemplate.BeforeServices, &depTemp.ServiceId)
				depTemp.AfterServices = append(depTemp.AfterServices, &hpaTemplate.ServiceId)
				depTemp.Embeds = append(depTemp.Embeds, hpaTemplate.ServiceId)
				hpaTemplate.Deleted = true
				serviceTemplates = append(serviceTemplates, hpaTemplate)
			}
		}
	}

	if depTemp.Name == "frontend" {
		fmt.Println("for debuggin")
	}
	//finding kubernetes service
	labels := getLabels(dep, dep.Kind) //it is better to write get label function
	kubeSvcList, err := conn.getKubernetesServices(ctx, labels, namespace)
	if err != nil {
		return
	}

	//container dependency finding
	for _, container := range dep.Spec.Template.Spec.Containers {
		err = conn.resolveContainerDependency(ctx, kubeSvcList, container, depTemp, namespace, labels)
		if err != nil {
			return
		}
	}

	//volume dependency finding
	for _, vol := range dep.Spec.Template.Spec.Volumes {
		if vol.Secret != nil {
			secretname := vol.Secret.SecretName
			secret, err := conn.getSecret(ctx, secretname, namespace)
			if err != nil {
				return
			}
			secretTemp, err := conn.getCpConvertedTemplate(secret, secret.Kind)
			if err != nil {
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
				return
			}
			configmapTemp, err := conn.getCpConvertedTemplate(configmap, configmap.Kind)
			if err != nil {
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
				return
			}
			pvcTemp, err := conn.getCpConvertedTemplate(pvc, pvc.Kind)
			if err != nil {
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
					return
				}
				storageClassTemp, err := conn.getCpConvertedTemplate(storageClass, storageClass.Kind)
				if err != nil {
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
					configmapTemp, err := conn.getCpConvertedTemplate(configmap, configmap.Kind)
					if err != nil {
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
						return
					}
					secretTemp, err := conn.getCpConvertedTemplate(secret, secret.Kind)
					if err != nil {
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

	utils.Info.Printf("Resolved dependency :%v", getLogData(types.AppDiscoveryLog{
		ProjectId:   conn.ProjectId,
		ServiceType: string(constants.Deployment),
		ServiceName: dep.Name,
		Namespace:   dep.Namespace,
	}))
}

func (conn *GrpcConn) discoverIstioComponents(ctx context.Context, svcTemp *svcTypes.ServiceTemplate, kubesvcTemp *svcTypes.ServiceTemplate, namespace string) error {
	err := conn.discoverIstioDestinationRules(ctx, svcTemp, kubesvcTemp, namespace)
	if err != nil {
		return err
	}

	err = conn.discoverIstioVirtualServices(ctx, svcTemp, kubesvcTemp, namespace)
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
		svcEntryTemp, err := conn.getIstioCpConvertedTemplate(svcEntry, svcEntry.Kind)
		if err != nil {
			return err
		}

		if !isAlreadyExist(svcEntryTemp.Namespace, svcEntryTemp.ServiceSubType, svcEntryTemp.Name) {
			serviceTemplates = append(serviceTemplates, svcEntryTemp)
		}
	}
	return nil
}

func (conn *GrpcConn) discoverIstioDestinationRules(ctx context.Context, svcTemp *svcTypes.ServiceTemplate, kubesvcTemp *svcTypes.ServiceTemplate, namespace string) error {

	drList, err := conn.getAllDestinationRules(ctx, namespace)
	if err != nil {
		return err
	}
	for _, dr := range drList.Items {
		if dr.Spec.Host == kubesvcTemp.Name {
			drTemp, err := conn.getIstioCpConvertedTemplate(dr, dr.Kind)
			if err != nil {
				return err
			}
			drTemp.AfterServices = append(drTemp.AfterServices, &svcTemp.ServiceId)
			svcTemp.BeforeServices = append(svcTemp.BeforeServices, &drTemp.ServiceId)
			serviceTemplates = append(serviceTemplates, drTemp)
			break
		}
	}

	return nil
}

func (conn *GrpcConn) discoverIstioVirtualServices(ctx context.Context, svcTemp *svcTypes.ServiceTemplate, kubesvcTemp *svcTypes.ServiceTemplate, namespace string) error {

	vsList, err := conn.getAllVirtualServices(ctx, namespace)
	if err != nil {
		return err
	}
	for _, vs := range vsList.Items {
		vsTemp, err := conn.getIstioCpConvertedTemplate(vs, vs.Kind)
		if err != nil {
			return err
		}
		for _, http := range vs.Spec.Http {
			for _, route := range http.Route {
				if !isAlreadyExist(vsTemp.Namespace, vsTemp.ServiceSubType, vsTemp.Name) && route.Destination.Host == kubesvcTemp.Name {
					vsTemp.AfterServices = append(vsTemp.AfterServices, &svcTemp.ServiceId)
					svcTemp.BeforeServices = append(svcTemp.BeforeServices, &vsTemp.ServiceId)
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

			gatewayTemp, err := conn.getIstioCpConvertedTemplate(istioGateway, istioGateway.Kind)
			if err != nil {
				return err
			}

			if !isAlreadyExist(gatewayTemp.Namespace, gatewayTemp.ServiceSubType, gatewayTemp.Name) {
				gatewayTemp.AfterServices = append(gatewayTemp.AfterServices, &svcTemp.ServiceId)
				svcTemp.BeforeServices = append(svcTemp.BeforeServices, &gatewayTemp.ServiceId)
				serviceTemplates = append(serviceTemplates, gatewayTemp)
			}
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

		utils.Error.Printf("Error while getting list :%v", getLogData(types.AppDiscoveryLog{
			ProjectId:    conn.ProjectId,
			ServiceType:  string(constants.ServiceEntry),
			Namespace:    namespace,
			ErrorMessage: "error from grpc server :" + err.Error(),
		}))

		return nil, errors.New("error from grpc server :" + err.Error())
	}

	var serviceEntryList *istioClient.ServiceEntryList
	err = json.Unmarshal(response.Resource, &serviceEntryList)
	if err != nil {

		utils.Error.Printf("Error while getting list :%v", getLogData(types.AppDiscoveryLog{
			ProjectId:    conn.ProjectId,
			ServiceType:  string(constants.ServiceEntry),
			Namespace:    namespace,
			ErrorMessage: err.Error(),
		}))

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

		utils.Error.Printf("Error while getting service :%v", getLogData(types.AppDiscoveryLog{
			ProjectId:    conn.ProjectId,
			ServiceType:  string(constants.Gateway),
			ServiceName:  gatewayName,
			Namespace:    namespace,
			ErrorMessage: "error from grpc server :" + err.Error(),
		}))

		return nil, errors.New("error from grpc server :" + err.Error())
	}

	var istioGateway *istioClient.Gateway
	err = json.Unmarshal(response.Resource, &istioGateway)
	if err != nil {

		utils.Error.Printf("Error while getting service :%v", getLogData(types.AppDiscoveryLog{
			ProjectId:    conn.ProjectId,
			ServiceType:  string(constants.Gateway),
			ServiceName:  gatewayName,
			Namespace:    namespace,
			ErrorMessage: err.Error(),
		}))

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

		utils.Error.Printf("Error while getting list :%v", getLogData(types.AppDiscoveryLog{
			ProjectId:    conn.ProjectId,
			ServiceType:  string(constants.DestinationRule),
			Namespace:    namespace,
			ErrorMessage: "error from grpc server :" + err.Error(),
		}))

		return nil, errors.New("error from grpc server :" + err.Error())
	}

	var destinationRuleList *istioClient.DestinationRuleList
	err = json.Unmarshal(response.Resource, &destinationRuleList)
	if err != nil {

		utils.Error.Printf("Error while getting list :%v", getLogData(types.AppDiscoveryLog{
			ProjectId:    conn.ProjectId,
			ServiceType:  string(constants.DestinationRule),
			Namespace:    namespace,
			ErrorMessage: err.Error(),
		}))

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

		utils.Error.Printf("Error while getting list :%v", getLogData(types.AppDiscoveryLog{
			ProjectId:    conn.ProjectId,
			ServiceType:  string(constants.VirtualService),
			Namespace:    namespace,
			ErrorMessage: "error from grpc server :" + err.Error(),
		}))

		return nil, errors.New("error from grpc server :" + err.Error())
	}

	var virtualServiceList *istioClient.VirtualServiceList
	err = json.Unmarshal(response.Resource, &virtualServiceList)
	if err != nil {

		utils.Error.Printf("Error while getting list :%v", getLogData(types.AppDiscoveryLog{
			ProjectId:    conn.ProjectId,
			ServiceType:  string(constants.VirtualService),
			Namespace:    namespace,
			ErrorMessage: err.Error(),
		}))

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
		utils.Info.Println("checking existance of 'istio-sytem' namepace :", err)
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
						clstrroleTemp, err := conn.getCpConvertedTemplate(clstrrole, clstrrole.Kind)
						if err != nil {
							return nil, err
						} else {

							clstrrolebindTemp, err := conn.getCpConvertedTemplate(clstrrolebind, clstrrolebind.Kind)
							if err != nil {
								return nil, err
							}
							if !isAlreadyExist(clstrrolebindTemp.Namespace, clstrrolebindTemp.ServiceSubType, clstrrolebindTemp.Name) {
								clstrrolebindTemp.BeforeServices = append(clstrrolebindTemp.BeforeServices, &clstrroleTemp.ServiceId)
								clstrroleTemp.AfterServices = append(clstrroleTemp.AfterServices, &clstrrolebindTemp.ServiceId)

								clstrrolebindTemp.BeforeServices = append(clstrrolebindTemp.BeforeServices, &svcAccTemp.ServiceId)
								svcAccTemp.AfterServices = append(svcAccTemp.AfterServices, &clstrrolebindTemp.ServiceId)
								svcAccTemp.Embeds = append(svcAccTemp.Embeds, clstrroleTemp.ServiceId)
								svcAccTemp.Embeds = append(svcAccTemp.Embeds, clstrrolebindTemp.ServiceId)

								rbacServiceTemplates = append(rbacServiceTemplates, clstrrolebindTemp)
								rbacServiceTemplates = append(rbacServiceTemplates, clstrroleTemp)
							} else {
								clstrrolebindTemp.BeforeServices = append(clstrrolebindTemp.BeforeServices, &svcAccTemp.ServiceId)
								svcAccTemp.AfterServices = append(svcAccTemp.AfterServices, &clstrrolebindTemp.ServiceId)
								svcAccTemp.Embeds = append(svcAccTemp.Embeds, clstrrolebindTemp.ServiceId)
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
						roleTemp, err := conn.getCpConvertedTemplate(role, role.Kind)
						if err != nil {
							return nil, err
						} else {
							rolebindTemp, err := conn.getCpConvertedTemplate(rolebinding, rolebinding.Kind)
							if err != nil {
								return nil, err
							}
							if !isAlreadyExist(rolebindTemp.Namespace, rolebindTemp.ServiceSubType, rolebindTemp.Name) {
								rolebindTemp.BeforeServices = append(rolebindTemp.BeforeServices, &roleTemp.ServiceId)
								roleTemp.AfterServices = append(roleTemp.AfterServices, &rolebindTemp.ServiceId)

								rolebindTemp.BeforeServices = append(rolebindTemp.BeforeServices, &svcAccTemp.ServiceId)
								svcAccTemp.AfterServices = append(svcAccTemp.AfterServices, &rolebindTemp.ServiceId)
								svcAccTemp.Embeds = append(svcAccTemp.Embeds, rolebindTemp.ServiceId)
								svcAccTemp.Embeds = append(svcAccTemp.Embeds, roleTemp.ServiceId)

								rbacServiceTemplates = append(rbacServiceTemplates, rolebindTemp)
								rbacServiceTemplates = append(rbacServiceTemplates, roleTemp)
							} else {
								rolebindTemp.BeforeServices = append(rolebindTemp.BeforeServices, &svcAccTemp.ServiceId)
								svcAccTemp.AfterServices = append(svcAccTemp.AfterServices, &rolebindTemp.ServiceId)
								svcAccTemp.Embeds = append(svcAccTemp.Embeds, rolebindTemp.ServiceId)
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

		utils.Error.Printf("Error while getting list :%v", getLogData(types.AppDiscoveryLog{
			ProjectId:    conn.ProjectId,
			ServiceType:  string(constants.CronJob),
			Namespace:    namespace,
			ErrorMessage: "error from grpc server :" + err.Error(),
		}))

		return nil, errors.New("error from grpc server :" + err.Error())
	}

	var cronjobList *v1beta1.CronJobList
	err = json.Unmarshal(response.Resource, &cronjobList)
	if err != nil {

		utils.Error.Printf("Error while getting list :%v", getLogData(types.AppDiscoveryLog{
			ProjectId:    conn.ProjectId,
			ServiceType:  string(constants.CronJob),
			Namespace:    namespace,
			ErrorMessage: err.Error(),
		}))

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

		utils.Error.Printf("Error while getting list :%v", getLogData(types.AppDiscoveryLog{
			ProjectId:    conn.ProjectId,
			ServiceType:  string(constants.DaemonSet),
			Namespace:    namespace,
			ErrorMessage: "error from grpc server :" + err.Error(),
		}))

		return nil, errors.New("error from grpc server :" + err.Error())
	}

	var daemonsetList *v1.DaemonSetList
	err = json.Unmarshal(response.Resource, &daemonsetList)
	if err != nil {

		utils.Error.Printf("Error while getting list :%v", getLogData(types.AppDiscoveryLog{
			ProjectId:    conn.ProjectId,
			ServiceType:  string(constants.DaemonSet),
			Namespace:    namespace,
			ErrorMessage: err.Error(),
		}))

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

		utils.Error.Printf("Error while getting list :%v", getLogData(types.AppDiscoveryLog{
			ProjectId:    conn.ProjectId,
			ServiceType:  string(constants.StatefulSet),
			Namespace:    namespace,
			ErrorMessage: "error from grpc server :" + err.Error(),
		}))

		return nil, errors.New("error from grpc server :" + err.Error())
	}

	var statefulsetList *v1.StatefulSetList
	err = json.Unmarshal(response.Resource, &statefulsetList)
	if err != nil {

		utils.Error.Printf("Error while getting list :%v", getLogData(types.AppDiscoveryLog{
			ProjectId:    conn.ProjectId,
			ServiceType:  string(constants.StatefulSet),
			Namespace:    namespace,
			ErrorMessage: err.Error(),
		}))

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
		Args:      []string{"get", "deployments", "-n", namespace, "-o", "json"},
	})
	if err != nil {

		utils.Error.Printf("Error while getting list :%v", getLogData(types.AppDiscoveryLog{
			ProjectId:    conn.ProjectId,
			ServiceType:  string(constants.Deployment),
			Namespace:    namespace,
			ErrorMessage: "error from grpc server :" + err.Error(),
		}))

		return nil, errors.New("error from grpc server :" + err.Error())
	}

	var deploymentList *v1.DeploymentList
	err = json.Unmarshal(response.Resource, &deploymentList)
	if err != nil {

		utils.Error.Printf("Error while getting list :%v", getLogData(types.AppDiscoveryLog{
			ProjectId:    conn.ProjectId,
			ServiceType:  string(constants.Deployment),
			Namespace:    namespace,
			ErrorMessage: err.Error(),
		}))

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

		utils.Error.Printf("Error while getting list :%v", getLogData(types.AppDiscoveryLog{
			ProjectId:    conn.ProjectId,
			ServiceType:  string(constants.Job),
			Namespace:    namespace,
			ErrorMessage: "error from grpc server :" + err.Error(),
		}))

		return nil, errors.New("error from grpc server :" + err.Error())
	}

	var jobList *batch.JobList
	err = json.Unmarshal(response.Resource, &jobList)
	if err != nil {

		utils.Error.Printf("Error while getting list :%v", getLogData(types.AppDiscoveryLog{
			ProjectId:    conn.ProjectId,
			ServiceType:  string(constants.Job),
			Namespace:    namespace,
			ErrorMessage: err.Error(),
		}))

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

		utils.Error.Printf("Error while getting list :%v", getLogData(types.AppDiscoveryLog{
			ProjectId:    conn.ProjectId,
			ServiceType:  string(constants.HPA),
			Namespace:    namespace,
			ErrorMessage: "error from grpc server :" + err.Error(),
		}))

		return nil, errors.New("error from grpc server :" + err.Error())
	}

	var hpaList *autoscale.HorizontalPodAutoscalerList
	err = json.Unmarshal(response.Resource, &hpaList)
	if err != nil {

		utils.Error.Printf("Error while getting list :%v", getLogData(types.AppDiscoveryLog{
			ProjectId:    conn.ProjectId,
			ServiceType:  string(constants.HPA),
			Namespace:    namespace,
			ErrorMessage: err.Error(),
		}))

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

		utils.Error.Printf("Error while getting service :%v", getLogData(types.AppDiscoveryLog{
			ProjectId:    conn.ProjectId,
			ServiceType:  string(constants.StorageClass),
			ServiceName:  storgaClassName,
			Namespace:    namespace,
			ErrorMessage: "error from grpc server :" + err.Error(),
		}))

		return nil, errors.New("error from grpc server :" + err.Error())
	}

	var storageClass *storage.StorageClass
	err = json.Unmarshal(response.Resource, &storageClass)
	if err != nil {

		utils.Error.Printf("Error while getting service :%v", getLogData(types.AppDiscoveryLog{
			ProjectId:    conn.ProjectId,
			ServiceType:  string(constants.StorageClass),
			ServiceName:  storgaClassName,
			Namespace:    namespace,
			ErrorMessage: err.Error(),
		}))

		return nil, err
	}

	return storageClass, nil
}

func (conn *GrpcConn) getKubernetesServices(ctx context.Context, labels map[string]string, namespace string) (*v2.ServiceList, error) {

	response, err := pb.NewK8SResourceClient(conn.Connection).GetK8SResource(ctx, &pb.K8SResourceRequest{
		ProjectId: conn.ProjectId,
		CompanyId: conn.CompanyId,
		Token:     conn.token,
		Command:   "kubectl",
		Args:      []string{"get", "svc", "-n", namespace, "-o", "json"},
	})
	if err != nil {

		utils.Error.Printf("Error while getting list :%v", getLogData(types.AppDiscoveryLog{
			ProjectId:    conn.ProjectId,
			ServiceType:  string(constants.Service),
			Namespace:    namespace,
			ErrorMessage: "error from grpc server :" + err.Error(),
		}))

		return nil, errors.New("error from grpc server :" + err.Error())
	}

	var kubeServiceList *v2.ServiceList
	err = json.Unmarshal(response.Resource, &kubeServiceList)
	if err != nil {

		utils.Error.Printf("Error while getting list :%v", getLogData(types.AppDiscoveryLog{
			ProjectId:    conn.ProjectId,
			ServiceType:  string(constants.Service),
			Namespace:    namespace,
			ErrorMessage: err.Error(),
		}))

		return nil, err
	}

	isLabelMatched := false
	var serviceList = new(v2.ServiceList)
	for _, item := range kubeServiceList.Items {
		for kubeKey, kubeLabel := range item.Spec.Selector {
			for key, value := range labels {
				if kubeKey == key && kubeLabel == value {
					isLabelMatched = true
					serviceList.Items = append(serviceList.Items, item)
					break
				}
			}

			if isLabelMatched {
				break
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
		Args:      []string{"get", "roles", rolename, "-n", namespace, "-o", "json"},
	})
	if err != nil {

		utils.Error.Printf("Error while getting service :%v", getLogData(types.AppDiscoveryLog{
			ProjectId:    conn.ProjectId,
			ServiceType:  string(constants.Role),
			ServiceName:  rolename,
			Namespace:    namespace,
			ErrorMessage: "error from grpc server :" + err.Error(),
		}))

		return nil, errors.New("error from grpc server :" + err.Error())
	}

	var role *rbac.Role
	err = json.Unmarshal(response.Resource, &role)
	if err != nil {

		utils.Error.Printf("Error while getting service :%v", getLogData(types.AppDiscoveryLog{
			ProjectId:    conn.ProjectId,
			ServiceType:  string(constants.Role),
			ServiceName:  rolename,
			Namespace:    namespace,
			ErrorMessage: err.Error(),
		}))

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

		utils.Error.Printf("Error while getting list :%v", getLogData(types.AppDiscoveryLog{
			ProjectId:    conn.ProjectId,
			ServiceType:  string(constants.RoleBinding),
			Namespace:    namespace,
			ErrorMessage: "error from grpc server :" + err.Error(),
		}))

		return nil, errors.New("error from grpc server :" + err.Error())
	}

	var rolebindings *rbac.RoleBindingList
	err = json.Unmarshal(response.Resource, &rolebindings)
	if err != nil {

		utils.Error.Printf("Error while getting list :%v", getLogData(types.AppDiscoveryLog{
			ProjectId:    conn.ProjectId,
			ServiceType:  string(constants.RoleBinding),
			Namespace:    namespace,
			ErrorMessage: err.Error(),
		}))

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
		Args:      []string{"get", "pvc", pvcname, "-n", namespace, "-o", "json"},
	})
	if err != nil {

		utils.Error.Printf("Error while getting service :%v", getLogData(types.AppDiscoveryLog{
			ProjectId:    conn.ProjectId,
			ServiceType:  string(constants.PersistentVolumeClaim),
			ServiceName:  pvcname,
			Namespace:    namespace,
			ErrorMessage: "error from grpc server :" + err.Error(),
		}))

		return nil, errors.New("error from grpc server :" + err.Error())
	}

	var pvc *v2.PersistentVolumeClaim
	err = json.Unmarshal(response.Resource, &pvc)
	if err != nil {

		utils.Error.Printf("Error while getting service :%v", getLogData(types.AppDiscoveryLog{
			ProjectId:    conn.ProjectId,
			ServiceType:  string(constants.PersistentVolumeClaim),
			ServiceName:  pvcname,
			Namespace:    namespace,
			ErrorMessage: err.Error(),
		}))

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

		utils.Error.Printf("Error while getting service :%v", getLogData(types.AppDiscoveryLog{
			ProjectId:    conn.ProjectId,
			ServiceType:  string(constants.ConfigMap),
			ServiceName:  configmapname,
			Namespace:    namespace,
			ErrorMessage: "error from grpc server :" + err.Error(),
		}))

		return nil, errors.New("error from grpc server :" + err.Error())
	}

	var configmap *v2.ConfigMap
	err = json.Unmarshal(response.Resource, &configmap)
	if err != nil {

		utils.Error.Printf("Error while getting service :%v", getLogData(types.AppDiscoveryLog{
			ProjectId:    conn.ProjectId,
			ServiceType:  string(constants.ConfigMap),
			ServiceName:  configmapname,
			Namespace:    namespace,
			ErrorMessage: err.Error(),
		}))

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
		Args:      []string{"get", "sa", svcname, "-n", namespace, "-o", "json"},
	})
	if err != nil {

		utils.Error.Printf("Error while getting service :%v", getLogData(types.AppDiscoveryLog{
			ProjectId:    conn.ProjectId,
			ServiceType:  string(constants.ServiceAccount),
			ServiceName:  svcname,
			Namespace:    namespace,
			ErrorMessage: "error from grpc server :" + err.Error(),
		}))

		return nil, errors.New("error from grpc server :" + err.Error())
	}

	var svcAcc *api.ServiceAccount
	err = json.Unmarshal(response.Resource, &svcAcc)
	if err != nil {

		utils.Error.Printf("Error while getting service :%v", getLogData(types.AppDiscoveryLog{
			ProjectId:    conn.ProjectId,
			ServiceType:  string(constants.ServiceAccount),
			ServiceName:  svcname,
			Namespace:    namespace,
			ErrorMessage: err.Error(),
		}))

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
		Args:      []string{"get", "secrets", secretname, "-n", namespace, "-o", "json"},
	})
	if err != nil {

		utils.Error.Printf("Error while getting service :%v", getLogData(types.AppDiscoveryLog{
			ProjectId:    conn.ProjectId,
			ServiceType:  string(constants.Secret),
			ServiceName:  secretname,
			Namespace:    namespace,
			ErrorMessage: "error from grpc server :" + err.Error(),
		}))

		return nil, errors.New("error from grpc server :" + err.Error())
	}

	var scrt *v2.Secret
	err = json.Unmarshal(response.Resource, &scrt)
	if err != nil {

		utils.Error.Printf("Error while getting service :%v", getLogData(types.AppDiscoveryLog{
			ProjectId:    conn.ProjectId,
			ServiceType:  string(constants.Secret),
			ServiceName:  secretname,
			Namespace:    namespace,
			ErrorMessage: err.Error(),
		}))

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

		utils.Error.Printf("Error while getting list :%v", getLogData(types.AppDiscoveryLog{
			ProjectId:    conn.ProjectId,
			ServiceType:  string(constants.ClusterRoleBinding),
			ErrorMessage: "error from grpc server :" + err.Error(),
		}))

		return nil, errors.New("error from grpc server :" + err.Error())
	}

	var clusterrolebindings *rbac.ClusterRoleBindingList
	err = json.Unmarshal(response.Resource, &clusterrolebindings)
	if err != nil {

		utils.Error.Printf("Error while getting list :%v", getLogData(types.AppDiscoveryLog{
			ProjectId:    conn.ProjectId,
			ServiceType:  string(constants.ClusterRoleBinding),
			ErrorMessage: err.Error(),
		}))

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
		Args:      []string{"get", "clusterrole", clusterrolename, "-o", "json"},
	})
	if err != nil {

		utils.Error.Printf("Error while getting service :%v", getLogData(types.AppDiscoveryLog{
			ProjectId:    conn.ProjectId,
			ServiceType:  string(constants.ClusterRole),
			ServiceName:  clusterrolename,
			ErrorMessage: "error from grpc server :" + err.Error(),
		}))

		return nil, errors.New("error from grpc server :" + err.Error())
	}

	var clusterrole *rbac.ClusterRole
	err = json.Unmarshal(response.Resource, &clusterrole)
	if err != nil {

		utils.Error.Printf("Error while getting service :%v", getLogData(types.AppDiscoveryLog{
			ProjectId:    conn.ProjectId,
			ServiceType:  string(constants.ClusterRole),
			ServiceName:  clusterrolename,
			ErrorMessage: err.Error(),
		}))

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

func (conn *GrpcConn) getCpConvertedTemplate(data interface{}, kind string) (*svcTypes.ServiceTemplate, error) {

	var template *svcTypes.ServiceTemplate
	switch constants.K8sKind(kind) {
	case constants.Deployment:
		CpDeployment, err := convertToCPDeployment(data)
		if err != nil {

			ErrDep := data.(v1.Deployment)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.Deployment),
				ServiceName:  ErrDep.Name,
				Namespace:    ErrDep.Namespace,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		bytes, err := json.Marshal(CpDeployment)
		if err != nil {

			ErrDep := data.(v1.Deployment)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.Deployment),
				ServiceName:  ErrDep.Name,
				Namespace:    ErrDep.Namespace,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}

		err = json.Unmarshal(bytes, &template)
		if err != nil {

			ErrDep := data.(v1.Deployment)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.Deployment),
				ServiceName:  ErrDep.Name,
				Namespace:    ErrDep.Namespace,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		id := strconv.Itoa(rand.Int())
		template.ServiceId = id
		template.Version = "v1"
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
		CpJob, err := convertToCPJob(data.(*batch.Job))
		if err != nil {

			ErrJob := data.(batch.Job)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.Job),
				ServiceName:  ErrJob.Name,
				Namespace:    ErrJob.Namespace,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		bytes, err := json.Marshal(CpJob)
		if err != nil {

			ErrJob := data.(batch.Job)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.Job),
				ServiceName:  ErrJob.Name,
				Namespace:    ErrJob.Namespace,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		err = json.Unmarshal(bytes, &template)
		if err != nil {

			ErrJob := data.(batch.Job)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.Job),
				ServiceName:  ErrJob.Name,
				Namespace:    ErrJob.Namespace,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		id := strconv.Itoa(rand.Int())
		template.ServiceId = id
		template.Version = "v1"
	case constants.DaemonSet:
		CpDaemonset, err := convertToCPDaemonSet(data)
		if err != nil {

			ErrDmSet := data.(v1.DaemonSet)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.DaemonSet),
				ServiceName:  ErrDmSet.Name,
				Namespace:    ErrDmSet.Namespace,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		bytes, err := json.Marshal(CpDaemonset)
		if err != nil {

			ErrDmSet := data.(v1.DaemonSet)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.DaemonSet),
				ServiceName:  ErrDmSet.Name,
				Namespace:    ErrDmSet.Namespace,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		err = json.Unmarshal(bytes, &template)
		if err != nil {

			ErrDmSet := data.(v1.DaemonSet)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.DaemonSet),
				ServiceName:  ErrDmSet.Name,
				Namespace:    ErrDmSet.Namespace,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		id := strconv.Itoa(rand.Int())
		template.ServiceId = id
		template.Version = "v1"
	case constants.StatefulSet:
		CpStatefuleSet, err := convertToCPStatefulSet(data)
		if err != nil {

			ErrSts := data.(v1.StatefulSet)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.StatefulSet),
				ServiceName:  ErrSts.Name,
				Namespace:    ErrSts.Namespace,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		bytes, err := json.Marshal(CpStatefuleSet)
		if err != nil {

			ErrSts := data.(v1.StatefulSet)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.StatefulSet),
				ServiceName:  ErrSts.Name,
				Namespace:    ErrSts.Namespace,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		err = json.Unmarshal(bytes, &template)
		if err != nil {

			ErrSts := data.(v1.StatefulSet)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.StatefulSet),
				ServiceName:  ErrSts.Name,
				Namespace:    ErrSts.Namespace,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		id := strconv.Itoa(rand.Int())
		template.ServiceId = id
		template.Version = "v1"
	case constants.Service:
		bytes, err := json.Marshal(data)
		if err != nil {

			ErrKubeSvc := data.(v2.Service)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.Service),
				ServiceName:  ErrKubeSvc.Name,
				Namespace:    ErrKubeSvc.Namespace,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		var k8Service v2.Service
		err = json.Unmarshal(bytes, &k8Service)
		if err != nil {

			ErrKubeSvc := data.(v2.Service)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.Service),
				ServiceName:  ErrKubeSvc.Name,
				Namespace:    ErrKubeSvc.Namespace,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}

		CpKubeService, err := convertToCPKubernetesService(&k8Service)
		if err != nil {

			ErrKubeSvc := data.(v2.Service)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.Service),
				ServiceName:  ErrKubeSvc.Name,
				Namespace:    ErrKubeSvc.Namespace,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		bytes, err = json.Marshal(CpKubeService)
		if err != nil {

			ErrKubeSvc := data.(v2.Service)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.Service),
				ServiceName:  ErrKubeSvc.Name,
				Namespace:    ErrKubeSvc.Namespace,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		err = json.Unmarshal(bytes, &template)
		if err != nil {

			ErrKubeSvc := data.(v2.Service)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.Service),
				ServiceName:  ErrKubeSvc.Name,
				Namespace:    ErrKubeSvc.Namespace,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		id := strconv.Itoa(rand.Int())
		template.ServiceId = id
		template.Version = "v1"
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

		CpConfigMap, err := ConvertToCPConfigMap(data.(*v2.ConfigMap))
		if err != nil {

			ErrConfigMap := data.(v2.ConfigMap)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.ConfigMap),
				ServiceName:  ErrConfigMap.Name,
				Namespace:    ErrConfigMap.Namespace,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		bytes, err := json.Marshal(CpConfigMap)
		if err != nil {

			ErrConfigMap := data.(v2.ConfigMap)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.ConfigMap),
				ServiceName:  ErrConfigMap.Name,
				Namespace:    ErrConfigMap.Namespace,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		err = json.Unmarshal(bytes, &template)
		if err != nil {

			ErrConfigMap := data.(v2.ConfigMap)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.ConfigMap),
				ServiceName:  ErrConfigMap.Name,
				Namespace:    ErrConfigMap.Namespace,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		id := strconv.Itoa(rand.Int())
		template.ServiceId = id
		template.Version = "v1"
		if isAlreadyExist(template.Namespace, template.ServiceSubType, template.Name) {
			template = GetExistingService(template.Namespace, template.ServiceSubType, template.Name)
		}
	case constants.Secret:
		bytes, err := json.Marshal(data)
		if err != nil {

			ErrSecret := data.(v2.Secret)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.Secret),
				ServiceName:  ErrSecret.Name,
				Namespace:    ErrSecret.Namespace,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		var secret v2.Secret
		err = json.Unmarshal(bytes, &secret)
		if err != nil {

			ErrSecret := data.(v2.Secret)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.Secret),
				ServiceName:  ErrSecret.Name,
				Namespace:    ErrSecret.Namespace,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		CpSecret, err := ConvertToCPSecret(&secret)
		if err != nil {

			ErrSecret := data.(v2.Secret)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.Secret),
				ServiceName:  ErrSecret.Name,
				Namespace:    ErrSecret.Namespace,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		bytes, err = json.Marshal(CpSecret)
		if err != nil {

			ErrSecret := data.(v2.Secret)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.Secret),
				ServiceName:  ErrSecret.Name,
				Namespace:    ErrSecret.Namespace,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		err = json.Unmarshal(bytes, &template)
		if err != nil {

			ErrSecret := data.(v2.Secret)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.Secret),
				ServiceName:  ErrSecret.Name,
				Namespace:    ErrSecret.Namespace,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		id := strconv.Itoa(rand.Int())
		template.ServiceId = id
		template.Version = "v1"
		if isAlreadyExist(template.Namespace, template.ServiceSubType, template.Name) {
			template = GetExistingService(template.Namespace, template.ServiceSubType, template.Name)
		}
	case constants.ServiceAccount:
		bytes, err := json.Marshal(data)
		if err != nil {

			ErrSvcAcc := data.(v2.ServiceAccount)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.ServiceAccount),
				ServiceName:  ErrSvcAcc.Name,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		var serviceaccount v2.ServiceAccount
		err = json.Unmarshal(bytes, &serviceaccount)
		if err != nil {

			ErrSvcAcc := data.(v2.ServiceAccount)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.ServiceAccount),
				ServiceName:  ErrSvcAcc.Name,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		CpServiceAccount, err := convertToCPServiceAccount(&serviceaccount)
		if err != nil {

			ErrSvcAcc := data.(v2.ServiceAccount)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.ServiceAccount),
				ServiceName:  ErrSvcAcc.Name,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		bytes, err = json.Marshal(CpServiceAccount)
		if err != nil {

			ErrSvcAcc := data.(v2.ServiceAccount)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.ServiceAccount),
				ServiceName:  ErrSvcAcc.Name,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		err = json.Unmarshal(bytes, &template)
		if err != nil {

			ErrSvcAcc := data.(v2.ServiceAccount)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.ServiceAccount),
				ServiceName:  ErrSvcAcc.Name,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		id := strconv.Itoa(rand.Int())
		template.ServiceId = id
		template.Version = "v1"
		if isAlreadyExist(template.Namespace, template.ServiceSubType, template.Name) {
			template = GetExistingService(template.Namespace, template.ServiceSubType, template.Name)
		}
	case constants.Role:
		bytes, err := json.Marshal(data)
		if err != nil {

			ErrRole := data.(rbac.Role)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.Role),
				ServiceName:  ErrRole.Name,
				Namespace:    ErrRole.Namespace,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		var role rbac.Role
		err = json.Unmarshal(bytes, &role)
		if err != nil {

			ErrRole := data.(rbac.Role)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.Role),
				ServiceName:  ErrRole.Name,
				Namespace:    ErrRole.Namespace,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		CpRole, err := ConvertToCPRole(&role)
		if err != nil {

			ErrRole := data.(rbac.Role)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.Role),
				ServiceName:  ErrRole.Name,
				Namespace:    ErrRole.Namespace,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		bytes, err = json.Marshal(CpRole)
		if err != nil {

			ErrRole := data.(rbac.Role)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.Role),
				ServiceName:  ErrRole.Name,
				Namespace:    ErrRole.Namespace,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		err = json.Unmarshal(bytes, &template)
		if err != nil {

			ErrRole := data.(rbac.Role)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.Role),
				ServiceName:  ErrRole.Name,
				Namespace:    ErrRole.Namespace,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		id := strconv.Itoa(rand.Int())
		template.ServiceId = id
		template.Version = "v1"
		if isAlreadyExist(template.Namespace, template.ServiceSubType, template.Name) {
			template = GetExistingService(template.Namespace, template.ServiceSubType, template.Name)
		}
	case constants.RoleBinding:
		bytes, err := json.Marshal(data)
		if err != nil {

			ErrRoleBinding := data.(rbac.RoleBinding)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.RoleBinding),
				ServiceName:  ErrRoleBinding.Name,
				Namespace:    ErrRoleBinding.Namespace,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		var roleBinding rbac.RoleBinding
		err = json.Unmarshal(bytes, &roleBinding)
		if err != nil {

			ErrRoleBinding := data.(rbac.RoleBinding)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.RoleBinding),
				ServiceName:  ErrRoleBinding.Name,
				Namespace:    ErrRoleBinding.Namespace,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		CpRoleBinding, err := ConvertToCPRoleBinding(&roleBinding)
		if err != nil {

			ErrRoleBinding := data.(rbac.RoleBinding)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.RoleBinding),
				ServiceName:  ErrRoleBinding.Name,
				Namespace:    ErrRoleBinding.Namespace,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		bytes, err = json.Marshal(CpRoleBinding)
		if err != nil {

			ErrRoleBinding := data.(rbac.RoleBinding)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.RoleBinding),
				ServiceName:  ErrRoleBinding.Name,
				Namespace:    ErrRoleBinding.Namespace,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		err = json.Unmarshal(bytes, &template)
		if err != nil {

			ErrRoleBinding := data.(rbac.RoleBinding)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.RoleBinding),
				ServiceName:  ErrRoleBinding.Name,
				Namespace:    ErrRoleBinding.Namespace,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		id := strconv.Itoa(rand.Int())
		template.ServiceId = id
		template.Version = "v1"
		if isAlreadyExist(template.Namespace, template.ServiceSubType, template.Name) {
			template = GetExistingService(template.Namespace, template.ServiceSubType, template.Name)
		}
	case constants.ClusterRole:
		bytes, err := json.Marshal(data)
		if err != nil {

			ErrClusterRole := data.(rbac.ClusterRole)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.ClusterRole),
				ServiceName:  ErrClusterRole.Name,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		var clusterRole rbac.ClusterRole
		err = json.Unmarshal(bytes, &clusterRole)
		if err != nil {

			ErrClusterRole := data.(rbac.ClusterRole)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.ClusterRole),
				ServiceName:  ErrClusterRole.Name,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		CpClusterRole, err := ConvertToCPClusterRole(&clusterRole)
		if err != nil {

			ErrClusterRole := data.(rbac.ClusterRole)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.ClusterRole),
				ServiceName:  ErrClusterRole.Name,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		bytes, err = json.Marshal(CpClusterRole)
		if err != nil {

			ErrClusterRole := data.(rbac.ClusterRole)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.ClusterRole),
				ServiceName:  ErrClusterRole.Name,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		err = json.Unmarshal(bytes, &template)
		if err != nil {

			ErrClusterRole := data.(rbac.ClusterRole)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.ClusterRole),
				ServiceName:  ErrClusterRole.Name,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		id := strconv.Itoa(rand.Int())
		template.ServiceId = id
		template.Version = "v1"
		if isAlreadyExist(template.Namespace, template.ServiceSubType, template.Name) {
			template = GetExistingService(template.Namespace, template.ServiceSubType, template.Name)
		}
	case constants.ClusterRoleBinding:
		bytes, err := json.Marshal(data)
		if err != nil {

			ErrClusterRoleBind := data.(rbac.ClusterRoleBinding)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.ClusterRoleBinding),
				ServiceName:  ErrClusterRoleBind.Name,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		var clusterRoleBinding rbac.ClusterRoleBinding
		err = json.Unmarshal(bytes, &clusterRoleBinding)
		if err != nil {

			ErrClusterRoleBind := data.(rbac.ClusterRoleBinding)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.ClusterRoleBinding),
				ServiceName:  ErrClusterRoleBind.Name,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		CpClusterRoleBinding, err := ConvertToCPClusterRoleBinding(&clusterRoleBinding)
		if err != nil {

			ErrClusterRoleBind := data.(rbac.ClusterRoleBinding)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.ClusterRoleBinding),
				ServiceName:  ErrClusterRoleBind.Name,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		bytes, err = json.Marshal(CpClusterRoleBinding)
		if err != nil {

			ErrClusterRoleBind := data.(rbac.ClusterRoleBinding)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.ClusterRoleBinding),
				ServiceName:  ErrClusterRoleBind.Name,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		err = json.Unmarshal(bytes, &template)
		if err != nil {

			ErrClusterRoleBind := data.(rbac.ClusterRoleBinding)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.ClusterRoleBinding),
				ServiceName:  ErrClusterRoleBind.Name,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		id := strconv.Itoa(rand.Int())
		template.ServiceId = id
		template.Version = "v1"
		if isAlreadyExist(template.Namespace, template.ServiceSubType, template.Name) {
			template = GetExistingService(template.Namespace, template.ServiceSubType, template.Name)
		}
	case constants.PersistentVolume:
		bytes, err := json.Marshal(data)
		if err != nil {

			ErrPv := data.(v2.PersistentVolume)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.PersistentVolume),
				ServiceName:  ErrPv.Name,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		var persistenVolume v2.PersistentVolume
		err = json.Unmarshal(bytes, &persistenVolume)
		if err != nil {

			ErrPv := data.(v2.PersistentVolume)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.PersistentVolume),
				ServiceName:  ErrPv.Name,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		CpPersistentVolume, err := convertToCPPersistentVolume(&persistenVolume)
		if err != nil {

			ErrPv := data.(v2.PersistentVolume)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.PersistentVolume),
				ServiceName:  ErrPv.Name,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		bytes, err = json.Marshal(CpPersistentVolume)
		if err != nil {

			ErrPv := data.(v2.PersistentVolume)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.PersistentVolume),
				ServiceName:  ErrPv.Name,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		err = json.Unmarshal(bytes, &template)
		if err != nil {

			ErrPv := data.(v2.PersistentVolume)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.PersistentVolume),
				ServiceName:  ErrPv.Name,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		id := strconv.Itoa(rand.Int())
		template.ServiceId = id
		template.Version = "v1"
	case constants.PersistentVolumeClaim:
		bytes, err := json.Marshal(data)
		if err != nil {

			ErrPVC := data.(v2.PersistentVolumeClaim)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.PersistentVolumeClaim),
				ServiceName:  ErrPVC.Name,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		var persistenVolumeClaim v2.PersistentVolumeClaim
		err = json.Unmarshal(bytes, &persistenVolumeClaim)
		if err != nil {

			ErrPVC := data.(v2.PersistentVolumeClaim)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.PersistentVolumeClaim),
				ServiceName:  ErrPVC.Name,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		CpPVC, err := convertToCPPersistentVolumeClaim(&persistenVolumeClaim)
		if err != nil {

			ErrPVC := data.(v2.PersistentVolumeClaim)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.PersistentVolumeClaim),
				ServiceName:  ErrPVC.Name,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		bytes, err = json.Marshal(CpPVC)
		if err != nil {

			ErrPVC := data.(v2.PersistentVolumeClaim)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.PersistentVolumeClaim),
				ServiceName:  ErrPVC.Name,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		err = json.Unmarshal(bytes, &template)
		if err != nil {

			ErrPVC := data.(v2.PersistentVolumeClaim)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.PersistentVolumeClaim),
				ServiceName:  ErrPVC.Name,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		id := strconv.Itoa(rand.Int())
		template.ServiceId = id
		template.Version = "v1"
		if isAlreadyExist(template.Namespace, template.ServiceSubType, template.Name) {
			template = GetExistingService(template.Namespace, template.ServiceSubType, template.Name)
		}
	case constants.StorageClass:
		bytes, err := json.Marshal(data)
		if err != nil {

			ErrSC := data.(storage.StorageClass)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.StorageClass),
				ServiceName:  ErrSC.Name,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		var storageClass storage.StorageClass
		err = json.Unmarshal(bytes, &storageClass)
		if err != nil {

			ErrSC := data.(storage.StorageClass)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.StorageClass),
				ServiceName:  ErrSC.Name,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		CpStorageClass, err := convertToCPStorageClass(&storageClass)
		if err != nil {

			ErrSC := data.(storage.StorageClass)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.StorageClass),
				ServiceName:  ErrSC.Name,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		bytes, err = json.Marshal(CpStorageClass)
		if err != nil {

			ErrSC := data.(storage.StorageClass)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.StorageClass),
				ServiceName:  ErrSC.Name,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		err = json.Unmarshal(bytes, &template)
		if err != nil {

			ErrSC := data.(storage.StorageClass)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.StorageClass),
				ServiceName:  ErrSC.Name,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		id := strconv.Itoa(rand.Int())
		template.ServiceId = id
		template.Version = "v1"
		if isAlreadyExist(template.Namespace, template.ServiceSubType, template.Name) {
			template = GetExistingService(template.Namespace, template.ServiceSubType, template.Name)
		}
	default:
		utils.Error.Printf("Kind does not exist in defined switch cases :%v", kind)
		return nil, errors.New("type does not exit")
	}

	return template, nil

}

func (conn *GrpcConn) getIstioCpConvertedTemplate(data interface{}, kind string) (*svcTypes.ServiceTemplate, error) {
	var template *svcTypes.ServiceTemplate

	switch constants.MeshKind(kind) {
	case constants.DestinationRule:
		bytes, err := json.Marshal(data)
		if err != nil {

			ErrDR := data.(istioClient.DestinationRule)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.DestinationRule),
				ServiceName:  ErrDR.Name,
				Namespace:    ErrDR.Namespace,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		var dr istioClient.DestinationRule
		err = json.Unmarshal(bytes, &dr)
		if err != nil {

			ErrDR := data.(istioClient.DestinationRule)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.DestinationRule),
				ServiceName:  ErrDR.Name,
				Namespace:    ErrDR.Namespace,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		CpDr, err := convertToCPDestinationRule(&dr)
		if err != nil {

			ErrDR := data.(istioClient.DestinationRule)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.DestinationRule),
				ServiceName:  ErrDR.Name,
				Namespace:    ErrDR.Namespace,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		bytes, err = json.Marshal(CpDr)
		if err != nil {

			ErrDR := data.(istioClient.DestinationRule)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.DestinationRule),
				ServiceName:  ErrDR.Name,
				Namespace:    ErrDR.Namespace,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		err = json.Unmarshal(bytes, &template)
		if err != nil {

			ErrDR := data.(istioClient.DestinationRule)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.DestinationRule),
				ServiceName:  ErrDR.Name,
				Namespace:    ErrDR.Namespace,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		id := strconv.Itoa(rand.Int())
		template.ServiceId = id
		template.Version = "v1"
		if isAlreadyExist(template.Namespace, template.ServiceSubType, template.Name) {
			template = GetExistingService(template.Namespace, template.ServiceSubType, template.Name)
		}
	case constants.VirtualService:
		bytes, err := json.Marshal(data)
		if err != nil {

			ErrVS := data.(istioClient.VirtualService)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.VirtualService),
				ServiceName:  ErrVS.Name,
				Namespace:    ErrVS.Namespace,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		var vs istioClient.VirtualService
		err = json.Unmarshal(bytes, &vs)
		if err != nil {

			ErrVS := data.(istioClient.VirtualService)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.VirtualService),
				ServiceName:  ErrVS.Name,
				Namespace:    ErrVS.Namespace,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		CpVs, err := convertToCPVirtualService(&vs)
		if err != nil {

			ErrVS := data.(istioClient.VirtualService)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.VirtualService),
				ServiceName:  ErrVS.Name,
				Namespace:    ErrVS.Namespace,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		bytes, err = json.Marshal(CpVs)
		if err != nil {

			ErrVS := data.(istioClient.VirtualService)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.VirtualService),
				ServiceName:  ErrVS.Name,
				Namespace:    ErrVS.Namespace,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		err = json.Unmarshal(bytes, &template)
		if err != nil {

			ErrVS := data.(istioClient.VirtualService)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.VirtualService),
				ServiceName:  ErrVS.Name,
				Namespace:    ErrVS.Namespace,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		id := strconv.Itoa(rand.Int())
		template.ServiceId = id
		template.Version = "v1"
		if isAlreadyExist(template.Namespace, template.ServiceSubType, template.Name) {
			template = GetExistingService(template.Namespace, template.ServiceSubType, template.Name)
		}
	case constants.Gateway:
		bytes, err := json.Marshal(data)
		if err != nil {

			ErrGTW := data.(istioClient.Gateway)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.Gateway),
				ServiceName:  ErrGTW.Name,
				Namespace:    ErrGTW.Namespace,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		var gateway istioClient.Gateway
		err = json.Unmarshal(bytes, &gateway)
		if err != nil {

			ErrGTW := data.(istioClient.Gateway)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.Gateway),
				ServiceName:  ErrGTW.Name,
				Namespace:    ErrGTW.Namespace,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		CpGateway, err := convertToCPGateway(&gateway)
		if err != nil {

			ErrGTW := data.(istioClient.Gateway)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.Gateway),
				ServiceName:  ErrGTW.Name,
				Namespace:    ErrGTW.Namespace,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		bytes, err = json.Marshal(CpGateway)
		if err != nil {

			ErrGTW := data.(istioClient.Gateway)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.Gateway),
				ServiceName:  ErrGTW.Name,
				Namespace:    ErrGTW.Namespace,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		err = json.Unmarshal(bytes, &template)
		if err != nil {

			ErrGTW := data.(istioClient.Gateway)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.Gateway),
				ServiceName:  ErrGTW.Name,
				Namespace:    ErrGTW.Namespace,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		id := strconv.Itoa(rand.Int())
		template.ServiceId = id
		template.Version = "v1"
		if isAlreadyExist(template.Namespace, template.ServiceSubType, template.Name) {
			template = GetExistingService(template.Namespace, template.ServiceSubType, template.Name)
		}
	case constants.ServiceEntry:
		bytes, err := json.Marshal(data)
		if err != nil {

			ErrSvcEntry := data.(istioClient.ServiceEntry)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.ServiceEntry),
				ServiceName:  ErrSvcEntry.Name,
				Namespace:    ErrSvcEntry.Namespace,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		var svcEntry istioClient.ServiceEntry
		err = json.Unmarshal(bytes, &svcEntry)
		if err != nil {

			ErrSvcEntry := data.(istioClient.ServiceEntry)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.ServiceEntry),
				ServiceName:  ErrSvcEntry.Name,
				Namespace:    ErrSvcEntry.Namespace,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		CpSvcEntry, err := convertToCPServiceEntry(&svcEntry)
		if err != nil {

			ErrSvcEntry := data.(istioClient.ServiceEntry)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.ServiceEntry),
				ServiceName:  ErrSvcEntry.Name,
				Namespace:    ErrSvcEntry.Namespace,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		bytes, err = json.Marshal(CpSvcEntry)
		if err != nil {

			ErrSvcEntry := data.(istioClient.ServiceEntry)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.ServiceEntry),
				ServiceName:  ErrSvcEntry.Name,
				Namespace:    ErrSvcEntry.Namespace,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		err = json.Unmarshal(bytes, &template)
		if err != nil {

			ErrSvcEntry := data.(istioClient.ServiceEntry)
			utils.Error.Printf("Error while CP conversion : %v", getLogData(types.AppDiscoveryLog{
				ProjectId:    conn.ProjectId,
				ServiceType:  string(constants.ServiceEntry),
				ServiceName:  ErrSvcEntry.Name,
				Namespace:    ErrSvcEntry.Namespace,
				ErrorMessage: err.Error(),
			}))

			return nil, err
		}
		id := strconv.Itoa(rand.Int())
		template.ServiceId = id
		template.Version = "v1"
		if isAlreadyExist(template.Namespace, template.ServiceSubType, template.Name) {
			template = GetExistingService(template.Namespace, template.ServiceSubType, template.Name)
		}
	default:
		utils.Error.Printf("Kind does not exist in defined switch cases :%v", kind)
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

func getLogData(data interface{}) string {
	bytes, _ := json.Marshal(data)
	return string(bytes)
}

func getLabels(data interface{}, kind string) map[string]string {
	labels := make(map[string]string)

	switch constants.K8sKind(kind) {
	case constants.Deployment:
		dep := data.(v1.Deployment)
		for key, value := range dep.Spec.Template.Labels {
			labels[key] = value
		}
	case constants.StatefulSet:
		sts := data.(v1.StatefulSet)
		for key, value := range sts.Spec.Template.Labels {
			labels[key] = value
		}
	case constants.DaemonSet:
		daemontset := data.(v1.DaemonSet)
		for key, value := range daemontset.Spec.Template.Labels {
			labels[key] = value
		}

	}
	return labels
}

func (conn *GrpcConn) resolveContainerDependency(ctx context.Context, kubeSvcList *v2.ServiceList, container v2.Container, svcTemp *svcTypes.ServiceTemplate, namespace string, labels map[string]string) error {

	//resolving dependencies of kubernetes service
	if kubeSvcList != nil && len(kubeSvcList.Items) > 0 {
		for _, kubeSvc := range kubeSvcList.Items {
			if isPortMatched(&kubeSvc, &container) {
				k8serviceTemp, err := conn.getCpConvertedTemplate(kubeSvc, kubeSvc.Kind)
				if err != nil {
					return err
				}

				if conn.isIstioEnabled(ctx) {
					//Istio components discovery
					err = conn.discoverIstioComponents(ctx, svcTemp, k8serviceTemp, namespace)
					if err != nil {
						utils.Error.Println(err)
						return err
					}
				} else {
					//istio components creation
					istioSvcTemps, err := CreateIstioComponents(k8serviceTemp, labels)
					if err != nil {
						utils.Error.Println(err)
						return err
					}
					for _, istioSvc := range istioSvcTemps {
						serviceTemplates = append(serviceTemplates, istioSvc)
					}
					//istio components creation
				}

				if !isAlreadyExist(k8serviceTemp.Namespace, k8serviceTemp.ServiceSubType, k8serviceTemp.Name) {
					k8serviceTemp.AfterServices = append(k8serviceTemp.AfterServices, &svcTemp.ServiceId)
					svcTemp.BeforeServices = append(svcTemp.BeforeServices, &k8serviceTemp.ServiceId)
					for _, key := range k8serviceTemp.AfterServices {
						svcTemp.Embeds = append(svcTemp.Embeds, *key)
					}
					svcTemp.Embeds = append(svcTemp.Embeds, k8serviceTemp.ServiceId)
					serviceTemplates = append(serviceTemplates, k8serviceTemp)
				} else {
					isSameService := false
					for _, serviceId := range k8serviceTemp.BeforeServices {
						if *serviceId == svcTemp.ServiceId {
							isSameService = true
							break
						}
					}
					if !isSameService {
						k8serviceTemp.AfterServices = append(k8serviceTemp.AfterServices, &svcTemp.ServiceId)
						svcTemp.BeforeServices = append(svcTemp.BeforeServices, &k8serviceTemp.ServiceId)
						for _, key := range k8serviceTemp.AfterServices {
							svcTemp.Embeds = append(svcTemp.Embeds, *key)
						}
						svcTemp.Embeds = append(svcTemp.Embeds, k8serviceTemp.ServiceId)
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
			secretTemp, err := conn.getCpConvertedTemplate(secret, secret.Kind)
			if err != nil {
				return err
			}
			if !isAlreadyExist(secretTemp.Namespace, secretTemp.ServiceSubType, secretTemp.Name) {
				svcTemp.BeforeServices = append(svcTemp.BeforeServices, &secretTemp.ServiceId)
				secretTemp.AfterServices = append(secretTemp.AfterServices, &svcTemp.ServiceId)
				serviceTemplates = append(serviceTemplates, secretTemp)
			} else {
				isSameService := false
				for _, serviceId := range secretTemp.AfterServices {
					if *serviceId == svcTemp.ServiceId {
						isSameService = true
						break
					}
				}
				if !isSameService {
					svcTemp.BeforeServices = append(svcTemp.BeforeServices, &secretTemp.ServiceId)
					secretTemp.AfterServices = append(secretTemp.AfterServices, &svcTemp.ServiceId)
				}
			}
		} else if env.ValueFrom != nil && env.ValueFrom.ConfigMapKeyRef != nil {
			configmapname := env.ValueFrom.ConfigMapKeyRef.Name
			configmap, err := conn.getConfigMap(ctx, configmapname, namespace)
			if err != nil {
				return err
			}
			configmapTemp, err := conn.getCpConvertedTemplate(configmap, configmap.Kind)
			if err != nil {
				return err
			}
			if !isAlreadyExist(configmapTemp.Namespace, configmapTemp.ServiceSubType, configmapTemp.Name) {
				svcTemp.BeforeServices = append(svcTemp.BeforeServices, &configmapTemp.ServiceId)
				configmapTemp.AfterServices = append(configmapTemp.AfterServices, &svcTemp.ServiceId)
				serviceTemplates = append(serviceTemplates, configmapTemp)
			} else {
				isSameService := false
				for _, serviceId := range configmapTemp.AfterServices {
					if *serviceId == svcTemp.ServiceId {
						isSameService = true
						break
					}
				}
				if !isSameService {
					svcTemp.BeforeServices = append(svcTemp.BeforeServices, &configmapTemp.ServiceId)
					configmapTemp.AfterServices = append(configmapTemp.AfterServices, &svcTemp.ServiceId)
				}
			}
		}
	}

	return nil
}
