package pubsubEventData

import base_data "game-message-core/grpc/baseData"

type UserTaskRewardEvent struct {
	MsgVersion     int64                        `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	UserId         int64                        `json:"userId"`
	Exps           []base_data.GrpcTalentExp    `json:"exps"`
	Items          []base_data.GrpcItemBaseInfo `json:"items"`
	TaskListReward bool                         `json:"taskListReward"`
}
