package main

import (
	"Istio/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)



func main() {
	/*client, err := kubernetes.NewForConfig(&rest.Config{Host: "https://3.84.228.162:6443", Username: "cloudplex", Password: "64bdySICej", TLSClientConfig: rest.TLSClientConfig{Insecure: true}})
	fmt.Println(err)
	pods, err := client.CoreV1().Pods("default").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for i := range pods.Items {
		fmt.Println(pods.Items[i].Name, pods.Items[i].Namespace)
	}
	os.Exit(0)*/
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/istioservicedeployer", controllers.ServiceRequest)
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8654", r))
}
