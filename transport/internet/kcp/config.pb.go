// Code generated by protoc-gen-go.
// source: v2ray.com/core/transport/internet/kcp/config.proto
// DO NOT EDIT!

/*
Package kcp is a generated protocol buffer package.

It is generated from these files:
	v2ray.com/core/transport/internet/kcp/config.proto

It has these top-level messages:
	MTU
	TTI
	UplinkCapacity
	DownlinkCapacity
	WriteBuffer
	ReadBuffer
	ConnectionReuse
	Config
*/
package kcp

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
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

// Maximum Transmission Unit, in bytes.
type MTU struct {
	Value uint32 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *MTU) Reset()                    { *m = MTU{} }
func (m *MTU) String() string            { return proto.CompactTextString(m) }
func (*MTU) ProtoMessage()               {}
func (*MTU) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Transmission Time Interview, in milli-sec.
type TTI struct {
	Value uint32 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *TTI) Reset()                    { *m = TTI{} }
func (m *TTI) String() string            { return proto.CompactTextString(m) }
func (*TTI) ProtoMessage()               {}
func (*TTI) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// Uplink capacity, in MB.
type UplinkCapacity struct {
	Value uint32 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *UplinkCapacity) Reset()                    { *m = UplinkCapacity{} }
func (m *UplinkCapacity) String() string            { return proto.CompactTextString(m) }
func (*UplinkCapacity) ProtoMessage()               {}
func (*UplinkCapacity) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// Downlink capacity, in MB.
type DownlinkCapacity struct {
	Value uint32 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *DownlinkCapacity) Reset()                    { *m = DownlinkCapacity{} }
func (m *DownlinkCapacity) String() string            { return proto.CompactTextString(m) }
func (*DownlinkCapacity) ProtoMessage()               {}
func (*DownlinkCapacity) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type WriteBuffer struct {
	// Buffer size in bytes.
	Size uint32 `protobuf:"varint,1,opt,name=size" json:"size,omitempty"`
}

