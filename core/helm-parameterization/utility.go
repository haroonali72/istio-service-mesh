package helm_parameterization

import (
	"encoding/json"
	"istio-service-mesh/core/helm-parameterization/types"
	v12 "k8s.io/api/core/v1"
	"sigs.k8s.io/yaml"
	"strings"
)

func appendName(actualName string, tplFile *[]byte) (string, error) {
	name := strings.Replace(NameHelmParameter, "{{ .Name }}", actualName, -1)
	var nameInterface interface{}
	_ = json.Unmarshal([]byte(name), &nameInterface)

	nameGenFunc := strings.ReplaceAll(NameFunction, "{{ .Name }}", actualName)
	//deploymentRaw.Name = name
	*tplFile = append(*tplFile, []byte(nameGenFunc)...)
	return name, nil
}

func appendServiceAccountName(actualName string, tplFile *[]byte) (string, error) {
	name := strings.Replace(NameHelmParameter, "{{ .Name }}", actualName, -1)
	var nameInterface interface{}
	_ = json.Unmarshal([]byte(name), &nameInterface)

	nameGenFunc := strings.ReplaceAll(ServiceAccountNameFunction, "{{ .Name }}", actualName)
	*tplFile = append(*tplFile, []byte(nameGenFunc)...)
	return name, nil
}

func appendRefName(actualName string) (string, error) {
	name := strings.Replace(NameHelmParameter, "{{ .Name }}", actualName, -1)
	var nameInterface interface{}
	_ = json.Unmarshal([]byte(name), &nameInterface)
	return name, nil

}

func appendLabels(labels map[string]string, name string, tplFile *[]byte) (string, error) {
	rawLabels, err := yaml.Marshal(labels)
	if err != nil {
		return "", err
	}
	labelFunc := strings.ReplaceAll(LabelFunction, "{{ .Name }}", name)
	labelFunc = strings.ReplaceAll(labelFunc, "{{ .Labels }}", string(rawLabels))

	labelTemplate := strings.Replace(LabelParameter, "{{ .Name }}", name, -1)
	labelTemplate = strings.Replace(labelTemplate, "{{ .Indent }}", "4", -1)
	//deploymentRaw.Labels = labelTemplate

	*tplFile = append(*tplFile, []byte(labelFunc)...)
	return labelTemplate, nil
}

func appendMatchLabels(matchLabel map[string]string, name string, tplFile *[]byte) (string, error) {
	rawLabels, err := yaml.Marshal(matchLabel)
	if err != nil {
		return "", err
	}
	labelFunc := strings.ReplaceAll(LabelFunction, "{{ .Name }}", name)
	labelFunc = strings.ReplaceAll(labelFunc, "{{ .Labels }}", string(rawLabels))
	labelFunc = strings.ReplaceAll(labelFunc, name+".labels", name+".matchLabels")

	labelTemplate := strings.Replace(MatchSelectorParameter, "{{ .Name }}", name, -1)
	labelTemplate = strings.Replace(labelTemplate, "{{ .Indent }}", "6", -1)
	//deploymentRaw.Spec.Selector.MatchLabels = labelTemplate

	*tplFile = append(*tplFile, []byte(labelFunc)...)
	return labelTemplate, nil
}

func appendPodLabels(labels map[string]string, name string) (string, error) {

	labelFunc := strings.ReplaceAll(LabelFunction, "{{ .Name }}", name)
	labelFunc = strings.ReplaceAll(labelFunc, "{{ .Labels }}", name)

	labelTemplate := strings.Replace(LabelParameter, "{{ .Name }}", name, -1)
	labelTemplate = strings.Replace(labelTemplate, "{{ .Indent }}", "8", -1)
	//deploymentRaw.Spec.Template.Labels = labelTemplate

	//*tplFile = append(*tplFile, []byte(labelFunc)...)
	return labelTemplate, nil
}

func appendReplicasTemplate(replicas int32, chartFile *types.CoreComponentsChartValues) (string, error) {
	chartFile.Replicas = replicas
	//deploymentRaw.Spec.Replicas = ReplicasHelmParameter
	return ReplicasHelmParameter, nil
}

