package base_data

import "game-message-core/proto"

// service data
type ServiceData struct {
	AppId           string                    `json:"appId"`
	ServiceType     proto.ServiceType         `json:"serviceType"`
	SceneSerSubType proto.SceneServiceSubType `json:"sceneSerSubType"`
	Owner           int64                     `json:"owner"`
	Url             string                    `json:"url"`
	MapId           int32                     `json:"mapId"`
	Online          int32                     `json:"online"`
	MaxOnline       int32                     `json:"maxOnline"`
	CreatedAt       int64                     `json:"createdAt"`
	UpdatedAt       int64                     `json:"updatedAt"`
}

func (p *ServiceData) Clear() {
	p.AppId = ""
	p.ServiceType = 0
	p.SceneSerSubType = 0
	p.Owner = 0
	p.Url = ""
	p.MapId = 0
	p.Online = 0
	p.MaxOnline = 0
	p.CreatedAt = 0
	p.UpdatedAt = 0
}
