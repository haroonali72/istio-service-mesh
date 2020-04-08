package helm_parameterization

import (
	"istio-service-mesh/core/helm-parameterization/types"
	v1 "k8s.io/api/apps/v1"
	"sigs.k8s.io/yaml"
	"strings"
)

type DeploymentStructs struct {
	ChartFile *types.CoreComponentsChartValues
}

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

	deploymentRaw.Labels, _ = appendLabels(deployment.Labels, deployment.Name, tplFile)
	deploymentRaw.Spec.Selector.MatchLabels, _ = appendMatchLabels(deployment.Spec.Selector.MatchLabels, deployment.Name, tplFile)
	deploymentRaw.Spec.Template.Labels, _ = appendPodLabels(deployment.Spec.Template.Labels, deployment.Name)
	deploymentRaw.Spec.Replicas, _ = appendReplicasTemplate(*deployment.Spec.Replicas, chartFile)
	deploymentRaw.Spec.Template.Spec.Containers[0].Resources, _ = appendResourceQuota(deployment.Spec.Template.Spec.Containers[0].Resources, chartFile)
	deploymentRaw.Spec.Template.Spec.Containers[0].LivenessProbe, deploymentRaw.Spec.Template.Spec.Containers[0].ReadinessProbe, _ = appendProbing(deployment.Spec.Template.Spec.Containers[0].LivenessProbe, deployment.Spec.Template.Spec.Containers[0].ReadinessProbe, chartFile)
	deploymentRaw.Spec.Template.Spec.Containers[0].Image, _ = appendImage(deployment.Spec.Template.Spec.Containers[0].Image, chartFile)
	deploymentRaw.Spec.Template.Spec.Containers[0].ImagePullPolicy, _ = appendImagePullPolicy(string(deployment.Spec.Template.Spec.Containers[0].ImagePullPolicy), chartFile)
	deploymentRaw.Spec.Template.Spec.ImagePullSecrets, _ = appendImagePullSecret(deployment.Spec.Template.Spec.ImagePullSecrets, chartFile)
	deploymentRaw.Spec.Template.Spec.Containers[0].Ports, _ = appendPorts(deployment.Spec.Template.Spec.Containers[0].Ports, chartFile)

	//add this at the end. This function will replace name with helm parameter
	deploymentRaw.Name, _ = appendName(deployment.Name, tplFile)

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
