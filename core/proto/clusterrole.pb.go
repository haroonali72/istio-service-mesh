// Code generated by protoc-gen-go. DO NOT EDIT.
// source: clusterrole.proto

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

type ClusterRole struct {
	ProjectId             string              `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	ServiceId             string              `protobuf:"bytes,2,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	Name                  string              `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	ServiceType           string              `protobuf:"bytes,4,opt,name=service_type,json=serviceType,proto3" json:"service_type,omitempty"`
	ServiceSubType        string              `protobuf:"bytes,5,opt,name=service_sub_type,json=serviceSubType,proto3" json:"service_sub_type,omitempty"`
	Status                string              `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	Token                 string              `protobuf:"bytes,7,opt,name=token,proto3" json:"token,omitempty"`
	ServiceDependencyInfo []string            `protobuf:"bytes,8,rep,name=service_dependency_info,json=serviceDependencyInfo,proto3" json:"service_dependency_info,omitempty"`
	ServiceAttributes     *ClusterRoleSvcAttr `protobuf:"bytes,9,opt,name=service_attributes,json=serviceAttributes,proto3" json:"service_attributes,omitempty"`
	CompanyId             string              `protobuf:"bytes,10,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	Version               string              `protobuf:"bytes,11,opt,name=version,proto3" json:"version,omitempty"`
	Namespace             string              `protobuf:"bytes,12,opt,name=namespace,proto3" json:"namespace,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}            `json:"-"`
	XXX_unrecognized      []byte              `json:"-"`
	XXX_sizecache         int32               `json:"-"`
}

func (m *ClusterRole) Reset()         { *m = ClusterRole{} }
func (m *ClusterRole) String() string { return proto.CompactTextString(m) }
func (*ClusterRole) ProtoMessage()    {}
func (*ClusterRole) Descriptor() ([]byte, []int) {
	return fileDescriptor_5829b5cc33db6b62, []int{0}
}

func (m *ClusterRole) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterRole.Unmarshal(m, b)
}
func (m *ClusterRole) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterRole.Marshal(b, m, deterministic)
}
func (m *ClusterRole) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterRole.Merge(m, src)
}
func (m *ClusterRole) XXX_Size() int {
	return xxx_messageInfo_ClusterRole.Size(m)
}
func (m *ClusterRole) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterRole.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterRole proto.InternalMessageInfo

func (m *ClusterRole) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *ClusterRole) GetServiceId() string {
	if m != nil {
		return m.ServiceId
	}
	return ""
}

func (m *ClusterRole) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ClusterRole) GetServiceType() string {
	if m != nil {
		return m.ServiceType
	}
	return ""
}

func (m *ClusterRole) GetServiceSubType() string {
	if m != nil {
		return m.ServiceSubType
	}
	return ""
}

func (m *ClusterRole) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *ClusterRole) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *ClusterRole) GetServiceDependencyInfo() []string {
	if m != nil {
		return m.ServiceDependencyInfo
	}
	return nil
}

func (m *ClusterRole) GetServiceAttributes() *ClusterRoleSvcAttr {
	if m != nil {
		return m.ServiceAttributes
	}
	return nil
}

func (m *ClusterRole) GetCompanyId() string {
	if m != nil {
		return m.CompanyId
	}
	return ""
}

func (m *ClusterRole) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ClusterRole) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

type ClusterRoleSvcAttr struct {
	Rules                []*Rules `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterRoleSvcAttr) Reset()         { *m = ClusterRoleSvcAttr{} }
func (m *ClusterRoleSvcAttr) String() string { return proto.CompactTextString(m) }
func (*ClusterRoleSvcAttr) ProtoMessage()    {}
func (*ClusterRoleSvcAttr) Descriptor() ([]byte, []int) {
	return fileDescriptor_5829b5cc33db6b62, []int{1}
}

func (m *ClusterRoleSvcAttr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterRoleSvcAttr.Unmarshal(m, b)
}
func (m *ClusterRoleSvcAttr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterRoleSvcAttr.Marshal(b, m, deterministic)
}
func (m *ClusterRoleSvcAttr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterRoleSvcAttr.Merge(m, src)
}
func (m *ClusterRoleSvcAttr) XXX_Size() int {
	return xxx_messageInfo_ClusterRoleSvcAttr.Size(m)
}
func (m *ClusterRoleSvcAttr) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterRoleSvcAttr.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterRoleSvcAttr proto.InternalMessageInfo

