package types

import "time"

type DestinationRules struct {
	ServiceId             string             `bson:"service_id" json:"service_id",valid:"required"`
	CompanyId             string             `bson:"company_id" json:"company_id",valid:"required"`
	Name                  string             `bson:"name" json:"name",valid:"required"`
	ServiceType           string             `bson:"service_type" json:"service_type",valid:"required"`
	ServiceSubType        string             `bson:"service_sub_type" json:"service_sub_type",valid:"required"`
	ServiceDependencyInfo interface{}        `bson:"service_dependency_info" json:"service_dependency_info",valid:"required"`
	Namespace             string             `bson:"namespace" json:"namespace",valid:"required"`
	ServiceAttributes     DRServiceAttribute `bson:"service_attributes" json:"service_attributes",valid:"required"`
}

type DRServiceAttribute struct {
	Host          string         `json:"host" bson:"host"`
	TrafficPolicy *TrafficPolicy `json:"traffic_policy" bson:"traffic_policy"`
	Subsets       *Subset        `json:"subsets" bson:"subsets"`
}
type TrafficPolicy struct {
	LoadBalancer      *LoadBalancer       `json:"load_balancer" bson:"load_balancer"`
	PortLevelSettings []*PortLevelSetting `json:"port_level_settings" bson:"port_level_settings"`
	ConnectionPool    *ConnectionPool     `json:"connection_pool" bson:"connection_pool"`
	OutlierDetection  *OutlierDetection   `json:"outlier_detection" bson:"outlier_detection"`
	DrTls             *DrTls              `json:"dr_tls" bson:"dr_tls"`
}
type LoadBalancer struct {
	Simple         string          `json:"simple" bson:"simple"`
	ConsistentHash *ConsistentHash `json:"consistent_hash" bson:"consistent_hash"`
}
type ConsistentHash struct {
	HTTPHeaderName  string      `json:"http_header_name" bson:"http_header_name"`
	UseSourceIP     bool        `json:"use_source_ip" bson:"use_source_ip"`
	MinimumRingSize string      `json:"minimum_ring_size" bson:"minimum_ring_size"`
	HttpCookie      *HttpCookie `json:"http_cookie" bson:"http_cookie"`
}
type HttpCookie struct {
	Name string `json:"name" bson:"name"`
	Path string `json:"path" bson:"path" `
	Ttl  int64  `json:"ttl" bson:"ttl"`
}
type PortLevelSetting struct {
	Port             *DrPort           `json:"dr_port" bson:"dr_port"`
	LoadBalancer     *LoadBalancer     `json:"load_balancer" bson:"load_balancer"`
	ConnectionPool   *ConnectionPool   `json:"connection_pool" bson:"connection_pool"`
	OutlierDetection *OutlierDetection `json:"outlier_detection" bson:"outlier_detection"`
	DrTls            *DrTls            `json:"dr_tls" bson:"dr_tls"`
}
type DrPort struct {
	Number int32 `json:"number" bson:"number"`
}
type ConnectionPool struct {
	Tcp  *DrTcp  `json:"dr_tcp" bson:"dr_tcp"`
	Http *DrHttp `json:"dr_http" bson:"dr_http"`
}
type DrTcp struct {
	MaxConnections int32          `json:"max_connections" bson:"max_connections"`
	ConnectTimeout *time.Duration `json:"connect_timeout" bson:"connect_timeout"` //time
	TcpKeepalive   *TcpKeepalive  `json:"tcp_keep_alive" bson:"tcp_keep_alive"`
}
type TcpKeepalive struct {
	Time     *time.Duration `json:"time" bson:"time"`         //time
	Interval *time.Duration `json:"interval" bson:"interval"` //time
	Probes   uint32         `json:"probes" bson:"probes"`
}
type DrHttp struct {
	HTTP1MaxPendingRequests                           int32 `json:"http_1_max_pending_requests" bson:"http_1_max_pending_requests"`
	HTTP2MaxRequests                                  int32 `json:"http_2_max_requests" bson:"http_2_max_requests"`
	MaxRequestsPerConnection                          int32 `json:"max_requests_per_connection" bson:"max_requests_per_connection"`
	MaxRetries                                        int32 `json:"max_retries" bson:"max_retries"`
	IdleTimeout                                       int32 `json:"idle_timeout" bson:"idle_timeout"` //time
	ConnectionPoolSettingsHTTPSettingsH2UpgradePolicy int32 `json:"connection_pool_settings_http_settings_h2_upgrade_policy" bson:"connection_pool_settings_http_settings_h2_upgrade_policy"`
}
type OutlierDetection struct {
	ConsecutiveErrors  int32          `json:"consecutive_errors" bson:"consecutive_errors"`
	Interval           *time.Duration `json:"interval" bson:"interval"`                     //time
	BaseEjectionTime   *time.Duration `json:"base_ejection_time" bson:"base_ejection_time"` //time
	MaxEjectionPercent int32          `json:"max_ejection_percent" bson:"max_ejection_percent"`
	MinHealthPercent   int32          `json:"min_health_percent" bson:"min_health_percent"`
}
type Subset struct {
	Name          []string           `json:"name" bson:"name"`
	Labels        *map[string]string `json:"labels" bson:"labels"` //map
	TrafficPolicy *TrafficPolicy     `json:"traffic_policy" bson:"traffic_policy" `
}
type Label struct {
	Version string `json:"version" bson:"version"`
}
type DrTls struct {
	Mode              string   `json:"mode" bson:"mode"` //mode
	ClientCertificate string   `json:"client_certificate" bson:"client_certificate"`
	PrivateKey        string   `json:"private_key" bson:"private_key"`
	CaCertificates    string   `json:"ca_certificates" bson:"ca_certificates"`
	SubjectAltNames   string   `json:"subject_alt_names" bson:"subject_alt_names"`
	Name              []string `json:"name" bson:"name"`
}

/*
const (
	MODE_DISABLE Mode= "DISABLE";
	MODE_SIMPLE  Mode= "SIMPLE";
	MODE_MUTUAL  Mode= "MUTUAL";
	MODE_ISTIO_MUTUAL Mode= "ISTIO_MUTUAL";
)
*/
/*
type Simple string

const (
	SIMPLE_ROUND_ROBIN Simple = "ROUND_ROBIN"
	SIMPLE_LEAST_CONN  Simple = "LEAST_CONN"
	SIMPLE_RANDOM      Simple = "RANDOM"
	SIMPLE_PASSTHROUGH Simple = "PASSTHROUGH"
)


H2UpgradePolicy
const(
ConnectionPoolSettings_HTTPSettings_DEFAULT =0
ConnectionPoolSettings_HTTPSettings_DO_NOT_UPGRADE =1
ConnectionPoolSettings_HTTPSettings_UPGRADE =2
)
*/
