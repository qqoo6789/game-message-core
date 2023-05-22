package xlsxTable

import (
	"encoding/json"
	"fmt"
	"game-message-core/proto"
	"time"
)

// ----------- make task xlsx used types --------------------------
type TaskXlsxOption struct {
	OptionType  proto.TaskOptionType `json:"optionType"`
	Value       string               `json:"value"`
	GuiderPoint int32                `json:"guiderPoint"`
}

type TaskXlsxOptions struct {
	Options []TaskXlsxOption `json:"options"`
}

//--------------- task db table used types --------------------------

type TaskRowOption struct {
	OptionType  proto.TaskOptionType `json:"optionType"`
	Param1      int32                `json:"param1"`
	Param2      int32                `json:"param2"`
	Param3      int32                `json:"param3"`
	Param4      int32                `json:"param4"`
	Chance      int32                `json:"chance"` // 权重值
	GuiderPoint int32                `json:"guiderPoint"`
}

type TaskRowOptionList struct {
	// 权重和
	ChanceSum int32 `json:"chanceSum"`
	// 子项列表
	Options []TaskRowOption `json:"options"`
}

type TaskLevelCondition struct {
	TalentType proto.TalentType `json:"talentType"`
	Level      int32            `json:"level"`
}
type TaskAcceptCondition struct {
	ConditionLvs []TaskLevelCondition `json:"conditionLvs"`
	PreTasks     []int32              `json:"preTasks"`
}

type TaskRewardExp struct {
	TalentType proto.TalentType `json:"talentType"`
	Exp        int32            `json:"exp"`
}
type TaskReward struct {
	RewardId   int32           `json:"rewardId"`
	RewardExps []TaskRewardExp `json:"rewardExps"`
}

type TaskTableRow struct {
	Id                   int32     `gorm:"primaryKey" json:"id"`
	Name                 string    `json:"name"`
	ConditionJson        string    `gorm:"type:text" json:"conditionJson"`
	RewardJson           string    `gorm:"type:text" json:"rewardJson"`
	DesignateOptionsJson string    `gorm:"type:text" json:"designateOptionsJson"`
	ChanceOptionsJson    string    `gorm:"type:text" json:"chanceOptionsJson"`
	CreatedAt            time.Time `json:"createdAt"` // 过期判断条件,

	DesignateOptions *TaskRowOptionList   `gorm:"-" json:"-"`
	ChanceOptions    *TaskRowOptionList   `gorm:"-" json:"-"`
	AcceptCondition  *TaskAcceptCondition `gorm:"-" json:"-"`
	RewardData       *TaskReward          `gorm:"-" json:"-"`
}

func (p *TaskTableRow) SetReward(reward *TaskReward) error {
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

func (p *TaskTableRow) GetReward() (*TaskReward, error) {
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

func (p *TaskTableRow) SetAcceptCondition(condition *TaskAcceptCondition) error {
	if condition == nil {
		p.AcceptCondition = nil
		p.ConditionJson = ""
		return nil
	}

	bs, err := json.Marshal(condition)
	if err != nil {
		return err
	}

	p.AcceptCondition = condition
	p.ConditionJson = string(bs)
	return nil
}

func (p *TaskTableRow) GetAcceptCondition() (*TaskAcceptCondition, error) {
	if p.AcceptCondition == nil && len(p.ConditionJson) > 2 {
		data := &TaskAcceptCondition{}
		err := json.Unmarshal([]byte(p.ConditionJson), data)
		if err != nil {
			return nil, err
		}
		p.AcceptCondition = data
	}
	return p.AcceptCondition, nil
}

func (p *TaskTableRow) SetDesignateOptions(designateOptions *TaskRowOptionList) error {
	if designateOptions == nil {
		p.DesignateOptions = nil
		p.DesignateOptionsJson = ""
	}

	bs, err := json.Marshal(designateOptions)
	if err != nil {
		return err
	}

	p.DesignateOptions = designateOptions
	p.DesignateOptionsJson = string(bs)
	return nil
}

func (p *TaskTableRow) GetDesignateOptions() (*TaskRowOptionList, error) {
	if p.DesignateOptions == nil && len(p.DesignateOptionsJson) > 2 {
		options := &TaskRowOptionList{}
		err := json.Unmarshal([]byte(p.DesignateOptionsJson), options)
		if err != nil {
			return nil, err
		}
		p.DesignateOptions = options
	}
	return p.DesignateOptions, nil
}

func (p *TaskTableRow) SetChanceOptions(chanceOptions *TaskRowOptionList) error {
	if chanceOptions == nil {
		p.ChanceOptions = nil
		p.ChanceOptionsJson = ""
		return nil
	}

	bs, err := json.Marshal(chanceOptions)
	if err != nil {
		return err
	}

	p.ChanceOptions = chanceOptions
	p.ChanceOptionsJson = string(bs)
	return nil
}

func (p *TaskTableRow) GetChanceOptions() (*TaskRowOptionList, error) {
	if p.ChanceOptions == nil && len(p.ChanceOptionsJson) > 2 {
		options := &TaskRowOptionList{}
		err := json.Unmarshal([]byte(p.ChanceOptionsJson), options)
		if err != nil {
			return nil, err
		}
		p.ChanceOptions = options
	}
	return p.ChanceOptions, nil
}

func (p *TaskTableRow) Check() error {
	chanceOptionsExist := false
	designateOptionsExist := false

	if p.ChanceOptions != nil && len(p.ChanceOptions.Options) > 0 {
		chanceOptionsExist = true
	}
	if p.DesignateOptions != nil && len(p.DesignateOptions.Options) > 0 {
		designateOptionsExist = true
	}
	if !chanceOptionsExist && !designateOptionsExist {
		return fmt.Errorf("task [%d] options is empty", p.Id)
	}
	return nil
}
