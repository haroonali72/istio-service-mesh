package types

type SolutionInfo struct {
	SolutionName string
	SolutionID   string
	ProjectID    string
	CompanyID    string
}

type DynamicConfigurationSolutionSchema struct {
	DoumentID       interface{}            `json:"_id" bson:"_id"`
	ApplicationId   interface{}            `json:"application_id" bson:"application_id"`
	ApplicationName interface{}            `json:"application_name" bson:"application_name"`
	ServicesData    map[string]interface{} `json:"solution_data" bson:"solution_data"`
}

type DynamicDataSolution struct {
	ApplicationId string                 `json:"application_id" bson:"application_id"`
	CompnayId     string                 `json:"company_id" bson:"company_id"`
	ServicesData  map[string]interface{} `json:"solution_data" bson:"solution_data"`
}

type LoadBalancerAPIResponse struct {
	Error       string `json:"error"`
	External_IP string `json:"external_ip"`
}

type CurrentServiceData struct {
	Key          string `json:"key" bson:"key"`
	Value        string `json:"value" bson:"value"`
	Dynamic      bool   `json:"dynamic" bson:"dynamic"`
	Type         string `json:"type" bson:"type"`
	LoadBalancer *bool  `json:"loadbalancer"`
}

type APISchema struct {
	ProjectId      string      `json:"project_id" bson:"project_id"`
	ServiceId      string      `json:"service_id" bson:"service_id"`
	ProjectService string      `json:"unique_id" bson:"unique_id"`
	ApiSchema      interface{} `json:"api_schema" bson:"api_schema"`
}

type APIServiceData struct {
	ServiceName string
	ServiceID   string
	ServiceData map[string]interface{}
}

type CurrentService struct {
	ServiceName string
	ServiceID   string
	ServiceData map[string]interface{}
}

type ExecutedServices struct {
	ExecutedServicesData map[string]interface{}
}

type ServiceAttribute struct {
	ServiceName string `json:"name" bson:"name"`
}
