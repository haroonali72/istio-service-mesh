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
	// 984 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0x5d, 0x6e, 0x23, 0x45,
	0x10, 0x5e, 0xc7, 0x89, 0x13, 0xd7, 0xd8, 0xc6, 0xe9, 0x5d, 0xbc, 0x83, 0x61, 0x37, 0x59, 0xaf,
	0x84, 0x2c, 0xb4, 0x64, 0xa5, 0xe1, 0x85, 0x1f, 0x09, 0x29, 0x24, 0x66, 0x65, 0xc1, 0x26, 0xa6,
	0x0d, 0xe4, 0x71, 0xd4, 0x9e, 0xa9, 0x38, 0x83, 0xc7, 0xd3, 0x43, 0x77, 0x4f, 0x56, 0x86, 0x17,
	0x6e, 0xc0, 0x51, 0x78, 0xe5, 0x06, 0x9c, 0x81, 0x4b, 0x70, 0x06, 0xd4, 0x3f, 0x63, 0x3b, 0xde,
	0x44, 0xbc, 0x84, 0x27, 0x4f, 0x7d, 0x5f, 0x55, 0x75, 0x55, 0x7d, 0xd5, 0x6d, 0xe8, 0xe4, 0x28,
	0x64, 0x22, 0x15, 0x66, 0xea, 0x9a, 0xa7, 0xc5, 0x1c, 0x8f, 0x72, 0xc1, 0x15, 0x27, 0x3b, 0xe6,
	0xa7, 0x4b, 0xa4, 0xe2, 0x82, 0x4d, 0x31, 0x4a, 0x99, 0x94, 0x96, 0xea, 0x36, 0x25, 0x8a, 0xeb,
	0x24, 0xc2, 0xd2, 0x9c, 0x62, 0x86, 0x82, 0xa5, 0xce, 0x6c, 0xac, 0xa7, 0xe9, 0xfd, 0xbd, 0x05,
	0x8f, 0x47, 0xcb, 0x13, 0x7e, 0x34, 0xd4, 0xd8, 0x86, 0x93, 0x27, 0x00, 0xb9, 0xe0, 0x3f, 0x61,
	0xa4, 0xc2, 0x24, 0xf6, 0x2b, 0x87, 0x95, 0x7e, 0x9d, 0xd6, 0x1d, 0x32, 0x8c, 0x35, 0xed, 0x0e,
	0xd2, 0xf4, 0x96, 0xa5, 0x1d, 0x32, 0x8c, 0x09, 0x81, 0xed, 0x8c, 0xcd, 0xd1, 0xaf, 0x1a, 0xc2,
	0x7c, 0x13, 0x1f, 0x76, 0xaf, 0xf5, 0x61, 0x3c, 0xf3, 0xb7, 0x0d, 0x5c, 0x9a, 0xe4, 0x19, 0x34,
	0xca, 0x64, 0x6a, 0x91, 0xa3, 0xbf, 0x63, 0x68, 0xcf, 0x61, 0xdf, 0x2f, 0x72, 0x24, 0x7d, 0x68,
	0x97, 0x2e, 0xb2, 0x98, 0x58, 0xb7, 0x9a, 0x71, 0x6b, 0x39, 0x7c, 0x5c, 0x4c, 0x8c, 0xe7, 0x23,
	0xd8, 0x51, 0x7c, 0x86, 0x99, 0xbf, 0x6b, 0x68, 0x6b, 0xe8, 0x7a, 0x23, 0x3e, 0xcf, 0x59, 0xb6,
	0xd0, 0xf5, 0xee, 0xd9, 0x7a, 0x1d, 0x32, 0x8c, 0xc9, 0x08, 0x48, 0x99, 0x9e, 0x29, 0x25, 0x92,
	0x49, 0xa1, 0x50, 0xfa, 0xf5, 0xc3, 0x4a, 0xdf, 0x0b, 0x9e, 0xd9, 0x69, 0x1d, 0x6d, 0x4e, 0xea,
	0x78, 0xe9, 0x48, 0xf7, 0x5d, 0xf0, 0x0a, 0xea, 0xfd, 0xb9, 0x0d, 0xdd, 0xbb, 0x23, 0xc8, 0x00,
	0x6a, 0x29, 0x9b, 0x60, 0x2a, 0xfd, 0xca, 0x61, 0xb5, 0xef, 0x05, 0x1f, 0xff, 0xe7, 0x21, 0x47,
	0xdf, 0x1a, 0xff, 0x41, 0xa6, 0xc4, 0x82, 0xba, 0x60, 0xf2, 0x05, 0xb4, 0x84, 0x96, 0x3f, 0x99,
	0x87, 0x39, 0x4f, 0x93, 0x68, 0x61, 0xa4, 0x68, 0x05, 0x8f, 0x5c, 0x3a, 0x6a, 0xc9, 0x91, 0xe1,
	0x68, 0x53, 0xac, 0x9b, 0x24, 0x00, 0x8f, 0x45, 0x11, 0x4a, 0x19, 0xce, 0x79, 0xac, 0xb5, 0xaa,
	0xf6, 0x5b, 0xc1, 0xbe, 0x8b, 0x3c, 0x36, 0xcc, 0x6b, 0x1e, 0x23, 0x05, 0xb6, 0xfc, 0x26, 0x5d,
	0xd8, 0x8b, 0x58, 0xce, 0xa2, 0x44, 0x2d, 0x9c, 0x8a, 0x4b, 0x9b, 0x5c, 0x80, 0xbf, 0xda, 0xd7,
	0xd0, 0x6e, 0x5a, 0x28, 0x79, 0x21, 0x22, 0x2b, 0xa9, 0x17, 0x3c, 0xb9, 0xa3, 0xcb, 0xb1, 0x71,
	0xa2, 0x6b, 0xeb, 0xbe, 0x8e, 0x93, 0x17, 0x50, 0x6e, 0x7a, 0x68, 0x56, 0x3d, 0x34, 0xbb, 0x65,
	0xe5, 0x6f, 0x3b, 0xe6, 0x44, 0x13, 0x67, 0x7a, 0xcf, 0x9e, 0x43, 0x73, 0xce, 0x8b, 0x4c, 0x85,
	0x3c, 0x57, 0x09, 0xcf, 0xa4, 0xbf, 0x7b, 0x58, 0xed, 0xd7, 0x69, 0xc3, 0x80, 0xe7, 0x16, 0x23,
	0x07, 0xe0, 0xb9, 0x02, 0x4d, 0xef, 0x76, 0x21, 0xc0, 0x42, 0xa6, 0xd1, 0x2f, 0xa1, 0x99, 0xf1,
	0x18, 0x43, 0x76, 0x79, 0x99, 0x64, 0xba, 0x5b, 0xbb, 0x0c, 0xef, 0xb9, 0x0e, 0x6c, 0x7d, 0x67,
	0x3c, 0xc6, 0x63, 0xe7, 0x40, 0x1b, 0xd9, 0x9a, 0xd5, 0xfd, 0x0c, 0xbc, 0x35, 0xc1, 0x48, 0x1b,
	0xaa, 0x33, 0x5c, 0xb8, 0x7b, 0xa4, 0x3f, 0xf5, 0x9e, 0x5e, 0xb3, 0xb4, 0x40, 0x77, 0x79, 0xac,
	0xf1, 0xf9, 0xd6, 0xa7, 0x95, 0xde, 0x5f, 0x15, 0xe8, 0xdc, 0x3e, 0x21, 0xf2, 0x1c, 0x6a, 0xd3,
	0x28, 0x0f, 0x73, 0x7b, 0x23, 0xbd, 0xa0, 0xe1, 0xca, 0x79, 0x75, 0x32, 0x1a, 0x9d, 0xd2, 0x9d,
	0x69, 0x94, 0x8f, 0x62, 0xf2, 0x21, 0xec, 0xb2, 0x37, 0x32, 0xc4, 0x89, 0x34, 0xb9, 0xbd, 0xa0,
	0x59, 0x6a, 0x7a, 0x31, 0x1e, 0x7c, 0x35, 0xa6, 0x35, 0xf6, 0x46, 0x0e, 0x26, 0x92, 0xbc, 0x04,
	0x60, 0xbf, 0x14, 0x02, 0xc3, 0x38, 0x91, 0x33, 0x73, 0x55, 0xbd, 0xa0, 0x5d, 0xba, 0x6a, 0xe2,
	0x34, 0x91, 0x33, 0x5a, 0x67, 0xe5, 0xe7, 0x2a, 0xe0, 0x32, 0x49, 0xd1, 0xc8, 0xbf, 0x11, 0xf0,
	0x75, 0x92, 0xa2, 0x0b, 0xd0, 0x9f, 0xbd, 0x5f, 0x61, 0xc7, 0x54, 0x46, 0x1e, 0xc3, 0x6e, 0x1e,
	0x5b, 0xd9, 0xec, 0x08, 0x6a, 0x79, 0x6c, 0xc4, 0x3a, 0x00, 0x4f, 0x27, 0x0b, 0xe5, 0x42, 0x2a,
	0x9c, 0xbb, 0x59, 0x80, 0x86, 0xc6, 0x06, 0x21, 0x1f, 0x40, 0x3d, 0x67, 0x42, 0x31, 0x2d, 0x9b,
	0xa9, 0xb1, 0x4a, 0x57, 0x80, 0x5e, 0x47, 0x81, 0x2c, 0xe6, 0x59, 0x6a, 0xd7, 0x71, 0x8f, 0x2e,
	0xed, 0xde, 0x6f, 0x15, 0xa8, 0xd9, 0x8e, 0xc9, 0xfb, 0x50, 0x77, 0x6a, 0x2f, 0xdf, 0xb2, 0x3d,
	0x0b, 0x0c, 0xe3, 0xff, 0xb3, 0x84, 0x7f, 0x2a, 0x50, 0x5f, 0x4e, 0x92, 0x1c, 0x83, 0x17, 0xb1,
	0xe8, 0x2a, 0xc9, 0xa6, 0x7a, 0xc3, 0x4c, 0x1d, 0xad, 0xe0, 0xe0, 0xc6, 0xc0, 0x99, 0x62, 0xda,
	0xf5, 0x64, 0xe5, 0x46, 0xd7, 0x63, 0xc8, 0x0b, 0xd8, 0x9e, 0x25, 0x59, 0x6c, 0x76, 0xbf, 0x15,
	0xf8, 0xb7, 0xc5, 0x7e, 0x93, 0x64, 0x31, 0x35, 0x5e, 0xe4, 0x29, 0xac, 0xb5, 0x71, 0x4b, 0x63,
	0x5d, 0xd8, 0xd3, 0xd2, 0x9f, 0xad, 0x5e, 0xea, 0xa5, 0x5d, 0xb6, 0x75, 0xbe, 0xd1, 0x96, 0xb6,
	0xf5, 0x4b, 0xae, 0xfd, 0x7e, 0xa0, 0x43, 0xf7, 0x54, 0x97, 0x66, 0xef, 0xf7, 0xb2, 0x61, 0x2d,
	0xbf, 0x3e, 0x5f, 0x62, 0x24, 0x50, 0x9d, 0xad, 0x84, 0x5f, 0x43, 0xf4, 0x60, 0xe5, 0x15, 0x13,
	0x68, 0xe8, 0xf2, 0x3f, 0xa4, 0x04, 0x6e, 0x54, 0x50, 0xdd, 0xa8, 0xa0, 0x0f, 0xef, 0xac, 0xf2,
	0xc8, 0x9c, 0x45, 0xe8, 0x5e, 0xa3, 0x4d, 0xb8, 0x37, 0x00, 0xf2, 0xf6, 0x5d, 0x25, 0x2f, 0x75,
	0xee, 0x9f, 0x8b, 0x44, 0x60, 0x79, 0x93, 0x1e, 0xba, 0x59, 0x6a, 0xb7, 0x31, 0xa6, 0x18, 0x29,
	0x2e, 0xe8, 0xd2, 0xe9, 0xa3, 0x53, 0x80, 0xd5, 0x8b, 0x48, 0xf6, 0xa1, 0x49, 0x91, 0xc5, 0x17,
	0x22, 0x51, 0x78, 0x9e, 0x45, 0xd8, 0x7e, 0x40, 0xda, 0xd0, 0xa0, 0xae, 0xba, 0xd7, 0x2c, 0x5b,
	0xb4, 0x2b, 0x37, 0x9c, 0x0c, 0xb4, 0x15, 0xfc, 0x51, 0x85, 0xf6, 0xe6, 0xcd, 0x26, 0x14, 0x3a,
	0x27, 0x02, 0x99, 0xc2, 0xb7, 0x98, 0xa7, 0x77, 0x3d, 0x97, 0xf6, 0xdf, 0xa6, 0xdb, 0x71, 0xbc,
	0xb3, 0x29, 0xca, 0x9c, 0x67, 0x12, 0x7b, 0x0f, 0xc8, 0x39, 0x3c, 0x7c, 0x85, 0xea, 0x1e, 0x13,
	0x52, 0xe8, 0x9c, 0x62, 0x8a, 0xf7, 0x5a, 0xe4, 0x77, 0xf0, 0xee, 0x88, 0xa9, 0xe8, 0xea, 0x7e,
	0xfb, 0x1e, 0x15, 0xf7, 0xd8, 0xf7, 0xa4, 0x66, 0x88, 0x4f, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff,
	0xc5, 0x9b, 0xb7, 0x00, 0x8a, 0x09, 0x00, 0x00,
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
