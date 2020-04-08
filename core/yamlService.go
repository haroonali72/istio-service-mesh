package core

import (
	"context"
	"encoding/json"
	"fmt"
	"istio-service-mesh/constants"
	helm_parameterization "istio-service-mesh/core/helm-parameterization"
	pb "istio-service-mesh/core/proto"
	"istio-service-mesh/utils"
	"regexp"
	"sigs.k8s.io/yaml"
)

func (s *Server) GetYamlService(ctx context.Context, req *pb.YamlServiceRequest) (*pb.YamlServiceResponse, error) {
	serviceResp := new(pb.YamlServiceResponse)
	switch req.Type {
	case constants.StorageClassServiceType:
		if err := ConvertSCToYaml(req, serviceResp); err != nil {
			return nil, err
		}

	case constants.PVCServiceType:
		if err := ConvertPVCToYaml(req, serviceResp); err != nil {
			return nil, err
		}
	case constants.PVServiceType:
		if err := ConvertPVToYaml(req, serviceResp); err != nil {
			return nil, err
		}
	case constants.GatewayServiceType:
		if err := ConvertGatewayToYaml(req, serviceResp); err != nil {
			return nil, err
		}
	case constants.KubernetesServiceType:
		if err := ConvertKubernetesServiceToYaml(req, serviceResp); err != nil {
			return nil, err
		}
	case constants.NetworkPolicyServiceType:
		if err := ConvertNetworkPolicyToYaml(req, serviceResp); err != nil {
			return nil, err
		}
	case constants.RoleServiceType:
		if err := ConvertRoleToYaml(req, serviceResp); err != nil {
			return nil, err
		}
	case constants.PolicyType:
		if err := ConvertPolicyToYaml(req, serviceResp); err != nil {
			return nil, err
		}
	case constants.RoleBindingServiceType:
		if err := ConvertRoleBindingToYaml(req, serviceResp); err != nil {
			return nil, err
		}
	case constants.ServiceAccountServiceType:
		if err := ConvertServiceAccountToYaml(req, serviceResp); err != nil {
			return nil, err
		}
	case constants.ClusterRoleServiceType:
		if err := ConvertClusterRoleToYaml(req, serviceResp); err != nil {
			return nil, err
		}
	case constants.ClusterRoleBindingServiceType:
		if err := ConvertClusterRoleBindingToYaml(req, serviceResp); err != nil {
			return nil, err
		}
	case constants.HpaServiceType:
		if err := ConvertHPAToYaml(req, serviceResp); err != nil {
			return nil, err
		}
	case constants.DeploymentServiceType:
		if err := ConvertDeploymentToYaml(req, serviceResp); err != nil {
			return nil, err
		}
	case constants.DaemonSetServiceType:
		if err := ConvertDaemonSeToYaml(req, serviceResp); err != nil {
			return nil, err
		}
	case constants.SecretServiceType:
		if err := ConvertSecretToYaml(req, serviceResp); err != nil {
			return nil, err
		}
	case constants.ConfigMapServiceType:
		if err := ConvertConfigMapToYaml(req, serviceResp); err != nil {
			return nil, err
		}
	case constants.ServiceEntryType:
		if err := ConvertServiceEntryToYaml(req, serviceResp); err != nil {
			return nil, err
		}
	case constants.VirtualServiceType:
		if err := ConvertVSToYaml(req, serviceResp); err != nil {
			return nil, err
		}

	case constants.DestinationRulesType:
		if err := ConvertDRToYaml(req, serviceResp); err != nil {
			return nil, err
		}
	case constants.StatefulSetServiceType:
		if err := ConvertStatefulToYaml(req, serviceResp); err != nil {
			return nil, err
		}
	case constants.JobServiceType:
		if err := ConvertJobToYaml(req, serviceResp); err != nil {
			return nil, err
		}
	case constants.CronJobServiceType:
		if err := ConvertCronJobToYaml(req, serviceResp); err != nil {
			return nil, err
		}
	}

	return serviceResp, nil
}

