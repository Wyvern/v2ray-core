// Code generated by protoc-gen-go.
// source: v2ray.com/core/transport/internet/config.proto
// DO NOT EDIT!

/*
Package internet is a generated protocol buffer package.

It is generated from these files:
	v2ray.com/core/transport/internet/config.proto

It has these top-level messages:
	NetworkSettings
	StreamConfig
	ProxyConfig
*/
package internet

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import v2ray_core_common_net "v2ray.com/core/common/net"
import v2ray_core_common_loader "v2ray.com/core/common/loader"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type NetworkSettings struct {
	// Type of network that this settings supports.
	Network v2ray_core_common_net.Network `protobuf:"varint,1,opt,name=network,enum=v2ray.core.common.net.Network" json:"network,omitempty"`
	// Specific settings.
	Settings *v2ray_core_common_loader.TypedSettings `protobuf:"bytes,2,opt,name=settings" json:"settings,omitempty"`
}

func (m *NetworkSettings) Reset()                    { *m = NetworkSettings{} }
func (m *NetworkSettings) String() string            { return proto.CompactTextString(m) }
func (*NetworkSettings) ProtoMessage()               {}
func (*NetworkSettings) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *NetworkSettings) GetSettings() *v2ray_core_common_loader.TypedSettings {
	if m != nil {
		return m.Settings
	}
	return nil
}

type StreamConfig struct {
	// Effective network.
	Network         v2ray_core_common_net.Network `protobuf:"varint,1,opt,name=network,enum=v2ray.core.common.net.Network" json:"network,omitempty"`
	NetworkSettings []*NetworkSettings            `protobuf:"bytes,2,rep,name=network_settings,json=networkSettings" json:"network_settings,omitempty"`
	// Type of security. Must be a message name of the settings proto.
	SecurityType     string                                    `protobuf:"bytes,3,opt,name=security_type,json=securityType" json:"security_type,omitempty"`
	SecuritySettings []*v2ray_core_common_loader.TypedSettings `protobuf:"bytes,4,rep,name=security_settings,json=securitySettings" json:"security_settings,omitempty"`
}

func (m *StreamConfig) Reset()                    { *m = StreamConfig{} }
func (m *StreamConfig) String() string            { return proto.CompactTextString(m) }
func (*StreamConfig) ProtoMessage()               {}
func (*StreamConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *StreamConfig) GetNetworkSettings() []*NetworkSettings {
	if m != nil {
		return m.NetworkSettings
	}
	return nil
}

func (m *StreamConfig) GetSecuritySettings() []*v2ray_core_common_loader.TypedSettings {
	if m != nil {
		return m.SecuritySettings
	}
	return nil
}

type ProxyConfig struct {
	Tag string `protobuf:"bytes,1,opt,name=tag" json:"tag,omitempty"`
}

func (m *ProxyConfig) Reset()                    { *m = ProxyConfig{} }
func (m *ProxyConfig) String() string            { return proto.CompactTextString(m) }
func (*ProxyConfig) ProtoMessage()               {}
func (*ProxyConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto.RegisterType((*NetworkSettings)(nil), "v2ray.core.transport.internet.NetworkSettings")
	proto.RegisterType((*StreamConfig)(nil), "v2ray.core.transport.internet.StreamConfig")
	proto.RegisterType((*ProxyConfig)(nil), "v2ray.core.transport.internet.ProxyConfig")
}

