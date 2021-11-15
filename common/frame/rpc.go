package frame

import (
	"errors"
	"rpc-demo/common"
	"rpc-demo/common/frame/service"
	"rpc-demo/common/rpc_service"
	"strconv"
)

func (this* Frame) StartRpc() bool {
	if this.normalConfig.ServerType!=common.CENTER_SERVER{
		//向中心服注册
		//向中心服请求服务器列表
	}

	//开始监听rpc端口
	//连接各个服务器

	return true
}

func (this* Frame) startRpcServer(serverID uint32,port string) bool {
	server :=rpc_service.CreateServer("tcp",port)
	server.Init()
	if err:=server.Start();err!=nil{
		this.log.Panicln(err)
		return false
	}
	this.rpcServerMap[serverID]=server
	return true
}

func (this* Frame) startRpcClient(serverID uint32,ip string,port string) bool {
	client :=rpc_service.CreateClient(ip,port)
	if err:=client.Connect();err!=nil{
		this.log.Panicln(err)
		return false
	}
	this.rpcClientMap[serverID]=client
	return true
}

func (this* Frame)registerRpcService(serviceType int) error{
	//中心服
	switch serviceType {
	case common.CENTER_SERVER:{
		service.RegisterCenterServerServiceServer(this.rpcServerMap[this.normalConfig.ServerID].GetConnect(),new(service.CenterServerService))
		break
	}
	case common.DISPATCH_SERVER:{
		for x,y :=range this.rpcServerMap{

		}
		break
	}
	default:{
		return errors.New("invalid service type:"+strconv.Itoa(serviceType))
	}


	}
	return nil;
}
func (this* Frame) CloseRpc(){
	for _,value:=range this.rpcClientMap{
		value.Close()
	}
	this.rpcClientMap=make(map[uint32]*rpc_service.RpcClient)

	for _,value:=range this.rpcServerMap{
		value.Close()
	}
	this.rpcServerMap=make(map[uint32]*rpc_service.RpcServer)
}