func appendResourceQuota(rq v12.ResourceRequirements, chartFile *types.CoreComponentsChartValues) (string, error) {
	chartFile.ResourceQuota = rq
	//deploymentRaw.Spec.Template.Spec.Containers[0].Resources = ResourcesParameter
	return ResourcesParameter, nil
}

func appendProbing(liveness, readiness *v12.Probe, chartFile *types.CoreComponentsChartValues) (string, string, error) {
	chartFile.LivenessProb = liveness
	chartFile.ReadinessProbe = readiness
	//deploymentRaw.Spec.Template.Spec.Containers[0].LivenessProbe = LivelinessProbParameter
	//deploymentRaw.Spec.Template.Spec.Containers[0].ReadinessProbe = ReadinessProbParameter
	return LivelinessProbParameter, ReadinessProbParameter, nil
}

func appendImage(image string, chartFile *types.CoreComponentsChartValues) (string, error) {
	chartFile.Image = image
	//deploymentRaw.Spec.Template.Spec.Containers[0].Image = ImageHelmParameter

	return ImageHelmParameter, nil

}

func appendImagePullPolicy(pullPolicy string, chartFile *types.CoreComponentsChartValues) (string, error) {
	chartFile.ImagePullPolicy = pullPolicy
	//
	return ImagePullPolicyParameter, nil
}

func appendImagePullSecret(imagePullSecret []v12.LocalObjectReference, chartFile *types.CoreComponentsChartValues) (string, error) {
	chartFile.ImagePullSecret = imagePullSecret
	return ImagePullSecret, nil
}
func appendPorts(ports []v12.ContainerPort, chartFile *types.CoreComponentsChartValues) (string, error) {
	chartFile.Ports = ports
	return PortsParameters, nil
}
func appendEnvs(envs []v12.EnvVar, chartFile *types.CoreComponentsChartValues) (string, error) {
	chartFile.Env = envs
	return EnvParameters, nil
}
func appendExtraStatements(deployment string, findString, appendString string) string {
	if index := strings.Index(deployment, "resources:"); index != -1 {

	}
	if index := strings.Index(deployment, findString); index != -1 {
		deployment = deployment[:index-3] + appendString + deployment[index-1:]
		index := strings.Index(deployment, findString)
		for i := index; i <= len(deployment); i++ {
			if deployment[i] == '\n' {
				deployment = deployment[:i] + EndParameter + deployment[i:]
				break
			}
		}

	}
	return deployment
}

func appendIfStatements(str string, findStr, appStr string) string {
	if index := strings.Index(str, findStr); index != -1 {
		str = appStr + "\n" + str
	}
	str = str + "\n" + "{{- end }}"
	return str
}

func appendHpaMinReplicas(minReplicas int32, chartValues *types.HPAChartValues) string {
	chartValues.MinReplicas = minReplicas
	return HpaMinReplicas
}

func appendHpaMaxReplicas(maxReplicas int32, chartValues *types.HPAChartValues) string {
	chartValues.MaxReplicas = maxReplicas
	return HpaMaxReplicas
}

func appendCpuUtilization(cpu int32, chartValues *types.HPAChartValues) string {
	chartValues.TargetCPUUtilizationPercentage = cpu
	return HpaCpuUtilization
}

func extraParametersReplacement(dep []byte, name string) string {
	depString := strings.ReplaceAll(string(dep), "'{{", "{{")
	depString = strings.ReplaceAll(depString, "}}'", "}}")
	depString = strings.ReplaceAll(depString, "{{ .Name }}", name)
	depString = appendExtraStatements(depString, "readinessProbe:", ReadinessProbIfCondition)
	depString = appendExtraStatements(depString, "resources:", ResourcesIfCondition)
	depString = appendExtraStatements(depString, "livenessProbe:", LivelinessProbIfCondition)
	depString = appendExtraStatements(depString, "imagePullSecrets:", ImagePullSecretIfCondition)
	return depString
}
