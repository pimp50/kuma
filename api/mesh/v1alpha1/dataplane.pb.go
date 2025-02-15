// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mesh/v1alpha1/dataplane.proto

package v1alpha1

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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

// Dataplane defines configuration of a side-car proxy.
type Dataplane struct {
	// Networking describes inbound and outbound interfaces of the dataplane.
	Networking *Dataplane_Networking `protobuf:"bytes,1,opt,name=networking,proto3" json:"networking,omitempty"`
	// Configuration for metrics that should be collected and exposed by the
	// dataplane.
	//
	// Settings defined here will override their respective defaults
	// defined at a Mesh level.
	Metrics              *Metrics `protobuf:"bytes,2,opt,name=metrics,proto3" json:"metrics,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Dataplane) Reset()         { *m = Dataplane{} }
func (m *Dataplane) String() string { return proto.CompactTextString(m) }
func (*Dataplane) ProtoMessage()    {}
func (*Dataplane) Descriptor() ([]byte, []int) {
	return fileDescriptor_7608682fd5ea84a4, []int{0}
}

func (m *Dataplane) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dataplane.Unmarshal(m, b)
}
func (m *Dataplane) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dataplane.Marshal(b, m, deterministic)
}
func (m *Dataplane) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dataplane.Merge(m, src)
}
func (m *Dataplane) XXX_Size() int {
	return xxx_messageInfo_Dataplane.Size(m)
}
func (m *Dataplane) XXX_DiscardUnknown() {
	xxx_messageInfo_Dataplane.DiscardUnknown(m)
}

var xxx_messageInfo_Dataplane proto.InternalMessageInfo

func (m *Dataplane) GetNetworking() *Dataplane_Networking {
	if m != nil {
		return m.Networking
	}
	return nil
}

func (m *Dataplane) GetMetrics() *Metrics {
	if m != nil {
		return m.Metrics
	}
	return nil
}

// Networking describes inbound and outbound interfaces of a dataplane.
type Dataplane_Networking struct {
	// Public IP on which the dataplane is accessible in the network.
	// Host names and DNS are not allowed.
	Address string `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	// Gateway describes configuration of gateway of the dataplane.
	Gateway *Dataplane_Networking_Gateway `protobuf:"bytes,3,opt,name=gateway,proto3" json:"gateway,omitempty"`
	// Inbound describes a list of inbound interfaces of the dataplane.
	Inbound []*Dataplane_Networking_Inbound `protobuf:"bytes,1,rep,name=inbound,proto3" json:"inbound,omitempty"`
	// Outbound describes a list of outbound interfaces of the dataplane.
	Outbound []*Dataplane_Networking_Outbound `protobuf:"bytes,2,rep,name=outbound,proto3" json:"outbound,omitempty"`
	// TransparentProxying describes configuration for transparent proxying.
	TransparentProxying  *Dataplane_Networking_TransparentProxying `protobuf:"bytes,4,opt,name=transparent_proxying,json=transparentProxying,proto3" json:"transparent_proxying,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                  `json:"-"`
	XXX_unrecognized     []byte                                    `json:"-"`
	XXX_sizecache        int32                                     `json:"-"`
}

func (m *Dataplane_Networking) Reset()         { *m = Dataplane_Networking{} }
func (m *Dataplane_Networking) String() string { return proto.CompactTextString(m) }
func (*Dataplane_Networking) ProtoMessage()    {}
func (*Dataplane_Networking) Descriptor() ([]byte, []int) {
	return fileDescriptor_7608682fd5ea84a4, []int{0, 0}
}

func (m *Dataplane_Networking) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dataplane_Networking.Unmarshal(m, b)
}
func (m *Dataplane_Networking) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dataplane_Networking.Marshal(b, m, deterministic)
}
func (m *Dataplane_Networking) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dataplane_Networking.Merge(m, src)
}
func (m *Dataplane_Networking) XXX_Size() int {
	return xxx_messageInfo_Dataplane_Networking.Size(m)
}
func (m *Dataplane_Networking) XXX_DiscardUnknown() {
	xxx_messageInfo_Dataplane_Networking.DiscardUnknown(m)
}

var xxx_messageInfo_Dataplane_Networking proto.InternalMessageInfo

func (m *Dataplane_Networking) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Dataplane_Networking) GetGateway() *Dataplane_Networking_Gateway {
	if m != nil {
		return m.Gateway
	}
	return nil
}

func (m *Dataplane_Networking) GetInbound() []*Dataplane_Networking_Inbound {
	if m != nil {
		return m.Inbound
	}
	return nil
}

func (m *Dataplane_Networking) GetOutbound() []*Dataplane_Networking_Outbound {
	if m != nil {
		return m.Outbound
	}
	return nil
}

func (m *Dataplane_Networking) GetTransparentProxying() *Dataplane_Networking_TransparentProxying {
	if m != nil {
		return m.TransparentProxying
	}
	return nil
}

// Inbound describes a service implemented by the dataplane.
type Dataplane_Networking_Inbound struct {
	// DEPRECATED: use networking.address, networking.inbound[].port and
	// networking.inbound[].servicePort Interface describes networking rules
	// for incoming traffic. The value is a string formatted as
	// <DATAPLANE_IP>:<DATAPLANE_PORT>:<WORKLOAD_PORT>, which means
	// that dataplane must listen on <DATAPLANE_IP>:<DATAPLANE_PORT>
	// and must dispatch to 127.0.0.1:<WORKLOAD_PORT>.
	//
	// E.g.,
	// "192.168.0.100:9090:8080" in case of IPv4 or
	// "[2001:db8::1]:7070:6060" in case of IPv6.
	Interface string `protobuf:"bytes,1,opt,name=interface,proto3" json:"interface,omitempty"`
	// Port of the inbound interface that will forward requests to the
	// service.
	Port uint32 `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	// Port of the service that requests will be forwarded to.
	ServicePort uint32 `protobuf:"varint,4,opt,name=servicePort,proto3" json:"servicePort,omitempty"`
	// Address on which inbound listener will be exposed. Defaults to
	// networking.address.
	Address string `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	// Tags associated with an application this dataplane is deployed next to,
	// e.g. service=web, version=1.0.
	// `service` tag is mandatory.
	Tags                 map[string]string `protobuf:"bytes,2,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Dataplane_Networking_Inbound) Reset()         { *m = Dataplane_Networking_Inbound{} }
func (m *Dataplane_Networking_Inbound) String() string { return proto.CompactTextString(m) }
func (*Dataplane_Networking_Inbound) ProtoMessage()    {}
func (*Dataplane_Networking_Inbound) Descriptor() ([]byte, []int) {
	return fileDescriptor_7608682fd5ea84a4, []int{0, 0, 0}
}

func (m *Dataplane_Networking_Inbound) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dataplane_Networking_Inbound.Unmarshal(m, b)
}
func (m *Dataplane_Networking_Inbound) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dataplane_Networking_Inbound.Marshal(b, m, deterministic)
}
func (m *Dataplane_Networking_Inbound) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dataplane_Networking_Inbound.Merge(m, src)
}
func (m *Dataplane_Networking_Inbound) XXX_Size() int {
	return xxx_messageInfo_Dataplane_Networking_Inbound.Size(m)
}
func (m *Dataplane_Networking_Inbound) XXX_DiscardUnknown() {
	xxx_messageInfo_Dataplane_Networking_Inbound.DiscardUnknown(m)
}

