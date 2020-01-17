package core

import (
	"context"
	"encoding/json"
	"errors"
	"golang.org/x/build/kubernetes/api"
	"google.golang.org/grpc"
	"istio-service-mesh/constants"
	pb "istio-service-mesh/core/proto"
	"istio-service-mesh/utils"
	v1 "k8s.io/api/apps/v1"
	v2 "k8s.io/api/core/v1"
	rbac "k8s.io/api/rbac/v1"
	"reflect"
)

func (s *Server) GetK8SResource(ctx context.Context, request *pb.K8SResourceRequest) (response *pb.K8SResourceResponse, err error) {
	response = new(pb.K8SResourceResponse)
	utils.Info.Println(reflect.TypeOf(ctx))

	if request.CompanyId == "" || request.ProjectId == "" {
		return &pb.K8SResourceResponse{}, errors.New("projectId or companyId must not be empty")
	}

	conn, err := grpc.DialContext(ctx, constants.K8sEngineGRPCURL, grpc.WithInsecure())
	if err != nil {
		utils.Error.Println(err)
		return &pb.K8SResourceResponse{}, err
	}
	defer conn.Close()

	response, err = pb.NewK8SResourceClient(conn).GetK8SResource(ctx, request)
	if err != nil {
		utils.Error.Println(err)
		return &pb.K8SResourceResponse{}, err
	}

	//if request.Name == "" {
	//	var dep []*v1.Deployment
	//	err = json.Unmarshal(response.Resource, dep)
	//	if err != nil {
	//		utils.Error.Println(err)
	//		return &pb.K8SResourceResponse{}, err
	//	}
	//} else {
	//	var dep *v1.Deployment
	//	err = json.Unmarshal(response.Resource, dep)
	//	if err != nil {
	//		utils.Error.Println(err)
	//		return &pb.K8SResourceResponse{}, err
	//	}
	//}

	var dep []*v1.Deployment
	err = json.Unmarshal(response.Resource, dep)
	if err != nil {
		utils.Error.Println(err)
		return &pb.K8SResourceResponse{}, err
	}

	return response, err
}

func getClusterRoleDependentInfo(ctx context.Context, conn *grpc.ClientConn, req *pb.K8SResourceRequest, deployments []*v1.Deployment) {
	projectId := req.ProjectId
	companyId := req.CompanyId
	token := req.Token
	for _, dep := range deployments {

		//checking for the service account if name not empty then getting cluster role and cluster role  binding against that service account
		if dep.Spec.Template.Spec.ServiceAccountName != "" {

			response, err := pb.NewK8SResourceClient(conn).GetK8SResource(ctx, &pb.K8SResourceRequest{
				ProjectId: projectId,
				CompanyId: companyId,
				Token:     token,
				Command:   "kubectl",
				Args:      []string{"get", "sa", dep.Spec.Template.Spec.ServiceAccountName, "-n", dep.Namespace, "-o", "json"},
			})
			if err != nil {
				utils.Error.Println(err)
			}

			var svcAcc *api.ServiceAccount
			err = json.Unmarshal(response.Resource, &svcAcc)
			if err != nil {
				utils.Error.Println(err)
			}

			//creating secrets for service account
			for _, secret := range svcAcc.Secrets {
				if secret.Name != "" {
					response, err = pb.NewK8SResourceClient(conn).GetK8SResource(ctx, &pb.K8SResourceRequest{
						ProjectId: projectId,
						CompanyId: companyId,
						Token:     token,
						Command:   "kubectl",
						Args:      []string{"get", "secrets", secret.Name},
					})
					if err != nil {
						utils.Error.Println(err)
					}

					var scrt []*v2.Secret
					err = json.Unmarshal(response.Resource, &scrt)
					if err != nil {
						utils.Error.Println(err)
					}
				}
			}

			response, err = pb.NewK8SResourceClient(conn).GetK8SResource(ctx, &pb.K8SResourceRequest{
				ProjectId: projectId,
				CompanyId: companyId,
				Token:     token,
				Command:   "kubectl",
				Args:      []string{"get", "clusterrolebinding"},
			})
			if err != nil {
				utils.Error.Println(err)
			}

			var clusterrolebindings []*rbac.ClusterRoleBinding
			err = json.Unmarshal(response.Resource, &clusterrolebindings)
			if err != nil {
				utils.Error.Println(err)
			}

			for _, clstrrolebind := range clusterrolebindings {
				for _, sub := range clstrrolebind.Subjects {
					if sub.Kind == "ServiceAccount" && sub.Name == dep.Spec.Template.Spec.ServiceAccountName {
						if clstrrolebind.RoleRef.Kind == "ClusterRole" {
							clusterrolename := clstrrolebind.RoleRef.Name
							response, err := pb.NewK8SResourceClient(conn).GetK8SResource(ctx, &pb.K8SResourceRequest{
								ProjectId: projectId,
								CompanyId: companyId,
								Token:     token,
								Command:   "kubectl",
								Args:      []string{"get", "clusterrole", clusterrolename},
							})
							if err != nil {
								utils.Error.Println(err)
							}

							var clusterrole *rbac.ClusterRole
							err = json.Unmarshal(response.Resource, &clusterrole)
							if err != nil {
								utils.Error.Println(err)
							}
						}
						break
					}
				}
			}
		}

		for _, container := range dep.Spec.Template.Spec.Containers {
			//discovering secret and config maps in deployment containers
			for _, env := range container.Env {
				if env.ValueFrom.SecretKeyRef.Name != "" {

					response, err := pb.NewK8SResourceClient(conn).GetK8SResource(ctx, &pb.K8SResourceRequest{
						ProjectId: projectId,
						CompanyId: companyId,
						Token:     token,
						Command:   "kubectl",
						Args:      []string{"get", "secrets", env.ValueFrom.SecretKeyRef.Name, "-o", "json"},
					})
					if err != nil {
						utils.Error.Println(err)
					}

					var secret *v2.Secret
					err = json.Unmarshal(response.Resource, &secret)
					if err != nil {
						utils.Error.Println(err)
					}

				} else if env.ValueFrom.ConfigMapKeyRef.Name != "" {
					response, err := pb.NewK8SResourceClient(conn).GetK8SResource(ctx, &pb.K8SResourceRequest{
						ProjectId: projectId,
						CompanyId: companyId,
						Token:     token,
						Command:   "kubectl",
						Args:      []string{"get", "configmaps", env.ValueFrom.ConfigMapKeyRef.Name, "-o", "json"},
					})
					if err != nil {
						utils.Error.Println(err)
					}

					var configmap *v2.ConfigMap
					err = json.Unmarshal(response.Resource, &configmap)
					if err != nil {
						utils.Error.Println(err)
					}
				} //else if env.ValueFrom.ResourceFieldRef.
			}
		}
	}
}
