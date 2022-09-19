package pubsubEventData

type PlayerDeathEventData struct {
	MsgVersion int64   `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	MapId      int32   `json:"mapId"`
	UserId     int64   `json:"userId"`
	PosX       float32 `json:"posX"`
	PosY       float32 `json:"posY"`
	PosZ       float32 `json:"posZ"`
	KillerType int32   `json:"killerType"`
	KillerId   int64   `json:"killerId"`
	KillerName string  `json:"killerName"`
}
