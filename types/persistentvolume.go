package types

import "time"

type PersistentVolumeService struct {
	Id                interface{}                       `json:"_id,omitempty" bson:"_id" valid:"-"`
	ServiceId         string                            `json:"service_id" bson:"service_id" binding:"required" valid:"alphanumspecial,length(4|30)~service_id must contain between 6 and 30 characters,lowercase~lowercase alphanumeric characters are allowed,required~service_id is missing in request"`
	Name              string                            `json:"name"  bson:"name" binding:"required" valid:"alphanumspecial,length(3|30),lowercase~lowercase alphanumeric characters are allowed,required"`
	Version           string                            `json:"version"  bson:"version"  binding:"required" valid:"alphanumspecial,length(1|10),lowercase~lowercase alphanumeric characters are allowed,required"`
	ServiceType       string                            `json:"service_type"  bson:"service_type" valid:"-"`
	ServiceSubType    string                            `json:"service_sub_type" bson:"service_type" valid:"-"`
	Namespace         string                            `json:"namespace" bson:"namespace" binding:"required" valid:"alphanumspecial,required"`
	CompanyId         string                            `json:"company_id,omitempty" bson:"company_id"`
	CreationDate      time.Time                         `json:"creation_date,omitempty" bson:"creation_date" valid:"-"`
	ServiceAttributes *PersistentVolumeServiceAttribute `json:"service_attributes"  bson:"company_id" binding:"required"`
}
type PersistentVolumeServiceAttribute struct {
	Labels                 map[string]string      `json:"labels"`
	ReclaimPolicy          ReclaimPolicy          `json:"reclaimPolicy"`
	PersistentVolumeSource PersistentVolumeSource `json:"persistentVolumeSource"`
	AccessMode             []AccessMode           `json:"accessMode"`
	Capcity                string                 `json:"capcity,omitempty"`
	StorageClassName       string                 `json:"storageClassName,omitempty"`
	MountOptions           []string               `json:"mountOptions,omitempty"`
	VolumeMode             *PersistentVolumeMode  `json:"volumeMode,omitempty" protobuf:"bytes,8,opt,name=volumeMode,casttype=PersistentVolumeMode"`
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
	GCPPD     GCPPD     `json:"gcpPd"`
	AWSEBS    AWSEBS    `json:"awsEbs"`
	AzureDisk AzureDisk `json:"azureDisk"`
	AzureFile AzureFile `json:"azureFile"`
}

type GCPPD struct {
	PdName     string `json:"pdName"`
	Filesystem string `json:"fileSystem"`
	Partition  int    `json:"partation"`
	ReadOnly   bool   `json:"readOnly"`
}

type AWSEBS struct {
	VolumeId   string `json:"volumeId"`
	Filesystem string `json:"fileSystem"`
	Partition  int    `json:"partation"`
	ReadOnly   bool   `json:"readOnly"`
}

type AzureDisk struct {
	CachingMode string `json:"cachingMode"`
	Filesystem  string `json:"fileSystem"`
	ReadOnly    bool   `json:"readOnly"`
	DiskName    string `json:"diskName"`
	DiskURI     string `json:"diskURI"`
}

type AzureFile struct {
	SecretName      string `json:"secretName"`
	ShareName       string `json:"shareName"`
	ReadOnly        bool   `json:"readOnly"`
	SecretNamespace string `json:"secretNamespace,omitempty"`
}
