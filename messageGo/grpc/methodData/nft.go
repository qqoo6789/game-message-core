package methodData

import (
	base_data "game-message-core/grpc/baseData"
)

type GrpcNftBaseData struct {
	NftId   string `json:"nftId"`
	ItemCid int32  `json:"itemCid"`
	Num     int32  `json:"num"`
}

func (p *GrpcNftBaseData) Clear() {
	p.NftId = ""
	p.ItemCid = 0
	p.Num = 0
}

// call mainService User use nft
type MainServiceActionUseNftInput struct {
	UserId int64           `json:"userId"`
	Nft    GrpcNftBaseData `json:"nft"`
}

func (p *MainServiceActionUseNftInput) Clear() {
	p.UserId = 0
	p.Nft.Clear()
}

type MainServiceActionUseNftOutput struct {
	Success   bool   `json:"success"`
	FailedMsg string `json:"failedMsg"`
}

func (p *MainServiceActionUseNftOutput) Clear() {
	p.Success = false
	p.FailedMsg = ""
}

// 通知mainService 扣除User nft input
type MainServiceActionTakeNftInput struct {
	UserId   int64             `json:"userId"`
	TakeNfts []GrpcNftBaseData `json:"takeNfts"`
}

func (p *MainServiceActionTakeNftInput) Clear() {
	p.UserId = 0
	if p.TakeNfts != nil {
		p.TakeNfts = p.TakeNfts[:0]
	}
}

type MainServiceActionTakeNftOutput struct {
	Success   bool   `json:"success"`
	FailedMsg string `json:"failedMsg"`
}

func (p *MainServiceActionTakeNftOutput) Clear() {
	p.Success = false
	p.FailedMsg = ""
}

// 通知mainService make User nft input
type MainServiceActionMintNftInput struct {
	UserId int64                      `json:"userId"`
	Item   base_data.GrpcItemBaseInfo `json:"item"`
}

func (p *MainServiceActionMintNftInput) Clear() {
	p.UserId = 0
	p.Item.Clear()
}

type MainServiceActionMintNftOutput struct {
	Success   bool   `json:"success"`
	FailedMsg string `json:"failedMsg"`
}

func (p *MainServiceActionMintNftOutput) Clear() {
	p.Success = false
	p.FailedMsg = ""
}

// 通知mainService 批量 make User nft input
type MainServiceActionMultiMintNftInput struct {
	UserId int64                        `json:"userId"`
	Items  []base_data.GrpcItemBaseInfo `json:"items"`
}

func (p *MainServiceActionMultiMintNftInput) Clear() {
	p.UserId = 0
	if p.Items != nil {
		p.Items = p.Items[:0]
	}
}

type MainServiceActionMultiMintNftOutput struct {
	Success   bool   `json:"success"`
	FailedMsg string `json:"failedMsg"`
}

func (p *MainServiceActionMultiMintNftOutput) Clear() {
	p.Success = false
	p.FailedMsg = ""
}
