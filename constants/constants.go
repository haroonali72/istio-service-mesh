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

	//Kubernetes Component
	Deployment  K8sKind = "Deployment"
	CronJob     K8sKind = "CronJob"
	Job         K8sKind = "Job"
	StatefulSet K8sKind = "StatefulSet"
	Service     K8sKind = "service"
	ConfigMap   K8sKind = "ConfigMap"
	Secret      K8sKind = "Secret"
	Daemonset   K8sKind = "daemonset"

	////Istio Components
	VirtualService  K8sKind = "VritualService"
	Gateway         K8sKind = "gateway"
	DestinationRule K8sKind = "DestinationRule"
	Policy          K8sKind = "Policy"
	ServiceEntry    K8sKind = "ServiceEntry"

	//RBAC
	Role               K8sKind = "Role"
	RoleBinding        K8sKind = "RoleBinding"
	ClusterRole        K8sKind = "ClusterRole"
	ClusterRoleBinding K8sKind = "ClusterRoleBinding"
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
