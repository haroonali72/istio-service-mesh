package services

import "bitbucket.org/cloudplex-devs/microservices-mesh-engine/types"

//Id                interface{}              `json:"_id,omitempty" bson:"_id" valid:"-"`
//ServiceId         string                   `json:"service_id" bson:"service_id" binding:"required" valid:"alphanumspecial,length(4|30)~service_id must contain between 6 and 30 characters,lowercase~lowercase alphanumeric characters are allowed,required~service_id is missing in request"`
//Name              string                   `json:"name"  bson:"name" binding:"required" valid:"alphanumspecial,length(3|30),lowercase~lowercase alphanumeric characters are allowed,required"`
//Version           string                   `json:"version"  bson:"version"  binding:"required" valid:"alphanumspecial,length(1|10),lowercase~lowercase alphanumeric characters are allowed,required"`
//ServiceType       constants.ServiceType    `json:"service_type"  bson:"service_type" valid:"-"`
//ServiceSubType    constants.ServiceSubType `json:"service_sub_type" bson:"service_type" valid:"-"`
//Namespace         string                   `json:"namespace" bson:"namespace" binding:"required" valid:"alphanumspecial,required"`
//CompanyId         string                   `json:"company_id,omitempty" bson:"company_id"`
//CreationDate      time.Time                `json:"creation_date,omitempty" bson:"creation_date" valid:"-"`
type CronJobService struct {
	types.ServiceBasicInfo `json:",inline" bson:",inline"`
	ServiceAttributes      *CronJobServiceAttribute `json:"service_attributes"  bson:"company_id" binding:"required"`
}
type CronJobServiceAttribute struct {
	CommonContainerAttributes `json:",inline,omitempty" bson:",inline,omitempty"`
	// The schedule in Cron format, see https://en.wikipedia.org/wiki/Cron.
	CronJobScheduleString string `json:"schedule" bson:"schedule" binding:"required"`
	// Optional deadline in seconds for starting the job if it misses scheduled
	// time for any reason.  Missed jobs executions will be counted as failed ones.
	// +optional
	StartingDeadLineSeconds *StartingDeadlineSeconds `json:"starting_deadline_seconds,omitempty"`
	// Specifies how to treat concurrent executions of a Job.
	// Valid values are:
	// - "Allow" (default): allows CronJobs to run concurrently;
	// - "Forbid": forbids concurrent runs, skipping next run if previous run hasn't finished yet;
	// - "Replace": cancels currently running job and replaces it with a new one
	// +optional
	ConcurrencyPolicy *ConcurrencyPolicy `json:"concurrency_policy,omitempty"`
	// This flag tells the controller to suspend subsequent executions, it does
	// not apply to already started executions.  Defaults to false.
	// +optional
	Suspend *Suspend `json:"suspend,omitempty"`
	// The number of failed finished jobs to retain.
	// This is a pointer to distinguish between explicit zero and not specified.
	// Defaults to 1.
	// +optional
	FailedJobsHistoryLimit *FailedJobsHistoryLimit `json:"failed_jobs_history_limit,omitempty" default:"1"`
	// The number of successful finished jobs to retain.
	// This is a pointer to distinguish between explicit zero and not specified.
	// Defaults to 3.
	// +optional
	SuccessfulJobsHistoryLimit *SuccessfulJobsHistoryLimit `json:"successfulJ_jobs_history_limit,omitempty" default:"3"`
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
