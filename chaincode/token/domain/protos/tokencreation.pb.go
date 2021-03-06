// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tokencreation.proto

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

type TokenCreationPB struct {
	TokenName            string   `protobuf:"bytes,1,opt,name=TokenName,proto3" json:"TokenName,omitempty"`
	MaxAmount            int64    `protobuf:"varint,2,opt,name=MaxAmount,proto3" json:"MaxAmount,omitempty"`
	CurrentAmount        int64    `protobuf:"varint,3,opt,name=CurrentAmount,proto3" json:"CurrentAmount,omitempty"`
	Creator              string   `protobuf:"bytes,4,opt,name=Creator,proto3" json:"Creator,omitempty"`
	Issuer               string   `protobuf:"bytes,5,opt,name=Issuer,proto3" json:"Issuer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TokenCreationPB) Reset()         { *m = TokenCreationPB{} }
func (m *TokenCreationPB) String() string { return proto.CompactTextString(m) }
func (*TokenCreationPB) ProtoMessage()    {}
func (*TokenCreationPB) Descriptor() ([]byte, []int) {
	return fileDescriptor_d61ccf08e6038651, []int{0}
}

func (m *TokenCreationPB) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TokenCreationPB.Unmarshal(m, b)
}
func (m *TokenCreationPB) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TokenCreationPB.Marshal(b, m, deterministic)
}
func (m *TokenCreationPB) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenCreationPB.Merge(m, src)
}
func (m *TokenCreationPB) XXX_Size() int {
	return xxx_messageInfo_TokenCreationPB.Size(m)
}
func (m *TokenCreationPB) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenCreationPB.DiscardUnknown(m)
}

var xxx_messageInfo_TokenCreationPB proto.InternalMessageInfo

func (m *TokenCreationPB) GetTokenName() string {
	if m != nil {
		return m.TokenName
	}
	return ""
}

func (m *TokenCreationPB) GetMaxAmount() int64 {
	if m != nil {
		return m.MaxAmount
	}
	return 0
}

func (m *TokenCreationPB) GetCurrentAmount() int64 {
	if m != nil {
		return m.CurrentAmount
	}
	return 0
}

func (m *TokenCreationPB) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *TokenCreationPB) GetIssuer() string {
	if m != nil {
		return m.Issuer
	}
	return ""
}

func init() {
	proto.RegisterType((*TokenCreationPB)(nil), "protos.TokenCreationPB")
}

func init() { proto.RegisterFile("tokencreation.proto", fileDescriptor_d61ccf08e6038651) }

var fileDescriptor_d61ccf08e6038651 = []byte{
	// 153 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0xc9, 0xcf, 0x4e,
	0xcd, 0x4b, 0x2e, 0x4a, 0x4d, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x03, 0x53, 0xc5, 0x4a, 0x4b, 0x19, 0xb9, 0xf8, 0x43, 0x40, 0xf2, 0xce, 0x50, 0xf9, 0x00,
	0x27, 0x21, 0x19, 0x2e, 0x4e, 0xb0, 0x90, 0x5f, 0x62, 0x6e, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06,
	0x67, 0x10, 0x42, 0x00, 0x24, 0xeb, 0x9b, 0x58, 0xe1, 0x98, 0x9b, 0x5f, 0x9a, 0x57, 0x22, 0xc1,
	0xa4, 0xc0, 0xa8, 0xc1, 0x1c, 0x84, 0x10, 0x10, 0x52, 0xe1, 0xe2, 0x75, 0x2e, 0x2d, 0x2a, 0x4a,
	0xcd, 0x2b, 0x81, 0xaa, 0x60, 0x06, 0xab, 0x40, 0x15, 0x14, 0x92, 0xe0, 0x62, 0x07, 0xdb, 0x97,
	0x5f, 0x24, 0xc1, 0x02, 0x36, 0x1f, 0xc6, 0x15, 0x12, 0xe3, 0x62, 0xf3, 0x2c, 0x2e, 0x2e, 0x4d,
	0x2d, 0x92, 0x60, 0x05, 0x4b, 0x40, 0x79, 0x49, 0x10, 0xf7, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xff, 0x8a, 0x5a, 0x45, 0xcd, 0x00, 0x00, 0x00,
}
