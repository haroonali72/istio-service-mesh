package helm_parameterization

import (
	"bitbucket.org/cloudplex-devs/istio-service-mesh/core/helm-parameterization/types"
	"k8s.io/api/core/v1"
	"sigs.k8s.io/yaml"
	"strings"
)

func SecretParameters(Secret *v1.Secret) (SecretYaml []byte, SecretParams []byte, functionsData []byte, err error) {
	result, err := yaml.Marshal(Secret)
	if err != nil {
		return nil, nil, nil, err
	}
	SecretRaw := new(types.SecretTemplate)
	err = yaml.Unmarshal(result, SecretRaw)
	if err != nil {
		return nil, nil, nil, err
	}

	SecretParams = []byte("\n" + Secret.Name + "Secret:\n  stringData:")
	for key, value := range Secret.StringData {
		SecretParams = append(SecretParams, []byte("\n    "+key+": "+value)...)
	}
	if len(Secret.StringData) > 0 {
		SecretRaw.StringData = "{{ index .Values \"" + Secret.Name + "Secret\" \"stringData\" | toYaml | trim | nindent 2 }}"
	}

	SecretParams = append(SecretParams, []byte("\n  data:")...)
	for key, value := range Secret.Data {
		SecretParams = append(SecretParams, []byte("\n    "+key+": ")...)
		SecretParams = append(SecretParams, value...)
	}
	if len(Secret.Data) > 0 {
		SecretRaw.Data = "{{ index .Values \"" + Secret.Name + "Secret\" \"data\" | toYaml | trim | nindent 2 }}"
	}
	if Secret.Type != "" {
		SecretParams = append(SecretParams, []byte("\n  type: "+Secret.Type)...)
		SecretRaw.Type = "{{ index .Values \"" + Secret.Name + "Secret\" \"type\" | toYaml | trim | nindent 2 }}"
	}
	SecretYaml, err = yaml.Marshal(SecretRaw)
	temp := strings.ReplaceAll(string(SecretYaml), "'{{", "{{")
	temp = strings.ReplaceAll(temp, "}}'", "}}")
	SecretYaml = []byte(temp)
	return
}
