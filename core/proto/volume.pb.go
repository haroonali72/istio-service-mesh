// Code generated by protoc-gen-go. DO NOT EDIT.
// source: volume.proto

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

type Volume struct {
	ProjectId             string      `protobuf:"bytes,1,opt,name=ProjectId,proto3" json:"ProjectId,omitempty"`
	ServiceId             string      `protobuf:"bytes,2,opt,name=ServiceId,proto3" json:"ServiceId,omitempty"`
	Name                  string      `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	ServiceType           string      `protobuf:"bytes,4,opt,name=ServiceType,proto3" json:"ServiceType,omitempty"`
	ServiceSubType        string      `protobuf:"bytes,5,opt,name=ServiceSubType,proto3" json:"ServiceSubType,omitempty"`
	Status                string      `protobuf:"bytes,6,opt,name=Status,proto3" json:"Status,omitempty"`
	Token                 string      `protobuf:"bytes,7,opt,name=token,proto3" json:"token,omitempty"`
	ServiceDependencyInfo []string    `protobuf:"bytes,8,rep,name=ServiceDependencyInfo,proto3" json:"ServiceDependencyInfo,omitempty"`
	ServiceAttributes     *VolumeAttr `protobuf:"bytes,9,opt,name=ServiceAttributes,proto3" json:"ServiceAttributes,omitempty"`
	CompanyId             string      `protobuf:"bytes,10,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}    `json:"-"`
	XXX_unrecognized      []byte      `json:"-"`
	XXX_sizecache         int32       `json:"-"`
}

func (m *Volume) Reset()         { *m = Volume{} }
func (m *Volume) String() string { return proto.CompactTextString(m) }
func (*Volume) ProtoMessage()    {}
func (*Volume) Descriptor() ([]byte, []int) {
	return fileDescriptor_498b213ad3bcd5ad, []int{0}
}

func (m *Volume) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Volume.Unmarshal(m, b)
}
func (m *Volume) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Volume.Marshal(b, m, deterministic)
}
func (m *Volume) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Volume.Merge(m, src)
}
func (m *Volume) XXX_Size() int {
	return xxx_messageInfo_Volume.Size(m)
}
func (m *Volume) XXX_DiscardUnknown() {
	xxx_messageInfo_Volume.DiscardUnknown(m)
}

var xxx_messageInfo_Volume proto.InternalMessageInfo

func (m *Volume) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *Volume) GetServiceId() string {
	if m != nil {
		return m.ServiceId
	}
	return ""
}

func (m *Volume) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Volume) GetServiceType() string {
	if m != nil {
		return m.ServiceType
	}
	return ""
}

func (m *Volume) GetServiceSubType() string {
	if m != nil {
		return m.ServiceSubType
	}
	return ""
}

func (m *Volume) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Volume) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Volume) GetServiceDependencyInfo() []string {
	if m != nil {
		return m.ServiceDependencyInfo
	}
	return nil
}

func (m *Volume) GetServiceAttributes() *VolumeAttr {
	if m != nil {
		return m.ServiceAttributes
	}
	return nil
}

func (m *Volume) GetCompanyId() string {
	if m != nil {
		return m.CompanyId
	}
	return ""
}

