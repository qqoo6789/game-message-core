package pubsubEventData

import (
	base_data "game-message-core/grpc/baseData"
)

type CreateTeamEvent struct {
	MsgVersion int64          `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	TeamData   base_data.Team `json:"teamData"`
}

func (p *CreateTeamEvent) Clear() {
	p.MsgVersion = 0
	p.TeamData.Clear()
}

type DisbandTeamEvent struct {
	MsgVersion int64 `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	TeamId     int64 `json:"teamId"`
}

func (p *DisbandTeamEvent) Clear() {
	p.MsgVersion = 0
	p.TeamId = 0
}

type JoinTeamEvent struct {
	MsgVersion int64                `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	TeamData   base_data.Team       `json:"teamData"`
	JoinMember base_data.TeamMember `json:"joinMember"`
}

func (p *JoinTeamEvent) Clear() {
	p.MsgVersion = 0
	p.TeamData.Clear()
	p.JoinMember.Clear()
}

type QuitTeamEvent struct {
	MsgVersion int64 `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	TeamId     int64 `json:"teamId"`
	Member     int64 `json:"member"`
}

func (p *QuitTeamEvent) Clear() {
	p.MsgVersion = 0
	p.TeamId = 0
	p.Member = 0
}
