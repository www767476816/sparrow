/*
 * @Author: your name
 * @Date: 2021-12-27 10:52:36
 * @LastEditTime: 2021-12-28 17:16:08
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: \sparrow\common\sql_service\sql_service.go
 */
package sql_service

import (
	"database/sql"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DataBaseInfo struct {
	*sql.DB
}

func Start(dbInfo *DataBaseInfo, dbName string, userName string, password string, dbAddr string, dbPort string, charset string) bool {
	connectInfo := userName + ":" + password + "@tcp(" + dbAddr + ":" + dbPort + ")/" + dbName + "?charset=" + charset + "&parseTime=true"
	var err error
	dbInfo.DB, err = sql.Open("mysql", connectInfo)
	if err != nil {
		panic(err)
	}
	//设置连接活跃时间单位秒
	dbInfo.DB.SetConnMaxLifetime(30)
	//设置最大连接数
	dbInfo.DB.SetMaxOpenConns(200)
	//设置空闲连接数
	dbInfo.DB.SetMaxIdleConns(100)
	//ping
	ping_err := dbInfo.DB.Ping()
	if ping_err != nil {
		panic(ping_err)
	}
	return true
}
