// Code generated by protoc-gen-go. DO NOT EDIT.
// source: job.proto

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

type JobService struct {
	ProjectId            string               `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	ServiceId            string               `protobuf:"bytes,2,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	Name                 string               `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Version              string               `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	ServiceType          string               `protobuf:"bytes,5,opt,name=service_type,json=serviceType,proto3" json:"service_type,omitempty"`
	ServiceSubType       string               `protobuf:"bytes,6,opt,name=service_sub_type,json=serviceSubType,proto3" json:"service_sub_type,omitempty"`
	Namespace            string               `protobuf:"bytes,7,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Token                string               `protobuf:"bytes,8,opt,name=token,proto3" json:"token,omitempty"`
	CompanyId            string               `protobuf:"bytes,9,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	ServiceAttributes    *JobServiceAttribute `protobuf:"bytes,10,opt,name=service_attributes,json=serviceAttributes,proto3" json:"service_attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *JobService) Reset()         { *m = JobService{} }
func (m *JobService) String() string { return proto.CompactTextString(m) }
func (*JobService) ProtoMessage()    {}
func (*JobService) Descriptor() ([]byte, []int) {
	return fileDescriptor_f32c477d91a04ead, []int{0}
}

func (m *JobService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobService.Unmarshal(m, b)
}
func (m *JobService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobService.Marshal(b, m, deterministic)
}
func (m *JobService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobService.Merge(m, src)
}
func (m *JobService) XXX_Size() int {
	return xxx_messageInfo_JobService.Size(m)
}
func (m *JobService) XXX_DiscardUnknown() {
	xxx_messageInfo_JobService.DiscardUnknown(m)
}

var xxx_messageInfo_JobService proto.InternalMessageInfo

func (m *JobService) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *JobService) GetServiceId() string {
	if m != nil {
		return m.ServiceId
	}
	return ""
}

func (m *JobService) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *JobService) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *JobService) GetServiceType() string {
	if m != nil {
		return m.ServiceType
	}
	return ""
}

func (m *JobService) GetServiceSubType() string {
	if m != nil {
		return m.ServiceSubType
	}
	return ""
}

func (m *JobService) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *JobService) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *JobService) GetCompanyId() string {
	if m != nil {
		return m.CompanyId
	}
	return ""
}

func (m *JobService) GetServiceAttributes() *JobServiceAttribute {
	if m != nil {
		return m.ServiceAttributes
	}
	return nil
}

