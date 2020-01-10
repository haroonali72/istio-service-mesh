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
		if err := json.Unmarshal(req.Service, &networkproto); err == nil {
			result, err := getStorageClass(&networkproto)
			if err != nil {
				utils.Error.Println(err)
			}
			if byteData, err := yaml.Marshal(result); err == nil {
				serviceResp.Service = byteData
			}
		} else {
			utils.Error.Println(err)
		}

	}
	return serviceResp, nil
}
