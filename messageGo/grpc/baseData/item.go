package base_data

import "game-message-core/proto"

type GrpcAvatarAttribute struct {
	Position proto.AvatarPosition `json:"position"`
	// 稀有度 unique,  mythic, epic, rare, common
	Rarity string `json:"rarity"`
	// cid
	ObjectId int32 `json:"objectId"`
	// 属性增量
	Data []GrpcAttributeData `json:"data"`
}

func (p *GrpcAvatarAttribute) Clear() {
	p.Position = 0
	p.Rarity = ""
	p.ObjectId = 0
	if p.Data != nil {
		p.Data = p.Data[:0]
	}
}

func (p *GrpcAvatarAttribute) Set(attr *proto.AvatarAttribute) {
	if attr == nil {
		return
	}
	p.Rarity = attr.Rarity
	p.Data = make([]GrpcAttributeData, 0, len(attr.Data))
	for _, data := range attr.Data {
		a := GrpcAttributeData{}
		a.Set(data)
		p.Data = append(p.Data, a)
	}
}
func (p *GrpcAvatarAttribute) ToProtoData() *proto.AvatarAttribute {
	attribute := &proto.AvatarAttribute{
		Rarity:   p.Rarity,
		ObjectId: p.ObjectId,
	}
	for _, data := range p.Data {
		attribute.Data = append(attribute.Data, data.ToProtoData())
	}
	return attribute
}

type GrpcItemBaseInfo struct {
	Cid     int32 `json:"cid"`
	Num     int32 `json:"num"`
	Quality int32 `json:"quality"`
}

func (p *GrpcItemBaseInfo) Clear() {
	p.Cid = 0
	p.Num = 0
	p.Quality = 1
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
	ConsumableType proto.NFTConsumableType `json:"consumableType"`
	Value          int32                   `json:"value"`
}

func (p *GrpcNFTConsumableInfo) Clear() {
	p.Quality = ""
	p.ConsumableType = 0
	p.Value = 0
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
