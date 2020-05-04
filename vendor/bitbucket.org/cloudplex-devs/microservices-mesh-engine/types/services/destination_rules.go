package services

import (
	"bitbucket.org/cloudplex-devs/microservices-mesh-engine/types"
	"time"
)

//ServiceId string `bson:"service_id" json:"service_id",valid:"required"`
//CompanyId string `bson:"company_id" json:"company_id",valid:"required"`
//Name      string `bson:"name" json:"name",valid:"required"`
//Version   string `json:"version"  bson:"version"  binding:"required" valid:"alphanumspecial,length(1|10),lowercase~lowercase alphanumeric characters are allowed,required"`
//
//ServiceType           constants.ServiceType    `bson:"service_type" json:"service_type",valid:"required"`
//ServiceSubType        constants.ServiceSubType `bson:"service_sub_type" json:"service_sub_type",valid:"required"`
//ServiceDependencyInfo interface{}              `bson:"service_dependency_info" json:"service_dependency_info",valid:"required"`
//Namespace             string                   `bson:"namespace" json:"namespace",valid:"required"`
type DestinationRules struct {
	types.ServiceBasicInfo `json:",inline" bson:",inline"`

	ServiceAttributes DRServiceAttribute `bson:"service_attributes" json:"service_attributes",valid:"required"`
}