type VolumeAttr struct {
	Size                 string     `protobuf:"bytes,1,opt,name=Size,proto3" json:"Size,omitempty"`
	Cloud                string     `protobuf:"bytes,2,opt,name=Cloud,proto3" json:"Cloud,omitempty"`
	MountPath            string     `protobuf:"bytes,3,opt,name=MountPath,proto3" json:"MountPath,omitempty"`
	Params               *Parameter `protobuf:"bytes,4,opt,name=Params,proto3" json:"Params,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *VolumeAttr) Reset()         { *m = VolumeAttr{} }
func (m *VolumeAttr) String() string { return proto.CompactTextString(m) }
func (*VolumeAttr) ProtoMessage()    {}
func (*VolumeAttr) Descriptor() ([]byte, []int) {
	return fileDescriptor_498b213ad3bcd5ad, []int{1}
}

func (m *VolumeAttr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VolumeAttr.Unmarshal(m, b)
}
func (m *VolumeAttr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VolumeAttr.Marshal(b, m, deterministic)
}
func (m *VolumeAttr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VolumeAttr.Merge(m, src)
}
func (m *VolumeAttr) XXX_Size() int {
	return xxx_messageInfo_VolumeAttr.Size(m)
}
func (m *VolumeAttr) XXX_DiscardUnknown() {
	xxx_messageInfo_VolumeAttr.DiscardUnknown(m)
}

var xxx_messageInfo_VolumeAttr proto.InternalMessageInfo

func (m *VolumeAttr) GetSize() string {
	if m != nil {
		return m.Size
	}
	return ""
}

func (m *VolumeAttr) GetCloud() string {
	if m != nil {
		return m.Cloud
	}
	return ""
}

func (m *VolumeAttr) GetMountPath() string {
	if m != nil {
		return m.MountPath
	}
	return ""
}

func (m *VolumeAttr) GetParams() *Parameter {
	if m != nil {
		return m.Params
	}
	return nil
}

type Parameter struct {
	Type                 string   `protobuf:"bytes,1,opt,name=Type,proto3" json:"Type,omitempty"`
	ReplicationType      string   `protobuf:"bytes,2,opt,name=ReplicationType,proto3" json:"ReplicationType,omitempty"`
	IOPS                 string   `protobuf:"bytes,3,opt,name=IOPS,proto3" json:"IOPS,omitempty"`
	Plugin               string   `protobuf:"bytes,4,opt,name=Plugin,proto3" json:"Plugin,omitempty"`
	SkuName              string   `protobuf:"bytes,5,opt,name=SkuName,proto3" json:"SkuName,omitempty"`
	Location             string   `protobuf:"bytes,6,opt,name=Location,proto3" json:"Location,omitempty"`
	StorageAccount       string   `protobuf:"bytes,7,opt,name=StorageAccount,proto3" json:"StorageAccount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Parameter) Reset()         { *m = Parameter{} }
func (m *Parameter) String() string { return proto.CompactTextString(m) }
func (*Parameter) ProtoMessage()    {}
func (*Parameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_498b213ad3bcd5ad, []int{2}
}

func (m *Parameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Parameter.Unmarshal(m, b)
}
func (m *Parameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Parameter.Marshal(b, m, deterministic)
}
func (m *Parameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Parameter.Merge(m, src)
}
func (m *Parameter) XXX_Size() int {
	return xxx_messageInfo_Parameter.Size(m)
}
func (m *Parameter) XXX_DiscardUnknown() {
	xxx_messageInfo_Parameter.DiscardUnknown(m)
}

var xxx_messageInfo_Parameter proto.InternalMessageInfo

func (m *Parameter) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Parameter) GetReplicationType() string {
	if m != nil {
		return m.ReplicationType
	}
	return ""
}

func (m *Parameter) GetIOPS() string {
	if m != nil {
		return m.IOPS
	}
	return ""
}

func (m *Parameter) GetPlugin() string {
	if m != nil {
		return m.Plugin
	}
	return ""
}

func (m *Parameter) GetSkuName() string {
	if m != nil {
		return m.SkuName
	}
	return ""
}

func (m *Parameter) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *Parameter) GetStorageAccount() string {
	if m != nil {
		return m.StorageAccount
	}
	return ""
}

func init() {
	proto.RegisterType((*Volume)(nil), "proto.Volume")
	proto.RegisterType((*VolumeAttr)(nil), "proto.VolumeAttr")
	proto.RegisterType((*Parameter)(nil), "proto.Parameter")
}

func init() { proto.RegisterFile("volume.proto", fileDescriptor_498b213ad3bcd5ad) }

