package constants

const (
	//------------------Services------------------------------------//

	//---------k8s service type -----------------/////
	DeploymentServiceType         = "deployment"
	StatefulSetServiceType        = "statefulset"
	DaemonSetServiceType          = "daemonset"
	JobServiceType                = "job"
	CronJobServiceType            = "cronjob"
	KubernetesServiceType         = "service"
	RoleServiceType               = "role"
	RoleBindingServiceType        = "role_binding"
	ClusterRoleServiceType        = "cluster_role"
	ClusterRoleBindingServiceType = "cluster_role_binding"
	ServiceAccountServiceType     = "service_account"
	Resources                     = "resources"
	SecretServiceType             = "secret"
	ConfigMapServiceType          = "config_map"
	PVCServiceType                = "persistent_volume_claim"
	PVServiceType                 = "persistent_volume"
	StorageClassServiceType       = "storage_class"
	NetworkPolicyServiceType      = "network_policy"
	HpaServiceType                = "hpa"

	/////---------------Istio service type---------------//////

	GatewayServiceType   = "gateway"
	VirtualServiceType   = "virtual_service"
	DestinationRulesType = "destination_rule"
	ServiceEntryType     = "service_entry"
	PolicyType           = "policy"
)
