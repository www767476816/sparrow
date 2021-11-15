package frame

import (
	"rpc-demo/common/log"
)

func (this* Frame) StartLog() bool {
	//日志
	this.log=new(log.LogInfo)
	this.log.Init(log.ConvertToLogLevel(this.selfConfig.LogConfig.LogLevel))
	return true
}

func (this* Frame) CloseLog(){
	this.log=nil
}
