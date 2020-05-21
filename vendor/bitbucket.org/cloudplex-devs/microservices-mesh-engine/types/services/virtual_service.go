package services

import (
	"bitbucket.org/cloudplex-devs/microservices-mesh-engine/types"
	"time"
)

//Id                interface{}              `json:"_id,omitempty" bson:"_id" valid:"-"`
//	ServiceId         string                   `json:"service_id" bson:"service_id"  valid:"required"`
//	CompanyId         string                   `json:"company_id" bson:"company_id" valid:"required"`
//	Version           string                   `json:"version" bson:"version"  valid:"required"`
//	ServiceType       constants.ServiceType    `json:"service_type" bson:"service_type"  valid:"required"`
//	ServiceSubType    constants.ServiceSubType `json:"service_sub_type" bson:"service_sub_type"  valid:"required"`
//	Name              string                   `json:"name" bson:"name"  valid:"required"`
//	Namespace         string                   `json:"namespace" bson:"namespace"  valid:"required"`
type VirtualService struct {
	types.ServiceBasicInfo `json:",inline" bson:",inline"`
	ServiceAttributes      *VSServiceAttribute `json:"service_attributes" bson:"service_attributes" binding:"required" valid:"required"`
}

type VSServiceAttribute struct {
	// The destination hosts to which traffic is being sent. Could
	// be a DNS name with wildcard prefix or an IP address.  Depending on the
	// platform, short-names can also be used instead of a FQDN (i.e. has no
	// dots in the name). In such a scenario, the FQDN of the host would be
	// derived based on the underlying platform.
	//
	// A single VirtualService can be used to describe all the traffic
	// properties of the corresponding hosts, including those for multiple
	// HTTP and TCP ports. Alternatively, the traffic properties of a host
	// can be defined using more than one VirtualService, with certain
	// caveats. Refer to the
	// [Operations Guide](https://istio.io/docs/ops/best-practices/traffic-management/#split-virtual-services)
	// for details.
	//default * if gateway is selected else name of the service
	//+optional
	Hosts    []string `json:"hosts,omitempty" bson:"hosts,omitempty"`
	Gateways []string `json:"gateways,omitempty" bson:"gateways,omitempty"`
	// An ordered list of route rules for HTTP traffic. HTTP routes will be
	// applied to platform service ports named 'http-*'/'http2-*'/'grpc-*', gateway
	// ports with protocol HTTP/HTTP2/GRPC/ TLS-terminated-HTTPS and service
	// entry ports using HTTP/HTTP2/GRPC protocols.  The first rule matching
	// an incoming request is used.
	Http []*Http `json:"http,omitempty" bson:"http,omitempty"`
	Tls  []*Tls  `json:"tls,omitempty" bson:"tls,omitempty"`
	Tcp  []*Tcp  `json:"tcp,omitempty" bson:"tcp,omitempty"`
}

type Http struct {
	// The name assigned to the route for debugging purposes. The
	// route's name will be concatenated with the match's name and will
	// be logged in the access logs for requests matching this
	// route/match.
	// +optional
	Name string `json:"name,omitempty" bson:"name,omitempty"`
	// Match conditions to be satisfied for the rule to be
	// activated. All conditions inside a single match block have AND
	// semantics, while the list of match blocks have OR semantics. The rule
	// is matched if any one of the match blocks succeed.
	// +optional
	HttpMatch []*HttpMatchRequest `json:"http_match,omitempty" bson:"http_match,omitempty"`
	// A HTTP rule can either redirect or forward (default) traffic. The
	// forwarding target can be one of several versions of a service (see
	// glossary in beginning of document). Weights associated with the
	// service version determine the proportion of traffic it receives.
	// +mandatory
	HttpRoute []*HttpRoute `json:"http_route,omitempty" bson:"http_route,omitempty"`
	// A HTTP rule can either redirect or forward (default) traffic. If
	// traffic passthrough option is specified in the rule,
	// route/redirect will be ignored. The redirect primitive can be used to
	// send a HTTP 301 redirect to a different URI or Authority.
	// +optional
	HttpRedirect *HttpRedirect `json:"http_redirect,omitempty" bson:"http_redirect,omitempty"`
	// Rewrite HTTP URIs and Authority headers. Rewrite cannot be used with
	// Redirect primitive. Rewrite will be performed before forwarding.
	// +optional
	HttpRewrite *HttpRewrite `json:"http_rewrite,omitempty" bson:"http_rewrite,omitempty"`
	//timeout in ms
	// +optional
	Timeout time.Duration `json:"timeout,omitempty" bson:"timeout,omitempty"`
	// Retry policy for HTTP requests.
	// +optional
	Retry *HttpRetry `json:"retry,omitempty" bson:"retry,omitempty"`
	// Fault injection policy to apply on HTTP traffic at the client side.
	// Note that timeouts or retries will not be enabled when faults are
	// enabled on the client side.
	// +optional
	FaultInjection *HttpFaultInjection `json:"fault_injection,omitempty" bson:"fault_injection,omitempty"`
	// Cross-Origin Resource Sharing policy (CORS). Refer to
	// [CORS](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS)
	// for further details about cross origin resource sharing.
	// +optional
	CorsPolicy *HttpCorsPolicy `json:"cors_policy,omitempty" bson:"cors_policy,omitempty"`
}

