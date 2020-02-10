// Code generated by protoc-gen-go. DO NOT EDIT.
// source: deployment.proto

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

type DeploymentService struct {
	ProjectId            string                       `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	ServiceId            string                       `protobuf:"bytes,2,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	Name                 string                       `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Version              string                       `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	ServiceType          string                       `protobuf:"bytes,5,opt,name=service_type,json=serviceType,proto3" json:"service_type,omitempty"`
	ServiceSubType       string                       `protobuf:"bytes,6,opt,name=service_sub_type,json=serviceSubType,proto3" json:"service_sub_type,omitempty"`
	Namespace            string                       `protobuf:"bytes,7,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Token                string                       `protobuf:"bytes,8,opt,name=token,proto3" json:"token,omitempty"`
	CompanyId            string                       `protobuf:"bytes,9,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	ServiceAttributes    *DeploymentServiceAttributes `protobuf:"bytes,10,opt,name=service_attributes,json=serviceAttributes,proto3" json:"service_attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *DeploymentService) Reset()         { *m = DeploymentService{} }
func (m *DeploymentService) String() string { return proto.CompactTextString(m) }
func (*DeploymentService) ProtoMessage()    {}
func (*DeploymentService) Descriptor() ([]byte, []int) {
	return fileDescriptor_fac0ec10f8e4d7ff, []int{0}
}

func (m *DeploymentService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeploymentService.Unmarshal(m, b)
}
func (m *DeploymentService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeploymentService.Marshal(b, m, deterministic)
}
func (m *DeploymentService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeploymentService.Merge(m, src)
}
func (m *DeploymentService) XXX_Size() int {
	return xxx_messageInfo_DeploymentService.Size(m)
}
func (m *DeploymentService) XXX_DiscardUnknown() {
	xxx_messageInfo_DeploymentService.DiscardUnknown(m)
}

var xxx_messageInfo_DeploymentService proto.InternalMessageInfo

func (m *DeploymentService) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *DeploymentService) GetServiceId() string {
	if m != nil {
		return m.ServiceId
	}
	return ""
}

func (m *DeploymentService) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DeploymentService) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *DeploymentService) GetServiceType() string {
	if m != nil {
		return m.ServiceType
	}
	return ""
}

func (m *DeploymentService) GetServiceSubType() string {
	if m != nil {
		return m.ServiceSubType
	}
	return ""
}

func (m *DeploymentService) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *DeploymentService) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *DeploymentService) GetCompanyId() string {
	if m != nil {
		return m.CompanyId
	}
	return ""
}

func (m *DeploymentService) GetServiceAttributes() *DeploymentServiceAttributes {
	if m != nil {
		return m.ServiceAttributes
	}
	return nil
}

