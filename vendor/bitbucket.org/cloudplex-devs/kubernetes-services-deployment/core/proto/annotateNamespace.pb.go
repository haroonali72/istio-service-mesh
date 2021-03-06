// Code generated by protoc-gen-go. DO NOT EDIT.
// source: annotateNamespace.proto

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

type Namespacerequest struct {
	InfraId              string   `protobuf:"bytes,1,opt,name=infra_id,json=infraId,proto3" json:"infra_id,omitempty"`
	CompanyId            string   `protobuf:"bytes,2,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	Token                string   `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	Namespace            string   `protobuf:"bytes,4,opt,name=namespace,proto3" json:"namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Namespacerequest) Reset()         { *m = Namespacerequest{} }
func (m *Namespacerequest) String() string { return proto.CompactTextString(m) }
func (*Namespacerequest) ProtoMessage()    {}
func (*Namespacerequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a309dda261f6d9c7, []int{0}
}

func (m *Namespacerequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Namespacerequest.Unmarshal(m, b)
}
func (m *Namespacerequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Namespacerequest.Marshal(b, m, deterministic)
}
func (m *Namespacerequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Namespacerequest.Merge(m, src)
}
func (m *Namespacerequest) XXX_Size() int {
	return xxx_messageInfo_Namespacerequest.Size(m)
}
func (m *Namespacerequest) XXX_DiscardUnknown() {
	xxx_messageInfo_Namespacerequest.DiscardUnknown(m)
}

var xxx_messageInfo_Namespacerequest proto.InternalMessageInfo

func (m *Namespacerequest) GetInfraId() string {
	if m != nil {
		return m.InfraId
	}
	return ""
}

func (m *Namespacerequest) GetCompanyId() string {
	if m != nil {
		return m.CompanyId
	}
	return ""
}

func (m *Namespacerequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Namespacerequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

type Namespaceresponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Namespaceresponse) Reset()         { *m = Namespaceresponse{} }
func (m *Namespaceresponse) String() string { return proto.CompactTextString(m) }
func (*Namespaceresponse) ProtoMessage()    {}
func (*Namespaceresponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a309dda261f6d9c7, []int{1}
}

func (m *Namespaceresponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Namespaceresponse.Unmarshal(m, b)
}
func (m *Namespaceresponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Namespaceresponse.Marshal(b, m, deterministic)
}
func (m *Namespaceresponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Namespaceresponse.Merge(m, src)
}
func (m *Namespaceresponse) XXX_Size() int {
	return xxx_messageInfo_Namespaceresponse.Size(m)
}
func (m *Namespaceresponse) XXX_DiscardUnknown() {
	xxx_messageInfo_Namespaceresponse.DiscardUnknown(m)
}

var xxx_messageInfo_Namespaceresponse proto.InternalMessageInfo

func (m *Namespaceresponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Namespacerequest)(nil), "proto.Namespacerequest")
	proto.RegisterType((*Namespaceresponse)(nil), "proto.Namespaceresponse")
}

func init() {
	proto.RegisterFile("annotateNamespace.proto", fileDescriptor_a309dda261f6d9c7)
}

var fileDescriptor_a309dda261f6d9c7 = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0xcc, 0xcb, 0xcb,
	0x2f, 0x49, 0x2c, 0x49, 0xf5, 0x4b, 0xcc, 0x4d, 0x2d, 0x2e, 0x48, 0x4c, 0x4e, 0xd5, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x0d, 0x8c, 0x5c, 0x02, 0x70, 0xa9, 0xa2, 0xd4,
	0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x49, 0x2e, 0x8e, 0xcc, 0xbc, 0xb4, 0xa2, 0xc4, 0xf8, 0xcc,
	0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x76, 0x30, 0xdf, 0x33, 0x45, 0x48, 0x96, 0x8b,
	0x2b, 0x39, 0x3f, 0xb7, 0x20, 0x31, 0xaf, 0x12, 0x24, 0xc9, 0x04, 0x96, 0xe4, 0x84, 0x8a, 0x00,
	0xa5, 0x45, 0xb8, 0x58, 0x4b, 0xf2, 0xb3, 0x53, 0xf3, 0x24, 0x98, 0xc1, 0x32, 0x10, 0x8e, 0x90,
	0x0c, 0x17, 0x67, 0x1e, 0xcc, 0x0e, 0x09, 0x16, 0x88, 0x1e, 0xb8, 0x80, 0x92, 0x2e, 0x97, 0x20,
	0x92, 0x0b, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x24, 0xb8, 0xd8, 0x81, 0x42, 0xc5, 0x89,
	0xe9, 0xa9, 0x30, 0x17, 0x40, 0xb9, 0x46, 0xb1, 0x5c, 0x82, 0x8e, 0xe8, 0x7e, 0x12, 0xf2, 0xc0,
	0x26, 0x28, 0x0e, 0xf1, 0xaa, 0x1e, 0xba, 0xff, 0xa4, 0x24, 0x30, 0x25, 0x20, 0xd6, 0x2a, 0x31,
	0x24, 0xb1, 0x81, 0xa5, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc8, 0x7a, 0xee, 0xe2, 0x39,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AnnotateNamespaceClient is the client API for AnnotateNamespace service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AnnotateNamespaceClient interface {
	AnnotateNamespace(ctx context.Context, in *Namespacerequest, opts ...grpc.CallOption) (*Namespaceresponse, error)
}

type annotateNamespaceClient struct {
	cc grpc.ClientConnInterface
}

func NewAnnotateNamespaceClient(cc grpc.ClientConnInterface) AnnotateNamespaceClient {
	return &annotateNamespaceClient{cc}
}

func (c *annotateNamespaceClient) AnnotateNamespace(ctx context.Context, in *Namespacerequest, opts ...grpc.CallOption) (*Namespaceresponse, error) {
	out := new(Namespaceresponse)
	err := c.cc.Invoke(ctx, "/proto.AnnotateNamespace/AnnotateNamespace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AnnotateNamespaceServer is the server API for AnnotateNamespace service.
type AnnotateNamespaceServer interface {
	AnnotateNamespace(context.Context, *Namespacerequest) (*Namespaceresponse, error)
}

// UnimplementedAnnotateNamespaceServer can be embedded to have forward compatible implementations.
type UnimplementedAnnotateNamespaceServer struct {
}

func (*UnimplementedAnnotateNamespaceServer) AnnotateNamespace(ctx context.Context, req *Namespacerequest) (*Namespaceresponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AnnotateNamespace not implemented")
}

func RegisterAnnotateNamespaceServer(s *grpc.Server, srv AnnotateNamespaceServer) {
	s.RegisterService(&_AnnotateNamespace_serviceDesc, srv)
}

func _AnnotateNamespace_AnnotateNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Namespacerequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnnotateNamespaceServer).AnnotateNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AnnotateNamespace/AnnotateNamespace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnnotateNamespaceServer).AnnotateNamespace(ctx, req.(*Namespacerequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AnnotateNamespace_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AnnotateNamespace",
	HandlerType: (*AnnotateNamespaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AnnotateNamespace",
			Handler:    _AnnotateNamespace_AnnotateNamespace_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "annotateNamespace.proto",
}
