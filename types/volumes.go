package types

type Cloud string

const (
	AWS   Cloud = "aws"
	Azure Cloud = "azure"
	GCP   Cloud = "gcp"
)

/*
type Parameters struct {
	Type            string `json:"type"`
	ReplicationType string `json:"replication_type"`
	Iops            string `json:"iops"`
	Plugin          string `json:"plugin"`
	SkuName         string `json:"sku_name"`
	Location        string `json:"location"`
	StorageAccount  string `json:"storage_account"`
}

type Volume struct {
	Name      string     `json:"name"`
	Size      int64      `json:"size"`
	Cloud     string     `json:"cloud"`
	Namespace string     `json:"-"`
	MountPath string     `json:"mount_path"`
	Params    Parameters `json:"params"`
}
*/
type KeytoPath struct {
	Key  string `json:"key"`
	Path string `json:"path"`
	Mode int32  `json:"mode"`
}

type SecretVolumeSource struct {
	Name        string      `json:"name"`
	Items       []KeytoPath `json:"items"`
	DefaultMode int32       `json:"default_mode"`
}

type ContainerVolumeMount struct {
	Name      string `json:"name"`
	ReadOnly  bool   `json:"read_only"`
	MountPath string `json:"mount_path"`
}

type ContainerVolume struct {
	Name               string              `json:"name"`
	ContainerSecret    *SecretVolumeSource `json:"container_secret"`
	ContainerConfigMap *SecretVolumeSource `json:"container_config_map"`
}
