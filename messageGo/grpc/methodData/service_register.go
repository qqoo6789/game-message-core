package methodData

import (
	base_data "game-message-core/grpc/baseData"
)

// 服务信息用于服务 注册和释放
type ServiceDataInput struct {
	MsgVersion int64                 `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	Service    base_data.ServiceData `json:"service"`    //
}

// 服务信息用于服务 注册和释放
type ServiceDataOutput struct {
	MsgVersion int64 `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	Success    bool  `json:"success"`
}
