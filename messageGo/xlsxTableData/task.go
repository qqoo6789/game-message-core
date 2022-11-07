package xlsxTable

import (
	"encoding/json"
	"errors"
	"fmt"
	"game-message-core/proto"
	"time"
)

// ----------- make task xlsx used types --------------------------
type TaskXlsxOption struct {
	OptionType proto.TaskOptionType `json:"optionType"`
	Value      string               `json:"value"`
}

type TaskXlsxOptions struct {
	Options []TaskXlsxOption `json:"options"`
}

//--------------- task db table used types --------------------------

type TaskRowOption struct {
	OptionType proto.TaskOptionType `json:"optionType"`
	Param1     int32                `json:"param1"`
	Param2     int32                `json:"param2"`
	Param3     int32                `json:"param3"`
	Param4     int32                `json:"param4"`
	Chance     int32                `json:"chance"` // 权重值
}

type TaskRowOptionList struct {
	// 权重和
	ChanceSum int32 `json:"chanceSum"`
	// 子项列表
	Options []TaskRowOption `json:"options"`
}

type TaskTableRow struct {
	UId                  uint   `gorm:"primaryKey;autoIncrement" json:"uid,string"`
	Id                   int32  `json:"id"`
	Level                int32  `json:"level"`
	Name                 string `json:"name"`
	SubSystem            string `json:"subSystem"`
	RewardId             int32  `json:"rewardId"`
	RewardExp            int32  `json:"rewardExp"`
	Difficulty           int32  `json:"difficulty"`
	DesignateOptionsJson string `gorm:"type:text" json:"designateOptionsJson"`
	ChanceOptionsJson    string `gorm:"type:text" json:"chanceOptionsJson"`

	CreatedAt time.Time `json:"createdAt"` // 过期判断条件,

	DesignateOptions *TaskRowOptionList `gorm:"-" json:"-"`
	ChanceOptions    *TaskRowOptionList `gorm:"-" json:"-"`
}

func (p *TaskTableRow) SetDesignateOptions(designateOptions *TaskRowOptionList) error {
	if designateOptions == nil {
		return errors.New("designateOptions is nil")
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
	if p.DesignateOptions == nil {
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
		return errors.New("designateOptions is nil")
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
	if p.ChanceOptions == nil {
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

	if p.ChanceOptions != nil || len(p.ChanceOptions.Options) > 0 {
		chanceOptionsExist = true
	}
	if p.DesignateOptions != nil || len(p.DesignateOptions.Options) > 0 {
		designateOptionsExist = true
	}
	if !chanceOptionsExist && !designateOptionsExist {
		return fmt.Errorf("task [%d] options is empty", p.Id)
	}
	return nil
}
