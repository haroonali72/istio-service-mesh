package types

type LegacyServiceAttributes struct {
	IP       string        `json:"ip" bson:"ip"`
	Ports    []PortsAttrib `json:"ports" bson:"ports"`
	FileName string        `json:"file_name" bson:"file_name"`
	URL      string        `json:"url" bson:"url"`
}
type PortsAttrib struct {
	Protocol string `json:"protocol"`
	Number   int32  `json:"number"`
	Name     string `json:"name"`
}
