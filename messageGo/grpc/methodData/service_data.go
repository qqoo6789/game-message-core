package methodData

import "game-message-core/proto"

// 服务信息用于服务 注册和释放
type ServiceData struct {
	MsgVersion  int64             `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	Id          int64             `json:"id"`
	Name        string            `json:"name"`
	AppId       string            `json:"appId"`
	ServiceType proto.ServiceType `json:"serviceType"`
	Host        string            `json:"host"`
	Port        int32             `json:"port"`
	MapId       int32             `json:"mapId"`
	Online      int32             `json:"online"`
	MaxOnline   int32             `json:"maxOnline"`
	CreatedAt   int64             `json:"createdAt"`
	UpdatedAt   int64             `json:"updatedAt"`
}
