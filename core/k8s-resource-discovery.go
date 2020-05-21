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
	batchv1 "k8s.io/api/batch/v1beta1"
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
var mutex sync.Mutex

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

	utils.Info.Printf("App discovered successfully for project %s", grpcConn.ProjectId)

	bytes, err := json.Marshal(serviceTemplates)
	if err != nil {
		return &pb.K8SResourceResponse{}, err
	} else if bytes == nil {
		return &pb.K8SResourceResponse{}, errors.New("there are no resources found in the given namespace(s)")
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
		err := conn.resolveRbacDecpendency(ctx, svcname, namespace, jobTemp)
		if err != nil {
			return
		}
	}

	//image pull secrets
	for _, objRef := range job.Spec.Template.Spec.ImagePullSecrets {
		err := conn.resolveSecretDepency(ctx, objRef.Name, namespace, jobTemp)
		if err != nil {
			return
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
				err := conn.resolveHpaDependency(jobTemp, hpa)
				if err != nil {
					return
				}
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
		err = conn.resolveContainerDependency(ctx, kubeSvcList, container, jobTemp, namespace, labels)
		if err != nil {
			return
		}
	}

	//volume dependency finding
	for _, vol := range job.Spec.Template.Spec.Volumes {
		if vol.Secret != nil {
			secretname := vol.Secret.SecretName
			err := conn.resolveSecretDepency(ctx, secretname, namespace, jobTemp)
			if err != nil {
				return
			}
		} else if vol.ConfigMap != nil {
			configmapname := vol.ConfigMap.Name
			err := conn.resolveConfigMapDependency(ctx, configmapname, namespace, jobTemp)
			if err != nil {
				return
			}
		} else if vol.PersistentVolumeClaim != nil {
			pvcname := vol.PersistentVolumeClaim.ClaimName
			err := conn.resolvePvcDependency(ctx, pvcname, namespace, jobTemp)
			if err != nil {
				return
			}
		}

		if vol.Projected != nil {
			for _, source := range vol.Projected.Sources {
				if source.ConfigMap != nil {
					configmapname := vol.ConfigMap.Name
					err := conn.resolveConfigMapDependency(ctx, configmapname, namespace, jobTemp)
					if err != nil {
						return
					}
				} else if source.Secret != nil {
					secretname := vol.Secret.SecretName
					err := conn.resolveSecretDepency(ctx, secretname, namespace, jobTemp)
					if err != nil {
						return
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
		err := conn.resolveRbacDecpendency(ctx, svcname, namespace, cronjobTemp)
		if err != nil {
			return
		}
	}

	//image pull secrets
	for _, objRef := range cronjob.Spec.JobTemplate.Spec.Template.Spec.ImagePullSecrets {
		err := conn.resolveSecretDepency(ctx, objRef.Name, namespace, cronjobTemp)
		if err != nil {
			return
		}
	}

	hpaList, err := conn.getAllHpas(ctx, namespace)
	if err != nil {
		return
	}
	if len(hpaList.Items) > 0 {
		for _, hpa := range hpaList.Items {
			if hpa.Spec.ScaleTargetRef.APIVersion == cronjob.APIVersion && strings.ToLower(hpa.Spec.ScaleTargetRef.Kind) == strings.ToLower(cronjob.Kind) && hpa.Spec.ScaleTargetRef.Name == cronjob.Name {
				err := conn.resolveHpaDependency(cronjobTemp, hpa)
				if err != nil {
					return
				}
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
		err = conn.resolveContainerDependency(ctx, kubeSvcList, container, cronjobTemp, namespace, labels)
		if err != nil {
			return
		}
	}

	//volume dependency finding
	for _, vol := range cronjob.Spec.JobTemplate.Spec.Template.Spec.Volumes {
		if vol.Secret != nil {
			secretname := vol.Secret.SecretName
			err := conn.resolveSecretDepency(ctx, secretname, namespace, cronjobTemp)
			if err != nil {
				return
			}
		} else if vol.ConfigMap != nil {
			configmapname := vol.ConfigMap.Name
			err := conn.resolveConfigMapDependency(ctx, configmapname, namespace, cronjobTemp)
			if err != nil {
				return
			}
		} else if vol.PersistentVolumeClaim != nil {
			pvcname := vol.PersistentVolumeClaim.ClaimName
			err := conn.resolvePvcDependency(ctx, pvcname, namespace, cronjobTemp)
			if err != nil {
				return
			}
		}

		if vol.Projected != nil {
			for _, source := range vol.Projected.Sources {
				if source.ConfigMap != nil {
					configmapname := vol.ConfigMap.Name
					err := conn.resolveConfigMapDependency(ctx, configmapname, namespace, cronjobTemp)
					if err != nil {
						return
					}
				} else if source.Secret != nil {
					secretname := vol.Secret.SecretName
					err := conn.resolveSecretDepency(ctx, secretname, namespace, cronjobTemp)
					if err != nil {
						return
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
		err := conn.resolveRbacDecpendency(ctx, svcname, namespace, daemonsetTemp)
		if err != nil {
			return
		}
	}

	//image pull secrets
	for _, objRef := range daemonset.Spec.Template.Spec.ImagePullSecrets {
		err := conn.resolveSecretDepency(ctx, objRef.Name, namespace, daemonsetTemp)
		if err != nil {
			return
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
		err = conn.resolveContainerDependency(ctx, kubeSvcList, container, daemonsetTemp, namespace, labels)
		if err != nil {
			return
		}
	}

	//volume dependency finding
	for _, vol := range daemonset.Spec.Template.Spec.Volumes {
		if vol.Secret != nil {
			secretname := vol.Secret.SecretName
			err := conn.resolveSecretDepency(ctx, secretname, namespace, daemonsetTemp)
			if err != nil {
				return
			}
		} else if vol.ConfigMap != nil {
			configmapname := vol.ConfigMap.Name
			err := conn.resolveConfigMapDependency(ctx, configmapname, namespace, daemonsetTemp)
			if err != nil {
				return
			}
		} else if vol.PersistentVolumeClaim != nil {
			pvcname := vol.PersistentVolumeClaim.ClaimName
			err := conn.resolvePvcDependency(ctx, pvcname, namespace, daemonsetTemp)
			if err != nil {
				return
			}
		}

		if vol.Projected != nil {
			for _, source := range vol.Projected.Sources {
				if source.ConfigMap != nil {
					configmapname := vol.ConfigMap.Name
					err := conn.resolveConfigMapDependency(ctx, configmapname, namespace, daemonsetTemp)
					if err != nil {
						return
					}
				} else if source.Secret != nil {
					secretname := vol.Secret.SecretName
					err := conn.resolveSecretDepency(ctx, secretname, namespace, daemonsetTemp)
					if err != nil {
						return
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
		err := conn.resolveRbacDecpendency(ctx, svcname, namespace, stsTemp)
		if err != nil {
			return
		}

	}

	//image pull secrets
	for _, objRef := range statefulset.Spec.Template.Spec.ImagePullSecrets {
		err := conn.resolveSecretDepency(ctx, objRef.Name, namespace, stsTemp)
		if err != nil {
			return
		}
	}

	hpaList, err := conn.getAllHpas(ctx, namespace)
	if err != nil {
		return
	}

	if len(hpaList.Items) > 0 {
		for _, hpa := range hpaList.Items {
			if hpa.Spec.ScaleTargetRef.APIVersion == statefulset.APIVersion && strings.ToLower(hpa.Spec.ScaleTargetRef.Kind) == strings.ToLower(statefulset.Kind) && hpa.Spec.ScaleTargetRef.Name == statefulset.Name {
				err := conn.resolveHpaDependency(stsTemp, hpa)
				if err != nil {
					return
				}
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
		err = conn.resolveContainerDependency(ctx, kubeSvcList, container, stsTemp, namespace, labels)
		if err != nil {
			return
		}
	}

	//volume dependency finding
	for _, vol := range statefulset.Spec.Template.Spec.Volumes {
		if vol.Secret != nil {
			secretname := vol.Secret.SecretName
			err := conn.resolveSecretDepency(ctx, secretname, namespace, stsTemp)
			if err != nil {
				return
			}
		} else if vol.ConfigMap != nil {
			configmapname := vol.ConfigMap.Name
			err := conn.resolveConfigMapDependency(ctx, configmapname, namespace, stsTemp)
			if err != nil {
				return
			}
		} else if vol.PersistentVolumeClaim != nil {
			pvcname := vol.PersistentVolumeClaim.ClaimName
			err := conn.resolvePvcDependency(ctx, pvcname, namespace, stsTemp)
			if err != nil {
				return
			}
		}

		if vol.Projected != nil {
			for _, source := range vol.Projected.Sources {
				if source.ConfigMap != nil {
					configmapname := vol.ConfigMap.Name
					err := conn.resolveConfigMapDependency(ctx, configmapname, namespace, stsTemp)
					if err != nil {
						return
					}
				} else if source.Secret != nil {
					secretname := vol.Secret.SecretName
					err := conn.resolveSecretDepency(ctx, secretname, namespace, stsTemp)
					if err != nil {
						return
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
		err := conn.resolveRbacDecpendency(ctx, svcname, namespace, depTemp)
		if err != nil {
			return
		}

	}

	//image pull secrets
	for _, objRef := range dep.Spec.Template.Spec.ImagePullSecrets {
		err := conn.resolveSecretDepency(ctx, objRef.Name, namespace, depTemp)
		if err != nil {
			return
		}
	}

	hpaList, err := conn.getAllHpas(ctx, namespace)
	if err != nil {
		return
	}
	if len(hpaList.Items) > 0 {
		for _, hpa := range hpaList.Items {
			if hpa.Spec.ScaleTargetRef.APIVersion == dep.APIVersion && strings.ToLower(hpa.Spec.ScaleTargetRef.Kind) == strings.ToLower(dep.Kind) && hpa.Spec.ScaleTargetRef.Name == dep.Name {
				err := conn.resolveHpaDependency(depTemp, hpa)
				if err != nil {
					return
				}
			}
		}
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
			err := conn.resolveSecretDepency(ctx, secretname, namespace, depTemp)
			if err != nil {
				return
			}
		} else if vol.ConfigMap != nil {
			configmapname := vol.ConfigMap.Name
			err := conn.resolveConfigMapDependency(ctx, configmapname, namespace, depTemp)
			if err != nil {
				return
			}
		} else if vol.PersistentVolumeClaim != nil {
			pvcname := vol.PersistentVolumeClaim.ClaimName
			err := conn.resolvePvcDependency(ctx, pvcname, namespace, depTemp)
			if err != nil {
				return
			}
		}

		if vol.Projected != nil {
			for _, source := range vol.Projected.Sources {
				if source.ConfigMap != nil {
					configmapname := vol.ConfigMap.Name
					err := conn.resolveConfigMapDependency(ctx, configmapname, namespace, depTemp)
					if err != nil {
						return
					}
				} else if source.Secret != nil {
					secretname := vol.Secret.SecretName
					err := conn.resolveSecretDepency(ctx, secretname, namespace, depTemp)
					if err != nil {
						return
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
			drTemp.Deleted = true
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
					vsTemp.Deleted = true
					serviceTemplates = append(serviceTemplates, vsTemp)

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

							addIngressConfigurations(svcTemp, vsTemp)

							gatewayTemp.AfterServices = append(gatewayTemp.AfterServices, &svcTemp.ServiceId)
							svcTemp.BeforeServices = append(svcTemp.BeforeServices, &gatewayTemp.ServiceId)
							serviceTemplates = append(serviceTemplates, gatewayTemp)
						}
					}
					break
				}
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
					}

					clstrroleTemp, err := conn.getCpConvertedTemplate(clstrrole, clstrrole.Kind)
					if err != nil {
						return nil, err
					}

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
						clstrroleTemp.Deleted = true
						clstrrolebindTemp.Deleted = true
						rbacServiceTemplates = append(rbacServiceTemplates, clstrrolebindTemp)
						rbacServiceTemplates = append(rbacServiceTemplates, clstrroleTemp)
					} else {
						clstrrolebindTemp.BeforeServices = append(clstrrolebindTemp.BeforeServices, &svcAccTemp.ServiceId)
						svcAccTemp.AfterServices = append(svcAccTemp.AfterServices, &clstrrolebindTemp.ServiceId)
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
					}

					roleTemp, err := conn.getCpConvertedTemplate(role, role.Kind)
					if err != nil {
						return nil, err
					}

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
						rolebindTemp.Deleted = true
						roleTemp.Deleted = true

						rbacServiceTemplates = append(rbacServiceTemplates, rolebindTemp)
						rbacServiceTemplates = append(rbacServiceTemplates, roleTemp)
					} else {
						rolebindTemp.BeforeServices = append(rolebindTemp.BeforeServices, &svcAccTemp.ServiceId)
						svcAccTemp.AfterServices = append(svcAccTemp.AfterServices, &rolebindTemp.ServiceId)
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

func (conn *GrpcConn) getPV(ctx context.Context, pvcname, namespace string) (*v2.PersistentVolume, error) {
	response, err := pb.NewK8SResourceClient(conn.Connection).GetK8SResource(ctx, &pb.K8SResourceRequest{
		ProjectId: conn.ProjectId,
		CompanyId: conn.CompanyId,
		Token:     conn.token,
		Command:   "kubectl",
		Args:      []string{"get", "pv", "-o", "json"},
	})
	if err != nil {

		utils.Error.Printf("Error while getting list :%v", getLogData(types.AppDiscoveryLog{
			ProjectId:    conn.ProjectId,
			ServiceType:  string(constants.PersistentVolume),
			ErrorMessage: "error from grpc server :" + err.Error(),
		}))

		return nil, errors.New("error from grpc server :" + err.Error())
	}

	var pvList *v2.PersistentVolumeList
	err = json.Unmarshal(response.Resource, &pvList)
	if err != nil {

		utils.Error.Printf("Error while getting list :%v", getLogData(types.AppDiscoveryLog{
			ProjectId:    conn.ProjectId,
			ServiceType:  string(constants.PersistentVolume),
			ErrorMessage: err.Error(),
		}))

		return nil, err
	}

	for _, pv := range pvList.Items {
		if pv.Spec.ClaimRef != nil && pv.Spec.ClaimRef.Name == pvcname && pv.Spec.ClaimRef.Namespace == namespace {
			return &pv, nil
		}
	}

	return nil, nil
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
		if replicas, ok := template.ServiceAttributes.(map[string]interface{})["replicas"]; ok {
			template.Replicas = int(replicas.(float64))
		}

		svcAttr := template.ServiceAttributes.(map[string]interface{})
		if _, ok := svcAttr["strategy"]; ok {
			svcAttr["strategy"] = struct{}{}
		}

		template.ServiceAttributes = svcAttr
		template.IsDiscovered = true
		template.ServiceId = id
		addVersion(template)
	case constants.CronJob:
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
		template.ServiceId = id
		template.IsDiscovered = true
		addVersion(template)
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
		if replicas, ok := template.ServiceAttributes.(map[string]interface{})["replicas"]; ok {
			template.Replicas = int(replicas.(float64))
		}
		id := strconv.Itoa(rand.Int())
		template.ServiceId = id
		template.IsDiscovered = true
		addVersion(template)
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
		template.IsDiscovered = true
		addVersion(template)
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
		if replicas, ok := template.ServiceAttributes.(map[string]interface{})["replicas"]; ok {
			template.Replicas = int(replicas.(float64))
		}
		id := strconv.Itoa(rand.Int())
		template.ServiceId = id
		template.IsDiscovered = true
		addVersion(template)
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

		svcAttr := template.ServiceAttributes.(map[string]interface{})
		if svcType, ok := svcAttr["type"]; ok {
			if svcType.(string) == "ClusterIP" {
				svcAttr["external_traffic_policy"] = "Local"
			} else {
				svcAttr["external_traffic_policy"] = "Cluster"
			}
		}
		template.ServiceAttributes = svcAttr
		id := strconv.Itoa(rand.Int())
		template.ServiceId = id
		template.IsDiscovered = true
		template.Version = "v1"
		if isAlreadyExist(template.Namespace, template.ServiceSubType, template.Name) {
			template = GetExistingService(template.Namespace, template.ServiceSubType, template.Name)
		}
	case constants.HPA:
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
		template.ServiceId = id
		template.IsDiscovered = true
		template.Version = "v1"
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
		template.IsDiscovered = true
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
		template.IsDiscovered = true
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
		template.IsDiscovered = true
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
		template.IsDiscovered = true
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
		template.IsDiscovered = true
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
		template.IsDiscovered = true
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
		template.IsDiscovered = true
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
		template.IsDiscovered = true
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
		template.IsDiscovered = true
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
		template.IsDiscovered = true
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
		template.IsDiscovered = true
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
		template.IsDiscovered = true
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
		template.IsDiscovered = true
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

		svcAttr := template.ServiceAttributes.(map[string]interface{})
		if resolution, ok := svcAttr["resolution"]; ok {
			if resolution.(string) == "" {
				svcAttr["resolution"] = "NONE"
			}
		} else {
			svcAttr["resolution"] = "NONE"
		}

		if location, ok := svcAttr["location"]; ok {
			if location.(string) == "" {
				location = "MESH_EXTERNAL"
			}
		} else {
			svcAttr["location"] = "MESH_EXTERNAL"
		}

		template.ServiceAttributes = svcAttr

		id := strconv.Itoa(rand.Int())
		template.ServiceId = id
		template.IsDiscovered = true
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
		istioVS.Name = cpKubeService.Name + "-vs"
		istioVS.Namespace = cpKubeService.Namespace
		istioVS.Version = cpKubeService.Version
		istioVS.ServiceType = meshConstants.MeshType
		istioVS.ServiceSubType = meshConstants.VirtualService
		istioVS.ServiceAttributes = new(meshTypes.VSServiceAttribute)
		for _, value := range cpKubeService.ServiceAttributes.Selector {

			istioVS.ServiceAttributes.Hosts = []string{value}
			http := new(meshTypes.Http)
			httpRoute := new(meshTypes.HttpRoute)
			//httpRoute.Weight = 100
			routeRule := new(meshTypes.RouteDestination)
			routeRule.Host = value
			routeRule.Subset = cpKubeService.Version
			routeRule.Port = cpKubeService.ServiceAttributes.Ports[0].Port
			httpRoute.Routes = append(httpRoute.Routes, routeRule)
			http.HttpRoute = append(http.HttpRoute, httpRoute)
			istioVS.ServiceAttributes.Http = append(istioVS.ServiceAttributes.Http, http)
		}

		destRule.ServiceType = meshConstants.MeshType
		destRule.ServiceSubType = meshConstants.DestinationRule
		destRule.Name = cpKubeService.Name + "-dr"
		destRule.Version = cpKubeService.Version
		destRule.Namespace = cpKubeService.Namespace
		for _, value := range cpKubeService.ServiceAttributes.Selector {
			destRule.ServiceAttributes.Host = value
			subset := new(meshTypes.Subset)
			subset.Name = cpKubeService.Version
			lab := make(map[string]string)
			lab["version"] = subset.Name
			subset.Labels = &lab
			destRule.ServiceAttributes.Subsets = append(destRule.ServiceAttributes.Subsets, subset)

		}

		//if cpKubeService.ServiceAttributes.Type == "LoadBalancer"{
		//	gateway := new(meshTypes.GatewayService)
		//	gateway.Name = cpKubeService.Name + "-gtw"
		//	gateway.Namespace = cpKubeService.Namespace
		//	gateway.Version = cpKubeService.Version
		//	gateway.ServiceType = meshConstants.MeshType
		//	gateway.ServiceSubType = meshConstants.Gateway
		//	gateway.ServiceAttributes = new(meshTypes.GatewayServiceAttributes)
		//	gateway.ServiceAttributes.Selectors = make(map[string]string)
		//	gateway.ServiceAttributes.Selectors["istio"]= "ingressgateway"
		//	server := new(meshTypes.Server)
		//	server.Port = new(meshTypes.Port)
		//	server.Port.Number = 80
		//	server.Port.Name = "http"
		//	server.Port.Protocol = meshTypes.Protocols_HTTP
		//	server.Hosts = append(server.Hosts, "*")
		//
		//	var GatewayTemplate *svcTypes.ServiceTemplate
		//	bytes, err = json.Marshal(gateway)
		//	if err != nil {
		//		utils.Error.Println(err)
		//		return nil, err
		//	}
		//
		//	err = json.Unmarshal(bytes, &GatewayTemplate)
		//	if err != nil {
		//		utils.Error.Println(err)
		//		return nil, err
		//	}
		//	id := strconv.Itoa(rand.Int())
		//	GatewayTemplate.ServiceId = id
		//	GatewayTemplate.Deleted = true
		//	svcComponents = append(svcComponents, GatewayTemplate)
		//
		//}

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

		VStemplate.Deleted = true
		//svcTemp.AfterServices = append(svcTemp.AfterServices, &VStemplate.ServiceId)
		//VStemplate.BeforeServices = append(VStemplate.BeforeServices, &svcTemp.ServiceId)

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
		DStemplate.Deleted = true
		//svcTemp.AfterServices = append(svcTemp.AfterServices, &DStemplate.ServiceId)
		//DStemplate.BeforeServices = append(DStemplate.BeforeServices, &svcTemp.ServiceId)

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
						svcTemp.BeforeServices = append(svcTemp.BeforeServices, &istioSvc.ServiceId)
						//svcTemp.Embeds = append(svcTemp.Embeds, istioSvc.ServiceId)
						istioSvc.AfterServices = append(istioSvc.AfterServices, &svcTemp.ServiceId)
						serviceTemplates = append(serviceTemplates, istioSvc)
					}
					//istio components creation
				}

				if !isAlreadyExist(k8serviceTemp.Namespace, k8serviceTemp.ServiceSubType, k8serviceTemp.Name) {
					addKubernetesServiceConfigurations(svcTemp, k8serviceTemp)
					k8serviceTemp.AfterServices = append(k8serviceTemp.AfterServices, &svcTemp.ServiceId)
					svcTemp.BeforeServices = append(svcTemp.BeforeServices, &k8serviceTemp.ServiceId)
					k8serviceTemp.Deleted = true
					for _, key := range svcTemp.BeforeServices {
						svcTemp.Embeds = append(svcTemp.Embeds, *key)
					}
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
						//in case if there are multuple deployments attached with same kubernetes service
						addKubernetesServiceConfigurations(svcTemp, k8serviceTemp)

						k8serviceTemp.AfterServices = append(k8serviceTemp.AfterServices, &svcTemp.ServiceId)
						svcTemp.BeforeServices = append(svcTemp.BeforeServices, &k8serviceTemp.ServiceId)
					}
				}
			}
		}
	}

	//discovering secret and config maps in deployment containers
	for _, env := range container.Env {
		if env.ValueFrom != nil && env.ValueFrom.SecretKeyRef != nil {
			secretname := env.ValueFrom.SecretKeyRef.Name
			err := conn.resolveSecretDepency(ctx, secretname, namespace, svcTemp)
			if err != nil {
				return err
			}
		} else if env.ValueFrom != nil && env.ValueFrom.ConfigMapKeyRef != nil {
			configmapname := env.ValueFrom.ConfigMapKeyRef.Name
			err := conn.resolveConfigMapDependency(ctx, configmapname, namespace, svcTemp)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func addRbacConfigurations(svcTemp *svcTypes.ServiceTemplate, rbacBindingTemp *svcTypes.ServiceTemplate) {
	svcAttr := svcTemp.ServiceAttributes.(map[string]interface{})
	svcAttr["is_rbac_enabled"] = true
	role := make(map[string]interface{})
	role["resource"] = rbacBindingTemp.Name
	role["type"] = rbacBindingTemp.ServiceSubType
	role["service_id"] = rbacBindingTemp.ServiceId
	svcAttr["role"] = role
	svcTemp.ServiceAttributes = svcAttr
}

func addScalingConfigurations(svcTemp *svcTypes.ServiceTemplate, hpaTemp *svcTypes.ServiceTemplate) {
	svcAttr := svcTemp.ServiceAttributes.(map[string]interface{})
	hpaAttr := hpaTemp.ServiceAttributes.(map[string]interface{})

	svcAttr["enable_scaling"] = true

	hpaConfigurations := make(map[string]interface{})
	hpaConfigurations["type"] = "new"
	if minRelicas, ok := hpaAttr["min_replicas"]; ok {
		hpaConfigurations["min_replicas"] = minRelicas
	}
	if maxRelicas, ok := hpaAttr["max_replicas"]; ok {
		hpaConfigurations["max_replicas"] = maxRelicas
	}

	var metrics []interface{}
	metric := make(map[string]interface{})
	metric["target_value_kind"] = "value"
	if cpuValue, ok := hpaAttr["target_cpu_utilization"]; ok {
		metric["target_value"] = cpuValue
	}
	metric["resource_kind"] = "cpu"
	metrics = append(metrics, metric)

	hpaConfigurations["metrics_values"] = metrics

	svcAttr["hpa_configurations"] = hpaConfigurations
	svcTemp.ServiceAttributes = svcAttr

}

func addIngressConfigurations(svcTemp *svcTypes.ServiceTemplate, vsTemp *svcTypes.ServiceTemplate) {
	svcAttr := svcTemp.ServiceAttributes.(map[string]interface{})
	vsAttr := vsTemp.ServiceAttributes.(map[string]interface{})

	svcAttr["enable_external_traffic"] = true

	var URIs []string
	if httpArr, ok := vsAttr["http"].([]interface{}); ok {
		for _, http := range httpArr {
			if httpMatchArr, ok := http.(map[string]interface{})["http_match"].([]interface{}); ok {
				for _, httpMatch := range httpMatchArr {
					if uriInterface, ok := httpMatch.(map[string]interface{})["uri"]; ok {
						uri := uriInterface.(map[string]interface{})
						if uri["type"].(string) == "exact" && uri["value"].(string) != "" {
							fmt.Println(uri["value"].(string))
							URIs = append(URIs, uri["value"].(string))
						}
					}
				}
			}
		}
	}

	svcAttr["uri"] = URIs
	svcTemp.ServiceAttributes = svcAttr
}

func addVolumeConfigurations(svcTemp *svcTypes.ServiceTemplate, pvTemp *svcTypes.ServiceTemplate, pvcTemp *svcTypes.ServiceTemplate) {
	svcAttr := svcTemp.ServiceAttributes.(map[string]interface{})

	if containterArry, ok := svcAttr["containers"].([]interface{}); ok {
		for _, container := range containterArry {
			if volumeMountArr, ok := container.(map[string]interface{})["volume_mounts"].([]interface{}); ok {
				for _, volMount := range volumeMountArr {
					volMount.(map[string]interface{})["service_id"] = pvTemp.ServiceId
					volMount.(map[string]interface{})["service_sub_type"] = pvTemp.ServiceSubType
					volMount.(map[string]interface{})["name"] = pvTemp.Name
					volMount.(map[string]interface{})["persistent_volume_claim_name"] = pvcTemp.Name
				}
			}
		}
	}

	svcTemp.ServiceAttributes = svcAttr
}

func addKubernetesServiceConfigurations(svcTemp *svcTypes.ServiceTemplate, kubeSvcTemp *svcTypes.ServiceTemplate) {
	svcAttr := svcTemp.ServiceAttributes.(map[string]interface{})
	kubeSvcAttr := kubeSvcTemp.ServiceAttributes.(map[string]interface{})

	if svcPortArry, ok := kubeSvcAttr["ports"].([]interface{}); ok {
		for _, v := range svcPortArry {
			if port, ok := v.(map[string]interface{})["target_port"].(map[string]interface{})["port_number"].(float64); ok {

				var portName string
				if val, ok := v.(map[string]interface{})["name"]; !ok {
					portName = "http-" + RandStringBytes(4)
				} else {
					portName = val.(string)
				}

				for index, _ := range svcAttr["containers"].([]interface{}) {
					if c, ok := svcAttr["containers"].([]interface{})[index].(map[string]interface{})["ports"].(map[string]interface{})[""]; ok {

						svcAttr["containers"].([]interface{})[index].(map[string]interface{})["ports"].(map[string]interface{})[portName] = svcAttr["containers"].([]interface{})[index].(map[string]interface{})["ports"].(map[string]interface{})[""]
						delete(svcAttr["containers"].([]interface{})[index].(map[string]interface{})["ports"].(map[string]interface{}), "")

						fmt.Println(c, port)
					} else {
						fmt.Println("port key exists")
					}
				}
			}
		}
	}

	svcTemp.ServiceAttributes = svcAttr
}

func (conn *GrpcConn) resolveHpaDependency(svcTemp *svcTypes.ServiceTemplate, hpa autoscale.HorizontalPodAutoscaler) error {
	hpaTemplate, err := conn.getCpConvertedTemplate(hpa, hpa.Kind)
	if err != nil {
		utils.Error.Println(err)
		return err
	}
	addScalingConfigurations(svcTemp, hpaTemplate)
	hpaTemplate.AfterServices = append(hpaTemplate.AfterServices, &svcTemp.ServiceId)
	svcTemp.BeforeServices = append(svcTemp.BeforeServices, &hpaTemplate.ServiceId)
	svcTemp.Embeds = append(svcTemp.Embeds, hpaTemplate.ServiceId)
	hpaTemplate.Deleted = true
	serviceTemplates = append(serviceTemplates, hpaTemplate)
	return nil
}

func (conn *GrpcConn) resolveSecretDepency(ctx context.Context, secretname, namespace string, svcTemp *svcTypes.ServiceTemplate) error {
	secret, err := conn.getSecret(ctx, secretname, namespace)
	if err != nil {
		return err
	}
	secretTemp, err := conn.getCpConvertedTemplate(secret, secret.Kind)
	if err != nil {
		utils.Error.Println(err)
		return err
	}
	if !isAlreadyExist(secretTemp.Namespace, secretTemp.ServiceSubType, secretTemp.Name) {
		svcTemp.BeforeServices = append(svcTemp.BeforeServices, &secretTemp.ServiceId)
		secretTemp.AfterServices = append(secretTemp.AfterServices, &svcTemp.ServiceId)
		secretTemp.Deleted = true
		svcTemp.Embeds = append(svcTemp.Embeds, secretTemp.ServiceId)
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

	return nil
}

func (conn *GrpcConn) resolveRbacDecpendency(ctx context.Context, svcAccName, namespace string, svcTemp *svcTypes.ServiceTemplate) error {
	mutex.Lock()
	defer mutex.Unlock()
	svcaccount, err := conn.getSvcAccount(ctx, svcAccName, namespace)
	if err != nil {
		return err
	}
	svcAccTemp, err := conn.getCpConvertedTemplate(svcaccount, svcaccount.Kind)
	if err != nil {
		utils.Error.Println(err)
		return err
	}
	if !isAlreadyExist(svcAccTemp.Namespace, svcAccTemp.ServiceSubType, svcAccTemp.Name) {
		rbacDependencies, err := conn.getK8sRbacResources(ctx, namespace, svcaccount, svcAccTemp)
		if err != nil {
			utils.Error.Println(err)
			return err
		}

		for _, rbacTemp := range rbacDependencies {
			if rbacTemp.ServiceSubType == meshConstants.ServiceAccount {
				svcTemp.BeforeServices = append(svcTemp.BeforeServices, &rbacTemp.ServiceId)
				rbacTemp.AfterServices = append(rbacTemp.AfterServices, &svcTemp.ServiceId)
				svcTemp.Embeds = append(svcTemp.Embeds, rbacTemp.ServiceId)
				rbacTemp.Deleted = true
				rbacTemp.IsEmbedded = true
			}

			if rbacTemp.ServiceSubType == meshConstants.Role {
				addRbacConfigurations(svcTemp, rbacTemp)
			}
			serviceTemplates = append(serviceTemplates, rbacTemp)
		}
	} else {
		service := GetExistingService(svcAccTemp.Namespace, svcAccTemp.ServiceSubType, svcAccTemp.Name)
		svcTemp.BeforeServices = append(svcTemp.BeforeServices, &service.ServiceId)
		service.AfterServices = append(service.AfterServices, &svcTemp.ServiceId)
		svcTemp.Embeds = append(svcTemp.Embeds, service.ServiceId)

	}

	return nil
}

func (conn *GrpcConn) resolveConfigMapDependency(ctx context.Context, configmapname, namespace string, svcTemp *svcTypes.ServiceTemplate) error {
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
		configmapTemp.Deleted = true
		svcTemp.Embeds = append(svcTemp.Embeds, configmapTemp.ServiceId)
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

	return nil
}

func (conn *GrpcConn) resolvePvcDependency(ctx context.Context, pvcname, namespace string, svcTemp *svcTypes.ServiceTemplate) error {
	pvc, err := conn.getPvc(ctx, pvcname, namespace)
	if err != nil {
		return err
	}
	pvcTemp, err := conn.getCpConvertedTemplate(pvc, pvc.Kind)
	if err != nil {
		return err
	}
	if !isAlreadyExist(pvcTemp.Namespace, pvcTemp.ServiceSubType, pvcTemp.Name) {
		svcTemp.BeforeServices = append(svcTemp.BeforeServices, &pvcTemp.ServiceId)
		pvcTemp.AfterServices = append(pvcTemp.AfterServices, &svcTemp.ServiceId)
		pvcTemp.Deleted = true
		svcTemp.Embeds = append(svcTemp.Embeds, pvcTemp.ServiceId)
		serviceTemplates = append(serviceTemplates, pvcTemp)
	} else {
		service := GetExistingService(pvcTemp.Namespace, pvcTemp.ServiceSubType, pvcTemp.Name)
		svcTemp.BeforeServices = append(svcTemp.BeforeServices, &service.ServiceId)
		service.AfterServices = append(service.AfterServices, &svcTemp.ServiceId)
	}

	pv, err := conn.getPV(ctx, pvc.Name, pvc.Namespace)
	if err != nil {
		return err
	}
	pvTemp, err := conn.getCpConvertedTemplate(pv, pv.Kind)
	if err != nil {
		return err
	}
	if !isAlreadyExist(pvTemp.Namespace, pvTemp.ServiceSubType, pvTemp.Name) {
		//adding PV and PVC parameters within K8s CP deployment type
		addVolumeConfigurations(svcTemp, pvTemp, pvcTemp)

		svcTemp.BeforeServices = append(svcTemp.BeforeServices, &pvTemp.ServiceId)
		pvTemp.AfterServices = append(pvTemp.AfterServices, &svcTemp.ServiceId)
		pvTemp.Deleted = true
		svcTemp.Embeds = append(svcTemp.Embeds, pvTemp.ServiceId)
		serviceTemplates = append(serviceTemplates, pvTemp)
	} else {
		//adding PV and PVC parameters within K8s CP deployment type
		addVolumeConfigurations(svcTemp, pvTemp, pvcTemp)

		service := GetExistingService(pvTemp.Namespace, pvTemp.ServiceSubType, pvTemp.Name)
		svcTemp.BeforeServices = append(svcTemp.BeforeServices, &service.ServiceId)
		service.AfterServices = append(service.AfterServices, &svcTemp.ServiceId)
	}

	if pvc.Spec.StorageClassName != nil && pv.Spec.StorageClassName != "" && *pvc.Spec.StorageClassName == pv.Spec.StorageClassName {
		storageClassName := *pvc.Spec.StorageClassName
		storageClass, err := conn.getStorageClass(ctx, storageClassName, namespace)
		if err != nil {
			return err
		}
		storageClassTemp, err := conn.getCpConvertedTemplate(storageClass, storageClass.Kind)
		if err != nil {
			return err
		}
		if !isAlreadyExist(storageClassTemp.Namespace, storageClassTemp.ServiceSubType, storageClassTemp.Name) {
			//attaching storage class with PVC
			pvcTemp.BeforeServices = append(pvcTemp.BeforeServices, &storageClassTemp.ServiceId)
			storageClassTemp.AfterServices = append(storageClassTemp.AfterServices, &pvcTemp.ServiceId)
			storageClassTemp.Deleted = true
			pvcTemp.Embeds = append(pvcTemp.Embeds, storageClassTemp.ServiceId)
			pvcTemp.IsEmbedded = true

			//attaching storage clase with PV
			pvTemp.BeforeServices = append(pvTemp.BeforeServices, &storageClassTemp.ServiceId)
			storageClassTemp.AfterServices = append(storageClassTemp.AfterServices, &pvTemp.ServiceId)

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

	return nil
}

func addVersion(svcTemp *svcTypes.ServiceTemplate) {
	strArr := strings.Split(svcTemp.Name, "-")
	if len(strArr) > 1 {
		if len(strArr) == 2 {
			svcTemp.Version = strArr[1]
		} else {
			svcTemp.Version = strArr[len(strArr)-1]
		}
	} else {
		svcTemp.Version = "v1"
	}
}

func RandStringBytes(n int) string {
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
