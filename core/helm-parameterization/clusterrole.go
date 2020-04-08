package helm_parameterization

import (
	"gopkg.in/yaml.v2"
	"istio-service-mesh/core/helm-parameterization/types"
	v1 "k8s.io/api/rbac/v1"
	"strings"
)

func ClusterRoleParameters(role v1.ClusterRole) (roleYaml []byte, functionsData []byte, err error) {
	result, err := yaml.Marshal(role)
	if err != nil {
		return nil, nil, err
	}

	roleRaw := new(types.ClusterRoleTemplate)
	err = yaml.Unmarshal(result, roleRaw)
	if err != nil {
		return nil, nil, err
	}
	tplFile := new([]byte)
	_ = tplFile

	roleRaw.Labels, _ = appendLabels(role.Labels, role.Name, tplFile)
	roleRaw.Name, _ = appendName(role.Name, tplFile)

	dep, err := yaml.Marshal(roleRaw)
	if err != nil {
		return nil, nil, err
	}

	depString := strings.ReplaceAll(string(dep), "'{{", "{{")
	depString = strings.ReplaceAll(depString, "}}'", "}}")

	return []byte(depString), *tplFile, nil

}
