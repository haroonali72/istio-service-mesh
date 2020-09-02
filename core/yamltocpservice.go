package core

import (
	"bitbucket.org/cloudplex-devs/istio-service-mesh/constants"
	"bitbucket.org/cloudplex-devs/istio-service-mesh/types"
	"bitbucket.org/cloudplex-devs/istio-service-mesh/utils"
	meshConstants "bitbucket.org/cloudplex-devs/microservices-mesh-engine/constants"
	pb "bitbucket.org/cloudplex-devs/microservices-mesh-engine/core/services/proto"
	meshTypes "bitbucket.org/cloudplex-devs/microservices-mesh-engine/types/services"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	v1alpha32 "istio.io/api/networking/v1alpha3"
	"istio.io/client-go/pkg/apis/networking/v1alpha3"
	apps "k8s.io/api/apps/v1"
	autoScalar "k8s.io/api/autoscaling/v1"
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
		err2 := yaml.Unmarshal(req.Service, &yamlService)
		if err2 == nil {
			if yamlService.Kind == constants.Gateway.String() {
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
			} else if yamlService.Kind == constants.ServiceEntry.String() {
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
			} else if yamlService.Kind == constants.DestinationRule.String() {
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
			} else if yamlService.Kind == constants.VirtualService.String() {
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
		//		return nil, errors.New(fmt.Sprintf("Error while decoding YAML object. Err was: %s", err))
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
	case *v1.Pod:
		pod, err := convertToCPPod(o)
		if err != nil {
			return nil, err
		}
		bytesData, err := json.Marshal(pod)
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

		//default:
		//	return nil, errors.New("object is not in our scope")
	}
	return serviceResp, nil

}

func convertToCPNetwokPolicy(np *net.NetworkPolicy) (*meshTypes.NetworkPolicyService, error) {
	networkPolicy := new(meshTypes.NetworkPolicyService)
	networkPolicy.Name = np.Name
	if np.Namespace == "" {
		networkPolicy.Namespace = "default"
	} else {
		networkPolicy.Namespace = np.Namespace
	}
	networkPolicy.ServiceType = meshConstants.Kubernetes
	networkPolicy.ServiceSubType = meshConstants.NetworkPolicy
	networkPolicy.ServiceAttributes = new(meshTypes.NetworkPolicyServiceAttribute)
	networkPolicy.ServiceAttributes.PodSelector = getCPLabelSelector(&np.Spec.PodSelector)
	for _, each := range np.Spec.Ingress {
		temp := meshTypes.IngressRule{}
		for _, ePort := range each.Ports {
			tp := meshTypes.NetworkPolicyPort{}
			tp.Protocol = (*meshTypes.Protocol)(ePort.Protocol)
			if ePort.Port.Type == intstr.Int {
				tp.Port.PortNumber = ePort.Port.IntVal
			} else {
				tp.Port.PortName = ePort.Port.StrVal
			}
			temp.Ports = append(temp.Ports, tp)
		}
		for _, from := range each.From {
			fm := meshTypes.NetworkPolicyPeer{}
			fm.PodSelector = getCPLabelSelector(from.PodSelector)
			fm.NamespaceSelector = getCPLabelSelector(from.NamespaceSelector)
			if from.IPBlock != nil {
				fm.IPBlock = new(meshTypes.IPBlock)
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
		temp := meshTypes.EgressRule{}
		for _, ePort := range each.Ports {
			tp := meshTypes.NetworkPolicyPort{}
			tp.Protocol = (*meshTypes.Protocol)(ePort.Protocol)
			if ePort.Port.Type == intstr.Int {
				tp.Port.PortNumber = ePort.Port.IntVal
			} else {
				tp.Port.PortName = ePort.Port.StrVal
			}
			temp.Ports = append(temp.Ports, tp)
		}
		for _, from := range each.To {
			fm := meshTypes.NetworkPolicyPeer{}
			fm.PodSelector = getCPLabelSelector(from.PodSelector)
			fm.NamespaceSelector = getCPLabelSelector(from.NamespaceSelector)
			if from.IPBlock != nil {
				fm.IPBlock = new(meshTypes.IPBlock)
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

func getCPLabelSelector(selector *metav1.LabelSelector) *meshTypes.LabelSelectorObj {
	if selector == nil {
		return nil
	}
	ls := new(meshTypes.LabelSelectorObj)
	ls.MatchLabels = selector.MatchLabels
	for _, each := range selector.MatchExpressions {
		temp := meshTypes.LabelSelectorRequirement{}
		temp.Key = each.Key
		temp.Operator = meshTypes.LabelSelectorOperator(each.Operator)
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

func convertToCPDeployment(deploy interface{}) (*meshTypes.DeploymentService, error) {
	byteData, _ := json.Marshal(deploy)
	service := apps.Deployment{}
	json.Unmarshal(byteData, &service)

	deployment := new(meshTypes.DeploymentService)
	deployment.ServiceAttributes = new(meshTypes.DeploymentServiceAttribute)

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

	deployment.ServiceType = meshConstants.Kubernetes
	deployment.ServiceSubType = meshConstants.Deployment
	if service.Labels["version"] != "" {
		deployment.Version = service.Labels["version"]
	}

	if service.Spec.Replicas != nil {
		deployment.ServiceAttributes.Replicas = service.Spec.Replicas

	} else {
		var a int32 = 1
		deployment.ServiceAttributes.Replicas = &a
	}

	deployment.ServiceAttributes.Labels = make(map[string]string)
	deployment.ServiceAttributes.Labels = service.Spec.Template.Labels
	deployment.ServiceAttributes.LabelSelector = new(meshTypes.LabelSelectorObj)
	deployment.ServiceAttributes.LabelSelector.MatchLabels = make(map[string]string)
	//deployment.ServiceAttributes.LabelSelector.MatchLabels["version"] = service.Labels["version"]
	//deployment.ServiceAttributes.LabelSelector.MatchLabels["name"] = service.Labels["name"]
	if service.Spec.Selector != nil && service.Spec.Selector.MatchLabels != nil {
		deployment.ServiceAttributes.LabelSelector.MatchLabels = service.Spec.Selector.MatchLabels
	}
	deployment.ServiceAttributes.Annotations = make(map[string]string)
	deployment.ServiceAttributes.Annotations = service.Spec.Template.Annotations
	deployment.ServiceAttributes.NodeSelector = make(map[string]string)
	deployment.ServiceAttributes.NodeSelector = service.Spec.Template.Spec.NodeSelector

	if service.Spec.Template.Spec.ServiceAccountName != "" {
		deployment.ServiceAttributes.ServiceAccountName = service.Spec.Template.Spec.ServiceAccountName
	}
	if service.Spec.Strategy.Type != "" {
		if service.Spec.Strategy.Type == apps.RecreateDeploymentStrategyType {
			var CpDepStrategy = new(meshTypes.DeploymentStrategy)
			CpDepStrategy.Type = meshTypes.RecreateDeploymentStrategyType
			deployment.ServiceAttributes.Strategy = CpDepStrategy
		} else if service.Spec.Strategy.Type == apps.RollingUpdateDeploymentStrategyType {
			deployment.ServiceAttributes.Strategy = new(meshTypes.DeploymentStrategy)
			deployment.ServiceAttributes.Strategy.Type = meshTypes.RollingUpdateDeploymentStrategyType
			if service.Spec.Strategy.RollingUpdate != nil {
				deployment.ServiceAttributes.Strategy.RollingUpdate = new(meshTypes.RollingUpdateDeployment)
				//deployment.ServiceAttributes.Strategy.RollingUpdate.MaxSurge = new(intstr.IntOrString)
				//deployment.ServiceAttributes.Strategy.RollingUpdate.MaxUnavailable = new(intstr.IntOrString)
				//deployment.ServiceAttributes.Strategy.RollingUpdate.MaxSurge.Type = service.Spec.Strategy.RollingUpdate.MaxSurge.Type
				//deployment.ServiceAttributes.Strategy.RollingUpdate.MaxSurge.IntVal = service.Spec.Strategy.RollingUpdate.MaxUnavailable.IntVal
				//deployment.ServiceAttributes.Strategy.RollingUpdate.MaxSurge.StrVal = service.Spec.Strategy.RollingUpdate.MaxUnavailable.StrVal
				//deployment.ServiceAttributes.Strategy.RollingUpdate.MaxUnavailable.Type = service.Spec.Strategy.RollingUpdate.MaxUnavailable.Type
				//deployment.ServiceAttributes.Strategy.RollingUpdate.MaxUnavailable.IntVal = service.Spec.Strategy.RollingUpdate.MaxUnavailable.IntVal
				//deployment.ServiceAttributes.Strategy.RollingUpdate.MaxUnavailable.StrVal = service.Spec.Strategy.RollingUpdate.MaxUnavailable.StrVal
			}
		}

	}

	for _, imageSecrets := range service.Spec.Template.Spec.ImagePullSecrets {
		tempImageSecrets := meshTypes.LocalObjectReference{Name: imageSecrets.Name}
		deployment.ServiceAttributes.ImagePullSecrets = append(deployment.ServiceAttributes.ImagePullSecrets, tempImageSecrets)
	}

	var volumeMountNames1 = make(map[string]bool)
	if containers, vm, err := getCPContainers(service.Spec.Template.Spec.Containers, service.Spec.Template.Spec.Volumes); err == nil {
		if len(containers) > 0 {
			deployment.ServiceAttributes.Containers = containers
			volumeMountNames1 = vm
		} else {
			utils.Error.Println("no containers exist")
			return nil, errors.New("no containers exist")
		}

	} else {
		utils.Error.Println(err)
		return nil, err
	}

	if containersList, volumeMounts, err := getCPContainers(service.Spec.Template.Spec.InitContainers, service.Spec.Template.Spec.Volumes); err == nil {
		if len(containersList) > 0 {
			deployment.ServiceAttributes.InitContainers = containersList
		}
		for k, v := range volumeMounts {
			volumeMountNames1[k] = v
		}

	} else {
		utils.Error.Println(err)
		return nil, err
	}

	if vols, err := getCPVolumes(service.Spec.Template.Spec.Volumes, volumeMountNames1); err == nil {
		if len(vols) > 0 {
			deployment.ServiceAttributes.Volumes = vols
		}

	} else {
		utils.Error.Println(err)
		return nil, err
	}

	if service.Spec.Template.Spec.Affinity != nil {
		if affinity, err := getCPAffinity(service.Spec.Template.Spec.Affinity); err == nil {
			deployment.ServiceAttributes.Affinity = affinity
		} else {
			utils.Error.Println(err)
			return nil, err
		}
	}
	return deployment, nil
}

func convertToCPPod(service *v1.Pod) (*meshTypes.PodService, error) {
	pod := new(meshTypes.PodService)
	pod.ServiceAttributes = new(meshTypes.PodServiceAttribute)

	if service.Name == "" {
		return nil, errors.New("Service name not found")
	} else {
		pod.Name = service.Name
	}

	if service.Namespace == "" {
		pod.Namespace = "default"
	} else {
		pod.Namespace = service.Namespace
	}

	pod.ServiceType = "k8s"
	pod.ServiceSubType = meshConstants.Pod
	if service.Labels["version"] != "" {
		pod.Version = service.Labels["version"]
	}

	pod.ServiceAttributes.Labels = make(map[string]string)
	pod.ServiceAttributes.Labels = service.Labels
	pod.ServiceAttributes.Annotations = make(map[string]string)
	pod.ServiceAttributes.Annotations = service.Annotations
	pod.ServiceAttributes.NodeSelector = make(map[string]string)
	pod.ServiceAttributes.NodeSelector = service.Spec.NodeSelector

	for _, imageSecrets := range service.Spec.ImagePullSecrets {
		tempImageSecrets := meshTypes.LocalObjectReference{Name: imageSecrets.Name}
		pod.ServiceAttributes.ImagePullSecrets = append(pod.ServiceAttributes.ImagePullSecrets, tempImageSecrets)
	}

	var volumeMountNames1 = make(map[string]bool)
	if containers, vm, err := getCPContainers(service.Spec.Containers, service.Spec.Volumes); err == nil {
		if len(containers) > 0 {
			pod.ServiceAttributes.Containers = containers
			volumeMountNames1 = vm
		} else {
			utils.Error.Println("no containers exist")
			return nil, errors.New("no containers exist")
		}

	} else {
		utils.Error.Println(err)
		return nil, err
	}

	if containersList, volumeMounts, err := getCPContainers(service.Spec.InitContainers, service.Spec.Volumes); err == nil {
		if len(containersList) > 0 {
			pod.ServiceAttributes.InitContainers = containersList
		}
		for k, v := range volumeMounts {
			volumeMountNames1[k] = v
		}

	} else {
		utils.Error.Println(err)
		return nil, err
	}

	if vols, err := getCPVolumes(service.Spec.Volumes, volumeMountNames1); err == nil {
		if len(vols) > 0 {
			pod.ServiceAttributes.Volumes = vols
		}

	} else {
		utils.Error.Println(err)
		return nil, err
	}

	if service.Spec.Affinity != nil {
		if affinity, err := getCPAffinity(service.Spec.Affinity); err == nil {
			pod.ServiceAttributes.Affinity = affinity
		} else {
			utils.Error.Println(err)
			return nil, err
		}
	}
	pod.ServiceAttributes.RestartPolicy = meshTypes.RestartPolicy(service.Spec.RestartPolicy)
	return pod, nil
}

func convertToCPDaemonSet(ds interface{}) (*meshTypes.DaemonSetService, error) {
	byteData, _ := json.Marshal(ds)
	service := apps.DaemonSet{}
	json.Unmarshal(byteData, &service)
	daemonSet := new(meshTypes.DaemonSetService)

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
	if service.Labels["version"] != "" {
		daemonSet.Version = service.Labels["version"]
	}

	daemonSet.ServiceType = meshConstants.Kubernetes
	daemonSet.ServiceSubType = meshConstants.DaemonSet
	daemonSet.ServiceAttributes = new(meshTypes.DaemonSetServiceAttribute)
	daemonSet.ServiceAttributes.Labels = make(map[string]string)
	daemonSet.ServiceAttributes.Labels = service.Spec.Template.Labels
	if service.Spec.Selector != nil {
		daemonSet.ServiceAttributes.LabelSelector = new(meshTypes.LabelSelectorObj)
		daemonSet.ServiceAttributes.LabelSelector.MatchLabels = make(map[string]string)
		daemonSet.ServiceAttributes.LabelSelector.MatchLabels = service.Spec.Selector.MatchLabels
	}

	daemonSet.ServiceAttributes.Annotations = make(map[string]string)
	daemonSet.ServiceAttributes.Annotations = service.Spec.Template.Annotations
	daemonSet.ServiceAttributes.NodeSelector = make(map[string]string)
	daemonSet.ServiceAttributes.NodeSelector = service.Spec.Template.Spec.NodeSelector

	//daemonSetUpdateStrategy
	if service.Spec.UpdateStrategy.Type != "" {
		daemonSet.ServiceAttributes.UpdateStrategy = new(meshTypes.DaemonSetUpdateStrategy)
		if service.Spec.UpdateStrategy.Type == apps.OnDeleteDaemonSetStrategyType {
			daemonSet.ServiceAttributes.UpdateStrategy.Type = meshTypes.OnDeleteDaemonSetStrategyType
		} else if service.Spec.UpdateStrategy.Type == apps.RollingUpdateDaemonSetStrategyType {
			daemonSet.ServiceAttributes.UpdateStrategy.Type = meshTypes.RollingUpdateDaemonSetStrategyType
			if service.Spec.UpdateStrategy.RollingUpdate != nil {
				daemonSet.ServiceAttributes.UpdateStrategy.RollingUpdate = new(meshTypes.RollingUpdateDaemonSet)
				daemonSet.ServiceAttributes.UpdateStrategy.RollingUpdate.MaxUnavailable = new(intstr.IntOrString)
				daemonSet.ServiceAttributes.UpdateStrategy.RollingUpdate.MaxUnavailable = service.Spec.UpdateStrategy.RollingUpdate.MaxUnavailable
			}
		}
	}

	//containers
	var volumeMountNames1 = make(map[string]bool)
	if containers, vm, err := getCPContainers(service.Spec.Template.Spec.Containers, service.Spec.Template.Spec.Volumes); err == nil {
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
	if containersList, volumeMounts, err := getCPContainers(service.Spec.Template.Spec.InitContainers, service.Spec.Template.Spec.Volumes); err == nil {
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

func convertToCPStatefulSet(sset interface{}) (*meshTypes.StatefulSetService, error) {

	byteData, _ := json.Marshal(sset)
	service := apps.StatefulSet{}
	json.Unmarshal(byteData, &service)
	statefulSet := new(meshTypes.StatefulSetService)

	if service.Name == "" {
		return nil, errors.New("service name not found")
	} else {
		statefulSet.Name = service.Name
	}
	statefulSet.Name = service.Name
	statefulSet.ServiceType = meshConstants.Kubernetes
	statefulSet.ServiceSubType = meshConstants.StatefulSet
	if service.Labels["version"] != "" {
		statefulSet.Version = service.Labels["version"]
	}

	if service.Namespace == "" {
		statefulSet.Namespace = "default"
	} else {
		statefulSet.Namespace = service.Namespace
	}

	statefulSet.ServiceAttributes = new(meshTypes.StatefulSetServiceAttribute)
	statefulSet.ServiceAttributes.Labels = make(map[string]string)
	statefulSet.ServiceAttributes.Labels = service.Spec.Template.Labels

	statefulSet.ServiceAttributes.Annotations = make(map[string]string)
	statefulSet.ServiceAttributes.Annotations = service.Spec.Template.Annotations
	if service.Spec.Selector != nil {
		statefulSet.ServiceAttributes.LabelSelector = new(meshTypes.LabelSelectorObj)
		statefulSet.ServiceAttributes.LabelSelector.MatchLabels = make(map[string]string)
		statefulSet.ServiceAttributes.LabelSelector.MatchLabels = service.Spec.Selector.MatchLabels
	}
	statefulSet.ServiceAttributes.NodeSelector = make(map[string]string)
	statefulSet.ServiceAttributes.NodeSelector = service.Spec.Template.Spec.NodeSelector

	//replicas
	if service.Spec.Replicas != nil {
		statefulSet.ServiceAttributes.Replicas = service.Spec.Replicas
	} else {
		var a int32 = 1
		statefulSet.ServiceAttributes.Replicas = &a
	}

	if service.Spec.ServiceName != "" {
		statefulSet.ServiceAttributes.ServiceName = service.Spec.ServiceName
	}
	//update strategy
	if service.Spec.UpdateStrategy.Type != "" {
		statefulSet.ServiceAttributes.UpdateStrategy = new(meshTypes.StateFulSetUpdateStrategy)
		if service.Spec.UpdateStrategy.Type == apps.OnDeleteStatefulSetStrategyType {
			statefulSet.ServiceAttributes.UpdateStrategy.Type = meshTypes.OnDeleteStatefulSetStrategyType
		} else if service.Spec.UpdateStrategy.Type == apps.RollingUpdateStatefulSetStrategyType {
			statefulSet.ServiceAttributes.UpdateStrategy.Type = meshTypes.RollingUpdateStatefulSetStrategyType
			if service.Spec.UpdateStrategy.RollingUpdate != nil {
				statefulSet.ServiceAttributes.UpdateStrategy.RollingUpdate = new(meshTypes.RollingUpdateStatefulSetStrategy)
				statefulSet.ServiceAttributes.UpdateStrategy.RollingUpdate.Partition = service.Spec.UpdateStrategy.RollingUpdate.Partition
			}
		}
	}
	//containers
	var volumeMountNames1 = make(map[string]bool)
	if containers, vm, err := getCpStsContainers(service.Spec.Template.Spec.Containers, service.Spec.VolumeClaimTemplates, service.Spec.Template.Spec.Volumes); err == nil {
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
	if containersList, volumeMounts, err := getCPContainers(service.Spec.Template.Spec.InitContainers, service.Spec.Template.Spec.Volumes); err == nil {
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
		//tempVC := new(meshTypes.PersistentVolumeClaimService)
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

func convertToCPJob(job *batch.Job) (*meshTypes.JobService, error) {
	cpJob := new(meshTypes.JobService)
	if job.Name == "" {
		return nil, errors.New("service name not found")
	} else {
		cpJob.Name = job.Name
	}
	cpJob.Name = job.Name
	cpJob.ServiceType = meshConstants.Kubernetes
	cpJob.ServiceSubType = meshConstants.Job
	if job.Namespace == "" {
		cpJob.Namespace = "default"
	} else {
		cpJob.Namespace = job.Namespace
	}
	if job.Labels["version"] != "" {
		cpJob.Version = job.Labels["version"]
	}

	var CpJobAttr = new(meshTypes.JobServiceAttribute)
	CpJobAttr.Labels = make(map[string]string)
	CpJobAttr.Labels = job.Spec.Template.Labels

	CpJobAttr.LabelSelector = new(meshTypes.LabelSelectorObj)
	if job.Spec.Selector != nil && job.Spec.Selector.MatchLabels != nil {
		CpJobAttr.LabelSelector.MatchLabels = make(map[string]string)
		CpJobAttr.LabelSelector.MatchLabels = job.Spec.Selector.MatchLabels

	}
	CpJobAttr.Annotations = make(map[string]string)
	CpJobAttr.Annotations = job.Spec.Template.Annotations

	CpJobAttr.NodeSelector = make(map[string]string)
	CpJobAttr.NodeSelector = job.Spec.Template.Spec.NodeSelector

	if job.Spec.Parallelism != nil {
		var CpJobParallelism = new(meshTypes.Parallelism)
		CpJobParallelism.Value = *job.Spec.Parallelism
		CpJobAttr.Parallelism = CpJobParallelism
	}
	if job.Spec.Completions != nil {
		var CpJobCompletions = new(meshTypes.Completions)
		CpJobCompletions.Value = *job.Spec.Completions
		CpJobAttr.Completions = CpJobCompletions
	}
	if job.Spec.ActiveDeadlineSeconds != nil {
		var CpJobActiveDeadlineSeconds = new(meshTypes.ActiveDeadlineSeconds)
		CpJobActiveDeadlineSeconds.Value = *job.Spec.ActiveDeadlineSeconds
		CpJobAttr.ActiveDeadlineSeconds = CpJobActiveDeadlineSeconds
	}
	if job.Spec.BackoffLimit != nil {
		var CpJobBackOffLimit = new(meshTypes.BackoffLimit)
		CpJobBackOffLimit.Value = *job.Spec.BackoffLimit
		CpJobAttr.BackoffLimit = CpJobBackOffLimit
	}
	if job.Spec.ManualSelector != nil {
		var CpJobManualSelector = new(meshTypes.ManualSelector)
		CpJobManualSelector.Value = *job.Spec.ManualSelector
		CpJobAttr.ManualSelector = CpJobManualSelector
	}

	//containers
	var volumeMountNames1 = make(map[string]bool)
	if containers, vm, err := getCPContainers(job.Spec.Template.Spec.Containers, job.Spec.Template.Spec.Volumes); err == nil {
		if len(containers) > 0 {
			CpJobAttr.Containers = containers
			volumeMountNames1 = vm
		} else {
			return nil, errors.New("no containers exist")
		}

	} else {
		return nil, err
	}

	//init containers
	if containersList, volumeMounts, err := getCPContainers(job.Spec.Template.Spec.InitContainers, job.Spec.Template.Spec.Volumes); err == nil {
		if len(containersList) > 0 {
			CpJobAttr.InitContainers = containersList
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
			CpJobAttr.Volumes = vols
		}

	} else {
		return nil, err
	}

	//affinity
	if job.Spec.Template.Spec.Affinity != nil {
		if affinity, err := getCPAffinity(job.Spec.Template.Spec.Affinity); err == nil {
			CpJobAttr.Affinity = affinity
		} else {
			return nil, err
		}
	}

	cpJob.ServiceAttributes = CpJobAttr
	return cpJob, nil
}

func convertToCPCronJob(job *batchv1.CronJob) (*meshTypes.CronJobService, error) {
	cpJob := new(meshTypes.CronJobService)
	if job.Name == "" {
		return nil, errors.New("service name not found")
	} else {
		cpJob.Name = job.Name
	}
	if job.Labels["version"] != "" {
		cpJob.Version = job.Labels["version"]
	}

	cpJob.ServiceType = meshConstants.Kubernetes
	cpJob.ServiceSubType = meshConstants.CronJob

	if job.Namespace == "" {
		cpJob.Namespace = "default"
	} else {
		cpJob.Namespace = job.Namespace
	}

	cpJob.ServiceAttributes = new(meshTypes.CronJobServiceAttribute)

	cpJob.ServiceAttributes.Labels = make(map[string]string)
	cpJob.ServiceAttributes.Labels = job.Labels
	cpJob.ServiceAttributes.Annotations = make(map[string]string)
	cpJob.ServiceAttributes.Annotations = job.Annotations

	var volumeMountNames1 = make(map[string]bool)
	if containers, vm, err := getCPContainers(job.Spec.JobTemplate.Spec.Template.Spec.Containers, job.Spec.JobTemplate.Spec.Template.Spec.Volumes); err == nil {
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
	if containersList, volumeMounts, err := getCPContainers(job.Spec.JobTemplate.Spec.Template.Spec.InitContainers, job.Spec.JobTemplate.Spec.Template.Spec.Volumes); err == nil {
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
	if vols, err := getCPVolumes(job.Spec.JobTemplate.Spec.Template.Spec.Volumes, volumeMountNames1); err == nil {
		if len(vols) > 0 {
			cpJob.ServiceAttributes.Volumes = vols
		}

	} else {
		return nil, err
	}

	if job.Spec.JobTemplate.Spec.Template.Spec.Affinity != nil {
		if affinity, err := getCPAffinity(job.Spec.JobTemplate.Spec.Template.Spec.Affinity); err == nil {
			cpJob.ServiceAttributes.Affinity = affinity
		} else {
			return nil, err
		}
	}

	if job.Spec.Schedule != "" {
		cpJob.ServiceAttributes.CronJobScheduleString = job.Spec.Schedule
	}
	if job.Spec.StartingDeadlineSeconds != nil {
		cpJob.ServiceAttributes.StartingDeadLineSeconds = &meshTypes.StartingDeadlineSeconds{
			Value: *job.Spec.StartingDeadlineSeconds,
		}
	}

	if job.Spec.FailedJobsHistoryLimit != nil {
		cpJob.ServiceAttributes.FailedJobsHistoryLimit = &meshTypes.FailedJobsHistoryLimit{Value: *job.Spec.FailedJobsHistoryLimit}
	}
	if job.Spec.SuccessfulJobsHistoryLimit != nil {
		cpJob.ServiceAttributes.SuccessfulJobsHistoryLimit = &meshTypes.SuccessfulJobsHistoryLimit{Value: *job.Spec.SuccessfulJobsHistoryLimit}
	}
	if job.Spec.Suspend != nil {
		cpJob.ServiceAttributes.Suspend = &meshTypes.Suspend{Value: *job.Spec.Suspend}
	}
	if job.Spec.ConcurrencyPolicy != "" {
		cpJob.ServiceAttributes.ConcurrencyPolicy = new(meshTypes.ConcurrencyPolicy)
		if job.Spec.ConcurrencyPolicy == batchv1.AllowConcurrent {
			value := meshTypes.ConcurrencyPolicyAllow
			cpJob.ServiceAttributes.ConcurrencyPolicy = &value
		} else if job.Spec.ConcurrencyPolicy == batchv1.ForbidConcurrent {
			value := meshTypes.ConcurrencyPolicyForbid
			cpJob.ServiceAttributes.ConcurrencyPolicy = &value
		} else {
			value := meshTypes.ConcurrencyPolicyReplace
			cpJob.ServiceAttributes.ConcurrencyPolicy = &value
		}
	}

	return cpJob, nil

}

//func getCPJobTemplateSpec(job batchv1.JobTemplateSpec) (*meshTypes.JobServiceAttribute, error) {
//	jobTemplate := new(meshTypes.JobServiceAttribute)
//	jobTemplate.Labels = make(map[string]string)
//	jobTemplate.Labels = job.Labels
//
//	jobTemplate.Annotations = make(map[string]string)
//	jobTemplate.Annotations = job.Spec.Template.Annotations
//	jobTemplate.LabelSelector = new(meshTypes.LabelSelectorObj)
//	jobTemplate.LabelSelector.MatchLabels = make(map[string]string)
//	if job.Spec.Selector != nil {
//		jobTemplate.LabelSelector.MatchLabels = job.Spec.Selector.MatchLabels
//	}
//	jobTemplate.NodeSelector = make(map[string]string)
//	jobTemplate.NodeSelector = job.Spec.Template.Spec.NodeSelector
//
//	var volumeMountNames1 = make(map[string]bool)
//
//	if containers, vm, err := getCPContainers(job.Spec.Template.Spec.Containers); err == nil {
//		if len(containers) > 0 {
//			jobTemplate.Containers = containers
//			volumeMountNames1 = vm
//		} else {
//			return nil, errors.New("no containers exist")
//		}
//
//	} else {
//		return nil, err
//	}
//
//	//init containers
//	if containersList, volumeMounts, err := getCPContainers(job.Spec.Template.Spec.InitContainers); err == nil {
//		if len(containersList) > 0 {
//			jobTemplate.InitContainers = containersList
//		}
//		for k, v := range volumeMounts {
//			volumeMountNames1[k] = v
//		}
//
//	} else {
//		return nil, err
//	}
//
//	if job.Spec.Template.Spec.Affinity != nil {
//		if affinity, err := getCPAffinity(job.Spec.Template.Spec.Affinity); err == nil {
//			jobTemplate.Affinity = affinity
//		} else {
//			return nil, err
//		}
//	}
//
//	//volumes
//	if vols, err := getCPVolumes(job.Spec.Template.Spec.Volumes, volumeMountNames1); err == nil {
//		if len(vols) > 0 {
//			jobTemplate.Volumes = vols
//		}
//
//	} else {
//		return nil, err
//	}
//	return jobTemplate, nil
//}

func convertToCPPersistentVolumeClaim(pvc *v1.PersistentVolumeClaim) (*meshTypes.PersistentVolumeClaimService, error) {
	persistentVolume := new(meshTypes.PersistentVolumeClaimService)
	persistentVolume.Name = pvc.Name
	persistentVolume.ServiceType = meshConstants.Kubernetes
	persistentVolume.ServiceSubType = meshConstants.PVC
	if pvc.Labels["version"] != "" {
		persistentVolume.Version = pvc.Labels["version"]
	}

	if pvc.Namespace != "" {
		persistentVolume.Namespace = pvc.Namespace
	} else {
		persistentVolume.Namespace = "default"
	}

	persistentVolume.ServiceAttributes = new(meshTypes.PersistentVolumeClaimServiceAttribute)
	if pvc.Spec.StorageClassName != nil {
		persistentVolume.ServiceAttributes.StorageClassName = *pvc.Spec.StorageClassName
	}
	if pvc.Spec.VolumeMode != nil {
		persistentVolume.ServiceAttributes.VolumeMode = (*meshTypes.PersistentVolumeMode)(pvc.Spec.VolumeMode)
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
		persistentVolume.ServiceAttributes.DataSource = new(meshTypes.TypedLocalObjectReference)
		persistentVolume.ServiceAttributes.DataSource.Name = pvc.Spec.DataSource.Name
		persistentVolume.ServiceAttributes.DataSource.Kind = pvc.Spec.DataSource.Kind
		if pvc.Spec.DataSource.APIGroup != nil {
			persistentVolume.ServiceAttributes.DataSource.APIGroup = pvc.Spec.DataSource.APIGroup
		}

	}
	for _, each := range pvc.Spec.AccessModes {
		var am meshTypes.AccessMode
		if each == v1.ReadWriteOnce {
			am = meshTypes.AccessModeReadWriteOnce
		} else if each == v1.ReadOnlyMany {
			am = meshTypes.AccessModeReadOnlyMany
		} else if each == v1.ReadWriteMany {
			am = meshTypes.AccessModeReadWriteMany
		} else {
			continue
		}

		persistentVolume.ServiceAttributes.AccessMode = append(persistentVolume.ServiceAttributes.AccessMode, am)
	}
	return persistentVolume, nil
}

func convertToCPPersistentVolume(pv *v1.PersistentVolume) (*meshTypes.PersistentVolumeService, error) {
	persistentVolume := new(meshTypes.PersistentVolumeService)
	persistentVolume.Name = pv.Name
	persistentVolume.ServiceType = meshConstants.Kubernetes
	persistentVolume.ServiceSubType = meshConstants.PV
	if pv.Labels["version"] != "" {
		persistentVolume.Version = pv.Labels["version"]
	}
	persistentVolume.ServiceAttributes = new(meshTypes.PersistentVolumeServiceAttribute)
	persistentVolume.ServiceAttributes.ReclaimPolicy = meshTypes.ReclaimPolicy(pv.Spec.PersistentVolumeReclaimPolicy)
	qu := pv.Spec.Capacity[v1.ResourceStorage]
	persistentVolume.ServiceAttributes.Capacity = qu.String()
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
		persistentVolume.ServiceAttributes.VolumeMode = (*meshTypes.PersistentVolumeMode)(pv.Spec.VolumeMode)
	}

	if pv.Spec.NodeAffinity != nil {
		persistentVolume.ServiceAttributes.NodeAffinity = new(meshTypes.VolumeNodeAffinity)
		if ns, err := getCPNodeSelector(pv.Spec.NodeAffinity.Required); err != nil {
			return nil, err
		} else {
			persistentVolume.ServiceAttributes.NodeAffinity.Required = *ns
		}

	}

	for _, each := range pv.Spec.AccessModes {
		var am meshTypes.AccessMode
		if each == v1.ReadWriteOnce {
			am = meshTypes.AccessModeReadWriteOnce
		} else if each == v1.ReadOnlyMany {
			am = meshTypes.AccessModeReadOnlyMany
		} else if each == v1.ReadWriteMany {
			am = meshTypes.AccessModeReadWriteMany
		} else {
			continue
		}

		persistentVolume.ServiceAttributes.AccessMode = append(persistentVolume.ServiceAttributes.AccessMode, am)
	}
	if pv.Spec.PersistentVolumeSource.AWSElasticBlockStore != nil {
		persistentVolume.ServiceAttributes.PersistentVolumeSource = new(meshTypes.PersistentVolumeSource)
		persistentVolume.ServiceAttributes.PersistentVolumeSource.AWSEBS = new(meshTypes.AWSEBS)
		persistentVolume.ServiceAttributes.PersistentVolumeSource.AWSEBS.VolumeId = pv.Spec.AWSElasticBlockStore.VolumeID
		persistentVolume.ServiceAttributes.PersistentVolumeSource.AWSEBS.ReadOnly = pv.Spec.AWSElasticBlockStore.ReadOnly
		persistentVolume.ServiceAttributes.PersistentVolumeSource.AWSEBS.Filesystem = pv.Spec.AWSElasticBlockStore.FSType
		persistentVolume.ServiceAttributes.PersistentVolumeSource.AWSEBS.Partition = int(pv.Spec.AWSElasticBlockStore.Partition)
	} else if pv.Spec.PersistentVolumeSource.GCEPersistentDisk != nil {
		persistentVolume.ServiceAttributes.PersistentVolumeSource = new(meshTypes.PersistentVolumeSource)
		persistentVolume.ServiceAttributes.PersistentVolumeSource.GCPPD = new(meshTypes.GCPPD)
		persistentVolume.ServiceAttributes.PersistentVolumeSource.GCPPD.PdName = pv.Spec.GCEPersistentDisk.PDName
		persistentVolume.ServiceAttributes.PersistentVolumeSource.GCPPD.ReadOnly = pv.Spec.GCEPersistentDisk.ReadOnly
		persistentVolume.ServiceAttributes.PersistentVolumeSource.GCPPD.Filesystem = pv.Spec.GCEPersistentDisk.FSType
		persistentVolume.ServiceAttributes.PersistentVolumeSource.GCPPD.Partition = int(pv.Spec.GCEPersistentDisk.Partition)
	} else if pv.Spec.PersistentVolumeSource.AzureFile != nil {
		persistentVolume.ServiceAttributes.PersistentVolumeSource = new(meshTypes.PersistentVolumeSource)
		persistentVolume.ServiceAttributes.PersistentVolumeSource.AzureFile = new(meshTypes.AzureFile)
		persistentVolume.ServiceAttributes.PersistentVolumeSource.AzureFile.ReadOnly = pv.Spec.AzureFile.ReadOnly
		persistentVolume.ServiceAttributes.PersistentVolumeSource.AzureFile.ShareName = pv.Spec.AzureFile.ShareName
		persistentVolume.ServiceAttributes.PersistentVolumeSource.AzureFile.SecretName = pv.Spec.AzureFile.SecretName
		persistentVolume.ServiceAttributes.PersistentVolumeSource.AzureFile.SecretNamespace = *pv.Spec.AzureFile.SecretNamespace
	} else if pv.Spec.PersistentVolumeSource.AzureDisk != nil {
		persistentVolume.ServiceAttributes.PersistentVolumeSource = new(meshTypes.PersistentVolumeSource)
		persistentVolume.ServiceAttributes.PersistentVolumeSource.AzureDisk = new(meshTypes.AzureDisk)
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
				persistentVolume.ServiceAttributes.PersistentVolumeSource.AzureDisk.CachingMode = meshTypes.AzureDataDiskCachingNone
			} else if *pv.Spec.AzureDisk.CachingMode == v1.AzureDataDiskCachingReadOnly {
				persistentVolume.ServiceAttributes.PersistentVolumeSource.AzureDisk.CachingMode = meshTypes.AzureDataDiskCachingReadOnly
			} else if *pv.Spec.AzureDisk.CachingMode == v1.AzureDataDiskCachingReadWrite {
				persistentVolume.ServiceAttributes.PersistentVolumeSource.AzureDisk.CachingMode = meshTypes.AzureDataDiskCachingReadWrite
			}

		}
		if pv.Spec.AzureDisk.Kind != nil {
			if *pv.Spec.AzureDisk.Kind == v1.AzureDedicatedBlobDisk {
				persistentVolume.ServiceAttributes.PersistentVolumeSource.AzureDisk.Kind = meshTypes.AzureDedicatedBlobDisk
			} else if *pv.Spec.AzureDisk.Kind == v1.AzureSharedBlobDisk {
				persistentVolume.ServiceAttributes.PersistentVolumeSource.AzureDisk.Kind = meshTypes.AzureSharedBlobDisk
			} else if *pv.Spec.AzureDisk.Kind == v1.AzureManagedDisk {
				persistentVolume.ServiceAttributes.PersistentVolumeSource.AzureDisk.Kind = meshTypes.AzureManagedDisk
			}
		}

	}

	return persistentVolume, nil
}

func convertToCPStorageClass(sc *storage.StorageClass) (*meshTypes.StorageClassService, error) {
	storageClass := new(meshTypes.StorageClassService)
	storageClass.Name = sc.Name
	storageClass.ServiceType = meshConstants.Kubernetes
	storageClass.ServiceSubType = meshConstants.StorageClass
	storageClass.ServiceAttributes = new(meshTypes.StorageClassServiceAttribute)
	if sc.Labels["version"] != "" {
		storageClass.Version = sc.Labels["version"]
	}
	if sc.ReclaimPolicy != nil {
		storageClass.ServiceAttributes.ReclaimPolicy = meshTypes.ReclaimPolicy(*sc.ReclaimPolicy)
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
		aT := meshTypes.TopologySelectorTerm{}
		for _, each2 := range each.MatchLabelExpressions {
			tr := meshTypes.TopologySelectorLabelRequirement{}
			tr.Key = each2.Key
			for _, value := range each2.Values {
				tr.Values = append(tr.Values, value)
			}
			aT.MatchLabelExpressions = append(aT.MatchLabelExpressions, tr)

		}
		storageClass.ServiceAttributes.AllowedTopologies = append(storageClass.ServiceAttributes.AllowedTopologies, aT)
	}

	if sc.VolumeBindingMode != nil {
		storageClass.ServiceAttributes.BindingMod = meshTypes.VolumeBindingMode(*sc.VolumeBindingMode)
	}
	storageClass.ServiceAttributes.Provisioner = sc.Provisioner
	if len(sc.Parameters) > 0 {
		storageClass.ServiceAttributes.Parameters = make(map[string]string)
	}

	for key, value := range sc.Parameters {
		storageClass.ServiceAttributes.Parameters[key] = value
	}

	return storageClass, nil
}

func ConvertToCPSecret(cm *v1.Secret) (*meshTypes.Secret, error) {
	var secret = new(meshTypes.Secret)
	secret.Name = cm.Name
	secret.Namespace = cm.Namespace
	if vr := cm.Labels["version"]; vr != "" {
		secret.Version = vr
	}
	secret.ServiceType = meshConstants.Kubernetes
	secret.ServiceSubType = meshConstants.Secret
	secret.ServiceAttributes = new(meshTypes.SecretServiceAttribute)
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

func ConvertToCPHPA(hpa *autoScalar.HorizontalPodAutoscaler) (*meshTypes.HPA, error) {
	var horizntalPodAutoscalar = new(meshTypes.HPA)
	horizntalPodAutoscalar.Name = hpa.Name
	horizntalPodAutoscalar.Namespace = hpa.Namespace
	horizntalPodAutoscalar.ServiceType = meshConstants.Kubernetes
	if vr := hpa.Labels["version"]; vr != "" {
		horizntalPodAutoscalar.Version = vr
	}
	horizntalPodAutoscalar.ServiceSubType = meshConstants.Hpa

	horizntalPodAutoscalar.ServiceAttributes.MaxReplicas = int(hpa.Spec.MaxReplicas)
	if hpa.Spec.MinReplicas != nil {
		horizntalPodAutoscalar.ServiceAttributes.MinReplicas = int(*hpa.Spec.MinReplicas)
	}
	horizntalPodAutoscalar.ServiceAttributes.CrossObjectVersion.Name = hpa.Spec.ScaleTargetRef.Name
	horizntalPodAutoscalar.ServiceAttributes.CrossObjectVersion.Type = hpa.Spec.ScaleTargetRef.Kind
	horizntalPodAutoscalar.ServiceAttributes.CrossObjectVersion.Version = hpa.Spec.ScaleTargetRef.APIVersion

	if hpa.Spec.TargetCPUUtilizationPercentage != nil {
		horizntalPodAutoscalar.ServiceAttributes.TargetCpuUtilization = hpa.Spec.TargetCPUUtilizationPercentage
	}

	/*var metrics []meshTypes.MetricValue
	for _, metric := range hpa.Spec.Metrics {
		cpMetric := meshTypes.MetricValue{}
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
	horizntalPodAutoscalar.ServiceAttributes.MetricValues = metrics*/

	return horizntalPodAutoscalar, nil
}

func ConvertToCPRole(k8ROle *rbac.Role) (*meshTypes.Role, error) {
	var role = new(meshTypes.Role)
	role.Name = k8ROle.Name
	role.Namespace = k8ROle.Namespace
	role.ServiceType = meshConstants.Kubernetes
	role.ServiceSubType = meshConstants.Role
	if vr := k8ROle.Labels["version"]; vr != "" {
		role.Version = vr
	}
	for _, each := range k8ROle.Rules {
		rolePolicy := meshTypes.Rule{}
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
			rolePolicy.Resources = append(rolePolicy.Resources, resource)
		}
		role.ServiceAttributes.Rules = append(role.ServiceAttributes.Rules, rolePolicy)
	}
	return role, nil
}

func ConvertToCPRoleBinding(k8sRoleBinding *rbac.RoleBinding) (*meshTypes.RoleBinding, error) {
	var rb = new(meshTypes.RoleBinding)
	rb.Name = k8sRoleBinding.Name
	rb.ServiceType = meshConstants.Kubernetes
	rb.ServiceSubType = meshConstants.RoleBinding
	if vr := k8sRoleBinding.Labels["version"]; vr != "" {
		rb.Version = vr
	}
	rb.Namespace = k8sRoleBinding.Namespace
	for _, each := range k8sRoleBinding.Subjects {
		var subject = meshTypes.Subject{}
		subject.Name = each.Name
		if each.Kind == "User" || each.Kind == "Group" {
			subject.Kind = each.Kind
		} else if each.Kind == constants.ServiceAccount.String() {
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

func ConvertToCPClusterRoleBinding(k8sClusterRoleBinding *rbac.ClusterRoleBinding) (*meshTypes.ClusterRoleBinding, error) {
	var crb = new(meshTypes.ClusterRoleBinding)
	crb.Name = k8sClusterRoleBinding.Name
	crb.ServiceType = meshConstants.Kubernetes
	crb.ServiceSubType = meshConstants.ClusterRoleBinding
	crb.ServiceAttributes.RoleRef.Name = k8sClusterRoleBinding.RoleRef.Name
	if k8sClusterRoleBinding.RoleRef.Kind == constants.ClusterRole.String() {
		crb.ServiceAttributes.RoleRef.Kind = meshConstants.ClusterRole.String()
	} else {
		crb.ServiceAttributes.RoleRef.Kind = meshConstants.Role.String()
	}
	if vr := k8sClusterRoleBinding.Labels["version"]; vr != "" {
		crb.Version = vr
	}
	crb.Namespace = k8sClusterRoleBinding.Namespace
	for _, each := range k8sClusterRoleBinding.Subjects {
		var subject = meshTypes.Subject{}
		subject.Name = each.Name
		if each.Kind == "User" || each.Kind == "Group" {
			subject.Kind = each.Kind
		} else if each.Kind == constants.ServiceAccount.String() {
			subject.Kind = each.Kind
			subject.Namespace = each.Namespace
		} else {
			return nil, errors.New("invalid subject kind" + each.Name + each.Kind)
		}
		crb.ServiceAttributes.Subjects = append(crb.ServiceAttributes.Subjects, subject)
	}
	return crb, nil
}

func ConvertToCPClusterRole(k8ROle *rbac.ClusterRole) (*meshTypes.ClusterRole, error) {
	var role = new(meshTypes.ClusterRole)
	role.Name = k8ROle.Name
	role.ServiceType = meshConstants.Kubernetes
	role.ServiceSubType = meshConstants.ClusterRole
	if vr := k8ROle.Labels["version"]; vr != "" {
		role.Version = vr
	}
	role.Namespace = k8ROle.Namespace
	for _, each := range k8ROle.Rules {
		rolePolicy := meshTypes.Rules{}
		for _, apigroup := range each.APIGroups {
			rolePolicy.ApiGroup = append(rolePolicy.ApiGroup, apigroup)
		}

		for _, verb := range each.Verbs {
			rolePolicy.Verbs = append(rolePolicy.Verbs, verb)
		}

		for _, resource := range each.Resources {
			rolePolicy.ResourceName = append(rolePolicy.ResourceName, resource)
		}
		//for _, resource := range each.ResourceNames {
		//	rolePolicy.ResourceName = append(rolePolicy.ResourceName, resource)
		//}
		for _, resourceUrls := range each.NonResourceURLs {
			rolePolicy.NonResourceUrls = append(rolePolicy.NonResourceUrls, resourceUrls)

		}
		role.ServiceAttributes.Rules = append(role.ServiceAttributes.Rules, rolePolicy)
	}
	return role, nil
}

func ConvertToCPConfigMap(cm *v1.ConfigMap) (*meshTypes.ConfigMap, error) {
	var configMap = new(meshTypes.ConfigMap)
	configMap.Name = cm.Name
	configMap.Namespace = cm.Namespace
	if vr := cm.Labels["version"]; vr != "" {
		configMap.Version = vr
	}
	configMap.ServiceType = meshConstants.Kubernetes
	configMap.ServiceSubType = meshConstants.ConfigMap
	configMap.ServiceAttributes = new(meshTypes.ConfigMapServiceAttribute)
	if len(cm.Data) > 0 {
		configMap.ServiceAttributes.Data = make(map[string]string)
	}
	for key, value := range cm.Data {
		configMap.ServiceAttributes.Data[key] = value
	}
	return configMap, nil
}

func convertToCPKubernetesService(svc *v1.Service) (*meshTypes.Service, error) {
	var service = new(meshTypes.Service)
	service.Name = svc.Name
	service.Namespace = svc.Namespace
	if vr := svc.Labels["version"]; vr != "" {
		service.Version = vr
	}
	service.ServiceType = meshConstants.Kubernetes
	service.ServiceSubType = meshConstants.KubernetesService
	if svc.Spec.Type != "" {
		service.ServiceAttributes.Type = string(svc.Spec.Type)
	} else {
		//set default type of k8s service
		service.ServiceAttributes.Type = "ClusterIP"
	}

	if svc.Spec.ExternalTrafficPolicy != "" {
		if svc.Spec.ExternalTrafficPolicy == v1.ServiceExternalTrafficPolicyTypeLocal && !(svc.Spec.Type == v1.ServiceTypeLoadBalancer || svc.Spec.Type == v1.ServiceTypeNodePort) {
			return nil, errors.New("for external traffic policy local service type should be LoadBalancer or NodePort")
		}
		service.ServiceAttributes.ExternalTrafficPolicy = string(svc.Spec.ExternalTrafficPolicy)
	} else {
		//set default External traffic policy of k8s service
		service.ServiceAttributes.ExternalTrafficPolicy = "Cluster"
	}
	if len(svc.Spec.Selector) > 0 {
		service.ServiceAttributes.Selector = make(map[string]string)
	}
	for key, value := range svc.Spec.Selector {
		service.ServiceAttributes.Selector[key] = value
	}

	for _, each := range svc.Spec.Ports {
		cpPort := meshTypes.KubePort{}
		if each.Name != "" {
			cpPort.Name = each.Name
		}
		if each.Port != 0 {
			cpPort.Port = each.Port
		}

		if !(svc.Spec.ClusterIP == "None") {
			if each.TargetPort.Type == intstr.String {
				cpPort.TargetPort.PortName = each.TargetPort.StrVal
			} else if each.TargetPort.Type == intstr.Int {
				cpPort.TargetPort.PortNumber = each.TargetPort.IntVal
			}
		} else {
			service.ServiceAttributes.ClusterIP = "None"
		}
		if each.Protocol != "" {
			cpPort.Protocol = string(each.Protocol)
		} else {
			cpPort.Protocol = "TCP"
		}
		if each.NodePort != 0 {
			cpPort.NodePort = each.NodePort
		}

		service.ServiceAttributes.Ports = append(service.ServiceAttributes.Ports, cpPort)
	}

	return service, nil
}

func convertToCPGateWayStruct(gw *v1alpha3.Gateway) (*meshTypes.GatewayService, error) {
	return nil, nil
}

func convertToCPVSStruct(gw *v1alpha3.VirtualService) (*meshTypes.VirtualService, error) {
	return nil, nil
}

func convertToCPDRStruct(gw *v1alpha3.DestinationRule) (*meshTypes.DestinationRules, error) {
	return nil, nil
}

func convertToCPServiceAccount(sa *v1.ServiceAccount) (*meshTypes.ServiceAccount, error) {
	var kube = new(meshTypes.ServiceAccount)
	kube.ServiceSubType = meshConstants.ServiceAccount
	kube.ServiceType = meshConstants.Kubernetes
	kube.Name = sa.Name
	kube.Namespace = sa.Namespace
	if vr := sa.Labels["version"]; vr != "" {
		kube.Version = vr
	}
	var CpSaAttr = new(meshTypes.ServiceAccountAttribute)
	for _, value := range sa.Secrets {
		CpSaAttr.Secrets = append(CpSaAttr.Secrets, value.Name)
	}
	for _, value := range sa.ImagePullSecrets {
		CpSaAttr.ImagePullSecretsName = append(CpSaAttr.ImagePullSecretsName, value.Name)
	}
	kube.ServiceAttributes = CpSaAttr
	return kube, nil

}

func getCPNodeSelector(nodeSelector *v1.NodeSelector) (*meshTypes.NodeSelector, error) {
	var temp *meshTypes.NodeSelector
	if nodeSelector != nil {
		temp = new(meshTypes.NodeSelector)
		var nodeSelectorTerms []meshTypes.NodeSelectorTerm
		for _, nodeSelectorTerm := range nodeSelector.NodeSelectorTerms {
			var tempMatchExpressions []meshTypes.NodeSelectorRequirement
			var tempMatchFields []meshTypes.NodeSelectorRequirement
			tempNodeSelectorTerm := meshTypes.NodeSelectorTerm{}
			for _, matchExpression := range nodeSelectorTerm.MatchExpressions {
				tempMatchExpression := meshTypes.NodeSelectorRequirement{}
				tempMatchExpression.Key = matchExpression.Key
				tempMatchExpression.Values = matchExpression.Values
				if matchExpression.Operator == v1.NodeSelectorOpIn {
					tempMatchExpression.Operator = meshTypes.NodeSelectorOpIn
				} else if matchExpression.Operator == v1.NodeSelectorOpNotIn {
					tempMatchExpression.Operator = meshTypes.NodeSelectorOpNotIn
				} else if matchExpression.Operator == v1.NodeSelectorOpExists {
					tempMatchExpression.Operator = meshTypes.NodeSelectorOpExists
				} else if matchExpression.Operator == v1.NodeSelectorOpDoesNotExist {
					tempMatchExpression.Operator = meshTypes.NodeSelectorOpDoesNotExists
				} else if matchExpression.Operator == v1.NodeSelectorOpGt {
					tempMatchExpression.Operator = meshTypes.NodeSelectorOpGt
				} else if matchExpression.Operator == v1.NodeSelectorOpLt {
					tempMatchExpression.Operator = meshTypes.NodeSelectorOpLt
				}

				tempMatchExpressions = append(tempMatchExpressions, tempMatchExpression)
			}
			for _, matchField := range nodeSelectorTerm.MatchFields {
				tempMatchField := meshTypes.NodeSelectorRequirement{}
				tempMatchField.Key = matchField.Key
				tempMatchField.Values = matchField.Values
				if matchField.Operator == v1.NodeSelectorOpIn {
					tempMatchField.Operator = meshTypes.NodeSelectorOpIn
				} else if matchField.Operator == v1.NodeSelectorOpNotIn {
					tempMatchField.Operator = meshTypes.NodeSelectorOpNotIn
				} else if matchField.Operator == v1.NodeSelectorOpExists {
					tempMatchField.Operator = meshTypes.NodeSelectorOpExists
				} else if matchField.Operator == v1.NodeSelectorOpDoesNotExist {
					tempMatchField.Operator = meshTypes.NodeSelectorOpDoesNotExists
				} else if matchField.Operator == v1.NodeSelectorOpGt {
					tempMatchField.Operator = meshTypes.NodeSelectorOpGt
				} else if matchField.Operator == v1.NodeSelectorOpLt {
					tempMatchField.Operator = meshTypes.NodeSelectorOpLt
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

func getCpStsContainers(conts []v1.Container, PVCs []v1.PersistentVolumeClaim, vol []v1.Volume) ([]*meshTypes.ContainerAttribute, map[string]bool, error) {
	volumeMountNames := make(map[string]bool)
	var containers []*meshTypes.ContainerAttribute
	persistentVolumesClaims := make(map[string]v1.PersistentVolumeClaim)
	volumes := make(map[string]v1.Volume)
	for _, each := range PVCs {
		persistentVolumesClaims[each.Name] = each
	}
	for _, each := range vol {
		volumes[each.Name] = each
	}
	for _, container := range conts {
		containerTemp := meshTypes.ContainerAttribute{}

		if container.ReadinessProbe != nil {
			if rp, err := getCPProbe(container.ReadinessProbe, container.Ports); err == nil {
				containerTemp.ReadinessProbe = rp
			} else {
				utils.Info.Println(err)
				return nil, nil, err
			}
		}

		if container.LivenessProbe != nil {
			if lp, err := getCPProbe(container.LivenessProbe, container.Ports); err == nil {
				containerTemp.LivenessProbe = lp
			} else {
				utils.Info.Println(err)
				return nil, nil, err
			}
		}

		if err := putCPCommandAndArguments(&containerTemp, container.Command, container.Args); err != nil {
			utils.Info.Println(err)
			return nil, nil, err
		}

		if err := putCPResource(&containerTemp, container.Resources.Limits, true); err != nil {
			utils.Info.Println(err)
			return nil, nil, err
		}

		if err := putCPResource(&containerTemp, container.Resources.Requests, false); err != nil {
			utils.Info.Println(err)
			return nil, nil, err
		}

		if container.SecurityContext != nil {
			if context, err := getCPSecurityContext(container.SecurityContext); err == nil {
				containerTemp.SecurityContext = context
			} else {
				utils.Info.Println(err)
				return nil, nil, err
			}
		}
		imgInfo := strings.Split(container.Image, ":")
		if len(imgInfo) == 2 {
			containerTemp.ImageName = imgInfo[0]
			if imgInfo[1] != "" {
				containerTemp.Tag = imgInfo[1]
			}

		} else {
			containerTemp.ImageName = container.Image
		}

		var volumeMounts []meshTypes.VolumeMount
		for _, volumeMount := range container.VolumeMounts {
			volumeMountNames[volumeMount.Name] = true
			temp := meshTypes.VolumeMount{}
			temp.Name = volumeMount.Name
			temp.MountPath = volumeMount.MountPath
			temp.SubPath = volumeMount.SubPath
			temp.SubPathExpr = volumeMount.SubPathExpr
			_, foundInPVCS := persistentVolumesClaims[volumeMount.Name]
			_, foundInVolumes := volumes[volumeMount.Name]
			if !foundInPVCS && !foundInVolumes {
				return nil, nil, errors.New(volumeMount.Name + " is not present in pod volume")
			}

			if volumeMount.MountPropagation != nil {
				if *volumeMount.MountPropagation == v1.MountPropagationNone {
					none := meshTypes.MountPropagationNone
					temp.MountPropagation = &none
				} else if *volumeMount.MountPropagation == v1.MountPropagationBidirectional {
					bi := meshTypes.MountPropagationBidirectional
					temp.MountPropagation = &bi
				} else if *volumeMount.MountPropagation == v1.MountPropagationHostToContainer {
					cont := meshTypes.MountPropagationHostToContainer
					temp.MountPropagation = &cont
				}

			}
			volumeMounts = append(volumeMounts, temp)

		}

		var ports []meshTypes.ContainerPort
		for _, port := range container.Ports {
			temp := meshTypes.ContainerPort{}
			if port.ContainerPort == 0 && port.HostPort != 0 {
				port.ContainerPort = port.HostPort
			}
			if port.Name != "" {
				temp.Name = port.Name
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
			ports = append(ports, temp)
		}

		var environmentVariables []meshTypes.EnvironmentVariable
		for _, envVariable := range container.Env {
			tempEnvVariable := meshTypes.EnvironmentVariable{}
			if envVariable.ValueFrom != nil {
				if envVariable.ValueFrom.ConfigMapKeyRef != nil {
					tempEnvVariable.Key = envVariable.Name
					tempEnvVariable.Value = "{{" + envVariable.ValueFrom.ConfigMapKeyRef.Name + ":" + envVariable.ValueFrom.ConfigMapKeyRef.Key + "}}"
					tempEnvVariable.Type = string(meshConstants.ConfigMap)
					tempEnvVariable.Dynamic = true
				} else if envVariable.ValueFrom.SecretKeyRef != nil {
					tempEnvVariable.Key = envVariable.Name
					tempEnvVariable.Value = "{{" + envVariable.ValueFrom.SecretKeyRef.Name + ":" + envVariable.ValueFrom.SecretKeyRef.Key + "}}"
					tempEnvVariable.Type = string(meshConstants.Secret)
					tempEnvVariable.Dynamic = true
				} else if envVariable.ValueFrom.FieldRef != nil {
					tempEnvVariable.Key = envVariable.Name
					tempEnvVariable.Value = envVariable.ValueFrom.FieldRef.FieldPath
					tempEnvVariable.Type = "FieldRef"
				}
				environmentVariables = append(environmentVariables, tempEnvVariable)
			} else {
				tempEnvVariable.Key = envVariable.Name
				tempEnvVariable.Value = envVariable.Value
				environmentVariables = append(environmentVariables, tempEnvVariable)
			}

		}

		containerTemp.Ports = ports
		containerTemp.EnvironmentVariables = environmentVariables
		containerTemp.VolumeMounts = volumeMounts

		containers = append(containers, &containerTemp)
	}
	return containers, volumeMountNames, nil
}

func getCPContainers(conts []v1.Container, volume []v1.Volume) ([]*meshTypes.ContainerAttribute, map[string]bool, error) {
	volumeMountNames := make(map[string]bool)
	var containers []*meshTypes.ContainerAttribute
	volumes := make(map[string]v1.Volume)
	for _, each := range volume {
		volumes[each.Name] = each
	}
	for _, container := range conts {
		containerTemp := meshTypes.ContainerAttribute{}
		containerTemp.ContainerName = container.Name
		if container.ReadinessProbe != nil {
			if rp, err := getCPProbe(container.ReadinessProbe, container.Ports); err == nil {
				containerTemp.ReadinessProbe = rp
			} else {
				utils.Info.Println(err)
				return nil, nil, err
			}
		}

		if container.LivenessProbe != nil {
			if lp, err := getCPProbe(container.LivenessProbe, container.Ports); err == nil {
				containerTemp.LivenessProbe = lp
			} else {
				utils.Info.Println(err)
				return nil, nil, err
			}
		}

		if err := putCPCommandAndArguments(&containerTemp, container.Command, container.Args); err != nil {
			utils.Info.Println(err)
			return nil, nil, err
		}

		if err := putCPResource(&containerTemp, container.Resources.Limits, true); err != nil {
			utils.Info.Println(err)
			return nil, nil, err
		}

		if err := putCPResource(&containerTemp, container.Resources.Requests, false); err != nil {
			utils.Info.Println(err)
			return nil, nil, err
		}

		if container.SecurityContext != nil {
			if context, err := getCPSecurityContext(container.SecurityContext); err == nil {
				containerTemp.SecurityContext = context
			} else {
				utils.Info.Println(err)
				return nil, nil, err
			}
		}
		imgInfo := strings.Split(container.Image, ":")
		if len(imgInfo) == 2 {
			containerTemp.ImageName = imgInfo[0]
			if imgInfo[1] != "" {
				containerTemp.Tag = imgInfo[1]
			}

		} else {
			containerTemp.ImageName = container.Image
		}

		var volumeMounts []meshTypes.VolumeMount
		for _, volumeMount := range container.VolumeMounts {
			volumeMountNames[volumeMount.Name] = true
			temp := meshTypes.VolumeMount{}
			temp.Name = volumeMount.Name
			temp.MountPath = volumeMount.MountPath
			temp.SubPath = volumeMount.SubPath
			temp.SubPathExpr = volumeMount.SubPathExpr
			tempVol, ok := volumes[volumeMount.Name]
			if !ok {
				return nil, nil, errors.New(volumeMount.Name + " is not present in pod volume")
			}
			if tempVol.ConfigMap != nil {
				temp.ConfigMap = new(meshTypes.ConfigMapVolumeMount)
				temp.ConfigMap.ConfigMapName = tempVol.ConfigMap.Name
				if tempVol.ConfigMap.Optional != nil {
					temp.ConfigMap.Optional = *tempVol.ConfigMap.Optional
				}
				for _, each := range tempVol.ConfigMap.Items {
					temp.ConfigMap.Items = append(temp.ConfigMap.Items, meshTypes.KeyItems{
						Key:  each.Key,
						Path: each.Path,
					})
				}
			} else if tempVol.Secret != nil {
				temp.Secret = new(meshTypes.ConfigMapVolumeMount)
				temp.Secret.ConfigMapName = tempVol.Secret.SecretName
				if tempVol.Secret.Optional != nil {
					temp.Secret.Optional = *tempVol.Secret.Optional
				}
				for _, each := range tempVol.Secret.Items {
					temp.Secret.Items = append(temp.Secret.Items, meshTypes.KeyItems{
						Key:  each.Key,
						Path: each.Path,
					})
				}
			} else {
				if tempVol.PersistentVolumeClaim != nil {
					temp.PvcSvcName = tempVol.PersistentVolumeClaim.ClaimName
				}
				if tempVol.EmptyDir != nil {
					temp.EmptyDir = new(meshTypes.EmptyDirVolumeMount)
					temp.EmptyDir.EmptyDirName = tempVol.Name
					temp.ServiceSubType = "emptyDir"
					temp.Name = tempVol.Name

				}
				if tempVol.HostPath != nil {
					temp.HostPath = new(meshTypes.HostPathVolumeMount)
					temp.HostPath.HostPathName = tempVol.Name
					temp.HostPath.Path = tempVol.HostPath.Path
					temp.ServiceSubType = "hostpath"
					temp.Name = tempVol.Name
				}
			}
			if volumeMount.MountPropagation != nil {
				if *volumeMount.MountPropagation == v1.MountPropagationNone {
					none := meshTypes.MountPropagationNone
					temp.MountPropagation = &none
				} else if *volumeMount.MountPropagation == v1.MountPropagationBidirectional {
					bi := meshTypes.MountPropagationBidirectional
					temp.MountPropagation = &bi
				} else if *volumeMount.MountPropagation == v1.MountPropagationHostToContainer {
					cont := meshTypes.MountPropagationHostToContainer
					temp.MountPropagation = &cont
				}

			}
			volumeMounts = append(volumeMounts, temp)

		}

		var ports []meshTypes.ContainerPort
		for _, port := range container.Ports {
			temp := meshTypes.ContainerPort{}
			temp.Name = port.Name
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
			ports = append(ports, temp)
		}

		var environmentVariables []meshTypes.EnvironmentVariable
		for _, envVariable := range container.Env {
			tempEnvVariable := meshTypes.EnvironmentVariable{}
			if envVariable.ValueFrom != nil {
				if envVariable.ValueFrom.ConfigMapKeyRef != nil {
					tempEnvVariable.Key = envVariable.Name
					tempEnvVariable.Value = "{{" + envVariable.ValueFrom.ConfigMapKeyRef.Name + ":" + envVariable.ValueFrom.ConfigMapKeyRef.Key + "}}"
					tempEnvVariable.Type = string(meshConstants.ConfigMap)
					tempEnvVariable.Dynamic = true
				} else if envVariable.ValueFrom.SecretKeyRef != nil {
					tempEnvVariable.Key = envVariable.Name
					tempEnvVariable.Value = "{{" + envVariable.ValueFrom.SecretKeyRef.Name + ":" + envVariable.ValueFrom.SecretKeyRef.Key + "}}"
					tempEnvVariable.Type = string(meshConstants.Secret)
					tempEnvVariable.Dynamic = true
				} else if envVariable.ValueFrom.FieldRef != nil {
					tempEnvVariable.Key = envVariable.Name
					tempEnvVariable.Value = envVariable.ValueFrom.FieldRef.FieldPath
					tempEnvVariable.Type = "FieldRef"
				}
				environmentVariables = append(environmentVariables, tempEnvVariable)
			} else {
				tempEnvVariable.Key = envVariable.Name
				tempEnvVariable.Value = envVariable.Value
				environmentVariables = append(environmentVariables, tempEnvVariable)
			}

		}

		containerTemp.Ports = ports
		containerTemp.EnvironmentVariables = environmentVariables
		containerTemp.VolumeMounts = volumeMounts

		containers = append(containers, &containerTemp)
	}
	return containers, volumeMountNames, nil
}

func getCPProbe(prob *v1.Probe, contPorts []v1.ContainerPort) (*meshTypes.Probe, error) {
	CpProbe := new(meshTypes.Probe)

	if prob.FailureThreshold > 0 {
		CpProbe.FailureThreshold = &prob.FailureThreshold

	}
	if prob.SuccessThreshold > 0 {
		CpProbe.SuccessThreshold = &prob.SuccessThreshold

	}
	CpProbe.InitialDelaySeconds = &prob.InitialDelaySeconds
	if prob.PeriodSeconds > 0 {
		CpProbe.PeriodSeconds = &prob.PeriodSeconds
	}
	if prob.TimeoutSeconds > 0 {
		CpProbe.TimeoutSeconds = &prob.TimeoutSeconds
	}

	if prob.Handler.Exec != nil {
		CpProbe.Handler = new(meshTypes.Handler)
		CpProbe.Handler.Type = "exec"
		CpProbe.Handler.Exec = new(meshTypes.ExecAction)
		for i := 0; i < len(prob.Handler.Exec.Command); i++ {
			CpProbe.Handler.Exec.Command = append(CpProbe.Handler.Exec.Command, prob.Handler.Exec.Command[i])
		}
	} else if prob.HTTPGet != nil {
		CpProbe.Handler = new(meshTypes.Handler)
		CpProbe.Handler.Type = "httpGet"
		CpProbe.Handler.HTTPGet = new(meshTypes.HTTPGetAction)
		if prob.HTTPGet.Port.Type == intstr.Int {
			if prob.HTTPGet.Port.IntVal > 0 && prob.HTTPGet.Port.IntVal < 65536 {
				CpProbe.Handler.HTTPGet.Port = int(prob.HTTPGet.Port.IntVal)
			} else {
				return nil, errors.New("not a valid port number for http_get")
			}
		} else if prob.HTTPGet.Port.Type == intstr.String {
			for _, value := range contPorts {
				if strings.EqualFold(value.Name, prob.HTTPGet.Port.StrVal) {
					CpProbe.Handler.HTTPGet.Port = int(value.ContainerPort)
				}
			}
		}
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

		if prob.HTTPGet.Scheme == v1.URISchemeHTTP || prob.HTTPGet.Scheme == v1.URISchemeHTTPS {
			if prob.HTTPGet.Scheme == v1.URISchemeHTTP {
				scheme := meshTypes.URISchemeHTTP
				CpProbe.Handler.HTTPGet.Scheme = &scheme
			} else if prob.HTTPGet.Scheme == v1.URISchemeHTTPS {
				scheme := meshTypes.URISchemeHTTPS
				CpProbe.Handler.HTTPGet.Scheme = &scheme
			}
		} else if prob.HTTPGet.Scheme == "" {
			CpProbe.Handler.HTTPGet.Scheme = nil
		} else {
			return nil, errors.New("invalid URI scheme")
		}

		for i := 0; i < len(prob.HTTPGet.HTTPHeaders); i++ {
			var cphttpheader meshTypes.HTTPHeader
			cphttpheader.Name = &prob.HTTPGet.HTTPHeaders[i].Name
			cphttpheader.Value = &prob.HTTPGet.HTTPHeaders[i].Value
			CpProbe.Handler.HTTPGet.HTTPHeaders = append(CpProbe.Handler.HTTPGet.HTTPHeaders, cphttpheader)
		}
	} else if prob.TCPSocket != nil {
		CpProbe.Handler = new(meshTypes.Handler)
		CpProbe.Handler.Type = "tcpSocket"
		CpProbe.Handler.TCPSocket = new(meshTypes.TCPSocketAction)
		if prob.TCPSocket.Port.Type == intstr.Int {
			if prob.TCPSocket.Port.IntVal > 0 && prob.TCPSocket.Port.IntVal < 65536 {
				CpProbe.Handler.TCPSocket.Port = int(prob.TCPSocket.Port.IntVal)
			} else {
				return nil, errors.New("not a valid port number for tcpsocket")
			}
		} else if prob.TCPSocket.Port.Type == intstr.String {
			for _, value := range contPorts {
				if strings.EqualFold(value.Name, prob.TCPSocket.Port.StrVal) {
					CpProbe.Handler.TCPSocket.Port = int(value.ContainerPort)
				}
			}

		}
		if prob.Handler.TCPSocket.Host != "" {
			CpProbe.Handler.TCPSocket.Host = &prob.Handler.TCPSocket.Host
		}

	} else {
		return nil, errors.New("no handler found")
	}
	return CpProbe, nil

}

func putCPCommandAndArguments(container *meshTypes.ContainerAttribute, command, args []string) error {
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

func putCPResource(container *meshTypes.ContainerAttribute, limitResources map[v1.ResourceName]resource.Quantity, isLimit bool) error {
	temp := make(map[string]string)
	for t, v := range limitResources {
		key := t.String()
		if key == meshTypes.ResourceTypeMemory || key == meshTypes.ResourceTypeCpu {
			quantity := v.String()
			temp[key] = quantity
		} else {
			return errors.New("Error Found: Invalid Request Resource Provided. Valid: 'cpu','memory'")
		}
	}
	if isLimit {
		container.LimitResources = temp
	} else {
		container.RequestResources = temp
	}

	return nil
}

func getCPSecurityContext(securityContext *v1.SecurityContext) (*meshTypes.SecurityContextStruct, error) {
	context := new(meshTypes.SecurityContextStruct)
	//if securityContext.Capabilities != nil {
	//	context.Capabilities = new(meshTypes.Capabilities)
	//	CpAdd := make([]meshTypes.Capability, len(securityContext.Capabilities.Add))
	//	for index, kubeAdd := range securityContext.Capabilities.Add {
	//		CpAdd[index] = meshTypes.Capability(kubeAdd)
	//	}
	//	context.Capabilities.Add = CpAdd
	//
	//	CpDrop := make([]meshTypes.Capability, len(securityContext.Capabilities.Drop))
	//	for index, kubeDrop := range securityContext.Capabilities.Drop {
	//		CpDrop[index] = meshTypes.Capability(kubeDrop)
	//	}
	//	context.Capabilities.Drop = CpDrop
	//}
	//if securityContext.AllowPrivilegeEscalation != nil {
	//	context.AllowPrivilegeEscalation = *securityContext.AllowPrivilegeEscalation
	//}
	//if securityContext.ReadOnlyRootFilesystem != nil && *securityContext.ReadOnlyRootFilesystem {
	//	context.ReadOnlyRootFileSystem = *securityContext.ReadOnlyRootFilesystem
	//}
	//if securityContext.Privileged != nil {
	//	context.Privileged = *securityContext.Privileged
	//}
	//if securityContext.ReadOnlyRootFilesystem != nil {
	//	context.ReadOnlyRootFileSystem = *securityContext.ReadOnlyRootFilesystem
	//}
	//
	//if securityContext.RunAsNonRoot != nil {
	//
	//}
	//if securityContext.RunAsUser != nil {
	//	context.RunAsUser = securityContext.RunAsUser
	//
	//}
	//
	//if securityContext.ProcMount != nil && *securityContext.ProcMount == v1.DefaultProcMount {
	//	context.ProcMount = meshTypes.DefaultProcMount
	//} else if securityContext.ProcMount != nil && *securityContext.ProcMount == v1.UnmaskedProcMount {
	//	context.ProcMount = meshTypes.UnmaskedProcMount
	//}
	//
	//if securityContext.SELinuxOptions != nil {
	//	context.SELinuxOptions = meshTypes.SELinuxOptionsStruct{
	//		User:  securityContext.SELinuxOptions.User,
	//		Role:  securityContext.SELinuxOptions.Role,
	//		Type:  securityContext.SELinuxOptions.Type,
	//		Level: securityContext.SELinuxOptions.Level,
	//	}
	//}
	return context, nil
}

func getCPAffinity(affinity *v1.Affinity) (*meshTypes.Affinity, error) {
	temp := new(meshTypes.Affinity)
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

func getCPNodeAffinity(nodeAffinity *v1.NodeAffinity) (*meshTypes.NodeAffinity, error) {
	temp := new(meshTypes.NodeAffinity)
	if nodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution != nil {
		if ns, err := getCPNodeSelector(nodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution); err != nil {
			return nil, err
		} else {
			temp.ReqDuringSchedulingIgnDuringExec = ns
		}
	}

	var tempPrefSchedulingTerms []meshTypes.PreferredSchedulingTerm
	for _, prefSchedulingTerm := range nodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
		tempPrefSchedulingTerm := meshTypes.PreferredSchedulingTerm{}

		tempPrefSchedulingTerm.Weight = prefSchedulingTerm.Weight
		var tempMatchExpressions []meshTypes.NodeSelectorRequirement
		var tempMatchFields []meshTypes.NodeSelectorRequirement

		for _, matchExpression := range prefSchedulingTerm.Preference.MatchExpressions {
			tempMatchExpression := meshTypes.NodeSelectorRequirement{}
			tempMatchExpression.Key = matchExpression.Key
			tempMatchExpression.Values = matchExpression.Values
			switch matchExpression.Operator {
			case v1.NodeSelectorOpIn:
				tempMatchExpression.Operator = meshTypes.NodeSelectorOpIn
			case v1.NodeSelectorOpNotIn:
				tempMatchExpression.Operator = meshTypes.NodeSelectorOpNotIn
			case v1.NodeSelectorOpExists:
				tempMatchExpression.Operator = meshTypes.NodeSelectorOpExists
			case v1.NodeSelectorOpDoesNotExist:
				tempMatchExpression.Operator = meshTypes.NodeSelectorOpDoesNotExists
			case v1.NodeSelectorOpLt:
				tempMatchExpression.Operator = meshTypes.NodeSelectorOpLt
			case v1.NodeSelectorOpGt:
				tempMatchExpression.Operator = meshTypes.NodeSelectorOpGt
			}
			tempMatchExpressions = append(tempMatchExpressions, tempMatchExpression)
		}
		for _, matchField := range prefSchedulingTerm.Preference.MatchFields {
			tempMatchField := meshTypes.NodeSelectorRequirement{}
			tempMatchField.Key = matchField.Key
			tempMatchField.Values = matchField.Values
			switch matchField.Operator {
			case v1.NodeSelectorOpIn:
				tempMatchField.Operator = meshTypes.NodeSelectorOpIn
			case v1.NodeSelectorOpNotIn:
				tempMatchField.Operator = meshTypes.NodeSelectorOpNotIn
			case v1.NodeSelectorOpExists:
				tempMatchField.Operator = meshTypes.NodeSelectorOpExists
			case v1.NodeSelectorOpDoesNotExist:
				tempMatchField.Operator = meshTypes.NodeSelectorOpDoesNotExists
			case v1.NodeSelectorOpLt:
				tempMatchField.Operator = meshTypes.NodeSelectorOpLt
			case v1.NodeSelectorOpGt:
				tempMatchField.Operator = meshTypes.NodeSelectorOpGt
			}

			tempMatchFields = append(tempMatchFields, tempMatchField)
		}
		tempPrefSchedulingTerm.Preference.MatchExpressions = tempMatchExpressions
		tempPrefSchedulingTerm.Preference.MatchFields = tempMatchFields

		tempPrefSchedulingTerms = append(tempPrefSchedulingTerms, tempPrefSchedulingTerm)

	}
	return temp, nil
}

func getCPPodAffinity(podAffinity *v1.PodAffinity) (*meshTypes.PodAffinity, error) {
	temp := new(meshTypes.PodAffinity)
	var tempPodAffinityTerms []meshTypes.PodAffinityTerm
	for _, podAffinityTerm := range podAffinity.RequiredDuringSchedulingIgnoredDuringExecution {
		tempPodAffinityTerm := meshTypes.PodAffinityTerm{}

		tempPodAffinityTerm.Namespaces = podAffinityTerm.Namespaces
		tempPodAffinityTerm.TopologyKey = podAffinityTerm.TopologyKey
		ls := getCPLabelSelector(podAffinityTerm.LabelSelector)
		tempPodAffinityTerm.LabelSelector = ls

		tempPodAffinityTerms = append(tempPodAffinityTerms, tempPodAffinityTerm)

	}
	temp.ReqDuringSchedulingIgnDuringExec = tempPodAffinityTerms
	var tempWeightedAffinityTerms []meshTypes.WeightedPodAffinityTerm
	for _, weighted := range podAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
		tempWeightedAffinityTerm := meshTypes.WeightedPodAffinityTerm{}

		tempWeightedAffinityTerm.Weight = weighted.Weight

		tempPodAffinityTerm := meshTypes.PodAffinityTerm{}
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

func getCPAntiPodAffinity(podAntiAffinity *v1.PodAntiAffinity) (*meshTypes.PodAntiAffinity, error) {

	temp := new(meshTypes.PodAntiAffinity)
	var tempPodAffinityTerms []meshTypes.PodAffinityTerm
	for _, podAffinityTerm := range podAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution {
		tempPodAffinityTerm := meshTypes.PodAffinityTerm{}

		tempPodAffinityTerm.Namespaces = podAffinityTerm.Namespaces
		tempPodAffinityTerm.TopologyKey = podAffinityTerm.TopologyKey
		ls := getCPLabelSelector(podAffinityTerm.LabelSelector)
		tempPodAffinityTerm.LabelSelector = ls
		tempPodAffinityTerms = append(tempPodAffinityTerms, tempPodAffinityTerm)

	}
	temp.ReqDuringSchedulingIgnDuringExec = tempPodAffinityTerms
	var tempWeightedAffinityTerms []meshTypes.WeightedPodAffinityTerm
	for _, weighted := range podAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
		tempWeightedAffinityTerm := meshTypes.WeightedPodAffinityTerm{}
		tempWeightedAffinityTerm.Weight = weighted.Weight
		tempPodAffinityTerm := meshTypes.PodAffinityTerm{}
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

func getCPVolumes(vols []v1.Volume, volumeMountNames map[string]bool) ([]meshTypes.Volume, error) {
	var volumes []meshTypes.Volume
	for _, volume := range vols {

		if !volumeMountNames[volume.Name] {
			continue
		}
		volumeMountNames[volume.Name] = false
		tempVolume := meshTypes.Volume{}
		tempVolume.Name = volume.Name

		if volume.VolumeSource.Secret != nil {
			tempVolume.VolumeSource.Secret = new(meshTypes.SecretVolumeSource)
			tempVolume.VolumeSource.Secret.SecretName = volume.VolumeSource.Secret.SecretName
			tempVolume.VolumeSource.Secret.DefaultMode = volume.VolumeSource.Secret.DefaultMode
			var secretItems []meshTypes.KeyToPath
			for _, item := range volume.VolumeSource.Secret.Items {
				secretItem := meshTypes.KeyToPath{
					Key:  item.Key,
					Path: item.Path,
					Mode: item.Mode,
				}
				secretItems = append(secretItems, secretItem)
			}
			tempVolume.VolumeSource.Secret.Items = secretItems
		}
		if volume.VolumeSource.ConfigMap != nil {
			tempVolume.VolumeSource.ConfigMap = new(meshTypes.ConfigMapVolumeSource)
			tempVolume.VolumeSource.ConfigMap.Name = volume.VolumeSource.ConfigMap.LocalObjectReference.Name

			tempVolume.VolumeSource.ConfigMap.DefaultMode = volume.VolumeSource.ConfigMap.DefaultMode
			var configMapItems []meshTypes.KeyToPath
			for _, item := range volume.VolumeSource.ConfigMap.Items {
				configMapItem := meshTypes.KeyToPath{
					Key:  item.Key,
					Path: item.Path,
					Mode: item.Mode,
				}
				configMapItems = append(configMapItems, configMapItem)
			}
			tempVolume.VolumeSource.ConfigMap.Items = configMapItems
		}

		if volume.VolumeSource.AWSElasticBlockStore != nil {
			tempVolume.VolumeSource.AWSElasticBlockStore = new(meshTypes.AWSElasticBlockStoreVolumeSource)
			tempVolume.VolumeSource.AWSElasticBlockStore.ReadOnly = volume.VolumeSource.AWSElasticBlockStore.ReadOnly
			tempVolume.VolumeSource.AWSElasticBlockStore.Partition = volume.VolumeSource.AWSElasticBlockStore.Partition
		}

		if volume.VolumeSource.EmptyDir != nil {
			tempVolume.VolumeSource.EmptyDir = new(meshTypes.EmptyDirVolumeSource)
			//quantity, _ := resource.ParseQuantity(volume.VolumeSource.EmptyDir.SizeLimit)
			tempVolume.VolumeSource.EmptyDir.SizeLimit = volume.VolumeSource.EmptyDir.SizeLimit
			if volume.VolumeSource.EmptyDir.Medium == v1.StorageMediumDefault {
				tempVolume.VolumeSource.EmptyDir.Medium = meshTypes.StorageMediumDefault

			}
			if volume.VolumeSource.EmptyDir.Medium == v1.StorageMediumMemory {
				tempVolume.VolumeSource.EmptyDir.Medium = meshTypes.StorageMediumMemory
			}

			if volume.VolumeSource.EmptyDir.Medium == v1.StorageMediumHugePages {
				tempVolume.VolumeSource.EmptyDir.Medium = meshTypes.StorageMediumHugePages
			}

		}

		if volume.VolumeSource.GCEPersistentDisk != nil {
			tempVolume.VolumeSource.GCEPersistentDisk = new(meshTypes.GCEPersistentDiskVolumeSource)
			tempVolume.VolumeSource.GCEPersistentDisk.Partition = volume.VolumeSource.GCEPersistentDisk.Partition
			tempVolume.VolumeSource.GCEPersistentDisk.ReadOnly = volume.VolumeSource.GCEPersistentDisk.ReadOnly
			tempVolume.VolumeSource.GCEPersistentDisk.PDName = volume.VolumeSource.GCEPersistentDisk.PDName
		}

		if volume.VolumeSource.AzureDisk != nil {
			tempVolume.VolumeSource.AzureFile = new(meshTypes.AzureFileVolumeSource)
			tempVolume.VolumeSource.AzureDisk.ReadOnly = volume.VolumeSource.AzureDisk.ReadOnly
			tempVolume.VolumeSource.AzureDisk.DataDiskURI = volume.VolumeSource.AzureDisk.DiskName

			if *volume.VolumeSource.AzureDisk.CachingMode == v1.AzureDataDiskCachingNone {
				temp := meshTypes.AzureDataDiskCachingNone
				tempVolume.VolumeSource.AzureDisk.CachingMode = &temp
			} else if *volume.VolumeSource.AzureDisk.CachingMode == v1.AzureDataDiskCachingReadWrite {
				temp := meshTypes.AzureDataDiskCachingReadWrite
				tempVolume.VolumeSource.AzureDisk.CachingMode = &temp
			} else if *volume.VolumeSource.AzureDisk.CachingMode == v1.AzureDataDiskCachingReadOnly {
				temp := meshTypes.AzureDataDiskCachingReadOnly
				tempVolume.VolumeSource.AzureDisk.CachingMode = &temp
			}

			if *volume.VolumeSource.AzureDisk.Kind == v1.AzureSharedBlobDisk {
				temp := meshTypes.AzureSharedBlobDisk
				tempVolume.VolumeSource.AzureDisk.Kind = &temp
			} else if *volume.VolumeSource.AzureDisk.Kind == v1.AzureDedicatedBlobDisk {
				temp := meshTypes.AzureDedicatedBlobDisk
				tempVolume.VolumeSource.AzureDisk.Kind = &temp
			} else if *volume.VolumeSource.AzureDisk.Kind == v1.AzureManagedDisk {
				temp := meshTypes.AzureManagedDisk
				tempVolume.VolumeSource.AzureDisk.Kind = &temp
			}
		}

		if volume.VolumeSource.AzureFile != nil {
			tempVolume.VolumeSource.AzureFile = new(meshTypes.AzureFileVolumeSource)
			tempVolume.VolumeSource.AzureFile.ReadOnly = volume.VolumeSource.AzureFile.ReadOnly
			tempVolume.VolumeSource.AzureFile.SecretName = volume.VolumeSource.AzureFile.SecretName
			tempVolume.VolumeSource.AzureFile.ShareName = volume.VolumeSource.AzureFile.ShareName

		}
		if volume.VolumeSource.HostPath != nil {
			tempVolume.VolumeSource.HostPath = new(meshTypes.HostPathVolumeSource)
			tempVolume.VolumeSource.HostPath.Path = volume.VolumeSource.HostPath.Path
			if volume.VolumeSource.HostPath.Type != nil {
				if *volume.VolumeSource.HostPath.Type == v1.HostPathUnset {
					hostPathType := meshTypes.HostPathUnset
					tempVolume.VolumeSource.HostPath.Type = &hostPathType
				} else if *volume.VolumeSource.HostPath.Type == v1.HostPathDirectoryOrCreate {
					hostPathType := meshTypes.HostPathDirectoryOrCreate
					tempVolume.VolumeSource.HostPath.Type = &hostPathType
				} else if *volume.VolumeSource.HostPath.Type == v1.HostPathFileOrCreate {
					hostPathType := meshTypes.HostPathFileOrCreate
					tempVolume.VolumeSource.HostPath.Type = &hostPathType
				}
			}

		}

		volumes = append(volumes, tempVolume)

	}

	return volumes, nil
}

func convertToCPVirtualService(input *v1alpha3.VirtualService) (*meshTypes.VirtualService, error) {
	var vServ = new(meshTypes.VirtualService)
	if input.Labels["version"] != "" {
		vServ.Version = input.Labels["version"]
	} else {
		vServ.Version = ""
	}
	vServ.ServiceType = meshConstants.MeshType
	vServ.ServiceSubType = meshConstants.VirtualService
	vServ.Name = input.Name
	vServ.Namespace = input.Namespace
	if vr := input.Labels["version"]; vr != "" {
		vServ.Version = vr
	}
	vServ.ServiceAttributes = new(meshTypes.VSServiceAttribute)
	vServ.ServiceAttributes.Hosts = input.Spec.Hosts
	vServ.ServiceAttributes.Gateways = input.Spec.Gateways

	for _, http := range input.Spec.Http {
		vSer := new(meshTypes.Http)

		for _, match := range http.Match {
			m := new(meshTypes.HttpMatchRequest)
			m.Name = match.Name
			if match.Uri != nil {
				m.Uri = new(meshTypes.HttpMatch)
				if match.Uri.GetPrefix() != "" {
					m.Uri.Type = "prefix"
					m.Uri.Value = match.Uri.GetPrefix()
				} else if match.Uri.GetExact() != "" {
					m.Uri.Type = "exact"
					m.Uri.Value = match.Uri.GetExact()
				} else if match.Uri.GetRegex() != "" {
					m.Uri.Type = "regex"
					m.Uri.Value = match.Uri.GetRegex()
				}
			}
			if match.Scheme != nil {
				m.Scheme = new(meshTypes.HttpMatch)
				if match.Scheme.GetPrefix() != "" {
					m.Scheme.Type = "prefix"
					m.Scheme.Value = match.Scheme.GetPrefix()
				} else if match.Scheme.GetExact() != "" {
					m.Scheme.Type = "exact"
					m.Scheme.Value = match.Scheme.GetExact()
				} else if match.Scheme.GetRegex() != "" {
					m.Scheme.Type = "regex"
					m.Scheme.Value = match.Scheme.GetRegex()
				}
			}
			if match.Method != nil {
				m.Method = new(meshTypes.HttpMatch)
				if match.Method.GetPrefix() != "" {
					m.Method.Type = "prefix"
					m.Method.Value = match.Method.GetPrefix()
				} else if match.Method.GetExact() != "" {
					m.Method.Type = "exact"
					m.Method.Value = match.Method.GetExact()
				} else if match.Method.GetRegex() != "" {
					m.Method.Type = "regex"
					m.Method.Value = match.Method.GetRegex()
				}
			}
			if match.Authority != nil {
				m.Authority = new(meshTypes.HttpMatch)
				if match.Authority.GetPrefix() != "" {
					m.Authority.Type = "prefix"
					m.Authority.Value = match.Authority.GetPrefix()
				} else if match.Authority.GetExact() != "" {
					m.Authority.Type = "exact"
					m.Authority.Value = match.Authority.GetExact()
				} else if match.Authority.GetRegex() != "" {
					m.Authority.Type = "regex"
					m.Authority.Value = match.Authority.GetRegex()
				}
			}
			if match.Headers != nil {
				m.Headers = make(map[string]*meshTypes.HttpMatch)
				for key, value := range match.Headers {

					tempHttpMatch := new(meshTypes.HttpMatch)
					fmt.Println(value.GetExact())
					if value.GetPrefix() != "" {
						tempHttpMatch.Type = "prefix"
						tempHttpMatch.Value = value.GetPrefix()
					} else if value.GetExact() != "" {
						tempHttpMatch.Type = "exact"
						tempHttpMatch.Value = value.GetExact()
					} else if value.GetRegex() != "" {
						tempHttpMatch.Type = "regex"
						tempHttpMatch.Value = value.GetRegex()
					}

					m.Headers[key] = tempHttpMatch
				}

			}

			vSer.HttpMatch = append(vSer.HttpMatch, m)
		}

		for _, route := range http.Route {
			r := new(meshTypes.HttpRoute)

			if route.Destination != nil {
				destRoute := new(meshTypes.RouteDestination)
				destRoute.Host = route.Destination.Host
				destRoute.Subset = route.Destination.Subset
				if route.Destination.Port != nil {
					destRoute.Port = int32(route.Destination.Port.Number)
				}
				r.Routes = append(r.Routes, destRoute)

			}
			r.Weight = &route.Weight
			vSer.HttpRoute = append(vSer.HttpRoute, r)
		}
		if http.Redirect != nil {
			vSer.HttpRedirect = new(meshTypes.HttpRedirect)
			vSer.HttpRedirect.Uri = http.Redirect.Uri
			vSer.HttpRedirect.Authority = http.Redirect.Authority
			vSer.HttpRedirect.RedirectCode = int32(http.Redirect.RedirectCode)
		}
		if http.Rewrite != nil {
			vSer.HttpRewrite = new(meshTypes.HttpRewrite)
			vSer.HttpRewrite.Uri = http.Rewrite.Uri
			vSer.HttpRewrite.Authority = http.Rewrite.Authority
		}
		if http.Timeout != nil {
			vSer.Timeout = time.Duration(http.Timeout.Seconds)
		}

		if http.Fault != nil {
			vSer.FaultInjection = new(meshTypes.HttpFaultInjection)
			if http.Fault.GetDelay() != nil {
				if http.Fault.GetDelay().GetFixedDelay() != nil {
					vSer.FaultInjection.DelayType = "FixedDelay"
					vSer.FaultInjection.DelayValue = time.Duration(http.Fault.Delay.GetFixedDelay().Seconds)
				} else if http.Fault.GetDelay().GetExponentialDelay() != nil {
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
			vSer.CorsPolicy = new(meshTypes.HttpCorsPolicy)
			vSer.CorsPolicy.AllowOrigin = http.CorsPolicy.AllowOrigin
			vSer.CorsPolicy.AllowMethod = http.CorsPolicy.AllowMethods
			vSer.CorsPolicy.AllowHeaders = http.CorsPolicy.AllowHeaders
			vSer.CorsPolicy.ExposeHeaders = http.CorsPolicy.ExposeHeaders
			vSer.CorsPolicy.MaxAge = time.Duration(http.CorsPolicy.MaxAge.Seconds)
			vSer.CorsPolicy.AllowCredentials = http.CorsPolicy.AllowCredentials.Value
		}
		if http.Retries != nil {
			vSer.Retry = new(meshTypes.HttpRetry)
			vSer.Retry.TotalAttempts = http.Retries.Attempts
			if http.Retries.PerTryTimeout != nil {
				vSer.Retry.PerTryTimeOut = http.Retries.PerTryTimeout.Seconds
			}
			vSer.Retry.RetryOn = http.Retries.RetryOn
		}

		vServ.ServiceAttributes.Http = append(vServ.ServiceAttributes.Http, vSer)
	}

	for _, serv := range input.Spec.Tls {
		tls := new(meshTypes.Tls)
		for _, match := range serv.Match {
			m := new(meshTypes.TlsMatchAttribute)
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
			r := new(meshTypes.TlsRoute)
			if route.Destination != nil {
				r.RouteDestination = new(meshTypes.RouteDestination)
				r.Weight = route.Weight
				if route.Destination.Port != nil {
					r.RouteDestination.Port = int32(route.Destination.Port.Number)
				}
				r.RouteDestination.Subset = route.Destination.Subset
				r.RouteDestination.Host = route.Destination.Host
				tls.Route = append(tls.Route, r)
			}

		}
		vServ.ServiceAttributes.Tls = append(vServ.ServiceAttributes.Tls, tls)
	}

	for _, serv := range input.Spec.Tcp {
		tcp := new(meshTypes.Tcp)
		for _, match := range serv.Match {
			m := new(meshTypes.TcpMatchRequest)
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
			d := new(meshTypes.TcpRoutes)
			if route.Destination != nil {
				d.Destination = new(meshTypes.RouteDestination)
				if route.Destination.Port != nil {
					d.Destination.Port = int32(route.Destination.Port.Number)
				}
				d.Destination.Subset = route.Destination.Subset
				d.Destination.Host = route.Destination.Host
			}
			d.Weight = route.Weight
			tcp.Routes = append(tcp.Routes, d)
		}
		vServ.ServiceAttributes.Tcp = append(vServ.ServiceAttributes.Tcp, tcp)
	}
	return vServ, nil
}

func convertToCPDestinationRule(input *v1alpha3.DestinationRule) (*meshTypes.DestinationRules, error) {
	vServ := new(meshTypes.DestinationRules)
	vServ.ServiceType = meshConstants.MeshType
	vServ.ServiceSubType = meshConstants.DestinationRule
	vServ.Name = input.Name
	vServ.Namespace = input.Namespace
	if vr := input.Labels["version"]; vr != "" {
		vServ.Version = vr
	}

	vServ.ServiceAttributes.Host = input.Spec.Host
	if input.Spec.TrafficPolicy != nil {
		vServ.ServiceAttributes.TrafficPolicy = new(meshTypes.TrafficPolicy)

		if input.Spec.TrafficPolicy.LoadBalancer != nil {
			vServ.ServiceAttributes.TrafficPolicy.LoadBalancer = new(meshTypes.LoadBalancer)
			loadBalType := strings.Split(input.Spec.TrafficPolicy.LoadBalancer.String(), ":")

			if loadBalType[0] == "simple" {
				vServ.ServiceAttributes.TrafficPolicy.LoadBalancer.Simple = input.Spec.TrafficPolicy.LoadBalancer.GetSimple().String()
			} else if loadBalType[0] == "consistent_hash" {
				vServ.ServiceAttributes.TrafficPolicy.LoadBalancer.ConsistentHash = new(meshTypes.ConsistentHash)
				if input.Spec.TrafficPolicy.LoadBalancer.GetConsistentHash().GetHttpCookie() != nil {
					vServ.ServiceAttributes.TrafficPolicy.LoadBalancer.ConsistentHash.HttpCookie = new(meshTypes.HttpCookie)
					vServ.ServiceAttributes.TrafficPolicy.LoadBalancer.ConsistentHash.HttpCookie.Name = input.Spec.TrafficPolicy.LoadBalancer.GetConsistentHash().GetHttpCookie().Name
					vServ.ServiceAttributes.TrafficPolicy.LoadBalancer.ConsistentHash.HttpCookie.Path = input.Spec.TrafficPolicy.LoadBalancer.GetConsistentHash().GetHttpCookie().Path
					vServ.ServiceAttributes.TrafficPolicy.LoadBalancer.ConsistentHash.HttpCookie.Ttl = input.Spec.TrafficPolicy.LoadBalancer.GetConsistentHash().GetHttpCookie().Ttl.Seconds

				} else if input.Spec.TrafficPolicy.LoadBalancer.GetConsistentHash().GetUseSourceIp() == true {
					vServ.ServiceAttributes.TrafficPolicy.LoadBalancer.ConsistentHash.HTTPHeaderName = input.Spec.TrafficPolicy.LoadBalancer.GetConsistentHash().GetHttpHeaderName()
				} else if input.Spec.TrafficPolicy.LoadBalancer.GetConsistentHash().GetHttpHeaderName() != "" {
					vServ.ServiceAttributes.TrafficPolicy.LoadBalancer.ConsistentHash.UseSourceIP = input.Spec.TrafficPolicy.LoadBalancer.GetConsistentHash().GetUseSourceIp()
				}
			}
		}
		if input.Spec.TrafficPolicy.ConnectionPool != nil {
			vServ.ServiceAttributes.TrafficPolicy.ConnectionPool = new(meshTypes.ConnectionPool)
			if input.Spec.TrafficPolicy.ConnectionPool.Tcp != nil {
				vServ.ServiceAttributes.TrafficPolicy.ConnectionPool.Tcp = new(meshTypes.DrTcp)
				if input.Spec.TrafficPolicy.ConnectionPool.Tcp.ConnectTimeout != nil {

					timeout := time.Duration(input.Spec.TrafficPolicy.ConnectionPool.Tcp.ConnectTimeout.GetNanos())
					vServ.ServiceAttributes.TrafficPolicy.ConnectionPool.Tcp.ConnectTimeout = &timeout
				}

				vServ.ServiceAttributes.TrafficPolicy.ConnectionPool.Tcp.MaxConnections = input.Spec.TrafficPolicy.ConnectionPool.Tcp.MaxConnections
				if input.Spec.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive != nil {
					vServ.ServiceAttributes.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive = new(meshTypes.TcpKeepalive)
					if input.Spec.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive.Interval != nil {
						keepAlive := time.Duration(input.Spec.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive.Interval.Seconds)
						vServ.ServiceAttributes.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive.Interval = &keepAlive
					}

					vServ.ServiceAttributes.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive.Probes = input.Spec.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive.Probes
					if input.Spec.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive.Time != nil {
						timealive := time.Duration(input.Spec.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive.Time.Seconds)
						vServ.ServiceAttributes.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive.Time = &timealive
					}

				}
			}
			if input.Spec.TrafficPolicy.ConnectionPool.Http != nil {
				vServ.ServiceAttributes.TrafficPolicy.ConnectionPool.Http = new(meshTypes.DrHttp)
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
			vServ.ServiceAttributes.TrafficPolicy.OutlierDetection = new(meshTypes.OutlierDetection)
			if input.Spec.TrafficPolicy.OutlierDetection.BaseEjectionTime != nil {
				injecTime := time.Duration(input.Spec.TrafficPolicy.OutlierDetection.BaseEjectionTime.GetSeconds())
				vServ.ServiceAttributes.TrafficPolicy.OutlierDetection.BaseEjectionTime = &injecTime
			}

			vServ.ServiceAttributes.TrafficPolicy.OutlierDetection.ConsecutiveErrors = input.Spec.TrafficPolicy.OutlierDetection.ConsecutiveErrors
			if input.Spec.TrafficPolicy.OutlierDetection.Interval != nil {
				interval := time.Duration(input.Spec.TrafficPolicy.OutlierDetection.Interval.GetSeconds())
				vServ.ServiceAttributes.TrafficPolicy.OutlierDetection.Interval = &interval
			}

			vServ.ServiceAttributes.TrafficPolicy.OutlierDetection.MaxEjectionPercent = input.Spec.TrafficPolicy.OutlierDetection.MaxEjectionPercent
			vServ.ServiceAttributes.TrafficPolicy.OutlierDetection.MinHealthPercent = input.Spec.TrafficPolicy.OutlierDetection.MinHealthPercent
		}

		for _, port := range input.Spec.TrafficPolicy.PortLevelSettings {

			setting := new(meshTypes.PortLevelSetting)

			if port.Port != nil {
				setting.Port = new(meshTypes.DrPort)
				setting.Port.Number = int32(port.Port.Number)
			}

			if port.ConnectionPool != nil {
				setting.ConnectionPool = new(meshTypes.ConnectionPool)
				if setting.ConnectionPool.Tcp != nil {
					setting.ConnectionPool.Tcp = new(meshTypes.DrTcp)
					if port.ConnectionPool.Tcp.ConnectTimeout != nil {
						timeout := time.Duration(port.ConnectionPool.Tcp.ConnectTimeout.Nanos)
						setting.ConnectionPool.Tcp.ConnectTimeout = &timeout
					}
					setting.ConnectionPool.Tcp.MaxConnections = port.ConnectionPool.Tcp.MaxConnections
					if port.ConnectionPool.Tcp.TcpKeepalive != nil {
						setting.ConnectionPool.Tcp.TcpKeepalive = new(meshTypes.TcpKeepalive)
						if port.ConnectionPool.Tcp.TcpKeepalive.Time != nil {
							t := time.Duration(port.ConnectionPool.Tcp.TcpKeepalive.Time.Seconds)
							setting.ConnectionPool.Tcp.TcpKeepalive.Time = &t
						}
						if port.ConnectionPool.Tcp.TcpKeepalive.Interval != nil {
							interval := time.Duration(port.ConnectionPool.Tcp.TcpKeepalive.Interval.Seconds)
							setting.ConnectionPool.Tcp.TcpKeepalive.Interval = &interval
						}

						setting.ConnectionPool.Tcp.TcpKeepalive.Probes = port.ConnectionPool.Tcp.TcpKeepalive.Probes

					}

				}
				if port.ConnectionPool.Http != nil {
					setting.ConnectionPool.Http = new(meshTypes.DrHttp)
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
				setting.LoadBalancer = new(meshTypes.LoadBalancer)
				if port.LoadBalancer.GetSimple().String() != "" {
					setting.LoadBalancer.Simple = port.LoadBalancer.GetSimple().String()
				} else if port.LoadBalancer.GetConsistentHash() != nil {
					setting.LoadBalancer.ConsistentHash = new(meshTypes.ConsistentHash)
					if port.LoadBalancer.GetConsistentHash().GetHttpHeaderName() != "" {
						setting.LoadBalancer.ConsistentHash.HTTPHeaderName = port.LoadBalancer.GetConsistentHash().GetHttpHeaderName()
					} else if port.LoadBalancer.GetConsistentHash().GetUseSourceIp() != false {
						setting.LoadBalancer.ConsistentHash.UseSourceIP = port.LoadBalancer.GetConsistentHash().GetUseSourceIp()
					} else if port.LoadBalancer.GetConsistentHash().GetHttpCookie() != nil {
						setting.LoadBalancer.ConsistentHash.HttpCookie = new(meshTypes.HttpCookie)
						setting.LoadBalancer.ConsistentHash.HttpCookie.Name = port.LoadBalancer.GetConsistentHash().GetHttpCookie().Name
						setting.LoadBalancer.ConsistentHash.HttpCookie.Path = port.LoadBalancer.GetConsistentHash().GetHttpCookie().Path

						setting.LoadBalancer.ConsistentHash.HttpCookie.Ttl = port.LoadBalancer.GetConsistentHash().GetHttpCookie().Ttl.Seconds

					}

					setting.LoadBalancer.ConsistentHash.MinimumRingSize = strconv.Itoa(int(port.LoadBalancer.GetConsistentHash().GetMinimumRingSize()))

				}
			}
			if port.OutlierDetection != nil {
				setting.OutlierDetection = new(meshTypes.OutlierDetection)
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
			vServ.ServiceAttributes.TrafficPolicy.DrTls = new(meshTypes.DrTls)
			vServ.ServiceAttributes.TrafficPolicy.DrTls.Mode = input.Spec.TrafficPolicy.Tls.GetMode().String()
			vServ.ServiceAttributes.TrafficPolicy.DrTls.ClientCertificate = input.Spec.TrafficPolicy.Tls.ClientCertificate
			vServ.ServiceAttributes.TrafficPolicy.DrTls.PrivateKey = input.Spec.TrafficPolicy.Tls.PrivateKey
			vServ.ServiceAttributes.TrafficPolicy.DrTls.CaCertificate = input.Spec.TrafficPolicy.Tls.CaCertificates
			vServ.ServiceAttributes.TrafficPolicy.DrTls.SubjectAltNames = input.Spec.TrafficPolicy.Tls.SubjectAltNames[0]
		}

	}
	for _, subset := range input.Spec.Subsets {
		ser := new(meshTypes.Subset)
		ser.Name = subset.Name
		if len(subset.Labels) > 0 {
			labels := make(map[string]string)
			labels = subset.Labels
			ser.Labels = &labels
		}

		if subset.TrafficPolicy != nil {
			ser.TrafficPolicy = new(meshTypes.TrafficPolicy)
			for _, port := range subset.TrafficPolicy.PortLevelSettings {
				setting := new(meshTypes.PortLevelSetting)
				if port.Port != nil {
					setting.Port = new(meshTypes.DrPort)
					setting.Port.Number = int32(port.Port.Number)
				}
				if port.ConnectionPool != nil {
					setting.ConnectionPool = new(meshTypes.ConnectionPool)
					if setting.ConnectionPool.Tcp != nil {
						setting.ConnectionPool.Tcp = new(meshTypes.DrTcp)
						if port.ConnectionPool.Tcp.ConnectTimeout != nil {
							timeout := time.Duration(port.ConnectionPool.Tcp.ConnectTimeout.Nanos)
							setting.ConnectionPool.Tcp.ConnectTimeout = &timeout
						}
						setting.ConnectionPool.Tcp.MaxConnections = port.ConnectionPool.Tcp.MaxConnections
						if port.ConnectionPool.Tcp.TcpKeepalive != nil {
							setting.ConnectionPool.Tcp.TcpKeepalive = new(meshTypes.TcpKeepalive)
							if port.ConnectionPool.Tcp.TcpKeepalive.Time != nil {
								t := time.Duration(port.ConnectionPool.Tcp.TcpKeepalive.Time.Seconds)
								setting.ConnectionPool.Tcp.TcpKeepalive.Time = &t
							}
							if port.ConnectionPool.Tcp.TcpKeepalive.Interval != nil {
								interval := time.Duration(port.ConnectionPool.Tcp.TcpKeepalive.Interval.Seconds)
								setting.ConnectionPool.Tcp.TcpKeepalive.Interval = &interval
							}

							setting.ConnectionPool.Tcp.TcpKeepalive.Probes = port.ConnectionPool.Tcp.TcpKeepalive.Probes

						}

					}
					if port.ConnectionPool.Http != nil {
						setting.ConnectionPool.Http = new(meshTypes.DrHttp)
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
					setting.LoadBalancer = new(meshTypes.LoadBalancer)
					if port.LoadBalancer.GetSimple().String() != "" {
						setting.LoadBalancer.Simple = port.LoadBalancer.GetSimple().String()
					} else if port.LoadBalancer.GetConsistentHash() != nil {
						setting.LoadBalancer.ConsistentHash = new(meshTypes.ConsistentHash)

						if port.LoadBalancer.GetConsistentHash().GetHttpCookie() != nil {
							setting.LoadBalancer.ConsistentHash.HttpCookie = new(meshTypes.HttpCookie)
							if port.LoadBalancer.GetConsistentHash().GetHttpCookie().Ttl != nil {
								setting.LoadBalancer.ConsistentHash.HttpCookie.Ttl = port.LoadBalancer.GetConsistentHash().GetHttpCookie().Ttl.Seconds

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
					setting.OutlierDetection = new(meshTypes.OutlierDetection)
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
					setting.DrTls = new(meshTypes.DrTls)
					vServ.ServiceAttributes.TrafficPolicy.DrTls.Mode = string(port.Tls.GetMode())
					setting.DrTls.ClientCertificate = port.Tls.ClientCertificate
					setting.DrTls.PrivateKey = port.Tls.PrivateKey
					setting.DrTls.CaCertificate = port.Tls.CaCertificates
					setting.DrTls.SubjectAltNames = port.Tls.SubjectAltNames[0]

				}
				vServ.ServiceAttributes.TrafficPolicy.PortLevelSettings = append(vServ.ServiceAttributes.TrafficPolicy.PortLevelSettings, setting)
			}
			if subset.TrafficPolicy.LoadBalancer != nil {
				ser.TrafficPolicy.LoadBalancer = new(meshTypes.LoadBalancer)
				if subset.TrafficPolicy.LoadBalancer.GetSimple().String() != "" {
					ser.TrafficPolicy.LoadBalancer.Simple = subset.TrafficPolicy.LoadBalancer.GetSimple().String()
				} else if subset.TrafficPolicy.LoadBalancer.GetConsistentHash() != nil {
					ser.TrafficPolicy.LoadBalancer.ConsistentHash = new(meshTypes.ConsistentHash)

					if subset.TrafficPolicy.LoadBalancer.GetConsistentHash().GetHttpCookie() != nil {
						if subset.TrafficPolicy.LoadBalancer.GetConsistentHash().GetHttpCookie().Ttl != nil {
							ser.TrafficPolicy.LoadBalancer.ConsistentHash.HttpCookie.Ttl = subset.TrafficPolicy.LoadBalancer.GetConsistentHash().GetHttpCookie().Ttl.Seconds

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
				ser.TrafficPolicy.ConnectionPool = new(meshTypes.ConnectionPool)
				if subset.TrafficPolicy.ConnectionPool.Tcp != nil {
					ser.TrafficPolicy.ConnectionPool.Tcp = new(meshTypes.DrTcp)
					if subset.TrafficPolicy.ConnectionPool.Tcp.ConnectTimeout != nil {
						timeout := time.Duration(subset.TrafficPolicy.ConnectionPool.Tcp.ConnectTimeout.Nanos)
						ser.TrafficPolicy.ConnectionPool.Tcp.ConnectTimeout = &timeout
					}

					ser.TrafficPolicy.ConnectionPool.Tcp.MaxConnections = subset.TrafficPolicy.ConnectionPool.Tcp.MaxConnections
					if subset.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive != nil {
						ser.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive = new(meshTypes.TcpKeepalive)
						if subset.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive.Time != nil {
							t := time.Duration(subset.TrafficPolicy.ConnectionPool.Tcp.TcpKeepalive.Time.Seconds)
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
					ser.TrafficPolicy.ConnectionPool.Http = new(meshTypes.DrHttp)
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
				ser.TrafficPolicy.OutlierDetection = new(meshTypes.OutlierDetection)
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
				ser.TrafficPolicy.DrTls = new(meshTypes.DrTls)
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

func convertToCPGateway(input *v1alpha3.Gateway) (*meshTypes.GatewayService, error) {
	gateway := new(meshTypes.GatewayService)
	gateway.Name = input.Name
	if input.Labels["version"] != "" {
		gateway.Version = input.Labels["version"]
	} else {
		gateway.Version = ""
	}
	gateway.ServiceType = meshConstants.MeshType
	gateway.ServiceSubType = meshConstants.Gateway
	gateway.Namespace = input.Namespace
	if vr := input.Labels["version"]; vr != "" {
		gateway.Version = vr
	}

	gateway.ServiceAttributes = new(meshTypes.GatewayServiceAttributes)
	gateway.ServiceAttributes.Selectors = make(map[string]string)
	gateway.ServiceAttributes.Selectors = input.Spec.Selector

	for _, serverInput := range input.Spec.Servers {
		server := new(meshTypes.Server)
		if serverInput.Tls != nil {
			server.Tls = new(meshTypes.TlsConfig)
			server.Tls.HttpsRedirect = serverInput.Tls.HttpsRedirect
			server.Tls.Mode = meshTypes.Mode(serverInput.Tls.Mode.String())
			server.Tls.ServerCertificate = serverInput.Tls.ServerCertificate
			server.Tls.CaCertificate = serverInput.Tls.CaCertificates
			server.Tls.PrivateKey = serverInput.Tls.PrivateKey
			for _, altNames := range serverInput.Tls.SubjectAltNames {
				server.Tls.SubjectAltName = append(serverInput.Tls.SubjectAltNames, altNames)
			}
			server.Tls.MinProtocolVersion = meshTypes.ProtocolVersion(serverInput.Tls.MinProtocolVersion.String())
			server.Tls.MaxProtocolVersion = meshTypes.ProtocolVersion(serverInput.Tls.MaxProtocolVersion.String())
		}
		if serverInput.Port != nil {
			server.Port = new(meshTypes.Port)
			server.Port.Name = serverInput.Port.Name
			server.Port.Number = serverInput.Port.Number
			server.Port.Protocol = meshTypes.Protocols(serverInput.Port.Protocol)
		}
		for _, host := range serverInput.Hosts {
			server.Hosts = append(server.Hosts, host)
		}

		gateway.ServiceAttributes.Servers = append(gateway.ServiceAttributes.Servers, server)

	}
	return gateway, nil

}

func convertToCPServiceEntry(input *v1alpha3.ServiceEntry) (*meshTypes.ServiceEntry, error) {
	svcEntry := new(meshTypes.ServiceEntry)
	svcEntry.Name = input.Name
	svcEntry.Namespace = input.Namespace
	svcEntry.ServiceType = meshConstants.MeshType
	svcEntry.ServiceSubType = meshConstants.ServiceEntry
	if input.Labels["version"] != "" {
		svcEntry.Version = input.Labels["version"]
	}
	svcEntry.ServiceAttributes = new(meshTypes.ServiceEntryAttributes)
	for _, host := range input.Spec.Hosts {
		svcEntry.ServiceAttributes.Hosts = append(svcEntry.ServiceAttributes.Hosts, host)
	}
	for _, address := range input.Spec.Addresses {
		svcEntry.ServiceAttributes.Addresses = append(svcEntry.ServiceAttributes.Addresses, address)
	}
	for _, port := range input.Spec.Ports {
		tempPort := new(meshTypes.ServiceEntryPort)
		tempPort.Name = port.Name
		tempPort.Protocol = port.Protocol
		tempPort.Number = port.Number
		svcEntry.ServiceAttributes.Ports = append(svcEntry.ServiceAttributes.Ports, tempPort)

	}
	for _, entryPoint := range input.Spec.Endpoints {
		tempEntryPoint := new(meshTypes.ServiceEntryEndpoint)
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
