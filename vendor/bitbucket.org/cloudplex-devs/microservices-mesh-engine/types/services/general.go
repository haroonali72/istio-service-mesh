package services

import "k8s.io/apimachinery/pkg/util/intstr"

type CommonContainerAttributes struct {
	IsInitContainerEnable bool `json:"enable_init,omitempty" bson:"enable_init,omitempty"`
	// List of containers belonging to the pod.
	// Containers cannot currently be added or removed.
	// There must be at least one container in a Pod.
	// Once Deployed cannot be updated.
	// +mandatory
	Containers []*ContainerAttribute `json:"containers" bson:"containers" jsonschema:"minItems=1"`
	// List of initialization containers belonging to the pod.
	// Init containers are executed in order prior to containers being started. If any
	// init container fails, the pod is considered to have failed and is handled according
	// to its restartPolicy. The name for an init container or normal container must be
	// unique among all containers.
	// Init containers may not have Lifecycle actions, Readiness probes, Liveness probes, or Startup probes.
	// The resourceRequirements of an init container are taken into account during scheduling
	// by finding the highest request/limit for each resource type, and then using the max of
	// of that value or the sum of the normal containers. Limits are applied to init containers
	// in a similar fashion.
	// Init containers cannot currently be added or removed.
	// Once deployed cannot be updated.
	// +optional
	InitContainers []*ContainerAttribute `json:"init_containers,omitempty" bson:"init_containers,omitempty"`
	// internal use only. Cloudplex automatically populated this array
	// +optional
	Volumes []Volume `json:"volumes,omitempty" bson:"volumes,omitempty"`
	// this option is to show info on Frontend
	// +optional
	MeshConfig                   *IstioConfig                  `json:"istio_config,omitempty" bson:"istio_config,omitempty"`
	LabelSelector                *LabelSelectorObj             `json:"label_selector,omitempty" bson:"label_selector,omitempty"`
	NodeSelector                 map[string]string             `json:"node_selector,omitempty" bson:"node_selector,omitempty"`
	Labels                       map[string]string             `json:"labels,omitempty" bson:"labels,omitempty"`
	Annotations                  map[string]string             `json:"annotations,omitempty" bson:"annotations,omitempty"`
	RbacRoles                    []K8sRbacAttribute            `json:"roles,omitempty" bson:"roles,omitempty"`
	IstioRoles                   []IstioRbacAttribute          `json:"istio_roles,omitempty" bson:"istio_roles,omitempty"`
	IsRbac                       bool                          `json:"is_rbac_enabled,omitempty" bson:"is_rbac_enabled,omitempty"`
	Affinity                     *Affinity                     `json:"affinity,omitempty" bson:"affinity,omitempty"`
	ImagePullSecrets             []LocalObjectReference        `json:"image_pull_secrets,omitempty" bson:"image_pull_secrets,omitempty"`
	ServiceAccountName           string                        `json:"service_account_name,omitempty" bson:"service_account_name,omitempty"`
	AutomountServiceAccountToken *AutomountServiceAccountToken `json:"automount_service_account_token,omitempty" bson:"automount_service_account_token,omitempty"`
	// this option is a metadata to attach gateway with container service
	// +optional
	EnableExternalTraffic bool `json:"enable_external_traffic,omitempty" bson:"enable_external_traffic,omitempty"`
}
type ImageRepositoryConfigurations struct {
	Url         string               `json:"url,omitempty"`
	Tag         string               `json:"tag,omitempty"`
	Credentials BasicAuthCredentials `json:"credentials,omitempty"`
	Profile     string               `json:"profile_id,omitempty" bson:"profile_id,omitempty"`
}

type BasicAuthCredentials struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}
type EnvironmentVariable struct {
	Key     string `json:"key" bson:"key"`
	Value   string `json:"value" bson:"value"`
	Dynamic bool   `json:"dynamic" bson:"dynamic"`
	Type    string `json:"type,omitempty" bson:"type,omitempty"`
}
type IstioConfig struct {
	Enable_External_Traffic bool `json:"enable_external_traffic,omitempty" bson:"enable_external_traffic,omitempty"`
}

