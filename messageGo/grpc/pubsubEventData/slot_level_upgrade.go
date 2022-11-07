package pubsubEventData

type SlotLevelUpgradeEvent struct {
	MsgVersion int64 `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	UserId     int64 `json:"userId"`
	SlotPos    int32 `json:"slotPos"`
	Level      int32 `json:"level"`
}
