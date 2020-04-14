package services

import "time"

type DaemonSetService struct {
	Id                interface{}                `json:"_id,omitempty" bson:"_id" valid:"-"`
	ServiceId         string                     `json:"service_id" bson:"service_id" binding:"required" valid:"alphanumspecial,length(4|30)~service_id must contain between 6 and 30 characters,lowercase~lowercase alphanumeric characters are allowed,required~service_id is missing in request"`
	Name              string                     `json:"name"  bson:"name" binding:"required" valid:"alphanumspecial,length(3|30),lowercase~lowercase alphanumeric characters are allowed,required"`
	Version           string                     `json:"version"  bson:"version"  binding:"required" valid:"alphanumspecial,length(1|10),lowercase~lowercase alphanumeric characters are allowed,required"`
	ServiceType       string                     `json:"service_type"  bson:"service_type" valid:"-"`
	ServiceSubType    string                     `json:"service_sub_type" bson:"service_type" valid:"-"`
	Namespace         string                     `json:"namespace" bson:"namespace" binding:"required" valid:"alphanumspecial,required"`
	CompanyId         string                     `json:"company_id,omitempty" bson:"company_id"`
	CreationDate      time.Time                  `json:"creation_date,omitempty" bson:"creation_date" valid:"-"`
	ServiceAttributes *DaemonSetServiceAttribute `json:"service_attributes,omitempty"  bson:"company_id" binding:"required"`
}
type DaemonSetServiceAttribute struct {
	Labels               map[string]string        `json:"labels,omitempty"`
	Annotations          map[string]string        `json:"annotations,omitempty"`
	LabelSelector        *LabelSelectorObj        `json:"label_selector"`
	UpdateStrategy       *DaemonSetUpdateStrategy `json:"update_strategy,omitempty"`
	MinReadySeconds      int32                    `json:"min_ready_seconds,omitempty"`
	RevisionHistoryLimit *RevisionHistoryLimit    `json:"revision_history_limit,omitempty"`
	Volumes              []Volume                 `json:"volumes,omitempty"`
	Containers           []*ContainerAttribute    `json:"containers,omitempty"`
	InitContainers       []*ContainerAttribute    `json:"initContainers,omitempty"`
	NodeSelector         map[string]string        `json:"node_selector"`
	MeshConfig           *IstioConfig             `json:"istio_config,omitempty"`
	Affinity             *Affinity                `json:"affinity,omitempty"`
	IsRbac               bool                     `json:"is_rbac_enabled"`
	RbacRoles            []K8sRbacAttribute       `json:"roles,omitempty"`
	IstioRoles           []IstioRbacAttribute     `json:"istio_roles,omitempty"`
}

type RevisionHistoryLimit struct {
	Value int32 `json:"value,omitempty"`
}