type LabelSelectorObj struct {
	MatchLabels      map[string]string          `json:"match_labels,omitempty" bson:"match_labels,omitempty"`
	MatchExpressions []LabelSelectorRequirement `json:"match_expressions,omitempty" bson:"match_expressions,omitempty"`
}
type LabelSelectorRequirement struct {
	Key      string                `json:"key" bson:"key" patchStrategy:"merge" patchMergeKey:"key" protobuf:"bytes,1,opt,name=key"`
	Operator LabelSelectorOperator `json:"operator" bson:"operator" protobuf:"bytes,2,opt,name=operator,casttype=LabelSelectorOperator"`
	Values   []string              `json:"values,omitempty" bson:"values,omitempty" protobuf:"bytes,3,rep,name=values"`
}

type LabelSelectorOperator string

const (
	LabelSelectorOpIn           LabelSelectorOperator = "In"
	LabelSelectorOpNotIn        LabelSelectorOperator = "NotIn"
	LabelSelectorOpExists       LabelSelectorOperator = "Exists"
	LabelSelectorOpDoesNotExist LabelSelectorOperator = "DoesNotExist"
)

type SecurityContextStruct struct {
	Capabilities             *Capabilities        `json:"capabilities,omitempty" bson:"capabilities,omitempty"`
	RunAsUser                *int64               `json:"run_as_user,omitempty" bson:"run_as_user,omitempty"`
	RunAsGroup               *int64               `json:"run_as_group,omitempty" bson:"run_as_group,omitempty"`
	RunAsNonRoot             bool                 `json:"run_as_non_root,omitempty" bson:"run_as_non_root,omitempty"`
	Privileged               bool                 `json:"privileged, omitempty" bson:"privileged, omitempty"`
	ProcMount                ProcMountType        `json:"proc_mount,omitempty" bson:"proc_mount,omitempty" `
	AllowPrivilegeEscalation bool                 `json:"allow_privilege_escalation,omitempty" bson:"allow_privilege_escalation,omitempty"`
	ReadOnlyRootFileSystem   bool                 `json:"read_only_root_filesystem,omitempty" bson:"read_only_root_filesystem,omitempty"`
	SELinuxOptions           SELinuxOptionsStruct `json:"se_linux_options,omitempty" bson:"se_linux_options,omitempty"`
}

type ProcMountType string

const (
	DefaultProcMount  ProcMountType = "Default"
	UnmaskedProcMount ProcMountType = "Unmasked"
)

type SELinuxOptionsStruct struct {
	User  string `json:"user,omitempty" bson:"user,omitempty"`
	Role  string `json:"role,omitempty" bson:"role,omitempty"`
	Type  string `json:"type,omitempty" bson:"type,omitempty"`
	Level string `json:"level,omitempty" bson:"level,omitempty"`
}

type Capability string

type Capabilities struct {
	Add  []Capability `json:"add,omitempty" bson:"add,omitempty"`
	Drop []Capability `json:"drop,omitempty" bson:"drop,omitempty"`
}

const (
	ResourceTypeMemory string = "memory"
	ResourceTypeCpu    string = "cpu"
)

//type ResourceType string
//
//const (
//	RecourceTypeMemory ResourceType = "memory"
//	RecourceTypeCpu    ResourceType = "cpu"
//)

