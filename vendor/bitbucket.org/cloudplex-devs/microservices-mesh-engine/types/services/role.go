package services

type Role struct {
	ServiceId             string           `bson:"service_id" json:"service_id",valid:"required"`
	CompanyId             string           `bson:"company_id" json:"company_id",valid:"required"`
	Name                  string           `bson:"name" json:"name",valid:"required"`
	Version               string           `bson:"version" json:"version",valid:"required"`
	ServiceType           string           `bson:"service_type" json:"service_type",valid:"required"`
	ServiceSubType        string           `bson:"service_sub_type" json:"service_sub_type",valid:"required,in(role)"`
	ServiceDependencyInfo interface{}      `bson:"service_dependency_info" json:"service_dependency_info",valid:"required"`
	Namespace             string           `bson:"namespace" json:"namespace",valid:"required"`
	ServiceAttributes     ServiceAttribute `bson:"service_attributes" json:"service_attributes",valid:"required"`
}

type ServiceAttribute struct {
	Rules []Rule `bson:"rules" json:"rules"`
}

type Rule struct {
	Resources []string `bson:"resources" json:"resources"`
	Verbs     []string `bson:"verbs" json:"verbs"`
	Api_group []string `bson:"api_groups" json:"api_groups"`
}
