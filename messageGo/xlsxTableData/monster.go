package xlsxTable

import (
	"time"
)

type MonsterAttType int32

const (
	MonsterAttTypeUnknown MonsterAttType = 0
	// 主动攻击型
	MonsterAttTypeInitiative MonsterAttType = 1
	// 被动攻击型
	MonsterAttTypePassive MonsterAttType = 2
	// 不攻击型
	MonsterAttTypeDumb MonsterAttType = 3
)

type MonsterTableRow struct {
	UId             uint           `gorm:"primaryKey;autoIncrement" json:"uid,string"`
	Cid             int32          `json:"cid"`
	Name            string         `json:"name"`
	BodyRadius      int32          `json:"bodyRadius"` // 身体半径(像素)
	AttType         MonsterAttType `json:"attType"`
	LockEnemyRadius int32          `json:"lockEnemyRadius"` // 锁敌范围
	CombatDist      int32          `json:"combatDist"`      // 脱战距离
	DropId          int32          `json:"dropId"`          // 掉落配置ID
	SkillSequence   string         `json:"skillSequence"`   // 技能列表
	Att             int32          `json:"att"`             // 攻击力
	AttSpeed        int32          `json:"attSpeed"`        // 攻击速率
	Def             int32          `json:"def"`             // 防御力
	HpLimit         int32          `json:"hpLimit"`         // 血量上限
	CritRate        int32          `json:"critRate"`        // 暴击(Critical Strikes)率
	CritDmg         int32          `json:"critDmg"`         // 暴击伤害
	HitRate         int32          `json:"hitRate"`         // 名中率
	MissRate        int32          `json:"missRate"`        // 闪避率
	MoveSpeed       int32          `json:"moveSpeed"`       // 移动速度
	PushDmg         int32          `json:"pushDmg"`         // 击退伤害
	PushDist        int32          `json:"pushDist"`        // 击退距离
	CreatedAt       time.Time      `json:"createdAt"`       // 过期判断条件
}