type JobServiceAttribute struct {
	Parallelism             *Parallelism             `protobuf:"bytes,1,opt,name=parallelism,proto3" json:"parallelism,omitempty"`
	Completions             *Completions             `protobuf:"bytes,2,opt,name=completions,proto3" json:"completions,omitempty"`
	ActiveDeadlineSeconds   *ActiveDeadlineSeconds   `protobuf:"bytes,3,opt,name=active_deadline_seconds,json=activeDeadlineSeconds,proto3" json:"active_deadline_seconds,omitempty"`
	BackoffLimit            *BackoffLimit            `protobuf:"bytes,4,opt,name=backoff_limit,json=backoffLimit,proto3" json:"backoff_limit,omitempty"`
	LabelSelector           *LabelSelectorObj        `protobuf:"bytes,5,opt,name=label_selector,json=labelSelector,proto3" json:"label_selector,omitempty"`
	ManualSelector          *ManualSelector          `protobuf:"bytes,6,opt,name=manual_selector,json=manualSelector,proto3" json:"manual_selector,omitempty"`
	IstioConfig             *IstioConfig             `protobuf:"bytes,7,opt,name=istio_config,json=istioConfig,proto3" json:"istio_config,omitempty"`
	Volumes                 []*Volume                `protobuf:"bytes,8,rep,name=volumes,proto3" json:"volumes,omitempty"`
	NodeSelector            map[string]string        `protobuf:"bytes,9,rep,name=node_selector,json=nodeSelector,proto3" json:"node_selector,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Affinity                *Affinity                `protobuf:"bytes,10,opt,name=affinity,proto3" json:"affinity,omitempty"`
	TtlSecondsAfterFinished *TTLSecondsAfterFinished `protobuf:"bytes,11,opt,name=ttl_seconds_after_finished,json=ttlSecondsAfterFinished,proto3" json:"ttl_seconds_after_finished,omitempty"`
	Labels                  map[string]string        `protobuf:"bytes,15,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Annotations             map[string]string        `protobuf:"bytes,16,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Roles                   []*K8SRbacAttribute      `protobuf:"bytes,20,rep,name=roles,proto3" json:"roles,omitempty"`
	IstioRoles              []*IstioRbacAttribute    `protobuf:"bytes,21,rep,name=istio_roles,json=istioRoles,proto3" json:"istio_roles,omitempty"`
	Containers              []*ContainerAttributes   `protobuf:"bytes,22,rep,name=containers,proto3" json:"containers,omitempty"`
	InitContainers          []*ContainerAttributes   `protobuf:"bytes,23,rep,name=init_containers,json=initContainers,proto3" json:"init_containers,omitempty"`
	IsRbacEnabled           bool                     `protobuf:"varint,24,opt,name=is_rbac_enabled,json=isRbacEnabled,proto3" json:"is_rbac_enabled,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                 `json:"-"`
	XXX_unrecognized        []byte                   `json:"-"`
	XXX_sizecache           int32                    `json:"-"`
}

func (m *JobServiceAttribute) Reset()         { *m = JobServiceAttribute{} }
func (m *JobServiceAttribute) String() string { return proto.CompactTextString(m) }
func (*JobServiceAttribute) ProtoMessage()    {}
func (*JobServiceAttribute) Descriptor() ([]byte, []int) {
	return fileDescriptor_f32c477d91a04ead, []int{1}
}

func (m *JobServiceAttribute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobServiceAttribute.Unmarshal(m, b)
}
func (m *JobServiceAttribute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobServiceAttribute.Marshal(b, m, deterministic)
}
func (m *JobServiceAttribute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobServiceAttribute.Merge(m, src)
}
func (m *JobServiceAttribute) XXX_Size() int {
	return xxx_messageInfo_JobServiceAttribute.Size(m)
}
func (m *JobServiceAttribute) XXX_DiscardUnknown() {
	xxx_messageInfo_JobServiceAttribute.DiscardUnknown(m)
}

var xxx_messageInfo_JobServiceAttribute proto.InternalMessageInfo

func (m *JobServiceAttribute) GetParallelism() *Parallelism {
	if m != nil {
		return m.Parallelism
	}
	return nil
}

func (m *JobServiceAttribute) GetCompletions() *Completions {
	if m != nil {
		return m.Completions
	}
	return nil
}

func (m *JobServiceAttribute) GetActiveDeadlineSeconds() *ActiveDeadlineSeconds {
	if m != nil {
		return m.ActiveDeadlineSeconds
	}
	return nil
}

func (m *JobServiceAttribute) GetBackoffLimit() *BackoffLimit {
	if m != nil {
		return m.BackoffLimit
	}
	return nil
}

func (m *JobServiceAttribute) GetLabelSelector() *LabelSelectorObj {
	if m != nil {
		return m.LabelSelector
	}
	return nil
}

func (m *JobServiceAttribute) GetManualSelector() *ManualSelector {
	if m != nil {
		return m.ManualSelector
	}
	return nil
}

func (m *JobServiceAttribute) GetIstioConfig() *IstioConfig {
	if m != nil {
		return m.IstioConfig
	}
	return nil
}

func (m *JobServiceAttribute) GetVolumes() []*Volume {
	if m != nil {
		return m.Volumes
	}
	return nil
}

func (m *JobServiceAttribute) GetNodeSelector() map[string]string {
	if m != nil {
		return m.NodeSelector
	}
	return nil
}

func (m *JobServiceAttribute) GetAffinity() *Affinity {
	if m != nil {
		return m.Affinity
	}
	return nil
}

func (m *JobServiceAttribute) GetTtlSecondsAfterFinished() *TTLSecondsAfterFinished {
	if m != nil {
		return m.TtlSecondsAfterFinished
	}
	return nil
}

func (m *JobServiceAttribute) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *JobServiceAttribute) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func (m *JobServiceAttribute) GetRoles() []*K8SRbacAttribute {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *JobServiceAttribute) GetIstioRoles() []*IstioRbacAttribute {
	if m != nil {
		return m.IstioRoles
	}
	return nil
}

func (m *JobServiceAttribute) GetContainers() []*ContainerAttributes {
	if m != nil {
		return m.Containers
	}
	return nil
}

func (m *JobServiceAttribute) GetInitContainers() []*ContainerAttributes {
	if m != nil {
		return m.InitContainers
	}
	return nil
}

func (m *JobServiceAttribute) GetIsRbacEnabled() bool {
	if m != nil {
		return m.IsRbacEnabled
	}
	return false
}

type Parallelism struct {
	Value                int32    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Parallelism) Reset()         { *m = Parallelism{} }
func (m *Parallelism) String() string { return proto.CompactTextString(m) }
func (*Parallelism) ProtoMessage()    {}
func (*Parallelism) Descriptor() ([]byte, []int) {
	return fileDescriptor_f32c477d91a04ead, []int{2}
}

func (m *Parallelism) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Parallelism.Unmarshal(m, b)
}
func (m *Parallelism) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Parallelism.Marshal(b, m, deterministic)
}
func (m *Parallelism) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Parallelism.Merge(m, src)
}
func (m *Parallelism) XXX_Size() int {
	return xxx_messageInfo_Parallelism.Size(m)
}
func (m *Parallelism) XXX_DiscardUnknown() {
	xxx_messageInfo_Parallelism.DiscardUnknown(m)
}

var xxx_messageInfo_Parallelism proto.InternalMessageInfo

func (m *Parallelism) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Completions struct {
	Value                int32    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Completions) Reset()         { *m = Completions{} }
func (m *Completions) String() string { return proto.CompactTextString(m) }
func (*Completions) ProtoMessage()    {}
func (*Completions) Descriptor() ([]byte, []int) {
	return fileDescriptor_f32c477d91a04ead, []int{3}
}

func (m *Completions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Completions.Unmarshal(m, b)
}
func (m *Completions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Completions.Marshal(b, m, deterministic)
}
func (m *Completions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Completions.Merge(m, src)
}
func (m *Completions) XXX_Size() int {
	return xxx_messageInfo_Completions.Size(m)
}
func (m *Completions) XXX_DiscardUnknown() {
	xxx_messageInfo_Completions.DiscardUnknown(m)
}

var xxx_messageInfo_Completions proto.InternalMessageInfo

func (m *Completions) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type BackoffLimit struct {
	Value                int32    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BackoffLimit) Reset()         { *m = BackoffLimit{} }
func (m *BackoffLimit) String() string { return proto.CompactTextString(m) }
func (*BackoffLimit) ProtoMessage()    {}
func (*BackoffLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_f32c477d91a04ead, []int{4}
}

func (m *BackoffLimit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BackoffLimit.Unmarshal(m, b)
}
func (m *BackoffLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BackoffLimit.Marshal(b, m, deterministic)
}
func (m *BackoffLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BackoffLimit.Merge(m, src)
}
func (m *BackoffLimit) XXX_Size() int {
	return xxx_messageInfo_BackoffLimit.Size(m)
}
func (m *BackoffLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_BackoffLimit.DiscardUnknown(m)
}

var xxx_messageInfo_BackoffLimit proto.InternalMessageInfo

func (m *BackoffLimit) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type ManualSelector struct {
	Value                bool     `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ManualSelector) Reset()         { *m = ManualSelector{} }
