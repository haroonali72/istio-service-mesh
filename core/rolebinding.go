package core

import (
	"context"
	"encoding/json"
	"errors"
	"google.golang.org/grpc"
	"istio-service-mesh/constants"
	pb "istio-service-mesh/core/proto"
	"istio-service-mesh/utils"
	v1 "k8s.io/api/rbac/v1"
	"strings"
)

func (s *Server) CreateRoleBindingService(ctx context.Context, req *pb.RoleBindingService) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	roleBinbRequest, err := getRequestRoleBindObject(req)
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

	raw, err := json.Marshal(roleBinbRequest)
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
func (s *Server) GetRoleBindingService(ctx context.Context, req *pb.RoleBindingService) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	roleBinbRequest, err := getRequestRoleBindObject(req)
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

	raw, err := json.Marshal(roleBinbRequest)
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
func (s *Server) DeleteRoleBindingService(ctx context.Context, req *pb.RoleBindingService) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	roleBinbRequest, err := getRequestRoleBindObject(req)
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

	raw, err := json.Marshal(roleBinbRequest)
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
func (s *Server) PatchRoleBindingService(ctx context.Context, req *pb.RoleBindingService) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	roleBinbRequest, err := getRequestRoleBindObject(req)
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

	raw, err := json.Marshal(roleBinbRequest)
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
func (s *Server) PutRoleBindingService(ctx context.Context, req *pb.RoleBindingService) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	roleBinbRequest, err := getRequestRoleBindObject(req)
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

	raw, err := json.Marshal(roleBinbRequest)
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

func getRoleBinding(input *pb.RoleBindingService) (*v1.RoleBinding, error) {
	//var roleBind = new(v1.RoleBinding)
	//labels := make(map[string]string)
	//labels["app"] = strings.ToLower(input.Name)
	//labels["version"] = strings.ToLower(input.Version)
	//roleBind.Kind = "RoleBinding"
	//roleBind.APIVersion = "rbac.authorization.k8s.io/v1"
	//roleBind.Name = input.Name
	//roleBind.Namespace = input.Namespace
	//roleBind.Labels = labels
	//
	//for _, subject := range input.ServiceAttributes.Subjects {
	//	var sub v1.Subject
	//	sub.Name = subject.Name
	//	sub.Kind = subject.Kind
	//	sub.APIGroup = subject.ApiGroup
	//	//sub.Namespace=subject.Namespace
	//	roleBind.Subjects = append(roleBind.Subjects, sub)
	//}
	//
	//roleBind.RoleRef.Kind = input.ServiceAttributes.Reference.Kind
	//roleBind.RoleRef.Name = input.ServiceAttributes.Reference.Name
	//roleBind.RoleRef.APIGroup = input.ServiceAttributes.Reference.ApiGroup
	//
	//return roleBind, nil

	var roleBind = new(v1.RoleBinding)
	if input.Name != "" {
		roleBind.Name = input.Name
	} else {
		return nil, errors.New("can not find name in cluster role binding")
	}

	labels := make(map[string]string)
	labels["app"] = strings.ToLower(input.Name)
	if input.Version != "" {
		labels["version"] = strings.ToLower(input.Version)
	}
	roleBind.Kind = "RoleBinding"
	roleBind.APIVersion = "rbac.authorization.k8s.io/v1"

	roleBind.Labels = labels
	for _, subject := range input.ServiceAttributes.Subjects {
		var reqsub v1.Subject
		if subject.Name != "" {
			reqsub.Name = subject.Name
		} else {
			return nil, errors.New("can not find name for subject")
		}
		if subject.Kind == "User" || subject.Kind == "Group" {
			reqsub.Kind = subject.Kind
			reqsub.APIGroup = "rbac.authorization.k8s.io"
		} else if subject.Kind == "ServiceAccount" {
			reqsub.Kind = subject.Kind
			if subject.Namespace != "" {
				reqsub.Namespace = subject.Namespace
			} else {
				return nil, errors.New("can not find name space for service account" + reqsub.Name)
			}

		} else {
			return nil, errors.New("can not find name space for service account" + reqsub.Name)
		}
		roleBind.Subjects = append(roleBind.Subjects, reqsub)
	}
	if input.ServiceAttributes.Reference != nil {
		if input.ServiceAttributes.Reference.Kind == "ClusterRole" || input.ServiceAttributes.Reference.Kind == "Role" {
			roleBind.RoleRef.Kind = input.ServiceAttributes.Reference.Kind
			roleBind.RoleRef.APIGroup = "rbac.authorization.k8s.io"
			if input.ServiceAttributes.Reference.Name != "" {
				roleBind.RoleRef.Name = input.ServiceAttributes.Reference.Name
			} else {
				return nil, errors.New("can not find Name in cluster role binding ref " + input.Name)
			}
		} else {
			return nil, errors.New("invalid kind  role binding ref " + input.Name)
		}
	} else {
		return nil, errors.New("can not find role ref in role binding" + input.Name)
	}

	return roleBind, nil

}
func getRequestRoleBindObject(req *pb.RoleBindingService) (*v1.RoleBinding, error) {
	roleReq, err := getRoleBinding(req)
	if err != nil {
		utils.Error.Println(err)

		return nil, err
	}
	return roleReq, nil
}
