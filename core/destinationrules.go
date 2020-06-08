package core

import (
	"context"
	"encoding/json"
	"github.com/gogo/protobuf/types"
	"time"

	//	types "github.com/gogo/protobuf/types"
	"bitbucket.org/cloudplex-devs/istio-service-mesh/constants"
	"bitbucket.org/cloudplex-devs/istio-service-mesh/utils"
	pb1 "bitbucket.org/cloudplex-devs/kubernetes-services-deployment/core/proto"
	pb "bitbucket.org/cloudplex-devs/microservices-mesh-engine/core/services/proto"
	"google.golang.org/grpc"
	"istio.io/api/networking/v1alpha3"
	istioClient "istio.io/client-go/pkg/apis/networking/v1alpha3"
	"strings"
)

func (s *Server) CreateDestinationRules(ctx context.Context, req *pb.DestinationRules) (*pb.ServiceResponse, error) {
	utils.Info.Println(ctx)
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	dsRequest, err := getDRRequestObject(req)

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

	raw, err := json.Marshal(dsRequest)
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
func (s *Server) GetDestinationRules(ctx context.Context, req *pb.DestinationRules) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	dsRequest, err := getDRRequestObject(req)

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

	raw, err := json.Marshal(dsRequest)
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
func (s *Server) DeleteDestinationRules(ctx context.Context, req *pb.DestinationRules) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	dsRequest, err := getDRRequestObject(req)

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

	raw, err := json.Marshal(dsRequest)
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
func (s *Server) PatchDestinationRules(ctx context.Context, req *pb.DestinationRules) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	dsRequest, err := getDRRequestObject(req)

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

	raw, err := json.Marshal(dsRequest)
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
func (s *Server) PutDestinationRules(ctx context.Context, req *pb.DestinationRules) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	dsRequest, err := getDRRequestObject(req)

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

	raw, err := json.Marshal(dsRequest)
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

func getDestinationRules(input *pb.DestinationRules) (*istioClient.DestinationRule, error) {

	var vServ = new(istioClient.DestinationRule)
	labels := make(map[string]string)
	labels["app"] = strings.ToLower(input.Name)
	labels["version"] = strings.ToLower(input.Version)
	vServ.Labels = labels
	vServ.Kind = constants.DestinationRule.String() //"DestinationRule"
	vServ.APIVersion = "networking.istio.io/v1alpha3"
	vServ.Name = input.Name
	vServ.Namespace = input.Namespace

	vService := v1alpha3.DestinationRule{}
	vService.Host = input.ServiceAttribute.Host

	if input.ServiceAttribute.TrafficPolicy != nil {
		vService.TrafficPolicy = &v1alpha3.TrafficPolicy{}
		if input.ServiceAttribute.TrafficPolicy.LoadBalancer != nil {
			vService.TrafficPolicy.LoadBalancer = &v1alpha3.LoadBalancerSettings{}
			if (input.ServiceAttribute.TrafficPolicy.LoadBalancer.Simple >= 0 || input.ServiceAttribute.TrafficPolicy.LoadBalancer.Simple <= 3) && input.ServiceAttribute.TrafficPolicy.LoadBalancer.ConsistentHash == nil {
				vService.TrafficPolicy.LoadBalancer.LbPolicy = &v1alpha3.LoadBalancerSettings_Simple{
					Simple: v1alpha3.LoadBalancerSettings_SimpleLB(int32(input.ServiceAttribute.TrafficPolicy.LoadBalancer.Simple)),
				}
			} else if input.ServiceAttribute.TrafficPolicy.LoadBalancer.ConsistentHash != nil {
				if input.ServiceAttribute.TrafficPolicy.LoadBalancer.ConsistentHash.HttpCookie != nil {
					t := time.Duration(input.ServiceAttribute.TrafficPolicy.LoadBalancer.ConsistentHash.HttpCookie.Ttl)
					vService.TrafficPolicy.LoadBalancer.LbPolicy = &v1alpha3.LoadBalancerSettings_ConsistentHash{
						ConsistentHash: &v1alpha3.LoadBalancerSettings_ConsistentHashLB{
							HashKey: &v1alpha3.LoadBalancerSettings_ConsistentHashLB_HttpCookie{
								HttpCookie: &v1alpha3.LoadBalancerSettings_ConsistentHashLB_HTTPCookie{
									Name: input.ServiceAttribute.TrafficPolicy.LoadBalancer.ConsistentHash.HttpCookie.Name,
									Path: input.ServiceAttribute.TrafficPolicy.LoadBalancer.ConsistentHash.HttpCookie.Path,
									Ttl:  &t,
								},
							},
						},
					}
				} else if input.ServiceAttribute.TrafficPolicy.LoadBalancer.ConsistentHash.HttpHeaderName != "" {
					vService.TrafficPolicy.LoadBalancer.LbPolicy = &v1alpha3.LoadBalancerSettings_ConsistentHash{
						ConsistentHash: &v1alpha3.LoadBalancerSettings_ConsistentHashLB{
							HashKey: &v1alpha3.LoadBalancerSettings_ConsistentHashLB_HttpHeaderName{
								HttpHeaderName: input.ServiceAttribute.TrafficPolicy.LoadBalancer.ConsistentHash.HttpHeaderName,
							},
						},
					}
				} else if input.ServiceAttribute.TrafficPolicy.LoadBalancer.ConsistentHash.UseSourceIp != false {
					vService.TrafficPolicy.LoadBalancer.LbPolicy = &v1alpha3.LoadBalancerSettings_ConsistentHash{
						ConsistentHash: &v1alpha3.LoadBalancerSettings_ConsistentHashLB{
							HashKey: &v1alpha3.LoadBalancerSettings_ConsistentHashLB_UseSourceIp{
								UseSourceIp: input.ServiceAttribute.TrafficPolicy.LoadBalancer.ConsistentHash.UseSourceIp,
							},
						},
					}
				}
			}
		}
		if input.ServiceAttribute.TrafficPolicy.ConnectionPool != nil {

			if input.ServiceAttribute.TrafficPolicy.ConnectionPool.DrTcp != nil {
				vService.TrafficPolicy.ConnectionPool = &v1alpha3.ConnectionPoolSettings{}
				vService.TrafficPolicy.ConnectionPool.Tcp = &v1alpha3.ConnectionPoolSettings_TCPSettings{}
				vService.TrafficPolicy.ConnectionPool.Tcp.ConnectTimeout = &types.Duration{Seconds: int64(input.ServiceAttribute.TrafficPolicy.ConnectionPool.DrTcp.ConnectTimeout)}
				vService.TrafficPolicy.ConnectionPool.Tcp.MaxConnections = input.ServiceAttribute.TrafficPolicy.ConnectionPool.DrTcp.MaxConnections
				if input.ServiceAttribute.TrafficPolicy.ConnectionPool.DrTcp.TcpKeepAlive != nil {
					vService.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive = &v1alpha3.ConnectionPoolSettings_TCPSettings_TcpKeepalive{}
					vService.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive.Interval = &types.Duration{Seconds: int64(input.ServiceAttribute.TrafficPolicy.ConnectionPool.DrTcp.TcpKeepAlive.Interval)}
					vService.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive.Probes = input.ServiceAttribute.TrafficPolicy.ConnectionPool.DrTcp.TcpKeepAlive.Probes
					vService.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive.Time = &types.Duration{Seconds: int64(input.ServiceAttribute.TrafficPolicy.ConnectionPool.DrTcp.TcpKeepAlive.Time)}
				}
			}
			if input.ServiceAttribute.TrafficPolicy.ConnectionPool.DrHttp != nil {
				vService.TrafficPolicy.ConnectionPool = &v1alpha3.ConnectionPoolSettings{}
				vService.TrafficPolicy.ConnectionPool.Http = &v1alpha3.ConnectionPoolSettings_HTTPSettings{}
				vService.TrafficPolicy.ConnectionPool.Http.H2UpgradePolicy = v1alpha3.ConnectionPoolSettings_HTTPSettings_H2UpgradePolicy(input.ServiceAttribute.TrafficPolicy.ConnectionPool.DrHttp.ConnectionPoolSettingsHttpSettingsH2UpgradePolicy)
				vService.TrafficPolicy.ConnectionPool.Http.Http1MaxPendingRequests = input.ServiceAttribute.TrafficPolicy.ConnectionPool.DrHttp.Http_1MaxPendingRequests
				vService.TrafficPolicy.ConnectionPool.Http.Http2MaxRequests = int32(input.ServiceAttribute.TrafficPolicy.ConnectionPool.DrHttp.GetHttp_2MaxRequests())
				if input.ServiceAttribute.TrafficPolicy.ConnectionPool.DrHttp.IdleTimeout > 0 {
					vService.TrafficPolicy.ConnectionPool.Http.IdleTimeout = &types.Duration{Nanos: input.ServiceAttribute.TrafficPolicy.ConnectionPool.DrHttp.IdleTimeout}
				}
				vService.TrafficPolicy.ConnectionPool.Http.MaxRequestsPerConnection = input.ServiceAttribute.TrafficPolicy.ConnectionPool.DrHttp.MaxRequestsPerConnection
				vService.TrafficPolicy.ConnectionPool.Http.MaxRetries = input.ServiceAttribute.TrafficPolicy.ConnectionPool.DrHttp.MaxRetries
			}

		}
		if input.ServiceAttribute.TrafficPolicy.OutlierDetection != nil {
			vService.TrafficPolicy.OutlierDetection = &v1alpha3.OutlierDetection{}
			vService.TrafficPolicy.OutlierDetection.BaseEjectionTime = &types.Duration{Seconds: int64(input.ServiceAttribute.TrafficPolicy.OutlierDetection.BaseEjectionTime)}
			vService.TrafficPolicy.OutlierDetection.ConsecutiveErrors = input.ServiceAttribute.TrafficPolicy.OutlierDetection.ConsecutiveErrors
			vService.TrafficPolicy.OutlierDetection.Interval = &types.Duration{Seconds: int64(input.ServiceAttribute.TrafficPolicy.OutlierDetection.Interval)}
			vService.TrafficPolicy.OutlierDetection.MaxEjectionPercent = input.ServiceAttribute.TrafficPolicy.OutlierDetection.MaxEjectionPercent
			vService.TrafficPolicy.OutlierDetection.MinHealthPercent = input.ServiceAttribute.TrafficPolicy.OutlierDetection.MinHealthPercent
		}
		for _, port := range input.ServiceAttribute.TrafficPolicy.PortLevelSettings {

			setting := &v1alpha3.TrafficPolicy_PortTrafficPolicy{}
			if port.DrPort != nil {
				setting.Port = &v1alpha3.PortSelector{}
				setting.Port.Number = uint32(port.DrPort.Number)
			}
			if port.ConnectionPool != nil {
				setting.ConnectionPool = &v1alpha3.ConnectionPoolSettings{}
				setting.ConnectionPool.Tcp = &v1alpha3.ConnectionPoolSettings_TCPSettings{}
				setting.ConnectionPool.Tcp.ConnectTimeout = &types.Duration{Nanos: port.ConnectionPool.DrTcp.ConnectTimeout}
				setting.ConnectionPool.Tcp.MaxConnections = port.ConnectionPool.DrTcp.MaxConnections
				setting.ConnectionPool.Tcp.TcpKeepalive = &v1alpha3.ConnectionPoolSettings_TCPSettings_TcpKeepalive{}
				setting.ConnectionPool.Tcp.TcpKeepalive.Interval = &types.Duration{Nanos: port.ConnectionPool.DrTcp.TcpKeepAlive.Interval}
				setting.ConnectionPool.Tcp.TcpKeepalive.Probes = port.ConnectionPool.DrTcp.TcpKeepAlive.Probes
				setting.ConnectionPool.Tcp.TcpKeepalive.Time = &types.Duration{Nanos: port.ConnectionPool.DrTcp.TcpKeepAlive.Time}

				setting.ConnectionPool.Http = &v1alpha3.ConnectionPoolSettings_HTTPSettings{}
				setting.ConnectionPool.Http.H2UpgradePolicy = v1alpha3.ConnectionPoolSettings_HTTPSettings_H2UpgradePolicy(port.ConnectionPool.DrHttp.Http_2MaxRequests)
				setting.ConnectionPool.Http.Http1MaxPendingRequests = port.ConnectionPool.DrHttp.Http_1MaxPendingRequests
				setting.ConnectionPool.Http.Http2MaxRequests = int32(port.ConnectionPool.DrHttp.GetHttp_2MaxRequests())
				setting.ConnectionPool.Http.IdleTimeout = &types.Duration{Nanos: port.ConnectionPool.DrHttp.IdleTimeout}
				setting.ConnectionPool.Http.MaxRequestsPerConnection = port.ConnectionPool.DrHttp.MaxRequestsPerConnection
				setting.ConnectionPool.Http.MaxRetries = port.ConnectionPool.DrHttp.MaxRetries
			}
			if port.LoadBalancer != nil {
				setting.LoadBalancer = &v1alpha3.LoadBalancerSettings{}
				if port.LoadBalancer.Simple >= 0 && port.LoadBalancer.Simple <= 3 {
					setting.LoadBalancer.LbPolicy = &v1alpha3.LoadBalancerSettings_Simple{
						Simple: v1alpha3.LoadBalancerSettings_SimpleLB(int32(port.LoadBalancer.Simple)),
					}
				} else if port.LoadBalancer.ConsistentHash != nil {
					if port.LoadBalancer.ConsistentHash.HttpCookie != nil {
						t := time.Duration(port.LoadBalancer.ConsistentHash.HttpCookie.Ttl)
						setting.LoadBalancer.LbPolicy = &v1alpha3.LoadBalancerSettings_ConsistentHash{
							ConsistentHash: &v1alpha3.LoadBalancerSettings_ConsistentHashLB{
								HashKey: &v1alpha3.LoadBalancerSettings_ConsistentHashLB_HttpCookie{
									HttpCookie: &v1alpha3.LoadBalancerSettings_ConsistentHashLB_HTTPCookie{
										Name: port.LoadBalancer.ConsistentHash.HttpCookie.Name,
										Path: port.LoadBalancer.ConsistentHash.HttpCookie.Path,
										Ttl:  &t,
									},
								},
							},
						}
					} else if port.LoadBalancer.ConsistentHash.HttpHeaderName != "" {
						setting.LoadBalancer.LbPolicy = &v1alpha3.LoadBalancerSettings_ConsistentHash{
							ConsistentHash: &v1alpha3.LoadBalancerSettings_ConsistentHashLB{
								HashKey: &v1alpha3.LoadBalancerSettings_ConsistentHashLB_HttpHeaderName{
									HttpHeaderName: port.LoadBalancer.ConsistentHash.HttpHeaderName,
								},
							},
						}
					} else if port.LoadBalancer.ConsistentHash.UseSourceIp != false {
						setting.LoadBalancer.LbPolicy = &v1alpha3.LoadBalancerSettings_ConsistentHash{
							ConsistentHash: &v1alpha3.LoadBalancerSettings_ConsistentHashLB{
								HashKey: &v1alpha3.LoadBalancerSettings_ConsistentHashLB_UseSourceIp{
									UseSourceIp: port.LoadBalancer.ConsistentHash.UseSourceIp,
								},
							},
						}
					}
				}
			}
			if port.OutlierDetection != nil {
				setting.OutlierDetection = &v1alpha3.OutlierDetection{}
				setting.OutlierDetection.BaseEjectionTime = &types.Duration{Nanos: port.OutlierDetection.BaseEjectionTime}
				setting.OutlierDetection.ConsecutiveErrors = port.OutlierDetection.ConsecutiveErrors
				setting.OutlierDetection.Interval = &types.Duration{Nanos: port.OutlierDetection.Interval}
				setting.OutlierDetection.MaxEjectionPercent = port.OutlierDetection.MaxEjectionPercent
				setting.OutlierDetection.MinHealthPercent = port.OutlierDetection.MinHealthPercent
			}

			vService.TrafficPolicy.PortLevelSettings = append(vService.TrafficPolicy.PortLevelSettings, setting)
		}
		if input.ServiceAttribute.TrafficPolicy.DrTls != nil {
			vService.TrafficPolicy.Tls = &v1alpha3.TLSSettings{}
			vService.TrafficPolicy.Tls.Mode = v1alpha3.TLSSettings_TLSmode(input.ServiceAttribute.TrafficPolicy.DrTls.GetMode())
			vService.TrafficPolicy.Tls.ClientCertificate = input.ServiceAttribute.TrafficPolicy.DrTls.ClientCertificate
			vService.TrafficPolicy.Tls.PrivateKey = input.ServiceAttribute.TrafficPolicy.DrTls.PrivateKey
			vService.TrafficPolicy.Tls.CaCertificates = input.ServiceAttribute.TrafficPolicy.DrTls.CaCertificate
			vService.TrafficPolicy.Tls.SubjectAltNames[0] = input.ServiceAttribute.TrafficPolicy.DrTls.SubjectAltNames
		}
	}
	for _, subset := range input.ServiceAttribute.Subsets {
		ser := &v1alpha3.Subset{}
		ser.Name = subset.Name
		ser.Labels = subset.Labels
		if subset.TrafficPolicy != nil {
			ser.TrafficPolicy = &v1alpha3.TrafficPolicy{}
			for _, port := range subset.TrafficPolicy.PortLevelSettings {
				setting := &v1alpha3.TrafficPolicy_PortTrafficPolicy{}
				if port.DrPort != nil {
					setting.Port = &v1alpha3.PortSelector{}
					setting.Port.Number = uint32(port.DrPort.Number)
				}
				if port.ConnectionPool != nil {
					setting.ConnectionPool = &v1alpha3.ConnectionPoolSettings{}
					if port.ConnectionPool.DrTcp != nil {
						setting.ConnectionPool.Tcp = &v1alpha3.ConnectionPoolSettings_TCPSettings{}
						setting.ConnectionPool.Tcp.ConnectTimeout = &types.Duration{Nanos: port.ConnectionPool.DrTcp.ConnectTimeout}
						setting.ConnectionPool.Tcp.MaxConnections = port.ConnectionPool.DrTcp.MaxConnections
						if port.ConnectionPool.DrTcp.TcpKeepAlive != nil {
							setting.ConnectionPool.Tcp.TcpKeepalive = &v1alpha3.ConnectionPoolSettings_TCPSettings_TcpKeepalive{}
							setting.ConnectionPool.Tcp.TcpKeepalive.Interval = &types.Duration{Seconds: int64(port.ConnectionPool.DrTcp.TcpKeepAlive.Interval)}
							setting.ConnectionPool.Tcp.TcpKeepalive.Probes = port.ConnectionPool.DrTcp.TcpKeepAlive.Probes
							setting.ConnectionPool.Tcp.TcpKeepalive.Time = &types.Duration{Seconds: int64(port.ConnectionPool.DrTcp.TcpKeepAlive.Time)}
						}
					}
					if port.ConnectionPool.DrHttp != nil {
						setting.ConnectionPool.Http = &v1alpha3.ConnectionPoolSettings_HTTPSettings{}
						setting.ConnectionPool.Http.H2UpgradePolicy = v1alpha3.ConnectionPoolSettings_HTTPSettings_H2UpgradePolicy(port.ConnectionPool.DrHttp.Http_2MaxRequests)
						setting.ConnectionPool.Http.Http1MaxPendingRequests = port.ConnectionPool.DrHttp.Http_1MaxPendingRequests
						//setting.ConnectionPool.Http.Http2MaxRequests = port.ConnectionPool.DrHttp.Http_2MaxRequests
						setting.ConnectionPool.Http.IdleTimeout = &types.Duration{Nanos: port.ConnectionPool.DrHttp.IdleTimeout}
						setting.ConnectionPool.Http.MaxRequestsPerConnection = port.ConnectionPool.DrHttp.MaxRequestsPerConnection
						setting.ConnectionPool.Http.MaxRetries = port.ConnectionPool.DrHttp.MaxRetries
					}
				}
				if port.LoadBalancer != nil {
					setting.LoadBalancer = &v1alpha3.LoadBalancerSettings{}
					if port.LoadBalancer.Simple >= 0 && port.LoadBalancer.Simple <= 3 {
						setting.LoadBalancer.LbPolicy = &v1alpha3.LoadBalancerSettings_Simple{
							Simple: v1alpha3.LoadBalancerSettings_SimpleLB(int32(port.LoadBalancer.Simple)),
						}
					} else if port.LoadBalancer.ConsistentHash != nil {
						if port.LoadBalancer.ConsistentHash.HttpCookie != nil {
							t := time.Duration(port.LoadBalancer.ConsistentHash.HttpCookie.Ttl)
							setting.LoadBalancer.LbPolicy = &v1alpha3.LoadBalancerSettings_ConsistentHash{
								ConsistentHash: &v1alpha3.LoadBalancerSettings_ConsistentHashLB{
									HashKey: &v1alpha3.LoadBalancerSettings_ConsistentHashLB_HttpCookie{
										HttpCookie: &v1alpha3.LoadBalancerSettings_ConsistentHashLB_HTTPCookie{
											Name: port.LoadBalancer.ConsistentHash.HttpCookie.Name,
											Path: port.LoadBalancer.ConsistentHash.HttpCookie.Path,
											Ttl:  &t,
										},
									},
								},
							}
						} else if port.LoadBalancer.ConsistentHash.HttpHeaderName != "" {
							setting.LoadBalancer.LbPolicy = &v1alpha3.LoadBalancerSettings_ConsistentHash{
								ConsistentHash: &v1alpha3.LoadBalancerSettings_ConsistentHashLB{
									HashKey: &v1alpha3.LoadBalancerSettings_ConsistentHashLB_HttpHeaderName{
										HttpHeaderName: port.LoadBalancer.ConsistentHash.HttpHeaderName,
									},
								},
							}
						} else if port.LoadBalancer.ConsistentHash.UseSourceIp != false {
							setting.LoadBalancer.LbPolicy = &v1alpha3.LoadBalancerSettings_ConsistentHash{
								ConsistentHash: &v1alpha3.LoadBalancerSettings_ConsistentHashLB{
									HashKey: &v1alpha3.LoadBalancerSettings_ConsistentHashLB_UseSourceIp{
										UseSourceIp: port.LoadBalancer.ConsistentHash.UseSourceIp,
									},
								},
							}
						}
					}
				}
				if port.OutlierDetection != nil {
					setting.OutlierDetection = &v1alpha3.OutlierDetection{}
					setting.OutlierDetection.BaseEjectionTime = &types.Duration{Nanos: port.OutlierDetection.BaseEjectionTime}
					setting.OutlierDetection.ConsecutiveErrors = port.OutlierDetection.ConsecutiveErrors
					setting.OutlierDetection.Interval = &types.Duration{Nanos: port.OutlierDetection.Interval}
					setting.OutlierDetection.MaxEjectionPercent = port.OutlierDetection.MaxEjectionPercent
					setting.OutlierDetection.MinHealthPercent = port.OutlierDetection.MinHealthPercent
				}
				if port.DrTls != nil {
					setting.Tls = &v1alpha3.TLSSettings{}
					vService.TrafficPolicy.Tls.Mode = v1alpha3.TLSSettings_TLSmode(port.DrTls.GetMode())
					setting.Tls.ClientCertificate = port.DrTls.ClientCertificate
					setting.Tls.PrivateKey = port.DrTls.PrivateKey
					setting.Tls.CaCertificates = port.DrTls.CaCertificate
					setting.Tls.SubjectAltNames = append(setting.Tls.SubjectAltNames, port.DrTls.SubjectAltNames)
					vService.TrafficPolicy.PortLevelSettings = append(vService.TrafficPolicy.PortLevelSettings, setting)
				}
			}
			if subset.TrafficPolicy.LoadBalancer != nil {
				ser.TrafficPolicy.LoadBalancer = &v1alpha3.LoadBalancerSettings{}
				if (subset.TrafficPolicy.LoadBalancer.Simple >= 0 && subset.TrafficPolicy.LoadBalancer.Simple <= 3) && subset.TrafficPolicy.LoadBalancer.ConsistentHash == nil {
					ser.TrafficPolicy.LoadBalancer.LbPolicy = &v1alpha3.LoadBalancerSettings_Simple{
						Simple: v1alpha3.LoadBalancerSettings_SimpleLB(int32(subset.TrafficPolicy.LoadBalancer.Simple)),
					}
				} else if subset.TrafficPolicy.LoadBalancer.ConsistentHash != nil {
					if subset.TrafficPolicy.LoadBalancer.ConsistentHash.HttpCookie != nil {
						t := time.Duration(subset.TrafficPolicy.LoadBalancer.ConsistentHash.HttpCookie.Ttl)
						ser.TrafficPolicy.LoadBalancer.LbPolicy = &v1alpha3.LoadBalancerSettings_ConsistentHash{
							ConsistentHash: &v1alpha3.LoadBalancerSettings_ConsistentHashLB{
								HashKey: &v1alpha3.LoadBalancerSettings_ConsistentHashLB_HttpCookie{
									HttpCookie: &v1alpha3.LoadBalancerSettings_ConsistentHashLB_HTTPCookie{
										Name: subset.TrafficPolicy.LoadBalancer.ConsistentHash.HttpCookie.Name,
										Path: subset.TrafficPolicy.LoadBalancer.ConsistentHash.HttpCookie.Path,
										Ttl:  &t,
									},
								},
							},
						}
					} else if subset.TrafficPolicy.LoadBalancer.ConsistentHash.HttpHeaderName != "" {
						ser.TrafficPolicy.LoadBalancer.LbPolicy = &v1alpha3.LoadBalancerSettings_ConsistentHash{
							ConsistentHash: &v1alpha3.LoadBalancerSettings_ConsistentHashLB{
								HashKey: &v1alpha3.LoadBalancerSettings_ConsistentHashLB_HttpHeaderName{
									HttpHeaderName: subset.TrafficPolicy.LoadBalancer.ConsistentHash.HttpHeaderName,
								},
							},
						}
					} else if subset.TrafficPolicy.LoadBalancer.ConsistentHash.UseSourceIp != false {
						ser.TrafficPolicy.LoadBalancer.LbPolicy = &v1alpha3.LoadBalancerSettings_ConsistentHash{
							ConsistentHash: &v1alpha3.LoadBalancerSettings_ConsistentHashLB{
								HashKey: &v1alpha3.LoadBalancerSettings_ConsistentHashLB_UseSourceIp{
									UseSourceIp: subset.TrafficPolicy.LoadBalancer.ConsistentHash.UseSourceIp,
								},
							},
						}
					}
				}
			}
			if subset.TrafficPolicy.ConnectionPool != nil {
				ser.TrafficPolicy.ConnectionPool = &v1alpha3.ConnectionPoolSettings{}
				if subset.TrafficPolicy.ConnectionPool.DrTcp != nil {
					ser.TrafficPolicy.ConnectionPool.Tcp = &v1alpha3.ConnectionPoolSettings_TCPSettings{}
					ser.TrafficPolicy.ConnectionPool.Tcp.ConnectTimeout = &types.Duration{Nanos: subset.TrafficPolicy.ConnectionPool.DrTcp.ConnectTimeout}
					ser.TrafficPolicy.ConnectionPool.Tcp.MaxConnections = subset.TrafficPolicy.ConnectionPool.DrTcp.MaxConnections
					if subset.TrafficPolicy.ConnectionPool.DrTcp.TcpKeepAlive != nil {
						ser.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive = &v1alpha3.ConnectionPoolSettings_TCPSettings_TcpKeepalive{}
						ser.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive.Interval = &types.Duration{Seconds: int64(subset.TrafficPolicy.ConnectionPool.DrTcp.TcpKeepAlive.Interval)}
						ser.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive.Probes = subset.TrafficPolicy.ConnectionPool.DrTcp.TcpKeepAlive.Probes
						ser.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive.Time = &types.Duration{Seconds: int64(subset.TrafficPolicy.ConnectionPool.DrTcp.TcpKeepAlive.Time)}
					}
				}
				if subset.TrafficPolicy.ConnectionPool.DrHttp != nil {
					ser.TrafficPolicy.ConnectionPool.Http = &v1alpha3.ConnectionPoolSettings_HTTPSettings{}
					ser.TrafficPolicy.ConnectionPool.Http.H2UpgradePolicy = v1alpha3.ConnectionPoolSettings_HTTPSettings_H2UpgradePolicy(subset.TrafficPolicy.ConnectionPool.DrHttp.Http_2MaxRequests)
					ser.TrafficPolicy.ConnectionPool.Http.Http1MaxPendingRequests = subset.TrafficPolicy.ConnectionPool.DrHttp.Http_1MaxPendingRequests
					ser.TrafficPolicy.ConnectionPool.Http.Http2MaxRequests = int32(subset.TrafficPolicy.ConnectionPool.DrHttp.GetHttp_2MaxRequests())
					ser.TrafficPolicy.ConnectionPool.Http.IdleTimeout = &types.Duration{Nanos: subset.TrafficPolicy.ConnectionPool.DrHttp.IdleTimeout}
					ser.TrafficPolicy.ConnectionPool.Http.MaxRequestsPerConnection = subset.TrafficPolicy.ConnectionPool.DrHttp.MaxRequestsPerConnection
					ser.TrafficPolicy.ConnectionPool.Http.MaxRetries = subset.TrafficPolicy.ConnectionPool.DrHttp.MaxRetries
				}
			}
			if subset.TrafficPolicy.OutlierDetection != nil {
				ser.TrafficPolicy.OutlierDetection = &v1alpha3.OutlierDetection{}
				ser.TrafficPolicy.OutlierDetection.BaseEjectionTime = &types.Duration{Nanos: subset.TrafficPolicy.OutlierDetection.BaseEjectionTime}
				ser.TrafficPolicy.OutlierDetection.ConsecutiveErrors = subset.TrafficPolicy.OutlierDetection.ConsecutiveErrors
				ser.TrafficPolicy.OutlierDetection.Interval = &types.Duration{Nanos: subset.TrafficPolicy.OutlierDetection.Interval}
				ser.TrafficPolicy.OutlierDetection.MaxEjectionPercent = subset.TrafficPolicy.OutlierDetection.MaxEjectionPercent
				ser.TrafficPolicy.OutlierDetection.MinHealthPercent = subset.TrafficPolicy.OutlierDetection.MinHealthPercent
			}
			if subset.TrafficPolicy.DrTls != nil {
				ser.TrafficPolicy.Tls = &v1alpha3.TLSSettings{}
				vService.TrafficPolicy.Tls.Mode = v1alpha3.TLSSettings_TLSmode(int32(subset.TrafficPolicy.DrTls.GetMode()))
				ser.TrafficPolicy.Tls.ClientCertificate = subset.TrafficPolicy.DrTls.ClientCertificate
				ser.TrafficPolicy.Tls.PrivateKey = subset.TrafficPolicy.DrTls.PrivateKey
				ser.TrafficPolicy.Tls.CaCertificates = subset.TrafficPolicy.DrTls.CaCertificate
				ser.TrafficPolicy.Tls.SubjectAltNames = append(ser.TrafficPolicy.Tls.SubjectAltNames, subset.TrafficPolicy.DrTls.SubjectAltNames)
			}
		}
		vService.Subsets = append(vService.Subsets, ser)
	}

	vServ.Spec = vService

	return vServ, nil
}

func getDRRequestObject(req *pb.DestinationRules) (*istioClient.DestinationRule, error) {
	drReq, err := getDestinationRules(req)
	if err != nil {
		utils.Error.Println(err)

		return nil, err
	}
	return drReq, nil
}
