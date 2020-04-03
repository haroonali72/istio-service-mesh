package helm_parameterization

import (
	"encoding/json"
	"istio-service-mesh/core/helm-parameterization/types"
	v1 "k8s.io/api/apps/v1"
	v12 "k8s.io/api/core/v1"
	"sigs.k8s.io/yaml"
	"strings"
)

func DeploymentParameters(deployment *v1.Deployment) (deploymentYaml []byte, deploymentParams []byte, functionsData []byte, err error) {
	result, err := yaml.Marshal(deployment)
	if err != nil {
		return nil, nil, nil, err
	}
	deploymentRaw := new(types.DeploymentTemplate)
	err = yaml.Unmarshal(result, deploymentRaw)
	if err != nil {
		return nil, nil, nil, err
	}
	tplFile := new([]byte)
	_ = tplFile
	chartFile := new(types.CoreComponentsChartValues)

	_ = appendLabels(deploymentRaw, tplFile)
	_ = appendMatchLabels(deploymentRaw, tplFile)
	_ = appendPodLabels(deploymentRaw, tplFile)
	_ = appendReplicasTemplate(deploymentRaw, chartFile, *deployment.Spec.Replicas)
	_ = appendResourceQuota(deploymentRaw, chartFile, deployment.Spec.Template.Spec.Containers[0].Resources)
	_ = appendProbing(deploymentRaw, chartFile, deployment.Spec.Template.Spec.Containers[0].LivenessProbe, deployment.Spec.Template.Spec.Containers[0].ReadinessProbe)
	_ = appendImage(deploymentRaw, chartFile, deployment.Spec.Template.Spec.Containers[0].Image)
	_ = appendImagePullPolicy(deploymentRaw, chartFile, string(deployment.Spec.Template.Spec.Containers[0].ImagePullPolicy))
	_ = appendImagePullSecret(deploymentRaw, chartFile)
	_ = appendPorts(deploymentRaw, chartFile)

	//add this at the end. This function will replace name with helm parameter
	_ = appendName(deploymentRaw, tplFile)
	dep, err := yaml.Marshal(deploymentRaw)
	if err != nil {
		return nil, nil, nil, err
	}

	chartRaw, err := yaml.Marshal(chartFile)
	if err != nil {
		return nil, nil, nil, err
	}

	depString := strings.ReplaceAll(string(dep), "'{{", "{{")

	depString = strings.ReplaceAll(depString, "}}'", "}}")
	depString = appendExtraStatements(depString, "readinessProbe:", ReadinessProbIfCondition)
	depString = appendExtraStatements(depString, "resources:", ResourcesIfCondition)
	depString = appendExtraStatements(depString, "livenessProbe:", LivelinessProbIfCondition)
	depString = appendExtraStatements(depString, "imagePullSecrets:", ImagePullSecretIfCondition)
	return []byte(depString), chartRaw, *tplFile, nil
}
func appendName(deploymentRaw *types.DeploymentTemplate, tplFile *[]byte) error {
	name := strings.Replace(NameHelmParameter, "{{ .Name }}", deploymentRaw.Name.(string), -1)
	var nameInterface interface{}
	_ = json.Unmarshal([]byte(name), &nameInterface)

	nameGenFunc := strings.ReplaceAll(NameFunction, "{{ .Name }}", deploymentRaw.Name.(string))
	deploymentRaw.Name = name
	*tplFile = append(*tplFile, []byte(nameGenFunc)...)
	return nil
}

func appendLabels(deploymentRaw *types.DeploymentTemplate, tplFile *[]byte) error {
	rawLabels, err := yaml.Marshal(deploymentRaw.Labels)
	if err != nil {
		return err
	}
	labelFunc := strings.ReplaceAll(LabelFunction, "{{ .Name }}", deploymentRaw.Name.(string))
	labelFunc = strings.ReplaceAll(labelFunc, "{{ .Labels }}", string(rawLabels))

	labelTemplate := strings.Replace(LabelParameter, "{{ .Name }}", deploymentRaw.Name.(string), -1)
	labelTemplate = strings.Replace(labelTemplate, "{{ .Indent }}", "4", -1)
	deploymentRaw.Labels = labelTemplate

	*tplFile = append(*tplFile, []byte(labelFunc)...)
	return nil
}

