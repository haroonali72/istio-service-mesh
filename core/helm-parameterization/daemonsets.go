package helm_parameterization

import (
	"istio-service-mesh/core/helm-parameterization/types"
	v1 "k8s.io/api/apps/v1"
	"sigs.k8s.io/yaml"
	"strings"
)

func DaemonSetsParameters(daemonset *v1.DaemonSet) (daemonsetYaml []byte, daemonsetParams []byte, functionsData []byte, err error) {
	result, err := yaml.Marshal(daemonset)
	if err != nil {
		return nil, nil, nil, err
	}
	daemonsetRaw := new(types.DaemonSetTemplate)
	err = yaml.Unmarshal(result, daemonsetRaw)
	if err != nil {
		return nil, nil, nil, err
	}
	tplFile := new([]byte)
	_ = tplFile
	chartFile := new(types.CoreComponentsChartValues)

	daemonsetRaw.Labels, _ = appendLabels(daemonset.Labels, daemonset.Name, tplFile)
	daemonsetRaw.Spec.Selector.MatchLabels, _ = appendMatchLabels(daemonset.Spec.Selector.MatchLabels, daemonset.Name, tplFile)
	daemonsetRaw.Spec.Template.Labels, _ = appendPodLabels(daemonset.Spec.Template.Labels, daemonset.Name)
	daemonsetRaw.Spec.Template.Spec.Containers[0].Resources, _ = appendResourceQuota(daemonset.Spec.Template.Spec.Containers[0].Resources, chartFile)
	daemonsetRaw.Spec.Template.Spec.Containers[0].LivenessProbe, daemonsetRaw.Spec.Template.Spec.Containers[0].ReadinessProbe, _ = appendProbing(daemonset.Spec.Template.Spec.Containers[0].LivenessProbe, daemonset.Spec.Template.Spec.Containers[0].ReadinessProbe, chartFile)
	daemonsetRaw.Spec.Template.Spec.Containers[0].Image, _ = appendImage(daemonset.Spec.Template.Spec.Containers[0].Image, chartFile)
	daemonsetRaw.Spec.Template.Spec.Containers[0].ImagePullPolicy, _ = appendImagePullPolicy(string(daemonset.Spec.Template.Spec.Containers[0].ImagePullPolicy), chartFile)
	daemonsetRaw.Spec.Template.Spec.ImagePullSecrets, _ = appendImagePullSecret(daemonset.Spec.Template.Spec.ImagePullSecrets, chartFile)
	daemonsetRaw.Spec.Template.Spec.Containers[0].Ports, _ = appendPorts(daemonset.Spec.Template.Spec.Containers[0].Ports, chartFile)

	//add this at the end. This function will replace name with helm parameter
	daemonsetRaw.Name, _ = appendName(daemonset.Name, tplFile)

	daemonsetYaml, err = yaml.Marshal(daemonsetRaw)
	if err != nil {
		return nil, nil, nil, err
	}

	chartRaw, err := yaml.Marshal(chartFile)
	if err != nil {
		return nil, nil, nil, err
	}

	daemonsetYamlStr := strings.ReplaceAll(string(daemonsetYaml), "'{{", "{{")
	daemonsetYamlStr = strings.ReplaceAll(daemonsetYamlStr, "}}'", "}}")

	daemonsetYamlStr = appendExtraStatements(daemonsetYamlStr, "readinessProbe:", ReadinessProbIfCondition)
	daemonsetYamlStr = appendExtraStatements(daemonsetYamlStr, "resources:", ResourcesIfCondition)
	daemonsetYamlStr = appendExtraStatements(daemonsetYamlStr, "livenessProbe:", LivelinessProbIfCondition)
	daemonsetYamlStr = appendExtraStatements(daemonsetYamlStr, "imagePullSecrets:", ImagePullSecretIfCondition)
	return []byte(daemonsetYamlStr), chartRaw, *tplFile, nil
}
