package pubsubEventData

import (
	base_data "game-message-core/grpc/baseData"
	"game-message-core/proto"
)

type UpdateTalentEvent struct {
	MsgVersion int64                            `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	UserId     int64                            `json:"userId"`
	Levels     []base_data.GrpcTalentLevel      `json:"levels"`
	Trees      []base_data.GrpcTalentTreeUpdate `json:"trees"`
}

func (p *UpdateTalentEvent) Clear() {
	p.MsgVersion = 0
	p.UserId = 0
	if p.Levels != nil {
		p.Levels = p.Levels[:0]
	}
	if p.Trees != nil {
		p.Trees = p.Trees[:0]
	}
}

type AddTalentExpEvent struct {
	MsgVersion     int64                     `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	FormService    base_data.ServiceData     `json:"formService"`
	UserId         int64                     `json:"userId"`
	AddExps        []base_data.GrpcTalentExp `json:"addExps"`
	FromEntityType proto.EntityType          `json:"fromEntityType"`
	FromEntityCid  int32                     `json:"fromEntityCid"`
	FromEntityName string                    `json:"fromEntityName"`
}

func (p *AddTalentExpEvent) Clear() {
	p.MsgVersion = 0
	p.FormService.Clear()
	p.UserId = 0
	p.FromEntityType = 0
	p.FromEntityCid = 0
	p.FromEntityName = ""
	if p.AddExps != nil {
		p.AddExps = p.AddExps[:0]
	}
}
