package pubsubEventData

import (
	base_data "game-message-core/grpc/baseData"
)

type UserGetNFTEvent struct {
	MsgVersion int64                        `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	UserId     int64                        `json:"userId"`
	Items      []base_data.GrpcItemBaseInfo `json:"items"`
}

func (p *UserGetNFTEvent) Clear() {
	p.MsgVersion = 0
	p.UserId = 0
	if p.Items != nil {
		p.Items = p.Items[:0]
	}
}
