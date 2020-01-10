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
	VolumeBindingMode_WaitForFirstCustomer VolumeBindingMode = 1
)

var VolumeBindingMode_name = map[int32]string{
	0: "Immediate",
	1: "WaitForFirstCustomer",
}

var VolumeBindingMode_value = map[string]int32{
	"Immediate":            0,
	"WaitForFirstCustomer": 1,
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
	VolumeBindingMode    VolumeBindingMode `protobuf:"varint,1,opt,name=volumeBindingMode,proto3,enum=proto.VolumeBindingMode" json:"volumeBindingMode,omitempty"`
	AllowVolumeExpansion bool              `protobuf:"varint,2,opt,name=allowVolumeExpansion,proto3" json:"allowVolumeExpansion,omitempty"`
	Provisioner          string            `protobuf:"bytes,3,opt,name=provisioner,proto3" json:"provisioner,omitempty"`
	ScParameters         *Parameters       `protobuf:"bytes,4,opt,name=scParameters,proto3" json:"scParameters,omitempty"`
	ReclaimPolicy        ReclaimPolicy     `protobuf:"varint,5,opt,name=reclaimPolicy,proto3,enum=proto.ReclaimPolicy" json:"reclaimPolicy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
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

func (m *StorageClassAttributes) GetAllowVolumeExpansion() bool {
	if m != nil {
		return m.AllowVolumeExpansion
	}
	return false
}

func (m *StorageClassAttributes) GetProvisioner() string {
	if m != nil {
		return m.Provisioner
	}
	return ""
}

func (m *StorageClassAttributes) GetScParameters() *Parameters {
	if m != nil {
		return m.ScParameters
	}
	return nil
}

func (m *StorageClassAttributes) GetReclaimPolicy() ReclaimPolicy {
	if m != nil {
		return m.ReclaimPolicy
	}
	return ReclaimPolicy_Retain
}

type Parameters struct {
	GcppdscParm          map[string]string `protobuf:"bytes,1,rep,name=gcppdscParm,proto3" json:"gcppdscParm,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	AwsebsscParm         map[string]string `protobuf:"bytes,2,rep,name=awsebsscParm,proto3" json:"awsebsscParm,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	AzurdiskscParm       map[string]string `protobuf:"bytes,3,rep,name=azurdiskscParm,proto3" json:"azurdiskscParm,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	AzurfilescParm       map[string]string `protobuf:"bytes,4,rep,name=azurfilescParm,proto3" json:"azurfilescParm,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Parameters) Reset()         { *m = Parameters{} }
func (m *Parameters) String() string { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()    {}
func (*Parameters) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3943d2b4f3af7d5, []int{3}
}

func (m *Parameters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Parameters.Unmarshal(m, b)
}
func (m *Parameters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Parameters.Marshal(b, m, deterministic)
}
func (m *Parameters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Parameters.Merge(m, src)
}
func (m *Parameters) XXX_Size() int {
	return xxx_messageInfo_Parameters.Size(m)
}
func (m *Parameters) XXX_DiscardUnknown() {
	xxx_messageInfo_Parameters.DiscardUnknown(m)
}

var xxx_messageInfo_Parameters proto.InternalMessageInfo

func (m *Parameters) GetGcppdscParm() map[string]string {
	if m != nil {
		return m.GcppdscParm
	}
	return nil
}

func (m *Parameters) GetAwsebsscParm() map[string]string {
	if m != nil {
		return m.AwsebsscParm
	}
	return nil
}

func (m *Parameters) GetAzurdiskscParm() map[string]string {
	if m != nil {
		return m.AzurdiskscParm
	}
	return nil
}

func (m *Parameters) GetAzurfilescParm() map[string]string {
	if m != nil {
		return m.AzurfilescParm
	}
	return nil
}

func init() {
	proto.RegisterEnum("proto.VolumeBindingMode", VolumeBindingMode_name, VolumeBindingMode_value)
	proto.RegisterEnum("proto.ReclaimPolicy", ReclaimPolicy_name, ReclaimPolicy_value)
	proto.RegisterType((*StorageClassService)(nil), "proto.StorageClassService")
	proto.RegisterType((*StorageClassServiceResponse)(nil), "proto.StorageClassServiceResponse")
	proto.RegisterType((*StorageClassAttributes)(nil), "proto.StorageClassAttributes")
	proto.RegisterType((*Parameters)(nil), "proto.Parameters")
	proto.RegisterMapType((map[string]string)(nil), "proto.Parameters.AwsebsscParmEntry")
	proto.RegisterMapType((map[string]string)(nil), "proto.Parameters.AzurdiskscParmEntry")
	proto.RegisterMapType((map[string]string)(nil), "proto.Parameters.AzurfilescParmEntry")
	proto.RegisterMapType((map[string]string)(nil), "proto.Parameters.GcppdscParmEntry")
}

func init() { proto.RegisterFile("storageclass.proto", fileDescriptor_a3943d2b4f3af7d5) }

var fileDescriptor_a3943d2b4f3af7d5 = []byte{
	// 677 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0xc7, 0xd7, 0x8f, 0x7d, 0xf4, 0xf4, 0x83, 0xd6, 0xab, 0xa6, 0xa8, 0x68, 0x52, 0x29, 0x42,
	0x54, 0xbb, 0xe8, 0x45, 0x10, 0x12, 0x9a, 0x10, 0x68, 0xec, 0xa3, 0xea, 0xc4, 0xa4, 0x2a, 0x43,
	0x70, 0x89, 0xdc, 0xe4, 0x30, 0xcc, 0x92, 0x38, 0xb2, 0x9d, 0x8e, 0x72, 0xcb, 0xd3, 0xf0, 0x16,
	0xbc, 0x02, 0x6f, 0x84, 0xe2, 0xb8, 0x34, 0x69, 0xab, 0x49, 0x43, 0xbd, 0x6a, 0xfc, 0xff, 0xff,
	0xfd, 0xcb, 0xf1, 0xe9, 0x89, 0x81, 0x48, 0xc5, 0x05, 0xbd, 0x41, 0xd7, 0xa7, 0x52, 0x0e, 0x22,
	0xc1, 0x15, 0x27, 0xdb, 0xfa, 0xa7, 0x53, 0x97, 0x28, 0xa6, 0xcc, 0xc5, 0x54, 0xed, 0xfd, 0x29,
	0xc2, 0xfe, 0x75, 0x1a, 0x3e, 0x4d, 0xc2, 0xd7, 0xa9, 0x4b, 0x0e, 0x01, 0x22, 0xc1, 0xbf, 0xa1,
	0xab, 0x3e, 0x33, 0xcf, 0x2a, 0x74, 0x0b, 0xfd, 0x8a, 0x53, 0x31, 0xca, 0xc8, 0x4b, 0x6c, 0xc3,
	0x49, 0xec, 0x62, 0x6a, 0x1b, 0x65, 0xe4, 0x11, 0x02, 0xe5, 0x90, 0x06, 0x68, 0x95, 0xb4, 0xa1,
	0x9f, 0x89, 0x05, 0xbb, 0x53, 0x14, 0x92, 0xf1, 0xd0, 0x2a, 0x6b, 0x79, 0xbe, 0x24, 0x4f, 0xa0,
	0x36, 0x87, 0xa9, 0x59, 0x84, 0xd6, 0xb6, 0xb6, 0xab, 0x46, 0xfb, 0x30, 0x8b, 0x90, 0xf4, 0xa1,
	0x39, 0x8f, 0xc8, 0x78, 0x92, 0xc6, 0x76, 0x74, 0xac, 0x61, 0xf4, 0xeb, 0x78, 0xa2, 0x93, 0x6d,
	0xd8, 0x56, 0xfc, 0x16, 0x43, 0x6b, 0x57, 0xdb, 0xe9, 0x22, 0xa9, 0xd7, 0xe5, 0x41, 0x44, 0xc3,
	0x59, 0x52, 0xef, 0x5e, 0x5a, 0xaf, 0x51, 0x46, 0x1e, 0x79, 0x0f, 0x64, 0x8e, 0xa7, 0x4a, 0x09,
	0x36, 0x89, 0x15, 0x4a, 0xab, 0xd2, 0x2d, 0xf4, 0xab, 0xf6, 0x61, 0xda, 0xa9, 0x41, 0xb6, 0x4b,
	0x27, 0xff, 0x42, 0x4e, 0xcb, 0x6c, 0x5c, 0x48, 0x3d, 0x17, 0x1e, 0xaf, 0x69, 0xa9, 0x83, 0x32,
	0xe2, 0xa1, 0xd4, 0x15, 0xa2, 0x10, 0x5c, 0x98, 0xae, 0xa6, 0x0b, 0x32, 0x80, 0xb2, 0x40, 0x19,
	0xe9, 0x5e, 0x56, 0xed, 0xce, 0x9a, 0x97, 0xce, 0x39, 0x3a, 0xd7, 0xfb, 0x55, 0x84, 0x83, 0xf5,
	0x25, 0x91, 0x0b, 0x68, 0x4d, 0xb9, 0x1f, 0x07, 0xf8, 0x8e, 0x85, 0x1e, 0x0b, 0x6f, 0xae, 0xb8,
	0x87, 0xfa, 0x65, 0x0d, 0xdb, 0x32, 0xdc, 0x8f, 0xcb, 0xbe, 0xb3, 0xba, 0x85, 0xd8, 0xd0, 0xa6,
	0xbe, 0xcf, 0xef, 0xd2, 0xf0, 0xf9, 0xf7, 0x88, 0x86, 0xfa, 0xef, 0x4b, 0x4a, 0xdc, 0x73, 0xd6,
	0x7a, 0xa4, 0x0b, 0xd5, 0x48, 0xf0, 0x29, 0x4b, 0x16, 0x28, 0xcc, 0x00, 0x64, 0x25, 0xf2, 0x12,
	0x6a, 0xd2, 0x1d, 0x53, 0x41, 0x03, 0x54, 0x28, 0xa4, 0x1e, 0x86, 0xaa, 0xdd, 0x32, 0x85, 0x2d,
	0x0c, 0x27, 0x17, 0x23, 0xc7, 0x50, 0x17, 0xc9, 0x3c, 0xb3, 0x60, 0xcc, 0x7d, 0xe6, 0xce, 0xf4,
	0x94, 0x34, 0xec, 0xb6, 0xd9, 0xe7, 0x64, 0x3d, 0x27, 0x1f, 0xed, 0xfd, 0x2e, 0x03, 0x64, 0x50,
	0x67, 0x50, 0xbd, 0x71, 0xa3, 0xc8, 0xd3, 0xfc, 0xc0, 0x2a, 0x74, 0x4b, 0xfd, 0xaa, 0xdd, 0x5b,
	0x29, 0x60, 0x30, 0x5c, 0x84, 0xce, 0x43, 0x25, 0x66, 0x4e, 0x76, 0x1b, 0x19, 0x42, 0x8d, 0xde,
	0x49, 0x9c, 0x48, 0x83, 0x29, 0x6a, 0xcc, 0xd3, 0x55, 0xcc, 0x49, 0x26, 0x95, 0x72, 0x72, 0x1b,
	0xc9, 0x15, 0x34, 0xe8, 0x8f, 0x58, 0x78, 0x4c, 0xde, 0x1a, 0x54, 0x49, 0xa3, 0x9e, 0xad, 0x41,
	0xe5, 0x72, 0x29, 0x6c, 0x69, 0xf3, 0x1c, 0xf7, 0x85, 0xf9, 0x68, 0x70, 0xe5, 0xfb, 0x70, 0x8b,
	0x5c, 0x06, 0xb7, 0x10, 0x3b, 0x6f, 0xa0, 0xb9, 0xdc, 0x07, 0xd2, 0x84, 0xd2, 0x2d, 0xce, 0xcc,
	0xfc, 0x26, 0x8f, 0xc9, 0x4c, 0x4f, 0xa9, 0x1f, 0xa3, 0xb9, 0x0a, 0xd2, 0xc5, 0x71, 0xf1, 0x55,
	0xa1, 0xf3, 0x16, 0x5a, 0x2b, 0x0d, 0x78, 0x10, 0xe0, 0x04, 0xf6, 0xd7, 0x1c, 0xfb, 0x7f, 0x10,
	0x4b, 0x47, 0x7d, 0x08, 0xe2, 0xe8, 0x35, 0xb4, 0x56, 0xbe, 0x19, 0x52, 0x87, 0xca, 0x28, 0x08,
	0xd0, 0x63, 0x54, 0x61, 0x73, 0x8b, 0x58, 0xd0, 0xfe, 0x44, 0x99, 0xba, 0xe0, 0xe2, 0x82, 0x09,
	0xa9, 0x4e, 0x63, 0xa9, 0x78, 0x80, 0xa2, 0x59, 0x38, 0x7a, 0x0e, 0xf5, 0xdc, 0x80, 0x12, 0x80,
	0x1d, 0x07, 0x15, 0x65, 0x61, 0x73, 0x2b, 0x79, 0x3e, 0x43, 0x1f, 0x15, 0x36, 0x0b, 0xf6, 0xcf,
	0x12, 0xd4, 0xb2, 0x5f, 0x35, 0xb9, 0x04, 0x72, 0x2a, 0x90, 0x2a, 0xcc, 0xa9, 0xf7, 0x5c, 0x0f,
	0x9d, 0x83, 0xb9, 0x97, 0xbf, 0x76, 0x7a, 0x5b, 0x64, 0x08, 0x8f, 0x86, 0xa8, 0x36, 0x00, 0xba,
	0x04, 0x92, 0x56, 0xbc, 0x01, 0xd6, 0x08, 0x5a, 0x63, 0xaa, 0xdc, 0xaf, 0x9b, 0x39, 0xdf, 0x38,
	0xde, 0xc0, 0xf9, 0x26, 0x3b, 0xda, 0x78, 0xf1, 0x37, 0x00, 0x00, 0xff, 0xff, 0x30, 0x7d, 0x1f,
	0x66, 0x47, 0x07, 0x00, 0x00,
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
