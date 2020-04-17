package services

import "k8s.io/apimachinery/pkg/util/intstr"

type ImageRepositoryConfigurations struct {
	Url         string               `json:"url,omitempty"`
	Tag         string               `json:"tag,omitempty"`
	Credentials BasicAuthCredentials `json:"credentials,omitempty"`
	Profile     string               `json:"profile_id,omitempty"`
}

type BasicAuthCredentials struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}
type EnvironmentVariable struct {
	Key     string `json:"key"`
	Value   string `json:"value"`
	Dynamic bool   `json:"dynamic"`
	Type    string `json:"type"`
}
type IstioConfig struct {
	Enable_External_Traffic bool `json:"enable_external_traffic"`
}

type LabelSelectorObj struct {
	MatchLabels      map[string]string          `json:"match_labels,omitempty"`
	MatchExpressions []LabelSelectorRequirement `json:"match_expressions,omitempty"`
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

type SecurityContextStruct struct {
	Capabilities             []Capabilities       `json:"capabilities"`
	RunAsUser                *int64               `json:"run_as_user"`
	RunAsGroup               *int64               `json:"run_as_group"`
	RunAsNonRoot             bool                 `json:"run_as_non_root"`
	Privileged               bool                 `json:"privileged"`
	ProcMount                ProcMountType        `json:"proc_mount"`
	AllowPrivilegeEscalation bool                 `json:"allow_privilege_escalation"`
	ReadOnlyRootFileSystem   bool                 `json:"read_only_root_filesystem"`
	SELinuxOptions           SELinuxOptionsStruct `json:"se_linux_options"`
}

type ProcMountType string

const (
	DefaultProcMount  ProcMountType = "Default"
	UnmaskedProcMount ProcMountType = "Unmasked"
)

type SELinuxOptionsStruct struct {
	User  string `json:"user,omitempty"`
	Role  string `json:"role,omitempty"`
	Type  string `json:"type,omitempty"`
	Level string `json:"level,omitempty"`
}

type Capability string

type Capabilities struct {
	Add  []Capability `json:"add"`
	Drop []Capability `json:"drop"`
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
	HTTPHeaders []HTTPHeader `json:"http_headers,omitempty" protobuf:"bytes,5,rep,name=http_headers"`
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
	HTTPGet *HTTPGetAction `json:"http_get,omitempty" protobuf:"bytes,2,opt,name=http_get"`
	// TCPSocket specifies an action involving a TCP port.
	// TCP hooks not yet supported
	// +optional
	TCPSocket *TCPSocketAction `json:"tcp_socket,omitempty" protobuf:"bytes,3,opt,name=tcp_socket"`
}

type Probe struct {
	// The action taken to determine the health of a container
	Handler *Handler `json:"handler,inline" protobuf:"bytes,1,opt,name=handler"`
	// Number of seconds after the container has started before liveness probes are initiated.
	// More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	// +optional
	InitialDelaySeconds *int32 `json:"initial_delay_seconds,omitempty" protobuf:"varint,2,opt,name=initial_delay_seconds"`
	// Number of seconds after which the probe times out.
	// Defaults to 1 second. Minimum value is 1.
	// More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	// +optional
	TimeoutSeconds *int32 `json:"timeout_seconds,omitempty" protobuf:"varint,3,opt,name=timeout_seconds"`
	// How often (in seconds) to perform the probe.
	// Default to 10 seconds. Minimum value is 1.
	// +optional
	PeriodSeconds *int32 `json:"period_seconds,omitempty" protobuf:"varint,4,opt,name=period_seconds"`
	// Minimum consecutive successes for the probe to be considered successful after having failed.
	// Defaults to 1. Must be 1 for liveness. Minimum value is 1.
	// +optional
	SuccessThreshold *int32 `json:"success_threshold,omitempty" protobuf:"varint,5,opt,name=success_threshold"`
	// Minimum consecutive failures for the probe to be considered failed after having succeeded.
	// Defaults to 3. Minimum value is 1.
	// +optional
	FailureThreshold *int32 `json:"failure_threshold,omitempty" protobuf:"varint,6,opt,name=failure_threshold"`
}

type ContainerAttribute struct {
	EnvironmentVariables          map[string]EnvironmentVariable `json:"environment_variables,omitempty"`
	ImageRepositoryConfigurations *ImageRepositoryConfigurations `json:"image_repository_configurations,omitempty" binding:"required"`
	Ports                         map[string]ContainerPort       `json:"ports,omitempty"`
	Tag                           string                         `json:"tag"`
	ImagePrefix                   string                         `json:"image_prefix"`
	ImageName                     string                         `json:"image_name,omitempty"`
	Command                       []string                       `json:"command,omitempty"`
	Args                          []string                       `json:"args,omitempty"`
	LimitResources                map[string]string              `json:"limit_resources,omitempty"`
	RequestResources              map[string]string              `json:"request_resources,omitempty"`
	LivenessProbe                 *Probe                         `json:"liveness_probe,omitempty"`
	ReadinessProbe                *Probe                         `json:"readiness_probe,omitempty"`
	SecurityContext               *SecurityContextStruct         `json:"security_context,omitempty"`
	VolumeMounts                  []VolumeMount                  `json:"volume_mounts,omitempty"`
}

type VolumeMount struct {
	Name             string                `json:"name"`
	ReadOnly         bool                  `json:"readonly,omitempty"`
	MountPath        string                `json:"mount_path,omitempty"`
	SubPath          string                `json:"sub_path,omitempty"`
	MountPropagation *MountPropagationMode `json:"mount_propagation,omitempty"`
	SubPathExpr      string                `json:"sub_path_expr,omitempty"`
	PvcSvcName       string                `json:"persistent_volume_claim_name,omitempty" bson:"persistent_volume_claim_name"`
}

type MountPropagationMode string

const (
	MountPropagationNone            MountPropagationMode = "None"
	MountPropagationHostToContainer MountPropagationMode = "HostToContainer"
	MountPropagationBidirectional   MountPropagationMode = "Bidirectional"
)

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

type ContainerPort struct {
	// Number of port to expose on the host.
	// If specified, this must be a valid port number, 0 < x < 65536.
	// If HostNetwork is specified, this must match ContainerPort.
	// Most containers do not need this.
	// +optional
	HostPort int32 `json:"host_port,omitempty"`
	// Number of port to expose on the pod's IP address.
	// This must be a valid port number, 0 < x < 65536.
	ContainerPort int32 `json:"container_port"`
	// Protocol for port. Must be UDP, TCP, or SCTP.
	// Defaults to "TCP".
	// +optional
	Protocol Protocol `json:"protocol,omitempty"`
	// What host IP to bind the external port to.
	// +optional
	HostIP string `json:"host_ip,omitempty"`
}

type Affinity struct {
	NodeAffinity    *NodeAffinity    `json:"node_Affinity,omitempty"`
	PodAffinity     *PodAffinity     `json:"pod_Affinity,omitempty"`
	PodAntiAffinity *PodAntiAffinity `json:"pod_anti_affinity,omitempty"`
}

type NodeAffinity struct {
	ReqDuringSchedulingIgnDuringExec *NodeSelector             `json:"req_during_scheduling_ign_during_exec,omitempty"`
	PrefDuringIgnDuringExec          []PreferredSchedulingTerm `json:"pref_during_ign_during_exec,omitempty"`
}

// A node selector represents the union of the results of one or more label queries
// over a set of nodes; that is, it represents the OR of the selectors represented
// by the node selector terms.
type NodeSelector struct {
	//Required. A list of node selector terms. The terms are ORed.
	NodeSelectorTerms []NodeSelectorTerm `json:"node_selector_terms"`
}

// A null or empty node selector term matches no objects. The requirements of
// them are ANDed.
// The TopologySelectorTerm type implements a subset of the NodeSelectorTerm.
type NodeSelectorTerm struct {
	// A list of node selector requirements by node's labels.
	// +optional
	MatchExpressions []NodeSelectorRequirement `json:"match_expressions,omitempty"`
	// A list of node selector requirements by node's fields.
	// +optional
	MatchFields []NodeSelectorRequirement `json:"match_fields,omitempty"`
}

// A node selector requirement is a selector that contains values, a key, and an operator
// that relates the key and values.
type NodeSelectorRequirement struct {
	// The label key that the selector applies to.
	Key string `json:"key"`
	// Represents a key's relationship to a set of values.
	// Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.
	Operator NodeSelectorOperator `json:"operator"`
	// An array of string values. If the operator is In or NotIn,
	// the values array must be non-empty. If the operator is Exists or DoesNotExist,
	// the values array must be empty. If the operator is Gt or Lt, the values
	// array must have a single element, which will be interpreted as an integer.
	// This array is replaced during a strategic merge patch.
	// +optional
	Values []string `json:"values,omitempty"`
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
	Weight int32 `json:"weight"`
	// A node selector term, associated with the corresponding weight.
	Preference NodeSelectorTerm `json:"preference"`
}

type PodAffinity struct {
	ReqDuringSchedulingIgnDuringExec []PodAffinityTerm         `json:"req_during_scheduling_ign_during_exec,omitempty"`
	PrefDuringIgnDuringExec          []WeightedPodAffinityTerm `json:"pref_during_ign_during_exec,omitempty"`
}

type PodAffinityTerm struct {
	LabelSelector *LabelSelectorObj `json:"label_selector,omitempty"`
	Namespaces    []string          `json:"namespaces,omitempty"`
	TopologyKey   string            `json:"topology_key,omitempty"`
}

type WeightedPodAffinityTerm struct {
	// weight associated with matching the corresponding podAffinityTerm,
	// in the range 1-100.
	Weight int32 `json:"weight"`
	// Required. A pod affinity term, associated with the corresponding weight.
	PodAffinityTerm PodAffinityTerm `json:"pod_affinity_term"`
}

type PodAntiAffinity struct {
	ReqDuringSchedulingIgnDuringExec []PodAffinityTerm         `json:"req_during_scheduling_ign_during_exec,omitempty"`
	PrefDuringIgnDuringExec          []WeightedPodAffinityTerm `json:"pref_during_ign_during_exec,omitempty"`
}

type DeploymentStrategy struct {
	Type          DeploymentStrategyType   `json:"type,omitempty"`
	RollingUpdate *RollingUpdateDeployment `json:"rolling_update,omitempty"`
}

type DeploymentStrategyType string

const (
	// Kill all existing pods before creating new ones.
	RecreateDeploymentStrategyType DeploymentStrategyType = "Recreate"

	// Replace the old ReplicaSets by new one using rolling update i.e gradually scale down the old ReplicaSets and scale up the new one.
	RollingUpdateDeploymentStrategyType DeploymentStrategyType = "RollingUpdate"
)

type RollingUpdateDeployment struct {
	MaxUnavailable *intstr.IntOrString `json:"max_unavailable,omitempty"`
	MaxSurge       *intstr.IntOrString `json:"max_surge,omitempty"`
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
	Value int32 `json:"value,omitempty"`
}

type TerminationGracePeriodSeconds struct {
	Value int32 `json:"value,omitempty"`
}

type ActiveDeadlineSeconds struct {
	Value int64 `json:"value,omitempty"`
}