func ConvertSCToYaml(req *pb.YamlServiceRequest, serviceResp *pb.YamlServiceResponse) error {

	sc := pb.StorageClassService{}
	if err := json.Unmarshal(req.Service, &sc); err != nil {
		utils.Error.Println(err)
		return err
	}
	result, err := getStorageClass(&sc)
	if err != nil {
		utils.Error.Println(err)
		return err
	}

	if byteData, err := yaml.Marshal(result); err != nil {
		utils.Error.Println(err)
		return err
	} else {
		serviceResp.Service = byteData
	}
	return nil
}

func ConvertPVCToYaml(req *pb.YamlServiceRequest, serviceResp *pb.YamlServiceResponse) error {

	pvc := pb.PersistentVolumeClaimService{}
	if err := json.Unmarshal(req.Service, &pvc); err != nil {
		utils.Error.Println(err)
		return err
	}
	result, err := getPersistentVolumeClaim(&pvc)
	if err != nil {
		utils.Error.Println(err)
		return err
	}
	a, b, c, d := helm_parameterization.PersistentVolumeClaimParameters(result)

	fmt.Println(a, b, c, d)

	if byteData, err := yaml.Marshal(result); err != nil {
		utils.Error.Println(err)
		return err
	} else {
		strdata := string(byteData)
		re := regexp.MustCompile("(?m)[\r\n]+^.*creationTimestamp.*$")
		res := re.ReplaceAllString(strdata, "")
		re = regexp.MustCompile("(?m)[\r\n]+^.*status.*$")
		res = re.ReplaceAllString(res, "")
		serviceResp.Service = []byte(res)
		serviceResp.Namespace = result.Namespace
	}
	return nil
}

func ConvertPVToYaml(req *pb.YamlServiceRequest, serviceResp *pb.YamlServiceResponse) error {

	pv := pb.PersistentVolumeService{}
	if err := json.Unmarshal(req.Service, &pv); err != nil {
		utils.Error.Println(err)
		return err
	}
	result, err := getPersistentVolume(&pv)
	if err != nil {
		utils.Error.Println(err)
		return err
	}
	a, b, c, d := helm_parameterization.PersistentVolumeParameters(result)

	fmt.Println(a, b, c, d)

	if byteData, err := yaml.Marshal(result); err != nil {
		utils.Error.Println(err)
		return err
	} else {
		serviceResp.Service = byteData
	}
	return nil
}

func ConvertDRToYaml(req *pb.YamlServiceRequest, serviceResp *pb.YamlServiceResponse) error {

	dr := pb.DestinationRules{}
	if err := json.Unmarshal(req.Service, &dr); err != nil {
		utils.Error.Println(err)
		return err
	}
	result, err := getDestinationRules(&dr)
	if err != nil {
		utils.Error.Println(err)
		return err
	}
	if byteData, err := yaml.Marshal(result); err != nil {
		utils.Error.Println(err)
		return err
	} else {
		serviceResp.Service = byteData
		serviceResp.Namespace = result.Namespace
	}
	return nil
}

func ConvertVSToYaml(req *pb.YamlServiceRequest, serviceResp *pb.YamlServiceResponse) error {

	vs := pb.VirtualService{}
	if err := json.Unmarshal(req.Service, &vs); err != nil {
		utils.Error.Println(err)
		return err
	}
	result, err := getVirtualService(&vs)
	if err != nil {
		utils.Error.Println(err)
		return err
	}
	if byteData, err := yaml.Marshal(result); err != nil {
		utils.Error.Println(err)
		return err
	} else {
		serviceResp.Service = byteData
		serviceResp.Namespace = result.Namespace
	}
	return nil
}

func ConvertGatewayToYaml(req *pb.YamlServiceRequest, serviceResp *pb.YamlServiceResponse) error {

	gateway := pb.GatewayService{}
	if err := json.Unmarshal(req.Service, &gateway); err != nil {
		utils.Error.Println(err)
		return err
	}
	result, err := getIstioGateway(&gateway)
	if err != nil {
		utils.Error.Println(err)
		return err
	}
	if byteData, err := yaml.Marshal(result); err != nil {
		utils.Error.Println(err)
		return err
	} else {
		serviceResp.Service = byteData
		serviceResp.Namespace = result.Namespace
	}
	return nil
}

