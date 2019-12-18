package core
import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"istio-service-mesh/constants"
	pb "istio-service-mesh/core/proto"
	"istio-service-mesh/utils"
	"istio.io/api/networking/v1alpha3"
	istioClient "istio.io/client-go/pkg/apis/networking/v1alpha3"
	"strings"
)

type RoleBindingServer struct{
}
func (s *RoleBindingServer)CreateRoleBinding(ctx context.Context,req *pb.RoleBindingService)(*pb.ServiceResponse,error){
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id: req.ServiceId,
		ServiceId: req.ServiceId,
		Name: req.Name,
	}
	ksdRequest ,err := getRoleBindingRequestObject(req)

	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp,err)
		return serviceResp,err
	}

	conn, err := grpc.DialContext(ctx,constants.K8sEngineGRPCURL,grpc.WithInsecure())
	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp,err)
		return serviceResp,err
	}
	defer conn.Close()


	raw, err := json.Marshal(ksdRequest)
	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp,err)
		return serviceResp,err
	}
	result, err := pb.NewServiceClient(conn).CreateService(ctx,&pb.ServiceRequest{
		ProjectId:req.ProjectId,
		Service: raw,
		CompanyId:req.CompanyId,
		Token: req.Token,
	})
	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp,err)
		return serviceResp,err
	}
	utils.Info.Println(result.Service)
	serviceResp.Status.Status = "successful"
	serviceResp.Status.StatusIndividual = append(serviceResp.Status.StatusIndividual ,"successful")

	return serviceResp,nil

}
func (s *RoleBindingServer)GetRoleBinding(ctx context.Context,req *pb.RoleBindingService)(*pb.ServiceResponse,error){
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id: req.ServiceId,
		ServiceId: req.ServiceId,
		Name: req.Name,
	}
	ksdRequest ,err := getRoleBindingRequestObject(req)

	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp,err)
		return serviceResp,err
	}

	conn, err := grpc.DialContext(ctx,constants.K8sEngineGRPCURL,grpc.WithInsecure())
	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp,err)
		return serviceResp,err
	}
	defer conn.Close()

	raw, err := json.Marshal(ksdRequest)
	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp,err)
		return serviceResp,err
	}
	result, err := pb.NewServiceClient(conn).GetService(ctx,&pb.ServiceRequest{
		ProjectId:req.ProjectId,
		Service: raw,
		CompanyId:req.CompanyId,
		Token: req.Token,
	})
	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp,err)
		return serviceResp,err
	}
	utils.Info.Println(result.Service)
	serviceResp.Status.Status = "successful"
	serviceResp.Status.StatusIndividual = append(serviceResp.Status.StatusIndividual ,"successful")

	return serviceResp,nil
}
func (s *RoleBindingServer)DeleteRoleBinding(ctx context.Context,req *pb.RoleBindingService)(*pb.ServiceResponse,error){
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id: req.ServiceId,
		ServiceId: req.ServiceId,
		Name: req.Name,
	}
	ksdRequest ,err := getRoleBindingRequestObject(req)

	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp,err)
		return serviceResp,err
	}

	conn, err := grpc.DialContext(ctx,constants.K8sEngineGRPCURL,grpc.WithInsecure())
	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp,err)
		return serviceResp,err
	}
	defer conn.Close()

	raw, err := json.Marshal(ksdRequest)
	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp,err)
		return serviceResp,err
	}
	result, err := pb.NewServiceClient(conn).DeleteService(ctx,&pb.ServiceRequest{
		ProjectId:req.ProjectId,
		Service: raw,
		CompanyId:req.CompanyId,
		Token: req.Token,
	})
	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp,err)
		return serviceResp,err
	}
	utils.Info.Println(result.Service)
	serviceResp.Status.Status = "successful"
	serviceResp.Status.StatusIndividual = append(serviceResp.Status.StatusIndividual ,"successful")

	return serviceResp,nil
}
func (s *RoleBindingServer)PatchRoleBinding(ctx context.Context,req *pb.RoleBindingService)(*pb.ServiceResponse,error){
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id: req.ServiceId,
		ServiceId: req.ServiceId,
		Name: req.Name,
	}
	ksdRequest ,err := getRoleBindingRequestObject(req)

	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp,err)
		return serviceResp,err
	}

	conn, err := grpc.DialContext(ctx,constants.K8sEngineGRPCURL,grpc.WithInsecure())
	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp,err)
		return serviceResp,err
	}
	defer conn.Close()

	raw, err := json.Marshal(ksdRequest)
	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp,err)
		return serviceResp,err
	}
	result, err := pb.NewServiceClient(conn).PatchService(ctx,&pb.ServiceRequest{
		ProjectId:req.ProjectId,
		Service: raw,
		CompanyId:req.CompanyId,
		Token: req.Token,
	})
	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp,err)
		return serviceResp,err
	}
	utils.Info.Println(result.Service)
	serviceResp.Status.Status = "successful"
	serviceResp.Status.StatusIndividual = append(serviceResp.Status.StatusIndividual ,"successful")

	return serviceResp,nil
}
func (s *RoleBindingServer)PutRoleBinding(ctx context.Context,req *pb.RoleBindingService)(*pb.ServiceResponse,error){
	serviceResp := new(pb.ServiceResponse)
	serviceResp.Status = &pb.ServiceStatus{
		Id: req.ServiceId,
		ServiceId: req.ServiceId,
		Name: req.Name,
	}
	ksdRequest ,err := getRoleRequestObject(req)

	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp,err)
		return serviceResp,err
	}

	conn, err := grpc.DialContext(ctx,constants.K8sEngineGRPCURL,grpc.WithInsecure())
	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp,err)
		return serviceResp,err
	}
	defer conn.Close()

	raw, err := json.Marshal(ksdRequest)
	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp,err)
		return serviceResp,err
	}
	result, err := pb.NewServiceClient(conn).PutService(ctx,&pb.ServiceRequest{
		ProjectId:req.ProjectId,
		Service: raw,
		CompanyId:req.CompanyId,
		Token: req.Token,
	})
	if err != nil {
		utils.Error.Println(err)
		getErrorResp(serviceResp,err)
		return serviceResp,err
	}
	utils.Info.Println(result.Service)
	serviceResp.Status.Status = "successful"
	serviceResp.Status.StatusIndividual = append(serviceResp.Status.StatusIndividual ,"successful")

	return serviceResp,nil
}


