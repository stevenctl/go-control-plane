// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/health_checker/redis/v2/redis.proto

package envoy_config_health_checker_redis_v2

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
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

type Redis struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Redis) Reset()         { *m = Redis{} }
func (m *Redis) String() string { return proto.CompactTextString(m) }
func (*Redis) ProtoMessage()    {}
func (*Redis) Descriptor() ([]byte, []int) {
	return fileDescriptor_055a998fcb839d64, []int{0}
}

func (m *Redis) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Redis.Unmarshal(m, b)
}
func (m *Redis) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Redis.Marshal(b, m, deterministic)
}
func (m *Redis) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Redis.Merge(m, src)
}
func (m *Redis) XXX_Size() int {
	return xxx_messageInfo_Redis.Size(m)
}
func (m *Redis) XXX_DiscardUnknown() {
	xxx_messageInfo_Redis.DiscardUnknown(m)
}

var xxx_messageInfo_Redis proto.InternalMessageInfo

func (m *Redis) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func init() {
	proto.RegisterType((*Redis)(nil), "envoy.config.health_checker.redis.v2.Redis")
}

func init() {
	proto.RegisterFile("envoy/config/health_checker/redis/v2/redis.proto", fileDescriptor_055a998fcb839d64)
}

var fileDescriptor_055a998fcb839d64 = []byte{
	// 174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x48, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0xcf, 0x48, 0x4d, 0xcc, 0x29, 0xc9, 0x88,
	0x4f, 0xce, 0x48, 0x4d, 0xce, 0x4e, 0x2d, 0xd2, 0x2f, 0x4a, 0x4d, 0xc9, 0x2c, 0xd6, 0x2f, 0x33,
	0x82, 0x30, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x54, 0xc0, 0x3a, 0xf4, 0x20, 0x3a, 0xf4,
	0x50, 0x75, 0xe8, 0x41, 0x14, 0x96, 0x19, 0x49, 0xc9, 0x96, 0xa6, 0x14, 0x24, 0xea, 0x27, 0xe6,
	0xe5, 0xe5, 0x97, 0x24, 0x96, 0x64, 0xe6, 0xe7, 0x15, 0xeb, 0x17, 0x97, 0x24, 0x96, 0x94, 0x42,
	0x0d, 0x51, 0x92, 0xe4, 0x62, 0x0d, 0x02, 0x29, 0x15, 0x12, 0xe0, 0x62, 0xce, 0x4e, 0xad, 0x94,
	0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x9d, 0xbc, 0x76, 0x35, 0x9c, 0xb8, 0xc8, 0xc6,
	0x24, 0xc0, 0xc4, 0x65, 0x94, 0x99, 0xaf, 0x07, 0xb6, 0xac, 0xa0, 0x28, 0xbf, 0xa2, 0x52, 0x8f,
	0x18, 0x7b, 0x9d, 0xb8, 0xc0, 0xc6, 0x06, 0x80, 0x2c, 0x09, 0x60, 0x4c, 0x62, 0x03, 0xdb, 0x66,
	0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xdf, 0xf5, 0x6d, 0x97, 0xe6, 0x00, 0x00, 0x00,
}