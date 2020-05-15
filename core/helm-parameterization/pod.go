package helm_parameterization

import (
	"bitbucket.org/cloudplex-devs/istio-service-mesh/core/helm-parameterization/types"
	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/yaml"
)

type PodStructs struct {
	ChartFile *types.CoreComponentsChartValues
}

func PodParameters(pod *v1.Pod) (podYaml []byte, podParams []byte, functionsData []byte, err error) {
	result, err := yaml.Marshal(pod)
	if err != nil {
		return nil, nil, nil, err
	}
	podRaw := new(types.PodSpecTemplate)
	err = yaml.Unmarshal(result, podRaw)
	if err != nil {
		return nil, nil, nil, err
	}
	tplFile := new([]byte)
	_ = tplFile
	chartFile := new(types.CoreComponentsChartValues)

	podRaw.Labels, _ = appendLabels(pod.Labels, pod.Name, tplFile)

	podRaw.Spec.Containers[0].Resources, _ = appendResourceQuota(pod.Spec.Containers[0].Resources, chartFile)
	podRaw.Spec.Containers[0].LivenessProbe, podRaw.Spec.Containers[0].ReadinessProbe, _ = appendProbing(pod.Spec.Containers[0].LivenessProbe, pod.Spec.Containers[0].ReadinessProbe, chartFile)
	podRaw.Spec.Containers[0].Image, _ = appendImage(pod.Spec.Containers[0].Image, chartFile)
	podRaw.Spec.Containers[0].ImagePullPolicy, _ = appendImagePullPolicy(string(pod.Spec.Containers[0].ImagePullPolicy), chartFile)
	podRaw.Spec.ImagePullSecrets, _ = appendImagePullSecret(pod.Spec.ImagePullSecrets, chartFile)
	podRaw.Spec.Containers[0].Ports, _ = appendPorts(pod.Spec.Containers[0].Ports, chartFile)
	podRaw.Spec.Containers[0].Env, _ = appendEnvs(pod.Spec.Containers[0].Env, chartFile)

	if len(pod.Annotations) > 1 {
		podRaw.Annotations, _ = appendAnnotations(pod.Annotations, pod.Name, tplFile)
	}
	//add this at the end. This function will replace name with helm parameter
	podRaw.Name, _ = appendName(pod.Name, tplFile)

	dep, err := yaml.Marshal(podRaw)
	if err != nil {
		return nil, nil, nil, err
	}

	chartRaw, err := yaml.Marshal(chartFile)
	if err != nil {
		return nil, nil, nil, err
	}

	depString := extraParametersReplacement(dep, pod.Name)
	return []byte(depString), chartRaw, *tplFile, nil
}
