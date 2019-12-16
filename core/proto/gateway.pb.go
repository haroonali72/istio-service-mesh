// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gateway.proto

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

type Mode int32

const (
	Mode_PASSTHROUGH      Mode = 0
	Mode_SIMPLE           Mode = 1
	Mode_MUTUAL           Mode = 2
	Mode_AUTO_PASSTHROUGH Mode = 3
	Mode_ISTIO_MUTUAL     Mode = 4
)

var Mode_name = map[int32]string{
	0: "PASSTHROUGH",
	1: "SIMPLE",
	2: "MUTUAL",
	3: "AUTO_PASSTHROUGH",
	4: "ISTIO_MUTUAL",
}

var Mode_value = map[string]int32{
	"PASSTHROUGH":      0,
	"SIMPLE":           1,
	"MUTUAL":           2,
	"AUTO_PASSTHROUGH": 3,
	"ISTIO_MUTUAL":     4,
}

func (x Mode) String() string {
	return proto.EnumName(Mode_name, int32(x))
}

func (Mode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{0}
}

type ProtocolVersion int32

const (
	ProtocolVersion_TLS_AUTO ProtocolVersion = 0
	ProtocolVersion_TLSV1_0  ProtocolVersion = 1
	ProtocolVersion_TLSV1_1  ProtocolVersion = 2
	ProtocolVersion_TLSV1_2  ProtocolVersion = 3
	ProtocolVersion_TLSV1_3  ProtocolVersion = 4
)

var ProtocolVersion_name = map[int32]string{
	0: "TLS_AUTO",
	1: "TLSV1_0",
	2: "TLSV1_1",
	3: "TLSV1_2",
	4: "TLSV1_3",
}

var ProtocolVersion_value = map[string]int32{
	"TLS_AUTO": 0,
	"TLSV1_0":  1,
	"TLSV1_1":  2,
	"TLSV1_2":  3,
	"TLSV1_3":  4,
}

func (x ProtocolVersion) String() string {
	return proto.EnumName(ProtocolVersion_name, int32(x))
}

func (ProtocolVersion) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{1}
}

type Protocols int32

const (
	Protocols_HTTP  Protocols = 0
	Protocols_HTTPS Protocols = 1
	Protocols_GRPC  Protocols = 2
	Protocols_HTTP2 Protocols = 3
	Protocols_MONGO Protocols = 4
	Protocols_TCP   Protocols = 5
	Protocols_TLS   Protocols = 6
)

var Protocols_name = map[int32]string{
	0: "HTTP",
	1: "HTTPS",
	2: "GRPC",
	3: "HTTP2",
	4: "MONGO",
	5: "TCP",
	6: "TLS",
}

var Protocols_value = map[string]int32{
	"HTTP":  0,
	"HTTPS": 1,
	"GRPC":  2,
	"HTTP2": 3,
	"MONGO": 4,
	"TCP":   5,
	"TLS":   6,
}

func (x Protocols) String() string {
	return proto.EnumName(Protocols_name, int32(x))
}

func (Protocols) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{2}
}

type NameRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NameRequest) Reset()         { *m = NameRequest{} }
func (m *NameRequest) String() string { return proto.CompactTextString(m) }
func (*NameRequest) ProtoMessage()    {}
func (*NameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{0}
}

