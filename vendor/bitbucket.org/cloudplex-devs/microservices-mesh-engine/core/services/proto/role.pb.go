// Code generated by protoc-gen-go. DO NOT EDIT.
// source: role.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

type RoleService struct {
	ServiceId             string                 `protobuf:"bytes,1,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	Token                 string                 `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	ProjectId             string                 `protobuf:"bytes,3,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	CompanyId             string                 `protobuf:"bytes,4,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	Name                  string                 `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	ServiceType           string                 `protobuf:"bytes,6,opt,name=service_type,json=serviceType,proto3" json:"service_type,omitempty"`
	ServiceSubType        string                 `protobuf:"bytes,7,opt,name=service_sub_type,json=serviceSubType,proto3" json:"service_sub_type,omitempty"`
	ServiceDependencyInfo *any.Any               `protobuf:"bytes,8,opt,name=service_dependency_info,json=serviceDependencyInfo,proto3" json:"service_dependency_info,omitempty"`
	Namespace             string                 `protobuf:"bytes,9,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Version               string                 `protobuf:"bytes,10,opt,name=version,proto3" json:"version,omitempty"`
	ServiceAttribute      *RoleServiceAttributes `protobuf:"bytes,11,opt,name=service_attribute,json=serviceAttribute,proto3" json:"service_attribute,omitempty"`
	HookConfiguration     *HookConfiguration     `protobuf:"bytes,12,opt,name=hook_configuration,json=hookConfiguration,proto3" json:"hook_configuration,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}               `json:"-"`
	XXX_unrecognized      []byte                 `json:"-"`
	XXX_sizecache         int32                  `json:"-"`
}

func (m *RoleService) Reset()         { *m = RoleService{} }
func (m *RoleService) String() string { return proto.CompactTextString(m) }
func (*RoleService) ProtoMessage()    {}
func (*RoleService) Descriptor() ([]byte, []int) {
	return fileDescriptor_48a3ff9f7c9032f8, []int{0}
}

func (m *RoleService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleService.Unmarshal(m, b)
}
func (m *RoleService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleService.Marshal(b, m, deterministic)
}
func (m *RoleService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleService.Merge(m, src)
}
func (m *RoleService) XXX_Size() int {
	return xxx_messageInfo_RoleService.Size(m)
}
func (m *RoleService) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleService.DiscardUnknown(m)
}

var xxx_messageInfo_RoleService proto.InternalMessageInfo

func (m *RoleService) GetServiceId() string {
	if m != nil {
		return m.ServiceId
	}
	return ""
}

func (m *RoleService) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *RoleService) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *RoleService) GetCompanyId() string {
	if m != nil {
		return m.CompanyId
	}
	return ""
}

func (m *RoleService) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RoleService) GetServiceType() string {
	if m != nil {
		return m.ServiceType
	}
	return ""
}

func (m *RoleService) GetServiceSubType() string {
	if m != nil {
		return m.ServiceSubType
	}
	return ""
}

func (m *RoleService) GetServiceDependencyInfo() *any.Any {
	if m != nil {
		return m.ServiceDependencyInfo
	}
	return nil
}

func (m *RoleService) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *RoleService) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *RoleService) GetServiceAttribute() *RoleServiceAttributes {
	if m != nil {
		return m.ServiceAttribute
	}
	return nil
}

func (m *RoleService) GetHookConfiguration() *HookConfiguration {
	if m != nil {
		return m.HookConfiguration
	}
	return nil
}

