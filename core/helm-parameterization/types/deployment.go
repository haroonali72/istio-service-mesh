package types

import (
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DeploymentTemplate struct {
	metav1.TypeMeta    `json:",inline" `
	ObjectMetaTemplate `json:"metadata" yaml:"metadata"`
	Spec               DeploymentSpecTemplate `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

type DeploymentSpecTemplate struct {
	Replicas interface{} `json:"replicas,omitempty" protobuf:"varint,1,opt,name=replicas"`

	Selector LabelSelectorTemplate `json:"selector" protobuf:"bytes,2,opt,name=selector"`

	Template PodSpecTemplate `json:"template" protobuf:"bytes,3,opt,name=template"`

	Strategy v1.DeploymentStrategy `json:"strategy,omitempty" patchStrategy:"retainKeys" protobuf:"bytes,4,opt,name=strategy"`

	MinReadySeconds int32 `json:"minReadySeconds,omitempty" protobuf:"varint,5,opt,name=minReadySeconds"`

	RevisionHistoryLimit *int32 `json:"revisionHistoryLimit,omitempty" protobuf:"varint,6,opt,name=revisionHistoryLimit"`

	Paused bool `json:"paused,omitempty" protobuf:"varint,7,opt,name=paused"`

	ProgressDeadlineSeconds *int32 `json:"progressDeadlineSeconds,omitempty" protobuf:"varint,9,opt,name=progressDeadlineSeconds"`
}