type DeploymentServiceAttributes struct {
	Containers                    map[string]*ContainerAttributes `protobuf:"bytes,1,rep,name=containers,proto3" json:"containers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	IstioConfig                   *IstioConfig                    `protobuf:"bytes,2,opt,name=istio_config,json=istioConfig,proto3" json:"istio_config,omitempty"`
	LabelSelector                 *LabelSelectorObj               `protobuf:"bytes,3,opt,name=label_selector,json=labelSelector,proto3" json:"label_selector,omitempty"`
	NodeSelector                  map[string]string               `protobuf:"bytes,4,rep,name=node_selector,json=nodeSelector,proto3" json:"node_selector,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Labels                        map[string]string               `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Annotations                   map[string]string               `protobuf:"bytes,6,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Roles                         []*K8SRbacAttribute             `protobuf:"bytes,8,rep,name=roles,proto3" json:"roles,omitempty"`
	IstioRoles                    []*IstioRbacAttribute           `protobuf:"bytes,9,rep,name=istio_roles,json=istioRoles,proto3" json:"istio_roles,omitempty"`
	Strategy                      *DeploymentStrategy             `protobuf:"bytes,10,opt,name=strategy,proto3" json:"strategy,omitempty"`
	Volumes                       []*Volume                       `protobuf:"bytes,11,rep,name=volumes,proto3" json:"volumes,omitempty"`
	Affinity                      *Affinity                       `protobuf:"bytes,12,opt,name=affinity,proto3" json:"affinity,omitempty"`
	InitContainers                map[string]*ContainerAttributes `protobuf:"bytes,13,rep,name=initContainers,proto3" json:"initContainers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Replicas                      *Replicas                       `protobuf:"bytes,14,opt,name=replicas,proto3" json:"replicas,omitempty"`
	TerminationGracePeriodSeconds *TerminationGracePeriodSeconds  `protobuf:"bytes,15,opt,name=TerminationGracePeriodSeconds,proto3" json:"TerminationGracePeriodSeconds,omitempty"`
	XXX_NoUnkeyedLiteral          struct{}                        `json:"-"`
	XXX_unrecognized              []byte                          `json:"-"`
	XXX_sizecache                 int32                           `json:"-"`
}

func (m *DeploymentServiceAttributes) Reset()         { *m = DeploymentServiceAttributes{} }
func (m *DeploymentServiceAttributes) String() string { return proto.CompactTextString(m) }
func (*DeploymentServiceAttributes) ProtoMessage()    {}
func (*DeploymentServiceAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_fac0ec10f8e4d7ff, []int{1}
}

func (m *DeploymentServiceAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeploymentServiceAttributes.Unmarshal(m, b)
}
func (m *DeploymentServiceAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeploymentServiceAttributes.Marshal(b, m, deterministic)
}
func (m *DeploymentServiceAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeploymentServiceAttributes.Merge(m, src)
}
func (m *DeploymentServiceAttributes) XXX_Size() int {
	return xxx_messageInfo_DeploymentServiceAttributes.Size(m)
}
func (m *DeploymentServiceAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_DeploymentServiceAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_DeploymentServiceAttributes proto.InternalMessageInfo

func (m *DeploymentServiceAttributes) GetContainers() map[string]*ContainerAttributes {
	if m != nil {
		return m.Containers
	}
	return nil
}

func (m *DeploymentServiceAttributes) GetIstioConfig() *IstioConfig {
	if m != nil {
		return m.IstioConfig
	}
	return nil
}

func (m *DeploymentServiceAttributes) GetLabelSelector() *LabelSelectorObj {
	if m != nil {
		return m.LabelSelector
	}
	return nil
}

func (m *DeploymentServiceAttributes) GetNodeSelector() map[string]string {
	if m != nil {
		return m.NodeSelector
	}
	return nil
}

func (m *DeploymentServiceAttributes) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *DeploymentServiceAttributes) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func (m *DeploymentServiceAttributes) GetRoles() []*K8SRbacAttribute {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *DeploymentServiceAttributes) GetIstioRoles() []*IstioRbacAttribute {
	if m != nil {
		return m.IstioRoles
	}
	return nil
}

func (m *DeploymentServiceAttributes) GetStrategy() *DeploymentStrategy {
	if m != nil {
		return m.Strategy
	}
	return nil
}

func (m *DeploymentServiceAttributes) GetVolumes() []*Volume {
	if m != nil {
		return m.Volumes
	}
	return nil
}

func (m *DeploymentServiceAttributes) GetAffinity() *Affinity {
	if m != nil {
		return m.Affinity
	}
	return nil
}

func (m *DeploymentServiceAttributes) GetInitContainers() map[string]*ContainerAttributes {
	if m != nil {
		return m.InitContainers
	}
	return nil
}

func (m *DeploymentServiceAttributes) GetReplicas() *Replicas {
	if m != nil {
		return m.Replicas
	}
	return nil
}

func (m *DeploymentServiceAttributes) GetTerminationGracePeriodSeconds() *TerminationGracePeriodSeconds {
	if m != nil {
		return m.TerminationGracePeriodSeconds
	}
	return nil
}

func init() {
	proto.RegisterType((*DeploymentService)(nil), "proto.DeploymentService")
	proto.RegisterType((*DeploymentServiceAttributes)(nil), "proto.DeploymentServiceAttributes")
	proto.RegisterMapType((map[string]string)(nil), "proto.DeploymentServiceAttributes.AnnotationsEntry")
	proto.RegisterMapType((map[string]*ContainerAttributes)(nil), "proto.DeploymentServiceAttributes.ContainersEntry")
	proto.RegisterMapType((map[string]*ContainerAttributes)(nil), "proto.DeploymentServiceAttributes.InitContainersEntry")
	proto.RegisterMapType((map[string]string)(nil), "proto.DeploymentServiceAttributes.LabelsEntry")
	proto.RegisterMapType((map[string]string)(nil), "proto.DeploymentServiceAttributes.NodeSelectorEntry")
}

func init() { proto.RegisterFile("deployment.proto", fileDescriptor_fac0ec10f8e4d7ff) }

var fileDescriptor_fac0ec10f8e4d7ff = []byte{
	// 755 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xdd, 0x6e, 0x23, 0x35,
	0x14, 0x26, 0x69, 0x93, 0x26, 0x67, 0xf2, 0x57, 0xb3, 0x02, 0x6f, 0x60, 0xa5, 0x12, 0x21, 0x11,
	0x09, 0x11, 0xa1, 0x2c, 0x8b, 0x96, 0xbd, 0x58, 0x54, 0x65, 0xdb, 0x2a, 0x02, 0x41, 0x71, 0x0a,
	0x52, 0x2f, 0x20, 0x72, 0x66, 0x4e, 0xcb, 0xb4, 0x13, 0x7b, 0x64, 0x3b, 0x91, 0xe6, 0x11, 0x79,
	0x18, 0x6e, 0x78, 0x02, 0x34, 0x1e, 0x67, 0x32, 0x4d, 0x4b, 0x49, 0x95, 0xbd, 0xca, 0x9c, 0xef,
	0xcf, 0xc7, 0xc7, 0x8e, 0xa1, 0x13, 0x60, 0x1c, 0xc9, 0x64, 0x8e, 0xc2, 0x0c, 0x62, 0x25, 0x8d,
	0x24, 0x15, 0xfb, 0xd3, 0x6d, 0x5e, 0xa3, 0x40, 0xc5, 0xa3, 0x81, 0x2b, 0x35, 0xaa, 0x65, 0xe8,
	0xa3, 0x2b, 0x1b, 0x4b, 0x19, 0x2d, 0xe6, 0xae, 0xea, 0xfd, 0x53, 0x86, 0xc3, 0x77, 0x79, 0xce,
	0x24, 0x53, 0x92, 0x17, 0x00, 0xb1, 0x92, 0x37, 0xe8, 0x9b, 0x69, 0x18, 0xd0, 0xd2, 0x51, 0xa9,
	0x5f, 0x67, 0x75, 0x87, 0x8c, 0x83, 0x94, 0x76, 0x99, 0x29, 0x5d, 0xce, 0x68, 0x87, 0x8c, 0x03,
	0x42, 0x60, 0x5f, 0xf0, 0x39, 0xd2, 0x3d, 0x4b, 0xd8, 0x6f, 0x42, 0xe1, 0x60, 0x89, 0x4a, 0x87,
	0x52, 0xd0, 0x7d, 0x0b, 0xaf, 0x4a, 0xf2, 0x19, 0x34, 0x56, 0x61, 0x26, 0x89, 0x91, 0x56, 0x2c,
	0xed, 0x39, 0xec, 0x22, 0x89, 0x91, 0xf4, 0xa1, 0xb3, 0x92, 0xe8, 0xc5, 0x2c, 0x93, 0x55, 0xad,
	0xac, 0xe5, 0xf0, 0xc9, 0x62, 0x66, 0x95, 0x9f, 0x42, 0x3d, 0x5d, 0x4e, 0xc7, 0xdc, 0x47, 0x7a,
	0x90, 0x35, 0x96, 0x03, 0xe4, 0x19, 0x54, 0x8c, 0xbc, 0x45, 0x41, 0x6b, 0x96, 0xc9, 0x8a, 0x74,
	0x37, 0xbe, 0x9c, 0xc7, 0x5c, 0x24, 0xe9, 0x6e, 0xea, 0x99, 0xc9, 0x21, 0xe3, 0x80, 0xfc, 0x02,
	0x64, 0xb5, 0x38, 0x37, 0x46, 0x85, 0xb3, 0x85, 0x41, 0x4d, 0xe1, 0xa8, 0xd4, 0xf7, 0x86, 0xbd,
	0x6c, 0x8a, 0x83, 0x7b, 0x13, 0x3c, 0xce, 0x95, 0xec, 0x50, 0x6f, 0x42, 0xbd, 0xbf, 0x00, 0x3e,
	0x79, 0xc4, 0x42, 0x58, 0xda, 0x91, 0x30, 0x3c, 0x14, 0xa8, 0x34, 0x2d, 0x1d, 0xed, 0xf5, 0xbd,
	0xe1, 0xf0, 0xff, 0x97, 0x1a, 0x8c, 0x72, 0xd3, 0x89, 0x30, 0x2a, 0x61, 0x85, 0x14, 0xf2, 0x0a,
	0x1a, 0xa1, 0x36, 0xa1, 0x9c, 0xfa, 0x52, 0x5c, 0x85, 0xd7, 0xf6, 0xd4, 0xbc, 0x21, 0x71, 0xa9,
	0xe3, 0x94, 0x1a, 0x59, 0x86, 0x79, 0xe1, 0xba, 0x20, 0x6f, 0xa1, 0x15, 0xf1, 0x19, 0x46, 0x53,
	0x8d, 0x11, 0xfa, 0x46, 0x2a, 0x7b, 0xaa, 0xde, 0xf0, 0x63, 0x67, 0xfc, 0x31, 0x25, 0x27, 0x8e,
	0xfb, 0x79, 0x76, 0xc3, 0x9a, 0x51, 0x11, 0x21, 0x97, 0xd0, 0x14, 0x32, 0xc0, 0xb5, 0x7d, 0xdf,
	0xee, 0xe6, 0x9b, 0x2d, 0x76, 0xf3, 0x93, 0x0c, 0x70, 0x95, 0x93, 0xed, 0xa7, 0x21, 0x0a, 0x10,
	0x39, 0x85, 0xaa, 0x5d, 0x4b, 0xd3, 0x8a, 0xcd, 0x1c, 0x6c, 0x91, 0x69, 0xdb, 0x75, 0xd3, 0x71,
	0x6e, 0xf2, 0x2b, 0x78, 0x5c, 0x08, 0x69, 0xb8, 0x09, 0xa5, 0xd0, 0xb4, 0x6a, 0xc3, 0x5e, 0x6e,
	0x11, 0x76, 0xbc, 0x76, 0x65, 0x89, 0xc5, 0x1c, 0xf2, 0x15, 0x54, 0x94, 0x8c, 0x50, 0xd3, 0x9a,
	0x0d, 0x5c, 0x0d, 0xec, 0x87, 0xd7, 0x9a, 0xcd, 0xb8, 0x9f, 0xc7, 0xb0, 0x4c, 0x45, 0xde, 0x40,
	0x36, 0xf7, 0x69, 0x66, 0xaa, 0x5b, 0xd3, 0xf3, 0xe2, 0xf1, 0xdc, 0xb5, 0x81, 0x55, 0x33, 0xeb,
	0x7d, 0x05, 0x35, 0x6d, 0x14, 0x37, 0x78, 0x9d, 0xb8, 0x8b, 0xf9, 0xfc, 0x7e, 0xfb, 0x4e, 0xc0,
	0x72, 0x29, 0xf9, 0x02, 0x0e, 0xb2, 0xb7, 0x40, 0x53, 0xcf, 0x2e, 0xd7, 0x74, 0xae, 0xdf, 0x2c,
	0xca, 0x56, 0x2c, 0xf9, 0x12, 0x6a, 0xfc, 0xea, 0x2a, 0x14, 0xa1, 0x49, 0x68, 0xc3, 0xe6, 0xb7,
	0x9d, 0xf2, 0xd8, 0xc1, 0x2c, 0x17, 0x90, 0x3f, 0xa0, 0x95, 0x7e, 0xac, 0xef, 0x22, 0x6d, 0xda,
	0xf0, 0x6f, 0xb7, 0x98, 0xe8, 0xf8, 0x8e, 0x31, 0x1b, 0xea, 0x46, 0x5a, 0xda, 0x8c, 0xc2, 0x38,
	0x0a, 0x7d, 0xae, 0x69, 0xeb, 0x4e, 0x33, 0xcc, 0xc1, 0x2c, 0x17, 0x90, 0x1b, 0x78, 0x71, 0x81,
	0x6a, 0x1e, 0x0a, 0x7b, 0x28, 0x67, 0x8a, 0xfb, 0x78, 0x8e, 0x2a, 0x94, 0xc1, 0x04, 0x7d, 0x29,
	0x02, 0x4d, 0xdb, 0x36, 0xe1, 0x73, 0x97, 0xf0, 0xa8, 0x96, 0x3d, 0x1e, 0xd5, 0xbd, 0x84, 0xf6,
	0x46, 0xef, 0xa4, 0x03, 0x7b, 0xb7, 0x98, 0xb8, 0x07, 0x34, 0xfd, 0x24, 0x5f, 0x43, 0x65, 0xc9,
	0xa3, 0x05, 0xba, 0xff, 0x5f, 0xd7, 0x2d, 0x9c, 0x1b, 0x0b, 0x0f, 0x47, 0x26, 0x7c, 0x53, 0x7e,
	0x5d, 0xea, 0x7e, 0x0f, 0x87, 0xf7, 0xfe, 0x0d, 0x0f, 0x84, 0x3f, 0x2b, 0x86, 0xd7, 0x8b, 0x01,
	0xdf, 0x81, 0x57, 0xb8, 0xfa, 0x4f, 0xb2, 0xbe, 0x85, 0xce, 0xe6, 0x45, 0x7f, 0x92, 0xff, 0x77,
	0xf8, 0xf0, 0x81, 0x63, 0x7d, 0x5f, 0xa3, 0x19, 0xfe, 0x5d, 0x06, 0x58, 0x5f, 0x29, 0x72, 0x0a,
	0x9d, 0x91, 0x42, 0x6e, 0xb0, 0x80, 0xd1, 0xff, 0xba, 0x79, 0xdd, 0x8f, 0x1c, 0xe3, 0x6a, 0x86,
	0x3a, 0x96, 0x42, 0x63, 0xef, 0x83, 0x34, 0xe7, 0x1d, 0x46, 0xb8, 0x73, 0xce, 0x08, 0x9a, 0x67,
	0x68, 0x76, 0x0c, 0x39, 0x81, 0xf6, 0x39, 0x37, 0xfe, 0x9f, 0xbb, 0xf7, 0x72, 0xbe, 0xd8, 0xb1,
	0x97, 0x59, 0xd5, 0x12, 0x2f, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xd8, 0x4d, 0xed, 0xc6, 0x7e,
	0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DeploymentClient is the client API for Deployment service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DeploymentClient interface {
	CreateDeployment(ctx context.Context, in *DeploymentService, opts ...grpc.CallOption) (*ServiceResponse, error)
	DeleteDeployment(ctx context.Context, in *DeploymentService, opts ...grpc.CallOption) (*ServiceResponse, error)
	GetDeployment(ctx context.Context, in *DeploymentService, opts ...grpc.CallOption) (*ServiceResponse, error)
	PatchDeployment(ctx context.Context, in *DeploymentService, opts ...grpc.CallOption) (*ServiceResponse, error)
	PutDeployment(ctx context.Context, in *DeploymentService, opts ...grpc.CallOption) (*ServiceResponse, error)
}

type deploymentClient struct {
	cc grpc.ClientConnInterface
}

func NewDeploymentClient(cc grpc.ClientConnInterface) DeploymentClient {
	return &deploymentClient{cc}
}

func (c *deploymentClient) CreateDeployment(ctx context.Context, in *DeploymentService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.Deployment/CreateDeployment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deploymentClient) DeleteDeployment(ctx context.Context, in *DeploymentService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.Deployment/DeleteDeployment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deploymentClient) GetDeployment(ctx context.Context, in *DeploymentService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.Deployment/GetDeployment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deploymentClient) PatchDeployment(ctx context.Context, in *DeploymentService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.Deployment/PatchDeployment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deploymentClient) PutDeployment(ctx context.Context, in *DeploymentService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.Deployment/PutDeployment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeploymentServer is the server API for Deployment service.
type DeploymentServer interface {
	CreateDeployment(context.Context, *DeploymentService) (*ServiceResponse, error)
	DeleteDeployment(context.Context, *DeploymentService) (*ServiceResponse, error)
	GetDeployment(context.Context, *DeploymentService) (*ServiceResponse, error)
	PatchDeployment(context.Context, *DeploymentService) (*ServiceResponse, error)
	PutDeployment(context.Context, *DeploymentService) (*ServiceResponse, error)
}

// UnimplementedDeploymentServer can be embedded to have forward compatible implementations.
type UnimplementedDeploymentServer struct {
}

func (*UnimplementedDeploymentServer) CreateDeployment(ctx context.Context, req *DeploymentService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDeployment not implemented")
}
func (*UnimplementedDeploymentServer) DeleteDeployment(ctx context.Context, req *DeploymentService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDeployment not implemented")
}
func (*UnimplementedDeploymentServer) GetDeployment(ctx context.Context, req *DeploymentService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeployment not implemented")
}
func (*UnimplementedDeploymentServer) PatchDeployment(ctx context.Context, req *DeploymentService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchDeployment not implemented")
}
func (*UnimplementedDeploymentServer) PutDeployment(ctx context.Context, req *DeploymentService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutDeployment not implemented")
}

func RegisterDeploymentServer(s *grpc.Server, srv DeploymentServer) {
	s.RegisterService(&_Deployment_serviceDesc, srv)
}

func _Deployment_CreateDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeploymentService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeploymentServer).CreateDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Deployment/CreateDeployment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeploymentServer).CreateDeployment(ctx, req.(*DeploymentService))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deployment_DeleteDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeploymentService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeploymentServer).DeleteDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Deployment/DeleteDeployment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeploymentServer).DeleteDeployment(ctx, req.(*DeploymentService))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deployment_GetDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeploymentService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeploymentServer).GetDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Deployment/GetDeployment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeploymentServer).GetDeployment(ctx, req.(*DeploymentService))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deployment_PatchDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeploymentService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeploymentServer).PatchDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Deployment/PatchDeployment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeploymentServer).PatchDeployment(ctx, req.(*DeploymentService))
	}
	return interceptor(ctx, in, info, handler)
}

func _Deployment_PutDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeploymentService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeploymentServer).PutDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Deployment/PutDeployment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeploymentServer).PutDeployment(ctx, req.(*DeploymentService))
	}
	return interceptor(ctx, in, info, handler)
}

var _Deployment_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Deployment",
	HandlerType: (*DeploymentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDeployment",
			Handler:    _Deployment_CreateDeployment_Handler,
		},
		{
			MethodName: "DeleteDeployment",
			Handler:    _Deployment_DeleteDeployment_Handler,
		},
		{
			MethodName: "GetDeployment",
			Handler:    _Deployment_GetDeployment_Handler,
		},
		{
			MethodName: "PatchDeployment",
			Handler:    _Deployment_PatchDeployment_Handler,
		},
		{
			MethodName: "PutDeployment",
			Handler:    _Deployment_PutDeployment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "deployment.proto",
}
