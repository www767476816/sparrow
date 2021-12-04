package base

import (
	"encoding/xml"
	"fmt"
	"net"
	"sparrow/common"
	"sparrow/common/config"
	"sparrow/dispatch_server/connect"
)

type SpecialConfig struct {
	XMLName   xml.Name   `xml:"root"`
	ListenPort  string      `xml:"special>listen_port"`
	MaxCount uint32  `xml:"special>max_count"`
	HeartBeatInterval  uint32 `xml:"special>heart_beat_interval"`
}
func (this* Base) InitTcp(configFile string)  {
	this.specialConfig=new(SpecialConfig)
	if !config.ReadConfig(configFile,this.specialConfig) {
		this.GetLog().Panic("read special config err")
	}
	this.currentConnID=1
}

func (this* Base) StartTcp()  {
	tcpListener,listenErr:=net.Listen("tcp",":"+this.specialConfig.ListenPort)
	if listenErr!=nil{
		this.GetLog().Panic(listenErr)
	}
	this.tcpListener=tcpListener
}
func (this* Base) newConnect(connID uint32,conn *connect.Connect){

	data :=make([]byte,common.NET_PACKAGE_SIZE)
	for{
		len,readErr:=conn.Read(data)
		if readErr!=nil{
			this.GetLog().Error(readErr)
			break
		}
		conn.ReceiveMsg(data,len)
	}

	fmt.Println(conn.RemoteAddr(),",断开连接")
	conn.Close()
	this.Lock()
	delete(this.clientConn, connID)
	this.Unlock()
}
func (this* Base) RunTcp()  {
	for{
		conn, accErr := this.tcpListener.Accept()
		if accErr != nil {
			this.GetLog().Error(accErr)
			continue
		}

		ip,port,parseErr:=net.SplitHostPort(conn.RemoteAddr().String())
		if parseErr != nil {
			this.GetLog().Error(parseErr)
			continue
		}
		this.Lock()
		newConnect :=connect.CreateConnect(this.currentConnID,conn,ip,port)
		this.clientConn[this.currentConnID]=newConnect
		this.currentConnID=this.currentConnID+1
		this.Unlock()

		go this.newConnect(conn)
	}
}

func (this* Base) CloseTcp()  {
	for _,item :=range this.clientConn{
		item.Close()
	}
	this.clientConn=make(map[uint32]*connect.Connect)

	this.tcpListener.Close()
	this.currentConnID=1
}


