package types

import (
	v12 "k8s.io/api/apps/v1"
	autoscaling "k8s.io/api/autoscaling/v2beta2"
	v13 "k8s.io/api/batch/v1"
	"k8s.io/api/batch/v2alpha1"
	v1 "k8s.io/api/core/v1"
)

type KSDResponse struct {
	Service OutputStruct `json:"service"`
}
type OutputStruct struct {
	Deployments  []DeploymentResp  `json:"deployment"`
	Kubernetes   []KubernetesResp  `json:"kubernetes-service"`
	Istio        []IstioResp       `json:"istio-component"`
	Secrets      []interface{}     `json:"secrets"`
	Nodes        []NodeResp        `json:"nodes"`
	StatefulSets []StatefulSetResp `json:"statefulset"`
	DaemonSets   []DaemonSetResp   `json:"daemonsets"`
	Jobs         []JobResp         `json:"jobs"`
	CronJobs     []CronJobResp     `json:"cronjobs"`
	HPAS         []HpasResp        `json:"hpas"`
}
type HpasResp struct {
	Error string                               `json:"error"`
	Hpas  *autoscaling.HorizontalPodAutoscaler `json:"data"`
}
type StatefulSetResp struct {
	Error        string           `json:"error"`
	StatefulSets *v12.StatefulSet `json:"data"`
}
type DaemonSetResp struct {
	Error      string         `json:"error"`
	DaemonSets *v12.DaemonSet `json:"data"`
}
type JobResp struct {
	Error string   `json:"error"`
	Jobs  *v13.Job `json:"data"`
}
type CronJobResp struct {
	Error    string            `json:"error"`
	CronJobs *v2alpha1.CronJob `json:"data"`
}

type NodeResp struct {
	Error string       `json:"error"`
	Nodes *v1.NodeList `json:"data"`
}
type DeploymentResp struct {
	Error       string          `json:"error"`
	Deployments *v12.Deployment `json:"data"`
}
type KubernetesResp struct {
	Error      string      `json:"error"`
	Kubernetes *v1.Service `json:"data"`
}
type IstioResp struct {
	Error string       `json:"error"`
	Istio *IstioObject `json:"data"`
}
