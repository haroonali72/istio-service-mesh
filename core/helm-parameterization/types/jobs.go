package types

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type JobTemplate struct {
	metav1.TypeMeta    `json:",inline" `
	ObjectMetaTemplate `json:"metadata" yaml:"metadata"`
	Spec               JobSpecTemplate `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

type JobSpecTemplate struct {
	Parallelism *int32 `json:"parallelism,omitempty" protobuf:"varint,1,opt,name=parallelism"`

	Completions *int32 `json:"completions,omitempty" protobuf:"varint,2,opt,name=completions"`

	ActiveDeadlineSeconds *int64 `json:"activeDeadlineSeconds,omitempty" protobuf:"varint,3,opt,name=activeDeadlineSeconds"`

	BackoffLimit *int32 `json:"backoffLimit,omitempty" protobuf:"varint,7,opt,name=backoffLimit"`

	Selector LabelSelectorTemplate `json:"selector" protobuf:"bytes,2,opt,name=selector"`

	Template PodSpecTemplate `json:"template" protobuf:"bytes,3,opt,name=template"`

	ManualSelector *bool `json:"manualSelector,omitempty" protobuf:"varint,5,opt,name=manualSelector"`

	TTLSecondsAfterFinished *int32 `json:"ttlSecondsAfterFinished,omitempty" protobuf:"varint,8,opt,name=ttlSecondsAfterFinished"`
}
