package methodData

import (
	base_data "game-message-core/grpc/baseData"
	"game-message-core/proto"
)

// 向manager service 查询服务信息
type ManagerActionSelectServiceInput struct {
	ServiceType     proto.ServiceType         `json:"serviceType"`
	SceneSerSubType proto.SceneServiceSubType `json:"sceneSerSubType"`
	OwnerId         int64                     `json:"ownerId"`
	MapId           int32                     `json:"mapId"`
}

func (p *ManagerActionSelectServiceInput) Clear() {
	p.ServiceType = 0
	p.SceneSerSubType = 0
	p.OwnerId = 0
	p.MapId = 0
}

type ManagerActionSelectServiceOutput struct {
	ErrorCode    int32                 `json:"errorCode"` //TODO: NEED USE GLOBAL ERROR CODE
	ErrorMessage string                `json:"errorMessage"`
	Service      base_data.ServiceData `json:"service"`
}

func (p *ManagerActionSelectServiceOutput) Clear() {
	p.ErrorCode = 0
	p.ErrorMessage = ""
	p.Service.Clear()
}

// 向manager service 查询服务信息
type MultiSelectServiceInput struct {
	ServiceType     proto.ServiceType         `json:"serviceType"`
	SceneSerSubType proto.SceneServiceSubType `json:"sceneSerSubType"`
	OwnerId         int64                     `json:"ownerId"`
	MapId           int32                     `json:"mapId"`
}

func (p *MultiSelectServiceInput) Clear() {
	p.ServiceType = 0
	p.SceneSerSubType = 0
	p.OwnerId = 0
	p.MapId = 0
}

type MultiSelectServiceOutput struct {
	ErrorCode    int32                   `json:"errorCode"` //TODO: NEED USE GLOBAL ERROR CODE
	ErrorMessage string                  `json:"errorMessage"`
	Services     []base_data.ServiceData `json:"services"`
}

func (p *MultiSelectServiceOutput) Clear() {
	p.ErrorCode = 0
	p.ErrorMessage = ""
	if p.Services != nil {
		p.Services = p.Services[:0]
	}
}
