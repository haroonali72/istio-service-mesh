package services

import "time"

type ServiceAccount struct {
	Id                interface{}              `json:"_id,omitempty" bson:"_id" valid:"-"`
	CompanyId         string                   `bson:"company_id" json:"company_id",valid:"required"`
	ServiceId         string                   `bson:"service_id" json:"service_id",valid:"required"`
	ServiceType       string                   `bson:"service_type" json:"service_type",valid:"required"`
	Version           string                   `bson:"version" json:"version",valid:"required"`
	Name              string                   `bson:"name" json:"name",valid:"required"`
	Namespace         string                   `bson:"namespace" json:"namespace",valid:"required"`
	ServiceSubType    string                   `json:"service_sub_type" bson:"service_type" valid:"-"`
	CreationDate      time.Time                `json:"creation_date,omitempty" bson:"creation_date" valid:"-"`
	ServiceAttributes *ServiceAccountAttribute `json:"service_attributes"  bson:"company_id" binding:"required"`
}

type ServiceAccountAttribute struct {
	Secrets              []string `json:"secrets,omitempty"`
	ImagePullSecretsName []string `json:"image_pull_secrets_name,omitempty"`
}
