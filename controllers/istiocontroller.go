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
	//"github.com/istio/api/networking/v1alpha3"
	"io/ioutil"
	"istio-service-mesh/constants"
	"istio-service-mesh/controllers/volumes"
	"istio-service-mesh/types"
	"istio-service-mesh/utils"
	policy "istio.io/api/authentication/v1alpha1"
	"istio.io/api/networking/v1alpha3"
	"istio.io/istio/pilot/pkg/model"
	v12 "k8s.io/api/apps/v1"
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
				Percentage:    &v1alpha3.Percent{Value: http.FaultInjection.FaultInjectionAbort.Percentage},
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
		p.Name = port.Protocol
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
	//SE.Location = v1alpha3.ServiceEntry_Location()

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
		labels["namespace"] = strings.ToLower(input.Namespace)
		istioServ.Metadata = labels
		istioServ.Kind = "ServiceEntry"
		istioServ.ApiVersion = "networking.istio.io/v1alpha3"
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
		labels["namespace"] = strings.ToLower(input.Namespace)
		istioServ.Metadata = labels
		istioServ.Kind = "VirtualService"
		istioServ.ApiVersion = "networking.istio.io/v1alpha3"
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
		labels["namespace"] = strings.ToLower(input.Namespace)
		istioServ.Metadata = labels
		istioServ.Kind = "DestinationRule"
		istioServ.ApiVersion = "networking.istio.io/v1alpha3"
		components = append(components, istioServ)
		return components, nil
	}
	return components, nil

}

