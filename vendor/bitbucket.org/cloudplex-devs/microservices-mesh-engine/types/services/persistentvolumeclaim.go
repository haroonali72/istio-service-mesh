package services

import "time"

type PersistentVolumeClaimService struct {
	Id                interface{}                            `json:"_id,omitempty" bson:"_id" valid:"-"`
	ServiceId         string                                 `json:"service_id" bson:"service_id" binding:"required" valid:"alphanumspecial,length(4|30)~service_id must contain between 6 and 30 characters,lowercase~lowercase alphanumeric characters are allowed,required~service_id is missing in request"`
	Name              string                                 `json:"name"  bson:"name" binding:"required" valid:"alphanumspecial,length(3|30),lowercase~lowercase alphanumeric characters are allowed,required"`
	Version           string                                 `json:"version"  bson:"version"  binding:"required" valid:"alphanumspecial,length(1|10),lowercase~lowercase alphanumeric characters are allowed,required"`
	ServiceType       string                                 `json:"service_type"  bson:"service_type" valid:"-"`
	ServiceSubType    string                                 `json:"service_sub_type" bson:"service_type" valid:"-"`
	Namespace         string                                 `json:"namespace" bson:"namespace" binding:"required" valid:"alphanumspecial,required"`
	CompanyId         string                                 `json:"company_id,omitempty" bson:"company_id"`
	CreationDate      time.Time                              `json:"creation_date,omitempty" bson:"creation_date" valid:"-"`
	ServiceAttributes *PersistentVolumeClaimServiceAttribute `json:"service_attributes"  bson:"company_id" binding:"required"`
}
type PersistentVolumeClaimServiceAttribute struct {
	LabelSelector    *LabelSelectorObj          `json:"label_selector,omitempty"`
	VolumeName       string                     `json:"volume_name,omitempty"`
	AccessMode       []AccessMode               `json:"access_mode,omitempty"`
	StorageClassName string                     `json:"storage_class_name,omitempty"`
	Request          string                     `json:"request_quantity,omitempty"`
	Limit            string                     `json:"limit_quantity,omitempty"`
	VolumeMode       *PersistentVolumeMode      `json:"volume_mode,omitempty" protobuf:"bytes,8,opt,name=volumeMode,casttype=PersistentVolumeMode"`
	DataSource       *TypedLocalObjectReference `json:"data_source,omitempty" protobuf:"bytes,7,opt,name=dataSource"`
}

// TypedLocalObjectReference contains enough information to let you locate the
// typed referenced object inside the same namespace.
type TypedLocalObjectReference struct {
	// APIGroup is the group for the resource being referenced.
	// If APIGroup is not specified, the specified Kind must be in the core API group.
	// For any other third-party types, APIGroup is required.
	// +optional
	APIGroup *string `json:"api_group,omitempty" protobuf:"bytes,1,opt,name=apiGroup"`
	// Kind is the type of resource being referenced
	Kind string `json:"kind" protobuf:"bytes,2,opt,name=kind"`
	// Name is the name of resource being referenced
	Name string `json:"name" protobuf:"bytes,3,opt,name=name"`
}