type ExecAction struct {
	// Command is the command line to execute inside the container, the working directory for the
	// command  is root ('/') in the container's filesystem. The command is simply exec'd, it is
	// not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use
	// a shell, you need to explicitly call out to that shell.
	// Exit status of 0 is treated as live/healthy and non-zero is unhealthy.
	// +optional
	Command []string `json:"command,omitempty" bson:"command,omitempty" protobuf:"bytes,1,rep,name=command"`
}
type HTTPHeader struct {
	// The header field name
	Name *string `json:"name" bson:"name" protobuf:"bytes,1,opt,name=name"`
	// The header field value
	Value *string `json:"value" bson:"value" protobuf:"bytes,2,opt,name=value"`
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
	Path *string `json:"path,omitempty" bson:"path,omitempty" protobuf:"bytes,1,opt,name=path"`
	// Name or number of the port to access on the container.
	// Number must be in the range 1 to 65535.
	// Name must be an IANA_SVC_NAME.
	Port int `json:"port" bson:"port" protobuf:"bytes,2,opt,name=port" jsonschema:"minimum=1,maximum=65535"`
	// Host name to connect to, defaults to the pod IP. You probably want to set
	// "Host" in httpHeaders instead.
	// +optional
	Host *string `json:"host,omitempty" bson:"host,omitempty" protobuf:"bytes,3,opt,name=host"`
	// Scheme to use for connecting to the host.
	// Defaults to HTTP.
	// +optional
	Scheme *string `json:"scheme,omitempty" bson:"scheme,omitempty" protobuf:"bytes,4,opt,name=scheme,casttype=URIScheme"`
	// Custom headers to set in the request. HTTP allows repeated headers.
	// +optional
	HTTPHeaders []HTTPHeader `json:"http_headers,omitempty" bson:"http_headers,omitempty" protobuf:"bytes,5,rep,name=http_headers"`
}
type TCPSocketAction struct {
	// Number or name of the port to access on the container.
	// Number must be in the range 1 to 65535.
	// Name must be an IANA_SVC_NAME.
	Port int `json:"port" bson:"port" protobuf:"bytes,1,opt,name=port" jsonschema:"minimum=1,maximum=65535"`
	// Optional: Host name to connect to, defaults to the pod IP.
	// +optional
	Host *string `json:"host,omitempty" bson:"host,omitempty" protobuf:"bytes,2,opt,name=host"`
}
type Handler struct {
	Type string `json:"handler_type" bson:"handler_type"`

	// One and only one of the following should be specified.
	// Exec specifies the action to take.
	// +optional
	Exec *ExecAction `json:"exec,omitempty" bson:"exec,omitempty" protobuf:"bytes,1,opt,name=exec"`
	// HTTPGet specifies the http request to perform.
	// +optional
	HTTPGet *HTTPGetAction `json:"http_get,omitempty" bson:"http_get,omitempty" protobuf:"bytes,2,opt,name=http_get"`
	// TCPSocket specifies an action involving a TCP port.
	// TCP hooks not yet supported
	// +optional
	TCPSocket *TCPSocketAction `json:"tcp_socket,omitempty"  bson:"tcp_socket,omitempty" protobuf:"bytes,3,opt,name=tcp_socket"`
}

type Probe struct {
	// The action taken to determine the health of a container
	Handler *Handler `json:"handler,inline" bson:"handler,inline" protobuf:"bytes,1,opt,name=handler"`
	// Number of seconds after the container has started before liveness probes are initiated.
	// More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	// +optional
	InitialDelaySeconds *int32 `json:"initial_delay_seconds,omitempty" bson:"initial_delay_seconds,omitempty" protobuf:"varint,2,opt,name=initial_delay_seconds"`
	// Number of seconds after which the probe times out.
	// Defaults to 1 second. Minimum value is 1.
	// More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	// +optional
	TimeoutSeconds *int32 `json:"timeout_seconds,omitempty" bson:"timeout_seconds,omitempty" protobuf:"varint,3,opt,name=timeout_seconds"`
	// How often (in seconds) to perform the probe.
	// Default to 10 seconds. Minimum value is 1.
	// +optional
	PeriodSeconds *int32 `json:"period_seconds,omitempty" bson:"period_seconds,omitempty" protobuf:"varint,4,opt,name=period_seconds"`
	// Minimum consecutive successes for the probe to be considered successful after having failed.
	// Defaults to 1. Must be 1 for liveness. Minimum value is 1.
	// +optional
	SuccessThreshold *int32 `json:"success_threshold,omitempty"  bson:"success_threshold,omitempty" protobuf:"varint,5,opt,name=success_threshold"`
	// Minimum consecutive failures for the probe to be considered failed after having succeeded.
	// Defaults to 3. Minimum value is 1.
	// +optional
	FailureThreshold *int32 `json:"failure_threshold,omitempty" bson:"failure_threshold,omitempty" protobuf:"varint,6,opt,name=failure_threshold"`
}