var fileDescriptor_498b213ad3bcd5ad = []byte{
	// 470 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x53, 0x4d, 0x6f, 0x13, 0x31,
	0x10, 0x25, 0x6d, 0xb3, 0x6d, 0x26, 0xe5, 0xa3, 0x16, 0xad, 0xac, 0x08, 0xa4, 0x28, 0x07, 0x94,
	0x53, 0x0f, 0x01, 0xc1, 0x11, 0x55, 0xad, 0x84, 0x22, 0xf1, 0xb1, 0xda, 0x45, 0x5c, 0x91, 0xb3,
	0x19, 0xda, 0xa5, 0x1b, 0x7b, 0xe5, 0xb5, 0x2b, 0x85, 0x03, 0x3f, 0x85, 0x3f, 0xc5, 0x7f, 0xe0,
	0x77, 0xb0, 0x1e, 0x4f, 0x1a, 0x08, 0x5c, 0xd2, 0x53, 0xfc, 0xde, 0x1b, 0x7b, 0x66, 0xdf, 0x9b,
	0xc0, 0xe1, 0x8d, 0xa9, 0xfc, 0x02, 0x4f, 0x6b, 0x6b, 0x9c, 0x11, 0x5d, 0xfa, 0x19, 0x1c, 0x37,
	0x68, 0x6f, 0xca, 0x02, 0x2d, 0x36, 0xb5, 0xd1, 0x0d, 0xab, 0xa3, 0x5f, 0x3b, 0x90, 0x7c, 0xa2,
	0x72, 0xf1, 0x04, 0x7a, 0xa9, 0x35, 0x5f, 0xb1, 0x70, 0xd3, 0xb9, 0xec, 0x0c, 0x3b, 0xe3, 0x5e,
	0xb6, 0x26, 0x82, 0x9a, 0xc7, 0x17, 0x5a, 0x75, 0x27, 0xaa, 0xb7, 0x84, 0x10, 0xb0, 0xf7, 0x5e,
	0x2d, 0x50, 0xee, 0x92, 0x40, 0x67, 0x31, 0x84, 0x3e, 0x17, 0x7c, 0x5c, 0xd6, 0x28, 0xf7, 0x48,
	0xfa, 0x93, 0x12, 0xcf, 0xe0, 0x01, 0xc3, 0xdc, 0xcf, 0xa8, 0xa8, 0x4b, 0x45, 0x1b, 0xac, 0x38,
	0x81, 0x24, 0x77, 0xca, 0xf9, 0x46, 0x26, 0xa4, 0x33, 0x12, 0x8f, 0xa1, 0xeb, 0xcc, 0x35, 0x6a,
	0xb9, 0x4f, 0x74, 0x04, 0xe2, 0x05, 0x1c, 0xf3, 0xfd, 0x0b, 0xac, 0x51, 0xcf, 0x51, 0x17, 0xcb,
	0xa9, 0xfe, 0x62, 0xe4, 0xc1, 0x70, 0xb7, 0xad, 0xfa, 0xbf, 0x28, 0x5e, 0xc3, 0x11, 0x0b, 0x67,
	0xce, 0xd9, 0x72, 0xe6, 0x1d, 0x36, 0xb2, 0xd7, 0xbe, 0xdb, 0x9f, 0x1c, 0x45, 0xaf, 0x4e, 0xa3,
	0x4f, 0x41, 0xce, 0xfe, 0xad, 0x15, 0x4f, 0x01, 0x0a, 0xb3, 0xa8, 0x95, 0x5e, 0x7e, 0x2e, 0xe7,
	0x12, 0xa2, 0x43, 0xcc, 0x4c, 0xe7, 0xa3, 0xef, 0x00, 0xeb, 0xfb, 0xc1, 0xaf, 0xbc, 0xfc, 0x86,
	0x6c, 0x33, 0x9d, 0xc3, 0xd7, 0x9c, 0x57, 0xc6, 0xaf, 0xdc, 0x8d, 0x20, 0xf8, 0xfe, 0xce, 0x78,
	0xed, 0x52, 0xe5, 0xae, 0xd8, 0xde, 0x35, 0x21, 0xc6, 0x90, 0xa4, 0xca, 0xaa, 0x45, 0x43, 0xf6,
	0xf6, 0x27, 0x8f, 0x78, 0x54, 0x22, 0xd1, 0xa1, 0xcd, 0x58, 0x1f, 0xfd, 0xec, 0xb4, 0xf1, 0xae,
	0xd8, 0xd0, 0x9f, 0xfc, 0xe6, 0xfe, 0xe4, 0xf2, 0x18, 0x1e, 0x66, 0x58, 0x57, 0x65, 0xa1, 0x5c,
	0x69, 0x34, 0xc9, 0x71, 0x92, 0x4d, 0x3a, 0xdc, 0x9e, 0x7e, 0x48, 0xf3, 0x55, 0xda, 0xe1, 0x1c,
	0x32, 0x4a, 0x2b, 0x7f, 0x59, 0x6a, 0x0e, 0x9a, 0x91, 0x90, 0xb0, 0x9f, 0x5f, 0x7b, 0x5a, 0x8e,
	0x18, 0xee, 0x0a, 0x8a, 0x01, 0x1c, 0xbc, 0x35, 0xf1, 0x55, 0xce, 0xf5, 0x16, 0xd3, 0x66, 0x38,
	0x63, 0xd5, 0x25, 0x9e, 0x15, 0x45, 0xf8, 0x5a, 0x8e, 0x78, 0x83, 0x9d, 0xfc, 0x68, 0xd7, 0x37,
	0x6e, 0xbb, 0x78, 0x05, 0x87, 0xe7, 0x16, 0x95, 0x43, 0x5e, 0xe7, 0xfb, 0x7f, 0xa5, 0x36, 0x38,
	0x61, 0xc8, 0xc1, 0x65, 0xfc, 0x37, 0x18, 0xdd, 0x6b, 0xf7, 0xa5, 0xf7, 0x06, 0xdd, 0xb6, 0xb7,
	0xda, 0x76, 0x17, 0x58, 0xe1, 0xf6, 0xed, 0x5e, 0x42, 0xbf, 0x8d, 0xae, 0xb8, 0xba, 0xc3, 0x98,
	0xa9, 0xdf, 0x76, 0xcc, 0x59, 0x42, 0xc2, 0xf3, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x19, 0xc3,
	0xb2, 0xf6, 0x14, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// VolumeClient is the client API for Volume service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VolumeClient interface {
	CreateVolume(ctx context.Context, in *Volume, opts ...grpc.CallOption) (*ServiceResponse, error)
	GetVolume(ctx context.Context, in *Volume, opts ...grpc.CallOption) (*ServiceResponse, error)
	DeleteVolume(ctx context.Context, in *Volume, opts ...grpc.CallOption) (*ServiceResponse, error)
	PatchVolume(ctx context.Context, in *Volume, opts ...grpc.CallOption) (*ServiceResponse, error)
	PutVolume(ctx context.Context, in *Volume, opts ...grpc.CallOption) (*ServiceResponse, error)
}

