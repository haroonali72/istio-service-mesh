package services

import "time"

type ClusterRoleBinding struct {
	Id        interface{} `json:"_id,omitempty" bson:"_id" valid:"-"`
	ServiceId string      `json:"service_id" bson:"service_id" binding:"required" valid:"alphanumspecial,length(6|30)~service_id must contain between 6 and 30 characters,lowercase~lowercase alphanumeric characters are allowed,required~service_id is missing in request"`
	Name      string      `json:"name" bson:"name" binding:"required" valid:"alphanumspecial,length(5|30),lowercase~lowercase alphanumeric characters are allowed,required"`
	Version   string      `json:"version"  bson:"version"  binding:"required" valid:"alphanumspecial,length(1|10),lowercase~lowercase alphanumeric characters are allowed,required"`
	Namespace string      `bson:"namespace" json:"namespace",valid:"required"`

	ServiceType           string                    `json:"service_type" bson:"service_type" binding:"required" valid:"alphanumspecial,required"`
	ServiceSubType        string                    `json:"service_sub_type" bson:"service_sub_type" binding:"required" valid:"alphanumspecial,required"`
	Status                string                    `json:"status" bson:"status" binding:"required" valid:"alphanumspecial,required"`
	CompanyId             string                    `json:"company_id" bson:"company_id" binding:"required" valid:"alphanumspecial,required"`
	ServiceDependencyInfo interface{}               `json:"service_dependency_info" bson:"service_dependency_info" `
	CreationDate          time.Time                 `json:"creation_date,omitempty" bson:"creation_date" valid:"-"`
	ServiceAttributes     ClusterRoleBindingSvcAttr `json:"service_attributes" bson:"service_attributes" binding:"required"`
}

type ClusterRoleBindingSvcAttr struct {
	Subjects           []Subject `json:"subjects" bson:"subjects"`
	NameClusterRoleRef string    `json:"name_cluster_role_ref"`
}

//type Subject struct {
//	Kind      string `json:"kind" bson:"kind"`
//	Name      string `json:"name" bson:"name"`
//
//}
//
//type RoleReferenceClusterRolebinding struct {
//	ApiGroup string `json:"api_group" bson:"api_group"`
//	Kind     string `json:"kind" bson:"kind"`
//	Name     string `json:"name" bson:"name"`
//}
