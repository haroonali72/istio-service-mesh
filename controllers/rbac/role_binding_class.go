package rbac

import (
	"istio-service-mesh/types"
	rbacV1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ProvisionRoleBinding(roleBinding types.RoleBinding) rbacV1.RoleBinding {
	rb := rbacV1.RoleBinding{
		ObjectMeta: metav1.ObjectMeta{Name: "sa-" + roleBinding.ServiceName + "-rolebinding"},
		Subjects: []rbacV1.Subject{
			{Kind: "ServiceAccount", Name: "sa-" + roleBinding.ServiceName},
		},
		RoleRef: rbacV1.RoleRef{Kind: "Role", Name: "sa-" + roleBinding.ServiceName + "-role"},
	}

	return rb
}
