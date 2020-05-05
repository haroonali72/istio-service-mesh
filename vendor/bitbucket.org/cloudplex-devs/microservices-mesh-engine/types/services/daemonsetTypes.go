package services

import (
	"bitbucket.org/cloudplex-devs/microservices-mesh-engine/types"
)

//Id                interface{}                `json:"_id,omitempty" bson:"_id" valid:"-"`
//ServiceId         string                     `json:"service_id" bson:"service_id" binding:"required" valid:"alphanumspecial,length(4|30)~service_id must contain between 6 and 30 characters,lowercase~lowercase alphanumeric characters are allowed,required~service_id is missing in request"`
//Name              string                     `json:"name"  bson:"name" binding:"required" valid:"alphanumspecial,length(3|30),lowercase~lowercase alphanumeric characters are allowed,required"`
//Version           string                     `json:"version"  bson:"version"  binding:"required" valid:"alphanumspecial,length(1|10),lowercase~lowercase alphanumeric characters are allowed,required"`
//ServiceType       constants.ServiceType      `json:"service_type"  bson:"service_type" valid:"-"`
//ServiceSubType    constants.ServiceSubType   `json:"service_sub_type" bson:"service_type" valid:"-"`
//Namespace         string                     `json:"namespace" bson:"namespace" binding:"required" valid:"alphanumspecial,required"`
//CompanyId         string                     `json:"company_id,omitempty" bson:"company_id"`
//CreationDate      time.Time                  `json:"creation_date,omitempty" bson:"creation_date" valid:"-"`
type DaemonSetService struct {
	types.ServiceBasicInfo `json:",inline" bson:",inline"`
	ServiceAttributes      *DaemonSetServiceAttribute `json:"service_attributes,omitempty"  bson:"company_id" binding:"required"`
}
type DaemonSetServiceAttribute struct {
	CommonContainerAttributes `json:",inline,omitempty" bson:",inline,omitempty"`
	// The minimum number of seconds for which a newly created DaemonSet pod should
	// be ready without any of its container crashing, for it to be considered
	// available. Defaults to 0 (pod will be considered available as soon as it
	// is ready).
	// +optional
	MinReadySeconds int32 `json:"min_ready_seconds,omitempty" default:"0"`
	// An update strategy to replace existing DaemonSet pods with new pods.
	// +optional
	UpdateStrategy *DaemonSetUpdateStrategy `json:"update_strategy,omitempty"`
	// The number of old history to retain to allow rollback.
	// This is a pointer to distinguish between explicit zero and not specified.
	// Defaults to 10.
	// +optional
	RevisionHistoryLimit *RevisionHistoryLimit `json:"revision_history_limit,omitempty" default:"10"`
}

type RevisionHistoryLimit struct {
	Value int32 `json:"value,omitempty"`
}
