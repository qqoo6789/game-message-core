using System;
using UnityEngine;


[Serializable]
public class EntityInfo
{
    public GameMessageCore.EntityType EntityType;
    public int EntityCid;
    public string EntityName;
    public long EntityId;
}

// 对应 proto.EntityProfile
[Serializable]
public class GrpcAttributeData
{
    public int Id;
    public int Value;
    public GameMessageCore.AttributeDisplayType DisplayType;

    public void Set(GameMessageCore.AttributeData attr)
    {
        if (attr == null)
        {
            return;
        }

        Id = attr.Id;
        Value = attr.Value;
        DisplayType = attr.DisplayType;
    }

    public GameMessageCore.AttributeData ToProtoData()
    {
        return new GameMessageCore.AttributeData()
        {
            Id = Id,
            Value = Value,
            DisplayType = DisplayType,
        };
    }
}


// 对应 proto.NftBuild
[Serializable]
public class GrpcNftBuild
{
    public long Id;
    public int Cid;
    public string FromNft;
    public long Owner;
    public int[] LandIds;
    public GrpcVector3 Position;
    public GrpcVector3 Dir;

    public int ElectricEnd;
    public int HarvestStartAt;
    public int HarvestAt;
    public int HarvestItemCount;
    public int CollectionStartAt;
    public int CollectionAt;
    public int CollectionItemCount;
    ///电量不足时建造保护期开始时间
    public int LandPlacementPowerZeroCooldownStartAt;
    //电量不足时建造保护期
    public int LandPlacementPowerZeroCooldownAt;

    public void Set(GameMessageCore.NftBuild build)
    {
        if (build == null)
        {
            return;
        }

        Id = build.Id;
        Cid = build.Cid;
        FromNft = build.FromNft;
        Owner = build.Owner;

        ElectricEnd = build.ElectricEnd;
        HarvestStartAt = build.HarvestStartAt;
        HarvestAt = build.HarvestAt;
        HarvestItemCount = build.HarvestItemCount;
        CollectionStartAt = build.CollectionStartAt;
        CollectionAt = build.CollectionAt;
        CollectionItemCount = build.CollectionItemCount;
        LandPlacementPowerZeroCooldownStartAt = build.LandPlacementPowerZeroCooldownStartAt;
        LandPlacementPowerZeroCooldownAt = build.LandPlacementPowerZeroCooldownAt;
        Position.Set(build.Position);
        Dir.Set(build.Dir);

        if (build.LandIds != null)
        {
            LandIds = new int[build.LandIds.Count];
            for (int i = 0; i < build.LandIds.Count; i++)
            {
                LandIds[i] = build.LandIds[i];
            }
        }
    }

    public GameMessageCore.NftBuild ToProtoData()
    {
        GameMessageCore.NftBuild build = new()
        {
            Id = Id,
            Cid = Cid,
            FromNft = FromNft,
            Owner = Owner,
            Position = Position.ToProtoData(),
            Dir = Dir.ToProtoData(),
            ElectricEnd = ElectricEnd,
            HarvestStartAt = HarvestStartAt,
            HarvestAt = HarvestAt,
            HarvestItemCount = HarvestItemCount,
            CollectionStartAt = CollectionStartAt,
            CollectionAt = CollectionAt,
            CollectionItemCount = CollectionItemCount,
            LandPlacementPowerZeroCooldownStartAt = LandPlacementPowerZeroCooldownStartAt,
            LandPlacementPowerZeroCooldownAt = LandPlacementPowerZeroCooldownAt,
        };
        foreach (int landId in LandIds)
        {
            build.LandIds.Add(landId);
        }
        return build;
    }
}

