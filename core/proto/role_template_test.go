package proto

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	services2 "microservices-mesh-engine/constants/services"
	"microservices-mesh-engine/models/mocks"
	"microservices-mesh-engine/types/services"
	"microservices-mesh-engine/utils"
	"os"
	"testing"
)

var role = `{
  "company_id": "5f4321wq",
  "service_id": "123qwerty",
  "name": "sample",
  "service_type": "k8s",
  "service_sub_type": "role",

  "service_dependency_info": [
    ""
  ],
  "namespace": "default",
  "service_attributes": {
    "rules": [{
      "name": ["abc", "xyz"],
      "verbs": [
        "get", "post"
      ],
      "api_group": [
        "xyz", "abc"
      ]
    }]
  }
}`

func init() {
	utils.LoggerInit(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
}

func TestCreateRole(t *testing.T) {
	db := mocks.Database{}
	svcManager := NewServiceManager(&db)
	service, err := sampleValidRole()
	if err != nil {
		t.Error(err)
		return
	}
	db.On("GetRoleTemplate", service.ServiceId, service.CompanyId).Return(nil, fmt.Errorf("data not found in database"))
	db.On("CreateRoleTemplate", &service).Return(&service, nil)
	_, err = svcManager.CreateRoleTemplate(&service)
	if err != nil {
		t.Error(err)
		return
	}
}

func sampleValidRole() (svc services.Role, err error) {
	err = json.Unmarshal([]byte(services2.serviceAccount), &svc)
	if err != nil {
		utils.Error.Println(err)
		return svc, err
	}

	return svc, nil
}