func (m *NameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NameRequest.Unmarshal(m, b)
}
func (m *NameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NameRequest.Marshal(b, m, deterministic)
}
func (m *NameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NameRequest.Merge(m, src)
}
func (m *NameRequest) XXX_Size() int {
	return xxx_messageInfo_NameRequest.Size(m)
}
func (m *NameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NameRequest proto.InternalMessageInfo

func (m *NameRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GatewayService struct {
	ProjectId            string                    `protobuf:"bytes,1,opt,name=ProjectId,proto3" json:"ProjectId,omitempty"`
	ServiceId            string                    `protobuf:"bytes,2,opt,name=serviceId,proto3" json:"serviceId,omitempty"`
	Name                 string                    `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Version              string                    `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	Type                 string                    `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	SubType              string                    `protobuf:"bytes,6,opt,name=sub_type,json=subType,proto3" json:"sub_type,omitempty"`
	Namespace            string                    `protobuf:"bytes,7,opt,name=namespace,proto3" json:"namespace,omitempty"`
	CompanyId            string                    `protobuf:"bytes,8,opt,name=companyId,proto3" json:"companyId,omitempty"`
	ServiceAttributes    *GatewayServiceAttributes `protobuf:"bytes,9,opt,name=serviceAttributes,proto3" json:"serviceAttributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *GatewayService) Reset()         { *m = GatewayService{} }
func (m *GatewayService) String() string { return proto.CompactTextString(m) }
func (*GatewayService) ProtoMessage()    {}
func (*GatewayService) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{1}
}

func (m *GatewayService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GatewayService.Unmarshal(m, b)
}
func (m *GatewayService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GatewayService.Marshal(b, m, deterministic)
}
func (m *GatewayService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GatewayService.Merge(m, src)
}
func (m *GatewayService) XXX_Size() int {
	return xxx_messageInfo_GatewayService.Size(m)
}
func (m *GatewayService) XXX_DiscardUnknown() {
	xxx_messageInfo_GatewayService.DiscardUnknown(m)
}

var xxx_messageInfo_GatewayService proto.InternalMessageInfo

func (m *GatewayService) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *GatewayService) GetServiceId() string {
	if m != nil {
		return m.ServiceId
	}
	return ""
}

func (m *GatewayService) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GatewayService) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *GatewayService) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *GatewayService) GetSubType() string {
	if m != nil {
		return m.SubType
	}
	return ""
}

func (m *GatewayService) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *GatewayService) GetCompanyId() string {
	if m != nil {
		return m.CompanyId
	}
	return ""
}

func (m *GatewayService) GetServiceAttributes() *GatewayServiceAttributes {
	if m != nil {
		return m.ServiceAttributes
	}
	return nil
}

type GatewayServiceResponse struct {
	Error                string          `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Resp                 *GatewayService `protobuf:"bytes,2,opt,name=resp,proto3" json:"resp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GatewayServiceResponse) Reset()         { *m = GatewayServiceResponse{} }
func (m *GatewayServiceResponse) String() string { return proto.CompactTextString(m) }
func (*GatewayServiceResponse) ProtoMessage()    {}
func (*GatewayServiceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{2}
}

func (m *GatewayServiceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GatewayServiceResponse.Unmarshal(m, b)
}
func (m *GatewayServiceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GatewayServiceResponse.Marshal(b, m, deterministic)
}
func (m *GatewayServiceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GatewayServiceResponse.Merge(m, src)
}
func (m *GatewayServiceResponse) XXX_Size() int {
	return xxx_messageInfo_GatewayServiceResponse.Size(m)
}
func (m *GatewayServiceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GatewayServiceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GatewayServiceResponse proto.InternalMessageInfo

func (m *GatewayServiceResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *GatewayServiceResponse) GetResp() *GatewayService {
	if m != nil {
		return m.Resp
	}
	return nil
}

type GatewayServiceAttributes struct {
	Selectors            map[string]string `protobuf:"bytes,1,rep,name=selectors,proto3" json:"selectors,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Servers              []*Server         `protobuf:"bytes,2,rep,name=servers,proto3" json:"servers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GatewayServiceAttributes) Reset()         { *m = GatewayServiceAttributes{} }
func (m *GatewayServiceAttributes) String() string { return proto.CompactTextString(m) }
func (*GatewayServiceAttributes) ProtoMessage()    {}
func (*GatewayServiceAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{3}
}

func (m *GatewayServiceAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GatewayServiceAttributes.Unmarshal(m, b)
}
func (m *GatewayServiceAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GatewayServiceAttributes.Marshal(b, m, deterministic)
}
func (m *GatewayServiceAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GatewayServiceAttributes.Merge(m, src)
}
func (m *GatewayServiceAttributes) XXX_Size() int {
	return xxx_messageInfo_GatewayServiceAttributes.Size(m)
}
func (m *GatewayServiceAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_GatewayServiceAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_GatewayServiceAttributes proto.InternalMessageInfo

func (m *GatewayServiceAttributes) GetSelectors() map[string]string {
	if m != nil {
		return m.Selectors
	}
	return nil
}

func (m *GatewayServiceAttributes) GetServers() []*Server {
	if m != nil {
		return m.Servers
	}
	return nil
}

type Server struct {
	Port                 *Port      `protobuf:"bytes,1,opt,name=port,proto3" json:"port,omitempty"`
	Hosts                []string   `protobuf:"bytes,2,rep,name=hosts,proto3" json:"hosts,omitempty"`
	Tls                  *TlsConfig `protobuf:"bytes,3,opt,name=tls,proto3" json:"tls,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Server) Reset()         { *m = Server{} }
func (m *Server) String() string { return proto.CompactTextString(m) }
func (*Server) ProtoMessage()    {}
func (*Server) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{4}
}

func (m *Server) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Server.Unmarshal(m, b)
}
func (m *Server) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Server.Marshal(b, m, deterministic)
}
func (m *Server) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Server.Merge(m, src)
}
func (m *Server) XXX_Size() int {
	return xxx_messageInfo_Server.Size(m)
}
func (m *Server) XXX_DiscardUnknown() {
	xxx_messageInfo_Server.DiscardUnknown(m)
}

var xxx_messageInfo_Server proto.InternalMessageInfo

func (m *Server) GetPort() *Port {
	if m != nil {
		return m.Port
	}
	return nil
}

func (m *Server) GetHosts() []string {
	if m != nil {
		return m.Hosts
	}
	return nil
}

func (m *Server) GetTls() *TlsConfig {
	if m != nil {
		return m.Tls
	}
	return nil
}

type Port struct {
	Name                 string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Nummber              uint32    `protobuf:"varint,2,opt,name=nummber,proto3" json:"nummber,omitempty"`
	Protocol             Protocols `protobuf:"varint,3,opt,name=protocol,proto3,enum=proto.Protocols" json:"protocol,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Port) Reset()         { *m = Port{} }
func (m *Port) String() string { return proto.CompactTextString(m) }
func (*Port) ProtoMessage()    {}
func (*Port) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{5}
}

func (m *Port) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Port.Unmarshal(m, b)
}
func (m *Port) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Port.Marshal(b, m, deterministic)
}
func (m *Port) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Port.Merge(m, src)
}
func (m *Port) XXX_Size() int {
	return xxx_messageInfo_Port.Size(m)
}
func (m *Port) XXX_DiscardUnknown() {
	xxx_messageInfo_Port.DiscardUnknown(m)
}

var xxx_messageInfo_Port proto.InternalMessageInfo

func (m *Port) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Port) GetNummber() uint32 {
	if m != nil {
		return m.Nummber
	}
	return 0
}

