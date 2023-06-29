package pubsubEventData

import (
	base_data "game-message-core/grpc/baseData"
	"game-message-core/proto"
)

type TaskListFinishEvent struct {
	MsgVersion   int64                        `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	UserId       int64                        `json:"userId"`
	TaskListType proto.TaskListType           `json:"taskListType"`
	RewardItem   []base_data.GrpcItemBaseInfo `json:"rewardItem"`
}

func (p *TaskListFinishEvent) Clear() {
	p.MsgVersion = 0
	p.UserId = 0
	p.TaskListType = 0
	if p.RewardItem != nil {
		p.RewardItem = p.RewardItem[:0]
	}
}
