package frame

import (
	"encoding/xml"
	"rpc-demo/common/config"
)

type logConfig struct {
	XMLName  xml.Name `xml:"log"`
	LogLevel string `xml:"level"`
}
type dateBase struct {
	XMLName  xml.Name `xml:"date_base"`
	Name string `xml:"name"`
	Ip string `xml:"ip"`
	Port string `xml:"port"`
	User string `xml:"user"`
	Password string `xml:"password"`
	Charset string `xml:"charset"`
}
type centerServer struct {
	XMLName  xml.Name `xml:"center_server"`
	Ip string `xml:"ip"`
	Port string `xml:"port"`
}
type netWork struct {
	XMLName  xml.Name `xml:"net_work"`
	ListenPort string `xml:"listen_port"`
}

type NormalConfig struct {
	XMLName   xml.Name   `xml:"root"`
	ServerID  uint32      `xml:"normal>server_id"`
	ServerType  int32      `xml:"normal>server_type"`
	LogConfig logConfig  `xml:"normal>log"`
	DateBase  []dateBase `xml:"normal>date_base"`
	CenterServer centerServer `xml:"normal>center_server"`
	NetWork   netWork    `xml:"normal>net_work"`
}

func (this* Frame) LoadConfig(configFile string) bool {
	//读取配置
	this.normalConfig=new(NormalConfig)
	if !config.ReadConfig(configFile,this.normalConfig) {
		return false
	}
	return true
}
func (this* Frame) CloseConfig(){
	this.normalConfig=nil
}

