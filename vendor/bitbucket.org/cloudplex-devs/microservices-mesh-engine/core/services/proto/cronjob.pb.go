// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cronjob.proto

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

type ConcurrencyPolicy int32

const (
	ConcurrencyPolicy_Allow   ConcurrencyPolicy = 0
	ConcurrencyPolicy_Forbid  ConcurrencyPolicy = 1
	ConcurrencyPolicy_Replace ConcurrencyPolicy = 2
)

var ConcurrencyPolicy_name = map[int32]string{
	0: "Allow",
	1: "Forbid",
	2: "Replace",
}

var ConcurrencyPolicy_value = map[string]int32{
	"Allow":   0,
	"Forbid":  1,
	"Replace": 2,
}

func (x ConcurrencyPolicy) String() string {
	return proto.EnumName(ConcurrencyPolicy_name, int32(x))
}

func (ConcurrencyPolicy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9693a1980b1a682f, []int{0}
}

type CronJobService struct {
	ProjectId               string                   `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	ServiceId               string                   `protobuf:"bytes,2,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	Name                    string                   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Version                 string                   `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	ServiceType             string                   `protobuf:"bytes,5,opt,name=service_type,json=serviceType,proto3" json:"service_type,omitempty"`
	ServiceSubType          string                   `protobuf:"bytes,6,opt,name=service_sub_type,json=serviceSubType,proto3" json:"service_sub_type,omitempty"`
	Namespace               string                   `protobuf:"bytes,7,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Token                   string                   `protobuf:"bytes,8,opt,name=token,proto3" json:"token,omitempty"`
	CompanyId               string                   `protobuf:"bytes,9,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	CronJobServiceAttribute *CronJobServiceAttribute `protobuf:"bytes,10,opt,name=CronJobServiceAttribute,proto3" json:"CronJobServiceAttribute,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                 `json:"-"`
	XXX_unrecognized        []byte                   `json:"-"`
	XXX_sizecache           int32                    `json:"-"`
}

func (m *CronJobService) Reset()         { *m = CronJobService{} }
func (m *CronJobService) String() string { return proto.CompactTextString(m) }
func (*CronJobService) ProtoMessage()    {}
func (*CronJobService) Descriptor() ([]byte, []int) {
	return fileDescriptor_9693a1980b1a682f, []int{0}
}

func (m *CronJobService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CronJobService.Unmarshal(m, b)
}
func (m *CronJobService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CronJobService.Marshal(b, m, deterministic)
}
func (m *CronJobService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CronJobService.Merge(m, src)
}
func (m *CronJobService) XXX_Size() int {
	return xxx_messageInfo_CronJobService.Size(m)
}
func (m *CronJobService) XXX_DiscardUnknown() {
	xxx_messageInfo_CronJobService.DiscardUnknown(m)
}

var xxx_messageInfo_CronJobService proto.InternalMessageInfo

func (m *CronJobService) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *CronJobService) GetServiceId() string {
	if m != nil {
		return m.ServiceId
	}
	return ""
}

func (m *CronJobService) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CronJobService) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *CronJobService) GetServiceType() string {
	if m != nil {
		return m.ServiceType
	}
	return ""
}

func (m *CronJobService) GetServiceSubType() string {
	if m != nil {
		return m.ServiceSubType
	}
	return ""
}

func (m *CronJobService) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *CronJobService) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *CronJobService) GetCompanyId() string {
	if m != nil {
		return m.CompanyId
	}
	return ""
}

func (m *CronJobService) GetCronJobServiceAttribute() *CronJobServiceAttribute {
	if m != nil {
		return m.CronJobServiceAttribute
	}
	return nil
}

