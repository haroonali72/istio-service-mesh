// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kubernetesservice.proto

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

type KubernetesService struct {
	ServiceId             string                 `protobuf:"bytes,1,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	Token                 string                 `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	ProjectId             string                 `protobuf:"bytes,3,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	CompanyId             string                 `protobuf:"bytes,4,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	Name                  string                 `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Version               string                 `protobuf:"bytes,6,opt,name=version,proto3" json:"version,omitempty"`
	ServiceType           string                 `protobuf:"bytes,7,opt,name=service_type,json=serviceType,proto3" json:"service_type,omitempty"`
	ServiceSubType        string                 `protobuf:"bytes,8,opt,name=service_sub_type,json=serviceSubType,proto3" json:"service_sub_type,omitempty"`
	ServiceDependencyInfo *any.Any               `protobuf:"bytes,9,opt,name=service_dependency_info,json=serviceDependencyInfo,proto3" json:"service_dependency_info,omitempty"`
	Namespace             string                 `protobuf:"bytes,10,opt,name=namespace,proto3" json:"namespace,omitempty"`
	KubeServiceAttributes *KubeServiceAttributes `protobuf:"bytes,11,opt,name=kube_service_attributes,json=kubeServiceAttributes,proto3" json:"kube_service_attributes,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}               `json:"-"`
	XXX_unrecognized      []byte                 `json:"-"`
	XXX_sizecache         int32                  `json:"-"`
}

func (m *KubernetesService) Reset()         { *m = KubernetesService{} }
func (m *KubernetesService) String() string { return proto.CompactTextString(m) }
func (*KubernetesService) ProtoMessage()    {}
func (*KubernetesService) Descriptor() ([]byte, []int) {
	return fileDescriptor_02d602f0285fe574, []int{0}
}

func (m *KubernetesService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KubernetesService.Unmarshal(m, b)
}
func (m *KubernetesService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KubernetesService.Marshal(b, m, deterministic)
}
func (m *KubernetesService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KubernetesService.Merge(m, src)
}
func (m *KubernetesService) XXX_Size() int {
	return xxx_messageInfo_KubernetesService.Size(m)
}
func (m *KubernetesService) XXX_DiscardUnknown() {
	xxx_messageInfo_KubernetesService.DiscardUnknown(m)
}

var xxx_messageInfo_KubernetesService proto.InternalMessageInfo

func (m *KubernetesService) GetServiceId() string {
	if m != nil {
		return m.ServiceId
	}
	return ""
}

func (m *KubernetesService) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *KubernetesService) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *KubernetesService) GetCompanyId() string {
	if m != nil {
		return m.CompanyId
	}
	return ""
}

func (m *KubernetesService) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *KubernetesService) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *KubernetesService) GetServiceType() string {
	if m != nil {
		return m.ServiceType
	}
	return ""
}

func (m *KubernetesService) GetServiceSubType() string {
	if m != nil {
		return m.ServiceSubType
	}
	return ""
}

func (m *KubernetesService) GetServiceDependencyInfo() *any.Any {
	if m != nil {
		return m.ServiceDependencyInfo
	}
	return nil
}

func (m *KubernetesService) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *KubernetesService) GetKubeServiceAttributes() *KubeServiceAttributes {
	if m != nil {
		return m.KubeServiceAttributes
	}
	return nil
}

type KubeServiceAttributes struct {
	KubePorts             []*KubePort       `protobuf:"bytes,1,rep,name=kube_ports,json=kubePorts,proto3" json:"kube_ports,omitempty"`
	Selector              map[string]string `protobuf:"bytes,2,rep,name=selector,proto3" json:"selector,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ClusterIp             string            `protobuf:"bytes,3,opt,name=cluster_ip,json=clusterIp,proto3" json:"cluster_ip,omitempty"`
	Type                  string            `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	ExternalTrafficPolicy string            `protobuf:"bytes,5,opt,name=external_traffic_policy,json=externalTrafficPolicy,proto3" json:"external_traffic_policy,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}          `json:"-"`
	XXX_unrecognized      []byte            `json:"-"`
	XXX_sizecache         int32             `json:"-"`
}

func (m *KubeServiceAttributes) Reset()         { *m = KubeServiceAttributes{} }
func (m *KubeServiceAttributes) String() string { return proto.CompactTextString(m) }
func (*KubeServiceAttributes) ProtoMessage()    {}
func (*KubeServiceAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_02d602f0285fe574, []int{1}
}

