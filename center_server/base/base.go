package base

import (
	"fmt"
	"sparrow/common"
	"sparrow/common/frame"
	"sparrow/protocol/msg_server"
)

type Base struct {
	*frame.Frame
	nextServerID uint32
}

func (this* Base) Init()  {
	fmt.Println("初始化开始")
	this.Frame=new(frame.Frame)

	this.Frame.Init("center_config.xml")
	this.Frame.SetServerID(common.CENTER_SERVER_ID)
	this.nextServerID=common.CENTER_SERVER_ID+1
	fmt.Println("初始化完成")
}
func (this* Base) Start() bool {
	fmt.Println("启动开始")
	if this.Frame.Start()==false{
		return false
	}
	this.RegisterRpcService()
	fmt.Println("启动完成")
	return true
}

func (this* Base) Run() {
	fmt.Println("运行开始")
	this.Frame.Run()
	fmt.Println("运行结束")
}

func (this* Base) Close() {
	fmt.Println("关闭")
	this.Frame.Close()
}

func (this* Base) addRpcService(serverID uint32,cli msg_server.RpcServiceClient) {
	this.Frame.RpcService[serverID]=cli
}
func (this* Base) getRpcService(serverID uint32) msg_server.RpcServiceClient {
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

