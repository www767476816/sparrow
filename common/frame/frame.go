package frame

import (
	"rpc-demo/common/log"
	"rpc-demo/common/rpc_service"
	"rpc-demo/common/sql_service"
)

type Frame struct {
	normalConfig* NormalConfig
	log* log.LogInfo
	dbList []*sql_service.DataBaseInfo
	rpcClientMap map[uint32]*rpc_service.RpcClient
	rpcServerMap map[uint32]*rpc_service.RpcServer
}

func (this* Frame) Init(configFile string) bool {
	//读取配置
	if !this.LoadConfig(configFile) {
		return false
	}
	//数据库
	this.InitDatabase()
	//先把rpc的数据贴上去

	return true
}

func (this* Frame) Start() bool {
	//启动数据库
	if !this.StartDatabase() {
		return false
	}
	return true
}

func (this* Frame) Run() error {

	return nil
}

func (this* Frame) Stop() error {

	return nil
}

func (this* Frame) Close() error {

	return nil
}
