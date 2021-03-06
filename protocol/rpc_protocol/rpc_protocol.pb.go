// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc_protocol.proto

/*
Package rpc_protocol is a generated protocol buffer package.

It is generated from these files:
	rpc_protocol.proto

It has these top-level messages:
	RegisterServiceRequest
	RegisterServiceResponse
	ServiceItem
	QueryServiceListResponse
	UpdateServiceListRequest
*/
package rpc_protocol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "google.golang.org/protobuf/types/known/emptypb"
import error_msg "sparrow/protocol/error_msg"
import msg_login_server "sparrow/protocol/msg_login_server"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// enum_error_code from public import error_msg.proto
type EnumErrorCode error_msg.EnumErrorCode

var EnumErrorCode_name = error_msg.EnumErrorCode_name
var EnumErrorCode_value = error_msg.EnumErrorCode_value

func (x EnumErrorCode) String() string { return (error_msg.EnumErrorCode)(x).String() }

const EnumErrorCode_SUCCESS = EnumErrorCode(error_msg.EnumErrorCode_SUCCESS)
const EnumErrorCode_FAILD = EnumErrorCode(error_msg.EnumErrorCode_FAILD)
const EnumErrorCode_LOGIN_ACCOUNTS_NOT_EXISTS = EnumErrorCode(error_msg.EnumErrorCode_LOGIN_ACCOUNTS_NOT_EXISTS)
const EnumErrorCode_LOGIN_PASSWORD_ERROR = EnumErrorCode(error_msg.EnumErrorCode_LOGIN_PASSWORD_ERROR)
const EnumErrorCode_REGISTER_ACCOUNTS_EXISTS = EnumErrorCode(error_msg.EnumErrorCode_REGISTER_ACCOUNTS_EXISTS)
const EnumErrorCode_REGISTER_ACCOUNTS_FORMAT_ERROR = EnumErrorCode(error_msg.EnumErrorCode_REGISTER_ACCOUNTS_FORMAT_ERROR)
const EnumErrorCode_REGISTER_PASSWORD_FORMAT_ERROR = EnumErrorCode(error_msg.EnumErrorCode_REGISTER_PASSWORD_FORMAT_ERROR)
const EnumErrorCode_ROLE_LIST_EMPTY = EnumErrorCode(error_msg.EnumErrorCode_ROLE_LIST_EMPTY)
const EnumErrorCode_ROLE_LIST_ERROR = EnumErrorCode(error_msg.EnumErrorCode_ROLE_LIST_ERROR)
const EnumErrorCode_ROLE_NAME_EXISTS = EnumErrorCode(error_msg.EnumErrorCode_ROLE_NAME_EXISTS)
const EnumErrorCode_ROLE_NAME_FORMAT_ERROR = EnumErrorCode(error_msg.EnumErrorCode_ROLE_NAME_FORMAT_ERROR)

// ClientLoginReq from public import msg_login_server.proto
type ClientLoginReq msg_login_server.ClientLoginReq

func (m *ClientLoginReq) Reset()         { (*msg_login_server.ClientLoginReq)(m).Reset() }
func (m *ClientLoginReq) String() string { return (*msg_login_server.ClientLoginReq)(m).String() }
func (*ClientLoginReq) ProtoMessage()    {}
func (m *ClientLoginReq) GetAccounts() string {
	return (*msg_login_server.ClientLoginReq)(m).GetAccounts()
}
func (m *ClientLoginReq) GetPassword() string {
	return (*msg_login_server.ClientLoginReq)(m).GetPassword()
}

// ClientLoginRes from public import msg_login_server.proto
type ClientLoginRes msg_login_server.ClientLoginRes

func (m *ClientLoginRes) Reset()         { (*msg_login_server.ClientLoginRes)(m).Reset() }
func (m *ClientLoginRes) String() string { return (*msg_login_server.ClientLoginRes)(m).String() }
func (*ClientLoginRes) ProtoMessage()    {}
func (m *ClientLoginRes) GetUid() uint64 { return (*msg_login_server.ClientLoginRes)(m).GetUid() }

// ClientRegisterReq from public import msg_login_server.proto
type ClientRegisterReq msg_login_server.ClientRegisterReq

