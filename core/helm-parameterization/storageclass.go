package helm_parameterization

import (
	"istio-service-mesh/core/helm-parameterization/types"
	"k8s.io/api/storage/v1"
	"sigs.k8s.io/yaml"
	"strconv"
	"strings"
)

func StorageClassParameters(storageClass *v1.StorageClass) (storageClassYaml []byte, storageClassParams []byte, functionsData []byte, err error) {
	result, err := yaml.Marshal(storageClass)
	if err != nil {
		return nil, nil, nil, err
	}
	storageClassRaw := new(types.StorageClassTemplate)
	err = yaml.Unmarshal(result, storageClassRaw)
	if err != nil {
		return nil, nil, nil, err
	}
	storageClassParams = []byte("\n" + storageClass.Name + "SC:")
	storageClassParams = append(storageClassParams, []byte("\n  mountOptions:")...)
	storageClassRaw.MountOptions = "{{ index .Values \"" + storageClass.Name + "SC\" \"mountOptions\" | toYaml | trim | nindent 2 }}"
	for _, each := range storageClass.MountOptions {
		storageClassParams = append(storageClassParams, []byte("\n    - "+each)...)
	}

	if storageClass.ReclaimPolicy != nil {
		storageClassParams = append(storageClassParams, []byte("\n  reclaimPolicy: "+*storageClass.ReclaimPolicy)...)
		//		storageClassRaw.ReclaimPolicy="{{ .Values."+storageClass.Name+"SC.reclaimPolicy }}"

		storageClassRaw.ReclaimPolicy = "{{ index .Values \"" + storageClass.Name + "SC\" \"reclaimPolicy\" }}"
	}

	storageClassParams = append(storageClassParams, []byte("\n  provisioner: "+storageClass.Provisioner)...)
	//	storageClassRaw.Provisioner="{{ .Values."+storageClass.Name+"SC.provisioner }}"
	storageClassRaw.Provisioner = "{{ index .Values \"" + storageClass.Name + "SC\" \"provisioner\" }}"
	storageClassParams = append(storageClassParams, []byte("\n  parameters:")...)

	for key, value := range storageClass.Parameters {
		storageClassParams = append(storageClassParams, []byte("\n    "+key+": "+value)...)
	}
	//	storageClassRaw.Parameters= "{{ .Values." + storageClass.Name + "SC.parameters }}"
	storageClassRaw.Parameters = "{{ index .Values \"" + storageClass.Name + "SC\" \"parameters\" | toYaml | trim | nindent 2 }}"

	if storageClass.VolumeBindingMode != nil {
		storageClassParams = append(storageClassParams, []byte("\n  volumeBindingMode: "+*storageClass.VolumeBindingMode)...)
		//storageClassRaw.VolumeBindingMode="{{ .Values."+storageClass.Name+"SC.volumeBindingMode }}"
		storageClassRaw.VolumeBindingMode = "{{ index .Values \"" + storageClass.Name + "SC\" \"volumeBindingMode\" }}"

	}
	if storageClass.AllowVolumeExpansion != nil {
		storageClassParams = append(storageClassParams, []byte("\n  allowVolumeExpansion: "+strconv.FormatBool(*storageClass.AllowVolumeExpansion))...)
		//storageClassRaw.AllowVolumeExpansion="{{ .Values."+storageClass.Name+"SC.allowVolumeExpansion }}"
		storageClassRaw.AllowVolumeExpansion = "{{ index .Values \"" + storageClass.Name + "SC\" \"allowVolumeExpansion\" }}"

	}
	storageClassYaml, err = yaml.Marshal(storageClassRaw)
	temp := strings.ReplaceAll(string(storageClassYaml), "'{{", "{{")
	temp = strings.ReplaceAll(temp, "}}'", "}}")
	storageClassYaml = []byte(temp)
	return
}
