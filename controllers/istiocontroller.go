package controllers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	googl_types "github.com/gogo/protobuf/types"
	"github.com/iancoleman/strcase"
	//"github.com/istio/api/networking/v1alpha3"
	"io/ioutil"
	"istio-service-mesh/constants"
	"istio-service-mesh/controllers/volumes"
	"istio-service-mesh/types"
	"istio-service-mesh/utils"
	"istio.io/api/networking/v1alpha3"
	"istio.io/istio/pilot/pkg/model"
	v12 "k8s.io/api/apps/v1"
	v13 "k8s.io/api/batch/v1"
	"k8s.io/api/batch/v2alpha1"
	"k8s.io/api/core/v1"
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
	json.Unmarshal(byteData, &serviceAttr)

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

		if http.FaultInjection.FaultInjectionAbort.Percentage != 0 && http.FaultInjection.FaultInjectionAbort.HttpStatus != 0 {
			abort := &v1alpha3.HTTPFaultInjection_Abort{
				Percentage: &v1alpha3.Percent{Value: http.FaultInjection.FaultInjectionAbort.Percentage},
				ErrorType:  &v1alpha3.HTTPFaultInjection_Abort_HttpStatus{HttpStatus: http.FaultInjection.FaultInjectionAbort.HttpStatus},
			}
			httpRoute.Fault.Abort = abort
		}
		if http.FaultInjection.FaultInjectionDelay.Percentage != 0 && http.FaultInjection.FaultInjectionDelay.FixedDelay != 0 {
			delay := &v1alpha3.HTTPFaultInjection_Delay{
				Percentage:    &v1alpha3.Percent{Value: http.FaultInjection.FaultInjectionAbort.Percentage},
				HttpDelayType: &v1alpha3.HTTPFaultInjection_Delay_FixedDelay{FixedDelay: &googl_types.Duration{Seconds: http.FaultInjection.FaultInjectionDelay.FixedDelay}},
			}
			httpRoute.Fault.Delay = delay
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
		fmt.Println(err)
	}

	fmt.Println(gotJSON)
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
func getIstioDestinationRule(service interface{}) (v1alpha3.DestinationRule, error) {
	destRule := v1alpha3.DestinationRule{}

	byteData, _ := json.Marshal(service)
	var serviceAttr types.IstioDestinationRuleAttributes
	json.Unmarshal(byteData, &serviceAttr)

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
		tp.ConnectionPool = &v1alpha3.ConnectionPoolSettings{
			Http: &v1alpha3.ConnectionPoolSettings_HTTPSettings{
				Http1MaxPendingRequests:  subset.Http1MaxPendingRequests,
				Http2MaxRequests:         subset.Http2MaxRequests,
				MaxRequestsPerConnection: subset.MaxRequestsPerConnection,
				MaxRetries:               subset.MaxRetries,
			},
		}
		ss.TrafficPolicy = &tp
		subsets = append(subsets, &ss)
	}
	destRule.Subsets = subsets
	destRule.Host = serviceAttr.Host
	destRule.Marshal()
	utils.Info.Println(destRule.String())
	return destRule, nil
}
func getIstioServiceEntry(service interface{}) (v1alpha3.ServiceEntry, error) {
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

	return SE, nil
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
		var istioServ types.IstioObject

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
		istioServ.Metadata = labels
		istioServ.Kind = "Gateway"
		istioServ.ApiVersion = "networking.istio.io/v1alpha3"

		return istioServ, nil
	}
	return istioServ, nil
}
func getIstioObject(input types.Service) (types.IstioObject, error) {
	var istioServ types.IstioObject

	switch input.SubType {

	case "service_entry":

		serv_entry, err := getIstioServiceEntry(input.ServiceAttributes)
		if err != nil {
			utils.Error.Println("There is error in deployment")
			return istioServ, err
		}
		istioServ.Spec = serv_entry
		labels := make(map[string]interface{})
		labels["name"] = strings.ToLower(input.Name)
		labels["app"] = strings.ToLower(input.Name)
		istioServ.Metadata = labels
		istioServ.Kind = "ServiceEntry"
		istioServ.ApiVersion = "networking.istio.io/v1alpha3"

	case "virtual_service":
		vr, err := getIstioVirtualService(input.ServiceAttributes)
		if err != nil {
			utils.Error.Println("There is error in deployment")
			return istioServ, err
		}
		/*d := jsonParser(vr, "\"Port\":{")
		d = jsonParser(d, "\"MatchType\":{")
		d = strings.Replace(d, "\"port\":{\"Number\"", "\"port\":{\"number\"", -1)
		utils.Info.Println(d)
		d = timeParser(d, "{\"seconds\":")
		utils.Info.Println(d)
		d = strings.Replace(d, "\"uri\":{\"Prefix\"", "\"uri\":{\"prefix\"", -1)*/

		//d = strings.Replace(d, "\"per_try_timeout\"", "\"perTryTimeout\"", -1)
		d := vr
		m, err := marshalUnMarshalOfIstioComponents(d)
		utils.Info.Println(err)
		istioServ.Spec = m
		labels := make(map[string]interface{})
		labels["name"] = strings.ToLower(input.Name)
		labels["app"] = strings.ToLower(input.Name)
		labels["version"] = strings.ToLower(input.Version)
		istioServ.Metadata = labels
		istioServ.Kind = "VirtualService"
		istioServ.ApiVersion = "networking.istio.io/v1alpha3"
		return istioServ, nil

	case "destination_rule":

		des_rule, err := getIstioDestinationRule(input.ServiceAttributes)
		if err != nil {
			utils.Error.Println("There is error in deployment")
			return istioServ, err
		}
		istioServ.Spec = des_rule
		labels := make(map[string]interface{})
		labels["name"] = strings.ToLower(input.Name)
		labels["app"] = strings.ToLower(input.Name)
		labels["version"] = strings.ToLower(input.Version)
		istioServ.Metadata = labels
		istioServ.Kind = "DestinationRule"
		istioServ.ApiVersion = "networking.istio.io/v1alpha3"
		return istioServ, nil
	}

	return istioServ, nil

}

