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
	autoscale "k8s.io/api/autoscaling/v1"
	"k8s.io/api/batch/v1beta1"
	v2 "k8s.io/api/core/v1"
	rbac "k8s.io/api/rbac/v1"
	"math/rand"
	"reflect"
	"sigs.k8s.io/yaml"
	"strconv"
)

type GrpcConn struct {
	Connection *grpc.ClientConn
	ProjectId  string
	CompanyId  string
	token      string
}

var servicesUniquenes map[string]*types.ServiceTemplate
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

	deploymentList, err := grpcConn.getAllDeployments(ctx, "cloudplex-system")
	if err != nil {
		return
	}
	//var deployList *v1.DeploymentList
	//bytes, err := json.Marshal(deploymentList)
	//if err != nil {
	//	utils.Error.Println(err)
	//	return &pb.K8SResourceResponse{}, err
	//}
	//
	//err = json.Unmarshal(bytes, &deployList)
	//if err != nil {
	//	utils.Error.Println(err)
	//	return &pb.K8SResourceResponse{}, err
	//}

	grpcConn.deploymentk8sToCp(ctx, deploymentList.Items)
	fmt.Println("done")

	//response, err = pb.NewK8SResourceClient(conn).GetK8SResource(ctx, request)
	//if err != nil {
	//	utils.Error.Println(err)
	//	return &pb.K8SResourceResponse{}, err
	//}

	//if request.Name == "" {
	//	var dep []*v1.Deployment
	//	err = json.Unmarshal(response.Resource, dep)
	//	if err != nil {
	//		utils.Error.Println(err)
	//		return &pb.K8SResourceResponse{}, err
	//	}
	//} else {
	//	var dep *v1.Deployment
	//	err = json.Unmarshal(response.Resource, dep)
	//	if err != nil {
	//		utils.Error.Println(err)
	//		return &pb.K8SResourceResponse{}, err
	//	}
	//}

	//var dep []*v1.Deployment
	//err = json.Unmarshal(response.Resource, dep)
	//if err != nil {
	//	utils.Error.Println(err)
	//	return &pb.K8SResourceResponse{}, err
	//}

	return response, err
}

//func (conn *GrpcConn) hpaK8sToCp(ctx context.Context, hpas []*autoscale.HorizontalPodAutoscaler) {
//	for _, hpa := range hpas {
//
//	}
//}

func (conn *GrpcConn) cronjobK8sToCp(ctx context.Context, cronjobs []v1beta1.CronJob) {
	for _, cronjob := range cronjobs {
		namespace := cronjob.Namespace
		if cronjob.Spec.JobTemplate.Spec.Template.Spec.ServiceAccountName != "" {
			//svcname := cronjob.Spec.JobTemplate.Spec.Template.Spec.ServiceAccountName
			//conn.getK8sRbacResources(ctx, svcname, namespace, constants.CronJob, cronjob)

		}

		hpaList, err := conn.getAllHpas(ctx, namespace)
		if err != nil {
			return
		}

		if len(hpaList.Items) > 0 {
			for _, hpa := range hpaList.Items {
				if hpa.Spec.ScaleTargetRef.APIVersion == cronjob.APIVersion && hpa.Spec.ScaleTargetRef.Kind == cronjob.Kind && hpa.Spec.ScaleTargetRef.Name == cronjob.Name {

				}
			}
		}

		//kubernetes service depecndency findings
		for key, value := range cronjob.Spec.JobTemplate.Spec.Template.Labels {
			resp, err := conn.getKubernetesServices(ctx, key, value, namespace)
			if err != nil {
				return
			}

			if len(resp.Items) > 0 {
				for _, kubeSvc := range resp.Items {
					fmt.Println(kubeSvc)
				}
			}
		}

		//container dependency finding
		for _, container := range cronjob.Spec.JobTemplate.Spec.Template.Spec.Containers {
			//discovering secret and config maps in deployment containers
			for _, env := range container.Env {
				if env.ValueFrom.SecretKeyRef != nil {
					secretname := env.ValueFrom.SecretKeyRef.Name
					resp, err := conn.getSecret(ctx, secretname, namespace)
					if err != nil {
						return
					}
					fmt.Println(resp)

				} else if env.ValueFrom.ConfigMapKeyRef != nil {
					configmapname := env.ValueFrom.ConfigMapKeyRef.Name
					resp, err := conn.getConfigMap(ctx, configmapname, namespace)
					if err != nil {
						return
					}
					fmt.Println(resp)
				}
			}
		}

		//volume dependency finding
		for _, vol := range cronjob.Spec.JobTemplate.Spec.Template.Spec.Volumes {
			if vol.Secret != nil {
				secretname := vol.Secret.SecretName
				resp, err := conn.getSecret(ctx, secretname, namespace)
				if err != nil {
					return
				}
				fmt.Println(resp)
			} else if vol.ConfigMap != nil {
				configmapname := vol.ConfigMap.Name
				resp, err := conn.getConfigMap(ctx, configmapname, namespace)
				if err != nil {
					return
				}
				fmt.Println(resp)
			} else if vol.PersistentVolumeClaim != nil {
				pvcname := vol.PersistentVolumeClaim.ClaimName
				resp, err := conn.getPvc(ctx, pvcname, namespace)
				if err != nil {
					return
				}
				fmt.Println(resp)
			} //else if vol.AWSElasticBlockStore.VolumeID

			for _, source := range vol.Projected.Sources {
				if source.ConfigMap != nil {
					configmapname := vol.ConfigMap.Name
					resp, err := conn.getConfigMap(ctx, configmapname, namespace)
					if err != nil {
						return
					}
					fmt.Println(resp)
				} else if source.Secret != nil {
					secretname := vol.Secret.SecretName
					resp, err := conn.getSecret(ctx, secretname, namespace)
					if err != nil {
						return
					}
					fmt.Println(resp)
				}
			}
		}

	}
}

