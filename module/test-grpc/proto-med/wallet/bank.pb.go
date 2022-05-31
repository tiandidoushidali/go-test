// Code generated by protoc-gen-go. DO NOT EDIT.
// source: med-wallet-go/bank.proto

package med_wallet_go

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type BankCardAuthReq struct {
	// 银行卡号
	CardNo string `protobuf:"bytes,1,opt,name=card_no,json=cardNo,proto3" json:"card_no,omitempty"`
	// 身份证
	IdCard string `protobuf:"bytes,2,opt,name=id_card,json=idCard,proto3" json:"id_card,omitempty"`
	// 开户姓名
	RealName             string   `protobuf:"bytes,3,opt,name=real_name,json=realName,proto3" json:"real_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BankCardAuthReq) Reset()         { *m = BankCardAuthReq{} }
func (m *BankCardAuthReq) String() string { return proto.CompactTextString(m) }
func (*BankCardAuthReq) ProtoMessage()    {}
func (*BankCardAuthReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfeb6510847d64f9, []int{0}
}

func (m *BankCardAuthReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BankCardAuthReq.Unmarshal(m, b)
}
func (m *BankCardAuthReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BankCardAuthReq.Marshal(b, m, deterministic)
}
func (m *BankCardAuthReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BankCardAuthReq.Merge(m, src)
}
func (m *BankCardAuthReq) XXX_Size() int {
	return xxx_messageInfo_BankCardAuthReq.Size(m)
}
func (m *BankCardAuthReq) XXX_DiscardUnknown() {
	xxx_messageInfo_BankCardAuthReq.DiscardUnknown(m)
}

var xxx_messageInfo_BankCardAuthReq proto.InternalMessageInfo

func (m *BankCardAuthReq) GetCardNo() string {
	if m != nil {
		return m.CardNo
	}
	return ""
}

func (m *BankCardAuthReq) GetIdCard() string {
	if m != nil {
		return m.IdCard
	}
	return ""
}

func (m *BankCardAuthReq) GetRealName() string {
	if m != nil {
		return m.RealName
	}
	return ""
}

type BankCardAuthResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BankCardAuthResp) Reset()         { *m = BankCardAuthResp{} }
func (m *BankCardAuthResp) String() string { return proto.CompactTextString(m) }
func (*BankCardAuthResp) ProtoMessage()    {}
func (*BankCardAuthResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfeb6510847d64f9, []int{1}
}

func (m *BankCardAuthResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BankCardAuthResp.Unmarshal(m, b)
}
func (m *BankCardAuthResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BankCardAuthResp.Marshal(b, m, deterministic)
}
func (m *BankCardAuthResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BankCardAuthResp.Merge(m, src)
}
func (m *BankCardAuthResp) XXX_Size() int {
	return xxx_messageInfo_BankCardAuthResp.Size(m)
}
func (m *BankCardAuthResp) XXX_DiscardUnknown() {
	xxx_messageInfo_BankCardAuthResp.DiscardUnknown(m)
}

var xxx_messageInfo_BankCardAuthResp proto.InternalMessageInfo

type BindBankCardReq struct {
	// 银行卡号
	CardNo string `protobuf:"bytes,1,opt,name=card_no,json=cardNo,proto3" json:"card_no,omitempty"`
	// 身份证
	IdCard string `protobuf:"bytes,2,opt,name=id_card,json=idCard,proto3" json:"id_card,omitempty"`
	// 开户姓名
	RealName             string   `protobuf:"bytes,3,opt,name=real_name,json=realName,proto3" json:"real_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BindBankCardReq) Reset()         { *m = BindBankCardReq{} }
func (m *BindBankCardReq) String() string { return proto.CompactTextString(m) }
func (*BindBankCardReq) ProtoMessage()    {}
func (*BindBankCardReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfeb6510847d64f9, []int{2}
}

func (m *BindBankCardReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BindBankCardReq.Unmarshal(m, b)
}
func (m *BindBankCardReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BindBankCardReq.Marshal(b, m, deterministic)
}
func (m *BindBankCardReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BindBankCardReq.Merge(m, src)
}
func (m *BindBankCardReq) XXX_Size() int {
	return xxx_messageInfo_BindBankCardReq.Size(m)
}
func (m *BindBankCardReq) XXX_DiscardUnknown() {
	xxx_messageInfo_BindBankCardReq.DiscardUnknown(m)
}

var xxx_messageInfo_BindBankCardReq proto.InternalMessageInfo

func (m *BindBankCardReq) GetCardNo() string {
	if m != nil {
		return m.CardNo
	}
	return ""
}

func (m *BindBankCardReq) GetIdCard() string {
	if m != nil {
		return m.IdCard
	}
	return ""
}

func (m *BindBankCardReq) GetRealName() string {
	if m != nil {
		return m.RealName
	}
	return ""
}

type BindBankCardResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BindBankCardResp) Reset()         { *m = BindBankCardResp{} }
func (m *BindBankCardResp) String() string { return proto.CompactTextString(m) }
func (*BindBankCardResp) ProtoMessage()    {}
func (*BindBankCardResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfeb6510847d64f9, []int{3}
}

