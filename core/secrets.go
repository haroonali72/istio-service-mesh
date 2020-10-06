package core

import (
	"bitbucket.org/cloudplex-devs/istio-service-mesh/constants"
	"bitbucket.org/cloudplex-devs/istio-service-mesh/utils"
	pb1 "bitbucket.org/cloudplex-devs/kubernetes-services-deployment/core/proto"
	pb "bitbucket.org/cloudplex-devs/microservices-mesh-engine/core/services/proto"
	"bitbucket.org/cloudplex-devs/microservices-mesh-engine/types/services"
	"context"
	"encoding/json"
	"google.golang.org/grpc"
	kb "k8s.io/api/core/v1"
	"strings"
)

func (s *Server) CreateSecretService(ctx context.Context, req *pb.SecretService) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getRequestSecretObject(req)
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
		InfraId:       req.InfraId,
		ApplicationId: req.ApplicationId,

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
func (s *Server) GetSecretService(ctx context.Context, req *pb.SecretService) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getRequestSecretObject(req)
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
		InfraId:       req.InfraId,
		ApplicationId: req.ApplicationId,

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
func (s *Server) DeleteSecretService(ctx context.Context, req *pb.SecretService) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getRequestSecretObject(req)
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
		InfraId:       req.InfraId,
		ApplicationId: req.ApplicationId,

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
func (s *Server) PatchSecretService(ctx context.Context, req *pb.SecretService) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getRequestSecretObject(req)
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
		InfraId:       req.InfraId,
		ApplicationId: req.ApplicationId,

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
func (s *Server) PutSecretService(ctx context.Context, req *pb.SecretService) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getRequestSecretObject(req)
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
		InfraId:       req.InfraId,
		ApplicationId: req.ApplicationId,

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

func getSecret(input *pb.SecretService) (*kb.Secret, error) {

	var kube = new(kb.Secret)
	kube.Kind = constants.Secret.String() //"Secret"
	kube.APIVersion = "v1"
	kube.Name = input.Name
	kube.Namespace = input.Namespace
	labels := make(map[string]string)
	labels["app"] = strings.ToLower(input.Name)
	labels["version"] = strings.ToLower(input.Version)
	kube.Labels = labels
	sa, err := getFromVault(input.ApplicationId, input.ServiceId, input.Token)
	if err != nil {
		utils.Error.Println(err)
		return nil, err
	}
	if sa != nil {

		switch sa.SecretType {
		case "Opaque":
			kube.Type = kb.SecretTypeOpaque
		case "ServiceAccountToken":
			kube.Type = kb.SecretType(kb.ServiceAccountTokenKey)
		case "ServiceAccountNameKey":
			kube.Type = kb.SecretType(kb.ServiceAccountNameKey)
		case "ServiceAccountUIDKey":
			kube.Type = kb.SecretType(kb.ServiceAccountUIDKey)
		case "ServiceAccountTokenKey":
			kube.Type = kb.SecretType(kb.ServiceAccountTokenKey)
		case "ServiceAccountKubeconfigKey":
			kube.Type = kb.SecretType(kb.ServiceAccountKubeconfigKey)
		case "ServiceAccountRootCAKey":
			kube.Type = kb.SecretType(kb.ServiceAccountRootCAKey)
		case "SecretTypeDockercfg":
			kube.Type = kb.SecretType(kb.SecretTypeDockercfg)
		case "DockerConfigKey":
			kube.Type = kb.SecretType(kb.DockerConfigKey)
		case "Tls":
			kube.Type = kb.SecretType(kb.SecretTypeTLS)
		default:
			kube.Type = kb.SecretType(sa.SecretType)
		}

		if len(sa.Data) > 0 {
			map2 := make(map[string][]byte)
			for key, value := range sa.Data {
				s := []byte(value)
				map2[key] = s
			}

			kube.Data = make(map[string][]byte)
			kube.Data = map2
		}

		if len(sa.StringData) > 0 {
			kube.StringData = make(map[string]string)
			for key, value := range sa.StringData {
				kube.StringData[key] = value
			}
		}
	}

	return kube, nil
}
func getRequestSecretObject(req *pb.SecretService) (*kb.Secret, error) {
	scrReq, err := getSecret(req)
	if err != nil {
		utils.Error.Println(err)

		return nil, err
	}
	return scrReq, nil
}
func getFromVault(applicationId, serviceID, token string) (*services.SecretServiceAttribute, error) {
	url := constants.VaultURL + constants.VAULT_GETSECRET
	url = strings.Replace(url, "{applicationId}", applicationId, -1)
	url = strings.Replace(url, "{serviceId}", serviceID, -1)
	_, resp, err := utils.Get(url, nil, map[string]string{"X-Auth-Token": token})
	if err == nil {
		attrib := new(services.SecretServiceAttribute)
		err = json.Unmarshal(resp, attrib)
		if err != nil {
			return nil, err
		}
		return attrib, nil
	} else {
		utils.Error.Println(err)
	}

	return nil, nil
}
