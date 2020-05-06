package services

import (
	"bitbucket.org/cloudplex-devs/microservices-mesh-engine/types"
)

//ServiceId             string                   `bson:"service_id" json:"service_id",valid:"required"`
//	CompanyId             string                   `bson:"company_id" json:"company_id",valid:"required"`
//	Name                  string                   `bson:"name" json:"name" valid:"required"`
//	Version               string                   `bson:"version" json:"version",valid:"required"`
//	ServiceType           constants.ServiceType    `bson:"service_type" json:"service_type",valid:"required,in(k8s|K8S|k8S|K8s)"`
//	ServiceSubType        constants.ServiceSubType `bson:"service_sub_type" json:"service_sub_type",valid:"required,in(kubernetesservice)"`
//	ServiceDependencyInfo interface{}              `bson:"service_dependency_info" json:"service_dependency_info",valid:"required"`
//	Namespace             string                   `bson:"namespace" json:"namespace",valid:"required"`
type Service struct {
	types.ServiceBasicInfo `json:",inline" bson:",inline"`
	ServiceAttributes      KubeServiceAttributes `bson:"service_attributes" json:"service_attributes",valid:"required"`
}
type KubeServiceAttributes struct {
	// The list of ports that are exposed by this service.
	// More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies
	Ports []KubePort `bson:"ports" json:"ports"`
	// Route service traffic to pods with label keys and values matching this
	// selector. If empty or not present, the service is assumed to have an
	// external process managing its endpoints, which Kubernetes will not
	// modify. Only applies to types ClusterIP, NodePort, and LoadBalancer.
	// Ignored if type is ExternalName.
	// More info: https://kubernetes.io/docs/concepts/services-networking/service/
	// +optional
	Selector map[string]string `bson:"selector,omitempty" json:"selector"`
	// clusterIP is the IP address of the service and is usually assigned
	// randomly by the master. If an address is specified manually and is not in
	// use by others, it will be allocated to the service; otherwise, creation
	// of the service will fail. This field can not be changed through updates.
	// Valid values are "None", empty string (""), or a valid IP address. "None"
	// can be specified for headless services when proxying is not required.
	// More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies
	// +optional
	ClusterIP string `bson:"cluster_ip,omitempty" json:"cluster_ip,omitempty"`
	// type determines how the Service is exposed. Defaults to ClusterIP. Valid
	// options are ClusterIP, NodePort, and LoadBalancer.
	// "ClusterIP" allocates a cluster-internal IP address for load-balancing to
	// endpoints. Endpoints are determined by the selector or if that is not
	// specified, by manual construction of an Endpoints object.
	// "NodePort" builds on ClusterIP and allocates a port on every node which
	// routes to the clusterIP.
	// "LoadBalancer" builds on NodePort and creates an
	// external load-balancer (if supported in the current cloud) which routes
	// to the clusterIP.
	// More info: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types
	// +optional
	Type string `bson:"type" json:"type",valid:"required,in(ClusterIP|NodePort|LoadBalancer)" default:"ClusterIP" jsonschema:"enum=ClusterIP,enum=NodePort,enum=LoadBalancer,default=ClusterIP" default:"ClusterIP"`
	// externalTrafficPolicy denotes if this Service desires to route external
	// traffic to node-local or cluster-wide endpoints. "Local" preserves the
	// client source IP and avoids a second hop for LoadBalancer and Nodeport
	// type services, but risks potentially imbalanced traffic spreading.
	// "Cluster" obscures the client source IP and may cause a second hop to
	// another node, but should have good overall load-spreading.
	// +optional
	ExternalTrafficPolicy string `bson:"external_traffic_policy,omitempty" json:"external_traffic_policy" valid:"required,in(Local|Cluster)" jsonschema:"enum=Local,enum=Cluster,default=Local" default:"Local"`
}
type KubePort struct {
	// The name of this port within the service. This must be a DNS_LABEL.
	// All ports within a ServiceSpec must have unique names.
	// prefix of name must be http-*, grpc-*,tcp-*,http2-*
	// +mandatory
	Name string `bson:"name" json:"name"`
	// The IP protocol for this port. Supports "TCP", "UDP", and "SCTP".
	// Default is TCP.
	// +optional
	Protocol string `bson:"protocol" json:"protocol" jsonschema:"enum=TCP,enum=UDP,enum=SCTP,default=TCP" default:"TCP"`
	// The port that will be exposed by this service.
	Port int32 `bson:"port" json:"port"  jsonschema:"minimum=1,maximum=65535"`
	// Number or name of the port to access on the pods targeted by the service.
	// Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.
	// If this is a string, it will be looked up as a named port in the
	// target Pod's container ports.
	// More info: https://kubernetes.io/docs/concepts/services-networking/service/#defining-a-service
	// +optional
	TargetPort PortItntOrString `bson:"target_port,omitempty" json:"target_port,omitempty"jsonschema:"minimum=1,maximum=65535" `
	// The port on each node on which this service is exposed when type=NodePort or LoadBalancer.
	// Usually assigned by the system. If specified, it will be allocated to the service
	// if unused or else creation of the service will fail.
	// supported range 30000-32767
	// More info: https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport
	// +optional
	NodePort int32 `bson:"node_port,omitempty" json:"node_port,omitempty"  jsonschema:"minimum=30000,maximum=32767"`
}
