// Code generated by protoc-gen-go.
// source: types/business_notify.proto
// DO NOT EDIT!

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 业务内容的通知
type BusinessNotify struct {
	Type int32 `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
}

func (m *BusinessNotify) Reset()                    { *m = BusinessNotify{} }
func (m *BusinessNotify) String() string            { return proto.CompactTextString(m) }
func (*BusinessNotify) ProtoMessage()               {}
func (*BusinessNotify) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *BusinessNotify) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func init() {
	proto.RegisterType((*BusinessNotify)(nil), "protocol.message.types.BusinessNotify")
}

func init() { proto.RegisterFile("types/business_notify.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 136 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0xa9, 0x2c, 0x48,
	0x2d, 0xd6, 0x4f, 0x2a, 0x2d, 0xce, 0xcc, 0x4b, 0x2d, 0x2e, 0x8e, 0xcf, 0xcb, 0x2f, 0xc9, 0x4c,
	0xab, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x03, 0x53, 0xc9, 0xf9, 0x39, 0x7a, 0xb9,
	0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0xa9, 0x7a, 0x60, 0xd5, 0x4a, 0x2a, 0x5c, 0x7c, 0x4e, 0x50, 0x0d,
	0x7e, 0x60, 0xf5, 0x42, 0x42, 0x5c, 0x2c, 0x20, 0x29, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xd6, 0x20,
	0x30, 0xdb, 0xc9, 0x85, 0x4b, 0x25, 0x39, 0x3f, 0x57, 0x2f, 0x37, 0x35, 0x25, 0x27, 0x33, 0x2f,
	0x3b, 0xb5, 0x48, 0x0f, 0xbb, 0x69, 0x51, 0x32, 0xe9, 0x45, 0x89, 0x05, 0xa9, 0xfa, 0x30, 0x59,
	0x7d, 0xa8, 0xac, 0x3e, 0x58, 0x36, 0x89, 0x0d, 0x2c, 0x6e, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff,
	0x06, 0xcd, 0xa8, 0xfb, 0xa9, 0x00, 0x00, 0x00,
}
