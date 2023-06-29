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

func (p *UserChangeServiceEvent) Clear() {
	p.MsgVersion = 0
	p.UserId = 0
	p.FormService.Clear()
	p.ToService.Clear()
	p.UserAgentAppId = ""
	p.UserSocketId = ""
	p.UserPosition.Clear()
	p.UserDir.Clear()
}
