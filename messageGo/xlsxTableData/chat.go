package xlsxTable

import "time"

// 频道类型
type ChatChannelType int32

const (
	ChatChannelTypeUnknown ChatChannelType = 0
	// 系统
	ChatChannelTypeSystem ChatChannelType = 1
	// 世界
	ChatChannelTypeWorld ChatChannelType = 2
	// 附近
	ChatChannelTypeNearby ChatChannelType = 3
	// 私聊
	ChatChannelTypePrivate ChatChannelType = 4
)

type ChatTableRow struct {
	UId       uint            `gorm:"primaryKey;autoIncrement" json:"uid,string"`
	ChatType  ChatChannelType `json:"chatType"`
	Cd        int32           `json:"cd"`
	CreatedAt time.Time       `json:"createdAt"` // 过期判断条件
}
