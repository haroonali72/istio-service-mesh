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

type Logger string

const (
	ServiceName = "istio-mesh-engine"

	//KSD
	KubernetesServicesDeployment = "/ksd/api/v1/solution"
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