func (m *ManualSelector) String() string { return proto.CompactTextString(m) }
func (*ManualSelector) ProtoMessage()    {}
func (*ManualSelector) Descriptor() ([]byte, []int) {
	return fileDescriptor_f32c477d91a04ead, []int{5}
}

func (m *ManualSelector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ManualSelector.Unmarshal(m, b)
}
func (m *ManualSelector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ManualSelector.Marshal(b, m, deterministic)
}
func (m *ManualSelector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManualSelector.Merge(m, src)
}
func (m *ManualSelector) XXX_Size() int {
	return xxx_messageInfo_ManualSelector.Size(m)
}
func (m *ManualSelector) XXX_DiscardUnknown() {
	xxx_messageInfo_ManualSelector.DiscardUnknown(m)
}

var xxx_messageInfo_ManualSelector proto.InternalMessageInfo

func (m *ManualSelector) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

type TTLSecondsAfterFinished struct {
	Value                int32    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TTLSecondsAfterFinished) Reset()         { *m = TTLSecondsAfterFinished{} }
func (m *TTLSecondsAfterFinished) String() string { return proto.CompactTextString(m) }
func (*TTLSecondsAfterFinished) ProtoMessage()    {}
func (*TTLSecondsAfterFinished) Descriptor() ([]byte, []int) {
	return fileDescriptor_f32c477d91a04ead, []int{6}
}

