package volumes

import (
	"istio-service-mesh/types"
	"istio-service-mesh/utils"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"strings"
)

func ProvisionVolumeClaim(volume types.Volume) v1.PersistentVolumeClaim {
	volumeClaim := v1.PersistentVolumeClaim{}

	volumeClaim.TypeMeta.Kind = "PersistentVolumeClaim"
	volumeClaim.TypeMeta.APIVersion = "v1"
	volumeClaim.Name = GetVolumeClaimName(volume.Name)
	volumeClaim.ObjectMeta.Name = volumeClaim.Name
	volumeClaim.Namespace = volume.Namespace

	volumeClaim.Spec.StorageClassName = utils.StringPtr(GetStorageClassName(volume.Name))
	volumeClaim.Spec.AccessModes = []v1.PersistentVolumeAccessMode{v1.ReadWriteOnce}

	volumeClaim.Spec.Resources.Requests = map[v1.ResourceName]resource.Quantity{
		v1.ResourceStorage: *resource.NewScaledQuantity(volume.Size, resource.Giga),
	}

	return volumeClaim
}

func GetVolumeClaimName(volumeName string) string {
	return strings.Replace(strings.ToLower(volumeName), " ", "-", -1) + "-claim"
}