func (m *KubeServiceAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KubeServiceAttributes.Unmarshal(m, b)
}
func (m *KubeServiceAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KubeServiceAttributes.Marshal(b, m, deterministic)
}
func (m *KubeServiceAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KubeServiceAttributes.Merge(m, src)
}
func (m *KubeServiceAttributes) XXX_Size() int {
	return xxx_messageInfo_KubeServiceAttributes.Size(m)
}
func (m *KubeServiceAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_KubeServiceAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_KubeServiceAttributes proto.InternalMessageInfo

func (m *KubeServiceAttributes) GetKubePorts() []*KubePort {
	if m != nil {
		return m.KubePorts
	}
	return nil
}

func (m *KubeServiceAttributes) GetSelector() map[string]string {
	if m != nil {
		return m.Selector
	}
	return nil
}

func (m *KubeServiceAttributes) GetClusterIp() string {
	if m != nil {
		return m.ClusterIp
	}
	return ""
}

func (m *KubeServiceAttributes) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *KubeServiceAttributes) GetExternalTrafficPolicy() string {
	if m != nil {
		return m.ExternalTrafficPolicy
	}
	return ""
}

type KubePort struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Protocol             string   `protobuf:"bytes,2,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Port                 int64    `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	TargetPort           int64    `protobuf:"varint,4,opt,name=target_port,json=targetPort,proto3" json:"target_port,omitempty"`
	NodePort             int64    `protobuf:"varint,5,opt,name=node_port,json=nodePort,proto3" json:"node_port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KubePort) Reset()         { *m = KubePort{} }
func (m *KubePort) String() string { return proto.CompactTextString(m) }
func (*KubePort) ProtoMessage()    {}
func (*KubePort) Descriptor() ([]byte, []int) {
	return fileDescriptor_02d602f0285fe574, []int{2}
}

func (m *KubePort) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KubePort.Unmarshal(m, b)
}
func (m *KubePort) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KubePort.Marshal(b, m, deterministic)
}
func (m *KubePort) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KubePort.Merge(m, src)
}
func (m *KubePort) XXX_Size() int {
	return xxx_messageInfo_KubePort.Size(m)
}
func (m *KubePort) XXX_DiscardUnknown() {
	xxx_messageInfo_KubePort.DiscardUnknown(m)
}

var xxx_messageInfo_KubePort proto.InternalMessageInfo

func (m *KubePort) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *KubePort) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *KubePort) GetPort() int64 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *KubePort) GetTargetPort() int64 {
	if m != nil {
		return m.TargetPort
	}
	return 0
}

func (m *KubePort) GetNodePort() int64 {
	if m != nil {
		return m.NodePort
	}
	return 0
}

type KubeSelector struct {
	Label1               string   `protobuf:"bytes,1,opt,name=label1,proto3" json:"label1,omitempty"`
	Label2               string   `protobuf:"bytes,2,opt,name=label2,proto3" json:"label2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KubeSelector) Reset()         { *m = KubeSelector{} }
func (m *KubeSelector) String() string { return proto.CompactTextString(m) }
func (*KubeSelector) ProtoMessage()    {}
func (*KubeSelector) Descriptor() ([]byte, []int) {
	return fileDescriptor_02d602f0285fe574, []int{3}
}

func (m *KubeSelector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KubeSelector.Unmarshal(m, b)
}
func (m *KubeSelector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KubeSelector.Marshal(b, m, deterministic)
}
func (m *KubeSelector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KubeSelector.Merge(m, src)
}
func (m *KubeSelector) XXX_Size() int {
	return xxx_messageInfo_KubeSelector.Size(m)
}
func (m *KubeSelector) XXX_DiscardUnknown() {
	xxx_messageInfo_KubeSelector.DiscardUnknown(m)
}

var xxx_messageInfo_KubeSelector proto.InternalMessageInfo

func (m *KubeSelector) GetLabel1() string {
	if m != nil {
		return m.Label1
	}
	return ""
}

func (m *KubeSelector) GetLabel2() string {
	if m != nil {
		return m.Label2
	}
	return ""
}

func init() {
	proto.RegisterType((*KubernetesService)(nil), "proto.KubernetesService")
	proto.RegisterType((*KubeServiceAttributes)(nil), "proto.KubeServiceAttributes")
	proto.RegisterMapType((map[string]string)(nil), "proto.KubeServiceAttributes.SelectorEntry")
	proto.RegisterType((*KubePort)(nil), "proto.KubePort")
	proto.RegisterType((*KubeSelector)(nil), "proto.KubeSelector")
}

func init() { proto.RegisterFile("kubernetesservice.proto", fileDescriptor_02d602f0285fe574) }

