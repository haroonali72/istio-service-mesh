package core

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	pb "istio-service-mesh/core/proto"
	"istio-service-mesh/types"
	"istio-service-mesh/utils"
	"istio.io/client-go/pkg/apis/networking/v1alpha3"
	apps "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	extensions "k8s.io/api/extensions/v1beta1"
	net "k8s.io/api/networking/v1"
	"k8s.io/api/rbac/v1beta1"
	storage "k8s.io/api/storage/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes/scheme"
	"strings"
)

func (s *Server) GetCPService(ctx context.Context, req *pb.YamlToCPServiceRequest) (*pb.YamlToCPServiceResponse, error) {
	serviceResp := new(pb.YamlToCPServiceResponse)
	decode := scheme.Codecs.UniversalDeserializer().Decode
	obj, _, err := decode(req.Service, nil, nil)

	if err != nil {
		yamlService := types.Yamlservice{}
		err := yaml.Unmarshal(req.Service, &yamlService)
		if err == nil {
			if yamlService.Kind == "Gateway" {
				gateway := new(v1alpha3.Gateway)
				err = yaml.Unmarshal(req.Service, gateway)
				if err != nil {
					utils.Error.Println(fmt.Sprintf("Error while decoding YAML object. Err was: %s", err))
					return nil, errors.New(fmt.Sprintf("Error while decoding YAML object. Err was: %s", err))
				}
				np, err := convertToCPGateWayStruct(gateway)
				if err != nil {
					return nil, err
				}
				bytesData, err := json.Marshal(np)
				if err != nil {
					return nil, err
				}
				serviceResp.Service = bytesData
				return serviceResp, nil
			} else if yamlService.Kind == "ServiceEntry" {
				se := new(v1alpha3.ServiceEntry)
				err = yaml.Unmarshal(req.Service, &se)
				if err != nil {
					utils.Error.Println(fmt.Sprintf("Error while decoding YAML object. Err was: %s", err))
					return nil, errors.New(fmt.Sprintf("Error while decoding YAML object. Err was: %s", err))
				}
				/*	np,err := convertToCPGateWayStruct(o)
					if err!=nil{
						return nil,err
					}
					bytesData, err:= json.Marshal(np)
					if err!=nil{
						return nil,err
					}
					serviceResp.Service=bytesData
					return serviceResp,nil*/
			} else if yamlService.Kind == "DestinationRule" {
				dr := new(v1alpha3.DestinationRule)
				err = yaml.Unmarshal(req.Service, dr)
				if err != nil {
					utils.Error.Println(fmt.Sprintf("Error while decoding YAML object. Err was: %s", err))
					return nil, errors.New(fmt.Sprintf("Error while decoding YAML object. Err was: %s", err))
				}
				np, err := convertToCPDRStruct(dr)
				if err != nil {
					return nil, err
				}
				bytesData, err := json.Marshal(np)
				if err != nil {
					return nil, err
				}
				serviceResp.Service = bytesData
				return serviceResp, nil
			} else if yamlService.Kind == "VirtualService" {
				vs := new(v1alpha3.VirtualService)
				err = yaml.Unmarshal(req.Service, vs)
				if err != nil {
					utils.Error.Println(fmt.Sprintf("Error while decoding YAML object. Err was: %s", err))
					return nil, errors.New(fmt.Sprintf("Error while decoding YAML object. Err was: %s", err))
				}
				np, err := convertToCPVSStruct(vs)
				if err != nil {
					return nil, err
				}
				bytesData, err := json.Marshal(np)
				if err != nil {
					return nil, err
				}
				serviceResp.Service = bytesData
				return serviceResp, nil
			}

		}
		utils.Error.Println(fmt.Sprintf("Error while decoding YAML object. Err was: %s", err))
		return nil, errors.New(fmt.Sprintf("Error while decoding YAML object. Err was: %s", err))
	}
	switch o := obj.(type) {
	case *net.NetworkPolicy:
		np, err := convertToCPNetwokPolicy(o)
		if err != nil {
			return nil, err
		}
		bytesData, err := json.Marshal(np)
		if err != nil {
			return nil, err
		}
		serviceResp.Service = bytesData
		return serviceResp, nil
	case *storage.StorageClass:
		sc, err := convertToCPStorageClass(o)
		if err != nil {
			return nil, err
		}
		bytesData, err := json.Marshal(sc)
		if err != nil {
			return nil, err
		}
		serviceResp.Service = bytesData
		return serviceResp, nil
	case *v1.PersistentVolume:
		pv, err := convertToCPPersistentVolume(o)
		if err != nil {
			return nil, err
		}
		bytesData, err := json.Marshal(pv)
		if err != nil {
			return nil, err
		}
		serviceResp.Service = bytesData
		return serviceResp, nil

	case *v1.PersistentVolumeClaim:
		pvc, err := convertToCPPersistentVolumeClaim(o)
		if err != nil {
			return nil, err
		}
		bytesData, err := json.Marshal(pvc)
		if err != nil {
			return nil, err
		}
		serviceResp.Service = bytesData
		return serviceResp, nil
	case *apps.Deployment:
		pvc, err := convertToCPDeployment(o)
		if err != nil {
			return nil, err
		}
		bytesData, err := json.Marshal(pvc)
		if err != nil {
			return nil, err
		}
		serviceResp.Service = bytesData
		return serviceResp, nil
	case *extensions.Deployment:
		pvc, err := convertToCPDeployment(o)
		if err != nil {
			return nil, err
		}
		bytesData, err := json.Marshal(pvc)
		if err != nil {
			return nil, err
		}
		serviceResp.Service = bytesData
		return serviceResp, nil
	case *v1beta1.Role:
	case *v1beta1.RoleBinding:
	case *v1beta1.ClusterRole:
	case *v1beta1.ClusterRoleBinding:
	case *v1.ServiceAccount:
	default:
		return nil, errors.New("object is not in our scope")
	}
	return serviceResp, nil

}

