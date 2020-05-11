package services

import (
	"bitbucket.org/cloudplex-devs/microservices-mesh-engine/types"
)

//Id                interface{}                    `json:"_id,omitempty" bson:"_id" valid:"-"`
//	ServiceId         string                         `json:"service_id" bson:"service_id" binding:"required" valid:"alphanumspecial,length(4|30)~service_id must contain between 6 and 30 characters,lowercase~lowercase alphanumeric characters are allowed,required~service_id is missing in request"`
//	Name              string                         `json:"name"  bson:"name" binding:"required" valid:"alphanumspecial,length(3|30),lowercase~lowercase alphanumeric characters are allowed,required"`
//	Version           string                         `json:"version"  bson:"version"  binding:"required" valid:"alphanumspecial,length(1|10),lowercase~lowercase alphanumeric characters are allowed,required"`
//	ServiceType       constants.ServiceType          `json:"service_type"  bson:"service_type" valid:"-"`
//	ServiceSubType    constants.ServiceSubType       `json:"service_sub_type" bson:"service_type" valid:"-"`
//	Namespace         string                         `json:"namespace" bson:"namespace" binding:"required" valid:"alphanumspecial,required"`
//	CompanyId         string                         `json:"company_id,omitempty" bson:"company_id"`
//	CreationDate      time.Time                      `json:"creation_date,omitempty" bson:"creation_date" valid:"-"`
type NetworkPolicyService struct {
	types.ServiceBasicInfo `json:",inline" bson:",inline"`
	ServiceAttributes      *NetworkPolicyServiceAttribute `json:"service_attributes"  bson:"company_id" binding:"required"`
}
type NetworkPolicyServiceAttribute struct {
	PodSelector *LabelSelectorObj `json:"pod_selector,omitempty" bson:"pod_selector,omitempty"` //empty means all po in np namespaces
	Ingress     []IngressRule     `json:"ingress,omitempty" bson:"ingress,omitempty"`
	Egress      []EgressRule      `json:"egress,omitempty" bson:"egress,omitempty"`
}
type IngressRule struct {
	Ports []NetworkPolicyPort `json:"ports,omitempty" bson:"ports,omitempty"`                                //empty means all ports allowed for this rules
	From  []NetworkPolicyPeer `json:"from,omitempty" bson:"from,omitempty" protobuf:"bytes,2,rep,name=from"` //empty means all sources are allowed
}

type EgressRule struct {
	Ports []NetworkPolicyPort `json:"ports,omitempty" bson:"ports,omitempty"`                            //empty means all ports allowed for this rules
	To    []NetworkPolicyPeer `json:"to,omitempty" bson:"to,omitempty" protobuf:"bytes,2,rep,name=from"` //empty means all destination are allowed
}

type NetworkPolicyPort struct {
	Protocol *Protocol        `json:"protocol,omitempty" bson:"protocol,omitempty" protobuf:"bytes,1,opt,name=protocol,casttype=k8s.io/api/core/v1.Protocol"` //default is TCP
	Port     PortItntOrString `json:"port,omitempty" bson:"port,omitempty"`
}

type Protocol string

const (
	ProtocolTCP  Protocol = "TCP"
	ProtocolUDP  Protocol = "UDP"
	ProtocolSCTP Protocol = "SCTP"
)

type PortItntOrString struct {
	PortNumber int32  `json:"port_number,omitempty" bson:"port_number,omitempty"  jsonschema:"minimum=0,maximum=65536"`
	PortName   string `json:"port_name,omitempty" bson:"port_name,omitempty"`
}

type IPBlock struct {
	CIDR   string   `json:"cidr" bson:"cidr" protobuf:"bytes,1,name=cidr"`
	Except []string `json:"except,omitempty" bson:"except,omitempty" protobuf:"bytes,2,rep,name=except"`
}

type NetworkPolicyPeer struct {
	PodSelector *LabelSelectorObj `json:"pod_selector,omitempty" bson:"pod_selector,omitempty" protobuf:"bytes,1,opt,name=podSelector"`

	NamespaceSelector *LabelSelectorObj `json:"namespace_selector,omitempty" bson:"namespace_selector,omitempty" protobuf:"bytes,2,opt,name=namespaceSelector"`

	IPBlock *IPBlock `json:"ip_block,omitempty" bson:"ip_block,omitempty" protobuf:"bytes,3,rep,name=ipBlock"` // If this field is set then neither of the other fields can be.

}
