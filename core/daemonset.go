package core

import (
	"bitbucket.org/cloudplex-devs/istio-service-mesh/constants"
	"bitbucket.org/cloudplex-devs/istio-service-mesh/utils"
	pb1 "bitbucket.org/cloudplex-devs/kubernetes-services-deployment/core/proto"
	pb "bitbucket.org/cloudplex-devs/microservices-mesh-engine/core/services/proto"
	"context"
	"encoding/json"
	"errors"
	"google.golang.org/grpc"
	v1 "k8s.io/api/apps/v1"
	v2 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func (s *Server) CreateDaemonSet(ctx context.Context, req *pb.DaemonSetService) (*pb.ServiceResponse, error) {
	utils.Info.Println(ctx)
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getDaemonSetRequestObject(req)
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

func (s *Server) GetDaemonSet(ctx context.Context, req *pb.DaemonSetService) (*pb.ServiceResponse, error) {
	utils.Info.Println(ctx)
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getDaemonSetRequestObject(req)
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
	if len(result.PodErrors) > 0 {
		getPodErrors(serviceResp, result.PodErrors)
		return serviceResp, nil
	}
	utils.Info.Println(result.Service)
	serviceResp.Status.Status = "successful"
	serviceResp.Status.StatusIndividual = append(serviceResp.Status.StatusIndividual, "successful")

	return serviceResp, nil
}

func (s *Server) DeleteDaemonSet(ctx context.Context, req *pb.DaemonSetService) (*pb.ServiceResponse, error) {
	utils.Info.Println(ctx)
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getDaemonSetRequestObject(req)
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

func (s *Server) PatchDaemonSet(ctx context.Context, req *pb.DaemonSetService) (*pb.ServiceResponse, error) {
	utils.Info.Println(ctx)
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getDaemonSetRequestObject(req)
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

func (s *Server) PutDaemonSet(ctx context.Context, req *pb.DaemonSetService) (*pb.ServiceResponse, error) {
	utils.Info.Println(ctx)
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getDaemonSetRequestObject(req)
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

func getDaemonSetRequestObject(service *pb.DaemonSetService) (*v1.DaemonSet, error) {
	daemonSet := new(v1.DaemonSet)
	if service.Name == "" {
		return nil, errors.New("service name not found")
	}

	daemonSet.APIVersion = "apps/v1"
	daemonSet.Kind = constants.DaemonSet.String() //"DaemonSet"

	if service.Namespace == "" {
		daemonSet.ObjectMeta.Namespace = "default"
	} else {
		daemonSet.ObjectMeta.Namespace = service.Namespace
	}

	if service.IsDiscovered {
		daemonSet.Name = service.Name
	} else {
		daemonSet.Name = service.Name + "-" + service.Version
	}
	daemonSet.Labels = make(map[string]string)
	daemonSet.Labels["keel.sh/policy"] = "force"
	for key, value := range service.ServiceAttributes.Labels {
		daemonSet.Labels[key] = value
	}

	daemonSet.Annotations = make(map[string]string)
	daemonSet.Annotations = service.ServiceAttributes.Annotations

	daemonSet.Spec.Selector = new(metav1.LabelSelector)
	daemonSet.Spec.Selector.MatchLabels = make(map[string]string)
	daemonSet.Spec.Selector.MatchLabels["app"] = service.Name
	if service.Version != "" {
		daemonSet.Spec.Selector.MatchLabels["version"] = service.Version
	}
	if service.ServiceAttributes.LabelSelector != nil {
		daemonSet.Spec.Selector.MatchLabels = service.ServiceAttributes.LabelSelector.MatchLabels
	} else {
		daemonSet.Spec.Selector.MatchLabels = service.ServiceAttributes.Labels
	}

	/*for key, value := range service.ServiceAttributes.LabelSelector.MatchLabels {
		daemonSet.Spec.Selector.MatchLabels[key] = value
	}*/
	daemonSet.Spec.Template.Labels = make(map[string]string)
	daemonSet.Spec.Template.Labels["app"] = service.Name
	if service.Version != "" {
		daemonSet.Spec.Template.Labels["version"] = service.Version
	}

	for key, value := range service.ServiceAttributes.Labels {
		daemonSet.Spec.Template.Labels[key] = value
	}

	daemonSet.Spec.Template.Annotations = make(map[string]string)
	daemonSet.Spec.Template.Annotations["sidecar.istio.io/inject"] = "true"
	daemonSet.Spec.Template.Spec.NodeSelector = make(map[string]string)
	daemonSet.Spec.Template.Spec.NodeSelector = service.ServiceAttributes.NodeSelector

	if service.ServiceAttributes.MinReadySeconds != 0 {
		daemonSet.Spec.MinReadySeconds = service.ServiceAttributes.MinReadySeconds
	}

	if service.ServiceAttributes.TerminationGracePeriodSeconds != nil {
		daemonSet.Spec.Template.Spec.TerminationGracePeriodSeconds = &service.ServiceAttributes.TerminationGracePeriodSeconds.Value
	}

	if service.ServiceAttributes.ActiveDeadlineSeconds != nil {
		daemonSet.Spec.Template.Spec.ActiveDeadlineSeconds = &service.ServiceAttributes.ActiveDeadlineSeconds.Value
	}

	if service.ServiceAttributes.RevisionHistoryLimit != nil {
		daemonSet.Spec.RevisionHistoryLimit = &service.ServiceAttributes.RevisionHistoryLimit.Value
	}

	if service.ServiceAttributes.UpdateStrategy != nil {
		if service.ServiceAttributes.UpdateStrategy.Type == pb.DaemonSetUpdateStrategyType_DaemonSetOnDelete {
			daemonSet.Spec.UpdateStrategy.Type = v1.OnDeleteDaemonSetStrategyType
		} else if service.ServiceAttributes.UpdateStrategy.Type == pb.DaemonSetUpdateStrategyType_DaemonSetRollingUpdate {
			daemonSet.Spec.UpdateStrategy.Type = v1.RollingUpdateDaemonSetStrategyType

			if service.ServiceAttributes.UpdateStrategy.RollingUpdate != nil {
				if service.ServiceAttributes.UpdateStrategy.RollingUpdate.GetIntVal() != 0 {
					daemonSet.Spec.UpdateStrategy.RollingUpdate = &v1.RollingUpdateDaemonSet{
						MaxUnavailable: &intstr.IntOrString{
							IntVal: service.ServiceAttributes.UpdateStrategy.RollingUpdate.GetIntVal(),
						},
					}
				} else if service.ServiceAttributes.UpdateStrategy.RollingUpdate.GetStrVal() != "" {
					daemonSet.Spec.UpdateStrategy.RollingUpdate = &v1.RollingUpdateDaemonSet{
						MaxUnavailable: &intstr.IntOrString{
							StrVal: service.ServiceAttributes.UpdateStrategy.RollingUpdate.GetStrVal(),
						},
					}
				}
			}
		}
	}

	volumeMountNames1 := make(map[string]bool)
	if containersList, volumeMounts, err := getContainers(service.ServiceAttributes.Containers, false); err == nil {
		if len(containersList) > 0 {
			daemonSet.Spec.Template.Spec.Containers = containersList
			volumeMountNames1 = volumeMounts
		} else {
			return nil, errors.New("no container exists")
		}

	} else {
		return nil, err
	}

	if containersList, volumeMounts, err := getContainers(service.ServiceAttributes.InitContainers, true); err == nil {
		if len(containersList) > 0 {
			daemonSet.Spec.Template.Spec.InitContainers = containersList
			for k, v := range volumeMounts {
				volumeMountNames1[k] = v
			}
		}

	} else {
		return nil, err
	}

	if volumes, err := getVolumes(service.ServiceAttributes.Volumes, volumeMountNames1); err == nil {
		if len(volumes) > 0 {
			daemonSet.Spec.Template.Spec.Volumes = volumes
		}

	} else {
		return nil, err
	}

	if service.ServiceAttributes.Affinity != nil {
		if aa, err := getAffinity(service.ServiceAttributes.Affinity); err != nil {
			return nil, err
		} else {
			daemonSet.Spec.Template.Spec.Affinity = aa
		}
	}

	if dockerSecret, exist := CreateDockerCfgSecret(service.ServiceAttributes.Containers[0], service.Token, service.Namespace); dockerSecret != nil && exist != false {
		daemonSet.Spec.Template.Spec.ImagePullSecrets = []v2.LocalObjectReference{v2.LocalObjectReference{
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

	return daemonSet, nil
}
