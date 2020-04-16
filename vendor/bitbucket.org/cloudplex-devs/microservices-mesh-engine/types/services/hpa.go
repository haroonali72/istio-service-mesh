package services

import "time"

type HPA struct {
	Id        interface{} `json:"_id,omitempty" bson:"_id" valid:"-"`
	ServiceId string      `json:"service_id" bson:"service_id" binding:"required" valid:"alphanumspecial,length(6|30)~service_id must contain between 6 and 30 characters,lowercase~lowercase alphanumeric characters are allowed,required~service_id is missing in request"`
	Name      string      `json:"name" bson:"name" binding:"required" valid:"alphanumspecial,length(5|30),lowercase~lowercase alphanumeric characters are allowed,required"`
	Version   string      `json:"version"  bson:"version"  binding:"required" valid:"alphanumspecial,length(1|10),lowercase~lowercase alphanumeric characters are allowed,required"`

	ServiceType           string      `json:"service_type" bson:"service_type" binding:"required" valid:"alphanumspecial,required"`
	CompanyId             string      `json:"company_id" bson:"company_id" binding:"required" valid:"alphanumspecial,required"`
	Status                string      `json:"status" bson:"status" binding:"required" valid:"alphanumspecial,required"`
	ServiceSubType        string      `json:"service_sub_type" bson:"service_sub_type" binding:"required" valid:"alphanumspecial,required"`
	ServiceDependencyInfo interface{} `json:"service_dependency_info" bson:"service_dependency_info" `
	Namespace             string      `json:"namespace" bson:"namespace"`
	CreationDate          time.Time   `json:"creation_date,omitempty" bson:"creation_date" valid:"-"`
	ServiceAttributes     HpaSvcAttr  `json:"service_attributes" bson:"service_attributes" binding:"required"`
}

type HpaSvcAttr struct {
	MinReplicas          int                `json:"min_replicas" bson:"min_replicas"`
	MaxReplicas          int                `json:"max_replicas" bson:"max_replicas"`
	CrossObjectVersion   CrossObjectVersion `json:"cross_object_version" bson:"cross_object_version"`
	TargetCpuUtilization *int32             `json:"target_cpu_utilization,omitempty" bson:"target_cpu_utilization"`
}

type CrossObjectVersion struct {
	Name    string `json:"name" bson:"name"`
	Version string `json:"version" bson:"version"`
	Type    string `json:"type" bson:"type"`
}

type MetricValue struct {
	TargetValueKind string `json:"target_value_kind" bson:"target_value_kind"`
	TargetValue     string `json:"target_value" bson:"target_value"`
	ResourceKind    string `json:"resource_kind" bson:"resource_kind"`
}
