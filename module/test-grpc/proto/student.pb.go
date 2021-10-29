// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: module/test-grpc/proto/student.proto

package test_student_v1

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type StudentInfoReq struct {
	Id                   int64    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty" form:"id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StudentInfoReq) Reset()         { *m = StudentInfoReq{} }
func (m *StudentInfoReq) String() string { return proto.CompactTextString(m) }
func (*StudentInfoReq) ProtoMessage()    {}
func (*StudentInfoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5ebc71363862919, []int{0}
}
func (m *StudentInfoReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StudentInfoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StudentInfoReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StudentInfoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StudentInfoReq.Merge(m, src)
}
func (m *StudentInfoReq) XXX_Size() int {
	return m.Size()
}
func (m *StudentInfoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_StudentInfoReq.DiscardUnknown(m)
}

var xxx_messageInfo_StudentInfoReq proto.InternalMessageInfo

func (m *StudentInfoReq) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type StudentInfoResp struct {
	Id                   int64    `protobuf:"varint,1,opt,name=Id,proto3" json:"id"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"name"`
	Age                  int64    `protobuf:"varint,3,opt,name=age,proto3" json:"age"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StudentInfoResp) Reset()         { *m = StudentInfoResp{} }
func (m *StudentInfoResp) String() string { return proto.CompactTextString(m) }
func (*StudentInfoResp) ProtoMessage()    {}
func (*StudentInfoResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5ebc71363862919, []int{1}
}
func (m *StudentInfoResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StudentInfoResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StudentInfoResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StudentInfoResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StudentInfoResp.Merge(m, src)
}
func (m *StudentInfoResp) XXX_Size() int {
	return m.Size()
}
func (m *StudentInfoResp) XXX_DiscardUnknown() {
	xxx_messageInfo_StudentInfoResp.DiscardUnknown(m)
}

var xxx_messageInfo_StudentInfoResp proto.InternalMessageInfo

func (m *StudentInfoResp) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *StudentInfoResp) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StudentInfoResp) GetAge() int64 {
	if m != nil {
		return m.Age
	}
	return 0
}

type StudentListReq struct {
	Page                 int64    `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty" form:"page"`
	PageSize             int64    `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty" form:"pageSize`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StudentListReq) Reset()         { *m = StudentListReq{} }
func (m *StudentListReq) String() string { return proto.CompactTextString(m) }
func (*StudentListReq) ProtoMessage()    {}
func (*StudentListReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5ebc71363862919, []int{2}
}
func (m *StudentListReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StudentListReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StudentListReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StudentListReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StudentListReq.Merge(m, src)
}
func (m *StudentListReq) XXX_Size() int {
	return m.Size()
}
func (m *StudentListReq) XXX_DiscardUnknown() {
	xxx_messageInfo_StudentListReq.DiscardUnknown(m)
}

var xxx_messageInfo_StudentListReq proto.InternalMessageInfo

func (m *StudentListReq) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *StudentListReq) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

type StudentListResp struct {
	List                 []*StudentEntity `protobuf:"bytes,1,rep,name=list,proto3" json:"list"`
	Total                int64            `protobuf:"varint,2,opt,name=total,proto3" json:"total"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *StudentListResp) Reset()         { *m = StudentListResp{} }
func (m *StudentListResp) String() string { return proto.CompactTextString(m) }
func (*StudentListResp) ProtoMessage()    {}
func (*StudentListResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5ebc71363862919, []int{3}
}
func (m *StudentListResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StudentListResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StudentListResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StudentListResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StudentListResp.Merge(m, src)
}
func (m *StudentListResp) XXX_Size() int {
	return m.Size()
}
func (m *StudentListResp) XXX_DiscardUnknown() {
	xxx_messageInfo_StudentListResp.DiscardUnknown(m)
}

var xxx_messageInfo_StudentListResp proto.InternalMessageInfo

func (m *StudentListResp) GetList() []*StudentEntity {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *StudentListResp) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

type StudentEntity struct {
	Id                   int64    `protobuf:"varint,1,opt,name=Id,proto3" json:"id"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"name"`
	Age                  int64    `protobuf:"varint,3,opt,name=Age,proto3" json:"age"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StudentEntity) Reset()         { *m = StudentEntity{} }
