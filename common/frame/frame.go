package frame

import (
	"github.com/go-redis/redis"
	"sparrow/common/log"
	"sparrow/common/rpc_service"
	"sparrow/common/sql_service"
	"sparrow/protocol/rpc_protocol"
)

type Frame struct {
	serverID uint32
	normalConfig* NormalConfig
	log* log.LogInfo
	dataBase []*sql_service.DataBaseInfo
	redisList []*redis.Client
	rpcServer *rpc_service.RpcServer
	rpcClientMap map[uint32]*rpc_service.RpcClient
	RpcService map[uint32]rpc_protocol.RpcServiceClient
}

func (this* Frame) Init(configFile string) bool {

	//读取配置
	if !this.LoadConfig(configFile) {
		return false
	}
	//日志
	this.StartLog()
	//数据库
	this.InitDatabase()
	//redis
	this.InitRedis()
	this.InitRpc()
	return true
}

func (this* Frame) Start() bool {
	//启动数据库
	if !this.StartDatabase() {
		return false
	}
	//启动redis
	this.StartRedis()
	//启动rpc
	this.startRpcServer(this.normalConfig.RpcPort,this.normalConfig.ServerType)
	return true
}

func (this* Frame) Run() {
	//rpc运行
	go this.rpcServer.Run()

}

func (this* Frame) Close() {
	this.CloseRpc()
	this.CloseRedis()
	this.CloseDatabase()
	this.CloseConfig()
	this.CloseLog()

}

func (this* Frame) GetServerID() uint32{
	return this.serverID
}
func (this* Frame) SetServerID(id uint32) {
	this.serverID=id
}
