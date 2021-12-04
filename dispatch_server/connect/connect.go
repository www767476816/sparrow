package connect

import (
	"fmt"
	"net"
	"sparrow/common"
)

type Connect struct {
	net.Conn
	id uint32
	ip string
	port string
}

func CreateConnect(id uint32,connect net.Conn,ip string,port string)*Connect{
	fmt.Println("新连接：id:",id,",ip:",ip,",port",port)
	newConnect:=new(Connect)
	newConnect.id=id
	newConnect.Conn=connect
	newConnect.ip=ip
	newConnect.port=port
	return newConnect
}

func (this* Connect) ReceiveMsg(data []byte,len int) {
	fmt.Println("connID:",this.id,",接收到数据，长度:",len,"数据：",string(data[0:len]))
	if !common.CheckPackageHead(data,len){
		base.Get
	}

}

func (this* Connect) SendMsg(data []byte) bool {

	return true
}
