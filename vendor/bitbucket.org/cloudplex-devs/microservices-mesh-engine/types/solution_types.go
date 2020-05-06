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

type SolutionMesh struct {
	ClusterName     *string     `json:"cluster_name"`
	ProjectId       *string     `json:"project_id"`
	Credentials     interface{} `json:"credentials"`
	KubeCredentials struct {
		KubernetesURL      string `json:"url"`
		KubernetesUsername string `json:"username"`
		KubernetesPassword string `json:"password"`
	} `json:"kubernetes_credentials"`
	Solution *SolutionMeshRequest `json:"solution_info"`
}
type SolutionMeshRequest struct {
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

type SolutionTemplate struct {
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
	SolutionId *string `json:"solution_id" bson:"solution_id" binding:"required" valid:"alphanumspecial,length(6|30)~The name must contain between 6 and 30 characters,lowercase~lowercase alphanumeric characters are allowed,required~Solution Id is missing in request"`
	// ProjectId is the field which will refer which project is linked
	// with this solution
	// +optional
	ProjectID *string `json:"project_id" bson:"project_id" valid:"-"`
	// Overall Status of the Solution
	// Status will give you status of solution
	// Valid Status list Deployed/ Deployment Failed/ New
	// +optional
	Status *string `json:"status" bson:"status" valid:"-" default:"new"`
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
}

type SolutionMetadataByProject struct {
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
	SolutionId *string `json:"solution_id" bson:"solution_id" binding:"required" valid:"alphanumspecial,length(6|30)~The name must contain between 6 and 30 characters,lowercase~lowercase alphanumeric characters are allowed,required~Solution Id is missing in request"`
	// Overall Status of the Solution
	// Status will give you status of solution
	// Valid Status list Deployed/ Deployment Failed/ New
	// +optional
	Status *string `json:"status" bson:"status" valid:"-" default:"new"`
}
