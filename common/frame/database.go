package frame

import (
	"sparrow/common/sql_service"
)

func (this* Frame) InitDatabase() {
	this.dataBase=make([]*sql_service.DataBaseInfo,0)
}

func (this* Frame) StartDatabase() bool {
	//启动数据库
	for _,config :=range this.normalConfig.DateBase{
		db:=new(sql_service.DataBaseInfo)
		if !sql_service.Start(db,config.Name,config.User,config.Password,config.Ip,config.Port,config.Charset){
			this.log.Panicln("database start error,",config)
			return false
		}
		this.dataBase=append(this.dataBase, db)
	}
	return true
}
func (this* Frame) CloseDatabase(){
	for _,db:= range this.dataBase{
		db.Close()
	}
	this.dataBase=this.dataBase[0:0]
}
