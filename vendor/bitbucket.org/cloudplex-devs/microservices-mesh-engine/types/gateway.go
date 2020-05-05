package types

type Gateway struct {
	ServiceID         string                   `json:"service_id"`
	Name              string                   `json:"name"`
	Namespace         string                   `json:"namespace"`
	ServiceType       string                   `json:"service_type"`
	ServiceSubType    string                   `json:"service_sub_type"`
	CompanyID         string                   `json:"company_id"`
	ServiceAttributes GatewayServiceAttributes `json:"service_attributes"`
}

type Port struct {
	Name     string `json:"name"`
	Number   uint32 `json:"number"`
	Protocol string `json:"protocol"`
}
type TLS struct {
	HTTPSRedirect      bool     `json:"https_redirect"`
	Mode               string   `json:"mode"`
	ServerCertificate  string   `json:"server_certificate"`
	PrivateKey         string   `json:"private_key"`
	CaCertificates     string   `json:"ca_certificates"`
	SubjectAltNames    []string `json:"subject_alt_names"`
	MinProtocolVersion string   `json:"min_protocol_version"`
	MaxProtocolVersion string   `json:"max_protocol_version"`
}
type Servers struct {
	Port  Port     `json:"port"`
	Hosts []string `json:"hosts"`
	TLS   TLS      `json:"tls"`
}
type GatewayServiceAttributes struct {
	Selectors map[string]string `json:"selectors"`
	Servers   []Servers         `json:"servers"`
}
