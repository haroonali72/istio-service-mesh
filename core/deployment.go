package core

import (
	"context"
	"encoding/json"
	"errors"
	"google.golang.org/grpc"
	"istio-service-mesh/constants"
	pb "istio-service-mesh/core/proto"
	"istio-service-mesh/types"
	"istio-service-mesh/utils"
	"k8s.io/api/apps/v1"
	v2 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"strings"
)

func (s *Server) CreateDeployment(ctx context.Context, req *pb.DeploymentService) (*pb.ServiceResponse, error) {
	utils.Info.Println(ctx)
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getDeploymentRequestObject(req)
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

func (s *Server) GetDeployment(ctx context.Context, req *pb.DeploymentService) (*pb.ServiceResponse, error) {
	utils.Info.Println(ctx)
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getDeploymentRequestObject(req)
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

func (s *Server) DeleteDeployment(ctx context.Context, req *pb.DeploymentService) (*pb.ServiceResponse, error) {
	utils.Info.Println(ctx)
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getDeploymentRequestObject(req)
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

func (s *Server) PatchDeployment(ctx context.Context, req *pb.DeploymentService) (*pb.ServiceResponse, error) {
	utils.Info.Println(ctx)
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getDeploymentRequestObject(req)
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

func (s *Server) PutDeployment(ctx context.Context, req *pb.DeploymentService) (*pb.ServiceResponse, error) {
	utils.Info.Println(ctx)
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getDeploymentRequestObject(req)
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

func getDeploymentRequestObject(service *pb.DeploymentService) (*v1.Deployment, error) {

	var deployment = new(v1.Deployment)
	if service.Name == "" {
		return nil, errors.New("Service name not found")
	}

	if service.Namespace == "" {
		deployment.ObjectMeta.Namespace = "default"
	} else {
		deployment.ObjectMeta.Namespace = service.Namespace
	}
	deployment.Name = service.Name + "-" + service.Version

	deployment.TypeMeta.Kind = "Deployment"
	deployment.TypeMeta.APIVersion = "apps/v1"

	deployment.Labels = make(map[string]string)
	deployment.Labels["keel.sh/policy"] = "force"
	for key, value := range service.ServiceAttributes.Labels {
		deployment.Labels[key] = value
	}

	deployment.Annotations = make(map[string]string)
	deployment.Annotations = service.ServiceAttributes.Annotations

	deployment.Spec.Selector = new(metav1.LabelSelector)
	deployment.Spec.Selector.MatchLabels = make(map[string]string)
	deployment.Spec.Selector.MatchLabels["app"] = service.Name
	deployment.Spec.Selector.MatchLabels["version"] = service.Version
	if service.ServiceAttributes.LabelSelector != nil {
		deployment.Spec.Selector.MatchLabels = service.ServiceAttributes.LabelSelector.MatchLabels
	}
	for key, value := range service.ServiceAttributes.LabelSelector.MatchLabels {
		deployment.Spec.Selector.MatchLabels[key] = value
	}

	deployment.Spec.Template.Labels = make(map[string]string)
	deployment.Spec.Template.Labels["app"] = service.Name
	deployment.Spec.Template.Labels["version"] = service.Version
	for key, value := range service.ServiceAttributes.Labels {
		deployment.Spec.Template.Labels[key] = value
	}

	deployment.Spec.Template.Annotations = make(map[string]string)
	deployment.Spec.Template.Annotations["sidecar.istio.io/inject"] = "true"
	deployment.Spec.Template.Spec.NodeSelector = make(map[string]string)
	deployment.Spec.Template.Spec.NodeSelector = service.ServiceAttributes.NodeSelector

	if service.ServiceAttributes.Replicas != nil {
		deployment.Spec.Replicas = &service.ServiceAttributes.Replicas.Value
	}

	if service.ServiceAttributes.TerminationGracePeriodSeconds != nil {
		deployment.Spec.Template.Spec.TerminationGracePeriodSeconds = &service.ServiceAttributes.TerminationGracePeriodSeconds.Value
	}

	for _, g := range service.ServiceAttributes.ImagePullSecrets {
		if g != nil {
			pullImageSecret := v2.LocalObjectReference{Name: g.Name}
			deployment.Spec.Template.Spec.ImagePullSecrets = append(deployment.Spec.Template.Spec.ImagePullSecrets, pullImageSecret)
		}
	}

	if service.ServiceAttributes.ServiceAccountName != "" {
		deployment.Spec.Template.Spec.ServiceAccountName = service.ServiceAttributes.ServiceAccountName
	}

	if service.ServiceAttributes.AutomountServiceAccountToken != nil {
		deployment.Spec.Template.Spec.AutomountServiceAccountToken = &service.ServiceAttributes.AutomountServiceAccountToken.Value
	}

	if service.ServiceAttributes.Strategy != nil {
		if service.ServiceAttributes.Strategy.Type == pb.DeploymentStrategyType_Recreate {
			deployment.Spec.Strategy.Type = v1.RecreateDeploymentStrategyType
		} else if service.ServiceAttributes.Strategy.Type == pb.DeploymentStrategyType_RollingUpdate {
			deployment.Spec.Strategy.Type = v1.RollingUpdateDeploymentStrategyType
			if service.ServiceAttributes.Strategy.RollingUpdate != nil {
				deployment.Spec.Strategy.RollingUpdate = &v1.RollingUpdateDeployment{
					MaxUnavailable: &intstr.IntOrString{
						IntVal: service.ServiceAttributes.Strategy.RollingUpdate.MaxUnavailable,
					},
					MaxSurge: &intstr.IntOrString{
						IntVal: service.ServiceAttributes.Strategy.RollingUpdate.MaxSurge,
					},
				}
			}

		}
	}

	var volumeMountNames1 = make(map[string]bool)
	if containersList, volumeMounts, err := getContainers(service.ServiceAttributes.Containers); err == nil {
		if len(containersList) > 0 {
			deployment.Spec.Template.Spec.Containers = containersList
			volumeMountNames1 = volumeMounts
		} else {
			return nil, errors.New("no container exists")
		}

	} else {
		return nil, err
	}

	if containersList, volumeMounts, err := getContainers(service.ServiceAttributes.InitContainers); err == nil {
		if len(containersList) > 0 {
			deployment.Spec.Template.Spec.InitContainers = containersList
			for k, v := range volumeMounts {
				volumeMountNames1[k] = v
			}
		}

	} else {
		return nil, err
	}

	if volumes, err := getVolumes(service.ServiceAttributes.Volumes, volumeMountNames1); err == nil {
		if len(volumes) > 0 {
			deployment.Spec.Template.Spec.Volumes = volumes
		}

	} else {
		return nil, err
	}

	if service.ServiceAttributes.Affinity != nil {
		if aa, err := getAffinity(service.ServiceAttributes.Affinity); err != nil {
			return nil, err
		} else {
			deployment.Spec.Template.Spec.Affinity = aa
		}
	}
	return deployment, nil
}

func getVolumes(vols []*pb.Volume, volumeMountNames map[string]bool) ([]v2.Volume, error) {

	var volumes []v2.Volume
	for _, volume := range vols {

		if !volumeMountNames[volume.Name] {
			continue
		}
		volumeMountNames[volume.Name] = false
		tempVolume := v2.Volume{}
		tempVolume.Name = volume.Name

		if volume.VolumeSource.Secret != nil {
			tempVolume.Secret = new(v2.SecretVolumeSource)
			tempVolume.Secret.SecretName = volume.VolumeSource.Secret.SecretName
			tempVolume.Secret.DefaultMode = &volume.VolumeSource.Secret.DefaultMode
			var secretItems []v2.KeyToPath
			for _, item := range volume.VolumeSource.Secret.Items {
				secretItem := v2.KeyToPath{
					Key:  item.Key,
					Path: item.Path,
					Mode: &item.Mode,
				}
				secretItems = append(secretItems, secretItem)
			}
			tempVolume.Secret.Items = secretItems
		}
		if volume.VolumeSource.ConfigMap != nil {
			tempVolume.ConfigMap = new(v2.ConfigMapVolumeSource)
			tempVolume.ConfigMap.Name = volume.VolumeSource.ConfigMap.LocalObjectReference.Name

			tempVolume.ConfigMap.DefaultMode = &volume.VolumeSource.ConfigMap.DefaultMode
			var configMapItems []v2.KeyToPath
			for _, item := range volume.VolumeSource.ConfigMap.Items {
				configMapItem := v2.KeyToPath{
					Key:  item.Key,
					Path: item.Path,
					Mode: &item.Mode,
				}
				configMapItems = append(configMapItems, configMapItem)
			}
			tempVolume.ConfigMap.Items = configMapItems
		}

		if volume.VolumeSource.AwsElasticBlockStore != nil {
			tempVolume.AWSElasticBlockStore = new(v2.AWSElasticBlockStoreVolumeSource)
			tempVolume.AWSElasticBlockStore.ReadOnly = volume.VolumeSource.AwsElasticBlockStore.ReadOnly
			tempVolume.AWSElasticBlockStore.Partition = volume.VolumeSource.AwsElasticBlockStore.Partition
		}

		if volume.VolumeSource.EmptyDir != nil {
			tempVolume.EmptyDir = new(v2.EmptyDirVolumeSource)
			quantity, _ := resource.ParseQuantity(volume.VolumeSource.EmptyDir.SizeLimit)
			tempVolume.EmptyDir.SizeLimit = &quantity
			if volume.VolumeSource.EmptyDir.Medium.String() == pb.StorageMedium_StorageMediumDefault.String() {
				tempVolume.EmptyDir.Medium = v2.StorageMediumDefault

			}
			if volume.VolumeSource.EmptyDir.Medium.String() == pb.StorageMedium_Memory.String() {
				tempVolume.EmptyDir.Medium = v2.StorageMediumMemory
			}

			if volume.VolumeSource.EmptyDir.Medium.String() == pb.StorageMedium_HugePages.String() {
				tempVolume.EmptyDir.Medium = v2.StorageMediumHugePages
			}

		}

		if volume.VolumeSource.GcePersistentDisk != nil {
			tempVolume.GCEPersistentDisk = new(v2.GCEPersistentDiskVolumeSource)
			tempVolume.GCEPersistentDisk.Partition = volume.VolumeSource.GcePersistentDisk.Partition
			tempVolume.GCEPersistentDisk.ReadOnly = volume.VolumeSource.GcePersistentDisk.ReadOnly
			tempVolume.GCEPersistentDisk.PDName = volume.VolumeSource.GcePersistentDisk.PdName
		}

		if volume.VolumeSource.AzureDisk != nil {
			tempVolume.AzureFile = new(v2.AzureFileVolumeSource)
			tempVolume.AzureDisk.ReadOnly = &volume.VolumeSource.AzureDisk.ReadOnly
			tempVolume.AzureDisk.DataDiskURI = volume.VolumeSource.AzureDisk.DiskURI

			if volume.VolumeSource.AzureDisk.CachingMode.String() == pb.AzureDataDiskCachingMode_ModeNone.String() {
				temp := v2.AzureDataDiskCachingNone
				tempVolume.AzureDisk.CachingMode = &temp
			} else if volume.VolumeSource.AzureDisk.CachingMode.String() == pb.AzureDataDiskCachingMode_ReadOnly.String() {
				temp := v2.AzureDataDiskCachingReadOnly
				tempVolume.AzureDisk.CachingMode = &temp
			} else if volume.VolumeSource.AzureDisk.CachingMode.String() == pb.AzureDataDiskCachingMode_ReadWrite.String() {
				temp := v2.AzureDataDiskCachingReadWrite
				tempVolume.AzureDisk.CachingMode = &temp
			}

			if volume.VolumeSource.AzureDisk.Kind.String() == pb.AzureDataDiskKind_Shared.String() {
				temp := v2.AzureSharedBlobDisk
				tempVolume.AzureDisk.Kind = &temp
			} else if volume.VolumeSource.AzureDisk.Kind.String() == pb.AzureDataDiskKind_Dedicated.String() {
				temp := v2.AzureDedicatedBlobDisk
				tempVolume.AzureDisk.Kind = &temp
			} else if volume.VolumeSource.AzureDisk.Kind.String() == pb.AzureDataDiskKind_Managed.String() {
				temp := v2.AzureManagedDisk
				tempVolume.AzureDisk.Kind = &temp
			}
		}

		if volume.VolumeSource.AzureFile != nil {
			tempVolume.AzureFile = new(v2.AzureFileVolumeSource)
			tempVolume.AzureFile.ReadOnly = volume.VolumeSource.AzureFile.ReadOnly
			tempVolume.AzureFile.SecretName = volume.VolumeSource.AzureFile.SecretName
			tempVolume.AzureFile.ShareName = volume.VolumeSource.AzureFile.ShareName

		}
		if volume.VolumeSource.HostPath != nil {
			tempVolume.HostPath = new(v2.HostPathVolumeSource)
			tempVolume.HostPath.Path = volume.VolumeSource.HostPath.Path
			hostPathType := volume.VolumeSource.HostPath.Type.String()
			hostPathTypeTemp := v2.HostPathType(hostPathType)
			tempVolume.HostPath.Type = &hostPathTypeTemp
		}

		volumes = append(volumes, tempVolume)

	}
	for key, _ := range volumeMountNames {
		if volumeMountNames[key] == true {
			return nil, errors.New("volume does not exist")
		}
	}
	return volumes, nil

}

func getContainers(conts map[string]*pb.ContainerAttributes) ([]v2.Container, map[string]bool, error) {

	volumeMountNames := make(map[string]bool)

	var containers []v2.Container

	for key, container := range conts {
		var containerTemp v2.Container
		containerTemp.Name = key
		if err := putCommandAndArguments(&containerTemp, container.Command, container.Args); err != nil {
			return nil, nil, err
		}
		err, isOk := checkRequestIsLessThanLimit(container)

		if err != nil {
			return nil, nil, err
		} else if isOk == false {
			return nil, nil, errors.New("Request Resource is greater limit resource")

		}

		if err := putReadinessProbe(&containerTemp, container.ReadinessProbe); err != nil {
			return nil, nil, err
		}
		if err := putLivenessProbe(&containerTemp, container.LivenessProbe); err != nil {
			return nil, nil, err
		}
		if err := putLimitResource(&containerTemp, container.LimitResources); err != nil {
			return nil, nil, err
		}
		if err := putRequestResource(&containerTemp, container.RequestResources); err != nil {
			return nil, nil, err
		}
		if container.SecurityContext != nil {
			if securityContext, err := configureSecurityContext(container.SecurityContext); err != nil {
				return nil, nil, err
			} else {

				containerTemp.SecurityContext = securityContext
			}
		}

		containerTemp.Image = container.ImagePrefix + container.ImageName
		if container.Tag != "" {
			containerTemp.Image += ":" + container.Tag
		}
		// volume mounts
		var volumeMounts []v2.VolumeMount
		for _, volumeMount := range container.VolumeMounts {
			volumeMountNames[volumeMount.Name] = true
			temp := v2.VolumeMount{}
			temp.Name = volumeMount.Name
			temp.MountPath = volumeMount.MountPath
			temp.SubPath = volumeMount.SubPath
			temp.SubPathExpr = volumeMount.SubPathExpr
			if volumeMount.MountPropagation.String() == pb.MountPropagationMode_None.String() {
				none := v2.MountPropagationNone
				temp.MountPropagation = &none

			}

			if volumeMount.MountPropagation.String() == pb.MountPropagationMode_HostToContainer.String() {
				htc := v2.MountPropagationNone
				temp.MountPropagation = &htc

			}
			if volumeMount.MountPropagation.String() == pb.MountPropagationMode_Bidirectional.String() {
				bi := v2.MountPropagationBidirectional
				temp.MountPropagation = &bi

			}
			volumeMounts = append(volumeMounts, temp)

		}

		var ports []v2.ContainerPort
		for key, port := range container.Ports {
			temp := v2.ContainerPort{}
			temp.Name = key
			if port.ContainerPort == 0 && port.HostPort == 0 {
				continue
			}
			if port.ContainerPort == 0 && port.HostPort != 0 {
				port.ContainerPort = port.HostPort
			}

			if port.ContainerPort > 0 && port.ContainerPort < 65536 {
				temp.ContainerPort = port.ContainerPort
			} else {
				utils.Info.Println("invalid prot number")
				continue
			}
			if port.HostPort != 0 {

				if port.HostPort > 0 && port.HostPort < 65536 {
					temp.HostPort = port.HostPort
				} else {
					utils.Info.Println("invalid prot number")
					continue
				}

			}
			ports = append(ports, temp)
		}

		var envVariables []v2.EnvVar
		for key, envVariable := range container.EnvironmentVariables {
			tempEnvVariable := v2.EnvVar{}
			if strings.EqualFold(key, "ConfigMap") {
				envVariableValue := strings.Split(envVariable.Value, ";")
				tempEnvVariable = v2.EnvVar{Name: envVariable.Key,
					ValueFrom: &v2.EnvVarSource{ConfigMapKeyRef: &v2.ConfigMapKeySelector{
						LocalObjectReference: v2.LocalObjectReference{Name: envVariableValue[0]},
						Key:                  envVariableValue[1],
					}}}

			} else if strings.EqualFold(key, "Secret") {
				envVariableValue := strings.Split(envVariable.Value, ";")
				tempEnvVariable = v2.EnvVar{Name: envVariable.Key,
					ValueFrom: &v2.EnvVarSource{SecretKeyRef: &v2.SecretKeySelector{
						LocalObjectReference: v2.LocalObjectReference{Name: envVariableValue[0]},
						Key:                  envVariableValue[1],
					}}}
			} else {
				tempEnvVariable = v2.EnvVar{Name: envVariable.Key, Value: envVariable.Value}
			}
			envVariables = append(envVariables, tempEnvVariable)
		}

		containerTemp.Ports = ports
		containerTemp.Env = envVariables
		containerTemp.VolumeMounts = volumeMounts

		containers = append(containers, containerTemp)

	}

	return containers, volumeMountNames, nil

}

func getAffinity(affinity *pb.Affinity) (*v2.Affinity, error) {
	temp := new(v2.Affinity)
	if affinity.NodeAffinity != nil {
		na, err := getNodeAffinity(affinity.NodeAffinity)
		if err != nil {
			return nil, err
		} else {
			temp.NodeAffinity = na
		}
	}

	if affinity.PodAffinity != nil {
		if pa, err := getPodAffinity(affinity.PodAffinity); err != nil {
			return nil, err
		} else {
			temp.PodAffinity = pa
		}
	}

	if affinity.PodAntiAffinity != nil {
		if paa, err := getAntiPodAffinity(affinity.PodAntiAffinity); err != nil {
			return nil, err
		} else {
			temp.PodAntiAffinity = paa
		}
	}
	return temp, nil
}

func getNodeAffinity(nodeAffinity *pb.NodeAffinity) (*v2.NodeAffinity, error) {
	temp := new(v2.NodeAffinity)
	if nodeAffinity.ReqDuringSchedulingIgnDuringExec != nil {
		if ns, err := getNodeSelector(nodeAffinity.ReqDuringSchedulingIgnDuringExec); err != nil {
			return nil, err
		} else {
			temp.RequiredDuringSchedulingIgnoredDuringExecution = ns
		}
	}

	var tempPrefSchedulingTerms []v2.PreferredSchedulingTerm
	for _, prefSchedulingTerm := range nodeAffinity.PrefDuringIgnDuringExec {
		tempPrefSchedulingTerm := v2.PreferredSchedulingTerm{}
		if prefSchedulingTerm != nil {

			tempPrefSchedulingTerm.Weight = prefSchedulingTerm.Weight
			var tempMatchExpressions []v2.NodeSelectorRequirement
			var tempMatchFields []v2.NodeSelectorRequirement

			if prefSchedulingTerm.Preference != nil {
				for _, matchExpression := range prefSchedulingTerm.Preference.MatchExpressions {
					tempMatchExpression := v2.NodeSelectorRequirement{}
					tempMatchExpression.Key = matchExpression.Key
					tempMatchExpression.Values = matchExpression.Values
					tempMatchExpression.Operator = v2.NodeSelectorOperator(strings.Trim(matchExpression.Operator.String(), "NodeSelectorOp"))
					tempMatchExpressions = append(tempMatchExpressions, tempMatchExpression)
				}
				for _, matchField := range prefSchedulingTerm.Preference.MatchFields {
					tempMatchField := v2.NodeSelectorRequirement{}
					tempMatchField.Key = matchField.Key
					tempMatchField.Values = matchField.Values
					tempMatchField.Operator = v2.NodeSelectorOperator(strings.Trim(matchField.Operator.String(), "NodeSelectorOp"))
					tempMatchFields = append(tempMatchFields, tempMatchField)
				}
				tempPrefSchedulingTerm.Preference.MatchExpressions = tempMatchExpressions
				tempPrefSchedulingTerm.Preference.MatchFields = tempMatchFields

			}
		}
		tempPrefSchedulingTerms = append(tempPrefSchedulingTerms, tempPrefSchedulingTerm)

	}
	return temp, nil
}

func getPodAffinity(podAffinity *pb.PodAffinity) (*v2.PodAffinity, error) {
	temp := new(v2.PodAffinity)
	var tempPodAffinityTerms []v2.PodAffinityTerm
	for _, podAffinityTerm := range podAffinity.ReqDuringSchedulingIgnDuringExec {
		tempPodAffinityTerm := v2.PodAffinityTerm{}
		if podAffinityTerm != nil {
			tempPodAffinityTerm.Namespaces = podAffinityTerm.Namespaces
			tempPodAffinityTerm.TopologyKey = podAffinityTerm.TopologyKey
			if ls, err := getLabelSelector(podAffinityTerm.LabelSelector); err != nil {
				return nil, err
			} else {
				tempPodAffinityTerm.LabelSelector = ls
			}
		}
		tempPodAffinityTerms = append(tempPodAffinityTerms, tempPodAffinityTerm)

	}
	temp.RequiredDuringSchedulingIgnoredDuringExecution = tempPodAffinityTerms
	var tempWeightedAffinityTerms []v2.WeightedPodAffinityTerm
	for _, weighted := range podAffinity.PrefDuringIgnDuringExec {
		tempWeightedAffinityTerm := v2.WeightedPodAffinityTerm{}
		if weighted != nil {
			tempWeightedAffinityTerm.Weight = weighted.Weight
			if weighted.PodAffinityTerm != nil {
				tempPodAffinityTerm := v2.PodAffinityTerm{}
				tempPodAffinityTerm.Namespaces = weighted.PodAffinityTerm.Namespaces
				tempPodAffinityTerm.TopologyKey = weighted.PodAffinityTerm.TopologyKey
				if ls, err := getLabelSelector(weighted.PodAffinityTerm.LabelSelector); err != nil {
					return nil, err
				} else {
					tempPodAffinityTerm.LabelSelector = ls
				}

				tempWeightedAffinityTerm.PodAffinityTerm = tempPodAffinityTerm
			}

		}
		tempWeightedAffinityTerms = append(tempWeightedAffinityTerms, tempWeightedAffinityTerm)
	}
	temp.PreferredDuringSchedulingIgnoredDuringExecution = tempWeightedAffinityTerms
	return temp, nil
}

func getAntiPodAffinity(podAntiAffinity *pb.PodAntiAffinity) (*v2.PodAntiAffinity, error) {
	temp := new(v2.PodAntiAffinity)
	var tempPodAffinityTerms []v2.PodAffinityTerm
	for _, podAffinityTerm := range podAntiAffinity.ReqDuringSchedulingIgnDuringExec {
		tempPodAffinityTerm := v2.PodAffinityTerm{}
		if podAffinityTerm != nil {
			tempPodAffinityTerm.Namespaces = podAffinityTerm.Namespaces
			tempPodAffinityTerm.TopologyKey = podAffinityTerm.TopologyKey

			if ls, err := getLabelSelector(podAffinityTerm.LabelSelector); err != nil {
				return nil, err
			} else {
				tempPodAffinityTerm.LabelSelector = ls
			}

		}
		tempPodAffinityTerms = append(tempPodAffinityTerms, tempPodAffinityTerm)

	}
	temp.RequiredDuringSchedulingIgnoredDuringExecution = tempPodAffinityTerms
	var tempWeightedAffinityTerms []v2.WeightedPodAffinityTerm
	for _, weighted := range podAntiAffinity.PrefDuringIgnDuringExec {
		tempWeightedAffinityTerm := v2.WeightedPodAffinityTerm{}
		if weighted != nil {
			tempWeightedAffinityTerm.Weight = weighted.Weight
			if weighted.PodAffinityTerm != nil {
				tempPodAffinityTerm := v2.PodAffinityTerm{}
				tempPodAffinityTerm.Namespaces = weighted.PodAffinityTerm.Namespaces
				tempPodAffinityTerm.TopologyKey = weighted.PodAffinityTerm.TopologyKey
				if ls, err := getLabelSelector(weighted.PodAffinityTerm.LabelSelector); err != nil {
					return nil, err
				} else {
					tempPodAffinityTerm.LabelSelector = ls
				}

				tempWeightedAffinityTerm.PodAffinityTerm = tempPodAffinityTerm
			}

		}
		tempWeightedAffinityTerms = append(tempWeightedAffinityTerms, tempWeightedAffinityTerm)
	}
	temp.PreferredDuringSchedulingIgnoredDuringExecution = tempWeightedAffinityTerms
	return temp, nil
}

//func setLabelSelector(temp *metav1.LabelSelector, labelSelector *pb.LabelSelectorObj) {
//	temp = new(metav1.LabelSelector)
//	temp.MatchLabels = make(map[string]string)
//	temp.MatchLabels = labelSelector.MatchLabels
//	var tempLabelSelectorRequirements []metav1.LabelSelectorRequirement
//	for _, labelSelectorRequirement := range labelSelector.MatchExpressions {
//		tempLabelSelectorRequirement := metav1.LabelSelectorRequirement{}
//		tempLabelSelectorRequirement.Key = labelSelectorRequirement.Key
//		tempLabelSelectorRequirement.Values = labelSelectorRequirement.Values
//		tempLabelSelectorRequirement.Operator = metav1.LabelSelectorOperator(labelSelectorRequirement.Operator.String())
//		tempLabelSelectorRequirements = append(tempLabelSelectorRequirements, tempLabelSelectorRequirement)
//	}
//	temp.MatchExpressions = tempLabelSelectorRequirements
//
//}

func getNodeSelector(nodeSelector *pb.NodeSelector) (*v2.NodeSelector, error) {
	if nodeSelector != nil {
		temp := new(v2.NodeSelector)
		var nodeSelectorTerms []v2.NodeSelectorTerm
		for _, nodeSelectorTerm := range nodeSelector.NodeSelectorTerms {
			var tempMatchExpressions []v2.NodeSelectorRequirement
			var tempMatchFields []v2.NodeSelectorRequirement
			tempNodeSelectorTerm := v2.NodeSelectorTerm{}
			if nodeSelectorTerm != nil {

				for _, matchExpression := range nodeSelectorTerm.MatchExpressions {
					tempMatchExpression := v2.NodeSelectorRequirement{}
					tempMatchExpression.Key = matchExpression.Key
					tempMatchExpression.Values = matchExpression.Values
					tempMatchExpression.Operator = v2.NodeSelectorOperator(strings.Trim(matchExpression.Operator.String(), "NodeSelectorOp"))
					tempMatchExpressions = append(tempMatchExpressions, tempMatchExpression)
				}
				for _, matchField := range nodeSelectorTerm.MatchFields {
					tempMatchField := v2.NodeSelectorRequirement{}
					tempMatchField.Key = matchField.Key
					tempMatchField.Values = matchField.Values
					tempMatchField.Operator = v2.NodeSelectorOperator(strings.Trim(matchField.Operator.String(), "NodeSelectorOp"))
					tempMatchFields = append(tempMatchFields, tempMatchField)
				}
			}
			tempNodeSelectorTerm.MatchFields = tempMatchFields
			tempNodeSelectorTerm.MatchExpressions = tempMatchExpressions
			nodeSelectorTerms = append(nodeSelectorTerms, tempNodeSelectorTerm)
		}
		temp.NodeSelectorTerms = nodeSelectorTerms
		return temp, nil
	}
	return nil, nil

}

func getInitContainers(service *pb.DeploymentService) ([]v2.Container, []string, error) {
	var volumeMountNames []string
	var containers []v2.Container

	for _, container := range service.ServiceAttributes.Containers {
		var containerTemp v2.Container
		containerTemp.Name = service.Name

		if err := putCommandAndArguments(&containerTemp, container.Command, container.Args); err != nil {
			return nil, volumeMountNames, err
		}
		err, isOk := checkRequestIsLessThanLimit(container)

		if err != nil {
			return nil, volumeMountNames, err
		} else if isOk == false {
			return nil, volumeMountNames, errors.New("Request Resource is greater limit resource")

		}
		if err := putLivenessProbe(&containerTemp, container.LivenessProbe); err != nil {
			return nil, volumeMountNames, err
		}
		if err := putLimitResource(&containerTemp, container.LimitResources); err != nil {
			return nil, volumeMountNames, err
		}
		if err := putRequestResource(&containerTemp, container.RequestResources); err != nil {
			return nil, volumeMountNames, err
		}
		if container.SecurityContext != nil {
			if securityContext, err := configureSecurityContext(container.SecurityContext); err != nil {
				return nil, volumeMountNames, err
			} else {
				containerTemp.SecurityContext = securityContext
			}
		}

		containerTemp.Image = container.ImagePrefix + container.ImageName
		if container.Tag != "" {
			containerTemp.Image += ":" + container.Tag
		}

		var ports []v2.ContainerPort
		for _, port := range container.Ports {
			temp := v2.ContainerPort{}
			if port.ContainerPort == 0 && port.HostPort == 0 {
				continue
			}
			if port.ContainerPort == 0 && port.HostPort != 0 {
				port.ContainerPort = port.HostPort
			}

			if port.ContainerPort > 0 && port.ContainerPort < 65536 {
				temp.ContainerPort = port.ContainerPort
			} else {
				utils.Info.Println("invalid prot number")
				continue
			}
			if port.HostPort != 0 {

				if port.HostPort > 0 && port.HostPort < 65536 {
					temp.HostPort = port.HostPort
				} else {
					utils.Info.Println("invalid prot number")
					continue
				}

			}
			ports = append(ports, temp)
		}

		var envVariables []v2.EnvVar
		for key, envVariable := range container.EnvironmentVariables {
			tempEnvVariable := v2.EnvVar{}
			if strings.EqualFold(key, "ConfigMap") {
				envVariableValue := strings.Split(envVariable.Value, ";")
				tempEnvVariable = v2.EnvVar{Name: envVariable.Key,
					ValueFrom: &v2.EnvVarSource{ConfigMapKeyRef: &v2.ConfigMapKeySelector{
						LocalObjectReference: v2.LocalObjectReference{Name: envVariableValue[0]},
						Key:                  envVariableValue[1],
					}}}

			} else if strings.EqualFold(key, "Secret") {
				envVariableValue := strings.Split(envVariable.Value, ";")
				tempEnvVariable = v2.EnvVar{Name: envVariable.Key,
					ValueFrom: &v2.EnvVarSource{SecretKeyRef: &v2.SecretKeySelector{
						LocalObjectReference: v2.LocalObjectReference{Name: envVariableValue[0]},
						Key:                  envVariableValue[1],
					}}}
			} else {
				tempEnvVariable = v2.EnvVar{Name: envVariable.Key, Value: envVariable.Value}
			}
			envVariables = append(envVariables, tempEnvVariable)
		}

		containerTemp.Ports = ports
		containerTemp.Env = envVariables

		containers = append(containers, containerTemp)

	}

	return containers, volumeMountNames, nil

}

func putCommandAndArguments(container *v2.Container, command, args []string) error {
	if len(command) > 0 && command[0] != "" {
		container.Command = command
		if len(args) > 0 {
			container.Args = args
		} else {
			container.Args = []string{}
		}

	} else if len(args) > 0 {
		container.Args = args
	}
	return nil
}

func checkRequestIsLessThanLimit(serviceAttr *pb.ContainerAttributes) (error, bool) {
	for t, v := range serviceAttr.LimitResources {
		r, found := serviceAttr.RequestResources[t]
		if found {
			rr, err := resource.ParseQuantity(r)
			if err != nil {
				return err, false
			}
			lr, err := resource.ParseQuantity(v)
			if err != nil {
				return err, false
			}
			rrint := rr.AsDec()
			lrint := lr.AsDec()
			if rrint.Cmp(lrint) == 1 {
				return nil, false
			}
		}
	}
	return nil, true
}

func putLivenessProbe(container *v2.Container, prob *pb.Probe) error {

	var temp v2.Probe
	if prob != nil {
		if prob.Handler != nil {
			temp.InitialDelaySeconds = prob.InitialDelaySeconds
			temp.FailureThreshold = prob.FailureThreshold
			temp.PeriodSeconds = prob.PeriodSeconds
			temp.SuccessThreshold = prob.SuccessThreshold
			temp.TimeoutSeconds = prob.TimeoutSeconds
			typeHandler := prob.Handler.HandlerType
			switch typeHandler {
			case "exec":
				if prob.Handler.Exec == nil {
					return errors.New("there is no liveness handler of exec type")
				}
				temp.Handler.Exec = &v2.ExecAction{}
				for i := 0; i < len(prob.Handler.Exec.Command); i++ {
					temp.Handler.Exec.Command = append(temp.Handler.Exec.Command, prob.Handler.Exec.Command[i])
				}

			case "http_get":
				if prob.Handler.HttpGet == nil {
					return errors.New("there is no liveness handler of httpGet type")
				}
				temp.Handler.HTTPGet = &v2.HTTPGetAction{}
				if prob.Handler.HttpGet.Port > 0 && prob.Handler.HttpGet.Port < 65536 {

					temp.HTTPGet.Host = prob.Handler.HttpGet.Host

					temp.HTTPGet.Path = prob.Handler.HttpGet.Path

					if strings.EqualFold(prob.Handler.HttpGet.Scheme, types.URISchemeHTTP) || strings.EqualFold(prob.Handler.HttpGet.Scheme, types.URISchemeHTTPS) {

						temp.HTTPGet.Scheme = v2.URIScheme(prob.Handler.HttpGet.Scheme)
					}

					if prob.Handler.HttpGet.HttpHeaders != nil {
						temp.HTTPGet.HTTPHeaders = []v2.HTTPHeader{}
						for i := 0; i < len(prob.Handler.HttpGet.HttpHeaders); i++ {
							tempheader := v2.HTTPHeader{prob.Handler.HttpGet.HttpHeaders[i].Name, prob.Handler.HttpGet.HttpHeaders[i].Value}
							temp.HTTPGet.HTTPHeaders = append(temp.HTTPGet.HTTPHeaders, tempheader)
						}
					}
					temp.HTTPGet.Port = intstr.FromInt(int(prob.Handler.HttpGet.Port))
				} else {
					return errors.New("Invalid Port number for http Get")
				}
			case "tcpSocket":
				if prob.Handler.TcpSocket == nil {
					return errors.New("there is no liveness handler of tcpSocket type")
				}
				temp.Handler.TCPSocket = &v2.TCPSocketAction{}
				if prob.Handler.TcpSocket.Port > 0 && prob.Handler.TcpSocket.Port < 65536 {
					temp.TCPSocket.Port = intstr.FromInt(int(prob.Handler.TcpSocket.Port))
					temp.TCPSocket.Host = prob.Handler.TcpSocket.Host

				} else {
					return errors.New("Invalid Port number for tcp socket")
				}

			default:
				return errors.New("There Must be liveness handler of valid type")

			}
		} else {
			return errors.New("Liveness prob header can not be nil")
		}
		container.LivenessProbe = &temp
	}
	return nil
}
func putReadinessProbe(container *v2.Container, prob *pb.Probe) error {
	var temp v2.Probe
	if prob != nil {
		if prob.Handler != nil {
			temp.InitialDelaySeconds = prob.InitialDelaySeconds
			temp.FailureThreshold = prob.FailureThreshold
			temp.PeriodSeconds = prob.PeriodSeconds
			temp.SuccessThreshold = prob.SuccessThreshold
			temp.TimeoutSeconds = prob.TimeoutSeconds

			switch typeHandler := prob.Handler.HandlerType; typeHandler {
			case "exec":
				if prob.Handler.Exec == nil {
					return errors.New("there is no readiness handler of exec type")
				}
				temp.Handler.Exec = &v2.ExecAction{}
				for i := 0; i < len(prob.Handler.Exec.Command); i++ {
					temp.Handler.Exec.Command = append(temp.Handler.Exec.Command, prob.Handler.Exec.Command[i])
				}

			case "http_get":
				if prob.Handler.HttpGet == nil {
					return errors.New("there is no readiness handler of httpGet type")
				}
				temp.Handler.HTTPGet = &v2.HTTPGetAction{}
				if prob.Handler.HttpGet.Port > 0 && prob.Handler.HttpGet.Port < 65536 {
					temp.HTTPGet.Host = prob.Handler.HttpGet.Host
					temp.HTTPGet.Path = prob.Handler.HttpGet.Path
					temp.HTTPGet.Port = intstr.FromInt(int(prob.Handler.HttpGet.Port))

					if prob.Handler.HttpGet.Scheme == types.URISchemeHTTP && prob.Handler.HttpGet.Scheme == types.URISchemeHTTPS {
						if prob.Handler.HttpGet.Scheme == types.URISchemeHTTP {

							temp.HTTPGet.Scheme = v2.URISchemeHTTP
						} else if prob.Handler.HttpGet.Scheme == types.URISchemeHTTPS {
							temp.HTTPGet.Scheme = v2.URISchemeHTTPS
						}
					}

					temp.HTTPGet.HTTPHeaders = []v2.HTTPHeader{}
					for i := 0; i < len(prob.Handler.HttpGet.HttpHeaders); i++ {
						if prob.Handler.HttpGet.HttpHeaders[i] != nil {
							tempHeader := v2.HTTPHeader{prob.Handler.HttpGet.HttpHeaders[i].Name, prob.Handler.HttpGet.HttpHeaders[i].Value}
							temp.HTTPGet.HTTPHeaders = append(temp.HTTPGet.HTTPHeaders, tempHeader)
						}
					}

				} else {
					return errors.New("Invalid Port number for http Get")
				}
			case "tcpSocket":
				if prob.Handler.TcpSocket == nil {
					return errors.New("there is no readiness handler of tcpSocket type")
				}
				temp.Handler.TCPSocket = &v2.TCPSocketAction{}
				if prob.Handler.TcpSocket.Port > 0 && prob.Handler.TcpSocket.Port < 65536 {
					temp.TCPSocket.Port = intstr.FromInt(int(prob.Handler.TcpSocket.Port))

					temp.TCPSocket.Host = prob.Handler.TcpSocket.Host

				} else {
					return errors.New("Invalid Port number for tcp socket")
				}

			default:
				return errors.New("There Must be readiness handler of valid type")

			}
		} else {
			return errors.New("Readiness prob handler can not be nil")
		}
		container.ReadinessProbe = &temp
	}
	return nil
}

func configureSecurityContext(securityContext *pb.SecurityContextStruct) (*v2.SecurityContext, error) {
	var context v2.SecurityContext
	context.Capabilities = &v2.Capabilities{}
	for _, capability := range securityContext.Capabilities {
		for _, add := range capability.Add {
			context.Capabilities.Add = append(context.Capabilities.Add, v2.Capability(add))
		}
		for _, dropCapability := range capability.Drop {
			context.Capabilities.Drop = append(context.Capabilities.Drop, v2.Capability(dropCapability))
		}
	}
	context.ReadOnlyRootFilesystem = &securityContext.ReadOnlyRootFilesystem
	context.Privileged = &securityContext.Privileged
	if securityContext.RunAsNonRoot && securityContext.RunAsUser == 0 {
		return nil, errors.New("RunAsNonRoot is Set, but RunAsUser value not given!")
	} else {
		context.RunAsNonRoot = &securityContext.RunAsNonRoot
		context.RunAsUser = &securityContext.RunAsUser
	}
	context.RunAsGroup = &securityContext.RunAsGroup
	context.AllowPrivilegeEscalation = &securityContext.AllowPrivilegeEscalation

	var procmount = securityContext.ProcMount.String()

	tempProcMount := v2.ProcMountType(procmount)
	context.ProcMount = &tempProcMount

	context.SELinuxOptions = &v2.SELinuxOptions{
		User:  securityContext.SeLinuxOptions.User,
		Role:  securityContext.SeLinuxOptions.Role,
		Type:  securityContext.SeLinuxOptions.Type,
		Level: securityContext.SeLinuxOptions.Level,
	}
	return &context, nil

}

func putRequestResource(container *v2.Container, requestResources map[string]string) error {
	temp := make(map[v2.ResourceName]resource.Quantity)
	for t, v := range requestResources {
		if t == types.ResourceTypeCpu || t == types.ResourceTypeMemory {
			quantity, err := resource.ParseQuantity(v)
			if err != nil {
				return err
			}
			//
			temp[v2.ResourceName(t)] = quantity
		} else {
			return errors.New("Error Found: Invalid Request Resource Provided. Valid: 'cpu','memory'")
		}
	}

	container.Resources.Requests = temp
	return nil
}

func putLimitResource(container *v2.Container, limitResources map[string]string) error {
	temp := make(map[v2.ResourceName]resource.Quantity)
	for t, v := range limitResources {
		if t == types.ResourceTypeMemory || t == types.ResourceTypeCpu {
			quantity, err := resource.ParseQuantity(v)
			if err != nil {
				return err
			}
			temp[v2.ResourceName(t)] = quantity
		} else {
			return errors.New("Error Found: Invalid Request Resource Provided. Valid: 'cpu','memory'")
		}
	}

	container.Resources.Limits = temp
	return nil
}
