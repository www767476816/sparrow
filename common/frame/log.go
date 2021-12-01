package frame

import (
	"sparrow/common/log"
)

func (this* Frame) StartLog() bool {
	//日志
	this.log=new(log.LogInfo)
	this.log.Init(log.ConvertToLogLevel(this.normalConfig.LogConfig.LogLevel))
	return true
}

func (this* Frame) CloseLog(){
	this.log=nil
}
func (this* Frame) GetLog() * log.LogInfo{
	return this.log
}

