package pubsubEventData

type TickOutPlayerEvent struct {
	MsgVersion        int64  `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	UserId            int64  `json:"userId"`
	AgentAppId        string `json:"agentAppId"`
	SceneServiceAppId string `json:"sceneServiceAppId"`
	SocketId          string `json:"socketId"`
}
