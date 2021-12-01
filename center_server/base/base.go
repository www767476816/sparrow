package base

import (
	"sparrow/common"
	"sparrow/common/frame"
	"sparrow/protocol/rpc_protocol"
)

type Base struct {
	*frame.Frame
	nextServerID uint32
}

func (this* Base) Init()  {
	this.Frame=new(frame.Frame)
	this.Frame.Init("login_config.xml")
	this.Frame.SetServerID(common.CENTER_SERVER_ID)
	this.nextServerID=common.CENTER_SERVER_ID+1
}
func (this* Base) Start() bool {
	if this.Frame.Start()==false{
		return false
	}
	this.RegisterRpcService()
	return true
}

func (this* Base) Run() {
	this.Frame.Run()
}

func (this* Base) Close() {
	this.Frame.Close()
}

func (this* Base) addRpcService(serverID uint32,cli rpc_protocol.RpcServiceClient) {
	this.Frame.RpcService[serverID]=cli
}
func (this* Base) getRpcService(serverID uint32) rpc_protocol.RpcServiceClient {
	return this.Frame.RpcService[serverID]
}

var(std = new(Base))

func Init(){
	std.Init()
}
func Start() bool{
	return std.Start()
}
func Run(){
	std.Run()
}
func Close(){
	std.Close()
}

