// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc_protocol.proto

/*
Package rpc_protocol is a generated protocol buffer package.

指定等会文件生成出来的package

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

// 注册服务
type RegisterServiceRequest struct {
	ServerType uint32 `protobuf:"varint,2,opt,name=server_type,json=serverType" json:"server_type,omitempty"`
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

type RegisterServiceResponse struct {
	ErrorCode int32  `protobuf:"varint,1,opt,name=error_code,json=errorCode" json:"error_code,omitempty"`
	ServerId  uint32 `protobuf:"varint,2,opt,name=server_id,json=serverId" json:"server_id,omitempty"`
}

func (m *RegisterServiceResponse) Reset()                    { *m = RegisterServiceResponse{} }
func (m *RegisterServiceResponse) String() string            { return proto.CompactTextString(m) }
func (*RegisterServiceResponse) ProtoMessage()               {}
func (*RegisterServiceResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RegisterServiceResponse) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func (m *RegisterServiceResponse) GetServerId() uint32 {
	if m != nil {
		return m.ServerId
	}
	return 0
}

// 请求服务列表
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
	ErrorCode  int32          `protobuf:"varint,1,opt,name=error_code,json=errorCode" json:"error_code,omitempty"`
	ServerList []*ServiceItem `protobuf:"bytes,2,rep,name=server_list,json=serverList" json:"server_list,omitempty"`
}

func (m *QueryServiceListResponse) Reset()                    { *m = QueryServiceListResponse{} }
func (m *QueryServiceListResponse) String() string            { return proto.CompactTextString(m) }
func (*QueryServiceListResponse) ProtoMessage()               {}
func (*QueryServiceListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *QueryServiceListResponse) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func (m *QueryServiceListResponse) GetServerList() []*ServiceItem {
	if m != nil {
		return m.ServerList
	}
	return nil
}

// 更新服务列表
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

// Client API for RpcProtocol service

type RpcProtocolClient interface {
	// 服务注册
	RegisterService(ctx context.Context, in *RegisterServiceRequest, opts ...grpc.CallOption) (*RegisterServiceResponse, error)
	// 服务请求
	QueryServiceList(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*QueryServiceListResponse, error)
	// 服务更新
	UpdateServiceList(ctx context.Context, in *UpdateServiceListRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
}

type rpcProtocolClient struct {
	cc *grpc.ClientConn
}

func NewRpcProtocolClient(cc *grpc.ClientConn) RpcProtocolClient {
	return &rpcProtocolClient{cc}
}

func (c *rpcProtocolClient) RegisterService(ctx context.Context, in *RegisterServiceRequest, opts ...grpc.CallOption) (*RegisterServiceResponse, error) {
	out := new(RegisterServiceResponse)
	err := grpc.Invoke(ctx, "/rpc_protocol.rpc_protocol/RegisterService", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcProtocolClient) QueryServiceList(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*QueryServiceListResponse, error) {
	out := new(QueryServiceListResponse)
	err := grpc.Invoke(ctx, "/rpc_protocol.rpc_protocol/QueryServiceList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcProtocolClient) UpdateServiceList(ctx context.Context, in *UpdateServiceListRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/rpc_protocol.rpc_protocol/UpdateServiceList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RpcProtocol service

type RpcProtocolServer interface {
	// 服务注册
	RegisterService(context.Context, *RegisterServiceRequest) (*RegisterServiceResponse, error)
	// 服务请求
	QueryServiceList(context.Context, *google_protobuf.Empty) (*QueryServiceListResponse, error)
	// 服务更新
	UpdateServiceList(context.Context, *UpdateServiceListRequest) (*google_protobuf.Empty, error)
}

func RegisterRpcProtocolServer(s *grpc.Server, srv RpcProtocolServer) {
	s.RegisterService(&_RpcProtocol_serviceDesc, srv)
}

func _RpcProtocol_RegisterService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcProtocolServer).RegisterService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc_protocol.rpc_protocol/RegisterService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcProtocolServer).RegisterService(ctx, req.(*RegisterServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcProtocol_QueryServiceList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcProtocolServer).QueryServiceList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc_protocol.rpc_protocol/QueryServiceList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcProtocolServer).QueryServiceList(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _RpcProtocol_UpdateServiceList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateServiceListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcProtocolServer).UpdateServiceList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc_protocol.rpc_protocol/UpdateServiceList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcProtocolServer).UpdateServiceList(ctx, req.(*UpdateServiceListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RpcProtocol_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc_protocol.rpc_protocol",
	HandlerType: (*RpcProtocolServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterService",
			Handler:    _RpcProtocol_RegisterService_Handler,
		},
		{
			MethodName: "QueryServiceList",
			Handler:    _RpcProtocol_QueryServiceList_Handler,
		},
		{
			MethodName: "UpdateServiceList",
			Handler:    _RpcProtocol_UpdateServiceList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc_protocol.proto",
}

func init() { proto.RegisterFile("rpc_protocol.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xcf, 0x4f, 0xc2, 0x30,
	0x18, 0xcd, 0x06, 0x1a, 0xf9, 0xf0, 0x67, 0x0f, 0x58, 0x47, 0x8c, 0xcb, 0xa2, 0x66, 0xa7, 0x91,
	0xe0, 0x49, 0xaf, 0xc6, 0x03, 0x89, 0x17, 0x87, 0x78, 0x94, 0xc8, 0xf6, 0x49, 0x9a, 0x80, 0xad,
	0x6d, 0x47, 0xb2, 0xbf, 0xc4, 0x7f, 0xd7, 0xd0, 0x0d, 0x43, 0x87, 0x0b, 0x7a, 0x6b, 0x5f, 0x5e,
	0xdf, 0xeb, 0xf7, 0xbe, 0x07, 0x44, 0x8a, 0x64, 0x2c, 0x24, 0xd7, 0x3c, 0xe1, 0xb3, 0xc8, 0x1c,
	0xc8, 0xfe, 0x3a, 0xe6, 0x75, 0xa7, 0x9c, 0x4f, 0x67, 0xd8, 0x33, 0xc0, 0x24, 0x7b, 0xef, 0xe1,
	0x5c, 0xe8, 0xbc, 0xa0, 0x06, 0xb7, 0xd0, 0x89, 0x71, 0xca, 0x94, 0x46, 0x39, 0x44, 0xb9, 0x60,
	0x09, 0xc6, 0xf8, 0x99, 0xa1, 0xd2, 0xe4, 0x02, 0xda, 0x0a, 0xe5, 0x02, 0xe5, 0x58, 0xe7, 0x02,
	0xa9, 0xeb, 0x3b, 0xe1, 0x41, 0x0c, 0x05, 0xf4, 0x9c, 0x0b, 0x0c, 0x46, 0x70, 0xba, 0xf1, 0x54,
	0x09, 0xfe, 0xa1, 0x90, 0x9c, 0x03, 0xa0, 0x94, 0x5c, 0x8e, 0x13, 0x9e, 0x22, 0x75, 0x7c, 0x27,
	0xdc, 0x89, 0x5b, 0x06, 0xb9, 0xe7, 0x29, 0x92, 0x2e, 0xb4, 0x4a, 0x69, 0x96, 0x96, 0xc2, 0x7b,
	0x05, 0x30, 0x48, 0x03, 0x0e, 0xed, 0x52, 0x6e, 0xa0, 0x71, 0x6e, 0x73, 0x1d, 0x9b, 0xbb, 0xf5,
	0x8f, 0xe4, 0x10, 0x5c, 0x26, 0x68, 0xc3, 0x77, 0xc2, 0x56, 0xec, 0x32, 0x41, 0x08, 0x34, 0x05,
	0x97, 0x9a, 0x36, 0x0d, 0x62, 0xce, 0x41, 0x06, 0xf4, 0x29, 0x43, 0x99, 0x97, 0xae, 0x8f, 0x4c,
	0xe9, 0xbf, 0x0e, 0x72, 0xf7, 0xe3, 0x3f, 0x63, 0x4a, 0x53, 0xd7, 0x6f, 0x84, 0xed, 0xfe, 0x59,
	0x64, 0xad, 0x64, 0x6d, 0x98, 0xd5, 0xd7, 0x96, 0x16, 0xc1, 0x0b, 0xd0, 0x91, 0x48, 0xdf, 0x34,
	0x5a, 0xbe, 0x45, 0xf6, 0x15, 0x5d, 0xe7, 0x1f, 0xba, 0xfd, 0x2f, 0x17, 0xac, 0xfd, 0x93, 0x57,
	0x38, 0xaa, 0xec, 0x89, 0x5c, 0xda, 0x52, 0xbf, 0x37, 0xc0, 0xbb, 0xda, 0xc2, 0x2a, 0x33, 0x8a,
	0xe1, 0xb8, 0x9a, 0x1f, 0xe9, 0x44, 0x45, 0xe9, 0xa2, 0x55, 0xe9, 0xa2, 0x87, 0x65, 0xe9, 0xbc,
	0x6b, 0x5b, 0xb2, 0x36, 0xf7, 0x21, 0x9c, 0x6c, 0x84, 0x43, 0x2a, 0x8f, 0xeb, 0xd2, 0xf3, 0x6a,
	0xcc, 0x27, 0xbb, 0xe6, 0x7e, 0xf3, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x75, 0x1b, 0x91, 0x7f, 0x33,
	0x03, 0x00, 0x00,
}