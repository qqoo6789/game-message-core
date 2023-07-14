package methodData

import "game-message-core/proto"

type TickOutPlayerInput struct {
	UserId            int64             `json:"userId"`
	AgentAppId        string            `json:"agentAppId"`
	SceneServiceAppId string            `json:"sceneServiceAppId"`
	SocketId          string            `json:"socketId"`
	TickOutCode       proto.TickOutType `json:"tickOutCode"`
}

func (p *TickOutPlayerInput) Clear() {
	p.UserId = 0
	p.AgentAppId = ""
	p.SceneServiceAppId = ""
	p.SocketId = ""
	p.TickOutCode = 0
}

type TickOutPlayerOutput struct {
	Success bool   `json:"success"`
	ErrMsg  string `json:"errMsg"`
}

func (p *TickOutPlayerOutput) Clear() {
	p.Success = false
	p.ErrMsg = ""
}
