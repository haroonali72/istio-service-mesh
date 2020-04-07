package types

type PodSpecTemplate struct {
	ObjectMetaTemplate `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec PodTemplate `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}
type PodTemplate struct {
	Volumes interface{} `json:"volumes,omitempty" patchStrategy:"merge,retainKeys" patchMergeKey:"name" protobuf:"bytes,1,rep,name=volumes"`

	InitContainers interface{} `json:"initContainers,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,20,rep,name=initContainers"`

	Containers []ContainerTemplate `json:"containers" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,2,rep,name=containers"`

	EphemeralContainers interface{} `json:"ephemeralContainers,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,34,rep,name=ephemeralContainers"`

	RestartPolicy interface{} `json:"restartPolicy,omitempty" protobuf:"bytes,3,opt,name=restartPolicy,casttype=RestartPolicy"`

	TerminationGracePeriodSeconds *int64 `json:"terminationGracePeriodSeconds,omitempty" protobuf:"varint,4,opt,name=terminationGracePeriodSeconds"`

	ActiveDeadlineSeconds *int64 `json:"activeDeadlineSeconds,omitempty" protobuf:"varint,5,opt,name=activeDeadlineSeconds"`

	DNSPolicy interface{} `json:"dnsPolicy,omitempty" protobuf:"bytes,6,opt,name=dnsPolicy,casttype=DNSPolicy"`

	NodeSelector map[string]string `json:"nodeSelector,omitempty" protobuf:"bytes,7,rep,name=nodeSelector"`

	ServiceAccountName string `json:"serviceAccountName,omitempty" protobuf:"bytes,8,opt,name=serviceAccountName"`

	DeprecatedServiceAccount string `json:"serviceAccount,omitempty" protobuf:"bytes,9,opt,name=serviceAccount"`

	AutomountServiceAccountToken *bool `json:"automountServiceAccountToken,omitempty" protobuf:"varint,21,opt,name=automountServiceAccountToken"`

	NodeName string `json:"nodeName,omitempty" protobuf:"bytes,10,opt,name=nodeName"`

	HostNetwork bool `json:"hostNetwork,omitempty" protobuf:"varint,11,opt,name=hostNetwork"`

	HostPID bool `json:"hostPID,omitempty" protobuf:"varint,12,opt,name=hostPID"`

	HostIPC bool `json:"hostIPC,omitempty" protobuf:"varint,13,opt,name=hostIPC"`

	ShareProcessNamespace *bool `json:"shareProcessNamespace,omitempty" protobuf:"varint,27,opt,name=shareProcessNamespace"`

	SecurityContext string `json:"securityContext,omitempty" protobuf:"bytes,14,opt,name=securityContext"`

	ImagePullSecrets interface{} `json:"imagePullSecrets,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,15,rep,name=imagePullSecrets"`

	Hostname string `json:"hostname,omitempty" protobuf:"bytes,16,opt,name=hostname"`

	Subdomain string `json:"subdomain,omitempty" protobuf:"bytes,17,opt,name=subdomain"`

	Affinity *Affinity `json:"affinity,omitempty" protobuf:"bytes,18,opt,name=affinity"`

	SchedulerName string `json:"schedulerName,omitempty" protobuf:"bytes,19,opt,name=schedulerName"`

	Tolerations interface{} `json:"tolerations,omitempty" protobuf:"bytes,22,opt,name=tolerations"`

	HostAliases       interface{} `json:"hostAliases,omitempty" patchStrategy:"merge" patchMergeKey:"ip" protobuf:"bytes,23,rep,name=hostAliases"`
	PriorityClassName string      `json:"priorityClassName,omitempty" protobuf:"bytes,24,opt,name=priorityClassName"`

	Priority       *int32      `json:"priority,omitempty" protobuf:"bytes,25,opt,name=priority"`
	DNSConfig      interface{} `json:"dnsConfig,omitempty" protobuf:"bytes,26,opt,name=dnsConfig"`
	ReadinessGates interface{} `json:"readinessGates,omitempty" protobuf:"bytes,28,opt,name=readinessGates"`

	RuntimeClassName   *string     `json:"runtimeClassName,omitempty" protobuf:"bytes,29,opt,name=runtimeClassName"`
	EnableServiceLinks *bool       `json:"enableServiceLinks,omitempty" protobuf:"varint,30,opt,name=enableServiceLinks"`
	PreemptionPolicy   interface{} `json:"preemptionPolicy,omitempty" protobuf:"bytes,31,opt,name=preemptionPolicy"`
	Overhead           interface{} `json:"overhead,omitempty" protobuf:"bytes,32,opt,name=overhead"`

	TopologySpreadConstraints interface{} `json:"topologySpreadConstraints,omitempty" patchStrategy:"merge" patchMergeKey:"topologyKey" protobuf:"bytes,33,opt,name=topologySpreadConstraints"`
}

