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

func getStorageClass(input *pb.StorageClassService) (*v1.StorageClass, error) {
	var sc = new(v1.StorageClass)
	sc.Name = input.Name
	sc.TypeMeta.Kind = "StorageClass"
	sc.TypeMeta.APIVersion = "storage.k8s.io/v1"
	if input.ServiceAttributes == nil {
		return nil, errors.New("not found")
	}
	if input.ServiceAttributes.AllowVolumeExpansion == "true" {
		vE := true
		sc.AllowVolumeExpansion = &vE
	} else if input.ServiceAttributes.AllowVolumeExpansion == "false" {
		vE := false
		sc.AllowVolumeExpansion = &vE
	}

	if volBindingMod := input.ServiceAttributes.VolumeBindingMode.String(); volBindingMod == pb.VolumeBindingMode_WaitForFirstConsumer.String() {
		vbm := v1.VolumeBindingWaitForFirstConsumer
		sc.VolumeBindingMode = &vbm
	} else if volBindingMod == pb.VolumeBindingMode_Immediate.String() {
		vbm := v1.VolumeBindingImmediate
		sc.VolumeBindingMode = &vbm
	}
	if reclaimPoilcy := input.ServiceAttributes.ReclaimPolicy.String(); reclaimPoilcy == pb.ReclaimPolicy_Retain.String() {
		rcp := core.PersistentVolumeReclaimRetain
		sc.ReclaimPolicy = &rcp
	} else if reclaimPoilcy == pb.ReclaimPolicy_Retain.String() {
		rcp := core.PersistentVolumeReclaimRetain
		sc.ReclaimPolicy = &rcp
	}

	for _, each := range input.ServiceAttributes.AllowedTopologies {
		aT := core.TopologySelectorTerm{}
		for _, each2 := range each.MatchLabelExpressions {
			tr := core.TopologySelectorLabelRequirement{}
			tr.Key = each2.Key
			for _, value := range each2.Values {
				tr.Values = append(tr.Values, value)
			}
			aT.MatchLabelExpressions = append(aT.MatchLabelExpressions, tr)

		}
		sc.AllowedTopologies = append(sc.AllowedTopologies, aT)
	}
	sc.Parameters = make(map[string]string)
	// SC  AWSEBS
	if len(input.ServiceAttributes.ScParameters.AwsEbsScParm) > 0 {
		sc.Provisioner = "kubernetes.io/aws-ebs"
		ebsType := input.ServiceAttributes.ScParameters.AwsEbsScParm["type"]
		if ebsType == "io1" {

			io1IopsperGb := input.ServiceAttributes.ScParameters.AwsEbsScParm["iopsPerGB"]
			if io1IopsperGb == "" {
				return nil, errors.New("can not find io1 IopsperGb in sc parameters")
			}
			sc.Parameters["iopsPerGB"] = io1IopsperGb
		}
		sc.Parameters["type"] = ebsType
		if input.ServiceAttributes.ScParameters.AwsEbsScParm["encrypted"] != "" {
			sc.Parameters["encrypted"] = input.ServiceAttributes.ScParameters.AwsEbsScParm["encrypted"]
		}
		if kmsKeyId := input.ServiceAttributes.ScParameters.AwsEbsScParm["kmsKeyId"]; kmsKeyId != "" {
			sc.Parameters["kmsKeyId"] = input.ServiceAttributes.ScParameters.AwsEbsScParm["kmsKeyId"]
		}
		if fsType := input.ServiceAttributes.ScParameters.AwsEbsScParm["fsType"]; fsType != "" {
			sc.Parameters["fsType"] = fsType
		}
		for _, each := range input.ServiceAttributes.MountOptions {
			sc.MountOptions = append(sc.MountOptions, each)
		}
		if input.ServiceAttributes.ScParameters.AwsEbsScParm["zone"] != "" {
			sc.Parameters["zone"] = input.ServiceAttributes.ScParameters.AwsEbsScParm["zone"]
		} else if input.ServiceAttributes.ScParameters.AwsEbsScParm["zones"] != "" {
			sc.Parameters["zones"] = input.ServiceAttributes.ScParameters.AwsEbsScParm["zones"]
		}
	}

	// SC  GCPPD
	if len(input.ServiceAttributes.ScParameters.GcpPdScParm) > 0 {
		sc.Provisioner = "kubernetes.io/gce-pd"

		if gcppdType := input.ServiceAttributes.ScParameters.GcpPdScParm["type"]; gcppdType == "pd-standard" {
			sc.Parameters["type"] = gcppdType
		}
		if regionalpd := input.ServiceAttributes.ScParameters.GcpPdScParm["replication-type"]; regionalpd == "regional-pd" {
			sc.Parameters["replication-type"] = regionalpd
		}
		for _, each := range input.ServiceAttributes.MountOptions {
			sc.MountOptions = append(sc.MountOptions, each)
		}
		if input.ServiceAttributes.ScParameters.GcpPdScParm["zone"] != "" {
			sc.Parameters["zone"] = input.ServiceAttributes.ScParameters.GcpPdScParm["zone"]
		} else if input.ServiceAttributes.ScParameters.GcpPdScParm["zones"] != "" {
			sc.Parameters["zones"] = input.ServiceAttributes.ScParameters.GcpPdScParm["zones"]
		}
	}
	// SC  Azuredisk
	if len(input.ServiceAttributes.ScParameters.AzureDiskScParm) > 0 {
		sc.Provisioner = "kubernetes.io/azure-disk"
		if skuName := input.ServiceAttributes.ScParameters.AzureDiskScParm["skuName"]; skuName != "" {
			sc.Parameters["skuName"] = skuName
		}
		if location := input.ServiceAttributes.ScParameters.AzureDiskScParm["location"]; location != "" {
			sc.Parameters["location"] = location
		}
		if sa := input.ServiceAttributes.ScParameters.AzureDiskScParm["storageAccount"]; sa != "" {
			sc.Parameters["storageAccount"] = sa
		}
		for _, each := range input.ServiceAttributes.MountOptions {
			sc.MountOptions = append(sc.MountOptions, each)
		}
	}
	// SC  AzureFile
	if len(input.ServiceAttributes.ScParameters.AzureFileScParm) > 0 {
		sc.Provisioner = "kubernetes.io/azure-file"
		if skuName := input.ServiceAttributes.ScParameters.AzureFileScParm["skuName"]; skuName != "" {
			sc.Parameters["skuName"] = skuName
		}
		if location := input.ServiceAttributes.ScParameters.AzureFileScParm["location"]; location != "" {
			sc.Parameters["location"] = location
		}
		if sa := input.ServiceAttributes.ScParameters.AzureFileScParm["storageAccount"]; sa != "" {
			sc.Parameters["storageAccount"] = sa
		}
		if scNs := input.ServiceAttributes.ScParameters.AzureFileScParm["secretNamespace"]; scNs != "" {
			sc.Parameters["secretNamespace"] = scNs
		}
		if sa := input.ServiceAttributes.ScParameters.AzureFileScParm["secretName"]; sa != "" {
			sc.Parameters["secretName"] = sa
		}
		if re := input.ServiceAttributes.ScParameters.AzureFileScParm["readOnly"]; re != "" {
			sc.Parameters["readOnly"] = re
		}
		for _, each := range input.ServiceAttributes.MountOptions {
			sc.MountOptions = append(sc.MountOptions, each)
		}
	}
	return sc, nil
}