package services

import (
	"bitbucket.org/cloudplex-devs/microservices-mesh-engine/types"
)

//ServiceId             string                   `bson:"service_id" json:"service_id",valid:"required"`
//	CompanyId             string                   `bson:"company_id" json:"company_id",valid:"required"`
//	Name                  string                   `bson:"name" json:"name",valid:"required"`
//	Version               string                   `bson:"version" json:"version",valid:"required"`
//	ServiceType           constants.ServiceType    `bson:"service_type" json:"service_type",valid:"required"`
//	ServiceSubType        constants.ServiceSubType `bson:"service_sub_type" json:"service_sub_type",valid:"required,in(rolebinding)"`
//	ServiceDependencyInfo interface{}              `bson:"service_dependency_info" json:"service_dependency_info",valid:"required"`
//	Namespace             string                   `bson:"namespace" json:"namespace",valid:"required"`
type RoleBinding struct {
	types.ServiceBasicInfo `json:",inline" bson:",inline"`
	ServiceAttributes      ServiceAttributee `bson:"service_attributes" json:"service_attributes",valid:"required"`
}

type ServiceAttributee struct {
	Subjects []Subject     `bson:"subjects" json:"subjects"`
	RoleRef  RoleReference `bson:"reference" json:"reference"`
	Hostname []string      `bson:"hostname" json:"hostname"`
}
type Subject struct {
	Kind      string `bson:"kind" json:"kind"`
	Name      string `bson:"name" json:"name"`
	Namespace string `json:"namespace" bson:"namespace"`
}

type RoleReference struct {
	Kind     string `bson:"kind" json:"kind"`
	Name     string `bson:"name" json:"name"`
	ApiGroup string `json:"api_group,omitempty" bson:"api_group,omitempty"`
}
