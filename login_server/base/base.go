package base

import (
	"golang.org/x/net/context"
	google_protobuf "google.golang.org/protobuf/types/known/emptypb"
	"sparrow/common"
	"sparrow/common/frame"
	"sparrow/protocol/error_msg"
	"sparrow/protocol/rpc_protocol"
)

type Base struct {
	*frame.Frame
}

func (this* Base) Init()  {
	this.Frame=new(frame.Frame)
	this.Frame.Init("login_config.xml")

}
func (this* Base) Start() bool {
	if this.Frame.Start()==false{
		return false
	}
	this.RegisterRpcService()
	//向中心服注册
	this.registerToCenter()
	//请求服务列表
	res,queryError:=this.getRpcService(common.CENTER_SERVER_ID).QueryServiceList(context.Background(),new(google_protobuf.Empty))
	if queryError!=nil{
		this.GetLog().Panic(queryError)
	}
	if res.ErrorCode!=error_msg.EnumErrorCode_SUCCESS{
		this.GetLog().Panic("register fail error code:",res.ErrorCode)
	}
	for _,item := range res.GetServerList(){
		std.AddRpcClient(item.ServerId,item.Ip,item.Port,item.ServerType)
		std.SetRpcClient(item.ServerId)
	}
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
func (this* Base) registerToCenter()  {
	this.AddRpcClient(common.CENTER_SERVER_ID,this.GetConfig().CenterServer.Ip,this.GetConfig().CenterServer.Port,common.CENTER_SERVER)
	this.SetRpcClient(common.CENTER_SERVER_ID)

	_,exist:=this.GetRpcClient()[common.CENTER_SERVER_ID]
	if !exist{
		this.GetLog().Panic("connect to center server error:The center server does not exist in the list!!!")
	}
	req := new(rpc_protocol.RegisterServiceRequest)
	req.ServerType=common.LOGIN_SERVER
	req.RpcPort=this.GetConfig().RpcPort
	res,err:=this.getRpcService(common.CENTER_SERVER_ID).RegisterService(context.Background(),req)
	if err!=nil{
		this.GetLog().Panic(err)
	}
	if res.ErrorCode!=error_msg.EnumErrorCode_SUCCESS{
		this.GetLog().Panic("register fail error code:",res.ErrorCode)
	}
	this.SetServerID(res.ServerId)
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

