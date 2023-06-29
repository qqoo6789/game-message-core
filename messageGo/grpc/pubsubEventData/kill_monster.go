package pubsubEventData

import base_data "game-message-core/grpc/baseData"

type KillMonsterEventData struct {
	MsgVersion        int64                        `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	SceneServiceAppId string                       `json:"sceneServiceAppId"`
	MapId             int32                        `json:"mapId"`
	OwnerId           int64                        `json:"ownerId"`
	KillerId          int64                        `json:"killerId"`
	PosX              float32                      `json:"posX"`
	PosY              float32                      `json:"posY"`
	PosZ              float32                      `json:"posZ"`
	MonsterCid        int32                        `json:"monsterCid"`
	MonsterName       string                       `json:"monsterName"`
	DropList          []base_data.GrpcItemBaseInfo `json:"dropList"`
}

func (p *KillMonsterEventData) Clear() {
	p.MsgVersion = 0
	p.SceneServiceAppId = ""
	p.MapId = 0
	p.OwnerId = 0
	p.KillerId = 0
	p.PosX = 0
	p.PosY = 0
	p.PosZ = 0
	p.MonsterCid = 0
	p.MonsterName = ""
	if p.DropList != nil {
		p.DropList = p.DropList[:0]
	}
}
