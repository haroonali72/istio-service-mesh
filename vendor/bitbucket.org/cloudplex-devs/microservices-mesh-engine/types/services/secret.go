package services

import (
	"bitbucket.org/cloudplex-devs/microservices-mesh-engine/types"
)

//Id                    interface{}              `json:"_id,omitempty" bson:"_id" valid:"-"`
//	ServiceId             string                   `bson:"service_id" json:"service_id",valid:"required"`
//	CompanyId             string                   `bson:"company_id" json:"company_id",valid:"required"`
//	Name                  string                   `bson:"name" json:"name",valid:"required"`
//	Version               string                   `bson:"version" json:"version",valid:"required"`
//	ServiceType           constants.ServiceType    `bson:"service_type" json:"service_type",valid:"required"`
//	ServiceSubType        constants.ServiceSubType `bson:"service_sub_type" json:"service_sub_type",valid:"required"`
//	ServiceDependencyInfo interface{}              `bson:"service_dependency_info" json:"service_dependency_info",valid:"required"`
//	Namespace             string                   `bson:"namespace" json:"namespace",valid:"required"`
type Secret struct {
	types.ServiceBasicInfo `json:",inline" bson:",inline"`
	ServiceAttributes      *SecretServiceAttribute `bson:"service_attributes" json:"service_attributes",valid:"required"`
}

type SecretServiceAttribute struct {
	Data       map[string][]byte `bson:"data,omitempty" json:"data,omitempty"`
	StringData map[string]string `bson:"secret_data,omitempty" json:"secret_data,omitempty",valid:"-"`
	SecretType string            `bson:"secret_type" json:"secret_type",valid:"required,in(Opaque|ServiceAccountToken|ServiceAccountNameKey|ServiceAccountUIDKey|ServiceAccountTokenKey|ServiceAccountKubeconfigKey|ServiceAccountRootCAKey|SecretTypeDockercfg|DockerConfigKey|Tls)"`
}
