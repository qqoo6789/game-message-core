package pubsubEventData

type ReloadConfigDataEvent struct {
	MsgVersion int64 `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
}
