// Code generated by protoc-gen-go.
// source: misc/group.proto
// DO NOT EDIT!

/*
Package misc is a generated protocol buffer package.

It is generated from these files:
	misc/group.proto
	misc/user.proto

It has these top-level messages:
	Group
	User
*/
package misc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Group struct {
	Id           int64    `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name         string   `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Avatar       []string `protobuf:"bytes,4,rep,name=avatar" json:"avatar,omitempty"`
	Amount       int32    `protobuf:"varint,5,opt,name=amount" json:"amount,omitempty"`
	Owner        int64    `protobuf:"varint,6,opt,name=owner" json:"owner,omitempty"`
	BusinessType int32    `protobuf:"varint,7,opt,name=businessType" json:"businessType,omitempty"`
	BusinessId   int64    `protobuf:"varint,8,opt,name=businessId" json:"businessId,omitempty"`
	AvatarUrl    string   `protobuf:"bytes,9,opt,name=avatarUrl" json:"avatarUrl,omitempty"`
	IsPublic     int32    `protobuf:"varint,10,opt,name=isPublic" json:"isPublic,omitempty"`
	IsForbidden  int32    `protobuf:"varint,11,opt,name=isForbidden" json:"isForbidden,omitempty"`
}

func (m *Group) Reset()                    { *m = Group{} }
func (m *Group) String() string            { return proto.CompactTextString(m) }
func (*Group) ProtoMessage()               {}
func (*Group) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Group) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Group) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Group) GetAvatar() []string {
	if m != nil {
		return m.Avatar
	}
	return nil
}

func (m *Group) GetAmount() int32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Group) GetOwner() int64 {
	if m != nil {
		return m.Owner
	}
	return 0
}

func (m *Group) GetBusinessType() int32 {
	if m != nil {
		return m.BusinessType
	}
	return 0
}

func (m *Group) GetBusinessId() int64 {
	if m != nil {
		return m.BusinessId
	}
	return 0
}

func (m *Group) GetAvatarUrl() string {
	if m != nil {
		return m.AvatarUrl
	}
	return ""
}

func (m *Group) GetIsPublic() int32 {
	if m != nil {
		return m.IsPublic
	}
	return 0
}

func (m *Group) GetIsForbidden() int32 {
	if m != nil {
		return m.IsForbidden
	}
	return 0
}

func init() {
	proto.RegisterType((*Group)(nil), "protocol.message.misc.Group")
}

func init() { proto.RegisterFile("misc/group.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 265 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0xe9, 0x76, 0x5b, 0xb7, 0xb3, 0x22, 0x32, 0xa8, 0x04, 0x15, 0x29, 0xeb, 0xa5, 0xa7,
	0xf4, 0xe0, 0x1b, 0xec, 0x41, 0xf1, 0x26, 0x45, 0x2f, 0xde, 0xd2, 0x26, 0x94, 0x60, 0x93, 0x94,
	0xa4, 0x55, 0x7c, 0x0c, 0xdf, 0x58, 0x3a, 0xdb, 0xd5, 0x15, 0x3c, 0x25, 0xdf, 0xf7, 0xcf, 0x4c,
	0xc2, 0xc0, 0xa9, 0xd1, 0xa1, 0x29, 0x5b, 0xef, 0xc6, 0x9e, 0xf7, 0xde, 0x0d, 0x0e, 0xcf, 0xe9,
	0x68, 0x5c, 0xc7, 0x8d, 0x0a, 0x41, 0xb4, 0x8a, 0x4f, 0x25, 0x9b, 0xaf, 0x05, 0x24, 0x0f, 0x53,
	0x19, 0x9e, 0xc0, 0x42, 0x4b, 0x16, 0xe5, 0x51, 0x11, 0x57, 0x0b, 0x2d, 0x11, 0x61, 0x69, 0x85,
	0x51, 0x2c, 0xce, 0xa3, 0x22, 0xab, 0xe8, 0x8e, 0x17, 0x90, 0x8a, 0x77, 0x31, 0x08, 0xcf, 0x96,
	0x79, 0x5c, 0x64, 0xd5, 0x4c, 0xe4, 0x8d, 0x1b, 0xed, 0xc0, 0x92, 0x3c, 0x2a, 0x92, 0x6a, 0x26,
	0x3c, 0x83, 0xc4, 0x7d, 0x58, 0xe5, 0x59, 0x4a, 0x63, 0x77, 0x80, 0x1b, 0x38, 0xae, 0xc7, 0xa0,
	0xad, 0x0a, 0xe1, 0xf9, 0xb3, 0x57, 0xec, 0x88, 0x7a, 0xfe, 0x38, 0xbc, 0x01, 0xd8, 0xf3, 0xa3,
	0x64, 0x2b, 0x6a, 0x3f, 0x30, 0x78, 0x0d, 0xd9, 0xee, 0xed, 0x17, 0xdf, 0xb1, 0x8c, 0xbe, 0xf8,
	0x2b, 0xf0, 0x12, 0x56, 0x3a, 0x3c, 0x8d, 0x75, 0xa7, 0x1b, 0x06, 0x34, 0xfd, 0x87, 0x31, 0x87,
	0xb5, 0x0e, 0xf7, 0xce, 0xd7, 0x5a, 0x4a, 0x65, 0xd9, 0x9a, 0xe2, 0x43, 0xb5, 0xdd, 0xc2, 0x6d,
	0xe3, 0x0c, 0x37, 0x4a, 0x76, 0xda, 0xbe, 0x29, 0xcf, 0xff, 0x5d, 0xdd, 0xeb, 0x55, 0xeb, 0x45,
	0xaf, 0xca, 0x7d, 0x58, 0xce, 0x61, 0x39, 0x85, 0x75, 0x4a, 0xfa, 0xee, 0x3b, 0x00, 0x00, 0xff,
	0xff, 0x07, 0x88, 0x1f, 0x35, 0x89, 0x01, 0x00, 0x00,
}
