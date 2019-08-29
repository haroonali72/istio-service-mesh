package main

import (
	"flag"
	"github.com/gorilla/mux"
	"io/ioutil"
	"istio-service-mesh/constants"
	"istio-service-mesh/controllers"
	"istio-service-mesh/utils"
	"log"
	"net/http"
	"os"
)

func main() {

	Port := flag.String("PORT", "", "Port of service")
	kubeEngine := flag.String("KUBERNETES_ENGINE_URL", "", "Kubernetes engine url")
	redisEngine := flag.String("REDIS_ENGINE_URL", "", "Redis url")
	loggingEngine := flag.String("LOGGING_ENGINE_URL", "", "Logger url")
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
	constants.RbacURL = *rbac
	utils.LoggerInit(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	controllers.Notifier.Init_notifier()

	r := mux.NewRouter()
	r.HandleFunc("/istioservicedeployer", controllers.ServiceRequest)
	r.HandleFunc("/importservice", controllers.ImportServiceRequest)
	log.Fatal(http.ListenAndServe(":"+constants.ServicePort, r))
}
