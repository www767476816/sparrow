package client

import (
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"sparrow/common"
	"sparrow/dispatch_server/base"
	"sparrow/dispatch_server/connect"
	"sparrow/protocol/msg_code"
	"sparrow/protocol/msg_login_server"
	"time"
)

type Client struct {
	conn *connect.Connect
}

func CreateClient(c *connect.Connect) *Client{
	result:=new(Client)
	result.conn=c
	return result
}

func (this* Client) Dispatch(msgCode msg_code.EnumMsgCode,data []byte)  {
	if msgCode<=msg_code.EnumMsgCode_MIN||msgCode>=msg_code.EnumMsgCode_MAX{
		base.GetLog().Error("msg code is error,code:",msgCode,",conn:",this.conn)
		return
	}
	serverType:=msgCode/10000
	switch serverType {
	case common.CENTER_SERVER:{
		this.toCenterServer(msgCode,data)
		break
	}
	case common.LOGIN_SERVER:{
		this.toLoginServer(msgCode,data)
		break
	}
	case common.LOGIC_SERVER:{
		this.toLogicServer(msgCode,data)
		break
	}
	case common.ITEM_SERVER:{
		this.toItemServer(msgCode,data)
		break
	}
	case common.CHAT_SERVER:{
		this.toChatServer(msgCode,data)
		break
	}
	default:{
		base.GetLog().Error("msg code type is error,code:",msgCode,",type:",serverType,",conn:",this.conn)
		break
	}
	}
}
func  (this* Client) StartHeartBeat() {
	ticker := time.NewTicker(time.Duration(base.GetSpecialConfig().HeartBeatInterval) * time.Second) // 运行时长

	for {
		select {
		case <-ticker.C:

		}}
	ticker.Stop()
}
//中心服
func (this* Client) toCenterServer(msgCode msg_code.EnumMsgCode,data []byte)  {

}
//登录服
func (this* Client) toLoginServer(msgCode msg_code.EnumMsgCode,data []byte)  {
	conn:=base.GetRpcService(common.CENTER_SERVER_ID)
	if conn==nil{
		return
	}
	switch msgCode {
	case msg_code.EnumMsgCode_CS_CLIENT_LOGIN_REQ:{
		req:=new(msg_login_server.ClientLoginReq)
		if err:=proto.Unmarshal(data,req);err!=nil{
			base.GetLog().Error("proto Unmarshal error,code:",msg_code.EnumMsgCode_CS_CLIENT_LOGIN_REQ,",data:",data)
			break
		}
		res,err:=conn.ClientLogin(context.Background(),req)
		if err!=nil{
			base.GetLog().Error("ClientLogin error,code:",msg_code.EnumMsgCode_CS_CLIENT_LOGIN_REQ,",req:",req,",res:",res)
			break
		}
		this.conn.SendMsg(msg_code.EnumMsgCode_SC_CLIENT_LOGIN_RES,res)
		break
	}
	case msg_code.EnumMsgCode_CS_CLIENT_REGISTER_REQ:{
		req:=new(msg_login_server.ClientRegisterReq)
		if err:=proto.Unmarshal(data,req);err!=nil{
			base.GetLog().Error("proto Unmarshal error,code:",msg_code.EnumMsgCode_CS_CLIENT_REGISTER_REQ,",data:",data)
			break
		}
		res,err:=conn.ClientRegister(context.Background(),req)
		if err!=nil{
			base.GetLog().Error("ClientLogin error,code:",msg_code.EnumMsgCode_CS_CLIENT_REGISTER_REQ,",req:",req,",res:",res)
			break
		}
		this.conn.SendMsg(msg_code.EnumMsgCode_SC_CLIENT_REGISTER_RES,res)
		break
	}
	case msg_code.EnumMsgCode_CS_QUERY_ROLE_LIST_REQ:{
		req:=new(msg_login_server.QueryRoleListReq)
		if err:=proto.Unmarshal(data,req);err!=nil{
			base.GetLog().Error("proto Unmarshal error,code:",msg_code.EnumMsgCode_CS_QUERY_ROLE_LIST_REQ,",data:",data)
			break
		}
		res,err:=conn.QueryRoleList(context.Background(),req)
		if err!=nil{
			base.GetLog().Error("ClientLogin error,code:",msg_code.EnumMsgCode_CS_QUERY_ROLE_LIST_REQ,",req:",req,",res:",res)
			break
		}
		this.conn.SendMsg(msg_code.EnumMsgCode_SC_QUERY_ROLE_LIST_RES,res)
		break
	}
	}
}
//逻辑服
func (this* Client) toLogicServer(msgCode msg_code.EnumMsgCode,data []byte)  {

}
//道具服
func (this* Client) toItemServer(msgCode msg_code.EnumMsgCode,data []byte)  {

}
//聊天服
func (this* Client) toChatServer(msgCode msg_code.EnumMsgCode,data []byte)  {

}