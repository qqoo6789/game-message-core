package methodData

import (
	base_data "game-message-core/grpc/baseData"
)

type GrpcNftBaseData struct {
	NftId   string `json:"nftId"`
	ItemCid int32  `json:"itemCid"`
	Num     int32  `json:"num"`
}

// call mainService User use nft
type MainServiceActionUseNftInput struct {
	UserId int64           `json:"userId"`
	Nft    GrpcNftBaseData `json:"nft"`
}
type MainServiceActionUseNftOutput struct {
	Success   bool   `json:"success"`
	FailedMsg string `json:"failedMsg"`
}

// 通知mainService 扣除User nft input
type MainServiceActionTakeNftInput struct {
	UserId   int64             `json:"userId"`
	TakeNfts []GrpcNftBaseData `json:"takeNfts"`
}
type MainServiceActionTakeNftOutput struct {
	Success   bool   `json:"success"`
	FailedMsg string `json:"failedMsg"`
}

// 通知mainService make User nft input
type MainServiceActionMintNftInput struct {
	UserId int64                      `json:"userId"`
	Item   base_data.GrpcItemBaseInfo `json:"item"`
}
type MainServiceActionMintNftOutput struct {
	Success   bool   `json:"success"`
	FailedMsg string `json:"failedMsg"`
}

// 通知mainService 批量 make User nft input
type MainServiceActionMultiMintNftInput struct {
	UserId int64                        `json:"userId"`
	Items  []base_data.GrpcItemBaseInfo `json:"items"`
}
type MainServiceActionMultiMintNftOutput struct {
	Success   bool   `json:"success"`
	FailedMsg string `json:"failedMsg"`
}