type HttpMatchRequest struct {
	// The name assigned to a match. The match's name will be
	// concatenated with the parent route's name and will be logged in
	// the access logs for requests matching this route.
	// +optional
	Name string `json:"name,omitempty" bson:"name,omitempty"`
	// URI to match
	// values are case-sensitive and formatted as follows:
	//
	// - exact: "value" for exact string match
	//
	// - prefix: "value" for prefix-based match
	//
	// - regex: "value" for ECMAscript style regex-based match
	// +optional
	Uri *HttpMatch `json:"uri,omitempty" bson:"uri,omitempty"`
	// URI to match
	// values are case-sensitive and formatted as follows:
	//
	// - exact: "value" for exact string match
	//
	// - prefix: "value" for prefix-based match
	//
	// - regex: "value" for ECMAscript style regex-based match
	// +optional
	Scheme *HttpMatch `json:"scheme,omitempty" bson:"scheme,omitempty"`
	// URI to match
	// values are case-sensitive and formatted as follows:
	//
	// - exact: "value" for exact string match
	//
	// - prefix: "value" for prefix-based match
	//
	// - regex: "value" for ECMAscript style regex-based match
	// +optional
	Method *HttpMatch `json:"method,omitempty" bson:"method,omitempty"`
	// URI to match
	// values are case-sensitive and formatted as follows:
	//
	// - exact: "value" for exact string match
	//
	// - prefix: "value" for prefix-based match
	//
	// - regex: "value" for ECMAscript style regex-based match
	// +optional
	Authority *HttpMatch `json:"authority,omitempty" bson:"authority,omitempty"`
	// URI to match
	// values are case-sensitive and formatted as follows:
	//
	// - exact: "value" for exact string match
	//
	// - prefix: "value" for prefix-based match
	//
	// - regex: "value" for ECMAscript style regex-based match
	// +optional
	Headers map[string]*HttpMatch `json:"headers,omitempty" bson:"headers,omitempty"`
}

type HttpMatch struct {
	// Support Values are exact/prefix/regex
	// default prefix
	// +optional
	Type string `json:"type,omitempty" bson:"type,omitempty" default:"prefix" jsonschema:"enum=exact,enum=prefix,enum=regex,default=prefix"`

	Value string `json:"value,omitempty" bson:"value,omitempty"`
}

type HttpRoute struct {
	// Destination uniquely identifies the instances of a service
	// to which the request/connection should be forwarded to.
	Routes []*RouteDestination `json:"routes,omitempty" bson:"routes,omitempty"`
	// The proportion of traffic to be forwarded to the service
	// version. (0-100). Sum of weights across destinations SHOULD BE == 100.
	// If there is only one destination in a rule, the weight value is 100
	// +optional
	Weight *int32 `json:"weight,omitempty" bson:"weight,omitempty" default:"100" jsonschema:"minimum:1,maximum=100"`
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
	// Delay requests before forwarding, emulating various failures such as
	// network issues, overloaded upstream service, etc.
	// valid values are FixedDelay/ExponentialDelay
	// default value is fixedDelay
	// +optional
	DelayType string `json:"delay_type,omitempty" bson:"delay_type,omitempty" valid:"in(FixedDelay|ExponentialDelay)" jsonschema:"enum=FixedDelay,enum=ExponentialDelay,default=FixedDelay" default:"FixedDelay"`
	// delay duration
	// fixed-length span of time represented as a count of seconds
	// and fractions of seconds at nanosecond resolution
	// for example: 2/4.5/3
	// +optional
	DelayValue time.Duration `json:"delay_value,omitempty" bson:"delay_value,omitempty"`
	// Percentage of requests on which the delay will be injected.
	// range is 0-100
	// default 1
	// +optional
	FaultPercentage float32 `json:"fault_percentage,omitempty" bson:"fault_percentage,omitempty" jsonschema:"minimum=0,maximum=100" default:"1"`
	// Abort Http request attempts and return error codes back to downstream
	// service, giving the impression that the upstream service is faulty.
	// supported values are HttpStatus/GrpcStatus/Http2Status
	// default value is HttpStatus
	// +optional
	AbortErrorType string `json:"abort_error_type,omitempty" bson:"abort_error_type" valid:"in(HttpStatus|GrpcStatus|Http2Status),omitempty" jsonschema:"enum=HttpStatus,enum=GrpcStatus,enum=Http2Status,default:HttpStatus" default:"HttpStatus"`
	// error code
	// +optional
	AbortErrorValue string `json:"abort_error_value,omitempty" bson:"abort_error_value,omitempty"`
	// Percentage of requests to be aborted with the error code provided.
	// range is 0-100
	// default 1
	// +optional
	AbortPercentage string `json:"abort_percentage,omitempty" bson:"abort_percentage,omitempty" jsonschema:"minimum=0,maximum=100" default:"1"`
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
	// The name of a service from the service registry. Service
	// names are looked up from the platform's service registry (e.g.,
	// Kubernetes services) and from the hosts
	// destinations that are not found in either of the two, will be dropped.
	// +optional
	Host string `json:"host,omitempty" bson:"host,omitempty"`
	// The name of a subset within the service. Applicable only to services
	// within the mesh. The subset must be defined in a corresponding
	// DestinationRule.
	// optional
	Subset string `json:"subset,omitempty" bson:"subset,omitempty"`
	// Specifies the port on the host that is being addressed. If a service
	// exposes only a single port it is not required to explicitly select the
	// port.
	// +optional
	Port int32 `json:"port,omitempty" bson:"port,omitempty" jsonschema:"minimum=0,maximum=65536"`
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
