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
	lenl := len(service.MatchLabels)
	lene := len(service.MatchExpressions)
	if lene == 0 && lenl == 0 {
		return nil, nil
	}
	ls := new(metav1.LabelSelector)
	if lenl > 0 {
		ls.MatchLabels = make(map[string]string)
	}
	for k, v := range service.MatchLabels {
		ls.MatchLabels[k] = v
	}
	for i := 0; i < len(service.MatchExpressions); i++ {
		if len(service.MatchExpressions[i].Key) > 0 {

			var temp = new(metav1.LabelSelectorRequirement)
			temp.Key = service.MatchExpressions[i].Key
			for _, each := range service.MatchExpressions[i].Values {
				temp.Values = append(temp.Values, each)
			}
			if service.MatchExpressions[i].Operator.String() == pb.LabelSelectorOperator_DoesNotExist.String() {
				temp.Operator = metav1.LabelSelectorOpDoesNotExist
			} else if service.MatchExpressions[i].Operator.String() == pb.LabelSelectorOperator_Exists.String() {
				temp.Operator = metav1.LabelSelectorOpExists
			} else if service.MatchExpressions[i].Operator.String() == pb.LabelSelectorOperator_In.String() {
				temp.Operator = metav1.LabelSelectorOpIn
			} else if service.MatchExpressions[i].Operator.String() == pb.LabelSelectorOperator_NotIn.String() {
				temp.Operator = metav1.LabelSelectorOpNotIn
			} else {
				return nil, errors.New("Invalid operator in label selector")
			}

			//m := jsonpb.Marshaler{}
			//tJson, err := m.MarshalToString(service.MatchExpressions[i])
			//if err != nil {
			//	return nil, err
			//}
			//um := jsonpb.Unmarshaler{}
			//err = um.Unmarshal(strings.NewReader(tJson), temp)
			//if err != nil {
			//	return nil, err
			//}
			ls.MatchExpressions = append(ls.MatchExpressions, *temp)
		} else {
			return nil, errors.New("key required in label selector match expression object")
		}
	}
	return ls, nil
}
