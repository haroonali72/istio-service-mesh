package core

import (
	"bitbucket.org/cloudplex-devs/istio-service-mesh/constants"
	"bitbucket.org/cloudplex-devs/istio-service-mesh/utils"
	pb1 "bitbucket.org/cloudplex-devs/kubernetes-services-deployment/core/proto"
	pb "bitbucket.org/cloudplex-devs/microservices-mesh-engine/core/services/proto"
	"context"
	"encoding/json"
	types "github.com/gogo/protobuf/types"
	"google.golang.org/grpc"
	"istio.io/api/networking/v1alpha3"
	istioClient "istio.io/client-go/pkg/apis/networking/v1alpha3"
	"strconv"
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

func getVirtualService(input *pb.VirtualService) (*istioClient.VirtualService, error) {
	var vServ = new(istioClient.VirtualService)
	labels := make(map[string]string)
	labels["app"] = strings.ToLower(input.Name)
	labels["version"] = strings.ToLower(input.Version)
	vServ.Labels = labels
	vServ.Kind = constants.VirtualService.String() //"VirtualService"
	vServ.APIVersion = "networking.istio.io/v1alpha3"
	vServ.Name = input.Name
	vServ.Namespace = input.Namespace

	vService := v1alpha3.VirtualService{}

	vService.Hosts = input.ServiceAttributes.Hosts
	vService.Gateways = input.ServiceAttributes.Gateways

	for _, http := range input.ServiceAttributes.Http {
		vSer := v1alpha3.HTTPRoute{}
		vSer.Name = http.Name

		for _, match := range http.HttpMatch {
			m := &v1alpha3.HTTPMatchRequest{}
			m.Name = match.Name
			if match.Uri != nil {
				if match.Uri.Type == "prefix" {
					m.Uri = &v1alpha3.StringMatch{
						MatchType: &v1alpha3.StringMatch_Prefix{
							Prefix: match.Uri.Value,
						},
					}
				} else if match.Uri.Type == "exact" {
					m.Uri = &v1alpha3.StringMatch{
						MatchType: &v1alpha3.StringMatch_Exact{
							Exact: match.Uri.Value,
						},
					}
				} else if match.Uri.Type == "regex" {
					m.Uri = &v1alpha3.StringMatch{
						MatchType: &v1alpha3.StringMatch_Regex{
							Regex: match.Uri.Value,
						},
					}
				}
			}
			if match.Scheme != nil {
				if match.Scheme.Type == "prefix" {
					m.Scheme = &v1alpha3.StringMatch{
						MatchType: &v1alpha3.StringMatch_Prefix{
							Prefix: match.Scheme.Value,
						},
					}
				} else if match.Scheme.Type == "exact" {
					m.Scheme = &v1alpha3.StringMatch{
						MatchType: &v1alpha3.StringMatch_Exact{
							Exact: match.Scheme.Value,
						},
					}
				} else if match.Scheme.Type == "regex" {
					m.Scheme = &v1alpha3.StringMatch{
						MatchType: &v1alpha3.StringMatch_Regex{
							Regex: match.Scheme.Value,
						},
					}
				}
			}
			if match.Method != nil {
				if match.Method.Type == "prefix" {
					m.Method = &v1alpha3.StringMatch{
						MatchType: &v1alpha3.StringMatch_Prefix{
							Prefix: match.Method.Value,
						},
					}
				} else if match.Method.Type == "exact" {
					m.Method = &v1alpha3.StringMatch{
						MatchType: &v1alpha3.StringMatch_Exact{
							Exact: match.Method.Value,
						},
					}
				} else if match.Method.Type == "regex" {
					m.Method = &v1alpha3.StringMatch{
						MatchType: &v1alpha3.StringMatch_Regex{
							Regex: match.Method.Value,
						},
					}
				}
			}
			if match.Authority != nil {
				if match.Authority.Type == "prefix" {
					m.Authority = &v1alpha3.StringMatch{
						MatchType: &v1alpha3.StringMatch_Prefix{
							Prefix: match.Authority.Value,
						},
					}
				} else if match.Authority.Type == "exact" {
					m.Authority = &v1alpha3.StringMatch{
						MatchType: &v1alpha3.StringMatch_Exact{
							Exact: match.Authority.Value,
						},
					}
				} else if match.Authority.Type == "regex" {
					m.Authority = &v1alpha3.StringMatch{
						MatchType: &v1alpha3.StringMatch_Regex{
							Regex: match.Authority.Value,
						},
					}
				}
			} else if match.Headers != nil {
				for key, value := range match.Headers {
					if value.Type == "prefix" {
						m.Headers = make(map[string]*v1alpha3.StringMatch)
						m.Headers[key] = &v1alpha3.StringMatch{
							MatchType: &v1alpha3.StringMatch_Prefix{
								Prefix: value.Value,
							},
						}
					} else if value.Type == "exact" {
						m.Headers = make(map[string]*v1alpha3.StringMatch)
						m.Headers[key] = &v1alpha3.StringMatch{
							MatchType: &v1alpha3.StringMatch_Exact{
								Exact: value.Value,
							},
						}

					} else if value.Type == "regex" {
						m.Headers = make(map[string]*v1alpha3.StringMatch)
						m.Headers[key] = &v1alpha3.StringMatch{
							MatchType: &v1alpha3.StringMatch_Regex{
								Regex: value.Value,
							},
						}
					}
				}
			}
			vSer.Match = append(vSer.Match, m)
		}

		for _, route := range http.HttpRoute {
			r := &v1alpha3.HTTPRouteDestination{}
			for _, rou := range route.Routes {

				r.Destination = &v1alpha3.Destination{}
				r.Destination.Port = &v1alpha3.PortSelector{}
				r.Destination.Port.Number = uint32(rou.Port)
				r.Destination.Host = rou.Host
				r.Destination.Subset = rou.Subset

			}
			r.Weight = route.Weight
			vSer.Route = append(vSer.Route, r)
		}
		if http.HttpRedirect != nil {
			vSer.Redirect = &v1alpha3.HTTPRedirect{}
			vSer.Redirect.Uri = http.HttpRedirect.Uri
			vSer.Redirect.Authority = http.HttpRedirect.Authority
			vSer.Redirect.RedirectCode = uint32(http.HttpRedirect.RedirectCode)
		}
		if http.HttpRewrite != nil {
			vSer.Rewrite = &v1alpha3.HTTPRewrite{}
			vSer.Rewrite.Uri = http.HttpRewrite.Uri
			vSer.Rewrite.Authority = http.HttpRewrite.Authority
		}
		if http.Timeout != 0 {
			vSer.Timeout = &types.Duration{Nanos: http.Timeout}
		}

		if http.FaultInjection != nil {

			if http.FaultInjection.DelayType == "FixedDelay" {
				vSer.Fault = &v1alpha3.HTTPFaultInjection{}
				vSer.Fault.Delay = &v1alpha3.HTTPFaultInjection_Delay{
					HttpDelayType: &v1alpha3.HTTPFaultInjection_Delay_FixedDelay{
						FixedDelay: &types.Duration{Seconds: int64(http.FaultInjection.DelayValue)},
					},
					Percentage: &v1alpha3.Percent{Value: float64(http.FaultInjection.FaultPercentage)},
				}
			} else if http.FaultInjection.DelayType == "ExponentialDelay" {
				vSer.Fault = &v1alpha3.HTTPFaultInjection{}
				vSer.Fault.Delay = &v1alpha3.HTTPFaultInjection_Delay{
					HttpDelayType: &v1alpha3.HTTPFaultInjection_Delay_FixedDelay{
						FixedDelay: &types.Duration{Seconds: int64(http.FaultInjection.DelayValue)},
					},
					Percentage: &v1alpha3.Percent{Value: float64(http.FaultInjection.FaultPercentage)},
				}
			}
			value, _ := strconv.ParseInt(http.FaultInjection.AbortPercentage, 10, 32)
			if http.FaultInjection.AbortErrorType == "HttpStatus" {
				vSer.Fault.Abort = &v1alpha3.HTTPFaultInjection_Abort{
					ErrorType: &v1alpha3.HTTPFaultInjection_Abort_HttpStatus{HttpStatus: int32(value)},
				}
			} else if http.FaultInjection.AbortErrorType == "GrpcStatus" {
				vSer.Fault.Abort = &v1alpha3.HTTPFaultInjection_Abort{
					ErrorType: &v1alpha3.HTTPFaultInjection_Abort_GrpcStatus{GrpcStatus: http.FaultInjection.AbortPercentage},
				}
			} else if http.FaultInjection.AbortErrorType == "Http2Status" {
				vSer.Fault.Abort = &v1alpha3.HTTPFaultInjection_Abort{
					ErrorType: &v1alpha3.HTTPFaultInjection_Abort_Http2Error{Http2Error: http.FaultInjection.AbortPercentage},
				}
			}
		}
		if http.CorsPolicy != nil {
			vSer.CorsPolicy = &v1alpha3.CorsPolicy{}
			vSer.CorsPolicy.AllowOrigin = http.CorsPolicy.AllowOrigin
			vSer.CorsPolicy.AllowMethods = http.CorsPolicy.AllowMethod
			vSer.CorsPolicy.AllowHeaders = http.CorsPolicy.AllowHeaders
			vSer.CorsPolicy.ExposeHeaders = http.CorsPolicy.ExposeHeaders
			vSer.CorsPolicy.MaxAge = &types.Duration{Nanos: http.CorsPolicy.MaxAge}
			vSer.CorsPolicy.AllowCredentials = &types.BoolValue{Value: http.CorsPolicy.AllowCredentials}
		}
		if http.Retry != nil {
			vSer.Retries = &v1alpha3.HTTPRetry{}
			vSer.Retries.Attempts = http.Retry.TotalAttempt
			vSer.Retries.PerTryTimeout = &types.Duration{Seconds: http.Retry.PerTryTimeout}
			vSer.Retries.RetryOn = http.Retry.RetryOn
		}
		vService.Http = append(vService.Http, &vSer)
	}

	for _, serv := range input.ServiceAttributes.Tls {
		tls := v1alpha3.TLSRoute{}
		for _, match := range serv.TlsMatch {
			m := &v1alpha3.TLSMatchAttributes{}
			for _, s := range match.SniHosts {
				m.SniHosts = append(m.SniHosts, s)
			}
			for _, d := range match.DestinationSubnets {
				m.DestinationSubnets = append(m.DestinationSubnets, d)
			}
			for _, g := range match.Gateways {
				m.Gateways = append(m.Gateways, g)
			}
			m.Port = uint32(match.Port)
			//m.SourceSubnet = match.SourceSubnet
			tls.Match = append(tls.Match, m)
		}

		for _, route := range serv.TlsRoute {
			r := &v1alpha3.RouteDestination{}
			r.Destination = &v1alpha3.Destination{}
			r.Weight = route.Weight
			r.Destination.Port = &v1alpha3.PortSelector{}
			r.Destination.Port.Number = uint32(route.RouteDestination.Port)
			r.Destination.Subset = route.RouteDestination.Subnet
			r.Destination.Host = route.RouteDestination.Host
			tls.Route = append(tls.Route, r)
		}
		vService.Tls = append(vService.Tls, &tls)
	}

	for _, serv := range input.ServiceAttributes.Tcp {
		tcp := v1alpha3.TCPRoute{}
		for _, match := range serv.TcpMatch {
			m := &v1alpha3.L4MatchAttributes{}
			m.SourceLabels = match.SourceLabels
			m.DestinationSubnets = match.DestinationSubnets
			m.Gateways = match.Gateways
			m.Port = uint32(match.Port)
			m.SourceSubnet = match.SourceSubnet
			tcp.Match = append(tcp.Match, m)
		}

		for _, route := range serv.TcpRoutes {
			d := &v1alpha3.RouteDestination{}
			d.Destination = &v1alpha3.Destination{}
			d.Destination.Port = &v1alpha3.PortSelector{}
			d.Destination.Port.Number = uint32(route.Destination.Port)
			d.Destination.Subset = route.Destination.Subnet
			d.Destination.Host = route.Destination.Host
			d.Weight = route.Weight
			tcp.Route = append(tcp.Route, d)
		}
		vService.Tcp = append(vService.Tcp, &tcp)
	}

	vServ.Spec = vService
	return vServ, nil
}
func getVirtualServiceSpec() (v1alpha3.VirtualService, error) {

	vService := v1alpha3.VirtualService{}

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