func (m *StudentEntity) String() string { return proto.CompactTextString(m) }
func (*StudentEntity) ProtoMessage()    {}
func (*StudentEntity) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5ebc71363862919, []int{4}
}
func (m *StudentEntity) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StudentEntity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StudentEntity.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StudentEntity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StudentEntity.Merge(m, src)
}
func (m *StudentEntity) XXX_Size() int {
	return m.Size()
}
func (m *StudentEntity) XXX_DiscardUnknown() {
	xxx_messageInfo_StudentEntity.DiscardUnknown(m)
}

var xxx_messageInfo_StudentEntity proto.InternalMessageInfo

func (m *StudentEntity) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *StudentEntity) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StudentEntity) GetAge() int64 {
	if m != nil {
		return m.Age
	}
	return 0
}

func init() {
	proto.RegisterType((*StudentInfoReq)(nil), "test.student.v1.StudentInfoReq")
	proto.RegisterType((*StudentInfoResp)(nil), "test.student.v1.StudentInfoResp")
	proto.RegisterType((*StudentListReq)(nil), "test.student.v1.StudentListReq")
	proto.RegisterType((*StudentListResp)(nil), "test.student.v1.StudentListResp")
	proto.RegisterType((*StudentEntity)(nil), "test.student.v1.StudentEntity")
}

func init() {
	proto.RegisterFile("module/test-grpc/proto/student.proto", fileDescriptor_a5ebc71363862919)
}

