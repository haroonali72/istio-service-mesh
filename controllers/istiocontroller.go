package controllers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	yaml2 "github.com/ghodss/yaml"
	googl_types "github.com/gogo/protobuf/types"
	"github.com/iancoleman/strcase"
	"istio-service-mesh/core"
	//"github.com/istio/api/networking/v1alpha3"
	"io/ioutil"
	"istio-service-mesh/constants"
	"istio-service-mesh/controllers/volumes"
	"istio-service-mesh/types"
	"istio-service-mesh/utils"
	policy "istio.io/api/authentication/v1alpha1"
	"istio.io/api/networking/v1alpha3"
	ist_rbac "istio.io/api/rbac/v1alpha1"
	"istio.io/istio/pilot/pkg/model"
	v12 "k8s.io/api/apps/v1"
	autoscaling "k8s.io/api/autoscaling/v2beta2"
	v13 "k8s.io/api/batch/v1"
	"k8s.io/api/batch/v2alpha1"
	"k8s.io/api/core/v1"
	rbacV1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"net/http"
	"strconv"
	"strings"
)

var Notifier utils.Notifier

func getIstioVirtualService(service interface{}) (string, error) {

	vService := v1alpha3.VirtualService{}

	var serviceAttr types.IstioVirtualServiceAttributes

	byteData, _ := json.Marshal(service)
	err := json.Unmarshal(byteData, &serviceAttr)
	if err != nil {
		utils.Error.Println(err)
		return "", err
	}
	var routes []*v1alpha3.HTTPRoute

	for _, http := range serviceAttr.HTTP {

		var httpRoute v1alpha3.HTTPRoute

		var destination []*v1alpha3.HTTPRouteDestination
		for _, route := range http.Routes {

			var httpD v1alpha3.HTTPRouteDestination
			httpD.Destination = &v1alpha3.Destination{Subset: route.Destination.Subset, Host: route.Destination.Host}
			if route.Destination.Port != 0 {
				httpD.Destination.Port = &v1alpha3.PortSelector{Port: &v1alpha3.PortSelector_Number{Number: uint32(route.Destination.Port)}}

			}
			if route.Weight > 0 {
				httpD.Weight = route.Weight
			}
			destination = append(destination, &httpD)
		}
		httpRoute.Route = destination

		var matches []*v1alpha3.HTTPMatchRequest
		for _, match := range http.Match {

			for _, uri := range match.Uris {
				matches = append(matches, &v1alpha3.HTTPMatchRequest{Uri: &v1alpha3.StringMatch{MatchType: &v1alpha3.StringMatch_Prefix{Prefix: uri}}})
			}
		}
		httpRoute.Match = matches
		/*	if http.RewriteUri != "" {
				var rewrite v1alpha3.HTTPRewrite
				rewrite.Uri = http.RewriteUri
				httpRoute.Rewrite = &rewrite
			}
			if http.RetriesUri != "" {
				var retries v1alpha3.HTTPRetry
				retries.RetryOn = http.RetriesUri
				httpRoute.Retries = &retries
			}*/
		if http.Timeout > 0 {
			time := googl_types.Duration{Seconds: http.Timeout}
			httpRoute.Timeout = &time
		}
		for _, retries := range http.Retries {

			var httpR v1alpha3.HTTPRetry
			if retries.Attempts > 0 && retries.Timeout > 0 {
				httpR.Attempts = int32(retries.Attempts)

				time := googl_types.Duration{Seconds: retries.Timeout}
				httpR.PerTryTimeout = &time

				httpRoute.Retries = &httpR
			}
		}
		set := false
		fault := &v1alpha3.HTTPFaultInjection{}
		if http.FaultInjection.FaultInjectionAbort.Percentage != 0 && http.FaultInjection.FaultInjectionAbort.HttpStatus != 0 {
			abort := &v1alpha3.HTTPFaultInjection_Abort{
				Percentage: &v1alpha3.Percent{Value: http.FaultInjection.FaultInjectionAbort.Percentage},
				ErrorType:  &v1alpha3.HTTPFaultInjection_Abort_HttpStatus{HttpStatus: http.FaultInjection.FaultInjectionAbort.HttpStatus},
			}
			fault.Abort = abort
			set = true
		}
		if http.FaultInjection.FaultInjectionDelay.Percentage != 0 && http.FaultInjection.FaultInjectionDelay.FixedDelay != 0 {
			delay := &v1alpha3.HTTPFaultInjection_Delay{
				Percentage:    &v1alpha3.Percent{Value: http.FaultInjection.FaultInjectionDelay.Percentage},
				HttpDelayType: &v1alpha3.HTTPFaultInjection_Delay_FixedDelay{FixedDelay: &googl_types.Duration{Seconds: http.FaultInjection.FaultInjectionDelay.FixedDelay}},
			}
			fault.Delay = delay
			set = true
		}
		if set {
			httpRoute.Fault = fault
		}
		routes = append(routes, &httpRoute)
	}
	vService.Http = routes
	vService.Hosts = serviceAttr.Hosts
	if serviceAttr.Gateways != nil {
		vService.Gateways = serviceAttr.Gateways
	}
	/*utils.Info.Println(vService.String())
	b, e := vService.Marshal()
	if e == nil {
		utils.Info.Println(string(b))
	}*/

	/*b, e := json.Marshal(vService)
	if e != nil {
		utils.Info.Println(e.Error())
	}
	utils.Info.Println(string(b))*/
	gotJSON, err := model.ToYAML(&vService)
	if err != nil {
		utils.Error.Println(err)
	}
	return gotJSON, nil
}

func getIstioGateway() (v1alpha3.Gateway, error) {
	gateway := v1alpha3.Gateway{}
	var hosts []string
	hosts = append(hosts, "*")
	var servers []*v1alpha3.Server
	var serv v1alpha3.Server
	serv.Port = &v1alpha3.Port{Name: strings.ToLower("HTTP"), Protocol: "HTTP", Number: uint32(80)}
	serv.Hosts = hosts
	servers = append(servers, &serv)

	/*var serv2 v1alpha3.Server
	serv2.Port = &v1alpha3.Port{Name: strings.ToLower("HTTPS"), Protocol: "HTTPS", Number: uint32(443)}
	serv2.Hosts = hosts
	servers = append(servers, &serv2)*/

	selector := make(map[string]string)

	selector["istio"] = "ingressgateway"
	gateway.Selector = selector
	gateway.Servers = servers
	return gateway, nil
}

func getIstioDestinationRule(service interface{}) (map[string]interface{}, error) {
	destRule := v1alpha3.DestinationRule{}

	byteData, _ := json.Marshal(service)
	var serviceAttr types.IstioDestinationRuleAttributes
	err := json.Unmarshal(byteData, &serviceAttr)
	if err != nil {
		utils.Info.Println(err.Error())
		return nil, err
	}
	var subsets []*v1alpha3.Subset

	for _, subset := range serviceAttr.Subsets {
		var ss v1alpha3.Subset
		var tp v1alpha3.TrafficPolicy
		ss.Name = subset.Name
		var labels = make(map[string]string)
		for _, label := range subset.Labels {
			labels[label.Key] = label.Value
		}
		ss.Labels = labels
		if subset.Http1MaxPendingRequests > 0 || subset.Http2MaxRequests > 0 || subset.MaxRequestsPerConnection > 0 || subset.MaxRetries > 0 {
			tp.ConnectionPool = &v1alpha3.ConnectionPoolSettings{
				Http: &v1alpha3.ConnectionPoolSettings_HTTPSettings{
					Http1MaxPendingRequests:  subset.Http1MaxPendingRequests,
					Http2MaxRequests:         subset.Http2MaxRequests,
					MaxRequestsPerConnection: subset.MaxRequestsPerConnection,
					MaxRetries:               subset.MaxRetries,
				},
			}
			ss.TrafficPolicy = &tp
		}

		subsets = append(subsets, &ss)
	}
	if len(subsets) > 0 {
		destRule.Subsets = subsets
	}
	destRule.Host = serviceAttr.Host

	switch serviceAttr.TrafficPolicy.TLS.Mode {
	case "ISTIO_MUTUAL":
		destRule.TrafficPolicy = &v1alpha3.TrafficPolicy{
			Tls: &v1alpha3.TLSSettings{
				Mode: v1alpha3.TLSSettings_ISTIO_MUTUAL,
			},
		}
	}
	yamlRaw, err := model.ToYAML(&destRule)
	if err != nil {
		utils.Error.Println(err)
	}
	return marshalUnMarshalOfIstioComponents(yamlRaw)
}
func createPolicy(serviceName string) (map[string]interface{}, error) {
	var p policy.Policy
	p.Targets = append(p.Targets, &policy.TargetSelector{Name: serviceName})
	p.Peers = append(p.Peers, &policy.PeerAuthenticationMethod{Params: &policy.PeerAuthenticationMethod_Mtls{
		Mtls: &policy.MutualTls{},
	},
	})
	yml, err := model.ToYAML(&p)
	if err != nil {
		return nil, err
	}
	return marshalUnMarshalOfIstioComponents(yml)

}
func getIstioServiceEntry(service interface{}) (types.IstioServiceEntryAttributes, v1alpha3.ServiceEntry, error) {
	SE := v1alpha3.ServiceEntry{}
	byteData, _ := json.Marshal(service)
	var serviceAttr types.IstioServiceEntryAttributes
	json.Unmarshal(byteData, &serviceAttr)
	var ports []*v1alpha3.Port
	for _, port := range serviceAttr.Ports {
		var p v1alpha3.Port
		p.Name = port.Name
		p.Protocol = port.Protocol
		p.Number = uint32(port.Port)
		ports = append(ports, &p)
	}

	SE.Ports = ports
	SE.Hosts = serviceAttr.Hosts
	switch serviceAttr.Resolution {
	case "DNS":
		SE.Resolution = v1alpha3.ServiceEntry_DNS
	case "STATIC":
		SE.Resolution = v1alpha3.ServiceEntry_STATIC
	case "NONE":
		SE.Resolution = v1alpha3.ServiceEntry_NONE
	}
	SE.Addresses = serviceAttr.Address
	if serviceAttr.Location == "mesh_external" {

		SE.Location = v1alpha3.ServiceEntry_MESH_EXTERNAL
	} else if serviceAttr.Location == "mesh_internal" {

		SE.Location = v1alpha3.ServiceEntry_MESH_INTERNAL
	}
	for i := range serviceAttr.Uri {
		SE.Endpoints = append(SE.Endpoints, &v1alpha3.ServiceEntry_Endpoint{
			Ports:    serviceAttr.Uri[i].Ports,
			Address:  serviceAttr.Uri[i].Address,
			Labels:   serviceAttr.Uri[i].Labels,
			Locality: serviceAttr.Uri[i].Locality,
			Network:  serviceAttr.Uri[i].Network,
			Weight:   serviceAttr.Uri[i].Weight,
		})
	}
	return serviceAttr, SE, nil
}
func getIstioConf(service types.Service) (types.IstioConfig, error) {
	b, e := json.Marshal(service.ServiceAttributes)
	if e != nil {
		return types.IstioConfig{}, e
	}
	var istioConfig types.IstioConfig
	e = json.Unmarshal(b, &istioConfig)
	if e != nil {
		return types.IstioConfig{}, e
	}
	return istioConfig, nil
}
func CheckGateway(input types.Service) (types.IstioObject, error) {
	var istioServ types.IstioObject

	istioConf, err := getIstioConf(input)
	if err != nil {
		fmt.Println("There is error in deployment")
		return istioServ, err
	}
	if istioConf.Enable_External_Traffic {
		//var istioServ types.IstioObject

		serv, err := getIstioGateway()
		if err != nil {
			utils.Error.Println("There is error in deployment")
			return istioServ, err
		}
		istioServ.Spec = serv
		istioServ.Spec = serv
		labels := make(map[string]interface{})
		labels["app"] = strings.ToLower(input.Name)
		labels["name"] = strings.ToLower(input.Name)
		labels["version"] = strings.ToLower(input.Version)
		labels["namespace"] = strings.ToLower(input.Namespace)
		istioServ.Metadata = labels
		istioServ.Kind = "Gateway"
		istioServ.ApiVersion = "networking.istio.io/v1alpha3"

		return istioServ, nil
	}
	return istioServ, nil
}
func getIstioObject(input types.Service) (components []types.IstioObject, err error) {
	var istioServ types.IstioObject

	switch input.SubType {

	case "service_entry":

		attrib, serv_entry, err := getIstioServiceEntry(input.ServiceAttributes)
		if err != nil {
			utils.Error.Println("There is error in deployment")
			return components, err
		}

		utils.Info.Println(attrib.IsMtlsEnable)
		if attrib.IsMtlsEnable {
			for i := range attrib.Hosts {
				seDestionation := types.IstioDestinationRuleAttributes{Host: attrib.Hosts[i]}
				seDestionation.TrafficPolicy.TLS.Mode = attrib.MtlsMode
				p, err := getIstioDestinationRule(seDestionation)
				if err != nil {
					utils.Error.Println(err)
				} else {
					var policyService types.IstioObject
					labels := make(map[string]interface{})
					labels["name"] = strings.ToLower(input.Name + "-" + strconv.Itoa(i) + "")
					labels["namespace"] = strings.ToLower(input.Namespace)
					policyService.Metadata = labels
					policyService.Kind = "DestinationRule"
					policyService.ApiVersion = "networking.istio.io/v1alpha3"
					policyService.Spec = p
					components = append(components, policyService)
				}
			}
		}
		istioServ.Spec = serv_entry
		labels := make(map[string]interface{})
		labels["name"] = strings.ToLower(input.Name)
		labels["app"] = strings.ToLower(input.Name)
		if input.Namespace == "" {
			input.Namespace = "default"
		}
		labels["namespace"] = strings.ToLower(input.Namespace)
		istioServ.Metadata = labels
		istioServ.Kind = "ServiceEntry"
		istioServ.ApiVersion = "networking.istio.io/v1alpha3"
		istioServ.Metadata["namespace"] = input.Namespace
		components = append(components, istioServ)
		return components, nil
	case "virtual_service":
		vr, err := getIstioVirtualService(input.ServiceAttributes)
		if err != nil {
			utils.Error.Println("There is error in deployment")
			return components, err
		}
		m, err := marshalUnMarshalOfIstioComponents(vr)
		utils.Info.Println(err)
		istioServ.Spec = m
		labels := make(map[string]interface{})
		labels["name"] = strings.ToLower(input.Name)
		labels["app"] = strings.ToLower(input.Name)
		labels["version"] = strings.ToLower(input.Version)
		if input.Namespace == "" {
			input.Namespace = "default"
		}
		labels["namespace"] = strings.ToLower(input.Namespace)
		istioServ.Metadata = labels
		istioServ.Kind = "VirtualService"
		istioServ.ApiVersion = "networking.istio.io/v1alpha3"
		istioServ.Metadata["namespace"] = input.Namespace
		components = append(components, istioServ)
		return components, nil

	case "destination_rule":

		des_rule, err := getIstioDestinationRule(input.ServiceAttributes)
		if err != nil {
			utils.Error.Println("There is error in deployment")
			return components, err
		}
		utils.Info.Println(des_rule)
		if _, ok := des_rule["trafficPolicy"]; ok {
			p, err := createPolicy(input.Name)
			if err != nil {
				utils.Error.Println(err)
			} else {
				var policyService types.IstioObject
				labels := make(map[string]interface{})
				labels["name"] = strings.ToLower(input.Name)
				if input.Namespace == "" {
					input.Namespace = "default"
				}
				labels["namespace"] = strings.ToLower(input.Namespace)
				policyService.Metadata = labels
				policyService.Kind = "Policy"
				policyService.ApiVersion = "authentication.istio.io/v1alpha1"
				policyService.Spec = p
				components = append(components, policyService)
			}
		}
		istioServ.Spec = des_rule
		labels := make(map[string]interface{})
		labels["name"] = strings.ToLower(input.Name)
		labels["app"] = strings.ToLower(input.Name)
		labels["version"] = strings.ToLower(input.Version)
		if input.Namespace == "" {
			input.Namespace = "default"
		}
		labels["namespace"] = strings.ToLower(input.Namespace)
		istioServ.Metadata = labels
		istioServ.Kind = "DestinationRule"
		istioServ.ApiVersion = "networking.istio.io/v1alpha3"
		istioServ.Metadata["namespace"] = input.Namespace
		components = append(components, istioServ)
		return components, nil
	}
	return components, nil

}
func getHPAObject(service types.Service) (autoscaling.HorizontalPodAutoscaler, error) {
	var hpa = autoscaling.HorizontalPodAutoscaler{}
	// Label Selector

	//keel labels
	hpa.Kind = "HorizontalPodAutoscaler"
	hpa.APIVersion = "autoscaling/v2beta2"
	hpaLabels := make(map[string]string)
	//deploymentLabels["keel.sh/match-tag"] = "true"
	hpaLabels["keel.sh/policy"] = "force"

	if service.Name == "" {
		utils.Error.Println("service name not found")
		return autoscaling.HorizontalPodAutoscaler{}, errors.New("Service name not found")
	}
	hpa.ObjectMeta.Name = service.Name + "-" + service.Version
	hpa.ObjectMeta.Labels = hpaLabels

	if service.Namespace == "" {
		hpa.ObjectMeta.Namespace = "default"
	} else {
		hpa.ObjectMeta.Namespace = service.Namespace
	}
	byteData, err := json.Marshal(service.ServiceAttributes)
	if err != nil {
		utils.Error.Println(err.Error())
		return autoscaling.HorizontalPodAutoscaler{}, err
	}

	utils.Info.Println(string(byteData))

	var serviceAttr types.HPAAttributes

	err = json.Unmarshal(byteData, &serviceAttr)
	if err != nil {
		utils.Error.Println(err.Error())
		return autoscaling.HorizontalPodAutoscaler{}, err
	}

	hpa.Spec.MinReplicas = &serviceAttr.MixReplicas
	hpa.Spec.MaxReplicas = serviceAttr.MaxReplicas
	crossObj := autoscaling.CrossVersionObjectReference{
		Kind:       serviceAttr.CrossObjectVersion.Type,
		Name:       serviceAttr.CrossObjectVersion.Name,
		APIVersion: serviceAttr.CrossObjectVersion.Version,
	}
	hpa.Spec.ScaleTargetRef = crossObj

	var metricsArr []autoscaling.MetricSpec
	for _, metrics := range serviceAttr.Metrics_ {
		met := autoscaling.MetricSpec{
			Type: autoscaling.ResourceMetricSourceType,
		}
		target := autoscaling.MetricTarget{}
		if metrics.TargetValueKind == "value" {
			target.Type = autoscaling.ValueMetricType
			target.Value = resource.NewScaledQuantity(metrics.TargetValue, ScaleUnit(metrics.TargetValueUnit))

		} else if metrics.TargetValueKind == "utilization" {
			target.Type = autoscaling.UtilizationMetricType
			v := int32(metrics.TargetValue)
			target.AverageUtilization = &v
		} else if metrics.TargetValueKind == "average" {
			target.Type = autoscaling.AverageValueMetricType
			target.AverageValue = resource.NewScaledQuantity(metrics.TargetValue, ScaleUnit(metrics.TargetValueUnit))
		}

		resource := autoscaling.ResourceMetricSource{}
		if metrics.ResourceKind == "cpu" {
			resource.Name = v1.ResourceCPU
		} else if metrics.ResourceKind == "memory" {
			resource.Name = v1.ResourceMemory
		} else if metrics.ResourceKind == "storage" {
			resource.Name = v1.ResourceEphemeralStorage
		}

		resource.Target = target

		met.Resource = &resource
		metricsArr = append(metricsArr, met)
	}

	hpa.Spec.Metrics = metricsArr
	hpa.Status = autoscaling.HorizontalPodAutoscalerStatus{
		Conditions: []autoscaling.HorizontalPodAutoscalerCondition{
			{
				Type:   autoscaling.AbleToScale,
				Status: v1.ConditionTrue,
			},
		},
	}
	return hpa, nil
}
func ScaleUnit(unit string) resource.Scale {

	if unit == "nano" {
		return resource.Nano
	} else if unit == "micro" {
		return resource.Micro
	} else if unit == "milli" {
		return resource.Milli
	} else if unit == "kilo" {
		return resource.Kilo
	} else if unit == "mega" {
		return resource.Mega
	} else if unit == "giga" {
		return resource.Giga
	} else if unit == "tera" {
		return resource.Tera
	} else if unit == "peta" {
		return resource.Peta
	} else {
		return resource.Exa
	}

}
func setLabelSelector(service types.Service, sel *metav1.LabelSelector) (*metav1.LabelSelector, error) {
	var serviceAttributes = types.DockerServiceAttributes{}
	if data, err := json.Marshal(service.ServiceAttributes); err == nil {

		if err = json.Unmarshal(data, &serviceAttributes); err != nil {
			utils.Error.Println(err)
			return nil, err
		}
	} else {
		return nil, err
	}
	lenl := len(serviceAttributes.LabelSelector.MatchLabel)
	lene := len(serviceAttributes.LabelSelector.MatchExpression)
	if lenl <= 0 && lene <= 0 {
		//		return nil,nil
		return &metav1.LabelSelector{make(map[string]string), nil}, nil
	}

	if (!(lenl > 0)) && lene > 0 {

		sel = &metav1.LabelSelector{nil, nil}
	} else if lene > 0 || lenl > 0 {
		sel = &metav1.LabelSelector{make(map[string]string), nil}

	}

	for k, v := range serviceAttributes.LabelSelector.MatchLabel {
		sel.MatchLabels[k] = v
	}
	for i := 0; i < len(serviceAttributes.LabelSelector.MatchExpression); i++ {
		if len(serviceAttributes.LabelSelector.MatchExpression[i].Key) > 0 && (serviceAttributes.LabelSelector.MatchExpression[i].Operator == types.LabelSelectorOpDoesNotExist ||
			serviceAttributes.LabelSelector.MatchExpression[i].Operator == types.LabelSelectorOpExists ||
			serviceAttributes.LabelSelector.MatchExpression[i].Operator == types.LabelSelectorOpIn ||
			serviceAttributes.LabelSelector.MatchExpression[i].Operator == types.LabelSelectorOpNotIn) {
			byteData, err := json.Marshal(serviceAttributes.LabelSelector.MatchExpression[i])
			if err != nil {
				return sel, err
			}
			var temp metav1.LabelSelectorRequirement

			err = json.Unmarshal(byteData, &temp)
			if err != nil {
				return nil, err
			}
			sel.MatchExpressions = append(sel.MatchExpressions, temp)
		} else {
			return nil, errors.New("Can not Apply Labels.Inavlid MatchExpression Label type")
		}
	}
	return sel, nil
}

