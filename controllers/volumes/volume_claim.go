package volumes

/*
func ProvisionVolumeClaim(volume types.Volume) v1.PersistentVolumeClaim {
	volumeClaim := v1.PersistentVolumeClaim{}

	var storageClassName string
	if volume.Cloud == string(types.DO) {
		storageClassName = "do-block-storage"
	} else {
		storageClassName = GetStorageClassName(volume.Name)
	}

	volumeClaim.TypeMeta.Kind = "PersistentVolumeClaim"
	volumeClaim.TypeMeta.APIVersion = "v1"
	volumeClaim.ObjectMeta.Name = volumeClaim.Name
	volumeClaim.ObjectMeta.Annotations = map[string]string{
		"volume.beta.kubernetes.io/storage-class": storageClassName,
	}

	volumeClaim.Name = GetVolumeClaimName(volume.Name)
	volumeClaim.Namespace = volume.Namespace
	volumeClaim.Spec.StorageClassName = &storageClassName
	volumeClaim.Spec.AccessModes = []v1.PersistentVolumeAccessMode{v1.ReadWriteOnce}

	size, _ := resource.ParseQuantity(strconv.Itoa(int(volume.Size)) + "Gi")
	volumeClaim.Spec.Resources = v1.ResourceRequirements{Requests: v1.ResourceList{v1.ResourceStorage: size}}

	return volumeClaim
}

func GetVolumeClaimName(volumeName string) string {
	return strings.Replace(strings.ToLower(volumeName), " ", "-", -1) + "-claim"
}
*/
