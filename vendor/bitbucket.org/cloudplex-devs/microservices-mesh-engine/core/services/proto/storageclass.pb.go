// Code generated by protoc-gen-go. DO NOT EDIT.
// source: storageclass.proto

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

type VolumeBindingMode int32

const (
	VolumeBindingMode_Immediate            VolumeBindingMode = 0
	VolumeBindingMode_WaitForFirstConsumer VolumeBindingMode = 1
)

var VolumeBindingMode_name = map[int32]string{
	0: "Immediate",
	1: "WaitForFirstConsumer",
}

var VolumeBindingMode_value = map[string]int32{
	"Immediate":            0,
	"WaitForFirstConsumer": 1,
}

func (x VolumeBindingMode) String() string {
	return proto.EnumName(VolumeBindingMode_name, int32(x))
}

func (VolumeBindingMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a3943d2b4f3af7d5, []int{0}
}

type ReclaimPolicy int32

const (
	ReclaimPolicy_Retain ReclaimPolicy = 0
	ReclaimPolicy_Delete ReclaimPolicy = 1
)

var ReclaimPolicy_name = map[int32]string{
	0: "Retain",
	1: "Delete",
}

var ReclaimPolicy_value = map[string]int32{
	"Retain": 0,
	"Delete": 1,
}

func (x ReclaimPolicy) String() string {
	return proto.EnumName(ReclaimPolicy_name, int32(x))
}

func (ReclaimPolicy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a3943d2b4f3af7d5, []int{1}
}

type StorageClassService struct {
	ProjectId            string                  `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	ServiceId            string                  `protobuf:"bytes,2,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	Name                 string                  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Version              string                  `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	ServiceType          string                  `protobuf:"bytes,5,opt,name=service_type,json=serviceType,proto3" json:"service_type,omitempty"`
	ServiceSubType       string                  `protobuf:"bytes,6,opt,name=service_sub_type,json=serviceSubType,proto3" json:"service_sub_type,omitempty"`
	Token                string                  `protobuf:"bytes,7,opt,name=token,proto3" json:"token,omitempty"`
	CompanyId            string                  `protobuf:"bytes,8,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	ServiceAttributes    *StorageClassAttributes `protobuf:"bytes,9,opt,name=service_attributes,json=serviceAttributes,proto3" json:"service_attributes,omitempty"`
	HookConfiguration    *HookConfiguration      `protobuf:"bytes,12,opt,name=hook_configuration,json=hookConfiguration,proto3" json:"hook_configuration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *StorageClassService) Reset()         { *m = StorageClassService{} }
func (m *StorageClassService) String() string { return proto.CompactTextString(m) }
func (*StorageClassService) ProtoMessage()    {}
func (*StorageClassService) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3943d2b4f3af7d5, []int{0}
}

func (m *StorageClassService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StorageClassService.Unmarshal(m, b)
}
func (m *StorageClassService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StorageClassService.Marshal(b, m, deterministic)
}
func (m *StorageClassService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StorageClassService.Merge(m, src)
}
func (m *StorageClassService) XXX_Size() int {
	return xxx_messageInfo_StorageClassService.Size(m)
}
func (m *StorageClassService) XXX_DiscardUnknown() {
	xxx_messageInfo_StorageClassService.DiscardUnknown(m)
}

var xxx_messageInfo_StorageClassService proto.InternalMessageInfo

func (m *StorageClassService) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *StorageClassService) GetServiceId() string {
	if m != nil {
		return m.ServiceId
	}
	return ""
}

func (m *StorageClassService) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StorageClassService) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *StorageClassService) GetServiceType() string {
	if m != nil {
		return m.ServiceType
	}
	return ""
}

func (m *StorageClassService) GetServiceSubType() string {
	if m != nil {
		return m.ServiceSubType
	}
	return ""
}

func (m *StorageClassService) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *StorageClassService) GetCompanyId() string {
	if m != nil {
		return m.CompanyId
	}
	return ""
}

