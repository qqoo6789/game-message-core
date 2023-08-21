package methodData

import (
	base_data "game-message-core/grpc/baseData"
)

type TeamDataInput struct {
	UserId int64 `json:"userId"`
}

func (p *TeamDataInput) Clear() {
	p.UserId = 0
}

type TeamDataOutput struct {
	Success  bool           `json:"success"`
	ErrMsg   string         `json:"errMsg"`
	Exist    bool           `json:"exist"`
	TeamData base_data.Team `json:"teamData"`
}

func (p *TeamDataOutput) Clear() {
	p.Success = false
	p.ErrMsg = ""
	p.Exist = false
	p.TeamData.Clear()
}
