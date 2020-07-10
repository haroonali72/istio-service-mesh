package services

import "bitbucket.org/cloudplex-devs/microservices-mesh-engine/types"

type ApiService struct {
	types.ServiceBasicInfo `json:",inline" bson:",inline"`
	ServiceAttributes      *types.APIService `json:"service_attributes"  bson:"service_attributes" binding:"required"`
}
