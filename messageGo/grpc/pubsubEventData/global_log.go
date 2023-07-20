package pubsubEventData

import (
	base_data "game-message-core/grpc/baseData"
	"game-message-core/proto"
)

// 日志采集
type GlobalLogEvent struct {
	MsgVersion   int64                 `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	FormService  base_data.ServiceData `json:"formService"`
	Action       proto.GameLogType     `json:"action"`
	UserId       int64                 `json:"userId"`
	ObjectType   string                `json:"objectType"`
	ObjectId     int64                 `json:"objectId"`
	ObjectName   string                `json:"objectName"`
	ObjectOwner  int64                 `json:"objectOwner"`
	ObjectCid    uint32                `json:"objectCid"`
	ObjectNumber int32                 `json:"objectNumber"`
	Pos          base_data.GrpcVector3 `json:"pos"`
	Desc         string                `json:"desc"`
}

func (p *GlobalLogEvent) Clear() {
	p.MsgVersion = 0
	p.FormService.Clear()
	p.Action = 0
	p.UserId = 0
	p.ObjectType = ""
	p.ObjectId = 0
	p.ObjectName = ""
	p.ObjectOwner = 0
	p.ObjectCid = 0
	p.ObjectNumber = 0
	p.Pos.Clear()
	p.Desc = ""
}
