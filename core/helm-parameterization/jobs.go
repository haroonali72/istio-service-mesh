package helm_parameterization

import (
	"istio-service-mesh/core/helm-parameterization/types"
	v12 "k8s.io/api/batch/v1"
	"sigs.k8s.io/yaml"
)

func JobParameters(job *v12.Job) (jobYaml []byte, jobParams []byte, functionsData []byte, err error) {
	result, err := yaml.Marshal(job)
	if err != nil {
		return nil, nil, nil, err
	}
	jobRaw := new(types.JobTemplate)
	err = yaml.Unmarshal(result, jobRaw)
	if err != nil {
		return nil, nil, nil, err
	}
	tplFile := new([]byte)
	_ = tplFile
	chartFile := new(types.CoreComponentsChartValues)

	jobRaw.Labels, _ = appendLabels(job.Labels, job.Name, tplFile)
	jobRaw.Spec.Selector.MatchLabels, _ = appendMatchLabels(job.Spec.Selector.MatchLabels, job.Name, tplFile)
	jobRaw.Spec.Template.Labels, _ = appendPodLabels(job.Spec.Template.Labels, job.Name)
	jobRaw.Spec.Template.Spec.Containers[0].Resources, _ = appendResourceQuota(job.Spec.Template.Spec.Containers[0].Resources, chartFile)
	jobRaw.Spec.Template.Spec.Containers[0].LivenessProbe, jobRaw.Spec.Template.Spec.Containers[0].ReadinessProbe, _ = appendProbing(job.Spec.Template.Spec.Containers[0].LivenessProbe, job.Spec.Template.Spec.Containers[0].ReadinessProbe, chartFile)
	jobRaw.Spec.Template.Spec.Containers[0].Image, _ = appendImage(job.Spec.Template.Spec.Containers[0].Image, chartFile)
	jobRaw.Spec.Template.Spec.Containers[0].ImagePullPolicy, _ = appendImagePullPolicy(string(job.Spec.Template.Spec.Containers[0].ImagePullPolicy), chartFile)
	jobRaw.Spec.Template.Spec.ImagePullSecrets, _ = appendImagePullSecret(job.Spec.Template.Spec.ImagePullSecrets, chartFile)
	jobRaw.Spec.Template.Spec.Containers[0].Ports, _ = appendPorts(job.Spec.Template.Spec.Containers[0].Ports, chartFile)
	jobRaw.Spec.Template.Spec.Containers[0].Env, _ = appendEnvs(job.Spec.Template.Spec.Containers[0].Env, chartFile)

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

	//jobYamlStr := strings.ReplaceAll(string(jobYaml), "'{{", "{{")
	//jobYamlStr = strings.ReplaceAll(jobYamlStr, "}}'", "}}")
	//
	//jobYamlStr = appendExtraStatements(jobYamlStr, "readinessProbe:", ReadinessProbIfCondition)
	//jobYamlStr = appendExtraStatements(jobYamlStr, "resources:", ResourcesIfCondition)
	//jobYamlStr = appendExtraStatements(jobYamlStr, "livenessProbe:", LivelinessProbIfCondition)
	//jobYamlStr = appendExtraStatements(jobYamlStr, "imagePullSecrets:", ImagePullSecretIfCondition)
	jobYamlStr := extraParametersReplacement(jobYaml, job.Name)
	return []byte(jobYamlStr), chartRaw, *tplFile, nil
}