var xxx_messageInfo_Dataplane_Networking_Inbound proto.InternalMessageInfo

func (m *Dataplane_Networking_Inbound) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

func (m *Dataplane_Networking_Inbound) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Dataplane_Networking_Inbound) GetServicePort() uint32 {
	if m != nil {
		return m.ServicePort
	}
	return 0
}

func (m *Dataplane_Networking_Inbound) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Dataplane_Networking_Inbound) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

// Outbound describes a service consumed by the dataplane.
type Dataplane_Networking_Outbound struct {
	// DEPRECATED: use networking.address and networking.outbound[].port
	// Interface describes networking rules for outgoing traffic.
	// The value is a string formatted as <DATAPLANE_IP>:<DATAPLANE_PORT>,
	// which means that dataplane must listen on
	// <DATAPLANE_IP>:<DATAPLANE_PORT> and must be dispatch to
	// <SERVICE>:<SERVICE_PORT>.
	//
	// E.g.,
	// "127.0.0.1:9090" in case of IPv4 or
	// "[::1]:8080" in case of IPv6 or
	// ":7070".
	Interface string `protobuf:"bytes,1,opt,name=interface,proto3" json:"interface,omitempty"`
	// Address on which the service will be available to this dataplane.
	// Defaults to 127.0.0.1
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	// Port on which the service will be available to this dataplane.
	Port uint32 `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
	// Service name.
	Service              string   `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Dataplane_Networking_Outbound) Reset()         { *m = Dataplane_Networking_Outbound{} }
