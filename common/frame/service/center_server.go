package service

import (
	"golang.org/x/net/context"
	google_protobuf "google.golang.org/protobuf/types/known/emptypb"
)

type CenterServerService struct {
}
func (ps *CenterServerService) RegisterService(ctx context.Context, request *RegisterServiceRequest) (*RegisterServiceResponse, error) {
	return &RegisterServiceResponse{}, nil
}
func (ps *CenterServerService) QueryServiceList(context.Context, *google_protobuf.Empty) (*QueryServiceListResponse, error){
	return &QueryServiceListResponse{}, nil
}
