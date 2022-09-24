using System;
using UnityEngine;


[Serializable]
public class GrpcPlayerFeature
{
    public int32 Eyebrow;
    public int32 Mouth;
    public int32 Eye;
    public int32 Face;
    public int32 Hair;
    public int32 Glove;
    public int32 Clothes;
    public int32 Pants;
	
    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }

    public void Set(GameMessageCore.PlayerFeature feature)
    {
        if (feature == null)
        {
            return;
        }
        Eyebrow = feature.Eyebrow;
        Mouth = feature.Mouth;
        Eye = feature.Eye;
        Face = feature.Face;
        Hair = feature.Hair;
        Glove = feature.Glove;
        Clothes = feature.Clothes;
        Pants = feature.Pants;
    }

    public GameMessageCore.PlayerFeature ToProtoData()
    {
        return new GameMessageCore.PlayerFeature()
        {
            Eyebrow = Eyebrow,
            Mouth = Mouth,
            Eye = Eye,
            Face = Face,
            Hair = Hair,
            Glove = Glove,
            Clothes = Clothes,
            Pants = Pants,
        };
    }
}


// 对应 proto.PlayerBaseData
[Serializable]
public class GrpcPlayerBaseData
{
    public long UserId;
    public string Name;
    public int RoleId;
    public string Gender;
    public string RoleIcon;
    public GrpcPlayerFeature Feature;
    public bool Guide;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }

    public void Set(GameMessageCore.PlayerBaseData baseData)
    {
        if (baseData == null)
        {
            return;
        }
        UserId = baseData.UserId;
        Name = baseData.Name;
        RoleId = baseData.RoleId;
        Gender = baseData.Gender;
        RoleIcon = baseData.RoleIcon;
        Guide = baseData.Guide;
        Feature.Set(baseData.Feature);
    }

    public GameMessageCore.PlayerBaseData ToProtoData()
    {
        return new GameMessageCore.PlayerBaseData
        {
            UserId = UserId,
            Name = Name,
            RoleId = RoleId,
            Gender = Gender,
            RoleIcon = RoleIcon,
            Feature = Feature.ToProtoData(),
            Guide = Guide,
        };
    }

}



