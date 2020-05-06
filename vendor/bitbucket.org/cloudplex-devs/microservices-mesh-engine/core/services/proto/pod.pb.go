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
	ProjectId            string                       `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	ServiceId            string                       `protobuf:"bytes,2,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	Name                 string                       `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Version              string                       `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	ServiceType          string                       `protobuf:"bytes,5,opt,name=service_type,json=serviceType,proto3" json:"service_type,omitempty"`
	ServiceSubType       string                       `protobuf:"bytes,6,opt,name=service_sub_type,json=serviceSubType,proto3" json:"service_sub_type,omitempty"`
	Namespace            string                       `protobuf:"bytes,7,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Token                string                       `protobuf:"bytes,8,opt,name=token,proto3" json:"token,omitempty"`
	CompanyId            string                       `protobuf:"bytes,9,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	ServiceAttributes    *PodServiceServiceAttributes `protobuf:"bytes,10,opt,name=service_attributes,json=serviceAttributes,proto3" json:"service_attributes,omitempty"`
	HookConfiguration    *HookConfiguration           `protobuf:"bytes,11,opt,name=hook_configuration,json=hookConfiguration,proto3" json:"hook_configuration,omitempty"`
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

func (m *PodService) GetProjectId() string {
	if m != nil {
		return m.ProjectId
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

func init() { proto.RegisterFile("pod.proto", fileDescriptor_106fb77aeb685f33) }

var fileDescriptor_106fb77aeb685f33 = []byte{
	// 890 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xef, 0x6e, 0x1b, 0x45,
	0x10, 0xaf, 0xe3, 0xda, 0xb1, 0xc7, 0x76, 0x62, 0x2f, 0x15, 0x2c, 0x6e, 0x2b, 0x4c, 0x40, 0x60,
	0x81, 0xf0, 0x07, 0x97, 0x88, 0x10, 0x24, 0x90, 0x65, 0xda, 0x60, 0x51, 0xa5, 0xe6, 0x1c, 0x90,
	0xf8, 0x74, 0x5a, 0xdf, 0x4d, 0x92, 0x6b, 0xce, 0xbb, 0xa7, 0xdd, 0x3d, 0x23, 0x3f, 0x01, 0xaf,
	0xc1, 0x43, 0xf0, 0x80, 0x68, 0xff, 0x9c, 0x73, 0x49, 0x51, 0xda, 0xf4, 0x53, 0x6e, 0x7e, 0xff,
	0x66, 0x6f, 0xbc, 0x37, 0x81, 0x66, 0x26, 0xe2, 0x51, 0x26, 0x85, 0x16, 0xa4, 0x66, 0xff, 0xf4,
	0x3b, 0x17, 0xc8, 0x51, 0xb2, 0x74, 0xe4, 0x4b, 0x85, 0x72, 0x9d, 0x44, 0xe8, 0xcb, 0xf6, 0x5a,
	0xa4, 0xf9, 0xaa, 0xa8, 0xba, 0x31, 0x66, 0xa9, 0xd8, 0xac, 0x90, 0x6b, 0x87, 0x1c, 0xfc, 0x5b,
	0x05, 0x98, 0x8b, 0x78, 0xe1, 0x4c, 0xe4, 0x29, 0x40, 0x26, 0xc5, 0x6b, 0x8c, 0x74, 0x98, 0xc4,
	0xb4, 0x32, 0xa8, 0x0c, 0x9b, 0x41, 0xd3, 0x23, 0xb3, 0xd8, 0xd0, 0x3e, 0xde, 0xd0, 0x3b, 0x8e,
	0xf6, 0xc8, 0x2c, 0x26, 0x04, 0x1e, 0x72, 0xb6, 0x42, 0x5a, 0xb5, 0x84, 0x7d, 0x26, 0x14, 0x76,
	0xd7, 0x28, 0x55, 0x22, 0x38, 0x7d, 0x68, 0xe1, 0xa2, 0x24, 0x9f, 0x42, 0xbb, 0x08, 0xd3, 0x9b,
	0x0c, 0x69, 0xcd, 0xd2, 0x2d, 0x8f, 0x9d, 0x6d, 0x32, 0x24, 0x43, 0xe8, 0x16, 0x12, 0x95, 0x2f,
	0x9d, 0xac, 0x6e, 0x65, 0x7b, 0x1e, 0x5f, 0xe4, 0x4b, 0xab, 0x7c, 0x02, 0x4d, 0xd3, 0x4e, 0x65,
	0x2c, 0x42, 0xba, 0xeb, 0x0e, 0xb6, 0x05, 0xc8, 0x23, 0xa8, 0x69, 0x71, 0x85, 0x9c, 0x36, 0x2c,
	0xe3, 0x0a, 0xf3, 0x36, 0x91, 0x58, 0x65, 0x8c, 0x6f, 0xcc, 0xdb, 0x34, 0x9d, 0xc9, 0x23, 0xb3,
	0x98, 0xfc, 0x06, 0xa4, 0x68, 0xce, 0xb4, 0x96, 0xc9, 0x32, 0xd7, 0xa8, 0x28, 0x0c, 0x2a, 0xc3,
	0xd6, 0xf8, 0xc0, 0x8d, 0x6f, 0x74, 0x3d, 0x3a, 0xff, 0x67, 0xb2, 0x55, 0x06, 0x3d, 0x75, 0x1b,
	0x22, 0x27, 0x40, 0x2e, 0x85, 0xb8, 0x0a, 0x23, 0xc1, 0xcf, 0x93, 0x8b, 0x5c, 0x32, 0x6d, 0xe6,
	0xd2, 0xb2, 0x91, 0xd4, 0x47, 0xfe, 0x22, 0xc4, 0xd5, 0xb4, 0xcc, 0x07, 0xbd, 0xcb, 0xdb, 0xd0,
	0xc1, 0xdf, 0x00, 0x8f, 0xef, 0xe8, 0x4d, 0x8e, 0xcd, 0xab, 0x71, 0xcd, 0x12, 0x8e, 0x52, 0xd1,
	0xca, 0xa0, 0x3a, 0x6c, 0x8d, 0xfb, 0xbe, 0xc1, 0xb4, 0x20, 0x4a, 0x67, 0x2d, 0xa9, 0xc9, 0x21,
	0xb4, 0x13, 0xa5, 0x13, 0xe1, 0x4f, 0x69, 0x7f, 0xe6, 0xd6, 0x98, 0x78, 0xf7, 0xcc, 0x50, 0xee,
	0x30, 0x41, 0x2b, 0xb9, 0x2e, 0xc8, 0x9f, 0xd0, 0xe1, 0x22, 0xc6, 0x50, 0x61, 0x8a, 0x91, 0x16,
	0x92, 0x3e, 0xb4, 0x5d, 0xbf, 0x7d, 0xfb, 0xa4, 0x46, 0xa7, 0x22, 0xc6, 0x85, 0xb7, 0x3d, 0xe7,
	0x5a, 0x6e, 0x82, 0x36, 0x2f, 0x41, 0xe4, 0x05, 0xd4, 0x53, 0xb6, 0xc4, 0x54, 0xd1, 0x9a, 0xcd,
	0x1c, 0xbd, 0x43, 0xe6, 0x4b, 0x6b, 0x70, 0x69, 0xde, 0x4d, 0x7e, 0x87, 0x16, 0xe3, 0x5c, 0x68,
	0x3b, 0x43, 0x45, 0xeb, 0x36, 0xec, 0xd9, 0x3b, 0x84, 0x4d, 0xae, 0x5d, 0x2e, 0xb1, 0x9c, 0x43,
	0xbe, 0x81, 0x9a, 0x14, 0x29, 0x2a, 0xda, 0xb0, 0x81, 0x1f, 0xf9, 0xc0, 0x5f, 0x8f, 0x54, 0xb0,
	0x64, 0xd1, 0x36, 0x26, 0x70, 0x2a, 0x72, 0x0c, 0x6e, 0x6e, 0xa1, 0x33, 0x35, 0xad, 0xe9, 0xe3,
	0xf2, 0x78, 0x6f, 0xda, 0xc0, 0xaa, 0x03, 0xeb, 0xfd, 0x01, 0xf6, 0x24, 0x2a, 0xcd, 0xa4, 0x0e,
	0xe7, 0x22, 0x4d, 0xa2, 0x8d, 0xbd, 0x8f, 0x7b, 0xe3, 0x47, 0xde, 0x1e, 0x38, 0xd2, 0x71, 0x41,
	0x47, 0x96, 0x4b, 0xf2, 0x25, 0xec, 0xba, 0x6d, 0xa0, 0x68, 0xcb, 0x36, 0xed, 0x78, 0xd7, 0x1f,
	0x16, 0x0d, 0x0a, 0x96, 0x7c, 0x0d, 0x0d, 0x76, 0x7e, 0x9e, 0xf0, 0x44, 0x6f, 0x68, 0xdb, 0xfe,
	0xfa, 0xfb, 0x5e, 0x39, 0xf1, 0x70, 0xb0, 0x15, 0x90, 0x29, 0xec, 0x9b, 0x87, 0xb0, 0x74, 0xdf,
	0x3a, 0x6f, 0xbd, 0x6f, 0x7b, 0xc6, 0x32, 0xbd, 0xbe, 0x73, 0x2b, 0x18, 0x9c, 0xa1, 0x5c, 0x25,
	0xdc, 0x8e, 0x34, 0xbc, 0x90, 0x2c, 0xc2, 0x30, 0x43, 0x99, 0x88, 0x38, 0x54, 0x18, 0x09, 0x1e,
	0x2b, 0xba, 0x6f, 0x4f, 0xf2, 0xb9, 0x4f, 0x2d, 0xc9, 0x4f, 0x8c, 0x7a, 0x6e, 0xc5, 0x0b, 0xa7,
	0x0d, 0x9e, 0xde, 0x49, 0x93, 0x19, 0x90, 0x64, 0xc5, 0x2e, 0x30, 0xcc, 0xf2, 0x34, 0x35, 0x0d,
	0x24, 0x6a, 0x45, 0xbb, 0xf6, 0xd8, 0x8f, 0x7d, 0x83, 0x97, 0x22, 0x62, 0xe9, 0xab, 0xa5, 0xd9,
	0x7c, 0x01, 0x9e, 0xa3, 0x44, 0x1e, 0x61, 0xd0, 0xb5, 0xb6, 0x79, 0x9e, 0xa6, 0x0b, 0x67, 0x22,
	0xa3, 0xed, 0x96, 0x98, 0x44, 0x91, 0xc8, 0xb9, 0x3e, 0x35, 0x1b, 0xb0, 0x67, 0x97, 0xc9, 0xff,
	0x30, 0xe4, 0x35, 0x7c, 0xc2, 0x72, 0x2d, 0x56, 0x06, 0x08, 0xb7, 0xfb, 0xc5, 0x09, 0x42, 0xb7,
	0xa4, 0x88, 0x7d, 0xd1, 0xcf, 0x8a, 0x91, 0x17, 0xea, 0xc5, 0x8d, 0xb0, 0x33, 0x23, 0x0d, 0x9e,
	0xb0, 0x3b, 0x58, 0xf2, 0x05, 0xec, 0x27, 0x2a, 0x94, 0x4b, 0x16, 0x85, 0xc8, 0xd9, 0x32, 0xc5,
	0x98, 0x7e, 0x30, 0xa8, 0x0c, 0x1b, 0x41, 0x27, 0xb1, 0x57, 0xf3, 0xb9, 0x03, 0xfb, 0x3f, 0x41,
	0xef, 0x8d, 0x4f, 0x90, 0x74, 0xa1, 0x7a, 0x85, 0x1b, 0xff, 0x3f, 0xc0, 0x3c, 0x9a, 0x2d, 0xba,
	0x66, 0x69, 0x8e, 0x7e, 0xf1, 0xbb, 0xe2, 0x78, 0xe7, 0xa8, 0xd2, 0xff, 0x1e, 0x5a, 0xa5, 0xef,
	0xed, 0x5e, 0xd6, 0x1f, 0xa1, 0x7b, 0xfb, 0xeb, 0xba, 0x8f, 0xff, 0xab, 0x43, 0xe8, 0xdc, 0xb8,
	0xf4, 0x04, 0xa0, 0x3e, 0x49, 0xff, 0x62, 0x1b, 0xd5, 0x7d, 0x40, 0x3a, 0xd0, 0x7c, 0xc5, 0x5f,
	0xb0, 0x24, 0xcd, 0x25, 0x76, 0x2b, 0xa4, 0x09, 0xb5, 0x53, 0x5c, 0xa3, 0xec, 0xee, 0x8c, 0xff,
	0xd9, 0x81, 0xea, 0x5c, 0xc4, 0xe4, 0x08, 0x9a, 0x53, 0x89, 0x4c, 0xa3, 0x29, 0x7a, 0x6f, 0xac,
	0x82, 0xfe, 0x87, 0x1e, 0xf2, 0x75, 0x80, 0x2a, 0x13, 0x5c, 0xe1, 0xc1, 0x03, 0xe3, 0xfc, 0x19,
	0x53, 0x7c, 0x0f, 0xe7, 0x21, 0xd4, 0x4f, 0x50, 0xdf, 0xdb, 0xf6, 0x1d, 0x34, 0xe6, 0x4c, 0x47,
	0x97, 0xef, 0xd3, 0x6f, 0x9e, 0xdf, 0xbb, 0xdf, 0xb2, 0x6e, 0x89, 0x67, 0xff, 0x05, 0x00, 0x00,
	0xff, 0xff, 0x8e, 0xbd, 0xdf, 0x79, 0x73, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

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
	cc *grpc.ClientConn
}

func NewPodClient(cc *grpc.ClientConn) PodClient {
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
