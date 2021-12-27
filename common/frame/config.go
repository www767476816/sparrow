package frame

import (
	"encoding/xml"
	"sparrow/common/config"
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
type redisConfig struct {
	XMLName  xml.Name `xml:"redis"`
	Ip string `xml:"ip"`
	Port string `xml:"port"`
	Password string `xml:"password"`
}
type centerServer struct {
	XMLName  xml.Name `xml:"center_server"`
	Ip string `xml:"ip"`
	Port string `xml:"port"`
}

type NormalConfig struct {
	XMLName   xml.Name   `xml:"root"`
	ServerType  uint32      `xml:"normal>server_type"`
	LogConfig logConfig  `xml:"normal>log"`
	DateBase  []dateBase `xml:"normal>date_base"`
	Redis redisConfig `xml:"normal>redis"`
	CenterServer centerServer `xml:"normal>center_server"`
	RpcPort   string    `xml:"normal>rpc_port"`
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
func (this* Frame) GetConfig() *NormalConfig{
	return this.normalConfig
}
