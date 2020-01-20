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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

func getPersistentVolumeClaim(input *pb.PersistentVolumeClaimService) (*core.PersistentVolumeClaim, error) {
	var pvc = new(core.PersistentVolumeClaim)
	pvc.Name = input.Name
	pvc.TypeMeta.Kind = "PersistentVolumeClaim"
	pvc.TypeMeta.APIVersion = "v1"
	pvc.Namespace = input.Namespace
	if input.ServiceAttributes.StorageClassName != "" {
		pvc.Spec.StorageClassName = &input.ServiceAttributes.StorageClassName
	} else if input.ServiceAttributes.VolumeName != "" {
		pvc.Spec.VolumeName = input.ServiceAttributes.VolumeName
		//label selector applied
		lenl := len(input.ServiceAttributes.LabelSelector.MatchLabel)
		lene := len(input.ServiceAttributes.LabelSelector.MatchExpression)

		if (!(lenl > 0)) && lene > 0 {

			pvc.Spec.Selector = &metav1.LabelSelector{nil, nil}
		} else if lene > 0 || lenl > 0 {
			pvc.Spec.Selector = &metav1.LabelSelector{MatchLabels: make(map[string]string)}

		}

		for k, v := range input.ServiceAttributes.LabelSelector.MatchLabel {
			pvc.Spec.Selector.MatchLabels[k] = v
		}
		for i := 0; i < len(input.ServiceAttributes.LabelSelector.MatchExpression); i++ {
			if len(input.ServiceAttributes.LabelSelector.MatchExpression[i].Key) > 0 && (input.ServiceAttributes.LabelSelector.MatchExpression[i].Operator == pb.LabelSelectorOperator_DoesNotExist ||
				input.ServiceAttributes.LabelSelector.MatchExpression[i].Operator == pb.LabelSelectorOperator_Exists ||
				input.ServiceAttributes.LabelSelector.MatchExpression[i].Operator == pb.LabelSelectorOperator_In ||
				input.ServiceAttributes.LabelSelector.MatchExpression[i].Operator == pb.LabelSelectorOperator_NotIn) {
				byteData, err := json.Marshal(input.ServiceAttributes.LabelSelector.MatchExpression[i])
				if err != nil {
					return nil, err
				}
				var temp metav1.LabelSelectorRequirement

				err = json.Unmarshal(byteData, &temp)
				if err != nil {
					return nil, err
				}
				pvc.Spec.Selector.MatchExpressions = append(pvc.Spec.Selector.MatchExpressions, temp)
			}
		}
	}
	if input.ServiceAttributes.String() == pb.PersistentVolumeMode_Filesystem.String() {
		pvm := core.PersistentVolumeFilesystem
		pvc.Spec.VolumeMode = &pvm
	} else if input.ServiceAttributes.VolumeMode.String() == pb.PersistentVolumeMode_Block.String() {
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
