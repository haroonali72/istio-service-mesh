package core

import (
	//	pb1 "bitbucket.org/cloudplex-devs/kubernetes-services-deployment/core/proto"
	//	"google.golang.org/grpc"
	//	"istio-service-mesh/constants"
	pb "bitbucket.org/cloudplex-devs/microservices-mesh-engine/core/services/proto"
	"context"
	"errors"
	"istio-service-mesh/utils"
	core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

func (s *Server) CreatePersistentVolume(ctx context.Context, req *pb.PersistentVolumeService) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	_, err := getPersistentVolume(req)

	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp, err)
		return serviceResp, err
	}

	//conn, err := grpc.DialContext(ctx, constants.K8sEngineGRPCURL, grpc.WithInsecure())
	//if err != nil {
	//	utils.Error.Println(err)
	//	getErrorResp(serviceResp, err)
	//	return serviceResp, err
	//}
	//defer conn.Close()
	//
	//raw, err := json.Marshal(ksdRequest)
	//if err != nil {
	//	utils.Error.Println(err)
	//	getErrorResp(serviceResp, err)
	//	return serviceResp, err
	//}
	//result, err := pb1.NewServiceClient(conn).CreateService(ctx, &pb1.ServiceRequest{
	//	ProjectId: req.ProjectId,
	//	Service:   raw,
	//	CompanyId: req.CompanyId,
	//	Token:     req.Token,
	//})
	//if err != nil {
	//	utils.Error.Println(err)
	//	getErrorResp(serviceResp, err)
	//	return serviceResp, err
	//}
	//utils.Info.Println(result.Service)
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
	_, err := getPersistentVolume(req)

	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp, err)
		return serviceResp, err
	}

	//conn, err := grpc.DialContext(ctx, constants.K8sEngineGRPCURL, grpc.WithInsecure())
	//if err != nil {
	//	utils.Error.Println(err)
	//	getErrorResp(serviceResp, err)
	//	return serviceResp, err
	//}
	//defer conn.Close()
	//
	//raw, err := json.Marshal(ksdRequest)
	//if err != nil {
	//	utils.Error.Println(err)
	//	getErrorResp(serviceResp, err)
	//	return serviceResp, err
	//}
	//result, err := pb1.NewServiceClient(conn).GetService(ctx, &pb1.ServiceRequest{
	//	ProjectId: req.ProjectId,
	//	Service:   raw,
	//	CompanyId: req.CompanyId,
	//	Token:     req.Token,
	//})
	//if err != nil {
	//	utils.Error.Println(err)
	//	getErrorResp(serviceResp, err)
	//	return serviceResp, err
	//}
	//utils.Info.Println(result.Service)
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
	_, err := getPersistentVolume(req)

	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp, err)
		return serviceResp, err
	}

	//conn, err := grpc.DialContext(ctx, constants.K8sEngineGRPCURL, grpc.WithInsecure())
	//if err != nil {
	//	utils.Error.Println(err)
	//	getErrorResp(serviceResp, err)
	//	return serviceResp, err
	//}
	//defer conn.Close()
	//
	//raw, err := json.Marshal(ksdRequest)
	//if err != nil {
	//	utils.Error.Println(err)
	//	getErrorResp(serviceResp, err)
	//	return serviceResp, err
	//}
	//result, err := pb1.NewServiceClient(conn).DeleteService(ctx, &pb1.ServiceRequest{
	//	ProjectId: req.ProjectId,
	//	Service:   raw,
	//	CompanyId: req.CompanyId,
	//	Token:     req.Token,
	//})
	//if err != nil {
	//	utils.Error.Println(err)
	//	getErrorResp(serviceResp, err)
	//	return serviceResp, err
	//}
	//utils.Info.Println(result.Service)
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
	_, err := getPersistentVolume(req)

	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp, err)
		return serviceResp, err
	}

	//conn, err := grpc.DialContext(ctx, constants.K8sEngineGRPCURL, grpc.WithInsecure())
	//if err != nil {
	//	utils.Error.Println(err)
	//	getErrorResp(serviceResp, err)
	//	return serviceResp, err
	//}
	//defer conn.Close()
	//
	//raw, err := json.Marshal(ksdRequest)
	//if err != nil {
	//	utils.Error.Println(err)
	//	getErrorResp(serviceResp, err)
	//	return serviceResp, err
	//}
	//result, err := pb1.NewServiceClient(conn).PatchService(ctx, &pb1.ServiceRequest{
	//	ProjectId: req.ProjectId,
	//	Service:   raw,
	//	CompanyId: req.CompanyId,
	//	Token:     req.Token,
	//})
	//if err != nil {
	//	utils.Error.Println(err)
	//	getErrorResp(serviceResp, err)
	//	return serviceResp, err
	//}
	//utils.Info.Println(result.Service)
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
	_, err := getPersistentVolume(req)

	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp, err)
		return serviceResp, err
	}

	//conn, err := grpc.DialContext(ctx, constants.K8sEngineGRPCURL, grpc.WithInsecure())
	//if err != nil {
	//	utils.Error.Println(err)
	//	getErrorResp(serviceResp, err)
	//	return serviceResp, err
	//}
	//defer conn.Close()
	//
	//raw, err := json.Marshal(ksdRequest)
	//if err != nil {
	//	utils.Error.Println(err)
	//	getErrorResp(serviceResp, err)
	//	return serviceResp, err
	//}
	//result, err := pb1.NewServiceClient(conn).PutService(ctx, &pb1.ServiceRequest{
	//	ProjectId: req.ProjectId,
	//	Service:   raw,
	//	CompanyId: req.CompanyId,
	//	Token:     req.Token,
	//})
	//if err != nil {
	//	utils.Error.Println(err)
	//	getErrorResp(serviceResp, err)
	//	return serviceResp, err
	//}
	//utils.Info.Println(result.Service)
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
		pv.Labels = make(map[string]string)
		pv.Labels = input.ServiceAttributes.Labels
	}
	if reclaimPoilcy := input.ServiceAttributes.ReclaimPolicy.String(); reclaimPoilcy == pb.ReclaimPolicy_Delete.String() {
		rcp := core.PersistentVolumeReclaimDelete
		pv.Spec.PersistentVolumeReclaimPolicy = rcp
	} else if reclaimPoilcy == pb.ReclaimPolicy_Retain.String() {
		rcp := core.PersistentVolumeReclaimRetain
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
	if input.ServiceAttributes.StorageClassName != "" {
		pv.Spec.StorageClassName = input.ServiceAttributes.StorageClassName
	}
	for _, each := range input.ServiceAttributes.MountOptions {
		pv.Spec.MountOptions = append(pv.Spec.MountOptions, each)
	}
	if input.ServiceAttributes.VolumeMode == string(core.PersistentVolumeFilesystem) {
		pvm := core.PersistentVolumeFilesystem
		pv.Spec.VolumeMode = &pvm
	} else if input.ServiceAttributes.VolumeMode == string(core.PersistentVolumeBlock) {
		pvm := core.PersistentVolumeBlock
		pv.Spec.VolumeMode = &pvm
	}
	if input.ServiceAttributes.NodeAffinity != nil {
		if ns, err := getNodeSelector(input.ServiceAttributes.NodeAffinity.Required); err != nil {
			return nil, err
		} else {
			pv.Spec.NodeAffinity = new(core.VolumeNodeAffinity)
			pv.Spec.NodeAffinity.Required = ns
		}

	}

	quantity, err := resource.ParseQuantity(input.ServiceAttributes.Capacity)
	if err != nil {
		return nil, errors.New("invalid storage capacity ")
	}
	pv.Spec.Capacity = make(map[core.ResourceName]resource.Quantity)
	pv.Spec.Capacity["storage"] = quantity
	if input.ServiceAttributes.PersistentVolumeSource != nil {

		if input.ServiceAttributes.PersistentVolumeSource.GcpPd != nil {
			pv.Spec.GCEPersistentDisk = new(core.GCEPersistentDiskVolumeSource)
			pv.Spec.GCEPersistentDisk.PDName = input.ServiceAttributes.PersistentVolumeSource.GcpPd.PdName
			pv.Spec.GCEPersistentDisk.ReadOnly = input.ServiceAttributes.PersistentVolumeSource.GcpPd.Readonly
			if input.ServiceAttributes.PersistentVolumeSource.GcpPd.FileSystem != "" {
				pv.Spec.GCEPersistentDisk.FSType = input.ServiceAttributes.PersistentVolumeSource.GcpPd.FileSystem
			}
			pv.Spec.GCEPersistentDisk.Partition = int32(input.ServiceAttributes.PersistentVolumeSource.GcpPd.Partation)
		} else if input.ServiceAttributes.PersistentVolumeSource.AwsEbs != nil {
			pv.Spec.AWSElasticBlockStore = new(core.AWSElasticBlockStoreVolumeSource)
			pv.Spec.AWSElasticBlockStore.VolumeID = input.ServiceAttributes.PersistentVolumeSource.AwsEbs.VolumeId
			pv.Spec.AWSElasticBlockStore.ReadOnly = input.ServiceAttributes.PersistentVolumeSource.AwsEbs.Readonly
			if input.ServiceAttributes.PersistentVolumeSource.AwsEbs.FileSystem != "" {
				pv.Spec.AWSElasticBlockStore.FSType = input.ServiceAttributes.PersistentVolumeSource.AwsEbs.FileSystem
			}
			pv.Spec.AWSElasticBlockStore.Partition = int32(input.ServiceAttributes.PersistentVolumeSource.AwsEbs.Partation)
		} else if input.ServiceAttributes.PersistentVolumeSource.AzureDisk != nil {
			pv.Spec.AzureDisk = new(core.AzureDiskVolumeSource)
			pv.Spec.AzureDisk.DiskName = input.ServiceAttributes.PersistentVolumeSource.AzureDisk.DiskName
			pv.Spec.AzureDisk.DataDiskURI = input.ServiceAttributes.PersistentVolumeSource.AzureDisk.DiskURI
			if input.ServiceAttributes.PersistentVolumeSource.AzureDisk.CachingMode.String() == pb.AzureDataDiskCachingMode_ModeNone.String() {
				temp := core.AzureDataDiskCachingNone
				pv.Spec.AzureDisk.CachingMode = &temp
			} else if input.ServiceAttributes.PersistentVolumeSource.AzureDisk.CachingMode.String() == pb.AzureDataDiskCachingMode_ReadOnly.String() {
				temp := core.AzureDataDiskCachingReadOnly
				pv.Spec.AzureDisk.CachingMode = &temp
			} else if input.ServiceAttributes.PersistentVolumeSource.AzureDisk.CachingMode.String() == pb.AzureDataDiskCachingMode_ReadWrite.String() {
				temp := core.AzureDataDiskCachingReadWrite
				pv.Spec.AzureDisk.CachingMode = &temp
			}
			if input.ServiceAttributes.PersistentVolumeSource.AzureDisk.Kind.String() == pb.AzureDataDiskKind_Shared.String() {
				temp := core.AzureSharedBlobDisk
				pv.Spec.AzureDisk.Kind = &temp
			} else if input.ServiceAttributes.PersistentVolumeSource.AzureDisk.Kind.String() == pb.AzureDataDiskKind_Dedicated.String() {
				temp := core.AzureDedicatedBlobDisk
				pv.Spec.AzureDisk.Kind = &temp
			} else if input.ServiceAttributes.PersistentVolumeSource.AzureDisk.Kind.String() == pb.AzureDataDiskKind_Managed.String() {
				temp := core.AzureManagedDisk
				pv.Spec.AzureDisk.Kind = &temp
			}
			pv.Spec.AzureDisk.ReadOnly = &input.ServiceAttributes.PersistentVolumeSource.AzureDisk.ReadOnly
			if input.ServiceAttributes.PersistentVolumeSource.AzureDisk.FileSystem != "" {
				pv.Spec.AzureDisk.FSType = &input.ServiceAttributes.PersistentVolumeSource.AzureDisk.FileSystem
			}

		} else if input.ServiceAttributes.PersistentVolumeSource.AzureFile != nil {
			pv.Spec.AzureFile = new(core.AzureFilePersistentVolumeSource)
			pv.Spec.AzureFile.SecretName = input.ServiceAttributes.PersistentVolumeSource.AzureFile.SecretName
			pv.Spec.AzureFile.ShareName = input.ServiceAttributes.PersistentVolumeSource.AzureFile.ShareName
			pv.Spec.AzureFile.ReadOnly = input.ServiceAttributes.PersistentVolumeSource.AzureFile.ReadOnly
			if input.ServiceAttributes.PersistentVolumeSource.AzureFile.SecretNamespace != "" {
				pv.Spec.AzureFile.SecretNamespace = &input.ServiceAttributes.PersistentVolumeSource.AzureFile.SecretNamespace
			}
		}
	}
	return pv, nil
}
