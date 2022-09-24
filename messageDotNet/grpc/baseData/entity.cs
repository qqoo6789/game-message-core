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

