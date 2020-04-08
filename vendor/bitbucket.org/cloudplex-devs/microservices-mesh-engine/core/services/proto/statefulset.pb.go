// Code generated by protoc-gen-go. DO NOT EDIT.
// source: statefulset.proto

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

type PodManagementPolicyType int32

const (
	PodManagementPolicyType_OrderedReady PodManagementPolicyType = 0
	PodManagementPolicyType_Parallel     PodManagementPolicyType = 1
)

var PodManagementPolicyType_name = map[int32]string{
	0: "OrderedReady",
	1: "Parallel",
}

var PodManagementPolicyType_value = map[string]int32{
	"OrderedReady": 0,
	"Parallel":     1,
}

func (x PodManagementPolicyType) String() string {
	return proto.EnumName(PodManagementPolicyType_name, int32(x))
}

func (PodManagementPolicyType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4f0a4b9f0c7482c8, []int{0}
}

type StatefulSetService struct {
	ProjectId            string                        `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	ServiceId            string                        `protobuf:"bytes,2,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	Name                 string                        `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Version              string                        `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	ServiceType          string                        `protobuf:"bytes,5,opt,name=service_type,json=serviceType,proto3" json:"service_type,omitempty"`
	ServiceSubType       string                        `protobuf:"bytes,6,opt,name=service_sub_type,json=serviceSubType,proto3" json:"service_sub_type,omitempty"`
	Namespace            string                        `protobuf:"bytes,7,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Token                string                        `protobuf:"bytes,8,opt,name=token,proto3" json:"token,omitempty"`
	CompanyId            string                        `protobuf:"bytes,9,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	ServiceAttributes    *StatefulSetServiceAttributes `protobuf:"bytes,10,opt,name=service_attributes,json=serviceAttributes,proto3" json:"service_attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *StatefulSetService) Reset()         { *m = StatefulSetService{} }
func (m *StatefulSetService) String() string { return proto.CompactTextString(m) }
func (*StatefulSetService) ProtoMessage()    {}
func (*StatefulSetService) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f0a4b9f0c7482c8, []int{0}
}

func (m *StatefulSetService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatefulSetService.Unmarshal(m, b)
}
func (m *StatefulSetService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatefulSetService.Marshal(b, m, deterministic)
}
func (m *StatefulSetService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatefulSetService.Merge(m, src)
}
func (m *StatefulSetService) XXX_Size() int {
	return xxx_messageInfo_StatefulSetService.Size(m)
}
func (m *StatefulSetService) XXX_DiscardUnknown() {
	xxx_messageInfo_StatefulSetService.DiscardUnknown(m)
}

var xxx_messageInfo_StatefulSetService proto.InternalMessageInfo

func (m *StatefulSetService) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *StatefulSetService) GetServiceId() string {
	if m != nil {
		return m.ServiceId
	}
	return ""
}

func (m *StatefulSetService) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StatefulSetService) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *StatefulSetService) GetServiceType() string {
	if m != nil {
		return m.ServiceType
	}
	return ""
}

func (m *StatefulSetService) GetServiceSubType() string {
	if m != nil {
		return m.ServiceSubType
	}
	return ""
}

func (m *StatefulSetService) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *StatefulSetService) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *StatefulSetService) GetCompanyId() string {
	if m != nil {
		return m.CompanyId
	}
	return ""
}

func (m *StatefulSetService) GetServiceAttributes() *StatefulSetServiceAttributes {
	if m != nil {
		return m.ServiceAttributes
	}
	return nil
}

