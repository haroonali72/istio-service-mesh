package volumes

import (
	"istio-service-mesh/types"
	core "k8s.io/api/core/v1"
	"k8s.io/api/storage/v1"
	"strconv"
	"strings"
)

func ProvisionStorageClass(volume types.Volume) *v1.StorageClass {
	storageClass := v1.StorageClass{}

	storageClass.TypeMeta.Kind = "StorageClass"
	storageClass.TypeMeta.APIVersion = "storage.k8s.io/v1"
	storageClass.Name = GetStorageClassName(volume.Name)
	storageClass.ObjectMeta.Name = storageClass.Name
	//storageClass.Namespace = volume.Namespace

	reclaimPolicy := core.PersistentVolumeReclaimRetain
	storageClass.ReclaimPolicy = &reclaimPolicy //default policy is to retain the volume

	switch strings.ToLower(volume.Cloud) {
	case string(types.AWS):
		provisionAwsClass(&storageClass, volume)
	case string(types.Azure):
		provisionAzureClass(&storageClass, volume)
	case string(types.GCP):
		provisionGCPClass(&storageClass, volume)
	case string(types.DO):
		return nil
	}

	return &storageClass
}

func GetStorageClassName(volumeName string) string {
	return strings.Replace(strings.ToLower(volumeName), " ", "-", -1) + "-class"
}

func provisionAwsClass(storageClass *v1.StorageClass, volume types.Volume) {
	storageClass.Provisioner = "kubernetes.io/aws-ebs"

	if strings.ToLower(volume.Params.Type) == "io1" && volume.Params.Iops != 0 {
		storageClass.Parameters = map[string]string{
			"type":      volume.Params.Type,
			"iopsPerGB": strconv.Itoa(volume.Params.Iops),
		}
	} else if strings.ToLower(volume.Params.Type) != "" {
		storageClass.Parameters = map[string]string{
			"type": volume.Params.Type,
		}
	}
}

func provisionAzureClass(storageClass *v1.StorageClass, volume types.Volume) {
	switch strings.ToLower(volume.Params.Plugin) {
	case "file":
		storageClass.Provisioner = "kubernetes.io/azure-file"
	case "disk":
		storageClass.Provisioner = "kubernetes.io/azure-disk"
	default:
		storageClass.Provisioner = "kubernetes.io/azure-file"
	}

	storageClass.Parameters = map[string]string{}
	if volume.Params.SkuName != "" {
		storageClass.Parameters["skuName"] = volume.Params.SkuName
	}
	if volume.Params.Location != "" {
		storageClass.Parameters["location"] = volume.Params.Location
	}
	if volume.Params.StorageAccount != "" {
		storageClass.Parameters["storageAccount"] = volume.Params.StorageAccount
	}
}

func provisionGCPClass(storageClass *v1.StorageClass, volume types.Volume) {
	storageClass.Provisioner = "kubernetes.io/gce-pd"

	storageClass.Parameters = map[string]string{}
	if volume.Params.Type != "" {
		storageClass.Parameters["type"] = volume.Params.Type
	}
	if volume.Params.ReplicationType != "" {
		storageClass.Parameters["replication-type"] = volume.Params.ReplicationType
	}
}
