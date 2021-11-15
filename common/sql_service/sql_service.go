package sql_service

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
)
type DataBaseInfo struct {
	*gorm.DB
}

func Start(dbInfo *DataBaseInfo,dbName string,userName string,password string,dbAddr string,dbPort string,charset string) bool {
	connectInfo :=userName+":"+password+"@tcp("+dbAddr+":"+dbPort+")/"+dbName+"?charset="+charset+"&parseTime=true"
	var err error
	dbInfo.DB,err=gorm.Open("mysql",connectInfo)
	if err!=nil{
		logrus.Panic(err)
	}
	//设置全局表名禁用复数
	dbInfo.SingularTable(true)
	//设置连接活跃时间单位秒
	dbInfo.DB.DB().SetConnMaxLifetime(30)
	//设置最大连接数
	dbInfo.DB.DB().SetMaxOpenConns(200)
	//设置空闲连接数
	dbInfo.DB.DB().SetMaxIdleConns(100)
	return true
}