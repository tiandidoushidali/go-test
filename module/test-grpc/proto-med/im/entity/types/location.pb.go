// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: med-common/app/service/med-xim/api/im/entity/types/location.proto

package xim

import (
	encoding_binary "encoding/binary"
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

// 位置消息
type Location struct {
	Longitude            float64  `protobuf:"fixed64,1,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Latitude             float64  `protobuf:"fixed64,2,opt,name=latitude,proto3" json:"latitude,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Location) Reset()         { *m = Location{} }
func (m *Location) String() string { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()    {}
func (*Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_d19b8a7120bf93c6, []int{0}
}
func (m *Location) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Location) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Location.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Location) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Location.Merge(m, src)
}
func (m *Location) XXX_Size() int {
	return m.Size()
}
func (m *Location) XXX_DiscardUnknown() {
	xxx_messageInfo_Location.DiscardUnknown(m)
}

var xxx_messageInfo_Location proto.InternalMessageInfo

func (m *Location) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *Location) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func init() {
	proto.RegisterType((*Location)(nil), "med.xim.entity.types.Location")
}

func init() {
	proto.RegisterFile("med-common/app/service/med-xim/api/im/entity/types/location.proto", fileDescriptor_d19b8a7120bf93c6)
}

var fileDescriptor_d19b8a7120bf93c6 = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0xcc, 0x4d, 0x4d, 0xd1,
	0x4d, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0xd3, 0x4f, 0x2c, 0x28, 0xd0, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb,
	0x4c, 0x4e, 0xd5, 0x07, 0x09, 0x57, 0x64, 0xe6, 0xea, 0x27, 0x16, 0x64, 0xea, 0x67, 0xe6, 0xea,
	0xa7, 0xe6, 0x95, 0x64, 0x96, 0x54, 0xea, 0x97, 0x54, 0x16, 0xa4, 0x16, 0xeb, 0xe7, 0xe4, 0x27,
	0x27, 0x96, 0x64, 0xe6, 0xe7, 0xe9, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x89, 0xe4, 0xa6, 0xa6,
	0xe8, 0x55, 0x64, 0xe6, 0xea, 0x41, 0x14, 0xe9, 0x81, 0x15, 0x29, 0xb9, 0x70, 0x71, 0xf8, 0x40,
	0xd5, 0x09, 0xc9, 0x70, 0x71, 0xe6, 0xe4, 0xe7, 0xa5, 0x67, 0x96, 0x94, 0xa6, 0xa4, 0x4a, 0x30,
	0x2a, 0x30, 0x6a, 0x30, 0x06, 0x21, 0x04, 0x84, 0xa4, 0xb8, 0x38, 0x72, 0x12, 0x4b, 0x20, 0x92,
	0x4c, 0x60, 0x49, 0x38, 0xdf, 0x49, 0xf0, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f,
	0x3c, 0x92, 0x63, 0x8c, 0x62, 0xae, 0xc8, 0xcc, 0x4d, 0x62, 0x03, 0xdb, 0x6a, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0x06, 0x29, 0xb6, 0x51, 0xba, 0x00, 0x00, 0x00,
}

func (m *Location) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Location) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Location) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Latitude != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Latitude))))
		i--
		dAtA[i] = 0x11
	}
	if m.Longitude != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Longitude))))
		i--
		dAtA[i] = 0x9
	}
	return len(dAtA) - i, nil
}

func encodeVarintLocation(dAtA []byte, offset int, v uint64) int {
	offset -= sovLocation(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Location) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Longitude != 0 {
		n += 9
	}
	if m.Latitude != 0 {
		n += 9
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovLocation(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLocation(x uint64) (n int) {
	return sovLocation(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Location) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLocation
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
			return fmt.Errorf("proto: Location: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Location: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Longitude", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Longitude = float64(math.Float64frombits(v))
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Latitude", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Latitude = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipLocation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLocation
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthLocation
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
func skipLocation(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLocation
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
					return 0, ErrIntOverflowLocation
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
					return 0, ErrIntOverflowLocation
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
				return 0, ErrInvalidLengthLocation
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLocation
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLocation
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLocation        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLocation          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLocation = fmt.Errorf("proto: unexpected end of group")
)