func (m *Port) GetProtocol() Protocols {
	if m != nil {
		return m.Protocol
	}
	return Protocols_HTTP
}

type TlsConfig struct {
	HttpsRedirect        bool            `protobuf:"varint,1,opt,name=httpsRedirect,proto3" json:"httpsRedirect,omitempty"`
	Mode                 Mode            `protobuf:"varint,2,opt,name=mode,proto3,enum=proto.Mode" json:"mode,omitempty"`
	ServerCertificate    string          `protobuf:"bytes,3,opt,name=serverCertificate,proto3" json:"serverCertificate,omitempty"`
	PrivateKey           string          `protobuf:"bytes,4,opt,name=privateKey,proto3" json:"privateKey,omitempty"`
	CaCertificate        string          `protobuf:"bytes,5,opt,name=caCertificate,proto3" json:"caCertificate,omitempty"`
	SubjectAltName       []string        `protobuf:"bytes,6,rep,name=subjectAltName,proto3" json:"subjectAltName,omitempty"`
	MinProtocolVersion   ProtocolVersion `protobuf:"varint,7,opt,name=minProtocolVersion,proto3,enum=proto.ProtocolVersion" json:"minProtocolVersion,omitempty"`
	MaxProtocolVersion   ProtocolVersion `protobuf:"varint,8,opt,name=maxProtocolVersion,proto3,enum=proto.ProtocolVersion" json:"maxProtocolVersion,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *TlsConfig) Reset()         { *m = TlsConfig{} }
func (m *TlsConfig) String() string { return proto.CompactTextString(m) }
func (*TlsConfig) ProtoMessage()    {}
func (*TlsConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{6}
}

func (m *TlsConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TlsConfig.Unmarshal(m, b)
}
func (m *TlsConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TlsConfig.Marshal(b, m, deterministic)
}
func (m *TlsConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TlsConfig.Merge(m, src)
}
func (m *TlsConfig) XXX_Size() int {
	return xxx_messageInfo_TlsConfig.Size(m)
}
func (m *TlsConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TlsConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TlsConfig proto.InternalMessageInfo

func (m *TlsConfig) GetHttpsRedirect() bool {
	if m != nil {
		return m.HttpsRedirect
	}
	return false
}

func (m *TlsConfig) GetMode() Mode {
	if m != nil {
		return m.Mode
	}
	return Mode_PASSTHROUGH
}

func (m *TlsConfig) GetServerCertificate() string {
	if m != nil {
		return m.ServerCertificate
	}
	return ""
}

func (m *TlsConfig) GetPrivateKey() string {
	if m != nil {
		return m.PrivateKey
	}
	return ""
}

func (m *TlsConfig) GetCaCertificate() string {
	if m != nil {
		return m.CaCertificate
	}
	return ""
}

func (m *TlsConfig) GetSubjectAltName() []string {
	if m != nil {
		return m.SubjectAltName
	}
	return nil
}

func (m *TlsConfig) GetMinProtocolVersion() ProtocolVersion {
	if m != nil {
		return m.MinProtocolVersion
	}
	return ProtocolVersion_TLS_AUTO
}

func (m *TlsConfig) GetMaxProtocolVersion() ProtocolVersion {
	if m != nil {
		return m.MaxProtocolVersion
	}
	return ProtocolVersion_TLS_AUTO
}

type ServiceResponse struct {
	Error                string         `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Status               *ServiceStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ServiceResponse) Reset()         { *m = ServiceResponse{} }