var fileDescriptor_a5ebc71363862919 = []byte{
	// 400 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0xcd, 0xae, 0xd2, 0x40,
	0x14, 0xce, 0xb4, 0x95, 0x9f, 0x41, 0x20, 0x99, 0x85, 0xa9, 0x44, 0x3b, 0xcd, 0xe8, 0x82, 0x0d,
	0x6d, 0xc4, 0x9d, 0x71, 0x43, 0x13, 0x17, 0x24, 0xc4, 0x98, 0xe1, 0x05, 0x6c, 0xe9, 0x50, 0x9b,
	0xd0, 0x4e, 0xa5, 0x53, 0x13, 0x7d, 0x27, 0xdf, 0xc3, 0xa5, 0x4f, 0xd0, 0x18, 0x96, 0x5d, 0xf2,
	0x04, 0x66, 0x66, 0x0a, 0x17, 0x6e, 0x2e, 0x77, 0x71, 0x37, 0x73, 0x7a, 0xce, 0xf9, 0xbe, 0xf3,
	0x9d, 0x9f, 0xc2, 0xb7, 0x19, 0x8f, 0xab, 0x1d, 0xf3, 0x05, 0x2b, 0xc5, 0x2c, 0xd9, 0x17, 0x1b,
	0xbf, 0xd8, 0x73, 0xc1, 0xfd, 0x52, 0x54, 0x31, 0xcb, 0x85, 0xa7, 0x3c, 0x34, 0x96, 0x69, 0xef,
	0x14, 0xfb, 0xf1, 0x6e, 0x32, 0x4b, 0x52, 0xf1, 0xad, 0x8a, 0xbc, 0x0d, 0xcf, 0xfc, 0x84, 0x27,
	0x5c, 0xb3, 0xa2, 0x6a, 0xab, 0x3c, 0x5d, 0x42, 0x7e, 0x69, 0x3e, 0xf1, 0xe1, 0x68, 0xad, 0xc9,
	0xcb, 0x7c, 0xcb, 0x29, 0xfb, 0x8e, 0x5e, 0x43, 0x63, 0x19, 0xdb, 0xc0, 0x05, 0x53, 0x33, 0x18,
	0x1e, 0x6b, 0xdc, 0xdf, 0xf2, 0x7d, 0xf6, 0x81, 0xa4, 0x31, 0xa1, 0xc6, 0x32, 0x26, 0x11, 0x1c,
	0x5f, 0x11, 0xca, 0x02, 0xbd, 0xb8, 0x60, 0x74, 0x9a, 0x1a, 0x1b, 0x69, 0x2c, 0xa1, 0xe8, 0x15,
	0xb4, 0x3e, 0x87, 0x19, 0xb3, 0x0d, 0x17, 0x4c, 0xfb, 0x41, 0xaf, 0xa9, 0xb1, 0x95, 0x87, 0x19,
	0xa3, 0x2a, 0x8a, 0x5e, 0x42, 0x33, 0x4c, 0x98, 0x6d, 0x2a, 0x5a, 0xb7, 0xa9, 0xb1, 0x74, 0xa9,
	0x7c, 0x08, 0x3b, 0x37, 0xb5, 0x4a, 0x4b, 0x21, 0x9b, 0x7a, 0x03, 0xad, 0x42, 0xa2, 0xb5, 0xc8,
	0xf8, 0x58, 0xe3, 0x81, 0x6e, 0x4b, 0x46, 0x09, 0x55, 0x49, 0xe4, 0xc1, 0x9e, 0xb4, 0xeb, 0xf4,
	0x97, 0xd6, 0x34, 0x03, 0x74, 0xac, 0xf1, 0xe8, 0x0e, 0x28, 0x33, 0xf4, 0x8c, 0x21, 0xc5, 0x79,
	0x14, 0x2d, 0x53, 0x16, 0xe8, 0x23, 0xb4, 0x76, 0x69, 0x29, 0x6c, 0xe0, 0x9a, 0xd3, 0xc1, 0xdc,
	0xf1, 0xee, 0x6d, 0xd7, 0x6b, 0xf1, 0x9f, 0x72, 0x91, 0x8a, 0x9f, 0x7a, 0x24, 0x89, 0xa7, 0xea,
	0x45, 0x18, 0x3e, 0x13, 0x5c, 0x84, 0xbb, 0x56, 0xbd, 0xdf, 0xd4, 0x58, 0x07, 0xa8, 0x36, 0xe4,
	0x2b, 0x1c, 0x5e, 0x55, 0x78, 0xfa, 0xea, 0x16, 0x0f, 0xac, 0x6e, 0x91, 0xb0, 0xf9, 0x6f, 0x00,
	0xbb, 0xad, 0x04, 0xfa, 0x02, 0x07, 0x17, 0xa7, 0x42, 0xf8, 0xd6, 0x34, 0xed, 0xe5, 0x27, 0xee,
	0xe3, 0x80, 0xb2, 0xb8, 0xa8, 0xb8, 0x52, 0xf3, 0xde, 0x22, 0xb4, 0x67, 0xbb, 0x5d, 0xf1, 0xb4,
	0xf0, 0xe0, 0xf9, 0x9f, 0x83, 0x03, 0xfe, 0x1e, 0x1c, 0xf0, 0xef, 0xe0, 0x80, 0xa8, 0xa3, 0x7e,
	0xca, 0xf7, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x66, 0xdc, 0xcc, 0x21, 0xfc, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StudentClient is the client API for Student service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StudentClient interface {
	// 学生详情
	StudentInfo(ctx context.Context, in *StudentInfoReq, opts ...grpc.CallOption) (*StudentInfoResp, error)
	// 学生列表
	StudentList(ctx context.Context, in *StudentListReq, opts ...grpc.CallOption) (*StudentListResp, error)
}

type studentClient struct {
	cc *grpc.ClientConn
}

func NewStudentClient(cc *grpc.ClientConn) StudentClient {
	return &studentClient{cc}
}

func (c *studentClient) StudentInfo(ctx context.Context, in *StudentInfoReq, opts ...grpc.CallOption) (*StudentInfoResp, error) {
	out := new(StudentInfoResp)
	err := c.cc.Invoke(ctx, "/test.student.v1.Student/StudentInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentClient) StudentList(ctx context.Context, in *StudentListReq, opts ...grpc.CallOption) (*StudentListResp, error) {
	out := new(StudentListResp)
	err := c.cc.Invoke(ctx, "/test.student.v1.Student/StudentList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StudentServer is the server API for Student service.
type StudentServer interface {
	// 学生详情
	StudentInfo(context.Context, *StudentInfoReq) (*StudentInfoResp, error)
	// 学生列表
	StudentList(context.Context, *StudentListReq) (*StudentListResp, error)
}

// UnimplementedStudentServer can be embedded to have forward compatible implementations.
type UnimplementedStudentServer struct {
}

func (*UnimplementedStudentServer) StudentInfo(ctx context.Context, req *StudentInfoReq) (*StudentInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StudentInfo not implemented")
}
func (*UnimplementedStudentServer) StudentList(ctx context.Context, req *StudentListReq) (*StudentListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StudentList not implemented")
}

func RegisterStudentServer(s *grpc.Server, srv StudentServer) {
	s.RegisterService(&_Student_serviceDesc, srv)
}

func _Student_StudentInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StudentInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServer).StudentInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.student.v1.Student/StudentInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServer).StudentInfo(ctx, req.(*StudentInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Student_StudentList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StudentListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServer).StudentList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.student.v1.Student/StudentList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServer).StudentList(ctx, req.(*StudentListReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Student_serviceDesc = grpc.ServiceDesc{
	ServiceName: "test.student.v1.Student",
	HandlerType: (*StudentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StudentInfo",
			Handler:    _Student_StudentInfo_Handler,
		},
		{
			MethodName: "StudentList",
			Handler:    _Student_StudentList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "module/test-grpc/proto/student.proto",
}

func (m *StudentInfoReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StudentInfoReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StudentInfoReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Id != 0 {
		i = encodeVarintStudent(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *StudentInfoResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StudentInfoResp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StudentInfoResp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Age != 0 {
		i = encodeVarintStudent(dAtA, i, uint64(m.Age))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintStudent(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintStudent(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *StudentListReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StudentListReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StudentListReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.PageSize != 0 {
		i = encodeVarintStudent(dAtA, i, uint64(m.PageSize))
		i--
		dAtA[i] = 0x10
	}
	if m.Page != 0 {
		i = encodeVarintStudent(dAtA, i, uint64(m.Page))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *StudentListResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StudentListResp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StudentListResp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Total != 0 {
		i = encodeVarintStudent(dAtA, i, uint64(m.Total))
		i--
		dAtA[i] = 0x10
	}
	if len(m.List) > 0 {
		for iNdEx := len(m.List) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.List[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintStudent(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *StudentEntity) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StudentEntity) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StudentEntity) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Age != 0 {
		i = encodeVarintStudent(dAtA, i, uint64(m.Age))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintStudent(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintStudent(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintStudent(dAtA []byte, offset int, v uint64) int {
	offset -= sovStudent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StudentInfoReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovStudent(uint64(m.Id))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *StudentInfoResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovStudent(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovStudent(uint64(l))
	}
	if m.Age != 0 {
		n += 1 + sovStudent(uint64(m.Age))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *StudentListReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Page != 0 {
		n += 1 + sovStudent(uint64(m.Page))
	}
	if m.PageSize != 0 {
		n += 1 + sovStudent(uint64(m.PageSize))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *StudentListResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.List) > 0 {
		for _, e := range m.List {
			l = e.Size()
			n += 1 + l + sovStudent(uint64(l))
		}
	}
	if m.Total != 0 {
		n += 1 + sovStudent(uint64(m.Total))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *StudentEntity) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovStudent(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovStudent(uint64(l))
	}
	if m.Age != 0 {
		n += 1 + sovStudent(uint64(m.Age))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovStudent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStudent(x uint64) (n int) {
	return sovStudent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StudentInfoReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStudent
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
			return fmt.Errorf("proto: StudentInfoReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StudentInfoReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStudent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStudent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStudent
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
func (m *StudentInfoResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStudent
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
			return fmt.Errorf("proto: StudentInfoResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StudentInfoResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStudent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int64(b&0x7F) << shift
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
					return ErrIntOverflowStudent
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
				return ErrInvalidLengthStudent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStudent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Age", wireType)
			}
			m.Age = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStudent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Age |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStudent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStudent
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
func (m *StudentListReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStudent
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
			return fmt.Errorf("proto: StudentListReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StudentListReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Page", wireType)
			}
			m.Page = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStudent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Page |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PageSize", wireType)
			}
			m.PageSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStudent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PageSize |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStudent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStudent
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
func (m *StudentListResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStudent
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
			return fmt.Errorf("proto: StudentListResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StudentListResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field List", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStudent
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
				return ErrInvalidLengthStudent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStudent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.List = append(m.List, &StudentEntity{})
			if err := m.List[len(m.List)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Total", wireType)
			}
			m.Total = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStudent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Total |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStudent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStudent
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
func (m *StudentEntity) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStudent
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
			return fmt.Errorf("proto: StudentEntity: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StudentEntity: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStudent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int64(b&0x7F) << shift
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
					return ErrIntOverflowStudent
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
				return ErrInvalidLengthStudent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStudent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Age", wireType)
			}
			m.Age = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStudent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Age |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStudent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStudent
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
func skipStudent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStudent
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
					return 0, ErrIntOverflowStudent
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
					return 0, ErrIntOverflowStudent
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
				return 0, ErrInvalidLengthStudent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStudent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStudent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStudent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStudent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStudent = fmt.Errorf("proto: unexpected end of group")
)
