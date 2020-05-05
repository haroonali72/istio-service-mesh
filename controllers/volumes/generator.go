package volumes

/*
func GenerateVolumeMounts(volumes []types.Volume) []v1.VolumeMount {
	volumeMounts := []v1.VolumeMount{}
	for _, volume := range volumes {
		volumeMounts = append(volumeMounts, v1.VolumeMount{
			Name:      volume.Name,
			MountPath: volume.MountPath,
		})

	}
	return volumeMounts
}

func GeneratePodVolumes(volumes []types.Volume) []v1.Volume {
	podVolumes := []v1.Volume{}
	for _, volume := range volumes {
		podVolumes = append(podVolumes, v1.Volume{
			Name: volume.Name,
			VolumeSource: v1.VolumeSource{
				PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{
					ClaimName: GetVolumeClaimName(volume.Name),
				},
			},
		})
	}
	return podVolumes
}
*/
