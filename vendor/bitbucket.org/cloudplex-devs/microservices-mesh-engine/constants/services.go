package constants

import "fmt"

type ServiceType string
type ServiceSubType string

const (
	Kubernetes ServiceType = "k8s"
	Serverless ServiceType = "serverless"
	MeshType   ServiceType = "service_mesh"
	OtherType  ServiceType = "other"
	AWSType    ServiceType = "aws_managed"
	AzureType  ServiceType = "azure_managed"
	GCPType    ServiceType = "gcp_managed"
	DOType     ServiceType = "do_managed"
	PubSubType ServiceType = "eventing"
	BuildType  ServiceType = "build"

	//------------------Services------------------------------------//

	// k8s service sub type
	Deployment         ServiceSubType = "deployment"
	StatefulSet        ServiceSubType = "statefulset"
	DaemonSet          ServiceSubType = "daemonset"
	Job                ServiceSubType = "job"
	CronJob            ServiceSubType = "cronjob"
	KubernetesService  ServiceSubType = "service"
	Role               ServiceSubType = "role"
	RoleBinding        ServiceSubType = "role_binding"
	ClusterRole        ServiceSubType = "cluster_role"
	ClusterRoleBinding ServiceSubType = "cluster_role_binding"
	ServiceAccount     ServiceSubType = "service_account"
	Resources          ServiceSubType = "resources"
	Secret             ServiceSubType = "secrets"
	ConfigMap          ServiceSubType = "config_map"
	PVC                ServiceSubType = "persistent_volume_claim"
	PV                 ServiceSubType = "persistent_volume"
	StorageClass       ServiceSubType = "storage_class"
	NetworkPolicy      ServiceSubType = "network_policy"
	Hpa                ServiceSubType = "hpa"
	InitContainer      ServiceSubType = "init_container"
	Pod                ServiceSubType = "pod"
	// Istio service sub type

	Gateway            ServiceSubType = "gateway"
	VirtualService     ServiceSubType = "virtual_service"
	DestinationRule    ServiceSubType = "destination_rule"
	ServiceEntry       ServiceSubType = "service_entry"
	MeshPolicy         ServiceSubType = "policy"
	PeerAuthentication ServiceSubType = "peer_authentication"

	// Serverless Service Sub Types
	ServerlessService ServiceSubType = "serverless"
	// Other Service Sub Type
	LegacyService ServiceSubType = "legacy"
	ApiService    ServiceSubType = "api"

	// AWS Services Sub Types
	AWSS3           ServiceSubType = "s3"
	AWSRDS          ServiceSubType = "aws_rds"
	AWSRedshift     ServiceSubType = "aws_redshift"
	AWSElasticCache ServiceSubType = "aws_elastic_cache"
	AWSEMR          ServiceSubType = "aws_emr"
	AWSEC2          ServiceSubType = "aws_ec2"
	AWSDynamoDB     ServiceSubType = "aws_dynamo_db"
	AWSSQS          ServiceSubType = "aws_sqs"

	// Azure Service SubTypes
	AzureStorageAccount ServiceSubType = "azure_storage_account"
	AzureRedisCache     ServiceSubType = "azure_rediscache"
	AzureCompute        ServiceSubType = "azure_compute"
	AzureMariaDB        ServiceSubType = "azure_maria_db"
	AzureMySQL          ServiceSubType = "azure_my_sql"
	AzurePostgresSQL    ServiceSubType = "azure_postgres_sql"

	// GCP Service Sub Types

	GCPCloudStorage ServiceSubType = "gcp_cloud_storage"
	GCPCloudSQL     ServiceSubType = "gcp_cloud_sql"
	GCPCompute      ServiceSubType = "gcp_compute"
	GCPCloudMemory  ServiceSubType = "gcp_cloud_memory"
	GCPBigData      ServiceSubType = "gcp_big_data"

	// DO Services Sub Types
	DORedis       ServiceSubType = "do_redis"
	DOPostgresSQl ServiceSubType = "do_postgres_sql"
	DoMySQL       ServiceSubType = "do_my_sql"

	// PUBSUB Service SubTypes
	PubSubEventing  ServiceSubType = "eventing"
	PubSubGCPSource ServiceSubType = "gcppubsubsource"

	//Build Service SubType
	BuildService ServiceSubType = "build"

	//Docker Registries
	AWSDockerRegistry   ServiceSubType = "aws_registry"
	AzureDockerRegistry ServiceSubType = "azure_registry"
	GCPDockerRegistry   ServiceSubType = "gcp_registry"
	DockerRegistry      ServiceSubType = "docker_registry"
)

