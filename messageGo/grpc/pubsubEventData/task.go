package pubsubEventData

import (
	base_data "game-message-core/grpc/baseData"
	"game-message-core/proto"
)

type TaskAcceptEvent struct {
	MsgVersion   int64              `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	UserId       int64              `json:"userId"`
	TaskId       int32              `json:"taskId"`
	InTaskListId int32              `json:"inTaskListId"`
	TaskListType proto.TaskListType `json:"taskListType"`
}

func (p *TaskAcceptEvent) Clear() {
	p.MsgVersion = 0
	p.UserId = 0
	p.InTaskListId = 0
	p.TaskId = 0
	p.TaskListType = 0
}

type TaskAbandonEvent struct {
	MsgVersion   int64              `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	UserId       int64              `json:"userId"`
	TaskId       int32              `json:"taskId"`
	InTaskListId int32              `json:"inTaskListId"`
	TaskListType proto.TaskListType `json:"taskListType"`
}

func (p *TaskAbandonEvent) Clear() {
	p.MsgVersion = 0
	p.UserId = 0
	p.InTaskListId = 0
	p.TaskId = 0
	p.TaskListType = 0
}

type TaskFinishEvent struct {
	MsgVersion   int64                        `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	UserId       int64                        `json:"userId"`
	TaskListType proto.TaskListType           `json:"taskListType"`
	InTaskListId int32                        `json:"inTaskListId"`
	TaskId       int32                        `json:"taskId"`
	RewardItem   []base_data.GrpcItemBaseInfo `json:"rewardItem"`
}

func (p *TaskFinishEvent) Clear() {
	p.MsgVersion = 0
	p.UserId = 0
	p.TaskListType = 0
	p.InTaskListId = 0
	p.TaskId = 0
	if p.RewardItem != nil {
		p.RewardItem = p.RewardItem[:0]
	}
}

type TaskListFinishEvent struct {
	MsgVersion   int64                        `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	UserId       int64                        `json:"userId"`
	TaskListType proto.TaskListType           `json:"taskListType"`
	TaskListId   int32                        `json:"taskListId"`
	RewardItem   []base_data.GrpcItemBaseInfo `json:"rewardItem"`
}

func (p *TaskListFinishEvent) Clear() {
	p.MsgVersion = 0
	p.UserId = 0
	p.TaskListType = 0
	p.TaskListId = 0
	if p.RewardItem != nil {
		p.RewardItem = p.RewardItem[:0]
	}
}
