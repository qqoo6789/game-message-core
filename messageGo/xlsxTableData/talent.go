package xlsxTable

import (
	"encoding/json"
	"fmt"
	"time"
)

type RealUpgradeData struct {
	TalentTreeLv  int32   `json:"talentTreeLv"`  // 技能树等级
	TalentExpType int32   `json:"talentExpType"` // 升级所需天赋经验类型 TalentType
	TalentExp     int32   `json:"talentExp"`     // 升级经验(对应天赋类型)
	PreNodes      []int32 `json:"preNodes"`      // 前置依赖节点列表
}

type TalentUpgradeData struct {
	TalentTreeLv  int32   `json:"talentTreeLv"`  // 技能树等级
	TalentExpType int32   `json:"talentExpType"` // 升级所需天赋经验类型 TalentType
	TalentExpList []int32 `json:"talentExpList"` // 升级经验(对应天赋类型)
	PreNodes      []int32 `json:"preNodes"`      // 前置依赖节点列表
}

func (p *TalentUpgradeData) GetRealUpgradeData(lv uint32) (*RealUpgradeData, error) {
	if lv == 0 {
		return nil, fmt.Errorf("invalid Level 0")
	}
	if len(p.TalentExpList) < int(lv) {
		return nil, fmt.Errorf("node lv is max level")
	}
	ud := &RealUpgradeData{
		TalentTreeLv:  p.TalentTreeLv,
		TalentExpType: p.TalentExpType,
		PreNodes:      p.PreNodes,
		TalentExp:     p.TalentExpList[int(lv-1)],
	}
	return ud, nil
}

// 掉落物配方.
type TalentTreeRow struct {
	NodeId      int32     `gorm:"primaryKey" json:"nodeId,int32"`
	TalentType  int32     `json:"talentType"`
	IsTrunk     bool      `json:"isTrunk"`
	Layer       int32     `json:"layer"`
	LvLimit     int32     `json:"lvLimit"`
	UpgradeJson string    `gorm:"type:text" json:"upgradeJson"`
	CreatedAt   time.Time `json:"createdAt"` // 过期判断条件

	UpgradeData *TalentUpgradeData `gorm:"-" json:"-"`
}

func (this *TalentTreeRow) SetUpgradeData(upData *TalentUpgradeData) error {
	bs, err := json.Marshal(upData)
	if err != nil {
		return err
	}
	this.UpgradeJson = string(bs)
	this.UpgradeData = upData
	return nil
}

func (this *TalentTreeRow) GetUpgradeData() (*TalentUpgradeData, error) {
	if len(this.UpgradeJson) < 2 {
		return nil, nil
	}

	upData := &TalentUpgradeData{}
	err := json.Unmarshal([]byte(this.UpgradeJson), upData)
	if err != nil {
		return nil, err
	}
	this.UpgradeData = upData
	return this.UpgradeData, err
}
