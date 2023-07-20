package base_data

import (
	"encoding/json"
	"game-message-core/proto"
)

// 对应 proto.Vector3
type GrpcVector3 struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
}

func (p *GrpcVector3) Clear() {
	p.X = 0
	p.Y = 0
	p.Z = 0
}

func (p *GrpcVector3) Set(v *proto.Vector3) {
	if v == nil {
		return
	}
	p.X = v.X
	p.Y = v.Y
	p.Z = v.Z
}

func (p *GrpcVector3) ToJson() string {
	if bs, err := json.Marshal(p); err == nil {
		return string(bs)
	}
	return ""
}

func (p *GrpcVector3) ToProtoData() proto.Vector3 {
	return proto.Vector3{X: p.X, Y: p.Y, Z: p.Z}
}
