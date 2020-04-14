package services

import "time"

type PersistentVolumeService struct {
	Id                interface{}                       `json:"_id,omitempty" bson:"_id" valid:"-"`
	ServiceId         string                            `json:"service_id" bson:"service_id" binding:"required" valid:"alphanumspecial,length(4|30)~service_id must contain between 6 and 30 characters,lowercase~lowercase alphanumeric characters are allowed,required~service_id is missing in request"`
	Name              string                            `json:"name"  bson:"name" binding:"required" valid:"alphanumspecial,length(3|30),lowercase~lowercase alphanumeric characters are allowed,required"`
	Version           string                            `json:"version"  bson:"version"  binding:"required" valid:"alphanumspecial,length(1|10),lowercase~lowercase alphanumeric characters are allowed,required"`
	ServiceType       string                            `json:"service_type"  bson:"service_type" valid:"-"`
	ServiceSubType    string                            `json:"service_sub_type" bson:"service_type" valid:"-"`
	CompanyId         string                            `json:"company_id,omitempty" bson:"company_id"`
	CreationDate      time.Time                         `json:"creation_date,omitempty" bson:"creation_date" valid:"-"`
	ServiceAttributes *PersistentVolumeServiceAttribute `json:"service_attributes"  bson:"company_id" binding:"required"`
}
type PersistentVolumeServiceAttribute struct {
	Labels                 map[string]string       `json:"labels,omitempty"`
	ReclaimPolicy          ReclaimPolicy           `json:"reclaim_policy,omitempty"`
	PersistentVolumeSource *PersistentVolumeSource `json:"persistent_volume_source,omitempty"`
	AccessMode             []AccessMode            `json:"access_mode"`
	Capacity               string                  `json:"capacity,omitempty"`
	StorageClassName       string                  `json:"storage_class_name,omitempty"`
	MountOptions           []string                `json:"mount_options,omitempty"`
	VolumeMode             *PersistentVolumeMode   `json:"volume_mode,omitempty" protobuf:"bytes,8,opt,name=volumeMode,casttype=PersistentVolumeMode"`
	NodeAffinity           *VolumeNodeAffinity     `json:"node_affinity,omitempty"`
}
type VolumeNodeAffinity struct {
	Required NodeSelector `json:"required,omitempty"`
}

// PersistentVolumeMode describes how a volume is intended to be consumed, either Block or Filesystem.
type PersistentVolumeMode string

const (
	// PersistentVolumeBlock means the volume will not be formatted with a filesystem and will remain a raw block device.
	PersistentVolumeBlock PersistentVolumeMode = "Block"
	// PersistentVolumeFilesystem means the volume will be or is formatted with a filesystem.
	PersistentVolumeFilesystem PersistentVolumeMode = "Filesystem"
)

type AccessMode string

const (
	AccessModeReadWriteOnce AccessMode = "ReadWriteOnce"
	AccessModeReadWriteMany AccessMode = "ReadWriteMany"
	AccessModeReadOnlyMany  AccessMode = "ReadOnlyMany"
)

type PersistentVolumeSource struct {
	GCPPD     *GCPPD     `json:"gcp_pd,omitempty"`
	AWSEBS    *AWSEBS    `json:"aws_ebs,omitempty"`
	AzureDisk *AzureDisk `json:"azure_disk,omitempty"`
	AzureFile *AzureFile `json:"azure_file,omitempty"`
}

type GCPPD struct {
	PdName     string `json:"pd_name"`
	Filesystem string `json:"file_system"`
	Partition  int    `json:"partition"`
	ReadOnly   bool   `json:"readonly"`
}

type AWSEBS struct {
	VolumeId   string `json:"volume_id"`
	Filesystem string `json:"file_dystem"`
	Partition  int    `json:"partition"`
	ReadOnly   bool   `json:"readonly"`
}

type AzureDisk struct {
	CachingMode AzureDataDiskCachingMode `json:"caching_mode,omitempty"`
	Filesystem  string                   `json:"file_system"`
	ReadOnly    bool                     `json:"readonly"`
	DiskName    string                   `json:"disk_name"`
	DiskURI     string                   `json:"disk_uri"`
	Kind        AzureDataDiskKind        `json:"kind,omitempty"`
}

type AzureFile struct {
	SecretName      string `json:"secret_name"`
	ShareName       string `json:"share_name"`
	ReadOnly        bool   `json:"readonly"`
	SecretNamespace string `json:"secret_namespace,omitempty"`
}
