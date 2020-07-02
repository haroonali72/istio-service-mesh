// Code generated by protoc-gen-go. DO NOT EDIT.
// source: peerauthenticaion.proto

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

type TlsMode int32

const (
	TlsMode_STRICT     TlsMode = 0
	TlsMode_PERMISSIVE TlsMode = 1
	TlsMode_DISABLE    TlsMode = 2
	TlsMode_UNSET      TlsMode = 3
)

var TlsMode_name = map[int32]string{
	0: "STRICT",
	1: "PERMISSIVE",
	2: "DISABLE",
	3: "UNSET",
}

var TlsMode_value = map[string]int32{
	"STRICT":     0,
	"PERMISSIVE": 1,
	"DISABLE":    2,
	"UNSET":      3,
}

func (x TlsMode) String() string {
	return proto.EnumName(TlsMode_name, int32(x))
}

func (TlsMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8fda6fe3c8d61939, []int{0}
}

type PeerAuthenticationService struct {
	ProjectId            string                               `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	ServiceId            string                               `protobuf:"bytes,2,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	Name                 string                               `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Version              string                               `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	ServiceType          string                               `protobuf:"bytes,5,opt,name=service_type,json=serviceType,proto3" json:"service_type,omitempty"`
	ServiceSubType       string                               `protobuf:"bytes,6,opt,name=service_sub_type,json=serviceSubType,proto3" json:"service_sub_type,omitempty"`
	Namespace            string                               `protobuf:"bytes,7,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Token                string                               `protobuf:"bytes,8,opt,name=token,proto3" json:"token,omitempty"`
	CompanyId            string                               `protobuf:"bytes,9,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	IsDiscovered         bool                                 `protobuf:"varint,10,opt,name=is_discovered,json=isDiscovered,proto3" json:"is_discovered,omitempty"`
	ServiceAttributes    *PeerAuthenticationServiceAttributes `protobuf:"bytes,11,opt,name=service_attributes,json=serviceAttributes,proto3" json:"service_attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *PeerAuthenticationService) Reset()         { *m = PeerAuthenticationService{} }
func (m *PeerAuthenticationService) String() string { return proto.CompactTextString(m) }
func (*PeerAuthenticationService) ProtoMessage()    {}
func (*PeerAuthenticationService) Descriptor() ([]byte, []int) {
	return fileDescriptor_8fda6fe3c8d61939, []int{0}
}

func (m *PeerAuthenticationService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeerAuthenticationService.Unmarshal(m, b)
}
func (m *PeerAuthenticationService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeerAuthenticationService.Marshal(b, m, deterministic)
}
func (m *PeerAuthenticationService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeerAuthenticationService.Merge(m, src)
}
func (m *PeerAuthenticationService) XXX_Size() int {
	return xxx_messageInfo_PeerAuthenticationService.Size(m)
}
func (m *PeerAuthenticationService) XXX_DiscardUnknown() {
	xxx_messageInfo_PeerAuthenticationService.DiscardUnknown(m)
}

var xxx_messageInfo_PeerAuthenticationService proto.InternalMessageInfo

func (m *PeerAuthenticationService) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *PeerAuthenticationService) GetServiceId() string {
	if m != nil {
		return m.ServiceId
	}
	return ""
}

func (m *PeerAuthenticationService) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PeerAuthenticationService) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *PeerAuthenticationService) GetServiceType() string {
	if m != nil {
		return m.ServiceType
	}
	return ""
}

func (m *PeerAuthenticationService) GetServiceSubType() string {
	if m != nil {
		return m.ServiceSubType
	}
	return ""
}

func (m *PeerAuthenticationService) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *PeerAuthenticationService) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *PeerAuthenticationService) GetCompanyId() string {
	if m != nil {
		return m.CompanyId
	}
	return ""
}

func (m *PeerAuthenticationService) GetIsDiscovered() bool {
	if m != nil {
		return m.IsDiscovered
	}
	return false
}

func (m *PeerAuthenticationService) GetServiceAttributes() *PeerAuthenticationServiceAttributes {
	if m != nil {
		return m.ServiceAttributes
	}
	return nil
}

type PeerAuthenticationServiceAttributes struct {
	Labels               map[string]string `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	TlsMode              TlsMode           `protobuf:"varint,2,opt,name=tls_mode,json=tlsMode,proto3,enum=proto.TlsMode" json:"tls_mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *PeerAuthenticationServiceAttributes) Reset()         { *m = PeerAuthenticationServiceAttributes{} }
func (m *PeerAuthenticationServiceAttributes) String() string { return proto.CompactTextString(m) }
func (*PeerAuthenticationServiceAttributes) ProtoMessage()    {}
func (*PeerAuthenticationServiceAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_8fda6fe3c8d61939, []int{1}
}

func (m *PeerAuthenticationServiceAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeerAuthenticationServiceAttributes.Unmarshal(m, b)
}
func (m *PeerAuthenticationServiceAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeerAuthenticationServiceAttributes.Marshal(b, m, deterministic)
}
func (m *PeerAuthenticationServiceAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeerAuthenticationServiceAttributes.Merge(m, src)
}
func (m *PeerAuthenticationServiceAttributes) XXX_Size() int {
	return xxx_messageInfo_PeerAuthenticationServiceAttributes.Size(m)
}
func (m *PeerAuthenticationServiceAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_PeerAuthenticationServiceAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_PeerAuthenticationServiceAttributes proto.InternalMessageInfo

func (m *PeerAuthenticationServiceAttributes) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *PeerAuthenticationServiceAttributes) GetTlsMode() TlsMode {
	if m != nil {
		return m.TlsMode
	}
	return TlsMode_STRICT
}

func init() {
	proto.RegisterEnum("proto.TlsMode", TlsMode_name, TlsMode_value)
	proto.RegisterType((*PeerAuthenticationService)(nil), "proto.PeerAuthenticationService")
	proto.RegisterType((*PeerAuthenticationServiceAttributes)(nil), "proto.PeerAuthenticationServiceAttributes")
	proto.RegisterMapType((map[string]string)(nil), "proto.PeerAuthenticationServiceAttributes.LabelsEntry")
}

func init() {
	proto.RegisterFile("peerauthenticaion.proto", fileDescriptor_8fda6fe3c8d61939)
}

var fileDescriptor_8fda6fe3c8d61939 = []byte{
	// 499 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x52, 0xcb, 0x6a, 0xdb, 0x4c,
	0x14, 0x8e, 0xad, 0xf8, 0x76, 0x94, 0x18, 0xfd, 0xc3, 0xdf, 0x66, 0x6a, 0x5a, 0x48, 0x9d, 0x4d,
	0x9a, 0x85, 0x17, 0x2e, 0x94, 0xb6, 0xd0, 0x85, 0x1b, 0x8b, 0x22, 0x48, 0x82, 0x91, 0x94, 0x40,
	0x57, 0x66, 0x24, 0x1d, 0x88, 0x1a, 0x59, 0x12, 0x9a, 0x91, 0xc1, 0x4f, 0xd4, 0x27, 0xe9, 0x53,
	0xf4, 0x65, 0x3a, 0x9a, 0x19, 0x9b, 0xd2, 0x0b, 0xcd, 0xc2, 0x2b, 0xcd, 0x7c, 0xb7, 0x73, 0xd1,
	0xc0, 0x49, 0x89, 0x58, 0xb1, 0x5a, 0xdc, 0x63, 0x2e, 0xd2, 0x98, 0xa5, 0x45, 0x3e, 0x29, 0xab,
	0x42, 0x14, 0xa4, 0xa3, 0x3e, 0xa3, 0x63, 0x8e, 0xd5, 0x3a, 0x8d, 0x51, 0xa3, 0xe3, 0xaf, 0x16,
	0x3c, 0x5b, 0x48, 0xc7, 0x6c, 0xe7, 0x10, 0xd2, 0x12, 0x68, 0x0d, 0x79, 0x01, 0x20, 0x65, 0x5f,
	0x30, 0x16, 0xcb, 0x34, 0xa1, 0xad, 0xd3, 0xd6, 0xf9, 0xc0, 0x1f, 0x18, 0xc4, 0x4b, 0x1a, 0xda,
	0xa4, 0x35, 0x74, 0x5b, 0xd3, 0x06, 0x91, 0x34, 0x81, 0xc3, 0x9c, 0xad, 0x90, 0x5a, 0x8a, 0x50,
	0x67, 0x42, 0xa1, 0xb7, 0xc6, 0x8a, 0xcb, 0x1a, 0xf4, 0x50, 0xc1, 0xdb, 0x2b, 0x79, 0x09, 0x47,
	0xdb, 0x30, 0xb1, 0x29, 0x91, 0x76, 0x14, 0x6d, 0x1b, 0x2c, 0x94, 0x10, 0x39, 0x07, 0x67, 0x2b,
	0xe1, 0x75, 0xa4, 0x65, 0x5d, 0x25, 0x1b, 0x1a, 0x3c, 0xa8, 0x23, 0xa5, 0x7c, 0x0e, 0x83, 0xa6,
	0x1c, 0x2f, 0x59, 0x8c, 0xb4, 0xa7, 0x1b, 0xdb, 0x01, 0xe4, 0x7f, 0xe8, 0x88, 0xe2, 0x01, 0x73,
	0xda, 0x57, 0x8c, 0xbe, 0x34, 0xd3, 0xc4, 0xc5, 0xaa, 0x64, 0xf9, 0xa6, 0x99, 0x66, 0xa0, 0x4d,
	0x06, 0x91, 0xd3, 0x9c, 0xc1, 0x71, 0xca, 0x97, 0x49, 0xca, 0xe3, 0x42, 0xb6, 0x8c, 0x09, 0x05,
	0xa9, 0xe8, 0xfb, 0x47, 0x29, 0x9f, 0xef, 0x30, 0xf2, 0x19, 0xc8, 0xb6, 0x43, 0x26, 0x44, 0x95,
	0x46, 0xb5, 0x40, 0x4e, 0x6d, 0xa9, 0xb4, 0xa7, 0x17, 0x7a, 0xe5, 0x93, 0xbf, 0xae, 0x7b, 0xb6,
	0x73, 0xf8, 0xff, 0xf1, 0x5f, 0xa1, 0xf1, 0xf7, 0x16, 0x9c, 0x3d, 0xc2, 0x4a, 0x6e, 0xa0, 0x9b,
	0xb1, 0x08, 0x33, 0x2e, 0xff, 0x97, 0x25, 0xcb, 0xbe, 0x79, 0x7c, 0xd9, 0xc9, 0x95, 0x32, 0xba,
	0xb9, 0xa8, 0x36, 0xbe, 0x49, 0x21, 0xaf, 0xa0, 0x2f, 0x32, 0xbe, 0x5c, 0x15, 0x09, 0xaa, 0x5f,
	0x3c, 0x9c, 0x0e, 0x4d, 0x62, 0x98, 0xf1, 0x6b, 0x89, 0xfa, 0x3d, 0xa1, 0x0f, 0xa3, 0x77, 0x60,
	0xff, 0x94, 0x40, 0x1c, 0xb0, 0x1e, 0x70, 0x63, 0x9e, 0x4d, 0x73, 0x6c, 0x16, 0xbf, 0x66, 0x59,
	0x8d, 0xe6, 0xad, 0xe8, 0xcb, 0xfb, 0xf6, 0xdb, 0xd6, 0xc5, 0x07, 0xe8, 0x99, 0x38, 0x02, 0xd0,
	0x0d, 0x42, 0xdf, 0xbb, 0x0c, 0x9d, 0x03, 0x32, 0x04, 0x58, 0xb8, 0xfe, 0xb5, 0x17, 0x04, 0xde,
	0x9d, 0xeb, 0xb4, 0x88, 0x0d, 0xbd, 0xb9, 0x17, 0xcc, 0x3e, 0x5e, 0xb9, 0x4e, 0x9b, 0x0c, 0xa0,
	0x73, 0x7b, 0x13, 0xb8, 0xa1, 0x63, 0x4d, 0xbf, 0x59, 0x40, 0x7e, 0x1f, 0x90, 0xdc, 0x01, 0xbd,
	0xac, 0x90, 0x09, 0xfc, 0x03, 0x77, 0xfa, 0xaf, 0xbd, 0x8c, 0x9e, 0x1a, 0x85, 0xb9, 0xfb, 0xf2,
	0xf5, 0x14, 0x39, 0xc7, 0xf1, 0x41, 0x93, 0x3b, 0xc7, 0x0c, 0xf7, 0x9e, 0x1b, 0xc0, 0x93, 0x4f,
	0x28, 0xf6, 0x1c, 0x7a, 0x0b, 0x27, 0x0b, 0x26, 0xe2, 0xfb, 0xfd, 0xf7, 0xba, 0xa8, 0xf7, 0xdc,
	0x6b, 0xd4, 0x55, 0xc4, 0xeb, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd7, 0xe6, 0xae, 0xc9, 0xc6,
	0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PeerAuthenticationClient is the client API for PeerAuthentication service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PeerAuthenticationClient interface {
	CreatePeerAuthentication(ctx context.Context, in *PeerAuthenticationService, opts ...grpc.CallOption) (*ServiceResponse, error)
	DeletePeerAuthentication(ctx context.Context, in *PeerAuthenticationService, opts ...grpc.CallOption) (*ServiceResponse, error)
	GetPeerAuthentication(ctx context.Context, in *PeerAuthenticationService, opts ...grpc.CallOption) (*ServiceResponse, error)
	PatchPeerAuthentication(ctx context.Context, in *PeerAuthenticationService, opts ...grpc.CallOption) (*ServiceResponse, error)
	PutPeerAuthentication(ctx context.Context, in *PeerAuthenticationService, opts ...grpc.CallOption) (*ServiceResponse, error)
}

type peerAuthenticationClient struct {
	cc grpc.ClientConnInterface
}

func NewPeerAuthenticationClient(cc grpc.ClientConnInterface) PeerAuthenticationClient {
	return &peerAuthenticationClient{cc}
}

func (c *peerAuthenticationClient) CreatePeerAuthentication(ctx context.Context, in *PeerAuthenticationService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.PeerAuthentication/CreatePeerAuthentication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerAuthenticationClient) DeletePeerAuthentication(ctx context.Context, in *PeerAuthenticationService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.PeerAuthentication/DeletePeerAuthentication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerAuthenticationClient) GetPeerAuthentication(ctx context.Context, in *PeerAuthenticationService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.PeerAuthentication/GetPeerAuthentication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerAuthenticationClient) PatchPeerAuthentication(ctx context.Context, in *PeerAuthenticationService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.PeerAuthentication/PatchPeerAuthentication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerAuthenticationClient) PutPeerAuthentication(ctx context.Context, in *PeerAuthenticationService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.PeerAuthentication/PutPeerAuthentication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PeerAuthenticationServer is the server API for PeerAuthentication service.
type PeerAuthenticationServer interface {
	CreatePeerAuthentication(context.Context, *PeerAuthenticationService) (*ServiceResponse, error)
	DeletePeerAuthentication(context.Context, *PeerAuthenticationService) (*ServiceResponse, error)
	GetPeerAuthentication(context.Context, *PeerAuthenticationService) (*ServiceResponse, error)
	PatchPeerAuthentication(context.Context, *PeerAuthenticationService) (*ServiceResponse, error)
	PutPeerAuthentication(context.Context, *PeerAuthenticationService) (*ServiceResponse, error)
}

// UnimplementedPeerAuthenticationServer can be embedded to have forward compatible implementations.
type UnimplementedPeerAuthenticationServer struct {
}

func (*UnimplementedPeerAuthenticationServer) CreatePeerAuthentication(ctx context.Context, req *PeerAuthenticationService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePeerAuthentication not implemented")
}
func (*UnimplementedPeerAuthenticationServer) DeletePeerAuthentication(ctx context.Context, req *PeerAuthenticationService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePeerAuthentication not implemented")
}
func (*UnimplementedPeerAuthenticationServer) GetPeerAuthentication(ctx context.Context, req *PeerAuthenticationService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPeerAuthentication not implemented")
}
func (*UnimplementedPeerAuthenticationServer) PatchPeerAuthentication(ctx context.Context, req *PeerAuthenticationService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchPeerAuthentication not implemented")
}
func (*UnimplementedPeerAuthenticationServer) PutPeerAuthentication(ctx context.Context, req *PeerAuthenticationService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutPeerAuthentication not implemented")
}

func RegisterPeerAuthenticationServer(s *grpc.Server, srv PeerAuthenticationServer) {
	s.RegisterService(&_PeerAuthentication_serviceDesc, srv)
}

func _PeerAuthentication_CreatePeerAuthentication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeerAuthenticationService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerAuthenticationServer).CreatePeerAuthentication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PeerAuthentication/CreatePeerAuthentication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerAuthenticationServer).CreatePeerAuthentication(ctx, req.(*PeerAuthenticationService))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeerAuthentication_DeletePeerAuthentication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeerAuthenticationService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerAuthenticationServer).DeletePeerAuthentication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PeerAuthentication/DeletePeerAuthentication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerAuthenticationServer).DeletePeerAuthentication(ctx, req.(*PeerAuthenticationService))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeerAuthentication_GetPeerAuthentication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeerAuthenticationService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerAuthenticationServer).GetPeerAuthentication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PeerAuthentication/GetPeerAuthentication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerAuthenticationServer).GetPeerAuthentication(ctx, req.(*PeerAuthenticationService))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeerAuthentication_PatchPeerAuthentication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeerAuthenticationService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerAuthenticationServer).PatchPeerAuthentication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PeerAuthentication/PatchPeerAuthentication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerAuthenticationServer).PatchPeerAuthentication(ctx, req.(*PeerAuthenticationService))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeerAuthentication_PutPeerAuthentication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeerAuthenticationService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerAuthenticationServer).PutPeerAuthentication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PeerAuthentication/PutPeerAuthentication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerAuthenticationServer).PutPeerAuthentication(ctx, req.(*PeerAuthenticationService))
	}
	return interceptor(ctx, in, info, handler)
}

var _PeerAuthentication_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.PeerAuthentication",
	HandlerType: (*PeerAuthenticationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePeerAuthentication",
			Handler:    _PeerAuthentication_CreatePeerAuthentication_Handler,
		},
		{
			MethodName: "DeletePeerAuthentication",
			Handler:    _PeerAuthentication_DeletePeerAuthentication_Handler,
		},
		{
			MethodName: "GetPeerAuthentication",
			Handler:    _PeerAuthentication_GetPeerAuthentication_Handler,
		},
		{
			MethodName: "PatchPeerAuthentication",
			Handler:    _PeerAuthentication_PatchPeerAuthentication_Handler,
		},
		{
			MethodName: "PutPeerAuthentication",
			Handler:    _PeerAuthentication_PutPeerAuthentication_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "peerauthenticaion.proto",
}
