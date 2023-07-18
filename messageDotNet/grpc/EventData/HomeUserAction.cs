using System;
using UnityEngine;


/// <summary>
/// 家园播种事件
/// </summary>
[Serializable]
public class HomeActionSowingEvent
{
    public long MsgVersion;
    public long UserId;
    public uint SeedCid;
    public long SoilId;
}

/// <summary>
/// 家园放置动物食物事件
/// </summary>
[Serializable]
public class HomeActionPutAnimalFoodEvent
{
    public long MsgVersion;
    public long UserId;
    public uint FoodCid;
    public long Num;
}



[Serializable]
public class HomeHarvestInfo
{
    public long SoilId;
    public GrpcItemBaseInfo[] Items;
}

/// <summary>
/// 家园收获土地产出事件
/// </summary>
[Serializable]
public class HomeHarvestEvent
{
    public long MsgVersion;
    public long UserId;
    public HomeHarvestInfo[] Harvests;
}