func convertToCPNetwokPolicy(np *net.NetworkPolicy) (*types.NetworkPolicyService, error) {
	networkPolicy := new(types.NetworkPolicyService)
	networkPolicy.Name = np.Name
	if np.Namespace == "" {
		networkPolicy.Namespace = "default"
	} else {
		networkPolicy.Namespace = np.Namespace
	}
	networkPolicy.ServiceType = "K8s"
	networkPolicy.ServiceSubType = "networkPolicy"
	networkPolicy.ServiceAttributes = new(types.NetworkPolicyServiceAttribute)
	networkPolicy.ServiceAttributes.PodSelector = getCPLabelSelector(&np.Spec.PodSelector)
	for _, each := range np.Spec.Ingress {
		temp := types.IngressRule{}
		for _, ePort := range each.Ports {
			tp := types.NetworkPolicyPort{}
			tp.Protocol = (*types.Protocol)(ePort.Protocol)
			if ePort.Port.Type == intstr.Int {
				tp.Port.PortNumber = ePort.Port.IntVal
			} else {
				tp.Port.PortName = ePort.Port.StrVal
			}
			temp.Ports = append(temp.Ports, tp)
		}
		for _, from := range each.From {
			fm := types.NetworkPolicyPeer{}
			fm.PodSelector = getCPLabelSelector(from.PodSelector)
			fm.NamespaceSelector = getCPLabelSelector(from.NamespaceSelector)
			if from.IPBlock != nil {
				fm.IPBlock = new(types.IPBlock)
				fm.IPBlock.CIDR = from.IPBlock.CIDR
				for _, cidrExcept := range from.IPBlock.Except {
					fm.IPBlock.Except = append(fm.IPBlock.Except, cidrExcept)
				}

			}
			temp.From = append(temp.From, fm)
		}

		networkPolicy.ServiceAttributes.Ingress = append(networkPolicy.ServiceAttributes.Ingress, temp)
	}

	//for egress

	for _, each := range np.Spec.Egress {
		temp := types.EgressRule{}
		for _, ePort := range each.Ports {
			tp := types.NetworkPolicyPort{}
			tp.Protocol = (*types.Protocol)(ePort.Protocol)
			if ePort.Port.Type == intstr.Int {
				tp.Port.PortNumber = ePort.Port.IntVal
			} else {
				tp.Port.PortName = ePort.Port.StrVal
			}
			temp.Ports = append(temp.Ports, tp)
		}
		for _, from := range each.To {
			fm := types.NetworkPolicyPeer{}
			fm.PodSelector = getCPLabelSelector(from.PodSelector)
			fm.NamespaceSelector = getCPLabelSelector(from.NamespaceSelector)
			if from.IPBlock != nil {
				fm.IPBlock = new(types.IPBlock)
				fm.IPBlock.CIDR = from.IPBlock.CIDR
				for _, cidrExcept := range from.IPBlock.Except {
					fm.IPBlock.Except = append(fm.IPBlock.Except, cidrExcept)
				}
			}
			temp.To = append(temp.To, fm)
		}

		networkPolicy.ServiceAttributes.Egress = append(networkPolicy.ServiceAttributes.Egress, temp)
	}

	return networkPolicy, nil

}

func getCPLabelSelector(selector *metav1.LabelSelector) *types.LabelSelectorObj {
	if selector == nil {
		return nil
	}
	ls := new(types.LabelSelectorObj)
	ls.MatchLabels = selector.MatchLabels
	for _, each := range selector.MatchExpressions {
		temp := types.LabelSelectorRequirement{}
		temp.Key = each.Key
		temp.Operator = types.LabelSelectorOperator(each.Operator)
		for _, value := range each.Values {
			temp.Values = append(temp.Values, value)
		}
		ls.MatchExpressions = append(ls.MatchExpressions, temp)
	}
	if len(ls.MatchLabels) == 0 && len(ls.MatchExpressions) == 0 {
		return nil
	}
	return ls
}

func convertToCPDeployment(deploy interface{}) (*types.DeploymentService, error) {
	byteData, _ := json.Marshal(deploy)
	service := apps.Deployment{}
	json.Unmarshal(byteData, &service)

	deployment := new(types.DeploymentService)
	deployment.ServiceAttributes = new(types.DeploymentServiceAttribute)

	if service.Name == "" {
		return nil, errors.New("Service name not found")
	} else {
		deployment.Name = service.Name
	}

	//if service.Namespace == "" {
	//	deployment.Namespace = "default"
	//} else {
	//	deployment.Namespace = service.Namespace
	//}

	deployment.ServiceType = "k8s"
	deployment.ServiceSubType = "Deployment"
	deployment.Version = service.Labels["version"]

	if service.Spec.Replicas != nil {
		deployment.ServiceAttributes.Replicas = new(types.Replicas)
		deployment.ServiceAttributes.Replicas.Value = *service.Spec.Replicas
	}

	deployment.ServiceAttributes.Labels = make(map[string]string)
	deployment.ServiceAttributes.Labels = service.Spec.Template.Labels
	deployment.ServiceAttributes.LabelSelector = new(types.LabelSelectorObj)
	deployment.ServiceAttributes.LabelSelector.MatchLabels = make(map[string]string)
	//deployment.ServiceAttributes.LabelSelector.MatchLabels["version"] = service.Labels["version"]
	//deployment.ServiceAttributes.LabelSelector.MatchLabels["name"] = service.Labels["name"]
	deployment.ServiceAttributes.LabelSelector.MatchLabels = service.Spec.Selector.MatchLabels

	deployment.ServiceAttributes.Annotations = make(map[string]string)
	deployment.ServiceAttributes.Annotations = service.Spec.Template.Annotations
	deployment.ServiceAttributes.NodeSelector = make(map[string]string)
	deployment.ServiceAttributes.NodeSelector = service.Spec.Template.Spec.NodeSelector

	if service.Spec.Strategy.Type != "" {
		if service.Spec.Strategy.Type == apps.RecreateDeploymentStrategyType {
			deployment.ServiceAttributes.Strategy.Type = types.RecreateDeploymentStrategyType
		} else if service.Spec.Strategy.Type == apps.RollingUpdateDeploymentStrategyType {
			deployment.ServiceAttributes.Strategy.Type = types.RollingUpdateDeploymentStrategyType
			if service.Spec.Strategy.RollingUpdate != nil {
				deployment.ServiceAttributes.Strategy.RollingUpdate = new(types.RollingUpdateDeployment)
				deployment.ServiceAttributes.Strategy.RollingUpdate.MaxSurge = new(intstr.IntOrString)
				deployment.ServiceAttributes.Strategy.RollingUpdate.MaxUnavailable = new(intstr.IntOrString)
				deployment.ServiceAttributes.Strategy.RollingUpdate.MaxSurge.IntVal = service.Spec.Strategy.RollingUpdate.MaxUnavailable.IntVal
				deployment.ServiceAttributes.Strategy.RollingUpdate.MaxSurge.StrVal = service.Spec.Strategy.RollingUpdate.MaxUnavailable.StrVal
				deployment.ServiceAttributes.Strategy.RollingUpdate.MaxUnavailable.IntVal = service.Spec.Strategy.RollingUpdate.MaxUnavailable.IntVal
				deployment.ServiceAttributes.Strategy.RollingUpdate.MaxUnavailable.StrVal = service.Spec.Strategy.RollingUpdate.MaxUnavailable.StrVal
			}
		}

	}

	var volumeMountNames1 = make(map[string]bool)
	if containers, vm, err := getCPContainers(service.Spec.Template.Spec.Containers); err == nil {
		if len(containers) > 0 {
			deployment.ServiceAttributes.Container = containers
			volumeMountNames1 = vm
		} else {
			return nil, errors.New("no containers exist")
		}

	} else {
		return nil, err
	}

	if containersList, volumeMounts, err := getCPContainers(service.Spec.Template.Spec.InitContainers); err == nil {
		if len(containersList) > 0 {
			deployment.ServiceAttributes.InitContainer = containersList
		}
		for k, v := range volumeMounts {
			volumeMountNames1[k] = v
		}

	} else {
		return nil, err
	}

	if vols, err := getCPVolumes(service.Spec.Template.Spec.Volumes, volumeMountNames1); err == nil {
		if len(vols) > 0 {
			deployment.ServiceAttributes.Volumes = vols
		}

	} else {
		return nil, err
	}

	if service.Spec.Template.Spec.Affinity != nil {
		if affinity, err := getCPAffinity(service.Spec.Template.Spec.Affinity); err == nil {
			deployment.ServiceAttributes.Affinity = affinity
		} else {
			return nil, err
		}
	}
	return deployment, nil
}