type RoleServiceAttributes struct {
	Rules                []*RoleRule `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *RoleServiceAttributes) Reset()         { *m = RoleServiceAttributes{} }
func (m *RoleServiceAttributes) String() string { return proto.CompactTextString(m) }
func (*RoleServiceAttributes) ProtoMessage()    {}
func (*RoleServiceAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_48a3ff9f7c9032f8, []int{1}
}

func (m *RoleServiceAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleServiceAttributes.Unmarshal(m, b)
}
func (m *RoleServiceAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleServiceAttributes.Marshal(b, m, deterministic)
}
func (m *RoleServiceAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleServiceAttributes.Merge(m, src)
}
func (m *RoleServiceAttributes) XXX_Size() int {
	return xxx_messageInfo_RoleServiceAttributes.Size(m)
}
func (m *RoleServiceAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleServiceAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_RoleServiceAttributes proto.InternalMessageInfo

func (m *RoleServiceAttributes) GetRules() []*RoleRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

type RoleRule struct {
	Resources            []string `protobuf:"bytes,1,rep,name=resources,proto3" json:"resources,omitempty"`
	Verbs                []string `protobuf:"bytes,2,rep,name=verbs,proto3" json:"verbs,omitempty"`
	ApiGroups            []string `protobuf:"bytes,3,rep,name=api_groups,json=apiGroups,proto3" json:"api_groups,omitempty"`
	ResourceNames        []string `protobuf:"bytes,4,rep,name=resource_names,json=resourceNames,proto3" json:"resource_names,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoleRule) Reset()         { *m = RoleRule{} }
func (m *RoleRule) String() string { return proto.CompactTextString(m) }
func (*RoleRule) ProtoMessage()    {}
func (*RoleRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_48a3ff9f7c9032f8, []int{2}
}

func (m *RoleRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleRule.Unmarshal(m, b)
}
func (m *RoleRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleRule.Marshal(b, m, deterministic)
}
func (m *RoleRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleRule.Merge(m, src)
}
func (m *RoleRule) XXX_Size() int {
	return xxx_messageInfo_RoleRule.Size(m)
}
func (m *RoleRule) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleRule.DiscardUnknown(m)
}

var xxx_messageInfo_RoleRule proto.InternalMessageInfo

func (m *RoleRule) GetResources() []string {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *RoleRule) GetVerbs() []string {
	if m != nil {
		return m.Verbs
	}
	return nil
}

func (m *RoleRule) GetApiGroups() []string {
	if m != nil {
		return m.ApiGroups
	}
	return nil
}

func (m *RoleRule) GetResourceNames() []string {
	if m != nil {
		return m.ResourceNames
	}
	return nil
}

func init() {
	proto.RegisterType((*RoleService)(nil), "proto.RoleService")
	proto.RegisterType((*RoleServiceAttributes)(nil), "proto.RoleServiceAttributes")
	proto.RegisterType((*RoleRule)(nil), "proto.RoleRule")
}

func init() { proto.RegisterFile("role.proto", fileDescriptor_48a3ff9f7c9032f8) }

