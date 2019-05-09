package rbac

import (
	"istio-service-mesh/types"
	rbacV1 "k8s.io/api/rbac/v1"
)

func ProvisionRole(role types.Role) rbacV1.Role {
	roleObj := rbacV1.Role{}
	roleObj.Namespace = role.Namespace
	roleObj.Name = "sa-" + role.ServiceName + "-role"

	if role.ApiGroup == "" {
		role.ApiGroup = rbacV1.APIGroupAll
	}

	rule := rbacV1.PolicyRule{APIGroups: []string{role.ApiGroup},
		Resources: []string{role.Resource},
		Verbs:     role.Verbs}
	roleObj.Rules = []rbacV1.PolicyRule{rule}

	return roleObj
}
