package main

import (
	"context"
	"fmt"
	"go.opencensus.io/plugin/ocgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"io/ioutil"
	"istio-service-mesh/constants"
	"istio-service-mesh/core"
	pb "istio-service-mesh/core/proto"
	"istio-service-mesh/utils"
	"log"
	"net"
	"os"
)

func main() {

	/*Port := flag.String("PORT", "", "Port of service")
	loggingEngine := flag.String("LOGGING_ENGINE_URL", "", "Logger url")
	kubeEngine := flag.String("KUBERNETES_ENGINE_URL", "", "Kubernetes engine url")
	redisEngine := flag.String("REDIS_ENGINE_URL", "", "Redis url")
	vaultEngine := flag.String("VAULT_ENGINE_URL", "", "vault url")
	rbac := flag.String("RBAC_URL", "", "Rbac url")

	flag.Parse()
	if *Port == "" {
		log.Fatal("PORT flag missing.\nTerminating....")
		os.Exit(1)
	}

	if *kubeEngine == "" {
		log.Fatal("KUBERNETES_ENGINE_URL flag missing.\nTerminating....")
		os.Exit(1)
	}

	if *redisEngine == "" {
		log.Fatal("REDIS_ENGINE_URL flag missing.\nTerminating....")
		os.Exit(1)
	}

	if *loggingEngine == "" {
		log.Fatal("LOGGING_ENGINE_URL flag missing.\nTerminating....")
		os.Exit(1)
	}
	if *vaultEngine == "" {
		log.Fatal("Vault_ENGINE_URL flag missing.\nTerminating....")
		os.Exit(1)
	}
	constants.ServicePort = *Port
	constants.KubernetesEngineURL = *kubeEngine
	constants.NotificationURL = *redisEngine
	constants.LoggingURL = *loggingEngine
	constants.VaultURL = *vaultEngine
	constants.RbacURL = *rbac*/

	utils.LoggerInit(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	err := utils.InitFlags()
	if err != nil {
		panic(err)
	}

	port := fmt.Sprintf(":%s", constants.ServicePort)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer(grpc.StatsHandler(&ocgrpc.ServerHandler{}))
	svc := &core.Server{}
	pb.RegisterGatewayServer(srv, svc)
	pb.RegisterClusterroleServer(srv, svc)
	pb.RegisterClusterrolebindingServer(srv, svc)
	pb.RegisterHpaServer(srv, svc)
	pb.RegisterVirtualServer(srv, svc)
	pb.RegisterDestinationrulesServer(srv, svc)
	pb.RegisterK8SResourceServer(srv, svc)
	go handleclient()
	pb.RegisterServiceEntryServer(srv, svc)
	pb.RegisterDeploymentServer(srv, svc)
	pb.RegisterStorageClassServer(srv, svc)
	pb.RegisterYamlServiceServer(srv, svc)
	pb.RegisterYamlToCPServiceServer(srv, svc)
	// Register reflection service on gRPC server.
	reflection.Register(srv)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	/*r := mux.NewRouter()
	r.HandleFunc("/istioservicedeployer", controllers.ServiceRequest)
	r.HandleFunc("/importservice", controllers.ImportServiceRequest)
	log.Fatal(http.ListenAndServe(":"+constants.ServicePort, r))*/
}

func handleclient() {
	conn, err := grpc.Dial("localhost:8654", grpc.WithInsecure())
	if err != nil {
		utils.Error.Println(err)
	}

	_, err = pb.NewK8SResourceClient(conn).GetK8SResource(context.Background(), &pb.K8SResourceRequest{
		ProjectId: "11",
		CompanyId: "1111",
		Token:     "dfsdfsdf",
	})
	if err != nil {
		utils.Error.Println(err)
	}

}
