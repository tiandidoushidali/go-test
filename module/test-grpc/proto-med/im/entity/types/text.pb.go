// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: med-common/app/service/med-xim/api/im/entity/types/text.proto

package xim

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	misc "go-test/module/test-grpc/proto-med/im/entity/misc"
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

// 文本消息
type Text struct {
	Content              string       `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	Anchors              []*Anchor    `protobuf:"bytes,2,rep,name=anchors,proto3" json:"anchors,omitempty"`
	AtUsers              []*misc.User `protobuf:"bytes,3,rep,name=atUsers,proto3" json:"atUsers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Text) Reset()         { *m = Text{} }
func (m *Text) String() string { return proto.CompactTextString(m) }
func (*Text) ProtoMessage()    {}
func (*Text) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c3d4086a4fab9e0, []int{0}
}
func (m *Text) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Text) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Text.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Text) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Text.Merge(m, src)
}
func (m *Text) XXX_Size() int {
	return m.Size()
}
func (m *Text) XXX_DiscardUnknown() {
	xxx_messageInfo_Text.DiscardUnknown(m)
}

var xxx_messageInfo_Text proto.InternalMessageInfo

func (m *Text) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Text) GetAnchors() []*Anchor {
	if m != nil {
		return m.Anchors
	}
	return nil
}

func (m *Text) GetAtUsers() []*misc.User {
	if m != nil {
		return m.AtUsers
	}
	return nil
}

func init() {
	proto.RegisterType((*Text)(nil), "med.xim.entity.types.Text")
}

func init() {
	proto.RegisterFile("med-common/app/service/med-xim/api/im/entity/types/text.proto", fileDescriptor_3c3d4086a4fab9e0)
}

var fileDescriptor_3c3d4086a4fab9e0 = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x8f, 0xb1, 0x4a, 0xc4, 0x40,
	0x10, 0x86, 0x59, 0x23, 0x06, 0xd7, 0xca, 0x60, 0x11, 0x0f, 0x09, 0x87, 0xd5, 0x35, 0xee, 0x80,
	0x07, 0x56, 0x8a, 0xe8, 0x23, 0x04, 0x6d, 0xec, 0xe2, 0xde, 0x80, 0x53, 0xcc, 0xee, 0x92, 0x1d,
	0x65, 0xef, 0x21, 0x7c, 0x2f, 0x4b, 0x1f, 0x41, 0xf2, 0x24, 0x92, 0x8d, 0x69, 0xc4, 0xc6, 0x2b,
	0x67, 0xf8, 0xfe, 0x6f, 0xfe, 0xd1, 0x37, 0x8c, 0x9b, 0x0b, 0xeb, 0x99, 0xbd, 0x83, 0x2e, 0x04,
	0x88, 0xd8, 0xbf, 0x91, 0x45, 0x18, 0xd7, 0x89, 0x18, 0xba, 0x40, 0x40, 0x0c, 0xe8, 0x84, 0x64,
	0x0b, 0xb2, 0x0d, 0x18, 0x41, 0x30, 0x89, 0x09, 0xbd, 0x17, 0x5f, 0x9d, 0x30, 0x6e, 0x4c, 0x22,
	0x36, 0x13, 0x60, 0x32, 0xb0, 0xb8, 0xdd, 0x41, 0xda, 0x39, 0xfb, 0xe2, 0xfb, 0x49, 0xbb, 0xb8,
	0xfe, 0x97, 0x80, 0x29, 0x5a, 0x78, 0x8d, 0xf8, 0x93, 0x3e, 0x7f, 0x57, 0x7a, 0xff, 0x01, 0x93,
	0x54, 0xb5, 0x2e, 0xad, 0x77, 0x82, 0x4e, 0x6a, 0xb5, 0x54, 0xab, 0xc3, 0x76, 0x1e, 0xab, 0x2b,
	0x5d, 0x4e, 0x07, 0x63, 0xbd, 0xb7, 0x2c, 0x56, 0x47, 0x97, 0x67, 0xe6, 0xaf, 0x4f, 0xcc, 0x5d,
	0x86, 0xda, 0x19, 0xae, 0xd6, 0xba, 0xec, 0xe4, 0x31, 0x62, 0x1f, 0xeb, 0x22, 0xe7, 0x4e, 0x7f,
	0xe7, 0xc6, 0x32, 0x66, 0x24, 0xda, 0x99, 0xbc, 0x3f, 0xfe, 0x18, 0x1a, 0xf5, 0x39, 0x34, 0xea,
	0x6b, 0x68, 0xd4, 0x53, 0x91, 0x88, 0x9f, 0x0f, 0x72, 0xd3, 0xf5, 0x77, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x68, 0xa1, 0x6f, 0xfa, 0x7f, 0x01, 0x00, 0x00,
}

func (m *Text) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Text) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Text) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.AtUsers) > 0 {
		for iNdEx := len(m.AtUsers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AtUsers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintText(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Anchors) > 0 {
		for iNdEx := len(m.Anchors) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Anchors[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintText(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Content) > 0 {
		i -= len(m.Content)
		copy(dAtA[i:], m.Content)
		i = encodeVarintText(dAtA, i, uint64(len(m.Content)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintText(dAtA []byte, offset int, v uint64) int {
	offset -= sovText(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Text) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Content)
	if l > 0 {
		n += 1 + l + sovText(uint64(l))
	}
	if len(m.Anchors) > 0 {
		for _, e := range m.Anchors {
			l = e.Size()
			n += 1 + l + sovText(uint64(l))
		}
	}
	if len(m.AtUsers) > 0 {
		for _, e := range m.AtUsers {
			l = e.Size()
			n += 1 + l + sovText(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovText(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozText(x uint64) (n int) {
	return sovText(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Text) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowText
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
			return fmt.Errorf("proto: Text: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Text: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowText
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
				return ErrInvalidLengthText
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthText
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Content = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Anchors", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowText
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
				return ErrInvalidLengthText
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthText
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Anchors = append(m.Anchors, &Anchor{})
			if err := m.Anchors[len(m.Anchors)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AtUsers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowText
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
				return ErrInvalidLengthText
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthText
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AtUsers = append(m.AtUsers, &misc.User{})
			if err := m.AtUsers[len(m.AtUsers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipText(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthText
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthText
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
func skipText(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowText
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
					return 0, ErrIntOverflowText
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
					return 0, ErrIntOverflowText
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
				return 0, ErrInvalidLengthText
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupText
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthText
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthText        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowText          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupText = fmt.Errorf("proto: unexpected end of group")
)