func setNodeSelector(service types.Service, sel map[string]string) (map[string]string, error) {
	var serviceAttributes types.DockerServiceAttributes
	if data, err := json.Marshal(service.ServiceAttributes); err == nil {

		if err = json.Unmarshal(data, &serviceAttributes); err != nil {
			utils.Error.Println(err)
			return nil, err
		}
	} else {
		return nil, err
	}

	if len(sel) <= 0 {
		sel = make(map[string]string)
	}
	for k, v := range serviceAttributes.NodeSelector {
		sel[k] = v
	}
	return sel, nil
}
func getDeploymentObject(service types.Service) (v12.Deployment, error) {
	var secrets, configMaps []string
	var deployment v12.Deployment
	if service.Name == "" {
		//Failed
		return v12.Deployment{}, errors.New("Service name not found")
	}
	deployment.Kind = "Deployment"
	deployment.APIVersion = "apps/v1"
	if service.Namespace == "" {
		deployment.ObjectMeta.Namespace = "default"
	} else {
		deployment.ObjectMeta.Namespace = service.Namespace
	}
	//add label to deployment object
	var err2 error
	deploymentLabels, err2 := getLabels(service)
	if err2 != nil {
		return v12.Deployment{}, err2
	}
	if deploymentLabels == nil {
		deploymentLabels = make(map[string]string)
	}

	//deploymentLabels["keel.sh/match-tag"] = "true"
	deploymentLabels["keel.sh/policy"] = "force"
	//deploymentLabels["keel.sh/trigger"] = "poll"

	deployment.ObjectMeta.Name = service.Name + "-" + service.Version
	deployment.ObjectMeta.Labels = deploymentLabels

	//add label to container in deployment object
	labels, err2 := getLabels(service)
	if err2 != nil {
		return v12.Deployment{}, err2
	}
	if labels == nil {
		labels = make(map[string]string)
	}

	labels["app"] = service.Name
	labels["version"] = strings.ToLower(service.Version)

	deployment.Spec.Template.Spec.NodeSelector, err2 = setNodeSelector(service, deployment.Spec.Template.Spec.NodeSelector)
	if err2 != nil {
		return v12.Deployment{}, err2
	}
	// adding replica
	{
		if service.Replicas >= 0 {
			deployment.Spec.Replicas = &service.Replicas
		}
	}
	//adding label selector
	deployment.Spec.Selector = &metav1.LabelSelector{make(map[string]string), nil}
	deployment.Spec.Selector.MatchLabels = labels
	deployment.Spec.Template.ObjectMeta.Labels = labels
	Annotations, err4 := getAnnotations(service)
	if err4 != nil {
		return v12.Deployment{}, err4
	}
	if Annotations == nil {
		Annotations = make(map[string]string)
	}
	Annotations["sidecar.istio.io/inject"] = "true"
	deployment.Spec.Template.ObjectMeta.Annotations = Annotations
	var err error
	deployment.Spec.Template.Spec.Containers, secrets, configMaps, err = getContainers(service)
	if err != nil {
		return v12.Deployment{}, err
	}
	isExistSecret := make(map[string]bool)
	isExistConfigMap := make(map[string]bool)
	for _, every := range secrets {
		isExistSecret[every] = true
		deployment.Spec.Template.Spec.Volumes = append(deployment.Spec.Template.Spec.Volumes, v1.Volume{
			Name: every,
			VolumeSource: v1.VolumeSource{
				Secret: &v1.SecretVolumeSource{
					SecretName: every,
				},
			},
		})
	}
	for _, every := range configMaps {
		isExistConfigMap[every] = true
		deployment.Spec.Template.Spec.Volumes = append(deployment.Spec.Template.Spec.Volumes, v1.Volume{
			Name: every,
			VolumeSource: v1.VolumeSource{
				ConfigMap: &v1.ConfigMapVolumeSource{
					LocalObjectReference: v1.LocalObjectReference{
						Name: every,
					},
				},
			},
		})
	}

	deployment.Spec.Template.Spec.InitContainers, secrets, configMaps, err = getInitContainers(service)
	if err != nil {
		return v12.Deployment{}, err
	}

	for _, every := range secrets {
		if _, ok := isExistSecret[every]; !ok {
			deployment.Spec.Template.Spec.Volumes = append(deployment.Spec.Template.Spec.Volumes, v1.Volume{
				Name: every,
				VolumeSource: v1.VolumeSource{
					Secret: &v1.SecretVolumeSource{
						SecretName: every,
					},
				},
			})
		}
	}
	for _, every := range configMaps {
		if _, ok := isExistConfigMap[every]; !ok {
			deployment.Spec.Template.Spec.Volumes = append(deployment.Spec.Template.Spec.Volumes, v1.Volume{
				Name: every,
				VolumeSource: v1.VolumeSource{
					ConfigMap: &v1.ConfigMapVolumeSource{
						LocalObjectReference: v1.LocalObjectReference{
							Name: every,
						},
					},
				},
			})
		}
	}

	return deployment, nil
}
func getDaemonSetObject(service types.Service) (v12.DaemonSet, error) {
	var secrets, configMaps []string
	var daemonset = v12.DaemonSet{}
	if service.Name == "" {
		//Failed
		return v12.DaemonSet{}, errors.New("Service name not found")
	}

	if service.Namespace == "" {
		daemonset.ObjectMeta.Namespace = "default"
	} else {
		daemonset.ObjectMeta.Namespace = service.Namespace
	}

	daemonset.Kind = "DaemonSet"
	daemonset.APIVersion = "apps/v1"
	// Label Selector
	//keel labels
	var err2 error
	deploymentLabels := make(map[string]string)
	//deploymentLabels["keel.sh/match-tag"] = "true"
	deploymentLabels, err2 = getLabels(service)
	if err2 != nil {
		return v12.DaemonSet{}, err2
	}
	deploymentLabels["keel.sh/policy"] = "force"
	//deploymentLabels["keel.sh/trigger"] = "poll"

	labels, err2 := getLabels(service)
	if err2 != nil {
		return v12.DaemonSet{}, err2
	}
	if labels == nil {
		labels = make(map[string]string)
	}

	labels["app"] = service.Name
	labels["version"] = strings.ToLower(service.Version)
	daemonset.ObjectMeta.Name = service.Name + "-" + service.Version
	daemonset.ObjectMeta.Labels = deploymentLabels

	/*	daemonset.Spec.Selector, err2 = setLabelSelector(service, daemonset.Spec.Selector)
		if err2 != nil {
			return v12.DaemonSet{}, err2
		}
	*/
	daemonset.Spec.Selector = &metav1.LabelSelector{make(map[string]string), nil}
	daemonset.Spec.Selector.MatchLabels = labels
	daemonset.Spec.Template.Spec.NodeSelector, err2 = setNodeSelector(service, daemonset.Spec.Template.Spec.NodeSelector)
	if err2 != nil {
		return v12.DaemonSet{}, err2
	}

	daemonset.Spec.Template.ObjectMeta.Labels = labels
	Annotations, err4 := getAnnotations(service)
	if err4 != nil {
		return v12.DaemonSet{}, err4
	}
	if Annotations == nil {
		Annotations = make(map[string]string)
	}

	Annotations["sidecar.istio.io/inject"] = "true"
	daemonset.Spec.Template.ObjectMeta.Annotations = Annotations

	var err error
	daemonset.Spec.Template.Spec.Containers, secrets, configMaps, err = getContainers(service)
	if err != nil {
		return v12.DaemonSet{}, err
	}
	isExistSecret := make(map[string]bool)
	isExistConfigMap := make(map[string]bool)
	for _, every := range secrets {
		isExistSecret[every] = true
		daemonset.Spec.Template.Spec.Volumes = append(daemonset.Spec.Template.Spec.Volumes, v1.Volume{
			Name: every,
			VolumeSource: v1.VolumeSource{
				Secret: &v1.SecretVolumeSource{
					SecretName: every,
				},
			},
		})
	}
	for _, every := range configMaps {
		isExistConfigMap[every] = true
		daemonset.Spec.Template.Spec.Volumes = append(daemonset.Spec.Template.Spec.Volumes, v1.Volume{
			Name: every,
			VolumeSource: v1.VolumeSource{
				ConfigMap: &v1.ConfigMapVolumeSource{
					LocalObjectReference: v1.LocalObjectReference{
						Name: every,
					},
				},
			},
		})
	}

	daemonset.Spec.Template.Spec.InitContainers, secrets, configMaps, err = getInitContainers(service)
	if err != nil {
		return v12.DaemonSet{}, err
	}

	for _, every := range secrets {
		if _, ok := isExistSecret[every]; !ok {
			daemonset.Spec.Template.Spec.Volumes = append(daemonset.Spec.Template.Spec.Volumes, v1.Volume{
				Name: every,
				VolumeSource: v1.VolumeSource{
					Secret: &v1.SecretVolumeSource{
						SecretName: every,
					},
				},
			})
		}
	}
	for _, every := range configMaps {
		if _, ok := isExistConfigMap[every]; !ok {
			daemonset.Spec.Template.Spec.Volumes = append(daemonset.Spec.Template.Spec.Volumes, v1.Volume{
				Name: every,
				VolumeSource: v1.VolumeSource{
					ConfigMap: &v1.ConfigMapVolumeSource{
						LocalObjectReference: v1.LocalObjectReference{
							Name: every,
						},
					},
				},
			})
		}
	}

	return daemonset, nil
}
func getCronJobObject(service types.Service) (v2alpha1.CronJob, error) {
	var secrets, configMaps []string
	var cronjob = v2alpha1.CronJob{}
	cronjob.Kind = "CronJob"
	cronjob.APIVersion = "batch/v1beta1"
	// Label Selector
	//keel labels

	deploymentLabels, err2 := getLabels(service)

	if err2 != nil {
		return v2alpha1.CronJob{}, err2
	}
	if deploymentLabels == nil {
		deploymentLabels = make(map[string]string)
	}

	//deploymentLabels["keel.sh/match-tag"] = "true"
	deploymentLabels["keel.sh/policy"] = "force"
	//deploymentLabels["keel.sh/trigger"] = "poll"

	var selector metav1.LabelSelector

	labels, err2 := getLabels(service)

	if err2 != nil {
		return v2alpha1.CronJob{}, err2
	}
	if labels == nil {
		labels = make(map[string]string)
	}

	labels["app"] = service.Name
	labels["version"] = strings.ToLower(service.Version)

	if service.Name == "" {
		//Failed
		return v2alpha1.CronJob{}, errors.New("Service name not found")
	}
	cronjob.ObjectMeta.Name = service.Name + "-" + service.Version
	cronjob.ObjectMeta.Labels = deploymentLabels
	selector.MatchLabels = labels

	if service.Namespace == "" {
		cronjob.ObjectMeta.Namespace = "default"
	} else {
		cronjob.ObjectMeta.Namespace = service.Namespace
	}

	cronjob.Spec.JobTemplate.Spec.Template.Spec.NodeSelector, err2 = setNodeSelector(service, cronjob.Spec.JobTemplate.Spec.Template.Spec.NodeSelector)
	if err2 != nil {
		return v2alpha1.CronJob{}, err2
	}

	cronjob.Spec.JobTemplate.Spec.Template.ObjectMeta.Labels = labels
	Annotations, err4 := getAnnotations(service)
	if err4 != nil {
		return v2alpha1.CronJob{}, err4
	}
	if Annotations == nil {
		Annotations = make(map[string]string)
	}
	Annotations["sidecar.istio.io/inject"] = "false"
	cronjob.Spec.JobTemplate.Spec.Template.ObjectMeta.Annotations = Annotations
	//

	byteData, _ := json.Marshal(service.ServiceAttributes)
	var serviceAttr types.DockerServiceAttributes
	json.Unmarshal(byteData, &serviceAttr)
	if len(serviceAttr.CronJobScheduleString) <= 0 {
		return v2alpha1.CronJob{}, errors.New("cron job schedule can not be zero")

	}
	if errrr := standardParser.Parse(serviceAttr.CronJobScheduleString); errrr != nil {
		return v2alpha1.CronJob{}, errors.New("invalid  cron job schedule")
	}
	cronjob.Spec.Schedule = serviceAttr.CronJobScheduleString

	var err error
	cronjob.Spec.JobTemplate.Spec.Template.Spec.Containers, secrets, configMaps, err = getContainers(service)
	if err != nil {
		return v2alpha1.CronJob{}, err
	}
	isExistSecret := make(map[string]bool)
	isExistConfigMap := make(map[string]bool)
	for _, every := range secrets {
		isExistSecret[every] = true
		cronjob.Spec.JobTemplate.Spec.Template.Spec.Volumes = append(cronjob.Spec.JobTemplate.Spec.Template.Spec.Volumes, v1.Volume{
			Name: every,
			VolumeSource: v1.VolumeSource{
				Secret: &v1.SecretVolumeSource{
					SecretName: every,
				},
			},
		})
	}
	for _, every := range configMaps {
		isExistConfigMap[every] = true
		cronjob.Spec.JobTemplate.Spec.Template.Spec.Volumes = append(cronjob.Spec.JobTemplate.Spec.Template.Spec.Volumes, v1.Volume{
			Name: every,
			VolumeSource: v1.VolumeSource{
				ConfigMap: &v1.ConfigMapVolumeSource{
					LocalObjectReference: v1.LocalObjectReference{
						Name: every,
					},
				},
			},
		})
	}

	cronjob.Spec.JobTemplate.Spec.Template.Spec.InitContainers, secrets, configMaps, err = getInitContainers(service)
	if err != nil {
		return v2alpha1.CronJob{}, err
	}

	for _, every := range secrets {
		if _, ok := isExistSecret[every]; !ok {
			cronjob.Spec.JobTemplate.Spec.Template.Spec.Volumes = append(cronjob.Spec.JobTemplate.Spec.Template.Spec.Volumes, v1.Volume{
				Name: every,
				VolumeSource: v1.VolumeSource{
					Secret: &v1.SecretVolumeSource{
						SecretName: every,
					},
				},
			})
		}
	}
	for _, every := range configMaps {
		if _, ok := isExistConfigMap[every]; !ok {
			cronjob.Spec.JobTemplate.Spec.Template.Spec.Volumes = append(cronjob.Spec.JobTemplate.Spec.Template.Spec.Volumes, v1.Volume{
				Name: every,
				VolumeSource: v1.VolumeSource{
					ConfigMap: &v1.ConfigMapVolumeSource{
						LocalObjectReference: v1.LocalObjectReference{
							Name: every,
						},
					},
				},
			})
		}
	}

	cronjob.Spec.JobTemplate.Spec.Template.Spec.RestartPolicy = v1.RestartPolicyNever

	return cronjob, nil
}
func getJobObject(service types.Service) (v13.Job, error) {
	var secrets, configMaps []string
	var job = v13.Job{}
	job.Kind = "Job"
	job.APIVersion = "batch/v1"
	// Label Selector

	//keel labels
	deploymentLabels := make(map[string]string)
	//deploymentLabels["keel.sh/match-tag"] = "true"
	deploymentLabels["keel.sh/policy"] = "force"
	//deploymentLabels["keel.sh/trigger"] = "poll"

	var selector metav1.LabelSelector
	var err2 error
	labels, err2 := getLabels(service)
	if err2 != nil {
		return v13.Job{}, err2
	}
	if labels == nil {
		labels = make(map[string]string)
	}
	labels["app"] = service.Name
	labels["version"] = strings.ToLower(service.Version)

	if service.Name == "" {
		//Failed
		return v13.Job{}, errors.New("Service name not found")
	}
	job.ObjectMeta.Name = service.Name + "-" + service.Version
	job.ObjectMeta.Labels = deploymentLabels
	selector.MatchLabels = labels

	if service.Namespace == "" {
		job.ObjectMeta.Namespace = "default"
	} else {
		job.ObjectMeta.Namespace = service.Namespace
	}
	job.Spec.Template.Spec.NodeSelector, err2 = setNodeSelector(service, job.Spec.Template.Spec.NodeSelector)
	if err2 != nil {
		return v13.Job{}, err2
	}
	job.Spec.Template.ObjectMeta.Labels = labels
	Annotations, err4 := getAnnotations(service)
	if err4 != nil {
		return v13.Job{}, err4
	}
	if Annotations == nil {
		Annotations = make(map[string]string)
	}
	Annotations["sidecar.istio.io/inject"] = "false"
	job.Spec.Template.ObjectMeta.Annotations = Annotations
	var err error
	job.Spec.Template.Spec.Containers, secrets, configMaps, err = getContainers(service)
	if err != nil {
		return v13.Job{}, err
	}
	isExistSecret := make(map[string]bool)
	isExistConfigMap := make(map[string]bool)
	for _, every := range secrets {
		isExistSecret[every] = true
		job.Spec.Template.Spec.Volumes = append(job.Spec.Template.Spec.Volumes, v1.Volume{
			Name: every,
			VolumeSource: v1.VolumeSource{
				Secret: &v1.SecretVolumeSource{
					SecretName: every,
				},
			},
		})
	}
	for _, every := range configMaps {
		isExistConfigMap[every] = true
		job.Spec.Template.Spec.Volumes = append(job.Spec.Template.Spec.Volumes, v1.Volume{
			Name: every,
			VolumeSource: v1.VolumeSource{
				ConfigMap: &v1.ConfigMapVolumeSource{
					LocalObjectReference: v1.LocalObjectReference{
						Name: every,
					},
				},
			},
		})
	}

	job.Spec.Template.Spec.InitContainers, secrets, configMaps, err = getInitContainers(service)
	if err != nil {
		return v13.Job{}, err
	}

	for _, every := range secrets {
		if _, ok := isExistSecret[every]; !ok {
			job.Spec.Template.Spec.Volumes = append(job.Spec.Template.Spec.Volumes, v1.Volume{
				Name: every,
				VolumeSource: v1.VolumeSource{
					Secret: &v1.SecretVolumeSource{
						SecretName: every,
					},
				},
			})
		}
	}
	for _, every := range configMaps {
		if _, ok := isExistConfigMap[every]; !ok {
			job.Spec.Template.Spec.Volumes = append(job.Spec.Template.Spec.Volumes, v1.Volume{
				Name: every,
				VolumeSource: v1.VolumeSource{
					ConfigMap: &v1.ConfigMapVolumeSource{
						LocalObjectReference: v1.LocalObjectReference{
							Name: every,
						},
					},
				},
			})
		}
	}

	job.Spec.Template.Spec.RestartPolicy = v1.RestartPolicyNever
	return job, nil
}
func getStatefulSetObject(service types.Service) (v12.StatefulSet, error) {
	var secrets, configMaps []string
	var statefulset = v12.StatefulSet{}
	if service.Name == "" {
		//Failed
		return v12.StatefulSet{}, errors.New("Service name not found")
	}

	if service.Namespace == "" {
		statefulset.ObjectMeta.Namespace = "default"
	} else {
		statefulset.ObjectMeta.Namespace = service.Namespace
	}
	statefulset.Kind = "StatefulSet"
	statefulset.APIVersion = "apps/v1"
	// Label Selector
	//keel labels
	var err2 error
	deploymentLabels := make(map[string]string)
	deploymentLabels, err2 = getLabels(service)
	if err2 != nil {
		return v12.StatefulSet{}, err2
	}
	//deploymentLabels["keel.sh/match-tag"] = "true"
	deploymentLabels["keel.sh/policy"] = "force"
	//deploymentLabels["keel.sh/trigger"] = "poll"

	labels, err2 := getLabels(service)
	if err2 != nil {
		return v12.StatefulSet{}, err2
	}
	if labels == nil {
		labels = make(map[string]string)
	}
	labels["app"] = service.Name
	labels["version"] = strings.ToLower(service.Version)

	statefulset.ObjectMeta.Name = service.Name + "-" + service.Version
	statefulset.ObjectMeta.Labels = deploymentLabels

	statefulset.Spec.Template.Spec.NodeSelector, err2 = setNodeSelector(service, statefulset.Spec.Template.Spec.NodeSelector)
	if err2 != nil {
		return v12.StatefulSet{}, err2
	}
	statefulset.Spec.Selector = &metav1.LabelSelector{make(map[string]string), nil}
	statefulset.Spec.Template.ObjectMeta.Labels = labels
	statefulset.Spec.Selector.MatchLabels = labels
	// adding replica
	{
		if service.Replicas >= 0 {
			statefulset.Spec.Replicas = &service.Replicas
		}
	}
	Annotations, err4 := getAnnotations(service)
	if err4 != nil {
		return v12.StatefulSet{}, err4
	}
	if Annotations == nil {
		Annotations = make(map[string]string)
	}
	Annotations["sidecar.istio.io/inject"] = "true"
	statefulset.Spec.Template.ObjectMeta.Annotations = Annotations

	var err error
	statefulset.Spec.Template.Spec.Containers, secrets, configMaps, err = getContainers(service)
	if err != nil {
		return v12.StatefulSet{}, err
	}
	isExistSecret := make(map[string]bool)
	isExistConfigMap := make(map[string]bool)
	for _, every := range secrets {
		isExistSecret[every] = true
		statefulset.Spec.Template.Spec.Volumes = append(statefulset.Spec.Template.Spec.Volumes, v1.Volume{
			Name: every,
			VolumeSource: v1.VolumeSource{
				Secret: &v1.SecretVolumeSource{
					SecretName: every,
				},
			},
		})
	}
	for _, every := range configMaps {
		isExistConfigMap[every] = true
		statefulset.Spec.Template.Spec.Volumes = append(statefulset.Spec.Template.Spec.Volumes, v1.Volume{
			Name: every,
			VolumeSource: v1.VolumeSource{
				ConfigMap: &v1.ConfigMapVolumeSource{
					LocalObjectReference: v1.LocalObjectReference{
						Name: every,
					},
				},
			},
		})
	}

	statefulset.Spec.Template.Spec.InitContainers, secrets, configMaps, err = getInitContainers(service)
	if err != nil {
		return v12.StatefulSet{}, err
	}
	for _, every := range secrets {
		isExistSecret[every] = true
		statefulset.Spec.Template.Spec.Volumes = append(statefulset.Spec.Template.Spec.Volumes, v1.Volume{
			Name: every,
			VolumeSource: v1.VolumeSource{
				Secret: &v1.SecretVolumeSource{
					SecretName: every,
				},
			},
		})
	}
	for _, every := range configMaps {
		isExistConfigMap[every] = true
		statefulset.Spec.Template.Spec.Volumes = append(statefulset.Spec.Template.Spec.Volumes, v1.Volume{
			Name: every,
			VolumeSource: v1.VolumeSource{
				ConfigMap: &v1.ConfigMapVolumeSource{
					LocalObjectReference: v1.LocalObjectReference{
						Name: every,
					},
				},
			},
		})
	}

	return statefulset, nil
}
func getConfigMapObject(service types.Service) (*v1.ConfigMap, error) {

	byteData, _ := json.Marshal(service.ServiceAttributes)
	var serviceAttr types.ConfigMap
	if err := json.Unmarshal(byteData, &serviceAttr); err != nil {
		utils.Error.Println(err)
		return nil, err
	}

	var configmap = v1.ConfigMap{}
	configmap.Kind = "ConfigMap"
	configmap.APIVersion = "v1"
	// Label Selector
	//keel labels
	//deploymentLabels := make(map[string]string)
	//deploymentLabels["keel.sh/match-tag"] = "true"
	//deploymentLabels["keel.sh/policy"] = "force"
	//deploymentLabels["keel.sh/trigger"] = "poll"

	//var selector metav1.LabelSelector
	//labels, _ := getLabels(service)
	labels := make(map[string]string)
	labels["app"] = service.Name
	labels["version"] = strings.ToLower(service.Version)

	if service.Name == "" {
		//Failed
		return &v1.ConfigMap{}, errors.New("Service name not found")
	}
	configmap.GenerateName = service.ID
	configmap.Name = service.ID

	//configmap.ObjectMeta.Name = service.Name
	configmap.ObjectMeta.Labels = labels
	//selector.MatchLabels = labels
	if service.Namespace == "" {
		configmap.Namespace = service.Namespace
		configmap.ObjectMeta.Namespace = "default"
	} else {
		configmap.Namespace = service.Namespace
		configmap.ObjectMeta.Namespace = service.Namespace
	}
	configmap.ObjectMeta.Name = service.ID
	//configmap.ObjectMeta.GenerateName = service.ID
	configmap.Data = make(map[string]string)
	if serviceAttr.Data != nil {
		for key, value := range serviceAttr.Data {
			configmap.Data[key] = value
		}
	}

	return &configmap, nil
}
func getServiceObject(input types.Service) (*v1.Service, error) {

	service := v1.Service{}
	service.Kind = "Service"
	service.APIVersion = "v1"
	service.Name = input.Name
	service.ObjectMeta.Name = input.Name

	if input.Namespace == "" {
		service.ObjectMeta.Namespace = "default"
	} else {
		service.ObjectMeta.Namespace = input.Namespace
	}
	service.Spec.Type = v1.ServiceTypeClusterIP

	labels, _ := getLabels(input)
	labels["app"] = service.Name
	service.Spec.Selector = labels
	byteData, _ := json.Marshal(input.ServiceAttributes)
	var serviceAttr types.DockerServiceAttributes
	err := json.Unmarshal(byteData, &serviceAttr)
	if err != nil {
		retunrsvc := v1.Service{}
		return &retunrsvc, err
	}
	var servicePorts []v1.ServicePort
	temp := v1.Service{}
	if serviceAttr.Ports == nil {
		return &temp, nil
	}
	for _, port := range serviceAttr.Ports {
		temp := v1.ServicePort{}
		if port.Container == "" && port.Host == "" {
			continue
		}
		if port.Container == "" && port.Host != "" {
			port.Container = port.Host
		}

		i, err := strconv.Atoi(port.Container)
		if err != nil {
			utils.Info.Println(err)
			continue
		}

		temp.Port = int32(i)
		if port.Name == "" {

			temp.Name = "http-" + strconv.Itoa(i)
		} else {
			temp.Name = port.Name
		}
		if port.Host != "" {
			i, err = strconv.Atoi(port.Host)
			if err != nil {
				utils.Info.Println(err)
				continue
			}
			temp.TargetPort = intstr.IntOrString{IntVal: int32(i)}
		}
		servicePorts = append(servicePorts, temp)
	}
	if len(servicePorts) == 0 {
		return nil, nil
	}
	service.Spec.Ports = servicePorts
	return &service, nil
}

