package pubsubEventData

import (
	base_data "game-message-core/grpc/baseData"
)

type ServiceUnregisterEvent struct {
	MsgVersion int64                 `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	Service    base_data.ServiceData `json:"service"`    //
}

func (p *ServiceUnregisterEvent) Clear() {
	p.MsgVersion = 0
	p.Service.Clear()
}
