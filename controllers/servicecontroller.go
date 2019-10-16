package controllers

import (
	"encoding/json"
	"errors"
	"github.com/iancoleman/strcase"
	"io/ioutil"
	"istio-service-mesh/constants"
	"istio-service-mesh/types"
	"istio-service-mesh/utils"
	"k8s.io/api/apps/v1"
	batchV1 "k8s.io/api/batch/v1"
	v1beta12 "k8s.io/api/batch/v1beta1"
	coreV1 "k8s.io/api/core/v1"
	rbacV1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	"net/http"
	"strconv"
	"strings"
)

func ImportServiceRequest(w http.ResponseWriter, r *http.Request) {

	//b,err:=json.Marshal(r.Body)
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	svc, errs := GetServices(b)
	utils.Info.Println(len(svc))
	result := struct {
		Services []types.Service `json:"services"`
		Errors   []string        `json:"errors"`
	}{
		Services: svc,
	}
	for i := range errs {
		result.Errors = append(result.Errors, errs[i].Error())
	}
	x, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(x)
}
func GetServices(rawData []byte) (svcs []types.Service, errs []error) {

	files, errs := parseK8sYaml(rawData)
	if len(files) == 0 {
		return svcs, errs
	}
	var deployments []types.Service
	var k8sRoles []types.K8sRbacAttribute
	var istioRoles []types.IstioRbacAttribute
	for kindName, obj := range files {
		svc := types.Service{}
		svc.Version = "v1"
		metdata := strings.Split(kindName, ";")
		kind := metdata[0]
		switch constants.K8sKind(kind) {
		case constants.Deployment:
			initcont := false
			var dep v1.Deployment
			err := json.Unmarshal(obj, &dep)
			if err != nil {
				errs = append(errs, err)
			} else {
				svc.Namespace = dep.Namespace
				svc.Name = dep.Name
				svc.ServiceType = "container"
				svc.SubType = strings.ToLower(string(constants.Deployment))
				if len(dep.Spec.Template.Spec.InitContainers) > 0 {
					svc1 := types.Service{
						Name:        dep.Spec.Template.Spec.InitContainers[0].Name, //becasue we only allow one init-cont
						Namespace:   dep.Namespace,
						ServiceType: "init_container",
						SubType:     "",
						Version:     "v1",
					}
					for _, initContainer := range dep.Spec.Template.Spec.InitContainers {
						attrib, err := getContainerData(&initContainer)
						if err != nil {
							utils.Error.Println(err)
							errs = append(errs, err)

						} else {
							svc1.ServiceAttributes = attrib
							svcs = append(svcs, svc1)
							initcont = true
						}
					}
				}
				if len(dep.Spec.Template.Spec.Containers) > 0 {
					attrib, err := getContainerData(&dep.Spec.Template.Spec.Containers[0])
					if err != nil {
						utils.Error.Println(err)
						errs = append(errs, err)

					} else {
						addData(&attrib, k8sRoles, istioRoles, nil)
						svc.ServiceAttributes = attrib
						if initcont {
							map1 := make(map[string]interface{})
							byteData, err := json.Marshal(svc.ServiceAttributes)
							if err != nil {
								utils.Error.Println(err)
								errs = append(errs, err)
							} else {
								err = json.Unmarshal(byteData, &map1)
								if err != nil {
									utils.Error.Println(err)
									errs = append(errs, err)
								} else {
									map1["enable_init"] = true
									svc.ServiceAttributes = map1
								}
							}

						}
						svcs = append(svcs, svc)
					}

				}
				deployments = append(deployments, svc)
			}
		case constants.StatefulSet:
			initcont := false
			var ss v1.StatefulSet
			err := json.Unmarshal(obj, &ss)
			if err != nil {
				errs = append(errs, err)
			} else {
				svc.Namespace = ss.Namespace
				svc.Name = ss.Name
				svc.ServiceType = "container"
				svc.SubType = strings.ToLower(string(constants.StatefulSet))
				if len(ss.Spec.Template.Spec.InitContainers) > 0 {
					svc1 := types.Service{
						Name:        "in_" + ss.Name,
						Namespace:   ss.Namespace,
						ServiceType: "init_container",
						SubType:     "",
						Version:     "v1",
					}
					for _, initContainer := range ss.Spec.Template.Spec.InitContainers {
						attrib, err := getContainerData(&initContainer)
						if err != nil {
							utils.Error.Println(err)
							errs = append(errs, err)

						} else {
							svc1.ServiceAttributes = attrib
							svcs = append(svcs, svc1)
							initcont = true
						}
					}
				}

				if len(ss.Spec.Template.Spec.Containers) > 0 {
					attrib, err := getContainerData(&ss.Spec.Template.Spec.Containers[0])
					if err != nil {
						utils.Error.Println(err)
						errs = append(errs, err)

					} else {
						addData(&attrib, k8sRoles, istioRoles, nil)
						svc.ServiceAttributes = attrib
						if initcont {
							map1 := make(map[string]interface{})
							byteData, err := json.Marshal(svc.ServiceAttributes)
							if err != nil {
								utils.Error.Println(err)
								errs = append(errs, err)
							} else {
								err = json.Unmarshal(byteData, &map1)
								if err != nil {
									utils.Error.Println(err)
									errs = append(errs, err)
								} else {
									map1["enable_init"] = true
									svc.ServiceAttributes = map1
								}
							}

						}
						svcs = append(svcs, svc)
					}

				}
				deployments = append(deployments, svc)
			}
		case constants.CronJob:
			initcont := false
			var ss v1beta12.CronJob
			err := json.Unmarshal(obj, &ss)
			if err != nil {
				errs = append(errs, err)
			} else {
				svc.Namespace = ss.Namespace
				svc.Name = ss.Name
				svc.ServiceType = "container"
				svc.SubType = strings.ToLower(string(constants.StatefulSet))
				if len(ss.Spec.JobTemplate.Spec.Template.Spec.InitContainers) > 0 {
					svc1 := types.Service{
						Name:        "in_" + ss.Name,
						Namespace:   ss.Namespace,
						ServiceType: "init_container",
						SubType:     "",
						Version:     "v1",
					}
					for _, initContainer := range ss.Spec.JobTemplate.Spec.Template.Spec.InitContainers {
						attrib, err := getContainerData(&initContainer)
						if err != nil {
							utils.Error.Println(err)
							errs = append(errs, err)

						} else {
							svc1.ServiceAttributes = attrib
							svcs = append(svcs, svc1)
							initcont = true
						}
					}
				}
				if len(ss.Spec.JobTemplate.Spec.Template.Spec.Containers) > 0 {
					attrib, err := getContainerData(&ss.Spec.JobTemplate.Spec.Template.Spec.Containers[0])
					if err != nil {
						utils.Error.Println(err)
						errs = append(errs, err)

					} else {
						addData(&attrib, k8sRoles, istioRoles, nil)
						svc.ServiceAttributes = attrib
						if initcont {
							map1 := make(map[string]interface{})
							byteData, err := json.Marshal(svc.ServiceAttributes)
							if err != nil {
								utils.Error.Println(err)
								errs = append(errs, err)
							} else {
								err = json.Unmarshal(byteData, &map1)
								if err != nil {
									utils.Error.Println(err)
									errs = append(errs, err)
								} else {
									map1["enable_init"] = true
									svc.ServiceAttributes = map1
								}
							}

						}
						svcs = append(svcs, svc)
					}
				}

				deployments = append(deployments, svc)
			}
		case constants.Job:
			initcont := false
			var ss batchV1.Job
			err := json.Unmarshal(obj, &ss)
			if err != nil {
				errs = append(errs, err)
			} else {
				svc.Namespace = ss.Namespace
				svc.Name = ss.Name
				svc.ServiceType = "container"
				svc.SubType = strings.ToLower(string(constants.StatefulSet))
				if len(ss.Spec.Template.Spec.InitContainers) > 0 {
					svc1 := types.Service{
						Name:        "in_" + ss.Name,
						Namespace:   ss.Namespace,
						ServiceType: "init_container",
						SubType:     "",
						Version:     "v1",
					}
					for _, initContainer := range ss.Spec.Template.Spec.InitContainers {
						attrib, err := getContainerData(&initContainer)
						if err != nil {
							utils.Error.Println(err)
							errs = append(errs, err)

						} else {
							svc1.ServiceAttributes = attrib
							svcs = append(svcs, svc1)
							initcont = true
						}
					}
				}

				if len(ss.Spec.Template.Spec.Containers) > 0 {
					attrib, err := getContainerData(&ss.Spec.Template.Spec.Containers[0])
					if err != nil {
						utils.Error.Println(err)
						errs = append(errs, err)

					} else {
						addData(&attrib, k8sRoles, istioRoles, nil)
						svc.ServiceAttributes = attrib
						if initcont {
							map1 := make(map[string]interface{})
							byteData, err := json.Marshal(svc.ServiceAttributes)
							if err != nil {
								utils.Error.Println(err)
								errs = append(errs, err)
							} else {
								err = json.Unmarshal(byteData, &map1)
								if err != nil {
									utils.Error.Println(err)
									errs = append(errs, err)
								} else {
									map1["enable_init"] = true
									svc.ServiceAttributes = map1
								}
							}
						}
						svcs = append(svcs, svc)
					}
				}
				deployments = append(deployments, svc)
			}
		case constants.Service:
		case constants.ConfigMap:
			var configMapSvc coreV1.ConfigMap
			err := json.Unmarshal(obj, &configMapSvc)
			if err != nil {
				errs = append(errs, err)
			} else {
				svc.Namespace = configMapSvc.Namespace
				svc.Name = configMapSvc.Name
				svc.SubType = strings.ToLower(string(constants.ConfigMap))
				svc.ServiceType = "container"
				var serviceAttr types.ConfigMap
				serviceAttr.Data = configMapSvc.Data
				svc.ServiceAttributes = serviceAttr
				svcs = append(svcs, svc)

			}
		case constants.Secret:
			var secrets coreV1.Secret
			err := json.Unmarshal(obj, &secrets)
			if err != nil {
				errs = append(errs, err)
			} else {
				svc.Namespace = secrets.Namespace
				svc.Name = secrets.Name
				svc.ServiceType = "secrets" //strings.ToLower(string(constants.Secret))
				svc.SubType = string(secrets.Type)
				var serviceAttr types.KubernetesSecret
				serviceAttr.Data = make(map[string]string)
				if secrets.Data != nil {
					for key, value := range secrets.Data {
						serviceAttr.Data[key] = string(value)
					}
				}
				if secrets.StringData != nil {
					serviceAttr.StringData = secrets.StringData
				}
				svc.ServiceAttributes = serviceAttr
				svcs = append(svcs, svc)

			}
		case constants.VirtualService:
		case constants.Gateway:
		case constants.DestinationRule:
		case constants.Policy:
		case constants.ServiceEntry:
		case constants.Role:
			var roleObj rbacV1.Role
			err := json.Unmarshal(obj, &roleObj)
			if err != nil {
				errs = append(errs, err)
			} else {
				roles := rbacObj(roleObj)
				if len(deployments) == 0 {
					k8sRoles = append(k8sRoles, roles...)
				}
				for i := range deployments {
					attrib := deployments[i].ServiceAttributes.(types.DockerServiceAttributes)
					attrib.RbacRoles = roles
					deployments[i].ServiceAttributes = attrib
				}
			}
		case constants.ClusterRole:
			var roleObj rbacV1.ClusterRole
			err := json.Unmarshal(obj, &roleObj)
			if err != nil {
				errs = append(errs, err)
			} else {
				roles := clusterRbacObj(roleObj)
				if len(deployments) == 0 {
					k8sRoles = append(k8sRoles, roles...)
				}
				for i := range deployments {
					attrib := deployments[i].ServiceAttributes.(types.DockerServiceAttributes)
					attrib.RbacRoles = roles
					deployments[i].ServiceAttributes = attrib
				}
			}
		}
	}
	return svcs, errs
}
func addData(ss interface{}, k8sroles []types.K8sRbacAttribute, istioroles []types.IstioRbacAttribute, secrets *types.KubernetesSecret) {
	svc := types.DockerServiceAttributes{}
	byteData, err := json.Marshal(ss)
	if err == nil {
		json.Unmarshal(byteData, svc)
	}
	if len(k8sroles) > 0 {
		svc.RbacRoles = k8sroles
		svc.IsRbac = true
	}
	if len(istioroles) > 0 {
		svc.IstioRoles = istioroles
		svc.IsRbac = true
	}
}
func parseK8sYaml(fileR []byte) (map[string][]byte, []error) {

	fileAsString := string(fileR[:])
	sepYamlfiles := strings.Split(fileAsString, "---")
	files := make(map[string][]byte)
	var errs []error
	for _, f := range sepYamlfiles {
		if f == "\n" || f == "" {
			// ignore empty cases
			continue
		}
		var raw map[string]interface{}
		jsonRawData, err := yaml.ToJSON([]byte(f))
		if err != nil {
			errs = append(errs, err)
			continue
		}
		err = json.Unmarshal(jsonRawData, &raw)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		kind := ""
		if _, ok := raw["kind"]; ok {
			if _, isType := raw["kind"].(string); isType {
				kind = raw["kind"].(string)
				//files[raw["kind"].(string)+";"] = jsonRawData
			}
		}
		name := ""
		if _, ok := raw["metadata"]; ok {
			if metdata, ok := raw["metadata"].(map[string]interface{}); ok {
				if _, exist := metdata["name"]; exist {
					name = metdata["name"].(string)
				}
			}
		}
		files[kind+";"+name] = jsonRawData

	}
	return files, errs
}
func getContainerData(c *coreV1.Container) (str interface{}, err error) {
	service := types.DockerServiceAttributes{}
	service.Command, service.Args = convertCommandAndArguments(c)
	service.LimitResources = convertLimitResource(c)
	service.RequestResources = convertRequestResource(c)
	if c.LivenessProbe != nil {
		limitprob, err := convertLivenessProbe(c)
		if err != nil {
			utils.Error.Println(err)
			return service, err
		}

		service.LivenessProb = &limitprob
	}
	if c.ReadinessProbe != nil {
		requestProb, err := convertReadinessProbe(c)
		if err != nil {
			utils.Error.Println(err)
			return service, err
		}
		service.RedinessProb = &requestProb
	}
	service.SecurityContext, err = revertSecurityContext(c.SecurityContext)
	if err != nil {
		return service, err
	}
	if imageData := strings.Split(c.Image, ":"); len(imageData) == 2 {
		service.ImageName = imageData[0]
		service.Tag = imageData[1]
	}

	for i := range c.Ports {
		containerPort := strconv.FormatInt(int64(c.Ports[i].ContainerPort), 10)
		hostPort := strconv.FormatInt(int64(c.Ports[i].HostPort), 10)
		if hostPort == "0" {
			hostPort = ""
		}
		if containerPort == "0" {
			containerPort = ""
		}
		service.Ports = append(service.Ports, &types.Port{
			Container: containerPort,
			Host:      hostPort,
			Name:      c.Ports[i].Name,
		})
	}

	for _, variable := range c.Env {
		envVar := struct {
			Key         string `json:"key"`
			Value       string `json:"value"`
			IsSecret    bool   `json:"secrets"`
			IsConfigMap bool   `json:"configmap"`
		}{
			variable.Name,
			variable.Value,
			false,
			false,
		}
		if variable.ValueFrom != nil {
			if variable.ValueFrom.SecretKeyRef != nil {
				envVar.Value = "{{" + variable.ValueFrom.SecretKeyRef.Name + ".service_attributes." + variable.ValueFrom.SecretKeyRef.Key + "}}"
				envVar.IsSecret = true
			} else if variable.ValueFrom.ConfigMapKeyRef != nil {
				envVar.Value = "{{" + variable.ValueFrom.ConfigMapKeyRef.Name + ".service_attributes." + variable.ValueFrom.ConfigMapKeyRef.Key + "}}"
				envVar.IsConfigMap = true
			}
		}
		service.EnvironmentVariables = append(service.EnvironmentVariables, &envVar)
	}
	mp := make(map[string]interface{})
	bytedata, err := json.Marshal(service)
	if err == nil {
		err = json.Unmarshal(bytedata, &mp)
	}
	return mp, nil
}
func convertCommandAndArguments(container *coreV1.Container) (command []string, args []string) {
	if len(container.Command) > 0 {
		command = container.Command
		args = container.Args
	}
	return command, args
}

