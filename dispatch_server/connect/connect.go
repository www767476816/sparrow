package connect

import (
	"fmt"
	"net"
	"sparrow/common"
	"sparrow/dispatch_server/base"
	"sparrow/dispatch_server/client"
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
func (this* Connect) GetID() uint32{
	return this.id
}
func (this* Connect) ReceiveMsg() {
	data :=make([]byte,common.NET_PACKAGE_SIZE)
	for{
		len,readErr:=this.Conn.Read(data)
		if readErr!=nil{
			base.GetLog().Error(readErr)
			break
		}
		fmt.Println("connID:",this.id,",接收到数据，长度:",len,"数据：",string(data[0:len]))
		if len>common.NET_PACKAGE_SIZE{
			data=data[0:common.NET_PACKAGE_SIZE]
			len=common.NET_PACKAGE_SIZE
		}
		if !common.CheckPackageHead(data,len){
			base.GetLog().Error("CheckPackageHead err,connect:",this)
			continue
		}
		client:=client.CreateClient(this)
		client.Dispatch(data)
	}

	fmt.Println(this.Conn.RemoteAddr(),",断开连接")
	base.CloseTcpByID(this.GetID())

}

func (this* Connect) SendMsg(msgCode int,data []byte) bool {
	len:=len(data)
	if len>(common.NET_PACKAGE_SIZE-common.PACKAGE_HEAD_LEN-common.MSG_CODE_LEN){
		base.GetLog().Error("data len is error,len:",len,",conn:",this)
		return false
	}
	headSlice:=common.IntToBytes(len)
	codeSlice:=common.IntToBytes(msgCode)
	headSlice=append(headSlice,codeSlice...)
	headSlice=append(headSlice,data...)
	sendLen,sendErr:=this.Write(headSlice)
	if sendErr!=nil{
		base.GetLog().Error("send data error,conn:",this,",err:",sendErr)
		return false
	}
	if sendLen!= (len+common.PACKAGE_HEAD_LEN)||sendErr!=nil{
		base.GetLog().Error("send data error,alllen:",len,"sendLen:",sendLen,",conn:",this)
		return false
	}
	return true
}

func getDataLen(data *[]byte) int{
	head:=(*data)[0:common.PACKAGE_HEAD_LEN]
	*data=(*data)[common.PACKAGE_HEAD_LEN:len(*data)]
	return common.BytesToInt(head)
}
func getMsgCode(data *[]byte) int{
	head:=(*data)[0:common.MSG_CODE_LEN]
	*data=(*data)[common.MSG_CODE_LEN:len(*data)]
	return common.BytesToInt(head)
}