func (m *ClientRegisterReq) Reset()         { (*msg_login_server.ClientRegisterReq)(m).Reset() }
func (m *ClientRegisterReq) String() string { return (*msg_login_server.ClientRegisterReq)(m).String() }
func (*ClientRegisterReq) ProtoMessage()    {}
func (m *ClientRegisterReq) GetAccounts() string {
	return (*msg_login_server.ClientRegisterReq)(m).GetAccounts()
}
func (m *ClientRegisterReq) GetPassword() string {
	return (*msg_login_server.ClientRegisterReq)(m).GetPassword()
}
func (m *ClientRegisterReq) GetChannel() int32 {
	return (*msg_login_server.ClientRegisterReq)(m).GetChannel()
}

// ClientRegisterRes from public import msg_login_server.proto
type ClientRegisterRes msg_login_server.ClientRegisterRes

func (m *ClientRegisterRes) Reset()         { (*msg_login_server.ClientRegisterRes)(m).Reset() }
func (m *ClientRegisterRes) String() string { return (*msg_login_server.ClientRegisterRes)(m).String() }
func (*ClientRegisterRes) ProtoMessage()    {}
func (m *ClientRegisterRes) GetUid() uint64 { return (*msg_login_server.ClientRegisterRes)(m).GetUid() }

// QueryRoleListReq from public import msg_login_server.proto
type QueryRoleListReq msg_login_server.QueryRoleListReq

func (m *QueryRoleListReq) Reset()         { (*msg_login_server.QueryRoleListReq)(m).Reset() }
func (m *QueryRoleListReq) String() string { return (*msg_login_server.QueryRoleListReq)(m).String() }
func (*QueryRoleListReq) ProtoMessage()    {}
func (m *QueryRoleListReq) GetUid() uint64 { return (*msg_login_server.QueryRoleListReq)(m).GetUid() }

// QueryRoleListRes from public import msg_login_server.proto
type QueryRoleListRes msg_login_server.QueryRoleListRes

func (m *QueryRoleListRes) Reset()         { (*msg_login_server.QueryRoleListRes)(m).Reset() }
func (m *QueryRoleListRes) String() string { return (*msg_login_server.QueryRoleListRes)(m).String() }
func (*QueryRoleListRes) ProtoMessage()    {}

// CreateRoleReq from public import msg_login_server.proto
type CreateRoleReq msg_login_server.CreateRoleReq

func (m *CreateRoleReq) Reset()          { (*msg_login_server.CreateRoleReq)(m).Reset() }
func (m *CreateRoleReq) String() string  { return (*msg_login_server.CreateRoleReq)(m).String() }
func (*CreateRoleReq) ProtoMessage()     {}
func (m *CreateRoleReq) GetName() string { return (*msg_login_server.CreateRoleReq)(m).GetName() }

// CreateRoleRes from public import msg_login_server.proto
type CreateRoleRes msg_login_server.CreateRoleRes

func (m *CreateRoleRes) Reset()         { (*msg_login_server.CreateRoleRes)(m).Reset() }
func (m *CreateRoleRes) String() string { return (*msg_login_server.CreateRoleRes)(m).String() }
func (*CreateRoleRes) ProtoMessage()    {}

// ????????????
type RegisterServiceRequest struct {
	ServerType uint32 `protobuf:"varint,1,opt,name=server_type,json=serverType" json:"server_type,omitempty"`
	RpcPort    string `protobuf:"bytes,2,opt,name=rpc_port,json=rpcPort" json:"rpc_port,omitempty"`
}

func (m *RegisterServiceRequest) Reset()                    { *m = RegisterServiceRequest{} }
func (m *RegisterServiceRequest) String() string            { return proto.CompactTextString(m) }
func (*RegisterServiceRequest) ProtoMessage()               {}
func (*RegisterServiceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RegisterServiceRequest) GetServerType() uint32 {
	if m != nil {
		return m.ServerType
	}
	return 0
}

func (m *RegisterServiceRequest) GetRpcPort() string {
	if m != nil {
		return m.RpcPort
	}
	return ""
}

type RegisterServiceResponse struct {
	ErrorCode error_msg.EnumErrorCode `protobuf:"varint,1,opt,name=error_code,json=errorCode,enum=error_msg.EnumErrorCode" json:"error_code,omitempty"`
	ServerId  uint32                  `protobuf:"varint,2,opt,name=server_id,json=serverId" json:"server_id,omitempty"`
}

