package core

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc"
	"istio-service-mesh/constants"
	pb "istio-service-mesh/core/proto"
	"istio-service-mesh/utils"
	cor "k8s.io/api/core/v1"
	net "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func (s *Server) CreateNetworkPolicy(ctx context.Context, req *pb.NetworkPolicyService) (*pb.ServiceResponse, error) {
	utils.Info.Println(ctx)
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getNetworkPolicy(req)

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

	/*converToResp(serviceResp,req.ProjectId,statusCode,resp)

	url := fmt.Sprintf("%s%s",constants.KubernetesEngineURL,constants.KUBERNETES_SERVICES_DEPLOYMENT)
	statusCode, resp, err := utils.Post(url,ksdRequest,getHeaders(ctx,req.ProjectId))

	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp,err)
		return serviceResp,err
	}
	converToResp(serviceResp,req.ProjectId,statusCode,resp)
	return serviceResp,nil*/
}
func (s *Server) GetNetworkPolicy(ctx context.Context, req *pb.NetworkPolicyService) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getNetworkPolicy(req)

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
func (s *Server) DeleteNetworkPolicy(ctx context.Context, req *pb.NetworkPolicyService) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getNetworkPolicy(req)

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
func (s *Server) PatchNetworkPolicy(ctx context.Context, req *pb.NetworkPolicyService) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getNetworkPolicy(req)

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
func (s *Server) PutNetworkPolicy(ctx context.Context, req *pb.NetworkPolicyService) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getNetworkPolicy(req)

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

func getNetworkPolicy(input *pb.NetworkPolicyService) (*net.NetworkPolicy, error) {
	var np = new(net.NetworkPolicy)
	np.Name = input.Name
	np.TypeMeta.Kind = "NetworkPolicy"
	np.TypeMeta.APIVersion = "networking.k8s.io/v1"
	np.Namespace = input.Namespace
	if ls, err := getLabelSelector(input.ServiceAttributes.PodSelector); err == nil {
		if ls != nil {
			np.Spec.PodSelector = *ls
		}
	} else {
		utils.Error.Println(err)
	}
	var err error
	if len(input.ServiceAttributes.Ingress) > 0 {
		for _, each := range input.ServiceAttributes.Ingress {
			temp := net.NetworkPolicyIngressRule{}
			for _, each2 := range each.From {
				temp2 := net.NetworkPolicyPeer{}
				temp2.PodSelector, err = getLabelSelector(each2.PodSelector)
				if err != nil {
					return nil, err
				}
				temp2.NamespaceSelector, err = getLabelSelector(each2.NamespaceSelector)
				if err != nil {
					return nil, err
				}
				if each2.IpBlock != nil {
					temp2.IPBlock = new(net.IPBlock)
					temp2.IPBlock.CIDR = each2.IpBlock.Cidr
					temp2.IPBlock.Except = each2.IpBlock.Except
				}
				temp.From = append(temp.From, temp2)
			}
			for _, each2 := range each.Ports {
				temp2 := net.NetworkPolicyPort{}
				if each2.Port.PortName != "" {
					port := intstr.FromString(each2.Port.PortName)
					temp2.Port = &port
				} else {
					port := intstr.FromInt(int(each2.Port.PortNumber))
					temp2.Port = &port
				}
				if each2.Protocol.String() == pb.Protocol_SCTP.String() {
					tcp := cor.ProtocolSCTP
					temp2.Protocol = &tcp
				} else if each2.Protocol.String() == pb.Protocol_UDP.String() {
					tcp := cor.ProtocolUDP
					temp2.Protocol = &tcp
				} else if each2.Protocol.String() == pb.Protocol_TCP.String() {
					tcp := cor.ProtocolTCP
					temp2.Protocol = &tcp
				}

				temp.Ports = append(temp.Ports, temp2)
			}
			np.Spec.Ingress = append(np.Spec.Ingress, temp)
		}
		np.Spec.PolicyTypes = append(np.Spec.PolicyTypes, net.PolicyTypeIngress)
	}
	if len(input.ServiceAttributes.Egress) > 0 {

		for _, each := range input.ServiceAttributes.Egress {
			temp := net.NetworkPolicyEgressRule{}
			for _, each2 := range each.To {
				temp2 := net.NetworkPolicyPeer{}
				temp2.PodSelector, err = getLabelSelector(each2.PodSelector)
				if err != nil {
					return nil, err
				}
				temp2.NamespaceSelector, err = getLabelSelector(each2.NamespaceSelector)
				if err != nil {
					return nil, err
				}
				if each2.IpBlock != nil {
					temp2.IPBlock = new(net.IPBlock)
					temp2.IPBlock.CIDR = each2.IpBlock.Cidr
					temp2.IPBlock.Except = each2.IpBlock.Except
				}
				temp.To = append(temp.To, temp2)
			}
			for _, each2 := range each.Ports {
				temp2 := net.NetworkPolicyPort{}
				if each2.Port.PortName != "" {
					port := intstr.FromString(each2.Port.PortName)
					temp2.Port = &port
				} else {
					port := intstr.FromInt(int(each2.Port.PortNumber))
					temp2.Port = &port
				}
				if each2.Protocol.String() == pb.Protocol_SCTP.String() {
					tcp := cor.ProtocolSCTP
					temp2.Protocol = &tcp
				} else if each2.Protocol.String() == pb.Protocol_UDP.String() {
					tcp := cor.ProtocolUDP
					temp2.Protocol = &tcp
				} else if each2.Protocol.String() == pb.Protocol_TCP.String() {
					tcp := cor.ProtocolTCP
					temp2.Protocol = &tcp
				}

				temp.Ports = append(temp.Ports, temp2)
			}
			np.Spec.Egress = append(np.Spec.Egress, temp)
		}
		np.Spec.PolicyTypes = append(np.Spec.PolicyTypes, net.PolicyTypeEgress)
	}
	return np, nil
}

func getLabelSelector(service *pb.LabelSelectorObj) (*metav1.LabelSelector, error) {
	if service == nil {
		return nil, nil
	}
	lenl := len(service.MatchLabel)
	lene := len(service.MatchExpression)
	if lene == 0 && lenl == 0 {
		return nil, nil
	}
	ls := new(metav1.LabelSelector)
	if lenl > 0 {
		ls.MatchLabels = make(map[string]string)
	}
	for k, v := range service.MatchLabel {
		ls.MatchLabels[k] = v
	}
	for i := 0; i < len(service.MatchExpression); i++ {
		if len(service.MatchExpression[i].Key) > 0 && (service.MatchExpression[i].Operator == pb.LabelSelectorOperator_DoesNotExist ||
			service.MatchExpression[i].Operator == pb.LabelSelectorOperator_Exists ||
			service.MatchExpression[i].Operator == pb.LabelSelectorOperator_In ||
			service.MatchExpression[i].Operator == pb.LabelSelectorOperator_NotIn) {
			byteData, err := json.Marshal(service.MatchExpression[i])
			if err != nil {
				return nil, err
			}
			var temp metav1.LabelSelectorRequirement

			err = json.Unmarshal(byteData, &temp)
			if err != nil {
				return nil, err
			}
			ls.MatchExpressions = append(ls.MatchExpressions, temp)
		}
	}
	return ls, nil
}
