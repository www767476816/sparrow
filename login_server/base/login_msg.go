/*
 * @Author: your name
 * @Date: 2021-12-27 17:17:22
 * @LastEditTime: 2021-12-29 20:37:50
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: \sparrow\login_server\base\login_msg.go
 */
package base

import (
	"sparrow/common"
	"sparrow/logic_server/base"
	"sparrow/protocol/error_msg"
	"sparrow/protocol/msg_common"
	"sparrow/protocol/msg_login_server"
	"time"

	"golang.org/x/net/context"
)

func (this *LoginRpc) ClientRegister(ctx context.Context, req *msg_login_server.ClientRegisterReq) (*msg_login_server.ClientRegisterRes, error) {
	res := new(msg_login_server.ClientRegisterRes)
	res.ErrorCode = error_msg.EnumErrorCode_SUCCESS

	if len(req.Accounts) > common.ACCOUNTS_LEN {
		res.ErrorCode = error_msg.EnumErrorCode_REGISTER_ACCOUNTS_FORMAT_ERROR
		return res, nil
	}
	if len(req.Password) > common.PASSWORD_LEN {
		res.ErrorCode = error_msg.EnumErrorCode_REGISTER_PASSWORD_FORMAT_ERROR
		return res, nil
	}
	db := base.GetDatabase()
	if db == nil {
		base.GetLog().Errorln("database not find")
		res.ErrorCode = error_msg.EnumErrorCode_FAILD
		return res, nil
	}
	result, err := db.DB.Query("SELECT uid FROM accounts_info WHERE account=", req.Accounts, " AND password=", req.Password)
	if err != nil {
		base.GetLog().Errorln(err)
		res.ErrorCode = error_msg.EnumErrorCode_FAILD
		return res, nil
	}
	defer result.Close()
	for result.Next() {
		read_err := result.Scan(&res.Uid)
		if read_err != nil {
			base.GetLog().Errorln(err)
			res.ErrorCode = error_msg.EnumErrorCode_FAILD
			return res, nil
		}
	}
	if res.Uid != 0 {
		res.ErrorCode = error_msg.EnumErrorCode_REGISTER_ACCOUNTS_EXISTS
		return res, nil
	}
	//注册
	global_db := base.GetGlobalDatabase()
	if global_db == nil {
		base.GetLog().Errorln("database not find")
		res.ErrorCode = error_msg.EnumErrorCode_FAILD
		return res, nil
	}

	trans, _ := global_db.DB.Begin()
	row := global_db.DB.QueryRow("SELECT id FROM id_pool WHERE invalid=0")
	row.Scan(&res.Uid)
	_, update_err := global_db.DB.Exec("UPDATE id_pool SET invalid=1,dispatch_time=", time.Now(), " WHERE id=", res.Uid)
	if update_err != nil {
		trans.Rollback()
		base.GetLog().Errorln("update id_pool error,", update_err)
		res.ErrorCode = error_msg.EnumErrorCode_FAILD
		return res, nil
	}
	trans.Commit()

	_, insert_err := db.DB.Exec("INSERT INTO accounts_info(uid,accounts,password,channel,register_time) VALUES(",
		res.Uid, ",", req.Accounts, ",", req.Password, ",", req.Channel, ",", time.Now())
	if insert_err != nil {
		base.GetLog().Errorln("INSERT accounts_info error,", insert_err)
		res.ErrorCode = error_msg.EnumErrorCode_FAILD
		return res, nil
	}

	return res, nil
}

