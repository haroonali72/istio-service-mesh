// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pod.proto

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

type RestartPolicy int32

const (
	RestartPolicy_Always    RestartPolicy = 0
	RestartPolicy_OnFailure RestartPolicy = 1
	RestartPolicy_Never     RestartPolicy = 2
)

var RestartPolicy_name = map[int32]string{
	0: "Always",
	1: "OnFailure",
	2: "Never",
}

var RestartPolicy_value = map[string]int32{
	"Always":    0,
	"OnFailure": 1,
	"Never":     2,
}

func (x RestartPolicy) String() string {
	return proto.EnumName(RestartPolicy_name, int32(x))
}

func (RestartPolicy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_106fb77aeb685f33, []int{0}
}

type PodService struct {
	ApplicationId        string                       `protobuf:"bytes,1,opt,name=application_id,json=applicationId,proto3" json:"application_id,omitempty"`
	ServiceId            string                       `protobuf:"bytes,2,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	Name                 string                       `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Version              string                       `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	ServiceType          string                       `protobuf:"bytes,5,opt,name=service_type,json=serviceType,proto3" json:"service_type,omitempty"`
	ServiceSubType       string                       `protobuf:"bytes,6,opt,name=service_sub_type,json=serviceSubType,proto3" json:"service_sub_type,omitempty"`
	Namespace            string                       `protobuf:"bytes,7,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Token                string                       `protobuf:"bytes,8,opt,name=token,proto3" json:"token,omitempty"`
	CompanyId            string                       `protobuf:"bytes,9,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	InfraId              string                       `protobuf:"bytes,10,opt,name=infra_id,json=infraId,proto3" json:"infra_id,omitempty"`
	ServiceAttributes    *PodServiceServiceAttributes `protobuf:"bytes,11,opt,name=service_attributes,json=serviceAttributes,proto3" json:"service_attributes,omitempty"`
	HookConfiguration    *HookConfiguration           `protobuf:"bytes,12,opt,name=hook_configuration,json=hookConfiguration,proto3" json:"hook_configuration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *PodService) Reset()         { *m = PodService{} }
func (m *PodService) String() string { return proto.CompactTextString(m) }
func (*PodService) ProtoMessage()    {}
func (*PodService) Descriptor() ([]byte, []int) {
	return fileDescriptor_106fb77aeb685f33, []int{0}
}

func (m *PodService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PodService.Unmarshal(m, b)
}
func (m *PodService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PodService.Marshal(b, m, deterministic)
}
func (m *PodService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PodService.Merge(m, src)
}
func (m *PodService) XXX_Size() int {
	return xxx_messageInfo_PodService.Size(m)
}
func (m *PodService) XXX_DiscardUnknown() {
	xxx_messageInfo_PodService.DiscardUnknown(m)
}

var xxx_messageInfo_PodService proto.InternalMessageInfo

func (m *PodService) GetApplicationId() string {
	if m != nil {
		return m.ApplicationId
	}
	return ""
}

func (m *PodService) GetServiceId() string {
	if m != nil {
		return m.ServiceId
	}
	return ""
}

func (m *PodService) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PodService) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *PodService) GetServiceType() string {
	if m != nil {
		return m.ServiceType
	}
	return ""
}

func (m *PodService) GetServiceSubType() string {
	if m != nil {
		return m.ServiceSubType
	}
	return ""
}

func (m *PodService) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *PodService) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *PodService) GetCompanyId() string {
	if m != nil {
		return m.CompanyId
	}
	return ""
}

func (m *PodService) GetInfraId() string {
	if m != nil {
		return m.InfraId
	}
	return ""
}

func (m *PodService) GetServiceAttributes() *PodServiceServiceAttributes {
	if m != nil {
		return m.ServiceAttributes
	}
	return nil
}

func (m *PodService) GetHookConfiguration() *HookConfiguration {
	if m != nil {
		return m.HookConfiguration
	}
	return nil
}

