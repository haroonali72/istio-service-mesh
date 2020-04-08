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

type PersistentVolumeMode int32

const (
	PersistentVolumeMode_Block      PersistentVolumeMode = 0
	PersistentVolumeMode_Filesystem PersistentVolumeMode = 1
)

var PersistentVolumeMode_name = map[int32]string{
	0: "Block",
	1: "Filesystem",
}

var PersistentVolumeMode_value = map[string]int32{
	"Block":      0,
	"Filesystem": 1,
}

func (x PersistentVolumeMode) String() string {
	return proto.EnumName(PersistentVolumeMode_name, int32(x))
}

func (PersistentVolumeMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_724324990baa6ac6, []int{1}
}

type PersistentVolumeService struct {
	ProjectId            string                      `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	ServiceId            string                      `protobuf:"bytes,2,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	Name                 string                      `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Version              string                      `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	ServiceType          string                      `protobuf:"bytes,5,opt,name=service_type,json=serviceType,proto3" json:"service_type,omitempty"`
	ServiceSubType       string                      `protobuf:"bytes,6,opt,name=service_sub_type,json=serviceSubType,proto3" json:"service_sub_type,omitempty"`
	Token                string                      `protobuf:"bytes,7,opt,name=token,proto3" json:"token,omitempty"`
	CompanyId            string                      `protobuf:"bytes,8,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	ServiceAttributes    *PersistentVolumeAttributes `protobuf:"bytes,9,opt,name=service_attributes,json=serviceAttributes,proto3" json:"service_attributes,omitempty"`
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

func (m *PersistentVolumeService) GetProjectId() string {
	if m != nil {
		return m.ProjectId
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

type PersistentVolumeAttributes struct {
	Labels                 map[string]string       `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ReclaimPolicy          ReclaimPolicy           `protobuf:"varint,2,opt,name=reclaim_policy,json=reclaimPolicy,proto3,enum=proto.ReclaimPolicy" json:"reclaim_policy,omitempty"`
	AccessMode             []AccessMode            `protobuf:"varint,3,rep,packed,name=access_mode,json=accessMode,proto3,enum=proto.AccessMode" json:"access_mode,omitempty"`
	Capcity                string                  `protobuf:"bytes,4,opt,name=capcity,proto3" json:"capcity,omitempty"`
	PersistentVolumeSource *PersistentVolumeSource `protobuf:"bytes,5,opt,name=persistent_volume_source,json=persistentVolumeSource,proto3" json:"persistent_volume_source,omitempty"`
	StorageClassName       string                  `protobuf:"bytes,6,opt,name=storage_class_name,json=storageClassName,proto3" json:"storage_class_name,omitempty"`
	MountOptions           []string                `protobuf:"bytes,7,rep,name=mount_options,json=mountOptions,proto3" json:"mount_options,omitempty"`
	VolumeMode             PersistentVolumeMode    `protobuf:"varint,8,opt,name=volume_mode,json=volumeMode,proto3,enum=proto.PersistentVolumeMode" json:"volume_mode,omitempty"`
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

func (m *PersistentVolumeAttributes) GetCapcity() string {
	if m != nil {
		return m.Capcity
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

func (m *PersistentVolumeAttributes) GetVolumeMode() PersistentVolumeMode {
	if m != nil {
		return m.VolumeMode
	}
	return PersistentVolumeMode_Block
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
	proto.RegisterEnum("proto.PersistentVolumeMode", PersistentVolumeMode_name, PersistentVolumeMode_value)
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
	// 1012 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xdd, 0x6e, 0x1b, 0x45,
	0x14, 0x8e, 0xe3, 0xd8, 0xb1, 0xcf, 0xda, 0xc6, 0x99, 0x86, 0x74, 0x71, 0x69, 0x9b, 0xba, 0x12,
	0xb2, 0xaa, 0x92, 0x8a, 0xe5, 0x86, 0x3f, 0x21, 0xa5, 0x49, 0xa8, 0x22, 0x68, 0x62, 0xc6, 0x40,
	0x2e, 0x57, 0xe3, 0xdd, 0x93, 0x64, 0xf1, 0x7a, 0x67, 0x99, 0x99, 0x4d, 0x65, 0xb8, 0xe1, 0x0d,
	0x78, 0x14, 0x1e, 0x83, 0x07, 0xe0, 0x8a, 0x97, 0xe0, 0x19, 0xd0, 0xfc, 0xac, 0xed, 0xb8, 0x89,
	0xb8, 0x49, 0xaf, 0x3c, 0xe7, 0xfb, 0xbe, 0x73, 0xf6, 0xfc, 0xcd, 0x18, 0x76, 0x72, 0x14, 0x32,
	0x91, 0x0a, 0x33, 0x75, 0xc5, 0xd3, 0x62, 0x8a, 0x7b, 0xb9, 0xe0, 0x8a, 0x93, 0x9a, 0xf9, 0xe9,
	0x11, 0xa9, 0xb8, 0x60, 0x17, 0x18, 0xa5, 0x4c, 0x4a, 0x4b, 0xf5, 0xda, 0x12, 0xc5, 0x55, 0x12,
	0x61, 0x69, 0x5e, 0x60, 0x86, 0x82, 0xa5, 0xce, 0x6c, 0x2d, 0x87, 0xe9, 0xff, 0xb3, 0x0e, 0xf7,
	0x87, 0xf3, 0x2f, 0xfc, 0x64, 0xa8, 0x91, 0x75, 0x27, 0x0f, 0x01, 0x72, 0xc1, 0x7f, 0xc6, 0x48,
	0x85, 0x49, 0xec, 0x57, 0x76, 0x2b, 0x83, 0x26, 0x6d, 0x3a, 0xe4, 0x38, 0xd6, 0xb4, 0xfb, 0x90,
	0xa6, 0xd7, 0x2d, 0xed, 0x90, 0xe3, 0x98, 0x10, 0xd8, 0xc8, 0xd8, 0x14, 0xfd, 0xaa, 0x21, 0xcc,
	0x99, 0xf8, 0xb0, 0x79, 0xa5, 0x3f, 0xc6, 0x33, 0x7f, 0xc3, 0xc0, 0xa5, 0x49, 0x9e, 0x40, 0xab,
	0x0c, 0xa6, 0x66, 0x39, 0xfa, 0x35, 0x43, 0x7b, 0x0e, 0xfb, 0x61, 0x96, 0x23, 0x19, 0x40, 0xb7,
	0x94, 0xc8, 0x62, 0x6c, 0x65, 0x75, 0x23, 0xeb, 0x38, 0x7c, 0x54, 0x8c, 0x8d, 0x72, 0x1b, 0x6a,
	0x8a, 0x4f, 0x30, 0xf3, 0x37, 0x0d, 0x6d, 0x0d, 0x9d, 0x6f, 0xc4, 0xa7, 0x39, 0xcb, 0x66, 0x3a,
	0xdf, 0x86, 0xcd, 0xd7, 0x21, 0xc7, 0x31, 0x19, 0x02, 0x29, 0xc3, 0x33, 0xa5, 0x44, 0x32, 0x2e,
	0x14, 0x4a, 0xbf, 0xb9, 0x5b, 0x19, 0x78, 0xc1, 0x13, 0xdb, 0xad, 0xbd, 0xd5, 0x4e, 0xed, 0xcf,
	0x85, 0x74, 0xcb, 0x39, 0x2f, 0xa0, 0xfe, 0xdf, 0x1b, 0xd0, 0xbb, 0xdd, 0x83, 0x1c, 0x41, 0x3d,
	0x65, 0x63, 0x4c, 0xa5, 0x5f, 0xd9, 0xad, 0x0e, 0xbc, 0xe0, 0xe3, 0xff, 0xfd, 0xc8, 0xde, 0x77,
	0x46, 0x7f, 0x94, 0x29, 0x31, 0xa3, 0xce, 0x99, 0x7c, 0x09, 0x1d, 0xa1, 0xc7, 0x9f, 0x4c, 0xc3,
	0x9c, 0xa7, 0x49, 0x34, 0x33, 0xa3, 0xe8, 0x04, 0xdb, 0x2e, 0x1c, 0xb5, 0xe4, 0xd0, 0x70, 0xb4,
	0x2d, 0x96, 0x4d, 0x12, 0x80, 0xc7, 0xa2, 0x08, 0xa5, 0x0c, 0xa7, 0x3c, 0xd6, 0xb3, 0xaa, 0x0e,
	0x3a, 0xc1, 0x96, 0xf3, 0xdc, 0x37, 0xcc, 0x6b, 0x1e, 0x23, 0x05, 0x36, 0x3f, 0xeb, 0x21, 0x46,
	0x2c, 0x8f, 0x12, 0x35, 0x2b, 0x87, 0xe8, 0x4c, 0x72, 0x06, 0xfe, 0x62, 0x5b, 0x43, 0xbb, 0x67,
	0xa1, 0xe4, 0x85, 0x88, 0xec, 0x40, 0xbd, 0xe0, 0xe1, 0x2d, 0x35, 0x8e, 0x8c, 0x88, 0x2e, 0x2d,
	0xfb, 0x32, 0x4e, 0x9e, 0x43, 0xb9, 0xe7, 0xa1, 0x59, 0xf4, 0xd0, 0x6c, 0x96, 0x1d, 0x7e, 0xd7,
	0x31, 0x07, 0x9a, 0x38, 0xd1, 0x5b, 0xf6, 0x14, 0xda, 0x53, 0x5e, 0x64, 0x2a, 0xe4, 0xb9, 0x4a,
	0x78, 0x26, 0xfd, 0xcd, 0xdd, 0xea, 0xa0, 0x49, 0x5b, 0x06, 0x3c, 0xb5, 0x18, 0xf9, 0x0a, 0x3c,
	0x97, 0xa0, 0xa9, 0xbc, 0x61, 0x7a, 0xf6, 0xe0, 0x96, 0xf4, 0x6c, 0x0f, 0xae, 0xe6, 0x67, 0xf2,
	0x35, 0xb4, 0x33, 0x1e, 0x63, 0xc8, 0xce, 0xcf, 0x93, 0x4c, 0x77, 0xc2, 0xee, 0xc9, 0x07, 0xce,
	0xdf, 0x7a, 0x9d, 0xf0, 0x18, 0xf7, 0x9d, 0x80, 0xb6, 0xb2, 0x25, 0xab, 0xf7, 0x39, 0x78, 0x4b,
	0xb3, 0x24, 0x5d, 0xa8, 0x4e, 0x70, 0xe6, 0xae, 0x98, 0x3e, 0xea, 0x15, 0xbe, 0x62, 0x69, 0x81,
	0xee, 0x5e, 0x59, 0xe3, 0x8b, 0xf5, 0xcf, 0x2a, 0xfd, 0xbf, 0x2a, 0xb0, 0x73, 0x73, 0xfb, 0xc8,
	0x53, 0xa8, 0x5f, 0x44, 0x79, 0x98, 0xdb, 0xcb, 0xea, 0x05, 0x2d, 0x97, 0xce, 0xab, 0x83, 0xe1,
	0xf0, 0x90, 0xd6, 0x2e, 0xa2, 0x7c, 0x18, 0x93, 0x8f, 0x60, 0x93, 0xbd, 0x91, 0x21, 0x8e, 0xa5,
	0x89, 0xed, 0x05, 0xed, 0x72, 0xdc, 0x67, 0xa3, 0xa3, 0x97, 0x23, 0x5a, 0x67, 0x6f, 0xe4, 0xd1,
	0x58, 0x92, 0x17, 0x00, 0xec, 0xd7, 0x42, 0x60, 0x18, 0x27, 0x72, 0x62, 0x6e, 0xb1, 0x17, 0x74,
	0x4b, 0xa9, 0x26, 0x0e, 0x13, 0x39, 0xa1, 0x4d, 0x56, 0x1e, 0x17, 0x0e, 0xe7, 0x49, 0x8a, 0x66,
	0x35, 0x56, 0x1c, 0xbe, 0x49, 0x52, 0x74, 0x0e, 0xfa, 0xd8, 0xff, 0x0d, 0x6a, 0x26, 0x33, 0x72,
	0x1f, 0x36, 0xf3, 0xd8, 0xce, 0xd4, 0xb6, 0xa0, 0x9e, 0xc7, 0x66, 0x92, 0x8f, 0xc1, 0xd3, 0xc1,
	0x42, 0x39, 0x93, 0x0a, 0xa7, 0xae, 0x17, 0xa0, 0xa1, 0x91, 0x41, 0xc8, 0x87, 0xd0, 0xcc, 0x99,
	0x50, 0x4c, 0xcf, 0xd4, 0xe4, 0x58, 0xa5, 0x0b, 0x80, 0xf4, 0xa0, 0x21, 0x90, 0xc5, 0x3c, 0x4b,
	0xed, 0xaa, 0x36, 0xe8, 0xdc, 0xee, 0xff, 0x5e, 0x81, 0xba, 0xad, 0x98, 0x3c, 0x80, 0xa6, 0x5b,
	0x85, 0xf9, 0x33, 0xd7, 0xb0, 0xc0, 0x71, 0xfc, 0x2e, 0x53, 0xf8, 0xb7, 0x02, 0xcd, 0x79, 0x27,
	0xc9, 0x3e, 0x78, 0x11, 0x8b, 0x2e, 0x93, 0xec, 0x42, 0x6f, 0x98, 0xc9, 0xa3, 0x13, 0x3c, 0xbe,
	0xd6, 0x70, 0xa6, 0x98, 0x96, 0x1e, 0x2c, 0x64, 0x74, 0xd9, 0x87, 0x3c, 0x87, 0x8d, 0x49, 0x92,
	0xc5, 0xe6, 0x62, 0x74, 0x02, 0xff, 0x26, 0xdf, 0x6f, 0x93, 0x2c, 0xa6, 0x46, 0x45, 0x1e, 0xc1,
	0x52, 0x19, 0x37, 0x14, 0xd6, 0x83, 0x86, 0x1e, 0xfd, 0xc9, 0xe2, 0x11, 0x9f, 0xdb, 0x65, 0x59,
	0xa7, 0x2b, 0x65, 0x69, 0x5b, 0xbf, 0x0f, 0x5a, 0xf7, 0x23, 0x3d, 0x76, 0xaf, 0x78, 0x69, 0xf6,
	0xff, 0x28, 0x0b, 0xd6, 0xe3, 0xd7, 0xdf, 0x97, 0x18, 0x09, 0x54, 0x27, 0x8b, 0xc1, 0x2f, 0x21,
	0xba, 0xb1, 0xf2, 0x92, 0x09, 0x34, 0x74, 0xf9, 0xf7, 0x52, 0x02, 0xd7, 0x32, 0xa8, 0xae, 0x64,
	0x30, 0x80, 0xf7, 0x16, 0x71, 0x64, 0xce, 0x22, 0x74, 0x2f, 0xd5, 0x2a, 0xdc, 0x3f, 0x02, 0xf2,
	0xf6, 0x5d, 0x25, 0x2f, 0x74, 0xec, 0x5f, 0x8a, 0x44, 0x60, 0x79, 0x93, 0xee, 0xb9, 0x5e, 0x6a,
	0xd9, 0x08, 0x53, 0x8c, 0x14, 0x17, 0x74, 0x2e, 0x7a, 0x76, 0x08, 0xb0, 0x78, 0x2c, 0xc9, 0x16,
	0xb4, 0x29, 0xb2, 0xf8, 0x4c, 0x24, 0x0a, 0x4f, 0xb3, 0x08, 0xbb, 0x6b, 0xa4, 0x0b, 0x2d, 0xea,
	0xb2, 0x7b, 0xcd, 0xb2, 0x59, 0xb7, 0x72, 0x4d, 0x64, 0xa0, 0xf5, 0x67, 0x9f, 0xc0, 0xf6, 0x4d,
	0x0f, 0x0f, 0x69, 0x42, 0xed, 0x65, 0xca, 0xa3, 0x49, 0x77, 0x8d, 0x74, 0x00, 0x74, 0xef, 0xec,
	0x32, 0x76, 0x2b, 0xc1, 0x9f, 0x55, 0xe8, 0xae, 0xfa, 0x10, 0x0a, 0x3b, 0x07, 0x02, 0x99, 0xc2,
	0xb7, 0x98, 0x47, 0xb7, 0x3d, 0xbf, 0xf6, 0xbf, 0xab, 0xb7, 0xe3, 0x78, 0x67, 0x53, 0x94, 0x39,
	0xcf, 0x24, 0xf6, 0xd7, 0xc8, 0x29, 0xdc, 0x7b, 0x85, 0xea, 0x0e, 0x03, 0x52, 0xd8, 0x39, 0xc4,
	0x14, 0xef, 0x34, 0xc9, 0xef, 0xe1, 0xfd, 0x21, 0x53, 0xd1, 0xe5, 0xdd, 0xd6, 0x3d, 0x2c, 0xee,
	0xb0, 0xee, 0x71, 0xdd, 0x10, 0x9f, 0xfe, 0x17, 0x00, 0x00, 0xff, 0xff, 0x81, 0xec, 0x67, 0x85,
	0xd8, 0x09, 0x00, 0x00,
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