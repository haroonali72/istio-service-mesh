package core

import (
	"context"
	"encoding/json"
	"errors"
	"google.golang.org/grpc"
	"istio-service-mesh/constants"
	pb "istio-service-mesh/core/proto"
	"istio-service-mesh/types"
	"istio-service-mesh/utils"
	"k8s.io/api/apps/v1"
	v2 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"strconv"
	"strings"
)

func (s *Server) CreateDeployment(ctx context.Context, req *pb.DeploymentService) (*pb.ServiceResponse, error) {
	utils.Info.Println(ctx)
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getDeploymentRequestObject(req)
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
	result, err := pb.NewServiceClient(conn).CreateService(ctx, &pb.ServiceRequest{
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

func (s *Server) GetDeployment(ctx context.Context, req *pb.DeploymentService) (*pb.ServiceResponse, error) {
	utils.Info.Println(ctx)
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getDeploymentRequestObject(req)
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
	result, err := pb.NewServiceClient(conn).GetService(ctx, &pb.ServiceRequest{
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

func (s *Server) DeleteDeployment(ctx context.Context, req *pb.DeploymentService) (*pb.ServiceResponse, error) {
	utils.Info.Println(ctx)
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getDeploymentRequestObject(req)
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
	result, err := pb.NewServiceClient(conn).DeleteService(ctx, &pb.ServiceRequest{
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

func (s *Server) PatchDeployment(ctx context.Context, req *pb.DeploymentService) (*pb.ServiceResponse, error) {
	utils.Info.Println(ctx)
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getDeploymentRequestObject(req)
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
	result, err := pb.NewServiceClient(conn).PatchService(ctx, &pb.ServiceRequest{
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

func (s *Server) PutDeployment(ctx context.Context, req *pb.DeploymentService) (*pb.ServiceResponse, error) {
	utils.Info.Println(ctx)
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id:        req.ServiceId,
		ServiceId: req.ServiceId,
		Name:      req.Name,
	}
	ksdRequest, err := getDeploymentRequestObject(req)
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
	result, err := pb.NewServiceClient(conn).PutService(ctx, &pb.ServiceRequest{
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

func getDeploymentRequestObject(service *pb.DeploymentService) (*v1.Deployment, error) {
	var secrets, configMaps []string
	var deployment = new(v1.Deployment)
	if service.Name == "" {
		return &v1.Deployment{}, errors.New("Service name not found")
	}
	if service.Namespace == "" {
		deployment.ObjectMeta.Namespace = "default"
	} else {
		deployment.ObjectMeta.Namespace = service.Namespace
	}
	deployment.Name = service.Name
	deployment.TypeMeta.Kind = "Deployment"
	deployment.TypeMeta.APIVersion = "apps/v1"

	deployment.Labels = make(map[string]string)
	deployment.Labels["keel.sh/policy"] = "force"
	deployment.Labels["version"] = service.Version

	deployment.Labels = service.ServiceAttributes.Labels
	deployment.Annotations = make(map[string]string)
	deployment.Annotations = service.ServiceAttributes.Annotations
	deployment.Spec.Selector = new(metav1.LabelSelector)
	deployment.Spec.Selector.MatchLabels = make(map[string]string)
	deployment.Spec.Selector.MatchLabels = service.ServiceAttributes.Labels

	deployment.Spec.Selector.MatchLabels["app"] = service.Name
	deployment.Spec.Selector.MatchLabels["version"] = service.Version

	deployment.Spec.Template.Labels = make(map[string]string)

	deployment.Spec.Template.Labels["app"] = service.Name
	deployment.Spec.Template.Labels["version"] = service.Version

	deployment.Spec.Template.Annotations = make(map[string]string)
	deployment.Spec.Template.Annotations["sidecar.istio.io/inject"] = "true"
	deployment.Spec.Template.Spec.NodeSelector = make(map[string]string)
	deployment.Spec.Template.Spec.NodeSelector = service.ServiceAttributes.NodeSelector

	deployment.Spec.Selector = &metav1.LabelSelector{make(map[string]string), nil}
	deployment.Spec.Selector.MatchLabels = service.ServiceAttributes.Labels
	deployment.Spec.Template.ObjectMeta.Labels = service.ServiceAttributes.Labels

	var err error
	deployment.Spec.Template.Spec.Containers, secrets, configMaps, err = getContainers(service)
	if err != nil {
		return &v1.Deployment{}, err
	}

	isExistSecret := make(map[string]bool)
	isExistConfigMap := make(map[string]bool)

	for _, every := range secrets {
		isExistSecret[every] = true
		deployment.Spec.Template.Spec.Volumes = append(deployment.Spec.Template.Spec.Volumes, v2.Volume{
			Name: every,
			VolumeSource: v2.VolumeSource{
				Secret: &v2.SecretVolumeSource{
					SecretName: every,
				},
			},
		})
	}

	for _, every := range configMaps {
		isExistConfigMap[every] = true
		deployment.Spec.Template.Spec.Volumes = append(deployment.Spec.Template.Spec.Volumes, v2.Volume{
			Name: every,
			VolumeSource: v2.VolumeSource{
				ConfigMap: &v2.ConfigMapVolumeSource{
					LocalObjectReference: v2.LocalObjectReference{
						Name: every,
					},
				},
			},
		})
	}

	deployment.Spec.Template.Spec.InitContainers, secrets, configMaps, err = getInitContainers(service)
	if err != nil {
		return &v1.Deployment{}, err
	}

	for _, every := range secrets {
		if _, ok := isExistSecret[every]; !ok {
			isExistSecret[every] = true
			deployment.Spec.Template.Spec.Volumes = append(deployment.Spec.Template.Spec.Volumes, v2.Volume{
				Name: every,
				VolumeSource: v2.VolumeSource{
					Secret: &v2.SecretVolumeSource{
						SecretName: every,
					},
				},
			})
		}
	}

	for _, every := range configMaps {
		if _, ok := isExistConfigMap[every]; !ok {
			isExistConfigMap[every] = true
			deployment.Spec.Template.Spec.Volumes = append(deployment.Spec.Template.Spec.Volumes, v2.Volume{
				Name: every,
				VolumeSource: v2.VolumeSource{
					ConfigMap: &v2.ConfigMapVolumeSource{
						LocalObjectReference: v2.LocalObjectReference{
							Name: every,
						},
					},
				},
			})
		}
	}

	return deployment, nil
}

func getContainers(service *pb.DeploymentService) ([]v2.Container, []string, []string, error) {

	var configMapsArray, secretsArray []string
	var container v2.Container
	container.Name = service.Name
	if err := putCommandAndArguments(&container, service.ServiceAttributes.Command, service.ServiceAttributes.Args); err != nil {
		return nil, secretsArray, configMapsArray, err
	}

	errr, isOk := checkRequestIsLessThanLimit(service.ServiceAttributes)
	if errr != nil {
		return nil, secretsArray, configMapsArray, errr
	} else if isOk == false {
		return nil, secretsArray, configMapsArray, errors.New("Request Resource is greater limit resource")

	}

	if err := putReadinessProbe(&container, service.ServiceAttributes.ReadinessProbe); err != nil {
		return nil, secretsArray, configMapsArray, err
	}

	if err := putLivenessProbe(&container, service.ServiceAttributes.LivenessProbe); err != nil {
		return nil, secretsArray, configMapsArray, err
	}
	if err := putLimitResource(&container, service.ServiceAttributes.LimitResources); err != nil {
		return nil, secretsArray, configMapsArray, err
	}
	if err := putRequestResource(&container, service.ServiceAttributes.RequestResources); err != nil {
		return nil, secretsArray, configMapsArray, err
	}

	if service.ServiceAttributes.SecurityContext != nil {
		if securityContext, err := configureSecurityContext(service.ServiceAttributes.SecurityContext); err != nil {
			return nil, secretsArray, configMapsArray, err
		} else {
			container.SecurityContext = securityContext
		}
	}

	container.Image = service.ServiceAttributes.ImagePrefix + service.ServiceAttributes.ImageName
	if service.ServiceAttributes.Tag != "" {
		container.Image += ":" + service.ServiceAttributes.Tag
	}

	var ports []v2.ContainerPort
	for _, port := range service.ServiceAttributes.Ports {
		temp := v2.ContainerPort{}
		if port.Container == "" && port.Host == "" {
			continue
		}
		if port.Container == "" && port.Host != "" {
			port.Container = port.Host
		}

		i, err := strconv.Atoi(port.Container)
		if err != nil {
			utils.Info.Println(err)
			continue
		}
		if i > 0 && i < 65536 {
			temp.ContainerPort = int32(i)
		} else {
			utils.Info.Println("invalid prot number")
			continue
		}
		if port.Host != "" {
			i, err = strconv.Atoi(port.Host)
			if err != nil {
				utils.Info.Println(err)
				continue
			}
			if i > 0 && i < 65536 {
				temp.HostPort = int32(i)
			} else {
				utils.Info.Println("invalid prot number")
				continue
			}

		}
		ports = append(ports, temp)
	}
	var envVariables []v2.EnvVar
	for key, envVariable := range service.ServiceAttributes.EnvironmentVariables {
		tempEnvVariable := v2.EnvVar{}
		if strings.EqualFold(key, "ConfigMap") {
			envVariableValue := strings.Split(envVariable.Value, ";")
			tempEnvVariable = v2.EnvVar{Name: envVariable.Key,
				ValueFrom: &v2.EnvVarSource{ConfigMapKeyRef: &v2.ConfigMapKeySelector{
					LocalObjectReference: v2.LocalObjectReference{Name: envVariableValue[0]},
					Key:                  envVariableValue[1],
				}}}

		} else if strings.EqualFold(key, "Secret") {
			envVariableValue := strings.Split(envVariable.Value, ";")
			tempEnvVariable = v2.EnvVar{Name: envVariable.Key,
				ValueFrom: &v2.EnvVarSource{SecretKeyRef: &v2.SecretKeySelector{
					LocalObjectReference: v2.LocalObjectReference{Name: envVariableValue[0]},
					Key:                  envVariableValue[1],
				}}}
		} else {
			tempEnvVariable = v2.EnvVar{Name: envVariable.Key, Value: envVariable.Value}
		}
		envVariables = append(envVariables, tempEnvVariable)
	}

	container.Ports = ports
	container.Env = envVariables
	var containers []v2.Container
	containers = append(containers, container)
	return containers, secretsArray, configMapsArray, nil

}

func getInitContainers(service *pb.DeploymentService) ([]v2.Container, []string, []string, error) {
	var configMapsArray, secretsArray []string
	if !service.ServiceAttributes.EnableInit {
		return nil, secretsArray, configMapsArray, nil
	}

	var container v2.Container
	if service.Name != "" {
		container.Name = service.Name
	} else {
		container.Name = "init-container-dummy"
	}
	if err := putCommandAndArguments(&container, service.ServiceAttributes.Command, service.ServiceAttributes.Args); err != nil {
		return nil, secretsArray, configMapsArray, err
	}
	if err := putLimitResource(&container, service.ServiceAttributes.LimitResources); err != nil {
		return nil, secretsArray, configMapsArray, err
	}
	if err := putRequestResource(&container, service.ServiceAttributes.RequestResources); err != nil {
		return nil, secretsArray, configMapsArray, err
	}

	if service.ServiceAttributes.SecurityContext != nil {
		if securityContext, err := configureSecurityContext(service.ServiceAttributes.SecurityContext); err != nil {
			return nil, secretsArray, configMapsArray, err
		} else {
			container.SecurityContext = securityContext
		}
	}

	container.Image = service.ServiceAttributes.ImagePrefix + service.ServiceAttributes.ImageName
	if service.ServiceAttributes.Tag != "" {
		container.Image += ":" + service.ServiceAttributes.Tag
	}
	var ports []v2.ContainerPort

	for _, port := range service.ServiceAttributes.Ports {
		temp := v2.ContainerPort{}
		if port.Container == "" && port.Host == "" {
			continue
		}
		if port.Container == "" && port.Host != "" {
			port.Container = port.Host
		}

		i, err := strconv.Atoi(port.Container)
		if err != nil {
			utils.Info.Println(err)
			continue
		}
		if i > 0 && i < 65536 {
			temp.ContainerPort = int32(i)
		} else {
			utils.Info.Println("invalid prot number")
			continue
		}
		if port.Host != "" {
			i, err = strconv.Atoi(port.Host)
			if err != nil {
				utils.Info.Println(err)
				continue
			}
			if i > 0 && i < 65536 {
				temp.HostPort = int32(i)
			} else {
				utils.Info.Println("invalid prot number")
				continue
			}

		}
		ports = append(ports, temp)
	}

	var envVariables []v2.EnvVar
	for key, envVariable := range service.ServiceAttributes.EnvironmentVariables {
		tempEnvVariable := v2.EnvVar{}
		if strings.EqualFold(key, "ConfigMap") {
			envVariableValue := strings.Split(envVariable.Value, ";")
			tempEnvVariable = v2.EnvVar{Name: envVariable.Key,
				ValueFrom: &v2.EnvVarSource{ConfigMapKeyRef: &v2.ConfigMapKeySelector{
					LocalObjectReference: v2.LocalObjectReference{Name: envVariableValue[0]},
					Key:                  envVariableValue[1],
				}}}

		} else if strings.EqualFold(key, "Secret") {
			envVariableValue := strings.Split(envVariable.Value, ";")
			tempEnvVariable = v2.EnvVar{Name: envVariable.Key,
				ValueFrom: &v2.EnvVarSource{SecretKeyRef: &v2.SecretKeySelector{
					LocalObjectReference: v2.LocalObjectReference{Name: envVariableValue[0]},
					Key:                  envVariableValue[1],
				}}}
		} else {
			tempEnvVariable = v2.EnvVar{Name: envVariable.Key, Value: envVariable.Value}
		}
		envVariables = append(envVariables, tempEnvVariable)
	}
	container.Ports = ports
	container.Env = envVariables
	var containers []v2.Container
	containers = append(containers, container)
	return containers, secretsArray, configMapsArray, nil

}

func putCommandAndArguments(container *v2.Container, command, args []string) error {
	if len(command) > 0 && command[0] != "" {
		container.Command = command
		if len(args) > 0 {
			container.Args = args
		} else {
			container.Args = []string{}
		}

	} else if len(args) > 0 {
		container.Args = args
	}
	return nil
}

func checkRequestIsLessThanLimit(serviceAttr *pb.DeploymentServiceAttributes) (error, bool) {
	for t, v := range serviceAttr.LimitResources {
		r, found := serviceAttr.RequestResources[t]
		if found {
			rr, err := resource.ParseQuantity(r)
			if err != nil {
				return err, false
			}
			lr, err := resource.ParseQuantity(v)
			if err != nil {
				return err, false
			}
			rrint := rr.AsDec()
			lrint := lr.AsDec()
			if rrint.Cmp(lrint) == 1 {
				return nil, false
			}
		}
	}
	return nil, true
}

func putLivenessProbe(container *v2.Container, prob *pb.Probe) error {

	var temp v2.Probe
	if prob != nil {
		if prob.Handler != nil {
			temp.InitialDelaySeconds = prob.InitialDelaySeconds
			temp.FailureThreshold = prob.FailureThreshold
			temp.PeriodSeconds = prob.PeriodSeconds
			temp.SuccessThreshold = prob.SuccessThreshold
			temp.TimeoutSeconds = prob.TimeoutSeconds
			typeHandler := prob.Handler.HandlerType
			switch typeHandler {
			case "exec":
				if prob.Handler.Exec == nil {
					return errors.New("there is no liveness handler of exec type")
				}
				temp.Handler.Exec = &v2.ExecAction{}
				for i := 0; i < len(prob.Handler.Exec.Command); i++ {
					temp.Handler.Exec.Command = append(temp.Handler.Exec.Command, prob.Handler.Exec.Command[i])
				}

			case "http_get":
				if prob.Handler.HttpGet == nil {
					return errors.New("there is no liveness handler of httpGet type")
				}
				temp.Handler.HTTPGet = &v2.HTTPGetAction{}
				if prob.Handler.HttpGet.Port > 0 && prob.Handler.HttpGet.Port < 65536 {

					temp.HTTPGet.Host = prob.Handler.HttpGet.Host

					temp.HTTPGet.Path = prob.Handler.HttpGet.Path

					if strings.EqualFold(prob.Handler.HttpGet.Scheme, types.URISchemeHTTP) || strings.EqualFold(prob.Handler.HttpGet.Scheme, types.URISchemeHTTPS) {

						temp.HTTPGet.Scheme = v2.URIScheme(prob.Handler.HttpGet.Scheme)
					} else {
						return errors.New("invalid urischeme ")
					}

					if prob.Handler.HttpGet.HttpHeaders != nil {
						temp.HTTPGet.HTTPHeaders = []v2.HTTPHeader{}
						for i := 0; i < len(prob.Handler.HttpGet.HttpHeaders); i++ {
							tempheader := v2.HTTPHeader{prob.Handler.HttpGet.HttpHeaders[i].Name, prob.Handler.HttpGet.HttpHeaders[i].Value}
							temp.HTTPGet.HTTPHeaders = append(temp.HTTPGet.HTTPHeaders, tempheader)
						}
					}
					temp.HTTPGet.Port = intstr.FromInt(int(prob.Handler.HttpGet.Port))
				} else {
					return errors.New("Invalid Port number for http Get")
				}
			case "tcpSocket":
				if prob.Handler.TcpSocket == nil {
					return errors.New("there is no liveness handler of tcpSocket type")
				}
				temp.Handler.TCPSocket = &v2.TCPSocketAction{}
				if prob.Handler.TcpSocket.Port > 0 && prob.Handler.TcpSocket.Port < 65536 {
					temp.TCPSocket.Port = intstr.FromInt(int(prob.Handler.TcpSocket.Port))
					temp.TCPSocket.Host = prob.Handler.TcpSocket.Host

				} else {
					return errors.New("Invalid Port number for tcp socket")
				}

			default:
				return errors.New("There Must be liveness handler of valid type")

			}
		} else {
			return errors.New("Liveness prob header can not be nil")
		}
		container.LivenessProbe = &temp
	}
	return nil
}
func putReadinessProbe(container *v2.Container, prob *pb.Probe) error {
	var temp v2.Probe
	if prob != nil {
		if prob.Handler != nil {
			temp.InitialDelaySeconds = prob.InitialDelaySeconds
			temp.FailureThreshold = prob.FailureThreshold
			temp.PeriodSeconds = prob.PeriodSeconds
			temp.SuccessThreshold = prob.SuccessThreshold
			temp.TimeoutSeconds = prob.TimeoutSeconds

			switch typeHandler := prob.Handler.HandlerType; typeHandler {
			case "exec":
				if prob.Handler.Exec == nil {
					return errors.New("there is no readiness handler of exec type")
				}
				temp.Handler.Exec = &v2.ExecAction{}
				for i := 0; i < len(prob.Handler.Exec.Command); i++ {
					temp.Handler.Exec.Command = append(temp.Handler.Exec.Command, prob.Handler.Exec.Command[i])
				}

			case "http_get":
				if prob.Handler.HttpGet == nil {
					return errors.New("there is no readiness handler of httpGet type")
				}
				temp.Handler.HTTPGet = &v2.HTTPGetAction{}
				if prob.Handler.HttpGet.Port > 0 && prob.Handler.HttpGet.Port < 65536 {
					temp.HTTPGet.Host = prob.Handler.HttpGet.Host
					temp.HTTPGet.Path = prob.Handler.HttpGet.Path

					if strings.EqualFold(prob.Handler.HttpGet.Scheme, types.URISchemeHTTP) || strings.EqualFold(prob.Handler.HttpGet.Scheme, types.URISchemeHTTPS) {

						temp.HTTPGet.Scheme = v2.URIScheme(prob.Handler.HttpGet.Scheme)
					} else {
						return errors.New("invalid urischeme ")
					}

					if prob.Handler.HttpGet.HttpHeaders != nil {
						temp.HTTPGet.HTTPHeaders = []v2.HTTPHeader{}
						for i := 0; i < len(prob.Handler.HttpGet.HttpHeaders); i++ {
							tempheader := v2.HTTPHeader{prob.Handler.HttpGet.HttpHeaders[i].Name, prob.Handler.HttpGet.HttpHeaders[i].Value}
							temp.HTTPGet.HTTPHeaders = append(temp.HTTPGet.HTTPHeaders, tempheader)
						}
					}
					temp.HTTPGet.Port = intstr.FromInt(int(prob.Handler.HttpGet.Port))
				} else {
					return errors.New("Invalid Port number for http Get")
				}
			case "tcpSocket":
				if prob.Handler.TcpSocket == nil {
					return errors.New("there is no readiness handler of tcpSocket type")
				}
				temp.Handler.TCPSocket = &v2.TCPSocketAction{}
				if prob.Handler.TcpSocket.Port > 0 && prob.Handler.TcpSocket.Port < 65536 {
					temp.TCPSocket.Port = intstr.FromInt(int(prob.Handler.TcpSocket.Port))

					temp.TCPSocket.Host = prob.Handler.TcpSocket.Host

				} else {
					return errors.New("Invalid Port number for tcp socket")
				}

			default:
				return errors.New("There Must be readiness handler of valid type")

			}
		} else {
			return errors.New("Readiness prob handler can not be nil")
		}
		container.ReadinessProbe = &temp
	}
	return nil
}

func configureSecurityContext(securityContext *pb.SecurityContextStruct) (*v2.SecurityContext, error) {
	var context v2.SecurityContext
	context.Capabilities = &v2.Capabilities{}
	for _, capability := range securityContext.Capabilities {
		for _, add := range capability.Add {
			context.Capabilities.Add = append(context.Capabilities.Add, v2.Capability(add))
		}
		for _, dropCapability := range capability.Drop {
			context.Capabilities.Drop = append(context.Capabilities.Drop, v2.Capability(dropCapability))
		}
	}
	context.ReadOnlyRootFilesystem = &securityContext.ReadOnlyRootFilesystem
	context.Privileged = &securityContext.Privileged
	if securityContext.RunAsNonRoot && securityContext.RunAsUser == 0 {
		return nil, errors.New("RunAsNonRoot is Set, but RunAsUser value not given!")
	} else {
		context.RunAsNonRoot = &securityContext.RunAsNonRoot
		context.RunAsUser = &securityContext.RunAsUser
	}
	context.RunAsGroup = &securityContext.RunAsGroup
	context.AllowPrivilegeEscalation = &securityContext.AllowPrivilegeEscalation

	var procmount = securityContext.ProcMount.String()

	tempProcMount := v2.ProcMountType(procmount)
	context.ProcMount = &tempProcMount

	context.SELinuxOptions = &v2.SELinuxOptions{
		User:  securityContext.SeLinuxOptions.User,
		Role:  securityContext.SeLinuxOptions.Role,
		Type:  securityContext.SeLinuxOptions.Type,
		Level: securityContext.SeLinuxOptions.Level,
	}
	return &context, nil

}

func getLabels(data *pb.DeploymentService) (map[string]string, error) {

	var serviceAttributes types.DockerServiceAttributes
	if data, err := json.Marshal(data.ServiceAttributes); err == nil {

		if err = json.Unmarshal(data, &serviceAttributes); err != nil {
			utils.Error.Println(err)
			return nil, err
		}
	} else {
		return nil, err
	}
	labels := make(map[string]string)
	for key, value := range serviceAttributes.Labels {
		labels[key] = value
	}
	return labels, nil
}

func resolveValue(value string) (string, string) {
	result := returnValuesArray(value)
	var splittedResult []string
	for _, every := range result {
		if len(every) > 0 {
			splittedResult = strings.Split(every, ";")
			break
		}
	}
	keyArray := strings.Split(splittedResult[1], ".")
	return splittedResult[0], keyArray[len(keyArray)-1]
}

func returnValuesArray(value string) []string {
	replacer := strings.NewReplacer(
		"{{", "|-|",
		"}}", "|-|")
	result := replacer.Replace(value)
	return strings.Split(result, "|-|")
}

func putRequestResource(container *v2.Container, requestResources map[string]string) error {
	temp := make(map[v2.ResourceName]resource.Quantity)
	for t, v := range requestResources {
		if t == types.ResourceTypeCpu || t == types.ResourceTypeMemory {
			quantity, err := resource.ParseQuantity(v)
			if err != nil {
				return err
			}
			//
			temp[v2.ResourceName(t)] = quantity
		} else {
			return errors.New("Error Found: Invalid Request Resource Provided. Valid: 'cpu','memory'")
		}
	}

	container.Resources.Requests = temp
	return nil
}

func putLimitResource(container *v2.Container, limitResources map[string]string) error {
	temp := make(map[v2.ResourceName]resource.Quantity)
	for t, v := range limitResources {
		if t == types.ResourceTypeMemory || t == types.ResourceTypeCpu {
			quantity, err := resource.ParseQuantity(v)
			if err != nil {
				return err
			}
			temp[v2.ResourceName(t)] = quantity
		} else {
			return errors.New("Error Found: Invalid Request Resource Provided. Valid: 'cpu','memory'")
		}
	}

	container.Resources.Limits = temp
	return nil
}
