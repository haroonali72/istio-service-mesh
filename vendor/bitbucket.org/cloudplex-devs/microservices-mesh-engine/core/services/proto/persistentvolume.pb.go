// Code generated by protoc-gen-go. DO NOT EDIT.
// source: persistentvolume.proto

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

type AccessMode int32

const (
	AccessMode_ReadWriteOnce AccessMode = 0
	AccessMode_ReadOnlyMany  AccessMode = 1
	AccessMode_ReadWriteMany AccessMode = 2
)

var AccessMode_name = map[int32]string{
	0: "ReadWriteOnce",
	1: "ReadOnlyMany",
	2: "ReadWriteMany",
}

var AccessMode_value = map[string]int32{
	"ReadWriteOnce": 0,
	"ReadOnlyMany":  1,
	"ReadWriteMany": 2,
}

func (x AccessMode) String() string {
	return proto.EnumName(AccessMode_name, int32(x))
}

func (AccessMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_724324990baa6ac6, []int{0}
}

type PersistentVolumeService struct {
	ApplicationId        string                      `protobuf:"bytes,1,opt,name=application_id,json=applicationId,proto3" json:"application_id,omitempty"`
	ServiceId            string                      `protobuf:"bytes,2,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	Name                 string                      `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Version              string                      `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	ServiceType          string                      `protobuf:"bytes,5,opt,name=service_type,json=serviceType,proto3" json:"service_type,omitempty"`
	ServiceSubType       string                      `protobuf:"bytes,6,opt,name=service_sub_type,json=serviceSubType,proto3" json:"service_sub_type,omitempty"`
	Token                string                      `protobuf:"bytes,7,opt,name=token,proto3" json:"token,omitempty"`
	CompanyId            string                      `protobuf:"bytes,8,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	ServiceAttributes    *PersistentVolumeAttributes `protobuf:"bytes,9,opt,name=service_attributes,json=serviceAttributes,proto3" json:"service_attributes,omitempty"`
	InfraId              string                      `protobuf:"bytes,11,opt,name=infra_id,json=infraId,proto3" json:"infra_id,omitempty"`
	HookConfiguration    *HookConfiguration          `protobuf:"bytes,12,opt,name=hook_configuration,json=hookConfiguration,proto3" json:"hook_configuration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *PersistentVolumeService) Reset()         { *m = PersistentVolumeService{} }
func (m *PersistentVolumeService) String() string { return proto.CompactTextString(m) }
func (*PersistentVolumeService) ProtoMessage()    {}
func (*PersistentVolumeService) Descriptor() ([]byte, []int) {
	return fileDescriptor_724324990baa6ac6, []int{0}
}

func (m *PersistentVolumeService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PersistentVolumeService.Unmarshal(m, b)
}
func (m *PersistentVolumeService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PersistentVolumeService.Marshal(b, m, deterministic)
}
func (m *PersistentVolumeService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PersistentVolumeService.Merge(m, src)
}
func (m *PersistentVolumeService) XXX_Size() int {
	return xxx_messageInfo_PersistentVolumeService.Size(m)
}
func (m *PersistentVolumeService) XXX_DiscardUnknown() {
	xxx_messageInfo_PersistentVolumeService.DiscardUnknown(m)
}

var xxx_messageInfo_PersistentVolumeService proto.InternalMessageInfo

func (m *PersistentVolumeService) GetApplicationId() string {
	if m != nil {
		return m.ApplicationId
	}
	return ""
}

func (m *PersistentVolumeService) GetServiceId() string {
	if m != nil {
		return m.ServiceId
	}
	return ""
}

func (m *PersistentVolumeService) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PersistentVolumeService) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *PersistentVolumeService) GetServiceType() string {
	if m != nil {
		return m.ServiceType
	}
	return ""
}

func (m *PersistentVolumeService) GetServiceSubType() string {
	if m != nil {
		return m.ServiceSubType
	}
	return ""
}

func (m *PersistentVolumeService) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *PersistentVolumeService) GetCompanyId() string {
	if m != nil {
		return m.CompanyId
	}
	return ""
}

func (m *PersistentVolumeService) GetServiceAttributes() *PersistentVolumeAttributes {
	if m != nil {
		return m.ServiceAttributes
	}
	return nil
}

func (m *PersistentVolumeService) GetInfraId() string {
	if m != nil {
		return m.InfraId
	}
	return ""
}

func (m *PersistentVolumeService) GetHookConfiguration() *HookConfiguration {
	if m != nil {
		return m.HookConfiguration
	}
	return nil
}

type PersistentVolumeAttributes struct {
	Labels                 map[string]string       `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ReclaimPolicy          ReclaimPolicy           `protobuf:"varint,2,opt,name=reclaim_policy,json=reclaimPolicy,proto3,enum=proto.ReclaimPolicy" json:"reclaim_policy,omitempty"`
	AccessMode             []AccessMode            `protobuf:"varint,3,rep,packed,name=access_mode,json=accessMode,proto3,enum=proto.AccessMode" json:"access_mode,omitempty"`
	Capacity               string                  `protobuf:"bytes,4,opt,name=capacity,proto3" json:"capacity,omitempty"`
	PersistentVolumeSource *PersistentVolumeSource `protobuf:"bytes,5,opt,name=persistent_volume_source,json=persistentVolumeSource,proto3" json:"persistent_volume_source,omitempty"`
	StorageClassName       string                  `protobuf:"bytes,6,opt,name=storage_class_name,json=storageClassName,proto3" json:"storage_class_name,omitempty"`
	MountOptions           []string                `protobuf:"bytes,7,rep,name=mount_options,json=mountOptions,proto3" json:"mount_options,omitempty"`
	VolumeMode             string                  `protobuf:"bytes,8,opt,name=volume_mode,json=volumeMode,proto3" json:"volume_mode,omitempty"`
	NodeAffinity           *VolumeNodeAffinity     `protobuf:"bytes,9,opt,name=node_affinity,json=nodeAffinity,proto3" json:"node_affinity,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                `json:"-"`
	XXX_unrecognized       []byte                  `json:"-"`
	XXX_sizecache          int32                   `json:"-"`
}

func (m *PersistentVolumeAttributes) Reset()         { *m = PersistentVolumeAttributes{} }
func (m *PersistentVolumeAttributes) String() string { return proto.CompactTextString(m) }
func (*PersistentVolumeAttributes) ProtoMessage()    {}
func (*PersistentVolumeAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_724324990baa6ac6, []int{1}
}

func (m *PersistentVolumeAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PersistentVolumeAttributes.Unmarshal(m, b)
}
func (m *PersistentVolumeAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PersistentVolumeAttributes.Marshal(b, m, deterministic)
}
func (m *PersistentVolumeAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PersistentVolumeAttributes.Merge(m, src)
}
func (m *PersistentVolumeAttributes) XXX_Size() int {
	return xxx_messageInfo_PersistentVolumeAttributes.Size(m)
}
func (m *PersistentVolumeAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_PersistentVolumeAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_PersistentVolumeAttributes proto.InternalMessageInfo

func (m *PersistentVolumeAttributes) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *PersistentVolumeAttributes) GetReclaimPolicy() ReclaimPolicy {
	if m != nil {
		return m.ReclaimPolicy
	}
	return ReclaimPolicy_Retain
}

func (m *PersistentVolumeAttributes) GetAccessMode() []AccessMode {
	if m != nil {
		return m.AccessMode
	}
	return nil
}

func (m *PersistentVolumeAttributes) GetCapacity() string {
	if m != nil {
		return m.Capacity
	}
	return ""
}

func (m *PersistentVolumeAttributes) GetPersistentVolumeSource() *PersistentVolumeSource {
	if m != nil {
		return m.PersistentVolumeSource
	}
	return nil
}

func (m *PersistentVolumeAttributes) GetStorageClassName() string {
	if m != nil {
		return m.StorageClassName
	}
	return ""
}

func (m *PersistentVolumeAttributes) GetMountOptions() []string {
	if m != nil {
		return m.MountOptions
	}
	return nil
}

func (m *PersistentVolumeAttributes) GetVolumeMode() string {
	if m != nil {
		return m.VolumeMode
	}
	return ""
}

func (m *PersistentVolumeAttributes) GetNodeAffinity() *VolumeNodeAffinity {
	if m != nil {
		return m.NodeAffinity
	}
	return nil
}

type PersistentVolumeSource struct {
	GcpPd                *GCPPD     `protobuf:"bytes,1,opt,name=gcp_pd,json=gcpPd,proto3" json:"gcp_pd,omitempty"`
	AwsEbs               *AWSEBS    `protobuf:"bytes,2,opt,name=aws_ebs,json=awsEbs,proto3" json:"aws_ebs,omitempty"`
	AzureDisk            *AzureDisk `protobuf:"bytes,3,opt,name=azure_disk,json=azureDisk,proto3" json:"azure_disk,omitempty"`
	AzureFile            *AzureFile `protobuf:"bytes,4,opt,name=azure_file,json=azureFile,proto3" json:"azure_file,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *PersistentVolumeSource) Reset()         { *m = PersistentVolumeSource{} }
func (m *PersistentVolumeSource) String() string { return proto.CompactTextString(m) }
func (*PersistentVolumeSource) ProtoMessage()    {}
func (*PersistentVolumeSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_724324990baa6ac6, []int{2}
}

func (m *PersistentVolumeSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PersistentVolumeSource.Unmarshal(m, b)
}
func (m *PersistentVolumeSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PersistentVolumeSource.Marshal(b, m, deterministic)
}
func (m *PersistentVolumeSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PersistentVolumeSource.Merge(m, src)
}
func (m *PersistentVolumeSource) XXX_Size() int {
	return xxx_messageInfo_PersistentVolumeSource.Size(m)
}
func (m *PersistentVolumeSource) XXX_DiscardUnknown() {
	xxx_messageInfo_PersistentVolumeSource.DiscardUnknown(m)
}

var xxx_messageInfo_PersistentVolumeSource proto.InternalMessageInfo

func (m *PersistentVolumeSource) GetGcpPd() *GCPPD {
	if m != nil {
		return m.GcpPd
	}
	return nil
}

func (m *PersistentVolumeSource) GetAwsEbs() *AWSEBS {
	if m != nil {
		return m.AwsEbs
	}
	return nil
}

func (m *PersistentVolumeSource) GetAzureDisk() *AzureDisk {
	if m != nil {
		return m.AzureDisk
	}
	return nil
}

func (m *PersistentVolumeSource) GetAzureFile() *AzureFile {
	if m != nil {
		return m.AzureFile
	}
	return nil
}

type GCPPD struct {
	PdName               string   `protobuf:"bytes,1,opt,name=pd_name,json=pdName,proto3" json:"pd_name,omitempty"`
	FileSystem           string   `protobuf:"bytes,2,opt,name=file_system,json=fileSystem,proto3" json:"file_system,omitempty"`
	Partation            int64    `protobuf:"varint,3,opt,name=partation,proto3" json:"partation,omitempty"`
	Readonly             bool     `protobuf:"varint,4,opt,name=readonly,proto3" json:"readonly,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GCPPD) Reset()         { *m = GCPPD{} }
func (m *GCPPD) String() string { return proto.CompactTextString(m) }
func (*GCPPD) ProtoMessage()    {}
func (*GCPPD) Descriptor() ([]byte, []int) {
	return fileDescriptor_724324990baa6ac6, []int{3}
}

func (m *GCPPD) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GCPPD.Unmarshal(m, b)
}
func (m *GCPPD) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GCPPD.Marshal(b, m, deterministic)
}
func (m *GCPPD) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GCPPD.Merge(m, src)
}
func (m *GCPPD) XXX_Size() int {
	return xxx_messageInfo_GCPPD.Size(m)
}
func (m *GCPPD) XXX_DiscardUnknown() {
	xxx_messageInfo_GCPPD.DiscardUnknown(m)
}

var xxx_messageInfo_GCPPD proto.InternalMessageInfo

func (m *GCPPD) GetPdName() string {
	if m != nil {
		return m.PdName
	}
	return ""
}

func (m *GCPPD) GetFileSystem() string {
	if m != nil {
		return m.FileSystem
	}
	return ""
}

func (m *GCPPD) GetPartation() int64 {
	if m != nil {
		return m.Partation
	}
	return 0
}

func (m *GCPPD) GetReadonly() bool {
	if m != nil {
		return m.Readonly
	}
	return false
}

type AWSEBS struct {
	VolumeId             string   `protobuf:"bytes,1,opt,name=volume_id,json=volumeId,proto3" json:"volume_id,omitempty"`
	FileSystem           string   `protobuf:"bytes,2,opt,name=file_system,json=fileSystem,proto3" json:"file_system,omitempty"`
	Partation            int64    `protobuf:"varint,3,opt,name=partation,proto3" json:"partation,omitempty"`
	Readonly             bool     `protobuf:"varint,4,opt,name=readonly,proto3" json:"readonly,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AWSEBS) Reset()         { *m = AWSEBS{} }
func (m *AWSEBS) String() string { return proto.CompactTextString(m) }
func (*AWSEBS) ProtoMessage()    {}
func (*AWSEBS) Descriptor() ([]byte, []int) {
	return fileDescriptor_724324990baa6ac6, []int{4}
}

func (m *AWSEBS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AWSEBS.Unmarshal(m, b)
}
func (m *AWSEBS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AWSEBS.Marshal(b, m, deterministic)
}
func (m *AWSEBS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AWSEBS.Merge(m, src)
}
func (m *AWSEBS) XXX_Size() int {
	return xxx_messageInfo_AWSEBS.Size(m)
}
func (m *AWSEBS) XXX_DiscardUnknown() {
	xxx_messageInfo_AWSEBS.DiscardUnknown(m)
}

var xxx_messageInfo_AWSEBS proto.InternalMessageInfo

func (m *AWSEBS) GetVolumeId() string {
	if m != nil {
		return m.VolumeId
	}
	return ""
}

func (m *AWSEBS) GetFileSystem() string {
	if m != nil {
		return m.FileSystem
	}
	return ""
}

func (m *AWSEBS) GetPartation() int64 {
	if m != nil {
		return m.Partation
	}
	return 0
}

func (m *AWSEBS) GetReadonly() bool {
	if m != nil {
		return m.Readonly
	}
	return false
}

type AzureDisk struct {
	CachingMode          AzureDataDiskCachingMode `protobuf:"varint,1,opt,name=cachingMode,proto3,enum=proto.AzureDataDiskCachingMode" json:"cachingMode,omitempty"`
	Kind                 AzureDataDiskKind        `protobuf:"varint,6,opt,name=kind,proto3,enum=proto.AzureDataDiskKind" json:"kind,omitempty"`
	FileSystem           string                   `protobuf:"bytes,2,opt,name=fileSystem,proto3" json:"fileSystem,omitempty"`
	DiskName             string                   `protobuf:"bytes,3,opt,name=diskName,proto3" json:"diskName,omitempty"`
	ReadOnly             bool                     `protobuf:"varint,4,opt,name=readOnly,proto3" json:"readOnly,omitempty"`
	DiskURI              string                   `protobuf:"bytes,5,opt,name=diskURI,proto3" json:"diskURI,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *AzureDisk) Reset()         { *m = AzureDisk{} }
func (m *AzureDisk) String() string { return proto.CompactTextString(m) }
func (*AzureDisk) ProtoMessage()    {}
func (*AzureDisk) Descriptor() ([]byte, []int) {
	return fileDescriptor_724324990baa6ac6, []int{5}
}

func (m *AzureDisk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AzureDisk.Unmarshal(m, b)
}
func (m *AzureDisk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AzureDisk.Marshal(b, m, deterministic)
}
func (m *AzureDisk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AzureDisk.Merge(m, src)
}
func (m *AzureDisk) XXX_Size() int {
	return xxx_messageInfo_AzureDisk.Size(m)
}
func (m *AzureDisk) XXX_DiscardUnknown() {
	xxx_messageInfo_AzureDisk.DiscardUnknown(m)
}

var xxx_messageInfo_AzureDisk proto.InternalMessageInfo

func (m *AzureDisk) GetCachingMode() AzureDataDiskCachingMode {
	if m != nil {
		return m.CachingMode
	}
	return AzureDataDiskCachingMode_ModeNone
}

func (m *AzureDisk) GetKind() AzureDataDiskKind {
	if m != nil {
		return m.Kind
	}
	return AzureDataDiskKind_Shared
}

func (m *AzureDisk) GetFileSystem() string {
	if m != nil {
		return m.FileSystem
	}
	return ""
}

func (m *AzureDisk) GetDiskName() string {
	if m != nil {
		return m.DiskName
	}
	return ""
}

func (m *AzureDisk) GetReadOnly() bool {
	if m != nil {
		return m.ReadOnly
	}
	return false
}

func (m *AzureDisk) GetDiskURI() string {
	if m != nil {
		return m.DiskURI
	}
	return ""
}

type AzureFile struct {
	SecretName           string   `protobuf:"bytes,1,opt,name=secretName,proto3" json:"secretName,omitempty"`
	ShareName            string   `protobuf:"bytes,2,opt,name=shareName,proto3" json:"shareName,omitempty"`
	ReadOnly             bool     `protobuf:"varint,3,opt,name=readOnly,proto3" json:"readOnly,omitempty"`
	SecretNamespace      string   `protobuf:"bytes,4,opt,name=secretNamespace,proto3" json:"secretNamespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AzureFile) Reset()         { *m = AzureFile{} }
func (m *AzureFile) String() string { return proto.CompactTextString(m) }
func (*AzureFile) ProtoMessage()    {}
func (*AzureFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_724324990baa6ac6, []int{6}
}

func (m *AzureFile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AzureFile.Unmarshal(m, b)
}
func (m *AzureFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AzureFile.Marshal(b, m, deterministic)
}
func (m *AzureFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AzureFile.Merge(m, src)
}
func (m *AzureFile) XXX_Size() int {
	return xxx_messageInfo_AzureFile.Size(m)
}
func (m *AzureFile) XXX_DiscardUnknown() {
	xxx_messageInfo_AzureFile.DiscardUnknown(m)
}

var xxx_messageInfo_AzureFile proto.InternalMessageInfo

func (m *AzureFile) GetSecretName() string {
	if m != nil {
		return m.SecretName
	}
	return ""
}

func (m *AzureFile) GetShareName() string {
	if m != nil {
		return m.ShareName
	}
	return ""
}

func (m *AzureFile) GetReadOnly() bool {
	if m != nil {
		return m.ReadOnly
	}
	return false
}

func (m *AzureFile) GetSecretNamespace() string {
	if m != nil {
		return m.SecretNamespace
	}
	return ""
}

//enum AzureDataDiskKind {
//    Shared=0;
//    Dedicated=1;
//    Managed=2;
//}
type VolumeNodeAffinity struct {
	Required             *NodeSelector `protobuf:"bytes,1,opt,name=required,proto3" json:"required,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *VolumeNodeAffinity) Reset()         { *m = VolumeNodeAffinity{} }
func (m *VolumeNodeAffinity) String() string { return proto.CompactTextString(m) }
func (*VolumeNodeAffinity) ProtoMessage()    {}
func (*VolumeNodeAffinity) Descriptor() ([]byte, []int) {
	return fileDescriptor_724324990baa6ac6, []int{7}
}

func (m *VolumeNodeAffinity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VolumeNodeAffinity.Unmarshal(m, b)
}
func (m *VolumeNodeAffinity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VolumeNodeAffinity.Marshal(b, m, deterministic)
}
func (m *VolumeNodeAffinity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VolumeNodeAffinity.Merge(m, src)
}
func (m *VolumeNodeAffinity) XXX_Size() int {
	return xxx_messageInfo_VolumeNodeAffinity.Size(m)
}
func (m *VolumeNodeAffinity) XXX_DiscardUnknown() {
	xxx_messageInfo_VolumeNodeAffinity.DiscardUnknown(m)
}

var xxx_messageInfo_VolumeNodeAffinity proto.InternalMessageInfo

func (m *VolumeNodeAffinity) GetRequired() *NodeSelector {
	if m != nil {
		return m.Required
	}
	return nil
}

func init() {
	proto.RegisterEnum("proto.AccessMode", AccessMode_name, AccessMode_value)
	proto.RegisterType((*PersistentVolumeService)(nil), "proto.PersistentVolumeService")
	proto.RegisterType((*PersistentVolumeAttributes)(nil), "proto.PersistentVolumeAttributes")
	proto.RegisterMapType((map[string]string)(nil), "proto.PersistentVolumeAttributes.LabelsEntry")
	proto.RegisterType((*PersistentVolumeSource)(nil), "proto.PersistentVolumeSource")
	proto.RegisterType((*GCPPD)(nil), "proto.GCPPD")
	proto.RegisterType((*AWSEBS)(nil), "proto.AWSEBS")
	proto.RegisterType((*AzureDisk)(nil), "proto.AzureDisk")
	proto.RegisterType((*AzureFile)(nil), "proto.AzureFile")
	proto.RegisterType((*VolumeNodeAffinity)(nil), "proto.VolumeNodeAffinity")
}

func init() {
	proto.RegisterFile("persistentvolume.proto", fileDescriptor_724324990baa6ac6)
}

var fileDescriptor_724324990baa6ac6 = []byte{
	// 1024 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x55, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xde, 0x34, 0x4d, 0x9a, 0x1c, 0x27, 0x21, 0x9d, 0x5d, 0xba, 0xde, 0xc0, 0xee, 0xb2, 0x41,
	0xa0, 0x0a, 0x41, 0x91, 0xc2, 0x0d, 0x3f, 0x12, 0x52, 0x49, 0x43, 0x59, 0xc1, 0xb6, 0x61, 0x02,
	0xf4, 0xd2, 0x9a, 0x38, 0x93, 0xc4, 0x8a, 0x63, 0x1b, 0xcf, 0xb8, 0x28, 0x70, 0xc3, 0x1b, 0xf0,
	0x28, 0xdc, 0xf2, 0x02, 0x88, 0xb7, 0xe1, 0x19, 0x38, 0xf3, 0x63, 0xc7, 0xcd, 0x6e, 0xc5, 0x4d,
	0xf7, 0xca, 0x3e, 0xdf, 0xf7, 0xcd, 0xcc, 0xf9, 0x9d, 0x81, 0xa3, 0x84, 0xa7, 0x22, 0x10, 0x92,
	0x47, 0xf2, 0x3a, 0x0e, 0xb3, 0x35, 0x3f, 0x49, 0xd2, 0x58, 0xc6, 0xa4, 0xa6, 0x3f, 0x3d, 0x22,
	0x64, 0x9c, 0xb2, 0x05, 0xf7, 0x43, 0x26, 0x84, 0xa1, 0x7a, 0x6d, 0xc1, 0xd3, 0xeb, 0xc0, 0xe7,
	0xb9, 0xb9, 0xe0, 0x11, 0x4f, 0x59, 0x68, 0xcd, 0x56, 0x79, 0x9b, 0xfe, 0xdf, 0x55, 0x78, 0x38,
	0x2e, 0x4e, 0xf8, 0x49, 0x53, 0x13, 0xb3, 0x9c, 0xbc, 0x07, 0x1d, 0x96, 0x24, 0x61, 0xe0, 0x33,
	0x19, 0xc4, 0x91, 0x17, 0xcc, 0xdc, 0xca, 0x3b, 0x95, 0xe3, 0x26, 0x6d, 0x97, 0xd0, 0xe7, 0x33,
	0xf2, 0x18, 0xc0, 0x1e, 0xa8, 0x24, 0x7b, 0x5a, 0xd2, 0xb4, 0x08, 0xd2, 0x04, 0xf6, 0x23, 0xb6,
	0xe6, 0x6e, 0x55, 0x13, 0xfa, 0x9f, 0xb8, 0x70, 0x70, 0xad, 0x0e, 0x8d, 0x23, 0x77, 0x5f, 0xc3,
	0xb9, 0x49, 0x9e, 0x41, 0x2b, 0xdf, 0x4c, 0x6e, 0x12, 0xee, 0xd6, 0x34, 0xed, 0x58, 0xec, 0x07,
	0x84, 0xc8, 0x31, 0x74, 0x73, 0x89, 0xc8, 0xa6, 0x46, 0x56, 0xd7, 0xb2, 0x8e, 0xc5, 0x27, 0xd9,
	0x54, 0x2b, 0x1f, 0x40, 0x4d, 0xc6, 0x2b, 0x1e, 0xb9, 0x07, 0x9a, 0x36, 0x86, 0xf2, 0xd7, 0x8f,
	0xd7, 0x09, 0x8b, 0x36, 0xca, 0xdf, 0x86, 0xf1, 0xd7, 0x22, 0xe8, 0xef, 0x18, 0x48, 0xbe, 0x3d,
	0x93, 0x32, 0x0d, 0xa6, 0x99, 0xe4, 0xc2, 0x6d, 0xa2, 0xcc, 0x19, 0x3c, 0x33, 0x59, 0x3b, 0xd9,
	0xcd, 0xd8, 0x69, 0x21, 0xa4, 0x87, 0x76, 0xf1, 0x16, 0x22, 0x8f, 0xa0, 0x11, 0x44, 0xf3, 0x94,
	0xa9, 0xe3, 0x1c, 0x13, 0xae, 0xb6, 0xf1, 0xb0, 0x73, 0x20, 0xcb, 0x38, 0x5e, 0x79, 0x7e, 0x1c,
	0xcd, 0x83, 0x45, 0x96, 0xea, 0x9c, 0xba, 0x2d, 0x7d, 0x98, 0x6b, 0x0f, 0xfb, 0x06, 0x05, 0xc3,
	0x32, 0x4f, 0x0f, 0x97, 0xbb, 0x50, 0xff, 0xaf, 0x7d, 0xe8, 0xdd, 0xee, 0x15, 0x19, 0x41, 0x3d,
	0x64, 0x53, 0x1e, 0x0a, 0x2c, 0x61, 0x15, 0xf7, 0xfe, 0xe8, 0x7f, 0x03, 0x39, 0xf9, 0x4e, 0xeb,
	0x47, 0x91, 0x4c, 0x37, 0xd4, 0x2e, 0x26, 0x5f, 0x40, 0x27, 0x55, 0xad, 0x16, 0xac, 0xbd, 0x24,
	0xc6, 0x16, 0xd8, 0xe8, 0x72, 0x77, 0x06, 0x0f, 0xec, 0x76, 0xd4, 0x90, 0x63, 0xcd, 0xd1, 0x76,
	0x5a, 0x36, 0xc9, 0x00, 0x1c, 0xe6, 0xfb, 0x5c, 0x08, 0x6f, 0x1d, 0xcf, 0x54, 0x3f, 0x54, 0x71,
	0xe5, 0xa1, 0x5d, 0x79, 0xaa, 0x99, 0x17, 0x48, 0x50, 0x60, 0xc5, 0x3f, 0xe9, 0x41, 0xc3, 0x67,
	0x09, 0xf3, 0x03, 0xb9, 0xb1, 0x9d, 0x52, 0xd8, 0xe4, 0x0a, 0xdc, 0xed, 0x6c, 0x78, 0xa6, 0xab,
	0x3d, 0x11, 0x67, 0xa9, 0x6f, 0xda, 0xc6, 0x19, 0x3c, 0xbe, 0x25, 0xca, 0x89, 0x16, 0xd1, 0xd2,
	0x68, 0x95, 0x71, 0xf2, 0x21, 0xe4, 0x53, 0xe5, 0xe9, 0xb1, 0xf2, 0x74, 0xff, 0x9a, 0x16, 0xeb,
	0x5a, 0x66, 0xa8, 0x88, 0x0b, 0xd5, 0xcb, 0xef, 0x42, 0x7b, 0x1d, 0x67, 0xe8, 0x41, 0x9c, 0xa8,
	0x4a, 0x08, 0x6c, 0xb6, 0x2a, 0x0a, 0x5b, 0x1a, 0xbc, 0x34, 0x18, 0x79, 0x0a, 0x8e, 0x75, 0x50,
	0xc7, 0x6e, 0x9a, 0x0e, 0x0c, 0xa4, 0x03, 0xfd, 0x12, 0xda, 0x11, 0x7e, 0x3d, 0x36, 0x9f, 0x07,
	0x91, 0x8a, 0xd6, 0x34, 0xdc, 0x23, 0x1b, 0x81, 0xf1, 0xef, 0x02, 0x15, 0xa7, 0x56, 0x40, 0x5b,
	0x51, 0xc9, 0xea, 0x7d, 0x06, 0x4e, 0xa9, 0x60, 0xa4, 0x0b, 0xd5, 0x15, 0xdf, 0xd8, 0x79, 0x55,
	0xbf, 0x6a, 0x16, 0xae, 0x59, 0x98, 0x71, 0x3b, 0xa0, 0xc6, 0xf8, 0x7c, 0xef, 0xd3, 0x4a, 0xff,
	0x9f, 0x0a, 0x1c, 0xbd, 0x3a, 0x43, 0x18, 0x5b, 0x7d, 0xe1, 0x27, 0x5e, 0x62, 0x26, 0xdf, 0x19,
	0xb4, 0xac, 0x3b, 0xe7, 0xc3, 0xf1, 0xf8, 0x8c, 0xd6, 0x90, 0x1b, 0xcf, 0xc8, 0xfb, 0x70, 0xc0,
	0x7e, 0x11, 0x1e, 0x9f, 0x0a, 0xbd, 0xb7, 0x33, 0x68, 0xe7, 0x35, 0xbd, 0x9a, 0x8c, 0xbe, 0x9a,
	0xd0, 0x3a, 0xb2, 0xa3, 0xa9, 0x20, 0x1f, 0x03, 0xb0, 0x5f, 0xb3, 0x94, 0x7b, 0xb3, 0x40, 0xac,
	0xf4, 0x75, 0xe0, 0x0c, 0xba, 0xb9, 0x54, 0x11, 0x67, 0x88, 0xd3, 0x26, 0xcb, 0x7f, 0xb7, 0x0b,
	0xe6, 0x41, 0xc8, 0x75, 0xf9, 0x77, 0x16, 0x7c, 0x8d, 0xb8, 0x5d, 0xa0, 0x7e, 0xfb, 0xbf, 0x41,
	0x4d, 0x7b, 0x46, 0x1e, 0xc2, 0x41, 0x32, 0x33, 0x65, 0x33, 0x29, 0xa8, 0x27, 0x33, 0x5d, 0x2c,
	0xac, 0x83, 0xda, 0xcc, 0x13, 0x1b, 0x0c, 0x76, 0x6d, 0x73, 0x01, 0x0a, 0x9a, 0x68, 0x84, 0xbc,
	0x0d, 0xcd, 0x84, 0xa5, 0xd2, 0xcc, 0xa1, 0xf2, 0xb1, 0x4a, 0xb7, 0x80, 0x6a, 0xc7, 0x94, 0xb3,
	0x59, 0x1c, 0x85, 0xa6, 0x1d, 0x1b, 0xb4, 0xb0, 0xfb, 0xbf, 0x57, 0xa0, 0x6e, 0x22, 0x26, 0x6f,
	0x41, 0xd3, 0x56, 0xbb, 0xb8, 0x33, 0x1b, 0x06, 0xc0, 0x91, 0x7f, 0x8d, 0x2e, 0xfc, 0x5b, 0x81,
	0x66, 0x91, 0x49, 0x72, 0x0a, 0x8e, 0xcf, 0xfc, 0x65, 0x10, 0x2d, 0x54, 0x87, 0x69, 0x3f, 0x3a,
	0x83, 0xa7, 0x37, 0x12, 0xce, 0x24, 0x53, 0xd2, 0xe1, 0x56, 0x46, 0xcb, 0x6b, 0x70, 0x12, 0xf6,
	0x57, 0x41, 0x34, 0xd3, 0xbd, 0xdf, 0x29, 0x2e, 0xa4, 0x1b, 0x6b, 0xbf, 0x45, 0x9e, 0x6a, 0x15,
	0x79, 0x02, 0xa5, 0x30, 0x5e, 0x11, 0x18, 0xba, 0xae, 0x4a, 0x7f, 0xb1, 0x7d, 0x0d, 0x0a, 0x3b,
	0x0f, 0xeb, 0x72, 0x27, 0x2c, 0x65, 0xab, 0xd7, 0x42, 0xe9, 0x7e, 0xa4, 0xcf, 0xed, 0x73, 0x90,
	0x9b, 0xfd, 0x3f, 0xf2, 0x80, 0x55, 0xf9, 0xd5, 0xf9, 0x82, 0xfb, 0x29, 0x97, 0x17, 0xdb, 0xc2,
	0x97, 0x10, 0x95, 0x58, 0xb1, 0x64, 0x29, 0xd7, 0x74, 0xfe, 0x4e, 0xe5, 0xc0, 0x0d, 0x0f, 0xaa,
	0x3b, 0x1e, 0x1c, 0xc3, 0x1b, 0xdb, 0x7d, 0x04, 0xde, 0x3f, 0xdc, 0xde, 0x46, 0xbb, 0x70, 0x7f,
	0x04, 0xe4, 0xe5, 0x59, 0xc5, 0x4e, 0xc6, 0xbd, 0x7e, 0xce, 0x82, 0x94, 0xe7, 0x93, 0x74, 0xdf,
	0xe6, 0x52, 0xc9, 0x26, 0x3c, 0xe4, 0x3e, 0x5e, 0x2b, 0xb4, 0x10, 0x7d, 0x70, 0x06, 0xb0, 0xbd,
	0x11, 0xc9, 0x21, 0xb4, 0x29, 0xba, 0x72, 0x95, 0x06, 0x92, 0x5f, 0x46, 0x3e, 0xef, 0xde, 0xc3,
	0x01, 0x6f, 0x51, 0xeb, 0xdd, 0x0b, 0x7c, 0xb7, 0xba, 0x95, 0x1b, 0x22, 0x0d, 0xed, 0x0d, 0xfe,
	0xac, 0x42, 0x77, 0x77, 0xb2, 0x09, 0x85, 0xa3, 0x21, 0x06, 0x26, 0xf9, 0x4b, 0xcc, 0x93, 0xdb,
	0xae, 0x4b, 0xf3, 0xa2, 0xf5, 0x8e, 0x2c, 0x6f, 0x6d, 0x8a, 0x21, 0xe3, 0xdd, 0xc6, 0xfb, 0xf7,
	0xc8, 0x25, 0xdc, 0x3f, 0xe7, 0xf2, 0x0e, 0x37, 0x44, 0x27, 0xcf, 0x30, 0x2b, 0x77, 0xea, 0xe4,
	0xf7, 0xf0, 0xe6, 0x98, 0x49, 0x7f, 0x79, 0xb7, 0x71, 0x8f, 0xb3, 0x3b, 0x8c, 0x7b, 0x5a, 0xd7,
	0xc4, 0x27, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x8e, 0x84, 0x47, 0x6b, 0xf6, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PersistentVolumeClient is the client API for PersistentVolume service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PersistentVolumeClient interface {
	CreatePersistentVolume(ctx context.Context, in *PersistentVolumeService, opts ...grpc.CallOption) (*ServiceResponse, error)
	GetPersistentVolume(ctx context.Context, in *PersistentVolumeService, opts ...grpc.CallOption) (*ServiceResponse, error)
	DeletePersistentVolume(ctx context.Context, in *PersistentVolumeService, opts ...grpc.CallOption) (*ServiceResponse, error)
	PatchPersistentVolume(ctx context.Context, in *PersistentVolumeService, opts ...grpc.CallOption) (*ServiceResponse, error)
	PutPersistentVolume(ctx context.Context, in *PersistentVolumeService, opts ...grpc.CallOption) (*ServiceResponse, error)
}

type persistentVolumeClient struct {
	cc grpc.ClientConnInterface
}

func NewPersistentVolumeClient(cc grpc.ClientConnInterface) PersistentVolumeClient {
	return &persistentVolumeClient{cc}
}

func (c *persistentVolumeClient) CreatePersistentVolume(ctx context.Context, in *PersistentVolumeService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.PersistentVolume/CreatePersistentVolume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *persistentVolumeClient) GetPersistentVolume(ctx context.Context, in *PersistentVolumeService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.PersistentVolume/GetPersistentVolume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *persistentVolumeClient) DeletePersistentVolume(ctx context.Context, in *PersistentVolumeService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.PersistentVolume/DeletePersistentVolume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *persistentVolumeClient) PatchPersistentVolume(ctx context.Context, in *PersistentVolumeService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.PersistentVolume/PatchPersistentVolume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *persistentVolumeClient) PutPersistentVolume(ctx context.Context, in *PersistentVolumeService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.PersistentVolume/PutPersistentVolume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PersistentVolumeServer is the server API for PersistentVolume service.
type PersistentVolumeServer interface {
	CreatePersistentVolume(context.Context, *PersistentVolumeService) (*ServiceResponse, error)
	GetPersistentVolume(context.Context, *PersistentVolumeService) (*ServiceResponse, error)
	DeletePersistentVolume(context.Context, *PersistentVolumeService) (*ServiceResponse, error)
	PatchPersistentVolume(context.Context, *PersistentVolumeService) (*ServiceResponse, error)
	PutPersistentVolume(context.Context, *PersistentVolumeService) (*ServiceResponse, error)
}

// UnimplementedPersistentVolumeServer can be embedded to have forward compatible implementations.
type UnimplementedPersistentVolumeServer struct {
}

func (*UnimplementedPersistentVolumeServer) CreatePersistentVolume(ctx context.Context, req *PersistentVolumeService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePersistentVolume not implemented")
}
func (*UnimplementedPersistentVolumeServer) GetPersistentVolume(ctx context.Context, req *PersistentVolumeService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPersistentVolume not implemented")
}
func (*UnimplementedPersistentVolumeServer) DeletePersistentVolume(ctx context.Context, req *PersistentVolumeService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePersistentVolume not implemented")
}
func (*UnimplementedPersistentVolumeServer) PatchPersistentVolume(ctx context.Context, req *PersistentVolumeService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchPersistentVolume not implemented")
}
func (*UnimplementedPersistentVolumeServer) PutPersistentVolume(ctx context.Context, req *PersistentVolumeService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutPersistentVolume not implemented")
}

func RegisterPersistentVolumeServer(s *grpc.Server, srv PersistentVolumeServer) {
	s.RegisterService(&_PersistentVolume_serviceDesc, srv)
}

func _PersistentVolume_CreatePersistentVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PersistentVolumeService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersistentVolumeServer).CreatePersistentVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PersistentVolume/CreatePersistentVolume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersistentVolumeServer).CreatePersistentVolume(ctx, req.(*PersistentVolumeService))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersistentVolume_GetPersistentVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PersistentVolumeService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersistentVolumeServer).GetPersistentVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PersistentVolume/GetPersistentVolume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersistentVolumeServer).GetPersistentVolume(ctx, req.(*PersistentVolumeService))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersistentVolume_DeletePersistentVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PersistentVolumeService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersistentVolumeServer).DeletePersistentVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PersistentVolume/DeletePersistentVolume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersistentVolumeServer).DeletePersistentVolume(ctx, req.(*PersistentVolumeService))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersistentVolume_PatchPersistentVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PersistentVolumeService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersistentVolumeServer).PatchPersistentVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PersistentVolume/PatchPersistentVolume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersistentVolumeServer).PatchPersistentVolume(ctx, req.(*PersistentVolumeService))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersistentVolume_PutPersistentVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PersistentVolumeService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersistentVolumeServer).PutPersistentVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PersistentVolume/PutPersistentVolume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersistentVolumeServer).PutPersistentVolume(ctx, req.(*PersistentVolumeService))
	}
	return interceptor(ctx, in, info, handler)
}

var _PersistentVolume_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.PersistentVolume",
	HandlerType: (*PersistentVolumeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePersistentVolume",
			Handler:    _PersistentVolume_CreatePersistentVolume_Handler,
		},
		{
			MethodName: "GetPersistentVolume",
			Handler:    _PersistentVolume_GetPersistentVolume_Handler,
		},
		{
			MethodName: "DeletePersistentVolume",
			Handler:    _PersistentVolume_DeletePersistentVolume_Handler,
		},
		{
			MethodName: "PatchPersistentVolume",
			Handler:    _PersistentVolume_PatchPersistentVolume_Handler,
		},
		{
			MethodName: "PutPersistentVolume",
			Handler:    _PersistentVolume_PutPersistentVolume_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "persistentvolume.proto",
}
