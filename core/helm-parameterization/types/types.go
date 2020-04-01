package types

type ObjectMetaTemplate struct {
	// Name is primarily intended for creation idempotence and configuration definition.
	Name        interface{} `json:"name,omitempty" yaml:"name,omitempty"`
	Namespace   string      `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	Labels      interface{} `json:"labels,omitempty" yaml:"labels,omitempty"`
	Annotations interface{} `json:"annotations,omitempty" yaml:"annotations,omitempty"`
}

type LabelSelectorTemplate struct {
	// matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels
	// map is equivalent to an element of matchExpressions, whose key field is "key", the
	// operator is "In", and the values array contains only "value". The requirements are ANDed.
	// +optional
	MatchLabels interface{} `json:"matchLabels,omitempty" protobuf:"bytes,1,rep,name=matchLabels"`
	// matchExpressions is a list of label selector requirements. The requirements are ANDed.
	// +optional
	MatchExpressions interface{} `json:"matchExpressions,omitempty" protobuf:"bytes,2,rep,name=matchExpressions"`
}