func (conn *GrpcConn) daemonsetK8sToCp(ctx context.Context, daemonsets []v1.DaemonSet) {
	for _, daemonset := range daemonsets {
		namespace := daemonset.Namespace
		if daemonset.Spec.Template.Spec.ServiceAccountName != "" {
			//svcname := daemonset.Spec.Template.Spec.ServiceAccountName
			//conn.getK8sRbacResources(ctx, svcname, namespace, constants.Daemonset, daemonset)

		}

		hpaList, err := conn.getAllHpas(ctx, namespace)
		if err != nil {
			return
		}

		if len(hpaList.Items) > 0 {
			for _, hpa := range hpaList.Items {
				if hpa.Spec.ScaleTargetRef.APIVersion == daemonset.APIVersion && hpa.Spec.ScaleTargetRef.Kind == daemonset.Kind && hpa.Spec.ScaleTargetRef.Name == daemonset.Name {

				}
			}
		}

		//kubernetes service depecndency findings
		for key, value := range daemonset.Spec.Template.Labels {
			resp, err := conn.getKubernetesServices(ctx, key, value, namespace)
			if err != nil {
				return
			}

			if len(resp.Items) > 0 {
				for _, kubeSvc := range resp.Items {
					fmt.Println(kubeSvc)
				}
			}
		}

		//container dependency finding
		for _, container := range daemonset.Spec.Template.Spec.Containers {
			//discovering secret and config maps in deployment containers
			for _, env := range container.Env {
				if env.ValueFrom.SecretKeyRef != nil {
					secretname := env.ValueFrom.SecretKeyRef.Name
					resp, err := conn.getSecret(ctx, secretname, namespace)
					if err != nil {
						return
					}
					fmt.Println(resp)

				} else if env.ValueFrom.ConfigMapKeyRef != nil {
					configmapname := env.ValueFrom.ConfigMapKeyRef.Name
					resp, err := conn.getConfigMap(ctx, configmapname, namespace)
					if err != nil {
						return
					}
					fmt.Println(resp)
				}
			}
		}

		//volume dependency finding
		for _, vol := range daemonset.Spec.Template.Spec.Volumes {
			if vol.Secret != nil {
				secretname := vol.Secret.SecretName
				resp, err := conn.getSecret(ctx, secretname, namespace)
				if err != nil {
					return
				}
				fmt.Println(resp)
			} else if vol.ConfigMap != nil {
				configmapname := vol.ConfigMap.Name
				resp, err := conn.getConfigMap(ctx, configmapname, namespace)
				if err != nil {
					return
				}
				fmt.Println(resp)
			} else if vol.PersistentVolumeClaim != nil {
				pvcname := vol.PersistentVolumeClaim.ClaimName
				resp, err := conn.getPvc(ctx, pvcname, namespace)
				if err != nil {
					return
				}
				fmt.Println(resp)
			} //else if vol.AWSElasticBlockStore.VolumeID

			for _, source := range vol.Projected.Sources {
				if source.ConfigMap != nil {
					configmapname := vol.ConfigMap.Name
					resp, err := conn.getConfigMap(ctx, configmapname, namespace)
					if err != nil {
						return
					}
					fmt.Println(resp)
				} else if source.Secret != nil {
					secretname := vol.Secret.SecretName
					resp, err := conn.getSecret(ctx, secretname, namespace)
					if err != nil {
						return
					}
					fmt.Println(resp)
				}
			}
		}

	}
}

