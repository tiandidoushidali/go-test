// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: med-common/app/service/med-xim/api/im/entity/commands/group_operation.proto

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

type GroupOperation struct {
	Operation            int32        `protobuf:"varint,1,opt,name=operation,proto3" json:"operation,omitempty"`
	Group                int64        `protobuf:"varint,2,opt,name=group,proto3" json:"group,omitempty"`
	Members              []*misc.User `protobuf:"bytes,3,rep,name=members,proto3" json:"members,omitempty"`
	Content              string       `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GroupOperation) Reset()         { *m = GroupOperation{} }
func (m *GroupOperation) String() string { return proto.CompactTextString(m) }
func (*GroupOperation) ProtoMessage()    {}
func (*GroupOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb095d92100260c4, []int{0}
}
func (m *GroupOperation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GroupOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GroupOperation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GroupOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupOperation.Merge(m, src)
}
func (m *GroupOperation) XXX_Size() int {
	return m.Size()
}
func (m *GroupOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupOperation.DiscardUnknown(m)
}

var xxx_messageInfo_GroupOperation proto.InternalMessageInfo

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

func (m *GroupOperation) GetMembers() []*misc.User {
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
	proto.RegisterType((*GroupOperation)(nil), "med.xim.entity.commands.GroupOperation")
}

func init() {
	proto.RegisterFile("med-common/app/service/med-xim/api/im/entity/commands/group_operation.proto", fileDescriptor_bb095d92100260c4)
}

var fileDescriptor_bb095d92100260c4 = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x8f, 0xb1, 0x4a, 0xc4, 0x40,
	0x10, 0x86, 0x59, 0xe3, 0x79, 0xdc, 0x0a, 0x16, 0x8b, 0xe0, 0x2a, 0x12, 0x82, 0x55, 0x10, 0xdc,
	0x01, 0xaf, 0xb5, 0xb2, 0xb1, 0xb0, 0x10, 0x02, 0x36, 0x36, 0x92, 0x4b, 0x06, 0x99, 0x62, 0x76,
	0x97, 0xdd, 0x3d, 0x89, 0xcf, 0xe1, 0x4b, 0x59, 0xfa, 0x08, 0x92, 0x27, 0x91, 0x24, 0x44, 0x41,
	0xb0, 0xb8, 0x72, 0x7e, 0x3e, 0xe6, 0xfb, 0x7f, 0x79, 0xcf, 0xd8, 0x5e, 0x35, 0x8e, 0xd9, 0x59,
	0xa8, 0xbd, 0x87, 0x88, 0xe1, 0x95, 0x1a, 0x84, 0x21, 0xee, 0x88, 0xa1, 0xf6, 0x04, 0xc4, 0x80,
	0x36, 0x51, 0x7a, 0x83, 0x01, 0xac, 0x6d, 0x1b, 0xe1, 0x25, 0xb8, 0xad, 0x7f, 0x76, 0x1e, 0x43,
	0x9d, 0xc8, 0x59, 0xe3, 0x83, 0x4b, 0x4e, 0x9d, 0x30, 0xb6, 0xa6, 0x23, 0x36, 0x13, 0x6e, 0x66,
	0xfc, 0xec, 0x66, 0x27, 0x0b, 0x53, 0x6c, 0x60, 0x1b, 0x31, 0x4c, 0x6f, 0x2f, 0xde, 0x85, 0x3c,
	0xba, 0x1b, 0x84, 0x0f, 0xb3, 0x4f, 0x9d, 0xcb, 0xd5, 0x8f, 0x5c, 0x8b, 0x42, 0x94, 0x8b, 0xea,
	0x37, 0x50, 0xc7, 0x72, 0x31, 0x16, 0xd4, 0x7b, 0x85, 0x28, 0xb3, 0x6a, 0x3a, 0xd4, 0x5a, 0x2e,
	0x19, 0x79, 0x83, 0x21, 0xea, 0xac, 0xc8, 0xca, 0xc3, 0xeb, 0x53, 0xf3, 0xa7, 0xef, 0x20, 0x36,
	0x8f, 0x11, 0x43, 0x35, 0x93, 0x4a, 0xcb, 0x65, 0xe3, 0x6c, 0x42, 0x9b, 0xf4, 0x7e, 0x21, 0xca,
	0x55, 0x35, 0x9f, 0xb7, 0x97, 0x1f, 0x7d, 0x2e, 0x3e, 0xfb, 0x5c, 0x7c, 0xf5, 0xb9, 0x90, 0xff,
	0x4d, 0x7f, 0xca, 0x3a, 0xe2, 0xcd, 0xc1, 0x38, 0x64, 0xfd, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x42,
	0x0d, 0x11, 0xad, 0x6e, 0x01, 0x00, 0x00,
}

func (m *GroupOperation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GroupOperation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GroupOperation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
		i = encodeVarintGroupOperation(dAtA, i, uint64(len(m.Content)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Members) > 0 {
		for iNdEx := len(m.Members) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Members[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGroupOperation(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Group != 0 {
		i = encodeVarintGroupOperation(dAtA, i, uint64(m.Group))
		i--
		dAtA[i] = 0x10
	}
	if m.Operation != 0 {
		i = encodeVarintGroupOperation(dAtA, i, uint64(m.Operation))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGroupOperation(dAtA []byte, offset int, v uint64) int {
	offset -= sovGroupOperation(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GroupOperation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Operation != 0 {
		n += 1 + sovGroupOperation(uint64(m.Operation))
	}
	if m.Group != 0 {
		n += 1 + sovGroupOperation(uint64(m.Group))
	}
	if len(m.Members) > 0 {
		for _, e := range m.Members {
			l = e.Size()
			n += 1 + l + sovGroupOperation(uint64(l))
		}
	}
	l = len(m.Content)
	if l > 0 {
		n += 1 + l + sovGroupOperation(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovGroupOperation(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGroupOperation(x uint64) (n int) {
	return sovGroupOperation(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GroupOperation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGroupOperation
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
			return fmt.Errorf("proto: GroupOperation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GroupOperation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Operation", wireType)
			}
			m.Operation = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroupOperation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Operation |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Group", wireType)
			}
			m.Group = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroupOperation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Group |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Members", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroupOperation
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
				return ErrInvalidLengthGroupOperation
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGroupOperation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Members = append(m.Members, &misc.User{})
			if err := m.Members[len(m.Members)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroupOperation
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
				return ErrInvalidLengthGroupOperation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGroupOperation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Content = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGroupOperation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGroupOperation
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGroupOperation
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
func skipGroupOperation(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGroupOperation
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
					return 0, ErrIntOverflowGroupOperation
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
					return 0, ErrIntOverflowGroupOperation
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
				return 0, ErrInvalidLengthGroupOperation
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGroupOperation
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGroupOperation
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGroupOperation        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGroupOperation          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGroupOperation = fmt.Errorf("proto: unexpected end of group")
)
