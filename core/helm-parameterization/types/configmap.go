package types

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type ConfigMapTemplate struct {
	metav1.TypeMeta    `json:",inline" `
	ObjectMetaTemplate `json:"metadata" yaml:"metadata"`
	Data               interface{} `json:"data,omitempty" protobuf:"bytes,2,rep,name=data"`
}
