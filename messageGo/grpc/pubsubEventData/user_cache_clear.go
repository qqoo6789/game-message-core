package pubsubEventData

type ClearUserCache struct {
	MsgVersion int64 `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	UserId     int64 `json:"userId"`
}

func (p *ClearUserCache) Clear() {
	p.MsgVersion = 0
	p.UserId = 0
}
