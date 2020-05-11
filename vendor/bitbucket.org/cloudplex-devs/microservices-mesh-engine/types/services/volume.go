package services

import "k8s.io/apimachinery/pkg/api/resource"

// Volume represents a named volume in a pod that may be accessed by any container in the pod.
type Volume struct {
	Name         string       `json:"name"  bson:"name" protobuf:"bytes,1,opt,name=name"`
	VolumeSource VolumeSource `json:"volumeSource" bson:"volumeSource" protobuf:"bytes,2,opt,name=volumeSource"`
}

// Represents the source of a volume to mount.
// Only one of its members may be specified.
type VolumeSource struct {
	HostPath              *HostPathVolumeSource              `json:"host_path,omitempty" bson:"host_path,omitempty"  protobuf:"bytes,1,opt,name=hostPath"`
	EmptyDir              *EmptyDirVolumeSource              `json:"empty_dir,omitempty" bson:"empty_dir,omitempty" protobuf:"bytes,2,opt,name=emptyDir"`
	GCEPersistentDisk     *GCEPersistentDiskVolumeSource     `json:"gce_persistent_disk,omitempty" bson:"gce_persistent_disk,omitempty" protobuf:"bytes,3,opt,name=gcePersistentDisk"`
	AWSElasticBlockStore  *AWSElasticBlockStoreVolumeSource  `json:"aws_elastic_block_store,omitempty" bson:"aws_elastic_block_store,omitempty" protobuf:"bytes,4,opt,name=awsElasticBlockStore"`
	Secret                *SecretVolumeSource                `json:"secret,omitempty" bson:"secret,omitempty" protobuf:"bytes,6,opt,name=secret"`
	PersistentVolumeClaim *PersistentVolumeClaimVolumeSource `json:"persistent_volume_claim,omitempty" bson:"persistent_volume_claim,omitempty" protobuf:"bytes,10,opt,name=persistentVolumeClaim"`
	AzureFile             *AzureFileVolumeSource             `json:"azure_file,omitempty" bson:"azure_file,omitempty" protobuf:"bytes,18,opt,name=azureFile"`
	ConfigMap             *ConfigMapVolumeSource             `json:"config_map,omitempty" bson:"config_map,omitempty" protobuf:"bytes,19,opt,name=configMap"`
	AzureDisk             *AzureDiskVolumeSource             `json:"azure_disk,omitempty" bson:"azure_disk,omitempty" protobuf:"bytes,22,opt,name=azureDisk"`
}

// Represents a host path mapped into a pod.
// Host path volumes do not support ownership management or SELinux relabeling.
type HostPathVolumeSource struct {
	Path string        `json:"path" bson:"path" protobuf:"bytes,1,opt,name=path"`
	Type *HostPathType `json:"type,omitempty" bson:"type,omitempty" protobuf:"bytes,2,opt,name=type"`
}

// Represents an empty directory for a pod.
// Empty directory volumes support ownership management and SELinux relabeling.
type EmptyDirVolumeSource struct {
	Medium    StorageMedium      `json:"medium,omitempty" bson:"medium,omitempty" protobuf:"bytes,1,opt,name=medium,casttype=StorageMedium"`
	SizeLimit *resource.Quantity `json:"size_limit,omitempty" bson:"size_limit,omitempty" protobuf:"bytes,2,opt,name=sizeLimit"`
}

// Represents a Persistent Disk resource in Google Compute Engine.
//
// A GCE PD must exist before mounting to a container. The disk must
// also be in the same GCE project and zone as the kubelet. A GCE PD
// can only be mounted as read/write once or read-only many times. GCE
// PDs support ownership management and SELinux relabeling.
type GCEPersistentDiskVolumeSource struct {
	PDName    string `json:"pd_name"  bson:"pd_name" protobuf:"bytes,1,opt,name=pdName"`
	FSType    string `json:"fs_type,omitempty" bson:"fs_type,omitempty" protobuf:"bytes,2,opt,name=fsType"`
	Partition int32  `json:"partition,omitempty" bson:"partition,omitempty" protobuf:"varint,3,opt,name=partition"`
	ReadOnly  bool   `json:"readonly,omitempty" bson:"readonly,omitempty" protobuf:"varint,4,opt,name=readOnly"`
}

