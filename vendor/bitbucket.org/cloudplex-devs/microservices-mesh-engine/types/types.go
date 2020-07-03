package types

import (
	"bitbucket.org/cloudplex-devs/microservices-mesh-engine/constants"
	"time"
)

type HelmHookImport struct {
	MetaData struct {
		Annotations struct {
			Type   *string `yaml:"helm.sh/hook,omitempty"`
			Weight string  `yaml:"helm.sh/hook-weight,omitempty"`
		} `yaml:"annotations"`
	} `yaml:"metadata"`
}

type APIError struct {
	ErrorCode    int
	ErrorMessage string
	CreatedAt    time.Time
}

type HTTPError struct {
	Code    int    `json:"code" example:"401"`
	Message string `json:"error" example:"unauthorized"`
}

type Status struct {
	// Status of the operation.
	// One of: "Success" or "Failure".
	Status string `json:"status" bson:"status"`
	// A human-readable description of the status of this operation.
	Message string `json:"message"`
}

type ResponseData struct {
	StatusCode int         `json:"status_code"`
	Body       interface{} `json:"body"`
	Error      error       `json:"error"`
	Status     string      `json:"status"`
}

type LoggingRequest struct {
	UserID      string      `json:"user_id"`
	CompanyId   string      `json:"company"`
	Message     interface{} `json:"message"`
	Id          string      `json:"id"`
	Environment string      `json:"environment"`
	Service     string      `json:"service"`
	Level       string      `json:"level"`
	Type        string      `json:"type"`
	PayLoad     string      `json:"payload"`
}

type ServicesResponse struct {
	Error string      `json:"error"`
	Data  ServiceResp `json:"service"`
}
type ServiceResp struct {
	ID               string   `json:"_id"`
	ServiceId        string   `json:"service_id"`
	Name             string   `json:"name"`
	StatusIndividual []string `json:"status_individual"`
	Status           string   `json:"status"`
	Reason           string   `json:"reason"`
	Namespace        string   `json:"namespace"`
	ServiceSubType   string   `json:"service_sub_type"`
	PodErrors        []string `json:"pod_errors"`
}

type LogRequest struct {
	//project Id of the project
	ProjectId string `json:"project_id"`
	// requester's service name (antelope, msme,raccoon)
	ServiceName string `json:"service_name"`
	//log severity (info,warn,debug,error)
	Severity string `json:"severity"`
	//requested user of platform (haseeb@cloudplex.io)
	UserId string `json:"user_id"`
	//company of the requested user (cloudplex.io)
	Company string `json:"company_id"`
	//message type (stdout,stderr)
	MessageType string `json:"message_type"`
	//actual message
	Message            interface{} `json:"message" `
	LoggingHttpRequest `json:"http_request"`
}
type LoggingHttpRequest struct {
	RequestId string `json:"request_id"`
	//url of the cloudplex server (e.g. apis.cloudplex.cf)
	Url string `json:"url"`
	//request method (GET/POST/PUT/PATCH/DELETE)
	Method string `json:"method" `
	//request path of backend service
	Path string `json:"path"`
	//request body
	Body string `json:"body"`
	//status code of service
	Status int `json:"status"`
}

type AuditTrailRequest struct {
	//Type of log e.g. audit-trail, backend-logging
	LogName constants.Logger `json:"log_name"`
	//project Id of the project
	ProjectId string `json:"project_id"`
	//resource name like project,network,cluster
	ResourceName string `json:"resource_name"`
	// requester's service name (antelope, msme,raccoon)
	ServiceName string `json:"service_name" binding:"required"`
	//log severity (info,warn,debug,error)
	Severity string `json:"severity" binding:"required"`
	//requested user of platform (haseeb@cloudplex.io)
	UserId string `json:"user_id" binding:"required"`
	//company of the requested user (cloudplex.io)
	Company string `json:"company_id" binding:"required"`
	//message type (stdout,stderr)
	MessageType string `json:"message_type"`
	//Response from actual service when/where log was generated
	Response interface{} `json:"response"`
	//actual message
	Message      interface{} `json:"message" binding:"required" `
	Http_Request struct {
		Request_Id string `json:"request_id"`
		//url of the cloudplex server (e.g. apis.cloudplex.cf)
		Url string `json:"url"`
		//request method (GET/POST/PUT/PATCH/DELETE)
		Method string `json:"method" `
		//request path of backend service
		Path string `json:"path"`
		//request body
		Body string `json:"body"`
		//status code of service
		Status int `json:"status"`
	} `json:"http_request"  binding:"required"`
}

/*type RBACSubscriptionResponse struct {
	Company struct {
		ActivePlan struct {
			Limits struct {
				AppCount *int `json:"appCount"`
				AppSize  *int `json:"APP_SIZE"`
			} `json:"limits"`
		} `json:"activePlan"`
	} `json:"company"`
}*/
type RBACSubscriptionResponse struct {
	Data []struct {
		ID         string `json:"_id"`
		Name       string `json:"name"`
		ActivePlan struct {
			Limits struct {
				AppCount   int `json:"appCount"`
				MaxAppSize int `json:"maxAppSize"`
			} `json:"limits"`
		} `json:"activePlan"`
	} `json:"data"`
}
type LoggingRequestFrontend struct {
	Message     interface{} `json:"message"`
	Id          string      `json:"id"`
	Environment string      `json:"environment"`
	Service     string      `json:"service"`
	Level       string      `json:"level"`
	//level is info erro
	CompanyId string `json:"company_id"`
	UserId    string `json:"userId"`
	Type      string `json:"type"`
}

type GitImport struct {
	ImportType string `json:"import_type"`
	URL        string `json:"url"`
}

type GitExport struct {
	ExportType string `json:"export_type"`
	ProfileId  string `json:"profile_id"`
	URL        string `json:"url"`
}

type GitMain struct {
	ProfileName string      `json:"profile_name" valid:"required"`
	ServiceId   string      `json:"service_id" valid:"required"`
	GitCreds    interface{} `json:"git_credentials"`
}