func (m *RegisterServiceResponse) Reset()                    { *m = RegisterServiceResponse{} }
func (m *RegisterServiceResponse) String() string            { return proto.CompactTextString(m) }
func (*RegisterServiceResponse) ProtoMessage()               {}
func (*RegisterServiceResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RegisterServiceResponse) GetErrorCode() error_msg.EnumErrorCode {
	if m != nil {
		return m.ErrorCode
	}
	return error_msg.EnumErrorCode_SUCCESS
}

func (m *RegisterServiceResponse) GetServerId() uint32 {
	if m != nil {
		return m.ServerId
	}
	return 0
}

// ??????????????????
type ServiceItem struct {
	ServerId   uint32 `protobuf:"varint,1,opt,name=server_id,json=serverId" json:"server_id,omitempty"`
	ServerType uint32 `protobuf:"varint,2,opt,name=server_type,json=serverType" json:"server_type,omitempty"`
	Ip         string `protobuf:"bytes,3,opt,name=ip" json:"ip,omitempty"`
	Port       string `protobuf:"bytes,4,opt,name=port" json:"port,omitempty"`
}

func (m *ServiceItem) Reset()                    { *m = ServiceItem{} }
func (m *ServiceItem) String() string            { return proto.CompactTextString(m) }
func (*ServiceItem) ProtoMessage()               {}
func (*ServiceItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ServiceItem) GetServerId() uint32 {
	if m != nil {
		return m.ServerId
	}
	return 0
}

func (m *ServiceItem) GetServerType() uint32 {
	if m != nil {
		return m.ServerType
	}
	return 0
}

func (m *ServiceItem) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *ServiceItem) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

type QueryServiceListResponse struct {
	ErrorCode  error_msg.EnumErrorCode `protobuf:"varint,1,opt,name=error_code,json=errorCode,enum=error_msg.EnumErrorCode" json:"error_code,omitempty"`
	ServerList []*ServiceItem          `protobuf:"bytes,2,rep,name=server_list,json=serverList" json:"server_list,omitempty"`
}

func (m *QueryServiceListResponse) Reset()                    { *m = QueryServiceListResponse{} }
func (m *QueryServiceListResponse) String() string            { return proto.CompactTextString(m) }
func (*QueryServiceListResponse) ProtoMessage()               {}
func (*QueryServiceListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *QueryServiceListResponse) GetErrorCode() error_msg.EnumErrorCode {
	if m != nil {
		return m.ErrorCode
	}
	return error_msg.EnumErrorCode_SUCCESS
}

func (m *QueryServiceListResponse) GetServerList() []*ServiceItem {
	if m != nil {
		return m.ServerList
	}
	return nil
}

// ??????????????????
type UpdateServiceListRequest struct {
	ServerList []*ServiceItem `protobuf:"bytes,1,rep,name=server_list,json=serverList" json:"server_list,omitempty"`
}

func (m *UpdateServiceListRequest) Reset()                    { *m = UpdateServiceListRequest{} }
func (m *UpdateServiceListRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateServiceListRequest) ProtoMessage()               {}
func (*UpdateServiceListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UpdateServiceListRequest) GetServerList() []*ServiceItem {
	if m != nil {
		return m.ServerList
	}
	return nil
}