// Represents a Persistent Disk resource in AWS.
//
// An AWS EBS disk must exist before mounting to a container. The disk
// must also be in the same AWS zone as the kubelet. An AWS EBS disk
// can only be mounted as read/write once. AWS EBS volumes support
// ownership management and SELinux relabeling.
type AWSElasticBlockStoreVolumeSource struct {
	VolumeID  string `json:"volume_id" bson:"volume_id" protobuf:"bytes,1,opt,name=volumeID"`
	FSType    string `json:"fs_type,omitempty" bson:"fs_type,omitempty" protobuf:"bytes,2,opt,name=fsType"`
	Partition int32  `json:"partition,omitempty" bson:"partition,omitempty" protobuf:"varint,3,opt,name=partition"`
	ReadOnly  bool   `json:"readonly,omitempty" bson:"readonly,omitempty" protobuf:"varint,4,opt,name=readOnly"`
}

// Adapts a Secret into a volume.
//
// The contents of the target Secret's Data field will be presented in a volume
// as files using the keys in the Data field as the file names.
// Secret volumes support ownership management and SELinux relabeling.
type SecretVolumeSource struct {
	SecretName  string      `json:"secret_name,omitempty" bson:"secret_name,omitempty" protobuf:"bytes,1,opt,name=secretName"`
	Items       []KeyToPath `json:"items,omitempty" bson:"items,omitempty" protobuf:"bytes,2,rep,name=items"`
	DefaultMode *int32      `json:"default_mode,omitempty" bson:"default_mode,omitempty" protobuf:"bytes,3,opt,name=defaultMode"`
	Optional    *bool       `json:"optional,omitempty" bson:"optional,omitempty" protobuf:"varint,4,opt,name=optional"`
}

const (
	SecretVolumeSourceDefaultMode int32 = 0644
)

// AzureFile represents an Azure File Service mount on the host and bind mount to the pod.
type AzureFileVolumeSource struct {
	SecretName string `json:"secret_name" bson:"secret_name" protobuf:"bytes,1,opt,name=secretName"`
	ShareName  string `json:"share_name" bson:"share_name" protobuf:"bytes,2,opt,name=shareName"`
	ReadOnly   bool   `json:"readonly,omitempty" bson:"readonly,omitempty" protobuf:"varint,3,opt,name=readOnly"`
}

// Adapts a ConfigMap into a volume.
//
// The contents of the target ConfigMap's Data field will be presented in a
// volume as files using the keys in the Data field as the file names, unless
// the items element is populated with specific mappings of keys to paths.
// ConfigMap volumes support ownership management and SELinux relabeling.
type ConfigMapVolumeSource struct {
	LocalObjectReference `json:",inline" bson:",inline" protobuf:"bytes,1,opt,name=localObjectReference"`
	Items                []KeyToPath `json:"items,omitempty" bson:"items,omitempty"  protobuf:"bytes,2,rep,name=items"`
	DefaultMode          *int32      `json:"default_mode,omitempty" bson:"default_mode,omitempty" protobuf:"varint,3,opt,name=defaultMode"`
	Optional             *bool       `json:"optional,omitempty" bson:"optional,omitempty" protobuf:"varint,4,opt,name=optional"`
}

const (
	ConfigMapVolumeSourceDefaultMode int32 = 0644
)

// Maps a string key to a path within a volume.
type KeyToPath struct {
	Key  string `json:"key" bson:"key" protobuf:"bytes,1,opt,name=key"`
	Path string `json:"path" bson:"path" protobuf:"bytes,2,opt,name=path"`
	Mode *int32 `json:"mode,omitempty" bson:"mode,omitempty" protobuf:"varint,3,opt,name=mode"`
}

