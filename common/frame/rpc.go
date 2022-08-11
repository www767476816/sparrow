package frame

import (
	"sparrow/common/rpc_service"
	"sparrow/protocol/msg_server"
)

func (this* Frame) InitRpc(){
	this.rpcServer=new(rpc_service.RpcServer)
	this.rpcClientMap=make(map[uint32]*rpc_service.RpcClient)
	this.RpcService=make(map[uint32]msg_server.RpcServiceClient)
}
func (this* Frame) GetRpcServer() *rpc_service.RpcServer{
	return this.rpcServer
}
func (this* Frame) GetRpcClient() map[uint32]*rpc_service.RpcClient{
	return this.rpcClientMap
}

func (this* Frame) startRpcServer(port string,serverType uint32) {
	server :=rpc_service.CreateServer("tcp",port,serverType)
	server.Init()
	if err:=server.Start();err!=nil{
		this.log.Panicln(err)
	}

	this.rpcServer=server
}
//添加服务
func (this* Frame)AddRpcClient(serverID uint32,ip string,port string,serverType uint32) {
	//如果之前有相同的连接就把之前的断掉
	var oldID uint32
	oldID=0
	for k,v := range this.rpcClientMap{
		if v.GetServerType()==serverType&&v.GetIP()==ip&&v.GetPort()==port{
			oldID=k
			v.Close()
			break
		}
	}
	if oldID!=0{
		delete(this.rpcClientMap, oldID)
	}
	//创建新连接
	client :=rpc_service.CreateClient(ip,port,serverType)
	if err:=client.Connect();err!=nil {
		this.log.Panicln(err)
	}

	this.rpcClientMap[serverID]=client
}
func (this* Frame) CloseRpc(){
	for _,value:=range this.rpcClientMap{
		value.Close()
	}
	this.rpcClientMap=make(map[uint32]*rpc_service.RpcClient)

	this.rpcServer.Close()
	this.rpcServer=nil
}

