package types

import (
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DaemonSetTemplate struct {
	metav1.TypeMeta    `json:",inline" `
	ObjectMetaTemplate `json:"metadata" yaml:"metadata"`
	Spec               DaemonSetSpecTemplate `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

type DaemonSetSpecTemplate struct {
	Selector LabelSelectorTemplate `json:"selector" protobuf:"bytes,2,opt,name=selector"`

	Template PodSpecTemplate `json:"template" protobuf:"bytes,3,opt,name=template"`

	MinReadySeconds int32 `json:"minReadySeconds,omitempty" protobuf:"varint,5,opt,name=minReadySeconds"`

	RevisionHistoryLimit *int32 `json:"revisionHistoryLimit,omitempty" protobuf:"varint,6,opt,name=revisionHistoryLimit"`

	// An update strategy to replace existing DaemonSet pods with new pods.
	// +optional
	UpdateStrategy v1.DaemonSetUpdateStrategy `json:"updateStrategy,omitempty" protobuf:"bytes,3,opt,name=updateStrategy"`
}
