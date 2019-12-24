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
	"k8s.io/api/storage/v1"
)

func (s *Server) CreateStorageClass(ctx context.Context, req *pb.StorageClassService) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getStorageClass(req)

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
func (s *Server) GetStorageClass(ctx context.Context, req *pb.StorageClassService) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getStorageClass(req)

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
func (s *Server) DeleteStorageClass(ctx context.Context, req *pb.StorageClassService) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getStorageClass(req)

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
func (s *Server) PatchStorageClass(ctx context.Context, req *pb.StorageClassService) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getStorageClass(req)

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
func (s *Server) PutStorageClass(ctx context.Context, req *pb.StorageClassService) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getStorageClass(req)

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

func getStorageClass(input *pb.StorageClassService) (*v1.StorageClass, error) {
	var sc = new(v1.StorageClass)
	sc.Name = input.Name
	sc.TypeMeta.Kind = "StorageClass"
	sc.TypeMeta.APIVersion = "storage.k8s.io/v1"
	if volBindingMod := input.ServiceAttributes.VolumeBindingMode.String(); volBindingMod == pb.VolumeBindingMode_WaitForFirstCustomer.String() {
		vbm := v1.VolumeBindingWaitForFirstConsumer
		sc.VolumeBindingMode = &vbm
	}
	if reclaimPoilcy := input.ServiceAttributes.ReclaimPolicy.String(); reclaimPoilcy == pb.ReclaimPolicy_Retain.String() {
		rcp := core.PersistentVolumeReclaimRetain
		sc.ReclaimPolicy = &rcp
	}
	sc.Parameters = make(map[string]string)
	// SC  AWSEBS
	if len(input.ServiceAttributes.ScParameters.AwsebsscParm) > 0 {
		sc.Provisioner = "kubernetes.io/aws-ebs"
		ebsType := input.ServiceAttributes.ScParameters.AwsebsscParm["type"]
		if ebsType == "io1" {

			io1IopsperGb := input.ServiceAttributes.ScParameters.AwsebsscParm["iopsPerGB"]
			if io1IopsperGb == "" {
				return nil, errors.New("can not find io1 IopsperGb in sc parameters")
			}
			sc.Parameters["iopsPerGB"] = io1IopsperGb
		}
		sc.Parameters["type"] = ebsType
		sc.Parameters["encrypted"] = input.ServiceAttributes.ScParameters.AwsebsscParm["encrypted"]
		if kmsKeyId := input.ServiceAttributes.ScParameters.AwsebsscParm["kmsKeyId"]; kmsKeyId != "" {
			sc.Parameters["kmsKeyId"] = input.ServiceAttributes.ScParameters.AwsebsscParm["kmsKeyId"]
		}
	}

	// SC  GCPPD
	if len(input.ServiceAttributes.ScParameters.GcppdscParm) > 0 {
		sc.Provisioner = "kubernetes.io/gce-pd"

		if gcppdType := input.ServiceAttributes.ScParameters.GcppdscParm["type"]; gcppdType == "pd-standard" {
			sc.Parameters["type"] = gcppdType
		}
		if regionalpd := input.ServiceAttributes.ScParameters.GcppdscParm["replication-type"]; regionalpd == "regional-pd" {
			sc.Parameters["replication-type"] = regionalpd
		}
	}
	// SC  Azuredisk
	if len(input.ServiceAttributes.ScParameters.AzurdiskscParm) > 0 {
		sc.Provisioner = "kubernetes.io/azure-disk"
		if skuName := input.ServiceAttributes.ScParameters.AzurdiskscParm["skuName"]; skuName != "" {
			sc.Parameters["skuName"] = skuName
		}
		if location := input.ServiceAttributes.ScParameters.AzurdiskscParm["location"]; location != "" {
			sc.Parameters["location"] = location
		}
		if sa := input.ServiceAttributes.ScParameters.AzurdiskscParm["storageAccount"]; sa != "" {
			sc.Parameters["storageAccount"] = sa
		}

	}
	// SC  AzureFile
	if len(input.ServiceAttributes.ScParameters.AzurfilescParm) > 0 {
		sc.Provisioner = "kubernetes.io/azure-file"
		if skuName := input.ServiceAttributes.ScParameters.AzurfilescParm["skuName"]; skuName != "" {
			sc.Parameters["skuName"] = skuName
		}
		if location := input.ServiceAttributes.ScParameters.AzurfilescParm["location"]; location != "" {
			sc.Parameters["location"] = location
		}
		if sa := input.ServiceAttributes.ScParameters.AzurfilescParm["storageAccount"]; sa != "" {
			sc.Parameters["storageAccount"] = sa
		}
		if scNs := input.ServiceAttributes.ScParameters.AzurfilescParm["secretNamespace"]; scNs != "" {
			sc.Parameters["secretNamespace"] = scNs
		}
		if sa := input.ServiceAttributes.ScParameters.AzurfilescParm["secretName"]; sa != "" {
			sc.Parameters["secretName"] = sa
		}

	}
	return sc, nil
}
