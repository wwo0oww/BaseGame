// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/net/login.proto

package net

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

type TestEchoACK struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=Msg,proto3" json:"Msg,omitempty"`
	Value                int32    `protobuf:"varint,2,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestEchoACK) Reset()         { *m = TestEchoACK{} }
func (m *TestEchoACK) String() string { return proto.CompactTextString(m) }
func (*TestEchoACK) ProtoMessage()    {}
func (*TestEchoACK) Descriptor() ([]byte, []int) {
	return fileDescriptor_83866164adb1960c, []int{0}
}

func (m *TestEchoACK) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestEchoACK.Unmarshal(m, b)
}
func (m *TestEchoACK) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestEchoACK.Marshal(b, m, deterministic)
}
func (m *TestEchoACK) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestEchoACK.Merge(m, src)
}
func (m *TestEchoACK) XXX_Size() int {
	return xxx_messageInfo_TestEchoACK.Size(m)
}
func (m *TestEchoACK) XXX_DiscardUnknown() {
	xxx_messageInfo_TestEchoACK.DiscardUnknown(m)
}

var xxx_messageInfo_TestEchoACK proto.InternalMessageInfo

func (m *TestEchoACK) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *TestEchoACK) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type MLoginTos struct {
	Op                   int32    `protobuf:"varint,1,opt,name=op,proto3" json:"op,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Pwd                  string   `protobuf:"bytes,3,opt,name=Pwd,proto3" json:"Pwd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MLoginTos) Reset()         { *m = MLoginTos{} }
func (m *MLoginTos) String() string { return proto.CompactTextString(m) }
func (*MLoginTos) ProtoMessage()    {}
func (*MLoginTos) Descriptor() ([]byte, []int) {
	return fileDescriptor_83866164adb1960c, []int{1}
}

func (m *MLoginTos) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MLoginTos.Unmarshal(m, b)
}
func (m *MLoginTos) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MLoginTos.Marshal(b, m, deterministic)
}
func (m *MLoginTos) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MLoginTos.Merge(m, src)
}
func (m *MLoginTos) XXX_Size() int {
	return xxx_messageInfo_MLoginTos.Size(m)
}
func (m *MLoginTos) XXX_DiscardUnknown() {
	xxx_messageInfo_MLoginTos.DiscardUnknown(m)
}

var xxx_messageInfo_MLoginTos proto.InternalMessageInfo

func (m *MLoginTos) GetOp() int32 {
	if m != nil {
		return m.Op
	}
	return 0
}

func (m *MLoginTos) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MLoginTos) GetPwd() string {
	if m != nil {
		return m.Pwd
	}
	return ""
}

type MLoginToc struct {
	Op                   int32    `protobuf:"varint,1,opt,name=op,proto3" json:"op,omitempty"`
	Errcode              int32    `protobuf:"varint,2,opt,name=errcode,proto3" json:"errcode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MLoginToc) Reset()         { *m = MLoginToc{} }
func (m *MLoginToc) String() string { return proto.CompactTextString(m) }
func (*MLoginToc) ProtoMessage()    {}
func (*MLoginToc) Descriptor() ([]byte, []int) {
	return fileDescriptor_83866164adb1960c, []int{2}
}

func (m *MLoginToc) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MLoginToc.Unmarshal(m, b)
}
func (m *MLoginToc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MLoginToc.Marshal(b, m, deterministic)
}
func (m *MLoginToc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MLoginToc.Merge(m, src)
}
func (m *MLoginToc) XXX_Size() int {
	return xxx_messageInfo_MLoginToc.Size(m)
}
func (m *MLoginToc) XXX_DiscardUnknown() {
	xxx_messageInfo_MLoginToc.DiscardUnknown(m)
}

var xxx_messageInfo_MLoginToc proto.InternalMessageInfo

func (m *MLoginToc) GetOp() int32 {
	if m != nil {
		return m.Op
	}
	return 0
}

func (m *MLoginToc) GetErrcode() int32 {
	if m != nil {
		return m.Errcode
	}
	return 0
}

func init() {
	proto.RegisterType((*TestEchoACK)(nil), "net.TestEchoACK")
	proto.RegisterType((*MLoginTos)(nil), "net.m_login_tos")
	proto.RegisterType((*MLoginToc)(nil), "net.m_login_toc")
}

func init() {
	proto.RegisterFile("proto/net/login.proto", fileDescriptor_83866164adb1960c)
}

var fileDescriptor_83866164adb1960c = []byte{
	// 173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0x4b, 0x2d, 0xd1, 0xcf, 0xc9, 0x4f, 0xcf, 0xcc, 0xd3, 0x03, 0xf3, 0x85, 0x98,
	0xf3, 0x52, 0x4b, 0x94, 0x4c, 0xb9, 0xb8, 0x43, 0x52, 0x8b, 0x4b, 0x5c, 0x93, 0x33, 0xf2, 0x1d,
	0x9d, 0xbd, 0x85, 0x04, 0xb8, 0x98, 0x7d, 0x8b, 0xd3, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83,
	0x40, 0x4c, 0x21, 0x11, 0x2e, 0xd6, 0xb0, 0xc4, 0x9c, 0xd2, 0x54, 0x09, 0x26, 0x05, 0x46, 0x0d,
	0xd6, 0x20, 0x08, 0x47, 0xc9, 0x99, 0x8b, 0x3b, 0x37, 0x1e, 0x6c, 0x58, 0x7c, 0x49, 0x7e, 0xb1,
	0x10, 0x1f, 0x17, 0x53, 0x7e, 0x01, 0x58, 0x17, 0x6b, 0x10, 0x53, 0x7e, 0x81, 0x90, 0x10, 0x17,
	0x8b, 0x5f, 0x62, 0x2e, 0x44, 0x0f, 0x67, 0x10, 0x98, 0x0d, 0x32, 0x3a, 0xa0, 0x3c, 0x45, 0x82,
	0x19, 0x62, 0x74, 0x40, 0x79, 0x8a, 0x92, 0x39, 0xb2, 0x21, 0xc9, 0x18, 0x86, 0x48, 0x70, 0xb1,
	0xa7, 0x16, 0x15, 0x25, 0xe7, 0xa7, 0xc0, 0xec, 0x86, 0x71, 0x93, 0xd8, 0xc0, 0x1e, 0x30, 0x06,
	0x04, 0x00, 0x00, 0xff, 0xff, 0xa4, 0xb7, 0x4e, 0x54, 0xd9, 0x00, 0x00, 0x00,
}