func getIstioRbacObjects(serviceAttr types.DockerServiceAttributes, serviceName string, nameSpace string) ([]types.IstioObject, error) {

	var istioObjects []types.IstioObject

	//var roles []ist_rbac.ServiceRole
	//var roleBindings []ist_rbac.ServiceRoleBinding
	for i, role := range serviceAttr.IstioRoles {
		name := strings.ToLower(serviceName + "-r" + strconv.Itoa(i) + "")
		rule := ist_rbac.AccessRule{}
		rule.Methods = role.Methods
		rule.Services = []string{serviceName}
		rule.Paths = role.Paths

		roleObj := ist_rbac.ServiceRole{}
		roleObj.Rules = []*ist_rbac.AccessRule{&rule}

		//roles = append(roles, roleObj)

		var istioRole types.IstioObject
		labels := make(map[string]interface{})
		labels["name"] = strings.ToLower(serviceName + "-r" + strconv.Itoa(i) + "")
		labels["namespace"] = strings.ToLower(nameSpace)
		istioRole.Metadata = labels
		istioRole.Kind = "ServiceRole"
		istioRole.ApiVersion = "rbac.istio.io/v1alpha1"
		istioRole.Spec = roleObj

		istioObjects = append(istioObjects, istioRole)

		// role binding

		roleBinding := ist_rbac.ServiceRoleBinding{}
		roleBinding.Role = name

		properties := make(map[string]string)
		properties["source.namespace"] = nameSpace
		subject := ist_rbac.Subject{Properties: properties}

		roleRef := ist_rbac.RoleRef{}
		roleRef.Name = name
		roleRef.Kind = "ServiceRole"
		roleBinding.Subjects = []*ist_rbac.Subject{&subject}
		roleBinding.RoleRef = &roleRef

		var istioRB types.IstioObject
		rbLabels := make(map[string]interface{})
		rbLabels["name"] = strings.ToLower(serviceName + "-rb" + strconv.Itoa(i) + "")
		rbLabels["namespace"] = strings.ToLower(nameSpace)
		istioRB.Metadata = rbLabels
		istioRB.Kind = "ServiceRoleBinding"
		istioRB.ApiVersion = "rbac.istio.io/v1alpha1"
		istioRB.Spec = roleBinding
		istioObjects = append(istioObjects, istioRB)

	}

	return istioObjects, nil
}

