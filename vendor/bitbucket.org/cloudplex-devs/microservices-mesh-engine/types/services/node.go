package services

import "bitbucket.org/cloudplex-devs/microservices-mesh-engine/types"

type NodeService struct {
	types.ServiceBasicInfo `json:",inline" bson:",inline"`
	ServiceAttributes      *NodeServiceAttributes `json:"service_attributes"  bson:"service_attributes" binding:"required"`
}

type NodeServiceAttributes struct {
	NodePool []string          `json:"nodepool" bson:"nodepool" binding:"required"`
	Labels   map[string]string `json:"nodelabel" bson:"nodelabel" binding:"required"`
}