func getDeploymentObject(service types.Service) (v12.Deployment, error) {
	var secrets, configMaps []string
	var deployment = v12.Deployment{}
	// Label Selector
	//keel labels
	deploymentLabels := make(map[string]string)
	//deploymentLabels["keel.sh/match-tag"] = "true"
	deploymentLabels["keel.sh/policy"] = "force"
	//deploymentLabels["keel.sh/trigger"] = "poll"
	var selector metav1.LabelSelector
	labels, _ := getLabels(service)

	labels["app"] = service.Name
	labels["version"] = strings.ToLower(service.Version)

	if service.Name == "" {
		//Failed
		return v12.Deployment{}, errors.New("Service name not found")
	}
	deployment.ObjectMeta.Name = service.Name + "-" + service.Version
	deployment.ObjectMeta.Labels = deploymentLabels
	selector.MatchLabels = labels

	if service.Namespace == "" {
		deployment.ObjectMeta.Namespace = "default"
	} else {
		deployment.ObjectMeta.Namespace = service.Namespace
	}
	deployment.Spec.Selector = &selector
	deployment.Spec.Template.ObjectMeta.Labels = labels
	Annotations, _ := getAnnotations(service)
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
	daemonset.Kind = "DaemonSet"
	daemonset.APIVersion = "apps/v1"
	// Label Selector
	//keel labels
	deploymentLabels := make(map[string]string)
	//deploymentLabels["keel.sh/match-tag"] = "true"
	deploymentLabels["keel.sh/policy"] = "force"
	//deploymentLabels["keel.sh/trigger"] = "poll"
	var selector metav1.LabelSelector
	labels, _ := getLabels(service)
	labels["app"] = service.Name
	labels["version"] = strings.ToLower(service.Version)

	if service.Name == "" {
		//Failed
		return v12.DaemonSet{}, errors.New("Service name not found")
	}
	daemonset.ObjectMeta.Name = service.Name + "-" + service.Version
	daemonset.ObjectMeta.Labels = deploymentLabels
	selector.MatchLabels = labels

	if service.Namespace == "" {
		daemonset.ObjectMeta.Namespace = "default"
	} else {
		daemonset.ObjectMeta.Namespace = service.Namespace
	}
	daemonset.Spec.Selector = &selector
	daemonset.Spec.Template.ObjectMeta.Labels = labels
	Annotations, _ := getAnnotations(service)
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
	deploymentLabels := make(map[string]string)
	//deploymentLabels["keel.sh/match-tag"] = "true"
	deploymentLabels["keel.sh/policy"] = "force"
	//deploymentLabels["keel.sh/trigger"] = "poll"

	var selector metav1.LabelSelector
	labels, _ := getLabels(service)
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
	//cronjob.Spec.JobTemplate.Spec.Selector = &selector
	cronjob.Spec.JobTemplate.Spec.Template.ObjectMeta.Labels = labels
	Annotations, _ := getAnnotations(service)
	Annotations["sidecar.istio.io/inject"] = "false"
	cronjob.Spec.JobTemplate.Spec.Template.ObjectMeta.Annotations = Annotations
	//

	byteData, _ := json.Marshal(service.ServiceAttributes)
	var serviceAttr types.DockerServiceAttributes
	json.Unmarshal(byteData, &serviceAttr)

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
	labels, _ := getLabels(service)
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
	//job.Spec.Selector = &selector
	job.Spec.Template.ObjectMeta.Labels = labels
	Annotations, _ := getAnnotations(service)
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
	statefulset.Kind = "StatefulSet"
	statefulset.APIVersion = "v1"
	// Label Selector
	//keel labels
	deploymentLabels := make(map[string]string)
	//deploymentLabels["keel.sh/match-tag"] = "true"
	deploymentLabels["keel.sh/policy"] = "force"
	//deploymentLabels["keel.sh/trigger"] = "poll"

	var selector metav1.LabelSelector
	labels, _ := getLabels(service)
	labels["app"] = service.Name
	labels["version"] = strings.ToLower(service.Version)

	if service.Name == "" {
		//Failed
		return v12.StatefulSet{}, errors.New("Service name not found")
	}
	statefulset.ObjectMeta.Name = service.Name + "-" + service.Version
	statefulset.ObjectMeta.Labels = deploymentLabels
	selector.MatchLabels = labels

	if service.Namespace == "" {
		statefulset.ObjectMeta.Namespace = "default"
	} else {
		statefulset.ObjectMeta.Namespace = service.Namespace
	}
	statefulset.Spec.Selector = &selector
	statefulset.Spec.Template.ObjectMeta.Labels = labels
	Annotations, _ := getAnnotations(service)
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
	// Label Selector
	//keel labels
	deploymentLabels := make(map[string]string)
	//deploymentLabels["keel.sh/match-tag"] = "true"
	deploymentLabels["keel.sh/policy"] = "force"
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
	configmap.ObjectMeta.Name = service.Name
	configmap.ObjectMeta.Labels = labels
	//selector.MatchLabels = labels

	if serviceAttr.Namespace != nil && *serviceAttr.Namespace == "" {
		configmap.ObjectMeta.Namespace = "default"
	} else if serviceAttr.Namespace != nil {
		configmap.ObjectMeta.Namespace = service.Namespace
	}

	if serviceAttr.Data != nil {
		for key, value := range serviceAttr.Data {
			configmap.Data[key] = value
		}
	}

	return &configmap, nil
}
func getServiceObject(input types.Service) (*v1.Service, error) {
	service := v1.Service{}
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
	json.Unmarshal(byteData, &serviceAttr)
	var servicePorts []v1.ServicePort

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
		temp.Name = "p" + strconv.Itoa(i)
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

func DeployIstio(input types.ServiceInput, requestType string) types.StatusRequest {

	var ret types.StatusRequest
	ret.ID = input.SolutionInfo.Service.ID
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

	res, err := CheckGateway(service)
	if err != nil {
		utils.Info.Println("There is error in deployment")
		ret.Status = append(ret.Status, "failed")
		ret.Reason = "Not a valid Istio Object. Error : " + err.Error()
		if requestType != "GET" {
			utils.SendLog(ret.Reason, "error", input.ProjectId)
		}
		return ret
	}
	if res.Spec != nil && res.Metadata != nil {
		finalObj.Services.Istio = append(finalObj.Services.Istio, res)
	}

	if service.ServiceType == "mesh" || service.ServiceType == "other" {

		res, err := getIstioObject(service)
		if err != nil {
			utils.Info.Println("There is error in deployment")
			ret.Status = append(ret.Status, "failed")
			ret.Reason = "Not a valid Istio Object. Error : " + err.Error()
			if requestType != "GET" {
				utils.SendLog(ret.Reason, "error", input.ProjectId)
			}
			return ret
		}
		finalObj.Services.Istio = append(finalObj.Services.Istio, res...)

	}

	if service.ServiceType == "secret" {
		if secret, err := getSecretObject(service); err != nil {
			utils.Info.Println("There is error in deployment")
			ret.Status = append(ret.Status, "failed")
			ret.Reason = "Not a valid Secret Object. Error : " + err.Error()
			if requestType != "GET" {
				utils.SendLog(ret.Reason, "error", input.ProjectId)
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
				utils.SendLog(ret.Reason, "error", input.ProjectId)
			}
			return ret
		} else {
			finalObj.Services.ConfigMap = append(finalObj.Services.ConfigMap, *configmap)
		}
	}

	secret, exists := CreateDockerCfgSecret(service)
	if exists {
		finalObj.Services.Secrets = append(finalObj.Services.Secrets, secret)
	}

	if service.ServiceType == "volume" {
		byteData, _ := json.Marshal(service.ServiceAttributes)
		var attributes types.VolumeAttributes
		err := json.Unmarshal(byteData, &attributes)

		if err == nil && attributes.Volume.Name != "" {
			//Creating a new storage-class and persistent-volume-claim for each volume
			attributes.Volume.Namespace = "default"
			if service.Namespace != "" {
				attributes.Volume.Namespace = service.Namespace
			}
			finalObj.Services.StorageClasses = append(finalObj.Services.StorageClasses, volumes.ProvisionStorageClass(attributes.Volume))
			finalObj.Services.PersistentVolumeClaims = append(finalObj.Services.PersistentVolumeClaims, volumes.ProvisionVolumeClaim(attributes.Volume))
		}
	} else if service.ServiceType == "container" {
		switch service.SubType {

		case "deployment":
			deployment, err := getDeploymentObject(service)
			if err != nil {
				ret.Status = append(ret.Status, "failed")
				ret.Reason = "Not a valid Deployment Object. Error : " + err.Error()
				if requestType != "GET" {
					utils.SendLog(ret.Reason, "error", input.ProjectId)
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
				var attributes types.VolumeAttributes
				err = json.Unmarshal(byteData, &attributes)

				if err == nil && attributes.Volume.Name != "" {
					volumesData := []types.Volume{attributes.Volume}
					deployment.Spec.Template.Spec.Containers[0].VolumeMounts = volumes.GenerateVolumeMounts(volumesData)
					deployment.Spec.Template.Spec.Volumes = volumes.GeneratePodVolumes(volumesData)
				}
			}

			finalObj.Services.Deployments = append(finalObj.Services.Deployments, deployment)

			//add rbac classes

			byteData, _ := json.Marshal(service.ServiceAttributes)
			var serviceAttr types.DockerServiceAttributes
			json.Unmarshal(byteData, &serviceAttr)
			utils.Info.Println("** rbac params **")
			utils.Info.Println(len(serviceAttr.RbacRoles))
			if serviceAttr.IsRbac {
				utils.Info.Println("** rbac is enabled **")
				serviceAccount, roles, roleBindings, err := getRbacObjects(serviceAttr, service.Name, service.Namespace)
				if err != nil {
					ret.Status = append(ret.Status, "failed")
					ret.Reason = "Not a valid rbac Object. Error : " + err.Error()
					if requestType != "GET" {
						utils.SendLog(ret.Reason, "error", input.ProjectId)
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

		case "daemonset":
			daemonset, err := getDaemonSetObject(service)
			if err != nil {
				ret.Status = append(ret.Status, "failed")
				ret.Reason = "Not a valid DaemonSet Object. Error : " + err.Error()
				if requestType != "GET" {
					utils.SendLog(ret.Reason, "error", input.ProjectId)
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

			if serviceAttr.IsRbac {
				serviceAccount, roles, roleBindings, err := getRbacObjects(serviceAttr, service.Name, service.Namespace)
				if err != nil {
					ret.Status = append(ret.Status, "failed")
					ret.Reason = "Not a valid rbac Object. Error : " + err.Error()
					if requestType != "GET" {
						utils.SendLog(ret.Reason, "error", input.ProjectId)
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

		case "cronjob":
			cronjob, err := getCronJobObject(service)
			if err != nil {
				ret.Status = append(ret.Status, "failed")
				ret.Reason = "Not a valid CronJob Object. Error : " + err.Error()
				if requestType != "GET" {
					utils.SendLog(ret.Reason, "error", input.ProjectId)
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
				serviceAccount, roles, roleBindings, err := getRbacObjects(serviceAttr, service.Name, service.Namespace)
				if err != nil {
					ret.Status = append(ret.Status, "failed")
					ret.Reason = "Not a valid rbac Object. Error : " + err.Error()
					if requestType != "GET" {
						utils.SendLog(ret.Reason, "error", input.ProjectId)
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

		case "job":
			job, err := getJobObject(service)
			if err != nil {
				ret.Status = append(ret.Status, "failed")
				ret.Reason = "Not a valid Job Object. Error : " + err.Error()
				if requestType != "GET" {
					utils.SendLog(ret.Reason, "error", input.ProjectId)
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
				serviceAccount, roles, roleBindings, err := getRbacObjects(serviceAttr, service.Name, service.Namespace)
				if err != nil {
					ret.Status = append(ret.Status, "failed")
					ret.Reason = "Not a valid rbac Object. Error : " + err.Error()
					if requestType != "GET" {
						utils.SendLog(ret.Reason, "error", input.ProjectId)
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

		case "statefulset":
			statefulset, err := getStatefulSetObject(service)
			if err != nil {
				ret.Status = append(ret.Status, "failed")
				ret.Reason = "Not a valid StatefulSet Object. Error : " + err.Error()
				if requestType != "GET" {
					utils.SendLog(ret.Reason, "error", input.ProjectId)
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

			if serviceAttr.IsRbac {
				serviceAccount, roles, roleBindings, err := getRbacObjects(serviceAttr, service.Name, service.Namespace)
				if err != nil {
					ret.Status = append(ret.Status, "failed")
					ret.Reason = "Not a valid rbac Object. Error : " + err.Error()
					if requestType != "GET" {
						utils.SendLog(ret.Reason, "error", input.ProjectId)
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

		//Getting Kubernetes Service Object
		serv, err := getServiceObject(service)
		if err != nil {
			ret.Status = append(ret.Status, "failed")
			ret.Reason = "Not a valid Service Object. Error : " + err.Error()
			if requestType != "GET" {
				utils.SendLog(ret.Reason, "error", input.ProjectId)
			}
			return ret
		}
		if serv != nil {
			finalObj.Services.Kubernetes = append(finalObj.Services.Kubernetes, *serv)
		}
	}

	//Send request to Kubernetes
	x, err := json.Marshal(finalObj)
	if err != nil {
		utils.Info.Println(err)
		ret.Status = append(ret.Status, "failed")
		ret.Reason = "Service Object parsing failed : " + err.Error()
		if requestType != "GET" {
			utils.SendLog(ret.Reason, "error", input.ProjectId)
		}
		return ret
	}
	utils.Info.Println("kubernetes request payload", string(x))

	if requestType != "POST" {
		ret, resp := GetFromKube(x, input.ProjectId, ret, requestType)
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
								ret.Status = append(ret.Status, "in progress")

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
			utils.SendLog(ret.Reason, "error", input.ProjectId)
		}
		return ret
	}
	if requestType != "GET" {
		//Send failure request
		return ForwardToKube(x, input.ProjectId, requestType, ret)
	}
	return ret

}

func GetFromKube(requestBody []byte, env_id string, ret types.StatusRequest, requestType string) (types.StatusRequest, types.ResponseRequest) {
	url := constants.KubernetesEngineURL
	var res types.ResponseRequest
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		utils.Info.Println(err)
		if requestType != "GET" {
			utils.SendLog("Connection to kubernetes microservice failed "+err.Error(), "info", env_id)
		}
		ret.Status = append(ret.Status, "failed")
		ret.Reason = "Connection to kubernetes deployment microservice failed Error : " + err.Error()
		if requestType != "GET" {
			utils.SendLog(ret.Reason, "error", env_id)
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
				utils.SendLog("Connection to kubernetes microservice failed "+err.Error(), "info", env_id)
				utils.SendLog(ret.Reason, "error", env_id)
			}
			return ret, res
		} else {

			utils.Info.Println(string(result))
			if requestType != "GET" {
				utils.SendLog(string(result), "info", env_id)
			}
			if statusCode != 200 {
				var resrf types.ResponseServiceRequestFailure
				err = json.Unmarshal(result, &resrf)
				if err != nil {
					ret.Status = append(ret.Status, "failed")
					ret.Reason = "kubernetes deployment microservice Response Parsing failed.Error : " + err.Error()
					if requestType != "GET" {

						utils.SendLog(ret.Reason, "error", env_id)
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

						utils.SendLog(ret.Reason, "error", env_id)
					}
					return ret, res
				}
				return ret, res
			}
		}
		return ret, res
	}
}
func ForwardToKube(requestBody []byte, env_id string, requestType string, ret types.StatusRequest) types.StatusRequest {

	url := constants.KubernetesEngineURL

	utils.Info.Println("forward to kube: " + url)
	utils.Info.Println("request type: " + requestType)

	req, err := http.NewRequest(requestType, url, bytes.NewBuffer(requestBody))

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		utils.Info.Println(err)
		ret.Status = append(ret.Status, "failed")
		ret.Reason = "Connection to kubernetes deployment microservice failed Error : " + err.Error()

		if requestType != "GET" {
			utils.SendLog("Connection to kubernetes microservice failed "+err.Error(), "info", env_id)

			utils.SendLog(ret.Reason, "error", env_id)
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
				utils.SendLog("Response Parsing failed "+err.Error(), "error", env_id)
				utils.SendLog(ret.Reason, "error", env_id)
			}
			return ret
		} else {
			utils.Info.Println(string(result))
			if requestType != "GET" {
				utils.SendLog(string(result), "info", env_id)
			}
			if statusCode != 200 {
				var resrf types.ResponseServiceRequestFailure
				err = json.Unmarshal(result, &resrf)
				if err != nil {
					ret.Status = append(ret.Status, "failed")
					ret.Reason = "kubernetes deployment microservice Response Parsing failed.Error : " + err.Error()
					if requestType != "GET" {

						utils.SendLog(ret.Reason, "error", env_id)
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
		return
	}

	var notification types.Notifier
	notification.Component = "Solution"
	notification.Id = input.SolutionInfo.Service.ID

	var status types.StatusRequest
	status.ID = input.SolutionInfo.Service.ID
	status.Name = input.SolutionInfo.Service.Name

	result := DeployIstio(input, r.Method)

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

		utils.Info.Println("Deployment Successful\n")
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
func CreateDockerCfgSecret(service types.Service) (v1.Secret, bool) {

	byteData, _ := json.Marshal(service.ServiceAttributes)
	var serviceAttr types.DockerServiceAttributes
	json.Unmarshal(byteData, &serviceAttr)

	if serviceAttr.ImageRepositoryConfigurations.Url == "" {
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
	email := "email@email.com"
	server := serviceAttr.ImageRepositoryConfigurations.Url

	tokens := strings.Split(server, "/")
	_ = tokens
	registry := server

	dockerConf := map[string]map[string]string{
		registry: {
			"username": username,
			"password": password,
			"email":    email,
			"auth":     base64.StdEncoding.EncodeToString([]byte(username + ":" + password)),
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
		Name:      service.Name,
		Namespace: service.Namespace,
	}

	secret.TypeMeta = typeMeta
	secret.ObjectMeta = objectMeta
	if serviceAttr.Data != nil {
		secret.Data = serviceAttr.Data
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
		Name:      service.Name,
		Namespace: service.Namespace,
	}

	secret.TypeMeta = typeMeta
	secret.ObjectMeta = objectMeta
	if serviceAttr.Data != nil {
		secret.Data = serviceAttr.Data
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
		container.Args = args
	} else if len(args) > 0 {
		return errors.New("Error Found: Arguments provided without a command.")
	}
	return nil
}
func putLimitResource(container *v1.Container, limitResourceTypes, limitResourceQuantities []string) error {
	temp := make(map[v1.ResourceName]resource.Quantity)
	for i := 0; i < len(limitResourceTypes) && i < len(limitResourceQuantities); i++ {
		if limitResourceTypes[i] == "memory" || limitResourceTypes[i] == "cpu" {
			quantity, err := resource.ParseQuantity(limitResourceQuantities[i])
			if err != nil {
				return err
			}
			temp[v1.ResourceName(limitResourceTypes[i])] = quantity
		} else {
			return errors.New("Error Found: Invalid Limit Resource Provided. Valid: 'cpu','memory'")
		}
	}
	container.Resources.Limits = temp
	return nil
}
func putRequestResource(container *v1.Container, requestResourceTypes, requestResourceQuantities []string) error {
	temp := make(map[v1.ResourceName]resource.Quantity)
	for i := 0; i < len(requestResourceTypes) && i < len(requestResourceQuantities); i++ {
		if requestResourceTypes[i] == "memory" || requestResourceTypes[i] == "cpu" {
			quantity, err := resource.ParseQuantity(requestResourceQuantities[i])
			if err != nil {
				return err
			}
			temp[v1.ResourceName(requestResourceTypes[i])] = quantity
		} else {
			return errors.New("Error Found: Invalid Request Resource Provided. Valid: 'cpu','memory'")
		}
	}
	container.Resources.Requests = temp
	return nil
}
func putLivenessProbe(container *v1.Container, byteData []byte) error {

	strLowerCamel := convertKeys(byteData)
	var tempLivenessProbe tempProbing
	json.Unmarshal(strLowerCamel, &tempLivenessProbe)
	utils.Info.Println(tempLivenessProbe)
	container.LivenessProbe = tempLivenessProbe.LivenessProbe

	return nil
}
func putReadinessProbe(container *v1.Container, byteData []byte) error {

	strLowerCamel := convertKeys(byteData)
	var tempReadinessProbe tempProbing
	json.Unmarshal(strLowerCamel, &tempReadinessProbe)
	utils.Info.Println(tempReadinessProbe)
	container.ReadinessProbe = tempReadinessProbe.ReadinessProbe

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

type tempProbing struct {
	LivenessProbe  *v1.Probe `json:"livenessProbe"`
	ReadinessProbe *v1.Probe `json:"readinessProbe"`
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
	json.Unmarshal(byteData, &serviceAttr)
	container.Name = serviceAttr.ImageName
	if err := putCommandAndArguments(&container, serviceAttr.Command, serviceAttr.Args); err != nil {
		return nil, secretsArray, configMapsArray, err
	}
	if err := putLimitResource(&container, serviceAttr.LimitResourceTypes, serviceAttr.LimitResourceQuantities); err != nil {
		return nil, secretsArray, configMapsArray, err
	}
	if err := putRequestResource(&container, serviceAttr.RequestResourceTypes, serviceAttr.RequestResourceQuantities); err != nil {
		return nil, secretsArray, configMapsArray, err
	}
	if err := putLivenessProbe(&container, byteData); err != nil {
		return nil, secretsArray, configMapsArray, err
	}

	if securityContext, err := configureSecurityContext(serviceAttr.SecurityContext); err != nil {
		return nil, secretsArray, configMapsArray, err
	} else {
		container.SecurityContext = securityContext
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

		temp.ContainerPort = int32(i)
		if port.Host != "" {
			i, err = strconv.Atoi(port.Host)
			if err != nil {
				utils.Info.Println(err)
				continue
			}
			temp.HostPort = int32(i)
		}
		ports = append(ports, temp)
	}
	//configMapExist := make(map[string]bool)
	//secretExist := make(map[string]bool)

	var envVariables []v1.EnvVar
	for _, envVariable := range serviceAttr.EnvironmentVariables {
		tempEnvVariable := v1.EnvVar{}
		if envVariable.IsConfigMap {
			tempEnvVariable = v1.EnvVar{Name: envVariable.Key,
				ValueFrom: &v1.EnvVarSource{ConfigMapKeyRef: &v1.ConfigMapKeySelector{
					LocalObjectReference: v1.LocalObjectReference{Name: ""},
					Key:                  "",
				}}}
		} else if envVariable.IsSecret {
			tempEnvVariable = v1.EnvVar{Name: envVariable.Key,
				ValueFrom: &v1.EnvVarSource{SecretKeyRef: &v1.SecretKeySelector{
					LocalObjectReference: v1.LocalObjectReference{Name: ""},
					Key:                  "",
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
	json.Unmarshal(byteData, &serviceAttr)
	if err := putCommandAndArguments(&container, serviceAttr.Command, serviceAttr.Args); err != nil {
		return nil, secretsArray, configMapsArray, err
	}
	if err := putLimitResource(&container, serviceAttr.LimitResourceTypes, serviceAttr.LimitResourceQuantities); err != nil {
		return nil, secretsArray, configMapsArray, err
	}
	if err := putRequestResource(&container, serviceAttr.RequestResourceTypes, serviceAttr.RequestResourceQuantities); err != nil {
		return nil, secretsArray, configMapsArray, err
	}
	if err := putLivenessProbe(&container, byteData); err != nil {
		return nil, secretsArray, configMapsArray, err
	}
	if err := putReadinessProbe(&container, byteData); err != nil {
		return nil, secretsArray, configMapsArray, err
	}

	if securityContext, err := configureSecurityContext(serviceAttr.SecurityContext); err != nil {
		return nil, secretsArray, configMapsArray, err
	} else {
		container.SecurityContext = securityContext
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

		temp.ContainerPort = int32(i)
		if port.Host != "" {
			i, err = strconv.Atoi(port.Host)
			if err != nil {
				utils.Info.Println(err)
				continue
			}
			temp.HostPort = int32(i)
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
		splittedResult = strings.Split(every, ":")
		break
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

func makeSecrets(data types.KubernetesSecret) (*v1.Secret, error) {
	var secret v1.Secret
	if data.Name != nil {
		secret.Name = *data.Name
	}
	if data.Namespace != nil {
		secret.Namespace = *data.Namespace
	}
	if data.Type != nil {
		switch *data.Type {
		case "Opaque":
			secret.Type = v1.SecretTypeOpaque
		case "TLS":
			secret.Type = v1.SecretTypeTLS
		}
	}
	if data.Data != nil {
		for key, value := range data.Data {
			secret.Data[key] = value
		}
	}
	if data.StringData != nil {
		for key, value := range data.StringData {
			secret.StringData[key] = value
		}
	}
	return &secret, nil
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