func (m *TTLSecondsAfterFinished) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TTLSecondsAfterFinished.Unmarshal(m, b)
}
func (m *TTLSecondsAfterFinished) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TTLSecondsAfterFinished.Marshal(b, m, deterministic)
}
func (m *TTLSecondsAfterFinished) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TTLSecondsAfterFinished.Merge(m, src)
}
func (m *TTLSecondsAfterFinished) XXX_Size() int {
	return xxx_messageInfo_TTLSecondsAfterFinished.Size(m)
}
func (m *TTLSecondsAfterFinished) XXX_DiscardUnknown() {
	xxx_messageInfo_TTLSecondsAfterFinished.DiscardUnknown(m)
}

var xxx_messageInfo_TTLSecondsAfterFinished proto.InternalMessageInfo

func (m *TTLSecondsAfterFinished) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*JobService)(nil), "proto.JobService")
	proto.RegisterType((*JobServiceAttribute)(nil), "proto.JobServiceAttribute")
	proto.RegisterMapType((map[string]string)(nil), "proto.JobServiceAttribute.AnnotationsEntry")
	proto.RegisterMapType((map[string]string)(nil), "proto.JobServiceAttribute.LabelsEntry")
	proto.RegisterMapType((map[string]string)(nil), "proto.JobServiceAttribute.NodeSelectorEntry")
	proto.RegisterType((*Parallelism)(nil), "proto.Parallelism")
	proto.RegisterType((*Completions)(nil), "proto.Completions")
	proto.RegisterType((*BackoffLimit)(nil), "proto.BackoffLimit")
	proto.RegisterType((*ManualSelector)(nil), "proto.ManualSelector")
	proto.RegisterType((*TTLSecondsAfterFinished)(nil), "proto.TTLSecondsAfterFinished")
}

func init() {
	proto.RegisterFile("job.proto", fileDescriptor_f32c477d91a04ead)
}

