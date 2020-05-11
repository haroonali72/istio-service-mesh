package services

import (
	"bitbucket.org/cloudplex-devs/microservices-mesh-engine/types"
)

//Id                interface{}                  `json:"_id,omitempty" bson:"_id" valid:"-"`
//	ServiceId         string                       `json:"service_id" bson:"service_id" binding:"required" valid:"alphanumspecial,length(4|30)~service_id must contain between 6 and 30 characters,lowercase~lowercase alphanumeric characters are allowed,required~service_id is missing in request"`
//	Name              string                       `json:"name"  bson:"name" binding:"required" valid:"alphanumspecial,length(3|30),lowercase~lowercase alphanumeric characters are allowed,required"`
//	Version           string                       `json:"version"  bson:"version"  binding:"required" valid:"alphanumspecial,length(1|10),lowercase~lowercase alphanumeric characters are allowed,required"`
//	ServiceType       constants.ServiceType        `json:"service_type"  bson:"service_type" valid:"-"`
//	ServiceSubType    constants.ServiceSubType     `json:"service_sub_type" bson:"service_type" valid:"-"`
//	Namespace         string                       `json:"namespace" bson:"namespace" binding:"required" valid:"alphanumspecial,required"`
//	CompanyId         string                       `json:"company_id,omitempty" bson:"company_id"`
//	CreationDate      time.Time                    `json:"creation_date,omitempty" bson:"creation_date" valid:"-"`
type StatefulSetService struct {
	types.ServiceBasicInfo `json:",inline" bson:",inline"`
	ServiceAttributes      *StatefulSetServiceAttribute `json:"service_attributes"  bson:"company_id" binding:"required"`
}
type StatefulSetServiceAttribute struct {
	CommonContainerAttributes `json:",inline,omitempty" bson:",inline,omitempty"`
	// replicas is the desired number of replicas of the given Template.
	// These are replicas in the sense that they are instantiations of the
	// same Template, but individual replicas also have a consistent identity.
	// If unspecified, defaults to 1.
	// +optional
	Replicas *int32 `json:"replicas,omitempty" bson:"replicas,omitempty" default:"1"`
	// revisionHistoryLimit is the maximum number of revisions that will
	// be maintained in the StatefulSet's revision history. The revision history
	// consists of all revisions not represented by a currently applied
	// StatefulSetSpec version. The default value is 10.
	// +optional
	RevisionHistoryLimit *RevisionHistoryLimit `json:"revision_history_limit,omitempty" bson:"revision_history_limit,omitempty"`
	// updateStrategy indicates the StatefulSetUpdateStrategy that will be
	// employed to update Pods in the StatefulSet when a revision is made to
	// Template.
	UpdateStrategy *StateFulSetUpdateStrategy `json:"update_Strategy,omitempty" bson:"update_Strategy,omitempty"`
	// podManagementPolicy controls how pods are created during initial scale up,
	// when replacing pods on nodes, or when scaling down. The default policy is
	// "OrderedReady", where pods are created in increasing order (pod-0, then
	// pod-1, etc) and the controller will wait until each pod is ready before
	// continuing. When scaling down, the pods are removed in the opposite order.
	// The alternative policy is "Parallel" which will create pods in parallel
	// to match the desired scale without waiting, and on scale down will delete
	// all pods at once.
	// +optional
	PodManagementPolicy PodManagementPolicyType `json:"pod_management_policy,omitempty"  bson:"pod_management_policy,omitempty" swaggerType:"string" default:"OrderedReady" swaggerType:"string"`
	// serviceName is the name of the service that governs this StatefulSet.
	// This service must exist before the StatefulSet, and is responsible for
	// the network identity of the set. Pods get DNS/hostnames that follow the
	// pattern: pod-specific-string.serviceName.default.svc.cluster.local
	// where "pod-specific-string" is managed by the StatefulSet controller.
	// +optional
	ServiceName string `json:"service_name,omitempty" json:"service_name,omitempty"`
	// volumeClaimTemplates is a list of claims that pods are allowed to reference.
	// +optional
	VolumeClaimTemplates []PersistentVolumeClaimService `json:"volume_claim_templates,omitempty" bson:"volume_claim_templates,omitempty"`
}
