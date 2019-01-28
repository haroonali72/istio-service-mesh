package types


type Route struct {
	Host      string `json:"host"`
	Subset string `json:"subset"`
}
type Port struct {

	Host      string `json:"host"`
	Container string `json:"container"`
}

type SEPort struct {
	Name      string `json:"name"`
	Port      int32 `json:"port"`
	Protocol string `json:"protocol"`
}
type SEEndpoints struct {
	Address      string `json:"address"`
}
type VSDestination struct {

}
type VSRetries struct {
	Attempts           int         `json:"attempts"`
	Timeout           int         `json:"timeouts"`
}
type VSRoute struct {
	Host           string         `json:"host"`
	Subset           string         `json:"subset"`
	Port           int         `json:"port"`
	Weight           int32         `json:"weight"`
}
type VSHTTP struct {
	Routes           []VSRoute         `json:"routes"`
	RewriteUri       string         `json:"rewrite_uri"`
	RetriesUri	     string         `json:"retries_uri"`
	Timeout          int32         `json:"timeout"`
	Retries          []VSRetries         `json:"retries"`
}

type IstioVirtualServiceAttributes struct {
	Hosts           []string         `json:"hosts"`
	Gateways         []string         `json:"gateway"`
	HTTP           []VSHTTP         `json:"http"`
}

type IstioServiceEntryAttributes struct {
	Hosts           []string         `json:"hosts"`
	Address         []string         `json:"address"`
	Ports           []SEPort         `json:"ports"`
	Uri				[]SEEndpoints `json:"endpoints"`
	Location      	string           `json:"location"`
	Resolution      string           `json:"resolution"`
}
type GWServers struct {
	Hosts           []string         `json:"hosts"`
	Labels         []string         `json:"labels"`
	Port      string           `json:"port"`
	Protocol      string           `json:"protocol"`
	Name      string           `json:"name"`
}
type IstioGatewayAttributes struct {
	Servers    []GWServers			`json:"servers"`
	Selector  map[string]string `json:"selector"`
}
type DRSubsets struct {
	Name      string           `json:"name"`
	Labels  map[string]string `json:"labels"`

}
type IstioDestinationRuleAttributes struct {
	Host           string         `json:"host"`
	Subsets        []DRSubsets         `json:"subsets"`

}
type DockerServiceAttributes struct {
	DistributionType      string            `json:"distribution_type"`
	DefaultConfigurations string            `json:"default_configurations"`
	EnvironmentVariables  map[string]string `json:"environment_variables"`
	Ports                 []Port            `json:"ports"`
	Files                 []string          `json:"files"`
	Tag                   string            `json:"tag"`
	ImagePrefix           string            `json:"image_prefix"`
	ImageName             string            `json:"image_name"`
}


type ServiceDependency struct {
	Name              string            `json:"name"`
	DependencyType    string            `json:"dependency_type"`
	Hosts           []string            `json:"hosts"`
	Uri             []string            `json:"uri"`
	TimeOut           string            `json:"timeout"`
	Routes           []Route            `json:"routes"`
	Ports             []Port            `json:"ports"`
}
/*
type ServiceDependencyx struct {
	ServiceType       string            `json:"service_type"`
	Name              string            `json:"name"`
	Version           string `json:"version"`
	ServiceAttributes ServiceAttributes `json:"service_attributes"`
}*/
type Service struct {
	ServiceType           string            `json:"service_type"`
	SubType           string            `json:"service_sub_type"`
	Name                  string            `json:"name"`
	Version               string            `json:"version"`
	ServiceDependencyInfo []ServiceDependency `json:"service_dependency_info"`
	ServiceAttributes     interface{} 		`json:"service_attributes"`
	Namespace             string            `json:"namespace"`
	Hostnames             []string            `json:"hostnames"`

}
type SolutionInfo struct {
	Name                  string            `json:"name"`
	Version               string            `json:"version"`
	PoolId               string            `json:"pool_id"`
	Service []Service `json:"services"`
	KubernetesIp string            `json:"kubeip"`
	KubernetesPort string            `json:"kubeport"`
	KubernetesUsername string            `json:"kubeusername"`
	KubernetesPassword string            `json:"kubepassword"`
}

type ServiceInput struct {
	KubernetesPassword string            `json:"cluster_id"`
	SolutionInfo SolutionInfo `json:"solution_info"`
}
