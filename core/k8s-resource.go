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

type GrpcConn struct {
	Connection *grpc.ClientConn
	ProjectId  string
	CompanyId  string
	token      string
}

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

	grpcconn := &GrpcConn{
		Connection: conn,
		ProjectId:  request.ProjectId,
		CompanyId:  request.CompanyId,
		token:      request.Token,
	}

	//response, err = pb.NewK8SResourceClient(conn).GetK8SResource(ctx, request)
	//if err != nil {
	//	utils.Error.Println(err)
	//	return &pb.K8SResourceResponse{}, err
	//}

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

	//var dep []*v1.Deployment
	//err = json.Unmarshal(response.Resource, dep)
	//if err != nil {
	//	utils.Error.Println(err)
	//	return &pb.K8SResourceResponse{}, err
	//}

	return response, err
}

func (conn *GrpcConn) deploymentk8sToCp(ctx context.Context, req *pb.K8SResourceRequest, deployments []*v1.Deployment) {

	for _, dep := range deployments {

		namespace := dep.Namespace
		//checking for the service account if name not empty then getting cluster role and cluster role  binding against that service account
		if dep.Spec.Template.Spec.ServiceAccountName != "" {

			svcname := dep.Spec.Template.Spec.ServiceAccountName
			svcaccount, err := conn.getsvcaccount(ctx, svcname, namespace)
			if err != nil {
				return
			}

			//creating secrets for service account
			for _, secret := range svcaccount.Secrets {
				if secret.Name != "" {

					secretname := secret.Name
					if secret.Namespace != "" {
						namespace = secret.Namespace
					}
					secretResp, err := conn.getsecret(ctx, secretname, namespace)
					if err != nil {
						return
					}
				}
			}

			clusterrolebindings, err := conn.getclusterroelbindings(ctx)
			if err != nil {
				return
			}

			for _, clstrrolebind := range clusterrolebindings {
				for _, sub := range clstrrolebind.Subjects {
					if sub.Kind == "ServiceAccount" && sub.Name == svcname {
						if clstrrolebind.RoleRef.Kind == "ClusterRole" {
							clusterrolename := clstrrolebind.RoleRef.Name
							resp, err := conn.getclusterrole(ctx, clusterrolename)
							if err != nil {
								return
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
					secretname := env.ValueFrom.SecretKeyRef.Name
					resp, err := conn.getsecret(ctx, secretname, namespace)
					if err != nil {
						return
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

func (conn *GrpcConn) getsvcaccount(ctx context.Context, svcname, namespace string) (*api.ServiceAccount, error) {
	response, err := pb.NewK8SResourceClient(conn.Connection).GetK8SResource(ctx, &pb.K8SResourceRequest{
		ProjectId: conn.ProjectId,
		CompanyId: conn.CompanyId,
		Token:     conn.token,
		Command:   "kubectl",
		Args:      []string{"get", "sa", svcname, "-n", namespace, "-o", "json"},
	})
	if err != nil {
		utils.Error.Println(err)
		return nil, errors.New("error from grpc server :" + err.Error())
	}

	var svcAcc *api.ServiceAccount
	err = json.Unmarshal(response.Resource, &svcAcc)
	if err != nil {
		utils.Error.Println(err)
		return nil, err
	}

	return svcAcc, nil
}

func (conn *GrpcConn) getsecret(ctx context.Context, secretname, namespace string) (*v2.Secret, error) {
	response, err := pb.NewK8SResourceClient(conn.Connection).GetK8SResource(ctx, &pb.K8SResourceRequest{
		ProjectId: conn.ProjectId,
		CompanyId: conn.CompanyId,
		Token:     conn.token,
		Command:   "kubectl",
		Args:      []string{"get", "secrets", secretname, "-n", namespace, "-o", "json"},
	})
	if err != nil {
		utils.Error.Println(err)
		return nil, errors.New("error from grpc server :" + err.Error())
	}

	var scrt *v2.Secret
	err = json.Unmarshal(response.Resource, &scrt)
	if err != nil {
		utils.Error.Println(err)
		return nil, err
	}
	return scrt, nil
}

func (conn *GrpcConn) getclusterroelbindings(ctx context.Context) ([]*rbac.ClusterRoleBinding, error) {
	response, err := pb.NewK8SResourceClient(conn.Connection).GetK8SResource(ctx, &pb.K8SResourceRequest{
		ProjectId: conn.ProjectId,
		CompanyId: conn.CompanyId,
		Token:     conn.token,
		Command:   "kubectl",
		Args:      []string{"get", "clusterrolebindings"},
	})
	if err != nil {
		utils.Error.Println(err)
		return nil, errors.New("error from grpc server :" + err.Error())
	}

	var clusterrolebindings []*rbac.ClusterRoleBinding
	err = json.Unmarshal(response.Resource, &clusterrolebindings)
	if err != nil {
		utils.Error.Println(err)
		return nil, err
	}

	return clusterrolebindings, nil
}

func (conn *GrpcConn) getclusterrole(ctx context.Context, clusterrolename string) (*rbac.ClusterRole, error) {
	response, err := pb.NewK8SResourceClient(conn.Connection).GetK8SResource(ctx, &pb.K8SResourceRequest{
		ProjectId: conn.ProjectId,
		CompanyId: conn.CompanyId,
		Token:     conn.token,
		Command:   "kubectl",
		Args:      []string{"get", "clusterrole", clusterrolename},
	})
	if err != nil {
		utils.Error.Println(err)
		return nil, errors.New("error from grpc server :" + err.Error())
	}

	var clusterrole *rbac.ClusterRole
	err = json.Unmarshal(response.Resource, &clusterrole)
	if err != nil {
		utils.Error.Println(err)
		return nil, err
	}
	return clusterrole, nil
}