func (m *Dataplane_Networking_Outbound) String() string { return proto.CompactTextString(m) }
func (*Dataplane_Networking_Outbound) ProtoMessage()    {}
func (*Dataplane_Networking_Outbound) Descriptor() ([]byte, []int) {
	return fileDescriptor_7608682fd5ea84a4, []int{0, 0, 1}
}

func (m *Dataplane_Networking_Outbound) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dataplane_Networking_Outbound.Unmarshal(m, b)
}
func (m *Dataplane_Networking_Outbound) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dataplane_Networking_Outbound.Marshal(b, m, deterministic)
}
func (m *Dataplane_Networking_Outbound) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dataplane_Networking_Outbound.Merge(m, src)
}
func (m *Dataplane_Networking_Outbound) XXX_Size() int {
	return xxx_messageInfo_Dataplane_Networking_Outbound.Size(m)
}
func (m *Dataplane_Networking_Outbound) XXX_DiscardUnknown() {
	xxx_messageInfo_Dataplane_Networking_Outbound.DiscardUnknown(m)
}

var xxx_messageInfo_Dataplane_Networking_Outbound proto.InternalMessageInfo

func (m *Dataplane_Networking_Outbound) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

func (m *Dataplane_Networking_Outbound) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Dataplane_Networking_Outbound) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Dataplane_Networking_Outbound) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

// Gateway describes a service that ingress should not be proxied.
type Dataplane_Networking_Gateway struct {
	// Tags associated with a gateway (e.g., Kong, Contour, etc) this
	// dataplane is deployed next to, e.g. service=gateway, env=prod.
	// `service` tag is mandatory.
	Tags                 map[string]string `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Dataplane_Networking_Gateway) Reset()         { *m = Dataplane_Networking_Gateway{} }
func (m *Dataplane_Networking_Gateway) String() string { return proto.CompactTextString(m) }
func (*Dataplane_Networking_Gateway) ProtoMessage()    {}
func (*Dataplane_Networking_Gateway) Descriptor() ([]byte, []int) {
	return fileDescriptor_7608682fd5ea84a4, []int{0, 0, 2}
}

func (m *Dataplane_Networking_Gateway) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dataplane_Networking_Gateway.Unmarshal(m, b)
}
func (m *Dataplane_Networking_Gateway) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dataplane_Networking_Gateway.Marshal(b, m, deterministic)
}
func (m *Dataplane_Networking_Gateway) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dataplane_Networking_Gateway.Merge(m, src)
}
func (m *Dataplane_Networking_Gateway) XXX_Size() int {
	return xxx_messageInfo_Dataplane_Networking_Gateway.Size(m)
}
func (m *Dataplane_Networking_Gateway) XXX_DiscardUnknown() {
	xxx_messageInfo_Dataplane_Networking_Gateway.DiscardUnknown(m)
}

var xxx_messageInfo_Dataplane_Networking_Gateway proto.InternalMessageInfo

func (m *Dataplane_Networking_Gateway) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

// TransparentProxying describes configuration for transparent proxying.
type Dataplane_Networking_TransparentProxying struct {
	// Port on which all traffic is being transparently redirected.
	RedirectPort         uint32   `protobuf:"varint,1,opt,name=redirect_port,json=redirectPort,proto3" json:"redirect_port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Dataplane_Networking_TransparentProxying) Reset() {
	*m = Dataplane_Networking_TransparentProxying{}
}
func (m *Dataplane_Networking_TransparentProxying) String() string { return proto.CompactTextString(m) }
func (*Dataplane_Networking_TransparentProxying) ProtoMessage()    {}
func (*Dataplane_Networking_TransparentProxying) Descriptor() ([]byte, []int) {
	return fileDescriptor_7608682fd5ea84a4, []int{0, 0, 3}
}

func (m *Dataplane_Networking_TransparentProxying) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dataplane_Networking_TransparentProxying.Unmarshal(m, b)
}
func (m *Dataplane_Networking_TransparentProxying) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dataplane_Networking_TransparentProxying.Marshal(b, m, deterministic)
}
func (m *Dataplane_Networking_TransparentProxying) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dataplane_Networking_TransparentProxying.Merge(m, src)
}
func (m *Dataplane_Networking_TransparentProxying) XXX_Size() int {
	return xxx_messageInfo_Dataplane_Networking_TransparentProxying.Size(m)
}
func (m *Dataplane_Networking_TransparentProxying) XXX_DiscardUnknown() {
	xxx_messageInfo_Dataplane_Networking_TransparentProxying.DiscardUnknown(m)
}

