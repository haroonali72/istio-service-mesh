package services

import "bitbucket.org/cloudplex-devs/microservices-mesh-engine/types"

type BuildService struct {
	types.ServiceBasicInfo `json:",inline" bson:",inline"`
	ServiceAttributes      *BuildServiceAttributes `json:"service_attributes"  bson:"service_attributes" binding:"required"`
}

type BuildServiceAttributes struct {
	GitConfigurations             GitConfigurations             `json:"git_configurations" binding:"required"`
	ImageRepositoryConfigurations ImageRepositoryConfigurations `json:"image_repository_configurations" binding:"required"`
}

type GitConfigurations struct {
	Url            string `json:"url"`
	Branch         string `json:"branch"`
	ProfileID      string `json:"profile_id"`
	ProfileName    string `json:"profile_name"`
	DockerFilePath string `json:"docker_file_path"`
}
