package core

import (
	"bitbucket.org/cloudplex-devs/istio-service-mesh/constants"
	"bitbucket.org/cloudplex-devs/istio-service-mesh/utils"
	pb1 "bitbucket.org/cloudplex-devs/kubernetes-services-deployment/core/proto"
	pb "bitbucket.org/cloudplex-devs/microservices-mesh-engine/core/services/proto"
	"context"
	"encoding/json"
	"google.golang.org/grpc"
	"istio.io/api/networking/v1alpha3"
	istioClient "istio.io/client-go/pkg/apis/networking/v1alpha3"
	"strings"
)

type Server struct {
}

func (s *Server) CreateGateway(ctx context.Context, req *pb.GatewayService) (*pb.ServiceResponse, error) {
	utils.Info.Println(ctx)
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getRequestObject(req)
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
func (s *Server) GetGateway(ctx context.Context, req *pb.GatewayService) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getRequestObject(req)

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
func (s *Server) DeleteGateway(ctx context.Context, req *pb.GatewayService) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getRequestObject(req)

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
func (s *Server) PatchGateway(ctx context.Context, req *pb.GatewayService) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getRequestObject(req)

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
func (s *Server) PutGateway(ctx context.Context, req *pb.GatewayService) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getRequestObject(req)

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

func getIstioGateway(input *pb.GatewayService) (*istioClient.Gateway, error) {
	var istioServ = new(istioClient.Gateway)
	labels := make(map[string]string)
	labels["app"] = strings.ToLower(input.Name)
	labels["version"] = strings.ToLower(input.Version)
	istioServ.Labels = labels
	istioServ.Kind = constants.Gateway.String() //"Gateway"
	istioServ.APIVersion = "networking.istio.io/v1alpha3"
	istioServ.Name = input.Name
	istioServ.Namespace = input.Namespace
	gateway := v1alpha3.Gateway{}

	gateway.Selector = input.ServiceAttributes.Selectors

	for _, serverInput := range input.ServiceAttributes.Servers {
		server := new(v1alpha3.Server)
		if serverInput.Port != nil {
			server.Port = new(v1alpha3.Port)
			server.Port.Name = serverInput.Port.Name
			server.Port.Number = serverInput.Port.Number
			server.Port.Protocol = serverInput.Port.GetProtocol().String()
		}
		if serverInput.Tls != nil {
			server.Tls = new(v1alpha3.ServerTLSSettings)
			server.Tls.HttpsRedirect = serverInput.Tls.HttpsRedirect
			server.Tls.Mode = v1alpha3.ServerTLSSettings_TLSmode(int32(serverInput.Tls.Mode))
			server.Tls.ServerCertificate = serverInput.Tls.ServerCertificate
			server.Tls.CaCertificates = serverInput.Tls.CaCertificate
			server.Tls.PrivateKey = serverInput.Tls.PrivateKey
			server.Tls.SubjectAltNames = serverInput.Tls.SubjectAltName
			server.Tls.CredentialName = serverInput.Tls.CredentialName
			server.Tls.MinProtocolVersion = v1alpha3.ServerTLSSettings_TLSProtocol(int32(serverInput.Tls.MinProtocolVersion))
			server.Tls.MaxProtocolVersion = v1alpha3.ServerTLSSettings_TLSProtocol(int32(serverInput.Tls.MaxProtocolVersion))
		}
		server.Hosts = serverInput.Hosts
		gateway.Servers = append(gateway.Servers, server)
	}
	istioServ.Spec = gateway
	return istioServ, nil
}
func getIstioGatewaySpec() (v1alpha3.Gateway, error) {

	gateway := v1alpha3.Gateway{}
	var hosts []string
	hosts = append(hosts, "*")
	var servers []*v1alpha3.Server

	var serv v1alpha3.Server
	serv.Port = &v1alpha3.Port{Name: strings.ToLower("HTTP"), Protocol: "HTTP", Number: uint32(80)}
	serv.Hosts = hosts
	servers = append(servers, &serv)

	/*var serv2 v1alpha3.Server
	serv2.Port = &v1alpha3.Port{Name: strings.ToLower("HTTPS"), Protocol: "HTTPS", Number: uint32(443)}
	serv2.Hosts = hosts
	servers = append(servers, &serv2)*/

	selector := make(map[string]string)

	selector["istio"] = "ingressgateway"
	gateway.Selector = selector
	gateway.Servers = servers
	return gateway, nil
}

func getRequestObject(req *pb.GatewayService) (*istioClient.Gateway, error) {
	gtwReq, err := getIstioGateway(req)
	if err != nil {
		utils.Error.Println(err)

		return nil, err
	}
	return gtwReq, nil
}
