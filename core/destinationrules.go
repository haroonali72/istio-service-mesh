package core

import (
	"context"
	"encoding/json"
	"github.com/gogo/protobuf/types"
	"time"

	//	types "github.com/gogo/protobuf/types"
	"google.golang.org/grpc"
	"istio-service-mesh/constants"
	pb "istio-service-mesh/core/proto"
	"istio-service-mesh/utils"
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

func getDestinationRules(input *pb.DestinationRules) (*istioClient.DestinationRule, error) {

	var vServ = new(istioClient.DestinationRule)
	labels := make(map[string]string)
	labels["app"] = strings.ToLower(input.Name)
	labels["version"] = strings.ToLower("networking.v.io/v1alpha3")
	vServ.Labels = labels
	vServ.Kind = "DestinationRule"
	vServ.APIVersion = "networking.v.io/v1alpha3"
	vServ.Name = input.Name
	vServ.Namespace = input.Namespace

	vService := v1alpha3.DestinationRule{}
	vService.Host = input.ServiceAttribute.Host
	vService.TrafficPolicy = &v1alpha3.TrafficPolicy{}
	vService.TrafficPolicy.LoadBalancer = &v1alpha3.LoadBalancerSettings{}
	//if input.ServiceAttributes.TrafficPolicy.LoadBalancer.Simple != nil {
	//	vService.TrafficPolicy.LoadBalancer.LbPolicy = &v1alpha3.LoadBalancerSettings_Simple{
	//  Simple: v1alpha3.LoadBalancerSettings_SimpleLB(input.ServiceAttribute.TrafficPolicy.LoadBalancer.Simple),
	//	}
	//} else if input.ServiceAttributes.TrafficPolicy.LoadBalancer.ConsistentHash != null {
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

	vService.TrafficPolicy.ConnectionPool = &v1alpha3.ConnectionPoolSettings{}
	vService.TrafficPolicy.ConnectionPool.Tcp = &v1alpha3.ConnectionPoolSettings_TCPSettings{}

	vService.TrafficPolicy.ConnectionPool.Tcp.ConnectTimeout = &types.Duration{Nanos: input.ServiceAttribute.TrafficPolicy.ConnectionPool.DrTcp.ConnectTimeout}
	vService.TrafficPolicy.ConnectionPool.Tcp.MaxConnections = input.ServiceAttribute.TrafficPolicy.ConnectionPool.DrTcp.MaxConnections

	vService.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive = &v1alpha3.ConnectionPoolSettings_TCPSettings_TcpKeepalive{}
	vService.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive.Interval = &types.Duration{Nanos: input.ServiceAttribute.TrafficPolicy.ConnectionPool.DrTcp.TcpKeepAlive.Interval}
	vService.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive.Probes = input.ServiceAttribute.TrafficPolicy.ConnectionPool.DrTcp.TcpKeepAlive.Probes
	vService.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive.Time = &types.Duration{Nanos: input.ServiceAttribute.TrafficPolicy.ConnectionPool.DrTcp.TcpKeepAlive.Time}

	vService.TrafficPolicy.ConnectionPool.Http = &v1alpha3.ConnectionPoolSettings_HTTPSettings{}
	//	vService.TrafficPolicy.ConnectionPool.Http.H2UpgradePolicy = input.ServiceAttribute.Subsets.TrafficPolicy.ConnectionPool.DrHttp.Http_2MaxRequests
	vService.TrafficPolicy.ConnectionPool.Http.Http1MaxPendingRequests = input.ServiceAttribute.Subsets.TrafficPolicy.ConnectionPool.DrHttp.Http_1MaxPendingRequests
	vService.TrafficPolicy.ConnectionPool.Http.Http2MaxRequests = input.ServiceAttribute.Subsets.TrafficPolicy.ConnectionPool.DrHttp.Http_2MaxRequests
	vService.TrafficPolicy.ConnectionPool.Http.IdleTimeout = &types.Duration{Nanos: input.ServiceAttribute.Subsets.TrafficPolicy.ConnectionPool.DrHttp.IdleTimeout}
	vService.TrafficPolicy.ConnectionPool.Http.MaxRequestsPerConnection = input.ServiceAttribute.Subsets.TrafficPolicy.ConnectionPool.DrHttp.MaxRequestsPerConnection
	vService.TrafficPolicy.ConnectionPool.Http.MaxRetries = input.ServiceAttribute.Subsets.TrafficPolicy.ConnectionPool.DrHttp.MaxRetries

	vService.TrafficPolicy.OutlierDetection = &v1alpha3.OutlierDetection{}
	vService.TrafficPolicy.OutlierDetection.BaseEjectionTime = &types.Duration{Nanos: input.ServiceAttribute.TrafficPolicy.OutlierDetection.BaseEjectionTime}
	vService.TrafficPolicy.OutlierDetection.ConsecutiveErrors = input.ServiceAttribute.TrafficPolicy.OutlierDetection.ConsecutiveErrors
	vService.TrafficPolicy.OutlierDetection.Interval = &types.Duration{Nanos: input.ServiceAttribute.TrafficPolicy.OutlierDetection.Interval}
	vService.TrafficPolicy.OutlierDetection.MaxEjectionPercent = input.ServiceAttribute.TrafficPolicy.OutlierDetection.MaxEjectionPercent
	vService.TrafficPolicy.OutlierDetection.MinHealthPercent = input.ServiceAttribute.TrafficPolicy.OutlierDetection.MinHealthPercent

	for _, port := range input.ServiceAttribute.TrafficPolicy.PortLevelSettings {

		setting := &v1alpha3.TrafficPolicy_PortTrafficPolicy{}
		setting.Port = &v1alpha3.PortSelector{}
		setting.Port.Number = uint32(port.DrPort.Number)

		setting.ConnectionPool = &v1alpha3.ConnectionPoolSettings{}
		setting.ConnectionPool.Tcp = &v1alpha3.ConnectionPoolSettings_TCPSettings{}
		setting.ConnectionPool.Tcp.ConnectTimeout = &types.Duration{Nanos: port.ConnectionPool.DrTcp.ConnectTimeout}
		setting.ConnectionPool.Tcp.MaxConnections = port.ConnectionPool.DrTcp.MaxConnections
		setting.ConnectionPool.Tcp.TcpKeepalive = &v1alpha3.ConnectionPoolSettings_TCPSettings_TcpKeepalive{}
		setting.ConnectionPool.Tcp.TcpKeepalive.Interval = &types.Duration{Nanos: port.ConnectionPool.DrTcp.TcpKeepAlive.Interval}
		setting.ConnectionPool.Tcp.TcpKeepalive.Probes = port.ConnectionPool.DrTcp.TcpKeepAlive.Probes
		setting.ConnectionPool.Tcp.TcpKeepalive.Time = &types.Duration{Nanos: port.ConnectionPool.DrTcp.TcpKeepAlive.Time}

		setting.ConnectionPool.Http = &v1alpha3.ConnectionPoolSettings_HTTPSettings{}
		//	setting.ConnectionPool.Http.H2UpgradePolicy = port.ConnectionPool.DrHttp.Http_2MaxRequests
		setting.ConnectionPool.Http.Http1MaxPendingRequests = port.ConnectionPool.DrHttp.Http_1MaxPendingRequests
		setting.ConnectionPool.Http.Http2MaxRequests = port.ConnectionPool.DrHttp.Http_2MaxRequests
		setting.ConnectionPool.Http.IdleTimeout = &types.Duration{Nanos: port.ConnectionPool.DrHttp.IdleTimeout}
		setting.ConnectionPool.Http.MaxRequestsPerConnection = port.ConnectionPool.DrHttp.MaxRequestsPerConnection
		setting.ConnectionPool.Http.MaxRetries = port.ConnectionPool.DrHttp.MaxRetries

		setting.LoadBalancer = &v1alpha3.LoadBalancerSettings{}
		//if port.LoadBalancer.Simple != nil {
		//	setting.LoadBalancer.LbPolicy = &v1alpha3.LoadBalancerSettings_Simple{
		//  Simple: v1alpha3.LoadBalancerSettings_SimpleLB(port.LoadBalancer.Simple),
		//	}
		//} else if port.LoadBalancer.ConsistentHash != null {
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

		setting.OutlierDetection = &v1alpha3.OutlierDetection{}
		setting.OutlierDetection.BaseEjectionTime = &types.Duration{Nanos: port.OutlierDetection.BaseEjectionTime}
		setting.OutlierDetection.ConsecutiveErrors = port.OutlierDetection.ConsecutiveErrors
		setting.OutlierDetection.Interval = &types.Duration{Nanos: port.OutlierDetection.Interval}
		setting.OutlierDetection.MaxEjectionPercent = port.OutlierDetection.MaxEjectionPercent
		setting.OutlierDetection.MinHealthPercent = port.OutlierDetection.MinHealthPercent

		vService.TrafficPolicy.PortLevelSettings = append(vService.TrafficPolicy.PortLevelSettings, setting)
	}

	//vService.TrafficPolicy.Tls =&v1alpha3.TLSSettings{}
	//vService.TrafficPolicy.Tls.Mode = &v1alpha3.TLSSettings_TLSmode()
	//TLSSettings_TLSmode
	//vService.TrafficPolicy.Tls.Mode = v1alpha3.TLSSettings_TLSmode(input.ServiceAttributes.TrafficPolicy.DrTls.Mode)
	//	vService.TrafficPolicy.Tls.ClientCertificate = input.ServiceAttribute.TrafficPolicy.DrTls.ClientCertificate
	//	vService.TrafficPolicy.Tls.PrivateKey = input.ServiceAttribute.TrafficPolicy.DrTls.PrivateKey
	//	vService.TrafficPolicy.Tls.CaCertificates = input.ServiceAttribute.TrafficPolicy.DrTls.CaCertificate
	//vService.TrafficPolicy.Tls.Sni=input.ServiceAttributes.TrafficPolicy.Tls.Name
	//	vService.TrafficPolicy.Tls.SubjectAltNames[0] = input.ServiceAttribute.TrafficPolicy.DrTls.SubjectAltNames

	//	for i, subset := range input.ServiceAttributes.Subsets {
	ser := &v1alpha3.Subset{}
	//for _, n := range input.ServiceAttribute.Subsets.Name {
	ser.Name = input.ServiceAttribute.Subsets.Name[0]
	//	}

	ser.Labels = input.ServiceAttribute.Subsets.Labels
	//traffic policy

	ser.TrafficPolicy = &v1alpha3.TrafficPolicy{}

	for _, port := range input.ServiceAttribute.Subsets.TrafficPolicy.PortLevelSettings {

		setting := &v1alpha3.TrafficPolicy_PortTrafficPolicy{}
		setting.Port = &v1alpha3.PortSelector{}
		setting.Port.Number = uint32(port.DrPort.Number)

		setting.ConnectionPool = &v1alpha3.ConnectionPoolSettings{}
		setting.ConnectionPool.Tcp = &v1alpha3.ConnectionPoolSettings_TCPSettings{}
		setting.ConnectionPool.Tcp.ConnectTimeout = &types.Duration{Nanos: port.ConnectionPool.DrTcp.ConnectTimeout}
		setting.ConnectionPool.Tcp.MaxConnections = port.ConnectionPool.DrTcp.MaxConnections
		setting.ConnectionPool.Tcp.TcpKeepalive = &v1alpha3.ConnectionPoolSettings_TCPSettings_TcpKeepalive{}
		setting.ConnectionPool.Tcp.TcpKeepalive.Interval = &types.Duration{Nanos: port.ConnectionPool.DrTcp.TcpKeepAlive.Interval}
		setting.ConnectionPool.Tcp.TcpKeepalive.Probes = port.ConnectionPool.DrTcp.TcpKeepAlive.Probes
		setting.ConnectionPool.Tcp.TcpKeepalive.Time = &types.Duration{Nanos: port.ConnectionPool.DrTcp.TcpKeepAlive.Time}

		setting.ConnectionPool.Http = &v1alpha3.ConnectionPoolSettings_HTTPSettings{}
		//	setting.ConnectionPool.Http.H2UpgradePolicy = port.ConnectionPool.DrHttp.Http_2MaxRequests
		setting.ConnectionPool.Http.Http1MaxPendingRequests = port.ConnectionPool.DrHttp.Http_1MaxPendingRequests
		setting.ConnectionPool.Http.Http2MaxRequests = port.ConnectionPool.DrHttp.Http_2MaxRequests
		setting.ConnectionPool.Http.IdleTimeout = &types.Duration{Nanos: port.ConnectionPool.DrHttp.IdleTimeout}
		setting.ConnectionPool.Http.MaxRequestsPerConnection = port.ConnectionPool.DrHttp.MaxRequestsPerConnection
		setting.ConnectionPool.Http.MaxRetries = port.ConnectionPool.DrHttp.MaxRetries

		setting.LoadBalancer = &v1alpha3.LoadBalancerSettings{}
		//if port.LoadBalancer.Simple != nil {
		//	setting.LoadBalancer.LbPolicy = &v1alpha3.LoadBalancerSettings_Simple{
		//  Simple: v1alpha3.LoadBalancerSettings_SimpleLB(port.LoadBalancer.Simple),
		//	}
		//} else if port.LoadBalancer.ConsistentHash != null {
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

		setting.OutlierDetection = &v1alpha3.OutlierDetection{}
		setting.OutlierDetection.BaseEjectionTime = &types.Duration{Nanos: port.OutlierDetection.BaseEjectionTime}
		setting.OutlierDetection.ConsecutiveErrors = port.OutlierDetection.ConsecutiveErrors
		setting.OutlierDetection.Interval = &types.Duration{Nanos: port.OutlierDetection.Interval}
		setting.OutlierDetection.MaxEjectionPercent = port.OutlierDetection.MaxEjectionPercent
		setting.OutlierDetection.MinHealthPercent = port.OutlierDetection.MinHealthPercent

		setting.Tls = &v1alpha3.TLSSettings{}

		//TLSSettings_TLSmode
		//vService.TrafficPolicy.Tls.Mode = v1alpha3.TLSSettings_TLSmode(input.ServiceAttributes.TrafficPolicy.DrTls.Mode)
		setting.Tls.ClientCertificate = port.DrTls.ClientCertificate
		setting.Tls.PrivateKey = port.DrTls.PrivateKey
		setting.Tls.CaCertificates = port.DrTls.CaCertificate
		//setting.Tls.Sni=port.DrTls.Name
		setting.Tls.SubjectAltNames = append(setting.Tls.SubjectAltNames, port.DrTls.SubjectAltNames)

		vService.TrafficPolicy.PortLevelSettings = append(vService.TrafficPolicy.PortLevelSettings, setting)
	}

	ser.TrafficPolicy.LoadBalancer = &v1alpha3.LoadBalancerSettings{}
	//if input.ServiceAttributes.Subsets.TrafficPolicy.LoadBalancer.Simple != nil {
	//	ser.TrafficPolicy.LoadBalancer.LbPolicy = &v1alpha3.LoadBalancerSettings_Simple{
	//  Simple: v1alpha3.LoadBalancerSettings_SimpleLB(input.ServiceAttribute.Subsets.TrafficPolicy.LoadBalancer.Simple),
	//	}
	//} else if input.ServiceAttributes.TrafficPolicy.LoadBalancer.ConsistentHash != null {
	if input.ServiceAttribute.Subsets.TrafficPolicy.LoadBalancer.ConsistentHash.HttpCookie != nil {
		t := time.Duration(input.ServiceAttribute.Subsets.TrafficPolicy.LoadBalancer.ConsistentHash.HttpCookie.Ttl)
		ser.TrafficPolicy.LoadBalancer.LbPolicy = &v1alpha3.LoadBalancerSettings_ConsistentHash{
			ConsistentHash: &v1alpha3.LoadBalancerSettings_ConsistentHashLB{
				HashKey: &v1alpha3.LoadBalancerSettings_ConsistentHashLB_HttpCookie{
					HttpCookie: &v1alpha3.LoadBalancerSettings_ConsistentHashLB_HTTPCookie{
						Name: input.ServiceAttribute.Subsets.TrafficPolicy.LoadBalancer.ConsistentHash.HttpCookie.Name,
						Path: input.ServiceAttribute.Subsets.TrafficPolicy.LoadBalancer.ConsistentHash.HttpCookie.Path,
						Ttl:  &t,
					},
				},
			},
		}
	} else if input.ServiceAttribute.Subsets.TrafficPolicy.LoadBalancer.ConsistentHash.HttpHeaderName != "" {
		ser.TrafficPolicy.LoadBalancer.LbPolicy = &v1alpha3.LoadBalancerSettings_ConsistentHash{
			ConsistentHash: &v1alpha3.LoadBalancerSettings_ConsistentHashLB{
				HashKey: &v1alpha3.LoadBalancerSettings_ConsistentHashLB_HttpHeaderName{
					HttpHeaderName: input.ServiceAttribute.Subsets.TrafficPolicy.LoadBalancer.ConsistentHash.HttpHeaderName,
				},
			},
		}
	} else if input.ServiceAttribute.Subsets.TrafficPolicy.LoadBalancer.ConsistentHash.UseSourceIp != false {
		ser.TrafficPolicy.LoadBalancer.LbPolicy = &v1alpha3.LoadBalancerSettings_ConsistentHash{
			ConsistentHash: &v1alpha3.LoadBalancerSettings_ConsistentHashLB{
				HashKey: &v1alpha3.LoadBalancerSettings_ConsistentHashLB_UseSourceIp{
					UseSourceIp: input.ServiceAttribute.Subsets.TrafficPolicy.LoadBalancer.ConsistentHash.UseSourceIp,
				},
			},
		}
	}

	ser.TrafficPolicy.ConnectionPool = &v1alpha3.ConnectionPoolSettings{}
	ser.TrafficPolicy.ConnectionPool.Tcp = &v1alpha3.ConnectionPoolSettings_TCPSettings{}

	ser.TrafficPolicy.ConnectionPool.Tcp.ConnectTimeout = &types.Duration{Nanos: input.ServiceAttribute.Subsets.TrafficPolicy.ConnectionPool.DrTcp.ConnectTimeout}
	ser.TrafficPolicy.ConnectionPool.Tcp.MaxConnections = input.ServiceAttribute.Subsets.TrafficPolicy.ConnectionPool.DrTcp.MaxConnections

	ser.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive = &v1alpha3.ConnectionPoolSettings_TCPSettings_TcpKeepalive{}
	ser.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive.Interval = &types.Duration{Nanos: input.ServiceAttribute.Subsets.TrafficPolicy.ConnectionPool.DrTcp.TcpKeepAlive.Interval}
	ser.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive.Probes = input.ServiceAttribute.Subsets.TrafficPolicy.ConnectionPool.DrTcp.TcpKeepAlive.Probes
	ser.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive.Time = &types.Duration{Nanos: input.ServiceAttribute.Subsets.TrafficPolicy.ConnectionPool.DrTcp.TcpKeepAlive.Time}
	///http
	ser.TrafficPolicy.ConnectionPool.Http = &v1alpha3.ConnectionPoolSettings_HTTPSettings{}
	//	ser.TrafficPolicy.ConnectionPool.Http.H2UpgradePolicy = input.ServiceAttribute.Subsets.TrafficPolicy.ConnectionPool.DrHttp.Http_2MaxRequests
	ser.TrafficPolicy.ConnectionPool.Http.Http1MaxPendingRequests = input.ServiceAttribute.Subsets.TrafficPolicy.ConnectionPool.DrHttp.Http_1MaxPendingRequests
	ser.TrafficPolicy.ConnectionPool.Http.Http2MaxRequests = input.ServiceAttribute.Subsets.TrafficPolicy.ConnectionPool.DrHttp.Http_2MaxRequests
	ser.TrafficPolicy.ConnectionPool.Http.IdleTimeout = &types.Duration{Nanos: input.ServiceAttribute.Subsets.TrafficPolicy.ConnectionPool.DrHttp.IdleTimeout}
	ser.TrafficPolicy.ConnectionPool.Http.MaxRequestsPerConnection = input.ServiceAttribute.Subsets.TrafficPolicy.ConnectionPool.DrHttp.MaxRequestsPerConnection
	ser.TrafficPolicy.ConnectionPool.Http.MaxRetries = input.ServiceAttribute.Subsets.TrafficPolicy.ConnectionPool.DrHttp.MaxRetries

	ser.TrafficPolicy.OutlierDetection = &v1alpha3.OutlierDetection{}

	ser.TrafficPolicy.OutlierDetection.BaseEjectionTime = &types.Duration{Nanos: input.ServiceAttribute.Subsets.TrafficPolicy.OutlierDetection.BaseEjectionTime}
	ser.TrafficPolicy.OutlierDetection.ConsecutiveErrors = input.ServiceAttribute.Subsets.TrafficPolicy.OutlierDetection.ConsecutiveErrors
	ser.TrafficPolicy.OutlierDetection.Interval = &types.Duration{Nanos: input.ServiceAttribute.Subsets.TrafficPolicy.OutlierDetection.Interval}
	ser.TrafficPolicy.OutlierDetection.MaxEjectionPercent = input.ServiceAttribute.Subsets.TrafficPolicy.OutlierDetection.MaxEjectionPercent
	ser.TrafficPolicy.OutlierDetection.MinHealthPercent = input.ServiceAttribute.Subsets.TrafficPolicy.OutlierDetection.MinHealthPercent

	ser.TrafficPolicy.Tls = &v1alpha3.TLSSettings{}

	//TLSSettings_TLSmode
	//vService.TrafficPolicy.Tls.Mode = v1alpha3.TLSSettings_TLSmode(input.ServiceAttributes.TrafficPolicy.DrTls.Mode)
	ser.TrafficPolicy.Tls.ClientCertificate = input.ServiceAttribute.Subsets.TrafficPolicy.DrTls.ClientCertificate
	ser.TrafficPolicy.Tls.PrivateKey = input.ServiceAttribute.Subsets.TrafficPolicy.DrTls.PrivateKey
	ser.TrafficPolicy.Tls.CaCertificates = input.ServiceAttribute.Subsets.TrafficPolicy.DrTls.CaCertificate
	//ser.TrafficPolicy.Tls.Sni=input.ServiceAttribute.Subsets.TrafficPolicy.DrTls.Name
	ser.TrafficPolicy.Tls.SubjectAltNames = append(ser.TrafficPolicy.Tls.SubjectAltNames, input.ServiceAttribute.Subsets.TrafficPolicy.DrTls.SubjectAltNames)

	vService.Subsets = append(vService.Subsets, ser)
	vServ.Spec = vService
	//}
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