func convertToCPPersistentVolumeClaim(pvc *v1.PersistentVolumeClaim) (*types.PersistentVolumeClaimService, error) {
	persistentVolume := new(types.PersistentVolumeClaimService)
	persistentVolume.Name = pvc.Name
	persistentVolume.ServiceType = "k8s"
	persistentVolume.ServiceSubType = "PVC"
	persistentVolume.ServiceAttributes = new(types.PersistentVolumeClaimServiceAttribute)
	if pvc.Spec.StorageClassName != nil {
		persistentVolume.ServiceAttributes.StorageClassName = *pvc.Spec.StorageClassName
	}
	if pvc.Spec.VolumeMode != nil {
		persistentVolume.ServiceAttributes.VolumeMode = (*types.PersistentVolumeMode)(pvc.Spec.VolumeMode)
	}
	persistentVolume.ServiceAttributes.LabelSelector = getCPLabelSelector(pvc.Spec.Selector)
	persistentVolume.ServiceAttributes.VolumeName = pvc.Spec.VolumeName
	if len(pvc.Spec.Resources.Requests) > 0 {
		qu := pvc.Spec.Resources.Requests[v1.ResourceStorage]
		persistentVolume.ServiceAttributes.Request = qu.String()

	}

	if len(pvc.Spec.Resources.Limits) > 0 {
		qu := pvc.Spec.Resources.Limits[v1.ResourceStorage]
		persistentVolume.ServiceAttributes.Limit = qu.String()

	}
	if pvc.Spec.DataSource != nil {
		persistentVolume.ServiceAttributes.DataSource = new(types.TypedLocalObjectReference)
		persistentVolume.ServiceAttributes.DataSource.Name = pvc.Spec.DataSource.Name
		persistentVolume.ServiceAttributes.DataSource.Kind = pvc.Spec.DataSource.Kind
		if pvc.Spec.DataSource.APIGroup != nil {
			persistentVolume.ServiceAttributes.DataSource.APIGroup = pvc.Spec.DataSource.APIGroup
		}

	}
	for _, each := range pvc.Spec.AccessModes {
		var am types.AccessMode
		if each == v1.ReadWriteOnce {
			am = types.AccessModeReadWriteOnce
		} else if each == v1.ReadOnlyMany {
			am = types.AccessModeReadOnlyMany
		} else if each == v1.ReadWriteMany {
			am = types.AccessModeReadWriteMany
		} else {
			continue
		}

		persistentVolume.ServiceAttributes.AccessMode = append(persistentVolume.ServiceAttributes.AccessMode, am)
	}
	return persistentVolume, nil
}

func convertToCPPersistentVolume(pv *v1.PersistentVolume) (*types.PersistentVolumeService, error) {
	persistentVolume := new(types.PersistentVolumeService)
	persistentVolume.Name = pv.Name
	persistentVolume.ServiceType = "k8s"
	persistentVolume.ServiceSubType = "PV"
	persistentVolume.ServiceAttributes = new(types.PersistentVolumeServiceAttribute)
	persistentVolume.ServiceAttributes.ReclaimPolicy = types.ReclaimPolicy(pv.Spec.PersistentVolumeReclaimPolicy)
	qu := pv.Spec.Capacity[v1.ResourceStorage]
	persistentVolume.ServiceAttributes.Capcity = qu.String()
	if len(pv.Labels) > 0 {
		persistentVolume.ServiceAttributes.Labels = make(map[string]string)
	}
	for k, v := range pv.Labels {
		persistentVolume.ServiceAttributes.Labels[k] = v
	}
	persistentVolume.ServiceAttributes.StorageClassName = pv.Spec.StorageClassName
	for _, each := range pv.Spec.MountOptions {
		persistentVolume.ServiceAttributes.MountOptions = append(persistentVolume.ServiceAttributes.MountOptions, each)
	}
	if pv.Spec.VolumeMode != nil {
		persistentVolume.ServiceAttributes.VolumeMode = (*types.PersistentVolumeMode)(pv.Spec.VolumeMode)
	}

	if pv.Spec.NodeAffinity != nil {
		persistentVolume.ServiceAttributes.NodeAffinity = new(types.VolumeNodeAffinity)
		if ns, err := getCPNodeSelector(pv.Spec.NodeAffinity.Required); err != nil {
			return nil, err
		} else {
			persistentVolume.ServiceAttributes.NodeAffinity.Required = *ns
		}

	}

	for _, each := range pv.Spec.AccessModes {
		var am types.AccessMode
		if each == v1.ReadWriteOnce {
			am = types.AccessModeReadWriteOnce
		} else if each == v1.ReadOnlyMany {
			am = types.AccessModeReadOnlyMany
		} else if each == v1.ReadWriteMany {
			am = types.AccessModeReadWriteMany
		} else {
			continue
		}

		persistentVolume.ServiceAttributes.AccessMode = append(persistentVolume.ServiceAttributes.AccessMode, am)
	}
	if pv.Spec.PersistentVolumeSource.AWSElasticBlockStore != nil {
		persistentVolume.ServiceAttributes.PersistentVolumeSource.AWSEBS.VolumeId = pv.Spec.AWSElasticBlockStore.VolumeID
		persistentVolume.ServiceAttributes.PersistentVolumeSource.AWSEBS.ReadOnly = pv.Spec.AWSElasticBlockStore.ReadOnly
		persistentVolume.ServiceAttributes.PersistentVolumeSource.AWSEBS.Filesystem = pv.Spec.AWSElasticBlockStore.FSType
		persistentVolume.ServiceAttributes.PersistentVolumeSource.AWSEBS.Partition = int(pv.Spec.AWSElasticBlockStore.Partition)
	} else if pv.Spec.PersistentVolumeSource.GCEPersistentDisk != nil {
		persistentVolume.ServiceAttributes.PersistentVolumeSource.GCPPD.PdName = pv.Spec.GCEPersistentDisk.PDName
		persistentVolume.ServiceAttributes.PersistentVolumeSource.GCPPD.ReadOnly = pv.Spec.GCEPersistentDisk.ReadOnly
		persistentVolume.ServiceAttributes.PersistentVolumeSource.GCPPD.Filesystem = pv.Spec.GCEPersistentDisk.FSType
		persistentVolume.ServiceAttributes.PersistentVolumeSource.GCPPD.Partition = int(pv.Spec.GCEPersistentDisk.Partition)
	} else if pv.Spec.PersistentVolumeSource.AzureFile != nil {
		persistentVolume.ServiceAttributes.PersistentVolumeSource.AzureFile.ReadOnly = pv.Spec.AzureFile.ReadOnly
		persistentVolume.ServiceAttributes.PersistentVolumeSource.AzureFile.ShareName = pv.Spec.AzureFile.ShareName
		persistentVolume.ServiceAttributes.PersistentVolumeSource.AzureFile.SecretName = pv.Spec.AzureFile.SecretName
		persistentVolume.ServiceAttributes.PersistentVolumeSource.AzureFile.SecretNamespace = *pv.Spec.AzureFile.SecretNamespace
	} else if pv.Spec.PersistentVolumeSource.AzureDisk != nil {
		if pv.Spec.AzureDisk.ReadOnly != nil {
			persistentVolume.ServiceAttributes.PersistentVolumeSource.AzureDisk.ReadOnly = *pv.Spec.AzureDisk.ReadOnly
		}
		if pv.Spec.AzureDisk.FSType != nil {
			persistentVolume.ServiceAttributes.PersistentVolumeSource.AzureDisk.Filesystem = *pv.Spec.AzureDisk.FSType
		}
		persistentVolume.ServiceAttributes.PersistentVolumeSource.AzureDisk.DiskURI = pv.Spec.AzureDisk.DataDiskURI
		persistentVolume.ServiceAttributes.PersistentVolumeSource.AzureDisk.DiskName = pv.Spec.AzureDisk.DiskName
		if pv.Spec.AzureDisk.CachingMode != nil {
			if *pv.Spec.AzureDisk.CachingMode == v1.AzureDataDiskCachingNone {
				persistentVolume.ServiceAttributes.PersistentVolumeSource.AzureDisk.CachingMode = types.AzureDataDiskCachingNone
			} else if *pv.Spec.AzureDisk.CachingMode == v1.AzureDataDiskCachingReadOnly {
				persistentVolume.ServiceAttributes.PersistentVolumeSource.AzureDisk.CachingMode = types.AzureDataDiskCachingReadOnly
			} else if *pv.Spec.AzureDisk.CachingMode == v1.AzureDataDiskCachingReadWrite {
				persistentVolume.ServiceAttributes.PersistentVolumeSource.AzureDisk.CachingMode = types.AzureDataDiskCachingReadWrite
			}

		}
		if pv.Spec.AzureDisk.Kind != nil {
			if *pv.Spec.AzureDisk.Kind == v1.AzureDedicatedBlobDisk {
				persistentVolume.ServiceAttributes.PersistentVolumeSource.AzureDisk.Kind = types.AzureDedicatedBlobDisk
			} else if *pv.Spec.AzureDisk.Kind == v1.AzureSharedBlobDisk {
				persistentVolume.ServiceAttributes.PersistentVolumeSource.AzureDisk.Kind = types.AzureSharedBlobDisk
			} else if *pv.Spec.AzureDisk.Kind == v1.AzureManagedDisk {
				persistentVolume.ServiceAttributes.PersistentVolumeSource.AzureDisk.Kind = types.AzureManagedDisk
			}
		}

	}

	return persistentVolume, nil
}