var xxx_messageInfo_Dataplane_Networking_TransparentProxying proto.InternalMessageInfo

func (m *Dataplane_Networking_TransparentProxying) GetRedirectPort() uint32 {
	if m != nil {
		return m.RedirectPort
	}
	return 0
}

func init() {
	proto.RegisterType((*Dataplane)(nil), "kuma.mesh.v1alpha1.Dataplane")
	proto.RegisterType((*Dataplane_Networking)(nil), "kuma.mesh.v1alpha1.Dataplane.Networking")
	proto.RegisterType((*Dataplane_Networking_Inbound)(nil), "kuma.mesh.v1alpha1.Dataplane.Networking.Inbound")
	proto.RegisterMapType((map[string]string)(nil), "kuma.mesh.v1alpha1.Dataplane.Networking.Inbound.TagsEntry")
	proto.RegisterType((*Dataplane_Networking_Outbound)(nil), "kuma.mesh.v1alpha1.Dataplane.Networking.Outbound")
	proto.RegisterType((*Dataplane_Networking_Gateway)(nil), "kuma.mesh.v1alpha1.Dataplane.Networking.Gateway")
	proto.RegisterMapType((map[string]string)(nil), "kuma.mesh.v1alpha1.Dataplane.Networking.Gateway.TagsEntry")
	proto.RegisterType((*Dataplane_Networking_TransparentProxying)(nil), "kuma.mesh.v1alpha1.Dataplane.Networking.TransparentProxying")
}

func init() { proto.RegisterFile("mesh/v1alpha1/dataplane.proto", fileDescriptor_7608682fd5ea84a4) }

