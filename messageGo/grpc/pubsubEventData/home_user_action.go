package pubsubEventData

import base_data "game-message-core/grpc/baseData"

// 家园播种事件
type HomeActionSowingEvent struct {
	MsgVersion int64  `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	UserId     int64  `json:"userId"`
	SeedCid    uint32 `json:"seedCid"`
	SoilId     int64  `json:"soilId"`
}

func (p *HomeActionSowingEvent) Clear() {
	p.MsgVersion = 0
	p.UserId = 0
	p.SeedCid = 0
	p.SoilId = 0
}

// 家园放置动物食物事件
type HomeActionPutAnimalFoodEvent struct {
	MsgVersion int64  `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	UserId     int64  `json:"userId"`
	FoodCid    uint32 `json:"foodCid"`
	Num        uint32 `json:"num"`
}

func (p *HomeActionPutAnimalFoodEvent) Clear() {
	p.MsgVersion = 0
	p.UserId = 0
	p.FoodCid = 0
	p.Num = 0
}

type HomeHarvestInfo struct {
	SoilId int64                        `json:"soilId"`
	Items  []base_data.GrpcItemBaseInfo `json:"items"`
}

func (p *HomeHarvestInfo) Clear() {
	p.SoilId = 0
	for i := 0; i < len(p.Items); i++ {
		p.Items[i].Clear()
	}
	if len(p.Items) > 0 {
		p.Items = p.Items[:0]
	}
}

// 家园收获土地产出事件
type HomeHarvestEvent struct {
	MsgVersion int64             `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	UserId     int64             `json:"userId"`
	Harvests   []HomeHarvestInfo `json:"harvests"`
}

func (p *HomeHarvestEvent) Clear() {
	p.MsgVersion = 0
	p.UserId = 0
	for i := 0; i < len(p.Harvests); i++ {
		p.Harvests[i].Clear()
	}
	if len(p.Harvests) > 0 {
		p.Harvests = p.Harvests[:0]
	}
}
