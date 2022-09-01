package xlsxTable

import "time"

type ItemTable struct {
	UId       uint      `gorm:"primaryKey;autoIncrement" json:"uid,string"`
	ItemCid   int32     `json:"itemCid"`
	Name      string    `json:"name"`
	ItemType  int32     `json:"itemType"`
	UserType  int32     `json:"userType"`
	CanMint   bool      `json:"canMint"`
	CreatedAt time.Time `json:"createdAt"` // 过期判断条件
}