var fileDescriptor_02d602f0285fe574 = []byte{
	// 613 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x26, 0x75, 0x7f, 0x92, 0x49, 0x0b, 0x65, 0xd5, 0x34, 0x26, 0x14, 0x01, 0x39, 0x55, 0x1c,
	0x5c, 0x51, 0x24, 0x84, 0x40, 0x42, 0xaa, 0x28, 0xa0, 0x42, 0x91, 0x22, 0xb7, 0xf7, 0xc8, 0x76,
	0x26, 0xc1, 0xc4, 0xdd, 0xb5, 0xd6, 0xeb, 0x08, 0xbf, 0x03, 0x07, 0xde, 0x86, 0x07, 0xe0, 0xc5,
	0xd8, 0x9d, 0x5d, 0x9b, 0x44, 0x54, 0x9c, 0x72, 0xf2, 0xce, 0xf7, 0xcd, 0xb7, 0x33, 0x3b, 0xdf,
	0x18, 0xfa, 0xf3, 0x32, 0x46, 0xc9, 0x51, 0x61, 0x51, 0xa0, 0x5c, 0xa4, 0x09, 0x06, 0xb9, 0x14,
	0x4a, 0xb0, 0x2d, 0xfa, 0x0c, 0xf6, 0x56, 0xd0, 0xc1, 0x83, 0x99, 0x10, 0xb3, 0x0c, 0x4f, 0x28,
	0x8a, 0xcb, 0xe9, 0x49, 0xc4, 0x2b, 0x4b, 0x0d, 0x7f, 0x7b, 0x70, 0xff, 0x73, 0x73, 0xd9, 0x95,
	0x95, 0xb1, 0x47, 0x00, 0xee, 0x86, 0x71, 0x3a, 0xf1, 0x5b, 0x4f, 0x5a, 0xc7, 0x9d, 0xb0, 0xe3,
	0x90, 0x8b, 0x09, 0x3b, 0x80, 0x2d, 0x25, 0xe6, 0xc8, 0xfd, 0x0d, 0x62, 0x6c, 0x60, 0x44, 0xfa,
	0xce, 0x6f, 0x98, 0x28, 0x23, 0xf2, 0xac, 0xc8, 0x21, 0x5a, 0xa4, 0xe9, 0x44, 0xdc, 0xe4, 0xba,
	0xb4, 0xa1, 0x37, 0x2d, 0xed, 0x10, 0x4d, 0x33, 0xd8, 0xe4, 0xd1, 0x0d, 0xfa, 0x5b, 0x44, 0xd0,
	0x99, 0xf9, 0xb0, 0xb3, 0x40, 0x59, 0xa4, 0x82, 0xfb, 0xdb, 0x04, 0xd7, 0x21, 0x7b, 0x0a, 0xbb,
	0x75, 0x83, 0xaa, 0xca, 0xd1, 0xdf, 0x21, 0xba, 0xeb, 0xb0, 0x6b, 0x0d, 0xb1, 0x63, 0xd8, 0xaf,
	0x53, 0x8a, 0x32, 0xb6, 0x69, 0x6d, 0x4a, 0xbb, 0xeb, 0xf0, 0xab, 0x32, 0xa6, 0xcc, 0x4b, 0xe8,
	0xd7, 0x99, 0x13, 0xcc, 0x91, 0x4f, 0x90, 0x27, 0xba, 0x49, 0x3e, 0x15, 0x7e, 0x47, 0x0b, 0xba,
	0xa7, 0x07, 0x81, 0x1d, 0x60, 0x50, 0x0f, 0x30, 0x38, 0xe3, 0x55, 0xd8, 0x73, 0xa2, 0xf3, 0x46,
	0x73, 0xa1, 0x25, 0xec, 0x08, 0x3a, 0xa6, 0xf9, 0x22, 0x8f, 0x12, 0xf4, 0xc1, 0x3e, 0xb3, 0x01,
	0xd8, 0xb5, 0xf5, 0x6e, 0x5c, 0x17, 0x8c, 0x94, 0x92, 0x69, 0x5c, 0xea, 0xd9, 0xfb, 0x5d, 0xaa,
	0x75, 0x64, 0x8b, 0x04, 0xc6, 0x14, 0x67, 0xc7, 0x59, 0x93, 0x13, 0xf6, 0xe6, 0xb7, 0xc1, 0xc3,
	0x5f, 0x1b, 0xd0, 0xbb, 0x55, 0xc0, 0x02, 0x00, 0xaa, 0x97, 0x0b, 0xa9, 0x0a, 0xed, 0xa4, 0xa7,
	0x4b, 0xdc, 0x5b, 0x2a, 0x31, 0xd2, 0x78, 0xd8, 0x99, 0xbb, 0x53, 0xc1, 0x3e, 0x40, 0xbb, 0xc0,
	0x4c, 0x3b, 0x26, 0xa4, 0x76, 0xd7, 0x64, 0x3f, 0xfb, 0x5f, 0x43, 0xc1, 0x95, 0x4b, 0x7e, 0xcf,
	0x95, 0xac, 0xc2, 0x46, 0x4b, 0x6e, 0x67, 0x65, 0xa1, 0x50, 0x8e, 0xd3, 0xbc, 0x5e, 0x06, 0x87,
	0x5c, 0xe4, 0xc6, 0x6d, 0x32, 0xc4, 0xae, 0x01, 0x9d, 0xd9, 0x4b, 0xe8, 0xe3, 0x77, 0x4d, 0xf3,
	0x28, 0x1b, 0x2b, 0x19, 0x4d, 0xa7, 0x69, 0xa2, 0xdb, 0xce, 0xd2, 0xa4, 0x72, 0x4b, 0xd1, 0xab,
	0xe9, 0x6b, 0xcb, 0x8e, 0x88, 0x1c, 0xbc, 0x81, 0xbd, 0x95, 0x2e, 0xd8, 0x3e, 0x78, 0x73, 0xac,
	0xdc, 0xda, 0x9a, 0xa3, 0x59, 0xd8, 0x45, 0x94, 0x95, 0x58, 0x2f, 0x2c, 0x05, 0xaf, 0x37, 0x5e,
	0xb5, 0x86, 0x3f, 0x5a, 0xd0, 0xae, 0xe7, 0xd0, 0xec, 0x60, 0x6b, 0x69, 0x07, 0x07, 0xd0, 0xa6,
	0xf7, 0x27, 0x22, 0x73, 0xea, 0x26, 0x36, 0xf9, 0x66, 0xae, 0xf4, 0x3c, 0x2f, 0xa4, 0x33, 0x7b,
	0x0c, 0x5d, 0x15, 0xc9, 0x19, 0x2a, 0x1a, 0x39, 0x3d, 0xd0, 0x0b, 0xc1, 0x42, 0x54, 0xe4, 0xa1,
	0xde, 0x0f, 0x31, 0xb1, 0x8e, 0xd0, 0xc3, 0xbc, 0xb0, 0x6d, 0x00, 0x43, 0x0e, 0xdf, 0xc2, 0xae,
	0x9d, 0xb3, 0x1b, 0xe3, 0x21, 0x6c, 0x67, 0x51, 0x8c, 0xd9, 0x73, 0xd7, 0x93, 0x8b, 0x1a, 0xfc,
	0xd4, 0xf5, 0xe4, 0xa2, 0xd3, 0x9f, 0x1e, 0xc0, 0xdf, 0xdf, 0x99, 0x7d, 0x81, 0xfe, 0x3b, 0x89,
	0x91, 0xc2, 0x7f, 0x7f, 0x71, 0x7f, 0xc9, 0xd6, 0x15, 0x66, 0x70, 0xe8, 0x18, 0x17, 0x87, 0x7a,
	0x73, 0x05, 0x2f, 0x70, 0x78, 0x87, 0x7d, 0x82, 0x83, 0x8f, 0xa8, 0xd6, 0x76, 0xd7, 0xa8, 0x5c,
	0xd3, 0x5d, 0x97, 0x70, 0x38, 0x8a, 0x54, 0xf2, 0x75, 0x3d, 0xb7, 0xe9, 0xa1, 0x9d, 0xeb, 0xf9,
	0xaf, 0x69, 0x68, 0xf1, 0x36, 0x11, 0x2f, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0xf1, 0x7e, 0x98,
	0x9c, 0xb4, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KubernetesClient is the client API for Kubernetes service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KubernetesClient interface {
	CreateKubernetesService(ctx context.Context, in *KubernetesService, opts ...grpc.CallOption) (*ServiceResponse, error)
	GetKubernetesService(ctx context.Context, in *KubernetesService, opts ...grpc.CallOption) (*ServiceResponse, error)
	PutKubernetesService(ctx context.Context, in *KubernetesService, opts ...grpc.CallOption) (*ServiceResponse, error)
	PatchKubernetesService(ctx context.Context, in *KubernetesService, opts ...grpc.CallOption) (*ServiceResponse, error)
	DeleteKubernetesService(ctx context.Context, in *KubernetesService, opts ...grpc.CallOption) (*ServiceResponse, error)
}

type kubernetesClient struct {
	cc *grpc.ClientConn
}

func NewKubernetesClient(cc *grpc.ClientConn) KubernetesClient {
	return &kubernetesClient{cc}
}

func (c *kubernetesClient) CreateKubernetesService(ctx context.Context, in *KubernetesService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.Kubernetes/CreateKubernetesService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubernetesClient) GetKubernetesService(ctx context.Context, in *KubernetesService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.Kubernetes/GetKubernetesService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubernetesClient) PutKubernetesService(ctx context.Context, in *KubernetesService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.Kubernetes/PutKubernetesService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubernetesClient) PatchKubernetesService(ctx context.Context, in *KubernetesService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.Kubernetes/PatchKubernetesService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kubernetesClient) DeleteKubernetesService(ctx context.Context, in *KubernetesService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.Kubernetes/DeleteKubernetesService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KubernetesServer is the server API for Kubernetes service.
type KubernetesServer interface {
	CreateKubernetesService(context.Context, *KubernetesService) (*ServiceResponse, error)
	GetKubernetesService(context.Context, *KubernetesService) (*ServiceResponse, error)
	PutKubernetesService(context.Context, *KubernetesService) (*ServiceResponse, error)
	PatchKubernetesService(context.Context, *KubernetesService) (*ServiceResponse, error)
	DeleteKubernetesService(context.Context, *KubernetesService) (*ServiceResponse, error)
}

// UnimplementedKubernetesServer can be embedded to have forward compatible implementations.
type UnimplementedKubernetesServer struct {
}

func (*UnimplementedKubernetesServer) CreateKubernetesService(ctx context.Context, req *KubernetesService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateKubernetesService not implemented")
}
func (*UnimplementedKubernetesServer) GetKubernetesService(ctx context.Context, req *KubernetesService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKubernetesService not implemented")
}
func (*UnimplementedKubernetesServer) PutKubernetesService(ctx context.Context, req *KubernetesService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutKubernetesService not implemented")
}
func (*UnimplementedKubernetesServer) PatchKubernetesService(ctx context.Context, req *KubernetesService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchKubernetesService not implemented")
}
func (*UnimplementedKubernetesServer) DeleteKubernetesService(ctx context.Context, req *KubernetesService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteKubernetesService not implemented")
}

func RegisterKubernetesServer(s *grpc.Server, srv KubernetesServer) {
	s.RegisterService(&_Kubernetes_serviceDesc, srv)
}

func _Kubernetes_CreateKubernetesService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KubernetesService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubernetesServer).CreateKubernetesService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Kubernetes/CreateKubernetesService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubernetesServer).CreateKubernetesService(ctx, req.(*KubernetesService))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kubernetes_GetKubernetesService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KubernetesService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubernetesServer).GetKubernetesService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Kubernetes/GetKubernetesService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubernetesServer).GetKubernetesService(ctx, req.(*KubernetesService))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kubernetes_PutKubernetesService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KubernetesService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubernetesServer).PutKubernetesService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Kubernetes/PutKubernetesService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubernetesServer).PutKubernetesService(ctx, req.(*KubernetesService))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kubernetes_PatchKubernetesService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KubernetesService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubernetesServer).PatchKubernetesService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Kubernetes/PatchKubernetesService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubernetesServer).PatchKubernetesService(ctx, req.(*KubernetesService))
	}
	return interceptor(ctx, in, info, handler)
}

func _Kubernetes_DeleteKubernetesService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KubernetesService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KubernetesServer).DeleteKubernetesService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Kubernetes/DeleteKubernetesService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KubernetesServer).DeleteKubernetesService(ctx, req.(*KubernetesService))
	}
	return interceptor(ctx, in, info, handler)
}

var _Kubernetes_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Kubernetes",
	HandlerType: (*KubernetesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateKubernetesService",
			Handler:    _Kubernetes_CreateKubernetesService_Handler,
		},
		{
			MethodName: "GetKubernetesService",
			Handler:    _Kubernetes_GetKubernetesService_Handler,
		},
		{
			MethodName: "PutKubernetesService",
			Handler:    _Kubernetes_PutKubernetesService_Handler,
		},
		{
			MethodName: "PatchKubernetesService",
			Handler:    _Kubernetes_PatchKubernetesService_Handler,
		},
		{
			MethodName: "DeleteKubernetesService",
			Handler:    _Kubernetes_DeleteKubernetesService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kubernetesservice.proto",
}
