package types

import (
	v12 "k8s.io/api/apps/v1"
	"k8s.io/api/core/v1"
	storage "k8s.io/api/storage/v1"
	"time"
)

type Route struct {
	Host   string `json:"host"`
	Subset string `json:"subset"`
}
type Port struct {
	Host      string `json:"host"`
	Container string `json:"container"`
}

type SEPort struct {
	Name     string `json:"name"`
	Port     int32  `json:"port"`
	Protocol string `json:"protocol"`
}
type SEEndpoints struct {
	Address string `json:"address"`
}
type VSDestination struct {
}
type VSRetries struct {
	Attempts int `json:"attempts"`
	Timeout  int `json:"timeouts"`
}
type VSRoute struct {
	Destination struct {
		Host   string `json:"host"`
		Subset string `json:"subset"`
		Port   int    `json:"port"`
	} `json:"destination"`
	Weight int32 `json:"weight"`
}
type VSHTTP struct {
	Routes     []VSRoute   `json:"route"`
	RewriteUri string      `json:"rewrite_uri"`
	RetriesUri string      `json:"retries_uri"`
	Timeout    int32       `json:"timeout"`
	Retries    []VSRetries `json:"retries"`
}

type IstioVirtualServiceAttributes struct {
	Hosts    []string `json:"hosts"`
	Gateways []string `json:"gateways"`
	HTTP     []VSHTTP `json:"http"`
}

type IstioServiceEntryAttributes struct {
	Hosts      []string      `json:"hosts"`
	Address    []string      `json:"address"`
	Ports      []SEPort      `json:"ports"`
	Uri        []SEEndpoints `json:"endpoints"`
	Location   string        `json:"location"`
	Resolution string        `json:"resolution"`
}
type GWServers struct {
	Hosts    []string `json:"hosts"`
	Labels   []string `json:"labels"`
	Port     int      `json:"port"`
	Protocol string   `json:"protocol"`
	Name     string   `json:"name"`
}
type IstioGatewayAttributes struct {
	Servers  []GWServers       `json:"servers"`
	Selector map[string]string `json:"selector"`
}
type DRSubsets struct {
	Name   string `json:"name"`
	Labels []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"labels"`
}
type IstioDestinationRuleAttributes struct {
	Host    string      `json:"host"`
	Subsets []DRSubsets `json:"subsets"`
}
type DockerServiceAttributes struct {
	DistributionType      string `json:"distribution_type"`
	DefaultConfigurations string `json:"default_configurations"`
	EnvironmentVariables  []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"environment_variables"`
	ImageRepositoryConfigurations ImageRepositoryConfigurations `json:"image_repository_configurations" binding:"required"`
	Ports                         []Port                        `json:"ports"`
	Files                         []string                      `json:"files"`
	Tag                           string                        `json:"tag"`
	ImagePrefix                   string                        `json:"image_prefix"`
	ImageName                     string                        `json:"image_name"`

	Command []string `json:"command"`
	Args    []string `json:"args"`

	//resource types: cpu, memory
	LimitResourceTypes        []string `json:"limitResourceTypes"`
	LimitResourceQuantities   []string `json:"limitResourceQuantities"`
	RequestResourceTypes      []string `json:"requestResourceTypes"`
	RequestResourceQuantities []string `json:"requestResourceQuantities"`

	LivenessProbe  *v1.Probe `json:"livenessProbe"`
	ReadinessProbe *v1.Probe `json:"readinessProbe"`
}

// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: DestinationRule
// metadata:
//   name: bookinfo-ratings-port
// spec:

type IstioObject struct {
	ApiVersion string                 `json:"apiVersion"`
	Kind       string                 `json:"kind"`
	Metadata   map[string]interface{} `json:"metadata"`
	Spec       interface{}            `json:"spec"`
}
type ServiceDependency struct {
	Name           string   `json:"name"`
	DependencyType string   `json:"dependency_type"`
	Hosts          []string `json:"hosts"`
	Uri            []string `json:"uri"`
	TimeOut        string   `json:"timeout"`
	Routes         []Route  `json:"routes"`
	Ports          []Port   `json:"ports"`
}

