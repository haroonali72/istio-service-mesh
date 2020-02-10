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
		if input.ServiceAttributes.ScParameters.AwsebsscParm["encrypted"] != "" {
			sc.Parameters["encrypted"] = input.ServiceAttributes.ScParameters.AwsebsscParm["encrypted"]
		}
		if kmsKeyId := input.ServiceAttributes.ScParameters.AwsebsscParm["kmsKeyId"]; kmsKeyId != "" {
			sc.Parameters["kmsKeyId"] = input.ServiceAttributes.ScParameters.AwsebsscParm["kmsKeyId"]
		}
		if fsType := input.ServiceAttributes.ScParameters.AwsebsscParm["fsType"]; fsType != "" {
			sc.Parameters["fsType"] = fsType
		}
		for _, each := range input.ServiceAttributes.MountOptions {
			sc.MountOptions = append(sc.MountOptions, each)
		}
		if input.ServiceAttributes.ScParameters.AwsebsscParm["zone"] != "" {
			sc.Parameters["zone"] = input.ServiceAttributes.ScParameters.AwsebsscParm["zone"]
		} else if input.ServiceAttributes.ScParameters.AwsebsscParm["zones"] != "" {
			sc.Parameters["zones"] = input.ServiceAttributes.ScParameters.AwsebsscParm["zones"]
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
		for _, each := range input.ServiceAttributes.MountOptions {
			sc.MountOptions = append(sc.MountOptions, each)
		}
		if input.ServiceAttributes.ScParameters.GcppdscParm["zone"] != "" {
			sc.Parameters["zone"] = input.ServiceAttributes.ScParameters.GcppdscParm["zone"]
		} else if input.ServiceAttributes.ScParameters.GcppdscParm["zones"] != "" {
			sc.Parameters["zones"] = input.ServiceAttributes.ScParameters.GcppdscParm["zones"]
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
		for _, each := range input.ServiceAttributes.MountOptions {
			sc.MountOptions = append(sc.MountOptions, each)
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
		if re := input.ServiceAttributes.ScParameters.AzurfilescParm["readOnly"]; re != "" {
			sc.Parameters["readOnly"] = re
		}
		for _, each := range input.ServiceAttributes.MountOptions {
			sc.MountOptions = append(sc.MountOptions, each)
		}
	}
	return sc, nil
}
