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
	v1alpha32 "istio.io/api/networking/v1alpha3"
	"istio.io/client-go/pkg/apis/networking/v1alpha3"
	apps "k8s.io/api/apps/v1"
	autoScalar "k8s.io/api/autoscaling/v2beta2"
	batch "k8s.io/api/batch/v1"
	batchv1 "k8s.io/api/batch/v1beta1"
	v1 "k8s.io/api/core/v1"
	extensions "k8s.io/api/extensions/v1beta1"
	net "k8s.io/api/networking/v1"
	rbac "k8s.io/api/rbac/v1"
	storage "k8s.io/api/storage/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes/scheme"
	"regexp"
	yaml2 "sigs.k8s.io/yaml"
	"strconv"
	"strings"
	"time"
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
				gateway := v1alpha3.Gateway{}
				err = yaml2.Unmarshal(req.Service, &gateway)
				if err != nil {
					utils.Error.Println(fmt.Sprintf("Error while decoding YAML object. Err was: %s", err))
					return nil, errors.New(fmt.Sprintf("Error while decoding YAML object. Err was: %s", err))
				}
				np, err := convertToCPGateway(&gateway)
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
				se := v1alpha3.ServiceEntry{}
				err = yaml2.Unmarshal(req.Service, &se)
				if err != nil {
					utils.Error.Println(fmt.Sprintf("Error while decoding YAML object. Err was: %s", err))
					return nil, errors.New(fmt.Sprintf("Error while decoding YAML object. Err was: %s", err))
				}
				np, err := convertToCPServiceEntry(&se)
				if err != nil {
					return nil, err
				}
				bytesData, err := json.Marshal(np)
				if err != nil {
					return nil, err
				}
				serviceResp.Service = bytesData
				return serviceResp, nil
			} else if yamlService.Kind == "DestinationRule" {
				dr := v1alpha3.DestinationRule{}
				err = yaml2.Unmarshal(req.Service, &dr)
				if err != nil {
					utils.Error.Println(fmt.Sprintf("Error while decoding YAML object. Err was: %s", err))
					return nil, errors.New(fmt.Sprintf("Error while decoding YAML object. Err was: %s", err))
				}
				np, err := convertToCPDestinationRule(&dr)
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
				vs := v1alpha3.VirtualService{}
				err2 := yaml2.Unmarshal(req.Service, &vs)
				//err2:= yaml.Unmarshal(bytes,&vs)
				if err2 != nil {
					utils.Error.Println(fmt.Sprintf("Error while decoding YAML object. Err was: %s", err2))
					return nil, errors.New(fmt.Sprintf("Error while decoding YAML object. Err was: %s", err2))
				}
				np, err := convertToCPVirtualService(&vs)
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
		strData := string(bytesData)
		re := regexp.MustCompile("(?m)[\r\n]+^.*creationTimestamp.*$")
		res := re.ReplaceAllString(strData, "")
		serviceResp.Service = []byte(res)
		return serviceResp, nil
	case *v1.Service:
		pvc, err := convertToCPKubernetesService(o)
		if err != nil {
			return nil, err
		}
		bytesData, err := json.Marshal(pvc)
		if err != nil {
			return nil, err
		}
		serviceResp.Service = bytesData
		return serviceResp, nil
	case *v1.ConfigMap:
		pvc, err := ConvertToCPConfigMap(o)
		if err != nil {
			return nil, err
		}
		bytesData, err := json.Marshal(pvc)
		if err != nil {
			return nil, err
		}
		serviceResp.Service = bytesData
		return serviceResp, nil
	case *v1.Secret:
		pvc, err := ConvertToCPSecret(o)
		if err != nil {
			return nil, err
		}
		bytesData, err := json.Marshal(pvc)
		if err != nil {
			return nil, err
		}
		serviceResp.Service = bytesData
		return serviceResp, nil
	case *autoScalar.HorizontalPodAutoscaler:
		pvc, err := ConvertToCPHPA(o)
		if err != nil {
			return nil, err
		}
		bytesData, err := json.Marshal(pvc)
		if err != nil {
			return nil, err
		}
		serviceResp.Service = bytesData
		return serviceResp, nil
	case *rbac.Role:
		pvc, err := ConvertToCPRole(o)
		if err != nil {
			return nil, err
		}
		bytesData, err := json.Marshal(pvc)
		if err != nil {
			return nil, err
		}
		serviceResp.Service = bytesData
		return serviceResp, nil
	case *rbac.RoleBinding:
		pvc, err := ConvertToCPRoleBinding(o)
		if err != nil {
			return nil, err
		}
		bytesData, err := json.Marshal(pvc)
		if err != nil {
			return nil, err
		}
		serviceResp.Service = bytesData
		return serviceResp, nil
	case *rbac.ClusterRole:
		pvc, err := ConvertToCPClusterRole(o)
		if err != nil {
			return nil, err
		}
		bytesData, err := json.Marshal(pvc)
		if err != nil {
			return nil, err
		}
		serviceResp.Service = bytesData
		return serviceResp, nil
	case *rbac.ClusterRoleBinding:
		pvc, err := ConvertToCPClusterRoleBinding(o)
		if err != nil {
			return nil, err
		}
		bytesData, err := json.Marshal(pvc)
		if err != nil {
			return nil, err
		}
		serviceResp.Service = bytesData
		return serviceResp, nil
	case *v1.ServiceAccount:
		pvc, err := convertToCPServiceAccount(o)
		if err != nil {
			return nil, err
		}
		bytesData, err := json.Marshal(pvc)
		if err != nil {
			return nil, err
		}
		serviceResp.Service = bytesData
		return serviceResp, nil
	case *apps.DaemonSet:
		ds, err := convertToCPDaemonSet(o)
		if err != nil {
			return nil, err
		}
		bytesData, err := json.Marshal(ds)
		if err != nil {
			return nil, err
		}
		serviceResp.Service = bytesData
		return serviceResp, nil
	case *extensions.DaemonSet:
		ds, err := convertToCPDaemonSet(o)
		if err != nil {
			return nil, err
		}
		bytesData, err := json.Marshal(ds)
		if err != nil {
			return nil, err
		}
		serviceResp.Service = bytesData
		return serviceResp, nil
	case *apps.StatefulSet:
		ds, err := convertToCPStatefulSet(o)
		if err != nil {
			return nil, err
		}
		bytesData, err := json.Marshal(ds)
		if err != nil {
			return nil, err
		}
		serviceResp.Service = bytesData
		return serviceResp, nil
	case *batch.Job:
		ds, err := convertToCPJob(o)
		if err != nil {
			return nil, err
		}
		bytesData, err := json.Marshal(ds)
		if err != nil {
			return nil, err
		}
		serviceResp.Service = bytesData
		return serviceResp, nil
	case *batchv1.CronJob:
		ds, err := convertToCPCronJob(o)
		if err != nil {
			return nil, err
		}
		bytesData, err := json.Marshal(ds)
		if err != nil {
			return nil, err
		}
		serviceResp.Service = bytesData
		return serviceResp, nil
	case *v1alpha3.VirtualService:
		vs, err := convertToCPVirtualService(o)
		if err != nil {
			return nil, err
		}
		bytesData, err := json.Marshal(vs)
		if err != nil {
			return nil, err
		}
		serviceResp.Service = bytesData
		return serviceResp, nil

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

	if service.Namespace == "" {
		deployment.Namespace = "default"
	} else {
		deployment.Namespace = service.Namespace
	}

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
			deployment.ServiceAttributes.Strategy = new(types.DeploymentStrategy)
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

	for _, imageSecrets := range service.Spec.Template.Spec.ImagePullSecrets {
		tempImageSecrets := types.LocalObjectReference{Name: imageSecrets.Name}
		deployment.ServiceAttributes.ImagePullSecrets = append(deployment.ServiceAttributes.ImagePullSecrets, tempImageSecrets)
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

func convertToCPDaemonSet(ds interface{}) (*types.DaemonSetService, error) {
	byteData, _ := json.Marshal(ds)
	service := apps.DaemonSet{}
	json.Unmarshal(byteData, &service)
	daemonSet := new(types.DaemonSetService)

	if service.Name == "" {
		return nil, errors.New("Service name not found")
	} else {
		daemonSet.Name = service.Name
	}

	if service.Namespace == "" {
		daemonSet.Namespace = "default"
	} else {
		daemonSet.Namespace = service.Namespace
	}

	daemonSet.ServiceType = "k8s"
	daemonSet.ServiceSubType = "DaemonSet"
	daemonSet.ServiceAttributes = new(types.DaemonSetServiceAttribute)
	daemonSet.ServiceAttributes.Labels = make(map[string]string)
	daemonSet.ServiceAttributes.Labels = service.Spec.Template.Labels
	daemonSet.ServiceAttributes.LabelSelector = new(types.LabelSelectorObj)
	daemonSet.ServiceAttributes.LabelSelector.MatchLabels = make(map[string]string)
	daemonSet.ServiceAttributes.LabelSelector.MatchLabels = service.Spec.Selector.MatchLabels

	daemonSet.ServiceAttributes.Annotations = make(map[string]string)
	daemonSet.ServiceAttributes.Annotations = service.Spec.Template.Annotations
	daemonSet.ServiceAttributes.NodeSelector = make(map[string]string)
	daemonSet.ServiceAttributes.NodeSelector = service.Spec.Template.Spec.NodeSelector

	//daemonSetUpdateStrategy
	if service.Spec.UpdateStrategy.Type != "" {
		daemonSet.ServiceAttributes.UpdateStrategy = new(types.DaemonSetUpdateStrategy)
		if service.Spec.UpdateStrategy.Type == apps.OnDeleteDaemonSetStrategyType {
			daemonSet.ServiceAttributes.UpdateStrategy.Type = types.OnDeleteDaemonSetStrategyType
		} else if service.Spec.UpdateStrategy.Type == apps.RollingUpdateDaemonSetStrategyType {
			daemonSet.ServiceAttributes.UpdateStrategy.Type = types.RollingUpdateDaemonSetStrategyType
			if service.Spec.UpdateStrategy.RollingUpdate != nil {
				daemonSet.ServiceAttributes.UpdateStrategy.RollingUpdate = new(types.RollingUpdateDaemonSet)
				daemonSet.ServiceAttributes.UpdateStrategy.RollingUpdate.MaxUnavailable = new(intstr.IntOrString)
				daemonSet.ServiceAttributes.UpdateStrategy.RollingUpdate.MaxUnavailable = service.Spec.UpdateStrategy.RollingUpdate.MaxUnavailable
			}
		}
	}

	//containers
	var volumeMountNames1 = make(map[string]bool)
	if containers, vm, err := getCPContainers(service.Spec.Template.Spec.Containers); err == nil {
		if len(containers) > 0 {
			daemonSet.ServiceAttributes.Containers = containers
			volumeMountNames1 = vm
		} else {
			return nil, errors.New("no containers exist")
		}

	} else {
		return nil, err
	}

	//init containers
	if containersList, volumeMounts, err := getCPContainers(service.Spec.Template.Spec.InitContainers); err == nil {
		if len(containersList) > 0 {
			daemonSet.ServiceAttributes.InitContainers = containersList
		}
		for k, v := range volumeMounts {
			volumeMountNames1[k] = v
		}

	} else {
		return nil, err
	}

	//volumes
	if vols, err := getCPVolumes(service.Spec.Template.Spec.Volumes, volumeMountNames1); err == nil {
		if len(vols) > 0 {
			daemonSet.ServiceAttributes.Volumes = vols
		}

	} else {
		return nil, err
	}

	//affinity
	if service.Spec.Template.Spec.Affinity != nil {
		if affinity, err := getCPAffinity(service.Spec.Template.Spec.Affinity); err == nil {
			daemonSet.ServiceAttributes.Affinity = affinity
		} else {
			return nil, err
		}
	}

	return daemonSet, nil
}

func convertToCPStatefulSet(sset interface{}) (*types.StatefulSetService, error) {

	byteData, _ := json.Marshal(sset)
	service := apps.StatefulSet{}
	json.Unmarshal(byteData, &service)
	statefulSet := new(types.StatefulSetService)

	statefulSet.Name = service.Name
	statefulSet.ServiceType = "k8s"
	statefulSet.ServiceSubType = "StatefulSet"

	if service.Namespace == "" {
		statefulSet.Namespace = "default"
	} else {
		statefulSet.Namespace = service.Namespace
	}

	statefulSet.ServiceAttributes = new(types.StatefulSetServiceAttribute)
	statefulSet.ServiceAttributes.Labels = make(map[string]string)
	statefulSet.ServiceAttributes.Labels = service.Spec.Template.Labels

	statefulSet.ServiceAttributes.Annotations = make(map[string]string)
	statefulSet.ServiceAttributes.Annotations = service.Spec.Template.Annotations
	statefulSet.ServiceAttributes.LabelSelector = new(types.LabelSelectorObj)
	statefulSet.ServiceAttributes.LabelSelector.MatchLabels = make(map[string]string)
	statefulSet.ServiceAttributes.LabelSelector.MatchLabels = service.Spec.Selector.MatchLabels
	statefulSet.ServiceAttributes.NodeSelector = make(map[string]string)
	statefulSet.ServiceAttributes.NodeSelector = service.Spec.Template.Spec.NodeSelector

	//replicas
	if service.Spec.Replicas != nil {
		statefulSet.ServiceAttributes.Replicas = &types.Replicas{Value: *service.Spec.Replicas}
	}

	if service.Spec.ServiceName != "" {
		statefulSet.ServiceAttributes.ServiceName = service.Spec.ServiceName
	}
	//update strategy
	if service.Spec.UpdateStrategy.Type != "" {
		statefulSet.ServiceAttributes.UpdateStrategy = new(types.StateFulSetUpdateStrategy)
		if service.Spec.UpdateStrategy.Type == apps.OnDeleteStatefulSetStrategyType {
			statefulSet.ServiceAttributes.UpdateStrategy.Type = types.OnDeleteStatefulSetStrategyType
		} else if service.Spec.UpdateStrategy.Type == apps.RollingUpdateStatefulSetStrategyType {
			statefulSet.ServiceAttributes.UpdateStrategy.Type = types.RollingUpdateStatefulSetStrategyType
			if service.Spec.UpdateStrategy.RollingUpdate != nil {
				statefulSet.ServiceAttributes.UpdateStrategy.RollingUpdate = new(types.RollingUpdateStatefulSetStrategy)
				statefulSet.ServiceAttributes.UpdateStrategy.RollingUpdate.Partition = service.Spec.UpdateStrategy.RollingUpdate.Partition
			}
		}
	}
	//containers
	var volumeMountNames1 = make(map[string]bool)
	if containers, vm, err := getCPContainers(service.Spec.Template.Spec.Containers); err == nil {
		if len(containers) > 0 {
			statefulSet.ServiceAttributes.Containers = containers
			volumeMountNames1 = vm
		} else {
			return nil, errors.New("no containers exist")
		}

	} else {
		return nil, err
	}

	//init containers
	if containersList, volumeMounts, err := getCPContainers(service.Spec.Template.Spec.InitContainers); err == nil {
		if len(containersList) > 0 {
			statefulSet.ServiceAttributes.InitContainers = containersList
		}
		for k, v := range volumeMounts {
			volumeMountNames1[k] = v
		}

	} else {
		return nil, err
	}

	//volumes
	if vols, err := getCPVolumes(service.Spec.Template.Spec.Volumes, volumeMountNames1); err == nil {
		if len(vols) > 0 {
			statefulSet.ServiceAttributes.Volumes = vols
		}

	} else {
		return nil, err
	}

	//volumeClaimTemplates
	for _, vc := range service.Spec.VolumeClaimTemplates {
		//tempVC := new(types.PersistentVolumeClaimService)
		if tempVC, error := convertToCPPersistentVolumeClaim(&vc); error == nil {
			statefulSet.ServiceAttributes.VolumeClaimTemplates = append(statefulSet.ServiceAttributes.VolumeClaimTemplates, *tempVC)
		} else {
			return nil, error
		}
	}

	//affinity
	if service.Spec.Template.Spec.Affinity != nil {
		if affinity, err := getCPAffinity(service.Spec.Template.Spec.Affinity); err == nil {
			statefulSet.ServiceAttributes.Affinity = affinity
		} else {
			return nil, err
		}
	}

	return statefulSet, nil

}

func convertToCPJob(job *batch.Job) (*types.JobService, error) {
	cpJob := new(types.JobService)
	cpJob.Name = job.Name
	cpJob.ServiceType = "k8s"
	cpJob.ServiceSubType = "job"
	if job.Namespace == "" {
		cpJob.Namespace = "default"
	} else {
		cpJob.Namespace = job.Namespace
	}

	cpJob.ServiceAttributes.Labels = make(map[string]string)
	cpJob.ServiceAttributes.Labels = job.Spec.Template.Labels

	cpJob.ServiceAttributes.LabelSelector = new(types.LabelSelectorObj)
	cpJob.ServiceAttributes.LabelSelector.MatchLabels = make(map[string]string)
	cpJob.ServiceAttributes.LabelSelector.MatchLabels = job.Spec.Selector.MatchLabels
	cpJob.ServiceAttributes.Annotations = make(map[string]string)
	cpJob.ServiceAttributes.Annotations = job.Spec.Template.Annotations

	cpJob.ServiceAttributes.NodeSelector = make(map[string]string)
	cpJob.ServiceAttributes.NodeSelector = job.Spec.Template.Spec.NodeSelector

	if job.Spec.Parallelism != nil {
		cpJob.ServiceAttributes.Parallelism.Value = *job.Spec.Parallelism
	}
	if job.Spec.Completions != nil {
		cpJob.ServiceAttributes.Completions.Value = *job.Spec.Completions
	}
	if job.Spec.ActiveDeadlineSeconds != nil {
		cpJob.ServiceAttributes.ActiveDeadlineSeconds.Value = *job.Spec.ActiveDeadlineSeconds
	}
	if job.Spec.BackoffLimit != nil {
		cpJob.ServiceAttributes.BackoffLimit.Value = *job.Spec.BackoffLimit
	}
	if job.Spec.ManualSelector != nil {
		cpJob.ServiceAttributes.ManualSelector.Value = *job.Spec.ManualSelector
	}

	//containers
	var volumeMountNames1 = make(map[string]bool)
	if containers, vm, err := getCPContainers(job.Spec.Template.Spec.Containers); err == nil {
		if len(containers) > 0 {
			cpJob.ServiceAttributes.Containers = containers
			volumeMountNames1 = vm
		} else {
			return nil, errors.New("no containers exist")
		}

	} else {
		return nil, err
	}

	//init containers
	if containersList, volumeMounts, err := getCPContainers(job.Spec.Template.Spec.InitContainers); err == nil {
		if len(containersList) > 0 {
			cpJob.ServiceAttributes.InitContainers = containersList
		}
		for k, v := range volumeMounts {
			volumeMountNames1[k] = v
		}

	} else {
		return nil, err
	}

	//volumes
	if vols, err := getCPVolumes(job.Spec.Template.Spec.Volumes, volumeMountNames1); err == nil {
		if len(vols) > 0 {
			cpJob.ServiceAttributes.Volumes = vols
		}

	} else {
		return nil, err
	}

	//affinity
	if job.Spec.Template.Spec.Affinity != nil {
		if affinity, err := getCPAffinity(job.Spec.Template.Spec.Affinity); err == nil {
			cpJob.ServiceAttributes.Affinity = affinity
		} else {
			return nil, err
		}
	}
	return cpJob, nil
}

func convertToCPCronJob(job *batchv1.CronJob) (*types.CronJobService, error) {
	cpJob := new(types.CronJobService)
	cpJob.Name = job.Labels["app"]
	cpJob.Version = job.Labels["version"]
	cpJob.ServiceType = "k8s"
	cpJob.ServiceSubType = "job"

	if job.Namespace == "" {
		cpJob.Namespace = "default"
	} else {
		cpJob.Namespace = job.Namespace
	}

	cpJob.ServiceAttributes = new(types.CronJobServiceAttribute)

	cpJob.ServiceAttributes.Labels = make(map[string]string)
	cpJob.ServiceAttributes.Labels = job.Labels
	cpJob.ServiceAttributes.Annotations = make(map[string]string)
	cpJob.ServiceAttributes.Annotations = job.Annotations

	if jobTemplate, err := getCPJobTemplateSpec(job.Spec.JobTemplate); err != nil {
		if jobTemplate != nil {
			cpJob.ServiceAttributes.JobServiceAttribute = jobTemplate
		}
	} else {
		return nil, err
	}

	if job.Spec.Schedule != "" {
		cpJob.ServiceAttributes.CronJobScheduleString = job.Spec.Schedule
	}
	if job.Spec.StartingDeadlineSeconds != nil {
		cpJob.ServiceAttributes.StartingDeadLineSeconds = &types.StartingDeadlineSeconds{
			Value: *job.Spec.StartingDeadlineSeconds,
		}
	}

	if job.Spec.FailedJobsHistoryLimit != nil {
		cpJob.ServiceAttributes.FailedJobsHistoryLimit = &types.FailedJobsHistoryLimit{Value: *job.Spec.FailedJobsHistoryLimit}
	}
	if job.Spec.SuccessfulJobsHistoryLimit != nil {
		cpJob.ServiceAttributes.SuccessfulJobsHistoryLimit = &types.SuccessfulJobsHistoryLimit{Value: *job.Spec.SuccessfulJobsHistoryLimit}
	}
	if job.Spec.Suspend != nil {
		cpJob.ServiceAttributes.Suspend = &types.Suspend{Value: *job.Spec.Suspend}
	}
	if job.Spec.ConcurrencyPolicy != "" {
		cpJob.ServiceAttributes.ConcurrencyPolicy = new(types.ConcurrencyPolicy)
		if job.Spec.ConcurrencyPolicy == batchv1.AllowConcurrent {
			value := types.ConcurrencyPolicyAllow
			cpJob.ServiceAttributes.ConcurrencyPolicy = &value
		} else if job.Spec.ConcurrencyPolicy == batchv1.ForbidConcurrent {
			value := types.ConcurrencyPolicyForbid
			cpJob.ServiceAttributes.ConcurrencyPolicy = &value
		} else {
			value := types.ConcurrencyPolicyReplace
			cpJob.ServiceAttributes.ConcurrencyPolicy = &value
		}
	}

	return cpJob, nil

}

func getCPJobTemplateSpec(job batchv1.JobTemplateSpec) (*types.JobServiceAttribute, error) {
	jobTemplate := new(types.JobServiceAttribute)
	jobTemplate.Labels = make(map[string]string)
	jobTemplate.Labels = job.Labels

	jobTemplate.Annotations = make(map[string]string)
	jobTemplate.Annotations = job.Spec.Template.Annotations
	jobTemplate.LabelSelector = new(types.LabelSelectorObj)
	jobTemplate.LabelSelector.MatchLabels = make(map[string]string)
	jobTemplate.LabelSelector.MatchLabels = job.Spec.Selector.MatchLabels
	jobTemplate.NodeSelector = make(map[string]string)
	jobTemplate.NodeSelector = job.Spec.Template.Spec.NodeSelector

	var volumeMountNames1 = make(map[string]bool)
	if containers, vm, err := getCPContainers(job.Spec.Template.Spec.Containers); err == nil {
		if len(containers) > 0 {
			jobTemplate.Containers = containers
			volumeMountNames1 = vm
		} else {
			return nil, errors.New("no containers exist")
		}

	} else {
		return nil, err
	}

	//init containers
	if containersList, volumeMounts, err := getCPContainers(job.Spec.Template.Spec.InitContainers); err == nil {
		if len(containersList) > 0 {
			jobTemplate.InitContainers = containersList
		}
		for k, v := range volumeMounts {
			volumeMountNames1[k] = v
		}

	} else {
		return nil, err
	}

	if job.Spec.Template.Spec.Affinity != nil {
		if affinity, err := getCPAffinity(job.Spec.Template.Spec.Affinity); err == nil {
			jobTemplate.Affinity = affinity
		} else {
			return nil, err
		}
	}

	//volumes
	if vols, err := getCPVolumes(job.Spec.Template.Spec.Volumes, volumeMountNames1); err == nil {
		if len(vols) > 0 {
			jobTemplate.Volumes = vols
		}

	} else {
		return nil, err
	}
	return jobTemplate, nil
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

func ConvertToCPSecret(cm *v1.Secret) (*types.Secret, error) {
	var secret = new(types.Secret)
	secret.Name = cm.Name
	secret.Namespace = cm.Namespace
	if vr := cm.Labels["version"]; vr != "" {
		secret.Version = vr
	}
	secret.ServiceType = "k8s"
	secret.ServiceSubType = "secret"
	secret.ServiceAttributes = new(types.SecretServiceAttribute)
	if len(cm.Data) > 0 {
		secret.ServiceAttributes.Data = make(map[string][]byte)
		for key, value := range cm.Data {
			secret.ServiceAttributes.Data[key] = value
		}
	}

	if len(cm.StringData) > 0 {
		secret.ServiceAttributes.StringData = make(map[string]string)
		for key, value := range cm.StringData {
			secret.ServiceAttributes.StringData[key] = value
		}
	}

	secret.ServiceAttributes.SecretType = string(cm.Type)
	return secret, nil
}

func ConvertToCPHPA(hpa *autoScalar.HorizontalPodAutoscaler) (*types.HPA, error) {
	var horizntalPodAutoscalar = new(types.HPA)
	horizntalPodAutoscalar.Name = hpa.Name
	horizntalPodAutoscalar.Namespace = hpa.Namespace
	horizntalPodAutoscalar.ServiceType = "k8s"
	horizntalPodAutoscalar.ServiceSubType = "hpa"
	horizntalPodAutoscalar.ServiceAttributes.MaxReplicas = int(hpa.Spec.MaxReplicas)
	if hpa.Spec.MinReplicas != nil {
		horizntalPodAutoscalar.ServiceAttributes.MinReplicas = int(*hpa.Spec.MinReplicas)
	}
	horizntalPodAutoscalar.ServiceAttributes.CrossObjectVersion.Name = hpa.Spec.ScaleTargetRef.Name
	horizntalPodAutoscalar.ServiceAttributes.CrossObjectVersion.Type = hpa.Spec.ScaleTargetRef.Kind
	horizntalPodAutoscalar.ServiceAttributes.CrossObjectVersion.Version = hpa.Spec.ScaleTargetRef.APIVersion

	var metrics []types.MetricValue
	for _, metric := range hpa.Spec.Metrics {
		cpMetric := types.MetricValue{}
		cpMetric.ResourceKind = string(autoScalar.ResourceMetricSourceType)
		if metric.Resource != nil {
			if metric.Resource.Target.Type == autoScalar.ValueMetricType {
				cpMetric.TargetValueKind = string(autoScalar.ValueMetricType)
				cpMetric.TargetValue = metric.Resource.Target.Value.String()
			} else if metric.Resource.Target.Type == autoScalar.UtilizationMetricType {
				cpMetric.TargetValueKind = string(autoScalar.UtilizationMetricType)
				if metric.Resource.Target.AverageUtilization != nil {
					cpMetric.TargetValue = strconv.Itoa(int(*metric.Resource.Target.AverageUtilization))
				}
			} else if metric.Resource.Target.Type == autoScalar.AverageValueMetricType {
				cpMetric.TargetValueKind = string(autoScalar.AverageValueMetricType)
				cpMetric.TargetValue = metric.Resource.Target.AverageValue.String()
			}

			if metric.Resource.Name == v1.ResourceCPU {
				cpMetric.ResourceKind = string(v1.ResourceCPU)
			} else if metric.Resource.Name == v1.ResourceMemory {
				cpMetric.ResourceKind = string(v1.ResourceMemory)
			} else if metric.Resource.Name == v1.ResourceStorage {
				cpMetric.ResourceKind = string(v1.ResourceStorage)
			}
		}

		metrics = append(metrics, cpMetric)

	}
	horizntalPodAutoscalar.ServiceAttributes.MetricValues = metrics

	return horizntalPodAutoscalar, nil
}

func ConvertToCPRole(k8ROle *rbac.Role) (*types.Role, error) {
	var role = new(types.Role)
	role.Name = k8ROle.Name
	role.Namespace = k8ROle.Namespace
	role.ServiceType = "k8s"
	role.ServiceSubType = "role"
	for _, each := range k8ROle.Rules {
		rolePolicy := types.Rule{}
		for _, apigroup := range each.APIGroups {
			rolePolicy.Api_group = append(rolePolicy.Api_group, apigroup)
		}

		for _, verb := range each.Verbs {
			rolePolicy.Verbs = append(rolePolicy.Verbs, verb)
		}

		for _, resource := range each.Resources {
			rolePolicy.Resources = append(rolePolicy.Resources, resource)
		}
		for _, resource := range each.ResourceNames {
			rolePolicy.ResourceName = append(rolePolicy.ResourceName, resource)
		}
		role.ServiceAttributes.Rules = append(role.ServiceAttributes.Rules, rolePolicy)
	}
	return role, nil
}

func ConvertToCPRoleBinding(k8sRoleBinding *rbac.RoleBinding) (*types.RoleBinding, error) {
	var rb = new(types.RoleBinding)
	rb.Name = k8sRoleBinding.Name
	rb.ServiceType = "k8s"
	rb.ServiceSubType = "cluster_role_binding"
	for _, each := range k8sRoleBinding.Subjects {
		var subject = types.Subject{}
		subject.Name = each.Name
		if each.Kind == "User" || each.Kind == "Group" {
			subject.Kind = each.Kind
		} else if each.Kind == "ServiceAccount" {
			subject.Kind = each.Kind
			subject.Namespace = each.Namespace
		} else {
			return nil, errors.New("invalid subject kind" + each.Name + each.Kind)
		}
		rb.ServiceAttributes.Subjects = append(rb.ServiceAttributes.Subjects, subject)
	}
	rb.ServiceAttributes.RoleRef.Kind = k8sRoleBinding.RoleRef.Kind
	rb.ServiceAttributes.RoleRef.Name = k8sRoleBinding.RoleRef.Name
	return rb, nil
}

func ConvertToCPClusterRoleBinding(k8sClusterRoleBinding *rbac.ClusterRoleBinding) (*types.ClusterRoleBinding, error) {
	var crb = new(types.ClusterRoleBinding)
	crb.Name = k8sClusterRoleBinding.Name
	crb.ServiceType = "k8s"
	crb.ServiceSubType = "cluster_role_binding"
	crb.ServiceAttributes.NameClusterRoleRef = k8sClusterRoleBinding.RoleRef.Name
	for _, each := range k8sClusterRoleBinding.Subjects {
		var subject = types.Subject{}
		subject.Name = each.Name
		if each.Kind == "User" || each.Kind == "Group" {
			subject.Kind = each.Kind
		} else if each.Kind == "ServiceAccount" {
			subject.Kind = each.Kind
			subject.Namespace = each.Namespace
		} else {
			return nil, errors.New("invalid subject kind" + each.Name + each.Kind)
		}
		crb.ServiceAttributes.Subjects = append(crb.ServiceAttributes.Subjects, subject)
	}
	return crb, nil
}

func ConvertToCPClusterRole(k8ROle *rbac.ClusterRole) (*types.ClusterRole, error) {
	var role = new(types.ClusterRole)
	role.Name = k8ROle.Name
	role.ServiceType = "k8s"
	role.ServiceSubType = "cluster_role"
	for _, each := range k8ROle.Rules {
		rolePolicy := types.Rules{}
		for _, apigroup := range each.APIGroups {
			rolePolicy.ApiGroup = append(rolePolicy.ApiGroup, apigroup)
		}

		for _, verb := range each.Verbs {
			rolePolicy.Verbs = append(rolePolicy.Verbs, verb)
		}

		for _, resource := range each.Resources {
			rolePolicy.Resources = append(rolePolicy.Resources, resource)
		}
		for _, resource := range each.ResourceNames {
			rolePolicy.ResourceName = append(rolePolicy.ResourceName, resource)
		}
		role.ServiceAttributes.Rules = append(role.ServiceAttributes.Rules, rolePolicy)
	}
	return role, nil
}

func ConvertToCPConfigMap(cm *v1.ConfigMap) (*types.ConfigMap, error) {
	var configMap = new(types.ConfigMap)
	configMap.Name = cm.Name
	configMap.Namespace = cm.Namespace
	if vr := cm.Labels["version"]; vr != "" {
		configMap.Version = vr
	}
	configMap.ServiceType = "k8s"
	configMap.ServiceSubType = "configmap"
	configMap.ServiceAttributes = new(types.ConfigMapServiceAttribute)
	if len(cm.Data) > 0 {
		configMap.ServiceAttributes.Data = make(map[string]string)
	}
	for key, value := range cm.Data {
		configMap.ServiceAttributes.Data[key] = value
	}
	return configMap, nil
}

func convertToCPKubernetesService(svc *v1.Service) (*types.Service, error) {
	var service = new(types.Service)
	service.Name = svc.Name
	service.Namespace = svc.Namespace
	if vr := svc.Labels["version"]; vr != "" {
		service.Version = vr
	}
	service.ServiceType = "k8s"
	service.ServiceSubType = "kubernetesservice"
	service.ServiceAttributes.Type = string(svc.Spec.Type)
	if len(svc.Spec.Selector) > 0 {
		service.ServiceAttributes.Selector = make(map[string]string)
	}
	for key, value := range svc.Spec.Selector {
		service.ServiceAttributes.Selector[key] = value
	}
	service.ServiceAttributes.ExternalTrafficPolicy = string(svc.Spec.ExternalTrafficPolicy)
	for _, each := range svc.Spec.Ports {
		cpPort := types.KubePort{}
		cpPort.Name = each.Name
		cpPort.Port = each.Port
		if !(svc.Spec.ClusterIP == "None" || svc.Spec.ClusterIP == "") {
			if each.TargetPort.Type == intstr.String {
				cpPort.TargetPort.PortName = each.TargetPort.StrVal
			} else if each.TargetPort.Type == intstr.Int {
				cpPort.TargetPort.PortNumber = each.TargetPort.IntVal
			}
		} else {
			service.ServiceAttributes.ClusterIP = "None"
		}
		cpPort.Protocol = string(each.Protocol)
		cpPort.NodePort = each.NodePort

		service.ServiceAttributes.Ports = append(service.ServiceAttributes.Ports, cpPort)
	}

	return service, nil
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

func convertToCPServiceAccount(sa *v1.ServiceAccount) (*types.ServiceAccount, error) {
	var kube = new(types.ServiceAccount)
	kube.ServiceSubType = "serviceaccount"
	kube.ServiceType = "k8s"
	kube.Name = sa.Name
	kube.Namespace = sa.Namespace
	for _, value := range sa.Secrets {
		kube.ServiceAttributes.Secrets = append(kube.ServiceAttributes.Secrets, value.Name)
	}
	for _, value := range sa.ImagePullSecrets {
		kube.ServiceAttributes.ImagePullSecretsName = append(kube.ServiceAttributes.ImagePullSecretsName, value.Name)
	}
	return kube, nil

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
			tempVolume.VolumeSource.Secret = new(types.SecretVolumeSource)
			tempVolume.VolumeSource.Secret.SecretName = volume.VolumeSource.Secret.SecretName
			tempVolume.VolumeSource.Secret.DefaultMode = volume.VolumeSource.Secret.DefaultMode
			var secretItems []types.KeyToPath
			for _, item := range volume.VolumeSource.Secret.Items {
				secretItem := types.KeyToPath{
					Key:  item.Key,
					Path: item.Path,
					Mode: item.Mode,
				}
				secretItems = append(secretItems, secretItem)
			}
			tempVolume.VolumeSource.Secret.Items = secretItems
		}
		if volume.VolumeSource.ConfigMap != nil {
			tempVolume.VolumeSource.ConfigMap = new(types.ConfigMapVolumeSource)
			tempVolume.VolumeSource.ConfigMap.Name = volume.VolumeSource.ConfigMap.LocalObjectReference.Name

			tempVolume.VolumeSource.ConfigMap.DefaultMode = volume.VolumeSource.ConfigMap.DefaultMode
			var configMapItems []types.KeyToPath
			for _, item := range volume.VolumeSource.ConfigMap.Items {
				configMapItem := types.KeyToPath{
					Key:  item.Key,
					Path: item.Path,
					Mode: item.Mode,
				}
				configMapItems = append(configMapItems, configMapItem)
			}
			tempVolume.VolumeSource.ConfigMap.Items = configMapItems
		}

		if volume.VolumeSource.AWSElasticBlockStore != nil {
			tempVolume.VolumeSource.AWSElasticBlockStore = new(types.AWSElasticBlockStoreVolumeSource)
			tempVolume.VolumeSource.AWSElasticBlockStore.ReadOnly = volume.VolumeSource.AWSElasticBlockStore.ReadOnly
			tempVolume.VolumeSource.AWSElasticBlockStore.Partition = volume.VolumeSource.AWSElasticBlockStore.Partition
		}

		if volume.VolumeSource.EmptyDir != nil {
			tempVolume.VolumeSource.EmptyDir = new(types.EmptyDirVolumeSource)
			//quantity, _ := resource.ParseQuantity(volume.VolumeSource.EmptyDir.SizeLimit)
			tempVolume.VolumeSource.EmptyDir.SizeLimit = volume.VolumeSource.EmptyDir.SizeLimit
			if volume.VolumeSource.EmptyDir.Medium == v1.StorageMediumDefault {
				tempVolume.VolumeSource.EmptyDir.Medium = types.StorageMediumDefault

			}
			if volume.VolumeSource.EmptyDir.Medium == v1.StorageMediumMemory {
				tempVolume.VolumeSource.EmptyDir.Medium = types.StorageMediumMemory
			}

			if volume.VolumeSource.EmptyDir.Medium == v1.StorageMediumHugePages {
				tempVolume.VolumeSource.EmptyDir.Medium = types.StorageMediumHugePages
			}

		}

		if volume.VolumeSource.GCEPersistentDisk != nil {
			tempVolume.VolumeSource.GCEPersistentDisk = new(types.GCEPersistentDiskVolumeSource)
			tempVolume.VolumeSource.GCEPersistentDisk.Partition = volume.VolumeSource.GCEPersistentDisk.Partition
			tempVolume.VolumeSource.GCEPersistentDisk.ReadOnly = volume.VolumeSource.GCEPersistentDisk.ReadOnly
			tempVolume.VolumeSource.GCEPersistentDisk.PDName = volume.VolumeSource.GCEPersistentDisk.PDName
		}

		if volume.VolumeSource.AzureDisk != nil {
			tempVolume.VolumeSource.AzureFile = new(types.AzureFileVolumeSource)
			tempVolume.VolumeSource.AzureDisk.ReadOnly = volume.VolumeSource.AzureDisk.ReadOnly
			tempVolume.VolumeSource.AzureDisk.DataDiskURI = volume.VolumeSource.AzureDisk.DiskName

			if *volume.VolumeSource.AzureDisk.CachingMode == v1.AzureDataDiskCachingNone {
				temp := types.AzureDataDiskCachingNone
				tempVolume.VolumeSource.AzureDisk.CachingMode = &temp
			} else if *volume.VolumeSource.AzureDisk.CachingMode == v1.AzureDataDiskCachingReadWrite {
				temp := types.AzureDataDiskCachingReadWrite
				tempVolume.VolumeSource.AzureDisk.CachingMode = &temp
			} else if *volume.VolumeSource.AzureDisk.CachingMode == v1.AzureDataDiskCachingReadOnly {
				temp := types.AzureDataDiskCachingReadOnly
				tempVolume.VolumeSource.AzureDisk.CachingMode = &temp
			}

			if *volume.VolumeSource.AzureDisk.Kind == v1.AzureSharedBlobDisk {
				temp := types.AzureSharedBlobDisk
				tempVolume.VolumeSource.AzureDisk.Kind = &temp
			} else if *volume.VolumeSource.AzureDisk.Kind == v1.AzureDedicatedBlobDisk {
				temp := types.AzureDedicatedBlobDisk
				tempVolume.VolumeSource.AzureDisk.Kind = &temp
			} else if *volume.VolumeSource.AzureDisk.Kind == v1.AzureManagedDisk {
				temp := types.AzureManagedDisk
				tempVolume.VolumeSource.AzureDisk.Kind = &temp
			}
		}

		if volume.VolumeSource.AzureFile != nil {
			tempVolume.VolumeSource.AzureFile = new(types.AzureFileVolumeSource)
			tempVolume.VolumeSource.AzureFile.ReadOnly = volume.VolumeSource.AzureFile.ReadOnly
			tempVolume.VolumeSource.AzureFile.SecretName = volume.VolumeSource.AzureFile.SecretName
			tempVolume.VolumeSource.AzureFile.ShareName = volume.VolumeSource.AzureFile.ShareName

		}
		if volume.VolumeSource.HostPath != nil {
			tempVolume.VolumeSource.HostPath = new(types.HostPathVolumeSource)
			tempVolume.VolumeSource.HostPath.Path = volume.VolumeSource.HostPath.Path
			if volume.VolumeSource.HostPath.Type != nil {
				if *volume.VolumeSource.HostPath.Type == v1.HostPathUnset {
					hostPathType := types.HostPathUnset
					tempVolume.VolumeSource.HostPath.Type = &hostPathType
				}
			}

		}

		volumes = append(volumes, tempVolume)

	}

	return volumes, nil
}

func convertToCPVirtualService(input *v1alpha3.VirtualService) (*types.VirtualService, error) {
	var vServ = new(types.VirtualService)
	if input.Labels["version"] != "" {
		vServ.Version = input.Labels["version"]
	} else {
		vServ.Version = ""
	}
	vServ.ServiceType = "mesh"
	vServ.ServiceSubType = "virtual_service"
	vServ.Name = input.Name
	vServ.Namespace = input.Namespace
	vServ.ServiceAttributes = new(types.VSServiceAttribute)
	vServ.ServiceAttributes.Hosts = input.Spec.Hosts
	vServ.ServiceAttributes.Gateways = input.Spec.Gateways

	for _, http := range input.Spec.Http {
		vSer := new(types.Http)

		for _, match := range http.Match {
			m := new(types.HttpMatchRequest)
			m.Name = match.Name
			if match.Uri != nil {
				m.Uri = new(types.HttpMatch)
				matchArray := strings.Split(match.Uri.String(), ":")
				if matchArray[0] == "prefix" {
					m.Uri.Type = "prefix"
					m.Uri.Value = matchArray[1]
				} else if matchArray[0] == "exact" {
					m.Uri.Type = "exact"
					m.Uri.Value = matchArray[1]
				} else if matchArray[0] == "regex" {
					m.Uri.Type = "regex"
					m.Uri.Value = matchArray[1]
				}
			}
			if match.Scheme != nil {
				m.Scheme = new(types.HttpMatch)
				matchArray := strings.Split(match.Scheme.String(), ":")
				if matchArray[0] == "prefix" {
					m.Scheme.Type = "prefix"
					m.Scheme.Value = matchArray[1]
				} else if matchArray[0] == "exact" {
					m.Scheme.Type = "exact"
					m.Scheme.Value = matchArray[1]
				} else if matchArray[0] == "regex" {
					m.Scheme.Type = "regex"
					m.Scheme.Value = matchArray[1]
				}
			}
			if match.Method != nil {
				m.Method = new(types.HttpMatch)
				matchArray := strings.Split(match.Method.String(), ":")
				if matchArray[0] == "prefix" {
					m.Method.Type = "prefix"
					m.Method.Value = matchArray[1]
				} else if matchArray[0] == "exact" {
					m.Method.Type = "exact"
					m.Method.Value = matchArray[1]
				} else if matchArray[0] == "regex" {
					m.Method.Type = "regex"
					m.Method.Value = matchArray[1]
				}
			}
			if match.Authority != nil {
				m.Authority = new(types.HttpMatch)
				matchArray := strings.Split(match.Authority.String(), ":")
				if matchArray[0] == "prefix" {
					m.Authority.Type = "prefix"
					m.Authority.Value = matchArray[1]
				} else if matchArray[0] == "exact" {
					m.Authority.Type = "exact"
					m.Authority.Value = matchArray[1]
				} else if matchArray[0] == "regex" {
					m.Authority.Type = "regex"
					m.Authority.Value = matchArray[1]
				}
			}
			vSer.HttpMatch = append(vSer.HttpMatch, m)
		}

		for _, route := range http.Route {
			r := new(types.HttpRoute)

			if route.Destination != nil {
				destRoute := new(types.RouteDestination)
				destRoute.Host = route.Destination.Host
				destRoute.Subnet = route.Destination.Subset
				if route.Destination.Port != nil {
					destRoute.Port = int32(route.Destination.Port.Number)
				}
				r.Routes = append(r.Routes, destRoute)

			}
			r.Weight = route.Weight
			vSer.HttpRoute = append(vSer.HttpRoute, r)
		}
		if http.Redirect != nil {
			vSer.HttpRedirect = new(types.HttpRedirect)
			vSer.HttpRedirect.Uri = http.Redirect.Uri
			vSer.HttpRedirect.Authority = http.Redirect.Authority
			vSer.HttpRedirect.RedirectCode = int32(http.Redirect.RedirectCode)
		}
		if http.Rewrite != nil {
			vSer.HttpRewrite = new(types.HttpRewrite)
			vSer.HttpRewrite.Uri = http.Rewrite.Uri
			vSer.HttpRewrite.Authority = http.Rewrite.Authority
		}
		if http.Timeout != nil {
			vSer.Timeout = time.Duration(http.Timeout.Seconds)
		}

		if http.Fault != nil {
			vSer.FaultInjection = new(types.HttpFaultInjection)
			if http.Fault.GetDelay() != nil {
				if http.Fault.GetDelay().String() == "FixedDelay" {
					vSer.FaultInjection.DelayType = "FixedDelay"
					vSer.FaultInjection.DelayValue = time.Duration(http.Fault.Delay.GetFixedDelay().Seconds)
				} else if http.Fault.GetDelay().String() == "ExponentialDelay" {
					vSer.FaultInjection.DelayType = "ExponentialDelay"
					vSer.FaultInjection.DelayValue = time.Duration(http.Fault.Delay.GetExponentialDelay().Seconds)
					vSer.FaultInjection.FaultPercentage = float32(http.Fault.Delay.GetPercentage().Value)
				}
			}
			if http.Fault.Abort != nil {
				if http.Fault.Abort.String() == "HttpStatus" {
					vSer.FaultInjection.AbortErrorValue = strconv.Itoa(int(http.Fault.Abort.GetHttpStatus()))
					vSer.FaultInjection.AbortErrorType = "HttpStatus"
				} else if http.Fault.Abort.String() == "GrpcStatus" {
					vSer.FaultInjection.AbortErrorType = "GrpcStatus"
					vSer.FaultInjection.AbortErrorValue = http.Fault.Abort.GetGrpcStatus()
				} else if http.Fault.Abort.String() == "Http2Status" {
					vSer.FaultInjection.AbortErrorType = "Http2Status"
					vSer.FaultInjection.AbortErrorValue = http.Fault.Abort.GetHttp2Error()
				}

			}

		}

		if http.CorsPolicy != nil {
			vSer.CorsPolicy = new(types.HttpCorsPolicy)
			vSer.CorsPolicy.AllowOrigin = http.CorsPolicy.AllowOrigin
			vSer.CorsPolicy.AllowMethod = http.CorsPolicy.AllowMethods
			vSer.CorsPolicy.AllowHeaders = http.CorsPolicy.AllowHeaders
			vSer.CorsPolicy.ExposeHeaders = http.CorsPolicy.ExposeHeaders
			vSer.CorsPolicy.MaxAge = time.Duration(http.CorsPolicy.MaxAge.Seconds)
			vSer.CorsPolicy.AllowCredentials = http.CorsPolicy.AllowCredentials.Value
		}

		vServ.ServiceAttributes.Http = append(vServ.ServiceAttributes.Http, vSer)
	}

	for _, serv := range input.Spec.Tls {
		tls := new(types.Tls)
		for _, match := range serv.Match {
			m := new(types.TlsMatchAttribute)
			for _, s := range match.SniHosts {
				m.SniHosts = append(m.SniHosts, s)
			}
			for _, d := range match.DestinationSubnets {
				m.DestinationSubnets = append(m.DestinationSubnets, d)
			}
			for _, g := range match.Gateways {
				m.Gateways = append(m.Gateways, g)
			}
			m.Port = int32(match.Port)
			//m.SourceSubnet = match.SourceSubnet
			tls.Match = append(tls.Match, m)
		}

		for _, route := range serv.Route {
			r := new(types.TlsRoute)
			if route.Destination != nil {
				r.RouteDestination = new(types.RouteDestination)
				r.Weight = route.Weight
				if route.Destination.Port != nil {
					r.RouteDestination.Port = int32(route.Destination.Port.Number)
				}
				r.RouteDestination.Subnet = route.Destination.Subset
				r.RouteDestination.Host = route.Destination.Host
				tls.Route = append(tls.Route, r)
			}

		}
		vServ.ServiceAttributes.Tls = append(vServ.ServiceAttributes.Tls, tls)
	}

	for _, serv := range input.Spec.Tcp {
		tcp := new(types.Tcp)
		for _, match := range serv.Match {
			m := new(types.TcpMatchRequest)
			maps := make(map[string]string)
			if len(match.SourceLabels) > 0 {
				m.SourceLabels = new(map[string]string)
				m.SourceLabels = &maps
			}
			m.DestinationSubnets = match.DestinationSubnets
			m.Gateways = match.Gateways
			m.Port = int32(match.Port)
			m.SourceSubnet = match.SourceSubnet
			tcp.Match = append(tcp.Match, m)
		}

		for _, route := range serv.Route {
			d := new(types.TcpRoutes)
			if route.Destination != nil {
				d.Destination = new(types.RouteDestination)
				if route.Destination.Port != nil {
					d.Destination.Port = int32(route.Destination.Port.Number)
				}
				d.Destination.Subnet = route.Destination.Subset
				d.Destination.Host = route.Destination.Host
			}
			d.Weight = route.Weight
			tcp.Routes = append(tcp.Routes, d)
		}
		vServ.ServiceAttributes.Tcp = append(vServ.ServiceAttributes.Tcp, tcp)
	}
	return vServ, nil
}

func convertToCPDestinationRule(input *v1alpha3.DestinationRule) (*types.DestinationRules, error) {
	vServ := new(types.DestinationRules)
	vServ.ServiceType = "mesh"
	vServ.ServiceSubType = "DestinationRule"
	vServ.Name = input.Name
	vServ.Namespace = input.Namespace

	vServ.ServiceAttributes.Host = input.Spec.Host
	if input.Spec.TrafficPolicy != nil {
		vServ.ServiceAttributes.TrafficPolicy = new(types.TrafficPolicy)

		if input.Spec.TrafficPolicy.LoadBalancer != nil {
			vServ.ServiceAttributes.TrafficPolicy.LoadBalancer = new(types.LoadBalancer)
			loadBalType := strings.Split(input.Spec.TrafficPolicy.LoadBalancer.String(), ":")

			if loadBalType[0] == "simple" {
				vServ.ServiceAttributes.TrafficPolicy.LoadBalancer.Simple = input.Spec.TrafficPolicy.LoadBalancer.GetSimple().String()
			} else if loadBalType[0] == "consistent_hash" {
				vServ.ServiceAttributes.TrafficPolicy.LoadBalancer.ConsistentHash = new(types.ConsistentHash)
				if input.Spec.TrafficPolicy.LoadBalancer.GetConsistentHash().GetHttpCookie() != nil {
					vServ.ServiceAttributes.TrafficPolicy.LoadBalancer.ConsistentHash.HttpCookie = new(types.HttpCookie)
					vServ.ServiceAttributes.TrafficPolicy.LoadBalancer.ConsistentHash.HttpCookie.Name = input.Spec.TrafficPolicy.LoadBalancer.GetConsistentHash().GetHttpCookie().Name
					vServ.ServiceAttributes.TrafficPolicy.LoadBalancer.ConsistentHash.HttpCookie.Path = input.Spec.TrafficPolicy.LoadBalancer.GetConsistentHash().GetHttpCookie().Path
					vServ.ServiceAttributes.TrafficPolicy.LoadBalancer.ConsistentHash.HttpCookie.Ttl = input.Spec.TrafficPolicy.LoadBalancer.GetConsistentHash().GetHttpCookie().Ttl.Nanoseconds()

				} else if input.Spec.TrafficPolicy.LoadBalancer.GetConsistentHash().GetUseSourceIp() == true {
					vServ.ServiceAttributes.TrafficPolicy.LoadBalancer.ConsistentHash.HTTPHeaderName = input.Spec.TrafficPolicy.LoadBalancer.GetConsistentHash().GetHttpHeaderName()
				} else if input.Spec.TrafficPolicy.LoadBalancer.GetConsistentHash().GetHttpHeaderName() != "" {
					vServ.ServiceAttributes.TrafficPolicy.LoadBalancer.ConsistentHash.UseSourceIP = input.Spec.TrafficPolicy.LoadBalancer.GetConsistentHash().GetUseSourceIp()
				}
			}
		}
		if input.Spec.TrafficPolicy.ConnectionPool != nil {
			vServ.ServiceAttributes.TrafficPolicy.ConnectionPool = new(types.ConnectionPool)
			if input.Spec.TrafficPolicy.ConnectionPool.Tcp != nil {
				vServ.ServiceAttributes.TrafficPolicy.ConnectionPool.Tcp = new(types.DrTcp)
				if input.Spec.TrafficPolicy.ConnectionPool.Tcp.ConnectTimeout != nil {
					timeout := time.Duration(input.Spec.TrafficPolicy.ConnectionPool.Tcp.ConnectTimeout.Nanos)
					vServ.ServiceAttributes.TrafficPolicy.ConnectionPool.Tcp.ConnectTimeout = &timeout
				}

				vServ.ServiceAttributes.TrafficPolicy.ConnectionPool.Tcp.MaxConnections = input.Spec.TrafficPolicy.ConnectionPool.Tcp.MaxConnections
				if input.Spec.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive != nil {
					vServ.ServiceAttributes.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive = new(types.TcpKeepalive)
					if input.Spec.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive.Interval != nil {
						keepAlive := time.Duration(input.Spec.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive.Interval.Nanos)
						vServ.ServiceAttributes.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive.Interval = &keepAlive
					}

					vServ.ServiceAttributes.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive.Probes = input.Spec.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive.Probes
					if input.Spec.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive.Time != nil {
						timealive := time.Duration(input.Spec.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive.Time.Nanos)
						vServ.ServiceAttributes.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive.Time = &timealive
					}

				}
			}
			if input.Spec.TrafficPolicy.ConnectionPool.Http != nil {
				vServ.ServiceAttributes.TrafficPolicy.ConnectionPool.Http = new(types.DrHttp)
				if input.Spec.TrafficPolicy.ConnectionPool.Http.H2UpgradePolicy == v1alpha32.ConnectionPoolSettings_HTTPSettings_DEFAULT {
					vServ.ServiceAttributes.TrafficPolicy.ConnectionPool.Http.ConnectionPoolSettingsHTTPSettingsH2UpgradePolicy = int32(v1alpha32.ConnectionPoolSettings_HTTPSettings_DEFAULT)
				} else if input.Spec.TrafficPolicy.ConnectionPool.Http.H2UpgradePolicy == v1alpha32.ConnectionPoolSettings_HTTPSettings_DO_NOT_UPGRADE {
					vServ.ServiceAttributes.TrafficPolicy.ConnectionPool.Http.ConnectionPoolSettingsHTTPSettingsH2UpgradePolicy = int32(v1alpha32.ConnectionPoolSettings_HTTPSettings_DO_NOT_UPGRADE)

				} else {
					vServ.ServiceAttributes.TrafficPolicy.ConnectionPool.Http.ConnectionPoolSettingsHTTPSettingsH2UpgradePolicy = int32(v1alpha32.ConnectionPoolSettings_HTTPSettings_UPGRADE)
				}

				vServ.ServiceAttributes.TrafficPolicy.ConnectionPool.Http.HTTP1MaxPendingRequests = input.Spec.TrafficPolicy.ConnectionPool.Http.Http1MaxPendingRequests
				vServ.ServiceAttributes.TrafficPolicy.ConnectionPool.Http.HTTP2MaxRequests = input.Spec.TrafficPolicy.ConnectionPool.Http.Http2MaxRequests
				if input.Spec.TrafficPolicy.ConnectionPool.Http.IdleTimeout != nil {
					vServ.ServiceAttributes.TrafficPolicy.ConnectionPool.Http.IdleTimeout = input.Spec.TrafficPolicy.ConnectionPool.Http.IdleTimeout.Nanos
				}
				vServ.ServiceAttributes.TrafficPolicy.ConnectionPool.Http.MaxRequestsPerConnection = input.Spec.TrafficPolicy.ConnectionPool.Http.MaxRequestsPerConnection
				vServ.ServiceAttributes.TrafficPolicy.ConnectionPool.Http.MaxRetries = input.Spec.TrafficPolicy.ConnectionPool.Http.MaxRetries
			}

		}
		if input.Spec.TrafficPolicy.OutlierDetection != nil {
			vServ.ServiceAttributes.TrafficPolicy.OutlierDetection = new(types.OutlierDetection)
			if input.Spec.TrafficPolicy.OutlierDetection.BaseEjectionTime != nil {
				injecTime := time.Duration(input.Spec.TrafficPolicy.OutlierDetection.BaseEjectionTime.Nanos)
				vServ.ServiceAttributes.TrafficPolicy.OutlierDetection.BaseEjectionTime = &injecTime
			}

			vServ.ServiceAttributes.TrafficPolicy.OutlierDetection.ConsecutiveErrors = input.Spec.TrafficPolicy.OutlierDetection.ConsecutiveErrors
			if input.Spec.TrafficPolicy.OutlierDetection.Interval != nil {
				interval := time.Duration(input.Spec.TrafficPolicy.OutlierDetection.Interval.Nanos)
				vServ.ServiceAttributes.TrafficPolicy.OutlierDetection.Interval = &interval
			}

			vServ.ServiceAttributes.TrafficPolicy.OutlierDetection.MaxEjectionPercent = input.Spec.TrafficPolicy.OutlierDetection.MaxEjectionPercent
			vServ.ServiceAttributes.TrafficPolicy.OutlierDetection.MinHealthPercent = input.Spec.TrafficPolicy.OutlierDetection.MinHealthPercent
		}

		for _, port := range input.Spec.TrafficPolicy.PortLevelSettings {

			setting := new(types.PortLevelSetting)

			if port.Port != nil {
				setting.Port = new(types.DrPort)
				setting.Port.Number = int32(port.Port.Number)
			}

			if port.ConnectionPool != nil {
				setting.ConnectionPool = new(types.ConnectionPool)
				if setting.ConnectionPool.Tcp != nil {
					setting.ConnectionPool.Tcp = new(types.DrTcp)
					if port.ConnectionPool.Tcp.ConnectTimeout != nil {
						timeout := time.Duration(port.ConnectionPool.Tcp.ConnectTimeout.Nanos)
						setting.ConnectionPool.Tcp.ConnectTimeout = &timeout
					}
					setting.ConnectionPool.Tcp.MaxConnections = port.ConnectionPool.Tcp.MaxConnections
					if port.ConnectionPool.Tcp.TcpKeepalive != nil {
						setting.ConnectionPool.Tcp.TcpKeepalive = new(types.TcpKeepalive)
						if port.ConnectionPool.Tcp.TcpKeepalive.Time != nil {
							t := time.Duration(port.ConnectionPool.Tcp.TcpKeepalive.Time.Nanos)
							setting.ConnectionPool.Tcp.TcpKeepalive.Time = &t
						}
						if port.ConnectionPool.Tcp.TcpKeepalive.Interval != nil {
							interval := time.Duration(port.ConnectionPool.Tcp.TcpKeepalive.Interval.Nanos)
							setting.ConnectionPool.Tcp.TcpKeepalive.Interval = &interval
						}

						setting.ConnectionPool.Tcp.TcpKeepalive.Probes = port.ConnectionPool.Tcp.TcpKeepalive.Probes

					}

				}
				if port.ConnectionPool.Http != nil {
					setting.ConnectionPool.Http = new(types.DrHttp)
					setting.ConnectionPool.Http.ConnectionPoolSettingsHTTPSettingsH2UpgradePolicy = int32(port.ConnectionPool.Http.H2UpgradePolicy)
					setting.ConnectionPool.Http.HTTP2MaxRequests = port.ConnectionPool.Http.Http2MaxRequests
					setting.ConnectionPool.Http.HTTP1MaxPendingRequests = port.ConnectionPool.Http.Http1MaxPendingRequests
					if port.ConnectionPool.Http.IdleTimeout != nil {
						setting.ConnectionPool.Http.IdleTimeout = port.ConnectionPool.Http.IdleTimeout.Nanos
					}
					setting.ConnectionPool.Http.MaxRequestsPerConnection = port.ConnectionPool.Http.MaxRequestsPerConnection
					setting.ConnectionPool.Http.MaxRetries = port.ConnectionPool.Http.MaxRetries

				}

			}
			if port.LoadBalancer != nil {
				setting.LoadBalancer = new(types.LoadBalancer)
				if port.LoadBalancer.GetSimple().String() != "" {
					setting.LoadBalancer.Simple = port.LoadBalancer.GetSimple().String()
				} else if port.LoadBalancer.GetConsistentHash() != nil {
					setting.LoadBalancer.ConsistentHash = new(types.ConsistentHash)
					if port.LoadBalancer.GetConsistentHash().GetHttpHeaderName() != "" {
						setting.LoadBalancer.ConsistentHash.HTTPHeaderName = port.LoadBalancer.GetConsistentHash().GetHttpHeaderName()
					} else if port.LoadBalancer.GetConsistentHash().GetUseSourceIp() != false {
						setting.LoadBalancer.ConsistentHash.UseSourceIP = port.LoadBalancer.GetConsistentHash().GetUseSourceIp()
					} else if port.LoadBalancer.GetConsistentHash().GetHttpCookie() != nil {
						setting.LoadBalancer.ConsistentHash.HttpCookie = new(types.HttpCookie)
						setting.LoadBalancer.ConsistentHash.HttpCookie.Name = port.LoadBalancer.GetConsistentHash().GetHttpCookie().Name
						setting.LoadBalancer.ConsistentHash.HttpCookie.Path = port.LoadBalancer.GetConsistentHash().GetHttpCookie().Path

						setting.LoadBalancer.ConsistentHash.HttpCookie.Ttl = port.LoadBalancer.GetConsistentHash().GetHttpCookie().Ttl.Nanoseconds()

					}

					setting.LoadBalancer.ConsistentHash.MinimumRingSize = strconv.Itoa(int(port.LoadBalancer.GetConsistentHash().GetMinimumRingSize()))

				}
			}
			if port.OutlierDetection != nil {
				setting.OutlierDetection = new(types.OutlierDetection)
				if port.OutlierDetection.Interval != nil {
					interval := time.Duration(port.OutlierDetection.Interval.Nanos)
					setting.OutlierDetection.Interval = &interval
				}
				if port.OutlierDetection.BaseEjectionTime != nil {
					ejec := time.Duration(port.OutlierDetection.BaseEjectionTime.Nanos)
					setting.OutlierDetection.BaseEjectionTime = &ejec
				}
				setting.OutlierDetection.ConsecutiveErrors = port.OutlierDetection.ConsecutiveErrors
				setting.OutlierDetection.MaxEjectionPercent = port.OutlierDetection.MaxEjectionPercent
				setting.OutlierDetection.MinHealthPercent = port.OutlierDetection.MinHealthPercent

			}

			vServ.ServiceAttributes.TrafficPolicy.PortLevelSettings = append(vServ.ServiceAttributes.TrafficPolicy.PortLevelSettings, setting)
		}
		if input.Spec.TrafficPolicy.Tls != nil {
			vServ.ServiceAttributes.TrafficPolicy.DrTls = new(types.DrTls)
			vServ.ServiceAttributes.TrafficPolicy.DrTls.Mode = input.Spec.TrafficPolicy.Tls.GetMode().String()
			vServ.ServiceAttributes.TrafficPolicy.DrTls.ClientCertificate = input.Spec.TrafficPolicy.Tls.ClientCertificate
			vServ.ServiceAttributes.TrafficPolicy.DrTls.PrivateKey = input.Spec.TrafficPolicy.Tls.PrivateKey
			vServ.ServiceAttributes.TrafficPolicy.DrTls.CaCertificate = input.Spec.TrafficPolicy.Tls.CaCertificates
			vServ.ServiceAttributes.TrafficPolicy.DrTls.SubjectAltNames = input.Spec.TrafficPolicy.Tls.SubjectAltNames[0]
		}

	}
	for _, subset := range input.Spec.Subsets {
		ser := new(types.Subset)
		ser.Name = subset.Name
		if len(subset.Labels) > 0 {
			labels := make(map[string]string)
			labels = subset.Labels
			ser.Labels = &labels
		}

		if subset.TrafficPolicy != nil {
			ser.TrafficPolicy = new(types.TrafficPolicy)
			for _, port := range subset.TrafficPolicy.PortLevelSettings {
				setting := new(types.PortLevelSetting)
				if port.Port != nil {
					setting.Port = new(types.DrPort)
					setting.Port.Number = int32(port.Port.Number)
				}
				if port.ConnectionPool != nil {
					setting.ConnectionPool = new(types.ConnectionPool)
					if setting.ConnectionPool.Tcp != nil {
						setting.ConnectionPool.Tcp = new(types.DrTcp)
						if port.ConnectionPool.Tcp.ConnectTimeout != nil {
							timeout := time.Duration(port.ConnectionPool.Tcp.ConnectTimeout.Nanos)
							setting.ConnectionPool.Tcp.ConnectTimeout = &timeout
						}
						setting.ConnectionPool.Tcp.MaxConnections = port.ConnectionPool.Tcp.MaxConnections
						if port.ConnectionPool.Tcp.TcpKeepalive != nil {
							setting.ConnectionPool.Tcp.TcpKeepalive = new(types.TcpKeepalive)
							if port.ConnectionPool.Tcp.TcpKeepalive.Time != nil {
								t := time.Duration(port.ConnectionPool.Tcp.TcpKeepalive.Time.Nanos)
								setting.ConnectionPool.Tcp.TcpKeepalive.Time = &t
							}
							if port.ConnectionPool.Tcp.TcpKeepalive.Interval != nil {
								interval := time.Duration(port.ConnectionPool.Tcp.TcpKeepalive.Interval.Nanos)
								setting.ConnectionPool.Tcp.TcpKeepalive.Interval = &interval
							}

							setting.ConnectionPool.Tcp.TcpKeepalive.Probes = port.ConnectionPool.Tcp.TcpKeepalive.Probes

						}

					}
					if port.ConnectionPool.Http != nil {
						setting.ConnectionPool.Http = new(types.DrHttp)
						setting.ConnectionPool.Http.ConnectionPoolSettingsHTTPSettingsH2UpgradePolicy = int32(port.ConnectionPool.Http.H2UpgradePolicy)
						setting.ConnectionPool.Http.HTTP2MaxRequests = port.ConnectionPool.Http.Http2MaxRequests
						setting.ConnectionPool.Http.HTTP1MaxPendingRequests = port.ConnectionPool.Http.Http1MaxPendingRequests
						if port.ConnectionPool.Http.IdleTimeout != nil {
							setting.ConnectionPool.Http.IdleTimeout = port.ConnectionPool.Http.IdleTimeout.Nanos
						}
						setting.ConnectionPool.Http.MaxRequestsPerConnection = port.ConnectionPool.Http.MaxRequestsPerConnection
						setting.ConnectionPool.Http.MaxRetries = port.ConnectionPool.Http.MaxRetries

					}

				}
				if port.LoadBalancer != nil {
					setting.LoadBalancer = new(types.LoadBalancer)
					if port.LoadBalancer.GetSimple().String() != "" {
						setting.LoadBalancer.Simple = port.LoadBalancer.GetSimple().String()
					} else if port.LoadBalancer.GetConsistentHash() != nil {
						setting.LoadBalancer.ConsistentHash = new(types.ConsistentHash)

						if port.LoadBalancer.GetConsistentHash().GetHttpCookie() != nil {
							setting.LoadBalancer.ConsistentHash.HttpCookie = new(types.HttpCookie)
							if port.LoadBalancer.GetConsistentHash().GetHttpCookie().Ttl != nil {
								setting.LoadBalancer.ConsistentHash.HttpCookie.Ttl = port.LoadBalancer.GetConsistentHash().GetHttpCookie().Ttl.Nanoseconds()

							}

							setting.LoadBalancer.ConsistentHash.HttpCookie.Name = port.LoadBalancer.GetConsistentHash().GetHttpCookie().GetName()
							setting.LoadBalancer.ConsistentHash.HttpCookie.Path = port.LoadBalancer.GetConsistentHash().GetHttpCookie().GetPath()
						} else if port.LoadBalancer.GetConsistentHash().GetHttpHeaderName() != "" {
							setting.LoadBalancer.ConsistentHash.HTTPHeaderName = port.LoadBalancer.GetConsistentHash().GetHttpHeaderName()
						} else if port.LoadBalancer.GetConsistentHash().GetUseSourceIp() != false {
							setting.LoadBalancer.ConsistentHash.UseSourceIP = port.LoadBalancer.GetConsistentHash().GetUseSourceIp()
						}
					}
				}
				if port.OutlierDetection != nil {
					setting.OutlierDetection = new(types.OutlierDetection)
					if port.OutlierDetection.Interval != nil {
						interval := time.Duration(port.OutlierDetection.Interval.Nanos)
						setting.OutlierDetection.Interval = &interval
					}
					if port.OutlierDetection.BaseEjectionTime != nil {
						ejec := time.Duration(port.OutlierDetection.BaseEjectionTime.Nanos)
						setting.OutlierDetection.BaseEjectionTime = &ejec
					}
					setting.OutlierDetection.ConsecutiveErrors = port.OutlierDetection.ConsecutiveErrors
					setting.OutlierDetection.MaxEjectionPercent = port.OutlierDetection.MaxEjectionPercent
					setting.OutlierDetection.MinHealthPercent = port.OutlierDetection.MinHealthPercent

				}
				if port.Tls != nil {
					setting.DrTls = new(types.DrTls)
					vServ.ServiceAttributes.TrafficPolicy.DrTls.Mode = string(port.Tls.GetMode())
					setting.DrTls.ClientCertificate = port.Tls.ClientCertificate
					setting.DrTls.PrivateKey = port.Tls.PrivateKey
					setting.DrTls.CaCertificate = port.Tls.CaCertificates
					setting.DrTls.SubjectAltNames = port.Tls.SubjectAltNames[0]

				}
				vServ.ServiceAttributes.TrafficPolicy.PortLevelSettings = append(vServ.ServiceAttributes.TrafficPolicy.PortLevelSettings, setting)
			}
			if subset.TrafficPolicy.LoadBalancer != nil {
				ser.TrafficPolicy.LoadBalancer = new(types.LoadBalancer)
				if subset.TrafficPolicy.LoadBalancer.GetSimple().String() != "" {
					ser.TrafficPolicy.LoadBalancer.Simple = subset.TrafficPolicy.LoadBalancer.GetSimple().String()
				} else if subset.TrafficPolicy.LoadBalancer.GetConsistentHash() != nil {
					ser.TrafficPolicy.LoadBalancer.ConsistentHash = new(types.ConsistentHash)

					if subset.TrafficPolicy.LoadBalancer.GetConsistentHash().GetHttpCookie() != nil {
						if subset.TrafficPolicy.LoadBalancer.GetConsistentHash().GetHttpCookie().Ttl != nil {
							ser.TrafficPolicy.LoadBalancer.ConsistentHash.HttpCookie.Ttl = subset.TrafficPolicy.LoadBalancer.GetConsistentHash().GetHttpCookie().Ttl.Nanoseconds()

						}

						ser.TrafficPolicy.LoadBalancer.ConsistentHash.HttpCookie.Name = subset.TrafficPolicy.LoadBalancer.GetConsistentHash().GetHttpCookie().GetName()
						ser.TrafficPolicy.LoadBalancer.ConsistentHash.HttpCookie.Path = subset.TrafficPolicy.LoadBalancer.GetConsistentHash().GetHttpCookie().GetPath()
					} else if subset.TrafficPolicy.LoadBalancer.GetConsistentHash().GetHttpHeaderName() != "" {
						ser.TrafficPolicy.LoadBalancer.ConsistentHash.HTTPHeaderName = subset.TrafficPolicy.LoadBalancer.GetConsistentHash().GetHttpHeaderName()
					} else if subset.TrafficPolicy.LoadBalancer.GetConsistentHash().GetUseSourceIp() != false {
						ser.TrafficPolicy.LoadBalancer.ConsistentHash.UseSourceIP = subset.TrafficPolicy.LoadBalancer.GetConsistentHash().GetUseSourceIp()
					}
				}

			}
			if subset.TrafficPolicy.ConnectionPool != nil {
				ser.TrafficPolicy.ConnectionPool = new(types.ConnectionPool)
				if subset.TrafficPolicy.ConnectionPool.Tcp != nil {
					ser.TrafficPolicy.ConnectionPool.Tcp = new(types.DrTcp)
					if subset.TrafficPolicy.ConnectionPool.Tcp.ConnectTimeout != nil {
						timeout := time.Duration(subset.TrafficPolicy.ConnectionPool.Tcp.ConnectTimeout.Nanos)
						ser.TrafficPolicy.ConnectionPool.Tcp.ConnectTimeout = &timeout
					}

					ser.TrafficPolicy.ConnectionPool.Tcp.MaxConnections = subset.TrafficPolicy.ConnectionPool.Tcp.MaxConnections
					if subset.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive != nil {
						ser.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive = new(types.TcpKeepalive)
						if subset.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive.Time != nil {
							t := time.Duration(subset.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive.Time.Nanos)
							ser.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive.Time = &t
						}
						if subset.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive.Interval != nil {
							interval := time.Duration(subset.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive.Interval.Seconds)
							ser.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive.Interval = &interval
						}

						ser.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive.Probes = subset.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive.Probes
						if subset.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive.Time != nil {
							ti := time.Duration(subset.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive.Time.Seconds)
							ser.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive.Time = &ti
						}

					}

				}
				if subset.TrafficPolicy.ConnectionPool.Http != nil {
					ser.TrafficPolicy.ConnectionPool.Http = new(types.DrHttp)
					ser.TrafficPolicy.ConnectionPool.Http.ConnectionPoolSettingsHTTPSettingsH2UpgradePolicy = int32(subset.TrafficPolicy.ConnectionPool.Http.H2UpgradePolicy)
					ser.TrafficPolicy.ConnectionPool.Http.HTTP2MaxRequests = subset.TrafficPolicy.ConnectionPool.Http.Http2MaxRequests
					ser.TrafficPolicy.ConnectionPool.Http.HTTP1MaxPendingRequests = subset.TrafficPolicy.ConnectionPool.Http.Http1MaxPendingRequests
					if subset.TrafficPolicy.ConnectionPool.Http.IdleTimeout != nil {
						ser.TrafficPolicy.ConnectionPool.Http.IdleTimeout = subset.TrafficPolicy.ConnectionPool.Http.IdleTimeout.Nanos
					}
					ser.TrafficPolicy.ConnectionPool.Http.MaxRequestsPerConnection = subset.TrafficPolicy.ConnectionPool.Http.MaxRequestsPerConnection
					ser.TrafficPolicy.ConnectionPool.Http.MaxRetries = subset.TrafficPolicy.ConnectionPool.Http.MaxRetries

				}

			}

			if subset.TrafficPolicy.OutlierDetection != nil {
				ser.TrafficPolicy.OutlierDetection = new(types.OutlierDetection)
				if subset.TrafficPolicy.OutlierDetection.Interval != nil {
					interval := time.Duration(subset.TrafficPolicy.OutlierDetection.Interval.Nanos)
					ser.TrafficPolicy.OutlierDetection.Interval = &interval
				}
				if subset.TrafficPolicy.OutlierDetection.BaseEjectionTime != nil {
					ejec := time.Duration(subset.TrafficPolicy.OutlierDetection.BaseEjectionTime.Nanos)
					ser.TrafficPolicy.OutlierDetection.BaseEjectionTime = &ejec
				}
				ser.TrafficPolicy.OutlierDetection.ConsecutiveErrors = subset.TrafficPolicy.OutlierDetection.ConsecutiveErrors
				ser.TrafficPolicy.OutlierDetection.MaxEjectionPercent = subset.TrafficPolicy.OutlierDetection.MaxEjectionPercent
				ser.TrafficPolicy.OutlierDetection.MinHealthPercent = subset.TrafficPolicy.OutlierDetection.MinHealthPercent

			}
			if subset.TrafficPolicy.Tls != nil {
				ser.TrafficPolicy.DrTls = new(types.DrTls)
				vServ.ServiceAttributes.TrafficPolicy.DrTls.Mode = string(subset.TrafficPolicy.Tls.GetMode())
				ser.TrafficPolicy.DrTls.ClientCertificate = subset.TrafficPolicy.Tls.ClientCertificate
				ser.TrafficPolicy.DrTls.PrivateKey = subset.TrafficPolicy.Tls.PrivateKey
				ser.TrafficPolicy.DrTls.CaCertificate = subset.TrafficPolicy.Tls.CaCertificates
				ser.TrafficPolicy.DrTls.SubjectAltNames = subset.TrafficPolicy.Tls.SubjectAltNames[0]
			}
		}
		vServ.ServiceAttributes.Subsets = append(vServ.ServiceAttributes.Subsets, ser)
	}

	return vServ, nil
}

func convertToCPGateway(input *v1alpha3.Gateway) (*types.GatewayService, error) {
	gateway := new(types.GatewayService)
	gateway.Name = input.Name
	if input.Labels["version"] != "" {
		gateway.Version = input.Labels["version"]
	} else {
		gateway.Version = ""
	}
	gateway.ServiceType = "mesh"
	gateway.ServiceSubType = "Gateway"
	gateway.Namespace = input.Namespace

	gateway.ServiceAttributes = new(types.GatewayServiceAttributes)
	gateway.ServiceAttributes.Selectors = make(map[string]string)
	gateway.ServiceAttributes.Selectors = input.Spec.Selector

	for _, serverInput := range input.Spec.Servers {
		server := new(types.Server)
		if serverInput.Tls != nil {
			server.Tls = new(types.TlsConfig)
			server.Tls.HttpsRedirect = serverInput.Tls.HttpsRedirect
			server.Tls.Mode = types.Mode(serverInput.Tls.Mode.String())
			server.Tls.ServerCertificate = serverInput.Tls.ServerCertificate
			server.Tls.CaCertificate = serverInput.Tls.CaCertificates
			server.Tls.PrivateKey = serverInput.Tls.PrivateKey
			for _, altNames := range serverInput.Tls.SubjectAltNames {
				server.Tls.SubjectAltName = append(serverInput.Tls.SubjectAltNames, altNames)
			}
			server.Tls.MinProtocolVersion = types.ProtocolVersion(serverInput.Tls.MinProtocolVersion.String())
			server.Tls.MaxProtocolVersion = types.ProtocolVersion(serverInput.Tls.MaxProtocolVersion.String())
		}
		if serverInput.Port != nil {
			server.Port = new(types.Port)
			server.Port.Name = serverInput.Port.Name
			server.Port.Nummber = serverInput.Port.Number
			server.Port.Protocol = types.Protocols(serverInput.Port.Protocol)
		}
		for _, host := range serverInput.Hosts {
			server.Hosts = append(server.Hosts, host)
		}

		gateway.ServiceAttributes.Servers = append(gateway.ServiceAttributes.Servers, server)

	}
	return gateway, nil

}

func convertToCPServiceEntry(input *v1alpha3.ServiceEntry) (*types.ServiceEntry, error) {
	svcEntry := new(types.ServiceEntry)
	svcEntry.Name = input.Name
	svcEntry.Namespace = input.Namespace
	svcEntry.ServiceType = "mesh"
	svcEntry.ServiceSubType = "service_entry"
	if input.Labels["version"] != "" {
		svcEntry.Version = input.Labels["version"]
	}
	svcEntry.ServiceAttributes = new(types.ServiceEntryAttributes)
	for _, host := range input.Spec.Hosts {
		svcEntry.ServiceAttributes.Hosts = append(svcEntry.ServiceAttributes.Hosts, host)
	}
	for _, address := range input.Spec.Addresses {
		svcEntry.ServiceAttributes.Addresses = append(svcEntry.ServiceAttributes.Addresses, address)
	}
	for _, port := range input.Spec.Ports {
		tempPort := new(types.ServiceEntryPort)
		tempPort.Name = port.Name
		tempPort.Protocol = port.Protocol
		tempPort.Number = port.Number
		svcEntry.ServiceAttributes.Ports = append(svcEntry.ServiceAttributes.Ports, tempPort)

	}
	for _, entryPoint := range input.Spec.Endpoints {
		tempEntryPoint := new(types.ServiceEntryEndpoint)
		tempEntryPoint.Address = entryPoint.Address
		tempEntryPoint.Locality = entryPoint.Locality
		tempEntryPoint.Network = entryPoint.Network
		tempEntryPoint.Ports = make(map[string]uint32)
		tempEntryPoint.Ports = entryPoint.Ports
		tempEntryPoint.Labels = make(map[string]string)
		tempEntryPoint.Labels = entryPoint.Labels

		svcEntry.ServiceAttributes.Endpoints = append(svcEntry.ServiceAttributes.Endpoints, tempEntryPoint)

	}

	svcEntry.ServiceAttributes.ExportTo = input.Spec.ExportTo
	for _, name := range input.Spec.SubjectAltNames {
		svcEntry.ServiceAttributes.SubjectAltNames = append(svcEntry.ServiceAttributes.SubjectAltNames, name)
	}

	return svcEntry, nil
}
