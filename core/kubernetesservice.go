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
	kb "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"strings"
)

func (s *Server) CreateKubernetesService(ctx context.Context, req *pb.KubernetesService) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getKubernetesService(req)
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
func (s *Server) GetKubernetesService(ctx context.Context, req *pb.KubernetesService) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getKubernetesService(req)
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
func (s *Server) DeleteKubernetesService(ctx context.Context, req *pb.KubernetesService) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getKubernetesService(req)
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
func (s *Server) PatchKubernetesService(ctx context.Context, req *pb.KubernetesService) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getKubernetesService(req)
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
func (s *Server) PutKubernetesService(ctx context.Context, req *pb.KubernetesService) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getKubernetesService(req)
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

func getKubernetesService(input *pb.KubernetesService) (*kb.Service, error) {
	var kube = new(kb.Service)
	kube.Kind = "Service"
	kube.APIVersion = "v1"
	kube.Name = input.Name
	if input.Namespace != "" {
		kube.Namespace = input.Namespace
	}
	labels := make(map[string]string)
	labels["app"] = strings.ToLower(input.Name)
	labels["version"] = strings.ToLower(input.Version)
	kube.Labels = labels
	portNames := make(map[string]bool)
	if input.KubeServiceAttributes == nil {
		return nil, errors.New("can not find service attribute object in service")
	}
	for _, port := range input.KubeServiceAttributes.KubePorts {
		spec := *new(kb.ServicePort)
		if port.Name != "" {
			spec.Name = port.Name
			if portNames[spec.Name] {
				return nil, errors.New("port Name Already in this service")
			}

			portNames[spec.Name] = true
		}
		if port.Port < 1 || port.Port > 65535 {
			return nil, errors.New("invalid Port Number Port number should be between 1 and 65535")
		}
		spec.Port = port.Port
		if port.Protocol == string(kb.ProtocolTCP) {
			spec.Protocol = kb.ProtocolTCP
		} else if port.Protocol == string(kb.ProtocolUDP) {
			spec.Protocol = kb.ProtocolUDP
		} else if port.Protocol == string(kb.ProtocolSCTP) {
			spec.Protocol = kb.ProtocolSCTP
		} else {
			return nil, errors.New("invalid protocol supported protocols are TCP, UDP and SCTP")
		}
		if port.TargetPort.PortName != "" {
			spec.TargetPort.StrVal = port.TargetPort.PortName
			spec.TargetPort.Type = intstr.String
		} else if port.Port > 0 && port.Port < 65536 {
			spec.TargetPort.IntVal = port.TargetPort.PortNumber
			spec.TargetPort.Type = intstr.Int
		}
		if port.NodePort >= 30000 && port.NodePort <= 32767 {
			spec.NodePort = port.NodePort
		}

		kube.Spec.Ports = append(kube.Spec.Ports, spec)
	}
	if len(input.KubeServiceAttributes.Selector) > 0 {
		kube.Spec.Selector = make(map[string]string)
	}
	for key, value := range input.KubeServiceAttributes.Selector {
		kube.Spec.Selector[key] = value
	}
	if input.KubeServiceAttributes.Type == string(kb.ServiceTypeClusterIP) {
		kube.Spec.Type = kb.ServiceTypeClusterIP
		if input.KubeServiceAttributes.ClusterIp == "None" {
			kube.Spec.ClusterIP = input.KubeServiceAttributes.ClusterIp
		} else if input.KubeServiceAttributes.ClusterIp != "" {
			kube.Spec.ClusterIP = input.KubeServiceAttributes.ClusterIp
		}
	} else if input.KubeServiceAttributes.Type == string(kb.ServiceTypeNodePort) {
		kube.Spec.Type = kb.ServiceTypeNodePort
		if input.KubeServiceAttributes.ClusterIp == "None" {
			kube.Spec.ClusterIP = input.KubeServiceAttributes.ClusterIp
		} else if input.KubeServiceAttributes.ClusterIp != "" {
			kube.Spec.ClusterIP = input.KubeServiceAttributes.ClusterIp
		}
	} else if input.KubeServiceAttributes.Type == string(kb.ServiceTypeLoadBalancer) {
		kube.Spec.Type = kb.ServiceTypeLoadBalancer
		if input.KubeServiceAttributes.ClusterIp == "None" {
			kube.Spec.ClusterIP = input.KubeServiceAttributes.ClusterIp
		} else if input.KubeServiceAttributes.ClusterIp != "" {
			kube.Spec.ClusterIP = input.KubeServiceAttributes.ClusterIp
		}
	}
	return kube, nil
}