var fileDescriptor_f32c477d91a04ead = []byte{
	// 875 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xdd, 0x6e, 0x1b, 0x45,
	0x14, 0xc6, 0x49, 0xed, 0xd8, 0x67, 0xfd, 0x93, 0x4c, 0x9b, 0x7a, 0xb1, 0x0a, 0x0a, 0x01, 0x05,
	0x4b, 0x85, 0x20, 0x19, 0x2a, 0x42, 0x2f, 0x82, 0x82, 0x5b, 0x90, 0x4b, 0x0a, 0x61, 0x13, 0x71,
	0xc3, 0xc5, 0x6a, 0x76, 0xf7, 0xb8, 0x9d, 0x64, 0x3c, 0x63, 0xed, 0x8c, 0x2d, 0xf9, 0x6d, 0x78,
	0x01, 0x9e, 0x8a, 0x17, 0x41, 0xf3, 0xb3, 0xd9, 0x4d, 0x9b, 0x14, 0x39, 0x57, 0xf6, 0xf9, 0xfe,
	0x76, 0x66, 0xce, 0x99, 0x81, 0xd6, 0xa5, 0x4c, 0x0e, 0xe7, 0xb9, 0xd4, 0x92, 0xd4, 0xed, 0xcf,
	0xa0, 0xf3, 0x06, 0x05, 0xe6, 0x94, 0x1f, 0xfa, 0x52, 0x61, 0xbe, 0x64, 0x29, 0xfa, 0xb2, 0xbd,
	0x94, 0x7c, 0x31, 0xf3, 0xd5, 0xfe, 0xbf, 0x1b, 0x00, 0xaf, 0x64, 0x72, 0xee, 0x24, 0xe4, 0x13,
	0x80, 0x79, 0x2e, 0x2f, 0x31, 0xd5, 0x31, 0xcb, 0xc2, 0xda, 0x5e, 0x6d, 0xd8, 0x8a, 0x5a, 0x1e,
	0x99, 0x64, 0x86, 0xf6, 0x61, 0x86, 0xde, 0x70, 0xb4, 0x47, 0x26, 0x19, 0x21, 0xf0, 0x40, 0xd0,
	0x19, 0x86, 0x9b, 0x96, 0xb0, 0xff, 0x49, 0x08, 0x5b, 0x4b, 0xcc, 0x15, 0x93, 0x22, 0x7c, 0x60,
	0xe1, 0xa2, 0x24, 0x9f, 0x41, 0xbb, 0x08, 0xd3, 0xab, 0x39, 0x86, 0x75, 0x4b, 0x07, 0x1e, 0xbb,
	0x58, 0xcd, 0x91, 0x0c, 0x61, 0xbb, 0x90, 0xa8, 0x45, 0xe2, 0x64, 0x0d, 0x2b, 0xeb, 0x7a, 0xfc,
	0x7c, 0x91, 0x58, 0xe5, 0x13, 0x68, 0x99, 0xcf, 0xa9, 0x39, 0x4d, 0x31, 0xdc, 0x72, 0x0b, 0xbb,
	0x06, 0xc8, 0x23, 0xa8, 0x6b, 0x79, 0x85, 0x22, 0x6c, 0x5a, 0xc6, 0x15, 0x66, 0x37, 0xa9, 0x9c,
	0xcd, 0xa9, 0x58, 0x99, 0xdd, 0xb4, 0x9c, 0xc9, 0x23, 0x93, 0x8c, 0x4c, 0x80, 0x14, 0x1f, 0xa7,
	0x5a, 0xe7, 0x2c, 0x59, 0x68, 0x54, 0x21, 0xec, 0xd5, 0x86, 0xc1, 0x68, 0xe0, 0x8e, 0xef, 0xb0,
	0x3c, 0xba, 0x93, 0x42, 0x12, 0xed, 0xa8, 0x77, 0x10, 0xb5, 0xff, 0x0f, 0xc0, 0xc3, 0x5b, 0xa4,
	0xe4, 0x3b, 0x08, 0xe6, 0x34, 0xa7, 0x9c, 0x23, 0x67, 0x6a, 0x66, 0xcf, 0x3b, 0x18, 0x11, 0x9f,
	0x7d, 0x56, 0x32, 0x51, 0x55, 0x66, 0x5c, 0x66, 0x95, 0x1c, 0x35, 0x93, 0x42, 0xd9, 0x36, 0x94,
	0xae, 0x71, 0xc9, 0x44, 0x55, 0x19, 0xb9, 0x80, 0x3e, 0x4d, 0x35, 0x5b, 0x62, 0x9c, 0x21, 0xcd,
	0x38, 0x13, 0x18, 0x2b, 0x4c, 0xa5, 0xc8, 0x94, 0xed, 0x57, 0x30, 0x7a, 0xe2, 0x13, 0x4e, 0xac,
	0xea, 0x85, 0x17, 0x9d, 0x3b, 0x4d, 0xb4, 0x4b, 0x6f, 0x83, 0xc9, 0x11, 0x74, 0x12, 0x9a, 0x5e,
	0xc9, 0xe9, 0x34, 0xe6, 0x6c, 0xc6, 0xb4, 0x6d, 0x72, 0x30, 0x7a, 0xe8, 0xb3, 0x7e, 0x72, 0xdc,
	0xa9, 0xa1, 0xa2, 0x76, 0x52, 0xa9, 0xc8, 0x31, 0x74, 0x39, 0x4d, 0x90, 0xc7, 0x0a, 0x39, 0xa6,
	0x5a, 0xe6, 0x76, 0x00, 0x82, 0x51, 0xdf, 0x5b, 0x4f, 0x0d, 0x79, 0xee, 0xb9, 0xdf, 0x93, 0xcb,
	0xa8, 0xc3, 0xab, 0x08, 0x39, 0x86, 0xde, 0x8c, 0x8a, 0x05, 0xad, 0x04, 0x34, 0x6c, 0xc0, 0xae,
	0x0f, 0x78, 0x6d, 0xd9, 0x42, 0x1f, 0x75, 0x67, 0x37, 0x6a, 0xf2, 0x0c, 0xda, 0x4c, 0x69, 0x26,
	0xe3, 0x54, 0x8a, 0x29, 0x7b, 0x63, 0x87, 0xa6, 0x3c, 0xc6, 0x89, 0xa1, 0xc6, 0x96, 0x89, 0x02,
	0x56, 0x16, 0xe4, 0x4b, 0xd8, 0x72, 0x17, 0x48, 0x85, 0xcd, 0xbd, 0xcd, 0x61, 0x30, 0xea, 0x78,
	0xc7, 0x9f, 0x16, 0x8d, 0x0a, 0x96, 0xfc, 0x01, 0x1d, 0x21, 0x33, 0x2c, 0x57, 0xd7, 0xb2, 0xf2,
	0xaf, 0xee, 0x9e, 0x9c, 0xc3, 0xdf, 0x64, 0x86, 0xc5, 0xfa, 0x5e, 0x0a, 0x9d, 0xaf, 0xa2, 0xb6,
	0xa8, 0x40, 0xe4, 0x29, 0x34, 0xe9, 0x74, 0xca, 0x04, 0xd3, 0x2b, 0x3f, 0x87, 0xbd, 0xa2, 0x67,
	0x1e, 0x8e, 0xae, 0x05, 0xe4, 0x2f, 0x18, 0x68, 0xcd, 0x8b, 0x1e, 0xc7, 0x74, 0xaa, 0x31, 0x8f,
	0x0d, 0xa7, 0xde, 0x62, 0x16, 0x06, 0xd6, 0xfe, 0xa9, 0xb7, 0x5f, 0x5c, 0x9c, 0xfa, 0x86, 0x9e,
	0x18, 0xd9, 0xcf, 0x5e, 0x15, 0xf5, 0xb5, 0xe6, 0xb7, 0x11, 0xe4, 0x18, 0x1a, 0xb6, 0x1b, 0x2a,
	0xec, 0xd9, 0x5d, 0x1d, 0x7c, 0x60, 0x57, 0xb6, 0x91, 0xca, 0xed, 0xc7, 0xbb, 0xc8, 0x6b, 0x08,
	0xa8, 0x10, 0x52, 0x53, 0x37, 0xc2, 0xdb, 0x36, 0xe4, 0xe9, 0x07, 0x42, 0x4e, 0x4a, 0xb5, 0x4b,
	0xaa, 0xfa, 0xc9, 0xd7, 0x50, 0xcf, 0x25, 0x47, 0x15, 0x3e, 0xb2, 0x41, 0xc5, 0x08, 0xfd, 0x7a,
	0xa4, 0xa2, 0x84, 0xa6, 0xe5, 0xd5, 0x74, 0x2a, 0xf2, 0x1c, 0x5c, 0x4b, 0x63, 0x67, 0xda, 0xb5,
	0xa6, 0x8f, 0xab, 0x9d, 0xbf, 0x69, 0x03, 0xab, 0x8e, 0xbc, 0x17, 0x52, 0x29, 0x34, 0x65, 0x02,
	0x73, 0x15, 0x3e, 0xb6, 0xd6, 0xc1, 0xf5, 0xdd, 0xf3, 0x44, 0x79, 0xf5, 0xa3, 0x8a, 0x9a, 0x8c,
	0xa1, 0x67, 0x7a, 0x13, 0x57, 0x02, 0xfa, 0xff, 0x1b, 0xd0, 0x35, 0x96, 0x71, 0x19, 0x72, 0x00,
	0x3d, 0xa6, 0xe2, 0x3c, 0xa1, 0x69, 0x8c, 0x82, 0x26, 0x1c, 0xb3, 0x30, 0xdc, 0xab, 0x0d, 0x9b,
	0x51, 0x87, 0xd9, 0xdd, 0xbe, 0x74, 0xe0, 0xe0, 0x47, 0xd8, 0x79, 0x6f, 0x9e, 0xc8, 0x36, 0x6c,
	0x5e, 0xe1, 0xca, 0x3f, 0xec, 0xe6, 0xaf, 0x79, 0x1a, 0x97, 0x94, 0x2f, 0xd0, 0xbf, 0xe6, 0xae,
	0x78, 0xbe, 0x71, 0x54, 0x1b, 0xfc, 0x00, 0x41, 0xa5, 0x75, 0x6b, 0x59, 0x8f, 0x61, 0xfb, 0xdd,
	0x86, 0xad, 0xe3, 0xdf, 0xff, 0x1c, 0x82, 0xca, 0xeb, 0x57, 0x0a, 0x8d, 0xb9, 0xee, 0x85, 0x46,
	0x54, 0x79, 0xec, 0xee, 0x10, 0x7d, 0x01, 0xed, 0xea, 0x1b, 0x74, 0x87, 0xea, 0x00, 0xba, 0x37,
	0x5f, 0x8b, 0x9b, 0xba, 0x66, 0xa1, 0xfb, 0x06, 0xfa, 0x77, 0x5c, 0x95, 0xdb, 0x83, 0x47, 0x7f,
	0x6f, 0xc0, 0xe6, 0x2b, 0x99, 0x90, 0x23, 0x68, 0x8d, 0x73, 0xa4, 0x1a, 0x4d, 0xb1, 0xf3, 0xde,
	0x9c, 0x0f, 0x1e, 0x7b, 0xc8, 0xd7, 0x11, 0xaa, 0xb9, 0x14, 0x0a, 0xf7, 0x3f, 0x32, 0xce, 0x17,
	0xc8, 0xf1, 0x1e, 0xce, 0x67, 0xd0, 0xf8, 0x05, 0xf5, 0xda, 0xb6, 0xef, 0xa1, 0x79, 0x46, 0x75,
	0xfa, 0xf6, 0x3e, 0xdf, 0x3b, 0x5b, 0xac, 0xfd, 0xbd, 0xa4, 0x61, 0x89, 0x6f, 0xff, 0x0b, 0x00,
	0x00, 0xff, 0xff, 0xb0, 0xfa, 0xf5, 0x74, 0xc8, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// JobClient is the client API for Job service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type JobClient interface {
	CreateJob(ctx context.Context, in *JobService, opts ...grpc.CallOption) (*ServiceResponse, error)
	DeleteJob(ctx context.Context, in *JobService, opts ...grpc.CallOption) (*ServiceResponse, error)
	GetJob(ctx context.Context, in *JobService, opts ...grpc.CallOption) (*ServiceResponse, error)
	PatchJob(ctx context.Context, in *JobService, opts ...grpc.CallOption) (*ServiceResponse, error)
	PutJob(ctx context.Context, in *JobService, opts ...grpc.CallOption) (*ServiceResponse, error)
}

