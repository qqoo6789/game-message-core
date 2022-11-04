package xlsxTable

import (
	"encoding/json"
	"errors"
	"game-message-core/proto"
	"time"
)

// ----------- make task xlsx used types --------------------------
type TaskXlsxRowOption struct {
	OptionType proto.TaskOptionType `json:"optionType"`
	Value      string               `json:"value"`
}

type TaskXlsxOptions struct {
	Options []TaskXlsxRowOption `json:"options"`
}

//--------------- task db table used types --------------------------

type TaskTableOptionParam struct {
	Param1 int32 `json:"param1"`
	Param2 int32 `json:"param2"`
	Param3 int32 `json:"param3"`
	Param4 int32 `json:"param4"`
	Param5 int32 `json:"param5"`
}

type TaskTableOption struct {
	OptionType      proto.TaskOptionType   `json:"optionType"`
	RandomExclusive int32                  `json:"randomSum"` // 通常为10000
	RandList        []TaskTableOptionParam `json:"randList"`
}

type TaskTableRowOptionList struct {
	Options []TaskTableOption `json:"options"`
}

type TaskTableRow struct {
	UId         uint   `gorm:"primaryKey;autoIncrement" json:"uid,string"`
	Id          int32  `json:"id"`
	Level       int32  `json:"level"`
	Name        string `json:"name"`
	SubSystem   string `json:"subSystem"`
	RewardId    int32  `json:"rewardId"`
	RewardExp   int32  `json:"rewardExp"`
	Difficulty  int32  `json:"difficulty"`
	NextTaskId  int32  `json:"nextTaskId"`
	OptionsJson string `gorm:"type:text" json:"optionsJson"`

	CreatedAt time.Time `json:"createdAt"` // 过期判断条件,

	Options *TaskTableRowOptionList `gorm:"-" json:"-"`
}

func (p *TaskTableRow) SetOptions(options *TaskTableRowOptionList) error {
	if options == nil {
		return errors.New("options is nil")
	}
	if len(options.Options) < 1 {
		return errors.New("options is empty")
	}

	bs, err := json.Marshal(options)
	if err != nil {
		return err
	}

	p.Options = options
	p.OptionsJson = string(bs)
	return nil
}

func (p *TaskTableRow) GetOptions() (*TaskTableRowOptionList, error) {
	if p.Options == nil {
		options := &TaskTableRowOptionList{}
		err := json.Unmarshal([]byte(p.OptionsJson), options)
		if err != nil {
			return nil, err
		}
		p.Options = options
	}
	return p.Options, nil
}
