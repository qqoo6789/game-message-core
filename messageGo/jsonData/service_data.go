package jsonData

import "game-message-core/proto"

type ServiceData struct {
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
