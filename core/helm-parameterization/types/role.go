package types

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type RoleTemplate struct {
	metav1.TypeMeta    `json:",inline"`
	ObjectMetaTemplate `json:"metadata" yaml:"metadata"`
	Rules              interface{} `json:"rules"`
}
