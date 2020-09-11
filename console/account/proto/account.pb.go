// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account.proto

package proto

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

// 登录请求体
type LoginRequest struct {
	//@inject_tag: json:"username"
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username"`
	//@inject_tag: json:"password"
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{0}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

// 登录响应体
type LoginResponse struct {
	//@inject_tag: json:"token"
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token"`
	//@inject_tag: json:"expires"
	Expires              string   `protobuf:"bytes,2,opt,name=expires,proto3" json:"expires"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{1}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *LoginResponse) GetExpires() string {
	if m != nil {
		return m.Expires
	}
	return ""
}

// 注册请求体
type RegisterRequest struct {
	//@inject_tag: validate:"required" json:"nickname"
	Nickname string `protobuf:"bytes,1,opt,name=nickname,proto3" json:"nickname" validate:"required"`
	//@inject_tag: validate:"required,gte=11" message:"required:手机号不能为空,gt:手机号不正确" json:"phone"
	Phone string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone" validate:"required,gte=11" message:"required:手机号不能为空,gt:手机号不正确"`
	//@inject_tag: validate:"required" json:"password"
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password" validate:"required"`
	//@inject_tag: validate:"required" json:"code"
	Code                 string   `protobuf:"bytes,4,opt,name=code,proto3" json:"code" validate:"required"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{2}
}

func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *RegisterRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *RegisterRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *RegisterRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

// 注册响应
type RegisterResponse struct {
	//@inject_tag: json:"id"
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResponse) Reset()         { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()    {}
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{3}
}

func (m *RegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResponse.Unmarshal(m, b)
}
func (m *RegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResponse.Marshal(b, m, deterministic)
}
func (m *RegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResponse.Merge(m, src)
}
func (m *RegisterResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterResponse.Size(m)
}
func (m *RegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResponse proto.InternalMessageInfo

func (m *RegisterResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

// 账号信息请求体
type InfoRequest struct {
	//@inject_tag: validate:"required" json:"authorization"
	Authorization        string   `protobuf:"bytes,1,opt,name=Authorization,proto3" json:"authorization" validate:"required"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InfoRequest) Reset()         { *m = InfoRequest{} }
func (m *InfoRequest) String() string { return proto.CompactTextString(m) }
func (*InfoRequest) ProtoMessage()    {}
func (*InfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{4}
}

func (m *InfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InfoRequest.Unmarshal(m, b)
}
func (m *InfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InfoRequest.Marshal(b, m, deterministic)
}
func (m *InfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InfoRequest.Merge(m, src)
}
func (m *InfoRequest) XXX_Size() int {
	return xxx_messageInfo_InfoRequest.Size(m)
}
func (m *InfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InfoRequest proto.InternalMessageInfo

func (m *InfoRequest) GetAuthorization() string {
	if m != nil {
		return m.Authorization
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "LoginResponse")
	proto.RegisterType((*RegisterRequest)(nil), "RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "RegisterResponse")
	proto.RegisterType((*InfoRequest)(nil), "InfoRequest")
}

func init() {
	proto.RegisterFile("account.proto", fileDescriptor_8e28828dcb8d24f0)
}

var fileDescriptor_8e28828dcb8d24f0 = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0x31, 0x4f, 0xf3, 0x30,
	0x10, 0x86, 0xd5, 0x7c, 0xcd, 0x97, 0xf6, 0x20, 0xa5, 0x58, 0x0c, 0x51, 0x26, 0x64, 0x21, 0xc4,
	0x64, 0x24, 0x3a, 0x32, 0xa0, 0x32, 0x20, 0x21, 0x31, 0x65, 0x64, 0x4b, 0x93, 0xa3, 0xb5, 0x2a,
	0x7c, 0xc6, 0x76, 0x04, 0xe2, 0xd7, 0xa3, 0x3a, 0x36, 0x24, 0x9d, 0x92, 0xe7, 0x4e, 0x7a, 0xef,
	0xc9, 0x1b, 0xc8, 0xeb, 0xa6, 0xa1, 0x4e, 0x39, 0xa1, 0x0d, 0x39, 0xe2, 0x4f, 0x70, 0xfa, 0x42,
	0x5b, 0xa9, 0x2a, 0xfc, 0xe8, 0xd0, 0x3a, 0x56, 0xc2, 0xac, 0xb3, 0x68, 0x54, 0xfd, 0x8e, 0xc5,
	0xe4, 0x72, 0x72, 0x33, 0xaf, 0x7e, 0xf9, 0xb0, 0xd3, 0xb5, 0xb5, 0x9f, 0x64, 0xda, 0x22, 0xe9,
	0x77, 0x91, 0xf9, 0x03, 0xe4, 0x21, 0xc7, 0x6a, 0x52, 0x16, 0xd9, 0x05, 0xa4, 0x8e, 0xf6, 0xa8,
	0x42, 0x4a, 0x0f, 0xac, 0x80, 0x0c, 0xbf, 0xb4, 0x34, 0x68, 0x43, 0x42, 0x44, 0x6e, 0xe1, 0xac,
	0xc2, 0xad, 0xb4, 0x0e, 0xcd, 0xc0, 0x45, 0xc9, 0x66, 0x3f, 0x74, 0x89, 0x7c, 0x88, 0xd7, 0x3b,
	0x52, 0x18, 0x62, 0x7a, 0x18, 0x19, 0xfe, 0x1b, 0x1b, 0x32, 0x06, 0xd3, 0x86, 0x5a, 0x2c, 0xa6,
	0x7e, 0xee, 0xdf, 0x39, 0x87, 0xe5, 0xdf, 0xd1, 0x20, 0xbe, 0x80, 0x44, 0xb6, 0xfe, 0x5e, 0x5a,
	0x25, 0xb2, 0xe5, 0x2b, 0x38, 0x79, 0x56, 0x6f, 0x14, 0xa5, 0xae, 0x20, 0x5f, 0x77, 0x6e, 0x47,
	0x46, 0x7e, 0xd7, 0x4e, 0x52, 0xfc, 0xbe, 0xf1, 0xf0, 0x6e, 0x03, 0xd9, 0xba, 0xef, 0x99, 0x5d,
	0x43, 0xea, 0x9b, 0x61, 0xb9, 0x18, 0x36, 0x5d, 0x2e, 0xc4, 0xb8, 0xb0, 0x5b, 0x98, 0x45, 0x17,
	0xb6, 0x14, 0x47, 0x5d, 0x94, 0xe7, 0xe2, 0x58, 0xf4, 0x71, 0xfe, 0x9a, 0x89, 0x7b, 0xff, 0x17,
	0x37, 0xff, 0xfd, 0x63, 0xf5, 0x13, 0x00, 0x00, 0xff, 0xff, 0xac, 0xd8, 0xac, 0xfe, 0xdd, 0x01,
	0x00, 0x00,
}