type StatefulSetServiceAttributes struct {
	Labels                        map[string]string               `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Annotations                   map[string]string               `protobuf:"bytes,2,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	LabelSelector                 *LabelSelectorObj               `protobuf:"bytes,3,opt,name=label_selector,json=labelSelector,proto3" json:"label_selector,omitempty"`
	Replicas                      *Replicas                       `protobuf:"bytes,4,opt,name=replicas,proto3" json:"replicas,omitempty"`
	RevisionHistoryLimit          *RevisionHistoryLimit           `protobuf:"bytes,5,opt,name=revision_history_limit,json=revisionHistoryLimit,proto3" json:"revision_history_limit,omitempty"`
	UpdateStrategy                *StateFulSetUpdateStrategy      `protobuf:"bytes,6,opt,name=update_strategy,json=updateStrategy,proto3" json:"update_strategy,omitempty"`
	Containers                    []*ContainerAttributes          `protobuf:"bytes,7,rep,name=containers,proto3" json:"containers,omitempty"`
	InitContainers                []*ContainerAttributes          `protobuf:"bytes,8,rep,name=init_containers,json=initContainers,proto3" json:"init_containers,omitempty"`
	NodeSelector                  map[string]string               `protobuf:"bytes,9,rep,name=node_selector,json=nodeSelector,proto3" json:"node_selector,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	IstioConfig                   *IstioConfig                    `protobuf:"bytes,10,opt,name=istio_config,json=istioConfig,proto3" json:"istio_config,omitempty"`
	IsRbacEnabled                 bool                            `protobuf:"varint,11,opt,name=is_rbac_enabled,json=isRbacEnabled,proto3" json:"is_rbac_enabled,omitempty"`
	Roles                         []*K8SRbacAttribute             `protobuf:"bytes,12,rep,name=roles,proto3" json:"roles,omitempty"`
	IstioRoles                    []*IstioRbacAttribute           `protobuf:"bytes,13,rep,name=istio_roles,json=istioRoles,proto3" json:"istio_roles,omitempty"`
	Volumes                       []*Volume                       `protobuf:"bytes,14,rep,name=volumes,proto3" json:"volumes,omitempty"`
	Affinity                      *Affinity                       `protobuf:"bytes,15,opt,name=affinity,proto3" json:"affinity,omitempty"`
	ServiceName                   string                          `protobuf:"bytes,16,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	TerminationGracePeriodSeconds *TerminationGracePeriodSeconds  `protobuf:"bytes,17,opt,name=termination_grace_period_seconds,json=terminationGracePeriodSeconds,proto3" json:"termination_grace_period_seconds,omitempty"`
	VolumeClaimTemplates          []*PersistentVolumeClaimService `protobuf:"bytes,18,rep,name=volume_claim_templates,json=volumeClaimTemplates,proto3" json:"volume_claim_templates,omitempty"`
	PodManagementPolicy           PodManagementPolicyType         `protobuf:"varint,19,opt,name=pod_management_policy,json=podManagementPolicy,proto3,enum=proto.PodManagementPolicyType" json:"pod_management_policy,omitempty"`
	XXX_NoUnkeyedLiteral          struct{}                        `json:"-"`
	XXX_unrecognized              []byte                          `json:"-"`
	XXX_sizecache                 int32                           `json:"-"`
}

func (m *StatefulSetServiceAttributes) Reset()         { *m = StatefulSetServiceAttributes{} }
func (m *StatefulSetServiceAttributes) String() string { return proto.CompactTextString(m) }
func (*StatefulSetServiceAttributes) ProtoMessage()    {}
func (*StatefulSetServiceAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f0a4b9f0c7482c8, []int{1}
}

func (m *StatefulSetServiceAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatefulSetServiceAttributes.Unmarshal(m, b)
}
func (m *StatefulSetServiceAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatefulSetServiceAttributes.Marshal(b, m, deterministic)
}
func (m *StatefulSetServiceAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatefulSetServiceAttributes.Merge(m, src)
}
func (m *StatefulSetServiceAttributes) XXX_Size() int {
	return xxx_messageInfo_StatefulSetServiceAttributes.Size(m)
}
func (m *StatefulSetServiceAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_StatefulSetServiceAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_StatefulSetServiceAttributes proto.InternalMessageInfo

func (m *StatefulSetServiceAttributes) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *StatefulSetServiceAttributes) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func (m *StatefulSetServiceAttributes) GetLabelSelector() *LabelSelectorObj {
	if m != nil {
		return m.LabelSelector
	}
	return nil
}

func (m *StatefulSetServiceAttributes) GetReplicas() *Replicas {
	if m != nil {
		return m.Replicas
	}
	return nil
}

func (m *StatefulSetServiceAttributes) GetRevisionHistoryLimit() *RevisionHistoryLimit {
	if m != nil {
		return m.RevisionHistoryLimit
	}
	return nil
}

func (m *StatefulSetServiceAttributes) GetUpdateStrategy() *StateFulSetUpdateStrategy {
	if m != nil {
		return m.UpdateStrategy
	}
	return nil
}

