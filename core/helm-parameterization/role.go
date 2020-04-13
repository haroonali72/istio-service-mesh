package helm_parameterization

import (
	"istio-service-mesh/core/helm-parameterization/types"
	v1 "k8s.io/api/rbac/v1"
	"sigs.k8s.io/yaml"
	"strings"
)

func RoleParameters(role *v1.Role) ([]byte, []byte, []byte, error) {
	result, err := yaml.Marshal(role)
	if err != nil {
		return nil, nil, nil, err
	}
	roleRaw := new(types.RoleTemplate)
	err = yaml.Unmarshal(result, &roleRaw)
	if err != nil {
		return nil, nil, nil, err
	}

	tplFile := new([]byte)
	_ = tplFile
	chartFile := new(types.RBACChartValues)

	roleRaw.Labels, _ = appendLabels(role.Labels, role.Name, tplFile)
	roleRaw.Name, _ = appendName(role.Name, tplFile)
	chartFile.Create = true
	roleYaml, err := yaml.Marshal(roleRaw)
	if err != nil {
		return nil, nil, nil, err
	}

	roleString := strings.ReplaceAll(string(roleYaml), "'{{", "{{")
	roleString = strings.ReplaceAll(roleString, "}}'", "}}")

	roleString = appendIfStatements(roleString, "apiVersion", KubernetesRBACIfCondition)
	valuesYaml, err := yaml.Marshal(chartFile)
	if err != nil {
		return nil, nil, nil, err
	}
	return []byte(roleString), valuesYaml, *tplFile, nil

}
