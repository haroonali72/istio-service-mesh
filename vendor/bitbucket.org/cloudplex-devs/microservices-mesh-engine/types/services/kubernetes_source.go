package services

import "bitbucket.org/cloudplex-devs/microservices-mesh-engine/types"

type KubernetesSourceService struct {
	types.ServiceBasicInfo `json:",inline" bson:",inline"`
	ServiceAttributes      KubernetesSourceAttribute `json:"service_attributes,omitempty" bson:"service_attributes,omitempty"`
}

type KubernetesSourceAttribute struct {
	KubeEventConfig KubeEventConfig `json:"kube_event_config,omitempty" bson:"kube_event_config,omitempty"`
}

type KubeEventConfig struct {
	Resources []string `json:"resources,omitempty" bson:"resources,omitempty"`
	Verbs     []string `json:"verbs,omitempty" bson:"verbs,omitempty"`
}