var fileDescriptor_48a3ff9f7c9032f8 = []byte{
	// 505 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x4d, 0x6b, 0x1b, 0x3f,
	0x10, 0xc6, 0xff, 0x7e, 0x4b, 0xe2, 0x71, 0xe2, 0x7f, 0x2c, 0x92, 0x56, 0x35, 0x29, 0xb8, 0x86,
	0x80, 0x4f, 0x1b, 0x70, 0xef, 0xa1, 0x26, 0x01, 0xd7, 0x50, 0x4a, 0xd8, 0xf4, 0x6e, 0xf6, 0x65,
	0x6c, 0x6f, 0xbd, 0x91, 0x84, 0xa4, 0x35, 0xec, 0xbd, 0x9f, 0xaf, 0x5f, 0xa8, 0x97, 0xa2, 0x37,
	0xd7, 0xb4, 0x3d, 0xd5, 0x27, 0x7b, 0x7e, 0xcf, 0xf3, 0x68, 0x46, 0xb3, 0x02, 0x90, 0xbc, 0xc4,
	0x48, 0x48, 0xae, 0x39, 0xe9, 0xd8, 0x9f, 0xe1, 0xc5, 0x1a, 0x19, 0xca, 0xa4, 0x8c, 0x7c, 0xa9,
	0x50, 0xee, 0x8a, 0xcc, 0x9b, 0x86, 0x6f, 0xd6, 0x9c, 0xaf, 0x4b, 0xbc, 0xb3, 0x55, 0x5a, 0xad,
	0xee, 0x12, 0x56, 0x3b, 0x69, 0xfc, 0xa3, 0x05, 0xbd, 0x98, 0x97, 0xf8, 0xec, 0x02, 0xe4, 0x2d,
	0x80, 0xcf, 0x2e, 0x8b, 0x9c, 0x36, 0x46, 0x8d, 0x49, 0x37, 0xee, 0x7a, 0xb2, 0xc8, 0xc9, 0x15,
	0x74, 0x34, 0xdf, 0x22, 0xa3, 0x4d, 0xab, 0xb8, 0xc2, 0x84, 0x84, 0xe4, 0x5f, 0x31, 0xd3, 0x26,
	0xd4, 0x72, 0x21, 0x4f, 0x16, 0xb9, 0x91, 0x33, 0xfe, 0x22, 0x12, 0x56, 0x1b, 0xb9, 0xed, 0x64,
	0x4f, 0x16, 0x39, 0x21, 0xd0, 0x66, 0xc9, 0x0b, 0xd2, 0x8e, 0x15, 0xec, 0x7f, 0xf2, 0x0e, 0xce,
	0xc3, 0x18, 0xba, 0x16, 0x48, 0x4f, 0xac, 0xd6, 0xf3, 0xec, 0x4b, 0x2d, 0x90, 0x4c, 0xe0, 0x32,
	0x58, 0x54, 0x95, 0x3a, 0xdb, 0xa9, 0xb5, 0xf5, 0x3d, 0x7f, 0xae, 0x52, 0xeb, 0xfc, 0x04, 0xaf,
	0x83, 0x33, 0x47, 0x81, 0x2c, 0x47, 0x96, 0xd5, 0xcb, 0x82, 0xad, 0x38, 0x3d, 0x1b, 0x35, 0x26,
	0xbd, 0xe9, 0x55, 0xe4, 0x16, 0x14, 0x85, 0x05, 0x45, 0x33, 0x56, 0xc7, 0xd7, 0x3e, 0xf4, 0xb8,
	0xcf, 0x2c, 0xd8, 0x8a, 0x93, 0x1b, 0xe8, 0x9a, 0x11, 0x95, 0x48, 0x32, 0xa4, 0x5d, 0x77, 0x99,
	0x3d, 0x20, 0x14, 0x4e, 0x77, 0x28, 0x55, 0xc1, 0x19, 0x05, 0xab, 0x85, 0x92, 0x2c, 0x60, 0x10,
	0xa6, 0x48, 0xb4, 0x96, 0x45, 0x5a, 0x69, 0xa4, 0x3d, 0xdb, 0xff, 0xc6, 0x35, 0x8e, 0x0e, 0x3e,
	0xc4, 0x2c, 0x58, 0x54, 0x1c, 0xae, 0xb9, 0x47, 0x64, 0x0e, 0x64, 0xc3, 0xf9, 0x76, 0x99, 0x71,
	0xb6, 0x2a, 0xd6, 0x95, 0x4c, 0xb4, 0xe9, 0x77, 0x6e, 0xcf, 0xa2, 0xfe, 0xac, 0x8f, 0x9c, 0x6f,
	0x1f, 0x0e, 0xf5, 0x78, 0xb0, 0xf9, 0x1d, 0x8d, 0xef, 0xe1, 0xfa, 0xaf, 0x3d, 0xc9, 0x2d, 0x74,
	0x64, 0x55, 0xa2, 0xa2, 0x8d, 0x51, 0x6b, 0xd2, 0x9b, 0xfe, 0x7f, 0x30, 0x60, 0x5c, 0x95, 0x18,
	0x3b, 0x75, 0xfc, 0xad, 0x01, 0x67, 0x81, 0x99, 0xc5, 0x48, 0x54, 0xbc, 0x92, 0x99, 0xcf, 0x75,
	0xe3, 0x5f, 0xc0, 0xbc, 0x9c, 0x1d, 0xca, 0x54, 0xd1, 0xa6, 0x55, 0x5c, 0x61, 0x9e, 0x46, 0x22,
	0x8a, 0xe5, 0x5a, 0xf2, 0x4a, 0x28, 0xda, 0x72, 0xa1, 0x44, 0x14, 0x73, 0x0b, 0xc8, 0x2d, 0xf4,
	0xc3, 0x09, 0x4b, 0xbb, 0x63, 0xda, 0xb6, 0x96, 0x8b, 0x40, 0x3f, 0x1b, 0x38, 0xfd, 0xde, 0x84,
	0xb6, 0x19, 0x83, 0xcc, 0x60, 0xf0, 0x20, 0x31, 0xd1, 0x78, 0xf8, 0xa4, 0xc9, 0x9f, 0xdb, 0x1d,
	0xbe, 0xf2, 0xcc, 0xd7, 0x31, 0x2a, 0xc1, 0x99, 0xc2, 0xf1, 0x7f, 0xe4, 0x1e, 0xfa, 0x73, 0xd4,
	0x47, 0xe5, 0x9f, 0xaa, 0x23, 0xf2, 0x1f, 0xe0, 0xf2, 0x29, 0xd1, 0xd9, 0xe6, 0xdf, 0x4f, 0x98,
	0xc1, 0xe0, 0x11, 0x4b, 0x3c, 0x62, 0x09, 0xe9, 0x89, 0x15, 0xde, 0xff, 0x0c, 0x00, 0x00, 0xff,
	0xff, 0x7a, 0x41, 0x67, 0xf8, 0x6a, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RoleClient is the client API for Role service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RoleClient interface {
	CreateRoleService(ctx context.Context, in *RoleService, opts ...grpc.CallOption) (*ServiceResponse, error)
	GetRoleService(ctx context.Context, in *RoleService, opts ...grpc.CallOption) (*ServiceResponse, error)
	PutRoleService(ctx context.Context, in *RoleService, opts ...grpc.CallOption) (*ServiceResponse, error)
	PatchRoleService(ctx context.Context, in *RoleService, opts ...grpc.CallOption) (*ServiceResponse, error)
	DeleteRoleService(ctx context.Context, in *RoleService, opts ...grpc.CallOption) (*ServiceResponse, error)
}

type roleClient struct {
	cc *grpc.ClientConn
}

func NewRoleClient(cc *grpc.ClientConn) RoleClient {
	return &roleClient{cc}
}

func (c *roleClient) CreateRoleService(ctx context.Context, in *RoleService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.Role/CreateRoleService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleClient) GetRoleService(ctx context.Context, in *RoleService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.Role/GetRoleService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleClient) PutRoleService(ctx context.Context, in *RoleService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.Role/PutRoleService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleClient) PatchRoleService(ctx context.Context, in *RoleService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.Role/PatchRoleService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleClient) DeleteRoleService(ctx context.Context, in *RoleService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.Role/DeleteRoleService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoleServer is the server API for Role service.
type RoleServer interface {
	CreateRoleService(context.Context, *RoleService) (*ServiceResponse, error)
	GetRoleService(context.Context, *RoleService) (*ServiceResponse, error)
	PutRoleService(context.Context, *RoleService) (*ServiceResponse, error)
	PatchRoleService(context.Context, *RoleService) (*ServiceResponse, error)
	DeleteRoleService(context.Context, *RoleService) (*ServiceResponse, error)
}

// UnimplementedRoleServer can be embedded to have forward compatible implementations.
type UnimplementedRoleServer struct {
}

func (*UnimplementedRoleServer) CreateRoleService(ctx context.Context, req *RoleService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoleService not implemented")
}
func (*UnimplementedRoleServer) GetRoleService(ctx context.Context, req *RoleService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoleService not implemented")
}
func (*UnimplementedRoleServer) PutRoleService(ctx context.Context, req *RoleService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutRoleService not implemented")
}
func (*UnimplementedRoleServer) PatchRoleService(ctx context.Context, req *RoleService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchRoleService not implemented")
}
func (*UnimplementedRoleServer) DeleteRoleService(ctx context.Context, req *RoleService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRoleService not implemented")
}

func RegisterRoleServer(s *grpc.Server, srv RoleServer) {
	s.RegisterService(&_Role_serviceDesc, srv)
}

func _Role_CreateRoleService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).CreateRoleService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Role/CreateRoleService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).CreateRoleService(ctx, req.(*RoleService))
	}
	return interceptor(ctx, in, info, handler)
}

