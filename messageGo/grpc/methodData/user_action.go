package methodData

type UserLeaveGameInput struct {
	MsgVersion int64  `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	AgentAppId string `json:"agentAppId"` // 网关 appId
	UserId     int64  `json:"userId"`     //
}

type UserLeaveGameOutput struct {
	Success bool   `json:"success"`
	ErrMsg  string `json:"errMsg"`
}