type volumeClient struct {
	cc *grpc.ClientConn
}

func NewVolumeClient(cc *grpc.ClientConn) VolumeClient {
	return &volumeClient{cc}
}

func (c *volumeClient) CreateVolume(ctx context.Context, in *Volume, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.volume/CreateVolume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volumeClient) GetVolume(ctx context.Context, in *Volume, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.volume/GetVolume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volumeClient) DeleteVolume(ctx context.Context, in *Volume, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.volume/DeleteVolume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volumeClient) PatchVolume(ctx context.Context, in *Volume, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.volume/PatchVolume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volumeClient) PutVolume(ctx context.Context, in *Volume, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.volume/PutVolume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VolumeServer is the server API for Volume service.
type VolumeServer interface {
	CreateVolume(context.Context, *Volume) (*ServiceResponse, error)
	GetVolume(context.Context, *Volume) (*ServiceResponse, error)
	DeleteVolume(context.Context, *Volume) (*ServiceResponse, error)
	PatchVolume(context.Context, *Volume) (*ServiceResponse, error)
	PutVolume(context.Context, *Volume) (*ServiceResponse, error)
}

// UnimplementedVolumeServer can be embedded to have forward compatible implementations.
type UnimplementedVolumeServer struct {
}

func (*UnimplementedVolumeServer) CreateVolume(ctx context.Context, req *Volume) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVolume not implemented")
}
func (*UnimplementedVolumeServer) GetVolume(ctx context.Context, req *Volume) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVolume not implemented")
}
func (*UnimplementedVolumeServer) DeleteVolume(ctx context.Context, req *Volume) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteVolume not implemented")
}
func (*UnimplementedVolumeServer) PatchVolume(ctx context.Context, req *Volume) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchVolume not implemented")
}
func (*UnimplementedVolumeServer) PutVolume(ctx context.Context, req *Volume) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutVolume not implemented")
}

func RegisterVolumeServer(s *grpc.Server, srv VolumeServer) {
	s.RegisterService(&_Volume_serviceDesc, srv)
}

func _Volume_CreateVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Volume)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumeServer).CreateVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.volume/CreateVolume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumeServer).CreateVolume(ctx, req.(*Volume))
	}
	return interceptor(ctx, in, info, handler)
}

func _Volume_GetVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Volume)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumeServer).GetVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.volume/GetVolume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumeServer).GetVolume(ctx, req.(*Volume))
	}
	return interceptor(ctx, in, info, handler)
}

func _Volume_DeleteVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Volume)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumeServer).DeleteVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.volume/DeleteVolume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumeServer).DeleteVolume(ctx, req.(*Volume))
	}
	return interceptor(ctx, in, info, handler)
}

func _Volume_PatchVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Volume)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumeServer).PatchVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.volume/PatchVolume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumeServer).PatchVolume(ctx, req.(*Volume))
	}
	return interceptor(ctx, in, info, handler)
}

func _Volume_PutVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Volume)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumeServer).PutVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.volume/PutVolume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumeServer).PutVolume(ctx, req.(*Volume))
	}
	return interceptor(ctx, in, info, handler)
}

var _Volume_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.volume",
	HandlerType: (*VolumeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateVolume",
			Handler:    _Volume_CreateVolume_Handler,
		},
		{
			MethodName: "GetVolume",
			Handler:    _Volume_GetVolume_Handler,
		},
		{
			MethodName: "DeleteVolume",
			Handler:    _Volume_DeleteVolume_Handler,
		},
		{
			MethodName: "PatchVolume",
			Handler:    _Volume_PatchVolume_Handler,
		},
		{
			MethodName: "PutVolume",
			Handler:    _Volume_PutVolume_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "volume.proto",
}
