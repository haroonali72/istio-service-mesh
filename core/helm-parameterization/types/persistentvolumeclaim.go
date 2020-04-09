package types

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type PersistentVolumeClaimTemplate struct {
	metav1.TypeMeta    `json:",inline" `
	ObjectMetaTemplate `json:"metadata" yaml:"metadata"`
	Spec               PersistentVolumeClaimTemplateSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}
type PersistentVolumeClaimTemplateSpec struct {
	// AccessModes contains the desired access modes the volume should have.
	// More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1
	// +optional
	AccessModes interface{} `json:"accessModes,omitempty" protobuf:"bytes,1,rep,name=accessModes,casttype=PersistentVolumeAccessMode"`
	// A label query over volumes to consider for binding.
	// +optional
	Selector *metav1.LabelSelector `json:"selector,omitempty" protobuf:"bytes,4,opt,name=selector"`
	// Resources represents the minimum resources the volume should have.
	// More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources
	// +optional
	Resources ResourceRequirements `json:"resources,omitempty" protobuf:"bytes,2,opt,name=resources"`
	// VolumeName is the binding reference to the PersistentVolume backing this claim.
	// +optional
	VolumeName string `json:"volumeName,omitempty" protobuf:"bytes,3,opt,name=volumeName"`
	// Name of the StorageClass required by the claim.
	// More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1
	// +optional
	StorageClassName string `json:"storageClassName,omitempty" protobuf:"bytes,5,opt,name=storageClassName"`
	// volumeMode defines what type of volume is required by the claim.
	// Value of Filesystem is implied when not included in claim spec.
	// This is a beta feature.
	// +optional
	VolumeMode string `json:"volumeMode,omitempty" protobuf:"bytes,6,opt,name=volumeMode,casttype=PersistentVolumeMode"`
}

// ResourceRequirements describes the compute resource requirements.
type ResourceRequirements struct {
	// Limits describes the maximum amount of compute resources allowed.
	// More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/
	// +optional
	Limits map[string]interface{} `json:"limits,omitempty" protobuf:"bytes,1,rep,name=limits,casttype=ResourceList,castkey=ResourceName"`
	// Requests describes the minimum amount of compute resources required.
	// If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,
	// otherwise to an implementation-defined value.
	// More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/
	// +optional
	Requests map[string]interface{} `json:"requests,omitempty" protobuf:"bytes,2,rep,name=requests,casttype=ResourceList,castkey=ResourceName"`
}
