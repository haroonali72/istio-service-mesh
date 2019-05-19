package types

type Cloud string

const (
	AWS   Cloud = "aws"
	Azure Cloud = "azure"
	GCP   Cloud = "gcp"
)

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