func init() {
	proto.RegisterType((*RegisterServiceRequest)(nil), "rpc_protocol.RegisterServiceRequest")
	proto.RegisterType((*RegisterServiceResponse)(nil), "rpc_protocol.RegisterServiceResponse")
	proto.RegisterType((*ServiceItem)(nil), "rpc_protocol.ServiceItem")
	proto.RegisterType((*QueryServiceListResponse)(nil), "rpc_protocol.QueryServiceListResponse")
	proto.RegisterType((*UpdateServiceListRequest)(nil), "rpc_protocol.UpdateServiceListRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RpcService service

type RpcServiceClient interface {
	// ????????????
	RegisterService(ctx context.Context, in *RegisterServiceRequest, opts ...grpc.CallOption) (*RegisterServiceResponse, error)
	// ????????????
	QueryServiceList(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*QueryServiceListResponse, error)
	// ????????????
	UpdateServiceList(ctx context.Context, in *UpdateServiceListRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	// ????????????
	ClientRegister(ctx context.Context, in *msg_login_server.ClientRegisterReq, opts ...grpc.CallOption) (*msg_login_server.ClientRegisterRes, error)
	// ????????????
	ClientLogin(ctx context.Context, in *msg_login_server.ClientLoginReq, opts ...grpc.CallOption) (*msg_login_server.ClientLoginRes, error)
	// ??????????????????
	QueryRoleList(ctx context.Context, in *msg_login_server.QueryRoleListReq, opts ...grpc.CallOption) (*msg_login_server.QueryRoleListRes, error)
	// ????????????
	CreateRole(ctx context.Context, in *msg_login_server.CreateRoleReq, opts ...grpc.CallOption) (*msg_login_server.CreateRoleRes, error)
}

type rpcServiceClient struct {
	cc *grpc.ClientConn
}

func NewRpcServiceClient(cc *grpc.ClientConn) RpcServiceClient {
	return &rpcServiceClient{cc}
}

func (c *rpcServiceClient) RegisterService(ctx context.Context, in *RegisterServiceRequest, opts ...grpc.CallOption) (*RegisterServiceResponse, error) {
	out := new(RegisterServiceResponse)
	err := grpc.Invoke(ctx, "/rpc_protocol.rpc_service/RegisterService", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcServiceClient) QueryServiceList(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*QueryServiceListResponse, error) {
	out := new(QueryServiceListResponse)
	err := grpc.Invoke(ctx, "/rpc_protocol.rpc_service/QueryServiceList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcServiceClient) UpdateServiceList(ctx context.Context, in *UpdateServiceListRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/rpc_protocol.rpc_service/UpdateServiceList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcServiceClient) ClientRegister(ctx context.Context, in *msg_login_server.ClientRegisterReq, opts ...grpc.CallOption) (*msg_login_server.ClientRegisterRes, error) {
	out := new(msg_login_server.ClientRegisterRes)
	err := grpc.Invoke(ctx, "/rpc_protocol.rpc_service/ClientRegister", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcServiceClient) ClientLogin(ctx context.Context, in *msg_login_server.ClientLoginReq, opts ...grpc.CallOption) (*msg_login_server.ClientLoginRes, error) {
	out := new(msg_login_server.ClientLoginRes)
	err := grpc.Invoke(ctx, "/rpc_protocol.rpc_service/ClientLogin", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcServiceClient) QueryRoleList(ctx context.Context, in *msg_login_server.QueryRoleListReq, opts ...grpc.CallOption) (*msg_login_server.QueryRoleListRes, error) {
	out := new(msg_login_server.QueryRoleListRes)
	err := grpc.Invoke(ctx, "/rpc_protocol.rpc_service/QueryRoleList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcServiceClient) CreateRole(ctx context.Context, in *msg_login_server.CreateRoleReq, opts ...grpc.CallOption) (*msg_login_server.CreateRoleRes, error) {
	out := new(msg_login_server.CreateRoleRes)
	err := grpc.Invoke(ctx, "/rpc_protocol.rpc_service/CreateRole", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RpcService service

type RpcServiceServer interface {
	// ????????????
	RegisterService(context.Context, *RegisterServiceRequest) (*RegisterServiceResponse, error)
	// ????????????
	QueryServiceList(context.Context, *google_protobuf.Empty) (*QueryServiceListResponse, error)
	// ????????????
	UpdateServiceList(context.Context, *UpdateServiceListRequest) (*google_protobuf.Empty, error)
	// ????????????
	ClientRegister(context.Context, *msg_login_server.ClientRegisterReq) (*msg_login_server.ClientRegisterRes, error)
	// ????????????
	ClientLogin(context.Context, *msg_login_server.ClientLoginReq) (*msg_login_server.ClientLoginRes, error)
	// ??????????????????
	QueryRoleList(context.Context, *msg_login_server.QueryRoleListReq) (*msg_login_server.QueryRoleListRes, error)
	// ????????????
	CreateRole(context.Context, *msg_login_server.CreateRoleReq) (*msg_login_server.CreateRoleRes, error)
}

func RegisterRpcServiceServer(s *grpc.Server, srv RpcServiceServer) {
	s.RegisterService(&_RpcService_serviceDesc, srv)
}

func _RpcService_RegisterService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServiceServer).RegisterService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc_protocol.rpc_service/RegisterService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServiceServer).RegisterService(ctx, req.(*RegisterServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcService_QueryServiceList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServiceServer).QueryServiceList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc_protocol.rpc_service/QueryServiceList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServiceServer).QueryServiceList(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcService_UpdateServiceList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateServiceListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServiceServer).UpdateServiceList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc_protocol.rpc_service/UpdateServiceList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServiceServer).UpdateServiceList(ctx, req.(*UpdateServiceListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcService_ClientRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(msg_login_server.ClientRegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServiceServer).ClientRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc_protocol.rpc_service/ClientRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServiceServer).ClientRegister(ctx, req.(*msg_login_server.ClientRegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcService_ClientLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(msg_login_server.ClientLoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServiceServer).ClientLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc_protocol.rpc_service/ClientLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServiceServer).ClientLogin(ctx, req.(*msg_login_server.ClientLoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcService_QueryRoleList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(msg_login_server.QueryRoleListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServiceServer).QueryRoleList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc_protocol.rpc_service/QueryRoleList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServiceServer).QueryRoleList(ctx, req.(*msg_login_server.QueryRoleListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcService_CreateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(msg_login_server.CreateRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServiceServer).CreateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc_protocol.rpc_service/CreateRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServiceServer).CreateRole(ctx, req.(*msg_login_server.CreateRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _RpcService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc_protocol.rpc_service",
	HandlerType: (*RpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterService",
			Handler:    _RpcService_RegisterService_Handler,
		},
		{
			MethodName: "QueryServiceList",
			Handler:    _RpcService_QueryServiceList_Handler,
		},
		{
			MethodName: "UpdateServiceList",
			Handler:    _RpcService_UpdateServiceList_Handler,
		},
		{
			MethodName: "ClientRegister",
			Handler:    _RpcService_ClientRegister_Handler,
		},
		{
			MethodName: "ClientLogin",
			Handler:    _RpcService_ClientLogin_Handler,
		},
		{
			MethodName: "QueryRoleList",
			Handler:    _RpcService_QueryRoleList_Handler,
		},
		{
			MethodName: "CreateRole",
			Handler:    _RpcService_CreateRole_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc_protocol.proto",
}

func init() { proto.RegisterFile("rpc_protocol.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 487 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x95, 0xd3, 0x08, 0x9a, 0x31, 0x49, 0xcb, 0x1c, 0x82, 0xeb, 0x1e, 0x1a, 0x99, 0x0f, 0xe5,
	0xe4, 0x48, 0xe1, 0x04, 0x47, 0x22, 0x0e, 0x95, 0x2a, 0x94, 0xba, 0x05, 0xa4, 0x1e, 0xb0, 0x5a,
	0x7b, 0xb0, 0x2c, 0xd9, 0xd9, 0xcd, 0xee, 0xa6, 0x52, 0x7e, 0x06, 0x7f, 0x94, 0xdf, 0x80, 0x76,
	0xd7, 0x69, 0x63, 0xbb, 0x51, 0x40, 0xe2, 0xb6, 0x7e, 0xfb, 0xf6, 0xcd, 0x7b, 0x33, 0x63, 0x40,
	0xc1, 0x93, 0x98, 0x0b, 0xa6, 0x58, 0xc2, 0x8a, 0xd0, 0x1c, 0xf0, 0xc5, 0x36, 0xe6, 0x9f, 0x66,
	0x8c, 0x65, 0x05, 0x4d, 0x0c, 0x70, 0xb7, 0xfa, 0x39, 0xa1, 0x92, 0xab, 0xb5, 0xa5, 0xfa, 0x47,
	0x24, 0x04, 0x13, 0x71, 0x29, 0xb3, 0x0a, 0x18, 0x96, 0x32, 0x8b, 0x0b, 0x96, 0xe5, 0x8b, 0x58,
	0x92, 0xb8, 0x27, 0x61, 0xf1, 0xe0, 0x1a, 0x86, 0x11, 0x65, 0xb9, 0x54, 0x24, 0xae, 0x48, 0xdc,
	0xe7, 0x09, 0x45, 0xb4, 0x5c, 0x91, 0x54, 0x78, 0x06, 0xae, 0x65, 0xc6, 0x6a, 0xcd, 0xc9, 0x73,
	0x46, 0xce, 0xb8, 0x1f, 0x81, 0x85, 0xae, 0xd7, 0x9c, 0xf0, 0x04, 0x0e, 0x8d, 0x21, 0x26, 0x94,
	0xd7, 0x19, 0x39, 0xe3, 0x5e, 0xf4, 0x5c, 0xf0, 0x64, 0xce, 0x84, 0x0a, 0x96, 0xf0, 0xaa, 0xa5,
	0x2a, 0x39, 0x5b, 0x48, 0xc2, 0x0f, 0x00, 0xd6, 0x5b, 0xc2, 0x52, 0xab, 0x3a, 0x98, 0xfa, 0xe1,
	0xa3, 0x5d, 0x5a, 0xac, 0xca, 0xf8, 0x91, 0x11, 0xf5, 0xcc, 0x79, 0xc6, 0x52, 0xc2, 0x53, 0xe8,
	0x55, 0x8e, 0xf2, 0xd4, 0x54, 0xec, 0x47, 0x87, 0x16, 0x38, 0x4f, 0x03, 0x06, 0x6e, 0x55, 0xea,
	0x5c, 0x51, 0x59, 0xe7, 0x3a, 0x75, 0x6e, 0x33, 0x5a, 0xa7, 0x15, 0x6d, 0x00, 0x9d, 0x9c, 0x7b,
	0x07, 0x26, 0x54, 0x27, 0xe7, 0x88, 0xd0, 0x35, 0x31, 0xbb, 0x06, 0x31, 0xe7, 0xe0, 0x97, 0x03,
	0xde, 0xe5, 0x8a, 0xc4, 0xba, 0x2a, 0x7b, 0x91, 0x4b, 0xf5, 0x3f, 0x52, 0x7e, 0x7c, 0x30, 0x57,
	0xe4, 0x52, 0x77, 0xf6, 0x60, 0xec, 0x4e, 0x4f, 0xc2, 0xda, 0x3e, 0x6c, 0x25, 0xdd, 0xf8, 0xd6,
	0xe5, 0x83, 0x6f, 0xe0, 0x7d, 0xe5, 0xe9, 0xad, 0xa2, 0x9a, 0x27, 0x3b, 0xcf, 0x86, 0xae, 0xf3,
	0x0f, 0xba, 0xd3, 0xdf, 0x5d, 0x70, 0x35, 0x51, 0xda, 0x7b, 0xfc, 0x01, 0x47, 0x8d, 0xf9, 0xe2,
	0x9b, 0xba, 0xd2, 0xd3, 0x4b, 0xe5, 0xbf, 0xdd, 0xc3, 0xaa, 0xda, 0x17, 0xc1, 0x71, 0xb3, 0xb5,
	0x38, 0x0c, 0xed, 0xc2, 0x87, 0x9b, 0x85, 0x0f, 0x3f, 0xeb, 0x85, 0xf7, 0xdf, 0xd5, 0x25, 0x77,
	0x8e, 0xe4, 0x0a, 0x5e, 0xb6, 0x7a, 0x83, 0x8d, 0xc7, 0xbb, 0x9a, 0xe7, 0xef, 0x28, 0x8e, 0x37,
	0x30, 0x98, 0x15, 0x39, 0x2d, 0xd4, 0x26, 0x09, 0xbe, 0x0e, 0x5b, 0x7f, 0x5a, 0x9d, 0x11, 0xd1,
	0xd2, 0xff, 0x0b, 0x92, 0xc4, 0x4b, 0x70, 0x2d, 0x78, 0xa1, 0x79, 0x38, 0xda, 0xf5, 0xc6, 0x5c,
	0x6b, 0xd5, 0x7d, 0x0c, 0x89, 0xdf, 0xa1, 0x6f, 0xfa, 0x13, 0xb1, 0xc2, 0xe6, 0x0f, 0xda, 0x4f,
	0x6a, 0x04, 0x2d, 0xbb, 0x9f, 0x23, 0xf1, 0x0b, 0xc0, 0x4c, 0xd0, 0xad, 0x22, 0x0d, 0xe2, 0xd9,
	0x13, 0x46, 0x1e, 0x6e, 0xb5, 0xe4, 0x1e, 0x82, 0xfc, 0x74, 0x7c, 0x33, 0x08, 0x27, 0xdb, 0xa3,
	0x99, 0x3b, 0xf3, 0xce, 0xdd, 0x33, 0xf3, 0xf5, 0xfe, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5a,
	0x43, 0x41, 0x41, 0x1d, 0x05, 0x00, 0x00,
}
