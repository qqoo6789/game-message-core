package pubsubEventData

import (
	base_data "game-message-core/grpc/baseData"
	"game-message-core/proto"
)

type UserUseNFTEvent struct {
	MsgVersion     int64                            `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	UserId         int64                            `json:"userId"`
	NftId          string                           `json:"nftId"`
	Cid            int32                            `json:"cid"`
	NftType        proto.NftTraitType               `json:"nftType"`
	Num            int32                            `json:"num"`
	ConsumableData *base_data.GrpcNFTConsumableInfo `json:"consumableData"`
	X              float32                          `json:"x"`
	Y              float32                          `json:"y"`
	Z              float32                          `json:"z"`
}

func (p *UserUseNFTEvent) Clear() {
	p.MsgVersion = 0
	p.UserId = 0
	p.NftId = ""
	p.Cid = 0
	p.NftType = 0
	p.Num = 0
	p.ConsumableData = nil
	p.X = 0
	p.Y = 0
	p.Z = 0
}
