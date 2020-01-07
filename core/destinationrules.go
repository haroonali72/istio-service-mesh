package core

import (
	"context"
	"encoding/json"
	"github.com/gogo/protobuf/types"

	//	types "github.com/gogo/protobuf/types"
	"google.golang.org/grpc"
	"istio-service-mesh/constants"
	pb "istio-service-mesh/core/proto"
	"istio-service-mesh/utils"
	"istio.io/api/networking/v1alpha3"
	istioClient "istio.io/client-go/pkg/apis/networking/v1alpha3"
	"strings"
	time "time"
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
	labels["version"] = strings.ToLower("version")
	vServ.Labels = labels
	vServ.Kind = "DestinationRules"
	vServ.APIVersion = "networking.v.io/v1alpha3"
	vServ.Name = input.Name
	vServ.Namespace = input.Namespace

	vService := v1alpha3.DestinationRule{}
	vService.Host = input.ServiceAttributes.Host

	if input.ServiceAttributes.TrafficPolicy.LoadBalancer.Simple != 0 {
		vService.TrafficPolicy.LoadBalancer.LbPolicy = &v1alpha3.LoadBalancerSettings_Simple{
			Simple: v1alpha3.LoadBalancerSettings_SimpleLB(input.ServiceAttributes.TrafficPolicy.LoadBalancer.Simple),
		}
	} else if input.ServiceAttributes.TrafficPolicy.LoadBalancer.ConsistentHash != null {
		vService.TrafficPolicy.LoadBalancer.LbPolicy = &v1alpha3.LoadBalancerSettings_ConsistentHash{
			ConsistentHash: &v1alpha3.LoadBalancerSettings_ConsistentHashLB{
				HashKey: &v1alpha3.LoadBalancerSettings_ConsistentHashLB_HttpCookie{
					HttpCookie: &v1alpha3.LoadBalancerSettings_ConsistentHashLB_HTTPCookie{
						Name: input.ServiceAttributes.TrafficPolicy.LoadBalancer.ConsistentHash.HttpCookie.Name,
						Path: input.ServiceAttributes.TrafficPolicy.LoadBalancer.ConsistentHash.HttpCookie.Path,
						Ttl:  &time.Duration(input.ServiceAttributes.TrafficPolicy.LoadBalancer.ConsistentHash.HttpCookie.Ttl),
					},
				},
			},
		}
	} else if input.ServiceAttributes.TrafficPolicy.LoadBalancer.ConsistentHash != null {
		vService.TrafficPolicy.LoadBalancer.LbPolicy = &v1alpha3.LoadBalancerSettings_ConsistentHash{
			ConsistentHash: &v1alpha3.LoadBalancerSettings_ConsistentHashLB{
				HashKey: &v1alpha3.LoadBalancerSettings_ConsistentHashLB_HttpHeaderName{
					HttpHeaderName: input.ServiceAttributes.TrafficPolicy.LoadBalancer.ConsistentHash.HttpHeaderName,
				},
			},
		}
	} else if input.ServiceAttributes.TrafficPolicy.LoadBalancer.ConsistentHash != null {
		vService.TrafficPolicy.LoadBalancer.LbPolicy = &v1alpha3.LoadBalancerSettings_ConsistentHash{
			ConsistentHash: &v1alpha3.LoadBalancerSettings_ConsistentHashLB{
				HashKey: &v1alpha3.LoadBalancerSettings_ConsistentHashLB_UseSourceIp{
					UseSourceIp: input.ServiceAttributes.TrafficPolicy.LoadBalancer.ConsistentHash.UseSourceIp,
				},
			},
		}
	}

	vService.TrafficPolicy.ConnectionPool.Tcp.ConnectTimeout = &types.Duration{Seconds: input.ServiceAttributes.TrafficPolicy.ConnectionPool.Tcp.ConnectTimeout}
	vService.TrafficPolicy.ConnectionPool.Tcp.MaxConnections = input.ServiceAttributes.TrafficPolicy.ConnectionPool.Tcp.MaxConnections
	vService.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive.Interval = &types.Duration{Seconds: input.ServiceAttributes.TrafficPolicy.ConnectionPool.Tcp.TcpKeepAlive.Interval}
	vService.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive.Probes = input.ServiceAttributes.TrafficPolicy.ConnectionPool.Tcp.TcpKeepAlive.Probes
	vService.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive.Time = &types.Duration{Seconds: input.ServiceAttributes.TrafficPolicy.ConnectionPool.Tcp.TcpKeepAlive.Time}

	vService.TrafficPolicy.OutlierDetection.BaseEjectionTime = &types.Duration{Seconds: input.ServiceAttributes.TrafficPolicy.OutlierDetection.BaseEjectionTime}
	vService.TrafficPolicy.OutlierDetection.ConsecutiveErrors = input.ServiceAttributes.TrafficPolicy.OutlierDetection.ConsecutiveErrors
	vService.TrafficPolicy.OutlierDetection.Interval = &types.Duration{Seconds: input.ServiceAttributes.TrafficPolicy.OutlierDetection.Interval}
	vService.TrafficPolicy.OutlierDetection.MaxEjectionPercent = input.ServiceAttributes.TrafficPolicy.OutlierDetection.MaxEjectionPercent
	vService.TrafficPolicy.OutlierDetection.MinHealthPercent = input.ServiceAttributes.TrafficPolicy.OutlierDetection.MinHealthPercent

	vService.TrafficPolicy.Tls.Mode = v1alpha3.TLSSettings_TLSmode(input.ServiceAttributes.TrafficPolicy.Tls.Mode)
	vService.TrafficPolicy.Tls.ClientCertificate = input.ServiceAttributes.TrafficPolicy.Tls.ClientCertificate
	vService.TrafficPolicy.Tls.PrivateKey = input.ServiceAttributes.TrafficPolicy.Tls.PrivateKey
	vService.TrafficPolicy.Tls.CaCertificates = input.ServiceAttributes.TrafficPolicy.Tls.CaCertificate
	//vService.TrafficPolicy.Tls.Sni=input.ServiceAttributes.TrafficPolicy.Tls.Name
	vService.TrafficPolicy.Tls.SubjectAltNames[0] = input.ServiceAttributes.TrafficPolicy.Tls.SubjectAltNames

	for i, subset := range input.ServiceAttributes.Subsets {
		for _, n := range subset.Name {
			vService.Subsets[i].Name = n
		}
		//vService.Subsets[i].Labels=
		//traffic policy
	}
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
