package services

import (
	"bitbucket.org/cloudplex-devs/microservices-mesh-engine/types"
)

//Id                interface{}                   `json:"_id,omitempty" bson:"_id" valid:"-"`
//	ServiceId         string                        `json:"service_id" bson:"service_id" binding:"required" valid:"alphanumspecial,length(4|30)~service_id must contain between 6 and 30 characters,lowercase~lowercase alphanumeric characters are allowed,required~service_id is missing in request"`
//	Name              string                        `json:"name"  bson:"name" binding:"required" valid:"alphanumspecial,length(3|30),lowercase~lowercase alphanumeric characters are allowed,required"`
//	Version           string                        `json:"version"  bson:"version"  binding:"required" valid:"alphanumspecial,length(1|10),lowercase~lowercase alphanumeric characters are allowed,required"`
//	ServiceType       constants.ServiceType         `json:"service_type"  bson:"service_type" valid:"-"`
//	ServiceSubType    constants.ServiceSubType      `json:"service_sub_type" bson:"service_type" valid:"-"`
//	CompanyId         string                        `json:"company_id,omitempty" bson:"company_id"`
//	CreationDate      time.Time                     `json:"creation_date,omitempty" bson:"creation_date" valid:"-"`
type StorageClassService struct {
	types.ServiceBasicInfo `json:",inline" bson:",inline"`
	ServiceAttributes      *StorageClassServiceAttribute `json:"service_attributes"  bson:"service_attributes" binding:"required"`
}
type StorageClassServiceAttribute struct {
	BindingMod           VolumeBindingMode      `json:"volume_binding_mode,omitempty" bson:"volume_binding_mode,omitempty"`
	AllowVolumeExpansion string                 `json:"allow_volume_expansion,omitempty" bson:"allow_volume_expansion,omitempty"`
	Provisioner          string                 `json:"provisioner,omitempty" bson:"provisioner,omitempty"`
	Parameters           map[string]string      `json:"parameters,omitempty" bson:"parameters,omitempty"`
	ReclaimPolicy        ReclaimPolicy          `json:"reclaim_policy,omitempty" bson:"reclaim_policy,omitempty"`
	MountOptions         []string               `json:"mount_options,omitempty" bson:"mount_options,omitempty"`
	AllowedTopologies    []TopologySelectorTerm `json:"allowed_topologies,omitempty" bson:"allowed_topologies,omitempty"`
}

type TopologySelectorTerm struct {
	MatchLabelExpressions []TopologySelectorLabelRequirement `json:"match_label_expressions,omitempty" bson:"match_label_expressions,omitempty"`
}
type TopologySelectorLabelRequirement struct {
	Key    string   `json:"key,omitempty" bson:"key,omitempty"`
	Values []string `json:"values,omitempty" bson:"values,omitempty"`
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
