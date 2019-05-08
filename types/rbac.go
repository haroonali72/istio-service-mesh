package types

type Role struct {
	Resource    string   `json:"resource"`
	Verbs       []string `json:"verbs"`
	ServiceName string   `json:"service"`
	Namespace   string   `json:"namespace"`
	ApiGroup    string   `json:"api_group"`
}

type RoleBinding struct {
	ServiceName string `json:"service"`
	Namespace   string `json:"namespace"`
}
type ServiceAccount struct {
	ServiceName string `json:"service"`
	Namespace   string `json:"namespace"`
}
