package core

import (
	pb1 "bitbucket.org/cloudplex-devs/kubernetes-services-deployment/core/proto"
	pb "bitbucket.org/cloudplex-devs/microservices-mesh-engine/core/services/proto"
	"context"
	"encoding/json"
	"errors"
	"google.golang.org/grpc"
	"istio-service-mesh/constants"
	"istio-service-mesh/utils"
	v1 "k8s.io/api/apps/v1"
	v12 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (s *Server) CreateStatefulSet(ctx context.Context, req *pb.StatefulSetService) (*pb.ServiceResponse, error) {
	utils.Info.Println(ctx)
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getStatefulSetRequestObject(req)
	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp, err)
		return serviceResp, err
	}
	conn, err := grpc.DialContext(ctx, constants.K8sEngineGRPCURL, grpc.WithInsecure())
	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp, err)
		return serviceResp, err
	}
	defer conn.Close()

	raw, err := json.Marshal(ksdRequest)
	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp, err)
		return serviceResp, err
	}
	result, err := pb1.NewServiceClient(conn).CreateService(ctx, &pb1.ServiceRequest{
		ProjectId: req.ProjectId,
		Service:   raw,
		CompanyId: req.CompanyId,
		Token:     req.Token,
	})

	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp, err)
		return serviceResp, err
	}
	utils.Info.Println(result.Service)
	serviceResp.Status.Status = "successful"
	serviceResp.Status.StatusIndividual = append(serviceResp.Status.StatusIndividual, "successful")

	return serviceResp, nil
}

func (s *Server) DeleteStatefulSet(ctx context.Context, req *pb.StatefulSetService) (*pb.ServiceResponse, error) {
	utils.Info.Println(ctx)
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getStatefulSetRequestObject(req)
	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp, err)
		return serviceResp, err
	}
	conn, err := grpc.DialContext(ctx, constants.K8sEngineGRPCURL, grpc.WithInsecure())
	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp, err)
		return serviceResp, err
	}
	defer conn.Close()

	raw, err := json.Marshal(ksdRequest)
	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp, err)
		return serviceResp, err
	}
	result, err := pb1.NewServiceClient(conn).DeleteService(ctx, &pb1.ServiceRequest{
		ProjectId: req.ProjectId,
		Service:   raw,
		CompanyId: req.CompanyId,
		Token:     req.Token,
	})

	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp, err)
		return serviceResp, err
	}
	utils.Info.Println(result.Service)
	serviceResp.Status.Status = "successful"
	serviceResp.Status.StatusIndividual = append(serviceResp.Status.StatusIndividual, "successful")

	return serviceResp, nil
}

func (s *Server) GetStatefulSet(ctx context.Context, req *pb.StatefulSetService) (*pb.ServiceResponse, error) {
	utils.Info.Println(ctx)
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getStatefulSetRequestObject(req)
	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp, err)
		return serviceResp, err
	}
	conn, err := grpc.DialContext(ctx, constants.K8sEngineGRPCURL, grpc.WithInsecure())
	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp, err)
		return serviceResp, err
	}
	defer conn.Close()

	raw, err := json.Marshal(ksdRequest)
	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp, err)
		return serviceResp, err
	}
	result, err := pb1.NewServiceClient(conn).GetService(ctx, &pb1.ServiceRequest{
		ProjectId: req.ProjectId,
		Service:   raw,
		CompanyId: req.CompanyId,
		Token:     req.Token,
	})

	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp, err)
		return serviceResp, err
	}
	utils.Info.Println(result.Service)
	serviceResp.Status.Status = "successful"
	serviceResp.Status.StatusIndividual = append(serviceResp.Status.StatusIndividual, "successful")

	return serviceResp, nil
}

func (s *Server) PutStatefulSet(ctx context.Context, req *pb.StatefulSetService) (*pb.ServiceResponse, error) {
	utils.Info.Println(ctx)
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getStatefulSetRequestObject(req)
	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp, err)
		return serviceResp, err
	}
	conn, err := grpc.DialContext(ctx, constants.K8sEngineGRPCURL, grpc.WithInsecure())
	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp, err)
		return serviceResp, err
	}
	defer conn.Close()

	raw, err := json.Marshal(ksdRequest)
	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp, err)
		return serviceResp, err
	}
	result, err := pb1.NewServiceClient(conn).PutService(ctx, &pb1.ServiceRequest{
		ProjectId: req.ProjectId,
		Service:   raw,
		CompanyId: req.CompanyId,
		Token:     req.Token,
	})

	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp, err)
		return serviceResp, err
	}
	utils.Info.Println(result.Service)
	serviceResp.Status.Status = "successful"
	serviceResp.Status.StatusIndividual = append(serviceResp.Status.StatusIndividual, "successful")

	return serviceResp, nil
}

