package xlsxTable

import "time"

type SlotLvTableRow struct {
	UId        uint      `gorm:"primaryKey;autoIncrement" json:"uid,string"`
	Position   int32     `json:"position"` // 道具槽位置 遵循 AvatarPosition 规则
	Lv         int32     `json:"lv"`
	UpExp      int32     `json:"upExp"`      // 升级下一级消耗的经验
	UpMeld     int32     `json:"upMeld"`     // 升级下一级消耗的meld
	HpLimit    int32     `json:"hpLimit"`    // 血量上限
	HpRecovery int32     `json:"hpRecovery"` // 血量恢复
	Att        int32     `json:"att"`        // 攻击力
	AttSpeed   int32     `json:"attSpeed"`   // 攻击速率
	Def        int32     `json:"def"`        // 防御力
	CritRate   int32     `json:"critRate"`   // 暴击(Critical Strikes)率
	CritDmg    int32     `json:"critDmg"`    // 暴击伤害
	HitRate    int32     `json:"hitRate"`    // 名中率
	MissRate   int32     `json:"missRate"`   // 闪避率
	MoveSpeed  int32     `json:"moveSpeed"`  // 移动速度
	CreatedAt  time.Time `json:"createdAt"`  // 过期判断条件
}
