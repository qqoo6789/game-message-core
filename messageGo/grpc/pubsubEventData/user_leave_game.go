package pubsubEventData

type UserLeaveGameEvent struct {
	MsgVersion int64  `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	AgentAppId string `json:"agentAppId"` // 网关 appId
	UserId     int64  `json:"userId"`     //
}

func (p *UserLeaveGameEvent) Clear() {
	p.MsgVersion = 0
	p.AgentAppId = ""
	p.UserId = 0
}
