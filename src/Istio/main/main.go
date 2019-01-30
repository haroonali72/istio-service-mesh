package main

import (
	"Istio/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)



func main() {


	//os.Exit(0)
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/istioservicedeployer", controllers.ServiceRequest)
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8654", r))
}
