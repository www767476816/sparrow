/*
 * @Author: your name
 * @Date: 2021-12-27 10:52:36
 * @LastEditTime: 2021-12-27 17:18:21
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: \sparrow\login_server\base\rpc.go
 */
package base

import (
	"errors"
	"fmt"
	"sparrow/protocol/msg_server"

	"golang.org/x/net/context"
	google_protobuf "google.golang.org/protobuf/types/known/emptypb"
)

func (this *Base) RegisterRpcService() {
	//注册服务器
	msg_server.RegisterRpcServiceServer(this.Frame.GetRpcServer().GetConnect(), new(LoginRpc))
}
func (this *Base) SetRpcClient(serverID uint32) bool {
	client, ok := this.GetRpcClient()[serverID]
	if ok == false {
		id_err := errors.New(fmt.Sprintf("serverID is invaild ,req id[%d]", serverID))
		std.GetLog().Error(id_err.Error())
		return false
	}
	rpcClient := msg_server.NewRpcServiceClient(client.GetConnect())
	this.addRpcService(serverID, rpcClient)
	return true
}

type LoginRpc struct {
}

func (this *LoginRpc) RegisterService(ctx context.Context, req *msg_server.RegisterServiceRequest) (*msg_server.RegisterServiceResponse, error) {

	return &msg_server.RegisterServiceResponse{}, nil
}

func (this *LoginRpc) QueryServiceList(ctx context.Context, req *google_protobuf.Empty) (*msg_server.QueryServiceListResponse, error) {

	return &msg_server.QueryServiceListResponse{}, nil
}
func (this *LoginRpc) UpdateServiceList(ctx context.Context, req *msg_server.UpdateServiceListRequest) (*google_protobuf.Empty, error) {

	for _, item := range req.GetServerList() {
		_, exist := std.GetRpcClient()[item.ServerId]
		if exist {
			continue
		}
		std.AddRpcClient(item.ServerId, item.Ip, item.Port, item.ServerType)
		std.SetRpcClient(item.ServerId)
	}
	return new(google_protobuf.Empty), nil
}