func ConvertServiceEntryToYaml(req *pb.YamlServiceRequest, serviceResp *pb.YamlServiceResponse) error {

	se := pb.ServiceEntryTemplate{}
	if err := json.Unmarshal(req.Service, &se); err != nil {
		utils.Error.Println(err)
		return err
	}
	result, err := getServiceEntryRequestObject(&se)
	if err != nil {
		utils.Error.Println(err)
		return err
	}
	if byteData, err := yaml.Marshal(result); err != nil {
		utils.Error.Println(err)
		return err
	} else {
		serviceResp.Service = byteData
		serviceResp.Namespace = result.Namespace
	}
	return nil
}

func ConvertConfigMapToYaml(req *pb.YamlServiceRequest, serviceResp *pb.YamlServiceResponse) error {

	cm := pb.ConfigMapService{}
	if err := json.Unmarshal(req.Service, &cm); err != nil {
		utils.Error.Println(err)
		return err
	}
	result, err := getConfigMapService(&cm)
	if err != nil {
		utils.Error.Println(err)
		return err
	}
	a, b, c, d := helm_parameterization.ConfigMapParameters(result)

	fmt.Println(a, b, c, d)
	if byteData, err := yaml.Marshal(result); err != nil {
		utils.Error.Println(err)
		return err
	} else {
		serviceResp.Service = byteData
		serviceResp.Namespace = result.Namespace
	}
	return nil
}

func ConvertSecretToYaml(req *pb.YamlServiceRequest, serviceResp *pb.YamlServiceResponse) error {

	secret := pb.SecretService{}
	if err := json.Unmarshal(req.Service, &secret); err != nil {
		utils.Error.Println(err)
		return err
	}
	result, err := getSecret(&secret)
	if err != nil {
		utils.Error.Println(err)
		return err
	}
	a, b, c, d := helm_parameterization.SecretParameters(result)

	fmt.Println(a, b, c, d)

	if byteData, err := yaml.Marshal(result); err != nil {
		utils.Error.Println(err)
		return err
	} else {
		serviceResp.Service = byteData
		serviceResp.Namespace = result.Namespace
	}
	return nil
}

func ConvertDaemonSeToYaml(req *pb.YamlServiceRequest, serviceResp *pb.YamlServiceResponse) error {

	ds := pb.DaemonSetService{}
	if err := json.Unmarshal(req.Service, &ds); err != nil {
		utils.Error.Println(err)
		return err
	}
	result, err := getDaemonSetRequestObject(&ds)
	if err != nil {
		utils.Error.Println(err)
		return err
	}
	if byteData, err := yaml.Marshal(result); err != nil {
		utils.Error.Println(err)
		return err
	} else {
		serviceResp.Service = byteData
		serviceResp.Namespace = result.Namespace
	}
	return nil
}

func ConvertDeploymentToYaml(req *pb.YamlServiceRequest, serviceResp *pb.YamlServiceResponse) error {

	deploy := pb.DeploymentService{}
	if err := json.Unmarshal(req.Service, &deploy); err != nil {
		utils.Error.Println(err)
		return err
	}
	result, err := getDeploymentRequestObject(&deploy)
	if err != nil {
		utils.Error.Println(err)
		return err
	}
	if byteData, err := yaml.Marshal(result); err != nil {
		utils.Error.Println(err)
		return err
	} else {
		serviceResp.Service = byteData
		serviceResp.Namespace = result.Namespace
	}
	return nil
}
func ConvertHPAToYaml(req *pb.YamlServiceRequest, serviceResp *pb.YamlServiceResponse) error {

	hpa := pb.HPA{}
	if err := json.Unmarshal(req.Service, &hpa); err != nil {
		utils.Error.Println(err)
		return err
	}
	result, err := getHpa(&hpa)
	if err != nil {
		utils.Error.Println(err)
		return err
	}
	if byteData, err := yaml.Marshal(result); err != nil {
		utils.Error.Println(err)
		return err
	} else {
		serviceResp.Service = byteData
		serviceResp.Namespace = result.Namespace
	}
	return nil
}