type ContainerAttribute struct {
	// Environment Variables for container service
	// List of environment variables to set in the container.
	// Example:
	//{
	//	"PORT":{ //<env-variable-name> sample value is PORT
	//		"key":"PORT",
	//		"value":"3000",
	//		"dynamic":false
	//	},
	//	"MYSQL_IP":{
	//		"key":"MYSQL_IP",
	//		"value":"{{<service-id>.name}}",
	//		"dynamic":true
	//	}
	//}
	//+optional

	ContainerName string `json:"container_name"bson:"container_name"`

	EnvironmentVariables []EnvironmentVariable `json:"environment_variables,omitempty" bson:"environment_variables,omitempty"`
	// image repository configurations. If you have private Docker Registry
	// create docker registry creation from secrets-management API and provide
	// profile_id information in this object
	// +optional
	ImageRepositoryConfigurations *ImageRepositoryConfigurations `json:"image_repository_configurations,omitempty" bson:"image_repository_configurations,omitempty" binding:"-"`
	// List of ports to expose from the container. Exposing a port here gives
	// the system additional information about the network connections a
	// container uses, but is primarily informational. Not specifying a port here
	// DOES NOT prevent that port from being exposed. Any port which is
	// listening on the default "0.0.0.0" address inside a container will be
	// accessible from the network.
	// Example:
	//{
	//	"<port-name>":{
	//		"container_port":9080,
	//		"protocol":"TCP",
	//	}
	//}
	// This syntax is to make dynamic parameter replacement simple
	// +optional
	Ports []ContainerPort `json:"ports,omitempty" bson:"ports,omitempty"`
	// Docker image tag
	// More info: https://kubernetes.io/docs/concepts/containers/images
	// +mandatory
	Tag string `json:"tag" bson:"tag" binding:"required"`
	// +deprecated
	//+optional
	ImagePrefix string `json:"image_prefix,omitempty" bson:"image_prefix,omitempty"`
	// Docker image name.
	// More info: https://kubernetes.io/docs/concepts/containers/images
	// container images in workload controllers like Deployments and StatefulSets.
	// +mandatory
	ImageName string `json:"image_name" bson:"image_name" binding:"required"`
	// Entrypoint array. Not executed within a shell.
	// The docker image's ENTRYPOINT is used if this is not provided.
	// Variable references $(VAR_NAME) are expanded using the container's environment. If a variable
	// cannot be resolved, the reference in the input string will be unchanged.
	// Cannot be updated.
	// More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
	// +optional
	Command []string `json:"command,omitempty" bson:"command,omitempty"`
	// Arguments to the entrypoint.
	// The docker image's CMD is used if this is not provided.
	// Variable references $(VAR_NAME) are expanded using the container's environment. If a variable
	// cannot be resolved, the reference in the input string will be unchanged.
	// More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
	// +optional
	Args []string `json:"args,omitempty" bson:"args,omitempty"`
	// Limits describes the maximum amount of compute resources allowed.
	// More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/
	// Example:
	// 	{
	//		"cpu":"200m",
	//		"memory":"200Mi",
	//	}
	// +optional
	LimitResources map[string]string `json:"limit_resources,omitempty" bson:"limit_resources,omitempty" jsonschema:"maxProperties=2"`
	// Requests describes the minimum amount of compute resources required.
	// If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,
	// otherwise to an implementation-defined value.
	// More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/
	// Example:
	// 	{
	//		"cpu":"200m",
	//		"memory":"200Mi",
	//	}
	// +optional
	RequestResources map[string]string `json:"request_resources,omitempty" bson:"request_resources,omitempty"  jsonschema:"maxProperties=2"`
	// Periodic probe of container liveness.
	// Container will be restarted if the probe fails.
	// Cannot be updated.
	// More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	// +optional
	LivenessProbe *Probe `json:"liveness_probe,omitempty" bson:"liveness_probe,omitempty"`
	// Periodic probe of container service readiness.
	// Container will be removed from service endpoints if the probe fails.
	// Cannot be updated.
	// More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	// +optional
	ReadinessProbe *Probe `json:"readiness_probe,omitempty" bson:"readiness_probe,omitempty"`
	// Security options the pod should run with.
	// More info: https://kubernetes.io/docs/concepts/policy/security-context/
	// More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
	// +optional
	SecurityContext *SecurityContextStruct `json:"security_context,omitempty" bson:"security_context,omitempty"`
	// Pod volumes to mount into the container's filesystem.
	// +optional
	VolumeMounts []VolumeMount `json:"volume_mounts,omitempty" bson:"volume_mounts,omitempty"`
	//enable pipeline
	// +optional
	IsEnabledPipeline bool `json:"is_enabled_pipeline,omitempty" bson:"is_enabled_pipeline,omitempty"`
	// CD deployment pipeline information
	// +optional
	DeploymentPipeline *DeploymentPipeline `json:"deployment_pipeline,omitempty" bson:"deployment_pipeline,omitempty"`
	// is deployment pipleline Canary/BlueGreen
	// +optional
	IsDeploymentPipeline bool `json:"is_deployment_pipeline,omitempty" bson:"is_deployment_pipeline,omitempty"`
	// +optional
	IsDirectDeployment bool `json:"is_direct_deployment,omitempty" bson:"is_direct_deployment,omitempty"`
}
type DeploymentPipeline struct {
	Name        string        `json:"name,omitempty" bson:"name,omitempty"`
	Description string        `json:"description,omitempty" bson:"description,omitempty"`
	Type        string        `json:"type,omitempty" bson:"type,omitempty"`
	Bluegreen   *Bluegreen    `json:"bluegreen,omitempty" bson:"bluegreen,omitempty"`
	Canary      *Canary       `json:"canary,omitempty" bson:"canary,omitempty"`
	Duration    string        `json:"duration,omitempty" bson:"duration,omitempty"`
	RunsHistory []RunsHistory `json:"runs_history,omitempty" bson:"runs_history,omitempty"`
	TotalRuns   int           `json:"total_runs,omitempty" bson:"total_runs,omitempty"`
	Status      string        `json:"status,omitempty" bson:"status,omitempty"`
}