func appendMatchLabels(deploymentRaw *types.DeploymentTemplate, tplFile *[]byte) error {
	rawLabels, err := yaml.Marshal(deploymentRaw.Spec.Selector.MatchLabels)
	if err != nil {
		return err
	}
	labelFunc := strings.ReplaceAll(LabelFunction, "{{ .Name }}", deploymentRaw.Name.(string))
	labelFunc = strings.ReplaceAll(labelFunc, "{{ .Labels }}", string(rawLabels))
	labelFunc = strings.ReplaceAll(labelFunc, deploymentRaw.Name.(string)+".labels", deploymentRaw.Name.(string)+".matchLabels")

	labelTemplate := strings.Replace(MatchSelectorParameter, "{{ .Name }}", deploymentRaw.Name.(string), -1)
	labelTemplate = strings.Replace(labelTemplate, "{{ .Indent }}", "6", -1)
	deploymentRaw.Spec.Selector.MatchLabels = labelTemplate

	*tplFile = append(*tplFile, []byte(labelFunc)...)
	return nil
}

func appendPodLabels(deploymentRaw *types.DeploymentTemplate, tplFile *[]byte) error {
	rawLabels, err := yaml.Marshal(deploymentRaw.Spec.Template.Labels)
	if err != nil {
		return err
	}
	labelFunc := strings.ReplaceAll(LabelFunction, "{{ .Name }}", deploymentRaw.Name.(string))
	labelFunc = strings.ReplaceAll(labelFunc, "{{ .Labels }}", string(rawLabels))

	labelTemplate := strings.Replace(LabelParameter, "{{ .Name }}", deploymentRaw.Name.(string), -1)
	labelTemplate = strings.Replace(labelTemplate, "{{ .Indent }}", "8", -1)
	deploymentRaw.Spec.Template.Labels = labelTemplate

	//*tplFile = append(*tplFile, []byte(labelFunc)...)
	return nil
}

func appendReplicasTemplate(deploymentRaw *types.DeploymentTemplate, chartFile *types.CoreComponentsChartValues, replicas int32) error {
	deploymentRaw.Spec.Replicas = ReplicasHelmParameter
	chartFile.Replicas = replicas
	return nil
}

func appendResourceQuota(deploymentRaw *types.DeploymentTemplate, chartFile *types.CoreComponentsChartValues, rq v12.ResourceRequirements) error {
	chartFile.ResourceQuota = rq
	deploymentRaw.Spec.Template.Spec.Containers[0].Resources = ResourcesParameter
	return nil
}

func appendProbing(deploymentRaw *types.DeploymentTemplate, chartFile *types.CoreComponentsChartValues, liveness, readiness *v12.Probe) error {
	chartFile.LivenessProb = liveness
	chartFile.ReadinessProbe = readiness
	deploymentRaw.Spec.Template.Spec.Containers[0].LivenessProbe = LivelinessProbParameter
	deploymentRaw.Spec.Template.Spec.Containers[0].ReadinessProbe = ReadinessProbParameter
	return nil
}

func appendImage(deploymentRaw *types.DeploymentTemplate, chartFile *types.CoreComponentsChartValues, image string) error {
	deploymentRaw.Spec.Template.Spec.Containers[0].Image = ImageHelmParameter
	chartFile.Image = image
	return nil

}

func appendImagePullPolicy(deploymentRaw *types.DeploymentTemplate, chartFile *types.CoreComponentsChartValues, pullPolicy string) error {
	deploymentRaw.Spec.Template.Spec.Containers[0].ImagePullPolicy = ImagePullPolicyParameter
	chartFile.ImagePullPolicy = pullPolicy
	return nil
}

func appendImagePullSecret(deploymentRaw *types.DeploymentTemplate, chartFile *types.CoreComponentsChartValues) error {
	chartFile.ImagePullSecret = deploymentRaw.Spec.Template.Spec.ImagePullSecrets
	deploymentRaw.Spec.Template.Spec.ImagePullSecrets = ImagePullSecret
	return nil
}
func appendPorts(deploymentRaw *types.DeploymentTemplate, chartFile *types.CoreComponentsChartValues) error {
	chartFile.Ports = deploymentRaw.Spec.Template.Spec.Containers[0].Ports
	deploymentRaw.Spec.Template.Spec.Containers[0].Ports = PortsParameters
	return nil
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
