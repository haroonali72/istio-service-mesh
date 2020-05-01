package services

import (
	"bitbucket.org/cloudplex-devs/microservices-mesh-engine/types"
)

//Id                interface{}              `json:"_id,omitempty" bson:"_id" valid:"-"`
//	ServiceId         string                   `json:"service_id" bson:"service_id" binding:"required" valid:"alphanumspecial,length(4|30)~service_id must contain between 6 and 30 characters,lowercase~lowercase alphanumeric characters are allowed,required~service_id is missing in request"`
//	Name              string                   `json:"name"  bson:"name" binding:"required" valid:"alphanumspecial,length(3|30),lowercase~lowercase alphanumeric characters are allowed,required"`
//	Version           string                   `json:"version"  bson:"version"  binding:"required" valid:"alphanumspecial,length(1|10),lowercase~lowercase alphanumeric characters are allowed,required"`
//	ServiceType       constants.ServiceType    `json:"service_type"  bson:"service_type" valid:"-"`
//	ServiceSubType    constants.ServiceSubType `json:"service_sub_type" bson:"service_type" valid:"-"`
//	Namespace         string                   `json:"namespace" bson:"namespace" binding:"required" valid:"alphanumspecial,required"`
//	CompanyId         string                   `json:"company_id,omitempty" bson:"company_id"`
//	CreationDate      time.Time                `json:"creation_date,omitempty" bson:"creation_date" valid:"-"`
type ServiceEntry struct {
	types.ServiceBasicInfo `json:",inline" bson:",inline"`
	ServiceAttributes      *ServiceEntryAttributes `json:"service_attributes"  bson:"company_id" binding:"required"`
}

type ServiceEntryAttributes struct {
	Hosts           []string                `json:"hosts" bson:"hosts" binding:"required"`
	Addresses       []string                `json:"addresses" bson:"addresses"`
	Location        Location                `json:"location" bson:"location"`
	Resolution      Resolution              `json:"resolution" bson:"resolution" binding:"required"`
	Ports           []*ServiceEntryPort     `json:"ports" bson:"ports" binding:"required"`
	Endpoints       []*ServiceEntryEndpoint `json:"endpoints" bson:"endpoints"`
	ExportTo        []string                `json:"export_to" bson:"export_to"`
	SubjectAltNames []string                `json:"subject_alt_names" bson:"subject_alt_names"`
}

type ServiceEntryPort struct {
	Name     string `json:"name" bson:"name"`
	Number   uint32 `json:"number" bson:"number" binding:"required"`
	Protocol string `json:"protocol" bson:"protocol" binding:"required"`
}

type ServiceEntryEndpoint struct {
	Address  string            `json:"address" bson:"address" binding:"required" `
	Ports    map[string]uint32 `json:"ports" bson:"ports"`
	Labels   map[string]string `json:"labels" bson:"labels"`
	Network  string            `json:"network" bson:"network"`
	Weight   string            `json:"weight" bson:"weight"`
	Locality string            `json:"locality" bson:"locality"`
}

type Location string

const Location_MESH_EXTERNAL Location = "MESH_EXTERNAL"
const Location_MESH_INTERNAL Location = "MESH_INTERNAL"

type Resolution string

const Resolution_NONE Resolution = "NONE"
const Resolution_STATIC Resolution = "STATIC"
const Resolution_DNS Resolution = "DNS"
