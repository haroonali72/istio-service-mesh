package services

import "bitbucket.org/cloudplex-devs/microservices-mesh-engine/types"

type GithubEventing struct {
	types.ServiceBasicInfo `json:",inline" bson:",inline"`
	ServiceAttributes      GithubEventAttribute `json:"service_attributes,omitempty" bson:"service_attributes,omitempty"`
}

type GithubEventAttribute struct {
	GithubEventConfig GithubEventConfig `json:"github_event_config,omitempty" bson:"github_event_config,omitempty"`
}

type GithubEventConfig struct {
	Username    string   `json:"username,omitempty" bson:"username"`
	Repository  string   `json:"repository,omitempty" bson:"repository"`
	AccessToken string   `json:"access_token,omitempty" bson:"access_token,omitempty"`
	SecretKey   string   `json:"secret_key,omitempty" bson:"secret_key,omitempty"`
	Events      []string `json:"events,omitempty" bson:"events,omitempty"`
}
