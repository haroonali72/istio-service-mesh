package types

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type HPATemplate struct {
	metav1.TypeMeta    `json:",inline" `
	ObjectMetaTemplate `json:"metadata" yaml:"metadata"`
	Spec               HPASpecTemplate `json:"spec,omitempty"`
}

type HPASpecTemplate struct {
	MinReplicas                    interface{}                         `json:"minReplicas,omitempty"`
	MaxReplicas                    interface{}                         `json:"maxReplicas,omitempty"`
	TargetCPUUtilizationPercentage int32                               `json:"targetCPUUtilizationPercentage,omitempty"`
	ScaleTargetRef                 CrossVersionObjectReferenceTemplate `json:"scaleTargetRef"`
}

type CrossVersionObjectReferenceTemplate struct {
	Kind       string      `json:"kind"`
	Name       interface{} `json:"name"`
	APIVersion string      `json:"apiVersion"`
}
