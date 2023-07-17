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

type AddTalentExpInfo struct {
	Owner      int64                     `json:"owner"`
	AddExps    []base_data.GrpcTalentExp `json:"addExps"`
	FromEntity base_data.EntityInfo      `json:"fromEntity"`
}

func (p *AddTalentExpInfo) Clear() {
	p.Owner = 0
	p.FromEntity.Clear()
	if p.AddExps != nil {
		p.AddExps = p.AddExps[:0]
	}
}

type AddTalentExpEvent struct {
	MsgVersion  int64                 `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	FormService base_data.ServiceData `json:"formService"`
	AddExp      AddTalentExpInfo      `json:"addExp"`
}

func (p *AddTalentExpEvent) Clear() {
	p.MsgVersion = 0
	p.FormService.Clear()
	p.AddExp.Clear()
}

type TakeTalentExpEvent struct {
	MsgVersion  int64                     `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	Owner       int64                     `json:"owner"`
	TakeExps    []base_data.GrpcTalentExp `json:"takeExps"`
	TakeDesc    string                    `json:"takeDesc"`
	FormService base_data.ServiceData     `json:"formService"`
}

func (p *TakeTalentExpEvent) Clear() {
	p.MsgVersion = 0
	p.Owner = 0
	p.TakeDesc = ""
	p.FormService.Clear()
	for i := 0; i < len(p.TakeExps); i++ {
		p.TakeExps[i].Clear()
	}
	if p.TakeExps != nil {
		p.TakeExps = p.TakeExps[:0]
	}
}
