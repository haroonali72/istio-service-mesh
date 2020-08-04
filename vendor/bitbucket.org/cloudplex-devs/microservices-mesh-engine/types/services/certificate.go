package services

import "bitbucket.org/cloudplex-devs/microservices-mesh-engine/types"

type CertificateService struct {
	types.ServiceBasicInfo `json:",inline" bson:",inline"`
	ServiceAttributes      *CertificateAttribute `json:"service_attributes,omitempty"  bson:"service_attributes" binding:"required"`
}

type CertificateAttribute struct {
	SecretName string    `json:"secret_name" bson:"secret_name" binding:"required"`
	CommonName string    `json:"common_name" bson:"common_name" binding:"required"`
	DnsNames   []string  `json:"dns_names" bson:"dns_names" binding:"required"`
	IssuerRef  IssuerRef `json:"issuer_ref" bson:"issuer_ref" binding:"required"`
}

type IssuerRef struct {
	Name string `json:"name,omitempty" bson:"name,omitempty" binding:"required"`
	Kind string `json:"kind,omitempty" bson:"kind,omitempty" binding:"required"`
}
