package types

import (
	istioClient "istio.io/client-go/pkg/apis/networking/v1alpha3"
	v12 "k8s.io/api/apps/v1"
	autoscaling "k8s.io/api/autoscaling/v2beta2"
	v13 "k8s.io/api/batch/v1"
	"k8s.io/api/batch/v2alpha1"
	"k8s.io/api/core/v1"
	rbacV1 "k8s.io/api/rbac/v1"
	storage "k8s.io/api/storage/v1"
	"time"
)

type Route struct {
	Host   string `json:"host"`
	Subset string `json:"subset"`
}

/*
type Port struct {
	Host      string `json:"host"`
	Container string `json:"container"`
	Name      string `json:"name"`
}
*/
type SEPort struct {
	Port     int32  `json:"port"`
	Protocol string `json:"protocol"`
	Name     string `json:"name"`
}

type SEEndpoints struct {
	Address  string            `json:"address"`
	Ports    map[string]uint32 `json:"ports,omitempty"`
	Labels   map[string]string `json:"labels,omitempty"`
	Network  string            `json:"network,omitempty"`
	Locality string            `json:"locality,omitempty"`
	Weight   uint32            `json:"weight,omitempty"`
}
type VSDestination struct {
}
type VSRetries struct {
	Attempts int   `json:"attempts"`
	Timeout  int64 `json:"per_request_timeout"`
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
	Routes []VSRoute `json:"route"`
	//RewriteUri string      `json:"rewrite_uri"`
	//RetriesUri string      `json:"retries_uri"`
	Timeout        int64          `json:"timeout"`
	Match          []URI          `json:"match"`
	Retries        []VSRetries    `json:"retries"`
	FaultInjection FaultInjection `json:"fault_injection"`
}
type URI struct {
	Uris []string `json:"uri"`
}

type IstioVirtualServiceAttributes struct {
	Hosts    []string `json:"hosts"`
	Gateways []string `json:"gateways"`
	HTTP     []VSHTTP `json:"http"`
}
type FaultInjection struct {
	FaultInjectionAbort FaultInjectionAbort `json:"fault_abort"`
	FaultInjectionDelay FaultInjectionDelay `json:"fault_delay"`
}
type FaultInjectionAbort struct {
	Percentage float64 `json:"percentage"`
	HttpStatus int32   `json:"http_status"`
}
type FaultInjectionDelay struct {
	Percentage float64 `json:"percentage"`
	FixedDelay int64   `json:"fix_delay"`
}
type IstioServiceEntryAttributes struct {
	Hosts            []string      `json:"hosts"`
	Address          []string      `json:"address"`
	Ports            []SEPort      `json:"ports"`
	Uri              []SEEndpoints `json:"endpoints"`
	Location         string        `json:"location"`
	Resolution       string        `json:"resolution"`
	IsMtlsEnable     bool          `json:"is_mtls_enable"`
	MtlsMode         string        `json:"mtls_mode"`
	MtlsCertificates interface{}   `json:"mtls_certificates"`
}

