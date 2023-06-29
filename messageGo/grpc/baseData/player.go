package base_data

import "game-message-core/proto"

type GrpcPlayerFeature struct {
	Eyebrow int32 `json:"eyebrow"`
	Mouth   int32 `json:"mouth"`
	Eye     int32 `json:"eye"`
	Face    int32 `json:"face"`
	Hair    int32 `json:"hair"`
	Glove   int32 `json:"glove"`
	Clothes int32 `json:"clothes"`
	Pants   int32 `json:"pants"`
	Skin    int32 `json:"skin"`
}

func (p *GrpcPlayerFeature) Clear() {
	p.Eyebrow = 0
	p.Mouth = 0
	p.Eye = 0
	p.Face = 0
	p.Hair = 0
	p.Glove = 0
	p.Clothes = 0
	p.Pants = 0
	p.Skin = 0
}

func (p *GrpcPlayerFeature) Set(feature *proto.PlayerFeature) {
	if feature == nil {
		return
	}
	p.Eyebrow = feature.Eyebrow
	p.Mouth = feature.Mouth
	p.Eye = feature.Eye
	p.Face = feature.Face
	p.Hair = feature.Hair
	p.Glove = feature.Glove
	p.Clothes = feature.Clothes
	p.Pants = feature.Pants
	p.Skin = feature.Skin
}
func (p *GrpcPlayerFeature) ToProtoData() *proto.PlayerFeature {
	return &proto.PlayerFeature{
		Eyebrow: p.Eyebrow,
		Mouth:   p.Mouth,
		Eye:     p.Eye,
		Face:    p.Face,
		Hair:    p.Hair,
		Glove:   p.Glove,
		Clothes: p.Clothes,
		Pants:   p.Pants,
		Skin:    p.Skin,
	}
}

// 对应 proto.PlayerBaseData
type GrpcPlayerBaseData struct {
	UserId   int64             `json:"userId"`
	Name     string            `json:"name"`
	RoleId   int32             `json:"roleId"`
	Gender   string            `json:"gender"`
	RoleIcon string            `json:"roleIcon"`
	Feature  GrpcPlayerFeature `json:"feature"`
	Guide    bool              `json:"guide"`
}

func (p *GrpcPlayerBaseData) Clear() {
	p.UserId = 0
	p.Name = ""
	p.RoleId = 0
	p.Gender = ""
	p.RoleIcon = ""
	p.Guide = false
	p.Feature.Clear()
}

func (p *GrpcPlayerBaseData) Set(baseData *proto.PlayerBaseData) {
	if baseData == nil {
		return
	}
	p.UserId = baseData.UserId
	p.Name = baseData.Name
	p.RoleId = baseData.RoleId
	p.Gender = baseData.Gender
	p.RoleIcon = baseData.RoleIcon
	p.Guide = baseData.Guide
	p.Feature.Set(baseData.Feature)
}
func (p *GrpcPlayerBaseData) ToProtoData() *proto.PlayerBaseData {
	pbData := &proto.PlayerBaseData{
		UserId:   p.UserId,
		Name:     p.Name,
		RoleId:   p.RoleId,
		Gender:   p.Gender,
		RoleIcon: p.RoleIcon,
		Feature:  p.Feature.ToProtoData(),
		Guide:    p.Guide,
	}

	return pbData
}