func (m *BindBankCardResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BindBankCardResp.Unmarshal(m, b)
}
func (m *BindBankCardResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BindBankCardResp.Marshal(b, m, deterministic)
}
func (m *BindBankCardResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BindBankCardResp.Merge(m, src)
}
func (m *BindBankCardResp) XXX_Size() int {
	return xxx_messageInfo_BindBankCardResp.Size(m)
}
func (m *BindBankCardResp) XXX_DiscardUnknown() {
	xxx_messageInfo_BindBankCardResp.DiscardUnknown(m)
}

var xxx_messageInfo_BindBankCardResp proto.InternalMessageInfo

type UnbindBankCardReq struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnbindBankCardReq) Reset()         { *m = UnbindBankCardReq{} }
func (m *UnbindBankCardReq) String() string { return proto.CompactTextString(m) }
func (*UnbindBankCardReq) ProtoMessage()    {}
func (*UnbindBankCardReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfeb6510847d64f9, []int{4}
}

func (m *UnbindBankCardReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnbindBankCardReq.Unmarshal(m, b)
}
func (m *UnbindBankCardReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnbindBankCardReq.Marshal(b, m, deterministic)
}
func (m *UnbindBankCardReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnbindBankCardReq.Merge(m, src)
}
func (m *UnbindBankCardReq) XXX_Size() int {
	return xxx_messageInfo_UnbindBankCardReq.Size(m)
}
func (m *UnbindBankCardReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UnbindBankCardReq.DiscardUnknown(m)
}

var xxx_messageInfo_UnbindBankCardReq proto.InternalMessageInfo

func (m *UnbindBankCardReq) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type UnbindBankCardResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnbindBankCardResp) Reset()         { *m = UnbindBankCardResp{} }
func (m *UnbindBankCardResp) String() string { return proto.CompactTextString(m) }
func (*UnbindBankCardResp) ProtoMessage()    {}
func (*UnbindBankCardResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfeb6510847d64f9, []int{5}
}

func (m *UnbindBankCardResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnbindBankCardResp.Unmarshal(m, b)
}
func (m *UnbindBankCardResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnbindBankCardResp.Marshal(b, m, deterministic)
}
func (m *UnbindBankCardResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnbindBankCardResp.Merge(m, src)
}
func (m *UnbindBankCardResp) XXX_Size() int {
	return xxx_messageInfo_UnbindBankCardResp.Size(m)
}
func (m *UnbindBankCardResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UnbindBankCardResp.DiscardUnknown(m)
}

var xxx_messageInfo_UnbindBankCardResp proto.InternalMessageInfo

