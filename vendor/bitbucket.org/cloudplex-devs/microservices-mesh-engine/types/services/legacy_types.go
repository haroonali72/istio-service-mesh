package services

import "bitbucket.org/cloudplex-devs/microservices-mesh-engine/types"

type LegacyService struct {
	types.ServiceBasicInfo `json:",inline" bson:",inline"`
	ServiceAttributes      *LegacyServiceAttributes `json:"service_attributes"  bson:"service_attributes" binding:"required"`
}
type LegacyServiceAttributes struct {
	IP       string        `json:"ip" bson:"ip"`
	Ports    []PortsAttrib `json:"ports" bson:"ports"`
	FileName string        `json:"file_name" bson:"file_name"`
	URL      string        `json:"url" bson:"url"`
}
type PortsAttrib struct {
	Protocol string `json:"protocol"`
	Number   int32  `json:"number"`
	Name     string `json:"name"`
}
