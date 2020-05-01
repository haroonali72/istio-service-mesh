package helm_parameterization

import (
	"bitbucket.org/cloudplex-devs/istio-service-mesh/core/helm-parameterization/types"
	core "k8s.io/api/core/v1"
	"sigs.k8s.io/yaml"
	"strconv"
	"strings"
)

func KubernetesServiceParameters(Service *core.Service) (ServiceYaml []byte, ServiceParams []byte, functionsData []byte, err error) {
	result, err := yaml.Marshal(Service)
	if err != nil {
		return nil, nil, nil, err
	}
	ServiceRaw := new(types.KubernetesServiceTemplate)
	err = yaml.Unmarshal(result, ServiceRaw)
	if err != nil {
		return nil, nil, nil, err
	}

	ServiceParams = []byte("\n" + Service.Name + "SVC:\n  ports:")
	for _, each := range Service.Spec.Ports {
		ServiceParams = append(ServiceParams, []byte("\n    - port: "+strconv.FormatInt(int64(each.Port), 10))...)
		if each.Name != "" {
			ServiceParams = append(ServiceParams, []byte("\n      name: "+each.Name)...)
		}
		if each.NodePort != 0 {
			ServiceParams = append(ServiceParams, []byte("\n      nodePort: "+strconv.FormatInt(int64(each.NodePort), 10))...)
		}
		if each.Protocol != "" {
			ServiceParams = append(ServiceParams, []byte("\n      protocol: "+each.Protocol)...)
		}
		if each.TargetPort.String() != "" {
			ServiceParams = append(ServiceParams, []byte("\n      targetPort: "+each.TargetPort.String())...)

		}
	}
	//ConfigMapRaw.Data="{{ index .Values \""+ ConfigMap.Name+"CM\" \"data\" | toYaml | trim | nindent 2 }}"
	ServiceRaw.Spec.Ports = "{{ index .Values \"" + Service.Name + "SVC\" \"ports\" | toYaml | trim | nindent 4 }}"
	if Service.Spec.Type != "" {
		ServiceParams = append(ServiceParams, []byte("\n  type: "+Service.Spec.Type)...)
		//		ServiceRaw.Spec.Type="{{ .Values."+ Service.Name+"SVC.type }}"
		ServiceRaw.Spec.Type = "{{ index .Values \"" + Service.Name + "SVC\" \"type\" }}"

	}
	if Service.Spec.ExternalTrafficPolicy != "" {
		ServiceParams = append(ServiceParams, []byte("\n  externalTrafficPolicy: "+Service.Spec.ExternalTrafficPolicy)...)
		//		ServiceRaw.Spec.ExternalTrafficPolicy="{{ .Values."+ Service.Name+"SVC.externalTrafficPolicy }}"
		ServiceRaw.Spec.ExternalTrafficPolicy = "{{ index .Values \"" + Service.Name + "SVC\" \"externalTrafficPolicy\" }}"
	}
	ServiceYaml, err = yaml.Marshal(ServiceRaw)
	temp := strings.ReplaceAll(string(ServiceYaml), "'{{", "{{")
	temp = strings.ReplaceAll(temp, "}}'", "}}")
	ServiceYaml = []byte(temp)
	return
}
