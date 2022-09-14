package methodData

// agent service forward client message request
type PullClientMessageInput struct {
	MsgVersion int64  `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	AgentAppId string `json:"agentAppId"` // 网关 appId
	UserId     int64  `json:"userId"`     //
	MsgBody    []byte `json:"msgBody"`    // proto message marshal bytes
}

// agent service forward client message response
type PullClientMessageOutput struct {
	Success bool   `json:"success"`
	ErrMsg  string `json:"errMsg"`
}

// other service broadcast proto message to client request
type BroadCastToClientInput struct {
	MsgVersion   int64  `json:"msgVersion"`   // 消息版本号 值为毫秒时间戳
	ServiceAppId string `json:"serviceAppId"` // 网关 appId
	UserId       int64  `json:"userId"`       //
	MsgBody      []byte `json:"msgBody"`      // proto message marshal bytes
}

// other service broadcast proto message to client response
type BroadCastToClientOutput struct {
	Success bool   `json:"success"`
	ErrMsg  string `json:"errMsg"`
}

// other service multiple broadcast proto message to client request
type MultipleBroadCastToClientInput struct {
	MsgVersion   int64   `json:"msgVersion"`   // 消息版本号 值为毫秒时间戳
	ServiceAppId string  `json:"serviceAppId"` // 网关 appId
	UserList     []int64 `json:"userList"`     //
	MsgBody      []byte  `json:"msgBody"`      // proto message marshal bytes
}

// other service multiple broadcast proto message to client response
type MultipleBroadCastToClientOutput struct {
	Success bool   `json:"success"`
	ErrMsg  string `json:"errMsg"`
}
