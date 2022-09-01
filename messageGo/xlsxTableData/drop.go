package xlsxTable

import "time"

// 掉落物配置.
type DropData struct {
	ObjectCid   int32 `json:"object_cid"`
	Num         int32 `json:"num"`
	Possibility int32 `json:"possibility"` // 宝箱触发万份比例
	Quality     int32 `json:"quality"`     // 品质  1~5
}

// 掉落物配方.
type DropTable struct {
	UId       uint        `gorm:"primaryKey;autoIncrement" json:"uid,string"`
	DropId    int32       `json:"dropId"`
	DropList  []*DropData `json:"dropList"`  // 宝箱列表
	CreatedAt time.Time   `json:"createdAt"` // 过期判断条件
}
