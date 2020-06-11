package services

import (
	"bitbucket.org/cloudplex-devs/microservices-mesh-engine/types"
)

//Id                interface{}          `json:"_id,omitempty" bson:"_id" valid:"-"`
//	ServiceId         string               `json:"service_id" bson:"service_id" binding:"required" valid:"alphanumspecial,length(4|30)~service_id must contain between 6 and 30 characters,lowercase~lowercase alphanumeric characters are allowed,required~service_id is missing in request"`
//	Name              string               `json:"name"  bson:"name" binding:"required" valid:"alphanumspecial,length(3|30),lowercase~lowercase alphanumeric characters are allowed,required"`
//	Version           string               `json:"version"  bson:"version"  binding:"required" valid:"alphanumspecial,length(1|10),lowercase~lowercase alphanumeric characters are allowed,required"`
//	ServiceType       string               `json:"service_type"  bson:"service_type" valid:"-"`
//	ServiceSubType    string               `json:"service_sub_type" bson:"service_type" valid:"-"`
//	Namespace         string               `json:"namespace" bson:"namespace" binding:"required" valid:"alphanumspecial,required"`
//	CompanyId         string               `json:"company_id,omitempty" bson:"company_id"`
//	CreationDate      time.Time            `json:"creation_date,omitempty" bson:"creation_date" valid:"-"`
type PodService struct {
	types.ServiceBasicInfo `json:",inline" bson:",inline"`
	ServiceAttributes      *PodServiceAttribute `json:"service_attributes, omitempty"  bson:"service_attributes" binding:"required"`
}
type PodServiceAttribute struct {
	CommonContainerAttributes `json:",inline,omitempty" bson:",inline,omitempty"`
	RestartPolicy             RestartPolicy `json:"restart_policy,omitempty" bson:"restart_policy,omitempty"`
}

type RestartPolicy string

const (
	RestartPolicyAlways    RestartPolicy = "Always"
	RestartPolicyOnFailure RestartPolicy = "OnFailure"
	RestartPolicyNever     RestartPolicy = "Never"
)
