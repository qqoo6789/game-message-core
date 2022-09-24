package pubsubEventData

import (
	base_data "game-message-core/grpc/baseData"
	"game-message-core/proto"
)

type UserEnterGameEvent struct {
	MsgVersion        int64                        `json:"msgVersion"`        // 消息版本号 值为毫秒时间戳
	SceneServiceAppId string                       `json:"sceneServiceAppId"` //
	MapId             int32                        `json:"mapId"`
	BaseData          base_data.GrpcPlayerBaseData `json:"baseData"`
	Position          proto.Vector3                `json:"position"`
	Dir               proto.Vector3                `json:"dir"`
}
