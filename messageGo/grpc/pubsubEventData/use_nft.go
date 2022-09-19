package pubsubEventData

import "game-message-core/proto"

type UserUseNFTEvent struct {
	MsgVersion     int64                    `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	UserId         int64                    `json:"userId"`
	NftId          string                   `json:"nftId"`
	Cid            int32                    `json:"cid"`
	NftType        proto.NFTType            `json:"nftType"`
	Num            int32                    `json:"num"`
	ConsumableData *proto.NFTConsumableInfo `json:"consumableData"`
	X              int32                    `json:"x"`
	Z              int32                    `json:"z"`
}
