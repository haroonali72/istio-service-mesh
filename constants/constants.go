package constants

var (
	LoggingURL          string
	IstioEngineURL      string
	KnativeEngineURL    string
	ServicePort         string
	KubernetesEngineURL string
	K8sEngineGRPCURL    string
	NotificationURL     string
	VaultURL            string
	RbacURL             string
)

type K8sKind string
type Logger string

const (
	SERVICE_NAME = "istio-mesh-engine"

	//KSD
	KUBERNETES_SERVICES_DEPLOYMENT = "/ksd/api/v1/solution"
	//Logging

	LOGGING_LEVEL_INFO  = "info"
	LOGGING_LEVEL_ERROR = "error"
	LOGGING_LEVEL_WARN  = "warn"

	BACKEND_LOGGING_ENDPOINT  = "/elephant/api/v1/backend/logging"
	FRONTEND_LOGGING_ENDPOINT = "/elephant/api/v1/frontend/logging"
	VAULT_BACKEND             = "/robin/api/v1/template/docker/credentials/"
	LOGGING_ENDPOINT          = "/api/v1/logger"
	//logger
	Backend_logging  Logger = "backendLogging"
	Frontend_logging Logger = "frontendLogging"

	/*//---------k8s service type -----------------/////
	DeploymentServiceType         = "deployment"
	StatefulSetServiceType        = "statefulSet"
	DaemonSetServiceType          = "daemonSet"
	JobServiceType                = "job"
	CronJobServiceType            = "cronJob"
	KubernetesServiceType         = "kubernetes"
	RoleServiceType               = "role"
	RoleBindingServiceType        = "role_binding"
	ClusterRoleServiceType        = "cluster_role"
	ClusterRoleBindingServiceType = "cluster_role_binding"
	ServiceAccountServiceType     = "service_account"
	Resources                     = "resources"
	SecretServiceType             = "secret"
	ConfigMapServiceType          = "config_map"
	PVCServiceType                = "pvc"
	PVServiceType                 = "pv"
	StorageClassServiceType       = "storage_class"
	NetworkPolicyServiceType      = "network_policy"
	HpaServiceType                = "hpa"

	/////---------------Istio service type---------------//////

	GatewayServiceType   = "gateway"
	VirtualServiceType   = "virtual"
	DestinationRulesType = "destination_rules"
	ServiceEntryType     = "service_entry"
	PolicyType           = "policy"

	/////---------------Istio service database collection--------------------////////
	GatewayServiceDataBase   = "gateway_service_template"
	VirtualServiceDataBase   = "virtual_service_template"
	DestinationRulesDataBase = "destination_rules_template"
	PolicyDataBase           = "policy_template"
	ServiceEntryDatabase     = "service_entry_template"*/

	//RBAc
	Rbac_Token_Info = "/security/api/rbac/token/info"
	Ksd_Get_Nobe    = "/all"
)

var ListOfNonKubernetesSupportedObjects = []string{
	"Gateway",
	"ServiceEntry",
	"DestinationRule",
	"VirtualService",
}
