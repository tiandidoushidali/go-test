// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: med-common/app/service/med-xim/api/im/event.proto

package xim

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// 用户注册事件
type SendMsgEvent struct {
	// 消息内容(message的序列化base64格式)
	Content              string   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendMsgEvent) Reset()         { *m = SendMsgEvent{} }
func (m *SendMsgEvent) String() string { return proto.CompactTextString(m) }
func (*SendMsgEvent) ProtoMessage()    {}
func (*SendMsgEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_401aadbf632e47c3, []int{0}
}
func (m *SendMsgEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SendMsgEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SendMsgEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SendMsgEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMsgEvent.Merge(m, src)
}
func (m *SendMsgEvent) XXX_Size() int {
	return m.Size()
}
func (m *SendMsgEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMsgEvent.DiscardUnknown(m)
}

var xxx_messageInfo_SendMsgEvent proto.InternalMessageInfo

func (m *SendMsgEvent) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

// 敏感词注册事件
type SensitiveWordMsgEvent struct {
	// 消息内容(message的序列化base64格式)
	Content              string   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SensitiveWordMsgEvent) Reset()         { *m = SensitiveWordMsgEvent{} }
func (m *SensitiveWordMsgEvent) String() string { return proto.CompactTextString(m) }
func (*SensitiveWordMsgEvent) ProtoMessage()    {}
func (*SensitiveWordMsgEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_401aadbf632e47c3, []int{1}
}
func (m *SensitiveWordMsgEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SensitiveWordMsgEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SensitiveWordMsgEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SensitiveWordMsgEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SensitiveWordMsgEvent.Merge(m, src)
}
func (m *SensitiveWordMsgEvent) XXX_Size() int {
	return m.Size()
}
func (m *SensitiveWordMsgEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_SensitiveWordMsgEvent.DiscardUnknown(m)
}

var xxx_messageInfo_SensitiveWordMsgEvent proto.InternalMessageInfo

func (m *SensitiveWordMsgEvent) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