func (conn *GrpcConn) statefulsetsK8sToCp(ctx context.Context, statefulsets []v1.StatefulSet) {

	for _, statefulset := range statefulsets {
		namespace := statefulset.Namespace
		if statefulset.Spec.Template.Spec.ServiceAccountName != "" {
			//svcname := statefulset.Spec.Template.Spec.ServiceAccountName
			//conn.getK8sRbacResources(ctx, svcname, namespace, constants.StatefulSet, statefulset)
		}

		hpaList, err := conn.getAllHpas(ctx, namespace)
		if err != nil {
			return
		}

		if len(hpaList.Items) > 0 {
			for _, hpa := range hpaList.Items {
				if hpa.Spec.ScaleTargetRef.APIVersion == statefulset.APIVersion && hpa.Spec.ScaleTargetRef.Kind == statefulset.Kind && hpa.Spec.ScaleTargetRef.Name == statefulset.Name {

				}
			}
		}

		//kubernetes service depecndency findings
		for key, value := range statefulset.Spec.Template.Labels {
			resp, err := conn.getKubernetesServices(ctx, key, value, namespace)
			if err != nil {
				return
			}

			if len(resp.Items) > 0 {
				for _, kubeSvc := range resp.Items {
					fmt.Println(kubeSvc)
				}
			}
		}

		//container dependency finding
		for _, container := range statefulset.Spec.Template.Spec.Containers {
			//discovering secret and config maps in deployment containers
			for _, env := range container.Env {
				if env.ValueFrom.SecretKeyRef != nil {
					secretname := env.ValueFrom.SecretKeyRef.Name
					resp, err := conn.getSecret(ctx, secretname, namespace)
					if err != nil {
						return
					}
					fmt.Println(resp)
				} else if env.ValueFrom.ConfigMapKeyRef != nil {
					configmapname := env.ValueFrom.ConfigMapKeyRef.Name
					resp, err := conn.getConfigMap(ctx, configmapname, namespace)
					if err != nil {
						return
					}
					fmt.Println(resp)
				}
			}
		}

		//volume dependency finding
		for _, vol := range statefulset.Spec.Template.Spec.Volumes {
			if vol.Secret != nil {
				secretname := vol.Secret.SecretName
				resp, err := conn.getSecret(ctx, secretname, namespace)
				if err != nil {
					return
				}
				fmt.Println(resp)
			} else if vol.ConfigMap != nil {
				configmapname := vol.ConfigMap.Name
				resp, err := conn.getConfigMap(ctx, configmapname, namespace)
				if err != nil {
					return
				}
				fmt.Println(resp)
			} else if vol.PersistentVolumeClaim != nil {
				pvcname := vol.PersistentVolumeClaim.ClaimName
				resp, err := conn.getPvc(ctx, pvcname, namespace)
				if err != nil {
					return
				}
				fmt.Println(resp)
			} //else if vol.AWSElasticBlockStore.VolumeID

			for _, source := range vol.Projected.Sources {
				if source.ConfigMap != nil {
					configmapname := vol.ConfigMap.Name
					resp, err := conn.getConfigMap(ctx, configmapname, namespace)
					if err != nil {
						return
					}
					fmt.Println(resp)
				} else if source.Secret != nil {
					secretname := vol.Secret.SecretName
					resp, err := conn.getSecret(ctx, secretname, namespace)
					if err != nil {
						return
					}
					fmt.Println(resp)
				}
			}
		}

	}

}