func convertToCPStorageClass(sc *storage.StorageClass) (*types.StorageClassService, error) {
	storageClass := new(types.StorageClassService)
	storageClass.Name = sc.Name
	storageClass.ServiceType = "k8s"
	storageClass.ServiceSubType = "SC"
	storageClass.ServiceAttributes = new(types.StorageClassServiceAttribute)
	if sc.ReclaimPolicy != nil {
		storageClass.ServiceAttributes.ReclaimPolicy = types.ReclaimPolicy(*sc.ReclaimPolicy)
	}
	if sc.AllowVolumeExpansion != nil {
		if *sc.AllowVolumeExpansion {
			storageClass.ServiceAttributes.AllowVolumeExpansion = "true"
		}
		if !*sc.AllowVolumeExpansion {
			storageClass.ServiceAttributes.AllowVolumeExpansion = "false"
		}
	}

	for _, each := range sc.MountOptions {
		storageClass.ServiceAttributes.MountOptions = append(storageClass.ServiceAttributes.MountOptions, each)
	}

	for _, each := range sc.AllowedTopologies {
		aT := types.TopologySelectorTerm{}
		for _, each2 := range each.MatchLabelExpressions {
			tr := types.TopologySelectorLabelRequirement{}
			tr.Key = each2.Key
			for _, value := range each2.Values {
				tr.Values = append(tr.Values, value)
			}
			aT.MatchLabelExpressions = append(aT.MatchLabelExpressions, tr)

		}
		storageClass.ServiceAttributes.AllowedTopologies = append(storageClass.ServiceAttributes.AllowedTopologies, aT)
	}

	if sc.Provisioner == "kubernetes.io/aws-ebs" {
		storageClass.ServiceAttributes.SCParameters.AwsEbsScParm = make(map[string]string)
		storageClass.ServiceAttributes.SCParameters.AwsEbsScParm["type"] = sc.Parameters["type"]
		if storageClass.ServiceAttributes.SCParameters.AwsEbsScParm["type"] == "io1" {
			storageClass.ServiceAttributes.SCParameters.AwsEbsScParm["iopsPerGB"] = sc.Parameters["iopsPerGB"]
		}
		if sc.Parameters["encrypted"] != "" {
			storageClass.ServiceAttributes.SCParameters.AwsEbsScParm["encrypted"] = sc.Parameters["encrypted"]
		}
		if sc.Parameters["kmsKeyId"] != "" {
			storageClass.ServiceAttributes.SCParameters.AwsEbsScParm["kmsKeyId"] = sc.Parameters["kmsKeyId"]
		}
		if sc.Parameters["zone"] != "" {
			storageClass.ServiceAttributes.SCParameters.AwsEbsScParm["zone"] = sc.Parameters["zone"]
		} else if sc.Parameters["zones"] != "" {
			storageClass.ServiceAttributes.SCParameters.AwsEbsScParm["zones"] = sc.Parameters["zones"]
		}
	} else if sc.Provisioner == "kubernetes.io/gce-pd" {
		storageClass.ServiceAttributes.SCParameters.GcpPdScParm = make(map[string]string)
		storageClass.ServiceAttributes.SCParameters.GcpPdScParm["type"] = sc.Parameters["type"]
		if sc.Parameters["replication-type"] != "" {
			storageClass.ServiceAttributes.SCParameters.GcpPdScParm["replication-type"] = sc.Parameters["replication-type"]
		}
		if sc.Parameters["zone"] != "" {
			storageClass.ServiceAttributes.SCParameters.GcpPdScParm["zone"] = sc.Parameters["zone"]
		} else if sc.Parameters["zones"] != "" {
			storageClass.ServiceAttributes.SCParameters.GcpPdScParm["zones"] = sc.Parameters["zones"]
		}
	} else if sc.Provisioner == "kubernetes.io/azure-disk" {
		storageClass.ServiceAttributes.SCParameters.AzureDiskScParm = make(map[string]string)
		if sc.Parameters["skuName"] != "" {
			storageClass.ServiceAttributes.SCParameters.AzureDiskScParm["skuName"] = sc.Parameters["skuName"]
		}
		if sc.Parameters["location"] != "" {
			storageClass.ServiceAttributes.SCParameters.AzureDiskScParm["location"] = sc.Parameters["location"]
		}
		if sc.Parameters["storageAccount"] != "" {
			storageClass.ServiceAttributes.SCParameters.AzureDiskScParm["storageAccount"] = sc.Parameters["storageAccount"]
		}
	} else if sc.Provisioner == "kubernetes.io/azure-file" {

		storageClass.ServiceAttributes.SCParameters.AzureFileScParm = make(map[string]string)
		if sc.Parameters["skuName"] != "" {
			storageClass.ServiceAttributes.SCParameters.AzureFileScParm["skuName"] = sc.Parameters["skuName"]
		}
		if sc.Parameters["location"] != "" {
			storageClass.ServiceAttributes.SCParameters.AzureFileScParm["location"] = sc.Parameters["location"]
		}
		if sc.Parameters["storageAccount"] != "" {
			storageClass.ServiceAttributes.SCParameters.AzureFileScParm["storageAccount"] = sc.Parameters["storageAccount"]
		}
		if sc.Parameters["secretNamespace"] != "" {
			storageClass.ServiceAttributes.SCParameters.AzureFileScParm["secretNamespace"] = sc.Parameters["secretNamespace"]
		}
		if sc.Parameters["secretName"] != "" {
			storageClass.ServiceAttributes.SCParameters.AzureFileScParm["secretName"] = sc.Parameters["secretName"]
		}
		if sc.Parameters["readOnly"] != "" {
			storageClass.ServiceAttributes.SCParameters.AzureFileScParm["readOnly"] = sc.Parameters["readOnly"]
		}
	}
	if sc.VolumeBindingMode != nil {
		storageClass.ServiceAttributes.BindingMod = types.VolumeBindingMode(*sc.VolumeBindingMode)
	}

	return storageClass, nil
}

