using System;
using UnityEngine;


[Serializable]
public class MintNftInfo
{
    public long Owner;
    public GrpcItemBaseInfo[] Items;
    public EntityInfo FromEntity;
}

/// <summary>
/// dapr call mainService mint User nft input
/// </summary>
[Serializable]
public class MintNftEvent
{
    // 消息版本号 值为毫秒时间戳
    public long MsgVersion;
    public ServiceData FormService;

    public MintNftInfo Nft;
}

/// <summary>
/// dapr call mainService mint User nft input
/// </summary>
[Serializable]
public class MultiMintNftEvent
{
    // 消息版本号 值为毫秒时间戳
    public long MsgVersion;
    public ServiceData FormService;
    public MintNftInfo[] Nfts;
}