func getDeploymentObject(service types.Service) (v12.Deployment, error) {
	var deployment = v12.Deployment{}
	// Label Selector

	//keel labels
	deploymentLabels := make(map[string]string)
	//deploymentLabels["keel.sh/match-tag"] = "true"
	deploymentLabels["keel.sh/policy"] = "force"
	//deploymentLabels["keel.sh/trigger"] = "poll"

	var selector metav1.LabelSelector
	labels := make(map[string]string)
	labels["app"] = service.Name
	labels["version"] = strings.ToLower(service.Version)

	if service.Name == "" {
		//Failed
		return v12.Deployment{}, errors.New("Service name not found")
	}
	deployment.ObjectMeta.Name = service.Name
	deployment.ObjectMeta.Labels = deploymentLabels
	selector.MatchLabels = labels

	if service.Namespace == "" {
		deployment.ObjectMeta.Namespace = "default"
	} else {
		deployment.ObjectMeta.Namespace = service.Namespace
	}
	deployment.Spec.Selector = &selector
	deployment.Spec.Template.ObjectMeta.Labels = labels
	deployment.Spec.Template.ObjectMeta.Annotations = map[string]string{
		"sidecar.istio.io/inject": "true",
	}
	//

	var container v1.Container
	container.Name = service.Name
	byteData, _ := json.Marshal(service.ServiceAttributes)
	var serviceAttr types.DockerServiceAttributes
	json.Unmarshal(byteData, &serviceAttr)
	err := putCommandAndArguments(&container, serviceAttr.Command, serviceAttr.Args)
	err = putLimitResource(&container, serviceAttr.LimitResourceTypes, serviceAttr.LimitResourceQuantities)
	err = putRequestResource(&container, serviceAttr.RequestResourceTypes, serviceAttr.RequestResourceQuantities)
	err = putLivenessProbe(&container, byteData)
	err = putReadinessProbe(&container, byteData)
	if err != nil {
		return v12.Deployment{}, err
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
		tempEnvVariable := v1.EnvVar{Name: envVariable.Key, Value: envVariable.Value}
		envVariables = append(envVariables, tempEnvVariable)
	}
	container.Ports = ports
	container.Env = envVariables
	var containers []v1.Container

	containers = append(containers, container)
	deployment.Spec.Template.Spec.Containers = containers

	return deployment, nil
}
func getDaemonSetObject(service types.Service) (v12.DaemonSet, error) {
	var daemonset = v12.DaemonSet{}
	// Label Selector

	//keel labels
	deploymentLabels := make(map[string]string)
	//deploymentLabels["keel.sh/match-tag"] = "true"
	deploymentLabels["keel.sh/policy"] = "force"
	//deploymentLabels["keel.sh/trigger"] = "poll"

	var selector metav1.LabelSelector
	labels := make(map[string]string)
	labels["app"] = service.Name
	labels["version"] = strings.ToLower(service.Version)

	if service.Name == "" {
		//Failed
		return v12.DaemonSet{}, errors.New("Service name not found")
	}
	daemonset.ObjectMeta.Name = service.Name
	daemonset.ObjectMeta.Labels = deploymentLabels
	selector.MatchLabels = labels

	if service.Namespace == "" {
		daemonset.ObjectMeta.Namespace = "default"
	} else {
		daemonset.ObjectMeta.Namespace = service.Namespace
	}
	daemonset.Spec.Selector = &selector
	daemonset.Spec.Template.ObjectMeta.Labels = labels
	daemonset.Spec.Template.ObjectMeta.Annotations = map[string]string{
		"sidecar.istio.io/inject": "true",
	}
	//

	var container v1.Container
	container.Name = service.Name
	byteData, _ := json.Marshal(service.ServiceAttributes)
	var serviceAttr types.DockerServiceAttributes
	json.Unmarshal(byteData, &serviceAttr)

	err := putCommandAndArguments(&container, serviceAttr.Command, serviceAttr.Args)
	err = putLimitResource(&container, serviceAttr.LimitResourceTypes, serviceAttr.LimitResourceQuantities)
	err = putRequestResource(&container, serviceAttr.RequestResourceTypes, serviceAttr.RequestResourceQuantities)
	err = putLivenessProbe(&container, byteData)
	err = putReadinessProbe(&container, byteData)
	if err != nil {
		return v12.DaemonSet{}, err
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
			fmt.Println(err)
			continue
		}

		temp.ContainerPort = int32(i)
		if port.Host != "" {
			i, err = strconv.Atoi(port.Host)
			if err != nil {
				fmt.Println(err)
				continue
			}
			temp.HostPort = int32(i)
		}
		ports = append(ports, temp)
	}
	var envVariables []v1.EnvVar
	for _, envVariable := range serviceAttr.EnvironmentVariables {
		tempEnvVariable := v1.EnvVar{Name: envVariable.Key, Value: envVariable.Value}
		envVariables = append(envVariables, tempEnvVariable)
	}
	container.Ports = ports
	container.Env = envVariables
	var containers []v1.Container

	containers = append(containers, container)
	daemonset.Spec.Template.Spec.Containers = containers

	return daemonset, nil
}
func getCronJobObject(service types.Service) (v2alpha1.CronJob, error) {
	var cronjob = v2alpha1.CronJob{}
	// Label Selector

	//keel labels
	deploymentLabels := make(map[string]string)
	//deploymentLabels["keel.sh/match-tag"] = "true"
	deploymentLabels["keel.sh/policy"] = "force"
	//deploymentLabels["keel.sh/trigger"] = "poll"

	var selector metav1.LabelSelector
	labels := make(map[string]string)
	labels["app"] = service.Name
	labels["version"] = strings.ToLower(service.Version)

	if service.Name == "" {
		//Failed
		return v2alpha1.CronJob{}, errors.New("Service name not found")
	}
	cronjob.ObjectMeta.Name = service.Name
	cronjob.ObjectMeta.Labels = deploymentLabels
	selector.MatchLabels = labels

	if service.Namespace == "" {
		cronjob.ObjectMeta.Namespace = "default"
	} else {
		cronjob.ObjectMeta.Namespace = service.Namespace
	}
	cronjob.Spec.JobTemplate.Spec.Selector = &selector
	cronjob.Spec.JobTemplate.Spec.Template.ObjectMeta.Labels = labels
	cronjob.Spec.JobTemplate.Spec.Template.ObjectMeta.Annotations = map[string]string{
		"sidecar.istio.io/inject": "true",
	}
	//

	var container v1.Container
	container.Name = service.Name
	byteData, _ := json.Marshal(service.ServiceAttributes)
	var serviceAttr types.DockerServiceAttributes
	json.Unmarshal(byteData, &serviceAttr)

	cronjob.Spec.Schedule = serviceAttr.CronJobScheduleString

	err := putCommandAndArguments(&container, serviceAttr.Command, serviceAttr.Args)
	err = putLimitResource(&container, serviceAttr.LimitResourceTypes, serviceAttr.LimitResourceQuantities)
	err = putRequestResource(&container, serviceAttr.RequestResourceTypes, serviceAttr.RequestResourceQuantities)
	err = putLivenessProbe(&container, byteData)
	err = putReadinessProbe(&container, byteData)
	if err != nil {
		return v2alpha1.CronJob{}, err
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
			fmt.Println(err)
			continue
		}

		temp.ContainerPort = int32(i)
		if port.Host != "" {
			i, err = strconv.Atoi(port.Host)
			if err != nil {
				fmt.Println(err)
				continue
			}
			temp.HostPort = int32(i)
		}
		ports = append(ports, temp)
	}
	var envVariables []v1.EnvVar
	for _, envVariable := range serviceAttr.EnvironmentVariables {
		tempEnvVariable := v1.EnvVar{Name: envVariable.Key, Value: envVariable.Value}
		envVariables = append(envVariables, tempEnvVariable)
	}
	container.Ports = ports
	container.Env = envVariables
	var containers []v1.Container

	containers = append(containers, container)
	cronjob.Spec.JobTemplate.Spec.Template.Spec.Containers = containers

	return cronjob, nil
}
func getJobObject(service types.Service) (v13.Job, error) {
	var job = v13.Job{}
	// Label Selector

	//keel labels
	deploymentLabels := make(map[string]string)
	//deploymentLabels["keel.sh/match-tag"] = "true"
	deploymentLabels["keel.sh/policy"] = "force"
	//deploymentLabels["keel.sh/trigger"] = "poll"

	var selector metav1.LabelSelector
	labels := make(map[string]string)
	labels["app"] = service.Name
	labels["version"] = strings.ToLower(service.Version)

	if service.Name == "" {
		//Failed
		return v13.Job{}, errors.New("Service name not found")
	}
	job.ObjectMeta.Name = service.Name
	job.ObjectMeta.Labels = deploymentLabels
	selector.MatchLabels = labels

	if service.Namespace == "" {
		job.ObjectMeta.Namespace = "default"
	} else {
		job.ObjectMeta.Namespace = service.Namespace
	}
	job.Spec.Selector = &selector
	job.Spec.Template.ObjectMeta.Labels = labels
	job.Spec.Template.ObjectMeta.Annotations = map[string]string{
		"sidecar.istio.io/inject": "true",
	}
	//

	var container v1.Container
	container.Name = service.Name
	byteData, _ := json.Marshal(service.ServiceAttributes)
	var serviceAttr types.DockerServiceAttributes
	json.Unmarshal(byteData, &serviceAttr)

	err := putCommandAndArguments(&container, serviceAttr.Command, serviceAttr.Args)
	err = putLimitResource(&container, serviceAttr.LimitResourceTypes, serviceAttr.LimitResourceQuantities)
	err = putRequestResource(&container, serviceAttr.RequestResourceTypes, serviceAttr.RequestResourceQuantities)
	err = putLivenessProbe(&container, byteData)
	err = putReadinessProbe(&container, byteData)
	if err != nil {
		return v13.Job{}, err
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
			fmt.Println(err)
			continue
		}

		temp.ContainerPort = int32(i)
		if port.Host != "" {
			i, err = strconv.Atoi(port.Host)
			if err != nil {
				fmt.Println(err)
				continue
			}
			temp.HostPort = int32(i)
		}
		ports = append(ports, temp)
	}
	var envVariables []v1.EnvVar
	for _, envVariable := range serviceAttr.EnvironmentVariables {
		tempEnvVariable := v1.EnvVar{Name: envVariable.Key, Value: envVariable.Value}
		envVariables = append(envVariables, tempEnvVariable)
	}
	container.Ports = ports
	container.Env = envVariables
	var containers []v1.Container

	containers = append(containers, container)
	job.Spec.Template.Spec.Containers = containers

	return job, nil
}
func getStatefulSetObject(service types.Service) (v12.StatefulSet, error) {
	var statefulset = v12.StatefulSet{}
	// Label Selector

	//keel labels
	deploymentLabels := make(map[string]string)
	//deploymentLabels["keel.sh/match-tag"] = "true"
	deploymentLabels["keel.sh/policy"] = "force"
	//deploymentLabels["keel.sh/trigger"] = "poll"

	var selector metav1.LabelSelector
	labels := make(map[string]string)
	labels["app"] = service.Name
	labels["version"] = strings.ToLower(service.Version)

	if service.Name == "" {
		//Failed
		return v12.StatefulSet{}, errors.New("Service name not found")
	}
	statefulset.ObjectMeta.Name = service.Name
	statefulset.ObjectMeta.Labels = deploymentLabels
	selector.MatchLabels = labels

	if service.Namespace == "" {
		statefulset.ObjectMeta.Namespace = "default"
	} else {
		statefulset.ObjectMeta.Namespace = service.Namespace
	}
	statefulset.Spec.Selector = &selector
	statefulset.Spec.Template.ObjectMeta.Labels = labels
	statefulset.Spec.Template.ObjectMeta.Annotations = map[string]string{
		"sidecar.istio.io/inject": "true",
	}
	//

	var container v1.Container
	container.Name = service.Name
	byteData, _ := json.Marshal(service.ServiceAttributes)
	var serviceAttr types.DockerServiceAttributes
	json.Unmarshal(byteData, &serviceAttr)

	err := putCommandAndArguments(&container, serviceAttr.Command, serviceAttr.Args)
	err = putLimitResource(&container, serviceAttr.LimitResourceTypes, serviceAttr.LimitResourceQuantities)
	err = putRequestResource(&container, serviceAttr.RequestResourceTypes, serviceAttr.RequestResourceQuantities)
	err = putLivenessProbe(&container, byteData)
	err = putReadinessProbe(&container, byteData)
	if err != nil {
		return v12.StatefulSet{}, err
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
			fmt.Println(err)
			continue
		}

		temp.ContainerPort = int32(i)
		if port.Host != "" {
			i, err = strconv.Atoi(port.Host)
			if err != nil {
				fmt.Println(err)
				continue
			}
			temp.HostPort = int32(i)
		}
		ports = append(ports, temp)
	}
	var envVariables []v1.EnvVar
	for _, envVariable := range serviceAttr.EnvironmentVariables {
		tempEnvVariable := v1.EnvVar{Name: envVariable.Key, Value: envVariable.Value}
		envVariables = append(envVariables, tempEnvVariable)
	}
	container.Ports = ports
	container.Env = envVariables
	var containers []v1.Container

	containers = append(containers, container)
	statefulset.Spec.Template.Spec.Containers = containers

	return statefulset, nil
}
func getConfigMapObject(service types.Service) (v1.ConfigMap, error) {
	var configmap = v1.ConfigMap{}
	// Label Selector

	//keel labels
	deploymentLabels := make(map[string]string)
	//deploymentLabels["keel.sh/match-tag"] = "true"
	deploymentLabels["keel.sh/policy"] = "force"
	//deploymentLabels["keel.sh/trigger"] = "poll"

	var selector metav1.LabelSelector
	labels := make(map[string]string)
	labels["app"] = service.Name
	labels["version"] = strings.ToLower(service.Version)

	if service.Name == "" {
		//Failed
		return v1.ConfigMap{}, errors.New("Service name not found")
	}
	configmap.ObjectMeta.Name = service.Name
	configmap.ObjectMeta.Labels = deploymentLabels
	selector.MatchLabels = labels

	if service.Namespace == "" {
		configmap.ObjectMeta.Namespace = "default"
	} else {
		configmap.ObjectMeta.Namespace = service.Namespace
	}
	/*statefulset.Spec.Selector = &selector
	statefulset.Spec.Template.ObjectMeta.Labels = labels
	statefulset.Spec.Template.ObjectMeta.Annotations = map[string]string{
		"sidecar.istio.io/inject": "true",
	}*/
	//

	var container v1.Container
	container.Name = service.Name
	byteData, _ := json.Marshal(service.ServiceAttributes)
	var serviceAttr types.DockerServiceAttributes
	json.Unmarshal(byteData, &serviceAttr)

	err := putCommandAndArguments(&container, serviceAttr.Command, serviceAttr.Args)
	err = putLimitResource(&container, serviceAttr.LimitResourceTypes, serviceAttr.LimitResourceQuantities)
	err = putRequestResource(&container, serviceAttr.RequestResourceTypes, serviceAttr.RequestResourceQuantities)
	err = putLivenessProbe(&container, byteData)
	err = putReadinessProbe(&container, byteData)
	if err != nil {
		return v1.ConfigMap{}, err
	}

	return configmap, nil
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

	labels := make(map[string]string)
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
	finalObj.Services.Istio = append(finalObj.Services.Istio, res)

	if service.ServiceType == "mesh" || service.ServiceType == "other" {

		res, err = getIstioObject(service)
		if err != nil {
			utils.Info.Println("There is error in deployment")
			ret.Status = append(ret.Status, "failed")
			ret.Reason = "Not a valid Istio Object. Error : " + err.Error()
			if requestType != "GET" {
				utils.SendLog(ret.Reason, "error", input.ProjectId)
			}
			return ret
		}
		finalObj.Services.Istio = append(finalObj.Services.Istio, res)

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
	utils.Info.Println(string(x))

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
	utils.Info.Println(string(x))
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
	fmt.Println(tempLivenessProbe)
	container.LivenessProbe = tempLivenessProbe.LivenessProbe

	return nil
}
func putReadinessProbe(container *v1.Container, byteData []byte) error {

	strLowerCamel := convertKeys(byteData)
	var tempReadinessProbe tempProbing
	json.Unmarshal(strLowerCamel, &tempReadinessProbe)
	fmt.Println(tempReadinessProbe)
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

func marshalUnMarshalOfIstioComponents(s string) (map[string]interface{}, error) {
	/*	s = strings.TrimSpace(s)
		s = "{" + s
		s = strings.Replace(s, " >", "}", -1)
		s = strings.Replace(s, "<", "{", -1)
		s = strings.Replace(s, " ", ",", -1)
		s = s + "}"
		for i := 0; i < len(s); i++ {
			t := '"'
			if s[i] == ':' {
				s = s[:i] + string(t) + s[i:]
				i++
			} else if s[i] == '{' || s[i] == ',' {
				s = s[:i+1] + string(t) + s[i+1:]
				i++
			}

		}*/
	utils.Info.Println(s)

	var dd map[string]interface{}
	err := json.Unmarshal([]byte(s), &dd)
	if err != nil {
		utils.Error.Println(err)
		return nil, err
	}
	return dd, nil
}
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
