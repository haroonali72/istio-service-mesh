package rbac

import (
	"istio-service-mesh/types"
	v1 "k8s.io/api/core/v1"
)

func ProvisionServiceAccount(serviceAccount types.ServiceAccount) v1.ServiceAccount {
	account := v1.ServiceAccount{}
	account.Name = "sa-" + serviceAccount.ServiceName
	account.Namespace = serviceAccount.Namespace
	return account
}
