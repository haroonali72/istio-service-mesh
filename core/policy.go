package core

import (
	pb1 "bitbucket.org/cloudplex-devs/kubernetes-services-deployment/core/proto"
	pb "bitbucket.org/cloudplex-devs/microservices-mesh-engine/core/services/proto"
	"context"
	"encoding/json"
	"google.golang.org/grpc"
	"istio-service-mesh/constants"
	"istio-service-mesh/utils"
	"istio.io/api/authentication/v1alpha1"
	istioClient "istio.io/client-go/pkg/apis/authentication/v1alpha1"
	"strings"
)

func (s *Server) CreatePolicy(ctx context.Context, req *pb.PolicyService) (*pb.ServiceResponse, error) {
	utils.Info.Println(ctx)
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getPolicyRequestObject(req)

	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp, err)
		return serviceResp, err
	}

	conn, err := grpc.DialContext(ctx, constants.K8sEngineGRPCURL, grpc.WithInsecure())
	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp, err)
		return serviceResp, err
	}
	defer conn.Close()

	raw, err := json.Marshal(ksdRequest)
	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp, err)
		return serviceResp, err
	}
	result, err := pb1.NewServiceClient(conn).CreateService(ctx, &pb1.ServiceRequest{
		ProjectId: req.ProjectId,
		Service:   raw,
		CompanyId: req.CompanyId,
		Token:     req.Token,
	})
	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp, err)
		return serviceResp, err
	}
	utils.Info.Println(result.Service)
	serviceResp.Status.Status = "successful"
	serviceResp.Status.StatusIndividual = append(serviceResp.Status.StatusIndividual, "successful")

	return serviceResp, nil

}
func (s *Server) GetPolicy(ctx context.Context, req *pb.PolicyService) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getPolicyRequestObject(req)

	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp, err)
		return serviceResp, err
	}

	conn, err := grpc.DialContext(ctx, constants.K8sEngineGRPCURL, grpc.WithInsecure())
	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp, err)
		return serviceResp, err
	}
	defer conn.Close()

	raw, err := json.Marshal(ksdRequest)
	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp, err)
		return serviceResp, err
	}
	result, err := pb1.NewServiceClient(conn).GetService(ctx, &pb1.ServiceRequest{
		ProjectId: req.ProjectId,
		Service:   raw,
		CompanyId: req.CompanyId,
		Token:     req.Token,
	})
	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp, err)
		return serviceResp, err
	}
	utils.Info.Println(result.Service)
	serviceResp.Status.Status = "successful"
	serviceResp.Status.StatusIndividual = append(serviceResp.Status.StatusIndividual, "successful")

	return serviceResp, nil
}
func (s *Server) DeletePolicy(ctx context.Context, req *pb.PolicyService) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getPolicyRequestObject(req)

	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp, err)
		return serviceResp, err
	}

	conn, err := grpc.DialContext(ctx, constants.K8sEngineGRPCURL, grpc.WithInsecure())
	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp, err)
		return serviceResp, err
	}
	defer conn.Close()

	raw, err := json.Marshal(ksdRequest)
	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp, err)
		return serviceResp, err
	}
	result, err := pb1.NewServiceClient(conn).DeleteService(ctx, &pb1.ServiceRequest{
		ProjectId: req.ProjectId,
		Service:   raw,
		CompanyId: req.CompanyId,
		Token:     req.Token,
	})
	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp, err)
		return serviceResp, err
	}
	utils.Info.Println(result.Service)
	serviceResp.Status.Status = "successful"
	serviceResp.Status.StatusIndividual = append(serviceResp.Status.StatusIndividual, "successful")

	return serviceResp, nil
}
func (s *Server) PatchPolicy(ctx context.Context, req *pb.PolicyService) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getPolicyRequestObject(req)

	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp, err)
		return serviceResp, err
	}

	conn, err := grpc.DialContext(ctx, constants.K8sEngineGRPCURL, grpc.WithInsecure())
	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp, err)
		return serviceResp, err
	}
	defer conn.Close()

	raw, err := json.Marshal(ksdRequest)
	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp, err)
		return serviceResp, err
	}
	result, err := pb1.NewServiceClient(conn).PatchService(ctx, &pb1.ServiceRequest{
		ProjectId: req.ProjectId,
		Service:   raw,
		CompanyId: req.CompanyId,
		Token:     req.Token,
	})
	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp, err)
		return serviceResp, err
	}
	utils.Info.Println(result.Service)
	serviceResp.Status.Status = "successful"
	serviceResp.Status.StatusIndividual = append(serviceResp.Status.StatusIndividual, "successful")

	return serviceResp, nil
}
func (s *Server) PutPolicy(ctx context.Context, req *pb.PolicyService) (*pb.ServiceResponse, error) {
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getPolicyRequestObject(req)

	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp, err)
		return serviceResp, err
	}

	conn, err := grpc.DialContext(ctx, constants.K8sEngineGRPCURL, grpc.WithInsecure())
	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp, err)
		return serviceResp, err
	}
	defer conn.Close()

	raw, err := json.Marshal(ksdRequest)
	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp, err)
		return serviceResp, err
	}
	result, err := pb1.NewServiceClient(conn).PutService(ctx, &pb1.ServiceRequest{
		ProjectId: req.ProjectId,
		Service:   raw,
		CompanyId: req.CompanyId,
		Token:     req.Token,
	})
	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp, err)
		return serviceResp, err
	}
	utils.Info.Println(result.Service)
	serviceResp.Status.Status = "successful"
	serviceResp.Status.StatusIndividual = append(serviceResp.Status.StatusIndividual, "successful")

	return serviceResp, nil
}

