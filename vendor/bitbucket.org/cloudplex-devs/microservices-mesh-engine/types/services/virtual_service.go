package services

import "time"

type VirtualService struct {
	Id                interface{}         `json:"_id,omitempty" bson:"_id" valid:"-"`
	ServiceId         string              `json:"service_id" bson:"service_id"  valid:"required"`
	CompanyId         string              `json:"company_id" bson:"company_id" valid:"required"`
	Version           string              `json:"version" bson:"version"  valid:"required"`
	ServiceType       string              `json:"service_type" bson:"service_type"  valid:"required"`
	ServiceSubType    string              `json:"service_sub_type" bson:"service_sub_type"  valid:"required"`
	Name              string              `json:"name" bson:"name"  valid:"required"`
	Namespace         string              `json:"namespace" bson:"namespace"  valid:"required"`
	ServiceAttributes *VSServiceAttribute `json:"service_attributes" bson:"service_attributes"  valid:"required"`
}

type VSServiceAttribute struct {
	Hosts    []string `json:"hosts,omitempty" bson:"hosts,omitempty"`
	Gateways []string `json:"gateways,omitempty" bson:"gateways,omitempty"`
	Http     []*Http  `json:"http,omitempty" bson:"http,omitempty"`
	Tls      []*Tls   `json:"tls,omitempty" bson:"tls,omitempty"`
	Tcp      []*Tcp   `json:"tcp,omitempty" bson:"tcp,omitempty"`
}

type Http struct {
	Name         string              `json:"name,omitempty" bson:"name,omitempty"`
	HttpMatch    []*HttpMatchRequest `json:"http_match,omitempty" bson:"http_match,omitempty"`
	HttpRoute    []*HttpRoute        `json:"http_route,omitempty" bson:"http_route,omitempty"`
	HttpRedirect *HttpRedirect       `json:"http_redirect,omitempty" bson:"http_redirect,omitempty"`
	HttpRewrite  *HttpRewrite        `json:"http_rewrite,omitempty" bson:"http_rewrite,omitempty"`
	//timeout in ms
	Timeout        time.Duration       `json:"timeout,omitempty" bson:"timeout,omitempty"`
	Retry          *HttpRetry          `json:"retry,omitempty" bson:"retry,omitempty"`
	FaultInjection *HttpFaultInjection `json:"fault_injection,omitempty" bson:"fault_injection,omitempty"`
	CorsPolicy     *HttpCorsPolicy     `json:"cors_policy,omitempty" bson:"cors_policy,omitempty"`
}

type HttpMatchRequest struct {
	Name      string     `json:"name,omitempty" bson:"name,omitempty"`
	Uri       *HttpMatch `json:"uri,omitempty" bson:"uri,omitempty"`
	Scheme    *HttpMatch `json:"scheme,omitempty" bson:"scheme,omitempty"`
	Method    *HttpMatch `json:"method,omitempty" bson:"method,omitempty"`
	Authority *HttpMatch `json:"authority,omitempty" bson:"authority,omitempty"`
}

type HttpMatch struct {
	Type  string `json:"type,omitempty" bson:"type,omitempty"`
	Value string `json:"value,omitempty" bson:"value,omitempty"`
}

type HttpRoute struct {
	Routes []*RouteDestination `json:"routes,omitempty" bson:"routes,omitempty"`
	Weight int32               `json:"weight,omitempty" bson:"weight,omitempty"`
}

type HttpRedirect struct {
	Uri          string `json:"uri,omitempty" bson:"uri,omitempty"`
	Authority    string `json:"authority,omitempty" bson:"authority,omitempty"`
	RedirectCode int32  `json:"redirect_code,omitempty" bson:"redirect_code,omitempty"`
}

type HttpRewrite struct {
	Uri       string `json:"uri,omitempty" bson:"uri,omitempty"`
	Authority string `json:"authority,omitempty" bson:"authority,omitempty"`
}

