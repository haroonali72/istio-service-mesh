package types

import "time"

type GatewayService struct {
	Id                interface{}               `json:"_id,omitempty" bson:"_id" valid:"-"`
	ServiceId         string                    `json:"service_id" bson:"service_id" binding:"required" valid:"alphanumspecial,length(4|30)~service_id must contain between 6 and 30 characters,lowercase~lowercase alphanumeric characters are allowed,required~service_id is missing in request"`
	Name              string                    `json:"name"  bson:"name" binding:"required" valid:"alphanumspecial,length(3|30),lowercase~lowercase alphanumeric characters are allowed,required"`
	Version           string                    `json:"version"  bson:"version"  binding:"required" valid:"alphanumspecial,length(1|10),lowercase~lowercase alphanumeric characters are allowed,required"`
	ServiceType       string                    `json:"service_type"  bson:"service_type" valid:"-"`
	ServiceSubType    string                    `json:"service_sub_type" bson:"service_type" valid:"-"`
	Namespace         string                    `json:"namespace" bson:"namespace" binding:"required" valid:"alphanumspecial,required"`
	CompanyId         string                    `json:"company_id,omitempty" bson:"company_id"`
	CreationDate      time.Time                 `json:"creation_date,omitempty" bson:"creation_date" valid:"-"`
	ServiceAttributes *GatewayServiceAttributes `json:"service_attributes"  bson:"company_id" binding:"required"`
}
type GatewayServiceAttributes struct {
	// One or more labels that indicate a specific set of pods/VMs
	// on which this gateway configuration should be applied. default: istio: ingressgateway
	Selectors map[string]string `json:"selectors" bson:"selectors" binding:"required"`
	// A list of server specifications.
	Servers []*Server `json:"servers" bson:"servers" binding:"required"`
}
type Server struct {
	// The Port on which the proxy should listen for incoming connections.
	Port  *Port    `json:"port" bson:"port" binding:"required"`
	Hosts []string `json:"hosts" bson:" hosts" binding:"required"`
	// Set of TLS related options that govern the server's behavior. Use
	// these options to control if all http requests should be redirected to
	// https, and the TLS modes to use.
	Tls *TlsConfig `json:"tls,omitempty" bson:"tls,omitempty"`
}
type Port struct {
	// Label assigned to the port.
	Name string `json:"name" bson:"name" binding:"required"`
	// A valid non-negative integer port number.
	Nummber uint32 `json:"nummber" bson:"nummber" binding:"required"`
	//supported protocols HTTP|HTTPS|GRPC|HTTP2|MONGO|TCPTLS
	Protocol Protocols `json:"protocol" bson:"protocol" binding:"required"`
}

type TlsConfig struct {
	HttpsRedirect bool `json:"httpsRedirect" bson:"httpsRedirect"`
	//supported modes  PASSTHROUGH|SIMPLE|MUTUAL|AUTO_PASSTHROUGH|ISTIO_MUTUAL
	Mode Mode `json:"mode,omitempty" bson:"mode,omitempty"`
	// REQUIRED if mode is SIMPLE or MUTUAL. The path to the file
	// holding the server-side TLS certificate to use.
	ServerCertificate string `json:"serverCertificate,omitempty" bson:"serverCertificate,omitempty"`
	// REQUIRED if mode is SIMPLE or MUTUAL. The path to the file
	// holding the server's private key.
	PrivateKey string `json:"privateKey,omitempty" bson:"privateKey,omitempty"`
	// REQUIRED if mode is MUTUAL. The path to a file containing
	// certificate authority certificates to use in verifying a presented
	// client side certificate.
	CaCertificate string `json:"caCertificate,omitempty" bson:"caCertificate,omitempty"`
	// A list of alternate names to verify the subject identity in the
	// certificate presented by the client.
	SubjectAltName []string `json:"subjectAltName,omitempty" bson:"subjectAltName,omitempty"`
	//Minimum TLS protocol version. supported values  TLS_AUTO|TLSV1_0|TLSV1_1|TLSV1_2|TLSV1_3
	MinProtocolVersion ProtocolVersion `json:"minProtocolVersion,omitempty" bson:"minProtocolVersion,omitempty"`
	//Maximum TLS protocol version. supported values  TLS_AUTO|TLSV1_0|TLSV1_1|TLSV1_2|TLSV1_3
	MaxProtocolVersion ProtocolVersion `json:"maxProtocolVersion,omitempty" bson:"maxProtocolVersion,omitempty"`
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

type Mode string

const (
	Mode_PASSTHROUGH      Mode = "PASSTHROUGH"
	Mode_SIMPLE           Mode = "SIMPLE"
	Mode_MUTUAL           Mode = "MUTUAL"
	Mode_AUTO_PASSTHROUGH Mode = "AUTO_PASSTHROUGH"
	Mode_ISTIO_MUTUAL     Mode = "ISTIO_MUTUAL"
)

type ProtocolVersion string

const (
	ProtocolVersion_TLS_AUTO ProtocolVersion = "TLS_AUTO"
	ProtocolVersion_TLSV1_0  ProtocolVersion = "TLSV1_0"
	ProtocolVersion_TLSV1_1  ProtocolVersion = "TLSV1_1"
	ProtocolVersion_TLSV1_2  ProtocolVersion = "TLSV1_2"
	ProtocolVersion_TLSV1_3  ProtocolVersion = "TLSV1_3"
)
