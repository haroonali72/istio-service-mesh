// Code generated by protoc-gen-go. DO NOT EDIT.
// source: secret.proto

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

type SecretService struct {
	ServiceId               string                   `protobuf:"bytes,1,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	Token                   string                   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	ProjectId               string                   `protobuf:"bytes,3,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	CompanyId               string                   `protobuf:"bytes,4,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	Name                    string                   `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Version                 string                   `protobuf:"bytes,6,opt,name=version,proto3" json:"version,omitempty"`
	ServiceType             string                   `protobuf:"bytes,7,opt,name=service_type,json=serviceType,proto3" json:"service_type,omitempty"`
	ServiceSubType          string                   `protobuf:"bytes,8,opt,name=service_sub_type,json=serviceSubType,proto3" json:"service_sub_type,omitempty"`
	Namespace               string                   `protobuf:"bytes,10,opt,name=namespace,proto3" json:"namespace,omitempty"`
	SecretServiceAttributes *SecretServiceAttributes `protobuf:"bytes,11,opt,name=secret_service_attributes,json=secretServiceAttributes,proto3" json:"secret_service_attributes,omitempty"`
	HookConfiguration       *HookConfiguration       `protobuf:"bytes,12,opt,name=hook_configuration,json=hookConfiguration,proto3" json:"hook_configuration,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                 `json:"-"`
	XXX_unrecognized        []byte                   `json:"-"`
	XXX_sizecache           int32                    `json:"-"`
}

func (m *SecretService) Reset()         { *m = SecretService{} }
func (m *SecretService) String() string { return proto.CompactTextString(m) }
func (*SecretService) ProtoMessage()    {}
func (*SecretService) Descriptor() ([]byte, []int) {
	return fileDescriptor_6acf428160d7a216, []int{0}
}

func (m *SecretService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SecretService.Unmarshal(m, b)
}
func (m *SecretService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SecretService.Marshal(b, m, deterministic)
}
func (m *SecretService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecretService.Merge(m, src)
}
func (m *SecretService) XXX_Size() int {
	return xxx_messageInfo_SecretService.Size(m)
}
func (m *SecretService) XXX_DiscardUnknown() {
	xxx_messageInfo_SecretService.DiscardUnknown(m)
}

var xxx_messageInfo_SecretService proto.InternalMessageInfo

func (m *SecretService) GetServiceId() string {
	if m != nil {
		return m.ServiceId
	}
	return ""
}

func (m *SecretService) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *SecretService) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *SecretService) GetCompanyId() string {
	if m != nil {
		return m.CompanyId
	}
	return ""
}

func (m *SecretService) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SecretService) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *SecretService) GetServiceType() string {
	if m != nil {
		return m.ServiceType
	}
	return ""
}

func (m *SecretService) GetServiceSubType() string {
	if m != nil {
		return m.ServiceSubType
	}
	return ""
}

func (m *SecretService) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *SecretService) GetSecretServiceAttributes() *SecretServiceAttributes {
	if m != nil {
		return m.SecretServiceAttributes
	}
	return nil
}

func (m *SecretService) GetHookConfiguration() *HookConfiguration {
	if m != nil {
		return m.HookConfiguration
	}
	return nil
}

type SecretServiceAttributes struct {
	SecretData           map[string]string `protobuf:"bytes,1,rep,name=secret_data,json=secretData,proto3" json:"secret_data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	SecretType           string            `protobuf:"bytes,2,opt,name=secret_type,json=secretType,proto3" json:"secret_type,omitempty"`
	Data                 map[string][]byte `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SecretServiceAttributes) Reset()         { *m = SecretServiceAttributes{} }
func (m *SecretServiceAttributes) String() string { return proto.CompactTextString(m) }
func (*SecretServiceAttributes) ProtoMessage()    {}
func (*SecretServiceAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_6acf428160d7a216, []int{1}
}

