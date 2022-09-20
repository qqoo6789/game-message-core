package pubsubEventData

import "game-message-core/proto"

type KillMonsterEventData struct {
	MsgVersion        int64                `json:"msgVersion"` // 消息版本号 值为毫秒时间戳
	SceneServiceAppId string               `json:"sceneServiceAppId"`
	MapId             int32                `json:"mapId"`
	UserId            int64                `json:"userId"`
	PosX              float32              `json:"posX"`
	PosY              float32              `json:"posY"`
	PosZ              float32              `json:"posZ"`
	MonsterCid        int32                `json:"monsterCid"`
	MonsterName       string               `json:"monsterName"`
	DropList          []proto.ItemBaseInfo `json:"dropList"`
	Exp               int32                `json:"exp"`
}
