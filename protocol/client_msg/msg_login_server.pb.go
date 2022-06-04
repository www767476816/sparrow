// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msg_login_server.proto

/*
Package client_msg is a generated protocol buffer package.

It is generated from these files:
	msg_login_server.proto

It has these top-level messages:
	ClientLoginReq
	ClientLoginRes
	ClientRegisterReq
	ClientRegisterRes
	QueryRoleListReq
	QueryRoleListRes
	CreateRoleReq
	CreateRoleRes
*/
package client_msg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
<<<<<<< HEAD:protocol/client_msg/msg_login_server.pb.go
import client_msg1 "sparrow/protocol/client_msg"
import client_msg2 "sparrow/protocol/client_msg"
=======
import msg_common "msg_common"
import error_msg "error_msg"
>>>>>>> 302d7dd3c10eca5474e92ca926dc8a331c8d7d3a:protocol/msg_login_server/msg_login_server.pb.go

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// 用户登录
type ClientLoginReq struct {
	Accounts string `protobuf:"bytes,1,opt,name=accounts" json:"accounts,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *ClientLoginReq) Reset()                    { *m = ClientLoginReq{} }
func (m *ClientLoginReq) String() string            { return proto.CompactTextString(m) }
func (*ClientLoginReq) ProtoMessage()               {}
func (*ClientLoginReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ClientLoginReq) GetAccounts() string {
	if m != nil {
		return m.Accounts
	}
	return ""
}

func (m *ClientLoginReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type ClientLoginRes struct {
	ErrorCode client_msg2.EnumErrorCode `protobuf:"varint,1,opt,name=error_code,json=errorCode,enum=client_msg.EnumErrorCode" json:"error_code,omitempty"`
	Uid       uint64                    `protobuf:"varint,2,opt,name=uid" json:"uid,omitempty"`
}

func (m *ClientLoginRes) Reset()                    { *m = ClientLoginRes{} }
func (m *ClientLoginRes) String() string            { return proto.CompactTextString(m) }
func (*ClientLoginRes) ProtoMessage()               {}
func (*ClientLoginRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ClientLoginRes) GetErrorCode() client_msg2.EnumErrorCode {
	if m != nil {
		return m.ErrorCode
	}
	return client_msg2.EnumErrorCode_SUCCESS
}

func (m *ClientLoginRes) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

// 用户注册
type ClientRegisterReq struct {
	Accounts string `protobuf:"bytes,1,opt,name=accounts" json:"accounts,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	Channel  int32  `protobuf:"varint,3,opt,name=channel" json:"channel,omitempty"`
}

func (m *ClientRegisterReq) Reset()                    { *m = ClientRegisterReq{} }
func (m *ClientRegisterReq) String() string            { return proto.CompactTextString(m) }
func (*ClientRegisterReq) ProtoMessage()               {}
func (*ClientRegisterReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ClientRegisterReq) GetAccounts() string {
	if m != nil {
		return m.Accounts
	}
	return ""
}

func (m *ClientRegisterReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *ClientRegisterReq) GetChannel() int32 {
	if m != nil {
		return m.Channel
	}
	return 0
}

type ClientRegisterRes struct {
	ErrorCode client_msg2.EnumErrorCode `protobuf:"varint,1,opt,name=error_code,json=errorCode,enum=client_msg.EnumErrorCode" json:"error_code,omitempty"`
	Uid       uint64                    `protobuf:"varint,2,opt,name=uid" json:"uid,omitempty"`
}

func (m *ClientRegisterRes) Reset()                    { *m = ClientRegisterRes{} }
func (m *ClientRegisterRes) String() string            { return proto.CompactTextString(m) }
func (*ClientRegisterRes) ProtoMessage()               {}
func (*ClientRegisterRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ClientRegisterRes) GetErrorCode() client_msg2.EnumErrorCode {
	if m != nil {
		return m.ErrorCode
	}
	return client_msg2.EnumErrorCode_SUCCESS
}

func (m *ClientRegisterRes) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