type PodServiceServiceAttributes struct {
	Containers                    []*ContainerAttributes         `protobuf:"bytes,1,rep,name=containers,proto3" json:"containers,omitempty"`
	IstioConfig                   *IstioConfig                   `protobuf:"bytes,2,opt,name=istio_config,json=istioConfig,proto3" json:"istio_config,omitempty"`
	NodeSelector                  map[string]string              `protobuf:"bytes,4,rep,name=node_selector,json=nodeSelector,proto3" json:"node_selector,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Labels                        map[string]string              `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Annotations                   map[string]string              `protobuf:"bytes,6,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Roles                         []*K8SRbacAttribute            `protobuf:"bytes,8,rep,name=roles,proto3" json:"roles,omitempty"`
	IstioRoles                    []*IstioRbacAttribute          `protobuf:"bytes,9,rep,name=istio_roles,json=istioRoles,proto3" json:"istio_roles,omitempty"`
	Restart_Policy                RestartPolicy                  `protobuf:"varint,10,opt,name=restart_Policy,json=restartPolicy,proto3,enum=proto.RestartPolicy" json:"restart_Policy,omitempty"`
	Volumes                       []*Volume                      `protobuf:"bytes,11,rep,name=volumes,proto3" json:"volumes,omitempty"`
	Affinity                      *Affinity                      `protobuf:"bytes,12,opt,name=affinity,proto3" json:"affinity,omitempty"`
	InitContainers                []*ContainerAttributes         `protobuf:"bytes,13,rep,name=init_containers,json=initContainers,proto3" json:"init_containers,omitempty"`
	TerminationGracePeriodSeconds *TerminationGracePeriodSeconds `protobuf:"bytes,15,opt,name=Termination_grace_period_seconds,json=TerminationGracePeriodSeconds,proto3" json:"Termination_grace_period_seconds,omitempty"`
	ImagePullSecrets              []*LocalObjectReference        `protobuf:"bytes,16,rep,name=image_pull_secrets,json=imagePullSecrets,proto3" json:"image_pull_secrets,omitempty"`
	ServiceAccountName            string                         `protobuf:"bytes,17,opt,name=serviceAccountName,proto3" json:"serviceAccountName,omitempty"`
	AutomountServiceAccountToken  *AutomountServiceAccountToken  `protobuf:"bytes,18,opt,name=automount_service_account_token,json=automountServiceAccountToken,proto3" json:"automount_service_account_token,omitempty"`
	IsRbacEnabled                 bool                           `protobuf:"varint,19,opt,name=is_rbac_enabled,json=isRbacEnabled,proto3" json:"is_rbac_enabled,omitempty"`
	XXX_NoUnkeyedLiteral          struct{}                       `json:"-"`
	XXX_unrecognized              []byte                         `json:"-"`
	XXX_sizecache                 int32                          `json:"-"`
}

func (m *PodServiceServiceAttributes) Reset()         { *m = PodServiceServiceAttributes{} }
func (m *PodServiceServiceAttributes) String() string { return proto.CompactTextString(m) }
func (*PodServiceServiceAttributes) ProtoMessage()    {}
func (*PodServiceServiceAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_106fb77aeb685f33, []int{1}
}

func (m *PodServiceServiceAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PodServiceServiceAttributes.Unmarshal(m, b)
}
func (m *PodServiceServiceAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PodServiceServiceAttributes.Marshal(b, m, deterministic)
}
func (m *PodServiceServiceAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PodServiceServiceAttributes.Merge(m, src)
}
func (m *PodServiceServiceAttributes) XXX_Size() int {
	return xxx_messageInfo_PodServiceServiceAttributes.Size(m)
}
func (m *PodServiceServiceAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_PodServiceServiceAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_PodServiceServiceAttributes proto.InternalMessageInfo

func (m *PodServiceServiceAttributes) GetContainers() []*ContainerAttributes {
	if m != nil {
		return m.Containers
	}
	return nil
}

func (m *PodServiceServiceAttributes) GetIstioConfig() *IstioConfig {
	if m != nil {
		return m.IstioConfig
	}
	return nil
}

func (m *PodServiceServiceAttributes) GetNodeSelector() map[string]string {
	if m != nil {
		return m.NodeSelector
	}
	return nil
}

func (m *PodServiceServiceAttributes) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *PodServiceServiceAttributes) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func (m *PodServiceServiceAttributes) GetRoles() []*K8SRbacAttribute {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *PodServiceServiceAttributes) GetIstioRoles() []*IstioRbacAttribute {
	if m != nil {
		return m.IstioRoles
	}
	return nil
}

func (m *PodServiceServiceAttributes) GetRestart_Policy() RestartPolicy {
	if m != nil {
		return m.Restart_Policy
	}
	return RestartPolicy_Always
}

func (m *PodServiceServiceAttributes) GetVolumes() []*Volume {
	if m != nil {
		return m.Volumes
	}
	return nil
}

func (m *PodServiceServiceAttributes) GetAffinity() *Affinity {
	if m != nil {
		return m.Affinity
	}
	return nil
}

func (m *PodServiceServiceAttributes) GetInitContainers() []*ContainerAttributes {
	if m != nil {
		return m.InitContainers
	}
	return nil
}

func (m *PodServiceServiceAttributes) GetTerminationGracePeriodSeconds() *TerminationGracePeriodSeconds {
	if m != nil {
		return m.TerminationGracePeriodSeconds
	}
	return nil
}

func (m *PodServiceServiceAttributes) GetImagePullSecrets() []*LocalObjectReference {
	if m != nil {
		return m.ImagePullSecrets
	}
	return nil
}

func (m *PodServiceServiceAttributes) GetServiceAccountName() string {
	if m != nil {
		return m.ServiceAccountName
	}
	return ""
}

func (m *PodServiceServiceAttributes) GetAutomountServiceAccountToken() *AutomountServiceAccountToken {
	if m != nil {
		return m.AutomountServiceAccountToken
	}
	return nil
}

func (m *PodServiceServiceAttributes) GetIsRbacEnabled() bool {
	if m != nil {
		return m.IsRbacEnabled
	}
	return false
}

func init() {
	proto.RegisterEnum("proto.RestartPolicy", RestartPolicy_name, RestartPolicy_value)
	proto.RegisterType((*PodService)(nil), "proto.PodService")
	proto.RegisterType((*PodServiceServiceAttributes)(nil), "proto.PodServiceServiceAttributes")
	proto.RegisterMapType((map[string]string)(nil), "proto.PodServiceServiceAttributes.AnnotationsEntry")
	proto.RegisterMapType((map[string]string)(nil), "proto.PodServiceServiceAttributes.LabelsEntry")
	proto.RegisterMapType((map[string]string)(nil), "proto.PodServiceServiceAttributes.NodeSelectorEntry")
}

func init() {
	proto.RegisterFile("pod.proto", fileDescriptor_106fb77aeb685f33)
}

var fileDescriptor_106fb77aeb685f33 = []byte{
	// 910 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x55, 0xdb, 0x6e, 0xdb, 0x46,
	0x10, 0x8d, 0xad, 0x48, 0x11, 0x47, 0xa2, 0x42, 0x6d, 0x83, 0x96, 0x51, 0x12, 0x34, 0x75, 0x6f,
	0x41, 0x8b, 0xea, 0xc1, 0xa9, 0x51, 0x37, 0x05, 0x5a, 0x08, 0x6a, 0xe2, 0x1a, 0x0d, 0x1c, 0x95,
	0x72, 0x0b, 0xf4, 0x89, 0x58, 0x91, 0x2b, 0x9b, 0x31, 0xb5, 0x4b, 0x2c, 0x97, 0x2e, 0xf4, 0x05,
	0xfd, 0x8d, 0xfe, 0x55, 0x7f, 0xa7, 0xb3, 0x17, 0xca, 0xb4, 0x53, 0x38, 0x71, 0x9e, 0xc8, 0x99,
	0x39, 0x67, 0x66, 0x76, 0x78, 0x76, 0x08, 0x5e, 0x21, 0xd2, 0x71, 0x21, 0x85, 0x12, 0xa4, 0x6d,
	0x1e, 0x23, 0xff, 0x84, 0x71, 0x26, 0x69, 0x3e, 0x76, 0x66, 0xc9, 0xe4, 0x79, 0x96, 0x30, 0x67,
	0xf6, 0xcf, 0x45, 0x5e, 0xad, 0x6a, 0x2b, 0x48, 0x59, 0x91, 0x8b, 0xf5, 0x8a, 0x71, 0x65, 0x3d,
	0x3b, 0xff, 0xb6, 0x00, 0x66, 0x22, 0x9d, 0x5b, 0x12, 0xf9, 0x1c, 0x06, 0xb4, 0x28, 0xf2, 0x2c,
	0xa1, 0x2a, 0x13, 0x3c, 0xce, 0xd2, 0x70, 0xeb, 0xf1, 0xd6, 0x13, 0x2f, 0xf2, 0x1b, 0xde, 0xc3,
	0x94, 0x3c, 0x02, 0x70, 0x65, 0x34, 0x64, 0xdb, 0x40, 0x3c, 0xe7, 0xc1, 0x30, 0x81, 0xdb, 0x9c,
	0xae, 0x58, 0xd8, 0x32, 0x01, 0xf3, 0x4e, 0x42, 0xb8, 0x73, 0xce, 0x64, 0x89, 0xfc, 0xf0, 0xb6,
	0x71, 0xd7, 0x26, 0xf9, 0x04, 0xfa, 0x75, 0x32, 0xb5, 0x2e, 0x58, 0xd8, 0x36, 0xe1, 0x9e, 0xf3,
	0x1d, 0xa3, 0x8b, 0x3c, 0x81, 0xa0, 0x86, 0x94, 0xd5, 0xc2, 0xc2, 0x3a, 0x06, 0x36, 0x70, 0xfe,
	0x79, 0xb5, 0x30, 0xc8, 0x87, 0xe0, 0xe9, 0x72, 0x65, 0x41, 0x13, 0x16, 0xde, 0xb1, 0x8d, 0x6d,
	0x1c, 0xe4, 0x1e, 0xb4, 0x95, 0x38, 0x63, 0x3c, 0xec, 0x9a, 0x88, 0x35, 0xf4, 0x69, 0x12, 0xb1,
	0x2a, 0x28, 0x5f, 0xeb, 0xd3, 0x78, 0x96, 0xe4, 0x3c, 0x78, 0x9a, 0xfb, 0xd0, 0xcd, 0xf8, 0x52,
	0x52, 0x1d, 0x04, 0xdb, 0xba, 0xb1, 0x31, 0xf4, 0x1b, 0x90, 0xba, 0x2f, 0xaa, 0x94, 0xcc, 0x16,
	0x95, 0x62, 0x65, 0xd8, 0x43, 0x50, 0x6f, 0x77, 0xc7, 0x4e, 0x78, 0x7c, 0x31, 0x5d, 0xf7, 0x98,
	0x6c, 0x90, 0xd1, 0xb0, 0xbc, 0xea, 0x22, 0x07, 0x40, 0x4e, 0x85, 0x38, 0x8b, 0x13, 0xc1, 0x97,
	0xd9, 0x49, 0x25, 0xcd, 0xc8, 0xc3, 0xbe, 0x49, 0x19, 0xba, 0x94, 0xbf, 0x20, 0x60, 0xda, 0x8c,
	0x47, 0xc3, 0xd3, 0xab, 0xae, 0x9d, 0xbf, 0x01, 0x1e, 0x5c, 0x53, 0x9b, 0x3c, 0xd3, 0xa7, 0xe6,
	0x8a, 0x66, 0xa8, 0x9e, 0x12, 0x3f, 0x73, 0x0b, 0x0b, 0x8c, 0x5c, 0x81, 0x69, 0x1d, 0x68, 0xf4,
	0xda, 0x40, 0x93, 0x3d, 0xe8, 0x67, 0x25, 0x56, 0x71, 0x5d, 0x1a, 0x05, 0xf4, 0x76, 0x89, 0x63,
	0x1f, 0xea, 0x90, 0x6d, 0x26, 0xea, 0x65, 0x17, 0x06, 0xf9, 0x13, 0x7c, 0x2e, 0x52, 0xfc, 0x86,
	0x2c, 0x67, 0x89, 0x12, 0x12, 0x95, 0xa0, 0xab, 0x7e, 0xfb, 0xf6, 0x49, 0x8d, 0x8f, 0x90, 0x37,
	0x77, 0xb4, 0xe7, 0x5c, 0xc9, 0x75, 0xd4, 0xe7, 0x0d, 0x17, 0x79, 0x01, 0x9d, 0x9c, 0x2e, 0x58,
	0x5e, 0xa2, 0x7c, 0x74, 0xce, 0xf1, 0x3b, 0xe4, 0x7c, 0x69, 0x08, 0x36, 0x9b, 0x63, 0x93, 0xdf,
	0xa1, 0x47, 0x39, 0x17, 0xca, 0xcc, 0xb0, 0x44, 0x91, 0xe9, 0x64, 0x4f, 0xdf, 0x21, 0xd9, 0xe4,
	0x82, 0x65, 0x33, 0x36, 0xf3, 0x90, 0x6f, 0xa0, 0x2d, 0x45, 0x8e, 0xda, 0xe8, 0x9a, 0x84, 0x1f,
	0xb9, 0x84, 0xbf, 0xee, 0x97, 0xd1, 0x82, 0x26, 0x9b, 0x34, 0x91, 0x45, 0xe1, 0xb7, 0xb1, 0x73,
	0x8b, 0x2d, 0xc9, 0x33, 0xa4, 0xfb, 0xcd, 0xf1, 0x5e, 0xa6, 0x81, 0x41, 0x47, 0x86, 0xfb, 0x03,
	0x0c, 0x24, 0x2b, 0x15, 0x95, 0x2a, 0x9e, 0x09, 0xbc, 0xb3, 0x6b, 0x23, 0xda, 0xc1, 0xee, 0x3d,
	0x47, 0x8f, 0x6c, 0xd0, 0xc6, 0x22, 0x5f, 0x36, 0x4d, 0xf2, 0x25, 0xde, 0x52, 0xb3, 0x30, 0xb4,
	0x8a, 0x75, 0x51, 0xdf, 0xb1, 0xfe, 0x30, 0xde, 0xa8, 0x8e, 0x92, 0xaf, 0xa1, 0x4b, 0x97, 0xcb,
	0x8c, 0x67, 0x6a, 0xed, 0xc4, 0x79, 0xd7, 0x21, 0x27, 0xce, 0x1d, 0x6d, 0x00, 0x64, 0x0a, 0x77,
	0xf5, 0x4b, 0xdc, 0xd0, 0x9b, 0xff, 0x56, 0xbd, 0x0d, 0x34, 0x65, 0x7a, 0xa1, 0xb9, 0x15, 0x3c,
	0x3e, 0x66, 0x72, 0x95, 0x71, 0xbb, 0x9a, 0x4e, 0x24, 0x5e, 0xe8, 0xb8, 0x60, 0x32, 0x13, 0x29,
	0x0a, 0x0a, 0x53, 0xa7, 0x65, 0x78, 0xd7, 0x74, 0xf2, 0x99, 0xcb, 0xda, 0x80, 0x1f, 0x68, 0xf4,
	0xcc, 0x80, 0xe7, 0x16, 0x1b, 0x3d, 0xba, 0x36, 0x4c, 0x0e, 0x81, 0x64, 0x2b, 0x7a, 0x82, 0x25,
	0xaa, 0x3c, 0xd7, 0x05, 0x24, 0x53, 0x65, 0x18, 0x98, 0xb6, 0x1f, 0xb8, 0x02, 0x2f, 0x45, 0x42,
	0xf3, 0x57, 0x8b, 0xd7, 0x28, 0xc1, 0x88, 0x2d, 0x99, 0x64, 0x3c, 0x61, 0x51, 0x60, 0x68, 0x33,
	0x64, 0xcd, 0x2d, 0x89, 0x8c, 0x37, 0x5b, 0x62, 0x92, 0x24, 0xa2, 0xe2, 0xea, 0x48, 0x2f, 0xc7,
	0xa1, 0x59, 0x25, 0xff, 0x13, 0x21, 0xaf, 0xe1, 0x63, 0x5a, 0x29, 0xb1, 0xd2, 0x8e, 0x78, 0xb3,
	0x5f, 0x2c, 0x20, 0xb6, 0xfb, 0x8b, 0x98, 0x83, 0x7e, 0x5a, 0x8f, 0xbc, 0x46, 0xcf, 0x2f, 0x25,
	0x3b, 0xd6, 0xd0, 0xe8, 0x21, 0xbd, 0x26, 0x4a, 0xbe, 0xc0, 0x4f, 0x53, 0xc6, 0x12, 0xd5, 0x14,
	0x33, 0x4e, 0x17, 0x39, 0x4b, 0xc3, 0x0f, 0x30, 0x77, 0x37, 0xf2, 0x33, 0x23, 0xcd, 0xe7, 0xd6,
	0x39, 0xfa, 0x09, 0x86, 0x6f, 0x5c, 0x41, 0x12, 0x40, 0xeb, 0x8c, 0xad, 0xdd, 0x2f, 0x42, 0xbf,
	0xea, 0x05, 0x7b, 0x4e, 0xf3, 0x8a, 0xb9, 0x7f, 0x82, 0x35, 0x9e, 0x6d, 0xef, 0x6f, 0x8d, 0xbe,
	0x87, 0x5e, 0xe3, 0xbe, 0xdd, 0x88, 0xfa, 0x23, 0x04, 0x57, 0x6f, 0xd7, 0x4d, 0xf8, 0x5f, 0xed,
	0x81, 0x7f, 0x49, 0xf4, 0x04, 0xa0, 0x33, 0xc9, 0xff, 0xa2, 0xeb, 0x32, 0xb8, 0x45, 0x7c, 0xf0,
	0x5e, 0xf1, 0x17, 0x34, 0xcb, 0x2b, 0xc9, 0x82, 0x2d, 0xe2, 0x41, 0xfb, 0x88, 0xe1, 0x9f, 0x29,
	0xd8, 0xde, 0xfd, 0x67, 0x1b, 0x5a, 0x78, 0xe3, 0xc9, 0x3e, 0x78, 0x53, 0xc9, 0xa8, 0x62, 0xda,
	0x18, 0xbe, 0xb1, 0x0a, 0x46, 0x1f, 0x3a, 0x97, 0xb3, 0xb1, 0x54, 0x81, 0x6d, 0xb2, 0x9d, 0x5b,
	0x9a, 0xf9, 0x33, 0x0e, 0xec, 0x3d, 0x98, 0x7b, 0xd0, 0x39, 0x60, 0xea, 0xc6, 0xb4, 0xef, 0xa0,
	0x3b, 0xa3, 0x2a, 0x39, 0x7d, 0x9f, 0x7a, 0xb3, 0xea, 0xc6, 0xf5, 0x16, 0x1d, 0x13, 0x78, 0xfa,
	0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xeb, 0x35, 0x25, 0xec, 0x96, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PodClient is the client API for Pod service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PodClient interface {
	CreatePod(ctx context.Context, in *PodService, opts ...grpc.CallOption) (*ServiceResponse, error)
	DeletePod(ctx context.Context, in *PodService, opts ...grpc.CallOption) (*ServiceResponse, error)
	GetPod(ctx context.Context, in *PodService, opts ...grpc.CallOption) (*ServiceResponse, error)
	PatchPod(ctx context.Context, in *PodService, opts ...grpc.CallOption) (*ServiceResponse, error)
	PutPod(ctx context.Context, in *PodService, opts ...grpc.CallOption) (*ServiceResponse, error)
}

type podClient struct {
	cc grpc.ClientConnInterface
}

func NewPodClient(cc grpc.ClientConnInterface) PodClient {
	return &podClient{cc}
}

func (c *podClient) CreatePod(ctx context.Context, in *PodService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.Pod/CreatePod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podClient) DeletePod(ctx context.Context, in *PodService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.Pod/DeletePod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podClient) GetPod(ctx context.Context, in *PodService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.Pod/GetPod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podClient) PatchPod(ctx context.Context, in *PodService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.Pod/PatchPod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podClient) PutPod(ctx context.Context, in *PodService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.Pod/PutPod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PodServer is the server API for Pod service.
type PodServer interface {
	CreatePod(context.Context, *PodService) (*ServiceResponse, error)
	DeletePod(context.Context, *PodService) (*ServiceResponse, error)
	GetPod(context.Context, *PodService) (*ServiceResponse, error)
	PatchPod(context.Context, *PodService) (*ServiceResponse, error)
	PutPod(context.Context, *PodService) (*ServiceResponse, error)
}

// UnimplementedPodServer can be embedded to have forward compatible implementations.
type UnimplementedPodServer struct {
}

func (*UnimplementedPodServer) CreatePod(ctx context.Context, req *PodService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePod not implemented")
}
func (*UnimplementedPodServer) DeletePod(ctx context.Context, req *PodService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePod not implemented")
}
func (*UnimplementedPodServer) GetPod(ctx context.Context, req *PodService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPod not implemented")
}
func (*UnimplementedPodServer) PatchPod(ctx context.Context, req *PodService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchPod not implemented")
}
func (*UnimplementedPodServer) PutPod(ctx context.Context, req *PodService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutPod not implemented")
}

func RegisterPodServer(s *grpc.Server, srv PodServer) {
	s.RegisterService(&_Pod_serviceDesc, srv)
}

func _Pod_CreatePod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PodService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PodServer).CreatePod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Pod/CreatePod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PodServer).CreatePod(ctx, req.(*PodService))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pod_DeletePod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PodService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PodServer).DeletePod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Pod/DeletePod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PodServer).DeletePod(ctx, req.(*PodService))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pod_GetPod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PodService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PodServer).GetPod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Pod/GetPod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PodServer).GetPod(ctx, req.(*PodService))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pod_PatchPod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PodService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PodServer).PatchPod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Pod/PatchPod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PodServer).PatchPod(ctx, req.(*PodService))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pod_PutPod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PodService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PodServer).PutPod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Pod/PutPod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PodServer).PutPod(ctx, req.(*PodService))
	}
	return interceptor(ctx, in, info, handler)
}

var _Pod_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Pod",
	HandlerType: (*PodServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePod",
			Handler:    _Pod_CreatePod_Handler,
		},
		{
			MethodName: "DeletePod",
			Handler:    _Pod_DeletePod_Handler,
		},
		{
			MethodName: "GetPod",
			Handler:    _Pod_GetPod_Handler,
		},
		{
			MethodName: "PatchPod",
			Handler:    _Pod_PatchPod_Handler,
		},
		{
			MethodName: "PutPod",
			Handler:    _Pod_PutPod_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pod.proto",
}
