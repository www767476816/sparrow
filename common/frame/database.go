package frame

import (
	"rpc-demo/common/sql_service"
)

func (this* Frame) InitDatabase() {
	this.dataBase=new(sql_service.DataBaseInfo)
}

func (this* Frame) StartDatabase() bool {
	//启动数据库
	dbConfig := this.normalConfig.DateBase
	if !sql_service.Start(this.dataBase,dbConfig.Name,dbConfig.User,dbConfig.Password,dbConfig.Ip,dbConfig.Port,dbConfig.Charset){
		return false
	}
	return true
}
func (this* Frame) CloseDatabase(){
	this.dataBase.Close()
	this.dataBase=nil
}
