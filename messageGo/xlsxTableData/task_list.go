package xlsxTable

import (
	"time"
)

type TaskListTableRow struct {
	UId             uint      `gorm:"primaryKey;autoIncrement" json:"uid,string"`
	Level           int32     `json:"level"`
	System          int32     `json:"system"`
	IncludeTaskJson string    `gorm:"type:text" json:"includeTask"`
	RewardItemsJson string    `gorm:"type:text" json:"rewardItemsJson"`
	RewardExp       int32     `json:"rewardExp"`
	ProgressReset   bool      `json:"progressReset"`
	NeedMELD        int32     `json:"needMELD"`
	CreatedAt       time.Time `json:"createdAt"` // 过期判断条件,

	IncludeTask *TaskObjectList `gorm:"-" json:"-"`
	RewardItems *TaskObjectList `gorm:"-" json:"-"`
}

func (p *TaskListTableRow) SetRewardItems(objs *TaskObjectList) error {
	bs, err := marshal(objs)
	p.RewardItemsJson = string(bs)
	return err
}
func (p *TaskListTableRow) GetRewardItems() *TaskObjectList {
	if p.RewardItems == nil {
		p.RewardItems, _ = unMarshal(p.RewardItemsJson)
	}
	return p.RewardItems
}

func (p *TaskListTableRow) SetIncludeTask(objs *TaskObjectList) error {
	bs, err := marshal(objs)
	p.IncludeTaskJson = string(bs)
	return err
}
func (p *TaskListTableRow) GetIncludeTask() *TaskObjectList {
	if p.IncludeTask == nil {
		p.IncludeTask, _ = unMarshal(p.IncludeTaskJson)
	}
	return p.IncludeTask
}