func (m *ClusterRoleSvcAttr) GetRules() []*Rules {
	if m != nil {
		return m.Rules
	}
	return nil
}

type Rules struct {
	ResourceName         []string `protobuf:"bytes,1,rep,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	Verbs                []string `protobuf:"bytes,2,rep,name=verbs,proto3" json:"verbs,omitempty"`
	ApiGroup             []string `protobuf:"bytes,3,rep,name=api_group,json=apiGroup,proto3" json:"api_group,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Rules) Reset()         { *m = Rules{} }
func (m *Rules) String() string { return proto.CompactTextString(m) }
func (*Rules) ProtoMessage()    {}
func (*Rules) Descriptor() ([]byte, []int) {
	return fileDescriptor_5829b5cc33db6b62, []int{2}
}

func (m *Rules) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rules.Unmarshal(m, b)
}
func (m *Rules) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rules.Marshal(b, m, deterministic)
}
func (m *Rules) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rules.Merge(m, src)
}
func (m *Rules) XXX_Size() int {
	return xxx_messageInfo_Rules.Size(m)
}
func (m *Rules) XXX_DiscardUnknown() {
	xxx_messageInfo_Rules.DiscardUnknown(m)
}

var xxx_messageInfo_Rules proto.InternalMessageInfo

func (m *Rules) GetResourceName() []string {
	if m != nil {
		return m.ResourceName
	}
	return nil
}

func (m *Rules) GetVerbs() []string {
	if m != nil {
		return m.Verbs
	}
	return nil
}

func (m *Rules) GetApiGroup() []string {
	if m != nil {
		return m.ApiGroup
	}
	return nil
}

func init() {
	proto.RegisterType((*ClusterRole)(nil), "proto.ClusterRole")
	proto.RegisterType((*ClusterRoleSvcAttr)(nil), "proto.ClusterRoleSvcAttr")
	proto.RegisterType((*Rules)(nil), "proto.Rules")
}

func init() { proto.RegisterFile("clusterrole.proto", fileDescriptor_5829b5cc33db6b62) }