type HttpRetry struct {
	TotalAttempts int32  `json:"total_attempt,omitempty" bson:"total_attempt,omitempty"`
	PerTryTimeOut int64  `json:"per_try_timeout,omitempty" bson:"per_try_timeout,omitempty"`
	RetryOn       string `json:"retry_on,omitempty" bson:"retry_on,omitempty"`
}

type HttpFaultInjection struct {
	DelayType       string        `json:"delay_type,omitempty" bson:"delay_type" valid:"in(FixedDelay|ExponentialDelay),omitempty"`
	DelayValue      time.Duration `json:"delay_value,omitempty" bson:"delay_value,omitempty"`
	FaultPercentage float32       `json:"fault_percentage,omitempty" bson:"fault_percentage,omitempty"`
	AbortErrorType  string        `json:"abort_error_type,omitempty" bson:"abort_error_type" valid:"in(HttpStatus|GrpcStatus|Http2Status),omitempty"`
	AbortErrorValue string        `json:"abort_error_value,omitempty" bson:"abort_error_value,omitempty"`
	AbortPercentage string        `json:"abort_percentage,omitempty" bson:"abort_percentage,omitempty"`
}

type HttpCorsPolicy struct {
	AllowOrigin   []string `json:"allow_origin,omitempty" bson:"allow_origin,omitempty"`
	AllowMethod   []string `json:"allow_method,omitempty" bson:"allow_method,omitempty"`
	AllowHeaders  []string `json:"allow_headers,omitempty" bson:"allow_headers,omitempty"`
	ExposeHeaders []string `json:"expose_headers,omitempty" bson:"expose_headers,omitempty"`
	//max age in ms
	MaxAge           time.Duration `json:"max_age,omitempty" bson:"max_age,omitempty"`
	AllowCredentials bool          `json:"allow_Credentials,omitempty" bson:"allow_credentials,omitempty"`
}

type Tls struct {
	Match []*TlsMatchAttribute `json:"tls_match,omitempty" bson:"tls_match,omitempty"`
	Route []*TlsRoute          `json:"tls_route,omitempty" bson:"tls_route,omitempty"`
}

type TlsMatchAttribute struct {
	SniHosts           []string `json:"sni_hosts,omitempty" bson:"sni_hosts,omitempty"`
	DestinationSubnets []string `json:"destination_subnets,omitempty" bson:"destination_subnets,omitempty"`
	Port               int32    `json:"port,omitempty" bson:"port,omitempty"`
	SourceSubnet       string   `json:"source_subnet,omitempty" bson:"source_subnet" valid:"in(ipv4|ipv6),omitempty"`
	Gateways           []string `json:"gateways,omitempty" bson:"gateways,omitempty"`
}

type TlsRoute struct {
	RouteDestination *RouteDestination `json:"route_destination,omitempty" bson:"route_destination,omitempty"`
	Weight           int32             `json:"weight,omitempty" bson:"weight,omitempty"`
}

type RouteDestination struct {
	Host   string `json:"host,omitempty" bson:"host,omitempty"`
	Subset string `json:"subset,omitempty" bson:"subset,omitempty"`
	Port   int32  `json:"port,omitempty" bson:"port,omitempty"`
}

type Tcp struct {
	Match  []*TcpMatchRequest `json:"tcp_match,omitempty" bson:"tcp_match,omitempty"`
	Routes []*TcpRoutes       `json:"tcp_routes,omitempty" bson:"tcp_routes,omitempty"`
}

type TcpMatchRequest struct {
	DestinationSubnets []string           `json:"destination_subnets,omitempty" bson:"destination_subnets,omitempty"`
	Port               int32              `json:"port,omitempty" bson:"port,omitempty"`
	SourceSubnet       string             `json:"source_subnet,omitempty" bson:"source_subnet,omitempty"`
	SourceLabels       *map[string]string `json:"source_labels,omitempty" bson:"source_labels,omitempty"`
	Gateways           []string           `json:"gateways,omitempty" bson:"gateways,omitempty"`
}

type TcpRoutes struct {
	Destination *RouteDestination `json:"destination,omitempty" bson:"destination,omitempty"`
	Weight      int32             `json:"weight,omitempty" bson:"weight,omitempty"`
}
