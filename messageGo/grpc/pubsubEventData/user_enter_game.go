package pubsubEventData

type UserEnterGameEvent struct {
	MsgVersion        int64   `json:"msgVersion"`        // 消息版本号 值为毫秒时间戳
	SceneServiceAppId string  `json:"sceneServiceAppId"` //
	AgentAppId        string  `json:"agentAppId"`
	UserSocketId      string  `json:"userSocketId"`
	UserId            int64   `json:"userId"`
	Name              string  `json:"name"`
	RoleId            int32   `json:"roleId"`
	Gender            string  `json:"gender"`
	RoleIcon          string  `json:"roleIcon"`
	Guide             bool    `json:"guide"`
	Eyebrow           int32   `json:"eyebrow"`
	Mouth             int32   `json:"mouth"`
	Eye               int32   `json:"eye"`
	Face              int32   `json:"face"`
	Hair              int32   `json:"hair"`
	Glove             int32   `json:"glove"`
	Clothes           int32   `json:"clothes"`
	Pants             int32   `json:"pants"`
	MapId             int32   `json:"mapId"`
	X                 float32 `json:"x"`
	Y                 float32 `json:"y"`
	Z                 float32 `json:"z"`
	DirX              float32 `json:"dirX"`
	DirY              float32 `json:"dirY"`
	DirZ              float32 `json:"dirZ"`
}
