package helm_parameterization

import (
	"istio-service-mesh/core/helm-parameterization/types"
	v1 "k8s.io/api/rbac/v1"
	"sigs.k8s.io/yaml"
	"strings"
)

func ClusterRoleBindingParameters(roleBinding *v1.ClusterRoleBinding) (roleBindingYaml []byte, values []byte, functionsData []byte, err error) {
	result, err := yaml.Marshal(roleBinding)
	if err != nil {
		return nil, nil, nil, err
	}

	roleBindingRaw := new(types.RoleBindingTemplate)
	err = yaml.Unmarshal(result, roleBindingRaw)
	if err != nil {
		return nil, nil, nil, err
	}
	tplFile := new([]byte)
	_ = tplFile

	chartFile := new(types.CoreComponentsChartValues)
	roleBindingRaw.Labels, _ = appendLabels(roleBinding.Labels, roleBinding.Name, tplFile)
	roleBindingRaw.Name, _ = appendName(roleBinding.Name, tplFile)
	roleBindingRaw.RoleRef.Name, _ = appendRoleRefName(roleBinding.RoleRef.Name)

	dep, err := yaml.Marshal(roleBindingRaw)
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
