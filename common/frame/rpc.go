package frame

import (
	"rpc-demo/common/rpc_service"
)


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

