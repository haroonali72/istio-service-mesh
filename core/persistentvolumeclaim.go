package core

import (
	pb1 "bitbucket.org/cloudplex-devs/kubernetes-services-deployment/core/proto"
	pb "bitbucket.org/cloudplex-devs/microservices-mesh-engine/core/services/proto"
	"context"
	"encoding/json"
	"errors"
	"google.golang.org/grpc"
	"istio-service-mesh/constants"
	"istio-service-mesh/utils"
	core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

func (s *Server) CreatePersistentVolumeClaim(ctx context.Context, req *pb.PersistentVolumeClaimService) (*pb.ServiceResponse, error) {
	utils.Info.Println(ctx)
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getPersistentVolumeClaim(req)

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

	/*converToResp(serviceResp,req.ProjectId,statusCode,resp)

	url := fmt.Sprintf("%s%s",constants.KubernetesEngineURL,constants.KUBERNETES_SERVICES_DEPLOYMENT)
	statusCode, resp, err := utils.Post(url,ksdRequest,getHeaders(ctx,req.ProjectId))

	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp,err)
		return serviceResp,err
	}
	converToResp(serviceResp,req.ProjectId,statusCode,resp)
	return serviceResp,nil*/
}
func (s *Server) GetPersistentVolumeClaim(ctx context.Context, req *pb.PersistentVolumeClaimService) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getPersistentVolumeClaim(req)

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
func (s *Server) DeletePersistentVolumeClaim(ctx context.Context, req *pb.PersistentVolumeClaimService) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getPersistentVolumeClaim(req)

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
func (s *Server) PatchPersistentVolumeClaim(ctx context.Context, req *pb.PersistentVolumeClaimService) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getPersistentVolumeClaim(req)

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
func (s *Server) PutPersistentVolumeClaim(ctx context.Context, req *pb.PersistentVolumeClaimService) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getPersistentVolumeClaim(req)

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

func getPersistentVolumeClaim(input *pb.PersistentVolumeClaimService) (*core.PersistentVolumeClaim, error) {
	var pvc = new(core.PersistentVolumeClaim)
	pvc.Name = input.Name
	pvc.TypeMeta.Kind = "PersistentVolumeClaim"
	pvc.TypeMeta.APIVersion = "v1"
	if input.Namespace == "" {
		pvc.Namespace = "default"
	} else {
		pvc.Namespace = input.Namespace
	}
	if input.ServiceAttributes.StorageClassName != "" {
		pvc.Spec.StorageClassName = &input.ServiceAttributes.StorageClassName
	} else {
		pvc.Spec.VolumeName = input.ServiceAttributes.VolumeName
		if ls, err := getLabelSelector(input.ServiceAttributes.LabelSelector); err == nil {
			if ls != nil {
				pvc.Spec.Selector = ls
			}

		} else {
			return nil, err
		}

	}
	if input.ServiceAttributes.DataSource != nil {
		pvc.Spec.DataSource = new(core.TypedLocalObjectReference)
		pvc.Spec.DataSource.Kind = input.ServiceAttributes.DataSource.Kind
		pvc.Spec.DataSource.Name = input.ServiceAttributes.DataSource.Name
		if input.ServiceAttributes.DataSource.ApiGroup != "" {
			pvc.Spec.DataSource.APIGroup = &input.ServiceAttributes.DataSource.ApiGroup

		}
	}
	//pvm := core.PersistentVolumeFilesystem
	//	pvc.Spec.VolumeMode = &pvm
	if input.ServiceAttributes.VolumeMode == string(core.PersistentVolumeFilesystem) {
		pvm := core.PersistentVolumeFilesystem
		pvc.Spec.VolumeMode = &pvm
	} else if input.ServiceAttributes.VolumeMode == string(core.PersistentVolumeBlock) {
		pvm := core.PersistentVolumeBlock
		pvc.Spec.VolumeMode = &pvm
	}
	for _, each := range input.ServiceAttributes.AccessMode {
		if each == pb.AccessMode_ReadOnlyMany {
			pvc.Spec.AccessModes = append(pvc.Spec.AccessModes, core.ReadOnlyMany)
		} else if each == pb.AccessMode_ReadWriteMany {
			pvc.Spec.AccessModes = append(pvc.Spec.AccessModes, core.ReadWriteMany)
		} else if each == pb.AccessMode_ReadWriteOnce {
			pvc.Spec.AccessModes = append(pvc.Spec.AccessModes, core.ReadWriteOnce)
		}
	}
	if input.ServiceAttributes.RequestQuantity != "" {
		quantity, err := resource.ParseQuantity(input.ServiceAttributes.RequestQuantity)
		if err != nil {
			return nil, errors.New("invalid storage capacity ")
		}
		pvc.Spec.Resources.Requests = make(map[core.ResourceName]resource.Quantity)

		pvc.Spec.Resources.Requests["storage"] = quantity
	}
	if input.ServiceAttributes.LimitQuantity != "" {
		quantity, err := resource.ParseQuantity(input.ServiceAttributes.LimitQuantity)
		if err != nil {
			return nil, errors.New("invalid storage capacity ")
		}
		pvc.Spec.Resources.Limits = make(map[core.ResourceName]resource.Quantity)

		pvc.Spec.Resources.Limits["storage"] = quantity
	}

	return pvc, nil
}