/*type GWServers struct {
	Port     int    `json:"port"`
	Protocol string `json:"protocol"`
}
type IstioGatewayAttributes struct {
	Servers []GWServers `json:"servers"`
}*/
type DRSubsets struct {
	Name   string `json:"name"`
	Labels []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"labels"`
	Http1MaxPendingRequests  int32 `json:"max_pending_requests"`
	Http2MaxRequests         int32 `json:"max_requests"`
	MaxRequestsPerConnection int32 `json:"max_requests_per_connection"`
	MaxRetries               int32 `json:"max_retries"`
}
type IstioDestinationRuleAttributes struct {
	Host          string      `json:"host"`
	Subsets       []DRSubsets `json:"subsets"`
	TrafficPolicy struct {
		TLS struct {
			Mode         string   `json:"mode"`
			Certificates struct{} `json:"certificates"`
		}
	} `json:"traffic_policy"`
}

/*type HPAServiceAttributes struct {
	HPA           HPAAttributes `json:"hpa_configurations"`
}*/
type HPAAttributes struct {
	MinReplicas        int32          `json:"min_replicas"`
	MaxReplicas        int32          `json:"max_replicas"`
	Metrics_           []Metrics      `json:"metrics_values"`
	CrossObjectVersion ScaleTargetRef `json:"cross_object_version"`
}
type Metrics struct {
	TargetValueKind string `json:"target_value_kind"`
	TargetValue     string `json:"target_value"`
	ResourceKind    string `json:"resource_kind"`
}
type ScaleTargetRef struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Type    string `json:"type"`
}

type KubernetesSecret struct {
	Name       *string           `json:"name"`
	Version    *string           `json:"version"`
	Namespace  *string           `json:"namespace"`
	Type       *string           `json:"type"`
	Data       map[string]string `json:"data"`
	StringData map[string]string `json:"string_data"`
}

//type ConfigMap struct {
//	Name      *string           `json:"name"`
//	Version   *string           `json:"version"`
//	Namespace *string           `json:"namespace"`
//	Data      map[string]string `json:"data"`
//}

// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: DestinationRule
// metadata:
//   name: bookinfo-ratings-port
// spec:

type VolumeAttributes struct {
	//	Volume Volume `json:"volume"`
}

type VolumeAttributesList struct {
	//	Volume []Volume `json:"volumes"`
}

type Yamlservice struct {
	Kind string `json:"kind"`
}
type RbacAttributes struct {
	RbacService Role `json:"role"`
}

type IstioObject struct {
	ApiVersion string                 `json:"apiVersion"`
	Kind       string                 `json:"kind"`
	Metadata   map[string]interface{} `json:"metadata"`
	Spec       interface{}            `json:"spec"`
}

//type ServiceDependency struct {
//	Name           string   `json:"name"`
//	DependencyType string   `json:"dependency_type"`
//	Hosts          []string `json:"hosts"`
//	Uri            []string `json:"uri"`
//	TimeOut        string   `json:"timeout"`
//	Routes         []Route  `json:"routes"`
//	Ports          []Port   `json:"ports"`
//}

/*
type ServiceDependencyx struct {
	ServiceType       string            `json:"service_type"`
	Name              string            `json:"name"`
	Version           string `json:"version"`
	ServiceAttributes ServiceAttributes `json:"service_attributes"`
}*/

//type Service struct {
//	ServiceType           string              `json:"service_type"`
//	SubType               string              `json:"service_sub_type"`
//	Name                  string              `json:"name"`
//	ID                    string              `json:"service_id"`
//	Version               string              `json:"version"`
//	ServiceDependencyInfo []ServiceDependency `json:"service_dependency_info"`
//	ServiceAttributes     interface{}         `json:"service_attributes"`
//	Namespace             string              `json:"namespace"`
//	GroupId               string              `json:"group_id"`
//	Hostnames             []string            `json:"hostnames"`
//	Replicas              int32               `json:"replicas"`
//}
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
	Deployments            []v12.Deployment                      `json:"deployment"`
	DaemonSets             []v12.DaemonSet                       `json:"daemonsets"`
	HPA                    []autoscaling.HorizontalPodAutoscaler `json:"hpas"`
	CronJobs               []v2alpha1.CronJob                    `json:"cronjob"`
	Jobs                   []v13.Job                             `json:"job"`
	StatefulSets           []v12.StatefulSet                     `json:"statefulset"`
	ConfigMap              []v1.ConfigMap                        `json:"configmap"`
	Kubernetes             []v1.Service                          `json:"kubernetes-service"`
	Istio                  []IstioObject                         `json:"istio-component"`
	IstioGateway           []*istioClient.Gateway                `json:"gateway"`
	StorageClasses         []storage.StorageClass                `json:"storage-classes"`
	RoleClasses            []rbacV1.Role                         `json:"role-classes"`
	RoleBindingClasses     []rbacV1.RoleBinding                  `json:"role-binding-classes"`
	ServiceAccountClasses  []v1.ServiceAccount                   `json:"service-account-classes"`
	PersistentVolumeClaims []v1.PersistentVolumeClaim            `json:"persistent-volume-claims"`
	Secrets                []interface{}                         `json:"secrets"`
	Nodes                  []v1.Node                             `json:"nodes"`
}
type NooeWrapper struct {
	Error    string      `json:"error"`
	Nodes    v1.NodeList `json:"data2"`
	KubeData string      `json:"data"`
}
type DeploymentWrapper struct {
	Error       string         `json:"error"`
	Deployments v12.Deployment `json:"data2"`
	KubeData    string         `json:"data"`
}
type KubernetesWrapper struct {
	Error      string     `json:"error"`
	Kubernetes v1.Service `json:"data2"`
	KubeData   string     `json:"data"`
}
type IstioWrapper struct {
	Error    string      `json:"error"`
	Istio    IstioObject `json:"data2"`
	KubeData string      `json:"data"`
}

type OutputResp struct {
	Deployments []DeploymentWrapper `json:"deployment"`
	Kubernetes  []KubernetesWrapper `json:"kubernetes-service"`
	Istio       []IstioWrapper      `json:"istio-component"`
	Secrets     []interface{}       `json:"secrets"`
	Nodes       []NooeWrapper       `json:"nodes"`
}
type ServiceOutput struct {
	ClusterInfo KubernetesCred `json:"kubernetes_credentials"`
	ConfigMap   []v1.ConfigMap `json:"configmap"`
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
	CompanyId   string      `json:"company_id"`
	Message     interface{} `json:"message"`
	Id          string      `json:"id"`
	Environment string      `json:"environment"`
	Service     string      `json:"service"`
	Level       string      `json:"level"`
	Type        string      `json:"type"`
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
	ID        string   `json:"_id"`
	ServiceId string   `json:"service_id"`
	Name      string   `json:"name"`
	Status    []string `json:"status_individual"`
	StatusF   string   `json:"status"`
	Reason    string   `json:"reason"`
}

type ResponseRequest struct {
	Service OutputResp `json:"service"`
}
type ResponseServiceRequestMessage struct {
}
type ResponseServiceRequestFailure struct {
	Error string `json:"error"`
}

type VaultCredentialsConfigurations struct {
	Credentials BasicAuthCredentails `json:"docker_credentials"`
	Profile     string               `json:"profile_name"`
}
type BasicAuthCredentails struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type LoggingHttpRequest struct {
	RequestId string `json:"request_id"`
	//url of the cloudplex server (e.g. apis.cloudplex.cf)
	Url string `json:"url"`
	//request method (GET/POST/PUT/PATCH/DELETE)
	Method string `json:"method" `
	//request path of backend service
	Path string `json:"path"`
	//request body
	Body string `json:"body"`
	//status code of service
	Status int `json:"status"`
}