type RunsHistory struct {
	ActivityType string      `json:"activity_type,omitempty" bson:"activity_type,omitempty"`
	Tag          string      `json:"tag,omitempty" bson:"tag,omitempty"`
	LastRun      interface{} `json:"last_run,omitempty" bson:"last_run,omitempty"`
}

type Canary struct {
	TotalStages  int      `json:"total_stages,omitempty" bson:"total_stages,omitempty"`
	Stages       []*Stage `json:"stages,omitempty" bson:"stages,omitempty"`
	CurrentStage int      `json:"current_stage,omitempty" bson:"current_stage,omitempty"`
	Status       string   `json:"status,omitempty" bson:"status,omitempty"`
}

type Stage struct {
	Name                    string `json:"name,omitempty" bson:"name,omitempty"`
	AnalysisType            string `json:"analysis_type,omitempty" bson:"analysis_type,omitempty"`
	TrafficWeightCurrent    int32  `json:"traffic_weight_current,omitempty" bson:"traffic_weight_current,omitempty"`
	TrafficWeightBaseline   int32  `json:"traffic_weight_baseline,omitempty" bson:"traffic_weight_baseline,omitempty"`
	TrafficWeightCanary     int32  `json:"traffic_weight_canary,omitempty" bson:"traffic_weight_canary,omitempty"`
	RequestSuccessThreshold int    `json:"request_success_threshold,omitempty" bson:"request_success_threshold,omitempty"`
	RequestLatencyThreshold int    `json:"request_latency_threshold,omitempty" bson:"request_latency_threshold,omitempty"`
	CronExpression          string `json:"cron_expression,omitempty" bson:"cron_expression,omitempty"`
	Status                  string `json:"status,omitempty" bson:"status,omitempty"`
}