type CronJobServiceAttribute struct {
	Schedule                   string                      `protobuf:"bytes,1,opt,name=schedule,proto3" json:"schedule,omitempty"`
	StartingDeadlineSeconds    *StartingDeadlineSeconds    `protobuf:"bytes,2,opt,name=starting_deadline_seconds,json=startingDeadlineSeconds,proto3" json:"starting_deadline_seconds,omitempty"`
	ConcurrencyPolicy          ConcurrencyPolicy           `protobuf:"varint,3,opt,name=concurrency_policy,json=concurrencyPolicy,proto3,enum=proto.ConcurrencyPolicy" json:"concurrency_policy,omitempty"`
	Suspend                    *Suspend                    `protobuf:"bytes,4,opt,name=suspend,proto3" json:"suspend,omitempty"`
	FailedJobsHistoryLimit     *FailedJobsHistoryLimit     `protobuf:"bytes,5,opt,name=failed_jobs_history_limit,json=failedJobsHistoryLimit,proto3" json:"failed_jobs_history_limit,omitempty"`
	SuccessfulJobsHistoryLimit *SuccessfulJobsHistoryLimit `protobuf:"bytes,6,opt,name=successful_jobs_history_limit,json=successfulJobsHistoryLimit,proto3" json:"successful_jobs_history_limit,omitempty"`
	Labels                     map[string]string           `protobuf:"bytes,16,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Annotations                map[string]string           `protobuf:"bytes,17,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	JobTemplate                *JobServiceAttribute        `protobuf:"bytes,18,opt,name=job_template,json=jobTemplate,proto3" json:"job_template,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}                    `json:"-"`
	XXX_unrecognized           []byte                      `json:"-"`
	XXX_sizecache              int32                       `json:"-"`
}

func (m *CronJobServiceAttribute) Reset()         { *m = CronJobServiceAttribute{} }
func (m *CronJobServiceAttribute) String() string { return proto.CompactTextString(m) }
func (*CronJobServiceAttribute) ProtoMessage()    {}
func (*CronJobServiceAttribute) Descriptor() ([]byte, []int) {
	return fileDescriptor_9693a1980b1a682f, []int{1}
}

func (m *CronJobServiceAttribute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CronJobServiceAttribute.Unmarshal(m, b)
}
func (m *CronJobServiceAttribute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CronJobServiceAttribute.Marshal(b, m, deterministic)
}
func (m *CronJobServiceAttribute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CronJobServiceAttribute.Merge(m, src)
}
func (m *CronJobServiceAttribute) XXX_Size() int {
	return xxx_messageInfo_CronJobServiceAttribute.Size(m)
}
func (m *CronJobServiceAttribute) XXX_DiscardUnknown() {
	xxx_messageInfo_CronJobServiceAttribute.DiscardUnknown(m)
}

var xxx_messageInfo_CronJobServiceAttribute proto.InternalMessageInfo

func (m *CronJobServiceAttribute) GetSchedule() string {
	if m != nil {
		return m.Schedule
	}
	return ""
}

func (m *CronJobServiceAttribute) GetStartingDeadlineSeconds() *StartingDeadlineSeconds {
	if m != nil {
		return m.StartingDeadlineSeconds
	}
	return nil
}

func (m *CronJobServiceAttribute) GetConcurrencyPolicy() ConcurrencyPolicy {
	if m != nil {
		return m.ConcurrencyPolicy
	}
	return ConcurrencyPolicy_Allow
}

func (m *CronJobServiceAttribute) GetSuspend() *Suspend {
	if m != nil {
		return m.Suspend
	}
	return nil
}

func (m *CronJobServiceAttribute) GetFailedJobsHistoryLimit() *FailedJobsHistoryLimit {
	if m != nil {
		return m.FailedJobsHistoryLimit
	}
	return nil
}

func (m *CronJobServiceAttribute) GetSuccessfulJobsHistoryLimit() *SuccessfulJobsHistoryLimit {
	if m != nil {
		return m.SuccessfulJobsHistoryLimit
	}
	return nil
}

func (m *CronJobServiceAttribute) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *CronJobServiceAttribute) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func (m *CronJobServiceAttribute) GetJobTemplate() *JobServiceAttribute {
	if m != nil {
		return m.JobTemplate
	}
	return nil
}

type StartingDeadlineSeconds struct {
	Value                int64    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartingDeadlineSeconds) Reset()         { *m = StartingDeadlineSeconds{} }
