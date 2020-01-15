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
	EnvironmentVariables          map[string]*EnvironmentVariable `protobuf:"bytes,1,rep,name=environment_variables,json=environmentVariables,proto3" json:"environment_variables,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ImageRepositoryConfigurations *ImageRepositoryConfigurations  `protobuf:"bytes,2,opt,name=image_repository_configurations,json=imageRepositoryConfigurations,proto3" json:"image_repository_configurations,omitempty"`
	Ports                         map[string]*ContainerPort       `protobuf:"bytes,3,rep,name=ports,proto3" json:"ports,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Tag                           string                          `protobuf:"bytes,4,opt,name=tag,proto3" json:"tag,omitempty"`
	ImagePrefix                   string                          `protobuf:"bytes,5,opt,name=image_prefix,json=imagePrefix,proto3" json:"image_prefix,omitempty"`
	ImageName                     string                          `protobuf:"bytes,6,opt,name=image_name,json=imageName,proto3" json:"image_name,omitempty"`
	IstioConfig                   *IstioConfig                    `protobuf:"bytes,7,opt,name=istio_config,json=istioConfig,proto3" json:"istio_config,omitempty"`
	LabelSelector                 *LabelSelectorObj               `protobuf:"bytes,8,opt,name=label_selector,json=labelSelector,proto3" json:"label_selector,omitempty"`
	NodeSelector                  map[string]string               `protobuf:"bytes,9,rep,name=node_selector,json=nodeSelector,proto3" json:"node_selector,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Command                       []string                        `protobuf:"bytes,10,rep,name=command,proto3" json:"command,omitempty"`
	Args                          []string                        `protobuf:"bytes,11,rep,name=args,proto3" json:"args,omitempty"`
	SecurityContext               *SecurityContextStruct          `protobuf:"bytes,12,opt,name=security_context,json=securityContext,proto3" json:"security_context,omitempty"`
	LimitResources                map[string]string               `protobuf:"bytes,13,rep,name=limit_resources,json=limitResources,proto3" json:"limit_resources,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	RequestResources              map[string]string               `protobuf:"bytes,14,rep,name=request_resources,json=requestResources,proto3" json:"request_resources,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Labels                        map[string]string               `protobuf:"bytes,15,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Annotations                   map[string]string               `protobuf:"bytes,16,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	LivenessProbe                 *Probe                          `protobuf:"bytes,17,opt,name=liveness_probe,json=livenessProbe,proto3" json:"liveness_probe,omitempty"`
	ReadinessProbe                *Probe                          `protobuf:"bytes,18,opt,name=readiness_probe,json=readinessProbe,proto3" json:"readiness_probe,omitempty"`
	IsRbacEnabled                 bool                            `protobuf:"varint,19,opt,name=is_rbac_enabled,json=isRbacEnabled,proto3" json:"is_rbac_enabled,omitempty"`
	Roles                         []*K8SRbacAttribute             `protobuf:"bytes,20,rep,name=roles,proto3" json:"roles,omitempty"`
	IstioRoles                    []*IstioRbacAttribute           `protobuf:"bytes,21,rep,name=istio_roles,json=istioRoles,proto3" json:"istio_roles,omitempty"`
	EnableInit                    bool                            `protobuf:"varint,22,opt,name=enable_init,json=enableInit,proto3" json:"enable_init,omitempty"`
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

func (m *DeploymentServiceAttributes) GetEnvironmentVariables() map[string]*EnvironmentVariable {
	if m != nil {
		return m.EnvironmentVariables
	}
	return nil
}

func (m *DeploymentServiceAttributes) GetImageRepositoryConfigurations() *ImageRepositoryConfigurations {
	if m != nil {
		return m.ImageRepositoryConfigurations
	}
	return nil
}

func (m *DeploymentServiceAttributes) GetPorts() map[string]*ContainerPort {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *DeploymentServiceAttributes) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *DeploymentServiceAttributes) GetImagePrefix() string {
	if m != nil {
		return m.ImagePrefix
	}
	return ""
}

