package services

import "bitbucket.org/cloudplex-devs/microservices-mesh-engine/types"

type ServerlessService struct {
	types.ServiceBasicInfo `json:",inline" bson:",inline"`
	ServiceAttributes      *ServerlessServiceAttributes `json:"service_attributes"  bson:"service_attributes" binding:"required"`
}

type ServerlessServiceAttributes struct {
	//Address                       []interface{}                 `json:"address" `
	Domains                       []Domain                      `json:"domains"`
	LocalVisibilityCheck          bool                          `json:"local_visibility_check"`
	EnvVars                       []EnvVars                     `json:"environment_variables"`
	Ports                         []Ports                       `json:"ports"`
	ImageRepositoryConfigurations ImageRepositoryConfigurations `json:"image_repository_configurations" binding:"required"`
	Label                         string                        `json:"label"`
	MinimumScale                  int                           `json:"minimum_scale"`
	MaximumScale                  int                           `json:"maximum_scale"`
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
