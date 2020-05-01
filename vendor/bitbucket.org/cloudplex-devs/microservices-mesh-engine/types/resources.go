package types

type Resources struct {
	CompanyId string `bson:"company_id" json:"company_id"`
	ServiceId string `bson:"service_id" json:"service_id"`
	Type      string `bson:"type" json:"type"`
	Name      string `bson:"name" json:"name"`
}