func (m *DeploymentServiceAttributes) GetImageName() string {
	if m != nil {
		return m.ImageName
	}
	return ""
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

func (m *DeploymentServiceAttributes) GetCommand() []string {
	if m != nil {
		return m.Command
	}
	return nil
}

func (m *DeploymentServiceAttributes) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *DeploymentServiceAttributes) GetSecurityContext() *SecurityContextStruct {
	if m != nil {
		return m.SecurityContext
	}
	return nil
}

func (m *DeploymentServiceAttributes) GetLimitResources() map[string]string {
	if m != nil {
		return m.LimitResources
	}
	return nil
}

func (m *DeploymentServiceAttributes) GetRequestResources() map[string]string {
	if m != nil {
		return m.RequestResources
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

func (m *DeploymentServiceAttributes) GetLivenessProbe() *Probe {
	if m != nil {
		return m.LivenessProbe
	}
	return nil
}

func (m *DeploymentServiceAttributes) GetReadinessProbe() *Probe {
	if m != nil {
		return m.ReadinessProbe
	}
	return nil
}

func (m *DeploymentServiceAttributes) GetIsRbacEnabled() bool {
	if m != nil {
		return m.IsRbacEnabled
	}
	return false
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

func (m *DeploymentServiceAttributes) GetEnableInit() bool {
	if m != nil {
		return m.EnableInit
	}
	return false
}

func init() {
	proto.RegisterType((*DeploymentService)(nil), "proto.DeploymentService")
	proto.RegisterType((*DeploymentServiceAttributes)(nil), "proto.DeploymentServiceAttributes")
	proto.RegisterMapType((map[string]string)(nil), "proto.DeploymentServiceAttributes.AnnotationsEntry")
	proto.RegisterMapType((map[string]*EnvironmentVariable)(nil), "proto.DeploymentServiceAttributes.EnvironmentVariablesEntry")
	proto.RegisterMapType((map[string]string)(nil), "proto.DeploymentServiceAttributes.LabelsEntry")
	proto.RegisterMapType((map[string]string)(nil), "proto.DeploymentServiceAttributes.LimitResourcesEntry")
	proto.RegisterMapType((map[string]string)(nil), "proto.DeploymentServiceAttributes.NodeSelectorEntry")
	proto.RegisterMapType((map[string]*ContainerPort)(nil), "proto.DeploymentServiceAttributes.PortsEntry")
	proto.RegisterMapType((map[string]string)(nil), "proto.DeploymentServiceAttributes.RequestResourcesEntry")
}

func init() { proto.RegisterFile("deployment.proto", fileDescriptor_fac0ec10f8e4d7ff) }

var fileDescriptor_fac0ec10f8e4d7ff = []byte{
	// 954 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0x6d, 0x6f, 0xdb, 0x36,
	0x10, 0x9e, 0x93, 0x3a, 0xad, 0x4f, 0xf1, 0x4b, 0xd8, 0xa4, 0x53, 0xb3, 0x16, 0xcd, 0x82, 0x61,
	0x08, 0x06, 0x34, 0x18, 0x9c, 0x75, 0xc8, 0x8a, 0xa1, 0x43, 0xe0, 0xa6, 0x45, 0xb0, 0x22, 0xf3,
	0xe4, 0x6d, 0xc0, 0x3e, 0x09, 0x94, 0x74, 0xf5, 0xd8, 0xca, 0xa4, 0x4a, 0x52, 0x46, 0xfd, 0x6f,
	0xf7, 0x07, 0xf6, 0x65, 0xbf, 0x60, 0xe0, 0x8b, 0x6c, 0x39, 0x75, 0x3b, 0x07, 0xf9, 0x64, 0xf1,
	0xb9, 0xe7, 0x79, 0x78, 0x47, 0xde, 0xd1, 0xd0, 0xcb, 0xb0, 0xc8, 0xc5, 0x6c, 0x82, 0x5c, 0x1f,
	0x17, 0x52, 0x68, 0x41, 0x9a, 0xf6, 0x67, 0xbf, 0x3d, 0x46, 0x8e, 0x92, 0xe6, 0xc7, 0x7e, 0xa9,
	0x50, 0x4e, 0x59, 0x8a, 0x6e, 0x79, 0xf8, 0xef, 0x06, 0xec, 0x3c, 0x9f, 0x2b, 0x47, 0x2e, 0x46,
	0x1e, 0x02, 0x14, 0x52, 0xbc, 0xc1, 0x54, 0xc7, 0x2c, 0x0b, 0x1b, 0x07, 0x8d, 0xa3, 0x56, 0xd4,
	0xf2, 0xc8, 0x45, 0x66, 0xc2, 0xde, 0xc5, 0x84, 0x37, 0x5c, 0xd8, 0x23, 0x17, 0x19, 0x21, 0x70,
	0x8b, 0xd3, 0x09, 0x86, 0x9b, 0x36, 0x60, 0xbf, 0x49, 0x08, 0xb7, 0xa7, 0x28, 0x15, 0x13, 0x3c,
	0xbc, 0x65, 0xe1, 0x6a, 0x49, 0xbe, 0x84, 0xed, 0xca, 0x4c, 0xcf, 0x0a, 0x0c, 0x9b, 0x36, 0x1c,
	0x78, 0xec, 0xb7, 0x59, 0x81, 0xe4, 0x08, 0x7a, 0x15, 0x45, 0x95, 0x89, 0xa3, 0x6d, 0x59, 0x5a,
	0xc7, 0xe3, 0xa3, 0x32, 0xb1, 0xcc, 0x07, 0xd0, 0x32, 0xdb, 0xa9, 0x82, 0xa6, 0x18, 0xde, 0x76,
	0x89, 0xcd, 0x01, 0xb2, 0x0b, 0x4d, 0x2d, 0xde, 0x22, 0x0f, 0xef, 0xd8, 0x88, 0x5b, 0x98, 0x6a,
	0x52, 0x31, 0x29, 0x28, 0x9f, 0x99, 0x6a, 0x5a, 0x4e, 0xe4, 0x91, 0x8b, 0x8c, 0xfc, 0x0a, 0xa4,
	0xda, 0x9c, 0x6a, 0x2d, 0x59, 0x52, 0x6a, 0x54, 0x21, 0x1c, 0x34, 0x8e, 0x82, 0xfe, 0xa1, 0x3b,
	0xc5, 0xe3, 0x0f, 0x4e, 0xf0, 0x6c, 0xce, 0x8c, 0x76, 0xd4, 0x55, 0xe8, 0xf0, 0xef, 0x0e, 0x7c,
	0xf1, 0x09, 0x09, 0x79, 0x07, 0x7b, 0xc8, 0xa7, 0x4c, 0x0a, 0x6e, 0xe2, 0xf1, 0x94, 0x4a, 0x46,
	0x93, 0x1c, 0x55, 0xd8, 0x38, 0xd8, 0x3c, 0x0a, 0xfa, 0x3f, 0xfe, 0xff, 0xae, 0xc7, 0xe7, 0x0b,
	0xfd, 0x1f, 0x95, 0xfc, 0x9c, 0x6b, 0x39, 0x8b, 0x76, 0x71, 0x45, 0x88, 0xe4, 0xf0, 0x88, 0x4d,
	0xe8, 0x18, 0x63, 0x89, 0x85, 0x50, 0x4c, 0x0b, 0x39, 0x8b, 0x53, 0xc1, 0x5f, 0xb3, 0x71, 0x29,
	0xa9, 0x66, 0x82, 0x2b, 0x7b, 0xcf, 0x41, 0xff, 0x2b, 0xbf, 0xf9, 0x85, 0x61, 0x47, 0x73, 0xf2,
	0x60, 0x89, 0x1b, 0x3d, 0x64, 0x9f, 0x0a, 0x93, 0x01, 0x34, 0x0b, 0x21, 0xb5, 0x0a, 0x37, 0x6d,
	0x41, 0x8f, 0xd7, 0x28, 0x68, 0x68, 0xf8, 0xae, 0x02, 0xa7, 0x25, 0x3d, 0xd8, 0xd4, 0x74, 0xec,
	0xdb, 0xc9, 0x7c, 0x9a, 0x56, 0x72, 0x45, 0x14, 0x12, 0x5f, 0xb3, 0xf7, 0x55, 0x2b, 0x59, 0x6c,
	0x68, 0x21, 0x73, 0xd9, 0x8e, 0x62, 0x3b, 0xd4, 0x35, 0x51, 0xcb, 0x22, 0x97, 0xa6, 0x4d, 0x9f,
	0xc0, 0x36, 0x53, 0x9a, 0x09, 0x5f, 0xbb, 0x6d, 0xa1, 0xa0, 0x4f, 0xaa, 0x9a, 0x4d, 0xc8, 0x95,
	0x12, 0x05, 0x6c, 0xb1, 0x20, 0xcf, 0xa0, 0x93, 0xd3, 0x04, 0xf3, 0x58, 0x61, 0x8e, 0xa9, 0x16,
	0xd2, 0x76, 0x58, 0xd0, 0xff, 0xdc, 0x0b, 0x5f, 0x99, 0xe0, 0xc8, 0xc7, 0x7e, 0x49, 0xde, 0x44,
	0xed, 0xbc, 0x8e, 0x90, 0x3f, 0xa1, 0xcd, 0x45, 0x86, 0x0b, 0x79, 0xcb, 0x9e, 0xcb, 0x77, 0x6b,
	0x9c, 0xcb, 0xa5, 0xc8, 0xb0, 0xf2, 0x71, 0xc7, 0xb3, 0xcd, 0x6b, 0x90, 0x19, 0xbc, 0x54, 0x4c,
	0x26, 0x94, 0x67, 0x21, 0x1c, 0x6c, 0x9a, 0xc1, 0xf3, 0x4b, 0x33, 0xa6, 0x54, 0x8e, 0x55, 0x18,
	0x58, 0xd8, 0x7e, 0x93, 0x97, 0x66, 0xd2, 0xd2, 0x52, 0x32, 0x6d, 0xaf, 0x5f, 0xe3, 0x7b, 0x1d,
	0x6e, 0xdb, 0x52, 0x1e, 0xf8, 0x5c, 0x46, 0x3e, 0x3c, 0x70, 0xd1, 0x91, 0x96, 0x65, 0xaa, 0xa3,
	0xae, 0x5a, 0x86, 0x49, 0x0c, 0xdd, 0x9c, 0x4d, 0x98, 0x8e, 0x25, 0x2a, 0x51, 0xca, 0x14, 0x55,
	0xd8, 0xb6, 0x35, 0x7d, 0xbf, 0x46, 0x4d, 0xaf, 0x8c, 0x32, 0xaa, 0x84, 0xae, 0xaa, 0x4e, 0xbe,
	0x04, 0x12, 0x84, 0x1d, 0x89, 0xef, 0x4a, 0x54, 0xf5, 0x2d, 0x3a, 0x76, 0x8b, 0xd3, 0x35, 0xb6,
	0x88, 0x9c, 0xf6, 0xca, 0x26, 0x3d, 0x79, 0x05, 0x26, 0x2f, 0x60, 0xcb, 0x5e, 0x95, 0x0a, 0xbb,
	0xd6, 0xfb, 0x78, 0x9d, 0xf4, 0xad, 0xc0, 0x39, 0x7a, 0x35, 0xf9, 0x1d, 0x02, 0xca, 0xb9, 0xd0,
	0x7e, 0x96, 0x7a, 0xd6, 0xec, 0x64, 0x0d, 0xb3, 0xb3, 0x85, 0xca, 0x39, 0xd6, 0x7d, 0xc8, 0x09,
	0x74, 0x72, 0x36, 0x45, 0x8e, 0x4a, 0xc5, 0x85, 0x14, 0x09, 0x86, 0x3b, 0xf6, 0xb6, 0xb6, 0xbd,
	0xf3, 0xd0, 0x60, 0x51, 0xbb, 0xe2, 0xd8, 0x25, 0x79, 0x02, 0x5d, 0x89, 0x34, 0x63, 0x35, 0x15,
	0x59, 0xa1, 0xea, 0xcc, 0x49, 0x4e, 0xf6, 0x35, 0x74, 0x99, 0x8a, 0x65, 0x42, 0xd3, 0x18, 0xb9,
	0x79, 0x35, 0xb2, 0xf0, 0xee, 0x41, 0xe3, 0xe8, 0x4e, 0xd4, 0x66, 0x2a, 0x4a, 0x68, 0x7a, 0xee,
	0x40, 0xf2, 0x18, 0x9a, 0x52, 0x98, 0xd7, 0x6a, 0xd7, 0x16, 0x59, 0xcd, 0xc0, 0xcf, 0xa7, 0x96,
	0x35, 0x2f, 0x2d, 0x72, 0x2c, 0xf2, 0x14, 0xdc, 0x28, 0xc5, 0x4e, 0xb4, 0x67, 0x45, 0xf7, 0xeb,
	0x13, 0xb7, 0x2c, 0x03, 0xcb, 0x8e, 0xac, 0xf6, 0x11, 0x04, 0x2e, 0x95, 0x98, 0x71, 0xa6, 0xc3,
	0x7b, 0x36, 0x1d, 0x70, 0xd0, 0x05, 0x67, 0x7a, 0x3f, 0x85, 0xfb, 0x1f, 0x7d, 0x09, 0xcd, 0x03,
	0xf2, 0x16, 0x67, 0xfe, 0xef, 0xcd, 0x7c, 0x92, 0x6f, 0xa1, 0x39, 0xa5, 0x79, 0x89, 0xfe, 0xad,
	0xdb, 0xf7, 0x59, 0xac, 0xb0, 0x88, 0x1c, 0xf1, 0xe9, 0xc6, 0x69, 0x63, 0xff, 0x12, 0x60, 0xf1,
	0x3a, 0xad, 0x70, 0xfd, 0x66, 0xd9, 0x75, 0xd7, 0xbb, 0x9a, 0x51, 0xa1, 0x8c, 0xa3, 0x34, 0xe2,
	0xba, 0xdf, 0x4f, 0xb0, 0xf3, 0xc1, 0x54, 0xaf, 0xb0, 0xdd, 0xad, 0xdb, 0xb6, 0xea, 0x06, 0x67,
	0x70, 0x77, 0xc5, 0x08, 0x5d, 0xcb, 0x62, 0x00, 0x7b, 0x2b, 0x47, 0xe4, 0x5a, 0x26, 0x3f, 0x40,
	0x50, 0x9b, 0x85, 0x6b, 0x49, 0x9f, 0x41, 0xef, 0x6a, 0xe7, 0x5f, 0x47, 0xdf, 0xff, 0x67, 0x03,
	0x60, 0x31, 0x56, 0xe4, 0x05, 0xf4, 0x06, 0x12, 0xa9, 0xc6, 0x1a, 0x16, 0x7e, 0x6c, 0xfa, 0xf6,
	0xef, 0xcd, 0xdf, 0x3a, 0xbb, 0x8e, 0x50, 0x15, 0x82, 0x2b, 0x3c, 0xfc, 0xcc, 0xf8, 0x3c, 0xc7,
	0x1c, 0x6f, 0xec, 0x33, 0x80, 0xf6, 0x4b, 0xd4, 0x37, 0x34, 0x39, 0x87, 0xee, 0x90, 0xea, 0xf4,
	0xaf, 0x9b, 0xe7, 0x32, 0x2c, 0x6f, 0x98, 0x4b, 0xb2, 0x65, 0x03, 0x27, 0xff, 0x05, 0x00, 0x00,
	0xff, 0xff, 0x98, 0xf8, 0xb9, 0x46, 0x87, 0x0a, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

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
	cc *grpc.ClientConn
}

func NewDeploymentClient(cc *grpc.ClientConn) DeploymentClient {
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