func (s *Server) PatchStatefulSet(ctx context.Context, req *pb.StatefulSetService) (*pb.ServiceResponse, error) {
	utils.Info.Println(ctx)
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getStatefulSetRequestObject(req)
	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp, err)
		return serviceResp, err
	}
	conn, err := grpc.DialContext(ctx, constants.K8sEngineGRPCURL, grpc.WithInsecure())
	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp, err)
		return serviceResp, err
	}
	defer conn.Close()

	raw, err := json.Marshal(ksdRequest)
	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp, err)
		return serviceResp, err
	}
	result, err := pb1.NewServiceClient(conn).PatchService(ctx, &pb1.ServiceRequest{
		ProjectId: req.ProjectId,
		Service:   raw,
		CompanyId: req.CompanyId,
		Token:     req.Token,
	})

	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp, err)
		return serviceResp, err
	}
	utils.Info.Println(result.Service)
	serviceResp.Status.Status = "successful"
	serviceResp.Status.StatusIndividual = append(serviceResp.Status.StatusIndividual, "successful")

	return serviceResp, nil
}

func getStatefulSetRequestObject(service *pb.StatefulSetService) (*v1.StatefulSet, error) {
	statefulSet := new(v1.StatefulSet)
	if service.Name == "" {
		return nil, errors.New("service name not found")
	}
	statefulSet.APIVersion = "apps/v1"
	statefulSet.Kind = "StatefulSet"

	if service.Namespace == "" {
		statefulSet.ObjectMeta.Namespace = "default"
	} else {
		statefulSet.ObjectMeta.Namespace = service.Namespace
	}

	statefulSet.Name = service.Name + "-" + service.Version
	statefulSet.Labels = make(map[string]string)
	statefulSet.Labels["keel.sh/policy"] = "force"
	for key, value := range service.ServiceAttributes.Labels {
		statefulSet.Labels[key] = value
	}

	statefulSet.Annotations = make(map[string]string)
	statefulSet.Annotations = service.ServiceAttributes.Annotations

	statefulSet.Spec.Selector = new(metav1.LabelSelector)
	statefulSet.Spec.Selector.MatchLabels = make(map[string]string)
	statefulSet.Spec.Selector.MatchLabels["app"] = service.Name
	statefulSet.Spec.Selector.MatchLabels["version"] = service.Version

	if service.ServiceAttributes.LabelSelector != nil {
		statefulSet.Spec.Selector.MatchLabels = service.ServiceAttributes.LabelSelector.MatchLabels
	} else {
		statefulSet.Spec.Selector.MatchLabels = service.ServiceAttributes.Labels
	}
	/*for key, value := range service.ServiceAttributes.LabelSelector.MatchLabels {
		statefulSet.Spec.Selector.MatchLabels[key] = value
	}*/
	statefulSet.Spec.Template.Labels = make(map[string]string)
	statefulSet.Spec.Template.Labels["app"] = service.Name
	statefulSet.Spec.Template.Labels["version"] = service.Version
	for key, value := range service.ServiceAttributes.Labels {
		statefulSet.Spec.Template.Labels[key] = value
	}

	statefulSet.Spec.Template.Annotations = make(map[string]string)
	statefulSet.Spec.Template.Annotations["sidecar.istio.io/inject"] = "true"

	if service.ServiceAttributes.Replicas > 0 {
		statefulSet.Spec.Replicas = &service.ServiceAttributes.Replicas
	}

	if service.ServiceAttributes.RevisionHistoryLimit != nil {
		statefulSet.Spec.RevisionHistoryLimit = &service.ServiceAttributes.RevisionHistoryLimit.Value
	}

	if service.ServiceAttributes.TerminationGracePeriodSeconds != nil {
		statefulSet.Spec.Template.Spec.TerminationGracePeriodSeconds = &service.ServiceAttributes.TerminationGracePeriodSeconds.Value
	}

	if service.ServiceAttributes.UpdateStrategy != nil {
		if service.ServiceAttributes.UpdateStrategy.Type == pb.StatefulSetUpdateStrategyType_StatefulSetOnDelete {
			statefulSet.Spec.UpdateStrategy.Type = v1.OnDeleteStatefulSetStrategyType
		} else if service.ServiceAttributes.UpdateStrategy.Type == pb.StatefulSetUpdateStrategyType_StatefulSetRollingUpdate {
			statefulSet.Spec.UpdateStrategy.Type = v1.RollingUpdateStatefulSetStrategyType

			if service.ServiceAttributes.UpdateStrategy.RollingUpdate != nil {
				if service.ServiceAttributes.UpdateStrategy.RollingUpdate.Partition != nil {
					statefulSet.Spec.UpdateStrategy.RollingUpdate = &v1.RollingUpdateStatefulSetStrategy{
						Partition: &service.ServiceAttributes.UpdateStrategy.RollingUpdate.Partition.Value,
					}
				}
			}
		}
	}

	if dockerSecret, exist := CreateDockerCfgSecret(service.ServiceAttributes.Containers[0], service.Token, service.Namespace); dockerSecret != nil && exist != false {
		statefulSet.Spec.Template.Spec.ImagePullSecrets = []v12.LocalObjectReference{v12.LocalObjectReference{
			Name: dockerSecret.Name,
		}}
		var ctx context.Context
		conn, err := grpc.DialContext(ctx, constants.K8sEngineGRPCURL, grpc.WithInsecure())
		if err != nil {
			utils.Error.Println(err)
		}

		defer conn.Close()

		raw, err := json.Marshal(dockerSecret)
		if err != nil {
			utils.Error.Println(err)
		}
		result, err := pb1.NewServiceClient(conn).CreateService(ctx, &pb1.ServiceRequest{
			ProjectId: service.ProjectId,
			Service:   raw,
			CompanyId: service.CompanyId,
			Token:     service.Token,
		})

		if err != nil {
			utils.Error.Println(err)
		}

		utils.Info.Println(result.Service)

	}

	statefulSet.Spec.ServiceName = service.ServiceAttributes.ServiceName
	if service.ServiceAttributes.PodManagementPolicy == pb.PodManagementPolicyType_OrderedReady {
		statefulSet.Spec.PodManagementPolicy = v1.OrderedReadyPodManagement
	} else if service.ServiceAttributes.PodManagementPolicy == pb.PodManagementPolicyType_Parallel {
		statefulSet.Spec.PodManagementPolicy = v1.ParallelPodManagement
	}

	volumeMountNames1 := make(map[string]bool)
	if containersList, volumeMounts, err := getContainers(service.ServiceAttributes.Containers); err == nil {
		if len(containersList) > 0 {
			statefulSet.Spec.Template.Spec.Containers = containersList
			volumeMountNames1 = volumeMounts
		} else {
			return nil, errors.New("no container exists")
		}

	} else {
		return nil, err
	}

	if containersList, volumeMounts, err := getContainers(service.ServiceAttributes.InitContainers); err == nil {
		if len(containersList) > 0 {
			statefulSet.Spec.Template.Spec.InitContainers = containersList
			for k, v := range volumeMounts {
				volumeMountNames1[k] = v
			}
		}

	} else {
		return nil, err
	}

	for _, persistentVolumeClaim := range service.ServiceAttributes.VolumeClaimTemplates {
		if pvc, error := getPersistentVolumeClaim(persistentVolumeClaim); error == nil {

			if !volumeMountNames1[pvc.Name] {
				continue
			}
			statefulSet.Spec.VolumeClaimTemplates = append(statefulSet.Spec.VolumeClaimTemplates, v12.PersistentVolumeClaim{
				Spec:       pvc.Spec,
				ObjectMeta: metav1.ObjectMeta{Name: pvc.Name},
			})
			volumeMountNames1[pvc.Name] = false
		} else {
			return nil, errors.New("error adding persistent volume claim")
		}
	}

	for key, _ := range volumeMountNames1 {
		if volumeMountNames1[key] == true {
			return nil, errors.New("volume does not exist")
		}
	}

	if volumes, err := getVolumes(service.ServiceAttributes.Volumes, volumeMountNames1); err == nil {
		if len(volumes) > 0 {
			statefulSet.Spec.Template.Spec.Volumes = volumes
		}

	} else {
		return nil, err
	}

	if service.ServiceAttributes.Affinity != nil {
		if aa, err := getAffinity(service.ServiceAttributes.Affinity); err != nil {
			return nil, err
		} else {
			statefulSet.Spec.Template.Spec.Affinity = aa
		}
	}

	return statefulSet, nil
}
