package core

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc"
	"istio-service-mesh/constants"
	types "github.com/gogo/protobuf/types"
	pb "istio-service-mesh/core/proto"
	"istio-service-mesh/utils"
	"istio.io/api/networking/v1alpha3"
	istioClient "istio.io/client-go/pkg/apis/networking/v1alpha3"
	"net/http"
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
	vService.Hosts= input.ServiceAttributes.Hosts
	vService.Gateways=input.ServiceAttributes.Gateways


	for _,http := range input.ServiceAttributes.Http {
		vSer := v1alpha3.HTTPRoute{}
		vSer.Name = http.Name
		for i, match := range http.MatchRequest {
			vSer.Match[i] = &v1alpha3.HTTPMatchRequest{}
			vSer.Match[i].Name = match.Name
			vSer.Match[i].Uri.MatchType = match.Uri.Type
			vSer.Match[i].Scheme = match.Scheme
			vSer.Match[i].Method = match.Method
			vSer.Match[i].Authority = match.Authority
		}
		for i, route := range http.Route {
			vSer.Route[i] = &v1alpha3.HTTPRouteDestination{}
			vSer.Route[i].Destination.Port.Number = uint32(route.Routes.Port)
			vSer.Route[i].Destination.Host = route.Routes.Host
			vSer.Route[i].Destination.Subset = route.Routes.Subset
			vSer.Route[i].Weight = route.Weight
		}

		vSer.Redirect = &v1alpha3.HTTPRedirect{}
		vSer.Redirect.Uri = http.HttpRedirect.Uri
		vSer.Redirect.Authority = http.HttpRedirect.Authority
		vSer.Redirect.RedirectCode = uint32(http.HttpRedirect.RedirectCode)

		vSer.Rewrite = &v1alpha3.HTTPRewrite{}
		vSer.Rewrite.Uri = http.HttpRewrite.Uri
		vSer.Rewrite.Authority = http.HttpRewrite.Authority

		vSer.Timeout.Seconds = http.Timeout

		vSer.Retries = &v1alpha3.HTTPRetry{}
		vSer.Retries.Attempts = http.Retry.TotalAttempts
		vSer.Retries.PerTryTimeout.Seconds = http.Retry.PerTryTimeout
		vSer.Retries.RetryOn = http.Retry.RetryOn

		vSer.Fault = &v1alpha3.HTTPFaultInjection{}
		vSer.Fault.Abort.ErrorType = http.FaultInjection.AbortErrorValue
		vSer.Fault.Abort.Percentage.Value = float64(http.FaultInjection.AbortPercentage)
		vSer.Fault.Delay.HttpDelayType = http.FaultInjection.DelayType
		vSer.Fault.Delay.Percentage.Value = float64(http.FaultInjection.DelayValue)

		//vService.Http[i].Mirror.Subset=http.Route.Routes.Subset
		//vService.Http[i].Mirror.Host=http.Route.Routes.Host
		//vService.Http[i].Mirror.Port.Number=uint32(http.Route.Routes.Port)

		vSer.CorsPolicy = &v1alpha3.CorsPolicy{}
		vSer.CorsPolicy.AllowOrigin = http.CorsPolicy.AllowOrigin
		vSer.CorsPolicy.AllowMethods = http.CorsPolicy.AllMethod
		vSer.CorsPolicy.AllowHeaders = http.CorsPolicy.AllowHeaders
		vSer.CorsPolicy.ExposeHeaders = http.CorsPolicy.ExposeHeaders
		vSer.CorsPolicy.MaxAge.Seconds = http.CorsPolicy.MaxAge
		vSer.CorsPolicy.AllowCredentials.Value = http.CorsPolicy.AllowCredentials
	}



	for i,serv := range input.ServiceAttributes.Tls {
		tls := v1alpha3.TLSRoute{}
		for i,match :=range serv.Match{
			tls.Match[i] = &v1alpha3.TLSMatchAttributes{}
			tls.Match[i].SniHosts =match.SniHosts
			tls.Match[i].DestinationSubnets= match.DestinationSubnets
			tls.Match[i].Gateways = match.Gateways
			tls.Match[i].Port =  uint32(match.Port)
			tls.Match[i].SourceSubnet =  match.SourceSubnet
		}

		for i,route :=range serv.Routes{
			tls.Route[i] := &v1alpha3.RouteDestination{}
			tls.Route[i].Weight = route.Weight
			tls.Route[i].Destination.Port.Number=uint32(route.RouteDestination.Port)
			tls.Route[i].Destination.Subset=route.RouteDestination.Subnet
			tls.Route[i].Destination.Host=route.RouteDestination.Host
		}
	}

	for i,serv := range input.ServiceAttributes.Tcp {
		tcp := v1alpha3.TCPRoute{}
		for i,match :=range serv.Match{
			tcp.Match[i] = &v1alpha3.L4MatchAttributes{}
			tcp.Match[i].SourceLabels =match.SourceLabels
			tcp.Match[i].DestinationSubnets= match.DestinationSubnets
			tcp.Match[i].Gateways = match.Gateways
			tcp.Match[i].Port =  uint32(match.Port)
			tcp.Match[i].SourceSubnet =  match.SourceSubnet
		}

		for i,route :=range serv.Routes{
			tls.Route[i] := &v1alpha3.RouteDestination{}
			tls.Route[i].Weight = route.Weight
			tls.Route[i].Destination.Port.Number=uint32(route.RouteDestination.Port)
			tls.Route[i].Destination.Subset=route.RouteDestination.Subnet
			tls.Route[i].Destination.Host=route.RouteDestination.Host
		}
	}
	return vServ, nil
}
func getVirtualServiceSpec() (v1alpha3.VirtualService, error) {

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
