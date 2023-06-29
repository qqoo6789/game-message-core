package base_data

import "game-message-core/proto"

type GrpcTalentExp struct {
	TalentType proto.TalentType `json:"talentType"`
	Exp        int64            `json:"exp"`
}

func (p *GrpcTalentExp) Clear() {
	p.TalentType = 0
	p.Exp = 0
}

type GrpcTalentLevel struct {
	TalentType  proto.TalentType `json:"talentType"`
	Level       uint32           `json:"level"`
	MasterLevel uint32           `json:"masterLevel"`
}

func (p *GrpcTalentLevel) Clear() {
	p.TalentType = 0
	p.Level = 0
	p.MasterLevel = 0
}

type GrpcTalentNodeData struct {
	NodeId uint32 `json:"nodeId"`
	Level  uint32 `json:"level"`
}

func (p *GrpcTalentNodeData) Clear() {
	p.Level = 0
	p.NodeId = 0
}

type GrpcTalentTree struct {
	TalentType proto.TalentType      `json:"talentType"`
	Nodes      []*GrpcTalentNodeData `json:"nodes"`
}

func (p *GrpcTalentTree) Clear() {
	p.TalentType = 0
	if p.Nodes != nil {
		p.Nodes = p.Nodes[:0]
	}
}

type GrpcTalentData struct {
	Levels []*GrpcTalentLevel `json:"levels"`
	Trees  []*GrpcTalentTree  `json:"trees"`
}

func (p *GrpcTalentData) Clear() {
	if p.Levels != nil {
		p.Levels = p.Levels[:0]
	}
	if p.Trees != nil {
		p.Trees = p.Trees[:0]
	}
}

type GrpcTalentTreeUpdate struct {
	TalentType  proto.TalentType     `json:"talentType"`
	AddNodes    []GrpcTalentNodeData `json:"addNodes"`
	UpdateNodes []GrpcTalentNodeData `json:"updateNodes"`
	RemoveNodes []uint32             `json:"removeNodes"`
}

func (p *GrpcTalentTreeUpdate) Clear() {
	p.TalentType = 0
	if p.AddNodes != nil {
		p.AddNodes = p.AddNodes[:0]
	}
	if p.UpdateNodes != nil {
		p.UpdateNodes = p.UpdateNodes[:0]
	}
	if p.RemoveNodes != nil {
		p.RemoveNodes = p.RemoveNodes[:0]
	}
}
