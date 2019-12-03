// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

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

type CommonRequest struct {
	RequestId            string   `protobuf:"bytes,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommonRequest) Reset()         { *m = CommonRequest{} }
func (m *CommonRequest) String() string { return proto.CompactTextString(m) }
func (*CommonRequest) ProtoMessage()    {}
func (*CommonRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

func (m *CommonRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommonRequest.Unmarshal(m, b)
}
func (m *CommonRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommonRequest.Marshal(b, m, deterministic)
}
func (m *CommonRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonRequest.Merge(m, src)
}
func (m *CommonRequest) XXX_Size() int {
	return xxx_messageInfo_CommonRequest.Size(m)
}
func (m *CommonRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CommonRequest proto.InternalMessageInfo

func (m *CommonRequest) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

type CommonResponse struct {
	StatusCode           string               `protobuf:"bytes,1,opt,name=StatusCode,proto3" json:"StatusCode,omitempty"`
	Error                *CommonErrorResponse `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CommonResponse) Reset()         { *m = CommonResponse{} }
func (m *CommonResponse) String() string { return proto.CompactTextString(m) }
func (*CommonResponse) ProtoMessage()    {}
func (*CommonResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}

func (m *CommonResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommonResponse.Unmarshal(m, b)
}
func (m *CommonResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommonResponse.Marshal(b, m, deterministic)
}
func (m *CommonResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonResponse.Merge(m, src)
}
func (m *CommonResponse) XXX_Size() int {
	return xxx_messageInfo_CommonResponse.Size(m)
}
func (m *CommonResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CommonResponse proto.InternalMessageInfo

func (m *CommonResponse) GetStatusCode() string {
	if m != nil {
		return m.StatusCode
	}
	return ""
}

func (m *CommonResponse) GetError() *CommonErrorResponse {
	if m != nil {
		return m.Error
	}
	return nil
}

type CommonErrorResponse struct {
	ErrorCode            string   `protobuf:"bytes,1,opt,name=ErrorCode,proto3" json:"ErrorCode,omitempty"`
	ErrorMessage         string   `protobuf:"bytes,2,opt,name=ErrorMessage,proto3" json:"ErrorMessage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommonErrorResponse) Reset()         { *m = CommonErrorResponse{} }
func (m *CommonErrorResponse) String() string { return proto.CompactTextString(m) }
func (*CommonErrorResponse) ProtoMessage()    {}
func (*CommonErrorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{2}
}

func (m *CommonErrorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommonErrorResponse.Unmarshal(m, b)
}
func (m *CommonErrorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommonErrorResponse.Marshal(b, m, deterministic)
}
func (m *CommonErrorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonErrorResponse.Merge(m, src)
}
func (m *CommonErrorResponse) XXX_Size() int {
	return xxx_messageInfo_CommonErrorResponse.Size(m)
}
func (m *CommonErrorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonErrorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CommonErrorResponse proto.InternalMessageInfo

func (m *CommonErrorResponse) GetErrorCode() string {
	if m != nil {
		return m.ErrorCode
	}
	return ""
}

func (m *CommonErrorResponse) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func init() {
	proto.RegisterType((*CommonRequest)(nil), "protos.CommonRequest")
	proto.RegisterType((*CommonResponse)(nil), "protos.CommonResponse")
	proto.RegisterType((*CommonErrorResponse)(nil), "protos.CommonErrorResponse")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206) }

var fileDescriptor_555bd8c177793206 = []byte{
	// 172 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0x53, 0xc5, 0x4a, 0xba, 0x5c,
	0xbc, 0xce, 0x60, 0xf1, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x19, 0x2e, 0xce, 0x22,
	0x08, 0xd3, 0x33, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0x21, 0xa0, 0x94, 0xcc, 0xc5,
	0x07, 0x53, 0x5e, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0x24, 0xc7, 0xc5, 0x15, 0x5c, 0x92, 0x58,
	0x52, 0x5a, 0xec, 0x9c, 0x9f, 0x92, 0x0a, 0xd5, 0x80, 0x24, 0x22, 0x64, 0xc8, 0xc5, 0xea, 0x5a,
	0x54, 0x94, 0x5f, 0x24, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x6d, 0x24, 0x0d, 0xb1, 0xbf, 0x58, 0x0f,
	0x62, 0x0c, 0x58, 0x0a, 0x66, 0x56, 0x10, 0x44, 0xa5, 0x52, 0x38, 0x97, 0x30, 0x16, 0x59, 0x90,
	0xcb, 0xc0, 0x02, 0x48, 0x16, 0x21, 0x04, 0x84, 0x94, 0xb8, 0x78, 0xc0, 0x1c, 0xdf, 0xd4, 0xe2,
	0xe2, 0xc4, 0xf4, 0x54, 0xb0, 0x75, 0x9c, 0x41, 0x28, 0x62, 0x49, 0x10, 0x4f, 0x1b, 0x03, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x26, 0x32, 0x05, 0x30, 0x0b, 0x01, 0x00, 0x00,
}