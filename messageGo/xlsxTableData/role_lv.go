package xlsxTable

import "time"

type RoleLvTable struct {
	UId          uint      `gorm:"primaryKey;autoIncrement" json:"uid,string"`
	Lv           int32     `json:"lv"`
	Exp          int32     `json:"exp"`
	DeathExpLoss int32     `json:"deathExpLoss"` // 身体半径(像素)
	HpLimit      int32     `json:"hpLimit"`      // 血量上限
	HpRecovery   int32     `json:"hpRecovery"`   // 血量恢复
	Att          int32     `json:"att"`          // 攻击力
	AttSpeed     int32     `json:"attSpeed"`     // 攻击速率
	Def          int32     `json:"def"`          // 防御力
	CritRate     int32     `json:"critRate"`     // 暴击(Critical Strikes)率
	CritDmg      int32     `json:"critDmg"`      // 暴击伤害
	HitRate      int32     `json:"hitRate"`      // 名中率
	MissRate     int32     `json:"missRate"`     // 闪避率
	MoveSpeed    int32     `json:"moveSpeed"`    // 移动速度
	CreatedAt    time.Time `json:"createdAt"`    // 过期判断条件
}
