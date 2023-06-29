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
	MsgVersion   int64                     `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	ServiceAppId string                    `json:"serviceAppId"`
	MapId        int32                     `json:"mapId"`
	UserId       int64                     `json:"userId"`
	AddExps      []base_data.GrpcTalentExp `json:"addExps"`
}

func (p *AddTalentExpEvent) Clear() {
	p.MsgVersion = 0
	p.ServiceAppId = ""
	p.MapId = 0
	p.UserId = 0
	if p.AddExps != nil {
		p.AddExps = p.AddExps[:0]
	}
}
