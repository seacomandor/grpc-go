// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc/gcp/transport_security_common.proto

package grpc_gcp

import (
	fmt "fmt"
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

// The security level of the created channel. The list is sorted in increasing
// level of security. This order must always be maintained.
type SecurityLevel int32

const (
	SecurityLevel_SECURITY_NONE         SecurityLevel = 0
	SecurityLevel_INTEGRITY_ONLY        SecurityLevel = 1
	SecurityLevel_INTEGRITY_AND_PRIVACY SecurityLevel = 2
)

var SecurityLevel_name = map[int32]string{
	0: "SECURITY_NONE",
	1: "INTEGRITY_ONLY",
	2: "INTEGRITY_AND_PRIVACY",
}

var SecurityLevel_value = map[string]int32{
	"SECURITY_NONE":         0,
	"INTEGRITY_ONLY":        1,
	"INTEGRITY_AND_PRIVACY": 2,
}

func (x SecurityLevel) String() string {
	return proto.EnumName(SecurityLevel_name, int32(x))
}

func (SecurityLevel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b97e31e3cc23582a, []int{0}
}

// Max and min supported RPC protocol versions.
type RpcProtocolVersions struct {
	// Maximum supported RPC version.
	MaxRpcVersion *RpcProtocolVersions_Version `protobuf:"bytes,1,opt,name=max_rpc_version,json=maxRpcVersion,proto3" json:"max_rpc_version,omitempty"`
	// Minimum supported RPC version.
	MinRpcVersion        *RpcProtocolVersions_Version `protobuf:"bytes,2,opt,name=min_rpc_version,json=minRpcVersion,proto3" json:"min_rpc_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *RpcProtocolVersions) Reset()         { *m = RpcProtocolVersions{} }
func (m *RpcProtocolVersions) String() string { return proto.CompactTextString(m) }
func (*RpcProtocolVersions) ProtoMessage()    {}
func (*RpcProtocolVersions) Descriptor() ([]byte, []int) {
	return fileDescriptor_b97e31e3cc23582a, []int{0}
}

func (m *RpcProtocolVersions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RpcProtocolVersions.Unmarshal(m, b)
}
func (m *RpcProtocolVersions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RpcProtocolVersions.Marshal(b, m, deterministic)
}
func (m *RpcProtocolVersions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RpcProtocolVersions.Merge(m, src)
}
func (m *RpcProtocolVersions) XXX_Size() int {
	return xxx_messageInfo_RpcProtocolVersions.Size(m)
}
func (m *RpcProtocolVersions) XXX_DiscardUnknown() {
	xxx_messageInfo_RpcProtocolVersions.DiscardUnknown(m)
}

var xxx_messageInfo_RpcProtocolVersions proto.InternalMessageInfo

func (m *RpcProtocolVersions) GetMaxRpcVersion() *RpcProtocolVersions_Version {
	if m != nil {
		return m.MaxRpcVersion
	}
	return nil
}

func (m *RpcProtocolVersions) GetMinRpcVersion() *RpcProtocolVersions_Version {
	if m != nil {
		return m.MinRpcVersion
	}
	return nil
}

// RPC version contains a major version and a minor version.
type RpcProtocolVersions_Version struct {
	Major                uint32   `protobuf:"varint,1,opt,name=major,proto3" json:"major,omitempty"`
	Minor                uint32   `protobuf:"varint,2,opt,name=minor,proto3" json:"minor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RpcProtocolVersions_Version) Reset()         { *m = RpcProtocolVersions_Version{} }
func (m *RpcProtocolVersions_Version) String() string { return proto.CompactTextString(m) }
func (*RpcProtocolVersions_Version) ProtoMessage()    {}
func (*RpcProtocolVersions_Version) Descriptor() ([]byte, []int) {
	return fileDescriptor_b97e31e3cc23582a, []int{0, 0}
}

func (m *RpcProtocolVersions_Version) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RpcProtocolVersions_Version.Unmarshal(m, b)
}
func (m *RpcProtocolVersions_Version) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RpcProtocolVersions_Version.Marshal(b, m, deterministic)
}
func (m *RpcProtocolVersions_Version) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RpcProtocolVersions_Version.Merge(m, src)
}
func (m *RpcProtocolVersions_Version) XXX_Size() int {
	return xxx_messageInfo_RpcProtocolVersions_Version.Size(m)
}
func (m *RpcProtocolVersions_Version) XXX_DiscardUnknown() {
	xxx_messageInfo_RpcProtocolVersions_Version.DiscardUnknown(m)
}

var xxx_messageInfo_RpcProtocolVersions_Version proto.InternalMessageInfo

func (m *RpcProtocolVersions_Version) GetMajor() uint32 {
	if m != nil {
		return m.Major
	}
	return 0
}

func (m *RpcProtocolVersions_Version) GetMinor() uint32 {
	if m != nil {
		return m.Minor
	}
	return 0
}

