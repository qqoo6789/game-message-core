package pubsubEventData

import (
	base_data "game-message-core/grpc/baseData"
)

type SavePlayerEventData struct {
	MsgVersion  int64                 `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	FormService base_data.ServiceData `json:"formService"`
	UserId      int64                 `json:"userId"`
	PosX        float32               `json:"posX"`
	PosY        float32               `json:"posY"`
	PosZ        float32               `json:"posZ"`
	CurHP       int32                 `json:"curHP"`
	DirX        float32               `json:"dirX"`
	DirY        float32               `json:"dirY"`
	DirZ        float32               `json:"dirZ"`
}

func (p *SavePlayerEventData) Clear() {
	p.MsgVersion = 0
	p.FormService.Clear()
	p.UserId = 0
	p.PosX = 0
	p.PosY = 0
	p.PosZ = 0
	p.CurHP = 0
	p.DirX = 0
	p.DirY = 0
	p.DirZ = 0
}