func (conn *GrpcConn) deploymentk8sToCp(ctx context.Context, deployments []v1.Deployment) ([]*types.ServiceTemplate, error) {

	for _, dep := range deployments {
		depTemp, err := getCpConvertedTemplate(dep, dep.Kind)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}

		namespace := dep.Namespace
		//checking for the service account if name not empty then getting cluster role and cluster role  binding against that service account
		if dep.Spec.Template.Spec.ServiceAccountName != "" {
			svcname := dep.Spec.Template.Spec.ServiceAccountName
			svcaccount, err := conn.getSvcAccount(ctx, svcname, namespace)
			if err != nil {
				return nil, err
			}
			svcAccTemp, err := getCpConvertedTemplate(svcaccount, svcaccount.Kind)
			if err != nil {
				utils.Error.Println(err)
				return nil, err
			}
			if !isAlreadyExist(svcAccTemp.NameSpace, svcAccTemp.ServiceSubType, svcAccTemp.Name) {
				rbacDependencies, err := conn.getK8sRbacResources(ctx, namespace, svcaccount, svcAccTemp)
				if err != nil {
					utils.Error.Println(err)
					return nil, err
				}

				var strpointer = new(string)
				*strpointer = "service_account"
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

		hpaList, err := conn.getAllHpas(ctx, namespace)
		if err != nil {
			return nil, err
		}
		if len(hpaList.Items) > 0 {
			for _, hpa := range hpaList.Items {
				if hpa.Spec.ScaleTargetRef.APIVersion == dep.APIVersion && hpa.Spec.ScaleTargetRef.Kind == dep.Kind && hpa.Spec.ScaleTargetRef.Name == dep.Name {
					hpaTemplate, err := getCpConvertedTemplate(hpa, hpa.Kind)
					if err != nil {
						utils.Error.Println(err)
						return nil, err
					}
					hpaTemplate.BeforeServices = append(hpaTemplate.BeforeServices, depTemp.ServiceId)
					depTemp.AfterServices = append(depTemp.AfterServices, hpaTemplate.ServiceId)
					serviceTemplates = append(serviceTemplates, hpaTemplate)
				}
			}
		}

		//kubernetes service depecndency findings
		for key, value := range dep.Spec.Template.Labels {
			resp, err := conn.getKubernetesServices(ctx, key, value, namespace)
			if err != nil {
				utils.Error.Println(err)
				return nil, err
			}

			if len(resp.Items) > 0 {
				for _, kubeSvc := range resp.Items {
					k8serviceTemp, err := getCpConvertedTemplate(kubeSvc, kubeSvc.Kind)
					if err != nil {
						utils.Error.Println(err)
						return nil, err
					}
					if !isAlreadyExist(k8serviceTemp.NameSpace, k8serviceTemp.ServiceSubType, k8serviceTemp.Name) {
						k8serviceTemp.BeforeServices = append(k8serviceTemp.BeforeServices, depTemp.ServiceId)
						depTemp.AfterServices = append(depTemp.AfterServices, k8serviceTemp.ServiceId)
						serviceTemplates = append(serviceTemplates, k8serviceTemp)
					} else {
						for _, service := range serviceTemplates {
							if *service.NameSpace == *k8serviceTemp.NameSpace && *service.ServiceSubType == *k8serviceTemp.ServiceSubType && *service.Name == *k8serviceTemp.Name {
								service.BeforeServices = append(service.BeforeServices, depTemp.ServiceId)
								depTemp.AfterServices = append(depTemp.AfterServices, service.ServiceId)
							}
						}
					}
				}
			}
		}

		//container dependency finding
		for _, container := range dep.Spec.Template.Spec.Containers {
			//discovering secret and config maps in deployment containers
			for _, env := range container.Env {
				if env.ValueFrom != nil && env.ValueFrom.SecretKeyRef != nil {
					secretname := env.ValueFrom.SecretKeyRef.Name
					secret, err := conn.getSecret(ctx, secretname, namespace)
					if err != nil {
						utils.Error.Println(err)
						return nil, err
					}
					secretTemp, err := getCpConvertedTemplate(secret, secret.Kind)
					if err != nil {
						utils.Error.Println(err)
						return nil, err
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
							serviceTemplates = append(serviceTemplates, secretTemp)
						}
					}
				} else if env.ValueFrom != nil && env.ValueFrom.ConfigMapKeyRef != nil {
					configmapname := env.ValueFrom.ConfigMapKeyRef.Name
					configmap, err := conn.getConfigMap(ctx, configmapname, namespace)
					if err != nil {
						utils.Error.Println(err)
						return nil, err
					}
					configmapTemp, err := getCpConvertedTemplate(configmap, configmap.Kind)
					if err != nil {
						utils.Error.Println(err)
						return nil, err
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
							serviceTemplates = append(serviceTemplates, configmapTemp)
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
					return nil, err
				}
				secretTemp, err := getCpConvertedTemplate(secret, secret.Kind)
				if err != nil {
					utils.Error.Println(err)
					return nil, err
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
						serviceTemplates = append(serviceTemplates, secretTemp)
					}
				}
			} else if vol.ConfigMap != nil {
				configmapname := vol.ConfigMap.Name
				configmap, err := conn.getConfigMap(ctx, configmapname, namespace)
				if err != nil {
					utils.Error.Println(err)
					return nil, err
				}
				configmapTemp, err := getCpConvertedTemplate(configmap, configmap.Kind)
				if err != nil {
					utils.Error.Println(err)
					return nil, err
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
						serviceTemplates = append(serviceTemplates, configmapTemp)
					}
				}
			} else if vol.PersistentVolumeClaim != nil {
				pvcname := vol.PersistentVolumeClaim.ClaimName
				pvc, err := conn.getPvc(ctx, pvcname, namespace)
				if err != nil {
					utils.Error.Println(err)
					return nil, err
				}
				pvcTemp, err := getCpConvertedTemplate(pvc, pvc.Kind)
				if err != nil {
					utils.Error.Println(err)
					return nil, err
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
			}

			if vol.Projected != nil {
				for _, source := range vol.Projected.Sources {
					if source.ConfigMap != nil {
						configmapname := vol.ConfigMap.Name
						configmap, err := conn.getConfigMap(ctx, configmapname, namespace)
						if err != nil {
							utils.Error.Println(err)
							return nil, err
						}
						configmapTemp, err := getCpConvertedTemplate(configmap, configmap.Kind)
						if err != nil {
							utils.Error.Println(err)
							return nil, err
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
								serviceTemplates = append(serviceTemplates, configmapTemp)
							}
						}
					} else if source.Secret != nil {
						secretname := vol.Secret.SecretName
						secret, err := conn.getSecret(ctx, secretname, namespace)
						if err != nil {
							utils.Error.Println(err)
							return nil, err
						}
						secretTemp, err := getCpConvertedTemplate(secret, secret.Kind)
						if err != nil {
							utils.Error.Println(err)
							return nil, err
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
								serviceTemplates = append(serviceTemplates, secretTemp)
							}
						}
					}
				}
			}
		}

		serviceTemplates = append(serviceTemplates, depTemp)
	}
	return serviceTemplates, nil
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

	var serviceTemplates []*types.ServiceTemplate
	//creating secrets for service account
	for _, secret := range k8svcAcc.Secrets {
		if secret.Name != "" {
			secretname := secret.Name
			if secret.Namespace != "" {
				namespace = secret.Namespace
			}
			scrt, err := conn.getSecret(ctx, secretname, namespace)
			if err != nil {
				return nil, err
			} else {
				secretTemp, err := getCpConvertedTemplate(scrt, scrt.Kind)
				if err != nil {
					utils.Error.Println(err)
					return nil, err
				}
				//this is doubtful
				secretTemp.BeforeServices = append(secretTemp.BeforeServices, svcAccTemp.ServiceId)
				serviceTemplates = append(serviceTemplates, secretTemp)
				svcAccTemp.AfterServices = append(svcAccTemp.AfterServices, secretTemp.ServiceId)
				//this is doubtful
			}
		}
	}

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

								serviceTemplates = append(serviceTemplates, clstrrolebindTemp)
								serviceTemplates = append(serviceTemplates, clstrroleTemp)
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

								serviceTemplates = append(serviceTemplates, rolebindTemp)
								serviceTemplates = append(serviceTemplates, roleTemp)
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

	serviceTemplates = append(serviceTemplates, svcAccTemp)
	return serviceTemplates, nil
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

