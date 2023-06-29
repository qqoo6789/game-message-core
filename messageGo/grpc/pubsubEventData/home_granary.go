package pubsubEventData

import (
	base_data "game-message-core/grpc/baseData"
)

// 存放物品到家园仓库
type GranaryStockpileEvent struct {
	MsgVersion   int64                        `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	HomeOwner    int64                        `json:"homeOwner"`
	Items        []base_data.GrpcItemBaseInfo `json:"items"`
	OccupantId   int64                        `json:"occupantId"`   // 使用者Id(owner | 进入家园的其他人员)
	OccupantName string                       `json:"occupantName"` // 使用者Name
}

func (p *GranaryStockpileEvent) Clear() {
	p.MsgVersion = 0
	p.HomeOwner = 0
	p.OccupantId = 0
	p.OccupantName = ""
	if p.Items != nil {
		p.Items = p.Items[:0]
	}
}
