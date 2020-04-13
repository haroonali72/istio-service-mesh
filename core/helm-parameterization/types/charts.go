package types

type CoreComponentsChartValues struct {
	Replicas      int32       `yaml:"replicas,omitempty" json:"replicas,omitempty"`
	ResourceQuota interface{} `json:"resources,omitempty" yaml:"resources,omitempty"`
	ImageInfo     `json:"image,omitempty" yaml:"image,omitempty"`
	Probe         `json:"prob,omitempty" yaml:"prob,omitempty"`
	Ports         interface{} `json:"ports,omitempty" yaml:"ports,omitempty"`

	CronExpression string `json:"cronExpression,omitempty"  yaml:"cronExpression,omitempty"`
}

type ImageInfo struct {
	ImagePullPolicy string      `json:"imagePullPolicy,omitempty" yaml:"imagePullPolicy,omitempty"`
	ImagePullSecret interface{} `json:"imagePullSecrets,omitempty" yaml:"imagePullSecrets,omitempty"`
	Image           string      `json:"image,omitempty" yaml:"image,omitempty"`
}
type Probe struct {
	LivenessProb   interface{} `json:"livenessProbe,omitempty" yaml:"livenessProbe,omitempty"`
	ReadinessProbe interface{} `json:"readinessProbe,omitempty" yaml:"readinessProbe,omitempty"`
}

type ServiceAccountChart struct {
	ServiceAccount `json:"serviceAccount,omitempty" yaml:"serviceAccount,omitempty"`
}

type ServiceAccount struct {
	Create bool   `json:"create,omitempty" yaml:"create,omitempty"`
	Name   string `json:"name,omitempty" yaml:"name,omitempty"`
}

type HPAChartValues struct {
	AutoScalingInfo `json:"autoscaling,omitempty" yaml:"autoscaling,omitempty"`
}

type AutoScalingInfo struct {
	MinReplicas                    int32 `json:"minReplicas,omitempty" yaml:"minReplicas,omitempty"`
	MaxReplicas                    int32 `json:"maxReplicas,omitempty" yaml:"maxReplicas,omitempty"`
	TargetCPUUtilizationPercentage int32 `json:"targetCPUUtilizationPercentage,omitempty" yaml:"targetCPUUtilizationPercentage,omitempty"`
	Enabled                        bool  `json:"enabled" yaml:"enabled"`
}

type RBACChartValues struct {
	RBACInfo `json:"rbac,omitempty" yaml:"rbac,omitempty"`
}

type RBACInfo struct {
	Create bool `json:"create,omitempty" yaml:"create,omitempty"`
}
