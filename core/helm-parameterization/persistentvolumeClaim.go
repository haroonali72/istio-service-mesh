package helm_parameterization

import (
	"bitbucket.org/cloudplex-devs/istio-service-mesh/core/helm-parameterization/types"
	core "k8s.io/api/core/v1"
	"sigs.k8s.io/yaml"
	"strings"
)

func PersistentVolumeClaimParameters(persistentVolumeClaim *core.PersistentVolumeClaim) (persistentVolumeClaimYaml []byte, persistentVolumeClaimParams []byte, functionsData []byte, err error) {
	result, err := yaml.Marshal(persistentVolumeClaim)
	if err != nil {
		return nil, nil, nil, err
	}
	persistentVolumeClaimRaw := new(types.PersistentVolumeClaimTemplate)
	err = yaml.Unmarshal(result, persistentVolumeClaimRaw)
	if err != nil {
		return nil, nil, nil, err
	}

	persistentVolumeClaimParams = []byte("\n" + persistentVolumeClaim.Name + "PVC:\n  accessModes:")
	for _, each := range persistentVolumeClaim.Spec.AccessModes {
		persistentVolumeClaimParams = append(persistentVolumeClaimParams, []byte("\n    - "+each)...)
	}
	//ConfigMapRaw.Data="{{ index .Values \""+ ConfigMap.Name+"CM\" \"data\" | toYaml | trim | nindent 2 }}"
	persistentVolumeClaimRaw.Spec.AccessModes = "{{ index .Values \"" + persistentVolumeClaim.Name + "PVC\" \"accessModes\" | toYaml | trim | nindent 2 }}"
	qunatity, ok := persistentVolumeClaim.Spec.Resources.Requests["storage"]
	if ok {
		persistentVolumeClaimParams = append(persistentVolumeClaimParams, []byte("\n  resources:")...)
		persistentVolumeClaimParams = append(persistentVolumeClaimParams, []byte("\n    requests: "+qunatity.String())...)
		//	persistentVolumeClaimRaw.Spec.Resources.Requests["storage"]="{{ .Values."+ persistentVolumeClaim.Name+"PVC.resources.limits }}"
		persistentVolumeClaimRaw.Spec.Resources.Requests["storage"] = "{{ index .Values \"" + persistentVolumeClaim.Name + "PVC\" \"resources\" \"requests\" }}"
	}
	qunatity, ok = persistentVolumeClaim.Spec.Resources.Limits["storage"]
	if ok {
		persistentVolumeClaimParams = append(persistentVolumeClaimParams, []byte("\n  resources:")...)
		persistentVolumeClaimParams = append(persistentVolumeClaimParams, []byte("\n    limits: "+qunatity.String())...)
		//persistentVolumeClaimRaw.Spec.Resources.Limits["storage"]="{{ .Values."+ persistentVolumeClaim.Name+"PVC.resources.limits }}"
		persistentVolumeClaimRaw.Spec.Resources.Limits["storage"] = "{{ index .Values \"" + persistentVolumeClaim.Name + "PVC\" \"resources\" \"limits\" }}"
	}
	if persistentVolumeClaim.Spec.VolumeMode != nil {
		persistentVolumeClaimParams = append(persistentVolumeClaimParams, []byte("\n  volumeMode: "+*persistentVolumeClaim.Spec.VolumeMode)...)
		persistentVolumeClaimRaw.Spec.VolumeMode = "{{ index .Values \"" + persistentVolumeClaim.Name + "PVC\" \"volumeMode\" }}"
		//		persistentVolumeClaimRaw.Spec.VolumeMode="{{ .Values."+ persistentVolumeClaim.Name+"PVC.volumeMode }}"
	}
	persistentVolumeClaimYaml, err = yaml.Marshal(persistentVolumeClaimRaw)
	temp := strings.ReplaceAll(string(persistentVolumeClaimYaml), "'{{", "{{")
	temp = strings.ReplaceAll(temp, "}}'", "}}")
	persistentVolumeClaimYaml = []byte(temp)
	return
}
