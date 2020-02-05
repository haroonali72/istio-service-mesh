package types

import "k8s.io/apimachinery/pkg/util/intstr"

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

type ProcMountType string

const (
	DefaultProcMount  ProcMountType = "Default"
	UnmaskedProcMount ProcMountType = "Unmasked"
)

type Capability string

type Capabilities struct {
	Add  []Capability `json:"add"`
	Drop []Capability `json:"drop"`
}

const (
	ResourceTypeMemory string = "memory"
	ResourceTypeCpu    string = "cpu"
)

type Protocol string

const (
	ProtocolTCP  Protocol = "TCP"
	ProtocolUDP  Protocol = "UDP"
	ProtocolSCTP Protocol = "SCTP"
)

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
	VolumeMounts                  []VolumeMount                  `json:"volumeMounts,omitempty"`
}

type VolumeMount struct {
	Name             string                `json:"name"`
	ReadOnly         bool                  `json:"readOnly,omitempty"`
	MountPath        string                `json:"mountPath,omitempty"`
	SubPath          string                `json:"subPath,omitempty"`
	MountPropagation *MountPropagationMode `json:"mountPropagation,omitempty"`
	SubPathExpr      string                `json:"subPathExpr,omitempty"`
}

type MountPropagationMode string

const (
	MountPropagationNone            MountPropagationMode = "None"
	MountPropagationHostToContainer MountPropagationMode = "HostToContainer"
	MountPropagationBidirectional   MountPropagationMode = "Bidirectional"
)

type ContainerPort struct {
	// Number of port to expose on the host.
	// If specified, this must be a valid port number, 0 < x < 65536.
	// If HostNetwork is specified, this must match ContainerPort.
	// Most containers do not need this.
	// +optional
	HostPort int32 `json:"hostPort,omitempty"`
	// Number of port to expose on the pod's IP address.
	// This must be a valid port number, 0 < x < 65536.
	ContainerPort int32 `json:"containerPort"`
	// Protocol for port. Must be UDP, TCP, or SCTP.
	// Defaults to "TCP".
	// +optional
	Protocol Protocol `json:"protocol,omitempty"`
	// What host IP to bind the external port to.
	// +optional
	HostIP string `json:"hostIP,omitempty"`
}

type Affinity struct {
	NodeAffinity    *NodeAffinity    `json:"nodeAffinity,omitempty"`
	PodAffinity     *PodAffinity     `json:"podAffinity,omitempty"`
	PodAntiAffinity *PodAntiAffinity `json:"podAntiAffinity,omitempty"`
}

type NodeAffinity struct {
	ReqDuringSchedulingIgnDuringExec *NodeSelector             `json:"reqDuringSchedulingIgnDuringExec,omitempty"`
	PrefDuringIgnDuringExec          []PreferredSchedulingTerm `json:"prefDuringIgnDuringExec,omitempty"`
}

// A node selector represents the union of the results of one or more label queries
// over a set of nodes; that is, it represents the OR of the selectors represented
// by the node selector terms.
type NodeSelector struct {
	//Required. A list of node selector terms. The terms are ORed.
	NodeSelectorTerms []NodeSelectorTerm `json:"nodeSelectorTerms"`
}

// A null or empty node selector term matches no objects. The requirements of
// them are ANDed.
// The TopologySelectorTerm type implements a subset of the NodeSelectorTerm.
type NodeSelectorTerm struct {
	// A list of node selector requirements by node's labels.
	// +optional
	MatchExpressions []NodeSelectorRequirement `json:"matchExpressions,omitempty"`
	// A list of node selector requirements by node's fields.
	// +optional
	MatchFields []NodeSelectorRequirement `json:"matchFields,omitempty"`
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
	ReqDuringSchedulingIgnDuringExec []PodAffinityTerm         `json:"reqDuringSchedulingIgnDuringExec,omitempty"`
	PrefDuringIgnDuringExec          []WeightedPodAffinityTerm `json:"prefDuringIgnDuringExec,omitempty"`
}

type PodAffinityTerm struct {
	LabelSelector *LabelSelectorObj `json:"labelSelector,omitempty"`
	Namespaces    []string          `json:"namespaces,omitempty"`
	TopologyKey   string            `json:"topologyKey,omitempty"`
}

type WeightedPodAffinityTerm struct {
	// weight associated with matching the corresponding podAffinityTerm,
	// in the range 1-100.
	Weight int32 `json:"weight"`
	// Required. A pod affinity term, associated with the corresponding weight.
	PodAffinityTerm PodAffinityTerm `json:"podAffinityTerm"`
}

type PodAntiAffinity struct {
	ReqDuringSchedulingIgnDuringExec []PodAffinityTerm         `json:"reqDuringSchedulingIgnDuringExec,omitempty"`
	PrefDuringIgnDuringExec          []WeightedPodAffinityTerm `json:"prefDuringIgnDuringExec,omitempty"`
}

type DeploymentStrategy struct {
	Type          DeploymentStrategyType   `json:"type,omitempty"`
	RollingUpdate *RollingUpdateDeployment `json:"rollingUpdate,omitempty"`
}

type DeploymentStrategyType string

const (
	// Kill all existing pods before creating new ones.
	RecreateDeploymentStrategyType DeploymentStrategyType = "Recreate"

	// Replace the old ReplicaSets by new one using rolling update i.e gradually scale down the old ReplicaSets and scale up the new one.
	RollingUpdateDeploymentStrategyType DeploymentStrategyType = "RollingUpdate"
)

type RollingUpdateDeployment struct {
	MaxUnavailable *intstr.IntOrString `json:"maxUnavailable,omitempty"`
	MaxSurge       *intstr.IntOrString `json:"maxSurge,omitempty"`
}

type DaemonSetUpdateStrategy struct {
	Type          DaemonSetUpdateStrategyType `json:"type,omitempty"`
	RollingUpdate *RollingUpdateDaemonSet     `json:"rollingUpdate,omitempty"`
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
	RollingUpdate *RollingUpdateStatefulSetStrategy `json:"rollingUpdate,omitempty"`
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
	Value int64 `json:"value, omitempty"`
}

type ActiveDeadlineSeconds struct {
	Value int64 `json:"value, omitempty"`
}

type RevisionHistoryLimit struct {
	Value int32 `json:"value,omitempty"`
}

type RestartPolicy string

const (
	RestartPolicyAlways    RestartPolicy = "Always"
	RestartPolicyOnFailure RestartPolicy = "OnFailure"
	RestartPolicyNever     RestartPolicy = "Never"
)