func getRbacObjects(serviceAttr types.DockerServiceAttributes, serviceName string, nameSpace string) (v1.ServiceAccount, []rbacV1.Role, []rbacV1.RoleBinding, error) {
	account := v1.ServiceAccount{}
	account.Name = "sa-" + serviceName
	account.Namespace = nameSpace
	account.APIVersion = "v1"
	account.Kind = "ServiceAccount"
	var roles []rbacV1.Role
	var roleBindings []rbacV1.RoleBinding
	for _, role := range serviceAttr.RbacRoles {
		roleObj := rbacV1.Role{}
		roleObj.Namespace = nameSpace
		roleObj.Name = "sa-" + serviceName + "-role"
		roleObj.Kind = "Role"
		roleObj.APIVersion = "rbac.authorization.k8s.io/v1"
		rule := rbacV1.PolicyRule{APIGroups: role.ApiGroup,
			Resources: []string{role.Resource},
			Verbs:     role.Verbs}
		roleObj.Rules = []rbacV1.PolicyRule{rule}
		roles = append(roles, roleObj)

		// role binding

		rb := rbacV1.RoleBinding{
			ObjectMeta: metav1.ObjectMeta{Name: roleObj.Name + "binding"},
			Subjects: []rbacV1.Subject{
				{Kind: "ServiceAccount", Name: "sa-" + serviceName},
			},
			RoleRef: rbacV1.RoleRef{Kind: "Role", Name: roleObj.Name},
		}
		rb.Kind = "RoleBinding"
		rb.APIVersion = "rbac.authorization.k8s.io/v1"

		roleBindings = append(roleBindings, rb)

	}

	return account, roles, roleBindings, nil
}
func getAllNodes(service types.Service, ret types.StatusRequest, cpContext *core.Context) (types.ResponseRequest, error) {
	var nodes v1.Node
	nodes.Kind = "Node"
	nodes.APIVersion = "v1"
	var serviceOutput types.ServiceOutput
	serviceOutput.Services.Nodes = append(serviceOutput.Services.Nodes, nodes)

	if pId, ok := cpContext.Keys["project_id"]; ok {
		serviceOutput.ProjectId = fmt.Sprintf("%v", pId)
	}
	cpContext.Keys["service_type"] = "Node"
	x, err := json.Marshal(serviceOutput)
	if err != nil {
		return types.ResponseRequest{}, err
	}
	utils.Info.Println("kubernetes request payload", string(x))
	cpContext.Keys["service_type"] = "Node"
	resp, res := GetFromKube(x, serviceOutput.ProjectId, ret, "GET", cpContext)
	if resp.Reason != "" {
		return types.ResponseRequest{}, errors.New(resp.Reason)
	}
	return res, nil
}
func patchNodes(service types.Service, res types.ResponseRequest, ret types.StatusRequest, cpContext *core.Context) error {
	sericeAttrinutes := make(map[string]interface{})
	byteData, err := json.Marshal(service.ServiceAttributes)
	if err != nil {
		return err
	}
	err = json.Unmarshal(byteData, &sericeAttrinutes)
	if err != nil {
		return err
	}
	var existingLabel []string
	byteData, err = json.Marshal(sericeAttrinutes["nodepool"])
	if err != nil {
		return err
	}
	err = json.Unmarshal(byteData, &existingLabel)
	if err != nil {
		return err
	}

	nodeLabel := make(map[string]string)
	byteData, err = json.Marshal(sericeAttrinutes["nodelabel"])
	if err != nil {
		return err
	}
	err = json.Unmarshal(byteData, &nodeLabel)
	if err != nil {
		return err
	}
	if len(existingLabel) == 0 {
		return errors.New("can not find nodepool in request")
	}
	if nodeLabel == nil {
		return errors.New("no label to add")
	}
	var finalNodes []v1.Node
	for i := 0; i < len(res.Service.Nodes[0].Nodes.Items); i++ {

		nl := res.Service.Nodes[0].Nodes.Items[i].Labels
		for j := 0; j < len(existingLabel); j++ {
			if existingLabel[j] == nl["nodepool"] {
				finalNodes = append(finalNodes, res.Service.Nodes[0].Nodes.Items[i])
			}
		}

	}

	var temp4 types.ServiceOutput
	k := 0
	for i := 0; i < len(finalNodes); i++ {

		temp4.Services.Nodes = append(temp4.Services.Nodes, v1.Node{})
		temp4.Services.Nodes[k].ObjectMeta = metav1.ObjectMeta{
			Name:        finalNodes[i].Name,
			UID:         finalNodes[i].UID,
			Generation:  0,
			Labels:      nodeLabel,
			ClusterName: "",
		}

		temp4.Services.Nodes[k].Kind = "Node"
		temp4.Services.Nodes[k].APIVersion = "v1"
		k++

	}
	if pId, ok := cpContext.Keys["project_id"]; ok {
		temp4.ProjectId = fmt.Sprintf("%v", pId)
	}
	byteData, err = json.Marshal(temp4)
	if err != nil {
		return err
	}
	utils.Info.Println("kubernetes request payload", string(byteData))

	resp := ForwardToKube(byteData, temp4.ProjectId, "PATCH", ret, cpContext)

	if resp.Reason != "" {
		return errors.New(resp.Reason)
	}
	return nil
}
func DeployIstio(input types.ServiceInput, requestType string, cpContext *core.Context) types.StatusRequest {

	var ret types.StatusRequest
	ret.ID = input.SolutionInfo.Service.ID
	ret.ServiceId = input.SolutionInfo.Service.ID
	ret.Name = input.SolutionInfo.Service.Name

	var finalObj types.ServiceOutput
	if input.Creds.KubernetesURL != "" {
		finalObj.ClusterInfo.KubernetesURL = input.Creds.KubernetesURL
		finalObj.ClusterInfo.KubernetesUsername = input.Creds.KubernetesUsername
		finalObj.ClusterInfo.KubernetesPassword = input.Creds.KubernetesPassword
	}

	finalObj.ProjectId = input.ProjectId
	//for _,service :=range input.SolutionInfo.Service{
	service := input.SolutionInfo.Service
	//**Making Service Object*//

	respons, err := CheckGateway(service)
	if err != nil {
		utils.Info.Println("There is error in deployment")
		ret.Status = append(ret.Status, "failed")
		ret.Reason = "Not a valid Istio Object. Error : " + err.Error()
		if requestType != "GET" {
			typeArray := []string{"backendLogging", "frontendLogging"}
			cpContext.SendLog(ret.Reason, constants.LOGGING_LEVEL_ERROR, typeArray)
		}
		return ret
	}
	if respons.Spec != nil && respons.Metadata != nil {

		if strings.EqualFold(requestType, "patch") {
			ok, err := notAlreadyExistIstioObject(respons, cpContext, ret)
			if err != nil {
				typeArray := []string{"frontendLogging"}
				cpContext.SendLog(err.Error(), constants.LOGGING_LEVEL_ERROR, typeArray)
				return ret
			}
			if ok {
				finalObj.Services.Istio = append(finalObj.Services.Istio, respons)
			}

		} else {
			finalObj.Services.Istio = append(finalObj.Services.Istio, respons)
		}
	}

	if service.ServiceType == "mesh" || service.ServiceType == "other" {

		respnse, err := getIstioObject(service)
		if err != nil {
			utils.Info.Println("There is error in deployment")
			ret.Status = append(ret.Status, "failed")
			ret.Reason = "Not a valid Istio Object. Error : " + err.Error()
			if requestType != "GET" {
				typeArray := []string{"backendLogging", "frontendLogging"}
				cpContext.SendLog(ret.Reason, constants.LOGGING_LEVEL_ERROR, typeArray)

			}
			return ret
		}
		if strings.EqualFold(requestType, "patch") {
			for _, each := range respnse {
				if each.Spec != nil && each.Metadata != nil {

					ok, err := notAlreadyExistIstioObject(each, cpContext, ret)
					if err != nil {
						typeArray := []string{"frontendLogging"}
						cpContext.SendLog(err.Error(), constants.LOGGING_LEVEL_ERROR, typeArray)
						return ret
					}
					if ok {
						finalObj.Services.Istio = append(finalObj.Services.Istio, each)
					}

				} else {
					finalObj.Services.Istio = append(finalObj.Services.Istio, each)
				}
			}
		} else {
			finalObj.Services.Istio = append(finalObj.Services.Istio, respnse...)
		}

	}

	if service.ServiceType == "node" {
		res, err := getAllNodes(service, ret, cpContext)
		if err != nil {
			ret.Status = append(ret.Status, "failed")
			ret.Reason = "Can Not Get Nodes : " + err.Error()
			if requestType != "GET" {
				typeArray := []string{"backendLogging", "frontendLogging"}
				cpContext.SendLog(ret.Reason, constants.LOGGING_LEVEL_ERROR, typeArray)

			}
			return ret
		}
		err = patchNodes(service, res, ret, cpContext)
		if err != nil {
			ret.Status = append(ret.Status, "failed")
			ret.Reason = "Can Not Patch Nodes : " + err.Error()
			if requestType != "GET" {
				typeArray := []string{"backendLogging", "frontendLogging"}
				cpContext.SendLog(ret.Reason, constants.LOGGING_LEVEL_ERROR, typeArray)

			}
			return ret
		}

	}

	if service.ServiceType == "secrets" {
		if secret, err := getSecretObject(service); err != nil {
			utils.Info.Println("There is error in deployment")
			ret.Status = append(ret.Status, "failed")
			ret.Reason = "Not a valid Secret Object. Error : " + err.Error()
			if requestType != "GET" {
				typeArray := []string{"backendLogging", "frontendLogging"}
				cpContext.SendLog(ret.Reason, constants.LOGGING_LEVEL_ERROR, typeArray)
			}
			return ret
		} else {
			finalObj.Services.Secrets = append(finalObj.Services.Secrets, secret)
		}
	} else if service.ServiceType == "configmap" {
		if configmap, err := getConfigMapObject(service); err != nil {
			utils.Info.Println("There is error in deployment")
			ret.Status = append(ret.Status, "failed")
			ret.Reason = "Not a valid configMap Object. Error : " + err.Error()
			if requestType != "GET" {
				typeArray := []string{"backendLogging", "frontendLogging"}
				cpContext.SendLog(ret.Reason, constants.LOGGING_LEVEL_ERROR, typeArray)
			}
			return ret
		} else {
			finalObj.Services.ConfigMap = append(finalObj.Services.ConfigMap, *configmap)
		}
	}

	secret, exists := CreateDockerCfgSecret(service, input.ProjectId, cpContext)
	if exists {
		finalObj.Services.Secrets = append(finalObj.Services.Secrets, secret)
	}

	if service.ServiceType == "volume" {

		byteData, _ := json.Marshal(service.ServiceAttributes)
		var attributes types.VolumeAttributes
		err := json.Unmarshal(byteData, &attributes)
		fmt.Println("zunnoarinn" + attributes.Volume.Name)
		if err == nil && attributes.Volume.Name != "" {
			//Creating a new storage-class and persistent-volume-claim for each volume
			attributes.Volume.Namespace = "default"
			if service.Namespace != "" {
				attributes.Volume.Namespace = service.Namespace
			}

			finalObj.Services.StorageClasses = append(finalObj.Services.StorageClasses, volumes.ProvisionStorageClass(attributes.Volume))
			temppvc := volumes.ProvisionVolumeClaim(attributes.Volume)
			temppvc.Namespace = service.Namespace
			if service.Namespace == "" {
				temppvc.Namespace = "default"
			}
			finalObj.Services.PersistentVolumeClaims = append(finalObj.Services.PersistentVolumeClaims, temppvc)
		}
	} else if service.ServiceType == "container" {
		switch service.SubType {
		case "hpa":
			hpa, err := getHPAObject(service)
			if err != nil {
				ret.Status = append(ret.Status, "failed")
				ret.Reason = "Not a valid Deployment Object. Error : " + err.Error()
				if requestType != "GET" {
					typeArray := []string{"frontendLogging"}
					cpContext.SendLog(ret.Reason, constants.LOGGING_LEVEL_ERROR, typeArray)
				}

				return ret
			}
			hpa.Kind = "HorizontalPodAutoscaler"
			hpa.APIVersion = "autoscaling/v2beta2"
			finalObj.Services.HPA = append(finalObj.Services.HPA, hpa)

		case "deployment":
			deployment, err := getDeploymentObject(service)
			if err != nil {
				ret.Status = append(ret.Status, "failed")
				ret.Reason = "Not a valid Deployment Object. Error : " + err.Error()
				if requestType != "GET" {
					typeArray := []string{"frontendLogging"}
					cpContext.SendLog(ret.Reason, constants.LOGGING_LEVEL_ERROR, typeArray)
				}
				return ret
			}
			if exists {
				//Assigning Secret
				deployment.Spec.Template.Spec.ImagePullSecrets = append(deployment.Spec.Template.Spec.ImagePullSecrets, v1.LocalObjectReference{Name: secret.ObjectMeta.Name})
			}

			//secrets

			//Attaching persistent volumes if any in two-steps
			//Mounting each volume to container and adding corresponding volume to pod
			if len(deployment.Spec.Template.Spec.Containers) > 0 {
				byteData, _ := json.Marshal(service.ServiceAttributes)
				var attributes types.VolumeAttributesList
				err = json.Unmarshal(byteData, &attributes)

				if err == nil && len(attributes.Volume) > 0 {

					deployment.Spec.Template.Spec.Containers[0].VolumeMounts = volumes.GenerateVolumeMounts(attributes.Volume)
					deployment.Spec.Template.Spec.Volumes = volumes.GeneratePodVolumes(attributes.Volume)
				}
			}
			utils.Info.Println(deployment.Name)
			finalObj.Services.Deployments = append(finalObj.Services.Deployments, deployment)
			//add rbac classes

			byteData, _ := json.Marshal(service.ServiceAttributes)
			var serviceAttr types.DockerServiceAttributes
			json.Unmarshal(byteData, &serviceAttr)
			utils.Info.Println("** rbac params **")
			if serviceAttr.RbacRoles != nil {
				utils.Info.Println(len(serviceAttr.RbacRoles))
			}
			if serviceAttr.IstioRoles != nil {
				utils.Info.Println(len(serviceAttr.IstioRoles))
			}
			if serviceAttr.IsRbac {
				utils.Info.Println("** rbac is enabled **")
				if serviceAttr.RbacRoles != nil {
					if len(serviceAttr.RbacRoles) > 0 {
						serviceAccount, roles, roleBindings, err := getRbacObjects(serviceAttr, service.Name, service.Namespace)
						if err != nil {
							ret.Status = append(ret.Status, "failed")
							ret.Reason = "Not a valid rbac Object. Error : " + err.Error()
							if requestType != "GET" {
								typeArray := []string{"backendLogging", "frontendLogging"}
								cpContext.SendLog(ret.Reason, constants.LOGGING_LEVEL_ERROR, typeArray)
							}
							return ret
						}

						//add service account
						finalObj.Services.ServiceAccountClasses = append(finalObj.Services.ServiceAccountClasses, serviceAccount)

						// add roles and role bindings
						for _, role := range roles {
							finalObj.Services.RoleClasses = append(finalObj.Services.RoleClasses, role)
						}

						for _, roleBinding := range roleBindings {
							finalObj.Services.RoleBindingClasses = append(finalObj.Services.RoleBindingClasses, roleBinding)
						}
					}
				}

				if serviceAttr.IstioRoles != nil && len(serviceAttr.IstioRoles) > 0 {
					istioObjects, err := getIstioRbacObjects(serviceAttr, service.Name, service.Namespace)
					if err != nil {
						ret.Status = append(ret.Status, "failed")
						ret.Reason = "Not a valid rbac Object. Error : " + err.Error()
						if requestType != "GET" {
							typeArray := []string{"backendLogging", "frontendLogging"}
							cpContext.SendLog(ret.Reason, constants.LOGGING_LEVEL_ERROR, typeArray)
						}
						return ret
					}

					utils.Info.Println("isto rbac object's kinds")
					for _, istioObj := range istioObjects {
						utils.Info.Println(istioObj.Kind)
						finalObj.Services.Istio = append(finalObj.Services.Istio, istioObj)
					}
					utils.Info.Println("")
				}
			}

		case "daemonset":
			daemonset, err := getDaemonSetObject(service)
			if err != nil {
				ret.Status = append(ret.Status, "failed")
				ret.Reason = "Not a valid DaemonSet Object. Error : " + err.Error()
				if requestType != "GET" {
					typeArray := []string{"backendLogging", "frontendLogging"}
					cpContext.SendLog(ret.Reason, constants.LOGGING_LEVEL_ERROR, typeArray)

				}
				return ret
			}
			if exists {
				//Assigning Secret
				daemonset.Spec.Template.Spec.ImagePullSecrets = append(daemonset.Spec.Template.Spec.ImagePullSecrets, v1.LocalObjectReference{Name: secret.ObjectMeta.Name})
			}
			finalObj.Services.DaemonSets = append(finalObj.Services.DaemonSets, daemonset)

			//add rbac classes

			byteData, _ := json.Marshal(service)
			var serviceAttr types.DockerServiceAttributes
			json.Unmarshal(byteData, &serviceAttr)

			utils.Info.Println("** rbac params **")
			if serviceAttr.RbacRoles != nil {
				utils.Info.Println(len(serviceAttr.RbacRoles))
			}
			if serviceAttr.IstioRoles != nil {
				utils.Info.Println(len(serviceAttr.IstioRoles))
			}
			if serviceAttr.IsRbac {
				utils.Info.Println("** rbac is enabled **")
				if serviceAttr.RbacRoles != nil && len(serviceAttr.RbacRoles) > 0 {
					serviceAccount, roles, roleBindings, err := getRbacObjects(serviceAttr, service.Name, service.Namespace)
					if err != nil {
						ret.Status = append(ret.Status, "failed")
						ret.Reason = "Not a valid rbac Object. Error : " + err.Error()
						if requestType != "GET" {
							typeArray := []string{"backendLogging", "frontendLogging"}
							cpContext.SendLog(ret.Reason, constants.LOGGING_LEVEL_ERROR, typeArray)

						}
						return ret
					}

					//add service account
					finalObj.Services.ServiceAccountClasses = append(finalObj.Services.ServiceAccountClasses, serviceAccount)

					// add roles and role bindings
					for _, role := range roles {
						finalObj.Services.RoleClasses = append(finalObj.Services.RoleClasses, role)
					}

					for _, roleBinding := range roleBindings {
						finalObj.Services.RoleBindingClasses = append(finalObj.Services.RoleBindingClasses, roleBinding)
					}
				}

				if serviceAttr.IstioRoles != nil && len(serviceAttr.IstioRoles) > 0 {
					istioObjects, err := getIstioRbacObjects(serviceAttr, service.Name, service.Namespace)
					if err != nil {
						ret.Status = append(ret.Status, "failed")
						ret.Reason = "Not a valid rbac Object. Error : " + err.Error()
						if requestType != "GET" {
							typeArray := []string{"backendLogging", "frontendLogging"}
							cpContext.SendLog(ret.Reason, constants.LOGGING_LEVEL_ERROR, typeArray)

						}
						return ret
					}
					for _, istioObj := range istioObjects {
						finalObj.Services.Istio = append(finalObj.Services.Istio, istioObj)
					}
				}
			}

		case "cronjob":
			cronjob, err := getCronJobObject(service)
			if err != nil {
				ret.Status = append(ret.Status, "failed")
				ret.Reason = "Not a valid CronJob Object. Error : " + err.Error()
				if requestType != "GET" {
					typeArray := []string{"backendLogging", "frontendLogging"}
					cpContext.SendLog(ret.Reason, constants.LOGGING_LEVEL_ERROR, typeArray)

				}
				return ret
			}
			if exists {
				//Assigning Secret
				cronjob.Spec.JobTemplate.Spec.Template.Spec.ImagePullSecrets = append(cronjob.Spec.JobTemplate.Spec.Template.Spec.ImagePullSecrets, v1.LocalObjectReference{Name: secret.ObjectMeta.Name})
			}
			finalObj.Services.CronJobs = append(finalObj.Services.CronJobs, cronjob)

			//add rbac classes

			byteData, _ := json.Marshal(service)
			var serviceAttr types.DockerServiceAttributes
			json.Unmarshal(byteData, &serviceAttr)

			if serviceAttr.IsRbac {
				if serviceAttr.RbacRoles != nil {
					serviceAccount, roles, roleBindings, err := getRbacObjects(serviceAttr, service.Name, service.Namespace)
					if err != nil {
						ret.Status = append(ret.Status, "failed")
						ret.Reason = "Not a valid rbac Object. Error : " + err.Error()
						if requestType != "GET" {
							typeArray := []string{"backendLogging", "frontendLogging"}
							cpContext.SendLog(ret.Reason, constants.LOGGING_LEVEL_ERROR, typeArray)

						}
						return ret
					}

					//add service account
					finalObj.Services.ServiceAccountClasses = append(finalObj.Services.ServiceAccountClasses, serviceAccount)

					// add roles and role bindings
					for _, role := range roles {
						finalObj.Services.RoleClasses = append(finalObj.Services.RoleClasses, role)
					}

					for _, roleBinding := range roleBindings {
						finalObj.Services.RoleBindingClasses = append(finalObj.Services.RoleBindingClasses, roleBinding)
					}
				}
			}

		case "job":
			job, err := getJobObject(service)
			if err != nil {
				ret.Status = append(ret.Status, "failed")
				ret.Reason = "Not a valid Job Object. Error : " + err.Error()
				if requestType != "GET" {
					typeArray := []string{"backendLogging", "frontendLogging"}
					cpContext.SendLog(ret.Reason, constants.LOGGING_LEVEL_ERROR, typeArray)

				}
				return ret
			}
			if exists {
				//Assigning Secret
				job.Spec.Template.Spec.ImagePullSecrets = append(job.Spec.Template.Spec.ImagePullSecrets, v1.LocalObjectReference{Name: secret.ObjectMeta.Name})
			}
			finalObj.Services.Jobs = append(finalObj.Services.Jobs, job)

			//add rbac classes

			byteData, _ := json.Marshal(service)
			var serviceAttr types.DockerServiceAttributes
			json.Unmarshal(byteData, &serviceAttr)

			if serviceAttr.IsRbac {
				if serviceAttr.RbacRoles != nil {
					serviceAccount, roles, roleBindings, err := getRbacObjects(serviceAttr, service.Name, service.Namespace)
					if err != nil {
						ret.Status = append(ret.Status, "failed")
						ret.Reason = "Not a valid rbac Object. Error : " + err.Error()
						if requestType != "GET" {
							typeArray := []string{"backendLogging", "frontendLogging"}
							cpContext.SendLog(ret.Reason, constants.LOGGING_LEVEL_ERROR, typeArray)
						}
						return ret
					}

					//add service account
					finalObj.Services.ServiceAccountClasses = append(finalObj.Services.ServiceAccountClasses, serviceAccount)

					// add roles and role bindings
					for _, role := range roles {
						finalObj.Services.RoleClasses = append(finalObj.Services.RoleClasses, role)
					}

					for _, roleBinding := range roleBindings {
						finalObj.Services.RoleBindingClasses = append(finalObj.Services.RoleBindingClasses, roleBinding)
					}
				}
			}

		case "statefulset":
			statefulset, err := getStatefulSetObject(service)
			if err != nil {
				ret.Status = append(ret.Status, "failed")
				ret.Reason = "Not a valid StatefulSet Object. Error : " + err.Error()
				if requestType != "GET" {
					typeArray := []string{"backendLogging", "frontendLogging"}
					cpContext.SendLog(ret.Reason, constants.LOGGING_LEVEL_ERROR, typeArray)

				}
				return ret
			}
			//Attaching persistent volumes if any in two-steps
			//Mounting each volume to container and adding corresponding volume to pod
			if len(statefulset.Spec.Template.Spec.Containers) > 0 {
				byteData, _ := json.Marshal(service.ServiceAttributes)
				var attributes types.VolumeAttributes
				err = json.Unmarshal(byteData, &attributes)

				if err == nil && attributes.Volume.Name != "" {
					volumesData := []types.Volume{attributes.Volume}
					statefulset.Spec.Template.Spec.Containers[0].VolumeMounts = volumes.GenerateVolumeMounts(volumesData)
					statefulset.Spec.Template.Spec.Volumes = volumes.GeneratePodVolumes(volumesData)
				}
			}
			if exists {
				//Assigning Secret
				statefulset.Spec.Template.Spec.ImagePullSecrets = append(statefulset.Spec.Template.Spec.ImagePullSecrets, v1.LocalObjectReference{Name: secret.ObjectMeta.Name})
			}
			finalObj.Services.StatefulSets = append(finalObj.Services.StatefulSets, statefulset)

			//add rbac classes

			byteData, _ := json.Marshal(service)
			var serviceAttr types.DockerServiceAttributes
			json.Unmarshal(byteData, &serviceAttr)

			utils.Info.Println("** rbac params **")
			if serviceAttr.RbacRoles != nil {
				utils.Info.Println(len(serviceAttr.RbacRoles))
			}
			if serviceAttr.IstioRoles != nil {
				utils.Info.Println(len(serviceAttr.IstioRoles))
			}
			if serviceAttr.IsRbac {
				utils.Info.Println("** rbac is enabled **")
				if serviceAttr.RbacRoles != nil && len(serviceAttr.RbacRoles) > 0 {
					serviceAccount, roles, roleBindings, err := getRbacObjects(serviceAttr, service.Name, service.Namespace)
					if err != nil {
						ret.Status = append(ret.Status, "failed")
						ret.Reason = "Not a valid rbac Object. Error : " + err.Error()
						if requestType != "GET" {
							typeArray := []string{"backendLogging", "frontendLogging"}
							cpContext.SendLog(ret.Reason, constants.LOGGING_LEVEL_ERROR, typeArray)

						}
						return ret
					}

					//add service account
					finalObj.Services.ServiceAccountClasses = append(finalObj.Services.ServiceAccountClasses, serviceAccount)

					// add roles and role bindings
					for _, role := range roles {
						finalObj.Services.RoleClasses = append(finalObj.Services.RoleClasses, role)
					}

					for _, roleBinding := range roleBindings {
						finalObj.Services.RoleBindingClasses = append(finalObj.Services.RoleBindingClasses, roleBinding)
					}
				}

				if serviceAttr.IstioRoles != nil && len(serviceAttr.IstioRoles) > 0 {
					istioObjects, err := getIstioRbacObjects(serviceAttr, service.Name, service.Namespace)
					if err != nil {
						ret.Status = append(ret.Status, "failed")
						ret.Reason = "Not a valid rbac Object. Error : " + err.Error()
						if requestType != "GET" {
							typeArray := []string{"backendLogging", "frontendLogging"}
							cpContext.SendLog(ret.Reason, constants.LOGGING_LEVEL_ERROR, typeArray)

						}
						return ret
					}
					for _, istioObj := range istioObjects {
						finalObj.Services.Istio = append(finalObj.Services.Istio, istioObj)
					}
				}
			}

		}
		//todo: create function and call it
		fmt.Println("Hi: ", requestType, service.GroupId)
		if service.GroupId == "" {
			//Getting Kubernetes Service Object
			serv, err := getServiceObject(service)
			if err != nil {
				ret.Status = append(ret.Status, "failed")
				ret.Reason = "Not a valid Service Object. Error : " + err.Error()
				if requestType != "GET" {
					typeArray := []string{"backendLogging", "frontendLogging"}
					cpContext.SendLog(ret.Reason, constants.LOGGING_LEVEL_ERROR, typeArray)
				}
				return ret
			}
			if serv != nil && serv.Spec.Ports != nil {
				if strings.EqualFold(requestType, "patch") {
					var serviceOutput types.ServiceOutput
					serviceOutput.Services.Kubernetes = append(serviceOutput.Services.Kubernetes, *serv)

					if pId, ok := cpContext.Keys["project_id"]; ok {
						serviceOutput.ProjectId = fmt.Sprintf("%v", pId)
					}
					x, err := json.Marshal(serviceOutput)
					if err != nil {
						ret.Status = append(ret.Status, "failed")
						ret.Reason = "Not a valid service Object. Error : " + err.Error()
						if requestType != "GET" {
							typeArray := []string{"frontendLogging"}
							cpContext.SendLog(ret.Reason, constants.LOGGING_LEVEL_ERROR, typeArray)
						}
						return ret
					}
					utils.Info.Println("kubernetes request payload", string(x))
					resp, res := GetFromKube(x, serviceOutput.ProjectId, ret, "GET", cpContext)
					if len(res.Service.Kubernetes) > 0 && strings.Contains(res.Service.Kubernetes[0].Error, "not found") {
						resp = ForwardToKube(x, serviceOutput.ProjectId, "POST", ret, cpContext)
						if resp.Reason != "" {
							ret.Status = append(ret.Status, "failed")
							ret.Reason = "Not a valid Deployment Object. Error : " + resp.Reason
							if requestType != "GET" {
								typeArray := []string{"frontendLogging"}
								cpContext.SendLog(ret.Reason, constants.LOGGING_LEVEL_ERROR, typeArray)
							}
							return ret
						}
					} else {
						finalObj.Services.Kubernetes = append(finalObj.Services.Kubernetes, *serv)
					}
				} else {
					finalObj.Services.Kubernetes = append(finalObj.Services.Kubernetes, *serv)
				}

			}
		} else if service.GroupId != "" && requestType != "DELETE" {
			serv, err := getServiceObject(service)
			if err != nil {
				ret.Status = append(ret.Status, "failed")
				ret.Reason = "Not a valid Service Object. Error : " + err.Error()
				if requestType != "GET" {
					typeArray := []string{"backendLogging", "frontendLogging"}
					cpContext.SendLog(ret.Reason, constants.LOGGING_LEVEL_ERROR, typeArray)
					//utils.SendLog(ret.Reason, "error", input.ProjectId)
					//cpContext.SendBackendLogs(ret.Reason, constants.LOGGING_LEVEL_ERROR)

				}
				return ret
			}
			if serv != nil && serv.Spec.Ports != nil {
				if strings.EqualFold(requestType, "patch") {
					var serviceOutput types.ServiceOutput
					serviceOutput.Services.Kubernetes = append(serviceOutput.Services.Kubernetes, *serv)

					if pId, ok := cpContext.Keys["project_id"]; ok {
						serviceOutput.ProjectId = fmt.Sprintf("%v", pId)
					}
					x, err := json.Marshal(serviceOutput)
					if err != nil {
						ret.Status = append(ret.Status, "failed")
						ret.Reason = "Not a valid Deployment Object. Error : " + err.Error()
						if requestType != "GET" {
							typeArray := []string{"frontendLogging"}
							cpContext.SendLog(ret.Reason, constants.LOGGING_LEVEL_ERROR, typeArray)
						}
						return ret
					}
					utils.Info.Println("kubernetes request payload", string(x))
					resp, res := GetFromKube(x, serviceOutput.ProjectId, ret, "GET", cpContext)
					if len(res.Service.Kubernetes) > 0 && strings.Contains(res.Service.Kubernetes[0].Error, "not found") {
						resp = ForwardToKube(x, serviceOutput.ProjectId, "POST", ret, cpContext)
						if resp.Reason != "" {
							ret.Status = append(ret.Status, "failed")
							ret.Reason = "Not a valid Deployment Object. Error : " + resp.Reason
							if requestType != "GET" {
								typeArray := []string{"frontendLogging"}
								cpContext.SendLog(ret.Reason, constants.LOGGING_LEVEL_ERROR, typeArray)
							}
							return ret
						}
					} else {
						finalObj.Services.Kubernetes = append(finalObj.Services.Kubernetes, *serv)
					}
				} else {
					finalObj.Services.Kubernetes = append(finalObj.Services.Kubernetes, *serv)
				}

			}
		}
	}

	//Send request to Kubernetes
	x, err := json.Marshal(finalObj)
	if err != nil {
		utils.Info.Println(err)
		ret.Status = append(ret.Status, "failed")
		ret.Reason = "Service Object parsing failed : " + err.Error()
		if requestType != "GET" {
			typeArray := []string{"backendLogging", "frontendLogging"}
			cpContext.SendLog(ret.Reason, constants.LOGGING_LEVEL_ERROR, typeArray)

		}
		return ret
	}
	utils.Info.Printf("kubernetes request payload, requestTyp: %s, data: %s", requestType, string(x))

	if requestType != "POST" {
		ret, resp := GetFromKube(x, input.ProjectId, ret, requestType, cpContext)
		if ret.Reason == "" {
			//Successful in getting object
			if requestType == "GET" {
				if resp.Service.Kubernetes != nil {
					for _, k := range resp.Service.Kubernetes {
						if k.Error != "" {
							ret.Reason = ret.Reason + k.Error
							ret.Status = append(ret.Status, "failed")
							continue
						}
						ret.Status = append(ret.Status, "successful")
					}
				}
				if resp.Service.Deployments != nil {
					for _, d := range resp.Service.Deployments {
						if d.Error != "" {
							ret.Reason = ret.Reason + d.Error
							ret.Status = append(ret.Status, "failed")
							continue
						}
						for _, c := range d.Deployments.Status.Conditions {
							if c.Type == v12.DeploymentAvailable {
								ret.Status = append(ret.Status, "successful")

							} else if c.Type == v12.DeploymentProgressing {
								ret.Status = append(ret.Status, "successful") //will decide later

							} else {
								ret.Status = append(ret.Status, "failed")
								ret.Reason = ret.Reason + c.Reason
							}
						}
					}
				}
				if resp.Service.Istio != nil {
					for _, i := range resp.Service.Istio {
						if i.Error != "" {
							ret.Reason = ret.Reason + i.Error
							ret.Status = append(ret.Status, "failed")
							continue
						}
						ret.Status = append(ret.Status, "successful")
					}
				}
				return ret
			} else if requestType == "PATCH" {
				if resp.Service.Kubernetes != nil {
					var services []v1.Service
					for i, k := range resp.Service.Kubernetes {
						if k.Error != "" {
							ret.Reason = ret.Reason + k.Error
							ret.Status = append(ret.Status, "failed")
							continue
						}
						_ = i //Replace 0 with i in case of multiple services
						k.Kubernetes.Spec = finalObj.Services.Kubernetes[0].Spec
						services = append(services, k.Kubernetes)
					}
					finalObj.Services.Kubernetes = services

				}
				if resp.Service.Deployments != nil {
					var deployments []v12.Deployment

					for _, d := range resp.Service.Deployments {
						if d.Error != "" {
							ret.Reason = ret.Reason + d.Error
							ret.Status = append(ret.Status, "failed")
							continue
						}
						d.Deployments.Spec = finalObj.Services.Deployments[0].Spec
						deployments = append(deployments, d.Deployments)
					}
					finalObj.Services.Deployments = deployments
				}
				if resp.Service.Istio != nil {
					var istios []types.IstioObject
					for _, i := range resp.Service.Istio {
						if i.Error != "" {
							ret.Reason = ret.Reason + i.Error
							ret.Status = append(ret.Status, "failed")
							continue
						}
						i.Istio.Spec = finalObj.Services.Istio[0].Spec
						istios = append(istios, i.Istio)
					}
					finalObj.Services.Istio = istios
				}
			}
		} else {
			return ret
		}

	}

	x, err = json.Marshal(finalObj)
	if err != nil {
		utils.Info.Println(err)
		ret.Status = append(ret.Status, "failed")
		ret.Reason = "Service Object parsing failed : " + err.Error()
		if requestType != "GET" {
			typeArray := []string{"backendLogging", "frontendLogging"}
			cpContext.SendLog(ret.Reason, constants.LOGGING_LEVEL_ERROR, typeArray)

		}
		return ret
	}
	if requestType != "GET" {
		//Send failure request
		return ForwardToKube(x, input.ProjectId, requestType, ret, cpContext)
	}
	return ret

}

