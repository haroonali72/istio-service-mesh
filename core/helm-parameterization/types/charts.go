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
