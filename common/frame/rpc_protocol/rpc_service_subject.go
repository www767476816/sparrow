package rpc_protocol

import (
	"golang.org/x/net/context"
	google_protobuf "google.golang.org/protobuf/types/known/emptypb"
)

type RpcProtocolService struct {
}
func (ps *RpcProtocolService) RegisterService(ctx context.Context, request *RegisterServiceRequest) (*RegisterServiceResponse, error) {
	return &RegisterServiceResponse{}, nil
}
func (ps *RpcProtocolService) QueryServiceList(context.Context, *google_protobuf.Empty) (*QueryServiceListResponse, error){
	return &QueryServiceListResponse{}, nil
}
func (ps *RpcProtocolService) UpdateServiceList(ctx context.Context, request *UpdateServiceListRequest) (*google_protobuf.Empty, error) {
	return &google_protobuf.Empty{}, nil
}