func GetFromKube(requestBody []byte, env_id string, ret types.StatusRequest, requestType string, cpContext *core.Context) (types.StatusRequest, types.ResponseRequest) {
	url := constants.KubernetesEngineURL
	var res types.ResponseRequest
	if cpContext.Keys["service_type"] == "Node" {
		url += constants.Ksd_Get_Nobe
		cpContext.Keys["service_type"] = ""
	}
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	//Adding Headers
	if cpContext.Exists("project_id") {
		req.Header.Set("project_id", cpContext.GetString("project_id"))
	}

	if cpContext.Exists("company_id") {
		req.Header.Set("company_id", cpContext.GetString("company_id"))
	}
	if cpContext.Exists("user") {
		req.Header.Set("user", cpContext.GetString("user"))
	}
	if cpContext.Exists("token") {
		req.Header.Set("token", cpContext.GetString("token"))
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		utils.Info.Println(err)
		if requestType != "GET" {
			typeArray := []string{"backendLogging", "frontendLogging"}
			cpContext.SendLog("Connection to kubernetes microservice failed "+err.Error(), constants.LOGGING_LEVEL_INFO, typeArray)
		}
		ret.Status = append(ret.Status, "failed")
		ret.Reason = "Connection to kubernetes deployment microservice failed Error : " + err.Error()
		if requestType != "GET" {
			typeArray := []string{"backendLogging", "frontendLogging"}
			cpContext.SendLog(ret.Reason, constants.LOGGING_LEVEL_ERROR, typeArray)

		}

		return ret, res

	} else {
		statusCode := resp.StatusCode

		//Info.Printf("notification status code %d\n", statusCode)
		result, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			ret.Status = append(ret.Status, "failed")
			ret.Reason = "kubernetes deployment microservice Response Parsing failed.Error : " + err.Error()
			if requestType != "GET" {
				typeArray := []string{"backendLogging", "frontendLogging"}
				cpContext.SendLog("Connection to kubernetes microservice failed "+err.Error(), constants.LOGGING_LEVEL_INFO, typeArray)
				cpContext.SendLog(ret.Reason, constants.LOGGING_LEVEL_ERROR, typeArray)
			}
			return ret, res
		} else {

			utils.Info.Println(string(result))
			if requestType != "GET" {
				typeArray := []string{"backendLogging", "frontendLogging"}
				cpContext.SendLog(string(result), constants.LOGGING_LEVEL_INFO, typeArray)

			}
			if statusCode != 200 {
				var resrf types.ResponseServiceRequestFailure
				err = json.Unmarshal(result, &resrf)
				if err != nil {
					ret.Status = append(ret.Status, "failed")
					ret.Reason = "kubernetes deployment microservice Response Parsing failed.Error : " + err.Error()
					if requestType != "GET" {
						typeArray := []string{"backendLogging", "frontendLogging"}
						cpContext.SendLog(ret.Reason, constants.LOGGING_LEVEL_ERROR, typeArray)

					}
					return ret, res
				}
				ret.Status = append(ret.Status, "failed")
				ret.Reason = resrf.Error
				return ret, res
			} else {
				err = json.Unmarshal(result, &res)
				if err != nil {
					ret.Status = append(ret.Status, "failed")
					ret.Reason = "kubernetes deployment microservice Response Parsing failed.Error : " + err.Error()
					if requestType != "GET" {
						typeArray := []string{"backendLogging", "frontendLogging"}
						cpContext.SendLog(ret.Reason, constants.LOGGING_LEVEL_ERROR, typeArray)

					}
					return ret, res

				} else {

					for itr, _ := range res.Service.Nodes {
						err = json.Unmarshal([]byte(res.Service.Nodes[itr].KubeData), &res.Service.Nodes[itr].Nodes)
						if err != nil {
							utils.Error.Println(err.Error())
						}
					}

					for itr, _ := range res.Service.Deployments {
						err = json.Unmarshal([]byte(res.Service.Deployments[itr].KubeData), &res.Service.Deployments[itr].Deployments)
						if err != nil {
							utils.Error.Println(err.Error())
						}
					}

					for itr, _ := range res.Service.Istio {
						err = json.Unmarshal([]byte(res.Service.Istio[itr].KubeData), &res.Service.Istio[itr].Istio)
						if err != nil {
							utils.Error.Println(err.Error())
						}
					}

					for itr, _ := range res.Service.Kubernetes {
						err = json.Unmarshal([]byte(res.Service.Kubernetes[itr].KubeData), &res.Service.Kubernetes[itr].Kubernetes)
						if err != nil {
							utils.Error.Println(err.Error())
						}
					}

				}
				return ret, res
			}
		}
		return ret, res
	}
}
func ForwardToKube(requestBody []byte, env_id string, requestType string, ret types.StatusRequest, cpContext *core.Context) types.StatusRequest {

	url := constants.KubernetesEngineURL
	var res types.KSDResponse
	utils.Info.Println("forward to kube: " + url)
	utils.Info.Println("request type: " + requestType)

	req, err := http.NewRequest(requestType, url, bytes.NewBuffer(requestBody))

	req.Header.Set("Content-Type", "application/json")
	//Adding Headers
	if cpContext.Exists("project_id") {
		req.Header.Set("project_id", cpContext.GetString("project_id"))
	}
	if cpContext.Exists("solutionId") {
		req.Header.Set("solutionId", cpContext.GetString("solutionId"))
	}

	if cpContext.Exists("company_id") {
		req.Header.Set("company_id", cpContext.GetString("company_id"))
	}
	if cpContext.Exists("user") {
		req.Header.Set("user", cpContext.GetString("user"))
	}
	client := &http.Client{}
	//issue here
	resp, err := client.Do(req)
	if err != nil {
		utils.Info.Println(err)
		ret.Status = append(ret.Status, "failed")
		ret.Reason = "Connection to kubernetes deployment microservice failed Error : " + err.Error()

		if requestType != "GET" {
			typeArray := []string{"backendLogging", "frontendLogging"}
			cpContext.SendLog(ret.Reason, constants.LOGGING_LEVEL_ERROR, typeArray)
			cpContext.SendLog("Connection to kubernetes microservice failed "+err.Error(), constants.LOGGING_LEVEL_INFO, typeArray)

		}
		return ret

	} else {
		statusCode := resp.StatusCode

		//Info.Printf("notification status code %d\n", statusCode)
		result, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			utils.Info.Println(err)
			ret.Status = append(ret.Status, "failed")
			ret.Reason = "kubernetes deployment microservice Response Parsing failed.Error : " + err.Error()
			if requestType != "GET" {

				typeArray := []string{"backendLogging", "frontendLogging"}
				cpContext.SendLog(ret.Reason, constants.LOGGING_LEVEL_INFO, typeArray)
				cpContext.SendLog("Response Parsing failed "+err.Error(), constants.LOGGING_LEVEL_ERROR, typeArray)
			}
			return ret
		} else {
			utils.Info.Println(string(result))
			if requestType != "GET" {
				var rt map[string]interface{}
				utils.Error.Println(json.Unmarshal(result, &rt))
				rr, err := json.MarshalIndent(rt, "", "	")
				utils.Error.Println(json.Unmarshal(result, &res))
				for _, each := range res.Service.Kubernetes {
					if strings.Contains(each.Error, "already exists") {
						continue
					} else if each.Error != "" {
						ret.Status = append(ret.Status, "failed")
						ret.Reason = ret.Reason + ";" + each.Error
					} else if each.Error == "" && each.Kubernetes == nil {
						ret.Status = append(ret.Status, "failed")
						ret.Reason = ret.Reason + "; error from ksd data null error null"
					}
				}
				for _, each := range res.Service.Nodes {
					if strings.Contains(each.Error, "already exists") {
						continue
					} else if each.Error != "" {
						ret.Status = append(ret.Status, "failed")
						ret.Reason = ret.Reason + ";" + each.Error
					} else if each.Error == "" && each.Nodes == nil {
						ret.Status = append(ret.Status, "failed")
						ret.Reason = ret.Reason + "; error from ksd data null error null"
					}

				}
				for _, each := range res.Service.HPAS {
					if strings.Contains(each.Error, "already exists") {
						continue
					} else if each.Error != "" {
						ret.Status = append(ret.Status, "failed")
						ret.Reason = ret.Reason + ";" + each.Error
					} else if each.Error == "" && each.Hpas == nil {
						ret.Status = append(ret.Status, "failed")
						ret.Reason = ret.Reason + "; error from ksd data null error null"
					}
				}

				for _, each := range res.Service.Istio {
					if strings.Contains(each.Error, "already exists") {
						continue
					} else if each.Error != "" {
						ret.Status = append(ret.Status, "failed")
						ret.Reason = ret.Reason + ";" + each.Error
					} else if each.Error == "" && each.Istio == nil {
						ret.Status = append(ret.Status, "failed")
						ret.Reason = ret.Reason + "; error from ksd data null error null"
					}
				}
				for _, each := range res.Service.Deployments {
					if strings.Contains(each.Error, "already exists") {
						continue
					} else if each.Error != "" {
						ret.Status = append(ret.Status, "failed")
						ret.Reason = ret.Reason + ";" + each.Error
					} else if each.Error == "" && each.Deployments == nil {
						ret.Status = append(ret.Status, "failed")
						ret.Reason = ret.Reason + "; error from ksd data null error null"
					}
				}
				for _, each := range res.Service.StatefulSets {
					if strings.Contains(each.Error, "already exists") {
						continue
					} else if each.Error != "" {
						ret.Status = append(ret.Status, "failed")
						ret.Reason = ret.Reason + ";" + each.Error
					} else if each.Error == "" && each.StatefulSets == nil {
						ret.Status = append(ret.Status, "failed")
						ret.Reason = ret.Reason + "; error from ksd data null error null"
					}
				}
				for _, each := range res.Service.DaemonSets {
					if strings.Contains(each.Error, "already exists") {
						continue
					} else if each.Error != "" {
						ret.Status = append(ret.Status, "failed")
						ret.Reason = ret.Reason + ";" + each.Error
					} else if each.Error == "" && each.DaemonSets == nil {
						ret.Status = append(ret.Status, "failed")
						ret.Reason = ret.Reason + "; error from ksd data null error null"
					}
				}
				for _, each := range res.Service.Jobs {
					if strings.Contains(each.Error, "already exists") {
						continue
					} else if each.Error != "" {
						ret.Status = append(ret.Status, "failed")
						ret.Reason = ret.Reason + ";" + each.Error
					} else if each.Error == "" && each.Jobs == nil {
						ret.Status = append(ret.Status, "failed")
						ret.Reason = ret.Reason + "; error from ksd data null error null"
					}
				}
				for _, each := range res.Service.CronJobs {
					if strings.Contains(each.Error, "already exists") {
						continue
					} else if each.Error != "" {
						ret.Status = append(ret.Status, "failed")
						ret.Reason = ret.Reason + ";" + each.Error
					} else if each.Error == "" && each.CronJobs == nil {
						ret.Status = append(ret.Status, "failed")
						ret.Reason = ret.Reason + "; error from ksd data null error null"
					}
				}

				utils.Info.Printf("%s", rr)
				utils.Error.Println(err)
				typeArray := []string{"backendLogging", "frontendLogging"}
				cpContext.SendLog(string(rr), constants.LOGGING_LEVEL_INFO, typeArray)

			}
			if statusCode != 200 {
				var resrf types.ResponseServiceRequestFailure
				err = json.Unmarshal(result, &resrf)
				if err != nil {
					ret.Status = append(ret.Status, "failed")
					ret.Reason = "kubernetes deployment microservice Response Parsing failed.Error : " + err.Error()
					if requestType != "GET" {
						typeArray := []string{"backendLogging", "frontendLogging"}
						cpContext.SendLog(string(result), constants.LOGGING_LEVEL_ERROR, typeArray)
					}
					return ret
				}
				ret.Status = append(ret.Status, "failed")
				ret.Reason = resrf.Error
				return ret
			}
			ret.Status = append(ret.Status, "successful")
		}
	}
	return ret
}

