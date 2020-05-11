package services

import (
	"bitbucket.org/cloudplex-devs/microservices-mesh-engine/types"
)

//Id                interface{}               `json:"_id,omitempty" bson:"_id" valid:"-"`
//ServiceId         string                    `json:"service_id" bson:"service_id" binding:"required"`
//Name              string                    `json:"name"  bson:"name" binding:"required" `
//Version           string                    `json:"version"  bson:"version"  binding:"required"`
//ServiceType       constants.ServiceType     `json:"service_type"  bson:"service_type" valid:"-"`
//ServiceSubType    constants.ServiceSubType  `json:"service_sub_type" bson:"service_type" valid:"-"`
//Namespace         string                    `json:"namespace" bson:"namespace"`
//CompanyId         string                    `json:"company_id,omitempty" bson:"company_id"`
//CreationDate      time.Time                 `json:"creation_date,omitempty" bson:"creation_date" valid:"-"`
type GatewayService struct {
	types.ServiceBasicInfo `json:",inline" bson:",inline"`
	ServiceAttributes      *GatewayServiceAttributes `json:"service_attributes"  bson:"service_attributes" binding:"required"`
}
type GatewayServiceAttributes struct {
	// One or more labels that indicate a specific set of pods/VMs
	// on which this gateway configuration should be applied. default: istio: ingressgateway
	Selectors map[string]string `json:"selectors" bson:"selectors"`
	// A list of server specifications.
	Servers []*Server `json:"servers" bson:"servers" binding:"required"`
}
type Server struct {
	// The Port on which the proxy should listen for incoming connections.
	Port *Port `json:"port" bson:"port" binding:"required"`
	// One or more hosts exposed by this gateway.
	// While typically applicable to
	// HTTP services, it can also be used for TCP services using TLS with SNI.
	// A host is specified as a dnsName with an optional namespace/ prefix.
	// The dnsName should be specified using FQDN format
	// Set the dnsName to * to select all VirtualService hosts from the
	// specified namespace (e.g.,prod/*).
	//
	// +optional
	Hosts []string `json:"hosts" bson:" hosts" `
	// Set of TLS related options that govern the server's behavior. Use
	// these options to control if all http requests should be redirected to
	// https, and the TLS modes to use.
	Tls *TlsConfig `json:"tls,omitempty" bson:"tls,omitempty"`
}
type Port struct {
	// Label assigned to the port.
	// default value is http
	// +optional
	Name string `json:"name" bson:"name" binding:"required" default:"http"`
	// A valid non-negative integer port number.
	// default value is 80
	// +optional
	Number uint32 `json:"number,omitempty" bson:"number,omitempty" default:"80" jsonschema:"minimum=0,maximum=65536"`
	//supported protocols HTTP|HTTPS|GRPC|HTTP2|MONGO|TCP|TLS
	// default: HTTP
	// +optional
	Protocol Protocols `json:"protocol" bson:"protocol" binding:"required" swaggerType:"string" jsonschema:"enum=HTTP,enum=HTTPS,enum=GRPC,enum=MONGO,enum=TCP,enum=TLS,default=HTTP" default:"HTTP"`
}

type TlsConfig struct {
	// If set to true, the load balancer will send a 301 redirect for
	// all http connections, asking the clients to use HTTPS. Not
	// applicable in Sidecar API.
	// default value is false
	// +optional
	HttpsRedirect bool `json:"https_redirect" bson:"https_redirect" default:"false"`
	//supported modes  PASSTHROUGH|SIMPLE|MUTUAL|AUTO_PASSTHROUGH|ISTIO_MUTUAL
	// +optional
	Mode Mode `json:"mode,omitempty" bson:"mode,omitempty" swaggerType:"string" jsonschema:"enum=PASSTHROUGH,enum=SIMPLE,enum=MUTUAL,enum=AUTO_PASSTHROUGH,enum=ISTIO_MUTUAL"`
	// REQUIRED if mode is SIMPLE or MUTUAL. The path to the file
	// holding the server-side TLS certificate to use.
	// +optional
	ServerCertificate string `json:"server_certificate,omitempty" bson:"server_certificate,omitempty"`
	// REQUIRED if mode is SIMPLE or MUTUAL. The path to the file
	// holding the server's private key.
	// +optional
	PrivateKey string `json:"private_key,omitempty" bson:"private_key,omitempty"`
	// REQUIRED if mode is MUTUAL. The path to a file containing
	// certificate authority certificates to use in verifying a presented
	// client side certificate.
	// +optional
	CaCertificate string `json:"ca_certificate,omitempty" bson:"ca_certificate,omitempty"`
	// A list of alternate names to verify the subject identity in the
	// certificate presented by the client.
	// +optional
	SubjectAltName []string `json:"subject_alt_name,omitempty" bson:"subject_alt_name,omitempty"`
	//Minimum TLS protocol version. supported values  TLS_AUTO|TLSV1_0|TLSV1_1|TLSV1_2|TLSV1_3
	// +optional
	MinProtocolVersion ProtocolVersion `json:"min_protocol_version,omitempty" bson:"min_protocol_version,omitempty" swaggerType:"string"  jsonschema:"enum=TLS_AUTO,enum=TLSV1_0,enum=TLSV1_2,enum=TLSV1_3" `
	//Maximum TLS protocol version. supported values  TLS_AUTO|TLSV1_0|TLSV1_1|TLSV1_2|TLSV1_3
	// +optional
	MaxProtocolVersion ProtocolVersion `json:"max_protocol_Version,omitempty" bson:"max_protocol_version,omitempty" swaggerType:"string" jsonschema:"enum=TLS_AUTO,enum=TLSV1_0,enum=TLSV1_2,enum=TLSV1_3" `
}

type Protocols string

const (
	Protocols_HTTP  Protocols = "HTTP"
	Protocols_HTTPS Protocols = "HTTPS"
	Protocols_GRPC  Protocols = "GRPC"
	Protocols_HTTP2 Protocols = "HTTP2"
	Protocols_MONGO Protocols = "MONGO"
	Protocols_TCP   Protocols = "TCP"
	Protocols_TLS   Protocols = "TLS"
)

func (p *Protocols) String() string {
	return string(*p)
}

type Mode string

const (
	Mode_PASSTHROUGH      Mode = "PASSTHROUGH"
	Mode_SIMPLE           Mode = "SIMPLE"
	Mode_MUTUAL           Mode = "MUTUAL"
	Mode_AUTO_PASSTHROUGH Mode = "AUTO_PASSTHROUGH"
	Mode_ISTIO_MUTUAL     Mode = "ISTIO_MUTUAL"
)

func (m *Mode) String() string {
	return string(*m)
}

type ProtocolVersion string

const (
	ProtocolVersion_TLS_AUTO ProtocolVersion = "TLS_AUTO"
	ProtocolVersion_TLSV1_0  ProtocolVersion = "TLSV1_0"
	ProtocolVersion_TLSV1_1  ProtocolVersion = "TLSV1_1"
	ProtocolVersion_TLSV1_2  ProtocolVersion = "TLSV1_2"
	ProtocolVersion_TLSV1_3  ProtocolVersion = "TLSV1_3"
)

func (p *ProtocolVersion) String() string {
	return string(*p)
}
