package types

import (
	"time"
)

type DeploymentService struct {
	Id                interface{}                 `json:"_id,omitempty" bson:"_id" valid:"-"`
	ServiceId         string                      `json:"service_id" bson:"service_id" binding:"required" valid:"alphanumspecial,length(4|30)~service_id must contain between 6 and 30 characters,lowercase~lowercase alphanumeric characters are allowed,required~service_id is missing in request"`
	Name              string                      `json:"name"  bson:"name" binding:"required" valid:"alphanumspecial,length(3|30),lowercase~lowercase alphanumeric characters are allowed,required"`
	Version           string                      `json:"version"  bson:"version"  binding:"required" valid:"alphanumspecial,length(1|10),lowercase~lowercase alphanumeric characters are allowed,required"`
	ServiceType       string                      `json:"service_type"  bson:"service_type" valid:"-"`
	ServiceSubType    string                      `json:"service_sub_type" bson:"service_type" valid:"-"`
	Namespace         string                      `json:"namespace" bson:"namespace" binding:"required" valid:"alphanumspecial,required"`
	CompanyId         string                      `json:"company_id,omitempty" bson:"company_id"`
	CreationDate      time.Time                   `json:"creation_date,omitempty" bson:"creation_date" valid:"-"`
	ServiceAttributes *DeploymentServiceAttribute `json:"service_attributes"  bson:"company_id" binding:"required"`
}
type DeploymentServiceAttribute struct {
	Container     map[string]ContainerAttribute `json:"containers,omitempty"`
	MeshConfig    *IstioConfig                  `json:"istio_config,omitempty"`
	NodeSelector  map[string]string             `json:"node_selector,omitempty"`
	Labels        map[string]string             `json:"labels,omitempty"`
	LabelSelector *LabelSelectorObj             `json:"label_selector,omitempty"`
	Annotations   map[string]string             `json:"annotations,omitempty"`
	RbacRoles     []K8sRbacAttribute            `json:"roles,omitempty"`
	IstioRoles    []IstioRbacAttribute          `json:"istio_roles,omitempty"`
	Volumes       []Volume                      `json:"volumes,omitempty"`
	Affinity      *Affinity                     `json:"affinity,omitempty"`
	Strategy      DeploymentStrategy            `json:"strategy,omitempty"`
	InitContainer map[string]ContainerAttribute `json:"initContainers,omitempty"`
	Replicas      Replica                       `json:"replicas,omitempty"`
}

type Replica struct {
	Replica int32 `json:"replica,omitempty"`
}
