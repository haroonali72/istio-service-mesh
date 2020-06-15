package services

import (
	"bitbucket.org/cloudplex-devs/microservices-mesh-engine/types"
)

//Id                interface{}              `json:"_id,omitempty" bson:"_id" valid:"-"`
//	CompanyId         string                   `bson:"company_id" json:"company_id",valid:"required"`
//	ServiceId         string                   `bson:"service_id" json:"service_id",valid:"required"`
//	ServiceType       constants.ServiceType    `bson:"service_type" json:"service_type",valid:"required"`
//	Version           string                   `bson:"version" json:"version",valid:"required"`
//	Name              string                   `bson:"name" json:"name",valid:"required"`
//	Namespace         string                   `bson:"namespace" json:"namespace",valid:"required"`
//	ServiceSubType    constants.ServiceSubType `json:"service_sub_type" bson:"service_type" valid:"-"`
//	CreationDate      time.Time                `json:"creation_date,omitempty" bson:"creation_date" valid:"-"`
type ServiceAccount struct {
	types.ServiceBasicInfo `json:",inline" bson:",inline"`
	ServiceAttributes      *ServiceAccountAttribute `json:"service_attributes"  bson:"service_attributes" binding:"required"`
}

type ServiceAccountAttribute struct {
	Secrets              []string `json:"secrets,omitempty" bson:"secrets,omitempty"`
	ImagePullSecretsName []string `json:"image_pull_secrets_name,omitempty" bson:"image_pull_secrets_name,omitempty"`
}
