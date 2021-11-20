package common

const (
	CENTER_SERVER = 1			//中心服
	DISPATCH_SERVER = 2			//分发服
	LOGIN_SERVER = 3			//登录服
	LOGIC_SERVER = 4			//逻辑服
	ITEM_SERVER = 5				//道具服
	CHAT_SERVER = 6				//聊天服
)
const CENTER_SERVER_ID = 1
func ConvertUserID2DBIndex(userID uint64) int32{
	return int32(userID%16)
}
func ConvertUserID2DTablendex(userID uint64) int32{
	return int32(userID%16)
}