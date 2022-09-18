package methodData

import "game-message-core/proto"

// 向manager service 查询服务信息
type ManagerActionSelectServiceInput struct {
	MsgVersion  int64             `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	ServiceType proto.ServiceType `json:"serviceType"`
	MapId       int32             `json:"mapId"`
}

type ManagerActionSelectServiceOutput struct {
	ErrorCode    int32             `json:"errorCode"` //TODO: NEED USE GLOBAL ERROR CODE
	ErrorMessage string            `json:"errorMessage"`
	ServiceType  proto.ServiceType `json:"serviceType"`
	ServiceAppId string            `json:"serviceAppId"`
	MapId        int32             `json:"mapId"`
	Host         string            `json:"host"`
	Port         int32             `json:"port"`
	Online       int32             `json:"online"`
	MaxOnline    int32             `json:"maxOnline"`
	CreatedAt    int64             `json:"createdAt"`
	UpdateAt     int64             `json:"updatedAt"`
}
