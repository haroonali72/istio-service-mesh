package controllers

import (
"Istio/types"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/istio/api/networking/v1alpha3"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/util/intstr"
	"net/http"
	"strconv"
	v12 "k8s.io/api/apps/v1"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

func getIstioVirtualService(service types.Service)(v1alpha3.VirtualService , error){
	vService := v1alpha3.VirtualService{}
	byteData, _ := json.Marshal(service.ServiceAttributes)
	var serviceAttr types.IstioVirtualServiceAttributes
	json.Unmarshal(byteData, &serviceAttr)
	var routes [] *v1alpha3.HTTPRoute

	for _ , http := range serviceAttr.HTTP {

		var httpRoute v1alpha3.HTTPRoute
		var destination [] *v1alpha3.HTTPRouteDestination
		for _ , route := range http.Routes {
			var httpD v1alpha3.HTTPRouteDestination
			httpD.Destination = &v1alpha3.Destination{Subset: route.Subset,Host: route.Host , Port:&v1alpha3.PortSelector{&v1alpha3.PortSelector_Number{Number:uint32(route.Port)}}}
			if(route.Weight > 0){
				httpD.Weight = route.Weight
			}
			destination = append(destination,&httpD)
		}
		httpRoute.Route = destination
		if(http.RewriteUri != ""){
			var rewrite v1alpha3.HTTPRewrite
			rewrite.Uri = http.RewriteUri
			httpRoute.Rewrite = &rewrite
		}
		if(http.RetriesUri != ""){
			var retries v1alpha3.HTTPRetry
			retries.RetryOn = http.RetriesUri
			httpRoute.Retries = &retries
		}
		if(http.Timeout > 0 ){
			//var timeout int32
			//httpRoute.Timeout = google_protobuf.(timeout)
		}
		routes = append(routes,&httpRoute)
	}
	vService.Http = routes
	vService.Hosts = serviceAttr.Hosts
	vService.Gateways = serviceAttr.Gateways
	fmt.Println(vService.String())
	return vService , nil
}

func getIstioGateway(service types.Service)(v1alpha3.Gateway , error){
	gateway := v1alpha3.Gateway{}
	byteData, _ := json.Marshal(service.ServiceAttributes)
	var serviceAttr types.IstioGatewayAttributes
	json.Unmarshal(byteData, &serviceAttr)
	var servers []*v1alpha3.Server
	for _ , server := range serviceAttr.Servers {
		var serv v1alpha3.Server
		i, err := strconv.Atoi(server.Port)
		if err != nil {
			fmt.Println(err)
			continue
		}
		serv.Port = &v1alpha3.Port{Name:server.Name,Protocol:server.Protocol,Number:uint32(i)}
		serv.Hosts = server.Hosts
		servers = append(servers,&serv)
	}
	gateway.Selector = serviceAttr.Selector
	gateway.Servers = servers
	return gateway, nil
}
func getIstioDestinationRule(service types.Service)(v1alpha3.DestinationRule , error){
	destRule := v1alpha3.DestinationRule{}
	byteData, _ := json.Marshal(service.ServiceAttributes)
	var serviceAttr types.IstioDestinationRuleAttributes
	json.Unmarshal(byteData, &serviceAttr)
	var subsets []*v1alpha3.Subset

	for _ , subset := range serviceAttr.Subsets {
		var ss v1alpha3.Subset
		ss.Name = subset.Name
		ss.Labels = subset.Labels
		subsets = append(subsets,&ss)
	}
	destRule.Subsets = subsets
	destRule.Host = serviceAttr.Host
	return destRule , nil
}
func getIstioServiceEntry(service types.Service)(v1alpha3.ServiceEntry , error){
	SE := v1alpha3.ServiceEntry{}
	//x := v1.Service{}
	//x.Spec = SE

	byteData, _ := json.Marshal(service.ServiceAttributes)
	var serviceAttr types.IstioServiceEntryAttributes
	json.Unmarshal(byteData, &serviceAttr)
	var ports []*v1alpha3.Port
	for _ , port := range serviceAttr.Ports {
		var p v1alpha3.Port
		p.Name = port.Name
		p.Protocol = port.Protocol
		p.Number = uint32(port.Port)
		ports = append(ports,&p)
	}

	SE.Ports = ports
	SE.Hosts = serviceAttr.Hosts
	//SE.Location = serviceAttr.Location
	SE.Addresses = serviceAttr.Address
	//SE.Location = v1alpha3.ServiceEntry_Location()


	return SE , nil
}
func getIstioObject(input types.Service)(string,error){
	switch input.SubType {
	case "virtual-service":
		serv , err := getIstioVirtualService(input)
		if err != nil {
			fmt.Println("There is error in deployment")
			return string(err.Error()) , err
		}
		x, err := json.Marshal(serv)
		if err != nil {
			fmt.Println(err)
			return "" , err
		}
		return string(x) , nil
	case "gateway":
		serv , err := getIstioGateway(input)
		if err != nil {
			fmt.Println("There is error in deployment")
			return string(err.Error()) , err
		}
		x, err := json.Marshal(serv)
		if err != nil {
			fmt.Println(err)
			return "" , err
		}
		return string(x) , nil
	case "destination-rule":
		serv , err := getIstioDestinationRule(input)
		if err != nil {
			fmt.Println("There is error in deployment")
			return string(err.Error()) , err
		}
		x, err := json.Marshal(serv)
		if err != nil {
			fmt.Println(err)
			return "" , err
		}
		return string(x) , nil
	case "service-entry":
		serv , err := getIstioServiceEntry(input)
		if err != nil {
			fmt.Println("There is error in deployment")
			return string(err.Error()) , err
		}
		x, err := json.Marshal(serv)
		if err != nil {
			fmt.Println(err)
			return "" , err
		}
		return string(x) , nil
	}

	return "" , nil
	/*
	x := v1alpha3.DestinationRule{}
	vService := v1alpha3.VirtualService{}
	if input.Hostnames != nil && len(input.Hostnames) > 0 {
		vService.Hosts = input.Hostnames
	} else {
		vService.Hosts = []string{input.Name}
	}
	var httpRoute v1alpha3.HTTPRoute
	var destination [] *v1alpha3.HTTPRouteDestination

	for _ , port := range input.ServiceAttributes.Ports {
		var httpD v1alpha3.HTTPRouteDestination
		if port.Container == "" && port.Host == ""{
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
		httpD.Destination = &v1alpha3.Destination{Host: input.Name , Port:&v1alpha3.PortSelector{&v1alpha3.PortSelector_Number{Number:uint32(i)}}}
		destination = append(destination,&httpD)
	}
	httpRoute.Route = destination
	vService.

	var httpD v1alpha3.HTTPRouteDestination
	httpD

	//v1alpha3.VirtualService{}
	name := input.Name
	for _,depedency := range input.ServiceDependencyInfo {
		if(depedency.DependencyType == "external"){

		} else if(depedency.DependencyType == "routing"){
			for _,route := range depedency.Routes{
				route.
			}
		}

	}
	*/
}
func getDeploymentObject(service types.Service)(v12.Deployment , error){
	var deployment = v12.Deployment{}
	// Label Selector

	var selector metav1.LabelSelector
	labels := make(map[string]string)
	labels["app"] = service.Name

	if service.Name == ""{
		//Failed
		return v12.Deployment{} , errors.New("Service name not found")
	}
	deployment.ObjectMeta.Name = service.Name
	selector.MatchLabels = labels

	if(service.Namespace == ""){
		deployment.ObjectMeta.Namespace = "default"
	}else{
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

	container.Image = serviceAttr.ImagePrefix+ serviceAttr.ImageName
	var ports []v1.ContainerPort
	for _ , port := range serviceAttr.Ports {
		temp := v1.ContainerPort{}
		if port.Container == "" && port.Host == ""{
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
		ports = append(ports,temp)
	}
	container.Ports = ports

	var containers []v1.Container

	containers = append(containers,container)
	deployment.Spec.Template.Spec.Containers = containers


	return deployment , nil
}

func getServiceObject(input types.Service)(v1.Service,error)  {
	service := v1.Service{}
	service.Name = input.Name
	service.ObjectMeta.Name = input.Name

	if(service.Namespace == ""){
		service.ObjectMeta.Namespace = "default"
	}else{
		service.ObjectMeta.Namespace = service.Namespace
	}
	service.Spec.Type = v1.ServiceTypeClusterIP

	labels := make(map[string]string)
	labels["app"] = service.Name
	service.Spec.Selector =labels
	byteData, _ := json.Marshal(input.ServiceAttributes)
	var serviceAttr types.DockerServiceAttributes
	json.Unmarshal(byteData, &serviceAttr)
	var servicePorts []v1.ServicePort

	for _ , port := range serviceAttr.Ports {
		temp := v1.ServicePort{}
		if port.Container == "" && port.Host == ""{
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
		temp.TargetPort = intstr.IntOrString{IntVal:int32(i)}
		servicePorts = append(servicePorts,temp)
	}
	service.Spec.Ports = servicePorts
	return service,nil
}
func DeployIstio(input types.ServiceInput)(string , error){

	var results []string

	for _,service :=range input.SolutionInfo.Service{
		//**Making Service Object*//

		if service.ServiceType == "mesh"{
			res , err := getIstioObject(service)
			if err != nil {
				fmt.Println("There is error in deployment")
				return string(err.Error()) , err
			}
			results = append(results , res )
			
		} else if service.ServiceType == "docker"{
			//Getting Deployment Object
			deployment , err := getDeploymentObject(service)
			if err != nil {
				fmt.Println("There is error in deployment")
				return string(err.Error()) , err
			}
			x, err := json.Marshal(deployment)
			if err != nil {
				fmt.Println(err)
			}
			results = append(results , string(x) )

			//Getting Kubernetes Service Object
			serv , err := getServiceObject(service)
			if err != nil {
				fmt.Println("There is error in deployment")
				return string(err.Error()) , err
			}
			x, err = json.Marshal(serv)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(x)
			results = append(results , string(x) )
		}


	}

	return  "["+strings.Join(results,",")+"]",nil

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

	result , err := DeployIstio(input)
	if err != nil{
		fmt.Println(err.Error())
		w.Write([]byte(string(err.Error())))
	}else{

		fmt.Println("Deployment Successful\n")
		w.Write([]byte(result))
	}
}
