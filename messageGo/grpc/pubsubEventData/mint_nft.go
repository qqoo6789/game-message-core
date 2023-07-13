package pubsubEventData

import (
	base_data "game-message-core/grpc/baseData"
	"game-message-core/proto"
)

type MintNftInfo struct {
	Owner          int64                        `json:"owner"`
	Items          []base_data.GrpcItemBaseInfo `json:"items"`
	FromEntityType proto.EntityType             `json:"fromEntityType"`
	FromEntityCid  int32                        `json:"fromEntityCid"`
	FromEntityName string                       `json:"fromEntityName"`
}

func (p *MintNftInfo) Clear() {
	p.Owner = 0
	p.FromEntityType = 0
	p.FromEntityCid = 0
	p.FromEntityName = ""
	if p.Items != nil {
		p.Items = p.Items[:0]
	}
}

// 通知mainService make User nft input
type MintNftEvent struct {
	MsgVersion  int64                 `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	FormService base_data.ServiceData `json:"formService"`
	Nft         MintNftInfo           `json:"nft"`
}

func (p *MintNftEvent) Clear() {
	p.MsgVersion = 0
	p.FormService.Clear()
	p.Nft.Clear()
}

// 通知mainService 批量 make User nft input
type MultiMintNftEvent struct {
	MsgVersion  int64                 `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	FormService base_data.ServiceData `json:"formService"`
	Nfts        []MintNftInfo         `json:"nfts"`
}

func (p *MultiMintNftEvent) Clear() {
	p.FormService.Clear()
	for i, _ := range p.Nfts {
		p.Nfts[i].Clear()
	}
	if p.Nfts != nil {
		p.Nfts = p.Nfts[:0]
	}
}
