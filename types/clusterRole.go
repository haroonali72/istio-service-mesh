package types

import "time"

type ClusterRole struct {
	Id                    interface{}        `json:"_id,omitempty" bson:"_id" valid:"-"`
	ServiceId             string             `json:"service_id" bson:"service_id" binding:"required" valid:"alphanumspecial,length(6|30)~service_id must contain between 6 and 30 characters,lowercase~lowercase alphanumeric characters are allowed,required~service_id is missing in request"`
	Name                  string             `json:"name" bson:"name" binding:"required" valid:"alphanumspecial,length(5|30),lowercase~lowercase alphanumeric characters are allowed,required"`
	ServiceType           string             `json:"service_type" bson:"service_type" binding:"required" valid:"alphanumspecial,required"`
	ServiceSubType        string             `json:"service_sub_type" bson:"service_sub_type" binding:"required" valid:"alphanumspecial,required"`
	Status                string             `json:"status" bson:"status" binding:"required" valid:"alphanumspecial,required"`
	CompanyId             string             `json:"company_id" bson:"company_id" binding:"required" valid:"alphanumspecial,required"`
	ServiceDependencyInfo interface{}        `json:"service_dependency_info" bson:"service_dependency_info" `
	CreationDate          time.Time          `json:"creation_date,omitempty" bson:"creation_date" valid:"-"`
	ServiceAttributes     ClusterRoleSvcAttr `json:"service_attributes" bson:"service_attributes" binding:"required"`
}

type ClusterRoleSvcAttr struct {
	Rules []Rules `json:"rules" bson:"rules"`
}

type Rules struct {
	Resources    []string `bson:"resources" json:"resources"`
	ResourceName []string `json:"resource_name" bson:"resource_name"`
	Verbs        []string `json:"verbs" bson:"verbs"`
	ApiGroup     []string `json:"api_group" bson:"api_group"`
}
