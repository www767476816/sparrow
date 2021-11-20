package frame

import (
	"github.com/go-redis/redis"
	"rpc-demo/common/log"
	"rpc-demo/common/rpc_service"
	"rpc-demo/common/sql_service"
)

type Frame struct {
	serverID uint32
	normalConfig* NormalConfig
	log* log.LogInfo
	dbList []*sql_service.DataBaseInfo
	rpcServer *rpc_service.RpcServer
	redisList []*redis.Client
}

func (this* Frame) Init(configFile string) bool {
	//读取配置
	if !this.LoadConfig(configFile) {
		return false
	}
	//数据库
	this.InitDatabase()
	//redis
	this.InitRedis()

	return true
}

func (this* Frame) Start() bool {
	//启动数据库
	if !this.StartDatabase() {
		return false
	}
	//启动数据库
	this.StartRedis()
	return true
}

func (this* Frame) Run() error {

	return nil
}

func (this* Frame) Close() {
	this.CloseRedis()
	this.CloseDatabase()
	this.CloseConfig()
	this.CloseLog()
}
