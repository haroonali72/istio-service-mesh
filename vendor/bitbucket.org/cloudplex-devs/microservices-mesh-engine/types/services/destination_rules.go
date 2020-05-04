package services

import "time"

type DestinationRules struct {
	ServiceId string `bson:"service_id" json:"service_id",valid:"required"`
	CompanyId string `bson:"company_id" json:"company_id",valid:"required"`
	Name      string `bson:"name" json:"name",valid:"required"`
	Version   string `json:"version"  bson:"version"  binding:"required" valid:"alphanumspecial,length(1|10),lowercase~lowercase alphanumeric characters are allowed,required"`

	ServiceType           string             `bson:"service_type" json:"service_type",valid:"required"`
	ServiceSubType        string             `bson:"service_sub_type" json:"service_sub_type",valid:"required"`
	ServiceDependencyInfo interface{}        `bson:"service_dependency_info" json:"service_dependency_info",valid:"required"`
	Namespace             string             `bson:"namespace" json:"namespace",valid:"required"`
	ServiceAttributes     DRServiceAttribute `bson:"service_attributes" json:"service_attributes",valid:"required"`
}

type DRServiceAttribute struct {
	Host          string         `json:"host,omitempty" bson:"host,omitempty"`
	TrafficPolicy *TrafficPolicy `json:"traffic_policy,omitempty" bson:"traffic_policy,omitempty"`
	Subsets       []*Subset      `json:"subsets,omitempty" bson:"subsets,omitempty"`
}
type TrafficPolicy struct {
	LoadBalancer      *LoadBalancer       `json:"load_balancer,omitempty" bson:"load_balancer,omitempty"`
	PortLevelSettings []*PortLevelSetting `json:"port_level_settings,omitempty" bson:"port_level_settings,omitempty"`
	ConnectionPool    *ConnectionPool     `json:"connection_pool,omitempty" bson:"connection_pool,omitempty"`
	OutlierDetection  *OutlierDetection   `json:"outlier_detection,omitempty" bson:"outlier_detection,omitempty"`
	DrTls             *DrTls              `json:"dr_tls,omitempty" bson:"dr_tls,omitempty"`
}
type LoadBalancer struct {
	Simple         string          `json:"simple,omitempty" bson:"simple,omitempty"`
	ConsistentHash *ConsistentHash `json:"consistent_hash,omitempty" bson:"consistent_hash,omitempty"`
}
type ConsistentHash struct {
	HTTPHeaderName  string      `json:"http_header_name,omitempty" bson:"http_header_name,omitempty"`
	UseSourceIP     bool        `json:"use_source_ip,omitempty" bson:"use_source_ip,omitempty"`
	MinimumRingSize string      `json:"minimum_ring_size,omitempty" bson:"minimum_ring_size,omitempty"`
	HttpCookie      *HttpCookie `json:"http_cookie,omitempty" bson:"http_cookie,omitempty"`
}
type HttpCookie struct {
	Name string `json:"name,omitempty" bson:"name,omitempty"`
	Path string `json:"path,omitempty" bson:"path,omitempty" `
	Ttl  int64  `json:"ttl,omitempty" bson:"ttl,omitempty"`
}
type PortLevelSetting struct {
	Port             *DrPort           `json:"dr_port,omitempty" bson:"dr_port,omitempty"`
	LoadBalancer     *LoadBalancer     `json:"load_balancer,omitempty" bson:"load_balancer,omitempty"`
	ConnectionPool   *ConnectionPool   `json:"connection_pool,omitempty" bson:"connection_pool,omitempty"`
	OutlierDetection *OutlierDetection `json:"outlier_detection,omitempty" bson:"outlier_detection,omitempty"`
	DrTls            *DrTls            `json:"dr_tls,omitempty" bson:"dr_tls,omitempty"`
}
type DrPort struct {
	Number int32 `json:"number,omitempty" bson:"number,omitempty"`
}
type ConnectionPool struct {
	Tcp  *DrTcp  `json:"dr_tcp,omitempty" bson:"dr_tcp,omitempty"`
	Http *DrHttp `json:"dr_http,omitempty" bson:"dr_http,omitempty"`
}
type DrTcp struct {
	MaxConnections int32          `json:"max_connections,omitempty" bson:"max_connections,omitempty"`
	ConnectTimeout *time.Duration `json:"connect_timeout,omitempty" bson:"connect_timeout,omitempty"`
	TcpKeepalive   *TcpKeepalive  `json:"tcp_keep_alive,omitempty" bson:"tcp_keep_alive,omitempty"`
}
type TcpKeepalive struct {
	Time     *time.Duration `json:"time,omitempty" bson:"time,omitempty"`
	Interval *time.Duration `json:"interval,omitempty" bson:"interval,omitempty"`
	Probes   uint32         `json:"probes,omitempty" bson:"probes,omitempty"`
}
type DrHttp struct {
	HTTP1MaxPendingRequests                           int32 `json:"http_1_max_pending_requests,omitempty" bson:"http_1_max_pending_requests,omitempty"`
	HTTP2MaxRequests                                  int32 `json:"http_2_max_requests,omitempty" bson:"http_2_max_requests,omitempty"`
	MaxRequestsPerConnection                          int32 `json:"max_requests_per_connection,omitempty" bson:"max_requests_per_connection,omitempty"`
	MaxRetries                                        int32 `json:"max_retries,omitempty" bson:"max_retries,omitempty"`
	IdleTimeout                                       int32 `json:"idle_timeout,omitempty" bson:"idle_timeout,omitempty"` //time
	ConnectionPoolSettingsHTTPSettingsH2UpgradePolicy int32 `json:"connection_pool_settings_http_settings_h2_upgrade_policy,omitempty" bson:"connection_pool_settings_http_settings_h2_upgrade_policy,omitempty"`
}

type OutlierDetection struct {
	ConsecutiveErrors  int32          `json:"consecutive_errors,omitempty" bson:"consecutive_errors,omitempty"`
	Interval           *time.Duration `json:"interval,omitempty" bson:"interval,omitempty"`
	BaseEjectionTime   *time.Duration `json:"base_ejection_time,omitempty" bson:"base_ejection_time,omitempty"`
	MaxEjectionPercent int32          `json:"max_ejection_percent,omitempty" bson:"max_ejection_percent,omitempty"`
	MinHealthPercent   int32          `json:"min_health_percent,omitempty" bson:"min_health_percent,omitempty"`
}
type Subset struct {
	Name          string             `json:"name,omitempty" bson:"name,omitempty"`
	Labels        *map[string]string `json:"labels,omitempty" bson:"labels,omitempty"`
	TrafficPolicy *TrafficPolicy     `json:"traffic_policy,omitempty" bson:"traffic_policy,omitempty"`
}
type Label struct {
	Version string `json:"version,omitempty" bson:"version,omitempty"`
}
type DrTls struct {
	Mode              string   `json:"mode,omitempty" bson:"mode" valid:"in(ISTIO_MUTUAL|MUTUAL|DISABLE|SIMPLE),omitempty"`
	ClientCertificate string   `json:"client_certificate,omitempty" bson:"client_certificate,omitempty"`
	PrivateKey        string   `json:"private_key,omitempty" bson:"private_key,omitempty"`
	CaCertificate     string   `json:"ca_certificate,omitempty" bson:"ca_certificate,omitempty"`
	SubjectAltNames   string   `json:"subject_alt_names,omitempty" bson:"subject_alt_names,omitempty"`
	Name              []string `json:"name,omitempty" bson:"name,omitempty"`
}