type Bluegreen struct {
	TrafficWeightBluegreen int32  `json:"traffic_weight_bluegreen,omitempty" bson:"traffic_weight_bluegreen,omitempty"`
	TrafficWeightCurrent   int32  `json:"traffic_weight_current,omitempty" bson:"traffic_weight_current,omitempty"`
	Status                 string `json:"status,omitempty" bson:"status,omitempty"`
	RollBack               bool   `json:"roll_back,omitempty" bson:"roll_back,omitempty"`
}
type VolumeMount struct {
	Name             string                `json:"name" bson:"name" binding:"required" valid:"alphanumspecial,length(5|30),lowercase~lowercase alphanumeric characters are allowed,required"`
	ReadOnly         bool                  `json:"readonly,omitempty" bson:"readonly,omitempty"`
	MountPath        string                `json:"mount_path" bson:"mount_path"`
	SubPath          string                `json:"sub_path,omitempty" bson:"sub_path,omitempty"`
	MountPropagation *MountPropagationMode `json:"mount_propagation,omitempty" bson:"mount_propagation,omitempty"`
	SubPathExpr      string                `json:"sub_path_expr,omitempty" bson:"sub_path_expr,omitempty"`
	PvcSvcName       string                `json:"persistent_volume_claim_name,omitempty" bson:"persistent_volume_claim_name,omitempty"`
	ConfigMap        *ConfigMapVolumeMount `json:"configmap,omitempty"`
	Secret           *ConfigMapVolumeMount `json:"secrets,omitempty"`
}

type ConfigMapVolumeMount struct {
	ConfigMapName string     `json:"name"`
	Optional      bool       `json:"optional"`
	Items         []KeyItems `json:"items"`
}
type KeyItems struct {
	Key  string `json:"key"`
	Path string `json:"path"`
}
type MountPropagationMode string

const (
	MountPropagationNone            MountPropagationMode = "None"
	MountPropagationHostToContainer MountPropagationMode = "HostToContainer"
	MountPropagationBidirectional   MountPropagationMode = "Bidirectional"
)

func (c *MountPropagationMode) String() string {
	return string(*c)
}

type K8sRbacAttribute struct {
	Resource string   `json:"resource,omitempty" bson:"resource,omitempty"`
	Verbs    []string `json:"verbs,omitempty" bson:"verbs,omitempty"`
	ApiGroup []string `json:"api_group,omitempty" bson:"api_group,omitempty"`
}
type IstioRbacAttribute struct {
	Services []string `json:"services,omitempty" bson:"services,omitempty"`
	Methods  []string `json:"methods,omitempty" bson:"methods,omitempty"`
	Paths    []string `json:"paths,omitempty" bson:"paths,omitempty"`
}

type ContainerPort struct {
	// Number of port to expose on the host.
	// If specified, this must be a valid port number, 0 < x < 65536.
	// If HostNetwork is specified, this must match ContainerPort.
	// Most containers do not need this.
	// +optional
	HostPort int32 `json:"host_port,omitempty" bson:"host_port,omitempty" jsonschema:"minimum=0,maximum=65536"`
	// Number of port to expose on the pod's IP address.
	// This must be a valid port number, 0 < x < 65536.
	ContainerPort int32 `json:"container_port,omitempty" bson:"container_port,omitempty" jsonschema:"minimum=0,maximum=65536"`
	// Protocol for port. Must be UDP, TCP, or SCTP.
	// Defaults to "TCP".
	// +optional
	Protocol Protocol `json:"protocol,omitempty" bson:"protocol,omitempty" jsonschema:"enum=TCP,enum=UDP,enum=SCTP,default=TCP" default:"TCP"`
	// What host IP to bind the external port to.
	// +optional
	HostIP string `json:"host_ip,omitempty" bson:"host_ip,omitempty"`
}

