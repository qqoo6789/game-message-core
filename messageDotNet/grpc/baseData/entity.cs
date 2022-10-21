using System;
using UnityEngine;

// 对应 proto.EntityProfile
[Serializable]
public class GrpcEntityProfile
{
    public int Lv;
    public long Exp;
    public int Att;
    public int AttSpeed;
    public int Def;
    public int HpCurrent;
    public int HpLimit;
    public int CritRate;
    public int CritDmg;
    public int HitRate;
    public int MissRate;
    public float MoveSpeed;
    public int PushDmg;
    public int PushDist;
    public int HpRecovery;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }

    public void Set(GameMessageCore.EntityProfile profile)
    {
        if (profile == null)
        {
            return;
        }

        Lv = profile.Lv;
        Exp = profile.Exp;
        Att = profile.Att;
        AttSpeed = profile.AttSpeed;
        Def = profile.Def;
        HpCurrent = profile.HpCurrent;
        HpLimit = profile.HpLimit;
        CritRate = profile.CritRate;
        CritDmg = profile.CritDmg;
        HitRate = profile.HitRate;
        MissRate = profile.MissRate;
        MoveSpeed = profile.MoveSpeed;
        PushDmg = profile.PushDmg;
        PushDist = profile.PushDist;
        HpRecovery = profile.HpRecovery;
    }

    public GameMessageCore.EntityProfile ToProtoData()
    {
        return new GameMessageCore.EntityProfile()
        {
            Lv = Lv,
            Exp = Exp,
            Att = Att,
            AttSpeed = AttSpeed,
            Def = Def,
            HpCurrent = HpCurrent,
            HpLimit = HpLimit,
            CritRate = CritRate,
            CritDmg = CritDmg,
            HitRate = HitRate,
            MissRate = MissRate,
            MoveSpeed = MoveSpeed,
            PushDmg = PushDmg,
            PushDist = PushDist,
            HpRecovery = HpRecovery,
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
    public int ProduceBeginAt;
    public int HarvestItemCount;
    public int CollectionItemCount;
    public int CollectionAt;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }

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
        ProduceBeginAt = build.ProduceBeginAt;
        HarvestItemCount = build.HarvestItemCount;
        CollectionItemCount = build.CollectionItemCount;
        CollectionAt = build.CollectionAt;
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
            ProduceBeginAt = ProduceBeginAt,
            HarvestItemCount = HarvestItemCount,
            CollectionItemCount = CollectionItemCount,
            CollectionAt = CollectionAt,
        };
        foreach (int landId in LandIds)
        {
            build.LandIds.Add(landId);
        }
        return build;
    }
}

