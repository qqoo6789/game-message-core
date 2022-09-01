package xlsxTable

import "time"

type GameValueTable struct {
	UId           uint      `gorm:"primaryKey;autoIncrement" json:"uid,string"`
	Id            int32     `json:"id"`
	Value         int32     `json:"value"`
	StringValue   string    `json:"stringValue"`
	ValueArray    string    `json:"valueArray"`
	StringArray   string    `json:"stringArray"`
	IntValueArray []int32   `json:"intValueArray"`
	CreatedAt     time.Time `json:"createdAt"` // 过期判断条件
}
