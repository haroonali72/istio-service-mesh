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
	net "k8s.io/api/networking/v1"
	"k8s.io/api/rbac/v1beta1"
	storage "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes/scheme"
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
		fmt.Println(o)
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
	ls.MatchLabel = selector.MatchLabels
	for _, each := range selector.MatchExpressions {
		temp := types.LabelSelectorRequirement{}
		temp.Key = each.Key
		temp.Operator = types.LabelSelectorOperator(each.Operator)
		for _, value := range each.Values {
			temp.Values = append(temp.Values, value)
		}
		ls.MatchExpression = append(ls.MatchExpression, temp)
	}
	if len(ls.MatchLabel) == 0 && len(ls.MatchExpression) == 0 {
		return nil
	}
	return ls
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
	qu := pvc.Spec.Resources.Requests[v1.ResourceStorage]
	persistentVolume.ServiceAttributes.Request = qu.String()
	qu = pvc.Spec.Resources.Limits[v1.ResourceStorage]
	persistentVolume.ServiceAttributes.Limit = qu.String()
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
