package controllers

import (
	"Istio/types"
	"Istio/utils"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/istio/api/networking/v1alpha3"
	"io/ioutil"
	v12 "k8s.io/api/apps/v1"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"net/http"
	"strconv"
	"strings"
)

var Notifier utils.Notifier

func getIstioVirtualService(service types.Service) (v1alpha3.VirtualService, error) {
	vService := v1alpha3.VirtualService{}
	byteData, _ := json.Marshal(service.ServiceAttributes)
	var serviceAttr types.IstioVirtualServiceAttributes
	json.Unmarshal(byteData, &serviceAttr)
	var routes []*v1alpha3.HTTPRoute

	for _, http := range serviceAttr.HTTP {

		var httpRoute v1alpha3.HTTPRoute
		var destination []*v1alpha3.HTTPRouteDestination
		for _, route := range http.Routes {
			var httpD v1alpha3.HTTPRouteDestination
			httpD.Destination = &v1alpha3.Destination{Subset: route.Subset, Host: route.Host, Port: &v1alpha3.PortSelector{&v1alpha3.PortSelector_Number{Number: uint32(route.Port)}}}
			if route.Weight > 0 {
				httpD.Weight = route.Weight
			}
			destination = append(destination, &httpD)
		}
		httpRoute.Route = destination
		if http.RewriteUri != "" {
			var rewrite v1alpha3.HTTPRewrite
			rewrite.Uri = http.RewriteUri
			httpRoute.Rewrite = &rewrite
		}
		if http.RetriesUri != "" {
			var retries v1alpha3.HTTPRetry
			retries.RetryOn = http.RetriesUri
			httpRoute.Retries = &retries
		}
		if http.Timeout > 0 {
			//var timeout int32
			//httpRoute.Timeout = google_protobuf.(timeout)
		}
		routes = append(routes, &httpRoute)
	}
	vService.Http = routes
	vService.Hosts = serviceAttr.Hosts
	vService.Gateways = serviceAttr.Gateways
	fmt.Println(vService.String())
	return vService, nil
}
func getIstioGateway(service types.Service) (v1alpha3.Gateway, error) {
	gateway := v1alpha3.Gateway{}
	byteData, _ := json.Marshal(service.ServiceAttributes)
	var serviceAttr types.IstioGatewayAttributes
	json.Unmarshal(byteData, &serviceAttr)
	var servers []*v1alpha3.Server
	for _, server := range serviceAttr.Servers {
		var serv v1alpha3.Server
		serv.Port = &v1alpha3.Port{Name: server.Name, Protocol: server.Protocol, Number: uint32(server.Port)}
		serv.Hosts = server.Hosts
		servers = append(servers, &serv)
	}
	gateway.Selector = serviceAttr.Selector
	gateway.Servers = servers
	return gateway, nil
}
func getIstioDestinationRule(service types.Service) (v1alpha3.DestinationRule, error) {
	destRule := v1alpha3.DestinationRule{}
	byteData, _ := json.Marshal(service.ServiceAttributes)
	var serviceAttr types.IstioDestinationRuleAttributes
	json.Unmarshal(byteData, &serviceAttr)
	var subsets []*v1alpha3.Subset

	for _, subset := range serviceAttr.Subsets {
		var ss v1alpha3.Subset
		ss.Name = subset.Name
		var labels = make(map[string]string)
		for _, label := range subset.Labels {
			labels[label.Key] = label.Value
		}
		ss.Labels = labels
		subsets = append(subsets, &ss)
	}
	destRule.Subsets = subsets
	destRule.Host = serviceAttr.Host
	destRule.Marshal()
	fmt.Println(destRule.String())
	return destRule, nil
}
func getIstioServiceEntry(service types.Service) (v1alpha3.ServiceEntry, error) {
	SE := v1alpha3.ServiceEntry{}
	//x := v1.Service{}
	//x.Spec = SE

	byteData, _ := json.Marshal(service.ServiceAttributes)
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
	//SE.Location = serviceAttr.Location
	SE.Addresses = serviceAttr.Address
	//SE.Location = v1alpha3.ServiceEntry_Location()

	return SE, nil
}
func getIstioObject(input types.Service) (types.IstioObject, error) {
	var istioServ types.IstioObject

	switch input.SubType {
	case "virtual-service":
		serv, err := getIstioVirtualService(input)
		if err != nil {
			fmt.Println("There is error in deployment")
			return istioServ, err
		}
		istioServ.Spec = serv
		labels := make(map[string]string)
		labels["app"] = strings.ToLower(input.Name)
		labels["name"] = strings.ToLower(input.Name)
		istioServ.Metadata = labels
		istioServ.Kind = "VirtualService"
		istioServ.ApiVersion = "networking.istio.io/v1alpha3"
		return istioServ, nil
	case "gateway":
		serv, err := getIstioGateway(input)
		if err != nil {
			fmt.Println("There is error in deployment")
			return istioServ, err
		}
		istioServ.Spec = serv
		labels := make(map[string]string)
		labels["app"] = strings.ToLower(input.Name)
		labels["name"] = strings.ToLower(input.Name)
		istioServ.Metadata = labels
		istioServ.Kind = "Gateway"
		istioServ.ApiVersion = "networking.istio.io/v1alpha3"
		return istioServ, nil

	case "destination-rule":
		serv, err := getIstioDestinationRule(input)
		if err != nil {
			fmt.Println("There is error in deployment")
			return istioServ, err
		}
		istioServ.Spec = serv
		labels := make(map[string]string)
		labels["app"] = strings.ToLower(input.Name)
		labels["name"] = strings.ToLower(input.Name)
		istioServ.Metadata = labels
		istioServ.Kind = "DestinationRule"
		istioServ.ApiVersion = "networking.istio.io/v1alpha3"
		return istioServ, nil

	case "service-entry":
		serv, err := getIstioServiceEntry(input)
		if err != nil {
			fmt.Println("There is error in deployment")
			return istioServ, err
		}
		istioServ.Spec = serv
		istioServ.Spec = serv
		labels := make(map[string]string)
		labels["name"] = strings.ToLower(input.Name)
		labels["app"] = strings.ToLower(input.Name)
		istioServ.Metadata = labels
		istioServ.Kind = "ServiceEntry"
		istioServ.ApiVersion = "networking.istio.io/v1alpha3"
		return istioServ, nil

	}
	return istioServ, nil
}
func getDeploymentObject(service types.Service) (v12.Deployment, error) {
	var deployment = v12.Deployment{}
	// Label Selector

	var selector metav1.LabelSelector
	labels := make(map[string]string)
	labels["app"] = service.Name

	if service.Name == "" {
		//Failed
		return v12.Deployment{}, errors.New("Service name not found")
	}
	deployment.ObjectMeta.Name = service.Name
	selector.MatchLabels = labels

	if service.Namespace == "" {
		deployment.ObjectMeta.Namespace = "default"
	} else {
		deployment.ObjectMeta.Namespace = service.Namespace
	}
	deployment.Spec.Selector = &selector
	deployment.Spec.Template.ObjectMeta.Labels = labels
	//

	var container v1.Container
	container.Name = service.Name
	byteData, _ := json.Marshal(service.ServiceAttributes)
	var serviceAttr types.DockerServiceAttributes
	json.Unmarshal(byteData, &serviceAttr)

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
		if port.Container != "" && port.Host == "" {
			port.Host = port.Container
		}

		i, err := strconv.Atoi(port.Container)
		if err != nil {
			fmt.Println(err)
			continue
		}
		temp.ContainerPort = int32(i)
		i, err = strconv.Atoi(port.Host)
		if err != nil {
			fmt.Println(err)
			continue
		}
		temp.HostPort = int32(i)
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

func getServiceObject(input types.Service) (v1.Service, error) {
	service := v1.Service{}
	service.Name = input.Name
	service.ObjectMeta.Name = input.Name

	if service.Namespace == "" {
		service.ObjectMeta.Namespace = "default"
	} else {
		service.ObjectMeta.Namespace = service.Namespace
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
		if port.Container != "" && port.Host == "" {
			port.Host = port.Container
		}

		i, err := strconv.Atoi(port.Container)
		if err != nil {
			fmt.Println(err)
			continue
		}
		temp.Port = int32(i)
		i, err = strconv.Atoi(port.Host)
		if err != nil {
			fmt.Println(err)
			continue
		}
		temp.TargetPort = intstr.IntOrString{IntVal: int32(i)}
		servicePorts = append(servicePorts, temp)
	}
	service.Spec.Ports = servicePorts
	return service, nil
}
func DeployIstio(input types.ServiceInput) (string, error) {
	var finalObj types.ServiceOutput

	finalObj.ClusterInfo.KubernetesURL = input.Creds.KubernetesURL
	finalObj.ClusterInfo.KubernetesUsername = input.Creds.KubernetesUsername
	finalObj.ClusterInfo.KubernetesPassword = input.Creds.KubernetesPassword

	//for _,service :=range input.SolutionInfo.Service{
	service := input.SolutionInfo.Service
	//**Making Service Object*//
	if service.ServiceType == "mesh" {

		res, err := getIstioObject(service)
		if err != nil {
			fmt.Println("There is error in deployment")
			return string(err.Error()), err
		}
		finalObj.Services.Istio = append(finalObj.Services.Istio, res)

	} else if service.ServiceType == "docker" {
		//Getting Deployment Object
		deployment, err := getDeploymentObject(service)
		if err != nil {
			fmt.Println("There is error in deployment")
			return string(err.Error()), err
		}
		finalObj.Services.Deployments = append(finalObj.Services.Deployments, deployment)

		//Getting Kubernetes Service Object
		serv, err := getServiceObject(service)
		if err != nil {
			fmt.Println("There is error in deployment")
			return string(err.Error()), err
		}
		finalObj.Services.Kubernetes = append(finalObj.Services.Kubernetes, serv)
	}

	//Send request to Kubernetes
	x, err := json.Marshal(finalObj)
	if err != nil {
		fmt.Println(err)
	}
	if !ForwardToKube(x, input.EnvId) {
		return string(x), errors.New("Kubernetes Deployment Failed")
	}
	return string(x), nil

}
func ForwardToKube(requestBody []byte, env_id string) bool {

	url := "http://kubernetes-service-engine:8089/api/v1/kubernetes/deploy"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	//tr := &http.Transport{
	//	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	//}
	client := &http.Client{}
	//client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		utils.SendLog("Connection to kubernetes microservice failed "+err.Error(), "info", env_id)

		return false
		/*
			Info.Println(err)
			Info.Println(reflect.TypeOf(resp))
		*/

	} else {
		statusCode := resp.StatusCode

		//Info.Printf("notification status code %d\n", statusCode)
		result, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			utils.SendLog("Response Parsing failed "+err.Error(), "error", env_id)
			return false
		} else {
			utils.Info.Println(string(result))
			utils.SendLog(string(result), "info", env_id)
			if statusCode != 200 {
				return false
			}
			/*var kubresponse types.KubeResponse
			err1 := json.Unmarshal(result, &kubresponse)
			if err1 != nil {
				utils.Info.Println(err1)
				utils.Error.Println("Notification Parsing failed")
				return false
			}else{
				if(kubresponse.Status == "service deployment failed"){
					return false
				}else {
					return true
				}
			}*/
		}
	}
	return true
}
func ServiceRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Body)
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Println(string(b))
	// Unmarshal
	var input types.ServiceInput
	err = json.Unmarshal(b, &input)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var notification types.Notifier
	notification.EnvId = input.EnvId
	notification.Id = input.SolutionInfo.Service.ID

	result, err := DeployIstio(input)
	if err != nil {
		fmt.Println(err.Error())
		w.Write([]byte(string(err.Error())))
		notification.Status = "fail"
	} else {
		fmt.Println("Deployment Successful\n")
		w.Write([]byte(result))
		notification.Status = "success"
	}
	b, err1 := json.Marshal(notification)
	if err1 != nil {
		utils.Info.Println(err1)
		utils.Error.Println("Notification Parsing failed")
	} else {
		Notifier.Notify(input.SolutionInfo.Name, string(b))
		utils.Info.Println(string(b))

	}
}
