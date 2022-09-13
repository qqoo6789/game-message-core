package methodData

// agent service forward client message request
type ClientProtoMessageInput struct {
	MsgVersion    int64  `json:"msgVersion"`    // 消息版本号 值为毫秒时间戳
	AgentAppId    string `json:"agentAppId"`    // 网关 appId
	UserId        int64  `json:"userId"`        //
	ProtoBytesReq []byte `json:"protoBytesReq"` // proto message marshal bytes
}

// agent service forward client message response
type ClientProtoMessageOutput struct {
	MsgVersion     int64  `json:"msgVersion"`     // 消息版本号 值为毫秒时间戳
	AgentAppId     string `json:"agentAppId"`     // 网关 appId
	UserId         int64  `json:"userId"`         //
	ProtoBytesResp []byte `json:"protoBytesResp"` // proto message marshal bytes
}