func convertToCPGateWayStruct(gw *v1alpha3.Gateway) (*types.GatewayService, error) {
	return nil, nil
}

func convertToCPVSStruct(gw *v1alpha3.VirtualService) (*types.VirtualService, error) {
	return nil, nil
}

func convertToCPDRStruct(gw *v1alpha3.DestinationRule) (*types.DestinationRules, error) {
	return nil, nil
}

func getCPNodeSelector(nodeSelector *v1.NodeSelector) (*types.NodeSelector, error) {
	var temp *types.NodeSelector
	if nodeSelector != nil {
		temp = new(types.NodeSelector)
		var nodeSelectorTerms []types.NodeSelectorTerm
		for _, nodeSelectorTerm := range nodeSelector.NodeSelectorTerms {
			var tempMatchExpressions []types.NodeSelectorRequirement
			var tempMatchFields []types.NodeSelectorRequirement
			tempNodeSelectorTerm := types.NodeSelectorTerm{}
			for _, matchExpression := range nodeSelectorTerm.MatchExpressions {
				tempMatchExpression := types.NodeSelectorRequirement{}
				tempMatchExpression.Key = matchExpression.Key
				tempMatchExpression.Values = matchExpression.Values
				if matchExpression.Operator == v1.NodeSelectorOpIn {
					tempMatchExpression.Operator = types.NodeSelectorOpIn
				} else if matchExpression.Operator == v1.NodeSelectorOpNotIn {
					tempMatchExpression.Operator = types.NodeSelectorOpNotIn
				} else if matchExpression.Operator == v1.NodeSelectorOpExists {
					tempMatchExpression.Operator = types.NodeSelectorOpExists
				} else if matchExpression.Operator == v1.NodeSelectorOpDoesNotExist {
					tempMatchExpression.Operator = types.NodeSelectorOpDoesNotExists
				} else if matchExpression.Operator == v1.NodeSelectorOpGt {
					tempMatchExpression.Operator = types.NodeSelectorOpGt
				} else if matchExpression.Operator == v1.NodeSelectorOpLt {
					tempMatchExpression.Operator = types.NodeSelectorOpLt
				}

				tempMatchExpressions = append(tempMatchExpressions, tempMatchExpression)
			}
			for _, matchField := range nodeSelectorTerm.MatchFields {
				tempMatchField := types.NodeSelectorRequirement{}
				tempMatchField.Key = matchField.Key
				tempMatchField.Values = matchField.Values
				if matchField.Operator == v1.NodeSelectorOpIn {
					tempMatchField.Operator = types.NodeSelectorOpIn
				} else if matchField.Operator == v1.NodeSelectorOpNotIn {
					tempMatchField.Operator = types.NodeSelectorOpNotIn
				} else if matchField.Operator == v1.NodeSelectorOpExists {
					tempMatchField.Operator = types.NodeSelectorOpExists
				} else if matchField.Operator == v1.NodeSelectorOpDoesNotExist {
					tempMatchField.Operator = types.NodeSelectorOpDoesNotExists
				} else if matchField.Operator == v1.NodeSelectorOpGt {
					tempMatchField.Operator = types.NodeSelectorOpGt
				} else if matchField.Operator == v1.NodeSelectorOpLt {
					tempMatchField.Operator = types.NodeSelectorOpLt
				}
				tempMatchFields = append(tempMatchFields, tempMatchField)
			}

			tempNodeSelectorTerm.MatchFields = tempMatchFields
			tempNodeSelectorTerm.MatchExpressions = tempMatchExpressions
			nodeSelectorTerms = append(nodeSelectorTerms, tempNodeSelectorTerm)
		}
		temp.NodeSelectorTerms = nodeSelectorTerms

	}
	return temp, nil

}

func getCPContainers(conts []v1.Container) (map[string]types.ContainerAttribute, map[string]bool, error) {
	volumeMountNames := make(map[string]bool)
	containers := make(map[string]types.ContainerAttribute)

	for _, container := range conts {
		containerTemp := types.ContainerAttribute{}

		if container.ReadinessProbe != nil {
			if rp, err := getCPProbe(container.ReadinessProbe); err == nil {
				containerTemp.ReadinessProbe = rp
			} else {
				return nil, nil, err
			}
		}

		if container.LivenessProbe != nil {
			if lp, err := getCPProbe(container.LivenessProbe); err == nil {
				containerTemp.LivenessProbe = lp
			} else {
				return nil, nil, err
			}
		}

		if err := putCPCommandAndArguments(&containerTemp, container.Command, container.Args); err != nil {
			return nil, nil, err
		}

		if err := putCPResource(&containerTemp, container.Resources.Limits); err != nil {
			return nil, nil, err
		}

		if err := putCPResource(&containerTemp, container.Resources.Requests); err != nil {
			return nil, nil, err
		}

		if container.SecurityContext != nil {
			if context, err := getCPSecurityContext(container.SecurityContext); err == nil {
				containerTemp.SecurityContext = context
			} else {
				return nil, nil, err
			}
		}
		containerTemp.ImageName = container.Image

		var volumeMounts []types.VolumeMount
		for _, volumeMount := range container.VolumeMounts {
			volumeMountNames[volumeMount.Name] = true
			temp := types.VolumeMount{}
			temp.Name = volumeMount.Name
			temp.MountPath = volumeMount.MountPath
			temp.SubPath = volumeMount.SubPath
			temp.SubPathExpr = volumeMount.SubPathExpr
			if volumeMount.MountPropagation != nil {
				if *volumeMount.MountPropagation == v1.MountPropagationNone {
					*temp.MountPropagation = types.MountPropagationNone
				} else if *volumeMount.MountPropagation == v1.MountPropagationBidirectional {
					*temp.MountPropagation = types.MountPropagationBidirectional
				} else if *volumeMount.MountPropagation == v1.MountPropagationHostToContainer {
					*temp.MountPropagation = types.MountPropagationHostToContainer
				}

			}
			volumeMounts = append(volumeMounts, temp)

		}

		ports := make(map[string]types.ContainerPort)
		for _, port := range container.Ports {
			temp := types.ContainerPort{}
			if port.ContainerPort == 0 && port.HostPort != 0 {
				port.ContainerPort = port.HostPort
			}

			if port.ContainerPort > 0 && port.ContainerPort < 65536 {
				temp.ContainerPort = port.ContainerPort
			} else {
				utils.Info.Println("invalid port number")
				continue
			}
			if port.HostPort != 0 {
				if port.HostPort > 0 && port.HostPort < 65536 {
					temp.HostPort = port.HostPort
				} else {
					utils.Info.Println("invalid port number")
					continue
				}

			}
			ports[port.Name] = temp
		}

		environmentVariables := make(map[string]types.EnvironmentVariable)
		for _, envVariable := range container.Env {
			tempEnvVariable := types.EnvironmentVariable{}
			if envVariable.ValueFrom != nil {
				if envVariable.ValueFrom.ConfigMapKeyRef != nil {
					tempEnvVariable.Value = strings.Join([]string{envVariable.ValueFrom.ConfigMapKeyRef.Name, envVariable.ValueFrom.ConfigMapKeyRef.Key}, ";")
					tempEnvVariable.Type = "ConfigMap"
					tempEnvVariable.Dynamic = true
				} else if envVariable.ValueFrom.SecretKeyRef != nil {
					tempEnvVariable.Value = strings.Join([]string{envVariable.ValueFrom.SecretKeyRef.Name, envVariable.ValueFrom.SecretKeyRef.Key}, ";")
					tempEnvVariable.Type = "Secret"
					tempEnvVariable.Dynamic = true
				}
				environmentVariables[tempEnvVariable.Type] = tempEnvVariable
			} else {
				tempEnvVariable.Key = envVariable.Name
				tempEnvVariable.Value = envVariable.Value
				environmentVariables[tempEnvVariable.Key] = tempEnvVariable
			}

		}

		containerTemp.Ports = ports
		containerTemp.EnvironmentVariables = environmentVariables
		containerTemp.VolumeMounts = volumeMounts

		containers[container.Name] = containerTemp
	}
	return containers, volumeMountNames, nil
}

