package helm_parameterization

import (
	"istio-service-mesh/core/helm-parameterization/types"
	v1 "k8s.io/api/rbac/v1"
	"sigs.k8s.io/yaml"
	"strings"
)

func ClusterRoleBindingParameters(roleBinding v1.ClusterRoleBinding) (roleBindingYaml []byte, functionsData []byte, err error) {
	result, err := yaml.Marshal(roleBinding)
	if err != nil {
		return nil, nil, err
	}

	roleBindingRaw := new(types.ClusterRoleBindingTemplate)
	err = yaml.Unmarshal(result, roleBindingRaw)
	if err != nil {
		return nil, nil, err
	}
	tplFile := new([]byte)
	_ = tplFile

	roleBindingRaw.Labels, _ = appendLabels(roleBinding.Labels, roleBinding.Name, tplFile)
	//roleBindingRaw.RoleRef.Name, _ = appendName()
	roleBindingRaw.Name, _ = appendName(roleBinding.Name, tplFile)

	dep, err := yaml.Marshal(roleBindingRaw)
	if err != nil {
		return nil, nil, err
	}

	depString := strings.ReplaceAll(string(dep), "'{{", "{{")
	depString = strings.ReplaceAll(depString, "}}'", "}}")

	return []byte(depString), *tplFile, nil

}