func (m *SecretServiceAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SecretServiceAttributes.Unmarshal(m, b)
}
func (m *SecretServiceAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SecretServiceAttributes.Marshal(b, m, deterministic)
}
func (m *SecretServiceAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecretServiceAttributes.Merge(m, src)
}
func (m *SecretServiceAttributes) XXX_Size() int {
	return xxx_messageInfo_SecretServiceAttributes.Size(m)
}
func (m *SecretServiceAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_SecretServiceAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_SecretServiceAttributes proto.InternalMessageInfo

func (m *SecretServiceAttributes) GetSecretData() map[string]string {
	if m != nil {
		return m.SecretData
	}
	return nil
}

func (m *SecretServiceAttributes) GetSecretType() string {
	if m != nil {
		return m.SecretType
	}
	return ""
}

func (m *SecretServiceAttributes) GetData() map[string][]byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*SecretService)(nil), "proto.SecretService")
	proto.RegisterType((*SecretServiceAttributes)(nil), "proto.SecretServiceAttributes")
	proto.RegisterMapType((map[string][]byte)(nil), "proto.SecretServiceAttributes.DataEntry")
	proto.RegisterMapType((map[string]string)(nil), "proto.SecretServiceAttributes.SecretDataEntry")
}

func init() {
	proto.RegisterFile("secret.proto", fileDescriptor_6acf428160d7a216)
}

