package types

import (
	v1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type RoleTemplate struct {
	metav1.TypeMeta    `json:",inline"`
	ObjectMetaTemplate `json:"metadata" yaml:"metadata"`
	Rules              []v1.PolicyRule `json:"rules" yaml:"rules"`
}
