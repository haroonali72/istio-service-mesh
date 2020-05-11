package services

import "bitbucket.org/cloudplex-devs/microservices-mesh-engine/types"

//Id                interface{}                 `json:"_id,omitempty" bson:"_id" valid:"-"`
//ServiceId         string                      `json:"service_id" bson:"service_id" binding:"required" valid:"alphanumspecial,length(4|30)~service_id must contain between 6 and 30 characters,lowercase~lowercase alphanumeric characters are allowed,required~service_id is missing in request"`
//Name              string                      `json:"name"  bson:"name" binding:"required" valid:"alphanumspecial,length(3|30),lowercase~lowercase alphanumeric characters are allowed,required"`
//Version           string                      `json:"version"  bson:"version"  binding:"required" valid:"alphanumspecial,length(1|10),lowercase~lowercase alphanumeric characters are allowed,required"`
//ServiceType       constants.ServiceType       `json:"service_type"  bson:"service_type" valid:"-"`
//ServiceSubType    constants.ServiceSubType    `json:"service_sub_type" bson:"service_type" valid:"-"`
//Namespace         string                      `json:"namespace" bson:"namespace" binding:"required" valid:"alphanumspecial,required"`
//CompanyId         string                      `json:"company_id,omitempty" bson:"company_id"`
//CreationDate      time.Time                   `json:"creation_date,omitempty" bson:"creation_date" valid:"-"`
type DeploymentService struct {
	types.ServiceBasicInfo `json:",inline" bson:",inline"`
	ServiceAttributes      *DeploymentServiceAttribute `json:"service_attributes"  bson:"service_attributes" binding:"required"`
}
type DeploymentServiceAttribute struct {
	// Number of desired pods. This is a pointer to distinguish between explicit
	// zero and not specified. Defaults to 1.
	// +optional
	Replicas *int32 `json:"replicas,omitempty" bson:"replicas,omitempty" default:"1"`
	// The deployment strategy to use to replace existing pods with new ones.
	// +optional
	Strategy                  *DeploymentStrategy `json:"strategy,omitempty" bson:"strategy,omitempty"`
	CommonContainerAttributes `json:",inline,omitempty" bson:",inline,omitempty"`
}

type AutomountServiceAccountToken struct {
	Value bool `json:"value,omitempty" bson:"value,omitempty"`
}
