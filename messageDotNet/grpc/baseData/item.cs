using System;
using UnityEngine;


[Serializable]
public class GrpcAvatarAttribute
{
    // 稀有度 unique,  mythic, epic, rare, common
    public string Rarity;
    // 耐久度
    public int Durability;
    // 属性增量
    public GrpcAttributeData[] Data;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }

    public void Set(GameMessageCore.AvatarAttribute attr)
    {
        if (attr == null)
        {
            return;
        }

        Rarity = attr.Rarity;
        Durability = attr.Durability;
        Data = new GrpcAttributeData[attr.Data.Count];
        for (int i = 0; i < attr.Data.Count; i++)
        {
            GrpcAttributeData grpcAt = new();
            grpcAt.Set(attr.Data[i]);
            Data[i] = grpcAt;
        }
    }

    public GameMessageCore.AvatarAttribute ToProtoData()
    {
        GameMessageCore.AvatarAttribute attribute = new()
        {
            Rarity = Rarity,
            Durability = Durability,
        };
        foreach (GrpcAttributeData at in Data)
        {
            attribute.Data.Add(at.ToProtoData());
        }
        return attribute;
    }

}

// 对应 proto.PlayerAvatar
[Serializable]
public class GrpcPlayerAvatar
{
    // 装备位置
    public GameMessageCore.AvatarPosition Position;
    // 装备的物品 id
    public int ObjectId;
    // 属性
    public GrpcAvatarAttribute Attribute;


    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }


    public void Set(GameMessageCore.PlayerAvatar pa)
    {
        if (pa == null)
        {
            return;
        }

        Position = pa.Position;
        ObjectId = pa.ObjectId;
        Attribute = new GrpcAvatarAttribute();
        Attribute.Set(pa.Attribute);
    }

    public GameMessageCore.PlayerAvatar ToProtoData()
    {

        return new GameMessageCore.PlayerAvatar()
        {
            Position = Position,
            ObjectId = ObjectId,
            Attribute = Attribute.ToProtoData(),
        };
    }

}


[Serializable]
public class GrpcItemBaseInfo
{

    public int Cid;
    public int Num;
    public int Quality;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }

    public void Set(GameMessageCore.ItemBaseInfo info)
    {
        if (info == null)
        {
            return;
        }
        Cid = info.Cid;
        Num = info.Num;
        Quality = info.Quality;
    }
    public GameMessageCore.ItemBaseInfo ToProtoData()
    {
        return new GameMessageCore.ItemBaseInfo()
        {
            Cid = Cid,
            Num = Num,
            Quality = Quality,
        };
    }

}

// 消耗品
[Serializable]
public class GrpcNFTConsumableInfo
{
    public string Quality;
    public GameMessageCore.NFTConsumableType ConsumableType;
    public int Value;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }

    public void Set(GameMessageCore.NFTConsumableInfo info)
    {
        if (info == null)
        {
            return;
        }

        Quality = info.Quality;
        ConsumableType = info.ConsumableType;
        Value = info.Value;
    }

    public GameMessageCore.NFTConsumableInfo ToProtoData()
    {
        return new GameMessageCore.NFTConsumableInfo()
        {
            Quality = Quality,
            ConsumableType = ConsumableType,
            Value = Value,
        };
    }

}
