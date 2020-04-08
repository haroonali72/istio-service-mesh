package helm_parameterization

import (
	"istio-service-mesh/core/helm-parameterization/types"
	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/yaml"
	"strings"
)

func ServiceAccountParameters(svcAccount v1.ServiceAccount) (svcYaml []byte, values []byte, functionsData []byte, err error) {
	result, err := yaml.Marshal(svcAccount)
	if err != nil {
		return nil, nil, nil, err
	}

	svcRaw := new(types.ServiceAccountTemplate)
	err = yaml.Unmarshal(result, svcRaw)
	if err != nil {
		return nil, nil, nil, err
	}
	tplFile := new([]byte)
	_ = tplFile

	chartFile := new([]byte)

	svcRaw.Labels, _ = appendLabels(svcAccount.Labels, svcAccount.Name, tplFile)
	svcRaw.Name, _ = appendServiceAccountName(svcAccount.Name, tplFile)

	dep, err := yaml.Marshal(svcRaw)
	if err != nil {
		return nil, nil, nil, err
	}

	valuesYaml, err := yaml.Marshal(chartFile)
	if err != nil {
		return nil, nil, nil, err
	}

	depString := strings.ReplaceAll(string(dep), "'{{", "{{")
	depString = strings.ReplaceAll(depString, "}}'", "}}")
	depString = appendIfStatements(depString, "apiVersion", KubernetesRBACIfCondition)

	return []byte(depString), valuesYaml, *tplFile, nil

}
