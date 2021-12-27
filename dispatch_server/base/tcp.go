package base

import (
	"encoding/xml"
	"net"
	"sparrow/common/config"
	"sparrow/dispatch_server/connect"
)

type SpecialConfig struct {
	XMLName   xml.Name   `xml:"root"`
	ListenPort  string      `xml:"special>listen_port"`
	MaxCount int  `xml:"special>max_count"`
	HeartBeatInterval  int `xml:"special>heart_beat_interval"`
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
		if len(this.clientConn)>= this.specialConfig.MaxCount{
			this.GetLog().Warning("conn count is max")
			conn.Close()
			continue
		}

		this.Lock()
		newConnect :=connect.CreateConnect(this.currentConnID,conn,ip,port)
		this.clientConn[this.currentConnID]=newConnect
		this.currentConnID=this.currentConnID+1
		this.Unlock()

		go newConnect.ReceiveMsg()
	}
}
func (this* Base) CloseTcpByID(id uint32)  {
	conn,exist:=this.clientConn[id]
	if !exist{
		return
	}
	this.Lock()
	conn.Close()
	delete(this.clientConn, conn.GetID())
	this.Unlock()
}

func (this* Base) CloseTcp()  {
	for _,item :=range this.clientConn{
		item.Close()
	}
	this.clientConn=make(map[uint32]*connect.Connect)

	this.tcpListener.Close()
	this.currentConnID=1
}


