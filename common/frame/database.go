/*
 * @Author: your name
 * @Date: 2021-12-27 10:52:36
 * @LastEditTime: 2021-12-28 18:01:56
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: \sparrow\common\frame\database.go
 */
package frame

import (
	"sparrow/common/sql_service"
)

func (this *Frame) InitDatabase() {
	this.dataBase = make([]*sql_service.DataBaseInfo, 0)
}

func (this *Frame) StartDatabase() bool {
	//启动数据库
	for _, config := range this.normalConfig.DateBase {
		db := new(sql_service.DataBaseInfo)
		if !sql_service.Start(db, config.Name, config.User, config.Password, config.Ip, config.Port, config.Charset) {
			this.log.Panicln("database start error,", config)
			return false
		}

		this.dataBase = append(this.dataBase, db)
	}
	return true
}
func (this *Frame) GetDatabase(index int) *sql_service.DataBaseInfo {
	if index > len(this.dataBase)-1 {
		return nil
	}
	return this.dataBase[index]
}
func (this *Frame) CloseDatabase() {
	for _, db := range this.dataBase {
		db.Close()
	}
	this.dataBase = this.dataBase[0:0]
}
