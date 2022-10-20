package pubsubEventData

import (
	base_data "game-message-core/grpc/baseData"
)

type NftBuildAddEvent struct {
	MsgVersion int64                  `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	Build      base_data.GrpcNftBuild `json:"build"`
}
type NftBuildUpdateEvent struct {
	MsgVersion int64                  `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	Build      base_data.GrpcNftBuild `json:"build"`
}

type NftBuildRemoveEvent struct {
	MsgVersion int64  `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	EntityId   int64  `json:"entityId"`
	FromNft    string `json:"fromNft"`
	Owner      int64  `json:"owner"`
}
