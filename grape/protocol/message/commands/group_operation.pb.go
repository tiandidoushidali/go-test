// Code generated by protoc-gen-go.
// source: commands/group_operation.proto
// DO NOT EDIT!

package commands

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import protocol_message_misc "go-test/grape/protocol/message/misc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type GroupOperation struct {
	Operation int32                         `protobuf:"varint,1,opt,name=operation" json:"operation,omitempty"`
	Group     int64                         `protobuf:"varint,2,opt,name=group" json:"group,omitempty"`
	Members   []*protocol_message_misc.User `protobuf:"bytes,3,rep,name=members" json:"members,omitempty"`
	Content   string                        `protobuf:"bytes,4,opt,name=content" json:"content,omitempty"`
}

func (m *GroupOperation) Reset()                    { *m = GroupOperation{} }
func (m *GroupOperation) String() string            { return proto.CompactTextString(m) }
func (*GroupOperation) ProtoMessage()               {}
func (*GroupOperation) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *GroupOperation) GetOperation() int32 {
	if m != nil {
		return m.Operation
	}
	return 0
}

func (m *GroupOperation) GetGroup() int64 {
	if m != nil {
		return m.Group
	}
	return 0
}

func (m *GroupOperation) GetMembers() []*protocol_message_misc.User {
	if m != nil {
		return m.Members
	}
	return nil
}

func (m *GroupOperation) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func init() {
	proto.RegisterType((*GroupOperation)(nil), "protocol.message.commands.GroupOperation")
}

func init() { proto.RegisterFile("commands/group_operation.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0xce, 0xcf, 0xcd,
	0x4d, 0xcc, 0x4b, 0x29, 0xd6, 0x4f, 0x2f, 0xca, 0x2f, 0x2d, 0x88, 0xcf, 0x2f, 0x48, 0x2d, 0x4a,
	0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x04, 0x53, 0xc9, 0xf9,
	0x39, 0x7a, 0xb9, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0xa9, 0x7a, 0x30, 0x0d, 0x52, 0xfc, 0xb9, 0x99,
	0xc5, 0xc9, 0xfa, 0xa5, 0xc5, 0xa9, 0x45, 0x10, 0xb5, 0x4a, 0x53, 0x19, 0xb9, 0xf8, 0xdc, 0x41,
	0xa6, 0xf8, 0xc3, 0x0c, 0x11, 0x92, 0xe1, 0xe2, 0x84, 0x9b, 0x28, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1,
	0x1a, 0x84, 0x10, 0x10, 0x12, 0xe1, 0x62, 0x05, 0xdb, 0x2a, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x1c,
	0x04, 0xe1, 0x08, 0x99, 0x72, 0xb1, 0xe7, 0xa6, 0xe6, 0x26, 0xa5, 0x16, 0x15, 0x4b, 0x30, 0x2b,
	0x30, 0x6b, 0x70, 0x1b, 0x49, 0xeb, 0x61, 0x38, 0x02, 0x64, 0xb5, 0x5e, 0x68, 0x71, 0x6a, 0x51,
	0x10, 0x4c, 0xad, 0x90, 0x04, 0x17, 0x7b, 0x72, 0x7e, 0x5e, 0x49, 0x6a, 0x5e, 0x89, 0x04, 0x8b,
	0x02, 0xa3, 0x06, 0x67, 0x10, 0x8c, 0xeb, 0xe4, 0xc5, 0xa5, 0x9e, 0x9c, 0x9f, 0xab, 0x97, 0x9b,
	0x9a, 0x92, 0x93, 0x99, 0x97, 0x0d, 0x73, 0x2e, 0x36, 0x3f, 0x45, 0xc9, 0xa7, 0x17, 0x25, 0x16,
	0xa4, 0xea, 0xc3, 0x14, 0xe8, 0x43, 0x15, 0xe8, 0xc3, 0x14, 0x24, 0xb1, 0x81, 0xa5, 0x8c, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x48, 0x7e, 0x3b, 0x1e, 0x38, 0x01, 0x00, 0x00,
}
