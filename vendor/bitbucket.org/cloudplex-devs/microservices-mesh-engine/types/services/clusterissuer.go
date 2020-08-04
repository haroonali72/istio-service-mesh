package services

import "bitbucket.org/cloudplex-devs/microservices-mesh-engine/types"

type ClusterIssuerService struct {
	types.ServiceBasicInfo `json:",inline" bson:",inline"`
	ServiceAttributes      *ClusterIssuerAttribute `json:"service_attributes,omitempty"  bson:"service_attributes" binding:"required"`
}

type ClusterIssuerAttribute struct {
	Email string `json:"email" bson:"email" binding:"required"`
}
