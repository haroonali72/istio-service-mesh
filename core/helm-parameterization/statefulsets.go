package helm_parameterization

import (
	"istio-service-mesh/core/helm-parameterization/types"
	v1 "k8s.io/api/apps/v1"
	"sigs.k8s.io/yaml"
)

func StatefulSetParameters(statefulset *v1.StatefulSet) (statefulSetYaml []byte, statefulSetParams []byte, functionsData []byte, err error) {
	result, err := yaml.Marshal(statefulset)
	if err != nil {
		return nil, nil, nil, err
	}
	statefulsetRaw := new(types.StatefulSetTemplate)
	err = yaml.Unmarshal(result, statefulsetRaw)
	if err != nil {
		return nil, nil, nil, err
	}
	tplFile := new([]byte)
	_ = tplFile
	chartFile := new(types.CoreComponentsChartValues)

	statefulsetRaw.Labels, _ = appendLabels(statefulset.Labels, statefulset.Name, tplFile)
	statefulsetRaw.Spec.Selector.MatchLabels, _ = appendMatchLabels(statefulset.Spec.Selector.MatchLabels, statefulset.Name, tplFile)
	statefulsetRaw.Spec.Template.Labels, _ = appendPodLabels(statefulset.Spec.Template.Labels, statefulset.Name)
	statefulsetRaw.Spec.Replicas, _ = appendReplicasTemplate(*statefulset.Spec.Replicas, chartFile)
	statefulsetRaw.Spec.Template.Spec.Containers[0].Resources, _ = appendResourceQuota(statefulset.Spec.Template.Spec.Containers[0].Resources, chartFile)
	statefulsetRaw.Spec.Template.Spec.Containers[0].LivenessProbe, statefulsetRaw.Spec.Template.Spec.Containers[0].ReadinessProbe, _ = appendProbing(statefulset.Spec.Template.Spec.Containers[0].LivenessProbe, statefulset.Spec.Template.Spec.Containers[0].ReadinessProbe, chartFile)
	statefulsetRaw.Spec.Template.Spec.Containers[0].Image, _ = appendImage(statefulset.Spec.Template.Spec.Containers[0].Image, chartFile)
	statefulsetRaw.Spec.Template.Spec.Containers[0].ImagePullPolicy, _ = appendImagePullPolicy(string(statefulset.Spec.Template.Spec.Containers[0].ImagePullPolicy), chartFile)
	statefulsetRaw.Spec.Template.Spec.ImagePullSecrets, _ = appendImagePullSecret(statefulset.Spec.Template.Spec.ImagePullSecrets, chartFile)
	statefulsetRaw.Spec.Template.Spec.Containers[0].Ports, _ = appendPorts(statefulset.Spec.Template.Spec.Containers[0].Ports, chartFile)
	statefulsetRaw.Spec.Template.Spec.Containers[0].Env, _ = appendEnvs(statefulset.Spec.Template.Spec.Containers[0].Env, chartFile)

	//add this at the end. This function will replace name with helm parameter
	statefulsetRaw.Name, _ = appendName(statefulset.Name, tplFile)

	statefulsetYaml, err := yaml.Marshal(statefulsetRaw)
	if err != nil {
		return nil, nil, nil, err
	}

	chartRaw, err := yaml.Marshal(chartFile)
	if err != nil {
		return nil, nil, nil, err
	}

	/*statefulsetYamlStr := strings.ReplaceAll(string(statefulsetYaml), "'{{", "{{")
	statefulsetYamlStr = strings.ReplaceAll(statefulsetYamlStr, "}}'", "}}")

	statefulsetYamlStr = appendExtraStatements(statefulsetYamlStr, "readinessProbe:", ReadinessProbIfCondition)
	statefulsetYamlStr = appendExtraStatements(statefulsetYamlStr, "resources:", ResourcesIfCondition)
	statefulsetYamlStr = appendExtraStatements(statefulsetYamlStr, "livenessProbe:", LivelinessProbIfCondition)
	statefulsetYamlStr = appendExtraStatements(statefulsetYamlStr, "imagePullSecrets:", ImagePullSecretIfCondition)*/

	statefulsetYamlStr := extraParametersReplacement(statefulsetYaml, statefulset.Name)
	return []byte(statefulsetYamlStr), chartRaw, *tplFile, nil
}
