package constants

var (
	LoggingURL          string
	IstioEngineURL      string
	KnativeEngineURL    string
	ServicePort         string
	KubernetesEngineURL string
	NotificationURL     string
)

const (
	SERVICE_NAME        = "istio-mesh-engine"
	LOGGING_ENDPOINT    = "/api/v1/logger"
	LOGGING_LEVEL_INFO  = "info"
	LOGGING_LEVEL_ERROR = "error"
	LOGGING_LEVEL_WARN  = "warn"
)
