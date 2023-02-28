package xlsxTable

import "time"

type ItemEatableRow struct {
	Id           int32     `gorm:"primaryKey;autoIncrement" json:"id"`
	Cd           int32     `json:"cd"`
	CdType       int32     `json:"cdType"`
	InteractType string    `json:"interactType"`
	CreatedAt    time.Time `json:"createdAt"` // 过期判断条件
}
