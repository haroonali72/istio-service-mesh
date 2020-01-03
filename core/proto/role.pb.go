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

func init() {
	proto.RegisterType((*RoleService)(nil), "proto.RoleService")
	proto.RegisterType((*RoleServiceAttributes)(nil), "proto.RoleServiceAttributes")
	proto.RegisterType((*RoleRule)(nil), "proto.RoleRule")
}

func init() { proto.RegisterFile("role.proto", fileDescriptor_48a3ff9f7c9032f8) }

var fileDescriptor_48a3ff9f7c9032f8 = []byte{
	// 442 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x92, 0x5b, 0x6b, 0x13, 0x41,
	0x14, 0xc7, 0x6d, 0x93, 0xb4, 0xcd, 0x89, 0xd6, 0x76, 0x68, 0x75, 0x0c, 0x15, 0x74, 0x41, 0xe8,
	0xd3, 0x16, 0xea, 0x7b, 0x31, 0x58, 0x28, 0x01, 0x1f, 0xca, 0xd6, 0x57, 0x29, 0x7b, 0x39, 0x8d,
	0xab, 0xdb, 0x99, 0x65, 0x2e, 0x85, 0x7c, 0x2d, 0x3f, 0x88, 0x9f, 0xc9, 0x33, 0xb7, 0x25, 0xa0,
	0x4f, 0xe6, 0x29, 0x39, 0xff, 0xff, 0xf9, 0x9d, 0x39, 0x97, 0x05, 0x50, 0xb2, 0xc3, 0xbc, 0x57,
	0xd2, 0x48, 0x36, 0xf1, 0x3f, 0xf3, 0x17, 0x1a, 0xd5, 0x53, 0x5b, 0x47, 0x75, 0xfe, 0x66, 0x25,
	0xe5, 0xaa, 0xc3, 0x0b, 0x1f, 0x55, 0xf6, 0xe1, 0xa2, 0x14, 0xeb, 0x60, 0x65, 0xbf, 0x46, 0x30,
	0x2b, 0x88, 0xbf, 0x0b, 0x00, 0x7b, 0x0b, 0x10, 0xd9, 0xfb, 0xb6, 0xe1, 0x3b, 0xef, 0x76, 0xce,
	0xa7, 0xc5, 0x34, 0x2a, 0xcb, 0x86, 0x9d, 0xc0, 0xc4, 0xc8, 0x9f, 0x28, 0xf8, 0xae, 0x77, 0x42,
	0xe0, 0x20, 0xaa, 0xf6, 0x03, 0x6b, 0xe3, 0xa0, 0x51, 0x80, 0xa2, 0x42, 0x10, 0xd9, 0xb5, 0x7c,
	0xec, 0xe9, 0x51, 0x67, 0x8f, 0x83, 0x1d, 0x15, 0xb2, 0x19, 0x8c, 0x45, 0xf9, 0x88, 0x7c, 0xe2,
	0x0d, 0xff, 0x9f, 0xbd, 0x87, 0xe7, 0xa9, 0x0d, 0xb3, 0xee, 0x91, 0xef, 0x79, 0x6f, 0x16, 0xb5,
	0xaf, 0x24, 0xb1, 0x73, 0x38, 0x4a, 0x29, 0xda, 0x56, 0x21, 0x6d, 0xdf, 0xa7, 0x1d, 0x46, 0xfd,
	0xce, 0x56, 0x3e, 0xf3, 0x0b, 0xbc, 0x4e, 0x99, 0x0d, 0xf6, 0x28, 0x1a, 0x14, 0x35, 0xb5, 0x22,
	0x1e, 0x24, 0x3f, 0x20, 0x60, 0x76, 0x79, 0x92, 0x87, 0x05, 0xe5, 0x69, 0x41, 0xf9, 0x42, 0xac,
	0x8b, 0xd3, 0x08, 0x5d, 0x0f, 0xcc, 0x92, 0x10, 0x76, 0x06, 0x53, 0xd7, 0xa2, 0xee, 0xcb, 0x1a,
	0xf9, 0x34, 0x0c, 0x33, 0x08, 0x8c, 0xc3, 0xfe, 0x13, 0x2a, 0xdd, 0x4a, 0xc1, 0xc1, 0x7b, 0x29,
	0x64, 0x4b, 0x38, 0x4e, 0x5d, 0x94, 0xc6, 0xa8, 0xb6, 0xb2, 0x06, 0xf9, 0xcc, 0xbf, 0x7f, 0x16,
	0x1e, 0xce, 0x37, 0x0e, 0xb1, 0x48, 0x29, 0xba, 0x48, 0x63, 0x0e, 0x52, 0x76, 0x05, 0xa7, 0xff,
	0x4c, 0x65, 0x1f, 0x60, 0xa2, 0x6c, 0x87, 0x9a, 0x0e, 0x37, 0xa2, 0xba, 0x2f, 0x37, 0xea, 0x16,
	0xa4, 0x17, 0xc1, 0xcd, 0xbe, 0xc1, 0x41, 0x92, 0xdc, 0x38, 0x0a, 0xb5, 0xb4, 0xaa, 0x8e, 0x18,
	0x8d, 0x33, 0x08, 0xee, 0xde, 0xd4, 0x7f, 0xa5, 0xe9, 0xde, 0xce, 0x09, 0x81, 0x3b, 0x68, 0xd9,
	0xb7, 0xf7, 0x2b, 0x25, 0x6d, 0xaf, 0xe9, 0xde, 0x1e, 0x22, 0xe5, 0xc6, 0x0b, 0x97, 0xbf, 0x77,
	0x61, 0xec, 0xea, 0xb3, 0x05, 0x1c, 0x7f, 0x56, 0x58, 0x1a, 0xdc, 0xfc, 0xc2, 0xd8, 0xdf, 0xc3,
	0xce, 0x5f, 0x45, 0x2d, 0xc6, 0x05, 0x2d, 0x53, 0x0a, 0x8d, 0xd9, 0x33, 0x76, 0x05, 0x87, 0x37,
	0x68, 0xb6, 0xe2, 0x6f, 0xed, 0x16, 0xfc, 0x27, 0x38, 0xba, 0x2d, 0x4d, 0xfd, 0xfd, 0xff, 0x2b,
	0xd0, 0x12, 0xae, 0xb1, 0xc3, 0x2d, 0x96, 0x50, 0xed, 0x79, 0xe3, 0xe3, 0x9f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x42, 0xe3, 0x52, 0x6d, 0xea, 0x03, 0x00, 0x00,
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
