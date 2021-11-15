package frame

import (
	"rpc-demo/common/sql_service"
)

func (this* Frame) InitDatabase() {
	this.dbList=make([]*sql_service.DataBaseInfo, len(this.normalConfig.DateBase))
}

func (this* Frame) StartDatabase() bool {
	//启动数据库
	for index,dbConfig := range this.normalConfig.DateBase{
		if !sql_service.Start(this.dbList[index],dbConfig.Name,dbConfig.User,dbConfig.Password,dbConfig.Ip,dbConfig.Port,dbConfig.Charset){
			return false
		}
	}
	return true
}
func (this* Frame) CloseDatabase(){
	for _,dbInfo:= range this.dbList{
		dbInfo.Close()
	}
	this.dbList=this.dbList[0:0]
}