// 获取角色列表
type QueryRoleListReq struct {
	Uid uint64 `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
}

func (m *QueryRoleListReq) Reset()                    { *m = QueryRoleListReq{} }
func (m *QueryRoleListReq) String() string            { return proto.CompactTextString(m) }
func (*QueryRoleListReq) ProtoMessage()               {}
func (*QueryRoleListReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *QueryRoleListReq) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

type QueryRoleListRes struct {
	ErrorCode client_msg2.EnumErrorCode `protobuf:"varint,1,opt,name=error_code,json=errorCode,enum=client_msg.EnumErrorCode" json:"error_code,omitempty"`
	RoleList  []*client_msg1.Role       `protobuf:"bytes,2,rep,name=role_list,json=roleList" json:"role_list,omitempty"`
}

func (m *QueryRoleListRes) Reset()                    { *m = QueryRoleListRes{} }
func (m *QueryRoleListRes) String() string            { return proto.CompactTextString(m) }
func (*QueryRoleListRes) ProtoMessage()               {}
func (*QueryRoleListRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *QueryRoleListRes) GetErrorCode() client_msg2.EnumErrorCode {
	if m != nil {
		return m.ErrorCode
	}
	return client_msg2.EnumErrorCode_SUCCESS
}

func (m *QueryRoleListRes) GetRoleList() []*client_msg1.Role {
	if m != nil {
		return m.RoleList
	}
	return nil
}

// 创建角色
type CreateRoleReq struct {
	Uid  uint64 `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Head string `protobuf:"bytes,3,opt,name=head" json:"head,omitempty"`
}

func (m *CreateRoleReq) Reset()                    { *m = CreateRoleReq{} }
func (m *CreateRoleReq) String() string            { return proto.CompactTextString(m) }
func (*CreateRoleReq) ProtoMessage()               {}
func (*CreateRoleReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *CreateRoleReq) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *CreateRoleReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateRoleReq) GetHead() string {
	if m != nil {
		return m.Head
	}
	return ""
}

type CreateRoleRes struct {
	ErrorCode error_msg.EnumErrorCode `protobuf:"varint,1,opt,name=error_code,json=errorCode,enum=error_msg.EnumErrorCode" json:"error_code,omitempty"`
	RoleInfo  *msg_common.Role        `protobuf:"bytes,2,opt,name=role_info,json=roleInfo" json:"role_info,omitempty"`
}

func (m *CreateRoleRes) Reset()                    { *m = CreateRoleRes{} }
func (m *CreateRoleRes) String() string            { return proto.CompactTextString(m) }
func (*CreateRoleRes) ProtoMessage()               {}
func (*CreateRoleRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *CreateRoleRes) GetErrorCode() error_msg.EnumErrorCode {
	if m != nil {
		return m.ErrorCode
	}
	return error_msg.EnumErrorCode_SUCCESS
}

func (m *CreateRoleRes) GetRoleInfo() *msg_common.Role {
	if m != nil {
		return m.RoleInfo
	}
	return nil
}

func init() {
<<<<<<< HEAD:protocol/client_msg/msg_login_server.pb.go
	proto.RegisterType((*ClientLoginReq)(nil), "client_msg.ClientLoginReq")
	proto.RegisterType((*ClientLoginRes)(nil), "client_msg.ClientLoginRes")
	proto.RegisterType((*ClientRegisterReq)(nil), "client_msg.ClientRegisterReq")
	proto.RegisterType((*ClientRegisterRes)(nil), "client_msg.ClientRegisterRes")
	proto.RegisterType((*QueryRoleListReq)(nil), "client_msg.QueryRoleListReq")
	proto.RegisterType((*QueryRoleListRes)(nil), "client_msg.QueryRoleListRes")
=======
	proto.RegisterType((*ClientLoginReq)(nil), "msg_login_server.ClientLoginReq")
	proto.RegisterType((*ClientLoginRes)(nil), "msg_login_server.ClientLoginRes")
	proto.RegisterType((*ClientRegisterReq)(nil), "msg_login_server.ClientRegisterReq")
	proto.RegisterType((*ClientRegisterRes)(nil), "msg_login_server.ClientRegisterRes")
	proto.RegisterType((*QueryRoleListReq)(nil), "msg_login_server.QueryRoleListReq")
	proto.RegisterType((*QueryRoleListRes)(nil), "msg_login_server.QueryRoleListRes")
	proto.RegisterType((*CreateRoleReq)(nil), "msg_login_server.CreateRoleReq")
	proto.RegisterType((*CreateRoleRes)(nil), "msg_login_server.CreateRoleRes")
>>>>>>> 302d7dd3c10eca5474e92ca926dc8a331c8d7d3a:protocol/msg_login_server/msg_login_server.pb.go
}

