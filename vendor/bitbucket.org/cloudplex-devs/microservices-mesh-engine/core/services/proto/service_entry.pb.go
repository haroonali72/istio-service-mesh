// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service_entry.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Location int32

const (
	Location_MESH_EXTERNAL Location = 0
	Location_MESH_INTERNAL Location = 1
)

var Location_name = map[int32]string{
	0: "MESH_EXTERNAL",
	1: "MESH_INTERNAL",
}

var Location_value = map[string]int32{
	"MESH_EXTERNAL": 0,
	"MESH_INTERNAL": 1,
}

func (x Location) String() string {
	return proto.EnumName(Location_name, int32(x))
}

func (Location) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_805cd353d1f0cbc3, []int{0}
}

type Resolution int32

const (
	Resolution_NONE   Resolution = 0
	Resolution_STATIC Resolution = 1
	Resolution_DNS    Resolution = 3
)

var Resolution_name = map[int32]string{
	0: "NONE",
	1: "STATIC",
	3: "DNS",
}

var Resolution_value = map[string]int32{
	"NONE":   0,
	"STATIC": 1,
	"DNS":    3,
}

func (x Resolution) String() string {
	return proto.EnumName(Resolution_name, int32(x))
}

func (Resolution) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_805cd353d1f0cbc3, []int{1}
}

type ServiceEntryTemplate struct {
	ProjectId            string                  `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	ServiceId            string                  `protobuf:"bytes,2,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	Name                 string                  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Version              string                  `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	ServiceType          string                  `protobuf:"bytes,5,opt,name=service_type,json=serviceType,proto3" json:"service_type,omitempty"`
	ServiceSubType       string                  `protobuf:"bytes,6,opt,name=service_sub_type,json=serviceSubType,proto3" json:"service_sub_type,omitempty"`
	Namespace            string                  `protobuf:"bytes,7,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Token                string                  `protobuf:"bytes,8,opt,name=token,proto3" json:"token,omitempty"`
	CompanyId            string                  `protobuf:"bytes,9,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	ServiceAttributes    *ServiceEntryAttributes `protobuf:"bytes,10,opt,name=service_attributes,json=serviceAttributes,proto3" json:"service_attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ServiceEntryTemplate) Reset()         { *m = ServiceEntryTemplate{} }
func (m *ServiceEntryTemplate) String() string { return proto.CompactTextString(m) }
func (*ServiceEntryTemplate) ProtoMessage()    {}
func (*ServiceEntryTemplate) Descriptor() ([]byte, []int) {
	return fileDescriptor_805cd353d1f0cbc3, []int{0}
}

func (m *ServiceEntryTemplate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceEntryTemplate.Unmarshal(m, b)
}
func (m *ServiceEntryTemplate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceEntryTemplate.Marshal(b, m, deterministic)
}
func (m *ServiceEntryTemplate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceEntryTemplate.Merge(m, src)
}
func (m *ServiceEntryTemplate) XXX_Size() int {
	return xxx_messageInfo_ServiceEntryTemplate.Size(m)
}
func (m *ServiceEntryTemplate) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceEntryTemplate.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceEntryTemplate proto.InternalMessageInfo

func (m *ServiceEntryTemplate) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *ServiceEntryTemplate) GetServiceId() string {
	if m != nil {
		return m.ServiceId
	}
	return ""
}

func (m *ServiceEntryTemplate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ServiceEntryTemplate) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ServiceEntryTemplate) GetServiceType() string {
	if m != nil {
		return m.ServiceType
	}
	return ""
}

func (m *ServiceEntryTemplate) GetServiceSubType() string {
	if m != nil {
		return m.ServiceSubType
	}
	return ""
}

func (m *ServiceEntryTemplate) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ServiceEntryTemplate) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *ServiceEntryTemplate) GetCompanyId() string {
	if m != nil {
		return m.CompanyId
	}
	return ""
}

func (m *ServiceEntryTemplate) GetServiceAttributes() *ServiceEntryAttributes {
	if m != nil {
		return m.ServiceAttributes
	}
	return nil
}

