package services

import "bitbucket.org/cloudplex-devs/microservices-mesh-engine/types"

type GcpPubSubEventing struct {
	types.ServiceBasicInfo `json:",inline" bson:",inline"`
	ServiceAttributes      GcpPubSubEventAttribute `json:"service_attributes,omitempty" bson:"service_attributes,omitempty"`
}
type GcpPubSubEventAttribute struct {
	GcpPubSubEventConfig GcpPubSubEventConfig `json:"gcppubsub_event_config,omitempty" bson:"gcppubsub_event_config,omitempty"`
}

type GcpPubSubEventConfig struct {
	Topic         string `json:"topic,omitempty" bson:"topic,omitempty"`
	GcpServiceKey string `json:"gcp_service_key" bson:"gcp_service_key,omitempty"`
}