func init() { proto.RegisterFile("v2ray.com/core/transport/internet/config.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x92, 0xcd, 0x4a, 0xfb, 0x50,
	0x10, 0xc5, 0x49, 0xfb, 0xe7, 0x6f, 0x7b, 0x5b, 0x6d, 0xcd, 0xaa, 0x08, 0x6a, 0xad, 0x8b, 0x66,
	0xe3, 0x04, 0xe2, 0xc6, 0xb5, 0xdd, 0x8b, 0xa4, 0xdd, 0xe8, 0xa6, 0xc4, 0xdb, 0xb1, 0x04, 0xcd,
	0x9d, 0x30, 0x19, 0x3f, 0xf2, 0x16, 0x3e, 0x81, 0xcf, 0x2a, 0xf9, 0xb8, 0xa1, 0x14, 0x2d, 0x82,
	0xbb, 0x61, 0x38, 0xe7, 0xdc, 0x33, 0x3f, 0xae, 0x82, 0xd7, 0x80, 0xa3, 0x1c, 0x34, 0x25, 0xbe,
	0x26, 0x46, 0x5f, 0x38, 0x32, 0x59, 0x4a, 0x2c, 0x7e, 0x6c, 0x04, 0xd9, 0xa0, 0xf8, 0x9a, 0xcc,
	0x63, 0xbc, 0x86, 0x94, 0x49, 0xc8, 0x3d, 0xb6, 0x7a, 0x46, 0x68, 0xb4, 0x60, 0xb5, 0x47, 0xd3,
	0xad, 0x38, 0x4d, 0x49, 0x42, 0xc6, 0x2f, 0x62, 0x0c, 0xca, 0x1b, 0xf1, 0x53, 0x95, 0xf3, 0x93,
	0xf0, 0x99, 0xa2, 0x15, 0xb2, 0x2f, 0x79, 0x8a, 0x95, 0x70, 0xf2, 0xe1, 0xa8, 0xc1, 0x4d, 0x65,
	0x9d, 0xa3, 0x48, 0x6c, 0xd6, 0x99, 0x7b, 0xa5, 0xf6, 0xea, 0xb4, 0x91, 0x33, 0x76, 0xbc, 0x83,
	0xe0, 0x04, 0x36, 0x6a, 0x55, 0x51, 0x60, 0x50, 0xa0, 0x36, 0x86, 0x56, 0xee, 0xce, 0x54, 0x27,
	0xab, 0x53, 0x46, 0xad, 0xb1, 0xe3, 0xf5, 0x82, 0xe9, 0x37, 0xd6, 0xaa, 0x05, 0x2c, 0xf2, 0x14,
	0x57, 0xf6, 0xd1, 0xb0, 0x31, 0x4e, 0x3e, 0x5b, 0xaa, 0x3f, 0x17, 0xc6, 0x28, 0x99, 0x95, 0x68,
	0xfe, 0xd0, 0xe7, 0x4e, 0x0d, 0xeb, 0x71, 0xb9, 0xd1, 0xab, 0xed, 0xf5, 0x02, 0x80, 0x9d, 0xa4,
	0x61, 0x8b, 0x49, 0x38, 0x30, 0x5b, 0x90, 0xce, 0xd5, 0x7e, 0x86, 0xfa, 0x85, 0x63, 0xc9, 0x97,
	0x05, 0xcf, 0x51, 0x7b, 0xec, 0x78, 0xdd, 0xb0, 0x6f, 0x97, 0xc5, 0x75, 0xee, 0x42, 0x1d, 0x36,
	0xa2, 0xa6, 0xc0, 0xbf, 0xb2, 0xc0, 0xaf, 0xc1, 0x0c, 0x6d, 0x82, 0xdd, 0x4c, 0x4e, 0x55, 0xef,
	0x96, 0xe9, 0x3d, 0xaf, 0xf1, 0x0c, 0x55, 0x5b, 0xa2, 0x75, 0x89, 0xa6, 0x1b, 0x16, 0xe3, 0xf5,
	0x85, 0x3a, 0xd3, 0x94, 0xec, 0xbe, 0xf0, 0xbe, 0x63, 0xa7, 0x87, 0xff, 0xe5, 0x57, 0xb8, 0xfc,
	0x0a, 0x00, 0x00, 0xff, 0xff, 0xa6, 0xb6, 0x2f, 0xcf, 0xad, 0x02, 0x00, 0x00,
}
