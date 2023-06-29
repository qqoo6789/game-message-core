package methodData

import (
	base_data "game-message-core/grpc/baseData"
)

// 更新玩家使用的装备
type UpdateUsedAvatarInput struct {
	MsgVersion   int64                           `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	UserId       int64                           `json:"userId"`
	UsingAvatars []base_data.GrpcAvatarAttribute `json:"usingAvatars"`
	CurProfile   []base_data.GrpcAttributeData   `json:"curProfile"`
}

func (p *UpdateUsedAvatarInput) Clear() {
	p.MsgVersion = 0
	p.UserId = 0
	if p.UsingAvatars != nil {
		p.UsingAvatars = p.UsingAvatars[:0]
	}
	if p.CurProfile != nil {
		p.CurProfile = p.CurProfile[:0]
	}
}

type UpdateUsedAvatarOutput struct {
	MsgVersion int64 `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	Success    bool  `json:"success"`
}

func (p *UpdateUsedAvatarOutput) Clear() {
	p.Success = false
	p.MsgVersion = 0
}

// 更新玩家战斗属性(升级/装备槽等级变更)
type UpdateUserProfileInput struct {
	MsgVersion int64                         `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	UserId     int64                         `json:"userId"`
	CurProfile []base_data.GrpcAttributeData `json:"curProfile"`
}

func (p *UpdateUserProfileInput) Clear() {
	p.MsgVersion = 0
	p.UserId = 0
	if p.CurProfile != nil {
		p.CurProfile = p.CurProfile[:0]
	}
}

type UpdateUserProfileOutput struct {
	MsgVersion int64 `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	Success    bool  `json:"success"`
}

func (p *UpdateUserProfileOutput) Clear() {
	p.Success = false
	p.MsgVersion = 0
}

// 获取user 详细数据
type GetUserDataInput struct {
	MsgVersion int64 `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	UserId     int64 `json:"userId"`
}

func (p *GetUserDataInput) Clear() {
	p.MsgVersion = 0
	p.UserId = 0
}

type GetUserDataOutput struct {
	MsgVersion int64                           `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	Success    bool                            `json:"success"`
	ErrMsg     string                          `json:"errMsg"`
	BaseData   base_data.GrpcPlayerBaseData    `json:"baseData"`
	Profile    []base_data.GrpcAttributeData   `json:"profile"`
	Feature    base_data.GrpcPlayerFeature     `json:"feature"`
	MapId      int32                           `json:"mapId"`
	Pos        base_data.GrpcVector3           `json:"pos"`
	Dir        base_data.GrpcVector3           `json:"dir"`
	Avatars    []base_data.GrpcAvatarAttribute `json:"avatars"`
	TalentData base_data.GrpcTalentData        `json:"talentData"`
}

func (p *GetUserDataOutput) Clear() {
	p.MsgVersion = 0
	p.Success = false
	p.ErrMsg = ""
	p.BaseData.Clear()
	p.Feature.Clear()
	p.MapId = 0
	p.Pos.Clear()
	p.Dir.Clear()
	p.TalentData.Clear()
	if p.Avatars != nil {
		p.Avatars = p.Avatars[:0]
	}
	if p.Profile != nil {
		p.Profile = p.Profile[:0]
	}
}
