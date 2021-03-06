// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: med-common/app/service/med-xim/api/im/entity/group.proto

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

// 群组信息数据结构
type Group struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Owner                uint32   `protobuf:"varint,3,opt,name=owner,proto3" json:"owner,omitempty"`
	BusinessType         uint32   `protobuf:"varint,4,opt,name=businessType,proto3" json:"businessType,omitempty"`
	BusinessId           uint32   `protobuf:"varint,5,opt,name=businessId,proto3" json:"businessId,omitempty"`
	IsPublic             int32    `protobuf:"varint,6,opt,name=isPublic,proto3" json:"isPublic,omitempty"`
	IsForbidden          int32    `protobuf:"varint,7,opt,name=isForbidden,proto3" json:"isForbidden,omitempty"`
	InsertTime           uint32   `protobuf:"varint,8,opt,name=insertTime,proto3" json:"insertTime,omitempty"`
	AvatarList           []string `protobuf:"bytes,9,rep,name=avatarList,proto3" json:"avatarList,omitempty"`
	Content              string   `protobuf:"bytes,10,opt,name=content,proto3" json:"content,omitempty"`
	Number               int32    `protobuf:"varint,11,opt,name=number,proto3" json:"number,omitempty"`
	Broadcast            string   `protobuf:"bytes,12,opt,name=broadcast,proto3" json:"broadcast,omitempty"`
	QrCodeUrl            string   `protobuf:"bytes,13,opt,name=qrCodeUrl,proto3" json:"qrCodeUrl,omitempty"`
	ShareH5Url           string   `protobuf:"bytes,14,opt,name=shareH5Url,proto3" json:"shareH5Url,omitempty"`
	JoinStatus           int32    `protobuf:"varint,15,opt,name=joinStatus,proto3" json:"joinStatus,omitempty"`
	IsNewOrigin          bool     `protobuf:"varint,16,opt,name=isNewOrigin,proto3" json:"isNewOrigin,omitempty"`
	OwnerUser            *User    `protobuf:"bytes,17,opt,name=ownerUser,proto3" json:"ownerUser,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Group) Reset()         { *m = Group{} }
func (m *Group) String() string { return proto.CompactTextString(m) }
func (*Group) ProtoMessage()    {}
func (*Group) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b557e13e3e3ad6, []int{0}
}
func (m *Group) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Group) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Group.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Group) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Group.Merge(m, src)
}
func (m *Group) XXX_Size() int {
	return m.Size()
}
func (m *Group) XXX_DiscardUnknown() {
	xxx_messageInfo_Group.DiscardUnknown(m)
}

var xxx_messageInfo_Group proto.InternalMessageInfo

func (m *Group) GetId() uint32 {
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

func (m *Group) GetOwner() uint32 {
	if m != nil {
		return m.Owner
	}
	return 0
}

func (m *Group) GetBusinessType() uint32 {
	if m != nil {
		return m.BusinessType
	}
	return 0
}

func (m *Group) GetBusinessId() uint32 {
	if m != nil {
		return m.BusinessId
	}
	return 0
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

func (m *Group) GetInsertTime() uint32 {
	if m != nil {
		return m.InsertTime
	}
	return 0
}

func (m *Group) GetAvatarList() []string {
	if m != nil {
		return m.AvatarList
	}
	return nil
}

func (m *Group) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Group) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *Group) GetBroadcast() string {
	if m != nil {
		return m.Broadcast
	}
	return ""
}

func (m *Group) GetQrCodeUrl() string {
	if m != nil {
		return m.QrCodeUrl
	}
	return ""
}

func (m *Group) GetShareH5Url() string {
	if m != nil {
		return m.ShareH5Url
	}
	return ""
}

func (m *Group) GetJoinStatus() int32 {
	if m != nil {
		return m.JoinStatus
	}
	return 0
}

func (m *Group) GetIsNewOrigin() bool {
	if m != nil {
		return m.IsNewOrigin
	}
	return false
}

func (m *Group) GetOwnerUser() *User {
	if m != nil {
		return m.OwnerUser
	}
	return nil
}

func init() {
	proto.RegisterType((*Group)(nil), "med.xim.entity.Group")
}

func init() {
	proto.RegisterFile("med-common/app/service/med-xim/api/im/entity/group.proto", fileDescriptor_17b557e13e3e3ad6)
}

var fileDescriptor_17b557e13e3e3ad6 = []byte{
	// 409 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcb, 0x6e, 0xd4, 0x30,
	0x14, 0x86, 0xe5, 0xb9, 0x75, 0xc6, 0xd3, 0x0e, 0xd4, 0xaa, 0xd0, 0x51, 0x85, 0x46, 0x51, 0x57,
	0xd9, 0x90, 0x48, 0x45, 0x08, 0xd6, 0x20, 0x71, 0x91, 0x10, 0xa0, 0xd0, 0x6e, 0xd8, 0x39, 0xf1,
	0x51, 0x39, 0xa8, 0xb6, 0x83, 0xed, 0xb4, 0xd3, 0x27, 0xe0, 0xd5, 0x58, 0xf2, 0x08, 0x68, 0x9e,
	0x04, 0xd9, 0xa1, 0x9d, 0x74, 0xd9, 0x5d, 0xfe, 0xef, 0x3b, 0xf9, 0x9d, 0xcb, 0xe1, 0xaf, 0x34,
	0xaa, 0x67, 0x8d, 0xd5, 0xda, 0x9a, 0x52, 0xb6, 0x6d, 0xe9, 0xd1, 0x5d, 0x51, 0x83, 0x65, 0xc4,
	0x1b, 0xd2, 0xa5, 0x6c, 0xa9, 0x24, 0x5d, 0xa2, 0x09, 0x14, 0x6e, 0xca, 0x0b, 0x67, 0xbb, 0xb6,
	0x68, 0x9d, 0x0d, 0x56, 0xac, 0x34, 0xaa, 0x62, 0x43, 0xba, 0xe8, 0xdd, 0xf1, 0xcb, 0x07, 0x35,
	0x75, 0x1e, 0x5d, 0x5f, 0x74, 0xf2, 0x6b, 0xc2, 0xa7, 0xef, 0x62, 0xb1, 0x58, 0xf1, 0x11, 0x29,
	0x60, 0x19, 0xcb, 0x0f, 0xaa, 0x11, 0x29, 0x21, 0xf8, 0xc4, 0x48, 0x8d, 0x30, 0xca, 0x58, 0xbe,
	0xa8, 0xd2, 0xb5, 0x38, 0xe2, 0x53, 0x7b, 0x6d, 0xd0, 0xc1, 0x38, 0x8d, 0xf5, 0x41, 0x9c, 0xf0,
	0xfd, 0xba, 0xf3, 0x64, 0xd0, 0xfb, 0xb3, 0x9b, 0x16, 0x61, 0x92, 0xe4, 0x3d, 0x26, 0xd6, 0x9c,
	0xdf, 0xe6, 0x0f, 0x0a, 0xa6, 0x69, 0x62, 0x40, 0xc4, 0x31, 0x9f, 0x93, 0xff, 0xd2, 0xd5, 0x97,
	0xd4, 0xc0, 0x2c, 0x63, 0xf9, 0xb4, 0xba, 0xcb, 0x22, 0xe3, 0x4b, 0xf2, 0x6f, 0xad, 0xab, 0x49,
	0x29, 0x34, 0xb0, 0x97, 0xf4, 0x10, 0xc5, 0x76, 0x32, 0x1e, 0x5d, 0x38, 0x23, 0x8d, 0x30, 0xef,
	0xdb, 0x77, 0x24, 0x7a, 0x79, 0x25, 0x83, 0x74, 0x1f, 0xc9, 0x07, 0x58, 0x64, 0xe3, 0x7c, 0x51,
	0x0d, 0x88, 0x00, 0xbe, 0xd7, 0x58, 0x13, 0xd0, 0x04, 0xe0, 0xe9, 0x75, 0x6f, 0xa3, 0x78, 0xc2,
	0x67, 0xa6, 0xd3, 0x35, 0x3a, 0x58, 0xa6, 0x63, 0xff, 0x27, 0xf1, 0x94, 0x2f, 0x6a, 0x67, 0xa5,
	0x6a, 0xa4, 0x0f, 0xb0, 0x9f, 0xee, 0xd9, 0x81, 0x68, 0x7f, 0xba, 0x37, 0x56, 0xe1, 0xb9, 0xbb,
	0x84, 0x83, 0xde, 0xde, 0x81, 0xf8, 0x34, 0xfe, 0xbb, 0x74, 0xf8, 0xfe, 0x45, 0xd4, 0xab, 0xa4,
	0x07, 0x24, 0xfa, 0x1f, 0x96, 0xcc, 0xd7, 0x20, 0x43, 0xe7, 0xe1, 0x51, 0x3a, 0x77, 0x40, 0xfa,
	0xef, 0xf1, 0x09, 0xaf, 0x3f, 0x3b, 0xba, 0x20, 0x03, 0x8f, 0x33, 0x96, 0xcf, 0xab, 0x21, 0x12,
	0xa7, 0x7c, 0x91, 0x7e, 0xcd, 0xb9, 0x47, 0x07, 0x87, 0x19, 0xcb, 0x97, 0xa7, 0x47, 0xc5, 0xfd,
	0x95, 0x29, 0xa2, 0xab, 0x76, 0x63, 0xaf, 0x0f, 0x7f, 0x6f, 0xd7, 0xec, 0xcf, 0x76, 0xcd, 0xfe,
	0x6e, 0xd7, 0xec, 0xdb, 0x78, 0x43, 0xba, 0x9e, 0xa5, 0x1d, 0x79, 0xfe, 0x2f, 0x00, 0x00, 0xff,
	0xff, 0x54, 0xcc, 0x8c, 0x8e, 0xa8, 0x02, 0x00, 0x00,
}

func (m *Group) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Group) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Group) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.OwnerUser != nil {
		{
			size, err := m.OwnerUser.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGroup(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x8a
	}
	if m.IsNewOrigin {
		i--
		if m.IsNewOrigin {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x80
	}
	if m.JoinStatus != 0 {
		i = encodeVarintGroup(dAtA, i, uint64(m.JoinStatus))
		i--
		dAtA[i] = 0x78
	}
	if len(m.ShareH5Url) > 0 {
		i -= len(m.ShareH5Url)
		copy(dAtA[i:], m.ShareH5Url)
		i = encodeVarintGroup(dAtA, i, uint64(len(m.ShareH5Url)))
		i--
		dAtA[i] = 0x72
	}
	if len(m.QrCodeUrl) > 0 {
		i -= len(m.QrCodeUrl)
		copy(dAtA[i:], m.QrCodeUrl)
		i = encodeVarintGroup(dAtA, i, uint64(len(m.QrCodeUrl)))
		i--
		dAtA[i] = 0x6a
	}
	if len(m.Broadcast) > 0 {
		i -= len(m.Broadcast)
		copy(dAtA[i:], m.Broadcast)
		i = encodeVarintGroup(dAtA, i, uint64(len(m.Broadcast)))
		i--
		dAtA[i] = 0x62
	}
	if m.Number != 0 {
		i = encodeVarintGroup(dAtA, i, uint64(m.Number))
		i--
		dAtA[i] = 0x58
	}
	if len(m.Content) > 0 {
		i -= len(m.Content)
		copy(dAtA[i:], m.Content)
		i = encodeVarintGroup(dAtA, i, uint64(len(m.Content)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.AvatarList) > 0 {
		for iNdEx := len(m.AvatarList) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AvatarList[iNdEx])
			copy(dAtA[i:], m.AvatarList[iNdEx])
			i = encodeVarintGroup(dAtA, i, uint64(len(m.AvatarList[iNdEx])))
			i--
			dAtA[i] = 0x4a
		}
	}
	if m.InsertTime != 0 {
		i = encodeVarintGroup(dAtA, i, uint64(m.InsertTime))
		i--
		dAtA[i] = 0x40
	}
	if m.IsForbidden != 0 {
		i = encodeVarintGroup(dAtA, i, uint64(m.IsForbidden))
		i--
		dAtA[i] = 0x38
	}
	if m.IsPublic != 0 {
		i = encodeVarintGroup(dAtA, i, uint64(m.IsPublic))
		i--
		dAtA[i] = 0x30
	}
	if m.BusinessId != 0 {
		i = encodeVarintGroup(dAtA, i, uint64(m.BusinessId))
		i--
		dAtA[i] = 0x28
	}
	if m.BusinessType != 0 {
		i = encodeVarintGroup(dAtA, i, uint64(m.BusinessType))
		i--
		dAtA[i] = 0x20
	}
	if m.Owner != 0 {
		i = encodeVarintGroup(dAtA, i, uint64(m.Owner))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintGroup(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintGroup(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGroup(dAtA []byte, offset int, v uint64) int {
	offset -= sovGroup(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Group) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovGroup(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovGroup(uint64(l))
	}
	if m.Owner != 0 {
		n += 1 + sovGroup(uint64(m.Owner))
	}
	if m.BusinessType != 0 {
		n += 1 + sovGroup(uint64(m.BusinessType))
	}
	if m.BusinessId != 0 {
		n += 1 + sovGroup(uint64(m.BusinessId))
	}
	if m.IsPublic != 0 {
		n += 1 + sovGroup(uint64(m.IsPublic))
	}
	if m.IsForbidden != 0 {
		n += 1 + sovGroup(uint64(m.IsForbidden))
	}
	if m.InsertTime != 0 {
		n += 1 + sovGroup(uint64(m.InsertTime))
	}
	if len(m.AvatarList) > 0 {
		for _, s := range m.AvatarList {
			l = len(s)
			n += 1 + l + sovGroup(uint64(l))
		}
	}
	l = len(m.Content)
	if l > 0 {
		n += 1 + l + sovGroup(uint64(l))
	}
	if m.Number != 0 {
		n += 1 + sovGroup(uint64(m.Number))
	}
	l = len(m.Broadcast)
	if l > 0 {
		n += 1 + l + sovGroup(uint64(l))
	}
	l = len(m.QrCodeUrl)
	if l > 0 {
		n += 1 + l + sovGroup(uint64(l))
	}
	l = len(m.ShareH5Url)
	if l > 0 {
		n += 1 + l + sovGroup(uint64(l))
	}
	if m.JoinStatus != 0 {
		n += 1 + sovGroup(uint64(m.JoinStatus))
	}
	if m.IsNewOrigin {
		n += 3
	}
	if m.OwnerUser != nil {
		l = m.OwnerUser.Size()
		n += 2 + l + sovGroup(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovGroup(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGroup(x uint64) (n int) {
	return sovGroup(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Group) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGroup
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
			return fmt.Errorf("proto: Group: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Group: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
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
				return ErrInvalidLengthGroup
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			m.Owner = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Owner |= uint32(b&0x7F) << shift
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
					return ErrIntOverflowGroup
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
					return ErrIntOverflowGroup
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsPublic", wireType)
			}
			m.IsPublic = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IsPublic |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsForbidden", wireType)
			}
			m.IsForbidden = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IsForbidden |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InsertTime", wireType)
			}
			m.InsertTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.InsertTime |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AvatarList", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
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
				return ErrInvalidLengthGroup
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AvatarList = append(m.AvatarList, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
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
				return ErrInvalidLengthGroup
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Content = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Number", wireType)
			}
			m.Number = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Number |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Broadcast", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
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
				return ErrInvalidLengthGroup
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Broadcast = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QrCodeUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
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
				return ErrInvalidLengthGroup
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.QrCodeUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShareH5Url", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
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
				return ErrInvalidLengthGroup
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ShareH5Url = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field JoinStatus", wireType)
			}
			m.JoinStatus = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.JoinStatus |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsNewOrigin", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsNewOrigin = bool(v != 0)
		case 17:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerUser", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGroup
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.OwnerUser == nil {
				m.OwnerUser = &User{}
			}
			if err := m.OwnerUser.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGroup(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGroup
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGroup
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
func skipGroup(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGroup
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
					return 0, ErrIntOverflowGroup
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
					return 0, ErrIntOverflowGroup
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
				return 0, ErrInvalidLengthGroup
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGroup
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGroup
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGroup        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGroup          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGroup = fmt.Errorf("proto: unexpected end of group")
)
