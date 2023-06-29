package methodData

import (
	base_data "game-message-core/grpc/baseData"
)

// call main service all player build entities by map id
type MainServiceActionGetAllBuildInput struct {
	MapId int32 `json:"mapId"`
}

func (p *MainServiceActionGetAllBuildInput) Clear() {
	p.MapId = 0
}

type MainServiceActionGetAllBuildOutput struct {
	Success   bool                     `json:"success"`
	ErrMsg    string                   `json:"errMsg"`
	AllBuilds []base_data.GrpcNftBuild `json:"allBuilds"`
}

func (p *MainServiceActionGetAllBuildOutput) Clear() {
	p.Success = false
	p.ErrMsg = ""
	if p.AllBuilds != nil {
		p.AllBuilds = p.AllBuilds[:0]
	}
}
