package main

import (
	"encoding/json"
	"io/ioutil"
	"istio-service-mesh/controllers"
	"istio-service-mesh/utils"
	"os"
)

func main() {
	utils.LoggerInit(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	d, err := os.Open("D:\\Platalytics\\Cloudplex V2\\platform yamls\\platform-files\\elasticsearch.yaml")
	if err != nil {
		utils.Error.Println(err)
		return
	}
	raw, err := ioutil.ReadAll(d)
	if err != nil {
		utils.Error.Println(err)
		return
	}
	svc, errs := controllers.GetServices(raw)
	utils.Error.Println(errs)
	utils.Info.Println(svc)
	for i := range svc {
		raw, _ = json.Marshal(svc[i])
		utils.Info.Println(string(raw))
	}
}
