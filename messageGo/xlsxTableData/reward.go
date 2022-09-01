package xlsxTable

import "time"

//材料
type RewardItem struct {
	Cid      int32 `json:"cid"`
	Quantity int32 `json:"quantity"`
	RandAttr bool  `json:"randAttr"`
}

type RewardTable struct {
	UId         uint          `gorm:"primaryKey;autoIncrement" json:"uid,string"`
	RewardId    int32         `json:"rewardId"`
	Exp         int32         `json:"exp"`
	RewardItems []*RewardItem `json:"rewardItems"`
	CreatedAt   time.Time     `json:"createdAt"` // 过期判断条件
}
