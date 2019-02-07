package main

import (
	"Istio/constants"
	"Istio/controllers"
	"Istio/utils"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	constants.LoggingURL = "http://elephant:3500"
	constants.NotificationURL = "redis:6379"

	controllers.Notifier.Init_notifier()
	utils.LoggerInit(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	r := mux.NewRouter()
	r.HandleFunc("/istioservicedeployer", controllers.ServiceRequest)
	log.Fatal(http.ListenAndServe(":8654", r))
}