func init() {
	proto.RegisterEnum("grpc.s.gcp.SecurityLevel", SecurityLevel_name, SecurityLevel_value)
	proto.RegisterType((*RpcProtocolVersions)(nil), "grpc.gcp.RpcProtocolVersions")
	proto.RegisterType((*RpcProtocolVersions_Version)(nil), "grpc.gcp.RpcProtocolVersions.Version")
}

func init() {
	proto.RegisterFile("grpc/gcp/transport_security_common.proto", fileDescriptor_b97e31e3cc23582a)
}

var fileDescriptor_b97e31e3cc23582a = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x41, 0x4b, 0x3b, 0x31,
	0x10, 0xc5, 0xff, 0x5b, 0xf8, 0xab, 0x44, 0x56, 0xeb, 0x6a, 0x41, 0xc5, 0x83, 0x08, 0x42, 0xf1,
	0x90, 0x05, 0xc5, 0xb3, 0xb4, 0xb5, 0x48, 0xa1, 0x6e, 0xeb, 0xb6, 0x16, 0xea, 0x25, 0xc4, 0x18,
	0x42, 0x24, 0x9b, 0x09, 0xb3, 0xb1, 0xd4, 0xaf, 0xec, 0xa7, 0x90, 0x4d, 0xbb, 0x14, 0xc1, 0x8b,
	0xb7, 0xbc, 0xc7, 0xcc, 0x6f, 0x32, 0xf3, 0x48, 0x5b, 0xa1, 0x13, 0xa9, 0x12, 0x2e, 0xf5, 0xc8,
	0x6d, 0xe9, 0x00, 0x3d, 0x2b, 0xa5, 0xf8, 0x40, 0xed, 0x3f, 0x99, 0x80, 0xa2, 0x00, 0x4b, 0x1d,
	0x82, 0x87, 0x64, 0xa7, 0xaa, 0xa4, 0x4a, 0xb8, 0x8b, 0xaf, 0x88, 0x1c, 0xe6, 0x4e, 0x8c, 0x2b,
	0x5b, 0x80, 0x99, 0x49, 0x2c, 0x35, 0xd8, 0x32, 0x79, 0x24, 0xfb, 0x05, 0x5f, 0x32, 0x74, 0x82,
	0x2d, 0x56, 0xde, 0x71, 0x74, 0x1e, 0xb5, 0x77, 0xaf, 0x2f, 0x69, 0xdd, 0x4b, 0x7f, 0xe9, 0xa3,
	0xeb, 0x47, 0x1e, 0x17, 0x7c, 0x99, 0x3b, 0xb1, 0x96, 0x01, 0xa7, 0xed, 0x0f, 0x5c, 0xe3, 0x6f,
	0x38, 0x6d, 0x37, 0xb8, 0xd3, 0x5b, 0xb2, 0x5d, 0x93, 0x8f, 0xc8, 0xff, 0x82, 0xbf, 0x03, 0x86,
	0xef, 0xc5, 0xf9, 0x4a, 0x04, 0x57, 0x5b, 0xc0, 0x30, 0xa5, 0x72, 0x2b, 0x71, 0xf5, 0x44, 0xe2,
	0xc9, 0xfa, 0x1e, 0x43, 0xb9, 0x90, 0x26, 0x39, 0x20, 0xf1, 0xa4, 0xdf, 0x7b, 0xce, 0x07, 0xd3,
	0x39, 0xcb, 0x46, 0x59, 0xbf, 0xf9, 0x2f, 0x49, 0xc8, 0xde, 0x20, 0x9b, 0xf6, 0x1f, 0x82, 0x37,
	0xca, 0x86, 0xf3, 0x66, 0x94, 0x9c, 0x90, 0xd6, 0xc6, 0xeb, 0x64, 0xf7, 0x6c, 0x9c, 0x0f, 0x66,
	0x9d, 0xde, 0xbc, 0xd9, 0xe8, 0x2e, 0x49, 0x4b, 0xc3, 0x6a, 0x07, 0x6e, 0x7c, 0x49, 0xb5, 0xf5,
	0x12, 0x2d, 0x37, 0xdd, 0xb3, 0x69, 0x9d, 0x41, 0x3d, 0xb2, 0x17, 0x12, 0x08, 0x2b, 0x8e, 0xa3,
	0x97, 0x3b, 0x05, 0xa0, 0x8c, 0xa4, 0x0a, 0x0c, 0xb7, 0x8a, 0x02, 0xaa, 0x34, 0xc4, 0x27, 0x50,
	0xbe, 0x49, 0xeb, 0x35, 0x37, 0x65, 0x5a, 0x11, 0xd3, 0x9a, 0x98, 0x86, 0xe8, 0x42, 0x11, 0x53,
	0xc2, 0xbd, 0x6e, 0x05, 0x7d, 0xf3, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x31, 0x14, 0xb4, 0x11, 0xf6,
	0x01, 0x00, 0x00,
}
