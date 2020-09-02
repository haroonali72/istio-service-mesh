package services

import "bitbucket.org/cloudplex-devs/microservices-mesh-engine/types"

type ServerlessService struct {
	types.ServiceBasicInfo `json:",inline" bson:",inline"`
	ServiceAttributes      *ServerlessServiceAttributes `json:"service_attributes"  bson:"service_attributes" binding:"required"`
}

type ServerlessServiceAttributes struct {
	Domains                       []Domain                      `json:"domains,omitempty"`
	LocalVisibilityCheck          bool                          `json:"local_visibility_check,omitempty"`
	EnvVars                       []EnvVars                     `json:"environment_variables,omitempty"`
	Ports                         []Ports                       `json:"ports,omitempty"`
	ImageRepositoryConfigurations ImageRepositoryConfigurations `json:"image_repository_configurations" binding:"required"`
	Label                         string                        `json:"label,omitempty"`
	MinimumScale                  int                           `json:"minimum_scale,omitempty"`
	MaximumScale                  int                           `json:"maximum_scale,omitempty"`
}

type Domain struct {
	DomainName string `json:"domain_name"`
	Label      string `json:"label"`
}

type Ports struct {
	Name     string `json:"name"`
	Port     int    `json:"port"`
	Protocol string `json:"protocol"`
}
type EnvVars struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
