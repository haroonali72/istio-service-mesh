package types

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
	BindingMod           VolumeBindingMode      `json:"volumeBindingMode,omitempty"`
	AllowVolumeExpansion string                 `json:"allowVolumeExpansion,omitempty"`
	Provisioner          string                 `json:"provisioner,omitempty"`
	SCParameters         Parameters             `json:"scParameters,omitempty"`
	ReclaimPolicy        ReclaimPolicy          `json:"reclaimPolicy,omitempty"`
	MountOptions         []string               `json:"mountOptions,omitempty"`
	AllowedTopologies    []TopologySelectorTerm `json:"allowedTopologies,omitempty"`
}

type TopologySelectorTerm struct {
	MatchLabelExpressions []TopologySelectorLabelRequirement `json:"matchLabelExpressions,omitempty"`
}
type TopologySelectorLabelRequirement struct {
	Key    string   `json:"key,omitempty"`
	Values []string `json:"values,omitempty"`
}
type Parameters struct {
	GcpPdScParm     map[string]string `json:"gcppdscParm,omitempty"`
	AwsEbsScParm    map[string]string `json:"awsebsscParm,omitempty"`
	AzureDiskScParm map[string]string `json:"azurdiskscParm,omitempty"`
	AzureFileScParm map[string]string `json:"azurfilescParm,omitempty"`
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
