package types

type DockerService struct {
	Id                    *string                  `bson:"_id" json:"_id"`
	Name                  *string                  `bson:"name" json:"name"`
	Registered_Name       *string                  `bson:"registered_name" json:"registered_name"`
	Version               *string                  `bson:"version" json:"version"`
	ImagePrefix           *string                  `bson:"image_prefix" json:"image_prefix"`
	ImageName             *string                  `bson:"image_name" json:"image_name"`
	Tag                   *string                  `bson:"tag" json:"tag"`
	CRTag                 string                   `bson:"cr_tag" json:"cr_tag"`
	Dockerfile            *string                  `bson:"docker_file_contents" json:"docker_file_contents"`
	Built                 *string                  `bson:"status" json:"status"`
	Instances             *int                     `bson:"instances,omitempty" json:"instances,omitempty"`
	ServiceLabel          map[string]string        `bson:"service_label,omitempty" json:"service_label,omitempty"`
	KubeNodeSelector      map[string]string        `bson:"node_selector,omitempty" json:"node_selector,omitempty"`
	Env_Vars              map[string]string        `bson:"environment_variables,omitempty" json:"environment_variables,omitempty"`
	Ports                 map[string]string        `bson:"ports,omitempty" json:"ports,omitempty"`
	AdditionalPorts       []DockerPorts            `bson:"additional_ports,omitempty" json:"additional_ports,omitempty"`
	PortRange             DockerPortsRange         `bson:"port_range,omitempty" json:"port_range,omitempty"`
	HighAvailability      bool                     `bson:"high_availability,omitempty" json:"high_availability,omitempty"`
	VolumesMountList      []map[string]string      `bson:"host_volume_mount,omitempty" json:"host_volume_mount,omitempty"`
	ImageRegistry         DockerRegistry           `bson:"image_registry" json:"image_registry"`
	BaseImageRegistry     DockerRegistry           `bson:"base_image_registry" json:"base_image_registry"`
	ExternalVolumes       []ExternalVolume         `json:"external_volume"`
	Default_Configuration []FilesAttributes        `bson:"default_configurations" json:"default_configurations"`
	Ingress               []map[string]string      `bson:"ingress_support" json:"ingress_support"`
	LoadBalancer          bool                     `bson:"load_balancer" json:"load_balancer"`
	CPU                   Resource                 `bson:"cpu" json:"cpu"`
	Memory                Resource                 `bson:"memory" json:"memory"`
	CMD                   string                   `bson:"commands" json:"commands"`
	ARGS                  []string                 `bson:"arguments" json:"arguments"`
	SecurityContext       ContainerSecurityContext `bson:"security_context,omitempty" json:"security_context,omitempty" `
	LogsPath              string                   `json:"logs_path" bson:"logs_path"`
}

type DockerPorts struct {
	NodePort      int    `json:"node_port"`
	HostPort      int    `json:"host_port"`
	ContainerPort int    `json:"container_port"`
	ServicePort   int    `json:"service_port"`
	PortType      string `json:"port_type"`
}
type DockerPortsRange struct {
	//NodePort      int    `json:"node_port"`
	PortUpper int    `json:"port_upper"`
	PortLower int    `json:"port_lower"`
	PortType  string `json:"range_type"`
	HostPort  bool   `json:"host_range"`
}

// TODO use some interface, to support multiple cloudproviders
type ExternalVolume struct {
	MountPoint            string `json:"volume_mount"`
	IOps                  int    `json:"iops"`
	VolumeType            string `json:"volume_type"`
	VolumeSize            int    `json:"ebs_volumes_size"`
	Ebs_type              string `json:"ebs_type"`
	VolumeID              string `json:"volume_id"`
	Delete_on_termination bool   `json:"delete_on_termination"`
	Encryption            bool   `json:"volume_encryption"`
}
type Quota struct {
	Value int    `json:"value"`
	Unit  string `json:"unit"`
}

type Resource struct {
	Limit   Quota `json:"limit"`
	Request Quota `json:"request"`
}

type DockerRegistry struct {
	ServerAddress string `json:"registry_url"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	Email         string `json:"email"`
}
type Service_Attrib struct {
	Default_Configuration []FilesAttributes `bson:"default_configurations"`
}
type FilesAttributes struct {
	FileName      string `bson:"file_name" json:"file_name"`
	FilePath      string `bson:"path" json:"path"`
	FileType      string `bson:"file_type" json:"file_type"`
	FileContentes string `bson:"file_contents" json:"file_contents"`
}
type ServiceStub struct {
	Id     string `bson:"id" json:"_id"`
	Name   string `bson:"name" json:"name"`
	Number int    `bson:"number_of_instances" json:"number_of_instances"`
	Type   string `bson:"type,omitempty" json:"type,omitempty"`
}

type ContainerSecurityContext struct {
	CapabilitiesAdd        []string `bson:"capabilities_add" json:"capabilities_add"`
	CapabilitiesDrop       []string `bson:"capabilities_drop" json:"capabilities_drop"`
	RunAsUser              int      `bson:"run_as_user" json:"run_as_user"`
	Privileged             bool     `bson:"privileged" json:"privileged"`
	RunAsNonRoot           bool     `bson:"run_as_non_root" json:"run_as_non_root"`
	ReadOnlyRootFilesystem bool     `bson:"read_only_root_filesystem" json:"read_only_root_filesystem"`
}