func getRoleBinding(input *pb.RoleBindingService) (*istioClient.Gateway, error) {
	var istioServ = new(istioClient.Gateway)
	labels := make(map[string]string)
	labels["app"] = strings.ToLower(input.Name)
	labels["name"] = strings.ToLower(input.Name)
	labels["version"] = strings.ToLower(input.Version)
	labels["namespace"] = strings.ToLower(input.Namespace)
	istioServ.Labels = labels
	istioServ.Kind = "Gateway"
	istioServ.APIVersion = "networking.istio.io/v1alpha3"

	gateway := v1alpha3.Gateway{}

	gateway.Selector = input.ServiceAttributes.Selectors

	raw, err := json.Marshal(input.ServiceAttributes)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal gateway input object. Error: %s",err.Error())
	}
	err = json.Unmarshal(raw,gateway)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal gateway input object to istio object. Error: %s",err.Error())
	}

	istioServ.Spec = gateway
	return istioServ, nil
}
func getRoleBindingSpec() (v1alpha3.Gateway, error) {

	gateway := v1alpha3.Gateway{}
	var hosts []string
	hosts = append(hosts, "*")
	var servers []*v1alpha3.Server

	var serv v1alpha3.Server
	serv.Port = &v1alpha3.Port{Name: strings.ToLower("HTTP"), Protocol: "HTTP", Number: uint32(80)}
	serv.Hosts = hosts
	servers = append(servers, &serv)

	/*var serv2 v1alpha3.Server
	serv2.Port = &v1alpha3.Port{Name: strings.ToLower("HTTPS"), Protocol: "HTTPS", Number: uint32(443)}
	serv2.Hosts = hosts
	servers = append(servers, &serv2)*/

	selector := make(map[string]string)

	selector["istio"] = "ingressgateway"
	gateway.Selector = selector
	gateway.Servers = servers
	return gateway, nil
}

func getRoleBindingRequestObject(req *pb.RoleBindingService)(*istioClient.Gateway, error){
	gtwReq, err := getIstioGateway(req)
	if err != nil {
		utils.Error.Println(err)

		return nil,err
	}
	return gtwReq,nil
}