func ConvertClusterRoleToYaml(req *pb.YamlServiceRequest, serviceResp *pb.YamlServiceResponse) error {

	cr := pb.ClusterRole{}
	if err := json.Unmarshal(req.Service, &cr); err != nil {
		utils.Error.Println(err)
		return err
	}
	result, err := getClusterRole(&cr)
	if err != nil {
		utils.Error.Println(err)
		return err
	}
	if byteData, err := yaml.Marshal(result); err != nil {
		utils.Error.Println(err)
		return err
	} else {
		serviceResp.Service = byteData
	}
	return nil
}

func ConvertRoleToYaml(req *pb.YamlServiceRequest, serviceResp *pb.YamlServiceResponse) error {

	role := pb.RoleService{}
	if err := json.Unmarshal(req.Service, &role); err != nil {
		utils.Error.Println(err)
		return err
	}
	result, err := getRole(&role)
	if err != nil {
		utils.Error.Println(err)
		return err
	}
	if byteData, err := yaml.Marshal(result); err != nil {
		utils.Error.Println(err)
		return err
	} else {
		serviceResp.Service = byteData
		serviceResp.Namespace = result.Namespace
	}
	return nil
}
func ConvertRoleBindingToYaml(req *pb.YamlServiceRequest, serviceResp *pb.YamlServiceResponse) error {
	roleBinding := pb.RoleBindingService{}
	if err := json.Unmarshal(req.Service, &roleBinding); err != nil {
		utils.Error.Println(err)
		return err
	}
	result, err := getRoleBinding(&roleBinding)
	if err != nil {
		utils.Error.Println(err)
		return err
	}
	if byteData, err := yaml.Marshal(result); err != nil {
		utils.Error.Println(err)
		return err
	} else {
		serviceResp.Service = byteData
		serviceResp.Namespace = result.Namespace
	}
	return nil
}
func ConvertClusterRoleBindingToYaml(req *pb.YamlServiceRequest, serviceResp *pb.YamlServiceResponse) error {

	crBinding := pb.ClusterRoleBinding{}
	if err := json.Unmarshal(req.Service, &crBinding); err != nil {
		utils.Error.Println(err)
		return err
	}
	result, err := getClusterRoleBinding(&crBinding)
	if err != nil {
		utils.Error.Println(err)
		return err
	}
	if byteData, err := yaml.Marshal(result); err != nil {
		utils.Error.Println(err)
		return err
	} else {
		serviceResp.Service = byteData
	}
	return nil
}

func ConvertServiceAccountToYaml(req *pb.YamlServiceRequest, serviceResp *pb.YamlServiceResponse) error {

	sa := pb.ServiceAccountService{}
	if err := json.Unmarshal(req.Service, &sa); err != nil {
		utils.Error.Println(err)
		return err
	}
	result, err := getServiceAccount(&sa)
	if err != nil {
		utils.Error.Println(err)
		return err
	}
	if byteData, err := yaml.Marshal(result); err != nil {
		utils.Error.Println(err)
		return err
	} else {
		serviceResp.Service = byteData
		serviceResp.Namespace = result.Namespace
	}
	return nil
}

func ConvertNetworkPolicyToYaml(req *pb.YamlServiceRequest, serviceResp *pb.YamlServiceResponse) error {

	policy := pb.NetworkPolicyService{}
	if err := json.Unmarshal(req.Service, &policy); err != nil {
		utils.Error.Println(err)
		return err
	}
	result, err := getNetworkPolicy(&policy)
	if err != nil {
		utils.Error.Println(err)
		return err
	}
	if byteData, err := yaml.Marshal(result); err != nil {
		utils.Error.Println(err)
		return err
	} else {
		re := regexp.MustCompile("(?m)[\r\n]+^.*status.*$")
		byteData = re.ReplaceAll(byteData, []byte{})
		serviceResp.Service = byteData
		serviceResp.Namespace = result.Namespace
	}
	return nil
}

