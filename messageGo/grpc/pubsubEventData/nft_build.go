package pubsubEventData

import (
	base_data "game-message-core/grpc/baseData"
)

type NftBuildAddEvent struct {
	MsgVersion int64                  `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	Build      base_data.GrpcNftBuild `json:"build"`
}

func (p *NftBuildAddEvent) Clear() {
	p.MsgVersion = 0
	p.Build.Clear()
}

type NftBuildUpdateEvent struct {
	MsgVersion int64                  `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	Build      base_data.GrpcNftBuild `json:"build"`
}

func (p *NftBuildUpdateEvent) Clear() {
	p.MsgVersion = 0
	p.Build.Clear()
}

type NftBuildRemoveEvent struct {
	MsgVersion int64  `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	EntityId   int64  `json:"entityId"`
	FromNft    string `json:"fromNft"`
	Owner      int64  `json:"owner"`
}

func (p *NftBuildRemoveEvent) Clear() {
	p.MsgVersion = 0
	p.EntityId = 0
	p.FromNft = ""
	p.Owner = 0
}
