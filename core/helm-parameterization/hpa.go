package helm_parameterization

import (
	"istio-service-mesh/core/helm-parameterization/types"
	v1 "k8s.io/api/autoscaling/v1"
	"sigs.k8s.io/yaml"
	"strings"
)

func HPAParameters(hpa *v1.HorizontalPodAutoscaler) (hpaYaml []byte, hpaParams []byte, functionsData []byte, err error) {

	result, err := yaml.Marshal(hpa)
	if err != nil {
		return nil, nil, nil, err
	}

	hpaRaw := new(types.HPATemplate)
	err = yaml.Unmarshal(result, hpaRaw)
	if err != nil {
		return nil, nil, nil, err
	}
	tplFile := new([]byte)
	_ = tplFile

	chartFile := new(types.HPAChartValues)

	hpaRaw.Labels, _ = appendLabels(hpa.Labels, hpa.Name, tplFile)
	hpaRaw.Spec.MinReplicas = appendHpaMinReplicas(*hpa.Spec.MinReplicas, chartFile)
	hpaRaw.Spec.MaxReplicas = appendHpaMaxReplicas(hpa.Spec.MaxReplicas, chartFile)

	if hpa.Spec.TargetCPUUtilizationPercentage != nil {
		hpaRaw.Spec.TargetCPUUtilizationPercentage = *hpa.Spec.TargetCPUUtilizationPercentage

	}
	hpaRaw.Spec.ScaleTargetRef.Name, _ = appendRefName(hpa.Spec.ScaleTargetRef.Name)
	hpaRaw.Name, _ = appendName(hpa.Name, tplFile)
	chartFile.Enabled = true

	dep, err := yaml.Marshal(hpaRaw)
	if err != nil {
		return nil, nil, nil, err
	}

	chartRaw, err := yaml.Marshal(chartFile)
	if err != nil {
		return nil, nil, nil, err
	}

	depString := strings.ReplaceAll(string(dep), "'{{", "{{")
	depString = strings.ReplaceAll(depString, "}}'", "}}")

	depString = appendIfStatements(depString, "apiVersion", HpaIfCondition)
	return []byte(depString), chartRaw, *tplFile, nil

}