func ServiceRequest(w http.ResponseWriter, r *http.Request) {

	/* Logging New Architecture */

	cpContext := new(core.Context)
	err := cpContext.ReadLoggingParameters(r)
	if err != nil {
		utils.Error.Println(err)

		//http.Error(w, err.Error(), 500)

	} else {
		cpContext.InitializeLogger(r.Host, r.Method, r.URL.Host, "")
	}
	//Logging Initializations End

	utils.Info.Println(r.Body)
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	utils.Info.Println(string(b))
	// Unmarshal
	var input types.ServiceInput
	err = json.Unmarshal(b, &input)
	if err != nil {
		http.Error(w, err.Error(), 500)
		cpContext.SendBackendLogs(err.Error(), constants.LOGGING_LEVEL_ERROR)
		return
	}

	cpContext.AddProjectId(input.ProjectId)
	var notification types.Notifier
	notification.Component = "Service"
	notification.Id = input.SolutionInfo.Service.ID

	var status types.StatusRequest
	status.ID = input.SolutionInfo.Service.ID
	status.Name = input.SolutionInfo.Service.Name

	result := DeployIstio(input, r.Method, cpContext)

	inProgress := false
	failed := false
	for _, status := range result.Status {
		if status == "in progress" {
			inProgress = true
		}
		if status == "failed" {
			failed = true
		}
	}
	if inProgress {
		result.StatusF = "in progress"
	}
	if failed {
		result.StatusF = "failed"
	}
	if !failed && !inProgress {
		result.StatusF = "successful"
	}
	if result.Reason != "" {
		result.StatusF = "failed"
		x, err := json.Marshal(result)
		if err == nil {
			w.Write(x)
		}
		notification.Status = "fail"
	} else {

		utils.Info.Println("Deployment Successful")
		notification.Status = "success"
		x, err := json.Marshal(result)
		if err == nil {
			w.Write(x)
		}

	}
	b, err1 := json.Marshal(notification)
	if err1 != nil {
		utils.Info.Println(err1)
		utils.Error.Println("Notification Parsing failed")
	} else {
		if r.Method != "GET" {
			Notifier.Notify(input.ProjectId, string(b))
			utils.Info.Println(string(b))
		}
	}
}

const (
	SecretKind = "Secret"
)

func getSecretObject(service types.Service) (*v1.Secret, error) {
	switch service.SubType {
	case "Opaque":
		if secret, ok := CreateOpaqueSecret(service); ok {
			return secret, nil
		} else {
			return nil, errors.New("Something went wrong in getting secret object")
		}
	case "TLS":
		if secret, ok := CreateTLSSecret(service); ok {
			return secret, nil
		} else {
			return nil, errors.New("Something went wrong in getting secret object")
		}
	}
	return nil, errors.New("Something went wrong")
}

