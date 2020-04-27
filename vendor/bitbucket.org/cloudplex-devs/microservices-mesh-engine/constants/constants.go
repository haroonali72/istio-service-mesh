package constants

type Logger string

var (
	LoggingURL                    string
	IstioEngineURL                string
	KnativeEngineURL              string
	IstioDeploymentEngineURL      string
	KubernetesEngineURL           string
	SolutionEngineURL             string
	KubernetesDeploymentEngineURL string
	LegacyEngineURL               string
	ServicePort                   string
	ClusterAPI                    string
	RedisUrl                      string
	RobinUrl                      string
	RbacURL                       string
	////Mongodb///
	Host             string
	UserName         string
	Password         string
	Auth             bool
	Database         string
	SubscriptionURL  string
	CACert           string
	ClientCert       string
	ClientPem        string
	D_Duck_Username  string
	D_Duck_Password  string
	D_Duck_ApiKey    string
	D_Duck_ApiSecret string
)

const (
	BASE_PATH    = ".temp-data/"
	SERVICE_NAME = "ms-mesh-engine"
	Component    = "Solution"
	//////////////////istio////////////////////////////////
	IstioServicePostEndpoint     = "/istioservicedeployer"
	IstioYamlToServiceEndpoint   = "/importservice"
	IstioServicePutEndpoint      = ""
	DeployIstio                  = "/api/v1/istio"
	GetIstioStatus               = "/api/v1/istio/status/{token}"
	GetIngressControllerEndpoint = "ksd/api/v1/kubeservice/istio-system/istio-ingressgateway/endpoint"
	/////////////////////////////////////////////////////////////////
	///////////////knative////////////////////////////
	KnativeServicePostEndpoint = "/api/v1/service"
	KnativeEventingEndpoint    = "/api/v1/eventing"
	KnativeBuildPostEndpoint   = "/api/v1/build"
	////////////////////////////////////////////////////////////////
	//////////////////logging/////////////////////////////////////
	AUDIT_TRAIL_ENDPOINT      = "/elephant/api/v1/audit/store"
	FRONTEND_LOGGING_ENDPOINT = "/elephant/api/v1/frontend/logging"
	//FRONTEND_LOGGING_ENDPOINT = ""
	BACKEND_LOGGING_ENDPOINT = "/elephant/api/v1/backend/logging"
	LOGGING_LEVEL_INFO       = "info"
	LOGGING_LEVEL_ERROR      = "error"
	LOGGING_LEVEL_WARN       = "warn"

	///////logger

	Backend_logging  Logger = "backendLogging"
	Frontend_logging Logger = "frontendLogging"
	Audit_Trails     Logger = "auditTrails"

	/////////////////////////////////////////////////////////////
	/////////kubernetes/////////////////
	KUBERNETES_GET_CREDENTIALS_ENDPOINT = "/api/v1/credentials/{envId}"
	KUBERNETES_MASTER_PORT              = "6443"
	//////////////////////////////////
	CLUSTER_GET_ENDPOINT = "/antelope/cluster/{cloud_provider}/status/"
	Solution_Endpoint    = "/solution/"

	//////////////////////////////////////////////////
	///////////////////Robin/////////////////////////
	ROBIN_DYNAMIC_DATA_API_ENDPOINT = "/api/v1/dynamicData"

	//////////////////////////////////////
	LEGACY_VM_EXTENSION_ENDPOINT = "/legacy/api/v1/vmextansion/setup/{projectId}"

	//database Collections/////
	Solution                   = "solutions"
	Services                   = "services"
	SolutionTemplate           = "templates"
	DynamicData                = "dynamicData"
	DynamicConfigurationSchema = "dynamicConfigurationSchema"
	DeletedServiceState        = "services_state"
	CustomerSolutionTemplate   = "customer_solution_templates"

	//---------k8s service database collection-------------  /////

	KubernetesServiceDataBase     = "kubernetes_service_template"
	SecretServiceDataBase         = "secret_service_template"
	RoleServiceDataBase           = "role_service_template"
	RoleBindingServiceDataBase    = "role__binding_service_template"
	ServiceAccountServiceDataBase = "service_account_template"
	ConfigMapServiceDataBase      = "configmap_template"
	ClusterRoleService            = "cluster_role_service_template"
	ClusterRoleBindingService     = "cluster_role_binding_service_template"
	HpaService                    = "hpa_service_template"
	//////////////////////////////////////////////////////////////////
	Deploymentservice            = "deployment_service"
	Statefulsetservice           = "statefulset_service"
	Daemonsetservice             = "daemonset_service"
	Jobservice                   = "job_service"
	CronJobservice               = "cron_job_service"
	StorageclassService          = "storage_class_service"
	PersistentVolumeService      = "persistent_volume_service"
	PersistentVolumeClaimService = "persistent_volume_claim_service"
	NetworkPolicyService         = "network_policy_service"
	DefaultResourcesDatabase     = "default_resources"
	ResourcesDatabase            = "resource"

	/////---------------Istio service database collection--------------------////////
	GatewayServiceDataBase   = "gateway_service_template"
	VirtualServiceDataBase   = "virtual_service_template"
	DestinationRulesDataBase = "destination_rules_template"
	PolicyDataBase           = "policy_template"
	ServiceEntryDatabase     = "service_entry_template"

	//------Dynamic Configurations constants---------------//
	ErrInterfaceToCurrentService  = "current service data either has some configurations issue or not properly formatted"
	ErrInterfaceToExecutedService = "executed services data either have some configurations issue or not properly formatted"
	ErrServiceNotExecuted         = "the service you are fetching data from is not executed or have some other issue"
	ErrSolutionConversion         = "solution schema has some issues"
	//-----------------------------------------------------//

	MODE_POST   = "post"
	MODE_GET    = "get"
	MODE_PATCH  = "patch"
	MODE_PUT    = "put"
	MODE_DELETE = "delete"
	//-----------------Rbac Configurations----------------//
	Rbac_Token_Info        = "/security/api/rbac/token/extract"
	Rbac_Verify_Credential = "/security/api/rbac/allowed?resource_id={resourceID}&resource_type={resourceType}&action={action}"
	Rbac_Delete_Policy     = "/security/api/rbac/policy?resource_id={resourceID}&resource_type={resourceType}"
	Rbac_Add_Policy        = "/security/api/rbac/policy"
	Rbac_List              = "/security/api/rbac/list?companyId={companyID}&resource_type={resourceType}"
	Rbac_Evaluate          = "/security/api/rbac/evaluate"
	Rbac_Subscription_Plan = "/security/api/rbac/company/plan?companyId={companyID}"

	PROJECT_SECRET_POST_ENDPOINT   = "/api/v1/project/{projectId}/solution/{solutionId}/secrets/{serviceId}"
	PROJECT_SECRET_GET_ENDPOINT    = PROJECT_SECRET_POST_ENDPOINT
	PROJECT_SECRET_PUT_ENDPOINT    = PROJECT_SECRET_GET_ENDPOINT
	PROJECT_SECRET_DELETE_ENDPOINT = PROJECT_SECRET_GET_ENDPOINT

	TEMPLATE_SECRET_POST_ENDPOINT   = "/api/v1/solution/{solutionId}/secrets/{serviceId}"
	TEMPLATE_SECRET_GET_ENDPOINT    = TEMPLATE_SECRET_POST_ENDPOINT
	TEMPLATE_SECRET_PUT_ENDPOINT    = TEMPLATE_SECRET_GET_ENDPOINT
	TEMPLATE_SECRET_DELETE_ENDPOINT = TEMPLATE_SECRET_GET_ENDPOINT
)

var ListOfIgnoredServicesInAppSize = []string{
	"node",
	"init_container",
	"volume",
	"configmap",
	"secrets",
}

const LetterBytes = "abcdefghijklmnopqrstuvwxyz"
