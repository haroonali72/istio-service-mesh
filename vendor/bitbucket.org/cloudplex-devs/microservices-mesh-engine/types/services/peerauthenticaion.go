package services

import "bitbucket.org/cloudplex-devs/microservices-mesh-engine/types"

type PeerAuthenticationService struct {
	types.ServiceBasicInfo `json:",inline" bson:",inline"`
	ServiceAttributes      *PeerAuthenticationAttribute `json:"service_attributes,omitempty"  bson:"service_attributes" binding:"required"`
}
type PeerAuthenticationAttribute struct {
	Labels  map[string]string `json:"labels,omitempty" bson:"labels,omitempty"`
	TlsMode TlsMode           `json:"tls_mode,omitempty" bson:"tls_mode,omitempty"`
}

type TlsMode string

const (
	STRICT     TlsMode = "STRICT"
	PERMISSIVE TlsMode = "PERMISSIVE"
	DISABLE    TlsMode = "DISABLE"
	UNSET      TlsMode = "UNSET"
)