func (conn *GrpcConn) getKubernetesServices(ctx context.Context, key, value, namespace string) (*v2.ServiceList, error) {

	response, err := pb.NewK8SResourceClient(conn.Connection).GetK8SResource(ctx, &pb.K8SResourceRequest{
		ProjectId: conn.ProjectId,
		CompanyId: conn.CompanyId,
		Token:     conn.token,
		Command:   "kubectl",
		Args:      []string{"get", "svc", "-l", key + "=" + value, "-n", namespace, "-o", "json"},
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

	return kubeServiceList, nil
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
		Args:      []string{"get", "clusterrolebindings"},
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

func isAlreadyExist(namespace, svcsubtype, name *string) bool {
	for _, val := range serviceTemplates {
		if val.NameSpace == namespace && val.ServiceSubType == svcsubtype && val.Name == name {
			return true
		}
	}
	return false
}

func getCpConvertedTemplate(data interface{}, kind string) (*types.ServiceTemplate, error) {

	var template *types.ServiceTemplate
	switch kind {
	case "Deployment":
		bytes, err := json.Marshal(data)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}

		//TODO needs to convert cp schema first
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

		//TODO needs to convert cp schema first
		err = json.Unmarshal(bytes, &template)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		id := strconv.Itoa(rand.Int())
		template.ServiceId = &id
	case "DaemonSet":
		bytes, err := json.Marshal(data)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}

		//TODO needs to convert cp schema first
		err = json.Unmarshal(bytes, &template)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		id := strconv.Itoa(rand.Int())
		template.ServiceId = &id
	case "StatefulSet":
		bytes, err := json.Marshal(data)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}

		//TODO needs to convert cp schema first
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

		//TODO needs to convert cp schema first
		err = json.Unmarshal(bytes, &template)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		id := strconv.Itoa(rand.Int())
		template.ServiceId = &id
	case "HorizontalPodAutoscaler":
		bytes, err := json.Marshal(data)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}

		//TODO needs to convert cp schema first
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

		//TODO needs to convert cp schema first
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

		//TODO needs to convert cp schema first
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

		//TODO needs to convert cp schema first
		err = json.Unmarshal(bytes, &template)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		id := strconv.Itoa(rand.Int())
		template.ServiceId = &id
	case "Role":
		bytes, err := json.Marshal(data)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}

		//TODO needs to convert cp schema first
		err = json.Unmarshal(bytes, &template)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		id := strconv.Itoa(rand.Int())
		template.ServiceId = &id
	case "RoleBinding":
		bytes, err := json.Marshal(data)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}

		//TODO needs to convert cp schema first
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

		//TODO needs to convert cp schema first
		err = json.Unmarshal(bytes, &template)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		id := strconv.Itoa(rand.Int())
		template.ServiceId = &id
	case "ClusterRoleBinding":
		bytes, err := json.Marshal(data)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}

		//TODO needs to convert cp schema first
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

		//TODO needs to convert cp schema first
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

		//TODO needs to convert cp schema first
		err = json.Unmarshal(bytes, &template)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		id := strconv.Itoa(rand.Int())
		template.ServiceId = &id
	case "StorageClass":
		bytes, err := json.Marshal(data)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}

		//TODO needs to convert cp schema first
		err = json.Unmarshal(bytes, &template)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		id := strconv.Itoa(rand.Int())
		template.ServiceId = &id
	default:
		utils.Info.Println("Kind does not exist in defined switch cases")
		return nil, errors.New("type does not exit")
	}

	return template, nil

}
