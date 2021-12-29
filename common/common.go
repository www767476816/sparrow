/*
 * @Author: your name
 * @Date: 2021-12-27 10:52:36
 * @LastEditTime: 2021-12-29 20:26:07
 * @LastEditors: Please set LastEditors
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: \sparrow\common\common.go
 */
package common

import (
	"bytes"
	"encoding/binary"
)

const (
	SERVER_TYPE_MIN = 0
	CENTER_SERVER   = 1 //中心服
	DISPATCH_SERVER = 2 //分发服
	LOGIN_SERVER    = 3 //登录服
	LOGIC_SERVER    = 4 //逻辑服
	ITEM_SERVER     = 5 //道具服
	CHAT_SERVER     = 6 //聊天服
	SERVER_TYPE_MAX = 7
)
const CENTER_SERVER_ID = 1
const INVALID_SERVER_ID = 0
const NET_PACKAGE_SIZE = 65535
const PACKAGE_HEAD_LEN = 4
const MSG_CODE_LEN = 4
const ACCOUNTS_LEN = 16
const PASSWORD_LEN = 9
const NICK_NAME_LEN = 6

func ConvertUserID2DBIndex(userID uint64) int32 {
	return int32(userID % 16)
}
func ConvertUserID2DTablendex(userID uint64) int32 {
	return int32(userID % 16)
}

func CheckPackageHead(data []byte, len int) bool {
	if len < PACKAGE_HEAD_LEN {
		return false
	}
	dataLen := binary.LittleEndian.Uint16(data[0:PACKAGE_HEAD_LEN])
	if dataLen > (NET_PACKAGE_SIZE - PACKAGE_HEAD_LEN - MSG_CODE_LEN) {
		return false
	}
	return true
}

//整形转换成字节
func IntToBytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

//字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}