// 消息已读事件
type MsgReadEvent struct {
	UserId               uint32   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	MapId                string   `protobuf:"bytes,2,opt,name=map_id,json=mapId,proto3" json:"map_id,omitempty"`
	MsgId                uint64   `protobuf:"varint,3,opt,name=msg_id,json=msgId,proto3" json:"msg_id,omitempty"`
	Time                 int64    `protobuf:"varint,4,opt,name=time,proto3" json:"time,omitempty"`
	Ext                  string   `protobuf:"bytes,5,opt,name=ext,proto3" json:"ext,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgReadEvent) Reset()         { *m = MsgReadEvent{} }
func (m *MsgReadEvent) String() string { return proto.CompactTextString(m) }
func (*MsgReadEvent) ProtoMessage()    {}
func (*MsgReadEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_401aadbf632e47c3, []int{2}
}
func (m *MsgReadEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgReadEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgReadEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgReadEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgReadEvent.Merge(m, src)
}
func (m *MsgReadEvent) XXX_Size() int {
	return m.Size()
}
func (m *MsgReadEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgReadEvent.DiscardUnknown(m)
}

var xxx_messageInfo_MsgReadEvent proto.InternalMessageInfo

func (m *MsgReadEvent) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *MsgReadEvent) GetMapId() string {
	if m != nil {
		return m.MapId
	}
	return ""
}

func (m *MsgReadEvent) GetMsgId() uint64 {
	if m != nil {
		return m.MsgId
	}
	return 0
}

func (m *MsgReadEvent) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *MsgReadEvent) GetExt() string {
	if m != nil {
		return m.Ext
	}
	return ""
}

type ImGroupEvent struct {
	Type                 int32    `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	GroupId              int32    `protobuf:"varint,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	UserId               uint32   `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	BusinessType         uint32   `protobuf:"varint,4,opt,name=business_type,json=businessType,proto3" json:"business_type,omitempty"`
	BusinessId           uint32   `protobuf:"varint,5,opt,name=business_id,json=businessId,proto3" json:"business_id,omitempty"`
	Ext                  string   `protobuf:"bytes,6,opt,name=ext,proto3" json:"ext,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImGroupEvent) Reset()         { *m = ImGroupEvent{} }
func (m *ImGroupEvent) String() string { return proto.CompactTextString(m) }
func (*ImGroupEvent) ProtoMessage()    {}
func (*ImGroupEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_401aadbf632e47c3, []int{3}
}
func (m *ImGroupEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ImGroupEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ImGroupEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ImGroupEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImGroupEvent.Merge(m, src)
}
func (m *ImGroupEvent) XXX_Size() int {
	return m.Size()
}
func (m *ImGroupEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_ImGroupEvent.DiscardUnknown(m)
}

var xxx_messageInfo_ImGroupEvent proto.InternalMessageInfo

func (m *ImGroupEvent) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ImGroupEvent) GetGroupId() int32 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *ImGroupEvent) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *ImGroupEvent) GetBusinessType() uint32 {
	if m != nil {
		return m.BusinessType
	}
	return 0
}

func (m *ImGroupEvent) GetBusinessId() uint32 {
	if m != nil {
		return m.BusinessId
	}
	return 0
}

func (m *ImGroupEvent) GetExt() string {
	if m != nil {
		return m.Ext
	}
	return ""
}

func init() {
	proto.RegisterType((*SendMsgEvent)(nil), "med.xim.SendMsgEvent")
	proto.RegisterType((*SensitiveWordMsgEvent)(nil), "med.xim.SensitiveWordMsgEvent")
	proto.RegisterType((*MsgReadEvent)(nil), "med.xim.MsgReadEvent")
	proto.RegisterType((*ImGroupEvent)(nil), "med.xim.ImGroupEvent")
}

func init() {
	proto.RegisterFile("med-common/app/service/med-xim/api/im/event.proto", fileDescriptor_401aadbf632e47c3)
}

var fileDescriptor_401aadbf632e47c3 = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcd, 0x4a, 0x23, 0x41,
	0x14, 0x85, 0xa9, 0xe9, 0xbf, 0x99, 0x3b, 0x1d, 0x98, 0x29, 0x08, 0xb6, 0x9b, 0x18, 0xe2, 0xa6,
	0x37, 0xa6, 0x09, 0xbe, 0x81, 0x20, 0xd2, 0x8b, 0x6c, 0x2a, 0x82, 0xe0, 0x26, 0x74, 0x52, 0x97,
	0xa6, 0x16, 0xf5, 0x43, 0x57, 0x25, 0xb4, 0xf8, 0x4e, 0x3e, 0x87, 0x4b, 0x1f, 0x41, 0xf2, 0x24,
	0x52, 0x15, 0x3b, 0xe8, 0xce, 0xdd, 0xbd, 0x5f, 0x9d, 0x7b, 0xce, 0x81, 0x82, 0x85, 0x44, 0x7e,
	0xb5, 0xd5, 0x52, 0x6a, 0x55, 0x35, 0xc6, 0x54, 0x16, 0xbb, 0xbd, 0xd8, 0x62, 0xe5, 0x71, 0x2f,
	0x64, 0xd5, 0x18, 0x51, 0x09, 0x59, 0xe1, 0x1e, 0x95, 0x9b, 0x9b, 0x4e, 0x3b, 0x4d, 0x33, 0x89,
	0x7c, 0xde, 0x0b, 0x39, 0x2b, 0x21, 0x5f, 0xa1, 0xe2, 0x4b, 0xdb, 0xde, 0xfa, 0x67, 0x5a, 0x40,
	0xb6, 0xd5, 0xca, 0xa1, 0x72, 0x05, 0x99, 0x92, 0xf2, 0x0f, 0x1b, 0xd6, 0xd9, 0x02, 0xc6, 0x2b,
	0x54, 0x56, 0x38, 0xb1, 0xc7, 0x07, 0xdd, 0xfd, 0xe4, 0xe4, 0x19, 0xf2, 0xa5, 0x6d, 0x19, 0x36,
	0xfc, 0xa8, 0x3c, 0x83, 0x6c, 0x67, 0xb1, 0x5b, 0x0b, 0x1e, 0x94, 0x23, 0x96, 0xfa, 0xb5, 0xe6,
	0x74, 0x0c, 0xa9, 0x6c, 0x8c, 0xe7, 0xbf, 0x82, 0x43, 0x22, 0x1b, 0xf3, 0x89, 0x6d, 0xeb, 0x71,
	0x34, 0x25, 0x65, 0xcc, 0x12, 0x69, 0xdb, 0x9a, 0x53, 0x0a, 0xb1, 0x13, 0x12, 0x8b, 0x78, 0x4a,
	0xca, 0x88, 0x85, 0x99, 0xfe, 0x83, 0x08, 0x7b, 0x57, 0x24, 0xe1, 0xdc, 0x8f, 0xb3, 0x17, 0x02,
	0x79, 0x2d, 0xef, 0x3a, 0xbd, 0x33, 0xc7, 0x74, 0x7f, 0xf6, 0x64, 0x30, 0x44, 0x27, 0x2c, 0xcc,
	0xf4, 0x1c, 0x7e, 0xb7, 0x5e, 0x31, 0x44, 0x27, 0x2c, 0x0b, 0x7b, 0xcd, 0xbf, 0x96, 0x8d, 0xbe,
	0x95, 0xbd, 0x84, 0xd1, 0x66, 0x67, 0x85, 0x42, 0x6b, 0xd7, 0xc1, 0x30, 0x0e, 0xcf, 0xf9, 0x00,
	0xef, 0xbd, 0xf1, 0x05, 0xfc, 0x3d, 0x89, 0x04, 0x0f, 0xbd, 0x46, 0x0c, 0x06, 0x54, 0xf3, 0xa1,
	0x70, 0x7a, 0x2a, 0x7c, 0xf3, 0xff, 0xf5, 0x30, 0x21, 0x6f, 0x87, 0x09, 0x79, 0x3f, 0x4c, 0xc8,
	0x63, 0xd4, 0x0b, 0xb9, 0x49, 0xc3, 0x6f, 0x5d, 0x7f, 0x04, 0x00, 0x00, 0xff, 0xff, 0x01, 0x7b,
	0xeb, 0xd6, 0xe2, 0x01, 0x00, 0x00,
}

func (m *SendMsgEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SendMsgEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SendMsgEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Content) > 0 {
		i -= len(m.Content)
		copy(dAtA[i:], m.Content)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Content)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SensitiveWordMsgEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SensitiveWordMsgEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SensitiveWordMsgEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Content) > 0 {
		i -= len(m.Content)
		copy(dAtA[i:], m.Content)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Content)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgReadEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgReadEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgReadEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Ext) > 0 {
		i -= len(m.Ext)
		copy(dAtA[i:], m.Ext)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Ext)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Time != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.Time))
		i--
		dAtA[i] = 0x20
	}
	if m.MsgId != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.MsgId))
		i--
		dAtA[i] = 0x18
	}
	if len(m.MapId) > 0 {
		i -= len(m.MapId)
		copy(dAtA[i:], m.MapId)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.MapId)))
		i--
		dAtA[i] = 0x12
	}
	if m.UserId != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.UserId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ImGroupEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ImGroupEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ImGroupEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Ext) > 0 {
		i -= len(m.Ext)
		copy(dAtA[i:], m.Ext)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Ext)))
		i--
		dAtA[i] = 0x32
	}
	if m.BusinessId != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.BusinessId))
		i--
		dAtA[i] = 0x28
	}
	if m.BusinessType != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.BusinessType))
		i--
		dAtA[i] = 0x20
	}
	if m.UserId != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.UserId))
		i--
		dAtA[i] = 0x18
	}
	if m.GroupId != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.GroupId))
		i--
		dAtA[i] = 0x10
	}
	if m.Type != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvent(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SendMsgEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Content)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *SensitiveWordMsgEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Content)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *MsgReadEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.UserId != 0 {
		n += 1 + sovEvent(uint64(m.UserId))
	}
	l = len(m.MapId)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	if m.MsgId != 0 {
		n += 1 + sovEvent(uint64(m.MsgId))
	}
	if m.Time != 0 {
		n += 1 + sovEvent(uint64(m.Time))
	}
	l = len(m.Ext)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ImGroupEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovEvent(uint64(m.Type))
	}
	if m.GroupId != 0 {
		n += 1 + sovEvent(uint64(m.GroupId))
	}
	if m.UserId != 0 {
		n += 1 + sovEvent(uint64(m.UserId))
	}
	if m.BusinessType != 0 {
		n += 1 + sovEvent(uint64(m.BusinessType))
	}
	if m.BusinessId != 0 {
		n += 1 + sovEvent(uint64(m.BusinessId))
	}
	l = len(m.Ext)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvent(x uint64) (n int) {
	return sovEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SendMsgEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SendMsgEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SendMsgEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Content = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEvent
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthEvent
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SensitiveWordMsgEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SensitiveWordMsgEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SensitiveWordMsgEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Content = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEvent
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthEvent
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgReadEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgReadEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgReadEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			m.UserId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UserId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MapId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MapId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgId", wireType)
			}
			m.MsgId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MsgId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			m.Time = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Time |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ext", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ext = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEvent
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthEvent
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ImGroupEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ImGroupEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ImGroupEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupId", wireType)
			}
			m.GroupId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GroupId |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			m.UserId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UserId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BusinessType", wireType)
			}
			m.BusinessType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BusinessType |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BusinessId", wireType)
			}
			m.BusinessId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BusinessId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ext", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ext = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEvent
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthEvent
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvent
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthEvent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvent = fmt.Errorf("proto: unexpected end of group")
)