type DRServiceAttribute struct {
	// The name of a service from the service registry. Service
	// names are looked up from the platform's service registry (e.g.,
	// Kubernetes services, Consul services, etc.) and from the hosts
	// declared by [ServiceEntries](https://istio.io/docs/reference/config/networking/service-entry/#ServiceEntry). Rules defined for
	// services that do not exist in the service registry will be ignored.
	Host string `json:"host,omitempty" bson:"host,omitempty"`
	// Traffic policies to apply (load balancing policy, connection pool
	// sizes, outlier detection).
	// +optional
	TrafficPolicy *TrafficPolicy `json:"traffic_policy,omitempty" bson:"traffic_policy,omitempty"`
	// One or more named sets that represent individual versions of a
	// service. Traffic policies can be overridden at subset level.
	Subsets []*Subset `json:"subsets,omitempty" bson:"subsets,omitempty"`
}
type TrafficPolicy struct {
	// Settings controlling the load balancer algorithms.
	LoadBalancer *LoadBalancer `json:"load_balancer,omitempty" bson:"load_balancer,omitempty"`
	// Traffic policies specific to individual ports. Note that port level
	// settings will override the destination-level settings. Traffic
	// settings specified at the destination-level will not be inherited when
	// overridden by port-level settings, i.e. default values will be applied
	// to fields omitted in port-level traffic policies.
	PortLevelSettings []*PortLevelSetting `json:"port_level_settings,omitempty" bson:"port_level_settings,omitempty"`
	// Settings controlling the volume of connections to an upstream service
	// circuit breaker configurations
	// +optional
	ConnectionPool *ConnectionPool `json:"connection_pool,omitempty" bson:"connection_pool,omitempty"`
	// Settings controlling eviction of unhealthy hosts from the load balancing pool
	OutlierDetection *OutlierDetection `json:"outlier_detection,omitempty" bson:"outlier_detection,omitempty"`
	// TLS related settings for connections to the upstream service.
	DrTls *DrTls `json:"dr_tls,omitempty" bson:"dr_tls,omitempty"`
}
type LoadBalancer struct {
	// Standard load balancing algorithms that require no tuning.
	// supported values are ROUND_ROBIN/LEAST_CONN/RANDOM/PASSTHROUGH
	// ROUND_ROBIN: Round Robin policy. Default
	// LEAST_CONN: The least request load balancer uses an O(1) algorithm which selects
	// two random healthy hosts and picks the host which has fewer active requests.
	// RANDOM: The random load balancer selects a random healthy host.
	// The random load balancer generally performs better than round robin if no
	// health checking policy is configured.
	// PASSTHROUGH: This option will forward the connection to the original IP address
	//requested by the caller without doing any form of load balancing.
	//This option must be used with care.
	// +optional
	Simple string `json:"simple,omitempty" bson:"simple,omitempty" jsonschema:"enum=ROUND_ROBIN,enum=LEAST_CONN,enum=RANDOM,enum=PASSTHROUGH,default=ROUND_ROBIN"`
	// Consistent Hash-based load balancing can be used to provide soft session affinity
	// based on HTTP headers, cookies or other properties.
	// This load balancing policy is applicable only for HTTP connections.
	//+optional
	ConsistentHash *ConsistentHash `json:"consistent_hash,omitempty" bson:"consistent_hash,omitempty"`
}
type ConsistentHash struct {
	// Hash based on a specific HTTP header.
	// +mandatory
	HTTPHeaderName string `json:"http_header_name,omitempty" bson:"http_header_name,omitempty"`
	// Hash based on the source IP address.
	// +mandatory
	UseSourceIP bool `json:"use_source_ip,omitempty" bson:"use_source_ip,omitempty"`
	// The minimum number of virtual nodes to use for the hash ring. Defaults to 1024.
	// Larger ring sizes result in more granular load distributions.
	// If the number of hosts in the load balancing pool is larger than the ring size,
	// each host will be assigned a single virtual node.
	// +optional
	MinimumRingSize string `json:"minimum_ring_size,omitempty" bson:"minimum_ring_size,omitempty"`
	// Hash based on HTTP cookie.
	// +mandatory
	HttpCookie *HttpCookie `json:"http_cookie,omitempty" bson:"http_cookie,omitempty"`
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
	// Settings common to both HTTP and TCP upstream connections.
	// +optional
	Tcp *DrTcp `json:"dr_tcp,omitempty" bson:"dr_tcp,omitempty"`
	// HTTP connection pool settings.
	// +optional
	Http *DrHttp `json:"dr_http,omitempty" bson:"dr_http,omitempty"`
}
type DrTcp struct {
	// Maximum number of HTTP1 /TCP connections to a destination host. Default 2^32-1.
	// +optional
	MaxConnections int32 `json:"max_connections,omitempty" bson:"max_connections,omitempty"`
	// TCP connection timeout.
	// +optional
	ConnectTimeout *time.Duration `json:"connect_timeout,omitempty" bson:"connect_timeout,omitempty"`
	// If set then set SO_KEEPALIVE on the socket to enable TCP Keepalives.
	// +optional
	TcpKeepalive *TcpKeepalive `json:"tcp_keep_alive,omitempty" bson:"tcp_keep_alive,omitempty"`
}
type TcpKeepalive struct {
	Time     *time.Duration `json:"time,omitempty" bson:"time,omitempty"`
	Interval *time.Duration `json:"interval,omitempty" bson:"interval,omitempty"`
	Probes   uint32         `json:"probes,omitempty" bson:"probes,omitempty"`
}
type DrHttp struct {
	// Maximum number of pending HTTP requests to a destination. Default 2^32-1.
	// +optional
	HTTP1MaxPendingRequests int32 `json:"http_1_max_pending_requests,omitempty" bson:"http_1_max_pending_requests,omitempty"`
	// Maximum number of requests to a backend. Default 2^32-1.
	// +optional
	HTTP2MaxRequests int32 `json:"http_2_max_requests,omitempty" bson:"http_2_max_requests,omitempty"`
	// Maximum number of requests per connection to a backend. Setting this parameter to 1 disables keep alive.
	// Default 0, meaning “unlimited”, up to 2^29.
	// +optional
	MaxRequestsPerConnection int32 `json:"max_requests_per_connection,omitempty" bson:"max_requests_per_connection,omitempty"`
	// Maximum number of retries that can be outstanding to all hosts in a cluster at a given time.
	// Defaults to 2^32-1.
	// +optional
	MaxRetries int32 `json:"max_retries,omitempty" bson:"max_retries,omitempty"`
	// The idle timeout for upstream connection pool connections.
	// The idle timeout is defined as the period in which there are no active requests.
	// If not set, the default is 1 hour. When the idle timeout is reached the connection will be closed.
	// Note that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive.
	// Applies to both HTTP1.1 and HTTP2 connections.
	// +optional
	IdleTimeout int32 `json:"idle_timeout,omitempty" bson:"idle_timeout,omitempty"` //time
	// Specify if http1.1 connection should be upgraded to http2 for the associated destination.
	// support values Default:0, DO_NOT_UPGRADE:1,UPGRADE:2
	// +optional
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
	// Name of the subset. The service name and the subset name can
	// be used for traffic splitting in a route rule.
	Name string `json:"name,omitempty" bson:"name,omitempty"`
	// Labels apply a filter over the endpoints of a service in the
	// service registry.
	// +optional
	Labels *map[string]string `json:"labels,omitempty" bson:"labels,omitempty"`
	// Traffic policies that apply to this subset. Subsets inherit the
	// traffic policies specified at the DestinationRule level. Settings
	// specified at the subset level will override the corresponding settings
	// specified at the DestinationRule level.
	TrafficPolicy *TrafficPolicy `json:"traffic_policy,omitempty" bson:"traffic_policy,omitempty"`
}
type Label struct {
	Version string `json:"version,omitempty" bson:"version,omitempty"`
}
type DrTls struct {
	// Indicates whether connections to this port should be secured using TLS.
	// The value of this field determines how TLS is enforced.
	// +mandatory
	Mode string `json:"mode,omitempty" bson:"mode" valid:"in(ISTIO_MUTUAL|MUTUAL|DISABLE|SIMPLE),omitempty"  jsonschema:"enum=ISTIO_MUTUAL,enum=MUTUAL,enum=DISABLE,enum=SIMPLE" `
	//REQUIRED if mode is MUTUAL.
	// The path to the file holding the client-side TLS certificate to use.
	// Should be empty if mode is ISTIO_MUTUAL.
	ClientCertificate string `json:"client_certificate,omitempty" bson:"client_certificate,omitempty"`
	//REQUIRED if mode is MUTUAL. The path to the file holding the client’s private key.
	// Should be empty if mode is ISTIO_MUTUAL.
	PrivateKey string `json:"private_key,omitempty" bson:"private_key,omitempty"`
	//OPTIONAL: The path to the file containing certificate authority certificates to use
	// in verifying a presented server certificate. If omitted, the proxy will not verify the
	// server’s certificate. Should be empty if mode is ISTIO_MUTUAL.
	CaCertificate string `json:"ca_certificate,omitempty" bson:"ca_certificate,omitempty"`
	// A list of alternate names to verify the subject identity in the certificate. If specified,
	// the proxy will verify that the server certificate’s subject alt name matches one of the specified values.
	// If specified, this list overrides the value of subjectaltnames from the ServiceEntry.
	SubjectAltNames string   `json:"subject_alt_names,omitempty" bson:"subject_alt_names,omitempty"`
	Name            []string `json:"name,omitempty" bson:"name,omitempty"`
}
