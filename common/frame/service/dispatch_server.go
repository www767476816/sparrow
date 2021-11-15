package service
import (
	"golang.org/x/net/context"
	google_protobuf "google.golang.org/protobuf/types/known/emptypb"
)

type DispatchServerService struct {
}
func (ps *DispatchServerService) UpdateServiceList(ctx context.Context, request *UpdateServiceListRequest) (*google_protobuf.Empty, error) {
	return &google_protobuf.Empty{}, nil
}
