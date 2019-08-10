package constants

var (
	LoggingURL          string
	IstioEngineURL      string
	KnativeEngineURL    string
	ServicePort         string
	KubernetesEngineURL string
	NotificationURL     string
)

type K8sKind string

const (
	SERVICE_NAME = "istio-mesh-engine"
	//Logging
	LOGGING_ENDPOINT         = "/api/v1/logger"
	LOGGING_LEVEL_INFO       = "info"
	LOGGING_LEVEL_ERROR      = "error"
	LOGGING_LEVEL_WARN       = "warn"
	BACKEND_LOGGING_ENDPOINT = "/elephant/api/v1/backend/logging"
	VAULT_BACKEND = "http://robin:8093/robin/api/v1/template/docker/credentials/"

	//Kubernetes Component
	Deployment  K8sKind = "Deployment"
	CronJob     K8sKind = "CronJob"
	Job         K8sKind = "Job"
	StatefulSet K8sKind = "StatefulSet"
	Service     K8sKind = "service"
	ConfigMap   K8sKind = "ConfigMap"
	Secret      K8sKind = "Secret"
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
)