func (m *StartingDeadlineSeconds) String() string { return proto.CompactTextString(m) }
func (*StartingDeadlineSeconds) ProtoMessage()    {}
func (*StartingDeadlineSeconds) Descriptor() ([]byte, []int) {
	return fileDescriptor_9693a1980b1a682f, []int{2}
}

func (m *StartingDeadlineSeconds) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartingDeadlineSeconds.Unmarshal(m, b)
}
func (m *StartingDeadlineSeconds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartingDeadlineSeconds.Marshal(b, m, deterministic)
}
func (m *StartingDeadlineSeconds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartingDeadlineSeconds.Merge(m, src)
}
func (m *StartingDeadlineSeconds) XXX_Size() int {
	return xxx_messageInfo_StartingDeadlineSeconds.Size(m)
}
func (m *StartingDeadlineSeconds) XXX_DiscardUnknown() {
	xxx_messageInfo_StartingDeadlineSeconds.DiscardUnknown(m)
}

var xxx_messageInfo_StartingDeadlineSeconds proto.InternalMessageInfo

func (m *StartingDeadlineSeconds) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Suspend struct {
	Value                bool     `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Suspend) Reset()         { *m = Suspend{} }
func (m *Suspend) String() string { return proto.CompactTextString(m) }
func (*Suspend) ProtoMessage()    {}
func (*Suspend) Descriptor() ([]byte, []int) {
	return fileDescriptor_9693a1980b1a682f, []int{3}
}

func (m *Suspend) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Suspend.Unmarshal(m, b)
}
func (m *Suspend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Suspend.Marshal(b, m, deterministic)
}
func (m *Suspend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Suspend.Merge(m, src)
}
func (m *Suspend) XXX_Size() int {
	return xxx_messageInfo_Suspend.Size(m)
}
func (m *Suspend) XXX_DiscardUnknown() {
	xxx_messageInfo_Suspend.DiscardUnknown(m)
}

var xxx_messageInfo_Suspend proto.InternalMessageInfo

func (m *Suspend) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

type SuccessfulJobsHistoryLimit struct {
	Value                int32    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SuccessfulJobsHistoryLimit) Reset()         { *m = SuccessfulJobsHistoryLimit{} }
func (m *SuccessfulJobsHistoryLimit) String() string { return proto.CompactTextString(m) }
func (*SuccessfulJobsHistoryLimit) ProtoMessage()    {}
func (*SuccessfulJobsHistoryLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_9693a1980b1a682f, []int{4}
}

func (m *SuccessfulJobsHistoryLimit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SuccessfulJobsHistoryLimit.Unmarshal(m, b)
}
func (m *SuccessfulJobsHistoryLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SuccessfulJobsHistoryLimit.Marshal(b, m, deterministic)
}
func (m *SuccessfulJobsHistoryLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SuccessfulJobsHistoryLimit.Merge(m, src)
}
func (m *SuccessfulJobsHistoryLimit) XXX_Size() int {
	return xxx_messageInfo_SuccessfulJobsHistoryLimit.Size(m)
}
func (m *SuccessfulJobsHistoryLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_SuccessfulJobsHistoryLimit.DiscardUnknown(m)
}

var xxx_messageInfo_SuccessfulJobsHistoryLimit proto.InternalMessageInfo

func (m *SuccessfulJobsHistoryLimit) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type FailedJobsHistoryLimit struct {
	Value                int32    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FailedJobsHistoryLimit) Reset()         { *m = FailedJobsHistoryLimit{} }
func (m *FailedJobsHistoryLimit) String() string { return proto.CompactTextString(m) }
func (*FailedJobsHistoryLimit) ProtoMessage()    {}
func (*FailedJobsHistoryLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_9693a1980b1a682f, []int{5}
}

func (m *FailedJobsHistoryLimit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FailedJobsHistoryLimit.Unmarshal(m, b)
}
func (m *FailedJobsHistoryLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FailedJobsHistoryLimit.Marshal(b, m, deterministic)
}
func (m *FailedJobsHistoryLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FailedJobsHistoryLimit.Merge(m, src)
}
func (m *FailedJobsHistoryLimit) XXX_Size() int {
	return xxx_messageInfo_FailedJobsHistoryLimit.Size(m)
}
func (m *FailedJobsHistoryLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_FailedJobsHistoryLimit.DiscardUnknown(m)
}

var xxx_messageInfo_FailedJobsHistoryLimit proto.InternalMessageInfo

func (m *FailedJobsHistoryLimit) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterEnum("proto.ConcurrencyPolicy", ConcurrencyPolicy_name, ConcurrencyPolicy_value)
	proto.RegisterType((*CronJobService)(nil), "proto.CronJobService")
	proto.RegisterType((*CronJobServiceAttribute)(nil), "proto.CronJobServiceAttribute")
	proto.RegisterMapType((map[string]string)(nil), "proto.CronJobServiceAttribute.AnnotationsEntry")
	proto.RegisterMapType((map[string]string)(nil), "proto.CronJobServiceAttribute.LabelsEntry")
	proto.RegisterType((*StartingDeadlineSeconds)(nil), "proto.StartingDeadlineSeconds")
	proto.RegisterType((*Suspend)(nil), "proto.Suspend")
	proto.RegisterType((*SuccessfulJobsHistoryLimit)(nil), "proto.SuccessfulJobsHistoryLimit")
	proto.RegisterType((*FailedJobsHistoryLimit)(nil), "proto.FailedJobsHistoryLimit")
}

func init() {
	proto.RegisterFile("cronjob.proto", fileDescriptor_9693a1980b1a682f)
}

var fileDescriptor_9693a1980b1a682f = []byte{
	// 697 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xcd, 0x4e, 0xdb, 0x4a,
	0x14, 0x26, 0x09, 0x49, 0xc8, 0x31, 0x44, 0x61, 0x74, 0x2f, 0x18, 0xeb, 0x72, 0x2f, 0x64, 0x15,
	0xb1, 0x08, 0x52, 0xee, 0xa2, 0x3f, 0x12, 0x6d, 0x29, 0x14, 0x0a, 0x62, 0x41, 0x1d, 0x16, 0xa8,
	0x1b, 0xcb, 0x1e, 0x1f, 0x8a, 0x83, 0x33, 0x63, 0x79, 0xc6, 0x54, 0x7e, 0x90, 0xbe, 0x55, 0x9f,
	0xa7, 0xeb, 0xca, 0x33, 0xe3, 0x84, 0x9f, 0xa4, 0x55, 0x61, 0x65, 0xcf, 0xf7, 0x37, 0x67, 0xe6,
	0x9c, 0x81, 0x15, 0x9a, 0x72, 0x36, 0xe2, 0x41, 0x3f, 0x49, 0xb9, 0xe4, 0xa4, 0xae, 0x3e, 0xce,
	0x8a, 0xc0, 0xf4, 0x36, 0xa2, 0xa8, 0x51, 0xa7, 0x35, 0x11, 0x74, 0x7f, 0x54, 0xa1, 0x7d, 0x90,
	0x72, 0x76, 0xca, 0x83, 0xa1, 0xd6, 0x90, 0x4d, 0x80, 0x24, 0xe5, 0x23, 0xa4, 0xd2, 0x8b, 0x42,
	0xbb, 0xb2, 0x55, 0xe9, 0xb5, 0xdc, 0x96, 0x41, 0x4e, 0xc2, 0x82, 0x36, 0x69, 0x05, 0x5d, 0xd5,
	0xb4, 0x41, 0x4e, 0x42, 0x42, 0x60, 0x91, 0xf9, 0x63, 0xb4, 0x6b, 0x8a, 0x50, 0xff, 0xc4, 0x86,
	0xe6, 0x2d, 0xa6, 0x22, 0xe2, 0xcc, 0x5e, 0x54, 0x70, 0xb9, 0x24, 0xdb, 0xb0, 0x5c, 0x86, 0xc9,
	0x3c, 0x41, 0xbb, 0xae, 0x68, 0xcb, 0x60, 0x17, 0x79, 0x82, 0xa4, 0x07, 0x9d, 0x52, 0x22, 0xb2,
	0x40, 0xcb, 0x1a, 0x4a, 0xd6, 0x36, 0xf8, 0x30, 0x0b, 0x94, 0xf2, 0x1f, 0x68, 0x15, 0xdb, 0x89,
	0xc4, 0xa7, 0x68, 0x37, 0x75, 0x61, 0x13, 0x80, 0xfc, 0x05, 0x75, 0xc9, 0x6f, 0x90, 0xd9, 0x4b,
	0x8a, 0xd1, 0x8b, 0xe2, 0x34, 0x94, 0x8f, 0x13, 0x9f, 0xe5, 0xc5, 0x69, 0x5a, 0xda, 0x64, 0x90,
	0x93, 0x90, 0x5c, 0xc2, 0xfa, 0xfd, 0xdb, 0xd9, 0x97, 0x32, 0x8d, 0x82, 0x4c, 0xa2, 0x0d, 0x5b,
	0x95, 0x9e, 0x35, 0xf8, 0x57, 0xdf, 0x63, 0x7f, 0x8e, 0xca, 0x9d, 0x67, 0xef, 0x7e, 0x6b, 0xcc,
	0x8d, 0x26, 0x0e, 0x2c, 0x09, 0x7a, 0x8d, 0x61, 0x16, 0xa3, 0xb9, 0xff, 0xc9, 0x9a, 0x7c, 0x86,
	0x0d, 0x21, 0xfd, 0x54, 0x46, 0xec, 0x8b, 0x17, 0xa2, 0x1f, 0xc6, 0x11, 0x43, 0x4f, 0x20, 0xe5,
	0x2c, 0x14, 0xaa, 0x1b, 0xd3, 0x9a, 0x86, 0x46, 0x77, 0x68, 0x64, 0x43, 0xad, 0x72, 0xd7, 0xc5,
	0x6c, 0x82, 0x1c, 0x03, 0xa1, 0x9c, 0xd1, 0x2c, 0x4d, 0x91, 0xd1, 0xdc, 0x4b, 0x78, 0x1c, 0xd1,
	0x5c, 0x75, 0xb2, 0x3d, 0xb0, 0xcb, 0x83, 0x4e, 0x05, 0xe7, 0x8a, 0x77, 0x57, 0xe9, 0x43, 0x88,
	0xf4, 0xa0, 0x29, 0x32, 0x91, 0x20, 0x0b, 0x55, 0xc3, 0xad, 0x41, 0xbb, 0x2c, 0x49, 0xa3, 0x6e,
	0x49, 0x93, 0x4b, 0xd8, 0xb8, 0xf2, 0xa3, 0x18, 0x43, 0x6f, 0xc4, 0x03, 0xe1, 0x5d, 0x47, 0x42,
	0xf2, 0x34, 0xf7, 0xe2, 0x68, 0x1c, 0x49, 0x35, 0x0d, 0xd6, 0x60, 0xd3, 0x78, 0x8f, 0x94, 0xee,
	0x94, 0x07, 0xe2, 0xa3, 0x56, 0x9d, 0x15, 0x22, 0x77, 0xed, 0x6a, 0x26, 0x4e, 0x42, 0xd8, 0x14,
	0x19, 0xa5, 0x28, 0xc4, 0x55, 0x16, 0xcf, 0x4a, 0x6f, 0xa8, 0xf4, 0xed, 0x49, 0x65, 0xa5, 0xf6,
	0xd1, 0x0e, 0x8e, 0x98, 0xcb, 0x91, 0xf7, 0xd0, 0x88, 0xfd, 0x00, 0x63, 0x61, 0x77, 0xb6, 0x6a,
	0x3d, 0x6b, 0xb0, 0xf3, 0xeb, 0x79, 0xe8, 0x9f, 0x29, 0xf1, 0x07, 0x26, 0xd3, 0xdc, 0x35, 0x4e,
	0xf2, 0x09, 0x2c, 0x9f, 0x31, 0x2e, 0x7d, 0x19, 0x71, 0x26, 0xec, 0x55, 0x15, 0xb4, 0xfb, 0x9b,
	0xa0, 0xfd, 0xa9, 0x43, 0xa7, 0xdd, 0xcd, 0x20, 0x7b, 0xb0, 0x3c, 0xe2, 0x81, 0x27, 0x71, 0x9c,
	0xc4, 0xbe, 0x44, 0x9b, 0xa8, 0xb3, 0x3a, 0x26, 0x73, 0xd6, 0xa0, 0x5a, 0x23, 0x1e, 0x5c, 0x18,
	0xb9, 0xf3, 0x0a, 0xac, 0x3b, 0x85, 0x92, 0x0e, 0xd4, 0x6e, 0x30, 0x37, 0xa3, 0x58, 0xfc, 0x16,
	0x8f, 0xe9, 0xd6, 0x8f, 0x33, 0x34, 0xef, 0x5f, 0x2f, 0x5e, 0x57, 0x5f, 0x56, 0x9c, 0x37, 0xd0,
	0x79, 0x58, 0xda, 0x9f, 0xf8, 0xbb, 0xbb, 0xb0, 0x3e, 0x67, 0x6e, 0xa7, 0xa6, 0x22, 0xa8, 0x66,
	0x4c, 0xdd, 0xff, 0xa0, 0x69, 0xa6, 0xea, 0xbe, 0x60, 0xa9, 0x14, 0x0c, 0xc0, 0x99, 0xdf, 0xdc,
	0xfb, 0x9e, 0x7a, 0xe9, 0xe9, 0xc3, 0xda, 0xec, 0x71, 0x9b, 0xad, 0xdf, 0x79, 0x01, 0xab, 0x8f,
	0x1e, 0x06, 0x69, 0x41, 0x7d, 0x3f, 0x8e, 0xf9, 0xd7, 0xce, 0x02, 0x01, 0x68, 0x1c, 0xf1, 0x34,
	0x88, 0xc2, 0x4e, 0x85, 0x58, 0xd0, 0x74, 0x31, 0x89, 0x7d, 0x8a, 0x9d, 0xea, 0xe0, 0x7b, 0x15,
	0x9a, 0xa6, 0xc5, 0xe4, 0x1d, 0xac, 0x1c, 0xa4, 0xe8, 0x4b, 0x2c, 0x81, 0xbf, 0x67, 0xce, 0x80,
	0xb3, 0x56, 0x8e, 0xac, 0x5e, 0xbb, 0x28, 0x12, 0xce, 0x04, 0x76, 0x17, 0x8a, 0x84, 0x43, 0x8c,
	0xf1, 0x19, 0x09, 0x7b, 0x00, 0xc7, 0x28, 0x9f, 0x6c, 0x7f, 0x0b, 0xcb, 0xe7, 0xbe, 0xa4, 0xd7,
	0xcf, 0xd9, 0xff, 0x3c, 0x7b, 0xf2, 0xfe, 0x41, 0x43, 0x11, 0xff, 0xff, 0x0c, 0x00, 0x00, 0xff,
	0xff, 0xd8, 0x29, 0x4d, 0x93, 0x07, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CronJobClient is the client API for CronJob service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CronJobClient interface {
	CreateCronJob(ctx context.Context, in *CronJobService, opts ...grpc.CallOption) (*ServiceResponse, error)
	DeleteCronJob(ctx context.Context, in *CronJobService, opts ...grpc.CallOption) (*ServiceResponse, error)
	GetCronJob(ctx context.Context, in *CronJobService, opts ...grpc.CallOption) (*ServiceResponse, error)
	PatchCronJob(ctx context.Context, in *CronJobService, opts ...grpc.CallOption) (*ServiceResponse, error)
	PutCronJob(ctx context.Context, in *CronJobService, opts ...grpc.CallOption) (*ServiceResponse, error)
}

type cronJobClient struct {
	cc grpc.ClientConnInterface
}

func NewCronJobClient(cc grpc.ClientConnInterface) CronJobClient {
	return &cronJobClient{cc}
}

func (c *cronJobClient) CreateCronJob(ctx context.Context, in *CronJobService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.CronJob/CreateCronJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cronJobClient) DeleteCronJob(ctx context.Context, in *CronJobService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.CronJob/DeleteCronJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cronJobClient) GetCronJob(ctx context.Context, in *CronJobService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.CronJob/GetCronJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cronJobClient) PatchCronJob(ctx context.Context, in *CronJobService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.CronJob/PatchCronJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cronJobClient) PutCronJob(ctx context.Context, in *CronJobService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.CronJob/PutCronJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CronJobServer is the server API for CronJob service.
type CronJobServer interface {
	CreateCronJob(context.Context, *CronJobService) (*ServiceResponse, error)
	DeleteCronJob(context.Context, *CronJobService) (*ServiceResponse, error)
	GetCronJob(context.Context, *CronJobService) (*ServiceResponse, error)
	PatchCronJob(context.Context, *CronJobService) (*ServiceResponse, error)
	PutCronJob(context.Context, *CronJobService) (*ServiceResponse, error)
}

// UnimplementedCronJobServer can be embedded to have forward compatible implementations.
type UnimplementedCronJobServer struct {
}

func (*UnimplementedCronJobServer) CreateCronJob(ctx context.Context, req *CronJobService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCronJob not implemented")
}
func (*UnimplementedCronJobServer) DeleteCronJob(ctx context.Context, req *CronJobService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCronJob not implemented")
}
func (*UnimplementedCronJobServer) GetCronJob(ctx context.Context, req *CronJobService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCronJob not implemented")
}
func (*UnimplementedCronJobServer) PatchCronJob(ctx context.Context, req *CronJobService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchCronJob not implemented")
}
func (*UnimplementedCronJobServer) PutCronJob(ctx context.Context, req *CronJobService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutCronJob not implemented")
}

func RegisterCronJobServer(s *grpc.Server, srv CronJobServer) {
	s.RegisterService(&_CronJob_serviceDesc, srv)
}

func _CronJob_CreateCronJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CronJobService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CronJobServer).CreateCronJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CronJob/CreateCronJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CronJobServer).CreateCronJob(ctx, req.(*CronJobService))
	}
	return interceptor(ctx, in, info, handler)
}

func _CronJob_DeleteCronJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CronJobService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CronJobServer).DeleteCronJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CronJob/DeleteCronJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CronJobServer).DeleteCronJob(ctx, req.(*CronJobService))
	}
	return interceptor(ctx, in, info, handler)
}

func _CronJob_GetCronJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CronJobService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CronJobServer).GetCronJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CronJob/GetCronJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CronJobServer).GetCronJob(ctx, req.(*CronJobService))
	}
	return interceptor(ctx, in, info, handler)
}

func _CronJob_PatchCronJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CronJobService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CronJobServer).PatchCronJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CronJob/PatchCronJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CronJobServer).PatchCronJob(ctx, req.(*CronJobService))
	}
	return interceptor(ctx, in, info, handler)
}

func _CronJob_PutCronJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CronJobService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CronJobServer).PutCronJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CronJob/PutCronJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CronJobServer).PutCronJob(ctx, req.(*CronJobService))
	}
	return interceptor(ctx, in, info, handler)
}

var _CronJob_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.CronJob",
	HandlerType: (*CronJobServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCronJob",
			Handler:    _CronJob_CreateCronJob_Handler,
		},
		{
			MethodName: "DeleteCronJob",
			Handler:    _CronJob_DeleteCronJob_Handler,
		},
		{
			MethodName: "GetCronJob",
			Handler:    _CronJob_GetCronJob_Handler,
		},
		{
			MethodName: "PatchCronJob",
			Handler:    _CronJob_PatchCronJob_Handler,
		},
		{
			MethodName: "PutCronJob",
			Handler:    _CronJob_PutCronJob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cronjob.proto",
}
