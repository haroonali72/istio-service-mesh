package helm_parameterization

import (
	"istio-service-mesh/core/helm-parameterization/types"
	v1 "k8s.io/api/rbac/v1"
	"sigs.k8s.io/yaml"
	"strings"
)

func ClusterRoleParameters(role *v1.ClusterRole) (roleYaml []byte, values []byte, functionsData []byte, err error) {
	result, err := yaml.Marshal(role)
	if err != nil {
		return nil, nil, nil, err
	}

	roleRaw := new(types.RoleTemplate)
	err = yaml.Unmarshal(result, roleRaw)
	if err != nil {
		return nil, nil, nil, err
	}
	tplFile := new([]byte)
	_ = tplFile

	chartFile := new([]byte)
	roleRaw.Labels, _ = appendLabels(role.Labels, role.Name, tplFile)
	roleRaw.Name, _ = appendName(role.Name, tplFile)

	dep, err := yaml.Marshal(roleRaw)
	if err != nil {
		return nil, nil, nil, err
	}

	depString := strings.ReplaceAll(string(dep), "'{{", "{{")
	depString = strings.ReplaceAll(depString, "}}'", "}}")

	depString = appendIfStatements(depString, "apiVersion", KubernetesRBACIfCondition)

	chartRaw, err := yaml.Marshal(chartFile)
	if err != nil {
		return nil, nil, nil, err
	}

	return []byte(depString), chartRaw, *tplFile, nil

}
