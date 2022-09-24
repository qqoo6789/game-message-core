package base_data

import "game-message-core/proto"

// 属性字段增量数据 1v1
type GrpcAttributeData struct {
	// type
	Type proto.AttributeType `json:"type"`
	// 增量数值
	Value int32 `json:"value"`
}

func (p *GrpcAttributeData) Set(attr *proto.AttributeData) {
	if attr == nil {
		return
	}
	p.Type = attr.Type
	p.Value = attr.Value
}
func (p *GrpcAttributeData) ToProtoData() *proto.AttributeData {
	return &proto.AttributeData{
		Type:  p.Type,
		Value: p.Value,
	}
}

type GrpcAvatarAttribute struct {
	// 稀有度 unique,  mythic, epic, rare, common
	Rarity string `json:"rarity"`
	// 耐久度
	Durability int32 `json:"durability"`
	// 属性增量
	Data []GrpcAttributeData `json:"data"`
}

func (p *GrpcAvatarAttribute) Set(attr *proto.AvatarAttribute) {
	if attr == nil {
		return
	}
	p.Rarity = attr.Rarity
	p.Durability = attr.Durability
	p.Data = make([]GrpcAttributeData, 0, len(attr.Data))
	for _, data := range attr.Data {
		a := GrpcAttributeData{}
		a.Set(data)
		p.Data = append(p.Data, a)
	}
}
func (p *GrpcAvatarAttribute) ToProtoData() *proto.AvatarAttribute {
	attribute := &proto.AvatarAttribute{
		Rarity:     p.Rarity,
		Durability: p.Durability,
	}
	for _, data := range p.Data {
		attribute.Data = append(attribute.Data, data.ToProtoData())
	}
	return attribute
}

// 对应 proto.PlayerAvatar
type GrpcPlayerAvatar struct {
	// 装备位置
	Position proto.AvatarPosition `json:"position"`
	// 装备的物品 id
	ObjectId int32 ` json:"object_id"`
	// 属性
	Attribute *GrpcAvatarAttribute `json:"attribute"`
}

func (p *GrpcPlayerAvatar) Set(pa *proto.PlayerAvatar) {
	if pa == nil {
		return
	}

	p.Position = pa.Position
	p.ObjectId = pa.ObjectId
	p.Attribute = &GrpcAvatarAttribute{}
	p.Attribute.Set(pa.Attribute)
}
func (p *GrpcPlayerAvatar) ToProtoData() *proto.PlayerAvatar {
	pa := &proto.PlayerAvatar{
		Position:  p.Position,
		ObjectId:  p.ObjectId,
		Attribute: p.Attribute.ToProtoData(),
	}
	return pa
}

type GrpcItemBaseInfo struct {
	Cid     int32 `json:"cid"`
	Num     int32 `json:"num"`
	Quality int32 `json:"quality"`
}

func (p *GrpcItemBaseInfo) Set(info *proto.ItemBaseInfo) {
	if info == nil {
		return
	}
	p.Cid = info.Cid
	p.Num = info.Num
	p.Quality = info.Quality
}
func (p *GrpcItemBaseInfo) ToProtoData() *proto.ItemBaseInfo {
	pa := &proto.ItemBaseInfo{
		Cid:     p.Cid,
		Num:     p.Num,
		Quality: p.Quality,
	}
	return pa
}

// 消耗品
type GrpcNFTConsumableInfo struct {
	Quality        string                  `json:"quality"`
	ConsumableType proto.NFTConsumableType `json:"consumable"`
	Value          int32                   `json:"value"`
}

func (p *GrpcNFTConsumableInfo) Set(info *proto.NFTConsumableInfo) {
	if info == nil {
		return
	}
	p.Quality = info.Quality
	p.ConsumableType = info.ConsumableType
	p.Value = info.Value
}
func (p *GrpcNFTConsumableInfo) ToProtoData() *proto.NFTConsumableInfo {
	return &proto.NFTConsumableInfo{
		Quality:        p.Quality,
		ConsumableType: p.ConsumableType,
		Value:          p.Value,
	}
}
