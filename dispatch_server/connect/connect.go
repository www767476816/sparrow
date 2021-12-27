package connect

import (
	"fmt"
	"net"
	"sparrow/common"
	"sparrow/dispatch_server/base"
	"sparrow/dispatch_server/client"
	"sparrow/protocol/msg_code"

	"github.com/golang/protobuf/proto"
)

type Connect struct {
	conn net.Conn
	id   uint32
	ip   string
	port string
}

func CreateConnect(id uint32, connect net.Conn, ip string, port string) *Connect {
	fmt.Println("新连接：id:", id, ",ip:", ip, ",port", port)
	newConnect := new(Connect)
	newConnect.id = id
	newConnect.conn = connect
	newConnect.ip = ip
	newConnect.port = port
	return newConnect
}
func (this *Connect) GetID() uint32 {
	return this.id
}
func (this *Connect) ReceiveMsg() {
	data := make([]byte, common.NET_PACKAGE_SIZE)
	client := client.CreateClient(this)
	for {
		len, readErr := this.conn.Read(data)
		if readErr != nil {
			base.GetLog().Error(readErr)
			break
		}
		fmt.Println("connID:", this.id, ",接收到数据，长度:", len, "数据：", string(data[0:len]))
		if len > common.NET_PACKAGE_SIZE {
			data = data[0:common.NET_PACKAGE_SIZE]
			len = common.NET_PACKAGE_SIZE
		}
		if !common.CheckPackageHead(data, len) {
			base.GetLog().Error("CheckPackageHead err,connect:", this)
			continue
		}
		realLen := getDataLen(&data)
		if realLen != len {
			base.GetLog().Error("read len not equal recv len,readLen:", realLen, ",recv len:", len)
			continue
		}
		msgCode := getMsgCode(&data)
		client.Dispatch(msg_code.EnumMsgCode(msgCode), data)
	}

	fmt.Println(this.conn.RemoteAddr(), ",断开连接")
	base.CloseTcpByID(this.GetID())

}

func (this *Connect) SendMsg(msgCode msg_code.EnumMsgCode, msg proto.Message) bool {
	data, mar_err := proto.Marshal(msg)
	if mar_err != nil {
		base.GetLog().Error("proto marshal error,code:", msgCode, ",msg:", msg)
		return false
	}
	len := len(data)
	if len > (common.NET_PACKAGE_SIZE - common.PACKAGE_HEAD_LEN - common.MSG_CODE_LEN) {
		base.GetLog().Error("data len is error,len:", len, ",conn:", this)
		return false
	}
	headSlice := common.IntToBytes(len)
	codeSlice := common.IntToBytes(int(msgCode))
	headSlice = append(headSlice, codeSlice...)
	headSlice = append(headSlice, data...)
	sendLen, sendErr := this.conn.Write(headSlice)
	if sendErr != nil {
		base.GetLog().Error("send data error,conn:", this, ",err:", sendErr)
		return false
	}
	if sendLen != (len+common.PACKAGE_HEAD_LEN) || sendErr != nil {
		base.GetLog().Error("send data error,alllen:", len, "sendLen:", sendLen, ",conn:", this)
		return false
	}
	return true
}
func (this *Connect) Close() {
	this.conn.Close()
}

//返回客户端发送的字节长度，注意：会把消息头删掉！！！
func getDataLen(data *[]byte) int {
	head := (*data)[0:common.PACKAGE_HEAD_LEN]
	*data = (*data)[common.PACKAGE_HEAD_LEN:len(*data)]
	return common.BytesToInt(head) + common.PACKAGE_HEAD_LEN
}

//返回消息标识，注意：会把消息标识删掉！！！
func getMsgCode(data *[]byte) int {
	head := (*data)[0:common.MSG_CODE_LEN]
	*data = (*data)[common.MSG_CODE_LEN:len(*data)]
	return common.BytesToInt(head)
}
