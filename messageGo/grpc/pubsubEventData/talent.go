package pubsubEventData

import (
	base_data "game-message-core/grpc/baseData"
)

type UpdateTalentEvent struct {
	MsgVersion  int64                          `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	UserId      int64                          `json:"userId"`
	Levels      []base_data.GrpcTalentLevel    `json:"levels"`
	AddNodes    []base_data.GrpcTalentNodeData `json:"addNodes"`
	UpdateNodes []base_data.GrpcTalentNodeData `json:"updateNodes"`
	RemoveNodes []base_data.GrpcTalentNodeData `json:"removeNodes"`
}