// A single application container that you want to run within a pod.
type ContainerTemplate struct {
	Name string `json:"name" protobuf:"bytes,1,opt,name=name"`

	Image string `json:"image,omitempty" protobuf:"bytes,2,opt,name=image"`

	Command interface{} `json:"command,omitempty" protobuf:"bytes,3,rep,name=command"`

	Args interface{} `json:"args,omitempty" protobuf:"bytes,4,rep,name=args"`

	WorkingDir string `json:"workingDir,omitempty" protobuf:"bytes,5,opt,name=workingDir"`

	Ports interface{} `json:"ports,omitempty" patchStrategy:"merge" patchMergeKey:"containerPort" protobuf:"bytes,6,rep,name=ports"`

	EnvFrom interface{} `json:"envFrom,omitempty" protobuf:"bytes,19,rep,name=envFrom"`

	Env interface{} `json:"env,omitempty" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,7,rep,name=env"`

	Resources interface{} `json:"resources,omitempty" protobuf:"bytes,8,opt,name=resources"`

	VolumeMounts interface{} `json:"volumeMounts,omitempty" patchStrategy:"merge" patchMergeKey:"mountPath" protobuf:"bytes,9,rep,name=volumeMounts"`

	VolumeDevices interface{} `json:"volumeDevices,omitempty" patchStrategy:"merge" patchMergeKey:"devicePath" protobuf:"bytes,21,rep,name=volumeDevices"`

	LivenessProbe interface{} `json:"livenessProbe,omitempty" protobuf:"bytes,10,opt,name=livenessProbe"`

	ReadinessProbe interface{} `json:"readinessProbe,omitempty" protobuf:"bytes,11,opt,name=readinessProbe"`

	StartupProbe interface{} `json:"startupProbe,omitempty" protobuf:"bytes,22,opt,name=startupProbe"`

	Lifecycle interface{} `json:"lifecycle,omitempty" protobuf:"bytes,12,opt,name=lifecycle"`

	TerminationMessagePath interface{} `json:"terminationMessagePath,omitempty" protobuf:"bytes,13,opt,name=terminationMessagePath"`

	TerminationMessagePolicy interface{} `json:"terminationMessagePolicy,omitempty" protobuf:"bytes,20,opt,name=terminationMessagePolicy,casttype=TerminationMessagePolicy"`

	ImagePullPolicy interface{} `json:"imagePullPolicy,omitempty" protobuf:"bytes,14,opt,name=imagePullPolicy,casttype=PullPolicy"`

	SecurityContext interface{} `json:"securityContext,omitempty" protobuf:"bytes,15,opt,name=securityContext"`

	Stdin interface{} `json:"stdin,omitempty" protobuf:"varint,16,opt,name=stdin"`

	StdinOnce interface{} `json:"stdinOnce,omitempty" protobuf:"varint,17,opt,name=stdinOnce"`

	TTY interface{} `json:"tty,omitempty" protobuf:"varint,18,opt,name=tty"`
}

// Affinity is a group of affinity scheduling rules.
type Affinity struct {
	// Describes node affinity scheduling rules for the pod.
	// +optional
	NodeAffinity interface{} `json:"nodeAffinity,omitempty" protobuf:"bytes,1,opt,name=nodeAffinity"`
	// Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).
	// +optional
	PodAffinity interface{} `json:"podAffinity,omitempty" protobuf:"bytes,2,opt,name=podAffinity"`
	// Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).
	// +optional
	PodAntiAffinity interface{} `json:"podAntiAffinity,omitempty" protobuf:"bytes,3,opt,name=podAntiAffinity"`
}
