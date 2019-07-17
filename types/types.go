package types

import (
	v12 "k8s.io/api/apps/v1"
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
type Port struct {
	Host      string `json:"host"`
	Container string `json:"container"`
	Name      string `json:"name"`
}

type SEPort struct {
	Port     int32  `json:"port"`
	Protocol string `json:"protocol"`
}
type SEEndpoints struct {
	Address string `json:"address"`
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
type RecourceType string

const (
	RecourceTypeMemory RecourceType = "memory"
	RecourceTypeCpu    RecourceType = "cpu"
)

type ExecAction struct {
	// Command is the command line to execute inside the container, the working directory for the
	// command  is root ('/') in the container's filesystem. The command is simply exec'd, it is
	// not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use
	// a shell, you need to explicitly call out to that shell.
	// Exit status of 0 is treated as live/healthy and non-zero is unhealthy.
	// +optional
	Command []string `json:"command,omitempty" protobuf:"bytes,1,rep,name=command"`
}
type HTTPHeader struct {
	// The header field name
	Name *string `json:"name" protobuf:"bytes,1,opt,name=name"`
	// The header field value
	Value *string `json:"value" protobuf:"bytes,2,opt,name=value"`
}

const (
	// URISchemeHTTP means that the scheme used will be http://
	URISchemeHTTP string = "HTTP"
	// URISchemeHTTPS means that the scheme used will be https://
	URISchemeHTTPS string = "HTTPS"
)

type HTTPGetAction struct {
	// Path to access on the HTTP server.
	// +optional
	Path *string `json:"path,omitempty" protobuf:"bytes,1,opt,name=path"`
	// Name or number of the port to access on the container.
	// Number must be in the range 1 to 65535.
	// Name must be an IANA_SVC_NAME.
	Port int `json:"port" protobuf:"bytes,2,opt,name=port"`
	// Host name to connect to, defaults to the pod IP. You probably want to set
	// "Host" in httpHeaders instead.
	// +optional
	Host *string `json:"host,omitempty" protobuf:"bytes,3,opt,name=host"`
	// Scheme to use for connecting to the host.
	// Defaults to HTTP.
	// +optional
	Scheme *string `json:"scheme,omitempty" protobuf:"bytes,4,opt,name=scheme,casttype=URIScheme"`
	// Custom headers to set in the request. HTTP allows repeated headers.
	// +optional
	HTTPHeaders []HTTPHeader `json:"httpHeaders,omitempty" protobuf:"bytes,5,rep,name=http_headers"`
}
type TCPSocketAction struct {
	// Number or name of the port to access on the container.
	// Number must be in the range 1 to 65535.
	// Name must be an IANA_SVC_NAME.
	Port int `json:"port" protobuf:"bytes,1,opt,name=port"`
	// Optional: Host name to connect to, defaults to the pod IP.
	// +optional
	Host *string `json:"host,omitempty" protobuf:"bytes,2,opt,name=host"`
}
type Handler struct {
	Type string `json:"handler_type"`

	// One and only one of the following should be specified.
	// Exec specifies the action to take.
	// +optional
	Exec *ExecAction `json:"exec,omitempty" protobuf:"bytes,1,opt,name=exec"`
	// HTTPGet specifies the http request to perform.
	// +optional
	HTTPGet *HTTPGetAction `json:"httpGet,omitempty" protobuf:"bytes,2,opt,name=http_get"`
	// TCPSocket specifies an action involving a TCP port.
	// TCP hooks not yet supported
	// TODO: implement a realistic TCP lifecycle hook
	// +optional
	TCPSocket *TCPSocketAction `json:"tcpSocket,omitempty" protobuf:"bytes,3,opt,name=tcp_socket"`
}

type Probe struct {
	// The action taken to determine the health of a container
	Handler *Handler `json:",inline" protobuf:"bytes,1,opt,name=handler"`
	// Number of seconds after the container has started before liveness probes are initiated.
	// More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	// +optional
	InitialDelaySeconds *int32 `json:"initialDelaySeconds,omitempty" protobuf:"varint,2,opt,name=initial_delay_seconds"`
	// Number of seconds after which the probe times out.
	// Defaults to 1 second. Minimum value is 1.
	// More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	// +optional
	TimeoutSeconds *int32 `json:"timeoutSeconds,omitempty" protobuf:"varint,3,opt,name=timeout_seconds"`
	// How often (in seconds) to perform the probe.
	// Default to 10 seconds. Minimum value is 1.
	// +optional
	PeriodSeconds *int32 `json:"periodSeconds,omitempty" protobuf:"varint,4,opt,name=period_seconds"`
	// Minimum consecutive successes for the probe to be considered successful after having failed.
	// Defaults to 1. Must be 1 for liveness. Minimum value is 1.
	// +optional
	SuccessThreshold *int32 `json:"successThreshold,omitempty" protobuf:"varint,5,opt,name=success_threshold"`
	// Minimum consecutive failures for the probe to be considered failed after having succeeded.
	// Defaults to 3. Minimum value is 1.
	// +optional
	FailureThreshold *int32 `json:"failureThreshold,omitempty" protobuf:"varint,6,opt,name=failure_threshold"`
}

type DockerServiceAttributes struct {
	DistributionType      string `json:"distribution_type"`
	DefaultConfigurations string `json:"default_configurations"`
	EnvironmentVariables  []struct {
		Key         string `json:"key"`
		Value       string `json:"value"`
		IsSecret    bool   `json:"secrets"`
		IsConfigMap bool   `json:"configmap"`
	} `json:"environment_variables"`
	ImageRepositoryConfigurations ImageRepositoryConfigurations `json:"image_repository_configurations" binding:"required"`
	Ports                         []Port                        `json:"ports"`
	Files                         []string                      `json:"files"`
	Tag                           string                        `json:"tag"`
	ImagePrefix                   string                        `json:"image_prefix"`
	ImageName                     string                        `json:"image_name"`
	MeshConfig                    IstioConfig                   `json:"istio_config"`
	LabelSelector                 LabelSelectorObj              `json:"label_selector"`
	NodeSelector                  map[string]string             `json:"node_selector"`
	Command                       []string                      `json:"command"`
	Args                          []string                      `json:"args"`
	SecurityContext               SecurityContextStruct         `json:"security_context"`
	//resource types: cpu, memory
	LimitResources        map[RecourceType]string `json:"limit_resources"`
	RequestResources      map[RecourceType]string `json:"request_resources"`
	Labels                map[string]string       `json:"labels"`
	Annotations           map[string]string       `json:"annotations"`
	CronJobScheduleString string                  `json:"cron_job_schedule_string"`
	LivenessProb          Probe                   `json:"liveness_probe"`
	RedinessProb          Probe                   `json:"readiness_probe"`

	IsRbac bool `json:"is_rbac_enabled"`

	RbacRoles []K8sRbacAttribute `json:"roles"`

	IstioRoles []IstioRbacAttribute `json:"istio_roles"`

	IsInitContainerEnable bool `json:"enable_init"`
}
type K8sRbacAttribute struct {
	Resource string   `json:"resource"`
	Verbs    []string `json:"verbs"`
	ApiGroup []string `json:"api_group"`
}
type IstioRbacAttribute struct {
	Services []string `json:"services"`
	Methods  []string `json:"methods"`
	Paths    []string `json:"paths"`
}
type KubernetesSecret struct {
	Name       *string           `json:"name"`
	Version    *string           `json:"version"`
	Namespace  *string           `json:"namespace"`
	Type       *string           `json:"type"`
	Data       map[string]string `json:"data"`
	StringData map[string]string `json:"string_data"`
}

type ConfigMap struct {
	Name      *string           `json:"name"`
	Version   *string           `json:"version"`
	Namespace *string           `json:"namespace"`
	Data      map[string]string `json:"data"`
}

type SecurityContextStruct struct {
	CapabilitiesAdd          []interface{}        `json:"capabilities_add"`
	CapabilitiesDrop         []interface{}        `json:"capabilities_drop"`
	RunAsUser                *int64               `json:"run_as_user"`
	RunAsGroup               *int64               `json:"run_as_group"`
	RunAsNonRoot             bool                 `json:"run_as_non_root"`
	Privileged               bool                 `json:"privileged"`
	ProcMount                interface{}          `json:"proc_mount"`
	AllowPrivilegeEscalation bool                 `json:"allow_privilege_escalation"`
	ReadOnlyRootFileSystem   bool                 `json:"read_only_root_filesystem"`
	SELinuxOptions           SELinuxOptionsStruct `json:"se_linux_options"`
}

type SELinuxOptionsStruct struct {
	User  string `json:"user,omitempty"`
	Role  string `json:"role,omitempty"`
	Type  string `json:"type,omitempty"`
	Level string `json:"level,omitempty"`
}

// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: DestinationRule
// metadata:
//   name: bookinfo-ratings-port
// spec:

type VolumeAttributes struct {
	Volume Volume `json:"volume"`
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
type LabelSelectorObj struct {
	MatchLabel      map[string]string          `json:"match_label"`
	MatchExpression []LabelSelectorRequirement `json:"match_expression"`
}
type LabelSelectorRequirement struct {
	Key      string                `json:"key" patchStrategy:"merge" patchMergeKey:"key" protobuf:"bytes,1,opt,name=key"`
	Operator LabelSelectorOperator `json:"operator" protobuf:"bytes,2,opt,name=operator,casttype=LabelSelectorOperator"`
	Values   []string              `json:"values,omitempty" protobuf:"bytes,3,rep,name=values"`
}

type LabelSelectorOperator string

const (
	LabelSelectorOpIn           LabelSelectorOperator = "In"
	LabelSelectorOpNotIn        LabelSelectorOperator = "NotIn"
	LabelSelectorOpExists       LabelSelectorOperator = "Exists"
	LabelSelectorOpDoesNotExist LabelSelectorOperator = "DoesNotExist"
)

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
}
type IstioConfig struct {
	Enable_External_Traffic bool `json:"enable_external_traffic"`
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
	DaemonSets             []v12.DaemonSet            `json:"daemonsets"`
	CronJobs               []v2alpha1.CronJob         `json:"cronjob"`
	Jobs                   []v13.Job                  `json:"job"`
	StatefulSets           []v12.StatefulSet          `json:"statefulset"`
	ConfigMap              []v1.ConfigMap             `json:"configmap"`
	Kubernetes             []v1.Service               `json:"kubernetes-service"`
	Istio                  []IstioObject              `json:"istio-component"`
	StorageClasses         []storage.StorageClass     `json:"storage-classes"`
	RoleClasses            []rbacV1.Role              `json:"role-classes"`
	RoleBindingClasses     []rbacV1.RoleBinding       `json:"role-binding-classes"`
	ServiceAccountClasses  []v1.ServiceAccount        `json:"service-account-classes"`
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