type GetUserCardListReq struct {
	Page                 int64    `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit                int64    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserCardListReq) Reset()         { *m = GetUserCardListReq{} }
func (m *GetUserCardListReq) String() string { return proto.CompactTextString(m) }
func (*GetUserCardListReq) ProtoMessage()    {}
func (*GetUserCardListReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfeb6510847d64f9, []int{6}
}

func (m *GetUserCardListReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserCardListReq.Unmarshal(m, b)
}
func (m *GetUserCardListReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserCardListReq.Marshal(b, m, deterministic)
}
func (m *GetUserCardListReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserCardListReq.Merge(m, src)
}
func (m *GetUserCardListReq) XXX_Size() int {
	return xxx_messageInfo_GetUserCardListReq.Size(m)
}
func (m *GetUserCardListReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserCardListReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserCardListReq proto.InternalMessageInfo

func (m *GetUserCardListReq) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *GetUserCardListReq) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type GetUserCardListResp struct {
	List                 []*GetUserCardListResp_Entity `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Total                int64                         `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *GetUserCardListResp) Reset()         { *m = GetUserCardListResp{} }
func (m *GetUserCardListResp) String() string { return proto.CompactTextString(m) }
func (*GetUserCardListResp) ProtoMessage()    {}
func (*GetUserCardListResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfeb6510847d64f9, []int{7}
}

func (m *GetUserCardListResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserCardListResp.Unmarshal(m, b)
}
func (m *GetUserCardListResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserCardListResp.Marshal(b, m, deterministic)
}
func (m *GetUserCardListResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserCardListResp.Merge(m, src)
}
func (m *GetUserCardListResp) XXX_Size() int {
	return xxx_messageInfo_GetUserCardListResp.Size(m)
}
func (m *GetUserCardListResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserCardListResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserCardListResp proto.InternalMessageInfo

func (m *GetUserCardListResp) GetList() []*GetUserCardListResp_Entity {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *GetUserCardListResp) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

type GetUserCardListResp_Entity struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId               int64    `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	BankName             string   `protobuf:"bytes,3,opt,name=bankName,proto3" json:"bankName,omitempty"`
	BankBranchName       string   `protobuf:"bytes,4,opt,name=bankBranchName,proto3" json:"bankBranchName,omitempty"`
	CardTypeName         string   `protobuf:"bytes,5,opt,name=cardTypeName,proto3" json:"cardTypeName,omitempty"`
	CardNo               string   `protobuf:"bytes,6,opt,name=cardNo,proto3" json:"cardNo,omitempty"`
	CardHolder           string   `protobuf:"bytes,7,opt,name=cardHolder,proto3" json:"cardHolder,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserCardListResp_Entity) Reset()         { *m = GetUserCardListResp_Entity{} }
func (m *GetUserCardListResp_Entity) String() string { return proto.CompactTextString(m) }
func (*GetUserCardListResp_Entity) ProtoMessage()    {}
func (*GetUserCardListResp_Entity) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfeb6510847d64f9, []int{7, 0}
}

func (m *GetUserCardListResp_Entity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserCardListResp_Entity.Unmarshal(m, b)
}
func (m *GetUserCardListResp_Entity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserCardListResp_Entity.Marshal(b, m, deterministic)
}
func (m *GetUserCardListResp_Entity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserCardListResp_Entity.Merge(m, src)
}
func (m *GetUserCardListResp_Entity) XXX_Size() int {
	return xxx_messageInfo_GetUserCardListResp_Entity.Size(m)
}
func (m *GetUserCardListResp_Entity) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserCardListResp_Entity.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserCardListResp_Entity proto.InternalMessageInfo

func (m *GetUserCardListResp_Entity) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GetUserCardListResp_Entity) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *GetUserCardListResp_Entity) GetBankName() string {
	if m != nil {
		return m.BankName
	}
	return ""
}

func (m *GetUserCardListResp_Entity) GetBankBranchName() string {
	if m != nil {
		return m.BankBranchName
	}
	return ""
}

func (m *GetUserCardListResp_Entity) GetCardTypeName() string {
	if m != nil {
		return m.CardTypeName
	}
	return ""
}

func (m *GetUserCardListResp_Entity) GetCardNo() string {
	if m != nil {
		return m.CardNo
	}
	return ""
}

func (m *GetUserCardListResp_Entity) GetCardHolder() string {
	if m != nil {
		return m.CardHolder
	}
	return ""
}

func init() {
	proto.RegisterType((*BankCardAuthReq)(nil), "med_wallet_go.BankCardAuthReq")
	proto.RegisterType((*BankCardAuthResp)(nil), "med_wallet_go.BankCardAuthResp")
	proto.RegisterType((*BindBankCardReq)(nil), "med_wallet_go.BindBankCardReq")
	proto.RegisterType((*BindBankCardResp)(nil), "med_wallet_go.BindBankCardResp")
	proto.RegisterType((*UnbindBankCardReq)(nil), "med_wallet_go.UnbindBankCardReq")
	proto.RegisterType((*UnbindBankCardResp)(nil), "med_wallet_go.UnbindBankCardResp")
	proto.RegisterType((*GetUserCardListReq)(nil), "med_wallet_go.GetUserCardListReq")
	proto.RegisterType((*GetUserCardListResp)(nil), "med_wallet_go.GetUserCardListResp")
	proto.RegisterType((*GetUserCardListResp_Entity)(nil), "med_wallet_go.GetUserCardListResp.Entity")
}

func init() { proto.RegisterFile("med-wallet-go/bank.proto", fileDescriptor_dfeb6510847d64f9) }

var fileDescriptor_dfeb6510847d64f9 = []byte{
	// 483 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x55, 0xd3, 0x2e, 0xdb, 0x2e, 0xa3, 0x83, 0xcb, 0x34, 0xa2, 0x20, 0x8d, 0x2e, 0x48, 0x68,
	0x3c, 0x34, 0x95, 0xc6, 0x13, 0x0f, 0x20, 0x51, 0x84, 0x00, 0x09, 0x15, 0x29, 0x50, 0x09, 0xf1,
	0x12, 0xdc, 0xd8, 0xea, 0xac, 0x26, 0x71, 0x66, 0x7b, 0xa0, 0x7e, 0x1c, 0xff, 0xc0, 0x2b, 0x7f,
	0x83, 0xec, 0xb4, 0x52, 0xed, 0x22, 0xfa, 0xb4, 0x37, 0xdf, 0x7b, 0xcf, 0xb9, 0x27, 0xf1, 0x39,
	0x32, 0x44, 0x15, 0xa3, 0xc3, 0x9f, 0xa4, 0x2c, 0x99, 0x1e, 0xce, 0xc5, 0x68, 0x46, 0xea, 0x45,
	0xda, 0x48, 0xa1, 0x05, 0xde, 0xad, 0x18, 0xcd, 0xdb, 0x49, 0x3e, 0x17, 0xc9, 0x77, 0x38, 0x1e,
	0x93, 0x7a, 0xf1, 0x86, 0x48, 0xfa, 0xfa, 0x46, 0x5f, 0x65, 0xec, 0x1a, 0x1f, 0xc2, 0x7e, 0x41,
	0x24, 0xcd, 0x6b, 0x11, 0x75, 0x06, 0x9d, 0x8b, 0xc3, 0x2c, 0x34, 0xe5, 0x44, 0x98, 0x01, 0xa7,
	0xb9, 0x29, 0xa2, 0xa0, 0x1d, 0x70, 0x6a, 0x88, 0xf8, 0x08, 0x0e, 0x25, 0x23, 0x65, 0x5e, 0x93,
	0x8a, 0x45, 0x5d, 0x3b, 0x3a, 0x30, 0x8d, 0x09, 0xa9, 0x58, 0x82, 0x70, 0xcf, 0x55, 0x50, 0x8d,
	0x55, 0xe5, 0x35, 0x5d, 0xf7, 0x6f, 0x49, 0xd5, 0x51, 0x50, 0x4d, 0xf2, 0x04, 0xee, 0x4f, 0xeb,
	0x99, 0xa7, 0xdb, 0x87, 0x80, 0x53, 0x2b, 0xd9, 0xcd, 0x02, 0x4e, 0x93, 0x13, 0x40, 0x1f, 0xa4,
	0x9a, 0xe4, 0x15, 0xe0, 0x3b, 0xa6, 0xa7, 0x8a, 0x49, 0xd3, 0xfa, 0xc8, 0x95, 0x36, 0x5c, 0x84,
	0x5e, 0x43, 0xe6, 0x6c, 0xc5, 0xb6, 0x67, 0x3c, 0x81, 0xbd, 0x92, 0x57, 0x5c, 0xdb, 0x8f, 0xed,
	0x66, 0x6d, 0x91, 0xfc, 0x0a, 0xe0, 0xc1, 0xd6, 0x02, 0xd5, 0xe0, 0x4b, 0xe8, 0x95, 0x5c, 0xe9,
	0xa8, 0x33, 0xe8, 0x5e, 0xdc, 0xb9, 0x7c, 0x96, 0x3a, 0xe6, 0xa4, 0xff, 0x60, 0xa4, 0x6f, 0x6b,
	0xcd, 0xf5, 0x32, 0xb3, 0x34, 0x23, 0xa6, 0x85, 0x26, 0xe5, 0x5a, 0xcc, 0x16, 0xf1, 0xef, 0x0e,
	0x84, 0x2d, 0xcc, 0xff, 0x3b, 0x3c, 0x85, 0xf0, 0x46, 0x31, 0xf9, 0x81, 0xae, 0x18, 0xab, 0x0a,
	0x63, 0x38, 0x30, 0x19, 0x99, 0x6c, 0x5c, 0xe5, 0xba, 0xc6, 0xa7, 0xd0, 0x37, 0xe7, 0xb1, 0x24,
	0x75, 0x71, 0x65, 0x11, 0x3d, 0x8b, 0xf0, 0xba, 0x98, 0xc0, 0x91, 0x71, 0xe9, 0xcb, 0xb2, 0x61,
	0x16, 0xb5, 0x67, 0x51, 0x4e, 0xcf, 0xe8, 0xb7, 0xb6, 0x46, 0xa1, 0x63, 0xf2, 0x19, 0x80, 0x39,
	0xbd, 0x17, 0x25, 0x65, 0x32, 0xda, 0xb7, 0xb3, 0x8d, 0xce, 0xe5, 0x9f, 0x00, 0x7a, 0xc6, 0x10,
	0xfc, 0x04, 0x47, 0x9b, 0x69, 0xc2, 0x33, 0xef, 0xca, 0xbc, 0x30, 0xc7, 0x8f, 0xff, 0x3b, 0x57,
	0x8d, 0x5d, 0xb8, 0xe1, 0xf6, 0xf6, 0x42, 0x37, 0x2f, 0xdb, 0x0b, 0xbd, 0xa8, 0xe0, 0x57, 0x38,
	0xf6, 0x7c, 0xc3, 0xf3, 0x5d, 0xbe, 0x5e, 0xc7, 0xc9, 0x6e, 0xeb, 0x71, 0x0a, 0x7d, 0x37, 0x9a,
	0x38, 0xf0, 0x58, 0x5b, 0xf1, 0x8e, 0xcf, 0x77, 0x20, 0x54, 0x33, 0xfe, 0x0c, 0xa7, 0xf6, 0x69,
	0x28, 0x44, 0x99, 0x3a, 0xcf, 0xc6, 0xb7, 0x17, 0x73, 0xae, 0x4d, 0xab, 0xe4, 0xf5, 0x82, 0xc9,
	0xb4, 0x10, 0xd5, 0x48, 0x31, 0xf9, 0x83, 0x17, 0x6c, 0xa8, 0xe8, 0x62, 0xe4, 0xbe, 0x31, 0x4e,
	0x35, 0x0b, 0xed, 0xca, 0xe7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x1b, 0x39, 0x3a, 0x79, 0x89,
	0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BankClient is the client API for Bank service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BankClient interface {
	// 银行三要素鉴权
	BankCardAuth(ctx context.Context, in *BankCardAuthReq, opts ...grpc.CallOption) (*BankCardAuthResp, error)
	// 绑定银行卡
	BindBankCard(ctx context.Context, in *BindBankCardReq, opts ...grpc.CallOption) (*BindBankCardResp, error)
	// 获取用户银行卡
	GetUserCardList(ctx context.Context, in *GetUserCardListReq, opts ...grpc.CallOption) (*GetUserCardListResp, error)
	// 解绑银行卡
	UnbindBankCard(ctx context.Context, in *UnbindBankCardReq, opts ...grpc.CallOption) (*UnbindBankCardResp, error)
}

type bankClient struct {
	cc *grpc.ClientConn
}

func NewBankClient(cc *grpc.ClientConn) BankClient {
	return &bankClient{cc}
}

func (c *bankClient) BankCardAuth(ctx context.Context, in *BankCardAuthReq, opts ...grpc.CallOption) (*BankCardAuthResp, error) {
	out := new(BankCardAuthResp)
	err := c.cc.Invoke(ctx, "/med_wallet_go.Bank/BankCardAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankClient) BindBankCard(ctx context.Context, in *BindBankCardReq, opts ...grpc.CallOption) (*BindBankCardResp, error) {
	out := new(BindBankCardResp)
	err := c.cc.Invoke(ctx, "/med_wallet_go.Bank/BindBankCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankClient) GetUserCardList(ctx context.Context, in *GetUserCardListReq, opts ...grpc.CallOption) (*GetUserCardListResp, error) {
	out := new(GetUserCardListResp)
	err := c.cc.Invoke(ctx, "/med_wallet_go.Bank/GetUserCardList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankClient) UnbindBankCard(ctx context.Context, in *UnbindBankCardReq, opts ...grpc.CallOption) (*UnbindBankCardResp, error) {
	out := new(UnbindBankCardResp)
	err := c.cc.Invoke(ctx, "/med_wallet_go.Bank/UnbindBankCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BankServer is the server API for Bank service.
type BankServer interface {
	// 银行三要素鉴权
	BankCardAuth(context.Context, *BankCardAuthReq) (*BankCardAuthResp, error)
	// 绑定银行卡
	BindBankCard(context.Context, *BindBankCardReq) (*BindBankCardResp, error)
	// 获取用户银行卡
	GetUserCardList(context.Context, *GetUserCardListReq) (*GetUserCardListResp, error)
	// 解绑银行卡
	UnbindBankCard(context.Context, *UnbindBankCardReq) (*UnbindBankCardResp, error)
}

// UnimplementedBankServer can be embedded to have forward compatible implementations.
type UnimplementedBankServer struct {
}

func (*UnimplementedBankServer) BankCardAuth(ctx context.Context, req *BankCardAuthReq) (*BankCardAuthResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BankCardAuth not implemented")
}
func (*UnimplementedBankServer) BindBankCard(ctx context.Context, req *BindBankCardReq) (*BindBankCardResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BindBankCard not implemented")
}
func (*UnimplementedBankServer) GetUserCardList(ctx context.Context, req *GetUserCardListReq) (*GetUserCardListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserCardList not implemented")
}
func (*UnimplementedBankServer) UnbindBankCard(ctx context.Context, req *UnbindBankCardReq) (*UnbindBankCardResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnbindBankCard not implemented")
}

func RegisterBankServer(s *grpc.Server, srv BankServer) {
	s.RegisterService(&_Bank_serviceDesc, srv)
}

func _Bank_BankCardAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BankCardAuthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServer).BankCardAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/med_wallet_go.Bank/BankCardAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServer).BankCardAuth(ctx, req.(*BankCardAuthReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bank_BindBankCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BindBankCardReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServer).BindBankCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/med_wallet_go.Bank/BindBankCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServer).BindBankCard(ctx, req.(*BindBankCardReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bank_GetUserCardList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserCardListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServer).GetUserCardList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/med_wallet_go.Bank/GetUserCardList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServer).GetUserCardList(ctx, req.(*GetUserCardListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bank_UnbindBankCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnbindBankCardReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServer).UnbindBankCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/med_wallet_go.Bank/UnbindBankCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServer).UnbindBankCard(ctx, req.(*UnbindBankCardReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Bank_serviceDesc = grpc.ServiceDesc{
	ServiceName: "med_wallet_go.Bank",
	HandlerType: (*BankServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BankCardAuth",
			Handler:    _Bank_BankCardAuth_Handler,
		},
		{
			MethodName: "BindBankCard",
			Handler:    _Bank_BindBankCard_Handler,
		},
		{
			MethodName: "GetUserCardList",
			Handler:    _Bank_GetUserCardList_Handler,
		},
		{
			MethodName: "UnbindBankCard",
			Handler:    _Bank_UnbindBankCard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "med-wallet-go/bank.proto",
}