package core

import (
	"context"
	"encoding/json"
	"errors"
	"google.golang.org/grpc"
	"istio-service-mesh/constants"
	pb "istio-service-mesh/core/proto"
	"istio-service-mesh/utils"
	core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

func (s *Server) CreatePersistentVolume(ctx context.Context, req *pb.PersistentVolumeService) (*pb.ServiceResponse, error) {
	utils.Info.Println(ctx)
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getPersistentVolume(req)

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
func (s *Server) GetPersistentVolume(ctx context.Context, req *pb.PersistentVolumeService) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getPersistentVolume(req)

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
func (s *Server) DeletePersistentVolume(ctx context.Context, req *pb.PersistentVolumeService) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getPersistentVolume(req)

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
func (s *Server) PatchPersistentVolume(ctx context.Context, req *pb.PersistentVolumeService) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getPersistentVolume(req)

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
func (s *Server) PutPersistentVolume(ctx context.Context, req *pb.PersistentVolumeService) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getPersistentVolume(req)

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

func getPersistentVolume(input *pb.PersistentVolumeService) (*core.PersistentVolume, error) {
	var pv = new(core.PersistentVolume)
	pv.Name = input.Name
	pv.TypeMeta.Kind = "PersistentVolume"
	pv.TypeMeta.APIVersion = "v1"
	if len(input.ServiceAttributes.Labels) > 0 {
		pv.Labels = input.ServiceAttributes.Labels
	}
	if reclaimPoilcy := input.ServiceAttributes.ReclaimPolicy.String(); reclaimPoilcy == pb.ReclaimPolicy_Delete.String() {
		rcp := core.PersistentVolumeReclaimDelete
		pv.Spec.PersistentVolumeReclaimPolicy = rcp
	}
	for _, each := range input.ServiceAttributes.AccessMode {
		if each == pb.AccessMode_ReadOnlyMany {
			pv.Spec.AccessModes = append(pv.Spec.AccessModes, core.ReadOnlyMany)
		} else if each == pb.AccessMode_ReadWriteMany {
			pv.Spec.AccessModes = append(pv.Spec.AccessModes, core.ReadWriteMany)
		} else if each == pb.AccessMode_ReadWriteOnce {
			pv.Spec.AccessModes = append(pv.Spec.AccessModes, core.ReadWriteOnce)
		}
	}
	quantity, err := resource.ParseQuantity(input.ServiceAttributes.Capcity)
	if err != nil {
		return nil, errors.New("invalid storage capacity ")
	}
	pv.Spec.Capacity["storage"] = quantity
	if input.ServiceAttributes.PersistentVolumeSource.GcpPd.PdName != "" {
		pv.Spec.GCEPersistentDisk.PDName = input.ServiceAttributes.PersistentVolumeSource.GcpPd.PdName
		pv.Spec.GCEPersistentDisk.ReadOnly = input.ServiceAttributes.PersistentVolumeSource.GcpPd.ReadOnly
	} else if input.ServiceAttributes.PersistentVolumeSource.AwsEbs.VolumeId != "" {
		pv.Spec.AWSElasticBlockStore.VolumeID = input.ServiceAttributes.PersistentVolumeSource.AwsEbs.VolumeId
		pv.Spec.AWSElasticBlockStore.ReadOnly = input.ServiceAttributes.PersistentVolumeSource.AwsEbs.ReadOnly
	} else if input.ServiceAttributes.PersistentVolumeSource.AzureDisk.DiskName != "" {
		pv.Spec.AzureDisk.DiskName = input.ServiceAttributes.PersistentVolumeSource.AzureDisk.DiskName
		pv.Spec.AzureDisk.DataDiskURI = input.ServiceAttributes.PersistentVolumeSource.AzureDisk.DiskURI
		if input.ServiceAttributes.PersistentVolumeSource.AzureDisk.CachingMode == "None" {
			temp := core.AzureDataDiskCachingNone
			pv.Spec.AzureDisk.CachingMode = &temp
		} else if input.ServiceAttributes.PersistentVolumeSource.AzureDisk.CachingMode == "ReadOnly" {
			temp := core.AzureDataDiskCachingReadOnly
			pv.Spec.AzureDisk.CachingMode = &temp
		} else if input.ServiceAttributes.PersistentVolumeSource.AzureDisk.CachingMode == "ReadWrite" {
			temp := core.AzureDataDiskCachingReadWrite
			pv.Spec.AzureDisk.CachingMode = &temp
		}
		pv.Spec.AzureDisk.ReadOnly = &input.ServiceAttributes.PersistentVolumeSource.AzureDisk.ReadOnly
	} else if input.ServiceAttributes.PersistentVolumeSource.AzureFile.SecretName != "" {
		pv.Spec.AzureFile.SecretName = input.ServiceAttributes.PersistentVolumeSource.AzureFile.SecretName
		pv.Spec.AzureFile.ShareName = input.ServiceAttributes.PersistentVolumeSource.AzureFile.ShareName
		pv.Spec.AzureFile.ReadOnly = input.ServiceAttributes.PersistentVolumeSource.AzureFile.ReadOnly
	}
	return pv, nil
}
