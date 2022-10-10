package methodData

type TakeNftData struct {
	NftId   string `json:"nftId"`
	ItemCid int32  `json:"itemCid"`
	Num     int32  `json:"num"`
}

// 通知mainService 扣除User nft input
type MainServiceActionTakeNftInput struct {
	MsgVersion int64         `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	UserId     int64         `json:"userId"`
	TakeNfts   []TakeNftData `json:"takeNfts"`
}
type MainServiceActionTakeNftOutput struct {
	Success   bool   `json:"success"`
	FailedMsg string `json:"failedMsg"`
}
