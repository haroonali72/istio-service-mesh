package main

import (
	"Istio/constants"
	"Istio/controllers"
	"Istio/utils"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {

	err := utils.InitFlags()
	if(err != nil){
		log.Fatal(err)
		os.Exit(0)
	}
	fmt.Printf(constants.NotificationURL)
	fmt.Printf(constants.KubernetesEngineURL)
	fmt.Printf(constants.LoggingURL)
	fmt.Printf(constants.ServicePort)

	controllers.Notifier.Init_notifier()
	utils.LoggerInit(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	r := mux.NewRouter()
	r.HandleFunc("/istioservicedeployer", controllers.ServiceRequest)
	log.Fatal(http.ListenAndServe(":"+constants.ServicePort, r))
}