/*
type ServiceDependencyx struct {
	ServiceType       string            `json:"service_type"`
	Name              string            `json:"name"`
	Version           string `json:"version"`
	ServiceAttributes ServiceAttributes `json:"service_attributes"`
}*/
type Service struct {
	ServiceType           string              `json:"service_type"`
	SubType               string              `json:"service_sub_type"`
	Name                  string              `json:"name"`
	ID                    string              `json:"_id"`
	Version               string              `json:"version"`
	ServiceDependencyInfo []ServiceDependency `json:"service_dependency_info"`
	ServiceAttributes     interface{}         `json:"service_attributes"`
	Namespace             string              `json:"namespace"`
	Hostnames             []string            `json:"hostnames"`
	Volumes               []Volume            `json:"volumes"`
}
type SolutionInfo struct {
	ID      string  `json:"_id"`
	Name    string  `json:"name"`
	Version string  `json:"version"`
	PoolId  string  `json:"pool_id"`
	Service Service `json:"service"`
	KIP     string  `json:"kubeip"`
	KPo     string  `json:"kubeport"`
	KU      string  `json:"kubeusername"`
	KP      string  `json:"kubepassword"`
}

type ServiceInput struct {
	ClusterId    string         `json:"cluster_id"`
	ClusterName  string         `json:"cluster_name"`
	ProjectId    string         `json:"project_id"`
	SolutionInfo SolutionInfo   `json:"solution_info"`
	Creds        KubernetesCred `json:"kubernetes_credentials"`
}

type Output struct {
	Name               string    `json:"name"`
	Version            string    `json:"version"`
	PoolId             string    `json:"pool_id"`
	Service            []Service `json:"services"`
	KubernetesIp       string    `json:"kubeip"`
	KubernetesPort     string    `json:"kubeport"`
	KubernetesUsername string    `json:"kubeusername"`
	KubernetesPassword string    `json:"kubepassword"`
}
type KubernetesCred struct {
	KubernetesURL      string `json:"url"`
	KubernetesUsername string `json:"username"`
	KubernetesPassword string `json:"password"`
}
type OutputServices struct {
	Deployments            []v12.Deployment           `json:"deployment"`
	Kubernetes             []v1.Service               `json:"kubernetes-service"`
	Istio                  []IstioObject              `json:"istio-component"`
	StorageClasses         []storage.StorageClass     `json:"storage-classes"`
	PersistentVolumeClaims []v1.PersistentVolumeClaim `json:"persistent-volume-claims"`
	Secrets                []interface{}              `json:"secrets"`
}

type DeploymentWrapper struct {
	Error       string         `json:"error"`
	Deployments v12.Deployment `json:"data"`
}
type KubernetesWrapper struct {
	Error      string     `json:"error"`
	Kubernetes v1.Service `json:"data"`
}
type IstioWrapper struct {
	Error string      `json:"error"`
	Istio IstioObject `json:"data"`
}
type OutputResp struct {
	Deployments []DeploymentWrapper `json:"deployment"`
	Kubernetes  []KubernetesWrapper `json:"kubernetes-service"`
	Istio       []IstioWrapper      `json:"istio-component"`
	Secrets     []interface{}       `json:"secrets"`
}

type ServiceOutput struct {
	ClusterInfo KubernetesCred `json:"kubernetes_credentials"`
	Services    OutputServices `json:"service"`
	ProjectId   string         `json:"project_id"`
}
type APIError struct {
	ErrorCode    int
	ErrorMessage string
	CreatedAt    time.Time
}

type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

type Status struct {
	Message string `json:"status" bson:"status"`
}

type ResponseData struct {
	StatusCode int         `json:"status_code"`
	Body       interface{} `json:"body"`
	Error      error       `json:"error"`
	Status     string      `json:"status"`
}

type LoggingRequest struct {
	Message     string `json:"message"`
	Id          string `json:"id"`
	Environment string `json:"environment"`
	Service     string `json:"service"`
	Level       string `json:"level"`
}
type Notifier struct {
	Id string `json:"_id"`
	//EnvId  string `json:"environment_id"`
	Status    string `json:"status"`
	Component string `json:"component"`
}
type KubeResponse struct {
	Error  string `json:"error"`
	Status string `json:"status"`
}

type StatusRequest struct {
	ID      string   `json:"_id"`
	Name    string   `json:"name"`
	Status  []string `json:"status_individual"`
	StatusF string   `json:"status"`
	Reason  string   `json:"reason"`
}

type ResponseRequest struct {
	Service OutputResp `json:"service"`
}
type ResponseServiceRequestMessage struct {
}
type ResponseServiceRequestFailure struct {
	Error string `json:"error"`
}

type ImageRepositoryConfigurations struct {
	Url         string               `json:"url"`
	Tag         string               `json:"tag"`
	Credentials BasicAuthCredentails `json:"credentials"`
}

type BasicAuthCredentails struct {
	Username string `json:"username"`
	Password string `json:"password"`
}