package xlsxTable

import (
	"encoding/json"
	"fmt"
	"time"
)

type TaskObject struct {
	Param1 int32 `json:"param1"`
	Param2 int32 `json:"param2"`
	Param3 int32 `json:"param3"`
}
type TaskObjectList struct {
	ChanceSum int32        `json:"chanceSum"`
	ParamList []TaskObject `json:"paramList"`
}
type TaskTableRow struct {
	UId             uint   `gorm:"primaryKey;autoIncrement" json:"uid,string"`
	Level           int32  `json:"level"`
	Name            string `json:"name"`
	SubSystem       string `json:"subSystem"`
	Kind            int32  `json:"kind"`
	NeedItemJson    string `gorm:"type:text" json:"needItemJson"`
	UseItemJson     string `gorm:"type:text" json:"useItemJson"`
	KillMonsterJson string `gorm:"type:text" json:"killMonsterJson"`
	TargetPosJson   string `gorm:"type:text" json:"targetPosJson"`
	QuizJson        string `gorm:"type:text" json:"quizJson"`
	RequestLand     int32  `json:"requestLand"`
	RewardItemsJson string `gorm:"type:text" json:"rewardItemsJson"`
	RewardExp       int32  `json:"rewardExp"`
	Difficulty      int32  `json:"difficulty"`

	CreatedAt time.Time `json:"createdAt"` // 过期判断条件,

	NeedItem    *TaskObjectList `gorm:"-" json:"-"`
	UseItem     *TaskObjectList `gorm:"-" json:"-"`
	KillMonster *TaskObjectList `gorm:"-" json:"-"`
	TargetPos   *TaskObjectList `gorm:"-" json:"-"`
	Quiz        *TaskObjectList `gorm:"-" json:"-"`
	RewardItems *TaskObjectList `gorm:"-" json:"-"`
}

func marshal(objs *TaskObjectList) ([]byte, error) {
	if objs == nil {
		return nil, fmt.Errorf("TaskObjectList is nil")
	}
	return json.Marshal(objs)
}
func unMarshal(data string) (*TaskObjectList, error) {
	if len(data) < 2 {
		return nil, nil
	}

	objs := &TaskObjectList{}
	if err := json.Unmarshal([]byte(data), objs); err != nil {
		return nil, err
	}
	return objs, nil
}

func (p *TaskTableRow) SetNeedItem(objs *TaskObjectList) error {
	bs, err := marshal(objs)
	p.NeedItemJson = string(bs)
	return err
}
func (p *TaskTableRow) GetNeedItem() *TaskObjectList {
	if p.NeedItem == nil {
		p.NeedItem, _ = unMarshal(p.NeedItemJson)
	}
	return p.NeedItem
}

func (p *TaskTableRow) SetUseItem(objs *TaskObjectList) error {
	bs, err := marshal(objs)
	p.UseItemJson = string(bs)
	return err
}
func (p *TaskTableRow) GetUseItem() *TaskObjectList {
	if p.UseItem == nil {
		p.UseItem, _ = unMarshal(p.UseItemJson)
	}
	return p.UseItem
}

func (p *TaskTableRow) SetKillMonster(objs *TaskObjectList) error {
	bs, err := marshal(objs)
	p.KillMonsterJson = string(bs)
	return err
}
func (p *TaskTableRow) GetKillMonster() *TaskObjectList {
	if p.KillMonster == nil {
		p.KillMonster, _ = unMarshal(p.KillMonsterJson)
	}
	return p.KillMonster
}

func (p *TaskTableRow) SetTargetPos(objs *TaskObjectList) error {
	bs, err := marshal(objs)
	p.TargetPosJson = string(bs)
	return err
}
func (p *TaskTableRow) GetTargetPos() *TaskObjectList {
	if p.TargetPos == nil {
		p.TargetPos, _ = unMarshal(p.TargetPosJson)
	}
	return p.TargetPos
}

func (p *TaskTableRow) SetQuiz(objs *TaskObjectList) error {
	bs, err := marshal(objs)
	p.QuizJson = string(bs)
	return err
}
func (p *TaskTableRow) GetQuiz() *TaskObjectList {
	if p.Quiz == nil {
		p.Quiz, _ = unMarshal(p.QuizJson)
	}
	return p.Quiz
}

func (p *TaskTableRow) SetRewardItems(objs *TaskObjectList) error {
	bs, err := marshal(objs)
	p.RewardItemsJson = string(bs)
	return err
}
func (p *TaskTableRow) GetRewardItems() *TaskObjectList {
	if p.RewardItems == nil {
		p.RewardItems, _ = unMarshal(p.RewardItemsJson)
	}
	return p.RewardItems
}
