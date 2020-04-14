package services

import "time"

type JobService struct {
	Id                interface{}          `json:"_id,omitempty" bson:"_id" valid:"-"`
	ServiceId         string               `json:"service_id" bson:"service_id" binding:"required" valid:"alphanumspecial,length(4|30)~service_id must contain between 6 and 30 characters,lowercase~lowercase alphanumeric characters are allowed,required~service_id is missing in request"`
	Name              string               `json:"name"  bson:"name" binding:"required" valid:"alphanumspecial,length(3|30),lowercase~lowercase alphanumeric characters are allowed,required"`
	Version           string               `json:"version"  bson:"version"  binding:"required" valid:"alphanumspecial,length(1|10),lowercase~lowercase alphanumeric characters are allowed,required"`
	ServiceType       string               `json:"service_type"  bson:"service_type" valid:"-"`
	ServiceSubType    string               `json:"service_sub_type" bson:"service_type" valid:"-"`
	Namespace         string               `json:"namespace" bson:"namespace" binding:"required" valid:"alphanumspecial,required"`
	CompanyId         string               `json:"company_id,omitempty" bson:"company_id"`
	CreationDate      time.Time            `json:"creation_date,omitempty" bson:"creation_date" valid:"-"`
	ServiceAttributes *JobServiceAttribute `json:"service_attributes"  bson:"company_id" binding:"required"`
}
type JobServiceAttribute struct {
	Parallelism             *Parallelism             `json:"parallelism,omitempty"`
	Completions             *Completions             `json:"completions,omitempty"`
	ActiveDeadlineSeconds   *ActiveDeadlineSeconds   `json:"active_deadline_seconds,omitempty"`
	BackoffLimit            *BackoffLimit            `json:"backoff_limit,omitempty"`
	ManualSelector          *ManualSelector          `json:"manual_selector, omitempty"`
	Affinity                *Affinity                `json:"affinity,omitempty"`
	TTLSecondsAfterFinished *TTLSecondsAfterFinished `json:"ttl_seconds_after_finished,omitempty"`
	LabelSelector           *LabelSelectorObj        `json:"label_selector,omitempty"`
	NodeSelector            map[string]string        `json:"node_selector"`
	Labels                  map[string]string        `json:"labels,omitempty"`
	Annotations             map[string]string        `json:"annotations,omitempty"`
	IsRbac                  bool                     `json:"is_rbac_enabled"`
	RbacRoles               []K8sRbacAttribute       `json:"roles,omitempty"`
	IstioRoles              []IstioRbacAttribute     `json:"istio_roles,omitempty"`
	Containers              []*ContainerAttribute    `json:"containers,omitempty"`
	InitContainers          []*ContainerAttribute    `json:"init_containers,omitempty"`
	Volumes                 []Volume                 `json:"volumes,omitempty"`
}

type Parallelism struct {
	Value int32 `json:"value,omitempty"`
}

type Completions struct {
	Value int32 `json:"value,omitempty"`
}

type BackoffLimit struct {
	Value int32 `json:"value,omitempty"`
}

type ManualSelector struct {
	Value bool `json:"value,omitempty"`
}

type TTLSecondsAfterFinished struct {
	Value int32 `json:"value,omitempty"`
}