var fileDescriptor_5829b5cc33db6b62 = []byte{
	// 453 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x92, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x49, 0x53, 0xa7, 0xf5, 0x38, 0x54, 0xcd, 0x08, 0xca, 0x52, 0x40, 0x2a, 0xe6, 0xd2,
	0x53, 0x0f, 0x41, 0x42, 0x9c, 0x10, 0x51, 0x2b, 0x95, 0x5e, 0x50, 0xe5, 0x70, 0x8f, 0x6c, 0x67,
	0x00, 0x43, 0xea, 0xb5, 0xf6, 0x4f, 0xa4, 0x3c, 0x17, 0xcf, 0xc2, 0xfb, 0x74, 0xf6, 0x8f, 0xa3,
	0x4a, 0xbd, 0x35, 0xa7, 0x64, 0xbe, 0xdf, 0x7c, 0xeb, 0xd9, 0x6f, 0x07, 0x26, 0xf5, 0xca, 0x6a,
	0x43, 0x4a, 0xc9, 0x15, 0x5d, 0x74, 0x4a, 0x1a, 0x89, 0x89, 0xff, 0x39, 0x7d, 0xa9, 0x49, 0xad,
	0x9b, 0x9a, 0x14, 0xe9, 0x4e, 0xb6, 0x3a, 0xd2, 0xfc, 0xdf, 0x10, 0xb2, 0xcb, 0xe0, 0x29, 0xd8,
	0x83, 0xef, 0x00, 0x18, 0xfc, 0xa1, 0xda, 0x2c, 0x9a, 0xa5, 0x18, 0x9c, 0x0d, 0xce, 0xd3, 0x22,
	0x8d, 0xca, 0xcd, 0xd2, 0xe1, 0x78, 0x8e, 0xc3, 0x7b, 0x01, 0x47, 0x85, 0x31, 0xc2, 0x7e, 0x5b,
	0xde, 0x91, 0x18, 0x7a, 0xe0, 0xff, 0xe3, 0x7b, 0x18, 0xf7, 0x16, 0xb3, 0xe9, 0x48, 0xec, 0x7b,
	0x96, 0x45, 0xed, 0x07, 0x4b, 0x78, 0x0e, 0xc7, 0x7d, 0x8b, 0xb6, 0x55, 0x68, 0x4b, 0x7c, 0xdb,
	0x51, 0xd4, 0xe7, 0xb6, 0xf2, 0x9d, 0x27, 0x30, 0xd2, 0xa6, 0x34, 0x56, 0x8b, 0x91, 0xe7, 0xb1,
	0xc2, 0x17, 0x90, 0x18, 0xf9, 0x97, 0x5a, 0x71, 0xe0, 0xe5, 0x50, 0xe0, 0x27, 0x78, 0xd5, 0x9f,
	0xbb, 0xa4, 0x8e, 0xda, 0x25, 0xb5, 0xf5, 0x66, 0xd1, 0xb4, 0x3f, 0xa5, 0x38, 0x3c, 0x1b, 0x72,
	0x5f, 0x1f, 0xca, 0xd5, 0x96, 0xde, 0x30, 0xc4, 0x6f, 0x80, 0xbd, 0xaf, 0x34, 0x46, 0x35, 0x95,
	0x35, 0xa4, 0x45, 0xca, 0x47, 0x67, 0xd3, 0xd7, 0x21, 0xb8, 0x8b, 0x07, 0xa1, 0xcd, 0xd7, 0xf5,
	0x8c, 0xdb, 0x8a, 0x49, 0x34, 0xcd, 0xb6, 0x1e, 0x97, 0x57, 0x2d, 0xef, 0xba, 0xb2, 0xdd, 0xb8,
	0xbc, 0x20, 0xe4, 0x15, 0x15, 0xce, 0x4b, 0xc0, 0xc1, 0x9a, 0x94, 0x6e, 0x64, 0x2b, 0x32, 0xcf,
	0xfa, 0x12, 0xdf, 0x42, 0xea, 0xd2, 0xd3, 0x5d, 0x59, 0x93, 0x18, 0x07, 0xdf, 0x56, 0xc8, 0x3f,
	0x03, 0x3e, 0xfe, 0x3e, 0xe6, 0x90, 0x28, 0xbb, 0xe2, 0x49, 0x07, 0x7c, 0xb9, 0x6c, 0x3a, 0x8e,
	0x93, 0x16, 0x4e, 0x2b, 0x02, 0xca, 0x17, 0x90, 0xf8, 0x1a, 0x3f, 0xc0, 0x73, 0x5e, 0x05, 0x69,
	0x15, 0x5f, 0xd2, 0xbf, 0xd9, 0xc0, 0x27, 0x32, 0xee, 0xc5, 0xef, 0xee, 0xed, 0x38, 0x56, 0x1e,
	0xa8, 0xd2, 0xfc, 0xd2, 0x0e, 0x86, 0x02, 0xdf, 0x40, 0x5a, 0x76, 0xcd, 0xe2, 0x97, 0x92, 0xb6,
	0xe3, 0xa7, 0x76, 0xe4, 0x90, 0x85, 0x6b, 0x57, 0x4f, 0xff, 0xef, 0x6d, 0x17, 0xca, 0x2d, 0x21,
	0xce, 0x60, 0x72, 0xa9, 0xa8, 0x34, 0xf4, 0x70, 0xcb, 0xf0, 0x71, 0x88, 0xa7, 0x27, 0x51, 0x9b,
	0x87, 0x10, 0x8b, 0xb8, 0xa7, 0xf9, 0x33, 0xfc, 0x02, 0x47, 0xd7, 0x64, 0x9e, 0xee, 0xe7, 0x11,
	0xae, 0x68, 0x45, 0xbb, 0x8c, 0xf0, 0x15, 0x8e, 0x6f, 0x4b, 0x53, 0xff, 0xde, 0xe9, 0x12, 0xb7,
	0xf6, 0xe9, 0x97, 0xa8, 0x46, 0x1e, 0x7c, 0xbc, 0x0f, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x9c, 0x33,
	0xd8, 0xe2, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ClusterroleClient is the client API for Clusterrole service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClusterroleClient interface {
	CreateClusterRole(ctx context.Context, in *ClusterRole, opts ...grpc.CallOption) (*ServiceResponse, error)
	GetClusterRole(ctx context.Context, in *ClusterRole, opts ...grpc.CallOption) (*ServiceResponse, error)
	DeleteClusterRole(ctx context.Context, in *ClusterRole, opts ...grpc.CallOption) (*ServiceResponse, error)
	PatchClusterRole(ctx context.Context, in *ClusterRole, opts ...grpc.CallOption) (*ServiceResponse, error)
	PutClusterRole(ctx context.Context, in *ClusterRole, opts ...grpc.CallOption) (*ServiceResponse, error)
}

type clusterroleClient struct {
	cc *grpc.ClientConn
}

func NewClusterroleClient(cc *grpc.ClientConn) ClusterroleClient {
	return &clusterroleClient{cc}
}

func (c *clusterroleClient) CreateClusterRole(ctx context.Context, in *ClusterRole, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.Clusterrole/CreateClusterRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterroleClient) GetClusterRole(ctx context.Context, in *ClusterRole, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.Clusterrole/GetClusterRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterroleClient) DeleteClusterRole(ctx context.Context, in *ClusterRole, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.Clusterrole/DeleteClusterRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterroleClient) PatchClusterRole(ctx context.Context, in *ClusterRole, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.Clusterrole/PatchClusterRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterroleClient) PutClusterRole(ctx context.Context, in *ClusterRole, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.Clusterrole/PutClusterRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClusterroleServer is the server API for Clusterrole service.
type ClusterroleServer interface {
	CreateClusterRole(context.Context, *ClusterRole) (*ServiceResponse, error)
	GetClusterRole(context.Context, *ClusterRole) (*ServiceResponse, error)
	DeleteClusterRole(context.Context, *ClusterRole) (*ServiceResponse, error)
	PatchClusterRole(context.Context, *ClusterRole) (*ServiceResponse, error)
	PutClusterRole(context.Context, *ClusterRole) (*ServiceResponse, error)
}

// UnimplementedClusterroleServer can be embedded to have forward compatible implementations.
type UnimplementedClusterroleServer struct {
}

func (*UnimplementedClusterroleServer) CreateClusterRole(ctx context.Context, req *ClusterRole) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateClusterRole not implemented")
}
func (*UnimplementedClusterroleServer) GetClusterRole(ctx context.Context, req *ClusterRole) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClusterRole not implemented")
}
func (*UnimplementedClusterroleServer) DeleteClusterRole(ctx context.Context, req *ClusterRole) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteClusterRole not implemented")
}
func (*UnimplementedClusterroleServer) PatchClusterRole(ctx context.Context, req *ClusterRole) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchClusterRole not implemented")
}
func (*UnimplementedClusterroleServer) PutClusterRole(ctx context.Context, req *ClusterRole) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutClusterRole not implemented")
}

