package core

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc"
	"istio-service-mesh/constants"
	pb "istio-service-mesh/core/proto"
	"istio-service-mesh/utils"
	"istio.io/api/networking/v1alpha3"
	istioClient "istio.io/client-go/pkg/apis/networking/v1alpha3"
	"strings"
)

func (s *Server) CreateVirtualService(ctx context.Context, req *pb.VirtualService) (*pb.ServiceResponse, error) {
	utils.Info.Println(ctx)
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	vsrvRequest, err := getVSRequestObject(req)

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

	raw, err := json.Marshal(vsrvRequest)
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
func (s *Server) GetVirtualService(ctx context.Context, req *pb.VirtualService) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	vsrvRequest, err := getVSRequestObject(req)

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

	raw, err := json.Marshal(vsrvRequest)
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
func (s *Server) DeleteVirtualService(ctx context.Context, req *pb.VirtualService) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	vsrvRequest, err := getVSRequestObject(req)

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

	raw, err := json.Marshal(vsrvRequest)
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
func (s *Server) PatchVirtualService(ctx context.Context, req *pb.VirtualService) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	vsrvRequest, err := getVSRequestObject(req)

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

	raw, err := json.Marshal(vsrvRequest)
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
func (s *Server) PutVirtualService(ctx context.Context, req *pb.VirtualService) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	vsrvRequest, err := getVSRequestObject(req)

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

	raw, err := json.Marshal(vsrvRequest)
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

func getVirtualService(input *pb.VirtualService) (*istioClient.VirtualService, error) {
	var vServ = new(istioClient.VirtualService)
	labels := make(map[string]string)
	labels["app"] = strings.ToLower(input.Name)
	labels["version"] = strings.ToLower(input.Version)
	vServ.Labels = labels
	vServ.Kind = "VirtualService"
	vServ.APIVersion = "networking.v.io/v1alpha3"
	vServ.Name = input.Name
	vServ.Namespace = input.Namespace
	vService := v1alpha3.VirtualService{}
	for i, _ := range input.ServiceAttributes.Hosts {
		vService.Hosts[i] = input.ServiceAttributes.Hosts[i]
	}
	for i, _ := range input.ServiceAttributes.Gateways {
		vService.Gateways[i] = input.ServiceAttributes.Gateways[i]

	}
	for i, _ := range input.ServiceAttributes.Http {
		vService.Http[i].Name = input.ServiceAttributes.Http[i].Name
		vService.Http[i].CorsPolicy.AllowOrigin = input.ServiceAttributes.Http[i].CorsPolicy.AllowOrigin
		vService.Http[i].CorsPolicy.AllowMethods = input.ServiceAttributes.Http[i].CorsPolicy.AllMethod
		vService.Http[i].CorsPolicy.AllowHeaders = input.ServiceAttributes.Http[i].CorsPolicy.AllowHeaders
		vService.Http[i].CorsPolicy.ExposeHeaders = input.ServiceAttributes.Http[i].CorsPolicy.ExposeHeaders
		vService.Http[i].CorsPolicy.MaxAge = input.ServiceAttributes.Http[i].CorsPolicy.MaxAge
		vService.Http[i].CorsPolicy.AllowCredentials = input.ServiceAttributes.Http[i].CorsPolicy.AllowCrentials

	}
	vService.Selector = input.ServiceAttributes.Selectors

	for _, serverInput := range input.ServiceAttributes.Servers {
		server := new(v1alpha3.Server)
		if serverInput.Port != nil {
			server.Port = new(v1alpha3.Port)
			server.Port.Name = serverInput.Port.Name
			server.Port.Number = serverInput.Port.Nummber
			server.Port.Protocol = serverInput.Port.GetProtocol().String()
		}
		if serverInput.Tls != nil {
			server.Tls = new(v1alpha3.Server_TLSOptions)
			server.Tls.HttpsRedirect = serverInput.Tls.HttpsRedirect
			server.Tls.Mode = v1alpha3.Server_TLSOptions_TLSmode(int32(serverInput.Tls.Mode))
			server.Tls.ServerCertificate = serverInput.Tls.ServerCertificate
			server.Tls.CaCertificates = serverInput.Tls.CaCertificate
			server.Tls.PrivateKey = serverInput.Tls.PrivateKey
			server.Tls.SubjectAltNames = serverInput.Tls.SubjectAltName
			server.Tls.MinProtocolVersion = v1alpha3.Server_TLSOptions_TLSProtocol(int32(serverInput.Tls.MinProtocolVersion))
			server.Tls.MaxProtocolVersion = v1alpha3.Server_TLSOptions_TLSProtocol(int32(serverInput.Tls.MaxProtocolVersion))
		}
		server.Hosts = serverInput.Hosts
		vService.Servers = append(vService.Servers, server)
	}
	vServ.Spec = vService
	return vServ, nil
}
func getIstioGatewaySpec() (v1alpha3.VirtualService, error) {

	vService := v1alpha3.Gateway{}
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
	vService.Selector = selector
	vService.Servers = servers
	return vService, nil
}

func getVSRequestObject(req *pb.VirtualService) (*istioClient.VirtualService, error) {
	gtwReq, err := getVirtualService(req)
	if err != nil {
		utils.Error.Println(err)

		return nil, err
	}
	return gtwReq, nil
}
