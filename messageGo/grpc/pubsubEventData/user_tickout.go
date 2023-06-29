package pubsubEventData

import "game-message-core/proto"

type TickOutPlayerEvent struct {
	MsgVersion        int64             `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	UserId            int64             `json:"userId"`
	AgentAppId        string            `json:"agentAppId"`
	SceneServiceAppId string            `json:"sceneServiceAppId"`
	SocketId          string            `json:"socketId"`
	TickOutCode       proto.TickOutType `json:"tickOutCode"`
}

func (p *TickOutPlayerEvent) Clear() {
	p.MsgVersion = 0
	p.UserId = 0
	p.AgentAppId = ""
	p.SceneServiceAppId = ""
	p.SocketId = ""
	p.TickOutCode = 0
}