type jobClient struct {
	cc grpc.ClientConnInterface
}

func NewJobClient(cc grpc.ClientConnInterface) JobClient {
	return &jobClient{cc}
}

func (c *jobClient) CreateJob(ctx context.Context, in *JobService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.Job/CreateJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) DeleteJob(ctx context.Context, in *JobService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.Job/DeleteJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) GetJob(ctx context.Context, in *JobService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.Job/GetJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) PatchJob(ctx context.Context, in *JobService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.Job/PatchJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobClient) PutJob(ctx context.Context, in *JobService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.Job/PutJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JobServer is the server API for Job service.
type JobServer interface {
	CreateJob(context.Context, *JobService) (*ServiceResponse, error)
	DeleteJob(context.Context, *JobService) (*ServiceResponse, error)
	GetJob(context.Context, *JobService) (*ServiceResponse, error)
	PatchJob(context.Context, *JobService) (*ServiceResponse, error)
	PutJob(context.Context, *JobService) (*ServiceResponse, error)
}

// UnimplementedJobServer can be embedded to have forward compatible implementations.
type UnimplementedJobServer struct {
}

func (*UnimplementedJobServer) CreateJob(ctx context.Context, req *JobService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateJob not implemented")
}
func (*UnimplementedJobServer) DeleteJob(ctx context.Context, req *JobService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteJob not implemented")
}
func (*UnimplementedJobServer) GetJob(ctx context.Context, req *JobService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJob not implemented")
}
func (*UnimplementedJobServer) PatchJob(ctx context.Context, req *JobService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchJob not implemented")
}
func (*UnimplementedJobServer) PutJob(ctx context.Context, req *JobService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutJob not implemented")
}

func RegisterJobServer(s *grpc.Server, srv JobServer) {
	s.RegisterService(&_Job_serviceDesc, srv)
}

func _Job_CreateJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).CreateJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Job/CreateJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).CreateJob(ctx, req.(*JobService))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_DeleteJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).DeleteJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Job/DeleteJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).DeleteJob(ctx, req.(*JobService))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_GetJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).GetJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Job/GetJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).GetJob(ctx, req.(*JobService))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_PatchJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).PatchJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Job/PatchJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).PatchJob(ctx, req.(*JobService))
	}
	return interceptor(ctx, in, info, handler)
}

func _Job_PutJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServer).PutJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Job/PutJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServer).PutJob(ctx, req.(*JobService))
	}
	return interceptor(ctx, in, info, handler)
}

var _Job_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Job",
	HandlerType: (*JobServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateJob",
			Handler:    _Job_CreateJob_Handler,
		},
		{
			MethodName: "DeleteJob",
			Handler:    _Job_DeleteJob_Handler,
		},
		{
			MethodName: "GetJob",
			Handler:    _Job_GetJob_Handler,
		},
		{
			MethodName: "PatchJob",
			Handler:    _Job_PatchJob_Handler,
		},
		{
			MethodName: "PutJob",
			Handler:    _Job_PutJob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "job.proto",
}
