package core

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"istio-service-mesh/constants"
	pb "istio-service-mesh/core/proto"
	"istio-service-mesh/utils"
	kb "k8s.io/api/core/v1"
	"strings"
)

func (s *Server) CreateConfigMapService(ctx context.Context, req *pb.ConfigMapService) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	abc
	ksdRequest, err := getRequestConfigMapObject(req)
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
func (s *Server) GetConfigMapService(ctx context.Context, req *pb.ConfigMapService) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getRequestConfigMapObject(req)
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
func (s *Server) DeleteConfigMapService(ctx context.Context, req *pb.ConfigMapService) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getRequestConfigMapObject(req)
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
func (s *Server) PatchConfigMapService(ctx context.Context, req *pb.ConfigMapService) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getRequestConfigMapObject(req)
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
func (s *Server) PutConfigMapService(ctx context.Context, req *pb.ConfigMapService) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getRequestConfigMapObject(req)
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

func getConfigMapService(input *pb.ConfigMapService) (*kb.ConfigMap, error) {

	var config = new(kb.ConfigMap)
	config.Kind = "ConfigMap"
	config.APIVersion = "v1"
	config.Name = input.Name
	config.Namespace = input.Namespace
	config.ResourceVersion = input.Version
	labels := make(map[string]string)
	labels["app"] = strings.ToLower(input.Name)
	labels["version"] = strings.ToLower(input.Version)
	config.Labels = labels

	config.Data = input.ServiceAttribute.Data

	map2 := make(map[string][]byte)
	for key, value := range input.ServiceAttribute.BinaryData {
		s := []byte(value)
		fmt.Println(s)
		map2[key] = s
	}
	config.BinaryData = map2

	return config, nil
}
func getRequestConfigMapObject(req *pb.ConfigMapService) (*kb.ConfigMap, error) {
	configReq, err := getConfigMapService(req)
	if err != nil {
		utils.Error.Println(err)

		return nil, err
	}
	return configReq, nil
}
