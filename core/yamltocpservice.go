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
	}

	storageClass.ServiceAttributes.BindingMod = types.VolumeBindingMode(*sc.VolumeBindingMode)
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
