package base

import (
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc/peer"
	google_protobuf "google.golang.org/protobuf/types/known/emptypb"
	"net"
	"sparrow/common"
	"sparrow/protocol/error_msg"
	"sparrow/protocol/rpc_protocol"
	"time"
)

func (this*Base) RegisterRpcService() {
	//注册服务器
	rpc_protocol.RegisterRpcServiceServer(this.Frame.GetRpcServer().GetConnect(), new(CenterRpc))
}
func (this*Base) SetRpcClient(serverID uint32) bool{
	client,ok :=this.GetRpcClient()[serverID]
	if ok==false{
		id_err:= errors.New(fmt.Sprintf("serverID is invaild ,req id[%d]",serverID))
		std.GetLog().Error(id_err.Error())
		return false
	}
	rpcClient:= rpc_protocol.NewRpcServiceClient(client.GetConnect())
	this.addRpcService(serverID,rpcClient)
	return true
}
type CenterRpc struct{

}

func (this*CenterRpc)RegisterService(ctx context.Context,req *rpc_protocol.RegisterServiceRequest) (*rpc_protocol.RegisterServiceResponse, error)  {
	fmt.Println("收到注册消息,类型:",req.ServerType,",时间:",time.Now())
	//加入clien列表
	p, _ := peer.FromContext(ctx)
	ip,_,err:=net.SplitHostPort(p.Addr.String())
	if err!=nil{
		std.GetLog().Error("split addr err,addr："+p.Addr.String()+",err:"+err.Error())
		return new(rpc_protocol.RegisterServiceResponse),err
	}
	//连接这个client
	if req.GetServerType()<=common.SERVER_TYPE_MIN||req.GetServerType()>=common.SERVER_TYPE_MAX {
		type_err:= errors.New(fmt.Sprintf("server type is invaild ,req type[%d]",req.GetServerType()))
		std.GetLog().Error(type_err.Error())
		return new(rpc_protocol.RegisterServiceResponse),type_err
	}
	serverID:=std.nextServerID
	std.nextServerID+=1
	std.AddRpcClient(serverID,ip,req.GetRpcPort(),req.GetServerType())
	std.SetRpcClient(serverID)
	//通知其他client
	serviceList:=new(rpc_protocol.UpdateServiceListRequest)
	for key,value :=range std.GetRpcClient(){
		if key==serverID{
			continue
		}
		serviceList.ServerList=append(serviceList.ServerList, &rpc_protocol.ServiceItem{key,value.GetServerType(),value.GetIP(),value.GetPort()})
	}
	for key,_ :=range std.GetRpcClient(){
		std.getRpcService(key).UpdateServiceList(context.Background(),serviceList)
	}
	//构造返回
	res:= new(rpc_protocol.RegisterServiceResponse)
	res.ServerId=serverID
	res.ErrorCode=error_msg.EnumErrorCode_SUCCESS

	fmt.Println("返回注册信息,id:",serverID,",时间:",time.Now())
	return res,nil
}

func (this*CenterRpc)QueryServiceList(ctx context.Context,req *google_protobuf.Empty) (*rpc_protocol.QueryServiceListResponse, error)  {
	serviceList:=new(rpc_protocol.QueryServiceListResponse)
	for key,value :=range std.GetRpcClient(){
		serviceList.ServerList=append(serviceList.ServerList, &rpc_protocol.ServiceItem{key,value.GetServerType(),value.GetIP(),value.GetPort()})
	}
	return serviceList,nil
}
func (this*CenterRpc)UpdateServiceList(ctx context.Context,req *rpc_protocol.UpdateServiceListRequest) (*google_protobuf.Empty, error)  {

	return new(google_protobuf.Empty),nil
}