package types

type Service struct {
	ServiceId             string                `bson:"service_id" json:"service_id",valid:"required`
	CompanyId             string                `bson:"company_id" json:"company_id",valid:"required`
	Name                  string                `bson:"name" json:"name", valid:"required"`
	Version               string                `bson:"version" json:"version",valid:"required"`
	ServiceType           string                `bson:"service_type" json:"service_type",valid:"required,in(k8s|K8S|k8S|K8s)"`
	ServiceSubType        string                `bson:"service_sub_type" json:"service_sub_type",valid:"required,in(kubernetesservice)"`
	ServiceDependencyInfo interface{}           `bson:"service_dependency_info" json:"service_dependency_info",valid:"required"`
	Namespace             string                `bson:"namespace" json:"namespace",valid:"required"`
	ServiceAttributes     KubeServiceAttributes `bson:"service_attributes" json:"service_attributes",valid:"required"`
}
type KubeServiceAttributes struct {
	Ports                 []KubePort        `bson:"kube_ports" json:"kube_ports"`
	Selector              map[string]string `bson:"selector,omitempty" json:"selector"`
	ClusterIP             string            `bson:"cluster_ip,omitempty" json:"cluster_ip"`
	Type                  string            `bson:"type,omitempty" json:"type",valid:"required,in(ClusterIp|NodePort|LoadBalancer)"`
	ExternalTrafficPolicy string            `bson:"external_traffic_policy,omitempty" json:"external_traffic_policy" valid:"required,in(Local|Cluster)"`
}
type KubePort struct {
	Name       string           `bson:"name,omitempty" json:"name"`
	Protocol   string           `bson:"protocol,omitempty" json:"protocol"`
	Port       int32            `bson:"port,omitempty" json:"port"`
	TargetPort PortItntOrString `bson:"target_port,omitempty" json:"target_port"`
	NodePort   int32            `bson:"node_port,omitempty" json:"node_port"`
}