func getPolicy(input *pb.PolicyService) (*istioClient.Policy, error) {
	var policyServ = new(istioClient.Policy)
	labels := make(map[string]string)
	labels["app"] = strings.ToLower(input.Name)
	labels["version"] = strings.ToLower(input.Version)
	policyServ.Labels = labels
	policyServ.Kind = "Policy"
	policyServ.APIVersion = "authentication.policy.io/v1alpha1"
	policyServ.Name = input.Name
	policyServ.Namespace = input.Namespace
	for _, t := range input.ServiceAttributes.Target {
		target := v1alpha1.TargetSelector{}
		target.Name = t.Name

		if t.Ports != nil {
			for _, p := range t.Ports {
				port := &v1alpha1.PortSelector{}
				if p.Name != "" {
					port.Port = &v1alpha1.PortSelector_Name{Name: p.Name}
				} else if p.Number != 0 {
					port.Port = &v1alpha1.PortSelector_Number{Number: uint32(p.Number)}
				}
				target.Ports = append(target.Ports, port)
			}
		}
		policyServ.Spec.Targets = append(policyServ.Spec.Targets, &target)
	}

	for _, t := range input.ServiceAttributes.Peers {
		peer := v1alpha1.PeerAuthenticationMethod{}

		if t.Jwt != nil {
			var trigger []*v1alpha1.Jwt_TriggerRule
			for _, t := range t.Jwt.TriggerRules {
				tr := &v1alpha1.Jwt_TriggerRule{}
				for _, e := range t.ExcludedPath {
					if e.Type == "Exact" {
						tr.ExcludedPaths = append(tr.ExcludedPaths, &v1alpha1.StringMatch{
							MatchType: &v1alpha1.StringMatch_Exact{
								Exact: e.Value,
							},
						})
					} else if e.Type == "Prefix" {
						tr.ExcludedPaths = append(tr.ExcludedPaths, &v1alpha1.StringMatch{
							MatchType: &v1alpha1.StringMatch_Prefix{
								Prefix: e.Value,
							},
						})
					} else if e.Type == "Regex" {
						tr.ExcludedPaths = append(tr.ExcludedPaths, &v1alpha1.StringMatch{
							MatchType: &v1alpha1.StringMatch_Regex{
								Regex: e.Value,
							},
						})
					} else if e.Type == "Suffix" {
						tr.ExcludedPaths = append(tr.ExcludedPaths, &v1alpha1.StringMatch{
							MatchType: &v1alpha1.StringMatch_Suffix{
								Suffix: e.Value,
							},
						})
					}
				}
				for _, e := range t.IncludedPath {
					if e.Type == "Exact" {
						tr.IncludedPaths = append(tr.IncludedPaths, &v1alpha1.StringMatch{
							MatchType: &v1alpha1.StringMatch_Exact{
								Exact: e.Value,
							},
						})
					} else if e.Type == "Prefix" {
						tr.IncludedPaths = append(tr.IncludedPaths, &v1alpha1.StringMatch{
							MatchType: &v1alpha1.StringMatch_Prefix{
								Prefix: e.Value,
							},
						})
					} else if e.Type == "Regex" {
						tr.IncludedPaths = append(tr.IncludedPaths, &v1alpha1.StringMatch{
							MatchType: &v1alpha1.StringMatch_Regex{
								Regex: e.Value,
							},
						})
					} else if e.Type == "Suffix" {
						tr.IncludedPaths = append(tr.IncludedPaths, &v1alpha1.StringMatch{
							MatchType: &v1alpha1.StringMatch_Suffix{
								Suffix: e.Value,
							},
						})
					}
				}

				trigger = append(trigger, tr)
			}

			peer.Params = &v1alpha1.PeerAuthenticationMethod_Jwt{
				Jwt: &v1alpha1.Jwt{
					Issuer: t.Jwt.Issuer,
					//	for _,a :=range t.Jwt.Audiences{
					Audiences:    t.Jwt.Audiences,
					Jwks:         t.Jwt.Jwks,
					JwksUri:      t.Jwt.JwksUri,
					JwtHeaders:   t.Jwt.JwtHeader,
					JwtParams:    t.Jwt.JwtParams,
					TriggerRules: trigger,
				},
			}
		} else if t.Mtls != nil {
			peer.Params = &v1alpha1.PeerAuthenticationMethod_Mtls{
				&v1alpha1.MutualTls{
					AllowTls: t.Mtls.AllowTls,
					Mode:     v1alpha1.MutualTls_Mode(int32(t.Mtls.GetMode())),
				},
			}
		}

		policyServ.Spec.Peers = append(policyServ.Spec.Peers, &peer)
	}

	policyServ.Spec.PeerIsOptional = input.ServiceAttributes.PeerIsOptional

	for _, o := range input.ServiceAttributes.Origin {

		origin := v1alpha1.OriginAuthenticationMethod{}
		if origin.Jwt != nil {
			origin.Jwt = &v1alpha1.Jwt{}
			origin.Jwt.Issuer = o.Jwt.Issuer
			for _, a := range o.Jwt.Audiences {
				origin.Jwt.Audiences = append(origin.Jwt.Audiences, a)
			}
			origin.Jwt.JwksUri = o.Jwt.JwksUri
			origin.Jwt.Jwks = o.Jwt.Jwks
			for _, h := range o.Jwt.JwtHeader {
				origin.Jwt.JwtHeaders = append(origin.Jwt.JwtHeaders, h)
			}
			origin.Jwt.JwtParams = o.Jwt.JwtParams
			origin.Jwt.JwtHeaders = o.Jwt.JwtHeader
			for _, t := range o.Jwt.TriggerRules {
				tr := &v1alpha1.Jwt_TriggerRule{}
				for _, e := range t.ExcludedPath {
					if e.Type == "Exact" {
						tr.ExcludedPaths = append(tr.ExcludedPaths, &v1alpha1.StringMatch{
							MatchType: &v1alpha1.StringMatch_Exact{
								Exact: e.Value,
							},
						})
					} else if e.Type == "Prefix" {
						tr.ExcludedPaths = append(tr.ExcludedPaths, &v1alpha1.StringMatch{
							MatchType: &v1alpha1.StringMatch_Prefix{
								Prefix: e.Value,
							},
						})
					} else if e.Type == "Regex" {
						tr.ExcludedPaths = append(tr.ExcludedPaths, &v1alpha1.StringMatch{
							MatchType: &v1alpha1.StringMatch_Regex{
								Regex: e.Value,
							},
						})
					} else if e.Type == "Suffix" {
						tr.ExcludedPaths = append(tr.ExcludedPaths, &v1alpha1.StringMatch{
							MatchType: &v1alpha1.StringMatch_Suffix{
								Suffix: e.Value,
							},
						})
					}
				}
				for _, e := range t.IncludedPath {
					if e.Type == "Exact" {
						tr.IncludedPaths = append(tr.IncludedPaths, &v1alpha1.StringMatch{
							MatchType: &v1alpha1.StringMatch_Exact{
								Exact: e.Value,
							},
						})
					} else if e.Type == "Prefix" {
						tr.IncludedPaths = append(tr.IncludedPaths, &v1alpha1.StringMatch{
							MatchType: &v1alpha1.StringMatch_Prefix{
								Prefix: e.Value,
							},
						})
					} else if e.Type == "Regex" {
						tr.IncludedPaths = append(tr.IncludedPaths, &v1alpha1.StringMatch{
							MatchType: &v1alpha1.StringMatch_Regex{
								Regex: e.Value,
							},
						})
					} else if e.Type == "Suffix" {
						tr.IncludedPaths = append(tr.IncludedPaths, &v1alpha1.StringMatch{
							MatchType: &v1alpha1.StringMatch_Suffix{
								Suffix: e.Value,
							},
						})
					}
				}
				origin.Jwt.TriggerRules = append(origin.Jwt.TriggerRules, tr)
			}
			policyServ.Spec.Origins = append(policyServ.Spec.Origins, &origin)
		}
	}

	policyServ.Spec.OriginIsOptional = input.ServiceAttributes.OriginIsOptional

	policyServ.Spec.PrincipalBinding = v1alpha1.PrincipalBinding(int32(input.ServiceAttributes.PrincipalBinding))

	return policyServ, nil
}

func getPolicyRequestObject(req *pb.PolicyService) (*istioClient.Policy, error) {
	gtwReq, err := getPolicy(req)
	if err != nil {
		utils.Error.Println(err)

		return nil, err
	}
	return gtwReq, nil
}
