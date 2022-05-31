// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: med-common/app/service/med-xim/api/im/entity/types/json.proto

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

type Json struct {
	Content              string   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	Type                 int32    `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Json) Reset()         { *m = Json{} }
func (m *Json) String() string { return proto.CompactTextString(m) }
func (*Json) ProtoMessage()    {}
func (*Json) Descriptor() ([]byte, []int) {
	return fileDescriptor_432e456b854b7155, []int{0}
}
func (m *Json) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Json) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Json.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Json) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Json.Merge(m, src)
}
func (m *Json) XXX_Size() int {
	return m.Size()
}
func (m *Json) XXX_DiscardUnknown() {
	xxx_messageInfo_Json.DiscardUnknown(m)
}

var xxx_messageInfo_Json proto.InternalMessageInfo

func (m *Json) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Json) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func init() {
	proto.RegisterType((*Json)(nil), "med.xim.entity.types.Json")
}

func init() {
	proto.RegisterFile("med-common/app/service/med-xim/api/im/entity/types/json.proto", fileDescriptor_432e456b854b7155)
}

var fileDescriptor_432e456b854b7155 = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0xcd, 0x4d, 0x4d, 0xd1,
	0x4d, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0xd3, 0x4f, 0x2c, 0x28, 0xd0, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb,
	0x4c, 0x4e, 0xd5, 0x07, 0x09, 0x57, 0x64, 0xe6, 0xea, 0x27, 0x16, 0x64, 0xea, 0x67, 0xe6, 0xea,
	0xa7, 0xe6, 0x95, 0x64, 0x96, 0x54, 0xea, 0x97, 0x54, 0x16, 0xa4, 0x16, 0xeb, 0x67, 0x15, 0xe7,
	0xe7, 0xe9, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x89, 0xe4, 0xa6, 0xa6, 0xe8, 0x55, 0x64, 0xe6,
	0xea, 0x41, 0x14, 0xe8, 0x81, 0x15, 0x28, 0x99, 0x70, 0xb1, 0x78, 0x15, 0xe7, 0xe7, 0x09, 0x49,
	0x70, 0xb1, 0x27, 0xe7, 0xe7, 0x95, 0xa4, 0xe6, 0x95, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06,
	0xc1, 0xb8, 0x42, 0x42, 0x5c, 0x2c, 0x20, 0xa5, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x60,
	0xb6, 0x93, 0xe0, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x18,
	0xc5, 0x5c, 0x91, 0x99, 0x9b, 0xc4, 0x06, 0xb6, 0xc5, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x54,
	0x3f, 0x6e, 0xac, 0xa6, 0x00, 0x00, 0x00,
}

func (m *Json) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Json) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Json) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Type != 0 {
		i = encodeVarintJson(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Content) > 0 {
		i -= len(m.Content)
		copy(dAtA[i:], m.Content)
		i = encodeVarintJson(dAtA, i, uint64(len(m.Content)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintJson(dAtA []byte, offset int, v uint64) int {
	offset -= sovJson(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Json) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Content)
	if l > 0 {
		n += 1 + l + sovJson(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovJson(uint64(m.Type))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovJson(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozJson(x uint64) (n int) {
	return sovJson(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Json) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowJson
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
			return fmt.Errorf("proto: Json: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Json: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJson
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
				return ErrInvalidLengthJson
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthJson
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Content = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJson
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
		default:
			iNdEx = preIndex
			skippy, err := skipJson(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthJson
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthJson
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
func skipJson(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowJson
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
					return 0, ErrIntOverflowJson
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
					return 0, ErrIntOverflowJson
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
				return 0, ErrInvalidLengthJson
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupJson
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthJson
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthJson        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowJson          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupJson = fmt.Errorf("proto: unexpected end of group")
)