func (m *ServiceResponse) String() string { return proto.CompactTextString(m) }
func (*ServiceResponse) ProtoMessage()    {}
func (*ServiceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{7}
}

func (m *ServiceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceResponse.Unmarshal(m, b)
}
func (m *ServiceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceResponse.Marshal(b, m, deterministic)
}
func (m *ServiceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceResponse.Merge(m, src)
}
func (m *ServiceResponse) XXX_Size() int {
	return xxx_messageInfo_ServiceResponse.Size(m)
}
func (m *ServiceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceResponse proto.InternalMessageInfo

func (m *ServiceResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ServiceResponse) GetStatus() *ServiceStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ServiceStatus struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ServiceId            string   `protobuf:"bytes,2,opt,name=serviceId,proto3" json:"serviceId,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	StatusIndividual     []string `protobuf:"bytes,4,rep,name=statusIndividual,proto3" json:"statusIndividual,omitempty"`
	Status               string   `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	Reason               string   `protobuf:"bytes,6,opt,name=reason,proto3" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceStatus) Reset()         { *m = ServiceStatus{} }
func (m *ServiceStatus) String() string { return proto.CompactTextString(m) }
func (*ServiceStatus) ProtoMessage()    {}
func (*ServiceStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{8}
}

func (m *ServiceStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceStatus.Unmarshal(m, b)
}
func (m *ServiceStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceStatus.Marshal(b, m, deterministic)
}
func (m *ServiceStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceStatus.Merge(m, src)
}
func (m *ServiceStatus) XXX_Size() int {
	return xxx_messageInfo_ServiceStatus.Size(m)
}
func (m *ServiceStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceStatus proto.InternalMessageInfo

func (m *ServiceStatus) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ServiceStatus) GetServiceId() string {
	if m != nil {
		return m.ServiceId
	}
	return ""
}

func (m *ServiceStatus) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ServiceStatus) GetStatusIndividual() []string {
	if m != nil {
		return m.StatusIndividual
	}
	return nil
}

func (m *ServiceStatus) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *ServiceStatus) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func init() {
	proto.RegisterEnum("proto.Mode", Mode_name, Mode_value)
	proto.RegisterEnum("proto.ProtocolVersion", ProtocolVersion_name, ProtocolVersion_value)
	proto.RegisterEnum("proto.Protocols", Protocols_name, Protocols_value)
	proto.RegisterType((*NameRequest)(nil), "proto.NameRequest")
	proto.RegisterType((*GatewayService)(nil), "proto.GatewayService")
	proto.RegisterType((*GatewayServiceResponse)(nil), "proto.GatewayServiceResponse")
	proto.RegisterType((*GatewayServiceAttributes)(nil), "proto.GatewayServiceAttributes")
	proto.RegisterMapType((map[string]string)(nil), "proto.GatewayServiceAttributes.SelectorsEntry")
	proto.RegisterType((*Server)(nil), "proto.Server")
	proto.RegisterType((*Port)(nil), "proto.Port")
	proto.RegisterType((*TlsConfig)(nil), "proto.TlsConfig")
	proto.RegisterType((*ServiceResponse)(nil), "proto.ServiceResponse")
	proto.RegisterType((*ServiceStatus)(nil), "proto.ServiceStatus")
}

func init() { proto.RegisterFile("gateway.proto", fileDescriptor_f1a937782ebbded5) }

var fileDescriptor_f1a937782ebbded5 = []byte{
	// 883 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0x6d, 0x6f, 0xe3, 0x44,
	0x10, 0xae, 0x1d, 0xe7, 0x6d, 0xd2, 0xa4, 0xcb, 0xaa, 0x54, 0xe6, 0x74, 0xe2, 0x8a, 0x85, 0xa0,
	0x54, 0x55, 0xc5, 0xe5, 0xbe, 0x20, 0x04, 0xe2, 0xa2, 0x70, 0xa4, 0x11, 0xc9, 0xc5, 0x5a, 0x3b,
	0x87, 0xf8, 0x14, 0x39, 0xce, 0xde, 0xd5, 0x90, 0xd8, 0x61, 0x77, 0x1d, 0xc8, 0x8f, 0xe2, 0x1b,
	0x3f, 0x03, 0x89, 0x7f, 0x84, 0xd0, 0xbe, 0xd8, 0x69, 0xda, 0xa2, 0x3b, 0xf5, 0x3e, 0x65, 0xe7,
	0x79, 0x66, 0x9e, 0x9d, 0x19, 0xcf, 0x6c, 0xa0, 0xfd, 0x26, 0x12, 0xf4, 0xf7, 0x68, 0x7b, 0xb9,
	0x66, 0x99, 0xc8, 0x70, 0x55, 0xfd, 0x78, 0x9f, 0x40, 0xeb, 0x65, 0xb4, 0xa2, 0x84, 0xfe, 0x96,
	0x53, 0x2e, 0x30, 0x06, 0x27, 0x8d, 0x56, 0xd4, 0xb5, 0x4e, 0xad, 0xb3, 0x26, 0x51, 0x67, 0xef,
	0x2f, 0x1b, 0x3a, 0x03, 0x1d, 0x1b, 0x50, 0xb6, 0x49, 0x62, 0x8a, 0x1f, 0x43, 0xd3, 0x67, 0xd9,
	0x2f, 0x34, 0x16, 0xc3, 0x85, 0xf1, 0xdd, 0x01, 0x92, 0xe5, 0xda, 0x71, 0xb8, 0x70, 0x6d, 0xcd,
	0x96, 0x40, 0x79, 0x45, 0x65, 0x77, 0x05, 0x76, 0xa1, 0xbe, 0xa1, 0x8c, 0x27, 0x59, 0xea, 0x3a,
	0x0a, 0x2e, 0x4c, 0xe9, 0x2d, 0xb6, 0x6b, 0xea, 0x56, 0xb5, 0xb7, 0x3c, 0xe3, 0x8f, 0xa0, 0xc1,
	0xf3, 0xf9, 0x4c, 0xe1, 0x35, 0xed, 0xce, 0xf3, 0x79, 0x28, 0xa9, 0xc7, 0xd0, 0x94, 0x82, 0x7c,
	0x1d, 0xc5, 0xd4, 0xad, 0xeb, 0xab, 0x4b, 0x40, 0xb2, 0x71, 0xb6, 0x5a, 0x47, 0xe9, 0x76, 0xb8,
	0x70, 0x1b, 0x9a, 0x2d, 0x01, 0x3c, 0x86, 0x0f, 0x4c, 0x96, 0x3d, 0x21, 0x58, 0x32, 0xcf, 0x05,
	0xe5, 0x6e, 0xf3, 0xd4, 0x3a, 0x6b, 0x75, 0x9f, 0xe8, 0xa6, 0x5d, 0xee, 0xb7, 0x61, 0xe7, 0x46,
	0xee, 0x46, 0x7a, 0x3f, 0xc3, 0xc9, 0xbe, 0x3b, 0xa1, 0x7c, 0x9d, 0xa5, 0x9c, 0xe2, 0x63, 0xa8,
	0x52, 0xc6, 0x32, 0x66, 0x3a, 0xa7, 0x0d, 0xfc, 0x05, 0x38, 0x8c, 0xf2, 0xb5, 0x6a, 0x58, 0xab,
	0xfb, 0xe1, 0xbd, 0x37, 0x12, 0xe5, 0xe2, 0xfd, 0x63, 0x81, 0xfb, 0x7f, 0xa9, 0xe0, 0x91, 0xec,
	0xfe, 0x92, 0xc6, 0x22, 0x63, 0xdc, 0xb5, 0x4e, 0x2b, 0x67, 0xad, 0xee, 0xe5, 0x5b, 0xd2, 0xbf,
	0x0c, 0x8a, 0x80, 0x17, 0xa9, 0x60, 0x5b, 0xb2, 0x13, 0xc0, 0x9f, 0x43, 0x5d, 0x96, 0x46, 0x19,
	0x77, 0x6d, 0xa5, 0xd5, 0x36, 0x5a, 0x81, 0x42, 0x49, 0xc1, 0x3e, 0xfa, 0x06, 0x3a, 0xfb, 0x2a,
	0x18, 0x41, 0xe5, 0x57, 0xba, 0x35, 0x45, 0xca, 0xa3, 0x2c, 0x7c, 0x13, 0x2d, 0x73, 0x6a, 0x86,
	0x42, 0x1b, 0x5f, 0xdb, 0x5f, 0x59, 0x5e, 0x0c, 0x35, 0x2d, 0x88, 0x9f, 0x80, 0xb3, 0xce, 0x98,
	0x50, 0x61, 0xad, 0x6e, 0xcb, 0xdc, 0xe6, 0x67, 0x4c, 0x10, 0x45, 0x48, 0x91, 0xeb, 0x8c, 0x0b,
	0x9d, 0x4f, 0x93, 0x68, 0x03, 0x7b, 0x50, 0x11, 0x4b, 0xae, 0x86, 0xaa, 0xd5, 0x45, 0x26, 0x2a,
	0x5c, 0xf2, 0x7e, 0x96, 0xbe, 0x4e, 0xde, 0x10, 0x49, 0x7a, 0x73, 0x70, 0xa4, 0xce, 0x7d, 0x43,
	0x2e, 0x27, 0x30, 0xcd, 0x57, 0xab, 0x39, 0x65, 0x2a, 0xb9, 0x36, 0x29, 0x4c, 0x7c, 0x01, 0x0d,
	0xa5, 0x16, 0x67, 0x4b, 0x25, 0xdf, 0x29, 0xe5, 0x7d, 0x03, 0x73, 0x52, 0x7a, 0x78, 0xff, 0xda,
	0xd0, 0x2c, 0xaf, 0xc5, 0x9f, 0x42, 0xfb, 0x5a, 0x88, 0x35, 0x27, 0x74, 0x91, 0x30, 0x1a, 0xeb,
	0xaa, 0x1a, 0x64, 0x1f, 0x94, 0x25, 0xaf, 0xb2, 0x85, 0xee, 0x4a, 0xa7, 0x2c, 0x79, 0x9c, 0x2d,
	0x28, 0x51, 0x04, 0xbe, 0xd0, 0x93, 0x49, 0x59, 0x9f, 0x32, 0x91, 0xbc, 0x4e, 0xe2, 0x48, 0x14,
	0xfb, 0x73, 0x97, 0xc0, 0x1f, 0x03, 0xac, 0x59, 0xb2, 0x89, 0x04, 0xfd, 0x91, 0x6e, 0xcd, 0x3e,
	0xdd, 0x40, 0x64, 0x52, 0x71, 0x74, 0x53, 0x49, 0xef, 0xd6, 0x3e, 0x88, 0x3f, 0x83, 0x0e, 0xcf,
	0xe7, 0x72, 0xa3, 0x7b, 0x4b, 0x21, 0x9f, 0x08, 0xb7, 0xa6, 0xfa, 0x7d, 0x0b, 0xc5, 0x3f, 0x00,
	0x5e, 0x25, 0x69, 0xd1, 0x8a, 0x57, 0x66, 0x8b, 0xeb, 0xaa, 0x94, 0x93, 0x5b, 0x8d, 0x32, 0x2c,
	0xb9, 0x27, 0x42, 0xe9, 0x44, 0x7f, 0xdc, 0xd6, 0x69, 0xbc, 0x45, 0xe7, 0x4e, 0x84, 0x37, 0x85,
	0xa3, 0x77, 0xdb, 0xb7, 0x0b, 0xa8, 0x71, 0x11, 0x89, 0x9c, 0x9b, 0x8d, 0x3b, 0xbe, 0x31, 0xd8,
	0x49, 0x4c, 0x03, 0xc5, 0x11, 0xe3, 0xe3, 0xfd, 0x69, 0x41, 0x7b, 0x8f, 0xc1, 0x1d, 0xb0, 0x93,
	0xe2, 0xf1, 0xb3, 0x93, 0x87, 0xbc, 0x7a, 0xe7, 0x80, 0xb4, 0xfa, 0x30, 0x5d, 0x24, 0x9b, 0x64,
	0x91, 0x47, 0x4b, 0xd7, 0x51, 0x4d, 0xbe, 0x83, 0xe3, 0x93, 0x32, 0x5b, 0xfd, 0xb5, 0x8c, 0x25,
	0x71, 0x46, 0x23, 0x9e, 0xa5, 0xe6, 0x25, 0x34, 0xd6, 0xf9, 0x4f, 0xe0, 0xc8, 0x01, 0xc2, 0x47,
	0xd0, 0xf2, 0x7b, 0x41, 0x10, 0x5e, 0x91, 0xc9, 0x74, 0x70, 0x85, 0x0e, 0x30, 0x40, 0x2d, 0x18,
	0x8e, 0xfd, 0xd1, 0x0b, 0x64, 0xc9, 0xf3, 0x78, 0x1a, 0x4e, 0x7b, 0x23, 0x64, 0xe3, 0x63, 0x40,
	0xbd, 0x69, 0x38, 0x99, 0xdd, 0xf4, 0xae, 0x60, 0x04, 0x87, 0xc3, 0x20, 0x1c, 0x4e, 0x66, 0xc6,
	0xcf, 0x39, 0x0f, 0xe0, 0xe8, 0xf6, 0xa7, 0x3b, 0x84, 0x46, 0x38, 0x0a, 0x66, 0x32, 0x1c, 0x1d,
	0xe0, 0x16, 0xd4, 0xc3, 0x51, 0xf0, 0xea, 0xe9, 0xec, 0x4b, 0x64, 0xed, 0x8c, 0xa7, 0xc8, 0xde,
	0x19, 0x5d, 0x54, 0xd9, 0x19, 0xcf, 0x90, 0x73, 0x4e, 0xd4, 0xff, 0x89, 0x5e, 0x26, 0xdc, 0x00,
	0xe7, 0x2a, 0x0c, 0x7d, 0x74, 0x80, 0x9b, 0x50, 0x95, 0xa7, 0x00, 0x59, 0x12, 0x1c, 0x10, 0xbf,
	0x8f, 0xec, 0x02, 0x94, 0x1a, 0x4d, 0xa8, 0x8e, 0x27, 0x2f, 0x07, 0x13, 0xe4, 0xe0, 0x3a, 0x54,
	0xc2, 0xbe, 0x8f, 0xaa, 0xea, 0x30, 0x0a, 0x50, 0xad, 0xfb, 0xb7, 0x0d, 0x75, 0xf3, 0xe0, 0xe1,
	0xe7, 0xd0, 0xee, 0x33, 0x1a, 0x09, 0x5a, 0x00, 0xf7, 0x3f, 0xaf, 0x8f, 0x4e, 0xf6, 0x67, 0xa0,
	0x98, 0x20, 0xef, 0x00, 0x7f, 0x0b, 0x30, 0xa0, 0xe2, 0xc1, 0xe1, 0xcf, 0xa1, 0xfd, 0x3d, 0x5d,
	0xd2, 0xf7, 0x48, 0xe0, 0x3b, 0x38, 0xf4, 0x23, 0x11, 0x5f, 0xbf, 0x4f, 0x05, 0x7e, 0xfe, 0xe0,
	0x0a, 0xe6, 0x35, 0x45, 0x3c, 0xfb, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xe4, 0xeb, 0x44, 0xcf, 0x47,
	0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GatewayClient is the client API for Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GatewayClient interface {
	CreateGateway(ctx context.Context, in *GatewayService, opts ...grpc.CallOption) (*ServiceResponse, error)
	GetGateway(ctx context.Context, in *GatewayService, opts ...grpc.CallOption) (*ServiceResponse, error)
	DeleteGateway(ctx context.Context, in *GatewayService, opts ...grpc.CallOption) (*ServiceResponse, error)
	PatchGateway(ctx context.Context, in *GatewayService, opts ...grpc.CallOption) (*ServiceResponse, error)
	PutGateway(ctx context.Context, in *GatewayService, opts ...grpc.CallOption) (*ServiceResponse, error)
}

type gatewayClient struct {
	cc *grpc.ClientConn
}

func NewGatewayClient(cc *grpc.ClientConn) GatewayClient {
	return &gatewayClient{cc}
}

func (c *gatewayClient) CreateGateway(ctx context.Context, in *GatewayService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.Gateway/CreateGateway", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) GetGateway(ctx context.Context, in *GatewayService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.Gateway/GetGateway", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) DeleteGateway(ctx context.Context, in *GatewayService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.Gateway/DeleteGateway", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) PatchGateway(ctx context.Context, in *GatewayService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.Gateway/PatchGateway", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayClient) PutGateway(ctx context.Context, in *GatewayService, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.Gateway/PutGateway", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServer is the server API for Gateway service.
type GatewayServer interface {
	CreateGateway(context.Context, *GatewayService) (*ServiceResponse, error)
	GetGateway(context.Context, *GatewayService) (*ServiceResponse, error)
	DeleteGateway(context.Context, *GatewayService) (*ServiceResponse, error)
	PatchGateway(context.Context, *GatewayService) (*ServiceResponse, error)
	PutGateway(context.Context, *GatewayService) (*ServiceResponse, error)
}

// UnimplementedGatewayServer can be embedded to have forward compatible implementations.
type UnimplementedGatewayServer struct {
}

func (*UnimplementedGatewayServer) CreateGateway(ctx context.Context, req *GatewayService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGateway not implemented")
}
func (*UnimplementedGatewayServer) GetGateway(ctx context.Context, req *GatewayService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGateway not implemented")
}
func (*UnimplementedGatewayServer) DeleteGateway(ctx context.Context, req *GatewayService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGateway not implemented")
}
func (*UnimplementedGatewayServer) PatchGateway(ctx context.Context, req *GatewayService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchGateway not implemented")
}
func (*UnimplementedGatewayServer) PutGateway(ctx context.Context, req *GatewayService) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutGateway not implemented")
}

func RegisterGatewayServer(s *grpc.Server, srv GatewayServer) {
	s.RegisterService(&_Gateway_serviceDesc, srv)
}

func _Gateway_CreateGateway_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GatewayService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).CreateGateway(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Gateway/CreateGateway",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).CreateGateway(ctx, req.(*GatewayService))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_GetGateway_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GatewayService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).GetGateway(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Gateway/GetGateway",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).GetGateway(ctx, req.(*GatewayService))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_DeleteGateway_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GatewayService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).DeleteGateway(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Gateway/DeleteGateway",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).DeleteGateway(ctx, req.(*GatewayService))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_PatchGateway_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GatewayService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).PatchGateway(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Gateway/PatchGateway",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).PatchGateway(ctx, req.(*GatewayService))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gateway_PutGateway_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GatewayService)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServer).PutGateway(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Gateway/PutGateway",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServer).PutGateway(ctx, req.(*GatewayService))
	}
	return interceptor(ctx, in, info, handler)
}

var _Gateway_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Gateway",
	HandlerType: (*GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGateway",
			Handler:    _Gateway_CreateGateway_Handler,
		},
		{
			MethodName: "GetGateway",
			Handler:    _Gateway_GetGateway_Handler,
		},
		{
			MethodName: "DeleteGateway",
			Handler:    _Gateway_DeleteGateway_Handler,
		},
		{
			MethodName: "PatchGateway",
			Handler:    _Gateway_PatchGateway_Handler,
		},
		{
			MethodName: "PutGateway",
			Handler:    _Gateway_PutGateway_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gateway.proto",
}