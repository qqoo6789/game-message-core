package pubsubEventData

import (
	base_data "game-message-core/grpc/baseData"
)

type UpdateTalentEvent struct {
	MsgVersion int64                            `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	UserId     int64                            `json:"userId"`
	Levels     []base_data.GrpcTalentLevel      `json:"levels"`
	Trees      []base_data.GrpcTalentTreeUpdate `json:"trees"`
}

type AddTalentExpEvent struct {
	MsgVersion        int64                     `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	SceneServiceAppId string                    `json:"sceneServiceAppId"`
	MapId             int32                     `json:"mapId"`
	UserId            int64                     `json:"userId"`
	AddExps           []base_data.GrpcTalentExp `json:"addExps"`
}