func (m *StatefulSetServiceAttributes) GetContainers() []*ContainerAttributes {
	if m != nil {
		return m.Containers
	}
	return nil
}

func (m *StatefulSetServiceAttributes) GetInitContainers() []*ContainerAttributes {
	if m != nil {
		return m.InitContainers
	}
	return nil
}

func (m *StatefulSetServiceAttributes) GetNodeSelector() map[string]string {
	if m != nil {
		return m.NodeSelector
	}
	return nil
}

func (m *StatefulSetServiceAttributes) GetIstioConfig() *IstioConfig {
	if m != nil {
		return m.IstioConfig
	}
	return nil
}

func (m *StatefulSetServiceAttributes) GetIsRbacEnabled() bool {
	if m != nil {
		return m.IsRbacEnabled
	}
	return false
}

func (m *StatefulSetServiceAttributes) GetRoles() []*K8SRbacAttribute {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *StatefulSetServiceAttributes) GetIstioRoles() []*IstioRbacAttribute {
	if m != nil {
		return m.IstioRoles
	}
	return nil
}

func (m *StatefulSetServiceAttributes) GetVolumes() []*Volume {
	if m != nil {
		return m.Volumes
	}
	return nil
}

func (m *StatefulSetServiceAttributes) GetAffinity() *Affinity {
	if m != nil {
		return m.Affinity
	}
	return nil
}

func (m *StatefulSetServiceAttributes) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *StatefulSetServiceAttributes) GetTerminationGracePeriodSeconds() *TerminationGracePeriodSeconds {
	if m != nil {
		return m.TerminationGracePeriodSeconds
	}
	return nil
}

func (m *StatefulSetServiceAttributes) GetVolumeClaimTemplates() []*PersistentVolumeClaimService {
	if m != nil {
		return m.VolumeClaimTemplates
	}
	return nil
}

func (m *StatefulSetServiceAttributes) GetPodManagementPolicy() PodManagementPolicyType {
	if m != nil {
		return m.PodManagementPolicy
	}
	return PodManagementPolicyType_OrderedReady
}

func init() {
	proto.RegisterEnum("proto.PodManagementPolicyType", PodManagementPolicyType_name, PodManagementPolicyType_value)
	proto.RegisterType((*StatefulSetService)(nil), "proto.StatefulSetService")
	proto.RegisterType((*StatefulSetServiceAttributes)(nil), "proto.StatefulSetServiceAttributes")
	proto.RegisterMapType((map[string]string)(nil), "proto.StatefulSetServiceAttributes.AnnotationsEntry")
	proto.RegisterMapType((map[string]string)(nil), "proto.StatefulSetServiceAttributes.LabelsEntry")
	proto.RegisterMapType((map[string]string)(nil), "proto.StatefulSetServiceAttributes.NodeSelectorEntry")
}

func init() {
	proto.RegisterFile("statefulset.proto", fileDescriptor_4f0a4b9f0c7482c8)
}

