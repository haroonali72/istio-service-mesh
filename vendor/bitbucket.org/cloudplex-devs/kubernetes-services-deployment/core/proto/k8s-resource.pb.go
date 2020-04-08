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

type KubernetesResourceRequest struct {
	ProjectId            string   `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	CompanyId            string   `protobuf:"bytes,2,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	Token                string   `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	Command              string   `protobuf:"bytes,4,opt,name=command,proto3" json:"command,omitempty"`
	Namespaces           []string `protobuf:"bytes,5,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
	Args                 []string `protobuf:"bytes,6,rep,name=args,proto3" json:"args,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KubernetesResourceRequest) Reset()         { *m = KubernetesResourceRequest{} }
func (m *KubernetesResourceRequest) String() string { return proto.CompactTextString(m) }
func (*KubernetesResourceRequest) ProtoMessage()    {}
func (*KubernetesResourceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fe78c8ff5de7b42, []int{0}
}

func (m *KubernetesResourceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KubernetesResourceRequest.Unmarshal(m, b)
}
func (m *KubernetesResourceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KubernetesResourceRequest.Marshal(b, m, deterministic)
}
func (m *KubernetesResourceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KubernetesResourceRequest.Merge(m, src)
}
func (m *KubernetesResourceRequest) XXX_Size() int {
	return xxx_messageInfo_KubernetesResourceRequest.Size(m)
}
func (m *KubernetesResourceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_KubernetesResourceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_KubernetesResourceRequest proto.InternalMessageInfo

func (m *KubernetesResourceRequest) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *KubernetesResourceRequest) GetCompanyId() string {
	if m != nil {
		return m.CompanyId
	}
	return ""
}

func (m *KubernetesResourceRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *KubernetesResourceRequest) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *KubernetesResourceRequest) GetNamespaces() []string {
	if m != nil {
		return m.Namespaces
	}
	return nil
}

func (m *KubernetesResourceRequest) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