var fileDescriptor_6acf428160d7a216 = []byte{
	// 470 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x71, 0x92, 0x92, 0x71, 0x0a, 0x61, 0x5b, 0xd1, 0x25, 0xe2, 0x27, 0xe4, 0xe4, 0x53,
	0x0e, 0xe1, 0x00, 0x42, 0x70, 0x80, 0x26, 0x2a, 0x39, 0x51, 0x39, 0x9c, 0xb8, 0x44, 0x1b, 0x7b,
	0x68, 0x4c, 0xd2, 0x5d, 0x6b, 0x77, 0x1d, 0xc9, 0x6f, 0x07, 0xef, 0xc2, 0x83, 0xa0, 0xfd, 0x89,
	0x69, 0xab, 0x02, 0x95, 0x72, 0xf2, 0xce, 0xf7, 0x33, 0xdf, 0xce, 0x68, 0x0d, 0x5d, 0x85, 0xa9,
	0x44, 0x3d, 0x2a, 0xa4, 0xd0, 0x82, 0xb4, 0xec, 0xa7, 0x7f, 0xa8, 0x50, 0x6e, 0xf3, 0x14, 0x47,
	0xbe, 0xbc, 0x40, 0x8e, 0x92, 0x6d, 0x5c, 0x39, 0xfc, 0x11, 0xc2, 0xe1, 0xdc, 0xba, 0xe6, 0x4e,
	0x46, 0x9e, 0x01, 0x78, 0xc7, 0x22, 0xcf, 0x68, 0x30, 0x08, 0xe2, 0x4e, 0xd2, 0xf1, 0xc8, 0x2c,
	0x23, 0xc7, 0xd0, 0xd2, 0x62, 0x8d, 0x9c, 0x36, 0x2c, 0xe3, 0x0a, 0x63, 0x2a, 0xa4, 0xf8, 0x8e,
	0xa9, 0x36, 0xa6, 0xd0, 0x99, 0x3c, 0x32, 0xcb, 0x0c, 0x9d, 0x8a, 0xcb, 0x82, 0xf1, 0xca, 0xd0,
	0x4d, 0x47, 0x7b, 0x64, 0x96, 0x11, 0x02, 0x4d, 0xce, 0x2e, 0x91, 0xb6, 0x2c, 0x61, 0xcf, 0x84,
	0xc2, 0xc1, 0x16, 0xa5, 0xca, 0x05, 0xa7, 0x6d, 0x0b, 0xef, 0x4a, 0xf2, 0xd2, 0xcc, 0xe9, 0x2e,
	0xa8, 0xab, 0x02, 0xe9, 0x81, 0xa5, 0x23, 0x8f, 0x7d, 0xa9, 0x0a, 0x24, 0x31, 0xf4, 0x76, 0x12,
	0x55, 0x2e, 0x9d, 0xec, 0xbe, 0x95, 0x3d, 0xf0, 0xf8, 0xbc, 0x5c, 0x5a, 0xe5, 0x53, 0xe8, 0x98,
	0x38, 0x55, 0xb0, 0x14, 0x29, 0xb8, 0x8b, 0xd5, 0x00, 0xf9, 0x0a, 0x4f, 0xdc, 0x4a, 0x17, 0xbb,
	0x76, 0x4c, 0x6b, 0x99, 0x2f, 0x4b, 0x8d, 0x8a, 0x46, 0x83, 0x20, 0x8e, 0xc6, 0xcf, 0xdd, 0x22,
	0x47, 0xd7, 0x96, 0xf8, 0xa1, 0x56, 0x25, 0x27, 0xea, 0x76, 0x82, 0x9c, 0x01, 0x59, 0x09, 0xb1,
	0x5e, 0xa4, 0x82, 0x7f, 0xcb, 0x2f, 0x4a, 0xc9, 0xb4, 0x99, 0xb5, 0x6b, 0x9b, 0x52, 0xdf, 0xf4,
	0x93, 0x10, 0xeb, 0xd3, 0xab, 0x7c, 0xf2, 0x68, 0x75, 0x13, 0x1a, 0xfe, 0x6c, 0xc0, 0xc9, 0x5f,
	0xd2, 0xc9, 0x67, 0x88, 0xfc, 0x00, 0x19, 0xd3, 0x8c, 0x06, 0x83, 0x30, 0x8e, 0xc6, 0xa3, 0x7f,
	0x5f, 0xd9, 0xe3, 0x13, 0xa6, 0xd9, 0x94, 0x6b, 0x59, 0x25, 0xa0, 0x6a, 0x80, 0xbc, 0xa8, 0x1b,
	0xda, 0xa5, 0xba, 0x47, 0xe0, 0x05, 0x76, 0xa1, 0xef, 0xa0, 0x69, 0xa3, 0x42, 0x1b, 0x15, 0xff,
	0x27, 0xea, 0x4f, 0x88, 0x75, 0xf5, 0xdf, 0xc3, 0xc3, 0x1b, 0xe9, 0xa4, 0x07, 0xe1, 0x1a, 0x2b,
	0xff, 0x10, 0xcd, 0xd1, 0x3c, 0xc1, 0x2d, 0xdb, 0x94, 0xbb, 0x74, 0x57, 0xbc, 0x6d, 0xbc, 0x09,
	0xfa, 0xaf, 0xa1, 0x73, 0x67, 0x63, 0xf7, 0x8a, 0x71, 0xfc, 0xab, 0x01, 0x6d, 0x17, 0x4c, 0xa6,
	0x70, 0x74, 0x2a, 0x91, 0x69, 0xbc, 0xfe, 0x5b, 0x1c, 0xdf, 0x36, 0x49, 0xff, 0x71, 0x8d, 0xda,
	0x3a, 0x41, 0x55, 0x08, 0xae, 0x70, 0x78, 0x8f, 0x7c, 0x84, 0xde, 0x99, 0xd1, 0xed, 0xd9, 0xe3,
	0xbc, 0xdc, 0xb3, 0xc7, 0x04, 0xc8, 0x39, 0xd3, 0xe9, 0x6a, 0xbf, 0x2e, 0x53, 0x38, 0x9a, 0xe0,
	0x06, 0xf7, 0x5c, 0xca, 0xb2, 0x6d, 0x89, 0x57, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x9f, 0x06,
	0x27, 0xff, 0xa9, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SecretClient is the client API for Secret service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SecretClient interface {
	CreateSecretService(ctx context.Context, in *SecretService, opts ...grpc.CallOption) (*ServiceResponse, error)
	GetSecretService(ctx context.Context, in *SecretService, opts ...grpc.CallOption) (*ServiceResponse, error)
	PutSecretService(ctx context.Context, in *SecretService, opts ...grpc.CallOption) (*ServiceResponse, error)
	PatchSecretService(ctx context.Context, in *SecretService, opts ...grpc.CallOption) (*ServiceResponse, error)
	DeleteSecretService(ctx context.Context, in *SecretService, opts ...grpc.CallOption) (*ServiceResponse, error)
}

type secretClient struct {
	cc grpc.ClientConnInterface
}

func NewSecretClient(cc grpc.ClientConnInterface) SecretClient {
	return &secretClient{cc}
}

func (c *secretClient) CreateSecretService(ctx context.Context, in *SecretService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.Secret/CreateSecretService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretClient) GetSecretService(ctx context.Context, in *SecretService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.Secret/GetSecretService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretClient) PutSecretService(ctx context.Context, in *SecretService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.Secret/PutSecretService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretClient) PatchSecretService(ctx context.Context, in *SecretService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.Secret/PatchSecretService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretClient) DeleteSecretService(ctx context.Context, in *SecretService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.Secret/DeleteSecretService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecretServer is the server API for Secret service.
type SecretServer interface {
	CreateSecretService(context.Context, *SecretService) (*ServiceResponse, error)
	GetSecretService(context.Context, *SecretService) (*ServiceResponse, error)
	PutSecretService(context.Context, *SecretService) (*ServiceResponse, error)
	PatchSecretService(context.Context, *SecretService) (*ServiceResponse, error)
	DeleteSecretService(context.Context, *SecretService) (*ServiceResponse, error)
}

// UnimplementedSecretServer can be embedded to have forward compatible implementations.
type UnimplementedSecretServer struct {
}

func (*UnimplementedSecretServer) CreateSecretService(ctx context.Context, req *SecretService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSecretService not implemented")
}
func (*UnimplementedSecretServer) GetSecretService(ctx context.Context, req *SecretService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSecretService not implemented")
}
func (*UnimplementedSecretServer) PutSecretService(ctx context.Context, req *SecretService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutSecretService not implemented")
}
func (*UnimplementedSecretServer) PatchSecretService(ctx context.Context, req *SecretService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchSecretService not implemented")
}
func (*UnimplementedSecretServer) DeleteSecretService(ctx context.Context, req *SecretService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSecretService not implemented")
}

func RegisterSecretServer(s *grpc.Server, srv SecretServer) {
	s.RegisterService(&_Secret_serviceDesc, srv)
}

func _Secret_CreateSecretService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SecretService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServer).CreateSecretService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Secret/CreateSecretService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServer).CreateSecretService(ctx, req.(*SecretService))
	}
	return interceptor(ctx, in, info, handler)
}

