package base

import (
	"fmt"
	"sparrow/common"
	"sparrow/common/frame"
	"sparrow/common/log"
	"sparrow/common/sql_service"
	"sparrow/protocol/error_code"
	"sparrow/protocol/msg_server"
	"time"

	"golang.org/x/net/context"
	google_protobuf "google.golang.org/protobuf/types/known/emptypb"
)

type Base struct {
	*frame.Frame
}

func (this *Base) Init() {
	fmt.Println("初始化开始")
	this.Frame = new(frame.Frame)

	this.Frame.Init("logic_config.xml")
	fmt.Println("初始化完成")
}
func (this *Base) Start() bool {
	fmt.Println("启动开始")
	if this.Frame.Start() == false {
		return false
	}
	this.RegisterRpcService()
	return true
}

func (this *Base) Run() {
	fmt.Println("运行开始")
	this.Frame.Run()

	//向中心服注册
	this.registerToCenter()
	//请求服务列表
	this.QueryServiceList()

	fmt.Println("运行结束")
}

func (this *Base) Close() {
	fmt.Println("关闭")
	this.Frame.Close()
}

func (this *Base) addRpcService(serverID uint32, cli msg_server.RpcServiceClient) {
	this.Frame.RpcService[serverID] = cli
}

func (this *Base) getRpcService(serverID uint32) msg_server.RpcServiceClient {
	return this.Frame.RpcService[serverID]
}
func (this *Base) registerToCenter() {
	this.AddRpcClient(common.CENTER_SERVER_ID, this.GetConfig().CenterServer.Ip, this.GetConfig().CenterServer.Port, common.CENTER_SERVER)
	this.SetRpcClient(common.CENTER_SERVER_ID)

	_, exist := this.GetRpcClient()[common.CENTER_SERVER_ID]
	if !exist {
		this.GetLog().Panic("connect to center server error:The center server does not exist in the list!!!")
	}
	req := new(msg_server.RegisterServiceRequest)
	req.ServerType = common.LOGIC_SERVER
	req.RpcPort = this.GetConfig().RpcPort
	fmt.Println("向中心服注册：", time.Now())
	res, err := this.getRpcService(common.CENTER_SERVER_ID).RegisterService(context.Background(), req)
	if err != nil {
		this.GetLog().Panic(err)
	}
	fmt.Println("注册成功,id：", res.ServerId, ",时间：", time.Now())
	if res.ErrorCode != error_code.EnumErrorCode_SUCCESS {
		this.GetLog().Panic("register fail error code:", res.ErrorCode)
	}
	this.SetServerID(res.ServerId)
}
func (this *Base) QueryServiceList() {
	res, queryError := this.getRpcService(common.CENTER_SERVER_ID).QueryServiceList(context.Background(), new(google_protobuf.Empty))
	if queryError != nil {
		this.GetLog().Panic(queryError)
	}
	if res.ErrorCode != error_code.EnumErrorCode_SUCCESS {
		this.GetLog().Panic("register fail error code:", res.ErrorCode)
	}
	for _, item := range res.GetServerList() {
		if item.ServerId == this.Frame.GetServerID() {
			continue
		}
		std.AddRpcClient(item.ServerId, item.Ip, item.Port, item.ServerType)
		std.SetRpcClient(item.ServerId)
	}
}

var (
	std = new(Base)
)

func Init() {
	std.Init()
}
func Start() bool {
	return std.Start()
}
func Run() {
	std.Run()
}
func Close() {
	std.Close()
}
func GetDatabase() *sql_service.DataBaseInfo {
	return std.Frame.GetDatabase(1)
}
func GetGlobalDatabase() *sql_service.DataBaseInfo {
	return std.Frame.GetDatabase(0)
}
func GetLog() *log.LogInfo {
	return std.GetLog()
}

func GetRpcService(serverID uint32) msg_server.RpcServiceClient {
	return std.getRpcService(serverID)
}
