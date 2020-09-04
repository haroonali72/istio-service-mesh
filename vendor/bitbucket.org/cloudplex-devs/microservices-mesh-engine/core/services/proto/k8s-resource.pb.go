// Code generated by protoc-gen-go. DO NOT EDIT.
// source: k8s-resource.proto

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

type K8SResourceRequest struct {
	ApplicationId        string   `protobuf:"bytes,1,opt,name=application_id,json=applicationId,proto3" json:"application_id,omitempty"`
	CompanyId            string   `protobuf:"bytes,2,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	Token                string   `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	Command              string   `protobuf:"bytes,4,opt,name=command,proto3" json:"command,omitempty"`
	Namespaces           []string `protobuf:"bytes,5,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
	Args                 []string `protobuf:"bytes,6,rep,name=args,proto3" json:"args,omitempty"`
	InfraId              string   `protobuf:"bytes,7,opt,name=infra_id,json=infraId,proto3" json:"infra_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *K8SResourceRequest) Reset()         { *m = K8SResourceRequest{} }
func (m *K8SResourceRequest) String() string { return proto.CompactTextString(m) }
func (*K8SResourceRequest) ProtoMessage()    {}
func (*K8SResourceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fe78c8ff5de7b42, []int{0}
}

func (m *K8SResourceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_K8SResourceRequest.Unmarshal(m, b)
}
func (m *K8SResourceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_K8SResourceRequest.Marshal(b, m, deterministic)
}
func (m *K8SResourceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_K8SResourceRequest.Merge(m, src)
}
func (m *K8SResourceRequest) XXX_Size() int {
	return xxx_messageInfo_K8SResourceRequest.Size(m)
}
func (m *K8SResourceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_K8SResourceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_K8SResourceRequest proto.InternalMessageInfo

func (m *K8SResourceRequest) GetApplicationId() string {
	if m != nil {
		return m.ApplicationId
	}
	return ""
}

func (m *K8SResourceRequest) GetCompanyId() string {
	if m != nil {
		return m.CompanyId
	}
	return ""
}

func (m *K8SResourceRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *K8SResourceRequest) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *K8SResourceRequest) GetNamespaces() []string {
	if m != nil {
		return m.Namespaces
	}
	return nil
}

func (m *K8SResourceRequest) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *K8SResourceRequest) GetInfraId() string {
	if m != nil {
		return m.InfraId
	}
	return ""
}

