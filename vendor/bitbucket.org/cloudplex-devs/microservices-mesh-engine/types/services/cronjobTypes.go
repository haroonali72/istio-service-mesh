package services

import "time"

type CronJobService struct {
	Id                interface{}              `json:"_id,omitempty" bson:"_id" valid:"-"`
	ServiceId         string                   `json:"service_id" bson:"service_id" binding:"required" valid:"alphanumspecial,length(4|30)~service_id must contain between 6 and 30 characters,lowercase~lowercase alphanumeric characters are allowed,required~service_id is missing in request"`
	Name              string                   `json:"name"  bson:"name" binding:"required" valid:"alphanumspecial,length(3|30),lowercase~lowercase alphanumeric characters are allowed,required"`
	Version           string                   `json:"version"  bson:"version"  binding:"required" valid:"alphanumspecial,length(1|10),lowercase~lowercase alphanumeric characters are allowed,required"`
	ServiceType       string                   `json:"service_type"  bson:"service_type" valid:"-"`
	ServiceSubType    string                   `json:"service_sub_type" bson:"service_type" valid:"-"`
	Namespace         string                   `json:"namespace" bson:"namespace" binding:"required" valid:"alphanumspecial,required"`
	CompanyId         string                   `json:"company_id,omitempty" bson:"company_id"`
	CreationDate      time.Time                `json:"creation_date,omitempty" bson:"creation_date" valid:"-"`
	ServiceAttributes *CronJobServiceAttribute `json:"service_attributes"  bson:"company_id" binding:"required"`
}
type CronJobServiceAttribute struct {
	MeshConfig                 *IstioConfig                `json:"istio_config,omitempty"`
	LabelSelector              *LabelSelectorObj           `json:"label_selector,omitempty"`
	NodeSelector               map[string]string           `json:"node_selector"`
	Labels                     map[string]string           `json:"labels,omitempty"`
	Annotations                map[string]string           `json:"annotations,omitempty"`
	IsRbac                     bool                        `json:"is_rbac_enabled"`
	RbacRoles                  []K8sRbacAttribute          `json:"roles,omitempty"`
	CronJobScheduleString      string                      `json:"schedule"`
	IstioRoles                 []IstioRbacAttribute        `json:"istio_roles,omitempty"`
	IsInitContainerEnable      bool                        `json:"enable_init,omitempty"`
	StartingDeadLineSeconds    *StartingDeadlineSeconds    `json:"starting_deadline_seconds,omitempty"`
	ConcurrencyPolicy          *ConcurrencyPolicy          `json:"concurrency_policy,omitempty"`
	Suspend                    *Suspend                    `json:"suspend,omitempty"`
	FailedJobsHistoryLimit     *FailedJobsHistoryLimit     `json:"failed_jobs_history_limit,omitempty"`
	SuccessfulJobsHistoryLimit *SuccessfulJobsHistoryLimit `json:"successfulJ_jobs_history_limit,omitempty"`
	JobTemplate                *JobServiceAttribute        `json:"job_template,omitempty"`
}

type StartingDeadlineSeconds struct {
	Value int64 `json:"value,omitempty"`
}
type Suspend struct {
	Value bool `json:"value,omitempty"`
}
type SuccessfulJobsHistoryLimit struct {
	Value int32 `json:"value,omitempty"`
}
type FailedJobsHistoryLimit struct {
	Value int32 `json:"value,omitempty"`
}

type ConcurrencyPolicy string

const ConcurrencyPolicyAllow ConcurrencyPolicy = "Allow"
const ConcurrencyPolicyForbid ConcurrencyPolicy = "Forbid"
const ConcurrencyPolicyReplace ConcurrencyPolicy = "Replace"
