package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	v12 "k8s.io/api/apps/v1"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"log"
	"net/http"
	"strconv"
)

type Route struct {
	Host      string `json:"host"`
	Subset string `json:"subset"`
}
type Port struct {
	Host      string `json:"host"`
	Container string `json:"container"`
}
type ServiceAttributes struct {
	DistributionType      string            `json:"distribution_type"`
	DefaultConfigurations string            `json:"default_configurations"`
	EnvironmentVariables  map[string]string `json:"environment_variables"`
	Ports                 []Port            `json:"ports"`
	Files                 []string          `json:"files"`
	Tag                   string            `json:"tag"`
	ImagePrefix           string            `json:"image_prefix"`
	ImageName             string            `json:"image_name"`
}


type ServiceDependency struct {
	Name              string            `json:"name"`
	DependencyType    string            `json:"dependency_type"`
	Hosts           []string            `json:"hosts"`
	Uri             []string            `json:"uri"`
	TimeOut           string            `json:"timeout"`
	Routes           []Route            `json:"routes"`
	Ports             []Port            `json:"ports"`

}
/*
type ServiceDependencyx struct {
	ServiceType       string            `json:"service_type"`
	Name              string            `json:"name"`
	Version           string `json:"version"`
	ServiceAttributes ServiceAttributes `json:"service_attributes"`
}*/
type Service struct {
	ServiceType           string            `json:"service_type"`
	Name                  string            `json:"name"`
	Version               string            `json:"version"`
	ServiceDependencyInfo []ServiceDependency `json:"service_dependency_info"`
	ServiceAttributes     ServiceAttributes `json:"service_attributes"`
	Namespace                  string            `json:"namespace"`
	Hostnames             []string            `json:"hostnames"`

}
type SolutionInfo struct {
	Name                  string            `json:"name"`
	Version               string            `json:"version"`
	PoolId               string            `json:"pool_id"`

	Service []Service `json:"services"`
	KubernetesIp string            `json:"kubeip"`
	KubernetesPort string            `json:"kubeport"`
	KubernetesUsername string            `json:"kubeusername"`
	KubernetesPassword string            `json:"kubepassword"`
}

type ServiceInput struct {
	KubernetesPassword string            `json:"cluster_id"`
	SolutionInfo SolutionInfo `json:"solution_info"`
}


func getDeploymentObject(service Service)(v12.Deployment , error){
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
	container.Image = service.ServiceAttributes.ImagePrefix+ service.ServiceAttributes.ImageName
	var ports []v1.ContainerPort
	for _ , port := range service.ServiceAttributes.Ports {
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

func getIstioObject(input Service){

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

func getServiceObject(input Service)(v1.Service,error)  {
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

	var servicePorts []v1.ServicePort

	for _ , port := range input.ServiceAttributes.Ports {
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
func DeployIstio(input ServiceInput)(error){

	for _,service :=range input.SolutionInfo.Service{
		//**Making Service Object*//

		//Getting Deployment Object
		deployment , err := getDeploymentObject(service)
		if err != nil {
			fmt.Println("There is error in deployment")
			return err
		}
		x, err := json.Marshal(deployment)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(x))
		_ = deployment
		//Getting Kubernetes Service Object
		serv , err := getServiceObject(service)
		if err != nil {
			fmt.Println("There is error in deployment")
			return err
		}
		x, err = json.Marshal(serv)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(x))
		_ = serv


	}

	return  nil

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
	var input ServiceInput
	err = json.Unmarshal(b, &input)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	x, err := json.Marshal(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(x))
	err = DeployIstio(input)
	if err != nil{
		fmt.Println(err.Error())
		w.Write([]byte(string(err.Error())))
	}else{
		fmt.Println("Deployment Successful\n"+string(x))
		w.Write([]byte(string("Deployment Successful\n"+string(x))))
	}
}

func main() {
	/*client, err := kubernetes.NewForConfig(&rest.Config{Host: "https://3.84.228.162:6443", Username: "cloudplex", Password: "64bdySICej", TLSClientConfig: rest.TLSClientConfig{Insecure: true}})
	fmt.Println(err)
	pods, err := client.CoreV1().Pods("default").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for i := range pods.Items {
		fmt.Println(pods.Items[i].Name, pods.Items[i].Namespace)
	}
	os.Exit(0)*/


	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", ServiceRequest)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
