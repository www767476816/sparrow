package frame

import (
	"golang.org/x/net/context"
	"rpc-demo/common"
	"rpc-demo/common/frame/rpc_protocol"
	"rpc-demo/common/rpc_service"
	"strconv"
)

func (this* Frame) StartRpc() bool {
	//开始监听rpc端口
	this.startRpcServer(this.normalConfig.RpcPort,this.normalConfig.ServerType)
	go this.rpcServer.Run()

	if this.normalConfig.ServerType!=common.CENTER_SERVER{
		//向中心服注册
		this.startRpcClient(this.normalConfig.CenterServer.Ip,this.normalConfig.CenterServer.Port,this.normalConfig.ServerType)
		//向中心服请求服务器列表
		this.QueryServerList()
	}
	return true
}

func (this* Frame) startRpcServer(port string,serverType uint32) {
	server :=rpc_service.CreateServer("tcp",port,serverType)
	server.Init()
	if err:=server.Start();err!=nil{
		this.log.Panicln(err)
	}
	rpc_protocol.RegisterRpcProtocolServer(server.GetConnect(),new(rpc_protocol.RpcProtocolService))
	this.rpcServer=server
}

func (this* Frame) startRpcClient(ip string,port string,serverType uint32) bool {
	client :=rpc_service.CreateClient(ip,port,serverType)
	if err:=client.Connect();err!=nil {
		this.log.Panicln(err)
		return false
	}
	client.SetServiceClient(rpc_protocol.NewRpcProtocolClient(client.GetConnect()))
	//向中心服注册
	registerResult, registerErr := client.GetServiceClient().RegisterService(context.Background(),&rpc_protocol.RegisterServiceRequest{ServerType: this.normalConfig.ServerType})
	if registerErr !=nil{
		this.log.Panicln(registerErr)
	}
	if registerResult.ErrorCode!=0 {
		this.log.Panicln("register_result err,err_code:"+strconv.Itoa(int(registerResult.ErrorCode)))
	}
	this.serverID= registerResult.ServerId
	this.rpcClientMap[common.CENTER_SERVER_ID]=client
	return true
}
//添加服务
func (this* Frame)addRpcService(serverID uint32,ip string,port string,serverType uint32) {
	client :=rpc_service.CreateClient(ip,port,serverType)
	if err:=client.Connect();err!=nil {
		this.log.Panicln(err)
	}
	client.SetServiceClient(rpc_protocol.NewRpcProtocolClient(client.GetConnect()))
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

func (this* Frame) QueryServerList() bool{
	query_result,query_err:=this.rpcClientMap[common.CENTER_SERVER_ID].GetServiceClient().QueryServiceList(context.Background(),nil)
	if query_err!=nil{
		this.log.Panicln(query_err)
		return false
	}
	if(query_result.ErrorCode!=0){
		this.log.Panicln("register_result err,err_code:"+strconv.Itoa(int(query_result.ErrorCode)))
		return false
	}
	for _,item :=range query_result.ServerList{
		this.addRpcService(item.ServerId,item.Ip,item.Port,item.ServerType)
	}
	return true
}