// PersistentVolumeClaimVolumeSource references the user's PVC in the same namespace.
// This volume finds the bound PV and mounts that volume for the pod. A
// PersistentVolumeClaimVolumeSource is, essentially, a wrapper around another
// type of volume that is owned by someone else (the system).
type PersistentVolumeClaimVolumeSource struct {
	ClaimName string `json:"claim_name" bson:"claim_name" protobuf:"bytes,1,opt,name=claimName"`
	ReadOnly  bool   `json:"readonly,omitempty" bson:"readonly,omitempty" protobuf:"varint,2,opt,name=readOnly"`
}

// LocalObjectReference contains enough information to let you locate the
// referenced object inside the same namespace.
type LocalObjectReference struct {
	Name string `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`
}

// AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.
type AzureDiskVolumeSource struct {
	DiskName    string                    `json:"disk_name" bson:"disk_name" protobuf:"bytes,1,opt,name=diskName"`
	DataDiskURI string                    `json:"disk_uri" bson:"disk_uri" protobuf:"bytes,2,opt,name=diskURI"`
	CachingMode *AzureDataDiskCachingMode `json:"caching_mode,omitempty" bson:"caching_mode,omitempty" protobuf:"bytes,3,opt,name=cachingMode,casttype=AzureDataDiskCachingMode"`
	FSType      *string                   `json:"fs_type,omitempty" bson:"fs_type,omitempty" protobuf:"bytes,4,opt,name=fsType"`
	ReadOnly    *bool                     `json:"readonly,omitempty" bson:"readonly,omitempty" protobuf:"varint,5,opt,name=readOnly"`
	Kind        *AzureDataDiskKind        `json:"kind,omitempty" bson:"kind,omitempty" protobuf:"bytes,6,opt,name=kind,casttype=AzureDataDiskKind"`
}

type AzureDataDiskCachingMode string
type AzureDataDiskKind string

const (
	AzureDataDiskCachingNone      AzureDataDiskCachingMode = "ModeNone"
	AzureDataDiskCachingReadOnly  AzureDataDiskCachingMode = "ReadOnly"
	AzureDataDiskCachingReadWrite AzureDataDiskCachingMode = "ReadWrite"

	AzureSharedBlobDisk    AzureDataDiskKind = "Shared"
	AzureDedicatedBlobDisk AzureDataDiskKind = "Dedicated"
	AzureManagedDisk       AzureDataDiskKind = "Managed"
)

type HostPathType string

const (
	// For backwards compatible, leave it empty if unset
	HostPathUnset HostPathType = ""
	// If nothing exists at the given path, an empty directory will be created there
	// as needed with file mode 0755, having the same group and ownership with Kubelet.
	HostPathDirectoryOrCreate HostPathType = "DirectoryOrCreate"
	// A directory must exist at the given path
	HostPathDirectory HostPathType = "Directory"
	// If nothing exists at the given path, an empty file will be created there
	// as needed with file mode 0644, having the same group and ownership with Kubelet.
	HostPathFileOrCreate HostPathType = "FileOrCreate"
	// A file must exist at the given path
	HostPathFile HostPathType = "File"
	// A UNIX socket must exist at the given path
	HostPathSocket HostPathType = "Socket"
	// A character device must exist at the given path
	HostPathCharDev HostPathType = "CharDevice"
	// A block device must exist at the given path
	HostPathBlockDev HostPathType = "BlockDevice"
)

// StorageMedium defines ways that storage can be allocated to a volume.
type StorageMedium string

const (
	StorageMediumDefault   StorageMedium = ""          // use whatever the default is for the node, assume anything we don't explicitly handle is this
	StorageMediumMemory    StorageMedium = "Memory"    // use memory (e.g. tmpfs on linux)
	StorageMediumHugePages StorageMedium = "HugePages" // use hugepages
)