var (
	supportedServiceType = []ServiceType{
		Kubernetes,
		Serverless,
		MeshType,
		OtherType,
		AWSType,
		AzureType,
		GCPType,
		DOType,
		BuildType,
		PubSubType,
	}

	kubernetesServiceSubTypes = []ServiceSubType{
		Deployment,
		StatefulSet,
		DaemonSet,
		CronJob,
		Job,
		KubernetesService,
		Role,
		RoleBinding,
		ServiceAccount,
		ClusterRole,
		ClusterRoleBinding,
		PV,
		PVC,
		StorageClass,
		Secret,
		ConfigMap,
		Hpa,
		NetworkPolicy,
		InitContainer,
		Pod,
		//registries
		AWSDockerRegistry,
		AzureDockerRegistry,
		GCPDockerRegistry,
		DockerRegistry,
	}
	meshServicesSubTypes = []ServiceSubType{
		Gateway,
		VirtualService,
		DestinationRule,
		//ServiceEntry,
		MeshPolicy,
	}
	serverlessServicesSubTypes = []ServiceSubType{
		ServerlessService,
	}
	otherServicesSubTypes = []ServiceSubType{
		LegacyService,
		ApiService,
		ServiceEntry,
	}
	awsServiceSubTypes = []ServiceSubType{
		AWSS3,
		AWSDynamoDB,
		AWSEC2,
		AWSElasticCache,
		AWSEMR,
		AWSRDS,
		AWSRedshift,
		AWSSQS,
	}
	azureServiceSubTypes = []ServiceSubType{
		AzureCompute,
		AzureMariaDB,
		AzureMySQL,
		AzurePostgresSQL,
		AzureRedisCache,
		AzureStorageAccount,
	}
	gcpServiceSubTypes = []ServiceSubType{
		GCPBigData,
		GCPCloudMemory,
		GCPCloudSQL,
		GCPCloudStorage,
		GCPCompute,
	}
	doServiceSubTypes = []ServiceSubType{
		DOPostgresSQl,
		DORedis,
		DoMySQL,
	}
	pubSubServiceSubTypes = []ServiceSubType{
		PubSubEventing,
		PubSubGCPSource,
	}
	buildServiceSubTypes = []ServiceSubType{
		BuildService,
	}

	supportedServiceSubTypes []ServiceSubType

	mappedServices = map[ServiceType][]ServiceSubType{
		Kubernetes: kubernetesServiceSubTypes,
		Serverless: serverlessServicesSubTypes,
		MeshType:   meshServicesSubTypes,
		OtherType:  otherServicesSubTypes,
		AWSType:    awsServiceSubTypes,
		AzureType:  azureServiceSubTypes,
		GCPType:    gcpServiceSubTypes,
		DOType:     doServiceSubTypes,
		PubSubType: pubSubServiceSubTypes,
		BuildType:  buildServiceSubTypes,
	}
)

func init() {
	supportedServiceSubTypes = append(kubernetesServiceSubTypes, meshServicesSubTypes...)
	supportedServiceSubTypes = append(supportedServiceSubTypes, serverlessServicesSubTypes...)
	supportedServiceSubTypes = append(supportedServiceSubTypes, otherServicesSubTypes...)
	supportedServiceSubTypes = append(supportedServiceSubTypes, awsServiceSubTypes...)
	supportedServiceSubTypes = append(supportedServiceSubTypes, azureServiceSubTypes...)
	supportedServiceSubTypes = append(supportedServiceSubTypes, gcpServiceSubTypes...)
	supportedServiceSubTypes = append(supportedServiceSubTypes, doServiceSubTypes...)
	supportedServiceSubTypes = append(supportedServiceSubTypes, pubSubServiceSubTypes...)
	supportedServiceSubTypes = append(supportedServiceSubTypes, buildServiceSubTypes...)
}
func (x ServiceType) String() string {
	return string(x)
}
func (x ServiceSubType) String() string {
	return string(x)
}
func isSupportedServiceType(svcType ServiceType, supportedServiceType []ServiceType) bool {
	for i := range supportedServiceType {
		if supportedServiceType[i] == svcType {
			return true
		}
	}
	return false
}
func isSupportedServiceSubType(svcSubType ServiceSubType, supportedServiceSubTypes []ServiceSubType) bool {

	for i := range supportedServiceSubTypes {
		if supportedServiceSubTypes[i] == svcSubType {
			return true
		}
	}
	return false
}
func IsSupportedService(svcType ServiceType, svcSubType ServiceSubType) error {
	if !isSupportedServiceType(svcType, supportedServiceType) {
		return fmt.Errorf("unsupported service_type %s, Valid Types are %+v", svcType, supportedServiceType)
	}
	if !isSupportedServiceSubType(svcSubType, supportedServiceSubTypes) {
		return fmt.Errorf("unsupported service_sub_type %s, Valid Types are %+v", svcSubType, supportedServiceSubTypes)
	}
	serviceSubTypes := mappedServices[svcType]
	if !isSupportedServiceSubType(svcSubType, serviceSubTypes) {
		return fmt.Errorf("unsupported service_type %s and service_sub_type %s combination. Refer Documentation for more information", svcType, svcSubType)
	}
	return nil
}
