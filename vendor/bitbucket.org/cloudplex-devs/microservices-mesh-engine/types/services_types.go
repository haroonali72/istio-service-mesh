package types

import (
	"bitbucket.org/cloudplex-devs/microservices-mesh-engine/constants"
	"time"
)

type ServiceMeshRequest struct {
	Replicas              int                  `json:"replicas"`
	Id                    *string              `json:"_id"`
	ServiceId             *string              `json:"service_id" bson:"service_id" `
	Name                  *string              `json:"name"`
	NameSpace             *string              `json:"namespace"`
	Version               *string              `json:"version"`
	UserId                *string              `json:"user_id"`
	ServiceType           *string              `json:"service_type"`
	ServiceCategory       *string              `json:"service_sub_type"`
	IsYaml                bool                 `json:"is_yaml"`
	GroupId               string               `json:"group_id" bson:"group_id" valid:"-"`
	ServiceAttributes     interface{}          `json:"service_attributes"`
	ServiceDependencyInfo []ServiceMeshRequest `json:"service_dependency_info"`
}
type HookConfiguration struct {
	Weight        int64  `json:"weight"`
	ServiceID     string `json:"service_id"`
	ServiceStatus string
	PreInstall    bool `json:"pre_install"`
	PreUpdate     bool `json:"pre_update"`
	PostUpdate    bool `json:"post_update"`
	PostInstall   bool `json:"post_install"`
	PreDelete     bool `json:"pre_delete"`
	PostDelete    bool `json:"post_delete"`
	PreRollBack   bool `json:"pre_rollback"`
	PostRollBack  bool `json:"post_rollback"`
}
type ServiceBasicInfo struct {
	// number of replicas of Deployment/StatefulSets
	// +optional
	Replicas int `json:"replicas" bson:"replicas" default:"1"`
	// auto populated key
	// +optional
	CompanyId string `json:"company_id,omitempty" bson:"company_id" valid:"-"`
	// auto generated id
	// +optional
	Id interface{} `json:"_id,omitempty" bson:"_id" valid:"-"`
	// ServiceId is unique key in a solution. service_id is primarily intended to use in dynamic parameters and dependency management
	// valid regex is ^[ A-Za-z0-9_-]*$
	// valid length is 6-100 character
	// cannot update if deployed
	// +mandatory
	ServiceId string `json:"service_id" bson:"service_id" binding:"required" valid:"alphanumspecial,length(6|100)~service_id must contain between 6 and 100 characters,lowercase~service_id is invalid. Valid regex is ^[ A-Za-z0-9_-]*$,required~service_id is missing in request"`
	// Name is required when creating resources on Kubernetes,
	// Name is primarily intended for creation idempotence and configuration
	// definition.
	// Cannot be updated.
	// valid regex is ^[ A-Za-z0-9_-]*$
	// valid length is 5-30 character
	// +mandatory
	Name string `json:"name"  bson:"name" binding:"required" valid:"alphanumspecial,length(5|30),lowercase~service name is invalid. Valid regex is ^[ A-Za-z0-9_-]*$,required"`
	// Version is required for versioning purpose
	// Version will be concatinated with name for container services,
	// Cannot be updated.
	// valid regex is ^[ A-Za-z0-9_-]*$
	// valid length is 1-10 character
	// default=v1
	// +mandatory
	Version string `json:"version" bson:"version"  binding:"required" valid:"alphanumspecial,length(1|10),lowercase~service version is invalid. Valid regex is ^[ A-Za-z0-9_-]*$,required" default:"v1"`
	// ServiceType is used to categorize score and objects.
	// Refer Documentation for more details
	// valid regex is ^[ A-Za-z0-9_-]*$
	// +mandatory
	ServiceType constants.ServiceType `json:"service_type" bson:"service_type"  binding:"required" valid:"alphanumspecial,lowercase~service_type is invalid. Valid regex is ^[ A-Za-z0-9_-]*$,required" swaggertype:"string" `
	// ServiceSubType is used to categorize score and objects.
	// Refer Documentation for more details
	// valid regex is ^[ A-Za-z0-9_-]*$
	// +mandatory
	ServiceSubType constants.ServiceSubType `json:"service_sub_type" bson:"service_sub_type" binding:"required" valid:"alphanumspecial,lowercase~service_sub_type is invalid. Valid regex is ^[ A-Za-z0-9_-]*$,required" swaggertype:"string"`
	// GroupId is used to group container type services
	// if you are grouping services then name of the services
	// must be same and version  must be different
	// +optional
	GroupId string `json:"group_id,omitempty" bson:"group_id" valid:"-"`
	// Deleted is used when you delete a service from a deployed solution
	// This key is used for internal references
	// Any Value assigned in this will be ignored
	// +optional
	Deleted bool `json:"deleted,omitempty" bson:"deleted" valid:"-"`
	// Status of the service
	// valid status are success/failed/new
	Status string `json:"status,omitempty" bson:"status" valid:"-" default:"new"`
	//+deprecated
	NumberOfInstances int `json:"number_of_instances,omitempty" bson:"number_of_instances" `
	// CreationTime is timestamp of when the service is created
	// auto generated time
	// +optional
	CreationDate time.Time `json:"creation_date,omitempty" bson:"creation_date" valid:"-"`
	// Namespace defines the space within each name must be unique. An empty namespace is
	// equivalent to the "default" namespace, but "default" is the canonical representation.
	// Not all objects are required to be scoped to a namespace - the value of this field for
	// those objects will be empty.
	//
	// Must be a DNS_LABEL.
	// Cannot be updated.
	// More info: http://kubernetes.io/docs/user-guide/namespaces
	// +optional
	Namespace string `json:"namespace,omitempty" valid:"-" default:"default"`
	// NamespaceColor defines namespace color on Cloudplex UI
	// +optional
	NameSpaceColor string `json:"namespace_color,omitempty" bson:"namespace_color" valid:"-"`

	ChunckId string `json:"pool_id,omitempty" valid:"-"`
	//BeforeServices is the list of ServiceIds which will deploy before this service
	// +optional
	BeforeServices []*string `json:"in,omitempty" bson:"in" valid:"-"`
	//AfterServices is the list of ServiceIds which will deploy after this service
	// +optional
	AfterServices []*string            `json:"out,omitempty" bson:"out" valid:"-"`
	Dependencies  map[string][]*string `json:"-" bson:"-" valid:"-"`
	// required for Cloudplex Frontend
	// +optional
	Attrs *struct {
		Body struct {
			RefWidth        string  `json:"ref_width" bson:"ref_width" valid:"-"`
			RefHeight       string  `json:"ref_height" bson:"ref_height" valid:"-"`
			Stroke          string  `json:"stroke" bson:"stroke" valid:"-"`
			Fill            string  `json:"fill" bson:"fill" valid:"-"`
			StrokeWidth     float64 `json:"stroke_width" bson:"stroke_width" valid:"-"`
			StrokeDasharray string  `json:"stroke_dasharray" bson:"stroke_dasharray" valid:"-"`
		} `json:"body,omitempty" bson:"body" valid:"-"`
		Image struct {
			RefWidth            string `json:"ref_width,omitempty" bson:"ref_width" valid:"-"`
			RefHeight           int    `json:"ref_height,omitempty" bson:"ref_height" valid:"-"`
			X                   int    `json:"x,omitempty" bson:"x" valid:"-"`
			Y                   int    `json:"y,omitempty" bson:"y" valid:"-"`
			PreserveAspectRatio string `json:"preserve_aspect_ratio,omitempty" bson:"preserve_aspect_ratio" valid:"-"`
			XlinkHref           string `json:"xlink_href,omitempty" bson:"xlink_href" valid:"-"`
		} `json:"image,omitempty" bson:"image" valid:"-"`
		Label struct {
			TextVerticalAnchor string `json:"text_vertical_anchor" bson:"text_vertical_anchor" valid:"-"`
			TextAnchor         string `json:"text_anchor" bson:"text_anchor" valid:"-"`
			RefX               string `json:"ref_x" bson:"ref_x" valid:"-"`
			RefX2              int    `json:"ref_x2" bson:"ref_x2" valid:"-"`
			RefY               int    `json:"ref_y" bson:"ref_y" valid:"-"`
			FontSize           int    `json:"font_size" bson:"font_size" valid:"-"`
			Fill               string `json:"fill" bson:"fill" valid:"-"`
			TextWrap           struct {
				Text     string `json:"text" bson:"text" valid:"-"`
				Width    int    `json:"width" bson:"width" valid:"-"`
				Height   int    `json:"height" bson:"height" valid:"-"`
				Ellipsis bool   `json:"ellipsis" bson:"ellipsis" valid:"-"`
			} `json:"text_wrap,omitempty" bson:"text_wrap" valid:"-"`
			FontFamily  string  `json:"font_family" bson:"font_family" valid:"-"`
			FontWeight  string  `json:"font_weight" bson:"font_weight" valid:"-"`
			StrokeWidth float64 `json:"stroke_width" bson:"stroke_width" valid:"-"`
		} `json:"label,omitempty" bson:"label"valid:"-"`
		ServiceType struct {
			TextVerticalAnchor string `json:"text_vertical_anchor" bson:"text_vertical_anchor" valid:"-"`
			TextAnchor         string `json:"text_anchor" bson:"text_anchor" valid:"-"`
			RefX               string `json:"ref_x" bson:"ref_x" valid:"-"`
			RefX2              int    `json:"ref_x2" bson:"ref_x2" valid:"-"`
			RefY               int    `json:"ref_y" bson:"ref_y" valid:"-"`
			FontSize           int    `json:"font_size" bson:"font_size" valid:"-"`
			Fill               string `json:"fill" bson:"fill" valid:"-"`
			TextWrap           struct {
				Text     string `json:"text" bson:"text" valid:"-"`
				Width    int    `json:"width" bson:"width" valid:"-"`
				Height   int    `json:"height" bson:"height" valid:"-"`
				Ellipsis bool   `json:"ellipsis" bson:"ellipsis" valid:"-"`
			} `json:"text_wrap,omitempty" bson:"text_wrap" valid:"-"`
			FontFamily  string  `json:"font_family" bson:"font_family" valid:"-"`
			FontWeight  string  `json:"font_weight" bson:"font_weight" valid:"-"`
			StrokeWidth float64 `json:"c" bson:"stroke_width" valid:"-"`
		} `json:"service_type,omitempty" bson:"service_type" valid:"-"`
		Root struct {
			DataTooltip                 string `json:"data_tooltip" bson:"data_tooltip" valid:"-"`
			DataTooltipPosition         string `json:"data_tooltip_position" bson:"data_tooltip_position"valid:"-"`
			DataTooltipPositionSelector string `json:"data_tooltip_position_selector" bson:"data_tooltip_position_selector" valid:"-"`
		} `json:"root,omitempty" bson:"root" valid:"-"`
	} `json:"attrs,omitempty" bson:"attrs" valid:"-"`
	// required for Cloudplex Frontend
	// +optional
	Position Position `json:"position,omitempty" bson:"position" valid:"-" jsonschema:"-"`
	// required for Cloudplex Frontend
	// +optional
	Size struct {
		Width  int `json:"width" bson:"width" valid:"-"`
		Height int `json:"height" bson:"height" valid:"-"`
	} `json:"size,omitempty" bson:"size" valid:"-"`
	// required for Cloudplex Frontend
	// +optional
	Angle int `json:"angle,omitempty" bson:"angle" valid:"-"`
	//required for Cloudplex Frontend visualization
	// This key will tell to embed service or not
	// +optional
	IsEmbedded bool `json:"is_embedded,omitempty" bson:"is_embedded,omitempty"`
	// required for Cloudplex Frontend visualization
	// embeds contains list of service_ids which you want to group
	// +optional
	Embeds []string `json:"embeds,omitempty" bson:"embeds,omitempty"`
}
type HookMain struct {
	Weight     int64    `json:"weight" bson:"weight"`
	ServiceID  string   `json:"service_id,omitempty" bson:"service_id" binding:"required" valid:"alphanumspecial,length(6|100)~service_id must contain between 6 and 100 characters,lowercase~lowercase alphanumeric characters are allowed,required~service_id is missing in request"`
	HooksTypes []string `json:"hook_types,omitempty" bson:"hook_types"`
}
type ServiceTemplate struct {
	HookConfiguration *HookConfiguration `json:"hook_configuration,omitempty"`
	Hooks             []HookMain         `json:"hooks,omitempty" bson:"hooks,omitempty" valid:"-"`
	ServiceBasicInfo  `json:",inline" bson:",inline"`
	// ServiceAttributes are the actual attributes required to deploy
	// on Kubernetes Cluster
	// +mandatory
	ServiceAttributes interface{} `json:"service_attributes" bson:"service_attributes" binding:"required" valid:"-"`
}
type Position struct {
	X int `json:"x" bson:"x" valid:"-"`
	Y int `json:"y" bson:"y" valid:"-"`
}
type APIService struct {
	RequestType        string            `json:"request_type" bson:"request_type"`
	URL                string            `json:"url" bson:"url" `
	AuthenticationType string            `json:"auth_type" bson:"auth_type"`
	SecureRequest      bool              `json:"is_secure" bson:"is_secure"`
	Data               interface{}       `json:"data" bson:"data"`
	Headers            map[string]string `json:"headers" bson:"headers"`
	AuthCredentials    Authentication    `json:"credentials" bson:"credentials"`
}
type Authentication struct {
	Username string
	Password string
}
type Lable struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type DRSubsets struct {
	Name                     string  `json:"name"`
	Labels                   []Lable `json:"labels"`
	Http1MaxPendingRequests  int32   `json:"max_pending_requests"`
	Http2MaxRequests         int32   `json:"max_requests"`
	MaxRequestsPerConnection int32   `json:"max_requests_per_connection"`
	MaxRetries               int32   `json:"max_retries"`
}

