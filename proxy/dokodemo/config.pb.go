// Code generated by protoc-gen-go.
// source: v2ray.com/core/proxy/dokodemo/config.proto
// DO NOT EDIT!

/*
Package dokodemo is a generated protocol buffer package.

It is generated from these files:
	v2ray.com/core/proxy/dokodemo/config.proto

It has these top-level messages:
	Config
*/
package dokodemo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import com_v2ray_core_common_net "v2ray.com/core/common/net"
import com_v2ray_core_common_net1 "v2ray.com/core/common/net"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Config struct {
	Address        *com_v2ray_core_common_net.AddressPB    `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	Port           uint32                                  `protobuf:"varint,2,opt,name=port" json:"port,omitempty"`
	NetworkList    *com_v2ray_core_common_net1.NetworkList `protobuf:"bytes,3,opt,name=network_list,json=networkList" json:"network_list,omitempty"`
	Timeout        uint32                                  `protobuf:"varint,4,opt,name=timeout" json:"timeout,omitempty"`
	FollowRedirect bool                                    `protobuf:"varint,5,opt,name=follow_redirect,json=followRedirect" json:"follow_redirect,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Config) GetAddress() *com_v2ray_core_common_net.AddressPB {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Config) GetNetworkList() *com_v2ray_core_common_net1.NetworkList {
	if m != nil {
		return m.NetworkList
	}
	return nil
}

func init() {
	proto.RegisterType((*Config)(nil), "com.v2ray.core.proxy.dokodemo.Config")
}

func init() { proto.RegisterFile("v2ray.com/core/proxy/dokodemo/config.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x90, 0x4d, 0x4b, 0x03, 0x31,
	0x10, 0x86, 0x59, 0xad, 0x6d, 0x49, 0xfd, 0x80, 0x9c, 0x82, 0x20, 0x14, 0x11, 0xbb, 0x78, 0x48,
	0xa0, 0xde, 0x05, 0xeb, 0x49, 0x10, 0x91, 0x3d, 0x7a, 0x29, 0x35, 0x3b, 0x95, 0xd0, 0x4d, 0x66,
	0x99, 0x1d, 0xad, 0xfd, 0xcb, 0xfe, 0x0a, 0x21, 0xd9, 0x20, 0xf4, 0xd0, 0x5b, 0x66, 0x78, 0xf2,
	0xbc, 0x6f, 0x22, 0xee, 0xbe, 0xe7, 0xb4, 0xda, 0x69, 0x8b, 0xde, 0x58, 0x24, 0x30, 0x2d, 0xe1,
	0xcf, 0xce, 0xd4, 0xb8, 0xc1, 0x1a, 0x3c, 0x1a, 0x8b, 0x61, 0xed, 0x3e, 0x75, 0x4b, 0xc8, 0x28,
	0xaf, 0x2c, 0x7a, 0x9d, 0x79, 0x02, 0x1d, 0x59, 0x9d, 0xd9, 0xcb, 0xd9, 0x9e, 0xca, 0xa2, 0xf7,
	0x18, 0x4c, 0x00, 0x36, 0xab, 0xba, 0x26, 0xe8, 0xba, 0xe4, 0x39, 0x04, 0x06, 0xe0, 0x2d, 0xd2,
	0x26, 0x81, 0xd7, 0xbf, 0x85, 0x18, 0x3e, 0xc5, 0x06, 0xf2, 0x41, 0x8c, 0x7a, 0x89, 0x2a, 0xa6,
	0x45, 0x39, 0x99, 0xdf, 0xe8, 0xbd, 0x36, 0xc9, 0xa2, 0x03, 0xb0, 0x7e, 0x4c, 0xe4, 0xdb, 0xa2,
	0xca, 0x97, 0xa4, 0x14, 0x83, 0x16, 0x89, 0xd5, 0xd1, 0xb4, 0x28, 0xcf, 0xaa, 0x78, 0x96, 0xcf,
	0xe2, 0xb4, 0xcf, 0x5b, 0x36, 0xae, 0x63, 0x75, 0x1c, 0xc5, 0xb7, 0x07, 0xc4, 0xaf, 0x09, 0x7f,
	0x71, 0x1d, 0x57, 0x93, 0xf0, 0x3f, 0x48, 0x25, 0x46, 0xec, 0x3c, 0xe0, 0x17, 0xab, 0x41, 0x4c,
	0xc8, 0xa3, 0x9c, 0x89, 0x8b, 0x35, 0x36, 0x0d, 0x6e, 0x97, 0x04, 0xb5, 0x23, 0xb0, 0xac, 0x4e,
	0xa6, 0x45, 0x39, 0xae, 0xce, 0xd3, 0xba, 0xea, 0xb7, 0x0b, 0xf1, 0x3e, 0xce, 0x5f, 0xf9, 0x31,
	0x8c, 0xef, 0xbf, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x85, 0x0f, 0xd2, 0x20, 0x9e, 0x01, 0x00,
	0x00,
}
