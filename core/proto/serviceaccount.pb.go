// Code generated by protoc-gen-go. DO NOT EDIT.
// source: serviceaccount.proto

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

type ServiceAccountService struct {
	ServiceId                string                    `protobuf:"bytes,1,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	Token                    string                    `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	ProjectId                string                    `protobuf:"bytes,3,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	CompanyId                string                    `protobuf:"bytes,5,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	Name                     string                    `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Namespace                string                    `protobuf:"bytes,7,opt,name=namespace,proto3" json:"namespace,omitempty"`
	ServiceType              string                    `protobuf:"bytes,8,opt,name=service_type,json=serviceType,proto3" json:"service_type,omitempty"`
	ServiceSubType           string                    `protobuf:"bytes,9,opt,name=service_sub_type,json=serviceSubType,proto3" json:"service_sub_type,omitempty"`
	Version                  string                    `protobuf:"bytes,10,opt,name=version,proto3" json:"version,omitempty"`
	ServiceAccountAttributes *ServiceAccountAttributes `protobuf:"bytes,11,opt,name=serviceAccountAttributes,proto3" json:"serviceAccountAttributes,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}                  `json:"-"`
	XXX_unrecognized         []byte                    `json:"-"`
	XXX_sizecache            int32                     `json:"-"`
}

func (m *ServiceAccountService) Reset()         { *m = ServiceAccountService{} }
func (m *ServiceAccountService) String() string { return proto.CompactTextString(m) }
func (*ServiceAccountService) ProtoMessage()    {}
func (*ServiceAccountService) Descriptor() ([]byte, []int) {
	return fileDescriptor_f552791109664fa6, []int{0}
}

func (m *ServiceAccountService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceAccountService.Unmarshal(m, b)
}
func (m *ServiceAccountService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceAccountService.Marshal(b, m, deterministic)
}
func (m *ServiceAccountService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceAccountService.Merge(m, src)
}
func (m *ServiceAccountService) XXX_Size() int {
	return xxx_messageInfo_ServiceAccountService.Size(m)
}
func (m *ServiceAccountService) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceAccountService.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceAccountService proto.InternalMessageInfo

func (m *ServiceAccountService) GetServiceId() string {
	if m != nil {
		return m.ServiceId
	}
	return ""
}

func (m *ServiceAccountService) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *ServiceAccountService) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *ServiceAccountService) GetCompanyId() string {
	if m != nil {
		return m.CompanyId
	}
	return ""
}

func (m *ServiceAccountService) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ServiceAccountService) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ServiceAccountService) GetServiceType() string {
	if m != nil {
		return m.ServiceType
	}
	return ""
}

func (m *ServiceAccountService) GetServiceSubType() string {
	if m != nil {
		return m.ServiceSubType
	}
	return ""
}

func (m *ServiceAccountService) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ServiceAccountService) GetServiceAccountAttributes() *ServiceAccountAttributes {
	if m != nil {
		return m.ServiceAccountAttributes
	}
	return nil
}

type ServiceAccountServiceResponse struct {
	Error                string                 `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Resp                 *ServiceAccountService `protobuf:"bytes,2,opt,name=resp,proto3" json:"resp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ServiceAccountServiceResponse) Reset()         { *m = ServiceAccountServiceResponse{} }
func (m *ServiceAccountServiceResponse) String() string { return proto.CompactTextString(m) }
func (*ServiceAccountServiceResponse) ProtoMessage()    {}
func (*ServiceAccountServiceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f552791109664fa6, []int{1}
}

func (m *ServiceAccountServiceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceAccountServiceResponse.Unmarshal(m, b)
}
func (m *ServiceAccountServiceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceAccountServiceResponse.Marshal(b, m, deterministic)
}
func (m *ServiceAccountServiceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceAccountServiceResponse.Merge(m, src)
}
func (m *ServiceAccountServiceResponse) XXX_Size() int {
	return xxx_messageInfo_ServiceAccountServiceResponse.Size(m)
}
func (m *ServiceAccountServiceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceAccountServiceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceAccountServiceResponse proto.InternalMessageInfo

func (m *ServiceAccountServiceResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ServiceAccountServiceResponse) GetResp() *ServiceAccountService {
	if m != nil {
		return m.Resp
	}
	return nil
}

type ServiceAccountAttributes struct {
	Secrets              []string `protobuf:"bytes,1,rep,name=secrets,proto3" json:"secrets,omitempty"`
	ImagePullSecretsName []string `protobuf:"bytes,2,rep,name=imagePullSecretsName,proto3" json:"imagePullSecretsName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceAccountAttributes) Reset()         { *m = ServiceAccountAttributes{} }
func (m *ServiceAccountAttributes) String() string { return proto.CompactTextString(m) }
func (*ServiceAccountAttributes) ProtoMessage()    {}
func (*ServiceAccountAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_f552791109664fa6, []int{2}
}

func (m *ServiceAccountAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceAccountAttributes.Unmarshal(m, b)
}
func (m *ServiceAccountAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceAccountAttributes.Marshal(b, m, deterministic)
}
func (m *ServiceAccountAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceAccountAttributes.Merge(m, src)
}
func (m *ServiceAccountAttributes) XXX_Size() int {
	return xxx_messageInfo_ServiceAccountAttributes.Size(m)
}
func (m *ServiceAccountAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceAccountAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceAccountAttributes proto.InternalMessageInfo

func (m *ServiceAccountAttributes) GetSecrets() []string {
	if m != nil {
		return m.Secrets
	}
	return nil
}

func (m *ServiceAccountAttributes) GetImagePullSecretsName() []string {
	if m != nil {
		return m.ImagePullSecretsName
	}
	return nil
}

func init() {
	proto.RegisterType((*ServiceAccountService)(nil), "proto.ServiceAccountService")
	proto.RegisterType((*ServiceAccountServiceResponse)(nil), "proto.ServiceAccountServiceResponse")
	proto.RegisterType((*ServiceAccountAttributes)(nil), "proto.ServiceAccountAttributes")
}

func init() { proto.RegisterFile("serviceaccount.proto", fileDescriptor_f552791109664fa6) }

var fileDescriptor_f552791109664fa6 = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xcf, 0x8e, 0xd3, 0x30,
	0x10, 0xc6, 0xc9, 0x66, 0xbb, 0x4b, 0x26, 0xb0, 0x42, 0x56, 0x41, 0x56, 0xb7, 0x15, 0x25, 0xa7,
	0x9e, 0x2a, 0x14, 0x9e, 0xa0, 0x02, 0x09, 0xf5, 0x82, 0xaa, 0xa4, 0x3d, 0x71, 0x40, 0x89, 0x33,
	0x6a, 0x03, 0x6d, 0x6c, 0xd9, 0x4e, 0xa5, 0xde, 0x78, 0x30, 0x1e, 0x0e, 0xf9, 0x4f, 0x90, 0x2a,
	0xa5, 0x9c, 0x72, 0x72, 0xe6, 0xfb, 0x7e, 0x33, 0xb6, 0x3f, 0x07, 0xc6, 0x0a, 0xe5, 0xb9, 0x66,
	0x58, 0x30, 0xc6, 0xdb, 0x46, 0x2f, 0x85, 0xe4, 0x9a, 0x93, 0x91, 0x5d, 0x26, 0xaf, 0xbd, 0xe9,
	0xd4, 0xe4, 0x77, 0x08, 0x6f, 0x73, 0xa7, 0xac, 0x1c, 0xee, 0x2b, 0x32, 0x03, 0xf0, 0xe8, 0x8f,
	0xba, 0xa2, 0xc1, 0x3c, 0x58, 0x44, 0x59, 0xe4, 0x95, 0x75, 0x45, 0xc6, 0x30, 0xd2, 0xfc, 0x17,
	0x36, 0xf4, 0xce, 0x3a, 0xae, 0x30, 0x4d, 0x42, 0xf2, 0x9f, 0xc8, 0xb4, 0x69, 0x0a, 0x5d, 0x93,
	0x57, 0xd6, 0x95, 0xb1, 0x19, 0x3f, 0x89, 0xa2, 0xb9, 0x18, 0x7b, 0xe4, 0x6c, 0xaf, 0xac, 0x2b,
	0x42, 0xe0, 0xbe, 0x29, 0x4e, 0x48, 0x1f, 0xac, 0x61, 0xbf, 0xc9, 0x14, 0x22, 0xb3, 0x2a, 0x51,
	0x30, 0xa4, 0x8f, 0xae, 0xe3, 0x9f, 0x40, 0x3e, 0xc0, 0xab, 0xee, 0x90, 0xfa, 0x22, 0x90, 0xbe,
	0xb4, 0x40, 0xec, 0xb5, 0xed, 0x45, 0x20, 0x59, 0xc0, 0x9b, 0x0e, 0x51, 0x6d, 0xe9, 0xb0, 0xc8,
	0x62, 0x4f, 0x5e, 0xcf, 0xdb, 0xd2, 0x92, 0x14, 0x1e, 0xcf, 0x28, 0x55, 0xcd, 0x1b, 0x0a, 0x16,
	0xe8, 0x4a, 0xf2, 0x1d, 0xa8, 0xba, 0x0a, 0x69, 0xa5, 0xb5, 0xac, 0xcb, 0x56, 0xa3, 0xa2, 0xf1,
	0x3c, 0x58, 0xc4, 0xe9, 0x7b, 0x97, 0xe7, 0x32, 0xbf, 0x81, 0x65, 0x37, 0x07, 0x24, 0x7b, 0x98,
	0xf5, 0xbe, 0x40, 0x86, 0x4a, 0xf0, 0x46, 0xa1, 0x89, 0x1a, 0xa5, 0xe4, 0xd2, 0x3f, 0x82, 0x2b,
	0xc8, 0x47, 0xb8, 0x97, 0xa8, 0x84, 0xcd, 0x3f, 0x4e, 0xa7, 0xbd, 0xfb, 0x77, 0x93, 0x2c, 0x99,
	0x1c, 0x80, 0xde, 0x3a, 0x9e, 0xb9, 0xbb, 0x42, 0x26, 0x51, 0x2b, 0x1a, 0xcc, 0x43, 0x73, 0x77,
	0x5f, 0x92, 0x14, 0xc6, 0xf5, 0xa9, 0xd8, 0xe3, 0xa6, 0x3d, 0x1e, 0x73, 0xa7, 0x7d, 0x33, 0x8f,
	0x74, 0x67, 0xb1, 0x5e, 0x2f, 0xfd, 0x13, 0xc2, 0xd3, 0xf5, 0x56, 0x64, 0x07, 0xcf, 0x9f, 0x25,
	0x16, 0x1a, 0xfb, 0xff, 0xb6, 0xff, 0x9e, 0x7f, 0xf2, 0xee, 0xda, 0xed, 0x92, 0x49, 0x5e, 0x90,
	0x0c, 0xe8, 0x57, 0xd4, 0x83, 0xcf, 0xdc, 0xb4, 0x03, 0xcf, 0xdc, 0xc2, 0x64, 0x53, 0x68, 0x76,
	0x18, 0x76, 0xea, 0x0e, 0x9e, 0xbf, 0xe0, 0x11, 0x07, 0x0e, 0xb5, 0x7c, 0xb0, 0xc6, 0xa7, 0xbf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xce, 0x76, 0x19, 0x68, 0x49, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ServiceAccountClient is the client API for ServiceAccount service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceAccountClient interface {
	CreateServiceAccountService(ctx context.Context, in *ServiceAccountService, opts ...grpc.CallOption) (*ServiceResponse, error)
	GetServiceAccountService(ctx context.Context, in *ServiceAccountService, opts ...grpc.CallOption) (*ServiceResponse, error)
	PutServiceAccountService(ctx context.Context, in *ServiceAccountService, opts ...grpc.CallOption) (*ServiceResponse, error)
	PatchServiceAccountService(ctx context.Context, in *ServiceAccountService, opts ...grpc.CallOption) (*ServiceResponse, error)
	DeleteServiceAccountService(ctx context.Context, in *ServiceAccountService, opts ...grpc.CallOption) (*ServiceResponse, error)
}

type serviceAccountClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceAccountClient(cc grpc.ClientConnInterface) ServiceAccountClient {
	return &serviceAccountClient{cc}
}

func (c *serviceAccountClient) CreateServiceAccountService(ctx context.Context, in *ServiceAccountService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.ServiceAccount/CreateServiceAccountService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceAccountClient) GetServiceAccountService(ctx context.Context, in *ServiceAccountService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.ServiceAccount/GetServiceAccountService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceAccountClient) PutServiceAccountService(ctx context.Context, in *ServiceAccountService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.ServiceAccount/PutServiceAccountService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceAccountClient) PatchServiceAccountService(ctx context.Context, in *ServiceAccountService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.ServiceAccount/PatchServiceAccountService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceAccountClient) DeleteServiceAccountService(ctx context.Context, in *ServiceAccountService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.ServiceAccount/DeleteServiceAccountService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceAccountServer is the server API for ServiceAccount service.
type ServiceAccountServer interface {
	CreateServiceAccountService(context.Context, *ServiceAccountService) (*ServiceResponse, error)
	GetServiceAccountService(context.Context, *ServiceAccountService) (*ServiceResponse, error)
	PutServiceAccountService(context.Context, *ServiceAccountService) (*ServiceResponse, error)
	PatchServiceAccountService(context.Context, *ServiceAccountService) (*ServiceResponse, error)
	DeleteServiceAccountService(context.Context, *ServiceAccountService) (*ServiceResponse, error)
}

// UnimplementedServiceAccountServer can be embedded to have forward compatible implementations.
type UnimplementedServiceAccountServer struct {
}

func (*UnimplementedServiceAccountServer) CreateServiceAccountService(ctx context.Context, req *ServiceAccountService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateServiceAccountService not implemented")
}
func (*UnimplementedServiceAccountServer) GetServiceAccountService(ctx context.Context, req *ServiceAccountService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServiceAccountService not implemented")
}
func (*UnimplementedServiceAccountServer) PutServiceAccountService(ctx context.Context, req *ServiceAccountService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutServiceAccountService not implemented")
}
func (*UnimplementedServiceAccountServer) PatchServiceAccountService(ctx context.Context, req *ServiceAccountService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchServiceAccountService not implemented")
}
func (*UnimplementedServiceAccountServer) DeleteServiceAccountService(ctx context.Context, req *ServiceAccountService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteServiceAccountService not implemented")
}

func RegisterServiceAccountServer(s *grpc.Server, srv ServiceAccountServer) {
	s.RegisterService(&_ServiceAccount_serviceDesc, srv)
}

func _ServiceAccount_CreateServiceAccountService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceAccountService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceAccountServer).CreateServiceAccountService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ServiceAccount/CreateServiceAccountService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceAccountServer).CreateServiceAccountService(ctx, req.(*ServiceAccountService))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceAccount_GetServiceAccountService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceAccountService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceAccountServer).GetServiceAccountService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ServiceAccount/GetServiceAccountService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceAccountServer).GetServiceAccountService(ctx, req.(*ServiceAccountService))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceAccount_PutServiceAccountService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceAccountService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceAccountServer).PutServiceAccountService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ServiceAccount/PutServiceAccountService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceAccountServer).PutServiceAccountService(ctx, req.(*ServiceAccountService))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceAccount_PatchServiceAccountService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceAccountService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceAccountServer).PatchServiceAccountService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ServiceAccount/PatchServiceAccountService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceAccountServer).PatchServiceAccountService(ctx, req.(*ServiceAccountService))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceAccount_DeleteServiceAccountService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceAccountService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceAccountServer).DeleteServiceAccountService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ServiceAccount/DeleteServiceAccountService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceAccountServer).DeleteServiceAccountService(ctx, req.(*ServiceAccountService))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServiceAccount_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ServiceAccount",
	HandlerType: (*ServiceAccountServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateServiceAccountService",
			Handler:    _ServiceAccount_CreateServiceAccountService_Handler,
		},
		{
			MethodName: "GetServiceAccountService",
			Handler:    _ServiceAccount_GetServiceAccountService_Handler,
		},
		{
			MethodName: "PutServiceAccountService",
			Handler:    _ServiceAccount_PutServiceAccountService_Handler,
		},
		{
			MethodName: "PatchServiceAccountService",
			Handler:    _ServiceAccount_PatchServiceAccountService_Handler,
		},
		{
			MethodName: "DeleteServiceAccountService",
			Handler:    _ServiceAccount_DeleteServiceAccountService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "serviceaccount.proto",
}