func getCPProbe(prob *v1.Probe) (*types.Probe, error) {
	CpProbe := new(types.Probe)

	CpProbe.FailureThreshold = prob.FailureThreshold
	CpProbe.InitialDelaySeconds = &prob.InitialDelaySeconds
	CpProbe.SuccessThreshold = prob.SuccessThreshold
	CpProbe.PeriodSeconds = prob.PeriodSeconds
	CpProbe.TimeoutSeconds = prob.TimeoutSeconds

	if prob.Handler.Exec != nil {
		CpProbe.Handler = new(types.Handler)
		CpProbe.Handler.Type = "Exec"
		CpProbe.Handler.Exec = new(types.ExecAction)
		for i := 0; i < len(prob.Handler.Exec.Command); i++ {
			CpProbe.Handler.Exec.Command = append(CpProbe.Handler.Exec.Command, prob.Handler.Exec.Command[i])
		}
	} else if prob.HTTPGet != nil {
		CpProbe.Handler = new(types.Handler)
		CpProbe.Handler.Type = "http_get"
		CpProbe.Handler.HTTPGet = new(types.HTTPGetAction)
		if prob.HTTPGet.Port.IntVal > 0 && prob.HTTPGet.Port.IntVal < 65536 {
			if prob.HTTPGet.Host == "" {
				CpProbe.Handler.HTTPGet.Host = nil
			} else {
				CpProbe.Handler.HTTPGet.Host = &prob.HTTPGet.Host
			}
			if prob.HTTPGet.Path == "" {
				CpProbe.Handler.HTTPGet.Path = nil
			} else {
				CpProbe.Handler.HTTPGet.Path = &prob.HTTPGet.Path
			}

			if prob.HTTPGet.Scheme == v1.URISchemeHTTP && prob.HTTPGet.Scheme == v1.URISchemeHTTPS {
				if prob.HTTPGet.Scheme == v1.URISchemeHTTP {
					scheme := types.URISchemeHTTP
					CpProbe.Handler.HTTPGet.Scheme = &scheme
				} else if prob.HTTPGet.Scheme == v1.URISchemeHTTPS {
					scheme := types.URISchemeHTTPS
					CpProbe.Handler.HTTPGet.Scheme = &scheme
				}
			} else if prob.HTTPGet.Scheme == "" {
				CpProbe.Handler.HTTPGet.Scheme = nil
			} else {
				return nil, errors.New("invalid URI scheme")
			}

			for i := 0; i < len(prob.HTTPGet.HTTPHeaders); i++ {
				CpProbe.Handler.HTTPGet.HTTPHeaders[i].Name = &prob.HTTPGet.HTTPHeaders[i].Name
				CpProbe.Handler.HTTPGet.HTTPHeaders[i].Value = &prob.HTTPGet.HTTPHeaders[i].Value
			}
			CpProbe.Handler.HTTPGet.Port = int(prob.HTTPGet.Port.IntVal)
		} else {
			return nil, errors.New("not a valid port number for http_get")
		}

	} else if prob.TCPSocket != nil {
		CpProbe.Handler = new(types.Handler)
		CpProbe.Handler.Type = "tcpSocket"
		CpProbe.Handler.TCPSocket = new(types.TCPSocketAction)
		if prob.TCPSocket.Port.IntVal > 0 && prob.TCPSocket.Port.IntVal < 65536 {
			CpProbe.Handler.TCPSocket.Port = int(prob.TCPSocket.Port.IntVal)
			CpProbe.Handler.TCPSocket.Host = &prob.TCPSocket.Host
		} else {
			return nil, errors.New("not a valid port number for tcp socket")
		}

	} else {
		return nil, errors.New("no handler found")
	}
	return CpProbe, nil

}

