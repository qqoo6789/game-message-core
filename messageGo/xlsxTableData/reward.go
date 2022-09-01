package xlsxTable

import (
	"encoding/json"
	"time"
)

//材料
type RewardItem struct {
	Cid      int32 `json:"cid"`
	Quantity int32 `json:"quantity"`
	RandAttr bool  `json:"randAttr"`
}
type RewardItemList struct {
	List []RewardItem `json:"list"`
}
type RewardTableRow struct {
	UId        uint      `gorm:"primaryKey;autoIncrement" json:"uid,string"`
	RewardId   int32     `json:"rewardId"`
	Exp        int32     `json:"exp"`
	RewardJson string    `gorm:"type:text" json:"rewardJson"`
	CreatedAt  time.Time `json:"createdAt"` // 过期判断条件

	RewardList *RewardItemList `gorm:"-" json:"-"`
}

func (this *RewardTableRow) SetRewardList(list RewardItemList) error {
	bs, err := json.Marshal(list)
	if err != nil {
		return err
	}
	this.RewardJson = string(bs)
	this.RewardList = &list
	return nil
}

func (this *RewardTableRow) GetRewardList() (*RewardItemList, error) {
	if len(this.RewardJson) < 2 {
		return nil, nil
	}

	list := &RewardItemList{}
	err := json.Unmarshal([]byte(this.RewardJson), list)
	if err != nil {
		return nil, err
	}
	this.RewardList = list
	return this.RewardList, err
}
