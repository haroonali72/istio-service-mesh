package constants

type K8sKind string
type MeshKind string

const (
	Deployment            K8sKind = "Deployment"
	StatefulSet           K8sKind = "StatefulSet"
	DaemonSet             K8sKind = "DaemonSet"
	Job                   K8sKind = "Job"
	CronJob               K8sKind = "CronJob"
	Service               K8sKind = "Service"
	ConfigMap             K8sKind = "ConfigMap"
	Secret                K8sKind = "Secret"
	ServiceAccount        K8sKind = "ServiceAccount"
	Role                  K8sKind = "Role"
	RoleBinding           K8sKind = "RoleBinding"
	ClusterRole           K8sKind = "ClusterRole"
	ClusterRoleBinding    K8sKind = "ClusterRoleBinding"
	PersistentVolume      K8sKind = "PersistentVolume"
	PersistentVolumeClaim K8sKind = "PersistentVolumeClaim"
	StorageClass          K8sKind = "StorageClass"
	HPA                   K8sKind = "HorizontalPodAutoscaler"

	Gateway         MeshKind = "Gateway"
	VirtualService  MeshKind = "VirtualService"
	DestinationRule MeshKind = "DestinationRule"
	MeshPolicy      MeshKind = "Policy"
	ServiceEntry    MeshKind = "ServiceEntr"
)

var (
	k8sTypes = []K8sKind{
		Deployment,
		StatefulSet,
		DaemonSet,
		Job,
		CronJob,
		Service,
		ConfigMap,
		ServiceAccount,
		RoleBinding,
		Role,
		ClusterRoleBinding,
		ClusterRole,
		PersistentVolume,
		PersistentVolumeClaim,
		StorageClass,
		Secret,
	}
)

func (k K8sKind) String() string {
	return string(k)
}
func (m MeshKind) String() string {
	return string(m)
}

func IsSupportK8sType(s K8sKind) bool {
	for i := range k8sTypes {
		if k8sTypes[i] == s {

			return true
		}
	}
	return false
}
