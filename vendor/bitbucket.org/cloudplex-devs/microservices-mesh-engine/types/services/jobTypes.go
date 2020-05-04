package services

import (
	"bitbucket.org/cloudplex-devs/microservices-mesh-engine/types"
)

//Id                interface{}              `json:"_id,omitempty" bson:"_id" valid:"-"`
//	ServiceId         string                   `json:"service_id" bson:"service_id" binding:"required" valid:"alphanumspecial,length(4|30)~service_id must contain between 6 and 30 characters,lowercase~lowercase alphanumeric characters are allowed,required~service_id is missing in request"`
//	Name              string                   `json:"name"  bson:"name" binding:"required" valid:"alphanumspecial,length(3|30),lowercase~lowercase alphanumeric characters are allowed,required"`
//	Version           string                   `json:"version"  bson:"version"  binding:"required" valid:"alphanumspecial,length(1|10),lowercase~lowercase alphanumeric characters are allowed,required"`
//	ServiceType       constants.ServiceType    `json:"service_type"  bson:"service_type" valid:"-"`
//	ServiceSubType    constants.ServiceSubType `json:"service_sub_type" bson:"service_type" valid:"-"`
//	Namespace         string                   `json:"namespace" bson:"namespace" binding:"required" valid:"alphanumspecial,required"`
//	CompanyId         string                   `json:"company_id,omitempty" bson:"company_id"`
//	CreationDate      time.Time                `json:"creation_date,omitempty" bson:"creation_date" valid:"-"`
type JobService struct {
	types.ServiceBasicInfo `json:",inline" bson:",inline"`
	ServiceAttributes      *JobServiceAttribute `json:"service_attributes"  bson:"company_id" binding:"required"`
}
type JobServiceAttribute struct {
	CommonContainerAttributes `json:",inline,omitempty" bson:",inline,omitempty"`
	// Specifies the maximum desired number of pods the job should
	// run at any given time. The actual number of pods running in steady state will
	// be less than this number when ((.spec.completions - .status.successful) < .spec.parallelism),
	// i.e. when the work left to do is less than max parallelism.
	// More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
	// +optional
	Parallelism *Parallelism `json:"parallelism,omitempty"`
	// Specifies the desired number of successfully finished pods the
	// job should be run with.  Setting to nil means that the success of any
	// pod signals the success of all pods, and allows parallelism to have any positive
	// value.  Setting to 1 means that parallelism is limited to 1 and the success of that
	// pod signals the success of the job.
	// More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/
	// +optional
	Completions *Completions `json:"completions,omitempty"`
	// Specifies the duration in seconds relative to the startTime that the job may be active
	// before the system tries to terminate it; value must be positive integer
	// +optional
	ActiveDeadlineSeconds *ActiveDeadlineSeconds `json:"active_deadline_seconds,omitempty"`
	// Specifies the number of retries before marking this job failed.
	// Defaults to 6
	// +optional
	BackoffLimit            *BackoffLimit            `json:"backoff_limit,omitempty"`
	ManualSelector          *ManualSelector          `json:"manual_selector, omitempty"`
	TTLSecondsAfterFinished *TTLSecondsAfterFinished `json:"ttl_seconds_after_finished,omitempty"`
}

type Parallelism struct {
	Value int32 `json:"value,omitempty"`
}

type Completions struct {
	Value int32 `json:"value,omitempty"`
}

type BackoffLimit struct {
	Value int32 `json:"value,omitempty"`
}

type ManualSelector struct {
	Value bool `json:"value,omitempty"`
}

type TTLSecondsAfterFinished struct {
	Value int32 `json:"value,omitempty"`
}
