package controllers

import (
	"encoding/json"
	"fmt"
	"istio-service-mesh/types"
	//v1 "k8s.io/api/core/v1"
	//"k8s.io/apimachinery/pkg/util/intstr"
	"testing"
)

func TestGetDeploymentObject(t *testing.T) {
	ls_ml := make(map[string]string)
	ls_ml["app"] = "myapp"
	ls_ml["lab1"] = "mylabel"
	var lsr = types.LabelSelectorRequirement{"key1", "In", []string{"value1", "value2"}}
	var ls_me = []types.LabelSelectorRequirement{lsr}

	ls := types.LabelSelectorObj{ls_ml, ls_me}
	commands := []string{"/bin/bash", "-c"}
	var rs_cpu = types.RecourceTypeCpu
	rs_mem := types.RecourceType("memory")
	var lm_rcs = make(map[types.RecourceType]string)
	lm_rcs[rs_cpu] = "2"
	lm_rcs[rs_mem] = "200G"
	var r_rcs = make(map[types.RecourceType]string)
	r_rcs[rs_cpu] = "0.5"
	r_rcs[rs_mem] = "20G"

	var exec_handler = types.ExecAction{commands}
	//	path:="/mypath"
	//	sch :=types.URISchemeHTTPS
	//	var http_hader=types.HTTPHeader{&path,&path}
	//	var headerarray =[]types.HTTPHeader{http_hader}
	//	var http_handler=types.HTTPGetAction{&path,3553,nil,&sch,headerarray}

	//var tcp_handler=v1.TCPSocketAction{intstr.FromInt(23456),""}
	var prob_handler = types.Handler{"exec", &exec_handler, nil, nil}
	var re = int32(3)
	var liv_prob = types.Probe{&prob_handler, &re, nil, nil, nil, nil}

	var ds_atr = types.DockerServiceAttributes{Command: commands, Args: nil, LimitResources: lm_rcs, RequestResources: r_rcs, LivenessProb: &liv_prob}
	serviceatr := make(map[string]interface{})
	serviceatr["init_container"] = ds_atr
	serviceatr["label_selector"] = ls
	serviceatr["node_selector"] = ls_ml
	Jon := map[string]interface{}{"name": "cn1", "service_attributes": serviceatr}
	var v types.Service
	if d, err := json.Marshal(Jon); err != nil {
		fmt.Println(err)
	} else if err = json.Unmarshal(d, &v); err != nil {
		fmt.Println(err)
	} else {
		_, err := getDeploymentObject(v)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("done ")
		}
	}
}
func TestGetDaemonSetObject(t *testing.T) {
	ls_ml := make(map[string]string)
	ls_ml["app"] = "myapp"
	ls_ml["lab1"] = "mylabel"
	var lsr = types.LabelSelectorRequirement{"key1", "In", []string{"value1", "value2"}}
	var ls_me = []types.LabelSelectorRequirement{lsr}

	ls := types.LabelSelectorObj{ls_ml, ls_me}
	commands := []string{"/bin/bash", "-c"}
	var rs_cpu = types.RecourceTypeCpu
	rs_mem := types.RecourceType("memory")
	var lm_rcs = make(map[types.RecourceType]string)
	lm_rcs[rs_cpu] = "2"
	lm_rcs[rs_mem] = "200G"
	var r_rcs = make(map[types.RecourceType]string)
	r_rcs[rs_cpu] = "0.5"
	r_rcs[rs_mem] = "20G"

	//var exec_handler=types.ExecAction{commands}
	path := "/mypath"
	sch := types.URISchemeHTTPS
	var http_hader = types.HTTPHeader{&path, &path}
	var headerarray = []types.HTTPHeader{http_hader}
	var http_handler = types.HTTPGetAction{&path, 3553, nil, &sch, headerarray}

	//var tcp_handler=v1.TCPSocketAction{intstr.FromInt(23456),""}
	var prob_handler = types.Handler{"httpGet", nil, &http_handler, nil}
	var re = int32(3)
	var liv_prob = types.Probe{&prob_handler, &re, nil, nil, nil, nil}

	var ds_atr = types.DockerServiceAttributes{Command: commands, Args: nil, LimitResources: lm_rcs, RequestResources: r_rcs, LivenessProb: &liv_prob}
	serviceatr := make(map[string]interface{})
	serviceatr["init_container"] = ds_atr
	serviceatr["label_selector"] = ls
	serviceatr["node_selector"] = ls_ml
	Jon := map[string]interface{}{"name": "cn1", "service_attributes": serviceatr}
	var v types.Service
	if d, err := json.Marshal(Jon); err != nil {
		fmt.Println(err)
	} else if err = json.Unmarshal(d, &v); err != nil {
		fmt.Println(err)
	} else {
		_, err := getDaemonSetObject(v)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("done ")
		}
	}
}
func TestGetCronJobObject(t *testing.T) {
	ls_ml := make(map[string]string)
	ls_ml["app"] = "myapp"
	ls_ml["lab1"] = "mylabel"
	var lsr = types.LabelSelectorRequirement{"key1", "In", []string{"value1", "value2"}}
	var ls_me = []types.LabelSelectorRequirement{lsr}

	ls := types.LabelSelectorObj{ls_ml, ls_me}
	commands := []string{"/bin/bash", "-c"}
	var rs_cpu = types.RecourceTypeCpu
	rs_mem := types.RecourceType("memory")
	var lm_rcs = make(map[types.RecourceType]string)
	lm_rcs[rs_cpu] = "2"
	lm_rcs[rs_mem] = "200G"
	var r_rcs = make(map[types.RecourceType]string)
	r_rcs[rs_cpu] = "0.5"
	r_rcs[rs_mem] = "20G"

	//var exec_handler=types.ExecAction{commands}
	//	path:="/mypath"
	//	sch :=types.URISchemeHTTPS
	//	var http_hader=types.HTTPHeader{&path,&path}
	//	var headerarray =[]types.HTTPHeader{http_hader}
	//	var http_handler=types.HTTPGetAction{&path,3553,nil,&sch,headerarray}

	var tcp_handler = types.TCPSocketAction{23456, nil}
	var prob_handler = types.Handler{"tcpSocket", nil, nil, &tcp_handler}
	var re = int32(3)
	var liv_prob = types.Probe{&prob_handler, &re, nil, nil, nil, nil}

	var ds_atr = types.DockerServiceAttributes{Command: commands, Args: nil, LimitResources: lm_rcs, RequestResources: r_rcs, LivenessProb: &liv_prob}
	serviceatr := make(map[string]interface{})
	serviceatr["init_container"] = ds_atr
	serviceatr["label_selector"] = ls
	serviceatr["node_selector"] = ls_ml
	Jon := map[string]interface{}{"name": "cn1", "service_attributes": serviceatr}
	var v types.Service
	if d, err := json.Marshal(Jon); err != nil {
		fmt.Println(err)
	} else if err = json.Unmarshal(d, &v); err != nil {
		fmt.Println(err)
	} else {
		_, err := getCronJobObject(v)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("done ")
		}
	}
}

