package base_data

import (
	"game-message-core/proto"
)

// 对应 proto.EntityProfile
type GrpcAttributeData struct {
	Id          int32                      `json:"id"`
	Value       int32                      `json:"value"`
	DisplayType proto.AttributeDisplayType `json:"displayType"`
}

func (p *GrpcAttributeData) Set(attr *proto.AttributeData) {
	if attr == nil {
		return
	}
	p.Id = attr.Id
	p.Value = attr.Value
}

func (p *GrpcAttributeData) ToProtoData() *proto.AttributeData {
	return &proto.AttributeData{
		Id:    p.Id,
		Value: p.Value,
	}
}

// 对应 proto.NftBuild
type GrpcNftBuild struct {
	Id       int64       `json:"id"`
	Cid      int32       `json:"cid"`
	FromNft  string      `json:"fromNft"`
	Owner    int64       `json:"owner"`
	LandIds  []int32     `json:"landIds"`
	Position GrpcVector3 `json:"position"`
	Dir      GrpcVector3 `json:"dir"`
	// 电量过期时间 单位秒
	ElectricEnd int32 `json:"electricEnd"`
	// 产出开始时间点.
	HarvestStartAt int32 `json:"harvestStartAt"`
	// 可以收集的时间
	HarvestAt int32 `json:"harvestAt"`
	// 可收获的物品数量统计(没电时转移到采集) 单位秒
	HarvestItemCount int32 `json:"harvestItemCount"`
	// 采集开始时间点
	CollectionStartAt int32 `json:"collectionStartAt"`
	// 下次可采集(偷取)的时间戳 单位秒
	CollectionAt int32 `json:"collectionAt"`
	// 可采集(偷取)物品数量统计
	CollectionItemCount int32 `json:"collectionItemCount"`
	// 电量不足时建造保护期开始时间
	LandPlacementPowerZeroCooldownStartAt int32 `json:"landPlacementPowerZeroCooldownStartAt"`
	// 电量不足时建造保护期
	LandPlacementPowerZeroCooldownAt int32 `json:"landPlacementPowerZeroCooldown"`
}

func (p *GrpcNftBuild) Set(build *proto.NftBuild) {
	if build == nil {
		return
	}
	p.Id = build.Id
	p.Cid = build.Cid
	p.FromNft = build.FromNft
	p.Owner = build.Owner
	p.LandIds = build.LandIds
	p.ElectricEnd = build.ElectricEnd
	p.HarvestStartAt = build.HarvestStartAt
	p.HarvestAt = build.HarvestAt
	p.HarvestItemCount = build.HarvestItemCount
	p.CollectionStartAt = build.CollectionStartAt
	p.CollectionAt = build.CollectionAt
	p.CollectionItemCount = build.CollectionItemCount
	p.LandPlacementPowerZeroCooldownStartAt = build.LandPlacementPowerZeroCooldownStartAt
	p.LandPlacementPowerZeroCooldownAt = build.LandPlacementPowerZeroCooldownAt
	p.Position.Set(build.Position)
	p.Dir.Set(build.Dir)
}

func (p *GrpcNftBuild) ToProtoData() *proto.NftBuild {
	pbBuild := &proto.NftBuild{
		Id:                                    p.Id,
		Cid:                                   p.Cid,
		FromNft:                               p.FromNft,
		Owner:                                 p.Owner,
		LandIds:                               p.LandIds,
		ElectricEnd:                           p.ElectricEnd,
		HarvestStartAt:                        p.HarvestStartAt,
		HarvestAt:                             p.HarvestAt,
		HarvestItemCount:                      p.HarvestItemCount,
		CollectionStartAt:                     p.CollectionStartAt,
		CollectionAt:                          p.CollectionAt,
		CollectionItemCount:                   p.CollectionItemCount,
		LandPlacementPowerZeroCooldownStartAt: p.LandPlacementPowerZeroCooldownStartAt,
		LandPlacementPowerZeroCooldownAt:      p.LandPlacementPowerZeroCooldownAt,
	}

	pbPos := p.Position.ToProtoData()
	pbDir := p.Dir.ToProtoData()
	pbBuild.Position = &pbPos
	pbBuild.Dir = &pbDir
	return pbBuild
}
