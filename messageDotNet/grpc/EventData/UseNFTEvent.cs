using System;
using UnityEngine;

/// <summary>
/// use nft event
/// </summary>
[Serializable]
public class UserUseNFTEvent
{ // 消息版本号 值为毫秒时间戳
    public long MsgVersion;
    public long UserId;
    public string NftId;
    public int Cid;
    public GameMessageCore.NftTraitType NftType;
    public int Num;
    public GrpcNFTConsumableInfo ConsumableData;
    public float X;
    public float Y;
    public float Y;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }

}