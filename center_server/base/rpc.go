/*
 * @Author: your name
 * @Date: 2021-12-27 10:52:36
 * @LastEditTime: 2021-12-27 17:21:18
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: \sparrow\center_server\base\rpc.go
 */
package base

import (
	"errors"
	"fmt"
	"net"
	"sparrow/common"
	"sparrow/protocol/error_code"
	"sparrow/protocol/msg_server"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc/peer"
	google_protobuf "google.golang.org/protobuf/types/known/emptypb"
)

func (this *Base) RegisterRpcService() {
	//注册服务器
	msg_server.RegisterRpcServiceServer(this.Frame.GetRpcServer().GetConnect(), new(CenterRpc))
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

type CenterRpc struct {
}

func (this *CenterRpc) RegisterService(ctx context.Context, req *msg_server.RegisterServiceRequest) (*msg_server.RegisterServiceResponse, error) {
	fmt.Println("收到注册消息,类型:", req.ServerType, ",时间:", time.Now())
	//加入clien列表
	p, _ := peer.FromContext(ctx)
	ip, _, err := net.SplitHostPort(p.Addr.String())
	if err != nil {
		std.GetLog().Error("split addr err,addr：" + p.Addr.String() + ",err:" + err.Error())
		return new(msg_server.RegisterServiceResponse), err
	}
	//连接这个client
	if req.GetServerType() <= common.SERVER_TYPE_MIN || req.GetServerType() >= common.SERVER_TYPE_MAX {
		type_err := errors.New(fmt.Sprintf("server type is invaild ,req type[%d]", req.GetServerType()))
		std.GetLog().Error(type_err.Error())
		return new(msg_server.RegisterServiceResponse), type_err
	}
	serverID := std.nextServerID
	std.nextServerID += 1
	std.AddRpcClient(serverID, ip, req.GetRpcPort(), req.GetServerType())
	std.SetRpcClient(serverID)
	//通知其他client
	serviceList := new(msg_server.UpdateServiceListRequest)
	for key, value := range std.GetRpcClient() {
		if key == serverID {
			continue
		}
		serviceList.ServerList = append(serviceList.ServerList, &msg_server.ServiceItem{key, value.GetServerType(), value.GetIP(), value.GetPort()})
	}
	for key, _ := range std.GetRpcClient() {
		std.getRpcService(key).UpdateServiceList(context.Background(), serviceList)
	}
	//构造返回
	res := new(msg_server.RegisterServiceResponse)
	res.ServerId = serverID
	res.ErrorCode = error_code.EnumErrorCode_SUCCESS

	fmt.Println("返回注册信息,id:", serverID, ",时间:", time.Now())
	return res, nil
}

func (this *CenterRpc) QueryServiceList(ctx context.Context, req *google_protobuf.Empty) (*msg_server.QueryServiceListResponse, error) {
	serviceList := new(msg_server.QueryServiceListResponse)
	for key, value := range std.GetRpcClient() {
		serviceList.ServerList = append(serviceList.ServerList, &msg_server.ServiceItem{key, value.GetServerType(), value.GetIP(), value.GetPort()})
	}
	return serviceList, nil
}
func (this *CenterRpc) UpdateServiceList(ctx context.Context, req *msg_server.UpdateServiceListRequest) (*google_protobuf.Empty, error) {

	return new(google_protobuf.Empty), nil
}