func TestGetJobObject(t *testing.T) {
	ls_ml := make(map[string]string)
	ls_ml["app"] = "myapp"
	ls_ml["lab1"] = "mylabel"
	var lsr = types.LabelSelectorRequirement{"key1", "In", []string{"value1", "value2"}}
	var ls_me = []types.LabelSelectorRequirement{lsr}

	ls := types.LabelSelectorObj{ls_ml, ls_me}
	commands := []string{"/bin/bash", "-c"}
	var rs_cpu = types.RecourceTypeCpu
	rs_mem := types.RecourceType("memory")
	var lm_rcs = make(map[types.RecourceType]string)
	lm_rcs[rs_cpu] = "2"
	lm_rcs[rs_mem] = "200G"
	var r_rcs = make(map[types.RecourceType]string)
	r_rcs[rs_cpu] = "0.5"
	r_rcs[rs_mem] = "20G"

	var exec_handler = types.ExecAction{commands}
	//	path:="/mypath"
	//	sch :=types.URISchemeHTTPS
	//	var http_hader=types.HTTPHeader{&path,&path}
	//	var headerarray =[]types.HTTPHeader{http_hader}
	//	var http_handler=types.HTTPGetAction{&path,3553,nil,&sch,headerarray}

	var tcp_handler = types.TCPSocketAction{23456, nil}
	var prob_handler = types.Handler{"exec", &exec_handler, nil, &tcp_handler}
	var re = int32(3)
	var liv_prob = types.Probe{&prob_handler, &re, nil, nil, nil, nil}

	var ds_atr = types.DockerServiceAttributes{Command: commands, Args: nil, LimitResources: lm_rcs, RequestResources: r_rcs, LivenessProb: &liv_prob}
	serviceatr := make(map[string]interface{})
	serviceatr["init_container"] = ds_atr
	serviceatr["label_selector"] = ls
	serviceatr["node_selector"] = ls_ml
	Jon := map[string]interface{}{"name": "cn1", "service_attributes": serviceatr}
	var v types.Service
	if d, err := json.Marshal(Jon); err != nil {
		fmt.Println(err)
	} else if err = json.Unmarshal(d, &v); err != nil {
		fmt.Println(err)
	} else {
		_, err := getJobObject(v)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("done ")
		}
	}
}
func TestGetStatefulSetObject(t *testing.T) {
	ls_ml := make(map[string]string)
	ls_ml["app"] = "myapp"
	ls_ml["lab1"] = "mylabel"
	var lsr = types.LabelSelectorRequirement{"key1", "In", []string{"value1", "value2"}}
	var ls_me = []types.LabelSelectorRequirement{lsr}

	ls := types.LabelSelectorObj{ls_ml, ls_me}
	commands := []string{"/bin/bash", "-c"}
	var rs_cpu = types.RecourceTypeCpu
	rs_mem := types.RecourceType("memory")
	var lm_rcs = make(map[types.RecourceType]string)
	lm_rcs[rs_cpu] = "2"
	lm_rcs[rs_mem] = "200G"
	var r_rcs = make(map[types.RecourceType]string)
	r_rcs[rs_cpu] = "0.5"
	r_rcs[rs_mem] = "20G"

	var exec_handler = types.ExecAction{commands}
	//	path:="/mypath"
	//	sch :=types.URISchemeHTTPS
	//	var http_hader=types.HTTPHeader{&path,&path}
	//	var headerarray =[]types.HTTPHeader{http_hader}
	//	var http_handler=types.HTTPGetAction{&path,3553,nil,&sch,headerarray}

	//var tcp_handler=v1.TCPSocketAction{intstr.FromInt(23456),""}
	var prob_handler = types.Handler{"exec", &exec_handler, nil, nil}
	var re = int32(3)
	var liv_prob = types.Probe{&prob_handler, &re, nil, nil, nil, nil}

	var ds_atr = types.DockerServiceAttributes{Command: commands, Args: nil, LimitResources: lm_rcs, RequestResources: r_rcs, LivenessProb: &liv_prob}
	serviceatr := make(map[string]interface{})
	serviceatr["init_container"] = ds_atr
	serviceatr["label_selector"] = ls
	serviceatr["node_selector"] = ls_ml
	Jon := map[string]interface{}{"name": "cn1", "service_attributes": serviceatr}
	var v types.Service
	if d, err := json.Marshal(Jon); err != nil {
		fmt.Println(err)
	} else if err = json.Unmarshal(d, &v); err != nil {
		fmt.Println(err)
	} else {
		_, err := getStatefulSetObject(v)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("done ")
		}
	}
}
