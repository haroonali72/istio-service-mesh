package services

import "bitbucket.org/cloudplex-devs/microservices-mesh-engine/types"

type CertificateService struct {
	types.ServiceBasicInfo `json:",inline" bson:",inline"`
	ServiceAttributes      *CertificateAttribute `json:"service_attributes,omitempty"  bson:"service_attributes" binding:"required"`
}

type CertificateAttribute struct {
	SecretName string    `json:"secret_name,omitempty" bson:"secret_name,omitempty"`
	CommonName string    `json:"common_name,omitempty" bson:"common_name,omitempty"`
	DnsNames   []string  `json:"dns_names,omitempty" bson:"dns_names,omitempty"`
	IssuerRef  IssuerRef `json:"issuer_ref,omitempty" bson:"issuer_ref,omitempty"`
}

type IssuerRef struct {
	Name string `json:"name,omitempty" bson:"name,omitempty"`
	Kind string `json:"kind,omitempty" bson:"kind,omitempty"`
}