func RegisterClusterroleServer(s *grpc.Server, srv ClusterroleServer) {
	s.RegisterService(&_Clusterrole_serviceDesc, srv)
}

func _Clusterrole_CreateClusterRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterRole)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterroleServer).CreateClusterRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Clusterrole/CreateClusterRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterroleServer).CreateClusterRole(ctx, req.(*ClusterRole))
	}
	return interceptor(ctx, in, info, handler)
}

func _Clusterrole_GetClusterRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterRole)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterroleServer).GetClusterRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Clusterrole/GetClusterRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterroleServer).GetClusterRole(ctx, req.(*ClusterRole))
	}
	return interceptor(ctx, in, info, handler)
}

func _Clusterrole_DeleteClusterRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterRole)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterroleServer).DeleteClusterRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Clusterrole/DeleteClusterRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterroleServer).DeleteClusterRole(ctx, req.(*ClusterRole))
	}
	return interceptor(ctx, in, info, handler)
}

func _Clusterrole_PatchClusterRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterRole)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterroleServer).PatchClusterRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Clusterrole/PatchClusterRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterroleServer).PatchClusterRole(ctx, req.(*ClusterRole))
	}
	return interceptor(ctx, in, info, handler)
}

func _Clusterrole_PutClusterRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterRole)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterroleServer).PutClusterRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Clusterrole/PutClusterRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterroleServer).PutClusterRole(ctx, req.(*ClusterRole))
	}
	return interceptor(ctx, in, info, handler)
}

var _Clusterrole_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Clusterrole",
	HandlerType: (*ClusterroleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateClusterRole",
			Handler:    _Clusterrole_CreateClusterRole_Handler,
		},
		{
			MethodName: "GetClusterRole",
			Handler:    _Clusterrole_GetClusterRole_Handler,
		},
		{
			MethodName: "DeleteClusterRole",
			Handler:    _Clusterrole_DeleteClusterRole_Handler,
		},
		{
			MethodName: "PatchClusterRole",
			Handler:    _Clusterrole_PatchClusterRole_Handler,
		},
		{
			MethodName: "PutClusterRole",
			Handler:    _Clusterrole_PutClusterRole_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "clusterrole.proto",
}
