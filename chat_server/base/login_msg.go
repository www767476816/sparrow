/*
 * @Author: your name
 * @Date: 2021-12-27 17:17:22
 * @LastEditTime: 2021-12-28 15:23:21
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: \sparrow\login_server\base\login_msg.go
 */
package base

import (
	"sparrow/protocol/msg_client"

	"golang.org/x/net/context"
)

func (this *ChatRpc) ClientRegister(ctx context.Context, req *msg_client.ClientRegisterReq) (*msg_client.ClientRegisterRes, error) {

	return &msg_client.ClientRegisterRes{}, nil
}

func (this *ChatRpc) ClientLogin(ctx context.Context, req *msg_client.ClientLoginReq) (*msg_client.ClientLoginRes, error) {

	return &msg_client.ClientLoginRes{}, nil
}
func (this *ChatRpc) QueryRoleList(ctx context.Context, req *msg_client.QueryRoleListReq) (*msg_client.QueryRoleListRes, error) {

	return new(msg_client.QueryRoleListRes), nil
}
func (this *ChatRpc) CreateRole(ctx context.Context, req *msg_client.CreateRoleReq) (*msg_client.CreateRoleRes, error) {

	return new(msg_client.CreateRoleRes), nil
}