// this will be used by revions to pull the image from registry
func CreateDockerCfgSecret(service types.Service, projectId string, cpContext *core.Context) (v1.Secret, bool) {

	byteData, _ := json.Marshal(service.ServiceAttributes)
	var serviceAttr types.DockerServiceAttributes
	err := json.Unmarshal(byteData, &serviceAttr)
	if err != nil {
		return v1.Secret{}, false
	}
	if serviceAttr.ImageRepositoryConfigurations == nil {
		return v1.Secret{}, false
	}
	profileId := serviceAttr.ImageRepositoryConfigurations.Profile
	if profileId != "" {
		var vault types.VaultCredentialsConfigurations
		req, err := http.Get(constants.VaultURL + constants.VAULT_BACKEND + profileId)
		if err == nil {
			result, err := ioutil.ReadAll(req.Body)
			if err == nil {
				err = json.Unmarshal(result, &vault)
				typeArray := []string{"frontendLogging"}
				cpContext.SendLog("creds fetched "+vault.Credentials.Username+":"+vault.Credentials.Password, constants.LOGGING_LEVEL_ERROR, typeArray)

				if err == nil {
					if vault.Credentials.Username != "" && vault.Credentials.Password != "" {
						serviceAttr.ImageRepositoryConfigurations.Credentials.Username = vault.Credentials.Username
						serviceAttr.ImageRepositoryConfigurations.Credentials.Password = vault.Credentials.Password
					}
				}
			}
		} else {
			typeArray := []string{"frontendLogging"}
			cpContext.SendLog("vault fetch failure "+err.Error(), "error", typeArray)

		}

	} else {
		//typeArray := []string{"frontendLogging"}
		//cpContext.SendLog("profile id empty ", "error", typeArray)

	}
	if serviceAttr.ImageRepositoryConfigurations.Credentials.Username == "" || serviceAttr.ImageRepositoryConfigurations.Credentials.Password == "" {
		return v1.Secret{}, false
	}
	secret := v1.Secret{}

	typeMeta := metav1.TypeMeta{
		Kind:       SecretKind,
		APIVersion: v1.SchemeGroupVersion.String(),
	}
	objectMeta := metav1.ObjectMeta{
		Name:      service.Name + "-cfg-secret",
		Namespace: service.Namespace,
	}

	username := serviceAttr.ImageRepositoryConfigurations.Credentials.Username
	password := serviceAttr.ImageRepositoryConfigurations.Credentials.Password
	email := "my-account-email@address.com"
	server := serviceAttr.ImageName

	tokens := strings.Split(server, "/")
	registry := tokens[0]
	if strings.TrimSpace(registry) == "docker.io" {
		registry = "https://index.docker.io/v1/"
	}
	dockerConf := map[string]map[string]map[string]string{
		"auths": {
			registry: {
				"username": username,
				"password": password,
				"email":    email,
				"auth":     base64.StdEncoding.EncodeToString([]byte(username + ":" + password)),
			},
		},
	}

	dockerConfMarshaled, _ := json.Marshal(dockerConf)

	data := map[string][]byte{
		".dockercfg": dockerConfMarshaled,
	}

	secret.TypeMeta = typeMeta
	secret.ObjectMeta = objectMeta
	secret.Data = data
	secret.Type = v1.SecretTypeDockercfg

	return secret, true
}

func CreateOpaqueSecret(service types.Service) (*v1.Secret, bool) {

	byteData, _ := json.Marshal(service.ServiceAttributes)
	var serviceAttr types.KubernetesSecret
	if err := json.Unmarshal(byteData, &serviceAttr); err != nil {
		utils.Error.Println(err)
		return nil, false
	}

	secret := v1.Secret{}

	typeMeta := metav1.TypeMeta{
		Kind:       SecretKind,
		APIVersion: v1.SchemeGroupVersion.String(),
	}
	objectMeta := metav1.ObjectMeta{
		Name:      service.ID,
		Namespace: service.Namespace,
	}

	secret.TypeMeta = typeMeta
	secret.ObjectMeta = objectMeta
	secret.Data = make(map[string][]byte)
	if serviceAttr.Data != nil {
		for key, value := range serviceAttr.Data {
			if decoded_value, err := base64.StdEncoding.DecodeString(value); err != nil {
				utils.Error.Println(err)
				secret.Data[key] = []byte(value)
			} else {
				secret.Data[key] = decoded_value
			}
			//secret.Data[key] = []byte(value)
		}
	}
	if serviceAttr.StringData != nil {
		secret.StringData = serviceAttr.StringData
	}
	secret.Type = v1.SecretTypeOpaque

	return &secret, true
}

func CreateTLSSecret(service types.Service) (*v1.Secret, bool) {

	byteData, _ := json.Marshal(service.ServiceAttributes)
	var serviceAttr types.KubernetesSecret
	if err := json.Unmarshal(byteData, &serviceAttr); err != nil {
		utils.Error.Println(err)
		return nil, false
	}

	secret := v1.Secret{}

	typeMeta := metav1.TypeMeta{
		Kind:       SecretKind,
		APIVersion: v1.SchemeGroupVersion.String(),
	}
	objectMeta := metav1.ObjectMeta{
		Name:      service.ID,
		Namespace: service.Namespace,
	}

	secret.TypeMeta = typeMeta
	secret.ObjectMeta = objectMeta
	secret.Data = make(map[string][]byte)
	if serviceAttr.Data != nil {
		for key, value := range serviceAttr.Data {
			if decoded_value, err := base64.StdEncoding.DecodeString(value); err != nil {
				utils.Error.Println(err)
				secret.Data[key] = []byte(value)
			} else {
				secret.Data[key] = decoded_value
			}
		}
	}
	if serviceAttr.StringData != nil {
		secret.StringData = serviceAttr.StringData
	}
	secret.Type = v1.SecretTypeTLS

	return &secret, true
}