var fileDescriptor_7608682fd5ea84a4 = []byte{
	// 511 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0xdf, 0x6e, 0xd3, 0x3e,
	0x14, 0xc7, 0xe5, 0x24, 0x5d, 0x92, 0xb3, 0x5f, 0xa5, 0x9f, 0xbc, 0x49, 0x44, 0x19, 0x48, 0x85,
	0xab, 0x8a, 0x0b, 0x97, 0x82, 0x10, 0x68, 0xe2, 0x2a, 0x62, 0xe2, 0x8f, 0x34, 0x98, 0xac, 0x5d,
	0x20, 0x6e, 0x26, 0xaf, 0x31, 0x6d, 0xd4, 0xd6, 0x89, 0x1c, 0xb7, 0x23, 0xaf, 0xc0, 0x23, 0x20,
	0x1e, 0x91, 0x27, 0x28, 0x17, 0x43, 0x71, 0xec, 0xb4, 0xd3, 0x26, 0x68, 0x2f, 0xb8, 0x73, 0xce,
	0x39, 0xdf, 0x8f, 0x7d, 0xbe, 0xe7, 0x28, 0xf0, 0x60, 0xce, 0xcb, 0xc9, 0x60, 0x39, 0x64, 0xb3,
	0x62, 0xc2, 0x86, 0x83, 0x94, 0x29, 0x56, 0xcc, 0x98, 0xe0, 0xa4, 0x90, 0xb9, 0xca, 0x31, 0x9e,
	0x2e, 0xe6, 0x8c, 0xd4, 0x35, 0xc4, 0xd6, 0xc4, 0x47, 0x37, 0x25, 0x73, 0xae, 0x64, 0x36, 0x2a,
	0x1b, 0x41, 0x7c, 0x6f, 0xc9, 0x66, 0x59, 0xca, 0x14, 0x1f, 0xd8, 0x43, 0x93, 0x78, 0xf4, 0x2b,
	0x80, 0xf0, 0xb5, 0xa5, 0xe3, 0xb7, 0x00, 0x82, 0xab, 0xab, 0x5c, 0x4e, 0x33, 0x31, 0x8e, 0x50,
	0x0f, 0xf5, 0xf7, 0x9f, 0xf6, 0xc9, 0xed, 0xcb, 0x48, 0x2b, 0x21, 0x1f, 0xda, 0x7a, 0xba, 0xa1,
	0xc5, 0xcf, 0xc1, 0x37, 0x2f, 0x88, 0x1c, 0x8d, 0x39, 0xba, 0x0b, 0x73, 0xda, 0x94, 0x50, 0x5b,
	0x1b, 0xff, 0xf4, 0x01, 0xd6, 0x44, 0x1c, 0x81, 0xcf, 0xd2, 0x54, 0xf2, 0xb2, 0x8c, 0x3a, 0x3d,
	0xd4, 0x0f, 0xa9, 0xfd, 0xc4, 0xef, 0xc1, 0x1f, 0x33, 0xc5, 0xaf, 0x58, 0x15, 0xb9, 0x9a, 0xff,
	0x64, 0xdb, 0x67, 0x92, 0x37, 0x8d, 0x8e, 0x5a, 0x40, 0xcd, 0xca, 0xc4, 0x65, 0xbe, 0x10, 0x69,
	0x84, 0x7a, 0xee, 0x4e, 0xac, 0x77, 0x8d, 0x8e, 0x5a, 0x00, 0x3e, 0x85, 0x20, 0x5f, 0xa8, 0x06,
	0xe6, 0x68, 0xd8, 0x70, 0x6b, 0xd8, 0x47, 0x23, 0xa4, 0x2d, 0x02, 0xe7, 0x70, 0xa8, 0x24, 0x13,
	0x65, 0xc1, 0x24, 0x17, 0xea, 0xa2, 0x90, 0xf9, 0xd7, 0xaa, 0x1e, 0x8d, 0xa7, 0x7b, 0x7e, 0xb5,
	0x35, 0xfa, 0x7c, 0x0d, 0x39, 0x33, 0x0c, 0x7a, 0xa0, 0x6e, 0x07, 0xe3, 0x6f, 0x0e, 0xf8, 0xa6,
	0x29, 0x7c, 0x1f, 0xc2, 0x4c, 0x28, 0x2e, 0xbf, 0xb0, 0x11, 0xd7, 0xcb, 0x10, 0xd2, 0x75, 0x00,
	0x63, 0xf0, 0x8a, 0x5c, 0x2a, 0x6d, 0x7f, 0x97, 0xea, 0x33, 0xee, 0xc1, 0x7e, 0xc9, 0xe5, 0x32,
	0x1b, 0xf1, 0xb3, 0x3a, 0xe5, 0xe9, 0xd4, 0x66, 0xe8, 0x0f, 0x13, 0xfd, 0x04, 0x9e, 0x62, 0xe3,
	0xd2, 0xb8, 0x76, 0xbc, 0xeb, 0x08, 0xc8, 0x39, 0x1b, 0x97, 0x27, 0x42, 0xc9, 0x2a, 0x09, 0x56,
	0x49, 0xe7, 0x3b, 0x72, 0x02, 0x44, 0x35, 0x31, 0x7e, 0x01, 0x61, 0x9b, 0xc4, 0xff, 0x83, 0x3b,
	0xe5, 0x95, 0x69, 0xa7, 0x3e, 0xe2, 0x43, 0xe8, 0x2c, 0xd9, 0x6c, 0xc1, 0xf5, 0xa2, 0x86, 0xb4,
	0xf9, 0x38, 0x76, 0x5e, 0xa2, 0xb8, 0x82, 0xc0, 0xce, 0xe4, 0x2f, 0x66, 0x6c, 0xb4, 0xe5, 0xde,
	0x6c, 0xcb, 0xda, 0xe4, 0x6d, 0xd8, 0xf4, 0x10, 0x7c, 0xe3, 0x49, 0x73, 0x67, 0xe2, 0xaf, 0x12,
	0x4f, 0x3a, 0x13, 0x44, 0x6d, 0x3c, 0xfe, 0x81, 0xc0, 0x37, 0x8b, 0xda, 0x3a, 0x83, 0x76, 0x74,
	0xc6, 0xe8, 0xff, 0x85, 0x33, 0x27, 0x70, 0x70, 0xc7, 0x4a, 0x61, 0x02, 0x5d, 0xc9, 0xd3, 0x4c,
	0xf2, 0x91, 0xba, 0xd0, 0x5d, 0xd7, 0xb0, 0x6e, 0x12, 0xae, 0x92, 0xbd, 0xc7, 0x5e, 0x74, 0x7d,
	0xed, 0xd2, 0xff, 0x6c, 0xbe, 0xde, 0x86, 0x04, 0x3e, 0x07, 0xb6, 0x85, 0xcb, 0x3d, 0xfd, 0x43,
	0x7a, 0xf6, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x05, 0x6a, 0x5b, 0x20, 0xfb, 0x04, 0x00, 0x00,
}
