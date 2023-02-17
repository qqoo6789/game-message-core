package xlsxTable

import (
	"time"
)

type EntityAttributeRow struct {
	Id           uint32    `gorm:"primaryKey" json:"id"`
	Name         string    `json:"name"`         // 属性名
	ProfileType  uint32    `json:"profileType"`  // 天赋类型
	ValueType    uint32    `json:"valueType"`    // 属性值类型
	DefaultValue uint32    `json:"defaultValue"` // 默认值
	CreatedAt    time.Time `json:"createdAt"`    // 过期判断条件
}