type K8SResourceResponse struct {
	Resource             []byte   `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *K8SResourceResponse) Reset()         { *m = K8SResourceResponse{} }
func (m *K8SResourceResponse) String() string { return proto.CompactTextString(m) }
func (*K8SResourceResponse) ProtoMessage()    {}
func (*K8SResourceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fe78c8ff5de7b42, []int{1}
}

func (m *K8SResourceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_K8SResourceResponse.Unmarshal(m, b)
}
func (m *K8SResourceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_K8SResourceResponse.Marshal(b, m, deterministic)
}
func (m *K8SResourceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_K8SResourceResponse.Merge(m, src)
}
func (m *K8SResourceResponse) XXX_Size() int {
	return xxx_messageInfo_K8SResourceResponse.Size(m)
}
func (m *K8SResourceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_K8SResourceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_K8SResourceResponse proto.InternalMessageInfo

func (m *K8SResourceResponse) GetResource() []byte {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *K8SResourceResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*K8SResourceRequest)(nil), "proto.K8sResourceRequest")
	proto.RegisterType((*K8SResourceResponse)(nil), "proto.K8sResourceResponse")
}

func init() {
	proto.RegisterFile("k8s-resource.proto", fileDescriptor_6fe78c8ff5de7b42)
}

var fileDescriptor_6fe78c8ff5de7b42 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x90, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0x09, 0x6d, 0xfa, 0xe7, 0x80, 0x0e, 0x07, 0x83, 0x1b, 0x09, 0x54, 0x45, 0x42, 0x62,
	0xa1, 0x03, 0x2c, 0xfd, 0x06, 0x55, 0xc4, 0x96, 0x89, 0x0d, 0x99, 0xe4, 0x40, 0x51, 0x88, 0x6d,
	0x6c, 0x77, 0xe0, 0xa3, 0xf2, 0x6d, 0x48, 0xce, 0x09, 0x04, 0xc1, 0x64, 0xbf, 0xf7, 0x3b, 0x9d,
	0xee, 0x3d, 0xc0, 0x7a, 0xe7, 0x6e, 0x2d, 0x39, 0x7d, 0xb0, 0x05, 0x6d, 0x8d, 0xd5, 0x5e, 0x63,
	0xcc, 0x4f, 0xfa, 0x19, 0x01, 0x3e, 0xec, 0x5c, 0xde, 0xc3, 0x9c, 0xde, 0x0f, 0xe4, 0x3c, 0x5e,
	0xc3, 0x4a, 0x1a, 0xf3, 0x56, 0x15, 0xd2, 0x57, 0x5a, 0x3d, 0x55, 0xa5, 0x88, 0x36, 0xd1, 0xcd,
	0x32, 0x3f, 0x1b, 0xb9, 0x59, 0x89, 0x97, 0x00, 0x85, 0x6e, 0x8c, 0x54, 0x1f, 0xdd, 0xc8, 0x31,
	0x8f, 0x2c, 0x7b, 0xa7, 0xc5, 0x17, 0x10, 0x7b, 0x5d, 0x93, 0x12, 0x13, 0x26, 0x41, 0xa0, 0x80,
	0x79, 0x3b, 0xd2, 0x48, 0x55, 0x8a, 0x29, 0xfb, 0x83, 0xc4, 0x2b, 0x00, 0x25, 0x1b, 0x72, 0x46,
	0x16, 0xe4, 0x44, 0xbc, 0x99, 0xb4, 0x70, 0xe4, 0x20, 0xc2, 0x54, 0xda, 0x57, 0x27, 0x66, 0x4c,
	0xf8, 0x8f, 0x6b, 0x58, 0x54, 0xea, 0xc5, 0xca, 0xee, 0x80, 0x79, 0x58, 0xc7, 0x3a, 0x2b, 0xd3,
	0x3d, 0x9c, 0xff, 0x8a, 0xe6, 0x8c, 0x56, 0x8e, 0x30, 0x81, 0xc5, 0xd0, 0x05, 0xa7, 0x3a, 0xcd,
	0xbf, 0x75, 0x77, 0x31, 0x59, 0xab, 0x6d, 0x9f, 0x25, 0x88, 0xbb, 0x47, 0x38, 0xa9, 0x7f, 0x16,
	0x61, 0x06, 0xab, 0x3d, 0xf9, 0xd1, 0x6a, 0x5c, 0x87, 0x52, 0xb7, 0x7f, 0x9b, 0x4c, 0x92, 0xff,
	0x50, 0xb8, 0x24, 0x3d, 0x7a, 0x9e, 0x31, 0xbc, 0xff, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x95, 0xdc,
	0xa3, 0x80, 0xa2, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// K8SResourceClient is the client API for K8SResource service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type K8SResourceClient interface {
	GetK8SResource(ctx context.Context, in *K8SResourceRequest, opts ...grpc.CallOption) (*K8SResourceResponse, error)
}

type k8SResourceClient struct {
	cc grpc.ClientConnInterface
}

func NewK8SResourceClient(cc grpc.ClientConnInterface) K8SResourceClient {
	return &k8SResourceClient{cc}
}

func (c *k8SResourceClient) GetK8SResource(ctx context.Context, in *K8SResourceRequest, opts ...grpc.CallOption) (*K8SResourceResponse, error) {
	out := new(K8SResourceResponse)
	err := c.cc.Invoke(ctx, "/proto.k8sResource/GetK8sResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// K8SResourceServer is the server API for K8SResource service.
type K8SResourceServer interface {
	GetK8SResource(context.Context, *K8SResourceRequest) (*K8SResourceResponse, error)
}

// UnimplementedK8SResourceServer can be embedded to have forward compatible implementations.
type UnimplementedK8SResourceServer struct {
}

func (*UnimplementedK8SResourceServer) GetK8SResource(ctx context.Context, req *K8SResourceRequest) (*K8SResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetK8SResource not implemented")
}

func RegisterK8SResourceServer(s *grpc.Server, srv K8SResourceServer) {
	s.RegisterService(&_K8SResource_serviceDesc, srv)
}

func _K8SResource_GetK8SResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(K8SResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(K8SResourceServer).GetK8SResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.k8sResource/GetK8SResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(K8SResourceServer).GetK8SResource(ctx, req.(*K8SResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _K8SResource_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.k8sResource",
	HandlerType: (*K8SResourceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetK8sResource",
			Handler:    _K8SResource_GetK8SResource_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "k8s-resource.proto",
}