func convertLimitResource(container *coreV1.Container) map[types.RecourceType]string {
	var limitResources = make(map[types.RecourceType]string)
	for rName, rValue := range container.Resources.Limits {
		if rName == coreV1.ResourceCPU || rName == coreV1.ResourceMemory || rName == coreV1.ResourceStorage || rName == coreV1.ResourceEphemeralStorage {
			limitResources[types.RecourceType(rName)] = rValue.String()
		}
	}
	return limitResources
}
func convertRequestResource(container *coreV1.Container) map[types.RecourceType]string {
	var requestResources = make(map[types.RecourceType]string)
	for rName, rValue := range container.Resources.Requests {
		if rName == coreV1.ResourceCPU || rName == coreV1.ResourceMemory || rName == coreV1.ResourceStorage || rName == coreV1.ResourceEphemeralStorage {
			requestResources[types.RecourceType(rName)] = rValue.String()
		}
	}
	return requestResources
}
func convertLivenessProbe(container *coreV1.Container) (livenessprob types.Probe, err error) {
	livenessprob = types.Probe{}
	livenessprob.SuccessThreshold = &container.LivenessProbe.SuccessThreshold
	livenessprob.FailureThreshold = &container.LivenessProbe.FailureThreshold
	livenessprob.TimeoutSeconds = &container.LivenessProbe.TimeoutSeconds
	livenessprob.PeriodSeconds = &container.LivenessProbe.PeriodSeconds
	livenessprob.InitialDelaySeconds = &container.LivenessProbe.InitialDelaySeconds
	livenessprob.Handler = &types.Handler{}
	if container.LivenessProbe.Exec != nil {
		livenessprob.Handler.Exec = (*types.ExecAction)(container.LivenessProbe.Exec)
		livenessprob.Handler.Type = "exec"
	} else if container.LivenessProbe.HTTPGet != nil {
		if port := container.LivenessProbe.HTTPGet.Port.IntValue(); port > 0 && port < 65536 {
			livenessprob.Handler.HTTPGet = &types.HTTPGetAction{}
			livenessprob.Handler.HTTPGet.Port = port
			livenessprob.Handler.HTTPGet.Path = &container.LivenessProbe.HTTPGet.Path
			livenessprob.Handler.HTTPGet.Host = &container.LivenessProbe.HTTPGet.Host
			livenessprob.Handler.HTTPGet.Scheme = (*string)(&container.LivenessProbe.HTTPGet.Scheme)
			for i := 0; i < len(container.LivenessProbe.HTTPGet.HTTPHeaders); i++ {
				var temp = types.HTTPHeader{&container.LivenessProbe.HTTPGet.HTTPHeaders[i].Name, &container.LivenessProbe.HTTPGet.HTTPHeaders[i].Value}
				livenessprob.Handler.HTTPGet.HTTPHeaders = append(livenessprob.Handler.HTTPGet.HTTPHeaders, temp)
			}
			livenessprob.Handler.Type = "httpGet"
		} else {
			return types.Probe{}, errors.New("Invalid Port in Http Get in Liveness Prob")
		}

	} else if container.LivenessProbe.TCPSocket != nil {
		if port := container.LivenessProbe.TCPSocket.Port.IntValue(); port > 0 && port < 65536 {
			livenessprob.Handler.TCPSocket = &types.TCPSocketAction{}
			livenessprob.Handler.TCPSocket.Port = port
			livenessprob.Handler.TCPSocket.Host = &container.LivenessProbe.TCPSocket.Host
			livenessprob.Handler.Type = "tcpSocket"
		} else {
			return types.Probe{}, errors.New("Invalid Port in Tcp Socket in Liveness Prob")

		}

	} else {
		return types.Probe{}, errors.New("handler of liveness prob can not be nill")
	}
	return livenessprob, err
}
func convertReadinessProbe(container *coreV1.Container) (readinessprob types.Probe, err error) {

	readinessprob = types.Probe{}
	readinessprob.SuccessThreshold = &container.ReadinessProbe.SuccessThreshold
	readinessprob.FailureThreshold = &container.ReadinessProbe.FailureThreshold
	readinessprob.TimeoutSeconds = &container.ReadinessProbe.TimeoutSeconds
	readinessprob.PeriodSeconds = &container.ReadinessProbe.PeriodSeconds
	readinessprob.InitialDelaySeconds = &container.ReadinessProbe.InitialDelaySeconds
	readinessprob.Handler = &types.Handler{}
	if container.ReadinessProbe.Exec != nil {
		readinessprob.Handler.Exec = (*types.ExecAction)(container.ReadinessProbe.Exec)
		readinessprob.Handler.Type = "exec"
	} else if container.ReadinessProbe.HTTPGet != nil {
		if port := container.ReadinessProbe.HTTPGet.Port.IntValue(); port > 0 && port < 65536 {
			readinessprob.Handler.HTTPGet = &types.HTTPGetAction{}
			readinessprob.Handler.HTTPGet.Port = port
			readinessprob.Handler.HTTPGet.Path = &container.ReadinessProbe.HTTPGet.Path
			readinessprob.Handler.HTTPGet.Host = &container.ReadinessProbe.HTTPGet.Host
			readinessprob.Handler.HTTPGet.Scheme = (*string)(&container.ReadinessProbe.HTTPGet.Scheme)
			for i := 0; i < len(container.ReadinessProbe.HTTPGet.HTTPHeaders); i++ {
				var temp = types.HTTPHeader{&container.ReadinessProbe.HTTPGet.HTTPHeaders[i].Name, &container.ReadinessProbe.HTTPGet.HTTPHeaders[i].Value}
				readinessprob.Handler.HTTPGet.HTTPHeaders = append(readinessprob.Handler.HTTPGet.HTTPHeaders, temp)
			}
			readinessprob.Handler.Type = "httpGet"
		} else {
			return types.Probe{}, errors.New("Invalid Port in Http Get in Liveness Prob")
		}

	} else if container.ReadinessProbe.TCPSocket != nil {
		if port := container.ReadinessProbe.TCPSocket.Port.IntValue(); port > 0 && port < 65536 {
			readinessprob.Handler.TCPSocket = &types.TCPSocketAction{}
			readinessprob.Handler.TCPSocket.Port = port
			readinessprob.Handler.TCPSocket.Host = &container.ReadinessProbe.TCPSocket.Host
			readinessprob.Handler.Type = "tcpSocket"
		} else {
			return types.Probe{}, errors.New("Invalid Port in Tcp Socket in Liveness Prob")

		}

	} else {
		return types.Probe{}, errors.New("handler of liveness prob can not be nill")
	}
	return readinessprob, err
}
func revertSecurityContext(scontext *coreV1.SecurityContext) (securityContext *types.SecurityContextStruct, err error) {
	if scontext == nil {
		return securityContext, nil
	}
	raw, err := json.Marshal(scontext)
	if err != nil {
		utils.Error.Println(err)
		return securityContext, err
	}
	raw = k8sToSvcKeys(raw)
	err = json.Unmarshal(raw, &securityContext)
	if err != nil {
		utils.Error.Println(err)
		return securityContext, err
	}
	if scontext.Capabilities != nil {
		for _, c := range scontext.Capabilities.Add {
			securityContext.CapabilitiesAdd = append(securityContext.CapabilitiesAdd, string(c))
		}
		for _, c := range scontext.Capabilities.Drop {
			securityContext.CapabilitiesDrop = append(securityContext.CapabilitiesDrop, string(c))
		}
	}
	return securityContext, nil

	/*if scontext.ReadOnlyRootFilesystem != nil {
		securityContext.ReadOnlyRootFileSystem = *scontext.ReadOnlyRootFilesystem
	}
	if scontext.Privileged != nil {
		securityContext.Privileged = *scontext.Privileged
	}
	if scontext.RunAsNonRoot != nil && scontext.RunAsUser != nil {
		securityContext.RunAsNonRoot = *scontext.RunAsNonRoot
		securityContext.RunAsUser = scontext.RunAsUser
	}

	securityContext.RunAsGroup = scontext.RunAsGroup
	securityContext.AllowPrivilegeEscalation = *scontext.AllowPrivilegeEscalation
	if scontext.ProcMount != nil  {
		securityContext.ProcMount = string(*scontext.ProcMount)
	}
	if scontext.SELinuxOptions != nil {
		raw ,err := json.Marshal(scontext.SELinuxOptions)
		if err != nil {

		}
	}*/

}
func k8sToSvcKeys(j json.RawMessage) json.RawMessage {
	m := make(map[string]json.RawMessage)
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		// Not a JSON object
		return j
	}

	for k, v := range m {
		fixed := strcase.ToSnake(k)
		delete(m, k)
		m[fixed] = k8sToSvcKeys(v)
	}
	b, err := json.Marshal(m)
	if err != nil {
		return j
	}

	return json.RawMessage(b)
}
func rbacObj(role rbacV1.Role) (rbacAttrib []types.K8sRbacAttribute) {
	for i := range role.Rules {
		attrib := types.K8sRbacAttribute{
			ApiGroup: role.Rules[i].APIGroups,
			Verbs:    role.Rules[i].Verbs,
		}
		if len(role.Rules[i].Resources) > 0 {
			for _, res := range role.Rules[i].Resources {
				attrib.Resource = res
				rbacAttrib = append(rbacAttrib, attrib)
			}
		}
	}
	return rbacAttrib
}
func clusterRbacObj(role rbacV1.ClusterRole) (rbacAttrib []types.K8sRbacAttribute) {
	for i := range role.Rules {
		attrib := types.K8sRbacAttribute{
			ApiGroup: role.Rules[i].APIGroups,
			Verbs:    role.Rules[i].Verbs,
		}
		if len(role.Rules[i].Resources) > 0 {
			for _, res := range role.Rules[i].Resources {
				attrib.Resource = res
				rbacAttrib = append(rbacAttrib, attrib)
			}
		}
	}
	return rbacAttrib
}
