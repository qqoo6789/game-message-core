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

func (p *UserEnterGameEvent) Clear() {
	p.MsgVersion = 0
	p.SceneServiceAppId = ""
	p.AgentAppId = ""
	p.UserSocketId = ""
	p.UserId = 0
	p.Name = ""
	p.RoleId = 0
	p.Gender = ""
	p.RoleIcon = ""
	p.Guide = false
	p.Eyebrow = 0
	p.Mouth = 0
	p.Eye = 0
	p.Face = 0
	p.Hair = 0
	p.Glove = 0
	p.Clothes = 0
	p.Pants = 0
	p.MapId = 0
	p.X = 0
	p.Y = 0
	p.Z = 0
	p.DirX = 0
	p.DirY = 0
	p.DirZ = 0
}
