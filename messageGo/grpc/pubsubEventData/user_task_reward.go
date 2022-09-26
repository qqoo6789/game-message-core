package pubsubEventData

type UserTaskRewardEvent struct {
	MsgVersion     int64 `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	UserId         int64 `json:"userId"`
	Exp            int32 `json:"exp"`
	ItemCid        int32 `json:"itemCid"`
	Num            int32 `json:"num"`
	Quality        int32 `json:"quality"`
	TaskListReward bool  `json:"taskListReward"`
}
