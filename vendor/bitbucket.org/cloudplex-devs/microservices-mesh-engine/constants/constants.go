package constants

type Logger string
type Publisher string

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
	RabbitMQURL                   string
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
	RabbitMqUserName string
	RabbitMqPassword string
)

const (
	AuthTokenKey      = "X-Auth-Token"
	BASE_PATH         = ".temp-data/"
	SERVICE_NAME      = "ms-mesh-engine"
	Component         = "Application"
	RabbitMQWorkQueue = "bb.app.work"
	RabbitMqDoneQueue = "bb.app.done"

	CICDWorkQueue   = "cd.app.work"
	CICDMqDoneQueue = "cd.app.done"
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

	TrantulaServingEndpoint  = "/kube/api/v1/install/serving/{PROJECT_ID}"
	TrantulaEventingEndpoint = "/kube/api/v1/install/eventing/{PROJECT_ID}"
	TrantulaBuildEndpoint    = "/kube/api/v1/install/tekton/{PROJECT_ID}"
	////////////////////////////////////////////////////////////////
	//////////////////logging/////////////////////////////////////
	AuditTrailEndpoint      = "/elephant/api/v1/audit/store"
	FrontendLoggingEndpoint = "/elephant/api/v1/frontend/logging"
	//FRONTEND_LOGGING_ENDPOINT = ""
	BackendLoggingEndpoint = "/elephant/api/v1/backend/logging"
	LoggingLevelInfo       = "info"
	LoggingLevelError      = "error"
	LOGGING_LEVEL_WARN     = "warn"

	// Publisher Type Current supported Types are redis, rabbitmq
	RedisPublisher    Publisher = "redis"
	RabbitMqPublisher Publisher = "rabbitmq"
	CICDPublisher     Publisher = "cicd_rabbitmq"
	///////logger

	BackendLogging  Logger = "backendLogging"
	FrontendLogging Logger = "frontendLogging"
	AuditTrails     Logger = "auditTrails"

	/////////////////////////////////////////////////////////////
	/////////kubernetes/////////////////
	KubernetesGetCredentialsEndpoint = "/api/v1/credentials/{envId}"
	KubernetesMasterPort             = "6443"
	//////////////////////////////////
	ClusterGetEndpoint = "/antelope/cluster/{cloud_provider}/status/"
	SolutionEndpoint   = "/solution/"

	//////////////////////////////////////////////////
	///////////////////Robin/////////////////////////
	RobinDynamicDataApiEndpoint = "/api/v1/dynamicData"

	//////////////////////////////////////
	LegacyVmExtensionEndpoint = "/legacy/api/v1/vmextansion/setup/{projectId}"

	//database Collections/////
	Solution                   = "solutions"
	Services                   = "services"
	SolutionTemplate           = "templates"
	DynamicData                = "dynamicData"
	DynamicConfigurationSchema = "dynamicConfigurationSchema"
	DeletedServiceState        = "services_state"
	CustomerSolutionTemplate   = "customer_solution_templates"
	HelmHubChartInfo           = "helm_hub_chart_info"

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
	Podservice                   = "pod_service"
	Statefulsetservice           = "statefulset_service"
	Daemonsetservice             = "daemonset_service"
	Jobservice                   = "job_service"
	CronJobservice               = "cron_job_service"
	StorageclassService          = "storage_class_service"
	PersistentVolumeService      = "persistent_volume_service"
	PersistentVolumeClaimService = "persistent_volume_claim_service"
	NetworkPolicyService         = "network_policy_service"
	Legacyservice                = "legacy_service"
	DefaultResourcesDatabase     = "default_resources"
	ResourcesDatabase            = "resource"
	Serverlessservice            = "serverless_service"
	Buildservice                 = "build_service"
	APIservice                   = "api_service"
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
	///--------------------- Eventing service database collection ----------------///

	KubernetesSourceDatabase = "kubernetes_source_template"
	GithubSourceDatabase     = "github_source_template"
	GcpPubSubSourceDatabase  = "gcppubsub_source_template"

	///-----------------------Node Service Collection-------------------------------/////
	NodeServiceDataBase = "node_template"

	ModePost   = "post"
	ModeGet    = "get"
	ModePatch  = "patch"
	ModePut    = "put"
	ModeDelete = "delete"
	//-----------------Rbac Configurations----------------//
	Rbac_Token_Info         = "/security/api/rbac/token/extract"
	Rbac_Verify_Credential  = "/security/api/rbac/allowed?resource_id={resourceID}&resource_type={resourceType}&action={action}"
	Rbac_Delete_Policy      = "/security/api/rbac/policy?resource_id={resourceID}&resource_type={resourceType}"
	Rbac_Add_Policy         = "/security/api/rbac/policy"
	Rbac_Allowed_Namespaces = "/security/api/rbac/application/namespaces/allowed"
	Rbac_List               = "/security/api/rbac/list?companyId={companyID}&resource_type={resourceType}"
	Rbac_Evaluate           = "/security/api/rbac/evaluate"
	Rbac_Subscription_Plan  = "/security/api/rbac/companies"

	ProjectSecretPostEndpoint   = "/api/v1/application/{applicationId}/secrets/{serviceId}"
	ProjectSecretGetEndpoint    = ProjectSecretPostEndpoint
	ProjectSecretPutEndpoint    = ProjectSecretGetEndpoint
	ProjectSecretDeleteEndpoint = ProjectSecretGetEndpoint

	TemplateSecretPostEndpoint   = "/api/v1/secrettemplate/{templateId}/secrets/{serviceId}"
	TemplateSecretGetEndpoint    = TemplateSecretPostEndpoint
	TemplateSecretPutEndpoint    = TemplateSecretGetEndpoint
	TemplateSecretDeleteEndpoint = TemplateSecretGetEndpoint
)

var ListOfIgnoredServicesInAppSize = []string{
	"node",
	"init_container",
	"volume",
	"configmap",
	"secrets",
}
var ListOfServicesInAppSize = []ServiceSubType{
	Deployment,
	StatefulSet,
	DaemonSet,
	CronJob,
	Job,
	ServerlessService,
	PubSubEventing,
	PubSubGCPSource,
	LegacyService,
}

const LetterBytes = "abcdefghijklmnopqrstuvwxyz"

type Hook string

const (
	PreInstallHook   Hook = "pre-install"
	PostInstallHook  Hook = "post-install"
	PreUpgradeHook   Hook = "pre-upgrade"
	PostUpgradeHook  Hook = "post-upgrade"
	PreDeleteHook    Hook = "pre-delete"
	PostDeleteHook   Hook = "post-delete"
	PreRollbackHook  Hook = "pre-rollback"
	PostRollbackHook Hook = "post-rollback"
)
