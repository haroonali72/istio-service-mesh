package types

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type PersistentVolumeTemplate struct {
	metav1.TypeMeta    `json:",inline" `
	ObjectMetaTemplate `json:"metadata" yaml:"metadata"`
	Spec               PersistentVolumeTemplateSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}
type PersistentVolumeTemplateSpec struct {
	Capacity                      map[string]interface{} `json:"capacity,omitempty" protobuf:"bytes,1,rep,name=capacity,casttype=ResourceList,castkey=ResourceName"`
	AccessModes                   interface{}            `json:"accessModes,omitempty" protobuf:"bytes,3,rep,name=accessModes,casttype=PersistentVolumeAccessMode"`
	PersistentVolumeReclaimPolicy interface{}            `json:"persistentVolumeReclaimPolicy,omitempty" protobuf:"bytes,5,opt,name=persistentVolumeReclaimPolicy,casttype=PersistentVolumeReclaimPolicy"`
	// Name of StorageClass to which this persistent volume belongs. Empty value
	// means that this volume does not belong to any StorageClass.
	// +optional
	StorageClassName string `json:"storageClassName,omitempty" protobuf:"bytes,6,opt,name=storageClassName"`
	// A list of mount options, e.g. ["ro", "soft"]. Not validated - mount will
	// simply fail if one is invalid.
	// More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes/#mount-options
	// +optional
	MountOptions interface{} `json:"mountOptions,omitempty" protobuf:"bytes,7,opt,name=mountOptions"`
	// volumeMode defines if a volume is intended to be used with a formatted filesystem
	// or to remain in raw block state. Value of Filesystem is implied when not included in spec.
	// This is a beta feature.
	// +optional
	VolumeMode interface{} `json:"volumeMode,omitempty" protobuf:"bytes,8,opt,name=volumeMode,casttype=PersistentVolumeMode"`
	// The actual volume backing the persistent volume.
	PersistentVolumeSource `json:",inline" protobuf:"bytes,2,opt,name=persistentVolumeSource"`
}
type PersistentVolumeSource struct {
	// GCEPersistentDisk represents a GCE Disk resource that is attached to a
	// kubelet's host machine and then exposed to the pod. Provisioned by an admin.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
	// +optional
	GCEPersistentDisk *GCPPD `json:"gcePersistentDisk,omitempty" protobuf:"bytes,1,opt,name=gcePersistentDisk"`
	// AWSElasticBlockStore represents an AWS Disk resource that is attached to a
	// kubelet's host machine and then exposed to the pod.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
	// +optional
	AWSElasticBlockStore *AWSEBS `json:"awsElasticBlockStore,omitempty" protobuf:"bytes,2,opt,name=awsElasticBlockStore"`
	// HostPath represents a directory on the host.
	// Provisioned by a developer or tester.
	// This is useful for single-node development and testing only!
	// On-host storage is not supported in any way and WILL NOT WORK in a multi-node cluster.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath
	// +optional
	HostPath interface{} `json:"hostPath,omitempty" protobuf:"bytes,3,opt,name=hostPath"`
	// Glusterfs represents a Glusterfs volume that is attached to a host and
	// exposed to the pod. Provisioned by an admin.
	// More info: https://examples.k8s.io/volumes/glusterfs/README.md
	// +optional
	Glusterfs interface{} `json:"glusterfs,omitempty" protobuf:"bytes,4,opt,name=glusterfs"`
	// NFS represents an NFS mount on the host. Provisioned by an admin.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
	// +optional
	NFS interface{} `json:"nfs,omitempty" protobuf:"bytes,5,opt,name=nfs"`
	// RBD represents a Rados Block Device mount on the host that shares a pod's lifetime.
	// More info: https://examples.k8s.io/volumes/rbd/README.md
	// +optional
	RBD interface{} `json:"rbd,omitempty" protobuf:"bytes,6,opt,name=rbd"`
	// ISCSI represents an ISCSI Disk resource that is attached to a
	// kubelet's host machine and then exposed to the pod. Provisioned by an admin.
	// +optional
	ISCSI interface{} `json:"iscsi,omitempty" protobuf:"bytes,7,opt,name=iscsi"`
	// Cinder represents a cinder volume attached and mounted on kubelets host machine.
	// More info: https://examples.k8s.io/mysql-cinder-pd/README.md
	// +optional
	Cinder interface{} `json:"cinder,omitempty" protobuf:"bytes,8,opt,name=cinder"`
	// CephFS represents a Ceph FS mount on the host that shares a pod's lifetime
	// +optional
	CephFS interface{} `json:"cephfs,omitempty" protobuf:"bytes,9,opt,name=cephfs"`
	// FC represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.
	// +optional
	FC interface{} `json:"fc,omitempty" protobuf:"bytes,10,opt,name=fc"`
	// Flocker represents a Flocker volume attached to a kubelet's host machine and exposed to the pod for its usage. This depends on the Flocker control service being running
	// +optional
	Flocker interface{} `json:"flocker,omitempty" protobuf:"bytes,11,opt,name=flocker"`
	// FlexVolume represents a generic volume resource that is
	// provisioned/attached using an exec based plugin.
	// +optional
	FlexVolume interface{} `json:"flexVolume,omitempty" protobuf:"bytes,12,opt,name=flexVolume"`
	// AzureFile represents an Azure File Service mount on the host and bind mount to the pod.
	// +optional
	AzureFile *AzureFile `json:"azureFile,omitempty" protobuf:"bytes,13,opt,name=azureFile"`
	// VsphereVolume represents a vSphere volume attached and mounted on kubelets host machine
	// +optional
	VsphereVolume interface{} `json:"vsphereVolume,omitempty" protobuf:"bytes,14,opt,name=vsphereVolume"`
	// Quobyte represents a Quobyte mount on the host that shares a pod's lifetime
	// +optional
	Quobyte interface{} `json:"quobyte,omitempty" protobuf:"bytes,15,opt,name=quobyte"`
	// AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.
	// +optional
	AzureDisk *AzureDisk `json:"azureDisk,omitempty" protobuf:"bytes,16,opt,name=azureDisk"`
	// PhotonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine
	PhotonPersistentDisk interface{} `json:"photonPersistentDisk,omitempty" protobuf:"bytes,17,opt,name=photonPersistentDisk"`
	// PortworxVolume represents a portworx volume attached and mounted on kubelets host machine
	// +optional
	PortworxVolume interface{} `json:"portworxVolume,omitempty" protobuf:"bytes,18,opt,name=portworxVolume"`
	// ScaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.
	// +optional
	ScaleIO interface{} `json:"scaleIO,omitempty" protobuf:"bytes,19,opt,name=scaleIO"`
	// Local represents directly-attached storage with node affinity
	// +optional
	Local interface{} `json:"local,omitempty" protobuf:"bytes,20,opt,name=local"`
	// StorageOS represents a StorageOS volume that is attached to the kubelet's host machine and mounted into the pod
	// More info: https://examples.k8s.io/volumes/storageos/README.md
	// +optional
	StorageOS interface{} `json:"storageos,omitempty" protobuf:"bytes,21,opt,name=storageos"`
	// CSI represents storage that is handled by an external CSI driver (Beta feature).
	// +optional
	CSI interface{} `json:"csi,omitempty" protobuf:"bytes,22,opt,name=csi"`
}

type GCPPD struct {
	PdName    string      `json:"pdName"`
	FSType    string      `json:"fsType,omitempty" protobuf:"bytes,2,opt,name=fsType"`
	Partition interface{} `json:"partition,omitempty"`
	ReadOnly  interface{} `json:"readOnly,omitempty"`
}

type AWSEBS struct {
	VolumeId  string      `json:"volumeId"`
	FSType    string      `json:"fsType,omitempty" protobuf:"bytes,2,opt,name=fsType"`
	Partition interface{} `json:"partition,omitempty"`
	ReadOnly  interface{} `json:"readOnly,omitempty"`
}

type AzureDisk struct {
	CachingMode string      `json:"cachingMode,omitempty"`
	FSType      string      `json:"fsType,omitempty" protobuf:"bytes,4,opt,name=fsType"`
	ReadOnly    interface{} `json:"readOnly"`
	DiskName    string      `json:"diskName"`
	DiskURI     string      `json:"diskURI"`
	Kind        string      `json:"kind,omitempty"`
}

type AzureFile struct {
	SecretName      string      `json:"secretName"`
	ShareName       string      `json:"shareName"`
	ReadOnly        interface{} `json:"readOnly"`
	SecretNamespace string      `json:"secretNamespace,omitempty"`
}