func (m *StorageClassService) GetServiceAttributes() *StorageClassAttributes {
	if m != nil {
		return m.ServiceAttributes
	}
	return nil
}

func (m *StorageClassService) GetHookConfiguration() *HookConfiguration {
	if m != nil {
		return m.HookConfiguration
	}
	return nil
}

type StorageClassServiceResponse struct {
	Error                string               `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Resp                 *StorageClassService `protobuf:"bytes,2,opt,name=resp,proto3" json:"resp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *StorageClassServiceResponse) Reset()         { *m = StorageClassServiceResponse{} }
func (m *StorageClassServiceResponse) String() string { return proto.CompactTextString(m) }
func (*StorageClassServiceResponse) ProtoMessage()    {}
func (*StorageClassServiceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3943d2b4f3af7d5, []int{1}
}

func (m *StorageClassServiceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StorageClassServiceResponse.Unmarshal(m, b)
}
func (m *StorageClassServiceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StorageClassServiceResponse.Marshal(b, m, deterministic)
}
func (m *StorageClassServiceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StorageClassServiceResponse.Merge(m, src)
}
func (m *StorageClassServiceResponse) XXX_Size() int {
	return xxx_messageInfo_StorageClassServiceResponse.Size(m)
}
func (m *StorageClassServiceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StorageClassServiceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StorageClassServiceResponse proto.InternalMessageInfo

func (m *StorageClassServiceResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *StorageClassServiceResponse) GetResp() *StorageClassService {
	if m != nil {
		return m.Resp
	}
	return nil
}

type StorageClassAttributes struct {
	VolumeBindingMode    VolumeBindingMode       `protobuf:"varint,1,opt,name=volumeBindingMode,proto3,enum=proto.VolumeBindingMode" json:"volumeBindingMode,omitempty"`
	AllowVolumeExpansion string                  `protobuf:"bytes,2,opt,name=allowVolumeExpansion,proto3" json:"allowVolumeExpansion,omitempty"`
	Provisioner          string                  `protobuf:"bytes,3,opt,name=provisioner,proto3" json:"provisioner,omitempty"`
	Parameters           map[string]string       `protobuf:"bytes,4,rep,name=parameters,proto3" json:"parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ReclaimPolicy        ReclaimPolicy           `protobuf:"varint,5,opt,name=reclaimPolicy,proto3,enum=proto.ReclaimPolicy" json:"reclaimPolicy,omitempty"`
	MountOptions         []string                `protobuf:"bytes,6,rep,name=mountOptions,proto3" json:"mountOptions,omitempty"`
	AllowedTopologies    []*TopologySelectorTerm `protobuf:"bytes,7,rep,name=allowedTopologies,proto3" json:"allowedTopologies,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *StorageClassAttributes) Reset()         { *m = StorageClassAttributes{} }
func (m *StorageClassAttributes) String() string { return proto.CompactTextString(m) }
func (*StorageClassAttributes) ProtoMessage()    {}
func (*StorageClassAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3943d2b4f3af7d5, []int{2}
}

func (m *StorageClassAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StorageClassAttributes.Unmarshal(m, b)
}
func (m *StorageClassAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StorageClassAttributes.Marshal(b, m, deterministic)
}
func (m *StorageClassAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StorageClassAttributes.Merge(m, src)
}
func (m *StorageClassAttributes) XXX_Size() int {
	return xxx_messageInfo_StorageClassAttributes.Size(m)
}
func (m *StorageClassAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_StorageClassAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_StorageClassAttributes proto.InternalMessageInfo

func (m *StorageClassAttributes) GetVolumeBindingMode() VolumeBindingMode {
	if m != nil {
		return m.VolumeBindingMode
	}
	return VolumeBindingMode_Immediate
}

func (m *StorageClassAttributes) GetAllowVolumeExpansion() string {
	if m != nil {
		return m.AllowVolumeExpansion
	}
	return ""
}

func (m *StorageClassAttributes) GetProvisioner() string {
	if m != nil {
		return m.Provisioner
	}
	return ""
}

func (m *StorageClassAttributes) GetParameters() map[string]string {
	if m != nil {
		return m.Parameters
	}
	return nil
}

func (m *StorageClassAttributes) GetReclaimPolicy() ReclaimPolicy {
	if m != nil {
		return m.ReclaimPolicy
	}
	return ReclaimPolicy_Retain
}

func (m *StorageClassAttributes) GetMountOptions() []string {
	if m != nil {
		return m.MountOptions
	}
	return nil
}

func (m *StorageClassAttributes) GetAllowedTopologies() []*TopologySelectorTerm {
	if m != nil {
		return m.AllowedTopologies
	}
	return nil
}

type TopologySelectorTerm struct {
	MatchLabelExpressions []*TopologySelectorLabelRequirement `protobuf:"bytes,1,rep,name=matchLabelExpressions,proto3" json:"matchLabelExpressions,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                            `json:"-"`
	XXX_unrecognized      []byte                              `json:"-"`
	XXX_sizecache         int32                               `json:"-"`
}

func (m *TopologySelectorTerm) Reset()         { *m = TopologySelectorTerm{} }
func (m *TopologySelectorTerm) String() string { return proto.CompactTextString(m) }
func (*TopologySelectorTerm) ProtoMessage()    {}
func (*TopologySelectorTerm) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3943d2b4f3af7d5, []int{3}
}

func (m *TopologySelectorTerm) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopologySelectorTerm.Unmarshal(m, b)
}
func (m *TopologySelectorTerm) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopologySelectorTerm.Marshal(b, m, deterministic)
}
func (m *TopologySelectorTerm) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopologySelectorTerm.Merge(m, src)
}
func (m *TopologySelectorTerm) XXX_Size() int {
	return xxx_messageInfo_TopologySelectorTerm.Size(m)
}
func (m *TopologySelectorTerm) XXX_DiscardUnknown() {
	xxx_messageInfo_TopologySelectorTerm.DiscardUnknown(m)
}

var xxx_messageInfo_TopologySelectorTerm proto.InternalMessageInfo

func (m *TopologySelectorTerm) GetMatchLabelExpressions() []*TopologySelectorLabelRequirement {
	if m != nil {
		return m.MatchLabelExpressions
	}
	return nil
}

type TopologySelectorLabelRequirement struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Values               []string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TopologySelectorLabelRequirement) Reset()         { *m = TopologySelectorLabelRequirement{} }
func (m *TopologySelectorLabelRequirement) String() string { return proto.CompactTextString(m) }
func (*TopologySelectorLabelRequirement) ProtoMessage()    {}
func (*TopologySelectorLabelRequirement) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3943d2b4f3af7d5, []int{4}
}

func (m *TopologySelectorLabelRequirement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopologySelectorLabelRequirement.Unmarshal(m, b)
}
func (m *TopologySelectorLabelRequirement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopologySelectorLabelRequirement.Marshal(b, m, deterministic)
}
func (m *TopologySelectorLabelRequirement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopologySelectorLabelRequirement.Merge(m, src)
}
func (m *TopologySelectorLabelRequirement) XXX_Size() int {
	return xxx_messageInfo_TopologySelectorLabelRequirement.Size(m)
}
func (m *TopologySelectorLabelRequirement) XXX_DiscardUnknown() {
	xxx_messageInfo_TopologySelectorLabelRequirement.DiscardUnknown(m)
}

var xxx_messageInfo_TopologySelectorLabelRequirement proto.InternalMessageInfo

func (m *TopologySelectorLabelRequirement) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *TopologySelectorLabelRequirement) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

func init() {
	proto.RegisterEnum("proto.VolumeBindingMode", VolumeBindingMode_name, VolumeBindingMode_value)
	proto.RegisterEnum("proto.ReclaimPolicy", ReclaimPolicy_name, ReclaimPolicy_value)
	proto.RegisterType((*StorageClassService)(nil), "proto.StorageClassService")
	proto.RegisterType((*StorageClassServiceResponse)(nil), "proto.StorageClassServiceResponse")
	proto.RegisterType((*StorageClassAttributes)(nil), "proto.StorageClassAttributes")
	proto.RegisterMapType((map[string]string)(nil), "proto.StorageClassAttributes.ParametersEntry")
	proto.RegisterType((*TopologySelectorTerm)(nil), "proto.TopologySelectorTerm")
	proto.RegisterType((*TopologySelectorLabelRequirement)(nil), "proto.TopologySelectorLabelRequirement")
}

func init() { proto.RegisterFile("storageclass.proto", fileDescriptor_a3943d2b4f3af7d5) }

var fileDescriptor_a3943d2b4f3af7d5 = []byte{
	// 714 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x41, 0x6f, 0x12, 0x41,
	0x14, 0x2e, 0x85, 0x52, 0x79, 0x40, 0x0b, 0x23, 0x36, 0x1b, 0x9a, 0x26, 0xc8, 0xa5, 0xa4, 0x89,
	0x1c, 0xf0, 0x62, 0x1a, 0x3d, 0x28, 0xb6, 0x95, 0xa6, 0x8d, 0x64, 0x69, 0xf4, 0x64, 0x9a, 0x61,
	0xf7, 0x49, 0x47, 0x76, 0x67, 0xd6, 0x99, 0x59, 0x2c, 0x67, 0xff, 0x85, 0x7f, 0xd2, 0xbf, 0x60,
	0x76, 0x76, 0xa8, 0xd0, 0xae, 0xf5, 0xc2, 0x69, 0xf7, 0x7d, 0xdf, 0xf7, 0xbe, 0x79, 0x6f, 0xde,
	0x1b, 0x20, 0x4a, 0x0b, 0x49, 0x27, 0xe8, 0x05, 0x54, 0xa9, 0x6e, 0x24, 0x85, 0x16, 0x64, 0xcb,
	0x7c, 0x9a, 0x55, 0x85, 0x72, 0xc6, 0x3c, 0xec, 0xda, 0x70, 0x82, 0x1c, 0x25, 0x0d, 0xd2, 0xb0,
	0xfd, 0x2b, 0x0f, 0x4f, 0x47, 0x69, 0x6e, 0x3f, 0xc9, 0x1d, 0xa5, 0x62, 0x72, 0x00, 0x10, 0x49,
	0xf1, 0x0d, 0x3d, 0x7d, 0xcd, 0x7c, 0x27, 0xd7, 0xca, 0x75, 0x4a, 0x6e, 0xc9, 0x22, 0x03, 0x3f,
	0xa1, 0xad, 0x6d, 0x42, 0x6f, 0xa6, 0xb4, 0x45, 0x06, 0x3e, 0x21, 0x50, 0xe0, 0x34, 0x44, 0x27,
	0x6f, 0x08, 0xf3, 0x4f, 0x1c, 0xd8, 0x9e, 0xa1, 0x54, 0x4c, 0x70, 0xa7, 0x60, 0xe0, 0x45, 0x48,
	0x9e, 0x43, 0x65, 0x61, 0xa6, 0xe7, 0x11, 0x3a, 0x5b, 0x86, 0x2e, 0x5b, 0xec, 0x6a, 0x1e, 0x21,
	0xe9, 0x40, 0x6d, 0x21, 0x51, 0xf1, 0x38, 0x95, 0x15, 0x8d, 0x6c, 0xc7, 0xe2, 0xa3, 0x78, 0x6c,
	0x94, 0x0d, 0xd8, 0xd2, 0x62, 0x8a, 0xdc, 0xd9, 0x36, 0x74, 0x1a, 0x24, 0xf5, 0x7a, 0x22, 0x8c,
	0x28, 0x9f, 0x27, 0xf5, 0x3e, 0x49, 0xeb, 0xb5, 0xc8, 0xc0, 0x27, 0x17, 0x40, 0x16, 0xf6, 0x54,
	0x6b, 0xc9, 0xc6, 0xb1, 0x46, 0xe5, 0x94, 0x5a, 0xb9, 0x4e, 0xb9, 0x77, 0x90, 0xde, 0x54, 0x77,
	0xf9, 0x96, 0xde, 0xde, 0x89, 0xdc, 0xba, 0x4d, 0xfc, 0x0b, 0x91, 0x33, 0x20, 0x37, 0x42, 0x4c,
	0xaf, 0x3d, 0xc1, 0xbf, 0xb2, 0x49, 0x2c, 0xa9, 0x4e, 0x9a, 0xae, 0x18, 0x37, 0xc7, 0xba, 0x7d,
	0x10, 0x62, 0xda, 0x5f, 0xe6, 0xdd, 0xfa, 0xcd, 0x7d, 0xa8, 0xed, 0xc1, 0x7e, 0xc6, 0x6c, 0x5c,
	0x54, 0x91, 0xe0, 0xca, 0xb4, 0x8a, 0x52, 0x0a, 0x69, 0xc7, 0x93, 0x06, 0xa4, 0x0b, 0x05, 0x89,
	0x2a, 0x32, 0x43, 0x29, 0xf7, 0x9a, 0x19, 0xd5, 0x2f, 0x7c, 0x8c, 0xae, 0xfd, 0x3b, 0x0f, 0x7b,
	0xd9, 0xbd, 0x91, 0x53, 0xa8, 0xcf, 0x44, 0x10, 0x87, 0xf8, 0x8e, 0x71, 0x9f, 0xf1, 0xc9, 0xa5,
	0xf0, 0xd1, 0x1c, 0xb6, 0x73, 0xd7, 0xc7, 0xa7, 0xfb, 0xbc, 0xfb, 0x30, 0x85, 0xf4, 0xa0, 0x41,
	0x83, 0x40, 0xfc, 0x48, 0xc5, 0x27, 0xb7, 0x11, 0xe5, 0x66, 0x0f, 0xd2, 0xbd, 0xc9, 0xe4, 0x48,
	0x0b, 0xca, 0x91, 0x14, 0x33, 0x96, 0x04, 0x28, 0xed, 0x26, 0x2d, 0x43, 0xe4, 0x12, 0x20, 0xa2,
	0x92, 0x86, 0xa8, 0x51, 0x2a, 0xa7, 0xd0, 0xca, 0x77, 0xca, 0xbd, 0x17, 0x8f, 0x0e, 0xab, 0x3b,
	0xbc, 0xd3, 0x9f, 0x70, 0x2d, 0xe7, 0xee, 0x92, 0x01, 0x39, 0x86, 0xaa, 0x4c, 0xde, 0x0f, 0x0b,
	0x87, 0x22, 0x60, 0xde, 0xdc, 0xac, 0xe1, 0x4e, 0xaf, 0x61, 0x1d, 0xdd, 0x65, 0xce, 0x5d, 0x95,
	0x92, 0x36, 0x54, 0x42, 0x11, 0x73, 0xfd, 0x31, 0x4a, 0xe6, 0xa6, 0x9c, 0x62, 0x2b, 0xdf, 0x29,
	0xb9, 0x2b, 0x18, 0x19, 0x40, 0xdd, 0x34, 0x8a, 0xfe, 0x95, 0x88, 0x44, 0x20, 0x26, 0x0c, 0x95,
	0xb3, 0x6d, 0xaa, 0xde, 0xb7, 0x67, 0x58, 0x62, 0x3e, 0xc2, 0x00, 0x3d, 0x2d, 0xe4, 0x15, 0xca,
	0xd0, 0x7d, 0x98, 0xd5, 0x7c, 0x03, 0xbb, 0xf7, 0x3a, 0x21, 0x35, 0xc8, 0x4f, 0x71, 0x6e, 0x37,
	0x21, 0xf9, 0x4d, 0xb6, 0x63, 0x46, 0x83, 0x18, 0xed, 0x2d, 0xa7, 0xc1, 0xf1, 0xe6, 0xab, 0x5c,
	0x3b, 0x86, 0x46, 0xd6, 0x49, 0xe4, 0x0b, 0x3c, 0x0b, 0xa9, 0xf6, 0x6e, 0x2e, 0xe8, 0x18, 0x83,
	0x93, 0xdb, 0x48, 0xa2, 0x52, 0xa6, 0x9d, 0x9c, 0xa9, 0xf2, 0xf0, 0x1f, 0x55, 0x1a, 0xb9, 0x8b,
	0xdf, 0x63, 0x26, 0x31, 0x44, 0xae, 0xdd, 0x6c, 0x97, 0xf6, 0x05, 0xb4, 0xfe, 0x97, 0x9a, 0xd1,
	0xc6, 0x1e, 0x14, 0x4d, 0xe5, 0xca, 0xd9, 0x34, 0x97, 0x6a, 0xa3, 0xa3, 0xd7, 0x50, 0x7f, 0xb0,
	0x7b, 0xa4, 0x0a, 0xa5, 0x41, 0x18, 0xa2, 0xcf, 0xa8, 0xc6, 0xda, 0x06, 0x71, 0xa0, 0xf1, 0x99,
	0x32, 0x7d, 0x2a, 0xe4, 0x29, 0x93, 0x4a, 0xf7, 0x05, 0x57, 0x71, 0x88, 0xb2, 0x96, 0x3b, 0x3a,
	0x84, 0xea, 0xca, 0x40, 0x09, 0x40, 0xd1, 0x45, 0x4d, 0x19, 0xaf, 0x6d, 0x24, 0xff, 0xef, 0x31,
	0x40, 0x8d, 0xb5, 0x5c, 0xef, 0x67, 0x1e, 0x2a, 0xcb, 0xcb, 0x44, 0xce, 0x81, 0xf4, 0x25, 0x52,
	0x8d, 0x2b, 0xe8, 0x23, 0xcf, 0xac, 0xb9, 0xb7, 0xe0, 0x56, 0x9f, 0x6f, 0x7b, 0x83, 0x9c, 0xc1,
	0xee, 0x19, 0xea, 0x35, 0x18, 0x9d, 0x03, 0x49, 0x2b, 0x5e, 0x83, 0xd7, 0x00, 0xea, 0xc3, 0x64,
	0x7e, 0xeb, 0xe9, 0x6f, 0x18, 0xaf, 0xa1, 0xbf, 0x71, 0xd1, 0x10, 0x2f, 0xff, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x2f, 0x00, 0x18, 0x00, 0xe7, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StorageClassClient is the client API for StorageClass service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StorageClassClient interface {
	CreateStorageClass(ctx context.Context, in *StorageClassService, opts ...grpc.CallOption) (*ServiceResponse, error)
	GetStorageClass(ctx context.Context, in *StorageClassService, opts ...grpc.CallOption) (*ServiceResponse, error)
	DeleteStorageClass(ctx context.Context, in *StorageClassService, opts ...grpc.CallOption) (*ServiceResponse, error)
	PatchStorageClass(ctx context.Context, in *StorageClassService, opts ...grpc.CallOption) (*ServiceResponse, error)
	PutStorageClass(ctx context.Context, in *StorageClassService, opts ...grpc.CallOption) (*ServiceResponse, error)
}

type storageClassClient struct {
	cc *grpc.ClientConn
}

func NewStorageClassClient(cc *grpc.ClientConn) StorageClassClient {
	return &storageClassClient{cc}
}

func (c *storageClassClient) CreateStorageClass(ctx context.Context, in *StorageClassService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.StorageClass/CreateStorageClass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClassClient) GetStorageClass(ctx context.Context, in *StorageClassService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.StorageClass/GetStorageClass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClassClient) DeleteStorageClass(ctx context.Context, in *StorageClassService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.StorageClass/DeleteStorageClass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClassClient) PatchStorageClass(ctx context.Context, in *StorageClassService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.StorageClass/PatchStorageClass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClassClient) PutStorageClass(ctx context.Context, in *StorageClassService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.StorageClass/PutStorageClass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StorageClassServer is the server API for StorageClass service.
type StorageClassServer interface {
	CreateStorageClass(context.Context, *StorageClassService) (*ServiceResponse, error)
	GetStorageClass(context.Context, *StorageClassService) (*ServiceResponse, error)
	DeleteStorageClass(context.Context, *StorageClassService) (*ServiceResponse, error)
	PatchStorageClass(context.Context, *StorageClassService) (*ServiceResponse, error)
	PutStorageClass(context.Context, *StorageClassService) (*ServiceResponse, error)
}

// UnimplementedStorageClassServer can be embedded to have forward compatible implementations.
type UnimplementedStorageClassServer struct {
}

func (*UnimplementedStorageClassServer) CreateStorageClass(ctx context.Context, req *StorageClassService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStorageClass not implemented")
}
func (*UnimplementedStorageClassServer) GetStorageClass(ctx context.Context, req *StorageClassService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStorageClass not implemented")
}
func (*UnimplementedStorageClassServer) DeleteStorageClass(ctx context.Context, req *StorageClassService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStorageClass not implemented")
}
func (*UnimplementedStorageClassServer) PatchStorageClass(ctx context.Context, req *StorageClassService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchStorageClass not implemented")
}
func (*UnimplementedStorageClassServer) PutStorageClass(ctx context.Context, req *StorageClassService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutStorageClass not implemented")
}

func RegisterStorageClassServer(s *grpc.Server, srv StorageClassServer) {
	s.RegisterService(&_StorageClass_serviceDesc, srv)
}

func _StorageClass_CreateStorageClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StorageClassService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageClassServer).CreateStorageClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.StorageClass/CreateStorageClass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageClassServer).CreateStorageClass(ctx, req.(*StorageClassService))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageClass_GetStorageClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StorageClassService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageClassServer).GetStorageClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.StorageClass/GetStorageClass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageClassServer).GetStorageClass(ctx, req.(*StorageClassService))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageClass_DeleteStorageClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StorageClassService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageClassServer).DeleteStorageClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.StorageClass/DeleteStorageClass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageClassServer).DeleteStorageClass(ctx, req.(*StorageClassService))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageClass_PatchStorageClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StorageClassService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageClassServer).PatchStorageClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.StorageClass/PatchStorageClass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageClassServer).PatchStorageClass(ctx, req.(*StorageClassService))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageClass_PutStorageClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StorageClassService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageClassServer).PutStorageClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.StorageClass/PutStorageClass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageClassServer).PutStorageClass(ctx, req.(*StorageClassService))
	}
	return interceptor(ctx, in, info, handler)
}

var _StorageClass_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.StorageClass",
	HandlerType: (*StorageClassServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateStorageClass",
			Handler:    _StorageClass_CreateStorageClass_Handler,
		},
		{
			MethodName: "GetStorageClass",
			Handler:    _StorageClass_GetStorageClass_Handler,
		},
		{
			MethodName: "DeleteStorageClass",
			Handler:    _StorageClass_DeleteStorageClass_Handler,
		},
		{
			MethodName: "PatchStorageClass",
			Handler:    _StorageClass_PatchStorageClass_Handler,
		},
		{
			MethodName: "PutStorageClass",
			Handler:    _StorageClass_PutStorageClass_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "storageclass.proto",
}
