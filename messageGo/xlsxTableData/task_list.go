package xlsxTable

import (
	"encoding/json"
	"fmt"
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

type TaskListTableRow struct {
	Id            int32     `gorm:"primaryKey" json:"id"`
	System        int32     `json:"system"`
	TalentType    int32     `json:"talentType"`
	Level         int32     `json:"level"`
	TaskPoolJson  string    `gorm:"type:text" json:"taskPoolJson"`
	RewardJson    string    `gorm:"type:text" json:"rewardJson"`
	ProgressReset bool      `json:"progressReset"`
	NeedMELD      int32     `json:"needMELD"`
	CreatedAt     time.Time `json:"createdAt"` // 过期判断条件,

	TaskPool   *TaskListTableTaskPool `gorm:"-" json:"-"`
	RewardData *TaskReward            `gorm:"-" json:"-"`
}

func (p *TaskListTableRow) SetReward(reward *TaskReward) error {
	if reward == nil {
		p.RewardData = nil
		p.RewardJson = ""
		return nil
	}

	bs, err := json.Marshal(reward)
	if err != nil {
		return err
	}
	p.RewardData = reward
	p.RewardJson = string(bs)
	return nil
}

func (p *TaskListTableRow) GetReward() (*TaskReward, error) {
	if p.RewardData == nil && len(p.RewardJson) > 2 {
		data := &TaskReward{}
		err := json.Unmarshal([]byte(p.RewardJson), data)
		if err != nil {
			return nil, err
		}
		p.RewardData = data
	}
	return p.RewardData, nil
}

func (p *TaskListTableRow) SetTaskPool(taskPool *TaskListTableTaskPool) error {
	if taskPool == nil {
		p.TaskPoolJson = ""
		p.TaskPool = nil
		return nil
	}

	bs, err := json.Marshal(taskPool)
	if err != nil {
		return err
	}
	p.TaskPoolJson = string(bs)
	p.TaskPool = taskPool
	return err
}

func (p *TaskListTableRow) GetTaskPool() (*TaskListTableTaskPool, error) {
	if p.TaskPool == nil && len(p.TaskPoolJson) > 2 {
		taskPool := &TaskListTableTaskPool{}
		err := json.Unmarshal([]byte(p.TaskPoolJson), taskPool)
		if err != nil {
			return nil, err
		}

		p.TaskPool = taskPool
	}
	return p.TaskPool, nil
}

func (p *TaskListTableRow) Check() error {
	taskPoolExist := false
	taskSequenceExist := false

	if p.TaskPool != nil && len(p.TaskPool.Param) > 0 {
		taskPoolExist = true
	}
	if !taskPoolExist && !taskSequenceExist {
		return fmt.Errorf("task list [%d] task list is empty", p.Id)
	}
	return nil
}
