package xlsxTable

import (
	"encoding/json"
	"time"
)

// 掉落物配置.
type DropData struct {
	ObjectCid   int32 `json:"object_cid"`
	Num         int32 `json:"num"`
	Possibility int32 `json:"possibility"` // 宝箱触发万份比例
	Quality     int32 `json:"quality"`     // 品质  1~5
}

type DropList struct {
	List []DropData `json:"list"`
}

// 掉落物配方.
type DropTableRow struct {
	UId          uint      `gorm:"primaryKey;autoIncrement" json:"uid,string"`
	DropId       int32     `json:"dropId"`
	DropListJson string    `gorm:"type:text" json:"dropList"` // 宝箱列表
	CreatedAt    time.Time `json:"createdAt"`                 // 过期判断条件

	DropList *DropList `gorm:"-" json:"-"`
}

func (this *DropTableRow) SetDropList(list DropList) error {
	bs, err := json.Marshal(list)
	if err != nil {
		return err
	}
	this.DropListJson = string(bs)
	this.DropList = &list
	return nil
}

func (this *DropTableRow) GetDropList() (*DropList, error) {
	if len(this.DropListJson) < 2 {
		return nil, nil
	}

	list := &DropList{}
	err := json.Unmarshal([]byte(this.DropListJson), list)
	if err != nil {
		return nil, err
	}
	this.DropList = list
	return this.DropList, err
}
