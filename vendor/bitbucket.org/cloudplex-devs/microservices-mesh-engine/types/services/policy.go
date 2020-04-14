package services

type PolicyService struct {
	Id                    interface{}             `json:"_id,omitempty" bson:"_id" valid:"-"`
	ServiceId             string                  `bson:"service_id" json:"service_id",valid:"required"`
	CompanyId             string                  `bson:"company_id" json:"company_id",valid:"required"`
	Name                  string                  `bson:"name" json:"name",valid:"required"`
	Version               string                  `bson:"version" json:"version",valid:"required"`
	ServiceType           string                  `bson:"service_type" json:"service_type",valid:"required"`
	ServiceSubType        string                  `bson:"service_sub_type" json:"service_sub_type",valid:"required"`
	ServiceDependencyInfo interface{}             `bson:"service_dependency_info" json:"service_dependency_info",valid:"required"`
	Namespace             string                  `bson:"namespace" json:"namespace",valid:"required"`
	ServiceAttributes     *PolicyServiceAttribute `bson:"service_attributes" json:"service_attributes",valid:"required"`
}

type PolicyServiceAttribute struct {
	Peers            []*PeerAuthenticationMethod   `bson:"peers" json:"peers"`
	Targets          []*TargetSelector             `bson:"target" json:"target"`
	PeerIsOptional   bool                          `bson:"peer_is_optional" json:"is_optional"`
	Origins          []*OriginAuthenticationMethod `bson:"origin" json:"origin"`
	OriginIsOptional bool                          `bson:"origin_is_optional" json:"origin_is_optional"`
	PrincipalBinding string                        `bson:"principal_binding" json:"principal_binding" valid:"in(USE_PEER|USE_ORIGIN)"`
}
type PeerAuthenticationMethod struct {
	Mtls *MutualTls `bson:"mtls" json:"mtls"`
	Jwt  *Jwt       `bson:"jwt" json:"jwt"`
}
type MutualTls struct {
	AllowTls bool   `bson:"allow_tls" json:"allow_tls"`
	Mode     string `bson:"mode" json:"mode" valid:"in(STRICT|PERMISSIVE)"`
}
type Jwt struct {
	Issuer       string            `bson:"issuer" json:"issuer"`
	Audiences    []string          `bson:"audiences" json:"audiences"`
	JwksUri      string            `bson:"jwks_uri" json:"jwks_uri"`
	Jwks         string            `bson:"jwks" json:"jwks"`
	JwtHeaders   []string          `bson:"jwt_headers" json:"jwt_header"`
	JwtParams    []string          `bson:"jwt_params" json:"jwt_params"`
	TriggerRules []*JwtTriggerRule `bson:"trigger_rules" json:"trigger_rules"`
}
type JwtTriggerRule struct {
	ExcludedPaths []*ExcludedPath `bson:"excluded_path" json:"excluded_path"`
	IncludedPaths []*IncludedPath `bson:"included_path" json:"included_path"`
}
type IncludedPath struct {
	Type  string `bson:"type" json:"type" valid:"in(Exact/Prefix/Suffix/Regex)"`
	Value string `bson:"value" json:"value"`
}
type ExcludedPath struct {
	Type  string `bson:"type" json:"type" valid:"in(Exact/Prefix/Suffix/Regex)"`
	Value string `bson:"value" json:"value"`
}

type OriginAuthenticationMethod struct {
	Jwt *Jwt `bson:"jtw" json:"jtw"`
}

type TargetSelector struct {
	Name  string          `bson:"name" json:"name"`
	Ports []*PortSelector `bson:"ports" json:"ports"`
}
type PortSelector struct {
	Number int    `bson:"number" json:"number"`
	Name   string `bson:"name" json:"name"`
}