func putCPCommandAndArguments(container *types.ContainerAttribute, command, args []string) error {
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

func putCPResource(container *types.ContainerAttribute, limitResources map[v1.ResourceName]resource.Quantity) error {
	temp := make(map[string]string)
	for t, v := range limitResources {
		key := t.String()
		if key == types.ResourceTypeMemory || key == types.ResourceTypeCpu {
			quantity := v.String()
			temp[key] = quantity
		} else {
			return errors.New("Error Found: Invalid Request Resource Provided. Valid: 'cpu','memory'")
		}
	}

	container.LimitResources = temp
	return nil
}

func getCPSecurityContext(securityContext *v1.SecurityContext) (*types.SecurityContextStruct, error) {
	context := new(types.SecurityContextStruct)
	if securityContext.Capabilities != nil {
		context.Capabilities = new(types.Capabilities)
		for i := 0; i < len(securityContext.Capabilities.Add); i++ {
			context.Capabilities.Add[i] = types.Capability(securityContext.Capabilities.Add[i])
		}
		for i := 0; i < len(securityContext.Capabilities.Drop); i++ {
			context.Capabilities.Add[i] = types.Capability(securityContext.Capabilities.Drop[i])
		}
	}
	if securityContext.AllowPrivilegeEscalation != nil {
		context.AllowPrivilegeEscalation = *securityContext.AllowPrivilegeEscalation
	}
	if securityContext.ReadOnlyRootFilesystem != nil {
		context.AllowPrivilegeEscalation = *securityContext.AllowPrivilegeEscalation
	}
	if securityContext.Privileged != nil {
		context.Privileged = *securityContext.Privileged
	}
	if securityContext.ReadOnlyRootFilesystem != nil {
		context.ReadOnlyRootFileSystem = *securityContext.ReadOnlyRootFilesystem
	}

	if securityContext.RunAsNonRoot != nil {

	}
	if securityContext.RunAsUser != nil {
		context.RunAsUser = securityContext.RunAsUser

	}

	if *securityContext.ProcMount == v1.DefaultProcMount {
		context.ProcMount = types.DefaultProcMount
	} else if *securityContext.ProcMount == v1.UnmaskedProcMount {
		context.ProcMount = types.UnmaskedProcMount
	}

	if securityContext.SELinuxOptions != nil {
		context.SELinuxOptions = types.SELinuxOptionsStruct{
			User:  securityContext.SELinuxOptions.User,
			Role:  securityContext.SELinuxOptions.Role,
			Type:  securityContext.SELinuxOptions.Type,
			Level: securityContext.SELinuxOptions.Level,
		}
	}
	return context, nil
}

func getCPAffinity(affinity *v1.Affinity) (*types.Affinity, error) {
	temp := new(types.Affinity)
	if affinity.NodeAffinity != nil {
		na, err := getCPNodeAffinity(affinity.NodeAffinity)
		if err != nil {
			return nil, err
		} else {
			temp.NodeAffinity = na
		}
	}
	if affinity.PodAffinity != nil {
		pa, err := getCPPodAffinity(affinity.PodAffinity)
		if err != nil {
			return nil, err
		} else {
			temp.PodAffinity = pa
		}
	}
	if affinity.PodAntiAffinity != nil {
		paa, err := getCPAntiPodAffinity(affinity.PodAntiAffinity)
		if err != nil {
			return nil, err
		} else {
			temp.PodAntiAffinity = paa
		}
	}
	return temp, nil
}

func getCPNodeAffinity(nodeAffinity *v1.NodeAffinity) (*types.NodeAffinity, error) {
	temp := new(types.NodeAffinity)
	if nodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution != nil {
		if ns, err := getCPNodeSelector(nodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution); err != nil {
			return nil, err
		} else {
			temp.ReqDuringSchedulingIgnDuringExec = ns
		}
	}

	var tempPrefSchedulingTerms []types.PreferredSchedulingTerm
	for _, prefSchedulingTerm := range nodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
		tempPrefSchedulingTerm := types.PreferredSchedulingTerm{}

		tempPrefSchedulingTerm.Weight = prefSchedulingTerm.Weight
		var tempMatchExpressions []types.NodeSelectorRequirement
		var tempMatchFields []types.NodeSelectorRequirement

		for _, matchExpression := range prefSchedulingTerm.Preference.MatchExpressions {
			tempMatchExpression := types.NodeSelectorRequirement{}
			tempMatchExpression.Key = matchExpression.Key
			tempMatchExpression.Values = matchExpression.Values
			switch matchExpression.Operator {
			case v1.NodeSelectorOpIn:
				tempMatchExpression.Operator = types.NodeSelectorOpIn
			case v1.NodeSelectorOpNotIn:
				tempMatchExpression.Operator = types.NodeSelectorOpNotIn
			case v1.NodeSelectorOpExists:
				tempMatchExpression.Operator = types.NodeSelectorOpExists
			case v1.NodeSelectorOpDoesNotExist:
				tempMatchExpression.Operator = types.NodeSelectorOpDoesNotExists
			case v1.NodeSelectorOpLt:
				tempMatchExpression.Operator = types.NodeSelectorOpLt
			case v1.NodeSelectorOpGt:
				tempMatchExpression.Operator = types.NodeSelectorOpGt
			}
			tempMatchExpressions = append(tempMatchExpressions, tempMatchExpression)
		}
		for _, matchField := range prefSchedulingTerm.Preference.MatchFields {
			tempMatchField := types.NodeSelectorRequirement{}
			tempMatchField.Key = matchField.Key
			tempMatchField.Values = matchField.Values
			switch matchField.Operator {
			case v1.NodeSelectorOpIn:
				tempMatchField.Operator = types.NodeSelectorOpIn
			case v1.NodeSelectorOpNotIn:
				tempMatchField.Operator = types.NodeSelectorOpNotIn
			case v1.NodeSelectorOpExists:
				tempMatchField.Operator = types.NodeSelectorOpExists
			case v1.NodeSelectorOpDoesNotExist:
				tempMatchField.Operator = types.NodeSelectorOpDoesNotExists
			case v1.NodeSelectorOpLt:
				tempMatchField.Operator = types.NodeSelectorOpLt
			case v1.NodeSelectorOpGt:
				tempMatchField.Operator = types.NodeSelectorOpGt
			}

			tempMatchFields = append(tempMatchFields, tempMatchField)
		}
		tempPrefSchedulingTerm.Preference.MatchExpressions = tempMatchExpressions
		tempPrefSchedulingTerm.Preference.MatchFields = tempMatchFields

		tempPrefSchedulingTerms = append(tempPrefSchedulingTerms, tempPrefSchedulingTerm)

	}
	return temp, nil
}

func getCPPodAffinity(podAffinity *v1.PodAffinity) (*types.PodAffinity, error) {
	temp := new(types.PodAffinity)
	var tempPodAffinityTerms []types.PodAffinityTerm
	for _, podAffinityTerm := range podAffinity.RequiredDuringSchedulingIgnoredDuringExecution {
		tempPodAffinityTerm := types.PodAffinityTerm{}

		tempPodAffinityTerm.Namespaces = podAffinityTerm.Namespaces
		tempPodAffinityTerm.TopologyKey = podAffinityTerm.TopologyKey
		ls := getCPLabelSelector(podAffinityTerm.LabelSelector)
		tempPodAffinityTerm.LabelSelector = ls

		tempPodAffinityTerms = append(tempPodAffinityTerms, tempPodAffinityTerm)

	}
	temp.ReqDuringSchedulingIgnDuringExec = tempPodAffinityTerms
	var tempWeightedAffinityTerms []types.WeightedPodAffinityTerm
	for _, weighted := range podAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
		tempWeightedAffinityTerm := types.WeightedPodAffinityTerm{}

		tempWeightedAffinityTerm.Weight = weighted.Weight

		tempPodAffinityTerm := types.PodAffinityTerm{}
		tempPodAffinityTerm.Namespaces = weighted.PodAffinityTerm.Namespaces
		tempPodAffinityTerm.TopologyKey = weighted.PodAffinityTerm.TopologyKey
		ls := getCPLabelSelector(weighted.PodAffinityTerm.LabelSelector)
		tempPodAffinityTerm.LabelSelector = ls

		tempWeightedAffinityTerm.PodAffinityTerm = tempPodAffinityTerm

		tempWeightedAffinityTerms = append(tempWeightedAffinityTerms, tempWeightedAffinityTerm)
	}
	temp.PrefDuringIgnDuringExec = tempWeightedAffinityTerms
	return temp, nil

}

