package methodData

// 服务信息用于服务 注册和释放
type PlayerLeaveGameInput struct {
	MsgVersion int64  `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	AgentAppId string `json:"agentAppId"` // 网关 appId
	UserId     int64  `json:"userId"`     //
}

// 服务信息用于服务 注册和释放
type PlayerLeaveGameOutput struct {
	MsgVersion int64  `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	AgentAppId string `json:"agentAppId"` // 网关 appId
	UserId     int64  `json:"userId"`     //
	Success    bool   `json:"success"`    //
}
