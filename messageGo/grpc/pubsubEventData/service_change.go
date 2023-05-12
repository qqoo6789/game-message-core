package pubsubEventData

import (
	base_data "game-message-core/grpc/baseData"
)

// 玩家完成服务切换事件
type UserChangeServiceEvent struct {
	MsgVersion     int64                 `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	UserId         int64                 `json:"userId"`
	FormService    base_data.ServiceData `json:"formService"`
	ToService      base_data.ServiceData `json:"toService"`
	UserAgentAppId string                `json:"userAgentAppId"`
	UserSocketId   string                `json:"userSocketId"`
	UserPosition   base_data.GrpcVector3 `json:"userPosition"`
	UserDir        base_data.GrpcVector3 `json:"userDir"`
}
