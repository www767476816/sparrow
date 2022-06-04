package client

import (
	"sparrow/common"
	"sparrow/dispatch_server/base"
	"sparrow/dispatch_server/connect"
	"sparrow/protocol/msg_code"
	"sparrow/protocol/msg_login_server"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
)

type Client struct {
	conn           *connect.Connect
	ticker         *time.Ticker
	startHeartBeat bool
	heartBeatNum   int
	locker         sync.Mutex
	close		   bool
}

func CreateClient(c *connect.Connect) *Client {
	result := new(Client)
	result.conn = c
	result.ticker = nil
	result.startHeartBeat = false
	result.heartBeatNum = 0
	result.close=false
	return result
}
func (this *Client) Close(){
	this.StopHeartBeat()
}
func (this *Client) Dispatch(msgCode msg_code.EnumMsgCode, data []byte) bool {
	if msgCode <= msg_code.EnumMsgCode_MIN || msgCode >= msg_code.EnumMsgCode_MAX {
		base.GetLog().Error("msg code is error,code:", msgCode, ",conn:", this.conn)
		return false
	}
	if this.close {
		base.GetLog().Error("client is close,code:", msgCode, ",conn:", this.conn)
		this.Close()
		return false
	}
	serverType := msgCode / 10000
	result := false
	switch serverType {
	case common.CENTER_SERVER:
		{
			result = this.toCenterServer(msgCode, data)
			break
		}
	case common.DISPATCH_SERVER:
		{
			result = this.toDispatchServer(msgCode, data)
			break
		}
	case common.LOGIN_SERVER:
		{
			result = this.toLoginServer(msgCode, data)
			break
		}
	case common.LOGIC_SERVER:
		{
			result = this.toLogicServer(msgCode, data)
			break
		}
	case common.ITEM_SERVER:
		{
			result = this.toItemServer(msgCode, data)
			break
		}
	case common.CHAT_SERVER:
		{
			result = this.toChatServer(msgCode, data)
			break
		}
	default:
		{
			base.GetLog().Error("msg code type is error,code:", msgCode, ",type:", serverType, ",conn:", this.conn)
			break
		}
	}
	if !result {
		base.GetLog().Error("msg error,code:", msgCode, ",type:", serverType, ",conn:", this.conn)
	}
	return result
}
func (this *Client) StartHeartBeat() {
	this.ticker := time.NewTicker(time.Duration(base.GetSpecialConfig().HeartBeatInterval) * time.Second) // 运行时长
	this.startHeartBeat=true
	this.heartBeatNum=0
	 go func(for range ticker.C {
		if this.heartBeatNum<=0 {
			base.GetLog().Error("client heart beat overtime,client:",this.conn)
			this.close=true
			return
		}
		this.locker.Lock()
		this.heartBeatNum=0
		this.locker.Unlock()
	})
}
func (this *Client) StopHeartBeat() {
	if this.startHeartBeat{
		this.ticker.Stop()
		this.locker.Lock()
		this.heartBeatNum=0
		this.locker.Unlock()
		this.startHeartBeat=false
	}
}

//中心服
func (this *Client) toCenterServer(msgCode msg_code.EnumMsgCode, data []byte) bool {
	result := false
	return result
}
//分发服
func (this *Client) toDispatchServer(msgCode msg_code.EnumMsgCode, data []byte) bool {
	result := false
	switch msgCode {
	case msg_code.EnumMsgCode_CS_HEART_BEAT_REQ:
		{
			this.locker.Lock()
			this.heartBeatNum=this.heartBeatNum+1
			this.locker.Unlock()
			result = this.conn.SendMsg(msg_code.EnumMsgCode_SC_HEART_BEAT_RES, proto.Message{})
			break
		}
	}
	return result
}

//登录服
func (this *Client) toLoginServer(msgCode msg_code.EnumMsgCode, data []byte) bool {
	conn := base.GetRpcService(common.CENTER_SERVER_ID)
	if conn == nil {
		return false
	}
	result := false

	switch msgCode {
	case msg_code.EnumMsgCode_CS_CLIENT_LOGIN_REQ:
		{
			req := new(msg_login_server.ClientLoginReq)
			if err := proto.Unmarshal(data, req); err != nil {
				base.GetLog().Error("proto Unmarshal error,code:", msg_code.EnumMsgCode_CS_CLIENT_LOGIN_REQ, ",data:", data)
				break
			}
			res, err := conn.ClientLogin(context.Background(), req)
			if err != nil {
				base.GetLog().Error("ClientLogin error,code:", msg_code.EnumMsgCode_CS_CLIENT_LOGIN_REQ, ",req:", req, ",res:", res)
				break
			}
			this.startHeartBeat()
			result = this.conn.SendMsg(msg_code.EnumMsgCode_SC_CLIENT_LOGIN_RES, res)
			break
		}
	case msg_code.EnumMsgCode_CS_CLIENT_REGISTER_REQ:
		{
			req := new(msg_login_server.ClientRegisterReq)
			if err := proto.Unmarshal(data, req); err != nil {
				base.GetLog().Error("proto Unmarshal error,code:", msg_code.EnumMsgCode_CS_CLIENT_REGISTER_REQ, ",data:", data)
				break
			}
			res, err := conn.ClientRegister(context.Background(), req)
			if err != nil {
				base.GetLog().Error("ClientLogin error,code:", msg_code.EnumMsgCode_CS_CLIENT_REGISTER_REQ, ",req:", req, ",res:", res)
				break
			}
			result = this.conn.SendMsg(msg_code.EnumMsgCode_SC_CLIENT_REGISTER_RES, res)
			break
		}
	case msg_code.EnumMsgCode_CS_QUERY_ROLE_LIST_REQ:
		{
			req := new(msg_login_server.QueryRoleListReq)
			if err := proto.Unmarshal(data, req); err != nil {
				base.GetLog().Error("proto Unmarshal error,code:", msg_code.EnumMsgCode_CS_QUERY_ROLE_LIST_REQ, ",data:", data)
				break
			}
			res, err := conn.QueryRoleList(context.Background(), req)
			if err != nil {
				base.GetLog().Error("ClientLogin error,code:", msg_code.EnumMsgCode_CS_QUERY_ROLE_LIST_REQ, ",req:", req, ",res:", res)
				break
			}
			result = this.conn.SendMsg(msg_code.EnumMsgCode_SC_QUERY_ROLE_LIST_RES, res)
			break
		}
	}
	return result
}

//逻辑服
func (this *Client) toLogicServer(msgCode msg_code.EnumMsgCode, data []byte) bool {
	result := false
	return result
}

//道具服
func (this *Client) toItemServer(msgCode msg_code.EnumMsgCode, data []byte) bool {
	result := false
	return result
}

//聊天服
func (this *Client) toChatServer(msgCode msg_code.EnumMsgCode, data []byte) bool {
	result := false
	return result
}
