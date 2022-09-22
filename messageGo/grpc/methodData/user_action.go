package methodData

import "game-message-core/proto"

type UserLeaveGameInput struct {
	MsgVersion int64  `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	AgentAppId string `json:"agentAppId"` // 网关 appId
	UserId     int64  `json:"userId"`     //
}
type UserLeaveGameOutput struct {
	Success bool   `json:"success"`
	ErrMsg  string `json:"errMsg"`
}

// 更新玩家使用的装备
type UpdateUsedAvatarInput struct {
	MsgVersion   int64                `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	UserId       int64                `json:"userId"`
	UsingAvatars []proto.PlayerAvatar `json:"usingAvatars"`
	CurProfile   proto.EntityProfile  `json:"curProfile"`
}
type UpdateUsedAvatarOutput struct {
	MsgVersion int64 `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	Success    bool  `json:"success"`
}

// 更新玩家战斗属性(升级/装备槽等级变更)
type UpdateUserProfileInput struct {
	MsgVersion int64               `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	UserId     int64               `json:"userId"`
	CurProfile proto.EntityProfile `json:"curProfile"`
}
type UpdateUserProfileOutput struct {
	MsgVersion int64 `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	Success    bool  `json:"success"`
}

// 获取user 详细数据
type GetUserDataInput struct {
	MsgVersion int64 `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	UserId     int64 `json:"userId"`
}
type GetUserDataOutput struct {
	MsgVersion int64                `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	Success    bool                 `json:"success"`
	ErrMsg     string               `json:"errMsg"`
	BaseData   proto.PlayerBaseData `json:"baseData"`
	Profile    proto.EntityProfile  `json:"profile"`
	MapId      int32                `json:"mapId"`
	Pos        proto.Vector3        `json:"pos"`
	Dir        proto.Vector3        `json:"dir"`
	Avatars    []proto.PlayerAvatar `json:"avatars"`
}
