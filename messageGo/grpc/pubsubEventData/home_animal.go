package pubsubEventData

import (
	base_data "game-message-core/grpc/baseData"
)

type FreedAnimalEvent struct {
	MsgVersion int64                        `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	UserId     int64                        `json:"userId"`
	Animal     base_data.GrpcAnimalBaseData `json:"animal"`
}

type CaptureAnimalEvent struct {
	MsgVersion    int64                        `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	ServiceAppId  string                       `json:"serviceAppId"`
	MapId         int32                        `json:"mapId"`
	UserId        int64                        `json:"userId"`
	FreedAnimal   base_data.GrpcAnimalBaseData `json:"freedAnimal"`
	CaptureAnimal base_data.GrpcAnimalBaseData `json:"captureAnimal"`
}
