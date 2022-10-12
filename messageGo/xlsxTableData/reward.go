package xlsxTable

import (
	"encoding/json"
	"time"
)

// 4010001,1,1,100;
// cid,品质(指定),数量(指定),权重
type RewardItem struct {
	Cid      int32 `json:"cid"`
	Quantity int32 `json:"quantity"`
	Num      int32 `json:"num"`
	Weight   int32 `json:"weight"`
}
type RewardItemList struct {
	TotalWeight int32        `json:"totalWeight"` //总权重
	Rewards     []RewardItem `json:"list"`
}

type RewardTime struct {
	Time   int32 `json:"time"`   // 次数
	Weight int32 `json:"weight"` // 权重
}

//奖励次数 0,9000;1,1000
type RewardTimeList struct {
	TotalWeight int32        `json:"totalWeight"` //总权重
	RewardTimes []RewardTime `json:"rewardTimes"`
}

type RewardTableRow struct {
	UId             uint      `gorm:"primaryKey;autoIncrement" json:"uid,string"`
	RewardId        int32     `json:"rewardId"`
	RewardTimesJson string    `gorm:"type:text" json:"rewardTimesJson"`
	RewardItemJson  string    `gorm:"type:text" json:"rewardItemJson"`
	CreatedAt       time.Time `json:"createdAt"` // 过期判断条件

	RewardList  *RewardItemList `gorm:"-" json:"-"`
	RewardTimes *RewardTimeList `gorm:"-" json:"-"`
}

func (this *RewardTableRow) SetRewardList(list RewardItemList) error {
	bs, err := json.Marshal(list)
	if err != nil {
		return err
	}
	this.RewardItemJson = string(bs)
	this.RewardList = &list
	return nil
}

func (this *RewardTableRow) GetRewardList() (*RewardItemList, error) {
	if len(this.RewardItemJson) < 2 {
		return nil, nil
	}

	list := &RewardItemList{}
	err := json.Unmarshal([]byte(this.RewardItemJson), list)
	if err != nil {
		return nil, err
	}
	this.RewardList = list
	return this.RewardList, err
}

func (this *RewardTableRow) SetRewardTimeList(list RewardTimeList) error {
	bs, err := json.Marshal(list)
	if err != nil {
		return err
	}
	this.RewardTimesJson = string(bs)
	this.RewardTimes = &list
	return nil
}

func (this *RewardTableRow) GetRewardTimeList() (*RewardTimeList, error) {
	if len(this.RewardTimesJson) < 2 {
		return nil, nil
	}

	list := &RewardTimeList{}
	err := json.Unmarshal([]byte(this.RewardTimesJson), list)
	if err != nil {
		return nil, err
	}
	this.RewardTimes = list
	return this.RewardTimes, err
}
