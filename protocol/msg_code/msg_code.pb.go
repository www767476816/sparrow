// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msg_code.proto

/*
Package msg_code is a generated protocol buffer package.

It is generated from these files:
	msg_code.proto

It has these top-level messages:
*/
package msg_code

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// 10001-19999中心服，20000-29999分发服，30000-39999登录服，40000-49999逻辑服，50000-59999道具服，60000-69999聊天服
type EnumMsgCode int32

const (
	EnumMsgCode_DEFAULT                EnumMsgCode = 0
	EnumMsgCode_MIN                    EnumMsgCode = 10000
	EnumMsgCode_CS_HEART_BEAT_REQ      EnumMsgCode = 20001
	EnumMsgCode_SC_HEART_BEAT_RES      EnumMsgCode = 20002
	EnumMsgCode_CS_CLIENT_LOGIN_REQ    EnumMsgCode = 30001
	EnumMsgCode_SC_CLIENT_LOGIN_RES    EnumMsgCode = 30002
	EnumMsgCode_CS_CLIENT_REGISTER_REQ EnumMsgCode = 30003
	EnumMsgCode_SC_CLIENT_REGISTER_RES EnumMsgCode = 30004
	EnumMsgCode_CS_QUERY_ROLE_LIST_REQ EnumMsgCode = 30005
	EnumMsgCode_SC_QUERY_ROLE_LIST_RES EnumMsgCode = 30006
	EnumMsgCode_CS_CREATE_ROLE_REQ     EnumMsgCode = 30007
	EnumMsgCode_SC_CREATE_ROLE_RES     EnumMsgCode = 30008
	EnumMsgCode_MAX                    EnumMsgCode = 100000
)

var EnumMsgCode_name = map[int32]string{
	0:      "DEFAULT",
	10000:  "MIN",
	20001:  "CS_HEART_BEAT_REQ",
	20002:  "SC_HEART_BEAT_RES",
	30001:  "CS_CLIENT_LOGIN_REQ",
	30002:  "SC_CLIENT_LOGIN_RES",
	30003:  "CS_CLIENT_REGISTER_REQ",
	30004:  "SC_CLIENT_REGISTER_RES",
	30005:  "CS_QUERY_ROLE_LIST_REQ",
	30006:  "SC_QUERY_ROLE_LIST_RES",
	30007:  "CS_CREATE_ROLE_REQ",
	30008:  "SC_CREATE_ROLE_RES",
	100000: "MAX",
}
var EnumMsgCode_value = map[string]int32{
	"DEFAULT":                0,
	"MIN":                    10000,
	"CS_HEART_BEAT_REQ":      20001,
	"SC_HEART_BEAT_RES":      20002,
	"CS_CLIENT_LOGIN_REQ":    30001,
	"SC_CLIENT_LOGIN_RES":    30002,
	"CS_CLIENT_REGISTER_REQ": 30003,
	"SC_CLIENT_REGISTER_RES": 30004,
	"CS_QUERY_ROLE_LIST_REQ": 30005,
	"SC_QUERY_ROLE_LIST_RES": 30006,
	"CS_CREATE_ROLE_REQ":     30007,
	"SC_CREATE_ROLE_RES":     30008,
	"MAX":                    100000,
}

func (x EnumMsgCode) String() string {
	return proto.EnumName(EnumMsgCode_name, int32(x))
}
func (EnumMsgCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterEnum("msg_code.EnumMsgCode", EnumMsgCode_name, EnumMsgCode_value)
}

func init() { proto.RegisterFile("msg_code.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0x2d, 0x4e, 0x8f,
	0x4f, 0xce, 0x4f, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0xb5, 0x8e,
	0x32, 0x71, 0xf1, 0xa6, 0xe6, 0x95, 0xe6, 0xc6, 0xc3, 0x44, 0x84, 0xb8, 0xb9, 0xd8, 0x5d, 0x5c,
	0xdd, 0x1c, 0x43, 0x7d, 0x42, 0x04, 0x18, 0x84, 0x38, 0xb8, 0x98, 0x7d, 0x3d, 0xfd, 0x04, 0x26,
	0xf8, 0x09, 0x89, 0x73, 0x09, 0x3a, 0x07, 0xc7, 0x7b, 0xb8, 0x3a, 0x06, 0x85, 0xc4, 0x3b, 0xb9,
	0x3a, 0x86, 0xc4, 0x07, 0xb9, 0x06, 0x0a, 0x2c, 0x9c, 0xc3, 0x08, 0x92, 0x08, 0x76, 0x46, 0x95,
	0x08, 0x16, 0x58, 0x34, 0x87, 0x51, 0x48, 0x92, 0x4b, 0xd8, 0x39, 0x38, 0xde, 0xd9, 0xc7, 0xd3,
	0xd5, 0x2f, 0x24, 0xde, 0xc7, 0xdf, 0xdd, 0xd3, 0x0f, 0xac, 0x67, 0xe3, 0x2b, 0xb0, 0x54, 0xb0,
	0x33, 0xba, 0x54, 0xb0, 0xc0, 0xa6, 0x57, 0x8c, 0x42, 0x32, 0x5c, 0x62, 0x08, 0x5d, 0x41, 0xae,
	0xee, 0x9e, 0xc1, 0x21, 0xae, 0x41, 0x60, 0x8d, 0x9b, 0x21, 0xb2, 0x08, 0x8d, 0x48, 0xb2, 0xc1,
	0x02, 0x5b, 0xe0, 0x7a, 0x03, 0x43, 0x5d, 0x83, 0x22, 0xe3, 0x83, 0xfc, 0x7d, 0x5c, 0xe3, 0x7d,
	0x3c, 0x83, 0x21, 0x0e, 0xdd, 0x0a, 0xd7, 0x8b, 0x29, 0x1b, 0x2c, 0xb0, 0xed, 0x15, 0xa3, 0x90,
	0x04, 0x97, 0x10, 0xc8, 0xde, 0x20, 0x57, 0xc7, 0x10, 0x57, 0x88, 0x34, 0x48, 0xdf, 0x76, 0x88,
	0x0c, 0xc8, 0x4e, 0x14, 0x99, 0x60, 0x81, 0x1d, 0xaf, 0x18, 0x85, 0x38, 0xb9, 0x98, 0x7d, 0x1d,
	0x23, 0x04, 0x16, 0xf4, 0xb2, 0x39, 0x49, 0x47, 0x49, 0x16, 0x17, 0x24, 0x16, 0x15, 0xe5, 0x97,
	0xeb, 0x83, 0x83, 0x38, 0x39, 0x3f, 0x47, 0x1f, 0x16, 0xa4, 0x49, 0x6c, 0x60, 0x21, 0x63, 0x40,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xa4, 0xe1, 0xbe, 0x0c, 0x87, 0x01, 0x00, 0x00,
}