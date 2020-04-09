package types

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type RoleBindingTemplate struct {
	metav1.TypeMeta    `json:",inline"`
	ObjectMetaTemplate `json:"metdata" yaml:"metdata"`
	Subjects           []Subject `json:"subjects" yaml:"subjects"`
	RoleRef            RoleRef   `json:"roleRef" yaml:"roleRef"`
}

type Subject struct {
	Kind      string      `json:"kind" yaml="kind"`
	APIGroup  string      `json:"apiGroup" yaml:"apiGroup"`
	Name      interface{} `json:"name" yaml:"name"`
	Namespace interface{} `json:"namespace" yaml:"namespace"`
}

type RoleRef struct {
	APIGroup string      `json:"apiGroup" yaml:"apiGroup"`
	Kind     string      `json:"kind" yaml:"kind"`
	Name     interface{} `json:"name" yaml:"name"`
}
