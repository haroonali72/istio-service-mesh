package types

import "time"

type StatefulSetService struct {
	Id                interface{}                  `json:"_id,omitempty" bson:"_id" valid:"-"`
	ServiceId         string                       `json:"service_id" bson:"service_id" binding:"required" valid:"alphanumspecial,length(4|30)~service_id must contain between 6 and 30 characters,lowercase~lowercase alphanumeric characters are allowed,required~service_id is missing in request"`
	Name              string                       `json:"name"  bson:"name" binding:"required" valid:"alphanumspecial,length(3|30),lowercase~lowercase alphanumeric characters are allowed,required"`
	Version           string                       `json:"version"  bson:"version"  binding:"required" valid:"alphanumspecial,length(1|10),lowercase~lowercase alphanumeric characters are allowed,required"`
	ServiceType       string                       `json:"service_type"  bson:"service_type" valid:"-"`
	ServiceSubType    string                       `json:"service_sub_type" bson:"service_type" valid:"-"`
	Namespace         string                       `json:"namespace" bson:"namespace" binding:"required" valid:"alphanumspecial,required"`
	CompanyId         string                       `json:"company_id,omitempty" bson:"company_id"`
	CreationDate      time.Time                    `json:"creation_date,omitempty" bson:"creation_date" valid:"-"`
	ServiceAttributes *StatefulSetServiceAttribute `json:"service_attributes"  bson:"company_id" binding:"required"`
}
type StatefulSetServiceAttribute struct {
	Labels                        map[string]string              `json:"labels,omitempty"`
	Annotations                   map[string]string              `json:"annotations,omitempty"`
	LabelSelector                 *LabelSelectorObj              `json:"label_selector"`
	Replicas                      *Replicas                      `json:"replicas,omitempty"`
	RevisionHistoryLimit          *RevisionHistoryLimit          `json:"revisionHistoryLimit,omitempty"`
	UpdateStrategy                *StateFulSetUpdateStrategy     `json:"updateStrategy,omitempty"`
	Containers                    map[string]ContainerAttribute  `json:"containers,omitempty"`
	InitContainers                map[string]ContainerAttribute  `json:"initContainers,omitempty"`
	PodManagementPolicy           PodManagementPolicyType        `json:"podManagementPolicy,omitempty"`
	MeshConfig                    *IstioConfig                   `json:"istio_config,omitempty"`
	NodeSelector                  map[string]string              `json:"node_selector"`
	IsRbac                        bool                           `json:"is_rbac_enabled"`
	RbacRoles                     []K8sRbacAttribute             `json:"roles,omitempty"`
	IstioRoles                    []IstioRbacAttribute           `json:"istio_roles,omitempty"`
	Volumes                       []Volume                       `json:"volumes,omitempty"`
	Affinity                      *Affinity                      `json:"affinity, omitempty"`
	ServiceName                   string                         `json:"serviceName"`
	TerminationGracePeriodSeconds *TerminationGracePeriodSeconds `json:"terminationGracePeriodSeconds,omitempty"`
	VolumeClaimTemplates          []PersistentVolumeClaimService `json:"volumeClaimTemplates, omitempty"`
}
