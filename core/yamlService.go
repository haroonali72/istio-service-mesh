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
		networkproto := pb.StorageClassService{}
		if err := json.Unmarshal(req.Service, &networkproto); err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		result, err := getStorageClass(&networkproto)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		if byteData, err := yaml.Marshal(result); err != nil {
			utils.Error.Println(err)
			return nil, err
		} else {
			serviceResp.Service = byteData
		}

	case "PVC":
		networkproto := pb.PersistentVolumeClaimService{}
		if err := json.Unmarshal(req.Service, &networkproto); err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		result, err := getPersistentVolumeClaim(&networkproto)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		if byteData, err := yaml.Marshal(result); err != nil {
			utils.Error.Println(err)
			return nil, err
		} else {
			serviceResp.Service = byteData
		}

	case "PV":
		networkproto := pb.PersistentVolumeService{}
		if err := json.Unmarshal(req.Service, &networkproto); err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		result, err := getPersistentVolume(&networkproto)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		if byteData, err := yaml.Marshal(result); err != nil {
			utils.Error.Println(err)
			return nil, err
		} else {
			serviceResp.Service = byteData
		}
	case "gateway":
		networkproto := pb.GatewayService{}
		if err := json.Unmarshal(req.Service, &networkproto); err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		result, err := getIstioGateway(&networkproto)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		if byteData, err := yaml.Marshal(result); err != nil {
			utils.Error.Println(err)
			return nil, err
		} else {
			serviceResp.Service = byteData
		}
	case "kubernetesservice":
		networkproto := pb.KubernetesService{}
		if err := json.Unmarshal(req.Service, &networkproto); err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		result, err := getKubernetesService(&networkproto)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		if byteData, err := yaml.Marshal(result); err != nil {
			utils.Error.Println(err)
			return nil, err
		} else {
			serviceResp.Service = byteData
		}
	case "networkPolicy":
		networkproto := pb.NetworkPolicyService{}
		if err := json.Unmarshal(req.Service, &networkproto); err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		result, err := getNetworkPolicy(&networkproto)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		if byteData, err := yaml.Marshal(result); err != nil {
			utils.Error.Println(err)
			return nil, err
		} else {
			serviceResp.Service = byteData
		}
	case "role":
		networkproto := pb.RoleService{}
		if err := json.Unmarshal(req.Service, &networkproto); err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		result, err := getRole(&networkproto)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		if byteData, err := yaml.Marshal(result); err != nil {
			utils.Error.Println(err)
			return nil, err
		} else {
			serviceResp.Service = byteData
		}
	case "rolebinding":
		networkproto := pb.RoleBindingService{}
		if err := json.Unmarshal(req.Service, &networkproto); err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		result, err := getRoleBinding(&networkproto)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		if byteData, err := yaml.Marshal(result); err != nil {
			utils.Error.Println(err)
			return nil, err
		} else {
			serviceResp.Service = byteData
		}
	case "serviceaccount":
		networkproto := pb.ServiceAccountService{}
		if err := json.Unmarshal(req.Service, &networkproto); err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		result, err := getServiceAccount(&networkproto)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		if byteData, err := yaml.Marshal(result); err != nil {
			utils.Error.Println(err)
			return nil, err
		} else {
			serviceResp.Service = byteData
		}
	case "cluster_role":
		networkproto := pb.ClusterRole{}
		if err := json.Unmarshal(req.Service, &networkproto); err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		result, err := getClusterRole(&networkproto)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		if byteData, err := yaml.Marshal(result); err != nil {
			utils.Error.Println(err)
			return nil, err
		} else {
			serviceResp.Service = byteData
		}
	case "cluster_role_binding":
		networkproto := pb.ClusterRoleBinding{}
		if err := json.Unmarshal(req.Service, &networkproto); err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		result, err := getClusterRoleBinding(&networkproto)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		if byteData, err := yaml.Marshal(result); err != nil {
			utils.Error.Println(err)
			return nil, err
		} else {
			serviceResp.Service = byteData
		}
	case "hpa":
		networkproto := pb.HPA{}
		if err := json.Unmarshal(req.Service, &networkproto); err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		result, err := getHpa(&networkproto)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		if byteData, err := yaml.Marshal(result); err != nil {
			utils.Error.Println(err)
			return nil, err
		} else {
			serviceResp.Service = byteData
		}
	case "Deployment":
		networkproto := pb.DeploymentService{}
		if err := json.Unmarshal(req.Service, &networkproto); err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		result, err := getDeploymentRequestObject(&networkproto)
		if err != nil {
			utils.Error.Println(err)
			return nil, err
		}
		if byteData, err := yaml.Marshal(result); err != nil {
			utils.Error.Println(err)
			return nil, err
		} else {
			serviceResp.Service = byteData
		}
	}

	return serviceResp, nil
}