type Affinity struct {
	NodeAffinity    *NodeAffinity    `json:"node_Affinity,omitempty" bson:"node_Affinity,omitempty"`
	PodAffinity     *PodAffinity     `json:"pod_Affinity,omitempty" bson:"pod_Affinity,omitempty"`
	PodAntiAffinity *PodAntiAffinity `json:"pod_anti_affinity,omitempty" bson:"pod_anti_affinity,omitempty"`
}

type NodeAffinity struct {
	ReqDuringSchedulingIgnDuringExec *NodeSelector             `json:"req_during_scheduling_ign_during_exec,omitempty" bson:"req_during_scheduling_ign_during_exec,omitempty"`
	PrefDuringIgnDuringExec          []PreferredSchedulingTerm `json:"pref_during_ign_during_exec,omitempty" bson:"pref_during_ign_during_exec,omitempty"`
}

// A node selector represents the union of the results of one or more label queries
// over a set of nodes; that is, it represents the OR of the selectors represented
// by the node selector terms.
type NodeSelector struct {
	//Required. A list of node selector terms. The terms are ORed.
	NodeSelectorTerms []NodeSelectorTerm `json:"node_selector_terms,omitempty" bson:"node_selector_terms,omitempty"`
}

// A null or empty node selector term matches no objects. The requirements of
// them are ANDed.
// The TopologySelectorTerm type implements a subset of the NodeSelectorTerm.
type NodeSelectorTerm struct {
	// A list of node selector requirements by node's labels.
	// +optional
	MatchExpressions []NodeSelectorRequirement `json:"match_expressions,omitempty" bson:"match_expressions,omitempty"`
	// A list of node selector requirements by node's fields.
	// +optional
	MatchFields []NodeSelectorRequirement `json:"match_fields,omitempty" bson:"match_fields,omitempty"`
}

// A node selector requirement is a selector that contains values, a key, and an operator
// that relates the key and values.
type NodeSelectorRequirement struct {
	// The label key that the selector applies to.
	Key string `json:"key" bson:"key"`
	// Represents a key's relationship to a set of values.
	// Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.
	Operator NodeSelectorOperator `json:"operator" bson:"operator"`
	// An array of string values. If the operator is In or NotIn,
	// the values array must be non-empty. If the operator is Exists or DoesNotExist,
	// the values array must be empty. If the operator is Gt or Lt, the values
	// array must have a single element, which will be interpreted as an integer.
	// This array is replaced during a strategic merge patch.
	// +optional
	Values []string `json:"values,omitempty" bson:"values,omitempty"`
}

// A node selector operator is the set of operators that can be used in
// a node selector requirement.
type NodeSelectorOperator string

const (
	NodeSelectorOpIn            NodeSelectorOperator = "NodeSelectorOpIn"
	NodeSelectorOpNotIn         NodeSelectorOperator = "NodeSelectorOpNotIn"
	NodeSelectorOpExists        NodeSelectorOperator = "NodeSelectorOpExists"
	NodeSelectorOpDoesNotExists NodeSelectorOperator = "NodeSelectorOpDoesNotExist"
	NodeSelectorOpGt            NodeSelectorOperator = "NodeSelectorOpGt"
	NodeSelectorOpLt            NodeSelectorOperator = "NodeSelectorOpLt"
)

// An empty preferred scheduling term matches all objects with implicit weight 0
// (i.e. it's a no-op). A null preferred scheduling term matches no objects (i.e. is also a no-op).
type PreferredSchedulingTerm struct {
	// Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.
	Weight int32 `json:"weight" bson:"weight" jsonschema:"minimum=1,maximum=100"`
	// A node selector term, associated with the corresponding weight.
	Preference NodeSelectorTerm `json:"preference" bson:"preference"`
}

type PodAffinity struct {
	ReqDuringSchedulingIgnDuringExec []PodAffinityTerm         `json:"req_during_scheduling_ign_during_exec,omitempty" bson:"req_during_scheduling_ign_during_exec,omitempty"`
	PrefDuringIgnDuringExec          []WeightedPodAffinityTerm `json:"pref_during_ign_during_exec,omitempty" bson:"pref_during_ign_during_exec,omitempty"`
}

