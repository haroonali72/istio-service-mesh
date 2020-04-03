package helm_parameterization

import (
	"istio-service-mesh/core/helm-parameterization/types"
	"k8s.io/api/batch/v1beta1"
	"sigs.k8s.io/yaml"
	"strings"
)

func CronJobParameters(job v1beta1.CronJob) (jobYaml []byte, jobParams []byte, functionsData []byte, err error) {
	result, err := yaml.Marshal(job)
	if err != nil {
		return nil, nil, nil, err
	}
	jobRaw := new(types.CronJobTemplate)
	err = yaml.Unmarshal(result, jobRaw)
	if err != nil {
		return nil, nil, nil, err
	}
	tplFile := new([]byte)
	_ = tplFile
	chartFile := new(types.CoreComponentsChartValues)

	jobRaw.Labels, _ = appendLabels(job.Labels, job.Name, tplFile)
	jobRaw.Spec.JobTemplate.Spec.Selector.MatchLabels, _ = appendMatchLabels(job.Spec.JobTemplate.Spec.Selector.MatchLabels, job.Name, tplFile)
	jobRaw.Spec.JobTemplate.Spec.Template.Labels, _ = appendPodLabels(job.Spec.JobTemplate.Spec.Template.Labels, job.Name)
	jobRaw.Spec.JobTemplate.Spec.Template.Spec.Containers[0].Resources, _ = appendResourceQuota(job.Spec.JobTemplate.Spec.Template.Spec.Containers[0].Resources, chartFile)
	jobRaw.Spec.JobTemplate.Spec.Template.Spec.Containers[0].LivenessProbe, jobRaw.Spec.JobTemplate.Spec.Template.Spec.Containers[0].ReadinessProbe, _ = appendProbing(job.Spec.JobTemplate.Spec.Template.Spec.Containers[0].LivenessProbe, job.Spec.JobTemplate.Spec.Template.Spec.Containers[0].ReadinessProbe, chartFile)
	jobRaw.Spec.JobTemplate.Spec.Template.Spec.Containers[0].Image, _ = appendImage(job.Spec.JobTemplate.Spec.Template.Spec.Containers[0].Image, chartFile)
	jobRaw.Spec.JobTemplate.Spec.Template.Spec.Containers[0].ImagePullPolicy, _ = appendImagePullPolicy(string(job.Spec.JobTemplate.Spec.Template.Spec.Containers[0].ImagePullPolicy), chartFile)
	jobRaw.Spec.JobTemplate.Spec.Template.Spec.ImagePullSecrets, _ = appendImagePullSecret(job.Spec.JobTemplate.Spec.Template.Spec.ImagePullSecrets, chartFile)
	jobRaw.Spec.JobTemplate.Spec.Template.Spec.Containers[0].Ports, _ = appendPorts(job.Spec.JobTemplate.Spec.Template.Spec.Containers[0].Ports, chartFile)
	jobRaw.Spec.Schedule, _ = appendCronExpression(job.Spec.Schedule, chartFile)
	//add this at the end. This function will replace name with helm parameter
	jobRaw.Name, _ = appendName(job.Name, tplFile)

	jobYaml, err = yaml.Marshal(jobRaw)
	if err != nil {
		return nil, nil, nil, err
	}

	chartRaw, err := yaml.Marshal(chartFile)
	if err != nil {
		return nil, nil, nil, err
	}

	jobYamlStr := strings.ReplaceAll(string(jobYaml), "'{{", "{{")
	jobYamlStr = strings.ReplaceAll(jobYamlStr, "}}'", "}}")

	jobYamlStr = appendExtraStatements(jobYamlStr, "readinessProbe:", ReadinessProbIfCondition)
	jobYamlStr = appendExtraStatements(jobYamlStr, "resources:", ResourcesIfCondition)
	jobYamlStr = appendExtraStatements(jobYamlStr, "livenessProbe:", LivelinessProbIfCondition)
	jobYamlStr = appendExtraStatements(jobYamlStr, "imagePullSecrets:", ImagePullSecretIfCondition)
	return []byte(jobYamlStr), chartRaw, *tplFile, nil
}
func appendCronExpression(schedule string, chartFile *types.CoreComponentsChartValues) (string, error) {
	chartFile.CronExpression = schedule
	return CronExpressionParameter, nil
}