func _Secret_GetSecretService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SecretService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServer).GetSecretService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Secret/GetSecretService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServer).GetSecretService(ctx, req.(*SecretService))
	}
	return interceptor(ctx, in, info, handler)
}

func _Secret_PutSecretService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SecretService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServer).PutSecretService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Secret/PutSecretService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServer).PutSecretService(ctx, req.(*SecretService))
	}
	return interceptor(ctx, in, info, handler)
}

func _Secret_PatchSecretService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SecretService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServer).PatchSecretService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Secret/PatchSecretService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServer).PatchSecretService(ctx, req.(*SecretService))
	}
	return interceptor(ctx, in, info, handler)
}

func _Secret_DeleteSecretService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SecretService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServer).DeleteSecretService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Secret/DeleteSecretService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServer).DeleteSecretService(ctx, req.(*SecretService))
	}
	return interceptor(ctx, in, info, handler)
}

var _Secret_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Secret",
	HandlerType: (*SecretServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSecretService",
			Handler:    _Secret_CreateSecretService_Handler,
		},
		{
			MethodName: "GetSecretService",
			Handler:    _Secret_GetSecretService_Handler,
		},
		{
			MethodName: "PutSecretService",
			Handler:    _Secret_PutSecretService_Handler,
		},
		{
			MethodName: "PatchSecretService",
			Handler:    _Secret_PatchSecretService_Handler,
		},
		{
			MethodName: "DeleteSecretService",
			Handler:    _Secret_DeleteSecretService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "secret.proto",
}