func (m *WriteBuffer) Reset()                    { *m = WriteBuffer{} }
func (m *WriteBuffer) String() string            { return proto.CompactTextString(m) }
func (*WriteBuffer) ProtoMessage()               {}
func (*WriteBuffer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type ReadBuffer struct {
	// Buffer size in bytes.
	Size uint32 `protobuf:"varint,1,opt,name=size" json:"size,omitempty"`
}

func (m *ReadBuffer) Reset()                    { *m = ReadBuffer{} }
func (m *ReadBuffer) String() string            { return proto.CompactTextString(m) }
func (*ReadBuffer) ProtoMessage()               {}
func (*ReadBuffer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type ConnectionReuse struct {
	Enable bool `protobuf:"varint,1,opt,name=enable" json:"enable,omitempty"`
}

func (m *ConnectionReuse) Reset()                    { *m = ConnectionReuse{} }
func (m *ConnectionReuse) String() string            { return proto.CompactTextString(m) }
func (*ConnectionReuse) ProtoMessage()               {}
func (*ConnectionReuse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type Config struct {
	Mtu              *MTU                                    `protobuf:"bytes,1,opt,name=mtu" json:"mtu,omitempty"`
	Tti              *TTI                                    `protobuf:"bytes,2,opt,name=tti" json:"tti,omitempty"`
	UplinkCapacity   *UplinkCapacity                         `protobuf:"bytes,3,opt,name=uplink_capacity,json=uplinkCapacity" json:"uplink_capacity,omitempty"`
	DownlinkCapacity *DownlinkCapacity                       `protobuf:"bytes,4,opt,name=downlink_capacity,json=downlinkCapacity" json:"downlink_capacity,omitempty"`
	Congestion       bool                                    `protobuf:"varint,5,opt,name=congestion" json:"congestion,omitempty"`
	WriteBuffer      *WriteBuffer                            `protobuf:"bytes,6,opt,name=write_buffer,json=writeBuffer" json:"write_buffer,omitempty"`
	ReadBuffer       *ReadBuffer                             `protobuf:"bytes,7,opt,name=read_buffer,json=readBuffer" json:"read_buffer,omitempty"`
	HeaderConfig     *v2ray_core_common_loader.TypedSettings `protobuf:"bytes,8,opt,name=header_config,json=headerConfig" json:"header_config,omitempty"`
	ConnectionReuse  *ConnectionReuse                        `protobuf:"bytes,9,opt,name=connection_reuse,json=connectionReuse" json:"connection_reuse,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Config) GetMtu() *MTU {
	if m != nil {
		return m.Mtu
	}
	return nil
}

func (m *Config) GetTti() *TTI {
	if m != nil {
		return m.Tti
	}
	return nil
}

func (m *Config) GetUplinkCapacity() *UplinkCapacity {
	if m != nil {
		return m.UplinkCapacity
	}
	return nil
}

func (m *Config) GetDownlinkCapacity() *DownlinkCapacity {
	if m != nil {
		return m.DownlinkCapacity
	}
	return nil
}

func (m *Config) GetWriteBuffer() *WriteBuffer {
	if m != nil {
		return m.WriteBuffer
	}
	return nil
}

func (m *Config) GetReadBuffer() *ReadBuffer {
	if m != nil {
		return m.ReadBuffer
	}
	return nil
}

func (m *Config) GetHeaderConfig() *v2ray_core_common_loader.TypedSettings {
	if m != nil {
		return m.HeaderConfig
	}
	return nil
}

func (m *Config) GetConnectionReuse() *ConnectionReuse {
	if m != nil {
		return m.ConnectionReuse
	}
	return nil
}

func init() {
	proto.RegisterType((*MTU)(nil), "v2ray.core.transport.internet.kcp.MTU")
	proto.RegisterType((*TTI)(nil), "v2ray.core.transport.internet.kcp.TTI")
	proto.RegisterType((*UplinkCapacity)(nil), "v2ray.core.transport.internet.kcp.UplinkCapacity")
	proto.RegisterType((*DownlinkCapacity)(nil), "v2ray.core.transport.internet.kcp.DownlinkCapacity")
	proto.RegisterType((*WriteBuffer)(nil), "v2ray.core.transport.internet.kcp.WriteBuffer")
	proto.RegisterType((*ReadBuffer)(nil), "v2ray.core.transport.internet.kcp.ReadBuffer")
	proto.RegisterType((*ConnectionReuse)(nil), "v2ray.core.transport.internet.kcp.ConnectionReuse")
	proto.RegisterType((*Config)(nil), "v2ray.core.transport.internet.kcp.Config")
}

func init() { proto.RegisterFile("v2ray.com/core/transport/internet/kcp/config.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 473 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x93, 0xd1, 0x6f, 0xd3, 0x30,
	0x10, 0xc6, 0x35, 0xba, 0x96, 0x71, 0xd9, 0xd6, 0x12, 0x21, 0x14, 0x81, 0x84, 0xb6, 0x4a, 0x6c,
	0xe3, 0x01, 0x47, 0x74, 0x2f, 0xf0, 0xda, 0xf2, 0x32, 0x89, 0x21, 0x30, 0xa9, 0x90, 0x26, 0xa1,
	0xe2, 0x3a, 0xd7, 0x62, 0xb5, 0xb1, 0x2d, 0xc7, 0x59, 0x55, 0xfe, 0x24, 0xfe, 0x4a, 0x64, 0xa7,
	0x5d, 0xdb, 0x48, 0xdb, 0xf2, 0x96, 0xf8, 0xbe, 0xfb, 0xd9, 0xfa, 0xee, 0x3b, 0xe8, 0xdd, 0xf6,
	0x0c, 0x5b, 0x12, 0xae, 0xb2, 0x98, 0x2b, 0x83, 0xb1, 0x35, 0x4c, 0xe6, 0x5a, 0x19, 0x1b, 0x0b,
	0x69, 0xd1, 0x48, 0xb4, 0xf1, 0x8c, 0xeb, 0x98, 0x2b, 0x39, 0x11, 0x53, 0xa2, 0x8d, 0xb2, 0x2a,
	0x3c, 0x5d, 0xf7, 0x18, 0x24, 0x77, 0x7a, 0xb2, 0xd6, 0x93, 0x19, 0xd7, 0xaf, 0xce, 0x2b, 0x58,
	0xae, 0xb2, 0x4c, 0xc9, 0x78, 0xae, 0x58, 0x8a, 0x26, 0xb6, 0x4b, 0x8d, 0x25, 0xab, 0xfb, 0x1a,
	0x1a, 0xd7, 0xc9, 0x30, 0x7c, 0x01, 0xcd, 0x5b, 0x36, 0x2f, 0x30, 0xda, 0x3b, 0xd9, 0xbb, 0x38,
	0xa2, 0xe5, 0x8f, 0x2b, 0x26, 0xc9, 0xd5, 0x3d, 0xc5, 0x33, 0x38, 0x1e, 0xea, 0xb9, 0x90, 0xb3,
	0x01, 0xd3, 0x8c, 0x0b, 0xbb, 0xbc, 0x47, 0x77, 0x01, 0x9d, 0xcf, 0x6a, 0x21, 0x6b, 0x28, 0x4f,
	0x21, 0xf8, 0x69, 0x84, 0xc5, 0x7e, 0x31, 0x99, 0xa0, 0x09, 0x43, 0xd8, 0xcf, 0xc5, 0xdf, 0xb5,
	0xc6, 0x7f, 0x77, 0x4f, 0x00, 0x28, 0xb2, 0xf4, 0x01, 0xc5, 0x3b, 0x68, 0x0f, 0x94, 0x94, 0xc8,
	0xad, 0x50, 0x92, 0x62, 0x91, 0x63, 0xf8, 0x12, 0x5a, 0x28, 0xd9, 0x78, 0x5e, 0x0a, 0x0f, 0xe8,
	0xea, 0xaf, 0xfb, 0xaf, 0x09, 0xad, 0x81, 0x37, 0x36, 0xfc, 0x08, 0x8d, 0xcc, 0x16, 0xbe, 0x1e,
	0xf4, 0xce, 0xc8, 0xa3, 0x06, 0x93, 0xeb, 0x64, 0x48, 0x5d, 0x8b, 0xeb, 0xb4, 0x56, 0x44, 0x4f,
	0x6a, 0x77, 0x26, 0xc9, 0x15, 0x75, 0x2d, 0xe1, 0x0d, 0xb4, 0x0b, 0x6f, 0xe0, 0x88, 0xaf, 0x7c,
	0x89, 0x1a, 0x9e, 0xf2, 0xa1, 0x06, 0x65, 0xd7, 0x7a, 0x7a, 0x5c, 0xec, 0x8e, 0xe2, 0x37, 0x3c,
	0x4f, 0x57, 0xa6, 0x6f, 0xe8, 0xfb, 0x9e, 0x7e, 0x59, 0x83, 0x5e, 0x1d, 0x18, 0xed, 0xa4, 0xd5,
	0x11, 0xbe, 0x01, 0xe0, 0x4a, 0x4e, 0x31, 0x77, 0x3e, 0x47, 0x4d, 0x6f, 0xec, 0xd6, 0x49, 0xf8,
	0x1d, 0x0e, 0x17, 0x6e, 0x98, 0xa3, 0xb1, 0x9f, 0x55, 0xd4, 0xf2, 0x97, 0x93, 0x1a, 0x97, 0x6f,
	0x65, 0x80, 0x06, 0x8b, 0xad, 0x40, 0x7c, 0x85, 0xc0, 0x20, 0x4b, 0xd7, 0xc4, 0xa7, 0x9e, 0xf8,
	0xbe, 0x06, 0x71, 0x13, 0x19, 0x0a, 0x66, 0x13, 0x9f, 0x2f, 0x70, 0xf4, 0x07, 0xdd, 0x42, 0x8c,
	0xca, 0xf5, 0x8a, 0x0e, 0x3c, 0xf1, 0x7c, 0x9b, 0x58, 0x2e, 0x0e, 0x29, 0x17, 0x87, 0x24, 0x4b,
	0x8d, 0xe9, 0x0f, 0xb4, 0x56, 0xc8, 0x69, 0x4e, 0x0f, 0xcb, 0xee, 0x55, 0x84, 0x7e, 0x41, 0x87,
	0xdf, 0x05, 0x6f, 0x64, 0x5c, 0xf2, 0xa2, 0x67, 0x1e, 0xd8, 0xab, 0xf1, 0xc4, 0x4a, 0x66, 0x69,
	0x9b, 0xef, 0x1e, 0xf4, 0x3f, 0xc1, 0x5b, 0xae, 0xb2, 0xc7, 0x49, 0xfd, 0xa0, 0x7c, 0xcf, 0x37,
	0xb7, 0xde, 0x37, 0x8d, 0x19, 0xd7, 0xe3, 0x96, 0x5f, 0xf5, 0xcb, 0xff, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x76, 0x78, 0x1a, 0x9c, 0x6c, 0x04, 0x00, 0x00,
}
