package types

import (
	"time"
)

type Solution struct {
	ProjectId *string   `json:"project_id" binding:"required"`
	Package   *Packages `json:"solution" binding:"required"`
}
type Packages struct {
	Id             *string            `json:"_id"`
	Name           *string            `json:"name" binding:"required"`
	Version        *string            `json:"version"`
	UserId         *string            `json:"user_id"`
	RepositoryInfo interface{}        `json:"repository_info" `
	Services       []*ServiceTemplate `json:"services" binding:"required"`
}

//solutionMesh
//ApplicationMesh
type ApplicationDepRequest struct {
	InfrastructureId string              `json:"infrastructure_id"`
	Application      *ApplicationRequest `json:"solution_info"`
}

//SolutionMeshRequest
//ApplicationMeshRequest
type ApplicationRequest struct {
	Id             *string             `json:"_id"`
	Name           *string             `json:"name"`
	Version        *string             `json:"version"`
	UserId         *string             `json:"user_id"`
	CreationDate   *time.Time          `json:"creation_date"`
	RepositoryInfo interface{}         `json:"repository_info"`
	Services       *ServiceMeshRequest `json:"service,omitempty"`
}

type AWSCredentials struct {
	AccessKey string `json:"access_key"`
	SecretKey string `json:"access_secret"`
	Region    string `json:"region"`
}

type ApplicationTemplate struct {
	User string `json:"user" bson:"user"`
	//brownfield_infra_id will be specifically used for discovered application
	BrownfieldInfraId string   `json:"brownfield_infra_id" bson:"brownfield_infra_id"`
	InfraIds          []string `json:"infra_ids" bson:"infra_ids"`
	// Name Used to link with application/project
	// Cannot be updated.
	// valid regex is ^[ A-Za-z0-9_-]*$
	// valid length is 5-30 character
	// +mandatory
	Name *string `json:"name" bson:"name" binding:"required" valid:"alphanumspecial,length(5|30)~Application name must contain between 5 and 30 characters or application name is invalid. Valid regex is ^[ A-Za-z0-9_-]*$ ,required~Application Name is missing in request"`
	// ApplicationId used to link with application/project
	// Cannot be updated.
	// valid regex is ^[ A-Za-z0-9_-]*$
	// valid length is 6-30 character
	// +mandatory
	ApplicationId *string `json:"application_id" bson:"application_id" binding:"required" valid:"alphanumspecial,length(6|30)~The name must contain between 6 and 30 characters,lowercase~application_id is invalid. Valid regex is ^[ A-Za-z0-9_-]*$~Application Id is missing in request"`
	// version used to link multiple version of applications/project
	// Cannot be updated.
	// valid regex is ^[ A-Za-z0-9_-]*$
	// +mandatory
	Version *string `json:"version" bson:"version" valid:"matches(^[ A-Za-z0-9._-]),optional"`

	Description string `json:"description" bson:"description"`

	Tags []string `json:"tags"bson:"tags"`
	// Overall Status of the Solution
	// Status will give you status of solution
	// Valid Status list Deployed/ Deployment Failed/ New
	// +optional
	Status      *string          `json:"status" bson:"status" valid:"-" default:"new"`
	InfraStatus []ServicesStatus `json:"infra_status" bson:"infra_status"`
	// Services in the solution
	// services can be empty but atleast one services is required
	// during deployment
	// +optional
	Services []*ServiceTemplate `json:"services" bson:"services" binding:"required" valid:"-"`
	// auto populated key
	// +optional
	V int `json:"__v" bson:"__v"`
	// auto populated key
	// +optional
	CompanyID *string `json:"company_id" bson:"company_id" valid:"-"`
	// CreationTime is timestamp of when the solution is created
	// auto generated time
	// +optional
	CreationDate time.Time `json:"creation_date,omitempty" bson:"creation_date" valid:"-"`
}

type ApplicationMetadata struct {
	// Name Used to link with application/project
	// Cannot be updated.
	// valid regex is ^[ A-Za-z0-9_-]*$
	// valid length is 5-30 character
	// +mandatory
	Name *string `json:"name" bson:"name" binding:"required" valid:"alphanumspecial,length(5|30)~Solution name must contain between 4 and 30 characters,required~Solution Name is missing in request"`
	// SolutionId used to link with application/project
	// Cannot be updated.
	// valid regex is ^[ A-Za-z0-9_-]*$
	// valid length is 6-30 character
	// +mandatory
	TemplateId *string `json:"template_id" bson:"template_id" binding:"required" valid:"alphanumspecial,length(6|30)~The name must contain between 6 and 30 characters,lowercase~lowercase alphanumeric characters are allowed,required~Solution Id is missing in request"`
	// Overall Status of the Solution
	// Status will give you status of solution
	// Valid Status list Deployed/ Deployment Failed/ New
	// +optional
	Status *string `json:"status" bson:"status" valid:"-" default:"new"`

	// version used to link multiple version of applications/project
	// Cannot be updated.
	// valid regex is ^[ A-Za-z0-9_-]*$
	// +mandatory
	AllVersion  []string               `json:"versions,omitempty" bson:"versions" `
	VersionInfo map[string]VersionInfo `json:"version_info,omitempty" json:"version_info"`
}
type VersionInfo struct {
	Tags         []string  `json:"tags,omitempty"bson:"tags"`
	CreationDate time.Time `json:"creation_date,omitempty" bson:"creation_date" valid:"-"`
	Description  string    `json:"description" bson:"description"`
}