func ConvertPolicyToYaml(req *pb.YamlServiceRequest, serviceResp *pb.YamlServiceResponse) error {

	policy := pb.PolicyService{}
	if err := json.Unmarshal(req.Service, &policy); err != nil {
		utils.Error.Println(err)
		return err
	}
	result, err := getPolicy(&policy)
	if err != nil {
		utils.Error.Println(err)
		return err
	}
	if byteData, err := yaml.Marshal(result); err != nil {
		utils.Error.Println(err)
		return err
	} else {
		re := regexp.MustCompile("(?m)[\r\n]+^.*status.*$")
		byteData = re.ReplaceAll(byteData, []byte{})
		serviceResp.Service = byteData
		serviceResp.Namespace = result.Namespace
	}
	return nil
}
func ConvertKubernetesServiceToYaml(req *pb.YamlServiceRequest, serviceResp *pb.YamlServiceResponse) error {

	svc := pb.KubernetesService{}
	if err := json.Unmarshal(req.Service, &svc); err != nil {
		utils.Error.Println(err)
		return err
	}
	result, err := getKubernetesService(&svc)
	if err != nil {
		utils.Error.Println(err)
		return err
	}
	a, b, c, d := helm_parameterization.KubernetesServiceParameters(result)

	fmt.Println(a, b, c, d)

	if byteData, err := yaml.Marshal(result); err != nil {
		utils.Error.Println(err)
		return err
	} else {
		re := regexp.MustCompile("(?m)[\r\n]+^.*status.*$")
		byteData = re.ReplaceAll(byteData, []byte{})
		re = regexp.MustCompile("(?m)[\r\n]+^.*loadBalancer: {}*$")
		byteData = re.ReplaceAll(byteData, []byte{})
		//res := re.ReplaceAllString(strbyteData, "")
		serviceResp.Service = byteData
		serviceResp.Namespace = result.Namespace
	}
	return nil
}

// statefulset
func ConvertStatefulToYaml(req *pb.YamlServiceRequest, serviceResp *pb.YamlServiceResponse) error {
	ds := pb.StatefulSetService{}
	if err := json.Unmarshal(req.Service, &ds); err != nil {
		utils.Error.Println(err)
		return err
	}
	result, err := getStatefulSetRequestObject(&ds)
	if err != nil {
		utils.Error.Println(err)
		return err
	}
	if byteData, err := yaml.Marshal(result); err != nil {
		utils.Error.Println(err)
		return err
	} else {
		serviceResp.Service = byteData
		serviceResp.Namespace = result.Namespace
	}
	return nil
}

func ConvertJobToYaml(req *pb.YamlServiceRequest, serviceResp *pb.YamlServiceResponse) error {
	ds := pb.JobService{}
	if err := json.Unmarshal(req.Service, &ds); err != nil {
		utils.Error.Println(err)
		return err
	}
	result, err := getJobRequestObject(&ds)
	if err != nil {
		utils.Error.Println(err)
		return err
	}
	if byteData, err := yaml.Marshal(result); err != nil {
		utils.Error.Println(err)
		return err
	} else {
		serviceResp.Service = byteData
		serviceResp.Namespace = result.Namespace
	}
	return nil
}

func ConvertCronJobToYaml(req *pb.YamlServiceRequest, serviceResp *pb.YamlServiceResponse) error {
	ds := pb.CronJobService{}
	if err := json.Unmarshal(req.Service, &ds); err != nil {
		utils.Error.Println(err)
		return err
	}
	result, err := getCronJobRequestObject(&ds)
	if err != nil {
		utils.Error.Println(err)
		return err
	}
	if byteData, err := yaml.Marshal(result); err != nil {
		utils.Error.Println(err)
		return err
	} else {
		serviceResp.Service = byteData
		serviceResp.Namespace = result.Namespace
	}
	return nil
}
func ConvertVirtualServiceToYaml(req *pb.YamlServiceRequest, serviceResp *pb.YamlServiceResponse) error {
	ds := pb.VirtualService{}
	if err := json.Unmarshal(req.Service, &ds); err != nil {
		utils.Error.Println(err)
		return err
	}
	result, err := getVirtualService(&ds)
	if err != nil {
		utils.Error.Println(err)
		return err
	}
	if byteData, err := yaml.Marshal(result); err != nil {
		utils.Error.Println(err)
		return err
	} else {
		serviceResp.Service = byteData
		serviceResp.Namespace = result.Namespace
	}
	return nil

}
