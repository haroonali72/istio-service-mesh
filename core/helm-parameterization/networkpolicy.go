package helm_parameterization

import (
	"istio-service-mesh/core/helm-parameterization/types"
	net "k8s.io/api/networking/v1"
	"sigs.k8s.io/yaml"
	"strconv"
	"strings"
)

func NetworkPolicyParameters(NetworkPolicy *net.NetworkPolicy) (NetworkPolicyYaml []byte, NetworkPolicyParams []byte, functionsData []byte, err error) {
	result, err := yaml.Marshal(NetworkPolicy)
	if err != nil {
		return nil, nil, nil, err
	}
	Raw := new(types.NetworkPolicyTemplate)
	err = yaml.Unmarshal(result, Raw)
	if err != nil {
		return nil, nil, nil, err
	}

	NetworkPolicyParams = []byte("\n" + NetworkPolicy.Name + "NP:")
	if len(NetworkPolicy.Spec.PolicyTypes) > 0 {
		NetworkPolicyParams = append(NetworkPolicyParams, []byte("\n  policyTypes:")...)
		Raw.Spec.PolicyTypes = "{{ index .Values \"" + NetworkPolicy.Name + "NP\" \"policyTypes\" | toYaml | trim | nindent 2}}"
	}
	for _, each := range NetworkPolicy.Spec.PolicyTypes {
		NetworkPolicyParams = append(NetworkPolicyParams, []byte("\n  - "+each)...)
	}

	if len(NetworkPolicy.Spec.Egress) > 0 {
		NetworkPolicyParams = append(NetworkPolicyParams, []byte("\n  egress:")...)
	}
	for index, each := range NetworkPolicy.Spec.Egress {
		NetworkPolicyParams = append(NetworkPolicyParams, []byte("\n    "+strconv.Itoa(index)+":")...)
		if len(each.Ports) > 0 {
			NetworkPolicyParams = append(NetworkPolicyParams, []byte("\n      ports:")...)
		}
		for index2, value := range each.Ports {
			NetworkPolicyParams = append(NetworkPolicyParams, []byte("\n        "+strconv.Itoa(index2)+":")...)
			if value.Protocol != nil {
				NetworkPolicyParams = append(NetworkPolicyParams, []byte("\n          protocol: "+*value.Protocol)...)
				Raw.Spec.Egress[index].Ports[index2].Protocol = "{{ index .Values \"" + NetworkPolicy.Name + "NP\" \"egress\" \"" + strconv.Itoa(index) + "\" \"ports\" \"" + strconv.Itoa(index2) + "\" \"protocol\" }}"
			}
			if value.Port != nil {
				NetworkPolicyParams = append(NetworkPolicyParams, []byte("\n          port: "+value.Port.String())...)
				Raw.Spec.Egress[index].Ports[index2].Port = "{{ index .Values \"" + NetworkPolicy.Name + "NP\" \"egress\" \"" + strconv.Itoa(index) + "\" \"ports\" \"" + strconv.Itoa(index2) + "\" \"port\" }}"
			}
		}
		if len(each.To) > 0 {
			NetworkPolicyParams = append(NetworkPolicyParams, []byte("\n      to:")...)
		}
		for index2, value := range each.To {
			NetworkPolicyParams = append(NetworkPolicyParams, []byte("\n        "+strconv.Itoa(index2)+":")...)
			if value.IPBlock == nil {
				continue
			}
			Raw.Spec.Egress[index].To[index2].IPBlock = "{{ index .Values \"" + NetworkPolicy.Name + "NP\" \"egress\" \"" + strconv.Itoa(index) + "\" \"to\" \"" + strconv.Itoa(index2) + "\" \"ipBlock\" | toYaml | trim | nindent 8 }}"
			NetworkPolicyParams = append(NetworkPolicyParams, []byte("\n          ipBlock:")...)
			NetworkPolicyParams = append(NetworkPolicyParams, []byte("\n            cidr: "+value.IPBlock.CIDR)...)
			if len(value.IPBlock.Except) > 0 {
				NetworkPolicyParams = append(NetworkPolicyParams, []byte("\n            except:")...)
			}
			for _, cidr := range value.IPBlock.Except {
				NetworkPolicyParams = append(NetworkPolicyParams, []byte("\n            - "+cidr)...)
			}
		}
		//		NetworkPolicyParams =append(NetworkPolicyParams,[]byte("\n    - "+each)...)
	}

	if len(NetworkPolicy.Spec.Ingress) > 0 {
		NetworkPolicyParams = append(NetworkPolicyParams, []byte("\n  ingress:")...)
	}
	for index, each := range NetworkPolicy.Spec.Ingress {
		NetworkPolicyParams = append(NetworkPolicyParams, []byte("\n    "+strconv.Itoa(index)+":")...)
		if len(each.Ports) > 0 {
			NetworkPolicyParams = append(NetworkPolicyParams, []byte("\n      ports:")...)
		}
		for index2, value := range each.Ports {
			NetworkPolicyParams = append(NetworkPolicyParams, []byte("\n        "+strconv.Itoa(index2)+":")...)
			if value.Protocol != nil {
				NetworkPolicyParams = append(NetworkPolicyParams, []byte("\n          protocol: "+*value.Protocol)...)
				Raw.Spec.Ingress[index].Ports[index2].Protocol = "{{ index .Values \"" + NetworkPolicy.Name + "NP\" \"ingress\" \"" + strconv.Itoa(index) + "\" \"ports\" \"" + strconv.Itoa(index2) + "\" \"protocol\" }}"
			}
			if value.Port != nil {
				NetworkPolicyParams = append(NetworkPolicyParams, []byte("\n          port: "+value.Port.String())...)
				Raw.Spec.Ingress[index].Ports[index2].Port = "{{ index .Values \"" + NetworkPolicy.Name + "NP\" \"ingress\" \"" + strconv.Itoa(index) + "\" \"ports\" \"" + strconv.Itoa(index2) + "\" \"port\" }}"
			}
		}
		if len(each.From) > 0 {
			NetworkPolicyParams = append(NetworkPolicyParams, []byte("\n      from:")...)
		}
		for index2, value := range each.From {
			NetworkPolicyParams = append(NetworkPolicyParams, []byte("\n        "+strconv.Itoa(index2)+":")...)
			if value.IPBlock == nil {
				continue
			}
			Raw.Spec.Ingress[index].From[index2].IPBlock = "{{ index .Values \"" + NetworkPolicy.Name + "NP\" \"ingress\" \"" + strconv.Itoa(index) + "\" \"from\" \"" + strconv.Itoa(index2) + "\" \"ipBlock\" | toYaml | trim | nindent 8 }}"
			NetworkPolicyParams = append(NetworkPolicyParams, []byte("\n          ipBlock:")...)
			NetworkPolicyParams = append(NetworkPolicyParams, []byte("\n            cidr: "+value.IPBlock.CIDR)...)
			if len(value.IPBlock.Except) > 0 {
				NetworkPolicyParams = append(NetworkPolicyParams, []byte("\n            except:")...)
			}
			for _, cidr := range value.IPBlock.Except {
				NetworkPolicyParams = append(NetworkPolicyParams, []byte("\n            - "+cidr)...)
			}
		}
		//		NetworkPolicyParams =append(NetworkPolicyParams,[]byte("\n    - "+each)...)
	}
	NetworkPolicyYaml, err = yaml.Marshal(Raw)
	temp := strings.ReplaceAll(string(NetworkPolicyYaml), "'{{", "{{")
	temp = strings.ReplaceAll(temp, "}}'", "}}")
	NetworkPolicyYaml = []byte(temp)
	return
}
