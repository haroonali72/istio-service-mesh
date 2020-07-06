package core

import (
	pb "bitbucket.org/cloudplex-devs/microservices-mesh-engine/core/services/proto"
	"context"
)

//func converToResp(resp *pb.ServiceResponse, projectId string, responseStatusCode int, responseBody []byte) {
//
//	if responseStatusCode == 200 || responseStatusCode == 201 {
//		resp.Status.Status = "successful"
//		resp.Status.StatusIndividual = append(resp.Status.StatusIndividual, "successful")
//	} else {
//		var finalErr struct {
//			Error string `json:"error"`
//		}
//		resp.Status.StatusIndividual = append(resp.Status.StatusIndividual, "failed")
//		resp.Status.Status = "failed"
//		err := json.Unmarshal(responseBody, &finalErr)
//		if err != nil {
//			resp.Status.Reason = err.Error()
//			return
//		}
//		resp.Status.Reason = finalErr.Error
//	}
//	return
//}

func getHeaders(ctx context.Context, projectId string) map[string]string {
	return map[string]string{
		"token":        ctx.Value("token").(string),
		"project_id":   projectId,
		"Content-Type": "application/json",
	}
}
func getErrorResp(resp *pb.ServiceResponse, err error) {
	resp.Error = err.Error()
	resp.Status.Status = "failed"
	resp.Status.StatusIndividual = append(resp.Status.StatusIndividual, "failed")
	resp.Status.Reason = err.Error()
}

func getPodErrors(resp *pb.ServiceResponse, podErrors []string) {
	resp.Error = "Pod is in the pending state"
	resp.Status.Status = "failed"
	resp.Status.StatusIndividual = append(resp.Status.StatusIndividual, "failed")
	resp.Status.PodErrors = podErrors
}
