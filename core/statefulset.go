package core

import (
	"context"
	"encoding/json"
	"errors"
	"google.golang.org/grpc"
	"istio-service-mesh/constants"
	pb "istio-service-mesh/core/proto"
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
	result, err := pb.NewServiceClient(conn).CreateService(ctx, &pb.ServiceRequest{
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
	result, err := pb.NewServiceClient(conn).DeleteService(ctx, &pb.ServiceRequest{
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
	result, err := pb.NewServiceClient(conn).GetService(ctx, &pb.ServiceRequest{
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
	result, err := pb.NewServiceClient(conn).PutService(ctx, &pb.ServiceRequest{
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
	result, err := pb.NewServiceClient(conn).PatchService(ctx, &pb.ServiceRequest{
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

	statefulSet.Name = service.Name
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
	for key, value := range service.ServiceAttributes.LabelSelector.MatchLabels {
		statefulSet.Spec.Selector.MatchLabels[key] = value
	}
	statefulSet.Spec.Template.Labels = make(map[string]string)
	statefulSet.Spec.Template.Labels["app"] = service.Name
	statefulSet.Spec.Template.Labels["version"] = service.Version
	for key, value := range service.ServiceAttributes.Labels {
		statefulSet.Spec.Template.Labels[key] = value
	}

	statefulSet.Spec.Template.Annotations = make(map[string]string)
	statefulSet.Spec.Template.Annotations["sidecar.istio.io/inject"] = "true"

	if service.ServiceAttributes.Replicas != nil {
		statefulSet.Spec.Replicas = &service.ServiceAttributes.Replicas.Value
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

	statefulSet.Spec.ServiceName = service.ServiceAttributes.ServiceName
	if service.ServiceAttributes.PodManagementPolicy == pb.PodManagementPolicyType_OrderedReady {
		statefulSet.Spec.PodManagementPolicy = v1.OrderedReadyPodManagement
	} else if service.ServiceAttributes.PodManagementPolicy == pb.PodManagementPolicyType_Parallel {
		statefulSet.Spec.PodManagementPolicy = v1.ParallelPodManagement
	}

	var PVCS []v12.PersistentVolumeClaim
	for _, persistentVolumeClaim := range service.ServiceAttributes.VolumeClaimTemplates {
		if pvc, error := getPersistentVolumeClaim(persistentVolumeClaim); error != nil {
			PVCS = append(PVCS, *pvc)
		} else {
			return nil, errors.New("error adding persistent volume claim")
		}
	}

	statefulSet.Spec.VolumeClaimTemplates = PVCS

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
