package helm_parameterization

import (
	"bitbucket.org/cloudplex-devs/istio-service-mesh/core/helm-parameterization/types"
	"k8s.io/api/core/v1"
	"sigs.k8s.io/yaml"
	"strings"
)

func ConfigMapParameters(ConfigMap *v1.ConfigMap) (configMapYaml []byte, configMapParams []byte, functionsData []byte, err error) {
	result, err := yaml.Marshal(ConfigMap)
	if err != nil {
		return nil, nil, nil, err
	}
	ConfigMapRaw := new(types.ConfigMapTemplate)
	err = yaml.Unmarshal(result, ConfigMapRaw)
	if err != nil {
		return nil, nil, nil, err
	}

	configMapParams = []byte("\n" + ConfigMap.Name + "CM:\n  data:")
	for key, value := range ConfigMap.Data {
		configMapParams = append(configMapParams, []byte("\n    "+key+": "+value)...)
	}

	ConfigMapRaw.Data = "{{ index .Values \"" + ConfigMap.Name + "CM\" \"data\" | toYaml | trim | nindent 2 }}"

	configMapYaml, err = yaml.Marshal(ConfigMapRaw)
	temp := strings.ReplaceAll(string(configMapYaml), "'{{", "{{")
	temp = strings.ReplaceAll(temp, "}}'", "}}")
	configMapYaml = []byte(temp)
	return
}
