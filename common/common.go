package common

import "encoding/binary"

const (
	SERVER_TYPE_MIN = 0
	CENTER_SERVER = 1			//中心服
	DISPATCH_SERVER = 2			//分发服
	LOGIN_SERVER = 3			//登录服
	LOGIC_SERVER = 4			//逻辑服
	ITEM_SERVER = 5				//道具服
	CHAT_SERVER = 6				//聊天服
	SERVER_TYPE_MAX = 7
)
const CENTER_SERVER_ID = 1
const INVALID_SERVER_ID = 0
const NET_PACKAGE_SIZE = 65535
func ConvertUserID2DBIndex(userID uint64) int32{
	return int32(userID%16)
}
func ConvertUserID2DTablendex(userID uint64) int32{
	return int32(userID%16)
}

func CheckPackageHead(data []byte,len int) bool{
	if len>NET_PACKAGE_SIZE || len<2 {
		return false
	}
	dataLen:=binary.LittleEndian.Uint16(data[0:2])
	if dataLen>NET_PACKAGE_SIZE{
		return false
	}
	return true
}