func getCPAntiPodAffinity(podAntiAffinity *v1.PodAntiAffinity) (*types.PodAntiAffinity, error) {

	temp := new(types.PodAntiAffinity)
	var tempPodAffinityTerms []types.PodAffinityTerm
	for _, podAffinityTerm := range podAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution {
		tempPodAffinityTerm := types.PodAffinityTerm{}

		tempPodAffinityTerm.Namespaces = podAffinityTerm.Namespaces
		tempPodAffinityTerm.TopologyKey = podAffinityTerm.TopologyKey
		ls := getCPLabelSelector(podAffinityTerm.LabelSelector)
		tempPodAffinityTerm.LabelSelector = ls
		tempPodAffinityTerms = append(tempPodAffinityTerms, tempPodAffinityTerm)

	}
	temp.ReqDuringSchedulingIgnDuringExec = tempPodAffinityTerms
	var tempWeightedAffinityTerms []types.WeightedPodAffinityTerm
	for _, weighted := range podAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
		tempWeightedAffinityTerm := types.WeightedPodAffinityTerm{}
		tempWeightedAffinityTerm.Weight = weighted.Weight
		tempPodAffinityTerm := types.PodAffinityTerm{}
		tempPodAffinityTerm.Namespaces = weighted.PodAffinityTerm.Namespaces
		tempPodAffinityTerm.TopologyKey = weighted.PodAffinityTerm.TopologyKey
		ls := getCPLabelSelector(weighted.PodAffinityTerm.LabelSelector)
		tempPodAffinityTerm.LabelSelector = ls
		tempWeightedAffinityTerm.PodAffinityTerm = tempPodAffinityTerm

		tempWeightedAffinityTerms = append(tempWeightedAffinityTerms, tempWeightedAffinityTerm)
	}
	temp.PrefDuringIgnDuringExec = tempWeightedAffinityTerms
	return temp, nil
}

func getCPVolumes(vols []v1.Volume, volumeMountNames map[string]bool) ([]types.Volume, error) {
	var volumes []types.Volume
	for _, volume := range vols {

		if !volumeMountNames[volume.Name] {
			continue
		}
		volumeMountNames[volume.Name] = false
		tempVolume := types.Volume{}
		tempVolume.Name = volume.Name

		if volume.VolumeSource.Secret != nil {
			tempVolume.Secret = new(types.SecretVolumeSource)
			tempVolume.Secret.SecretName = volume.VolumeSource.Secret.SecretName
			tempVolume.Secret.DefaultMode = volume.VolumeSource.Secret.DefaultMode
			var secretItems []types.KeyToPath
			for _, item := range volume.VolumeSource.Secret.Items {
				secretItem := types.KeyToPath{
					Key:  item.Key,
					Path: item.Path,
					Mode: item.Mode,
				}
				secretItems = append(secretItems, secretItem)
			}
			tempVolume.Secret.Items = secretItems
		}
		if volume.VolumeSource.ConfigMap != nil {
			tempVolume.ConfigMap = new(types.ConfigMapVolumeSource)
			tempVolume.ConfigMap.Name = volume.VolumeSource.ConfigMap.LocalObjectReference.Name

			tempVolume.ConfigMap.DefaultMode = volume.VolumeSource.ConfigMap.DefaultMode
			var configMapItems []types.KeyToPath
			for _, item := range volume.VolumeSource.ConfigMap.Items {
				configMapItem := types.KeyToPath{
					Key:  item.Key,
					Path: item.Path,
					Mode: item.Mode,
				}
				configMapItems = append(configMapItems, configMapItem)
			}
			tempVolume.ConfigMap.Items = configMapItems
		}

		if volume.VolumeSource.AWSElasticBlockStore != nil {
			tempVolume.AWSElasticBlockStore = new(types.AWSElasticBlockStoreVolumeSource)
			tempVolume.AWSElasticBlockStore.ReadOnly = volume.VolumeSource.AWSElasticBlockStore.ReadOnly
			tempVolume.AWSElasticBlockStore.Partition = volume.VolumeSource.AWSElasticBlockStore.Partition
		}

		if volume.VolumeSource.EmptyDir != nil {
			tempVolume.EmptyDir = new(types.EmptyDirVolumeSource)
			//quantity, _ := resource.ParseQuantity(volume.VolumeSource.EmptyDir.SizeLimit)
			tempVolume.EmptyDir.SizeLimit = volume.VolumeSource.EmptyDir.SizeLimit
			if volume.VolumeSource.EmptyDir.Medium == v1.StorageMediumDefault {
				tempVolume.EmptyDir.Medium = types.StorageMediumDefault

			}
			if volume.VolumeSource.EmptyDir.Medium == v1.StorageMediumMemory {
				tempVolume.EmptyDir.Medium = types.StorageMediumMemory
			}

			if volume.VolumeSource.EmptyDir.Medium == v1.StorageMediumHugePages {
				tempVolume.EmptyDir.Medium = types.StorageMediumHugePages
			}

		}

		if volume.VolumeSource.GCEPersistentDisk != nil {
			tempVolume.GCEPersistentDisk = new(types.GCEPersistentDiskVolumeSource)
			tempVolume.GCEPersistentDisk.Partition = volume.VolumeSource.GCEPersistentDisk.Partition
			tempVolume.GCEPersistentDisk.ReadOnly = volume.VolumeSource.GCEPersistentDisk.ReadOnly
			tempVolume.GCEPersistentDisk.PDName = volume.VolumeSource.GCEPersistentDisk.PDName
		}

		if volume.VolumeSource.AzureDisk != nil {
			tempVolume.AzureFile = new(types.AzureFileVolumeSource)
			tempVolume.AzureDisk.ReadOnly = volume.VolumeSource.AzureDisk.ReadOnly
			tempVolume.AzureDisk.DataDiskURI = volume.VolumeSource.AzureDisk.DiskName

			if *volume.VolumeSource.AzureDisk.CachingMode == v1.AzureDataDiskCachingNone {
				temp := types.AzureDataDiskCachingNone
				tempVolume.AzureDisk.CachingMode = &temp
			} else if *volume.VolumeSource.AzureDisk.CachingMode == v1.AzureDataDiskCachingReadWrite {
				temp := types.AzureDataDiskCachingReadWrite
				tempVolume.AzureDisk.CachingMode = &temp
			} else if *volume.VolumeSource.AzureDisk.CachingMode == v1.AzureDataDiskCachingReadOnly {
				temp := types.AzureDataDiskCachingReadOnly
				tempVolume.AzureDisk.CachingMode = &temp
			}

			if *volume.VolumeSource.AzureDisk.Kind == v1.AzureSharedBlobDisk {
				temp := types.AzureSharedBlobDisk
				tempVolume.AzureDisk.Kind = &temp
			} else if *volume.VolumeSource.AzureDisk.Kind == v1.AzureDedicatedBlobDisk {
				temp := types.AzureDedicatedBlobDisk
				tempVolume.AzureDisk.Kind = &temp
			} else if *volume.VolumeSource.AzureDisk.Kind == v1.AzureManagedDisk {
				temp := types.AzureManagedDisk
				tempVolume.AzureDisk.Kind = &temp
			}
		}

		if volume.VolumeSource.AzureFile != nil {
			tempVolume.AzureFile = new(types.AzureFileVolumeSource)
			tempVolume.AzureFile.ReadOnly = volume.VolumeSource.AzureFile.ReadOnly
			tempVolume.AzureFile.SecretName = volume.VolumeSource.AzureFile.SecretName
			tempVolume.AzureFile.ShareName = volume.VolumeSource.AzureFile.ShareName

		}
		if volume.VolumeSource.HostPath != nil {
			tempVolume.HostPath = new(types.HostPathVolumeSource)
			tempVolume.HostPath.Path = volume.VolumeSource.HostPath.Path
			hostPathType := *volume.VolumeSource.HostPath.Type
			hostPathTypeTemp := types.HostPathType(hostPathType)
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
