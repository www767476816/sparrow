/*
 * @Author: your name
 * @Date: 2021-12-27 17:17:22
 * @LastEditTime: 2021-12-27 17:23:26
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: \sparrow\login_server\base\login_msg.go
 */
package base

import (
	"sparrow/protocol/msg_login_server"

	"golang.org/x/net/context"
)

func (this *ItemRpc) ClientRegister(ctx context.Context, req *msg_login_server.ClientRegisterReq) (*msg_login_server.ClientRegisterRes, error) {

	return &msg_login_server.ClientRegisterRes{}, nil
}

func (this *ItemRpc) ClientLogin(ctx context.Context, req *msg_login_server.ClientLoginReq) (*msg_login_server.ClientLoginRes, error) {

	return &msg_login_server.ClientLoginRes{}, nil
}
func (this *ItemRpc) QueryRoleList(ctx context.Context, req *msg_login_server.QueryRoleListReq) (*msg_login_server.QueryRoleListRes, error) {

	return new(msg_login_server.QueryRoleListRes), nil
}
