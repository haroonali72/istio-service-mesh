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
	Hosts    []string `json:"hosts" bson:"hosts" `
	Gateways []string `json:"gateways" bson:"gateways"`
	Http     []*Http  `json:"http" bson:"http"`
	Tls      []*Tls   `json:"tls" bson:"tls"`
	Tcp      []*Tcp   `json:"tcp" bson:"tcp"`
}

type Http struct {
	Name         string              `json:"name" bson:"name"`
	HttpMatch    []*HttpMatchRequest `json:"http_match" bson:"http_match"`
	HttpRoute    []*HttpRoute        `json:"http_route" bson:"http_route"`
	HttpRedirect *HttpRedirect       `json:"http_redirect" bson:"http_redirect"`
	HttpRewrite  *HttpRewrite        `json:"http_rewrite" bson:"http_rewrite"`
	//timeout in ms
	Timeout        time.Duration       `json:"timeout" bson:"timeout"`
	Retry          *HttpRetry          `json:"retry" bson:"retry"`
	FaultInjection *HttpFaultInjection `json:"fault_injection" bson:"fault_injection"`
	CorsPolicy     *HttpCorsPolicy     `json:"cors_policy" bson:"cors_policy"`
}

type HttpMatchRequest struct {
	Name      string     `json:"name" bson:"name"`
	Uri       *HttpMatch `json:"uri" bson:"uri"`
	Scheme    *HttpMatch `json:"scheme" bson:"scheme"`
	Method    *HttpMatch `json:"method" bson:"method"`
	Authority *HttpMatch `json:"authority" bson:"authority"`
}

type HttpMatch struct {
	Type  string `json:"type" bson:"type"`
	Value string `json:"value" bson:"value"`
}

type HttpRoute struct {
	Routes []*RouteDestination `json:"routes" bson:"routes"`
	Weight int32               `json:"weight" bson:"weight"`
}

type HttpRedirect struct {
	Uri          string `json:"uri" bson:"uri"`
	Authority    string `json:"authority" bson:"authority"`
	RedirectCode int32  `json:"redirect_code" bson:"redirect_code"`
}

type HttpRewrite struct {
	Uri       string `json:"uri" bson:"uri"`
	Authority string `json:"authority" bson:"authority"`
}

type HttpRetry struct {
	TotalAttempts int32  `json:"total_attempt" bson:"total_attempt"`
	PerTryTimeOut int64  `json:"per_try_timeout" bson:"per_try_timeout"`
	RetryOn       string `json:"retry_on" bson:"retry_on"`
}

type HttpFaultInjection struct {
	DelayType       string        `json:"delay_type" bson:"delay_type" valid:"in(FixedDelay|ExponentialDelay)"`
	DelayValue      time.Duration `json:"delay_value" bson:"delay_value"`
	FaultPercentage float32       `json:"fault_percentage" bson:"fault_percentage"`
	AbortErrorType  string        `json:"abort_error_type" bson:"abort_error_type" valid:"in(HttpStatus|GrpcStatus|Http2Status)"`
	AbortErrorValue string        `json:"abort_error_value" bson:"abort_error_value"`
	AbortPercentage string        `json:"abort_percentage" bson:"abort_percentage"`
}

type HttpCorsPolicy struct {
	AllowOrigin   []string `json:"allow_origin" bson:"allow_origin"`
	AllowMethod   []string `json:"allow_method" bson:"allow_method"`
	AllowHeaders  []string `json:"allow_headers" bson:"allow_headers"`
	ExposeHeaders []string `json:"expose_headers" bson:"expose_headers"`
	//max age in ms
	MaxAge           time.Duration `json:"max_age" bson:"max_age"`
	AllowCredentials bool          `json:"allow_Credentials" bson:"allow_credentials"`
}

type Tls struct {
	Match []*TlsMatchAttribute `json:"tls_match" bson:"tls_match"`
	Route []*TlsRoute          `json:"tls_route" bson:"tls_route"`
}

type TlsMatchAttribute struct {
	SniHosts           []string `json:"sni_hosts" bson:"sni_hosts"`
	DestinationSubnets []string `json:"destination_subnets" bson:"destination_subnets"`
	Port               int32    `json:"port" bson:"port"`
	SourceSubnet       string   `json:"source_subnet" bson:"source_subnet" valid:"in(ipv4|ipv6)"`
	Gateways           []string `json:"gateways" bson:"gateways"`
}

type TlsRoute struct {
	RouteDestination *RouteDestination `json:"route_destination" bson:"route_destination"`
	Weight           int32             `json:"weight" bson:"weight"`
}

type RouteDestination struct {
	Host   string `json:"host" bson:"host"`
	Subnet string `json:"subnet" bson:"subnet"`
	Port   int32  `json:"port" bson:"port"`
}

type Tcp struct {
	Match  []*TcpMatchRequest `json:"tcp_match" bson:"tcp_match"`
	Routes []*TcpRoutes       `json:"tcp_routes" bson:"tcp_routes"`
}

type TcpMatchRequest struct {
	DestinationSubnets []string           `json:"destination_subnets" bson:"destination_subnets"`
	Port               int32              `json:"port" bson:"port"`
	SourceSubnet       string             `json:"source_subnet" bson:"source_subnet"`
	SourceLabels       *map[string]string `json:"source_labels" bson:"source_labels"`
	Gateways           []string           `json:"gateways" bson:"gateways"`
}

type TcpRoutes struct {
	Destination *RouteDestination `json:"destination" bson:"destination"`
	Weight      int32             `json:"weight" bson:"weight"`
}
