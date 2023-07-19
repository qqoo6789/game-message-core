package pubsubEventData

import base_data "game-message-core/grpc/baseData"

// 采集物品
type CollectEvent struct {
	MsgVersion   int64                      `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	FormService  base_data.ServiceData      `json:"formService"`
	UserId       int64                      `json:"userId"`
	TargetEntity base_data.EntityInfo       `json:"targetEntity"`
	State        string                     `json:"state"`
	DropItem     base_data.GrpcItemBaseInfo `json:"dropItem"`
}

func (p *CollectEvent) Clear() {
	p.MsgVersion = 0
	p.UserId = 0
	p.UserId = 0
	p.TargetEntity.Clear()
	p.State = ""
	p.DropItem.Clear()
}
