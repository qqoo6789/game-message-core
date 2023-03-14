package base_data

import "game-message-core/proto"

type GrpcTalentExp struct {
	TalentType proto.TalentType `json:"talentType"`
	Exp        int64            `json:"exp"`
}

type GrpcTalentLevel struct {
	TalentType  proto.TalentType `json:"talentType"`
	Level       uint32           `json:"level"`
	MasterLevel uint32           `json:"masterLevel"`
}

type GrpcTalentNodeData struct {
	NodeId uint32 `json:"nodeId"`
	Level  uint32 `json:"level"`
}

type GrpcTalentTree struct {
	TalentType proto.TalentType      `json:"talentType"`
	Nodes      []*GrpcTalentNodeData `json:"nodes"`
}

type GrpcTalentData struct {
	Levels []*GrpcTalentLevel `json:"levels"`
	Trees  []*GrpcTalentTree  `json:"trees"`
}

type GrpcTalentTreeUpdate struct {
	TalentType  proto.TalentType     `json:"talentType"`
	AddNodes    []GrpcTalentNodeData `json:"addNodes"`
	UpdateNodes []GrpcTalentNodeData `json:"updateNodes"`
	RemoveNodes []uint32             `json:"removeNodes"`
}