func putCommandAndArguments(container *v1.Container, command, args []string) error {
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
func putLimitResource(container *v1.Container, limitResources map[types.RecourceType]string) error {
	temp := make(map[v1.ResourceName]resource.Quantity)
	for t, v := range limitResources {
		if t == types.RecourceTypeCpu || t == types.RecourceTypeMemory {
			quantity, err := resource.ParseQuantity(v)
			if err != nil {
				return err
			}
			temp[v1.ResourceName(t)] = quantity
		} else {
			return errors.New("Error Found: Invalid Request Resource Provided. Valid: 'cpu','memory'")
		}
	}

	container.Resources.Limits = temp
	return nil
}
func putRequestResource(container *v1.Container, requestResources map[types.RecourceType]string) error {
	temp := make(map[v1.ResourceName]resource.Quantity)
	for t, v := range requestResources {
		if t == types.RecourceTypeCpu || t == types.RecourceTypeMemory {
			quantity, err := resource.ParseQuantity(v)
			if err != nil {
				return err
			}
			//
			temp[v1.ResourceName(t)] = quantity
		} else {
			return errors.New("Error Found: Invalid Request Resource Provided. Valid: 'cpu','memory'")
		}
	}

	container.Resources.Requests = temp
	return nil
}
func putLivenessProbe(container *v1.Container, prob *types.Probe) error {

	var temp v1.Probe
	if prob != nil {
		if prob.Handler != nil {
			if prob.InitialDelaySeconds != nil {
				temp.InitialDelaySeconds = *prob.InitialDelaySeconds
			}
			if prob.FailureThreshold != nil {
				temp.FailureThreshold = *prob.FailureThreshold
			}
			if prob.PeriodSeconds != nil {
				temp.PeriodSeconds = *prob.PeriodSeconds
			}
			if prob.SuccessThreshold != nil {
				temp.SuccessThreshold = *prob.SuccessThreshold
			}
			if prob.TimeoutSeconds != nil {
				temp.TimeoutSeconds = *prob.TimeoutSeconds
			}
			switch typeHandler := prob.Handler.Type; typeHandler {
			case "exec":
				if prob.Handler.Exec == nil {
					return errors.New("there is no liveness handler of exec type")
				}
				temp.Handler.Exec = &v1.ExecAction{}
				for i := 0; i < len(prob.Handler.Exec.Command); i++ {
					temp.Handler.Exec.Command = append(temp.Handler.Exec.Command, prob.Handler.Exec.Command[i])
				}

			case "httpGet":
				if prob.Handler.HTTPGet == nil {
					return errors.New("there is no liveness handler of httpGet type")
				}
				temp.Handler.HTTPGet = &v1.HTTPGetAction{}
				if prob.Handler.HTTPGet.Port > 0 && prob.Handler.HTTPGet.Port < 65536 {
					if prob.Handler.HTTPGet.Host != nil {
						temp.HTTPGet.Host = *prob.Handler.HTTPGet.Host
					}
					if prob.Handler.HTTPGet.Path != nil {
						temp.HTTPGet.Path = *prob.Handler.HTTPGet.Path

					}
					if prob.Handler.HTTPGet.Scheme != nil {
						if *prob.Handler.HTTPGet.Scheme == types.URISchemeHTTP || *prob.Handler.HTTPGet.Scheme == types.URISchemeHTTPS {

							temp.HTTPGet.Scheme = v1.URIScheme(*prob.Handler.HTTPGet.Scheme)
						} else {
							return errors.New("invalid urischeme ")
						}
					}
					if prob.Handler.HTTPGet.HTTPHeaders != nil {
						temp.HTTPGet.HTTPHeaders = []v1.HTTPHeader{}
						for i := 0; i < len(prob.Handler.HTTPGet.HTTPHeaders); i++ {
							if prob.Handler.HTTPGet.HTTPHeaders[i].Value == nil || prob.Handler.HTTPGet.HTTPHeaders[i].Name == nil {
								return errors.New("http header name and values are required")
							}
							tempheader := v1.HTTPHeader{*prob.Handler.HTTPGet.HTTPHeaders[i].Name, *prob.Handler.HTTPGet.HTTPHeaders[i].Value}
							temp.HTTPGet.HTTPHeaders = append(temp.HTTPGet.HTTPHeaders, tempheader)
						}
					}
					temp.HTTPGet.Port = intstr.FromInt(prob.Handler.HTTPGet.Port)
				} else {
					return errors.New("Invalid Port number for http Get")
				}
			case "tcpSocket":
				if prob.Handler.TCPSocket == nil {
					return errors.New("there is no liveness handler of tcpSocket type")
				}
				temp.Handler.TCPSocket = &v1.TCPSocketAction{}
				if prob.Handler.TCPSocket.Port > 0 && prob.Handler.TCPSocket.Port < 65536 {
					temp.TCPSocket.Port = intstr.FromInt(prob.Handler.TCPSocket.Port)
					if prob.Handler.TCPSocket.Host != nil {
						temp.TCPSocket.Host = *prob.Handler.TCPSocket.Host
					}
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

func putReadinessProbe(container *v1.Container, prob *types.Probe) error {
	var temp v1.Probe
	if prob != nil {
		if prob.Handler != nil {
			if prob.InitialDelaySeconds != nil {
				temp.InitialDelaySeconds = *prob.InitialDelaySeconds
			}
			if prob.FailureThreshold != nil {
				temp.FailureThreshold = *prob.FailureThreshold
			}
			if prob.PeriodSeconds != nil {
				temp.PeriodSeconds = *prob.PeriodSeconds
			}
			if prob.SuccessThreshold != nil {
				temp.SuccessThreshold = *prob.SuccessThreshold
			}
			if prob.TimeoutSeconds != nil {
				temp.TimeoutSeconds = *prob.TimeoutSeconds
			}
			switch typeHandler := prob.Handler.Type; typeHandler {
			case "exec":
				if prob.Handler.Exec == nil {
					return errors.New("there is no readiness handler of exec type")
				}
				temp.Handler.Exec = &v1.ExecAction{}
				for i := 0; i < len(prob.Handler.Exec.Command); i++ {
					temp.Handler.Exec.Command = append(temp.Handler.Exec.Command, prob.Handler.Exec.Command[i])
				}

			case "httpGet":
				if prob.Handler.HTTPGet == nil {
					return errors.New("there is no readiness handler of httpGet type")
				}
				temp.Handler.HTTPGet = &v1.HTTPGetAction{}
				if prob.Handler.HTTPGet.Port > 0 && prob.Handler.HTTPGet.Port < 65536 {
					if prob.Handler.HTTPGet.Host != nil {
						temp.HTTPGet.Host = *prob.Handler.HTTPGet.Host
					}
					if prob.Handler.HTTPGet.Path != nil {
						temp.HTTPGet.Path = *prob.Handler.HTTPGet.Path

					}
					if prob.Handler.HTTPGet.Scheme != nil {
						if *prob.Handler.HTTPGet.Scheme == types.URISchemeHTTP || *prob.Handler.HTTPGet.Scheme == types.URISchemeHTTPS {

							temp.HTTPGet.Scheme = v1.URIScheme(*prob.Handler.HTTPGet.Scheme)
						} else {
							return errors.New("invalid urischeme ")
						}
					}
					if prob.Handler.HTTPGet.HTTPHeaders != nil {
						temp.HTTPGet.HTTPHeaders = []v1.HTTPHeader{}
						for i := 0; i < len(prob.Handler.HTTPGet.HTTPHeaders); i++ {
							if prob.Handler.HTTPGet.HTTPHeaders[i].Value == nil || prob.Handler.HTTPGet.HTTPHeaders[i].Name == nil {
								return errors.New("http header name and values are required")
							}
							tempheader := v1.HTTPHeader{*prob.Handler.HTTPGet.HTTPHeaders[i].Name, *prob.Handler.HTTPGet.HTTPHeaders[i].Value}
							temp.HTTPGet.HTTPHeaders = append(temp.HTTPGet.HTTPHeaders, tempheader)
						}
					}
					temp.HTTPGet.Port = intstr.FromInt(prob.Handler.HTTPGet.Port)
				} else {
					return errors.New("Invalid Port number for http Get")
				}
			case "tcpSocket":
				if prob.Handler.TCPSocket == nil {
					return errors.New("there is no readiness handler of tcpSocket type")
				}
				temp.Handler.TCPSocket = &v1.TCPSocketAction{}
				if prob.Handler.TCPSocket.Port > 0 && prob.Handler.TCPSocket.Port < 65536 {
					temp.TCPSocket.Port = intstr.FromInt(prob.Handler.TCPSocket.Port)
					if prob.Handler.TCPSocket.Host != nil {
						temp.TCPSocket.Host = *prob.Handler.TCPSocket.Host
					}
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

func convertKeys(j json.RawMessage) json.RawMessage {
	m := make(map[string]json.RawMessage)
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		// Not a JSON object
		return j
	}

	for k, v := range m {
		fixed := fixKey(k)
		delete(m, k)
		m[fixed] = convertKeys(v)
	}

	b, err := json.Marshal(m)
	if err != nil {
		return j
	}

	return json.RawMessage(b)
}
func fixKey(key string) string {
	return strcase.ToLowerCamel(key)
}

/*
type tempProbing struct {
	LivenessProbe  *v1.Probe `json:"livenessProbe"`
	ReadinessProbe *v1.Probe `json:"readinessProbe"`
}
*/
func checkRequestIsLessThanLimit(serviceAttr types.DockerServiceAttributes) (error, bool) {
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
func getInitContainers(service types.Service) ([]v1.Container, []string, []string, error) {
	var configMapsArray, secretsArray []string
	fmt.Println(service)
	serviceAttributes := make(map[string]interface{})
	var initContainerServiceAttributes interface{}
	if data, err := json.Marshal(service.ServiceAttributes); err == nil {
		if err = json.Unmarshal(data, &serviceAttributes); err == nil {
			if _, ok := serviceAttributes["init_container"]; ok {
				initContainerServiceAttributes = serviceAttributes["init_container"]
			} else {
				fmt.Println("No Init Containers for " + service.Name)
				return nil, secretsArray, configMapsArray, nil
			}
		} else {
			return nil, secretsArray, configMapsArray, err
		}
	} else {
		return nil, secretsArray, configMapsArray, err
	}

	var container v1.Container
	byteData, _ := json.Marshal(initContainerServiceAttributes)
	var serviceAttr types.DockerServiceAttributes
	var err = json.Unmarshal(byteData, &serviceAttr)
	if err != nil {
		return nil, secretsArray, configMapsArray, err
	}
	if serviceAttr.Name != "" {
		container.Name = serviceAttr.Name
	} else {
		container.Name = "init-container-dummy"
	}
	if err := putCommandAndArguments(&container, serviceAttr.Command, serviceAttr.Args); err != nil {
		return nil, secretsArray, configMapsArray, err
	}

	if err := putLimitResource(&container, serviceAttr.LimitResources); err != nil {
		return nil, secretsArray, configMapsArray, err
	}
	if err := putRequestResource(&container, serviceAttr.RequestResources); err != nil {
		return nil, secretsArray, configMapsArray, err
	}

	//init container do not have readiness prob
	if serviceAttr.SecurityContext != nil {
		if securityContext, err := configureSecurityContext(*serviceAttr.SecurityContext); err != nil {
			return nil, secretsArray, configMapsArray, err
		} else {
			container.SecurityContext = securityContext
		}
	}
	container.Image = serviceAttr.ImagePrefix + serviceAttr.ImageName
	if serviceAttr.Tag != "" {
		container.Image += ":" + serviceAttr.Tag
	}
	var ports []v1.ContainerPort
	for _, port := range serviceAttr.Ports {
		temp := v1.ContainerPort{}
		if port.Container == "" && port.Host == "" {
			continue
		}
		if port.Container == "" && port.Host != "" {
			port.Container = port.Host
		}

		i, err := strconv.Atoi(port.Container)
		if err != nil {
			utils.Info.Println(err)
			continue
		}
		if i > 0 && i < 65536 {
			temp.ContainerPort = int32(i)
		} else {
			utils.Info.Println("invalid prot number")
			continue
		}
		if port.Host != "" {
			i, err = strconv.Atoi(port.Host)
			if err != nil {
				utils.Info.Println(err)
				continue
			}
			if i > 0 && i < 65536 {
				temp.HostPort = int32(i)
			} else {
				utils.Info.Println("invalid prot number")
				continue
			}
		}
		ports = append(ports, temp)
	}
	//configMapExist := make(map[string]bool)
	//secretExist := make(map[string]bool)

	var envVariables []v1.EnvVar
	for _, envVariable := range serviceAttr.EnvironmentVariables {
		tempEnvVariable := v1.EnvVar{}
		if envVariable.IsConfigMap {
			source, key := resolveValue(envVariable.Value)
			tempEnvVariable = v1.EnvVar{Name: envVariable.Key,
				ValueFrom: &v1.EnvVarSource{ConfigMapKeyRef: &v1.ConfigMapKeySelector{
					LocalObjectReference: v1.LocalObjectReference{Name: source},
					Key:                  key,
				}}}
		} else if envVariable.IsSecret {
			source, key := resolveValue(envVariable.Value)
			tempEnvVariable = v1.EnvVar{Name: envVariable.Key,
				ValueFrom: &v1.EnvVarSource{SecretKeyRef: &v1.SecretKeySelector{
					LocalObjectReference: v1.LocalObjectReference{Name: source},
					Key:                  key,
				}}}
		} else {
			tempEnvVariable = v1.EnvVar{Name: envVariable.Key, Value: envVariable.Value}
		}
		envVariables = append(envVariables, tempEnvVariable)
	}
	container.Ports = ports
	container.Env = envVariables
	var containers []v1.Container

	containers = append(containers, container)

	return containers, secretsArray, configMapsArray, nil
}
func getContainers(service types.Service) ([]v1.Container, []string, []string, error) {
	var configMapsArray, secretsArray []string
	var container v1.Container
	container.Name = service.Name
	byteData, _ := json.Marshal(service.ServiceAttributes)
	var serviceAttr types.DockerServiceAttributes
	if err := json.Unmarshal(byteData, &serviceAttr); err != nil {
		return nil, nil, nil, err
	}

	if err := putCommandAndArguments(&container, serviceAttr.Command, serviceAttr.Args); err != nil {
		return nil, secretsArray, configMapsArray, err
	}
	errr, isOk := checkRequestIsLessThanLimit(serviceAttr)
	if errr != nil {
		return nil, secretsArray, configMapsArray, errr
	} else if isOk == false {
		return nil, secretsArray, configMapsArray, errors.New("Request Resource is greater limit resource")

	}
	if err := putLimitResource(&container, serviceAttr.LimitResources); err != nil {
		return nil, secretsArray, configMapsArray, err
	}
	if err := putRequestResource(&container, serviceAttr.RequestResources); err != nil {
		return nil, secretsArray, configMapsArray, err
	}
	if err := putLivenessProbe(&container, serviceAttr.LivenessProb); err != nil {
		return nil, secretsArray, configMapsArray, err
	}
	if err := putReadinessProbe(&container, serviceAttr.RedinessProb); err != nil {
		return nil, secretsArray, configMapsArray, err
	}
	if serviceAttr.SecurityContext != nil {
		if securityContext, err := configureSecurityContext(*serviceAttr.SecurityContext); err != nil {
			return nil, secretsArray, configMapsArray, err
		} else {
			container.SecurityContext = securityContext
		}
	}
	container.Image = serviceAttr.ImagePrefix + serviceAttr.ImageName
	if serviceAttr.Tag != "" {
		container.Image += ":" + serviceAttr.Tag
	}
	var ports []v1.ContainerPort
	for _, port := range serviceAttr.Ports {
		temp := v1.ContainerPort{}
		if port.Container == "" && port.Host == "" {
			continue
		}
		if port.Container == "" && port.Host != "" {
			port.Container = port.Host
		}

		i, err := strconv.Atoi(port.Container)
		if err != nil {
			utils.Info.Println(err)
			continue
		}
		if i > 0 && i < 65536 {
			temp.ContainerPort = int32(i)
		} else {
			utils.Info.Println("invalid prot number")
			continue
		}
		if port.Host != "" {
			i, err = strconv.Atoi(port.Host)
			if err != nil {
				utils.Info.Println(err)
				continue
			}
			if i > 0 && i < 65536 {
				temp.HostPort = int32(i)
			} else {
				utils.Info.Println("invalid prot number")
				continue
			}

		}
		ports = append(ports, temp)
	}
	var envVariables []v1.EnvVar
	for _, envVariable := range serviceAttr.EnvironmentVariables {
		tempEnvVariable := v1.EnvVar{}
		if envVariable.IsConfigMap {
			source, key := resolveValue(envVariable.Value)
			tempEnvVariable = v1.EnvVar{Name: envVariable.Key,
				ValueFrom: &v1.EnvVarSource{ConfigMapKeyRef: &v1.ConfigMapKeySelector{
					LocalObjectReference: v1.LocalObjectReference{Name: source},
					Key:                  key,
				}}}
		} else if envVariable.IsSecret {
			source, key := resolveValue(envVariable.Value)
			tempEnvVariable = v1.EnvVar{Name: envVariable.Key,
				ValueFrom: &v1.EnvVarSource{SecretKeyRef: &v1.SecretKeySelector{
					LocalObjectReference: v1.LocalObjectReference{Name: source},
					Key:                  key,
				}}}
		} else {
			tempEnvVariable = v1.EnvVar{Name: envVariable.Key, Value: envVariable.Value}
		}
		envVariables = append(envVariables, tempEnvVariable)
	}
	container.Ports = ports
	container.Env = envVariables
	var containers []v1.Container
	containers = append(containers, container)
	return containers, secretsArray, configMapsArray, nil
}

func resolveValue(value string) (string, string) {
	result := returnValuesArray(value)
	var splittedResult []string
	for _, every := range result {
		if len(every) > 0 {
			splittedResult = strings.Split(every, ":")
			break
		}
	}
	keyArray := strings.Split(splittedResult[1], ".")
	return splittedResult[0], keyArray[len(keyArray)-1]
}

func returnValuesArray(value string) []string {
	replacer := strings.NewReplacer(
		"{{", "|-|",
		"}}", "|-|")
	result := replacer.Replace(value)
	return strings.Split(result, "|-|")
}

func marshalUnMarshalOfIstioComponents(s string) (map[string]interface{}, error) {
	jsonRaw, err := yaml2.YAMLToJSON([]byte(s))
	if err != nil {
		utils.Error.Println(err)
		return nil, err
	}
	var dd map[string]interface{}
	err = json.Unmarshal(jsonRaw, &dd)
	if err != nil {
		utils.Error.Println(err)
		return nil, err
	}
	return dd, nil
}

/*

@Deprecated Methods

*/
func jsonParser(str string, str2 string) string {

	for strings.Index(str, str2) != -1 {
		ind := strings.Index(str, str2)
		length := len(str)
		replaced := false
		for ind < length && !replaced {
			if str[ind] == '}' {
				str = str[:ind] + str[ind+1:]
				replaced = true
			}
			ind = ind + 1
		}
		str = strings.Replace(str, str2, "", 1)
	}
	return str

}
func timeParser(str string, str2 string) string {

	for strings.Index(str, str2) != -1 {
		ind := strings.Index(str, str2)
		length := len(str)
		replaced := false
		for ind < length && !replaced {
			if str[ind] == '}' {
				str = str[:ind] + "s\"" + str[ind+1:]
				replaced = true
			}
			ind = ind + 1
		}
		str = strings.Replace(str, str2, "\"", 1)
	}
	return str

}

func configureSecurityContext(securityContext types.SecurityContextStruct) (*v1.SecurityContext, error) {

	var context v1.SecurityContext
	context.Capabilities = &v1.Capabilities{}
	for _, addCapability := range securityContext.CapabilitiesAdd {
		context.Capabilities.Add = append(context.Capabilities.Add, v1.Capability(addCapability.(string)))
	}
	for _, dropCapability := range securityContext.CapabilitiesDrop {
		context.Capabilities.Drop = append(context.Capabilities.Drop, v1.Capability(dropCapability.(string)))
	}
	context.ReadOnlyRootFilesystem = &securityContext.ReadOnlyRootFileSystem
	context.Privileged = &securityContext.Privileged
	if securityContext.RunAsNonRoot && securityContext.RunAsUser == nil {
		return nil, errors.New("RunAsNonRoot is Set, but RunAsUser value not given!")
	} else {
		context.RunAsNonRoot = &securityContext.RunAsNonRoot
		context.RunAsUser = securityContext.RunAsUser
	}
	context.RunAsGroup = securityContext.RunAsGroup
	context.AllowPrivilegeEscalation = &securityContext.AllowPrivilegeEscalation
	if proMount, ok := securityContext.ProcMount.(string); ok {
		tempProcMount := v1.ProcMountType(proMount)
		context.ProcMount = &tempProcMount
	}
	context.SELinuxOptions = &v1.SELinuxOptions{
		User:  securityContext.SELinuxOptions.User,
		Role:  securityContext.SELinuxOptions.Role,
		Type:  securityContext.SELinuxOptions.Type,
		Level: securityContext.SELinuxOptions.Level,
	}
	return &context, nil
}

func ConfigMapObject(data types.ConfigMap) (*v1.ConfigMap, error) {
	var configMap v1.ConfigMap
	if data.Name != nil {
		configMap.Name = *data.Name
	}
	if data.Namespace != nil {
		configMap.Namespace = *data.Namespace
	}

	configMap.Labels["app"] = *data.Name

	if data.Data != nil {
		for key, value := range data.Data {
			configMap.Data[key] = value
		}
	}
	return &configMap, nil
}

func getVolumeMount(data types.ContainerVolumeMount) *v1.VolumeMount {
	var volumeMount v1.VolumeMount
	volumeMount.Name = data.Name
	if data.ReadOnly {
		volumeMount.ReadOnly = data.ReadOnly
	}
	volumeMount.MountPath = data.MountPath
	return &volumeMount
}

func getVolume(data types.ContainerVolume) *v1.Volume {
	var volume v1.Volume
	volume.Name = data.Name
	if data.ContainerConfigMap != nil {
		volume.ConfigMap.Name = data.ContainerConfigMap.Name
		volume.ConfigMap.LocalObjectReference.Name = data.ContainerConfigMap.Name
		if data.ContainerConfigMap.DefaultMode != 0 {
			volume.ConfigMap.DefaultMode = &data.ContainerConfigMap.DefaultMode
		}
		for _, item := range data.ContainerConfigMap.Items {
			var tempItem v1.KeyToPath
			tempItem.Key = item.Key
			if item.Mode != 0 {
				tempItem.Mode = &item.Mode
			}
			if item.Path != "" {
				tempItem.Path = item.Path
			}
			volume.ConfigMap.Items = append(volume.ConfigMap.Items, tempItem)
		}
	} else if data.ContainerSecret != nil {
		volume.Secret.SecretName = data.ContainerSecret.Name
		if data.ContainerSecret.DefaultMode != 0 {
			volume.Secret.DefaultMode = &data.ContainerSecret.DefaultMode
		}
		for _, item := range data.ContainerSecret.Items {
			var tempItem v1.KeyToPath
			tempItem.Key = item.Key
			if item.Mode != 0 {
				tempItem.Mode = &item.Mode
			}
			if item.Path != "" {
				tempItem.Path = item.Path
			}
			volume.Secret.Items = append(volume.Secret.Items, tempItem)
		}
	}
	return &volume
}

func getLabels(data types.Service) (map[string]string, error) {

	var serviceAttributes types.DockerServiceAttributes
	if data, err := json.Marshal(data.ServiceAttributes); err == nil {

		if err = json.Unmarshal(data, &serviceAttributes); err != nil {
			utils.Error.Println(err)
			return nil, err
		}
	} else {
		return nil, err
	}
	labels := make(map[string]string)
	for key, value := range serviceAttributes.Labels {
		labels[key] = value
	}
	return labels, nil
}

func getAnnotations(data types.Service) (map[string]string, error) {
	var serviceAttributes types.DockerServiceAttributes
	if data, err := json.Marshal(data.ServiceAttributes); err == nil {
		if err = json.Unmarshal(data, &serviceAttributes); err != nil {
			utils.Error.Println(err)
			return nil, err
		}
	} else {
		return nil, err
	}
	annotations := make(map[string]string)
	for key, value := range serviceAttributes.Annotations {
		annotations[key] = value
	}
	return annotations, nil
}
func (p Parser) Parse(spec string) error {
	if len(spec) == 0 {
		return fmt.Errorf("empty spec string")
	}

	// Split on whitespace.
	fields := strings.Fields(spec)

	// Validate & fill in any omitted or optional fields
	var err error
	fields, err = normalizeFields(fields, p.options)

	return err
}

// A custom Parser that can be configured.
type Parser struct {
	options ParseOption
}
type ParseOption int

const (
	Second         ParseOption = 1 << iota // Seconds field, default 0
	SecondOptional                         // Optional seconds field, default 0
	Minute                                 // Minutes field, default 0
	Hour                                   // Hours field, default 0
	Dom                                    // Day of month field, default *
	Month                                  // Month field, default *
	Dow                                    // Day of week field, default *
	DowOptional                            // Optional day of week field, default *
	Descriptor                             // Allow descriptors such as @monthly, @weekly, etc.
)

var places = []ParseOption{
	Minute,
	Hour,
	Dom,
	Month,
	Dow,
}

var defaults = []string{
	"0",
	"0",
	"0",
	"*",
	"*",
	"*",
}

func normalizeFields(fields []string, options ParseOption) ([]string, error) {
	// Validate optionals & add their field to options
	optionals := 0
	if options&SecondOptional > 0 {
		options |= Second
		optionals++
	}
	if options&DowOptional > 0 {
		options |= Dow
		optionals++
	}
	if optionals > 1 {
		return nil, fmt.Errorf("multiple optionals may not be configured")
	}

	// Figure out how many fields we need
	max := 0
	for _, place := range places {
		if options&place > 0 {
			max++
		}
	}
	min := max - optionals

	// Validate number of fields
	if count := len(fields); count < min || count > max {
		if min == max {
			return nil, fmt.Errorf("expected exactly %d fields, found %d: %s", min, count, fields)
		}
		return nil, fmt.Errorf("expected %d to %d fields, found %d: %s", min, max, count, fields)
	}

	// Populate the optional field if not provided
	if min < max && len(fields) == min {
		switch {
		case options&DowOptional > 0:
			fields = append(fields, defaults[5]) // TODO: improve access to default
		case options&SecondOptional > 0:
			fields = append([]string{defaults[0]}, fields...)
		default:
			return nil, fmt.Errorf("unknown optional field")
		}
	}

	// Populate all fields not part of options with their defaults
	n := 0
	expandedFields := make([]string, len(places))
	copy(expandedFields, defaults)
	for i, place := range places {
		if options&place > 0 {
			expandedFields[i] = fields[n]
			n++
		}
	}
	return expandedFields, nil
}

var standardParser = NewParser(
	Minute | Hour | Dom | Month | Dow,
)

func NewParser(options ParseOption) Parser {
	optionals := 0
	if options&DowOptional > 0 {
		optionals++
	}
	if options&SecondOptional > 0 {
		optionals++
	}
	if optionals > 1 {
		panic("multiple optionals may not be configured")
	}
	return Parser{options}
}
func notAlreadyExistIstioObject(respons types.IstioObject, cpContext *core.Context, ret types.StatusRequest) (bool, error) {
	var serviceOutput types.ServiceOutput
	serviceOutput.Services.Istio = append(serviceOutput.Services.Istio, respons)
	if pId, ok := cpContext.Keys["project_id"]; ok {
		serviceOutput.ProjectId = fmt.Sprintf("%v", pId)
	}
	x, err := json.Marshal(serviceOutput)
	if err != nil {
		return false, err
	}

	utils.Info.Println("kubernetes request payload", string(x))
	resp, res := GetFromKube(x, serviceOutput.ProjectId, ret, "GET", cpContext)
	if len(res.Service.Istio) > 0 && strings.Contains(res.Service.Istio[0].Error, "not found") {
		resp = ForwardToKube(x, serviceOutput.ProjectId, "POST", ret, cpContext)
		if resp.Reason != "" {
			ret.Status = append(ret.Status, "failed")
			ret.Reason = "Not a valid Istio Object. Error : " + resp.Reason
			return false, errors.New(resp.Reason)
		}
		return false, nil
	}
	return true, nil
}
