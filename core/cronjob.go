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
	_ "k8s.io/api/batch/v1"
	v1 "k8s.io/api/batch/v1beta1"
	v2 "k8s.io/api/core/v1"
)

func (s *Server) CreateCronJob(ctx context.Context, req *pb.CronJobService) (*pb.ServiceResponse, error) {
	utils.Info.Println(ctx)
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getCronJobRequestObject(req)
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
		InfraId:       req.InfraId,
		ApplicationId: req.ApplicationId,

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

func (s *Server) GetCronJob(ctx context.Context, req *pb.CronJobService) (*pb.ServiceResponse, error) {
	utils.Info.Println(ctx)
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getCronJobRequestObject(req)
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
		InfraId:       req.InfraId,
		ApplicationId: req.ApplicationId,

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

func (s *Server) DeleteCronJob(ctx context.Context, req *pb.CronJobService) (*pb.ServiceResponse, error) {
	utils.Info.Println(ctx)
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getCronJobRequestObject(req)
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
		InfraId:       req.InfraId,
		ApplicationId: req.ApplicationId,

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

func (s *Server) PutCronJob(ctx context.Context, req *pb.CronJobService) (*pb.ServiceResponse, error) {
	utils.Info.Println(ctx)
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getCronJobRequestObject(req)
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
		InfraId:       req.InfraId,
		ApplicationId: req.ApplicationId,

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

func (s *Server) PatchCronJob(ctx context.Context, req *pb.CronJobService) (*pb.ServiceResponse, error) {
	utils.Info.Println(ctx)
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getCronJobRequestObject(req)
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
		InfraId:       req.InfraId,
		ApplicationId: req.ApplicationId,

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

func getCronJobRequestObject(service *pb.CronJobService) (*v1.CronJob, error) {

	cjob := new(v1.CronJob)
	if service.Name == "" {
		return nil, errors.New("service name not found")
	}
	if service.Namespace == "" {
		cjob.ObjectMeta.Namespace = "default"
	} else {
		cjob.ObjectMeta.Namespace = service.Namespace
	}

	if service.IsDiscovered {
		cjob.Name = service.Name
	} else {
		cjob.Name = service.Name + "-" + service.Version
	}
	cjob.APIVersion = "batch/v1beta1"
	cjob.Kind = constants.CronJob.String() //"CronJob"
	cjob.Labels = make(map[string]string)
	cjob.Labels["keel.sh/policy"] = "force"
	for key, value := range service.CronJobServiceAttribute.Labels {
		cjob.Labels[key] = value
	}
	cjob.Annotations = make(map[string]string)
	cjob.Annotations = service.CronJobServiceAttribute.Annotations

	cjob.Spec.JobTemplate.Spec.Template.Annotations = make(map[string]string)

	cjob.Spec.JobTemplate.Spec.Template.Annotations["sidecar.istio.io/inject"] = "false"

	cjob.Spec.Schedule = service.CronJobServiceAttribute.Schedule

	//cjob.Spec.JobTemplate.Spec.Selector = new(metav1.LabelSelector)
	//cjob.Spec.JobTemplate.Spec.Selector.MatchLabels = make(map[string]string)
	//cjob.Spec.JobTemplate.Spec.Selector.MatchLabels["app"] = service.Name
	//cjob.Spec.JobTemplate.Spec.Selector.MatchLabels["version"] = service.Version
	//
	//if service.CronJobServiceAttribute.LabelSelector != nil {
	//	cjob.Spec.JobTemplate.Spec.Selector.MatchLabels = service.CronJobServiceAttribute.LabelSelector.MatchLabels
	//} else {
	//	cjob.Spec.JobTemplate.Spec.Selector.MatchLabels = service.CronJobServiceAttribute.Labels
	//}
	/*for key, value := range service.CronJobServiceAttribute.JobTemplate.LabelSelector.MatchLabels {
		cjob.Spec.JobTemplate.Spec.Selector.MatchLabels[key] = value
	}*/
	//cjob.Spec.JobTemplate.Labels = make(map[string]string)
	//cjob.Spec.JobTemplate.Labels["app"] = service.Name
	//cjob.Spec.JobTemplate.Labels["version"] = service.Version
	//for key, value := range service.CronJobServiceAttribute.Labels {
	//	cjob.Spec.JobTemplate.Labels[key] = value
	//}

	volumeMountNames1 := make(map[string]bool)
	if service.CronJobServiceAttribute != nil {
		if containersList, volumeMounts, err := getContainers(service.CronJobServiceAttribute.Containers, false); err == nil {
			if len(containersList) > 0 {
				cjob.Spec.JobTemplate.Spec.Template.Spec.Containers = containersList
				volumeMountNames1 = volumeMounts
			} else {
				return nil, errors.New("no container exists")
			}

		} else {
			return nil, err
		}

	}

	if service.CronJobServiceAttribute != nil {
		if containersList, volumeMounts, err := getContainers(service.CronJobServiceAttribute.InitContainers, true); err == nil {
			if len(containersList) > 0 {
				cjob.Spec.JobTemplate.Spec.Template.Spec.InitContainers = containersList
				for k, v := range volumeMounts {
					volumeMountNames1[k] = v
				}
			}

		} else {
			return nil, err
		}
	}

	if dockerSecret, exist := CreateDockerCfgSecret(service.CronJobServiceAttribute.Containers[0], service.Token, service.Namespace); dockerSecret != nil && exist != false {
		cjob.Spec.JobTemplate.Spec.Template.Spec.ImagePullSecrets = []v2.LocalObjectReference{v2.LocalObjectReference{
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
			InfraId:   service.InfraId,
			Service:   raw,
			CompanyId: service.CompanyId,
			Token:     service.Token,
		})

		if err != nil {
			utils.Error.Println(err)
		}

		utils.Info.Println(result.Service)

	}

	if service.CronJobServiceAttribute != nil {
		if volumes, err := getVolumes(service.CronJobServiceAttribute.Volumes, volumeMountNames1); err == nil {
			if len(volumes) > 0 {
				cjob.Spec.JobTemplate.Spec.Template.Spec.Volumes = volumes
			}

		} else {
			return nil, err
		}
	}
	cjob.Spec.JobTemplate.Spec.Template.Spec.RestartPolicy = v2.RestartPolicyOnFailure

	//if service.CronJobServiceAttribute != nil {
	//	if service.CronJobServiceAttribute.Affinity != nil {
	//		if aa, err := getAffinity(service.CronJobServiceAttribute.JobTemplate.Affinity); err != nil {
	//			return nil, err
	//		} else {
	//			cjob.Spec.JobTemplate.Spec.Template.Spec.Affinity = aa
	//		}
	//	}
	//}

	return cjob, nil

}
