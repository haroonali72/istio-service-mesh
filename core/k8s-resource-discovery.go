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

var services []*string

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
			svcname := cronjob.Spec.JobTemplate.Spec.Template.Spec.ServiceAccountName
			conn.getK8sRbacResources(ctx, svcname, namespace, constants.CronJob, cronjob)

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
			svcname := daemonset.Spec.Template.Spec.ServiceAccountName
			conn.getK8sRbacResources(ctx, svcname, namespace, constants.Daemonset, daemonset)

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
			svcname := statefulset.Spec.Template.Spec.ServiceAccountName
			conn.getK8sRbacResources(ctx, svcname, namespace, constants.StatefulSet, statefulset)
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

func (conn *GrpcConn) deploymentk8sToCp(ctx context.Context, deployments []v1.Deployment) {

	var serviceTemplates []types.ServiceTemplate

	for _, dep := range deployments {

		namespace := dep.Namespace
		//checking for the service account if name not empty then getting cluster role and cluster role  binding against that service account
		if dep.Spec.Template.Spec.ServiceAccountName != "" {

			svcname := dep.Spec.Template.Spec.ServiceAccountName
			conn.getK8sRbacResources(ctx, svcname, namespace, constants.Deployment, dep)
			//svcaccount, err := conn.getSvcAccount(ctx, svcname, namespace)
			//if err != nil {
			//	return
			//}
			//
			////creating secrets for service account
			//for _, secret := range svcaccount.Secrets {
			//	if secret.Name != "" {
			//
			//		secretname := secret.Name
			//		if secret.Namespace != "" {
			//			namespace = secret.Namespace
			//		}
			//		_, err := conn.getSecret(ctx, secretname, namespace)
			//		if err != nil {
			//			return
			//		}
			//	}
			//}
			//
			//clusterrolebindings, err := conn.getAllClusterRoleBindings(ctx)
			//if err != nil {
			//	return
			//}
			//
			//for _, clstrrolebind := range clusterrolebindings.Items {
			//	for _, sub := range clstrrolebind.Subjects {
			//		if sub.Kind == "ServiceAccount" && sub.Name == svcname {
			//			if clstrrolebind.RoleRef.Kind == "ClusterRole" {
			//				clusterrolename := clstrrolebind.RoleRef.Name
			//				_, err := conn.getClusterRole(ctx, clusterrolename)
			//				if err != nil {
			//					return
			//				}
			//
			//				//for mainting uniqueness of services
			//				svc := clstrrolebind.Namespace + "-" + clstrrolebind.Kind + "-" + clstrrolebind.Name
			//				services = append(services, &svc)
			//			}
			//			break
			//		}
			//	}
			//}
			//
			//rolebindings, err := conn.getAllRoleBindings(ctx, namespace)
			//if err != nil {
			//	return
			//}
			//
			//for _, rolebinding := range rolebindings.Items {
			//	for _, sub := range rolebinding.Subjects {
			//		if sub.Kind == "ServiceAccount" && sub.Name == svcname {
			//			if rolebinding.RoleRef.Kind == "Role" {
			//				rolename := rolebinding.RoleRef.Name
			//				_, err := conn.getRole(ctx, rolename, namespace)
			//				if err != nil {
			//					return
			//				}
			//
			//				//for mainting uniqueness of services
			//				svc := rolebinding.Namespace + "-" + rolebinding.Kind + "-" + rolebinding.Name
			//				services = append(services, &svc)
			//			}
			//			break
			//		}
			//	}
			//}

		}

		hpaList, err := conn.getAllHpas(ctx, namespace)
		if err != nil {
			return
		}

		if len(hpaList.Items) > 0 {
			for _, hpa := range hpaList.Items {
				if hpa.Spec.ScaleTargetRef.APIVersion == dep.APIVersion && hpa.Spec.ScaleTargetRef.Kind == dep.Kind && hpa.Spec.ScaleTargetRef.Name == dep.Name {

				}
			}
		}

		//kubernetes service depecndency findings
		for key, value := range dep.Spec.Template.Labels {
			resp, err := conn.getKubernetesServices(ctx, key, value, namespace)
			if err != nil {
				return
			}

			if len(resp.Items) > 0 {
				for _, kubeSvc := range resp.Items {
					//for mainting uniqueness of services
					svc := kubeSvc.Namespace + "-" + kubeSvc.Kind + "-" + kubeSvc.Name
					services = append(services, &svc)

					fmt.Println(kubeSvc)
				}
			}
		}

		//container dependency finding
		for _, container := range dep.Spec.Template.Spec.Containers {
			//discovering secret and config maps in deployment containers
			for _, env := range container.Env {
				if env.ValueFrom != nil && env.ValueFrom.SecretKeyRef != nil {
					secretname := env.ValueFrom.SecretKeyRef.Name
					_, err := conn.getSecret(ctx, secretname, namespace)
					if err != nil {
						return
					}

				} else if env.ValueFrom != nil && env.ValueFrom.ConfigMapKeyRef != nil {
					configmapname := env.ValueFrom.ConfigMapKeyRef.Name
					_, err := conn.getConfigMap(ctx, configmapname, namespace)
					if err != nil {
						return
					}
				}
			}
		}

		//volume dependency finding
		for _, vol := range dep.Spec.Template.Spec.Volumes {
			if vol.Secret != nil {
				secretname := vol.Secret.SecretName
				_, err := conn.getSecret(ctx, secretname, namespace)
				if err != nil {
					return
				}
			} else if vol.ConfigMap != nil {
				configmapname := vol.ConfigMap.Name
				_, err := conn.getConfigMap(ctx, configmapname, namespace)
				if err != nil {
					return
				}
			} else if vol.PersistentVolumeClaim != nil {
				pvcname := vol.PersistentVolumeClaim.ClaimName
				_, err := conn.getPvc(ctx, pvcname, namespace)
				if err != nil {
					return
				}
			} //else if vol.AWSElasticBlockStore.VolumeID

			if vol.Projected != nil {
				for _, source := range vol.Projected.Sources {
					if source.ConfigMap != nil {
						configmapname := vol.ConfigMap.Name
						_, err := conn.getConfigMap(ctx, configmapname, namespace)
						if err != nil {
							return
						}
					} else if source.Secret != nil {
						secretname := vol.Secret.SecretName
						_, err := conn.getSecret(ctx, secretname, namespace)
						if err != nil {
							return
						}
					}
				}
			}
		}

		cpDeployment, err := convertToCPDeployment(dep)
		if err != nil {
			utils.Error.Println(err.Error())
			return
		}

		bytes, err := json.Marshal(cpDeployment)
		if err != nil {
			utils.Error.Println(err.Error())
			return
		}

		var svctemplate types.ServiceTemplate
		err = json.Unmarshal(bytes, &svctemplate)
		if err != nil {
			utils.Error.Println(err.Error())
			return
		}

		serviceTemplates = append(serviceTemplates, svctemplate)
	}
}

