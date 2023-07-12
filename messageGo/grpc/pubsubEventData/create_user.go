package pubsubEventData

import (
	"game-message-core/proto"
)

type CreateUserEvent struct {
	MsgVersion int64              `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	UserId     int64              `json:"userId,omitempty"`
	NickName   string             `json:"nickName,omitempty"`
	RoleId     int32              `json:"roleId,omitempty"`
	Gender     string             `json:"gender,omitempty"`
	Icon       string             `json:"icon,omitempty"`
	OS         proto.OSType       `json:"osType,omitempty"`
	PlatFrom   proto.PlatformType `json:"platform,omitempty"`
}

func (p *CreateUserEvent) Clear() {
	p.MsgVersion = 0
	p.UserId = 0
	p.NickName = ""
	p.RoleId = 0
	p.Gender = ""
	p.Icon = ""
	p.OS = proto.OSType_OSTypeUnknow
	p.PlatFrom = proto.PlatformType_PlatformUnKnow
}