func _Role_GetRoleService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).GetRoleService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Role/GetRoleService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).GetRoleService(ctx, req.(*RoleService))
	}
	return interceptor(ctx, in, info, handler)
}

func _Role_PutRoleService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).PutRoleService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Role/PutRoleService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).PutRoleService(ctx, req.(*RoleService))
	}
	return interceptor(ctx, in, info, handler)
}

func _Role_PatchRoleService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).PatchRoleService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Role/PatchRoleService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).PatchRoleService(ctx, req.(*RoleService))
	}
	return interceptor(ctx, in, info, handler)
}

func _Role_DeleteRoleService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).DeleteRoleService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Role/DeleteRoleService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).DeleteRoleService(ctx, req.(*RoleService))
	}
	return interceptor(ctx, in, info, handler)
}

var _Role_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Role",
	HandlerType: (*RoleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRoleService",
			Handler:    _Role_CreateRoleService_Handler,
		},
		{
			MethodName: "GetRoleService",
			Handler:    _Role_GetRoleService_Handler,
		},
		{
			MethodName: "PutRoleService",
			Handler:    _Role_PutRoleService_Handler,
		},
		{
			MethodName: "PatchRoleService",
			Handler:    _Role_PatchRoleService_Handler,
		},
		{
			MethodName: "DeleteRoleService",
			Handler:    _Role_DeleteRoleService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "role.proto",
}