type ServiceEntryAttributes struct {
	Hosts                []string                `protobuf:"bytes,1,rep,name=hosts,proto3" json:"hosts,omitempty"`
	Addresses            []string                `protobuf:"bytes,2,rep,name=addresses,proto3" json:"addresses,omitempty"`
	Ports                []*ServiceEntryPort     `protobuf:"bytes,3,rep,name=ports,proto3" json:"ports,omitempty"`
	Location             Location                `protobuf:"varint,4,opt,name=location,proto3,enum=proto.Location" json:"location,omitempty"`
	Resolution           Resolution              `protobuf:"varint,5,opt,name=resolution,proto3,enum=proto.Resolution" json:"resolution,omitempty"`
	Endpoints            []*ServiceEntryEndpoint `protobuf:"bytes,6,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	ExportTo             []string                `protobuf:"bytes,7,rep,name=exportTo,proto3" json:"exportTo,omitempty"`
	SubjectAltNames      []string                `protobuf:"bytes,8,rep,name=subject_alt_names,json=subjectAltNames,proto3" json:"subject_alt_names,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ServiceEntryAttributes) Reset()         { *m = ServiceEntryAttributes{} }
func (m *ServiceEntryAttributes) String() string { return proto.CompactTextString(m) }
func (*ServiceEntryAttributes) ProtoMessage()    {}
func (*ServiceEntryAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_805cd353d1f0cbc3, []int{1}
}

func (m *ServiceEntryAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceEntryAttributes.Unmarshal(m, b)
}
func (m *ServiceEntryAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceEntryAttributes.Marshal(b, m, deterministic)
}
func (m *ServiceEntryAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceEntryAttributes.Merge(m, src)
}
func (m *ServiceEntryAttributes) XXX_Size() int {
	return xxx_messageInfo_ServiceEntryAttributes.Size(m)
}
func (m *ServiceEntryAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceEntryAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceEntryAttributes proto.InternalMessageInfo

func (m *ServiceEntryAttributes) GetHosts() []string {
	if m != nil {
		return m.Hosts
	}
	return nil
}

func (m *ServiceEntryAttributes) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *ServiceEntryAttributes) GetPorts() []*ServiceEntryPort {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *ServiceEntryAttributes) GetLocation() Location {
	if m != nil {
		return m.Location
	}
	return Location_MESH_EXTERNAL
}

func (m *ServiceEntryAttributes) GetResolution() Resolution {
	if m != nil {
		return m.Resolution
	}
	return Resolution_NONE
}

func (m *ServiceEntryAttributes) GetEndpoints() []*ServiceEntryEndpoint {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

func (m *ServiceEntryAttributes) GetExportTo() []string {
	if m != nil {
		return m.ExportTo
	}
	return nil
}

func (m *ServiceEntryAttributes) GetSubjectAltNames() []string {
	if m != nil {
		return m.SubjectAltNames
	}
	return nil
}

type ServiceEntryPort struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Number               uint32   `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	Protocol             string   `protobuf:"bytes,3,opt,name=protocol,proto3" json:"protocol,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceEntryPort) Reset()         { *m = ServiceEntryPort{} }
func (m *ServiceEntryPort) String() string { return proto.CompactTextString(m) }
func (*ServiceEntryPort) ProtoMessage()    {}
func (*ServiceEntryPort) Descriptor() ([]byte, []int) {
	return fileDescriptor_805cd353d1f0cbc3, []int{2}
}

func (m *ServiceEntryPort) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceEntryPort.Unmarshal(m, b)
}
func (m *ServiceEntryPort) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceEntryPort.Marshal(b, m, deterministic)
}
func (m *ServiceEntryPort) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceEntryPort.Merge(m, src)
}
func (m *ServiceEntryPort) XXX_Size() int {
	return xxx_messageInfo_ServiceEntryPort.Size(m)
}
func (m *ServiceEntryPort) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceEntryPort.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceEntryPort proto.InternalMessageInfo

func (m *ServiceEntryPort) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ServiceEntryPort) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *ServiceEntryPort) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

