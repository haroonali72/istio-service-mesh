package main

import (
	"Istio/constants"
	"Istio/controllers"
	"Istio/utils"
	"flag"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {


	controllers.Notifier.Init_notifier()
	utils.LoggerInit(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

	Port := flag.String("PORT", "", "Port of service")
	kubeEngine := flag.String("KUBERNETES_ENGINE_URL", "", "Kubernetes engine url")
	redisEngine := flag.String("REDIS_ENGINE_URL", "", "Redis url")
	loggingEngine := flag.String("LOGGING_ENGINE_URL", "", "Logger url")

	flag.Parse()
	if *Port == ""{
		log.Fatal("PORT flag missing.\nTerminating....")
		os.Exit(1)
	}

	if *kubeEngine == ""{
		log.Fatal("KUBERNETES_ENGINE_URL flag missing.\nTerminating....")
		os.Exit(1)
	}

	if *redisEngine == ""{
		log.Fatal("REDIS_ENGINE_URL flag missing.\nTerminating....")
		os.Exit(1)
	}

	if *loggingEngine == ""{
		log.Fatal("LOGGING_ENGINE_URL flag missing.\nTerminating....")
		os.Exit(1)
	}
	constants.ServicePort = *Port
	constants.KubernetesEngineURL = *kubeEngine
	constants.NotificationURL = *redisEngine
	constants.LoggingURL = *loggingEngine


	r := mux.NewRouter()
	r.HandleFunc("/istioservicedeployer", controllers.ServiceRequest)
	log.Fatal(http.ListenAndServe(":"+constants.ServicePort, r))
}
