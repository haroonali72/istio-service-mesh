package services

import "bitbucket.org/cloudplex-devs/microservices-mesh-engine/types"

//Id        interface{} `json:"_id,omitempty" bson:"_id" valid:"-"`
//ServiceId string      `json:"service_id" bson:"service_id" binding:"required" valid:"alphanumspecial,length(6|30)~service_id must contain between 6 and 30 characters,lowercase~lowercase alphanumeric characters are allowed,required~service_id is missing in request"`
//Name      string      `json:"name" bson:"name" binding:"required" valid:"alphanumspecial,length(5|30),lowercase~lowercase alphanumeric characters are allowed,required"`
//Version   string      `json:"version"  bson:"version"  binding:"required" valid:"alphanumspecial,length(1|10),lowercase~lowercase alphanumeric characters are allowed,required"`
//Namespace string      `bson:"namespace" json:"namespace",valid:"required"`
//
//ServiceType           constants.ServiceType    `json:"service_type" bson:"service_type" binding:"required" valid:"alphanumspecial,required"`
//ServiceSubType        constants.ServiceSubType `json:"service_sub_type" bson:"service_sub_type" binding:"required" valid:"alphanumspecial,required"`
//Status                string                   `json:"status" bson:"status" binding:"required" valid:"alphanumspecial,required"`
//CompanyId             string                   `json:"company_id" bson:"company_id" binding:"required" valid:"alphanumspecial,required"`
//ServiceDependencyInfo interface{}              `json:"service_dependency_info" bson:"service_dependency_info" `
//CreationDate          time.Time                `json:"creation_date,omitempty" bson:"creation_date" valid:"-"`
type ClusterRole struct {
	types.ServiceBasicInfo `json:",inline" bson:",inline"`
	ServiceAttributes      ClusterRoleSvcAttr `json:"service_attributes" bson:"service_attributes" binding:"required"`
}

type ClusterRoleSvcAttr struct {
	Rules []Rules `json:"rules" bson:"rules"`
}

type Rules struct {
	ResourceName    []string `json:"resource_name" bson:"resource_name"`
	Verbs           []string `json:"verbs" bson:"verbs"`
	ApiGroup        []string `json:"api_group,omitempty" bson:"api_group,omitempty"`
	NonResourceUrls []string `json:"non_resource_urls,omitempty" bson:"non_resource_urls,omitempty"`
}