type IstioDestinationRuleAttributes struct {
	Host    string      `json:"host"`
	Subsets []DRSubsets `json:"subsets"`
}
type VSDestination struct {
	Host   string `json:"host"`
	Subset string `json:"subset"`
	Port   int    `json:"port"`
}
type VSRetries struct {
	Attempts int `json:"attempts"`
	Timeout  int `json:"per_request_timeout"`
}
type VSRoute struct {
	Destination VSDestination `json:"destination"`
	Weight      int32         `json:"weight"`
}
type VSHTTP struct {
	Routes []VSRoute `json:"route"`
	//RewriteUri string      `json:"rewrite_uri"`
	//RetriesUri string      `json:"retries_uri"`
	Timeout        int32          `json:"timeout"`
	Retries        []VSRetries    `json:"retries"`
	Match          []URI          `json:"match"`
	FaultInjection FaultInjection `json:"fault_injection"`
}
type URI struct {
	URIS []string `json:"uri"`
}
type IstioVirtualServiceAttributes struct {
	Hosts    []string `json:"hosts"`
	Gateways []string `json:"gateways"`
	HTTP     []VSHTTP `json:"http"`
}
type ServiceAttributes struct {
	Weight                   int32               `json:"weight"`
	TimeOut                  int32               `json:"time_out"`
	Retries                  VSRetries           `json:"retries"`
	URIS                     []string            `json:"uri"`
	Port                     int                 `json:"vs_port"`
	Enable_External_Traffic  bool                `json:"enable_external_traffic"`
	Http1MaxPendingRequests  int32               `json:"max_pending_requests"`
	Http2MaxRequests         int32               `json:"max_requests"`
	MaxRequestsPerConnection int32               `json:"max_requests_per_connection"`
	MaxRetries               int32               `json:"max_retries"`
	FaultInjectionAbort      FaultInjectionAbort `json:"fault_abort"`
	FaultInjectionDelay      FaultInjectionDelay `json:"fault_delay"`
	ScalingEnable            bool                `json:"enable_scaling"`
	HPA                      HPAAttributes       `json:"hpa_configurations"`
}
type FaultInjection struct {
	FaultInjectionAbort FaultInjectionAbort `json:"fault_abort"`
	FaultInjectionDelay FaultInjectionDelay `json:"fault_delay"`
}
type FaultInjectionAbort struct {
	Percentage float64 `json:"percentage"`
	HttpStatus int32   `json:"http_status"`
}
type FaultInjectionDelay struct {
	Percentage float64 `json:"percentage"`
	FixedDelay int64   `json:"fix_delay"`
}

type HPAAttributes struct {
	Type               string         `json:"type"`
	MixReplicas        int32          `json:"min_replicas"`
	MaxReplicas        int32          `json:"max_replicas"`
	Metrics_           []Metrics      `json:"metrics_values"`
	CrossObjectVersion ScaleTargetRef `json:"cross_object_version"`
}
type Metrics struct {
	TargetValueKind string `json:"target_value_kind"`
	TargetValue     int64  `json:"target_value"`
	TargetValueUnit string `json:"target_value_unit"`
	ResourceKind    string `json:"resource_kind"`
}

type ScaleTargetRef struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Type    string `json:"type"`
}
