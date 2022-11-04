package xlsxTable

import (
	"encoding/json"
	"time"
)

type TaskListTableParam struct {
	TaskId int32 `json:"taskId"`
	Chance int32 `json:"chance"`
}
type TaskListTableIncludeTasks struct {
	ChanceSum int32                `json:"chanceSum"`
	Param     []TaskListTableParam `json:"param"`
}

type TaskListTableRow struct {
	UId             uint      `gorm:"primaryKey;autoIncrement" json:"uid,string"`
	Id              int32     `json:"id"`
	Level           int32     `json:"level"`
	System          int32     `json:"system"`
	IncludeTaskJson string    `gorm:"type:text" json:"includeTask"`
	RewardId        int32     `json:"rewardId"`
	RewardExp       int32     `json:"rewardExp"`
	ProgressReset   bool      `json:"progressReset"`
	NeedMELD        int32     `json:"needMELD"`
	CreatedAt       time.Time `json:"createdAt"` // 过期判断条件,

	IncludeTask *TaskListTableIncludeTasks `gorm:"-" json:"-"`
}

func (p *TaskListTableRow) SetIncludeTask(include *TaskListTableIncludeTasks) error {
	bs, err := json.Marshal(include)
	if err != nil {
		return err
	}
	p.IncludeTaskJson = string(bs)
	p.IncludeTask = include
	return err
}

func (p *TaskListTableRow) GetIncludeTask() (*TaskListTableIncludeTasks, error) {
	if p.IncludeTask == nil {
		include := &TaskListTableIncludeTasks{}
		err := json.Unmarshal([]byte(p.IncludeTaskJson), include)
		if err != nil {
			return nil, err
		}

		p.IncludeTask = include
	}
	return p.IncludeTask, nil
}
