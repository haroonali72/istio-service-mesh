package core

import (
	"bitbucket.org/cloudplex-devs/istio-service-mesh/constants"
	"bitbucket.org/cloudplex-devs/istio-service-mesh/utils"
	pb1 "bitbucket.org/cloudplex-devs/kubernetes-services-deployment/core/proto"
	pb "bitbucket.org/cloudplex-devs/microservices-mesh-engine/core/services/proto"
	"context"
	"encoding/json"
	"errors"
	cmacme "github.com/jetstack/cert-manager/pkg/apis/acme/v1alpha3"
	"github.com/jetstack/cert-manager/pkg/apis/certmanager/v1alpha3"
	"google.golang.org/grpc"
	"strings"
)

func (s *Server) CreateClusterIssuer(ctx context.Context, req *pb.ClusterIssuerService) (*pb.ServiceResponse, error) {
	utils.Info.Println(ctx)
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}

	clstrIssuerReq, err := getClusterIssuerRequestObject(req)

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

	raw, err := json.Marshal(clstrIssuerReq)
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
func (s *Server) GetClusterIssuer(ctx context.Context, req *pb.ClusterIssuerService) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	clstrIssuerReq, err := getClusterIssuerRequestObject(req)

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

	raw, err := json.Marshal(clstrIssuerReq)
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
func (s *Server) DeleteClusterIssuer(ctx context.Context, req *pb.ClusterIssuerService) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	clstrIssuerReq, err := getClusterIssuerRequestObject(req)

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

	raw, err := json.Marshal(clstrIssuerReq)
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
func (s *Server) PatchClusterIssuer(ctx context.Context, req *pb.ClusterIssuerService) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	clstrIssuerReq, err := getClusterIssuerRequestObject(req)

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

	raw, err := json.Marshal(clstrIssuerReq)
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
func (s *Server) PutClusterIssuer(ctx context.Context, req *pb.ClusterIssuerService) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	clstrIssuerReq, err := getClusterIssuerRequestObject(req)

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

	raw, err := json.Marshal(clstrIssuerReq)
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

func getClusterIssuer(input *pb.ClusterIssuerService) (*v1alpha3.ClusterIssuer, error) {
	var clstrIssueSvc = new(v1alpha3.ClusterIssuer)
	labels := make(map[string]string)
	labels["app"] = strings.ToLower(input.Name)
	labels["version"] = strings.ToLower(input.Version)
	clstrIssueSvc.Labels = labels
	clstrIssueSvc.Kind = constants.ClusterIssuer.String()
	clstrIssueSvc.APIVersion = "cert-manager.io/v1alpha2"
	if input.Name == "" {
		return nil, errors.New("service name must not be empty")
	} else {
		clstrIssueSvc.Name = input.Name
	}

	clstrIssueSvc.Spec.ACME = new(cmacme.ACMEIssuer)
	clstrIssueSvc.Spec.ACME.Server = "https://acme-v02.api.letsencrypt.org/directory"
	if input.ServiceAttributes.Email == "" {
		return nil, errors.New("email must not be empty for cluster issuer service")
	} else {
		clstrIssueSvc.Spec.ACME.Email = input.ServiceAttributes.Email
	}

	clstrIssueSvc.Spec.ACME.PrivateKey.Name = input.Name + "-secret"

	ingressClass := new(string)
	*ingressClass = "istio"
	clstrIssueSvc.Spec.ACME.Solvers = []cmacme.ACMEChallengeSolver{
		{
			HTTP01: &cmacme.ACMEChallengeSolverHTTP01{
				Ingress: &cmacme.ACMEChallengeSolverHTTP01Ingress{
					Class: ingressClass,
				},
			},
		},
	}

	return clstrIssueSvc, nil
}

func getClusterIssuerRequestObject(req *pb.ClusterIssuerService) (*v1alpha3.ClusterIssuer, error) {
	clstrIssuerReq, err := getClusterIssuer(req)
	if err != nil {
		utils.Error.Println(err)

		return nil, err
	}
	return clstrIssuerReq, nil
}
