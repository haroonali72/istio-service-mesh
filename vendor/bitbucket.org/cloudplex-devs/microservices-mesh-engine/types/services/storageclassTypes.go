package services

import "time"

type StorageClassService struct {
	Id                interface{}                   `json:"_id,omitempty" bson:"_id" valid:"-"`
	ServiceId         string                        `json:"service_id" bson:"service_id" binding:"required" valid:"alphanumspecial,length(4|30)~service_id must contain between 6 and 30 characters,lowercase~lowercase alphanumeric characters are allowed,required~service_id is missing in request"`
	Name              string                        `json:"name"  bson:"name" binding:"required" valid:"alphanumspecial,length(3|30),lowercase~lowercase alphanumeric characters are allowed,required"`
	Version           string                        `json:"version"  bson:"version"  binding:"required" valid:"alphanumspecial,length(1|10),lowercase~lowercase alphanumeric characters are allowed,required"`
	ServiceType       string                        `json:"service_type"  bson:"service_type" valid:"-"`
	ServiceSubType    string                        `json:"service_sub_type" bson:"service_type" valid:"-"`
	CompanyId         string                        `json:"company_id,omitempty" bson:"company_id"`
	CreationDate      time.Time                     `json:"creation_date,omitempty" bson:"creation_date" valid:"-"`
	ServiceAttributes *StorageClassServiceAttribute `json:"service_attributes"  bson:"company_id" binding:"required"`
}
type StorageClassServiceAttribute struct {
	BindingMod           VolumeBindingMode      `json:"volume_binding_mode,omitempty"`
	AllowVolumeExpansion string                 `json:"allow_volume_expansion,omitempty"`
	Provisioner          string                 `json:"provisioner,omitempty"`
	Parameters           map[string]string      `json:"parameters,omitempty"`
	ReclaimPolicy        ReclaimPolicy          `json:"reclaim_policy,omitempty"`
	MountOptions         []string               `json:"mount_options,omitempty"`
	AllowedTopologies    []TopologySelectorTerm `json:"allowed_topologies,omitempty"`
}

type TopologySelectorTerm struct {
	MatchLabelExpressions []TopologySelectorLabelRequirement `json:"match_label_expressions,omitempty"`
}
type TopologySelectorLabelRequirement struct {
	Key    string   `json:"key,omitempty"`
	Values []string `json:"values,omitempty"`
}
type VolumeBindingMode string

const (
	VolumeBindingModeImmediate            VolumeBindingMode = "Immediate"
	VolumeBindingModeWaitForFirstConsumer VolumeBindingMode = "WaitForFirstConsumer"
)

type ReclaimPolicy string

const (
	ReclaimPolicyRetain ReclaimPolicy = "Retain"
	ReclaimPolicyDelete ReclaimPolicy = "Delete"
)