type ServiceEntryEndpoint struct {
	Address              string            `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Ports                map[string]uint32 `protobuf:"bytes,2,rep,name=ports,proto3" json:"ports,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Labels               map[string]string `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Network              string            `protobuf:"bytes,4,opt,name=network,proto3" json:"network,omitempty"`
	Locality             string            `protobuf:"bytes,5,opt,name=locality,proto3" json:"locality,omitempty"`
	Weight               uint32            `protobuf:"varint,6,opt,name=weight,proto3" json:"weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ServiceEntryEndpoint) Reset()         { *m = ServiceEntryEndpoint{} }
func (m *ServiceEntryEndpoint) String() string { return proto.CompactTextString(m) }
func (*ServiceEntryEndpoint) ProtoMessage()    {}
func (*ServiceEntryEndpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_805cd353d1f0cbc3, []int{3}
}

func (m *ServiceEntryEndpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceEntryEndpoint.Unmarshal(m, b)
}
func (m *ServiceEntryEndpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceEntryEndpoint.Marshal(b, m, deterministic)
}
func (m *ServiceEntryEndpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceEntryEndpoint.Merge(m, src)
}
func (m *ServiceEntryEndpoint) XXX_Size() int {
	return xxx_messageInfo_ServiceEntryEndpoint.Size(m)
}
func (m *ServiceEntryEndpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceEntryEndpoint.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceEntryEndpoint proto.InternalMessageInfo

func (m *ServiceEntryEndpoint) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ServiceEntryEndpoint) GetPorts() map[string]uint32 {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *ServiceEntryEndpoint) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *ServiceEntryEndpoint) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func (m *ServiceEntryEndpoint) GetLocality() string {
	if m != nil {
		return m.Locality
	}
	return ""
}

func (m *ServiceEntryEndpoint) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func init() {
	proto.RegisterEnum("proto.Location", Location_name, Location_value)
	proto.RegisterEnum("proto.Resolution", Resolution_name, Resolution_value)
	proto.RegisterType((*ServiceEntryTemplate)(nil), "proto.ServiceEntryTemplate")
	proto.RegisterType((*ServiceEntryAttributes)(nil), "proto.ServiceEntryAttributes")
	proto.RegisterType((*ServiceEntryPort)(nil), "proto.ServiceEntryPort")
	proto.RegisterType((*ServiceEntryEndpoint)(nil), "proto.ServiceEntryEndpoint")
	proto.RegisterMapType((map[string]string)(nil), "proto.ServiceEntryEndpoint.LabelsEntry")
	proto.RegisterMapType((map[string]uint32)(nil), "proto.ServiceEntryEndpoint.PortsEntry")
}

func init() {
	proto.RegisterFile("service_entry.proto", fileDescriptor_805cd353d1f0cbc3)
}

var fileDescriptor_805cd353d1f0cbc3 = []byte{
	// 709 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcd, 0x6e, 0xda, 0x4a,
	0x14, 0x8e, 0x71, 0xf8, 0x3b, 0x84, 0x04, 0xe6, 0x46, 0xb9, 0x23, 0xee, 0x8d, 0xc4, 0x65, 0x71,
	0x8b, 0x12, 0x35, 0x6a, 0xe9, 0x26, 0xa9, 0x2a, 0x55, 0x28, 0x41, 0x0d, 0x0d, 0xa5, 0x91, 0x61,
	0x51, 0x75, 0x83, 0x0c, 0x3e, 0x6a, 0xdc, 0x18, 0x8f, 0xe5, 0x19, 0x27, 0xe5, 0x0d, 0xfa, 0x94,
	0x7d, 0x86, 0xaa, 0x4f, 0x50, 0xcd, 0x78, 0xc6, 0x90, 0x08, 0xa5, 0x1b, 0x56, 0xf8, 0x7c, 0xdf,
	0x77, 0x66, 0xbe, 0xf3, 0x33, 0xc0, 0x5f, 0x1c, 0xe3, 0x3b, 0x7f, 0x86, 0x13, 0x0c, 0x45, 0xbc,
	0x38, 0x89, 0x62, 0x26, 0x18, 0xc9, 0xab, 0x9f, 0x46, 0x55, 0x73, 0x29, 0xda, 0xfa, 0x99, 0x83,
	0xfd, 0x51, 0x8a, 0xf4, 0xa4, 0x78, 0x8c, 0xf3, 0x28, 0x70, 0x05, 0x92, 0x43, 0x80, 0x28, 0x66,
	0x5f, 0x71, 0x26, 0x26, 0xbe, 0x47, 0xad, 0xa6, 0xd5, 0x2e, 0x3b, 0x65, 0x8d, 0xf4, 0x3d, 0x49,
	0x9b, 0x4b, 0x7c, 0x8f, 0xe6, 0x52, 0x5a, 0x23, 0x7d, 0x8f, 0x10, 0xd8, 0x0e, 0xdd, 0x39, 0x52,
	0x5b, 0x11, 0xea, 0x9b, 0x50, 0x28, 0xde, 0x61, 0xcc, 0x7d, 0x16, 0xd2, 0x6d, 0x05, 0x9b, 0x90,
	0xfc, 0x07, 0x3b, 0xe6, 0x30, 0xb1, 0x88, 0x90, 0xe6, 0x15, 0x5d, 0xd1, 0xd8, 0x78, 0x11, 0x21,
	0x69, 0x43, 0xcd, 0x48, 0x78, 0x32, 0x4d, 0x65, 0x05, 0x25, 0xdb, 0xd5, 0xf8, 0x28, 0x99, 0x2a,
	0xe5, 0xbf, 0x50, 0x96, 0xd7, 0xf1, 0xc8, 0x9d, 0x21, 0x2d, 0xa6, 0xc6, 0x32, 0x80, 0xec, 0x43,
	0x5e, 0xb0, 0x5b, 0x0c, 0x69, 0x49, 0x31, 0x69, 0x20, 0xab, 0x99, 0xb1, 0x79, 0xe4, 0x86, 0x0b,
	0x59, 0x4d, 0x39, 0x4d, 0xd2, 0x48, 0xdf, 0x23, 0x03, 0x20, 0xe6, 0x72, 0x57, 0x88, 0xd8, 0x9f,
	0x26, 0x02, 0x39, 0x85, 0xa6, 0xd5, 0xae, 0x74, 0x0e, 0xd3, 0x46, 0x9e, 0xac, 0x36, 0xb1, 0x9b,
	0x89, 0x9c, 0xba, 0x4e, 0x5c, 0x42, 0xad, 0x1f, 0x39, 0x38, 0x58, 0xaf, 0x96, 0xee, 0x6e, 0x18,
	0x17, 0x9c, 0x5a, 0x4d, 0x5b, 0xba, 0x53, 0x81, 0xac, 0xc8, 0xf5, 0xbc, 0x18, 0x39, 0x47, 0x4e,
	0x73, 0x8a, 0x59, 0x02, 0xe4, 0x39, 0xe4, 0x23, 0x16, 0x0b, 0x4e, 0xed, 0xa6, 0xdd, 0xae, 0x74,
	0xfe, 0x5e, 0xe3, 0xe7, 0x9a, 0xc5, 0xc2, 0x49, 0x55, 0xe4, 0x18, 0x4a, 0x01, 0x9b, 0xb9, 0xc2,
	0x8c, 0x61, 0xb7, 0xb3, 0xa7, 0x33, 0x06, 0x1a, 0x76, 0x32, 0x01, 0x79, 0x09, 0x10, 0x23, 0x67,
	0x41, 0xa2, 0xe4, 0x79, 0x25, 0xaf, 0x6b, 0xb9, 0x93, 0x11, 0xce, 0x8a, 0x88, 0x9c, 0x41, 0x19,
	0x43, 0x2f, 0x62, 0x7e, 0x28, 0x38, 0x2d, 0x28, 0x4b, 0xff, 0xac, 0xb1, 0xd4, 0xd3, 0x1a, 0x67,
	0xa9, 0x26, 0x0d, 0x28, 0xe1, 0x37, 0xe9, 0x72, 0xcc, 0x68, 0x51, 0x95, 0x99, 0xc5, 0xe4, 0x08,
	0xea, 0x3c, 0x99, 0xaa, 0x75, 0x74, 0x03, 0x31, 0x51, 0x03, 0xa5, 0x25, 0x25, 0xda, 0xd3, 0x44,
	0x37, 0x10, 0x43, 0x09, 0xb7, 0x3e, 0x43, 0xed, 0x71, 0xf5, 0xd9, 0x42, 0x5a, 0x2b, 0x0b, 0x79,
	0x00, 0x85, 0x30, 0x99, 0x4f, 0x31, 0x56, 0xfb, 0x5b, 0x75, 0x74, 0x24, 0x7d, 0x28, 0xc3, 0x33,
	0x16, 0xe8, 0x05, 0xce, 0xe2, 0xd6, 0xaf, 0x47, 0xef, 0xc5, 0xd4, 0x21, 0xb7, 0x5b, 0xcf, 0x44,
	0xdf, 0x61, 0x42, 0xf2, 0xc6, 0x0c, 0x28, 0xa7, 0xba, 0xf1, 0xff, 0x13, 0xdd, 0x38, 0x91, 0x5e,
	0xb9, 0x82, 0xcc, 0xbc, 0xde, 0x42, 0x21, 0x70, 0xa7, 0x18, 0x98, 0xf9, 0x3e, 0x7b, 0x2a, 0x7d,
	0xa0, 0x94, 0x69, 0xbe, 0x4e, 0x93, 0xc6, 0x42, 0x14, 0xf7, 0x2c, 0xbe, 0x35, 0xcf, 0x4e, 0x87,
	0xb2, 0x4e, 0x39, 0xe9, 0xc0, 0x17, 0x0b, 0xfd, 0xe4, 0xb2, 0x58, 0xf6, 0xe6, 0x1e, 0xfd, 0x2f,
	0x37, 0x42, 0xbd, 0xb2, 0xaa, 0xa3, 0xa3, 0xc6, 0x29, 0xc0, 0xd2, 0x23, 0xa9, 0x81, 0x7d, 0x8b,
	0x0b, 0x5d, 0xb0, 0xfc, 0x94, 0x1b, 0x7c, 0xe7, 0x06, 0x09, 0xea, 0x96, 0xa6, 0xc1, 0xeb, 0xdc,
	0xa9, 0xd5, 0x38, 0x83, 0xca, 0x8a, 0xbd, 0x3f, 0xa5, 0x96, 0x57, 0x52, 0x8f, 0x5e, 0x40, 0xc9,
	0x2c, 0x27, 0xa9, 0x43, 0xf5, 0x43, 0x6f, 0x74, 0x39, 0xe9, 0x7d, 0x1a, 0xf7, 0x9c, 0x61, 0x77,
	0x50, 0xdb, 0xca, 0xa0, 0xfe, 0x50, 0x43, 0xd6, 0xd1, 0x31, 0xc0, 0x72, 0x3f, 0x49, 0x09, 0xb6,
	0x87, 0x1f, 0x87, 0xbd, 0xda, 0x16, 0x01, 0x28, 0x8c, 0xc6, 0xdd, 0x71, 0xff, 0xbc, 0x66, 0x91,
	0x22, 0xd8, 0x17, 0xc3, 0x51, 0xcd, 0xee, 0x7c, 0xb7, 0x61, 0x67, 0xb5, 0x9d, 0xe4, 0x0a, 0xc8,
	0x79, 0x8c, 0xae, 0xc0, 0x07, 0xe8, 0xba, 0x35, 0x36, 0x7f, 0x97, 0x8d, 0x83, 0x87, 0xa4, 0x83,
	0x3c, 0x62, 0x21, 0xc7, 0xd6, 0x16, 0xb9, 0x84, 0xbd, 0x77, 0x28, 0x36, 0x71, 0xd2, 0x15, 0x90,
	0x0b, 0x0c, 0x70, 0x33, 0xb6, 0xde, 0x43, 0xfd, 0xda, 0x15, 0xb3, 0x9b, 0x0d, 0x95, 0x78, 0x9d,
	0x6c, 0xa2, 0xc4, 0x69, 0x41, 0x11, 0xaf, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x97, 0x06, 0x6d,
	0xc1, 0xc2, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ServiceEntryClient is the client API for ServiceEntry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceEntryClient interface {
	CreateServiceEntry(ctx context.Context, in *ServiceEntryTemplate, opts ...grpc.CallOption) (*ServiceResponse, error)
	GetServiceEntry(ctx context.Context, in *ServiceEntryTemplate, opts ...grpc.CallOption) (*ServiceResponse, error)
	DeleteServiceEntry(ctx context.Context, in *ServiceEntryTemplate, opts ...grpc.CallOption) (*ServiceResponse, error)
	PatchServiceEntry(ctx context.Context, in *ServiceEntryTemplate, opts ...grpc.CallOption) (*ServiceResponse, error)
	PutServiceEntry(ctx context.Context, in *ServiceEntryTemplate, opts ...grpc.CallOption) (*ServiceResponse, error)
}

type serviceEntryClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceEntryClient(cc grpc.ClientConnInterface) ServiceEntryClient {
	return &serviceEntryClient{cc}
}

func (c *serviceEntryClient) CreateServiceEntry(ctx context.Context, in *ServiceEntryTemplate, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.ServiceEntry/CreateServiceEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceEntryClient) GetServiceEntry(ctx context.Context, in *ServiceEntryTemplate, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.ServiceEntry/GetServiceEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceEntryClient) DeleteServiceEntry(ctx context.Context, in *ServiceEntryTemplate, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.ServiceEntry/DeleteServiceEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceEntryClient) PatchServiceEntry(ctx context.Context, in *ServiceEntryTemplate, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.ServiceEntry/PatchServiceEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceEntryClient) PutServiceEntry(ctx context.Context, in *ServiceEntryTemplate, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.ServiceEntry/PutServiceEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceEntryServer is the server API for ServiceEntry service.
type ServiceEntryServer interface {
	CreateServiceEntry(context.Context, *ServiceEntryTemplate) (*ServiceResponse, error)
	GetServiceEntry(context.Context, *ServiceEntryTemplate) (*ServiceResponse, error)
	DeleteServiceEntry(context.Context, *ServiceEntryTemplate) (*ServiceResponse, error)
	PatchServiceEntry(context.Context, *ServiceEntryTemplate) (*ServiceResponse, error)
	PutServiceEntry(context.Context, *ServiceEntryTemplate) (*ServiceResponse, error)
}

// UnimplementedServiceEntryServer can be embedded to have forward compatible implementations.
type UnimplementedServiceEntryServer struct {
}

func (*UnimplementedServiceEntryServer) CreateServiceEntry(ctx context.Context, req *ServiceEntryTemplate) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateServiceEntry not implemented")
}
func (*UnimplementedServiceEntryServer) GetServiceEntry(ctx context.Context, req *ServiceEntryTemplate) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServiceEntry not implemented")
}
func (*UnimplementedServiceEntryServer) DeleteServiceEntry(ctx context.Context, req *ServiceEntryTemplate) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteServiceEntry not implemented")
}
func (*UnimplementedServiceEntryServer) PatchServiceEntry(ctx context.Context, req *ServiceEntryTemplate) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchServiceEntry not implemented")
}
func (*UnimplementedServiceEntryServer) PutServiceEntry(ctx context.Context, req *ServiceEntryTemplate) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutServiceEntry not implemented")
}

func RegisterServiceEntryServer(s *grpc.Server, srv ServiceEntryServer) {
	s.RegisterService(&_ServiceEntry_serviceDesc, srv)
}

func _ServiceEntry_CreateServiceEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceEntryTemplate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceEntryServer).CreateServiceEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ServiceEntry/CreateServiceEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceEntryServer).CreateServiceEntry(ctx, req.(*ServiceEntryTemplate))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceEntry_GetServiceEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceEntryTemplate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceEntryServer).GetServiceEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ServiceEntry/GetServiceEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceEntryServer).GetServiceEntry(ctx, req.(*ServiceEntryTemplate))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceEntry_DeleteServiceEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceEntryTemplate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceEntryServer).DeleteServiceEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ServiceEntry/DeleteServiceEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceEntryServer).DeleteServiceEntry(ctx, req.(*ServiceEntryTemplate))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceEntry_PatchServiceEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceEntryTemplate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceEntryServer).PatchServiceEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ServiceEntry/PatchServiceEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceEntryServer).PatchServiceEntry(ctx, req.(*ServiceEntryTemplate))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceEntry_PutServiceEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceEntryTemplate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceEntryServer).PutServiceEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ServiceEntry/PutServiceEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceEntryServer).PutServiceEntry(ctx, req.(*ServiceEntryTemplate))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServiceEntry_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ServiceEntry",
	HandlerType: (*ServiceEntryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateServiceEntry",
			Handler:    _ServiceEntry_CreateServiceEntry_Handler,
		},
		{
			MethodName: "GetServiceEntry",
			Handler:    _ServiceEntry_GetServiceEntry_Handler,
		},
		{
			MethodName: "DeleteServiceEntry",
			Handler:    _ServiceEntry_DeleteServiceEntry_Handler,
		},
		{
			MethodName: "PatchServiceEntry",
			Handler:    _ServiceEntry_PatchServiceEntry_Handler,
		},
		{
			MethodName: "PutServiceEntry",
			Handler:    _ServiceEntry_PutServiceEntry_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_entry.proto",
}