func (this *LoginRpc) ClientLogin(ctx context.Context, req *msg_login_server.ClientLoginReq) (*msg_login_server.ClientLoginRes, error) {
	res := new(msg_login_server.ClientLoginRes)
	res.ErrorCode = error_msg.EnumErrorCode_SUCCESS

	db := base.GetDatabase()
	if db == nil {
		base.GetLog().Errorln("database not find")
		res.ErrorCode = error_msg.EnumErrorCode_FAILD
		return res, nil
	}
	result, err := db.DB.Query("SELECT uid FROM accounts_info WHERE account=", req.Accounts, " AND password=", req.Password)
	if err != nil {
		base.GetLog().Errorln(err)
		res.ErrorCode = error_msg.EnumErrorCode_FAILD
		return res, nil
	}
	defer result.Close()
	for result.Next() {
		read_err := result.Scan(&res.Uid)
		if read_err != nil {
			base.GetLog().Errorln(err)
			res.ErrorCode = error_msg.EnumErrorCode_FAILD
			return res, nil
		}
	}
	if res.Uid == 0 {
		res.ErrorCode = error_msg.EnumErrorCode_LOGIN_ACCOUNTS_NOT_EXISTS
		return res, nil
	}
	_, update_err := db.DB.Exec("UPDATE accounts_info SET last_login_time=", time.Now(), " WHERE id=", res.Uid)
	if update_err != nil {
		base.GetLog().Errorln("update id_pool error,", update_err)
		res.ErrorCode = error_msg.EnumErrorCode_FAILD
		return res, nil
	}

	return res, nil
}
func (this *LoginRpc) QueryRoleList(ctx context.Context, req *msg_login_server.QueryRoleListReq) (*msg_login_server.QueryRoleListRes, error) {
	res := new(msg_login_server.QueryRoleListRes)
	res.ErrorCode = error_msg.EnumErrorCode_SUCCESS

	db := base.GetDatabase()
	if db == nil {
		base.GetLog().Errorln("database not find")
		res.ErrorCode = error_msg.EnumErrorCode_FAILD
		return res, nil
	}
	result, err := db.DB.Query("SELECT role_id,role_name,head FROM role_info WHERE uid=", req.Uid)
	if err != nil {
		base.GetLog().Errorln(err)
		res.ErrorCode = error_msg.EnumErrorCode_FAILD
		return res, nil
	}
	defer result.Close()
	for result.Next() {
		role := new(msg_common.Role)
		read_err := result.Scan(&role.RoleId, &role.RoleName, &role.Head)
		if read_err != nil {
			base.GetLog().Errorln(err)
			res.ErrorCode = error_msg.EnumErrorCode_FAILD
			return res, nil
		}
		res.RoleList = append(res.RoleList, role)
	}
	db.DB.Exec("UPDATE role_info SET last_login_time=", time.Now(), " WHERE uid=", req.Uid)
	return res, nil
}
func (this *LoginRpc) CreateRole(ctx context.Context, req *msg_login_server.CreateRoleReq) (*msg_login_server.CreateRoleRes, error) {
	res := new(msg_login_server.CreateRoleRes)
	res.ErrorCode = error_msg.EnumErrorCode_SUCCESS

	db := base.GetDatabase()
	if db == nil {
		base.GetLog().Errorln("database not find")
		res.ErrorCode = error_msg.EnumErrorCode_FAILD
		return res, nil
	}
	//校验用户
	account_result, err := db.DB.Exec("SELECT uid FROM accounts_info WHERE uid=", req.Uid)
	if err != nil {
		base.GetLog().Errorln(err)
		res.ErrorCode = error_msg.EnumErrorCode_FAILD
		return res, nil
	}
	if count, _ := account_result.RowsAffected(); count <= 0 {
		base.GetLog().Errorln(err)
		res.ErrorCode = error_msg.EnumErrorCode_ROLE_LIST_ACCOUNT_NOT_EXISTS
		return res, nil
	}
	//校验角色
	if len(req.Name) > common.NICK_NAME_LEN {
		base.GetLog().Errorln("nickname too len,", req.Name)
		res.ErrorCode = error_msg.EnumErrorCode_ROLE_NAME_FORMAT_ERROR
		return res, nil
	}
	role_result, err := db.DB.Exec("SELECT uid FROM role_info WHERE role_name=", req.Name)
	if err != nil {
		base.GetLog().Errorln(err)
		res.ErrorCode = error_msg.EnumErrorCode_FAILD
		return res, nil
	}
	if count, _ := role_result.RowsAffected(); count <= 0 {
		base.GetLog().Errorln(err)
		res.ErrorCode = error_msg.EnumErrorCode_ROLE_NAME_EXISTS
		return res, nil
	}
	//创建
	_, insert_err := db.DB.Exec("INSERT INTO role_info(uid,role_id,role_name,head,create_time,last_login_time) VALUES(",
		req.Uid, ",", req.Uid, ",", req.Name, ",", req.Head, ",", time.Now(), ",", time.Now())
	if insert_err != nil {
		base.GetLog().Errorln("INSERT role_info error,", insert_err)
		res.ErrorCode = error_msg.EnumErrorCode_FAILD
		return res, nil
	}
	res.RoleInfo = new(msg_common.Role)
	res.RoleInfo.Head = req.Head
	res.RoleInfo.RoleId = req.Uid
	res.RoleInfo.RoleName = req.Name
	return res, nil
}
