package xlsxTable

import (
	"encoding/json"
	"time"
)

type TaskPoolParam struct {
	TaskId int32 `json:"taskId"`
	Chance int32 `json:"chance"`
}
type TaskListTableTaskPool struct {
	ChanceSum int32           `json:"chanceSum"`
	Param     []TaskPoolParam `json:"param"`
}

type TaskListTableTaskSequence struct {
	Sequence []int32 `json:"sequence"`
}

type TaskListTableRow struct {
	UId              uint      `gorm:"primaryKey;autoIncrement" json:"uid,string"`
	Id               int32     `json:"id"`
	Level            int32     `json:"level"`
	System           int32     `json:"system"`
	TaskPoolJson     string    `gorm:"type:text" json:"taskPoolJson"`
	TaskSequenceJson string    `gorm:"type:text" json:"taskSequenceJson"`
	RewardId         int32     `json:"rewardId"`
	RewardExp        int32     `json:"rewardExp"`
	ProgressReset    bool      `json:"progressReset"`
	NeedMELD         int32     `json:"needMELD"`
	CreatedAt        time.Time `json:"createdAt"` // 过期判断条件,

	TaskPool     *TaskListTableTaskPool     `gorm:"-" json:"-"`
	TaskSequence *TaskListTableTaskSequence `gorm:"-" json:"-"`
}

func (p *TaskListTableRow) SetTaskPool(taskPool *TaskListTableTaskPool) error {
	bs, err := json.Marshal(taskPool)
	if err != nil {
		return err
	}
	p.TaskPoolJson = string(bs)
	p.TaskPool = taskPool
	return err
}

func (p *TaskListTableRow) GetTaskPool() (*TaskListTableTaskPool, error) {
	if p.TaskPool == nil {
		taskPool := &TaskListTableTaskPool{}
		err := json.Unmarshal([]byte(p.TaskPoolJson), taskPool)
		if err != nil {
			return nil, err
		}

		p.TaskPool = taskPool
	}
	return p.TaskPool, nil
}

func (p *TaskListTableRow) SetSequence(taskSequence *TaskListTableTaskSequence) error {
	bs, err := json.Marshal(taskSequence)
	if err != nil {
		return err
	}
	p.TaskSequenceJson = string(bs)
	p.TaskSequence = taskSequence
	return err
}

func (p *TaskListTableRow) GetSequence() (*TaskListTableTaskSequence, error) {
	if p.TaskSequence == nil {
		taskSequence := &TaskListTableTaskSequence{}
		err := json.Unmarshal([]byte(p.TaskSequenceJson), taskSequence)
		if err != nil {
			return nil, err
		}

		p.TaskSequence = taskSequence
	}
	return p.TaskSequence, nil
}