type PodAffinityTerm struct {
	LabelSelector *LabelSelectorObj `json:"label_selector,omitempty" bson:"label_selector,omitempty"`
	Namespaces    []string          `json:"namespaces,omitempty" bson:"namespaces,omitempty"`
	TopologyKey   string            `json:"topology_key,omitempty" bson:"topology_key,omitempty"`
}

type WeightedPodAffinityTerm struct {
	// weight associated with matching the corresponding podAffinityTerm,
	// in the range 1-100.
	Weight int32 `json:"weight" bson:"weight" jsonschema:"minimum=1,maximum=100"`
	// Required. A pod affinity term, associated with the corresponding weight.
	PodAffinityTerm PodAffinityTerm `json:"pod_affinity_term" bson:"pod_affinity_term"`
}

type PodAntiAffinity struct {
	ReqDuringSchedulingIgnDuringExec []PodAffinityTerm         `json:"req_during_scheduling_ign_during_exec,omitempty" bson:"req_during_scheduling_ign_during_exec,omitempty"`
	PrefDuringIgnDuringExec          []WeightedPodAffinityTerm `json:"pref_during_ign_during_exec,omitempty" bson:"pref_during_ign_during_exec,omitempty"`
}

type DeploymentStrategy struct {
	Type          DeploymentStrategyType   `json:"type,omitempty" bson:"type,omitempty"`
	RollingUpdate *RollingUpdateDeployment `json:"rolling_update,omitempty" bson:"rolling_update,omitempty"`
}

type DeploymentStrategyType string

const (
	// Kill all existing pods before creating new ones.
	RecreateDeploymentStrategyType DeploymentStrategyType = "Recreate"

	// Replace the old ReplicaSets by new one using rolling update i.e gradually scale down the old ReplicaSets and scale up the new one.
	RollingUpdateDeploymentStrategyType DeploymentStrategyType = "RollingUpdate"
)

type RollingUpdateDeployment struct {
	MaxUnavailable int32 `json:"max_unavailable,omitempty"`
	MaxSurge       int32 `json:"max_surge,omitempty"`
}

type DaemonSetUpdateStrategy struct {
	Type          DaemonSetUpdateStrategyType `json:"type,omitempty"`
	RollingUpdate *RollingUpdateDaemonSet     `json:"rolling_update,omitempty"`
}

type DaemonSetUpdateStrategyType string

const (
	RollingUpdateDaemonSetStrategyType DaemonSetUpdateStrategyType = "RollingUpdate"
	OnDeleteDaemonSetStrategyType      DaemonSetUpdateStrategyType = "OnDelete"
)

type RollingUpdateDaemonSet struct {
	MaxUnavailable *intstr.IntOrString `json:"maxUnavailable,omitempty"`
}

type StateFulSetUpdateStrategy struct {
	Type          StatefulSetUpdateStrategyType     `json:"type,omitempty"`
	RollingUpdate *RollingUpdateStatefulSetStrategy `json:"rolling_update,omitempty"`
}

type StatefulSetUpdateStrategyType string

const (
	RollingUpdateStatefulSetStrategyType StatefulSetUpdateStrategyType = "RollingUpdate"
	OnDeleteStatefulSetStrategyType      StatefulSetUpdateStrategyType = "OnDelete"
)

type RollingUpdateStatefulSetStrategy struct {
	Partition *int32 `json:"partition,omitempty" protobuf:"varint,1,opt,name=partition"`
}

type PodManagementPolicyType string

const (
	OrderedReadyPodManagement PodManagementPolicyType = "OrderedReady"
	ParallelPodManagement     PodManagementPolicyType = "Parallel"
)

type Replicas struct {
	Value int32 `json:"value,omitempty" bson:"value,omitempty"`
}

type TerminationGracePeriodSeconds struct {
	Value int32 `json:"value,omitempty" bson:"value,omitempty"`
}

type ActiveDeadlineSeconds struct {
	Value int64 `json:"value,omitempty" bson:"value,omitempty"`
}