type KubernetesResourceResponse struct {
	Resource             []byte   `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KubernetesResourceResponse) Reset()         { *m = KubernetesResourceResponse{} }
func (m *KubernetesResourceResponse) String() string { return proto.CompactTextString(m) }
func (*KubernetesResourceResponse) ProtoMessage()    {}
func (*KubernetesResourceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fe78c8ff5de7b42, []int{1}
}

func (m *KubernetesResourceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KubernetesResourceResponse.Unmarshal(m, b)
}
func (m *KubernetesResourceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KubernetesResourceResponse.Marshal(b, m, deterministic)
}
func (m *KubernetesResourceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KubernetesResourceResponse.Merge(m, src)
}
func (m *KubernetesResourceResponse) XXX_Size() int {
	return xxx_messageInfo_KubernetesResourceResponse.Size(m)
}
func (m *KubernetesResourceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_KubernetesResourceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_KubernetesResourceResponse proto.InternalMessageInfo

func (m *KubernetesResourceResponse) GetResource() []byte {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *KubernetesResourceResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*KubernetesResourceRequest)(nil), "proto.KubernetesResourceRequest")
	proto.RegisterType((*KubernetesResourceResponse)(nil), "proto.KubernetesResourceResponse")
}

func init() {
	proto.RegisterFile("k8s-resource.proto", fileDescriptor_6fe78c8ff5de7b42)
}

var fileDescriptor_6fe78c8ff5de7b42 = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x09, 0x6d, 0x0a, 0x39, 0x10, 0xc3, 0x89, 0xc1, 0x44, 0x02, 0x85, 0x4c, 0x5d, 0xe8,
	0x00, 0x4b, 0xdf, 0x00, 0x55, 0x95, 0x18, 0xbc, 0x30, 0x22, 0xd7, 0x39, 0x10, 0x44, 0xb1, 0x8d,
	0xcf, 0x19, 0x78, 0x38, 0xde, 0x0d, 0xc5, 0x4e, 0xab, 0x0c, 0xc0, 0x64, 0xff, 0xff, 0x77, 0xb2,
	0x3e, 0x1f, 0x60, 0xbb, 0xe6, 0x3b, 0x4f, 0x6c, 0x7b, 0xaf, 0x69, 0xe5, 0xbc, 0x0d, 0x16, 0xf3,
	0x78, 0xd4, 0xdf, 0x19, 0x5c, 0x6d, 0xfb, 0x1d, 0x79, 0x43, 0x81, 0x58, 0x8e, 0x33, 0x92, 0x3e,
	0x7b, 0xe2, 0x80, 0xd7, 0x00, 0xce, 0xdb, 0x0f, 0xd2, 0xe1, 0xe5, 0xbd, 0x11, 0x59, 0x95, 0x2d,
	0x0b, 0x59, 0x8c, 0xcd, 0xa6, 0x19, 0xb0, 0xb6, 0x9d, 0x53, 0xe6, 0x6b, 0xc0, 0xc7, 0x09, 0x8f,
	0xcd, 0xa6, 0xc1, 0x4b, 0xc8, 0x83, 0x6d, 0xc9, 0x88, 0x59, 0x24, 0x29, 0xa0, 0x80, 0x13, 0x6d,
	0xbb, 0x4e, 0x99, 0x46, 0xcc, 0x63, 0xbf, 0x8f, 0x78, 0x03, 0x60, 0x54, 0x47, 0xec, 0x94, 0x26,
	0x16, 0x79, 0x35, 0x5b, 0x16, 0x72, 0xd2, 0x20, 0xc2, 0x5c, 0xf9, 0x37, 0x16, 0x8b, 0x48, 0xe2,
	0xbd, 0x7e, 0x82, 0xf2, 0x37, 0x7d, 0x76, 0xd6, 0x30, 0x61, 0x09, 0xa7, 0xfb, 0x6f, 0x47, 0xfb,
	0x73, 0x79, 0xc8, 0x83, 0x1d, 0x79, 0x6f, 0xfd, 0xe8, 0x9d, 0xc2, 0xfd, 0x2b, 0x9c, 0xb5, 0xeb,
	0xc3, 0x43, 0xf8, 0x0c, 0x17, 0x8f, 0x14, 0xb6, 0x93, 0xa6, 0x4a, 0xfb, 0x5b, 0xfd, 0xb9, 0xb4,
	0xf2, 0xf6, 0x9f, 0x89, 0xe4, 0x55, 0x1f, 0xed, 0x16, 0x71, 0xe6, 0xe1, 0x27, 0x00, 0x00, 0xff,
	0xff, 0x63, 0x62, 0xe3, 0xbf, 0x9b, 0x01, 0x00, 0x00,
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
	GetK8SResource(ctx context.Context, in *KubernetesResourceRequest, opts ...grpc.CallOption) (*KubernetesResourceResponse, error)
}

type k8SResourceClient struct {
	cc grpc.ClientConnInterface
}

func NewK8SResourceClient(cc grpc.ClientConnInterface) K8SResourceClient {
	return &k8SResourceClient{cc}
}

func (c *k8SResourceClient) GetK8SResource(ctx context.Context, in *KubernetesResourceRequest, opts ...grpc.CallOption) (*KubernetesResourceResponse, error) {
	out := new(KubernetesResourceResponse)
	err := c.cc.Invoke(ctx, "/proto.k8sResource/GetK8sResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// K8SResourceServer is the server API for K8SResource service.
type K8SResourceServer interface {
	GetK8SResource(context.Context, *KubernetesResourceRequest) (*KubernetesResourceResponse, error)
}

// UnimplementedK8SResourceServer can be embedded to have forward compatible implementations.
type UnimplementedK8SResourceServer struct {
}

func (*UnimplementedK8SResourceServer) GetK8SResource(ctx context.Context, req *KubernetesResourceRequest) (*KubernetesResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetK8SResource not implemented")
}

func RegisterK8SResourceServer(s *grpc.Server, srv K8SResourceServer) {
	s.RegisterService(&_K8SResource_serviceDesc, srv)
}

func _K8SResource_GetK8SResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KubernetesResourceRequest)
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
		return srv.(K8SResourceServer).GetK8SResource(ctx, req.(*KubernetesResourceRequest))
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