package core

import (
	"context"
	"encoding/json"
	pb "istio-service-mesh/core/proto"
	"istio-service-mesh/utils"
	"sigs.k8s.io/yaml"
)

func (s *Server) GetYamlService(ctx context.Context, req *pb.YamlServiceRequest) (*pb.YamlServiceResponse, error) {
	serviceResp := new(pb.YamlServiceResponse)
	switch req.Type {
	case "SC":
		if err := ConvertSCToYaml(req, serviceResp); err != nil {
			return nil, err
		}

	case "PVC":
		if err := ConvertPVCToYaml(req, serviceResp); err != nil {
			return nil, err
		}
	case "PV":
		if err := ConvertPVToYaml(req, serviceResp); err != nil {
			return nil, err
		}
	case "gateway":
		if err := ConvertGatewayToYaml(req, serviceResp); err != nil {
			return nil, err
		}
	case "kubernetesservice":
		if err := ConvertKubernetesServiceToYaml(req, serviceResp); err != nil {
			return nil, err
		}
	case "networkPolicy":
		if err := ConvertNetworkPolicyToYaml(req, serviceResp); err != nil {
			return nil, err
		}
	case "role":
		if err := ConvertRoleToYaml(req, serviceResp); err != nil {
			return nil, err
		}
	case "policy":
		if err := ConvertPolicyToYaml(req, serviceResp); err != nil {
			return nil, err
		}
	case "rolebinding":
		if err := ConvertRoleBindingToYaml(req, serviceResp); err != nil {
			return nil, err
		}
	case "serviceaccount":
		if err := ConvertServiceAccountToYaml(req, serviceResp); err != nil {
			return nil, err
		}
	case "cluster_role":
		if err := ConvertClusterRoleToYaml(req, serviceResp); err != nil {
			return nil, err
		}
	case "cluster_role_binding":
		if err := ConvertClusterRoleBindingToYaml(req, serviceResp); err != nil {
			return nil, err
		}
	case "hpa":
		if err := ConvertHPAToYaml(req, serviceResp); err != nil {
			return nil, err
		}
	case "deployment":
		if err := ConvertDeploymentToYaml(req, serviceResp); err != nil {
			return nil, err
		}
	case "DaemonSet":
		if err := ConvertDaemonSeToYaml(req, serviceResp); err != nil {
			return nil, err
		}
	case "secret":
		if err := ConvertSecretToYaml(req, serviceResp); err != nil {
			return nil, err
		}
	case "configmap":
		if err := ConvertConfigMapToYaml(req, serviceResp); err != nil {
			return nil, err
		}
	case "service_entry":
		if err := ConvertServiceEntryToYaml(req, serviceResp); err != nil {
			return nil, err
		}
	case "virtual_service":
		if err := ConvertVSToYaml(req, serviceResp); err != nil {
			return nil, err
		}

	case "destination_rule":
		if err := ConvertDRToYaml(req, serviceResp); err != nil {
			return nil, err
		}
	case "StatefulSet":
		if err := ConvertStatefulToYaml(req, serviceResp); err != nil {
			return nil, err
		}
	case "Job":
		if err := ConvertJobToYaml(req, serviceResp); err != nil {
			return nil, err
		}
	case "CronJob":
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
	if byteData, err := yaml.Marshal(result); err != nil {
		utils.Error.Println(err)
		return err
	} else {
		serviceResp.Service = byteData
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
	if byteData, err := yaml.Marshal(result); err != nil {
		utils.Error.Println(err)
		return err
	} else {
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