var fileDescriptor_4f0a4b9f0c7482c8 = []byte{
	// 950 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xdf, 0x6f, 0x1b, 0x45,
	0x10, 0xae, 0x93, 0xe6, 0xd7, 0x9c, 0xed, 0x38, 0xdb, 0x90, 0x5e, 0xd3, 0x16, 0x99, 0x80, 0xc0,
	0x02, 0x11, 0x24, 0x43, 0xa5, 0x36, 0x0f, 0x45, 0x91, 0x09, 0x69, 0x44, 0x69, 0xcd, 0x3a, 0x54,
	0x82, 0x97, 0xd3, 0xfa, 0x6e, 0xe2, 0x6e, 0x7b, 0xde, 0x3d, 0xed, 0xae, 0x2d, 0xdd, 0xff, 0xc7,
	0x9f, 0xc4, 0x03, 0x8f, 0x68, 0x7f, 0x9c, 0x7d, 0x69, 0x68, 0x49, 0xd5, 0x3e, 0xd9, 0xf3, 0x7d,
	0xdf, 0x7c, 0x33, 0xb7, 0x3b, 0x37, 0x07, 0x3b, 0xda, 0x30, 0x83, 0x17, 0xb3, 0x5c, 0xa3, 0x39,
	0x2c, 0x94, 0x34, 0x92, 0xac, 0xb9, 0x9f, 0xfd, 0xd6, 0x04, 0x05, 0x2a, 0x96, 0x1f, 0x86, 0x50,
	0xa3, 0x9a, 0xf3, 0x14, 0x43, 0xd8, 0x9c, 0xcb, 0x7c, 0x36, 0xad, 0xa2, 0xbb, 0x05, 0x2a, 0xcd,
	0xb5, 0x41, 0x61, 0x3c, 0x9e, 0xe6, 0x8c, 0x4f, 0x3d, 0x79, 0xf0, 0xf7, 0x0a, 0x90, 0x51, 0xa8,
	0x32, 0x42, 0x33, 0xf2, 0x3e, 0xe4, 0x3e, 0x40, 0xa1, 0xe4, 0x2b, 0x4c, 0x4d, 0xc2, 0xb3, 0xb8,
	0xd1, 0x6d, 0xf4, 0xb6, 0xe8, 0x56, 0x40, 0xce, 0x32, 0x4b, 0x87, 0x8a, 0x96, 0x5e, 0xf1, 0x74,
	0x40, 0xce, 0x32, 0x42, 0xe0, 0xa6, 0x60, 0x53, 0x8c, 0x57, 0x1d, 0xe1, 0xfe, 0x93, 0x18, 0x36,
	0xe6, 0xb6, 0x0f, 0x29, 0xe2, 0x9b, 0x0e, 0xae, 0x42, 0xf2, 0x19, 0x34, 0x2b, 0x33, 0x53, 0x16,
	0x18, 0xaf, 0x39, 0x3a, 0x0a, 0xd8, 0x79, 0x59, 0x20, 0xe9, 0x41, 0xa7, 0x92, 0xe8, 0xd9, 0xd8,
	0xcb, 0xd6, 0x9d, 0xac, 0x1d, 0xf0, 0xd1, 0x6c, 0xec, 0x94, 0xf7, 0x60, 0xcb, 0x96, 0xd3, 0x05,
	0x4b, 0x31, 0xde, 0xf0, 0x8d, 0x2d, 0x00, 0xb2, 0x0b, 0x6b, 0x46, 0xbe, 0x46, 0x11, 0x6f, 0x3a,
	0xc6, 0x07, 0xf6, 0x69, 0x52, 0x39, 0x2d, 0x98, 0x28, 0xed, 0xd3, 0x6c, 0xf9, 0xa4, 0x80, 0x9c,
	0x65, 0x84, 0x02, 0xa9, 0x8a, 0x33, 0x63, 0x14, 0x1f, 0xcf, 0x0c, 0xea, 0x18, 0xba, 0x8d, 0x5e,
	0xd4, 0xff, 0xdc, 0x1f, 0xe3, 0xe1, 0xd5, 0x23, 0x3c, 0x5e, 0x48, 0xe9, 0x8e, 0x7e, 0x13, 0x3a,
	0xf8, 0x2b, 0x82, 0x7b, 0xef, 0xca, 0x21, 0xa7, 0xb0, 0x9e, 0xb3, 0x31, 0xe6, 0x3a, 0x6e, 0x74,
	0x57, 0x7b, 0x51, 0xff, 0xbb, 0x6b, 0x14, 0x3a, 0x7c, 0xea, 0x32, 0x4e, 0x84, 0x51, 0x25, 0x0d,
	0xe9, 0xe4, 0x05, 0x44, 0x4c, 0x08, 0x69, 0x98, 0xe1, 0x52, 0xe8, 0x78, 0xc5, 0xb9, 0xfd, 0x70,
	0x1d, 0xb7, 0xe3, 0x65, 0x9a, 0xb7, 0xac, 0x1b, 0x91, 0xc7, 0xd0, 0x76, 0x15, 0x12, 0x8d, 0x39,
	0xa6, 0x46, 0x2a, 0x77, 0xdb, 0x51, 0xff, 0x76, 0xb0, 0x76, 0xbd, 0x8c, 0x02, 0xf7, 0x7c, 0xfc,
	0x8a, 0xb6, 0xf2, 0x3a, 0x42, 0xbe, 0x81, 0x4d, 0x85, 0x45, 0xce, 0x53, 0xa6, 0xdd, 0x40, 0x44,
	0xfd, 0xed, 0x90, 0x49, 0x03, 0x4c, 0x17, 0x02, 0xf2, 0x1b, 0xec, 0x29, 0x9c, 0x73, 0x3b, 0x2e,
	0xc9, 0x4b, 0xae, 0x8d, 0x54, 0x65, 0x92, 0xf3, 0x29, 0x37, 0x6e, 0x58, 0xa2, 0xfe, 0xdd, 0x45,
	0xaa, 0x17, 0x3d, 0xf1, 0x9a, 0xa7, 0x56, 0x42, 0x77, 0xd5, 0x7f, 0xa0, 0xe4, 0x0c, 0xb6, 0x67,
	0x45, 0xc6, 0x0c, 0x26, 0xda, 0x28, 0x66, 0x70, 0x52, 0xba, 0x89, 0x8a, 0xfa, 0xdd, 0xfa, 0xd9,
	0xfc, 0xec, 0xce, 0xe6, 0x77, 0x27, 0x1c, 0x05, 0x1d, 0x6d, 0xcf, 0x2e, 0xc5, 0xe4, 0xc8, 0xce,
	0x8f, 0x30, 0x8c, 0x0b, 0x54, 0x3a, 0xde, 0x70, 0x27, 0xbc, 0x1f, 0x5c, 0x06, 0x15, 0x51, 0x9b,
	0x87, 0x9a, 0x9a, 0x0c, 0x60, 0x9b, 0x0b, 0x6e, 0x92, 0x9a, 0xc1, 0xe6, 0xff, 0x1a, 0xb4, 0x6d,
	0xca, 0x60, 0x69, 0xf2, 0x27, 0xb4, 0x84, 0xcc, 0x70, 0x79, 0x15, 0x5b, 0xce, 0xe2, 0xc1, 0x75,
	0x6e, 0xf9, 0x99, 0xcc, 0xb0, 0xba, 0x14, 0x7f, 0xcd, 0x4d, 0x51, 0x83, 0xc8, 0x03, 0x68, 0x72,
	0x6d, 0xb8, 0xb4, 0x1d, 0x5e, 0xf0, 0x49, 0x98, 0x7b, 0x12, 0xac, 0xcf, 0x2c, 0x35, 0x70, 0x0c,
	0x8d, 0xf8, 0x32, 0x20, 0x5f, 0xc2, 0x36, 0xd7, 0x89, 0x1a, 0xb3, 0x34, 0x41, 0xc1, 0xc6, 0x39,
	0x66, 0x71, 0xd4, 0x6d, 0xf4, 0x36, 0x69, 0x8b, 0x6b, 0x3a, 0x66, 0xe9, 0x89, 0x07, 0xc9, 0xb7,
	0xb0, 0xa6, 0x64, 0x8e, 0x3a, 0x6e, 0xba, 0x96, 0xab, 0xe9, 0xf9, 0xe5, 0xa1, 0x53, 0x2d, 0xfa,
	0xa4, 0x5e, 0x45, 0x8e, 0xc0, 0x57, 0x49, 0x7c, 0x52, 0xcb, 0x25, 0xdd, 0xa9, 0x37, 0x73, 0x39,
	0x0d, 0x9c, 0x9a, 0xba, 0xdc, 0xaf, 0x60, 0xc3, 0xef, 0x3f, 0x1d, 0xb7, 0x5d, 0x5e, 0x2b, 0xe4,
	0xbd, 0x70, 0x28, 0xad, 0x58, 0x3b, 0x9a, 0xec, 0xe2, 0xc2, 0x9e, 0x71, 0x19, 0x6f, 0x5f, 0x1a,
	0xcd, 0xe3, 0x00, 0xd3, 0x85, 0xa0, 0xbe, 0xbd, 0xdc, 0xce, 0xeb, 0x5c, 0xda, 0x5e, 0xcf, 0xec,
	0xea, 0x9b, 0x42, 0xd7, 0xa0, 0x9a, 0x72, 0xe1, 0x5e, 0x9d, 0x64, 0xa2, 0x58, 0x8a, 0x49, 0x81,
	0x8a, 0xcb, 0x2c, 0xd1, 0x98, 0x4a, 0x91, 0xe9, 0x78, 0xc7, 0xd5, 0xf9, 0x22, 0xd4, 0x39, 0x5f,
	0xca, 0x4f, 0xad, 0x7a, 0xe8, 0xc4, 0x23, 0xaf, 0xa5, 0xf7, 0xcd, 0xbb, 0x68, 0xf2, 0x07, 0xec,
	0xf9, 0x27, 0x49, 0xdc, 0xa2, 0x4f, 0x0c, 0x4e, 0x8b, 0x9c, 0xd9, 0x9d, 0x45, 0xdc, 0x63, 0x57,
	0x3b, 0x6b, 0xb8, 0xf8, 0x2c, 0xf8, 0x03, 0x18, 0x58, 0x75, 0x18, 0x10, 0xba, 0x3b, 0x5f, 0x62,
	0xe7, 0x95, 0x01, 0xa1, 0xf0, 0x49, 0x21, 0xb3, 0x64, 0xca, 0x04, 0x9b, 0xe0, 0x14, 0x85, 0x49,
	0x0a, 0x99, 0xf3, 0xb4, 0x8c, 0x6f, 0x75, 0x1b, 0xbd, 0x76, 0xff, 0xd3, 0xca, 0x59, 0x66, 0xbf,
	0x2e, 0x24, 0x43, 0xa7, 0xb0, 0xcb, 0x99, 0xde, 0x2a, 0xae, 0x12, 0xfb, 0x8f, 0x20, 0xaa, 0xed,
	0x2d, 0xd2, 0x81, 0xd5, 0xd7, 0x58, 0x86, 0x4f, 0x8e, 0xfd, 0x6b, 0x97, 0xf6, 0x9c, 0xe5, 0x33,
	0x0c, 0xdf, 0x19, 0x1f, 0x1c, 0xad, 0x3c, 0x6c, 0xec, 0x3f, 0x86, 0xce, 0x9b, 0x4b, 0xea, 0xbd,
	0xf2, 0x7f, 0x84, 0x9d, 0x2b, 0xe3, 0xff, 0x3e, 0x06, 0x5f, 0x3f, 0x82, 0xdb, 0x6f, 0x79, 0x56,
	0xd2, 0x81, 0xe6, 0x73, 0x95, 0xa1, 0xc2, 0x8c, 0x22, 0xcb, 0xca, 0xce, 0x0d, 0xd2, 0x84, 0xcd,
	0x21, 0x53, 0x2c, 0xcf, 0x31, 0xef, 0x34, 0xfa, 0xff, 0xac, 0x40, 0x54, 0x7b, 0x31, 0xc9, 0x13,
	0xd8, 0x19, 0x28, 0x74, 0x6b, 0x65, 0x09, 0xde, 0x79, 0xeb, 0x1b, 0xbc, 0xbf, 0x57, 0x51, 0xe1,
	0xc2, 0x50, 0x17, 0x52, 0x68, 0x3c, 0xb8, 0x61, 0x9d, 0x7e, 0xc2, 0x1c, 0x3f, 0x82, 0xd3, 0x09,
	0xb4, 0x4f, 0xd1, 0x7c, 0xb0, 0xcd, 0x29, 0x74, 0x86, 0xcc, 0xa4, 0x2f, 0x3f, 0x46, 0x3f, 0xc3,
	0xd9, 0x07, 0xf7, 0x33, 0x5e, 0x77, 0xc4, 0xf7, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xc1, 0x43,
	0xb1, 0x3f, 0x5f, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StatefulSetClient is the client API for StatefulSet service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StatefulSetClient interface {
	CreateStatefulSet(ctx context.Context, in *StatefulSetService, opts ...grpc.CallOption) (*ServiceResponse, error)
	DeleteStatefulSet(ctx context.Context, in *StatefulSetService, opts ...grpc.CallOption) (*ServiceResponse, error)
	GetStatefulSet(ctx context.Context, in *StatefulSetService, opts ...grpc.CallOption) (*ServiceResponse, error)
	PatchStatefulSet(ctx context.Context, in *StatefulSetService, opts ...grpc.CallOption) (*ServiceResponse, error)
	PutStatefulSet(ctx context.Context, in *StatefulSetService, opts ...grpc.CallOption) (*ServiceResponse, error)
}

type statefulSetClient struct {
	cc grpc.ClientConnInterface
}

func NewStatefulSetClient(cc grpc.ClientConnInterface) StatefulSetClient {
	return &statefulSetClient{cc}
}

func (c *statefulSetClient) CreateStatefulSet(ctx context.Context, in *StatefulSetService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.StatefulSet/CreateStatefulSet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statefulSetClient) DeleteStatefulSet(ctx context.Context, in *StatefulSetService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.StatefulSet/DeleteStatefulSet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statefulSetClient) GetStatefulSet(ctx context.Context, in *StatefulSetService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.StatefulSet/GetStatefulSet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statefulSetClient) PatchStatefulSet(ctx context.Context, in *StatefulSetService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.StatefulSet/PatchStatefulSet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statefulSetClient) PutStatefulSet(ctx context.Context, in *StatefulSetService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.StatefulSet/PutStatefulSet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StatefulSetServer is the server API for StatefulSet service.
type StatefulSetServer interface {
	CreateStatefulSet(context.Context, *StatefulSetService) (*ServiceResponse, error)
	DeleteStatefulSet(context.Context, *StatefulSetService) (*ServiceResponse, error)
	GetStatefulSet(context.Context, *StatefulSetService) (*ServiceResponse, error)
	PatchStatefulSet(context.Context, *StatefulSetService) (*ServiceResponse, error)
	PutStatefulSet(context.Context, *StatefulSetService) (*ServiceResponse, error)
}

// UnimplementedStatefulSetServer can be embedded to have forward compatible implementations.
type UnimplementedStatefulSetServer struct {
}

func (*UnimplementedStatefulSetServer) CreateStatefulSet(ctx context.Context, req *StatefulSetService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStatefulSet not implemented")
}
func (*UnimplementedStatefulSetServer) DeleteStatefulSet(ctx context.Context, req *StatefulSetService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStatefulSet not implemented")
}
func (*UnimplementedStatefulSetServer) GetStatefulSet(ctx context.Context, req *StatefulSetService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatefulSet not implemented")
}
func (*UnimplementedStatefulSetServer) PatchStatefulSet(ctx context.Context, req *StatefulSetService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchStatefulSet not implemented")
}
func (*UnimplementedStatefulSetServer) PutStatefulSet(ctx context.Context, req *StatefulSetService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutStatefulSet not implemented")
}

func RegisterStatefulSetServer(s *grpc.Server, srv StatefulSetServer) {
	s.RegisterService(&_StatefulSet_serviceDesc, srv)
}

func _StatefulSet_CreateStatefulSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatefulSetService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatefulSetServer).CreateStatefulSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.StatefulSet/CreateStatefulSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatefulSetServer).CreateStatefulSet(ctx, req.(*StatefulSetService))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatefulSet_DeleteStatefulSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatefulSetService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatefulSetServer).DeleteStatefulSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.StatefulSet/DeleteStatefulSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatefulSetServer).DeleteStatefulSet(ctx, req.(*StatefulSetService))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatefulSet_GetStatefulSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatefulSetService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatefulSetServer).GetStatefulSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.StatefulSet/GetStatefulSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatefulSetServer).GetStatefulSet(ctx, req.(*StatefulSetService))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatefulSet_PatchStatefulSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatefulSetService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatefulSetServer).PatchStatefulSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.StatefulSet/PatchStatefulSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatefulSetServer).PatchStatefulSet(ctx, req.(*StatefulSetService))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatefulSet_PutStatefulSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatefulSetService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatefulSetServer).PutStatefulSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.StatefulSet/PutStatefulSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatefulSetServer).PutStatefulSet(ctx, req.(*StatefulSetService))
	}
	return interceptor(ctx, in, info, handler)
}

var _StatefulSet_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.StatefulSet",
	HandlerType: (*StatefulSetServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateStatefulSet",
			Handler:    _StatefulSet_CreateStatefulSet_Handler,
		},
		{
			MethodName: "DeleteStatefulSet",
			Handler:    _StatefulSet_DeleteStatefulSet_Handler,
		},
		{
			MethodName: "GetStatefulSet",
			Handler:    _StatefulSet_GetStatefulSet_Handler,
		},
		{
			MethodName: "PatchStatefulSet",
			Handler:    _StatefulSet_PatchStatefulSet_Handler,
		},
		{
			MethodName: "PutStatefulSet",
			Handler:    _StatefulSet_PutStatefulSet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "statefulset.proto",
}
