package types

type SolutionInfo struct {
	SolutionName string
	SolutionID   string
	ProjectID    string
	CompanyID    string
}

type DynamicConfigurationSolutionSchema struct {
	DoumentID    interface{}            `json:"_id" bson:"_id"`
	SolutionId   interface{}            `json:"solution_id" bson:"solution_id"`
	SolutionName interface{}            `json:"solution_name" bson:"solution_name"`
	ServicesData map[string]interface{} `json:"solution_data" bson:"solution_data"`
}

type DynamicDataSolution struct {
	SolutionId   string                 `json:"solution_id" bson:"solution_id"`
	ProjectId    string                 `json:"project_id" bson:"project_id"`
	CompnayId    string                 `json:"company_id" bson:"company_id"`
	ServicesData map[string]interface{} `json:"solution_data" bson:"solution_data"`
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
