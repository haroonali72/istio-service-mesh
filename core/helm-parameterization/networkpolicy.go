package helm_parameterization

//
//import (
//	"istio-service-mesh/core/helm-parameterization/types"
//	net "k8s.io/api/networking/v1"
//	"sigs.k8s.io/yaml"
//	"strings"
//)
//func NetworkPolicyParameters(NetworkPolicy *net.NetworkPolicy ) (NetworkPolicyYaml []byte, NetworkPolicyParams []byte, functionsData []byte, err error) {
//	result, err := yaml.Marshal(NetworkPolicy)
//	if err != nil {
//		return nil, nil, nil, err
//	}
//	Raw := new(types.NetworkPolicyTemplate)
//	err = yaml.Unmarshal(result, Raw)
//	if err != nil {
//		return nil, nil, nil, err
//	}
//
//
//	NetworkPolicyParams = []byte("\n"+ NetworkPolicy.Name+"NP:")
//if len(NetworkPolicy.Spec.PolicyTypes)>0{
//	NetworkPolicyParams =append(NetworkPolicyParams,[]byte("\n  policyTypes:")...)
//	Raw.Spec.PolicyTypes="{{ .Values."+ NetworkPolicy.Name+"NP.policyTypes }}"
//}
//	for _,each:=range NetworkPolicy.Spec.PolicyTypes{
//		NetworkPolicyParams =append(NetworkPolicyParams,[]byte("\n  - "+each)...)
//	}
//
//	if len(NetworkPolicy.Spec.Ingress)>0{
//		NetworkPolicyParams =append(NetworkPolicyParams,[]byte("\n  ingress:")...)
//	//	Raw.Spec.Ingress=
//	}
//	for index,each:=range NetworkPolicy.Spec.Ingress{
//		NetworkPolicyParams =append(NetworkPolicyParams,[]byte("\n    "+string(index)+":")...)
//		if len(each.From)>0{
//			NetworkPolicyParams =append(NetworkPolicyParams,[]byte("\n    - from:")...)
//		}
//		for index2,value:=range each.From{
//			if value.IPBlock==nil {
//				continue
//			}
//			Raw.Spec.Ingress[index].From[index2].IPBlock="{{ .Values."+ NetworkPolicy.Name+"NP.ingress.from. }}"
//				NetworkPolicyParams =append(NetworkPolicyParams,[]byte("\n    - ipBlock:")...)
//				NetworkPolicyParams =append(NetworkPolicyParams,[]byte("\n      cidr:"+value.IPBlock.CIDR)...)
//			if len(value.IPBlock.Except)>0{
//				NetworkPolicyParams =append(NetworkPolicyParams,[]byte("\n      except:"+value.IPBlock.CIDR)...)
//			}
//			for _,cidr:=range value.IPBlock.Except{
//				NetworkPolicyParams =append(NetworkPolicyParams,[]byte("\n      - :"+cidr)...)
//			}
//		}
////		NetworkPolicyParams =append(NetworkPolicyParams,[]byte("\n    - "+each)...)
//	}
//	if NetworkPolicy.Spec.ExternalTrafficPolicy !=""{
//		NetworkPolicyParams =append(NetworkPolicyParams,[]byte("\n  externalTrafficPolicy: "+NetworkPolicy.Spec.ExternalTrafficPolicy)...)
//		Raw.Spec.ExternalTrafficPolicy="{{ .Values."+ NetworkPolicy.Name+"SVC.externalTrafficPolicy }}"
//	}
//	qunatity,ok:=persistentVServiceolumeClaim.Spec.Resources.Requests["storage"]
//	if ok{
//		persistentVolumeClaimParams=append(persistentVolumeClaimParams,[]byte("\n  resources:")...)
//		persistentVolumeClaimParams=append(persistentVolumeClaimParams,[]byte("\n    requests:\n      storage: "+qunatity.String())...)
//		persistentVolumeClaimRaw.Spec.Resources.Requests["storage"]="{{ .Values."+ persistentVolumeClaim.Name+"PVC.resources.limits }}"
//
//	}
//	qunatity,ok=persistentVolumeClaim.Spec.Resources.Limits["storage"]
//	if ok{
//		persistentVolumeClaimParams=append(persistentVolumeClaimParams,[]byte("\n  resources:")...)
//		persistentVolumeClaimParams=append(persistentVolumeClaimParams,[]byte("\n    limits:\n      storage: "+qunatity.String())...)
//		persistentVolumeClaimRaw.Spec.Resources.Limits["storage"]="{{ .Values."+ persistentVolumeClaim.Name+"PVC.resources.limits }}"
//	}
//	if persistentVolumeClaim.Spec.VolumeMode!=nil{
//		persistentVolumeClaimParams=append(persistentVolumeClaimParams,[]byte("\n  volumeMode: "+*persistentVolumeClaim.Spec.VolumeMode)...)
//		persistentVolumeClaimRaw.Spec.VolumeMode="{{ .Values."+ persistentVolumeClaim.Name+"PVC.volumeMode }}"
//	}
//	persistentVolumeClaimYaml,err=yaml.Marshal(persistentVolumeClaimRaw)
//	temp:=strings.ReplaceAll(string(persistentVolumeClaimYaml),"'{{","{{")
//	temp=strings.ReplaceAll(temp,"}}'","}}")
//	persistentVolumeClaimYaml= []byte(temp)
//	return
//}
//
