package base_data

import "game-message-core/proto"

// 对应 proto.EntityProfile
type GrpcEntityProfile struct {
	Lv         int32   `json:"lv"`
	Exp        int64   `json:"exp"`
	Att        int32   `json:"att"`
	AttSpeed   int32   `json:"attSpeed"`
	Def        int32   `json:"def"`
	HpCurrent  int32   `json:"hpCurrent"`
	HpLimit    int32   `json:"hpLimit"`
	CritRate   int32   `json:"critRate"`
	CritDmg    int32   `json:"critDmg"`
	HitRate    int32   `json:"hitRate"`
	MissRate   int32   `json:"missRate"`
	MoveSpeed  float32 `json:"moveSpeed"`
	PushDmg    int32   `json:"pushDmg"`
	PushDist   int32   `json:"pushDist"`
	HpRecovery int32   `json:"hpRecovery"`
}

func (p *GrpcEntityProfile) Set(profile *proto.EntityProfile) {
	if profile == nil {
		return
	}
	p.Lv = profile.Lv
	p.Exp = profile.Exp
	p.Att = profile.Att
	p.AttSpeed = profile.AttSpeed
	p.Def = profile.Def
	p.HpCurrent = profile.HpCurrent
	p.HpLimit = profile.HpLimit
	p.CritRate = profile.CritRate
	p.CritDmg = profile.CritDmg
	p.HitRate = profile.HitRate
	p.MissRate = profile.MissRate
	p.MoveSpeed = profile.MoveSpeed
	p.PushDmg = profile.PushDmg
	p.PushDist = profile.PushDist
	p.HpRecovery = profile.HpRecovery
}

func (p *GrpcEntityProfile) ToProtoData() *proto.EntityProfile {
	return &proto.EntityProfile{
		Lv:         p.Lv,
		Exp:        p.Exp,
		Att:        p.Att,
		AttSpeed:   p.AttSpeed,
		Def:        p.Def,
		HpCurrent:  p.HpCurrent,
		HpLimit:    p.HpLimit,
		CritRate:   p.CritRate,
		CritDmg:    p.CritDmg,
		HitRate:    p.HitRate,
		MissRate:   p.MissRate,
		MoveSpeed:  p.MoveSpeed,
		PushDmg:    p.PushDmg,
		PushDist:   p.PushDist,
		HpRecovery: p.HpRecovery,
	}
}