func init() { proto.RegisterFile("msg_login_server.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
<<<<<<< HEAD:protocol/client_msg/msg_login_server.pb.go
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xb1, 0x4f, 0xf3, 0x30,
	0x10, 0xc5, 0x95, 0xf6, 0xfb, 0xa0, 0x35, 0x52, 0x09, 0x19, 0x50, 0xd4, 0x0a, 0x29, 0x8a, 0x18,
	0xb2, 0x90, 0x4a, 0x65, 0x63, 0xa4, 0x0b, 0x43, 0x17, 0x3c, 0x32, 0x10, 0x05, 0xe7, 0x14, 0x2c,
	0x39, 0xbe, 0x70, 0x76, 0xa8, 0x18, 0xf8, 0xdf, 0x91, 0xe3, 0x96, 0x20, 0x3a, 0x96, 0xcd, 0x77,
	0xef, 0x77, 0xef, 0xdd, 0x49, 0x66, 0x97, 0x8d, 0xa9, 0x0b, 0x85, 0xb5, 0xd4, 0x85, 0x01, 0x7a,
	0x07, 0xca, 0x5b, 0x42, 0x8b, 0x11, 0x13, 0x4a, 0x82, 0xb6, 0x45, 0x63, 0xea, 0x79, 0xe8, 0x18,
	0x81, 0x4d, 0x83, 0xda, 0xab, 0xf3, 0x73, 0x20, 0x42, 0x72, 0xa2, 0x6f, 0xa4, 0x0f, 0x6c, 0xb6,
	0xee, 0x07, 0x36, 0xce, 0x8a, 0xc3, 0x5b, 0x34, 0x67, 0x93, 0x52, 0x08, 0xec, 0xb4, 0x35, 0x71,
	0x90, 0x04, 0xd9, 0x94, 0x7f, 0xd7, 0x4e, 0x6b, 0x4b, 0x63, 0xb6, 0x48, 0x55, 0x3c, 0xf2, 0xda,
	0xbe, 0x4e, 0x9f, 0x7f, 0x39, 0x99, 0xe8, 0x8e, 0x31, 0x1f, 0x27, 0xb0, 0x82, 0xde, 0x6b, 0xb6,
	0x5a, 0xe4, 0xc3, 0x7e, 0x39, 0xe8, 0xae, 0x29, 0x06, 0x84, 0x4f, 0xfb, 0xf7, 0x1a, 0x2b, 0x88,
	0x42, 0x36, 0xee, 0xa4, 0x0f, 0xf9, 0xc7, 0xdd, 0x33, 0x05, 0x76, 0xe1, 0xfd, 0x39, 0xd4, 0xd2,
	0x58, 0xa0, 0x23, 0x96, 0x8d, 0x62, 0x76, 0x2a, 0x5e, 0x4b, 0xad, 0x41, 0xc5, 0xe3, 0x24, 0xc8,
	0xfe, 0xf3, 0x7d, 0x99, 0x96, 0x87, 0x31, 0x7f, 0x7d, 0xc9, 0x35, 0x0b, 0x1f, 0x3b, 0xa0, 0x0f,
	0x8e, 0x0a, 0x36, 0xd2, 0x58, 0x77, 0xc8, 0x8e, 0x0a, 0x06, 0xea, 0xf3, 0x80, 0x3a, 0x6e, 0x8f,
	0x1b, 0x36, 0x25, 0x54, 0x50, 0x28, 0x69, 0x6c, 0x3c, 0x4a, 0xc6, 0xd9, 0xd9, 0x2a, 0xfc, 0x39,
	0xea, 0x72, 0xf8, 0x84, 0x76, 0x69, 0xf7, 0x57, 0x4f, 0x0b, 0xd3, 0x96, 0x44, 0xb8, 0x5d, 0xf6,
	0x3f, 0x45, 0xa0, 0x5a, 0x0e, 0xf4, 0xcb, 0x49, 0xdf, 0xbc, 0xfd, 0x0a, 0x00, 0x00, 0xff, 0xff,
	0x8b, 0x05, 0x6c, 0xa7, 0x87, 0x02, 0x00, 0x00,
=======
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xcf, 0x4b, 0xf3, 0x40,
	0x10, 0x25, 0x6d, 0xbf, 0xcf, 0x66, 0xc5, 0x1a, 0x17, 0x91, 0x90, 0x53, 0x09, 0x1e, 0x7a, 0x31,
	0x42, 0x3d, 0x79, 0xb5, 0x17, 0x0b, 0xbd, 0xb8, 0x47, 0x41, 0x62, 0x4c, 0xa6, 0xe9, 0x42, 0xb2,
	0x53, 0x67, 0x37, 0x4a, 0xc1, 0x3f, 0x5e, 0xb2, 0x69, 0xad, 0xa6, 0x7a, 0x6a, 0x6f, 0x33, 0x6f,
	0x5f, 0xde, 0x8f, 0x65, 0xc3, 0x2e, 0x4a, 0x9d, 0xc7, 0x05, 0xe6, 0x52, 0xc5, 0x1a, 0xe8, 0x0d,
	0x28, 0x5a, 0x12, 0x1a, 0xe4, 0x5e, 0x1b, 0x0f, 0x2c, 0x92, 0x62, 0x59, 0xa2, 0x6a, 0x38, 0xc1,
	0x29, 0x10, 0x21, 0xc5, 0xa5, 0xce, 0x1b, 0x20, 0xbc, 0x67, 0x83, 0x49, 0x21, 0x41, 0x99, 0x59,
	0xfd, 0xa1, 0x80, 0x57, 0x1e, 0xb0, 0x7e, 0x92, 0xa6, 0x58, 0x29, 0xa3, 0x7d, 0x67, 0xe8, 0x8c,
	0x5c, 0xf1, 0xb5, 0xd7, 0x67, 0xcb, 0x44, 0xeb, 0x77, 0xa4, 0xcc, 0xef, 0x34, 0x67, 0x9b, 0x3d,
	0x7c, 0x6a, 0x29, 0x69, 0x7e, 0xcb, 0x58, 0x63, 0x97, 0x62, 0x06, 0x56, 0x6b, 0x30, 0x0e, 0xa2,
	0x6d, 0x02, 0x50, 0x55, 0x19, 0x6f, 0x19, 0xc2, 0xb5, 0xf3, 0x04, 0x33, 0xe0, 0x1e, 0xeb, 0x56,
	0xb2, 0xf1, 0xe8, 0x89, 0x7a, 0x0c, 0x81, 0x9d, 0x35, 0xf2, 0x02, 0x72, 0xa9, 0x0d, 0xd0, 0x1e,
	0x59, 0xb9, 0xcf, 0x8e, 0xd2, 0x45, 0xa2, 0x14, 0x14, 0x7e, 0x77, 0xe8, 0x8c, 0xfe, 0x89, 0xcd,
	0x1a, 0x3e, 0xef, 0xda, 0x1c, 0xb8, 0xc8, 0x25, 0xf3, 0x1e, 0x2a, 0xa0, 0x95, 0xc0, 0x02, 0x66,
	0x52, 0x9b, 0xba, 0xc7, 0x9a, 0xe5, 0x6c, 0x59, 0x1f, 0x3b, 0xac, 0xbd, 0x62, 0x5c, 0x31, 0x97,
	0xb0, 0x80, 0xb8, 0x90, 0xda, 0xf8, 0x9d, 0x61, 0x77, 0x74, 0x3c, 0xf6, 0xa2, 0x6f, 0xaf, 0xa3,
	0xb6, 0x11, 0x7d, 0x5a, 0x9b, 0x85, 0x53, 0x76, 0x32, 0x21, 0x48, 0x0c, 0x58, 0xfc, 0xb7, 0x80,
	0x9c, 0xb3, 0x9e, 0x4a, 0x4a, 0x58, 0x5f, 0xad, 0x9d, 0x6b, 0x6c, 0x01, 0x49, 0x66, 0xef, 0xd4,
	0x15, 0x76, 0x0e, 0x57, 0x3f, 0xa5, 0x0e, 0xd2, 0x42, 0xaa, 0x39, 0x5a, 0xe3, 0x3f, 0x5b, 0x4c,
	0xd5, 0x1c, 0xef, 0xce, 0x1f, 0x79, 0x74, 0xdd, 0xfe, 0x29, 0x5e, 0xfe, 0xdb, 0x87, 0x7f, 0xf3,
	0x19, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x64, 0x3d, 0x7f, 0x47, 0x03, 0x00, 0x00,
>>>>>>> 302d7dd3c10eca5474e92ca926dc8a331c8d7d3a:protocol/msg_login_server/msg_login_server.pb.go
}
