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
	autoscaler "k8s.io/api/autoscaling/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"strings"
)

func (s *Server) CreateHPA(ctx context.Context, req *pb.HPA) (*pb.ServiceResponse, error) {
	utils.Info.Println(ctx)
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getHpa(req)

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
func (s *Server) GetHPA(ctx context.Context, req *pb.HPA) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getHpa(req)

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
func (s *Server) DeleteHPA(ctx context.Context, req *pb.HPA) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getHpa(req)

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
func (s *Server) PatchHPA(ctx context.Context, req *pb.HPA) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getHpa(req)

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
func (s *Server) PutHPA(ctx context.Context, req *pb.HPA) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getHpa(req)

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

func getHpa(input *pb.HPA) (*autoscaler.HorizontalPodAutoscaler, error) {
	var hpaSvc = new(autoscaler.HorizontalPodAutoscaler)
	labels := make(map[string]string)
	labels["app"] = strings.ToLower(input.Name)
	labels["version"] = strings.ToLower(input.Version)
	hpaSvc.Kind = constants.HPA.String() //"HorizontalPodAutoscaler"
	hpaSvc.APIVersion = "autoscaling/v1"
	if input.Name == "" {
		return &autoscaler.HorizontalPodAutoscaler{}, errors.New("hpa name must not be empty")
	}
	hpaSvc.Name = input.Name
	hpaSvc.Labels = labels
	if hpaSvc.Namespace == "" {
		hpaSvc.Namespace = input.Namespace
	}
	if input.ServiceAttributes.CrossObjectVersion.Type == "Deployment" {
		input.ServiceAttributes.CrossObjectVersion.Version = "apps/v1"
	} else if input.ServiceAttributes.CrossObjectVersion.Type == "CronJob" {
		input.ServiceAttributes.CrossObjectVersion.Version = "batch/v1beta1"
	} else if input.ServiceAttributes.CrossObjectVersion.Type == "StatefulSet" {
		input.ServiceAttributes.CrossObjectVersion.Version = "batch/v1beta1"
	} else if input.ServiceAttributes.CrossObjectVersion.Type == "" {
		return &autoscaler.HorizontalPodAutoscaler{}, errors.New("target object type must not be empty")
	}

	if input.ServiceAttributes.CrossObjectVersion.Name == "" {
		return &autoscaler.HorizontalPodAutoscaler{}, errors.New("target object name must not be empty")
	}

	targetOjb := autoscaler.CrossVersionObjectReference{
		Kind:       input.ServiceAttributes.CrossObjectVersion.Type,
		Name:       input.ServiceAttributes.CrossObjectVersion.Name,
		APIVersion: input.ServiceAttributes.CrossObjectVersion.Version,
	}

	hpaSvc.Spec.ScaleTargetRef = targetOjb
	if input.ServiceAttributes.MaxReplicas == 0 {
		return &autoscaler.HorizontalPodAutoscaler{}, errors.New("max replica value can not be zero")
	}
	hpaSvc.Spec.MaxReplicas = int32(input.ServiceAttributes.MaxReplicas)

	if input.ServiceAttributes.MinReplicas == 0 {
		input.ServiceAttributes.MinReplicas = 1
	}
	minreplicas := int32(input.ServiceAttributes.MinReplicas)
	hpaSvc.Spec.MinReplicas = &minreplicas

	if input.ServiceAttributes.TargetCpuUtilization != 0 {
		hpaSvc.Spec.TargetCPUUtilizationPercentage = &input.ServiceAttributes.TargetCpuUtilization
	}
	/*var metrics []v2beta2.MetricSpec
	for _, metric := range input.ServiceAttributes.MetricValues {
		met := v2beta2.MetricSpec{
			Type: v2beta2.ResourceMetricSourceType,
		}
		target := v2beta2.MetricTarget{}
		if metric.TargetValueKind == "Value" {
			target.Type = v2beta2.ValueMetricType
			if value, error := resource.ParseQuantity(metric.TargetValue); error != nil {
				return nil, errors.New("error setting target value")
			} else {
				target.Value = &value
			}

		} else if metric.TargetValueKind == "Utilization" {
			target.Type = v2beta2.UtilizationMetricType
			value, _ := strconv.Atoi(metric.TargetValue)
			ptrval := int32(value)
			target.AverageUtilization = &ptrval
		} else if metric.TargetValueKind == "Average" {
			target.Type = v2beta2.AverageValueMetricType
			if value, error := resource.ParseQuantity(metric.TargetValue); error != nil {
				return nil, errors.New("error setting target value")
			} else {
				target.AverageValue = &value
			}
		}

		resource := v2beta2.ResourceMetricSource{}
		if metric.ResourceKind == "cpu" {
			resource.Name = v1.ResourceCPU
		} else if metric.ResourceKind == "memory" {
			resource.Name = v1.ResourceMemory
		} else if metric.ResourceKind == "storage" {
			resource.Name = v1.ResourceEphemeralStorage
		}

		resource.Target = target

		met.Resource = &resource
		metrics = append(metrics, met)
	}
	hpaSvc.Spec.Metrics = metrics*/
	/*
		var metrics []v2beta2.MetricSpec
			for _, metric := range input.ServiceAttributes.MetricValues {
				met := v2beta2.MetricSpec{
					Type: v2beta2.ResourceMetricSourceType,
				}
				target := v2beta2.MetricTarget{}
				if metric.TargetValueKind == "value" {
					target.Type = v2beta2.ValueMetricType
					value, _ := strconv.Atoi(metric.TargetValue)
					target.Value = resource.NewScaledQuantity(int64(value), ScaleUnit(metric.TargetValueUnit))

				} else if metric.TargetValueKind == "utilization" {
					target.Type = v2beta2.UtilizationMetricType
					value, _ := strconv.Atoi(metric.TargetValue)
					ptrval := int32(value)
					target.AverageUtilization = &ptrval
				} else if metric.TargetValueKind == "average" {
					target.Type = v2beta2.AverageValueMetricType
					value, _ := strconv.Atoi(metric.TargetValue)
					target.AverageValue = resource.NewScaledQuantity(int64(value), ScaleUnit(metric.TargetValueUnit))
				}

				resource := v2beta2.ResourceMetricSource{}
				if metric.ResourceKind == "cpu" {
					resource.Name = v1.ResourceCPU
				} else if metric.ResourceKind == "memory" {
					resource.Name = v1.ResourceMemory
				} else if metric.ResourceKind == "storage" {
					resource.Name = v1.ResourceEphemeralStorage
				}

				resource.Target = target

				met.Resource = &resource
				metrics = append(metrics, met)
			}
			hpaSvc.Spec.Metrics = metrics

	*/
	return hpaSvc, nil
}

func ScaleUnit(unit string) resource.Scale {

	if unit == "nano" {
		return resource.Nano
	} else if unit == "micro" {
		return resource.Micro
	} else if unit == "milli" {
		return resource.Milli
	} else if unit == "kilo" {
		return resource.Kilo
	} else if unit == "mega" {
		return resource.Mega
	} else if unit == "giga" {
		return resource.Giga
	} else if unit == "tera" {
		return resource.Tera
	} else if unit == "peta" {
		return resource.Peta
	} else {
		return resource.Exa
	}

}