func (conn *GrpcConn) getK8sRbacResources(ctx context.Context, svcname, namespace string, k8sObjectType constants.K8sKind, K8sObj interface{}) ([]types.ServiceTemplate, error) {

	var serviceTemplates []types.ServiceTemplate
	svcaccount, err := conn.getSvcAccount(ctx, svcname, namespace)
	if err != nil {
		return []types.ServiceTemplate{}, err
	}

	var svcAccTemp types.ServiceTemplate
	bytes, err := json.Marshal(svcaccount)
	if err != nil {
		utils.Error.Println(err)
		return []types.ServiceTemplate{}, err
	}
	err = json.Unmarshal(bytes, &svcAccTemp)
	if err != nil {
		utils.Error.Println(err)
		return []types.ServiceTemplate{}, err
	}
	id := strconv.Itoa(rand.Int())
	svcAccTemp.ServiceId = &id

	//creating secrets for service account
	for _, secret := range svcaccount.Secrets {
		if secret.Name != "" {

			secretname := secret.Name
			if secret.Namespace != "" {
				namespace = secret.Namespace
			}
			_, err := conn.getSecret(ctx, secretname, namespace)
			if err != nil {
				utils.Error.Println(err)
				return []types.ServiceTemplate{}, err
			}
		}
	}

	clusterrolebindings, err := conn.getAllClusterRoleBindings(ctx)
	if err != nil {
		utils.Error.Println(err)
		return []types.ServiceTemplate{}, err
	}

	for _, clstrrolebind := range clusterrolebindings.Items {
		for _, sub := range clstrrolebind.Subjects {
			if sub.Kind == "ServiceAccount" && sub.Name == svcname {
				if clstrrolebind.RoleRef.Kind == "ClusterRole" {
					clusterrolename := clstrrolebind.RoleRef.Name
					_, err := conn.getClusterRole(ctx, clusterrolename)
					if err != nil {
						return
					}

					//for mainting uniqueness of services
					svc := clstrrolebind.Namespace + "-" + clstrrolebind.Kind + "-" + clstrrolebind.Name
					services = append(services, &svc)
				}
				break
			}
		}
	}

	rolebindings, err := conn.getAllRoleBindings(ctx, namespace)
	if err != nil {
		return
	}

	for _, rolebinding := range rolebindings.Items {
		for _, sub := range rolebinding.Subjects {
			if sub.Kind == "ServiceAccount" && sub.Name == svcname {
				if rolebinding.RoleRef.Kind == "Role" {
					rolename := rolebinding.RoleRef.Name
					_, err := conn.getRole(ctx, rolename, namespace)
					if err != nil {
						return
					}

					//for mainting uniqueness of services
					svc := rolebinding.Namespace + "-" + rolebinding.Kind + "-" + rolebinding.Name
					services = append(services, &svc)
				}
				break
			}
		}
	}
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

	svc := role.Namespace + "-" + role.Kind + "-" + role.Name
	services = append(services, &svc)

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

	svc := pvc.Namespace + "-" + pvc.Kind + "-" + pvc.Name
	services = append(services, &svc)

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

	if !isAlreadyExist(svcAcc.Namespace, svcAcc.Kind, svcAcc.Name) {
		svc := svcAcc.Namespace + "-" + svcAcc.Kind + "-" + svcAcc.Name
		services = append(services, &svc)
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

	if !isAlreadyExist(scrt.Namespace, scrt.Kind, scrt.Name) {
		svc := scrt.Namespace + "-" + scrt.Kind + "-" + scrt.Name
		services = append(services, &svc)
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

	svc := clusterrole.Namespace + "-" + clusterrole.Kind + "-" + clusterrole.Name
	services = append(services, &svc)

	return clusterrole, nil
}

func isAlreadyExist(namespace, kind, name string) bool {
	for _, val := range services {
		if *val == namespace+"-"+kind+"-"+name {
			return true
		}
	}
	return false
}
