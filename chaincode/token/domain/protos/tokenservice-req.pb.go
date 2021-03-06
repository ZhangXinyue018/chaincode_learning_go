// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tokenservice-req.proto

package protos

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type PingRequest struct {
	CommonRequest        *CommonRequest `protobuf:"bytes,1,opt,name=CommonRequest,proto3" json:"CommonRequest,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e6b36ae05b9afdc, []int{0}
}

func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRequest.Unmarshal(m, b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
}
func (m *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(m, src)
}
func (m *PingRequest) XXX_Size() int {
	return xxx_messageInfo_PingRequest.Size(m)
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

func (m *PingRequest) GetCommonRequest() *CommonRequest {
	if m != nil {
		return m.CommonRequest
	}
	return nil
}

type CreateTokenRequest struct {
	CommonRequest        *CommonRequest `protobuf:"bytes,1,opt,name=CommonRequest,proto3" json:"CommonRequest,omitempty"`
	TokenName            string         `protobuf:"bytes,2,opt,name=TokenName,proto3" json:"TokenName,omitempty"`
	MaxAmount            int64          `protobuf:"varint,3,opt,name=MaxAmount,proto3" json:"MaxAmount,omitempty"`
	Creator              string         `protobuf:"bytes,4,opt,name=Creator,proto3" json:"Creator,omitempty"`
	Issuer               string         `protobuf:"bytes,5,opt,name=Issuer,proto3" json:"Issuer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CreateTokenRequest) Reset()         { *m = CreateTokenRequest{} }
func (m *CreateTokenRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTokenRequest) ProtoMessage()    {}
func (*CreateTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e6b36ae05b9afdc, []int{1}
}

func (m *CreateTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTokenRequest.Unmarshal(m, b)
}
func (m *CreateTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTokenRequest.Marshal(b, m, deterministic)
}
func (m *CreateTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTokenRequest.Merge(m, src)
}
func (m *CreateTokenRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTokenRequest.Size(m)
}
func (m *CreateTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTokenRequest proto.InternalMessageInfo

func (m *CreateTokenRequest) GetCommonRequest() *CommonRequest {
	if m != nil {
		return m.CommonRequest
	}
	return nil
}

func (m *CreateTokenRequest) GetTokenName() string {
	if m != nil {
		return m.TokenName
	}
	return ""
}

func (m *CreateTokenRequest) GetMaxAmount() int64 {
	if m != nil {
		return m.MaxAmount
	}
	return 0
}

func (m *CreateTokenRequest) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *CreateTokenRequest) GetIssuer() string {
	if m != nil {
		return m.Issuer
	}
	return ""
}

type IssueTokenRequest struct {
	CommonRequest        *CommonRequest `protobuf:"bytes,1,opt,name=CommonRequest,proto3" json:"CommonRequest,omitempty"`
	UserName             string         `protobuf:"bytes,2,opt,name=UserName,proto3" json:"UserName,omitempty"`
	TokenName            string         `protobuf:"bytes,3,opt,name=TokenName,proto3" json:"TokenName,omitempty"`
	TokenAmount          int64          `protobuf:"varint,4,opt,name=TokenAmount,proto3" json:"TokenAmount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *IssueTokenRequest) Reset()         { *m = IssueTokenRequest{} }
func (m *IssueTokenRequest) String() string { return proto.CompactTextString(m) }
func (*IssueTokenRequest) ProtoMessage()    {}
func (*IssueTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e6b36ae05b9afdc, []int{2}
}

func (m *IssueTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IssueTokenRequest.Unmarshal(m, b)
}
func (m *IssueTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IssueTokenRequest.Marshal(b, m, deterministic)
}
func (m *IssueTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IssueTokenRequest.Merge(m, src)
}
func (m *IssueTokenRequest) XXX_Size() int {
	return xxx_messageInfo_IssueTokenRequest.Size(m)
}
func (m *IssueTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IssueTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IssueTokenRequest proto.InternalMessageInfo

func (m *IssueTokenRequest) GetCommonRequest() *CommonRequest {
	if m != nil {
		return m.CommonRequest
	}
	return nil
}

func (m *IssueTokenRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *IssueTokenRequest) GetTokenName() string {
	if m != nil {
		return m.TokenName
	}
	return ""
}

func (m *IssueTokenRequest) GetTokenAmount() int64 {
	if m != nil {
		return m.TokenAmount
	}
	return 0
}

type TransferTokenRequest struct {
	CommonRequest        *CommonRequest `protobuf:"bytes,1,opt,name=CommonRequest,proto3" json:"CommonRequest,omitempty"`
	FromUserName         string         `protobuf:"bytes,2,opt,name=FromUserName,proto3" json:"FromUserName,omitempty"`
	ToUserName           string         `protobuf:"bytes,3,opt,name=ToUserName,proto3" json:"ToUserName,omitempty"`
	TokenName            string         `protobuf:"bytes,4,opt,name=TokenName,proto3" json:"TokenName,omitempty"`
	TokenAmount          int64          `protobuf:"varint,5,opt,name=TokenAmount,proto3" json:"TokenAmount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TransferTokenRequest) Reset()         { *m = TransferTokenRequest{} }
func (m *TransferTokenRequest) String() string { return proto.CompactTextString(m) }
func (*TransferTokenRequest) ProtoMessage()    {}
func (*TransferTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e6b36ae05b9afdc, []int{3}
}

func (m *TransferTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferTokenRequest.Unmarshal(m, b)
}
func (m *TransferTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferTokenRequest.Marshal(b, m, deterministic)
}
func (m *TransferTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferTokenRequest.Merge(m, src)
}
func (m *TransferTokenRequest) XXX_Size() int {
	return xxx_messageInfo_TransferTokenRequest.Size(m)
}
func (m *TransferTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TransferTokenRequest proto.InternalMessageInfo

func (m *TransferTokenRequest) GetCommonRequest() *CommonRequest {
	if m != nil {
		return m.CommonRequest
	}
	return nil
}

func (m *TransferTokenRequest) GetFromUserName() string {
	if m != nil {
		return m.FromUserName
	}
	return ""
}

func (m *TransferTokenRequest) GetToUserName() string {
	if m != nil {
		return m.ToUserName
	}
	return ""
}

func (m *TransferTokenRequest) GetTokenName() string {
	if m != nil {
		return m.TokenName
	}
	return ""
}

func (m *TransferTokenRequest) GetTokenAmount() int64 {
	if m != nil {
		return m.TokenAmount
	}
	return 0
}

type GetTokenRequest struct {
	CommonRequest        *CommonRequest `protobuf:"bytes,1,opt,name=CommonRequest,proto3" json:"CommonRequest,omitempty"`
	UserName             string         `protobuf:"bytes,2,opt,name=UserName,proto3" json:"UserName,omitempty"`
	TokenName            string         `protobuf:"bytes,3,opt,name=TokenName,proto3" json:"TokenName,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetTokenRequest) Reset()         { *m = GetTokenRequest{} }
func (m *GetTokenRequest) String() string { return proto.CompactTextString(m) }
func (*GetTokenRequest) ProtoMessage()    {}
func (*GetTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e6b36ae05b9afdc, []int{4}
}

func (m *GetTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTokenRequest.Unmarshal(m, b)
}
func (m *GetTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTokenRequest.Marshal(b, m, deterministic)
}
func (m *GetTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTokenRequest.Merge(m, src)
}
func (m *GetTokenRequest) XXX_Size() int {
	return xxx_messageInfo_GetTokenRequest.Size(m)
}
func (m *GetTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTokenRequest proto.InternalMessageInfo

func (m *GetTokenRequest) GetCommonRequest() *CommonRequest {
	if m != nil {
		return m.CommonRequest
	}
	return nil
}

func (m *GetTokenRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *GetTokenRequest) GetTokenName() string {
	if m != nil {
		return m.TokenName
	}
	return ""
}

type PaginateTokenByUserNameRequest struct {
	CommonRequest        *CommonRequest `protobuf:"bytes,1,opt,name=CommonRequest,proto3" json:"CommonRequest,omitempty"`
	UserName             string         `protobuf:"bytes,2,opt,name=UserName,proto3" json:"UserName,omitempty"`
	PageSize             int32          `protobuf:"varint,3,opt,name=PageSize,proto3" json:"PageSize,omitempty"`
	BookMark             string         `protobuf:"bytes,4,opt,name=BookMark,proto3" json:"BookMark,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *PaginateTokenByUserNameRequest) Reset()         { *m = PaginateTokenByUserNameRequest{} }
func (m *PaginateTokenByUserNameRequest) String() string { return proto.CompactTextString(m) }
func (*PaginateTokenByUserNameRequest) ProtoMessage()    {}
func (*PaginateTokenByUserNameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e6b36ae05b9afdc, []int{5}
}

func (m *PaginateTokenByUserNameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaginateTokenByUserNameRequest.Unmarshal(m, b)
}
func (m *PaginateTokenByUserNameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaginateTokenByUserNameRequest.Marshal(b, m, deterministic)
}
func (m *PaginateTokenByUserNameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaginateTokenByUserNameRequest.Merge(m, src)
}
func (m *PaginateTokenByUserNameRequest) XXX_Size() int {
	return xxx_messageInfo_PaginateTokenByUserNameRequest.Size(m)
}
func (m *PaginateTokenByUserNameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PaginateTokenByUserNameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PaginateTokenByUserNameRequest proto.InternalMessageInfo

func (m *PaginateTokenByUserNameRequest) GetCommonRequest() *CommonRequest {
	if m != nil {
		return m.CommonRequest
	}
	return nil
}

func (m *PaginateTokenByUserNameRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *PaginateTokenByUserNameRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *PaginateTokenByUserNameRequest) GetBookMark() string {
	if m != nil {
		return m.BookMark
	}
	return ""
}

type PaginateTokenByTokenNameRequest struct {
	CommonRequest        *CommonRequest `protobuf:"bytes,1,opt,name=CommonRequest,proto3" json:"CommonRequest,omitempty"`
	TokenName            string         `protobuf:"bytes,2,opt,name=TokenName,proto3" json:"TokenName,omitempty"`
	PageSize             int32          `protobuf:"varint,3,opt,name=PageSize,proto3" json:"PageSize,omitempty"`
	BookMark             string         `protobuf:"bytes,4,opt,name=BookMark,proto3" json:"BookMark,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *PaginateTokenByTokenNameRequest) Reset()         { *m = PaginateTokenByTokenNameRequest{} }
func (m *PaginateTokenByTokenNameRequest) String() string { return proto.CompactTextString(m) }
func (*PaginateTokenByTokenNameRequest) ProtoMessage()    {}
func (*PaginateTokenByTokenNameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e6b36ae05b9afdc, []int{6}
}

func (m *PaginateTokenByTokenNameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaginateTokenByTokenNameRequest.Unmarshal(m, b)
}
func (m *PaginateTokenByTokenNameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaginateTokenByTokenNameRequest.Marshal(b, m, deterministic)
}
func (m *PaginateTokenByTokenNameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaginateTokenByTokenNameRequest.Merge(m, src)
}
func (m *PaginateTokenByTokenNameRequest) XXX_Size() int {
	return xxx_messageInfo_PaginateTokenByTokenNameRequest.Size(m)
}
func (m *PaginateTokenByTokenNameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PaginateTokenByTokenNameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PaginateTokenByTokenNameRequest proto.InternalMessageInfo

func (m *PaginateTokenByTokenNameRequest) GetCommonRequest() *CommonRequest {
	if m != nil {
		return m.CommonRequest
	}
	return nil
}

func (m *PaginateTokenByTokenNameRequest) GetTokenName() string {
	if m != nil {
		return m.TokenName
	}
	return ""
}

func (m *PaginateTokenByTokenNameRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *PaginateTokenByTokenNameRequest) GetBookMark() string {
	if m != nil {
		return m.BookMark
	}
	return ""
}

type PaginateTokenLogByFromUserNameRequest struct {
	CommonRequest        *CommonRequest `protobuf:"bytes,1,opt,name=CommonRequest,proto3" json:"CommonRequest,omitempty"`
	FromUserName         string         `protobuf:"bytes,2,opt,name=FromUserName,proto3" json:"FromUserName,omitempty"`
	PageSize             int32          `protobuf:"varint,3,opt,name=PageSize,proto3" json:"PageSize,omitempty"`
	BookMark             string         `protobuf:"bytes,4,opt,name=BookMark,proto3" json:"BookMark,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *PaginateTokenLogByFromUserNameRequest) Reset()         { *m = PaginateTokenLogByFromUserNameRequest{} }
func (m *PaginateTokenLogByFromUserNameRequest) String() string { return proto.CompactTextString(m) }
func (*PaginateTokenLogByFromUserNameRequest) ProtoMessage()    {}
func (*PaginateTokenLogByFromUserNameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e6b36ae05b9afdc, []int{7}
}

func (m *PaginateTokenLogByFromUserNameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaginateTokenLogByFromUserNameRequest.Unmarshal(m, b)
}
func (m *PaginateTokenLogByFromUserNameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaginateTokenLogByFromUserNameRequest.Marshal(b, m, deterministic)
}
func (m *PaginateTokenLogByFromUserNameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaginateTokenLogByFromUserNameRequest.Merge(m, src)
}
func (m *PaginateTokenLogByFromUserNameRequest) XXX_Size() int {
	return xxx_messageInfo_PaginateTokenLogByFromUserNameRequest.Size(m)
}
func (m *PaginateTokenLogByFromUserNameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PaginateTokenLogByFromUserNameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PaginateTokenLogByFromUserNameRequest proto.InternalMessageInfo

func (m *PaginateTokenLogByFromUserNameRequest) GetCommonRequest() *CommonRequest {
	if m != nil {
		return m.CommonRequest
	}
	return nil
}

func (m *PaginateTokenLogByFromUserNameRequest) GetFromUserName() string {
	if m != nil {
		return m.FromUserName
	}
	return ""
}

func (m *PaginateTokenLogByFromUserNameRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *PaginateTokenLogByFromUserNameRequest) GetBookMark() string {
	if m != nil {
		return m.BookMark
	}
	return ""
}

type PaginateTokenLogByToUserNameRequest struct {
	CommonRequest        *CommonRequest `protobuf:"bytes,1,opt,name=CommonRequest,proto3" json:"CommonRequest,omitempty"`
	ToUserName           string         `protobuf:"bytes,2,opt,name=ToUserName,proto3" json:"ToUserName,omitempty"`
	PageSize             int32          `protobuf:"varint,3,opt,name=PageSize,proto3" json:"PageSize,omitempty"`
	BookMark             string         `protobuf:"bytes,4,opt,name=BookMark,proto3" json:"BookMark,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *PaginateTokenLogByToUserNameRequest) Reset()         { *m = PaginateTokenLogByToUserNameRequest{} }
func (m *PaginateTokenLogByToUserNameRequest) String() string { return proto.CompactTextString(m) }
func (*PaginateTokenLogByToUserNameRequest) ProtoMessage()    {}
func (*PaginateTokenLogByToUserNameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e6b36ae05b9afdc, []int{8}
}

func (m *PaginateTokenLogByToUserNameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaginateTokenLogByToUserNameRequest.Unmarshal(m, b)
}
func (m *PaginateTokenLogByToUserNameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaginateTokenLogByToUserNameRequest.Marshal(b, m, deterministic)
}
func (m *PaginateTokenLogByToUserNameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaginateTokenLogByToUserNameRequest.Merge(m, src)
}
func (m *PaginateTokenLogByToUserNameRequest) XXX_Size() int {
	return xxx_messageInfo_PaginateTokenLogByToUserNameRequest.Size(m)
}
func (m *PaginateTokenLogByToUserNameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PaginateTokenLogByToUserNameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PaginateTokenLogByToUserNameRequest proto.InternalMessageInfo

func (m *PaginateTokenLogByToUserNameRequest) GetCommonRequest() *CommonRequest {
	if m != nil {
		return m.CommonRequest
	}
	return nil
}

func (m *PaginateTokenLogByToUserNameRequest) GetToUserName() string {
	if m != nil {
		return m.ToUserName
	}
	return ""
}

func (m *PaginateTokenLogByToUserNameRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *PaginateTokenLogByToUserNameRequest) GetBookMark() string {
	if m != nil {
		return m.BookMark
	}
	return ""
}

func init() {
	proto.RegisterType((*PingRequest)(nil), "protos.PingRequest")
	proto.RegisterType((*CreateTokenRequest)(nil), "protos.CreateTokenRequest")
	proto.RegisterType((*IssueTokenRequest)(nil), "protos.IssueTokenRequest")
	proto.RegisterType((*TransferTokenRequest)(nil), "protos.TransferTokenRequest")
	proto.RegisterType((*GetTokenRequest)(nil), "protos.GetTokenRequest")
	proto.RegisterType((*PaginateTokenByUserNameRequest)(nil), "protos.PaginateTokenByUserNameRequest")
	proto.RegisterType((*PaginateTokenByTokenNameRequest)(nil), "protos.PaginateTokenByTokenNameRequest")
	proto.RegisterType((*PaginateTokenLogByFromUserNameRequest)(nil), "protos.PaginateTokenLogByFromUserNameRequest")
	proto.RegisterType((*PaginateTokenLogByToUserNameRequest)(nil), "protos.PaginateTokenLogByToUserNameRequest")
}

func init() { proto.RegisterFile("tokenservice-req.proto", fileDescriptor_2e6b36ae05b9afdc) }

var fileDescriptor_2e6b36ae05b9afdc = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x95, 0xcf, 0x4a, 0xeb, 0x40,
	0x14, 0x87, 0x99, 0xfe, 0xbb, 0xed, 0x69, 0x2f, 0x97, 0x3b, 0xdc, 0x5b, 0x42, 0x91, 0x1a, 0x46,
	0x84, 0x6e, 0xec, 0x42, 0x97, 0xae, 0x6c, 0x41, 0x51, 0xac, 0x94, 0x58, 0x1f, 0x60, 0x2c, 0xc7,
	0x10, 0x4a, 0x32, 0x76, 0x26, 0x15, 0xeb, 0x13, 0xf8, 0x30, 0x6e, 0x04, 0x57, 0x6e, 0xf4, 0x2d,
	0x7c, 0x1d, 0xe9, 0x24, 0x4d, 0x33, 0xed, 0x42, 0x28, 0xb1, 0xb8, 0x6a, 0xcf, 0xef, 0xcc, 0x0c,
	0xdf, 0x97, 0xcc, 0x21, 0x50, 0x0f, 0xc5, 0x08, 0x03, 0x85, 0xf2, 0xce, 0x1b, 0xe2, 0x9e, 0xc4,
	0x71, 0xfb, 0x56, 0x8a, 0x50, 0xd0, 0x92, 0xfe, 0x51, 0x8d, 0xda, 0x50, 0xf8, 0xbe, 0x08, 0xa2,
	0x94, 0x9d, 0x41, 0xb5, 0xef, 0x05, 0xae, 0x83, 0xe3, 0x09, 0xaa, 0x90, 0x1e, 0xc2, 0xef, 0xae,
	0x6e, 0xc7, 0x81, 0x45, 0x6c, 0xd2, 0xaa, 0xee, 0xff, 0x8f, 0x56, 0xab, 0xb6, 0xd1, 0x74, 0xcc,
	0xb5, 0xec, 0x8d, 0x00, 0xed, 0x4a, 0xe4, 0x21, 0x0e, 0x66, 0x08, 0x59, 0x9c, 0x49, 0xb7, 0xa0,
	0xa2, 0x0f, 0xbb, 0xe0, 0x3e, 0x5a, 0x39, 0x9b, 0xb4, 0x2a, 0xce, 0x22, 0x98, 0x75, 0x7b, 0xfc,
	0xfe, 0xc8, 0x17, 0x93, 0x20, 0xb4, 0xf2, 0x36, 0x69, 0xe5, 0x9d, 0x45, 0x40, 0x2d, 0xf8, 0xa5,
	0x71, 0x84, 0xb4, 0x0a, 0x7a, 0xe7, 0xbc, 0xa4, 0x75, 0x28, 0x9d, 0x2a, 0x35, 0x41, 0x69, 0x15,
	0x75, 0x23, 0xae, 0xd8, 0x13, 0x81, 0xbf, 0xfa, 0x6f, 0x76, 0x02, 0x0d, 0x28, 0x5f, 0x29, 0x94,
	0x29, 0xfe, 0xa4, 0x36, 0xe5, 0xf2, 0xcb, 0x72, 0x36, 0x54, 0x75, 0x11, 0xeb, 0x15, 0xb4, 0x5e,
	0x3a, 0x62, 0x1f, 0x04, 0xfe, 0x0d, 0x24, 0x0f, 0xd4, 0x0d, 0xca, 0xec, 0x88, 0x19, 0xd4, 0x8e,
	0xa5, 0xf0, 0x97, 0xa8, 0x8d, 0x8c, 0x36, 0x01, 0x06, 0x22, 0x59, 0x11, 0xa1, 0xa7, 0x12, 0xd3,
	0xac, 0xf0, 0x85, 0x59, 0x71, 0xd5, 0xec, 0x91, 0xc0, 0x9f, 0x13, 0x0c, 0x7f, 0xc0, 0x6b, 0x60,
	0xcf, 0x04, 0x9a, 0x7d, 0xee, 0x7a, 0xc1, 0xfc, 0x5e, 0x77, 0xa6, 0xf3, 0x9d, 0xdf, 0x4e, 0xd6,
	0x80, 0x72, 0x9f, 0xbb, 0x78, 0xe9, 0x3d, 0x44, 0x60, 0x45, 0x27, 0xa9, 0x67, 0xbd, 0x8e, 0x10,
	0xa3, 0x1e, 0x97, 0xa3, 0xf8, 0x09, 0x27, 0x35, 0x7b, 0x21, 0xb0, 0xbd, 0xc4, 0x9c, 0x08, 0x6d,
	0x60, 0x2c, 0xd7, 0xc5, 0x7e, 0x27, 0xb0, 0x6b, 0x60, 0x9f, 0x0b, 0xb7, 0x33, 0x4d, 0x5f, 0xbc,
	0x8d, 0x5d, 0xf0, 0x75, 0x15, 0x5e, 0x09, 0xec, 0xac, 0x2a, 0x2c, 0x26, 0x23, 0x13, 0x01, 0x73,
	0xfa, 0x72, 0x2b, 0xd3, 0xb7, 0x26, 0xfc, 0x75, 0xf4, 0x89, 0x38, 0xf8, 0x0c, 0x00, 0x00, 0xff,
	0xff, 0x02, 0xa0, 0x63, 0x9f, 0x43, 0x06, 0x00, 0x00,
}
