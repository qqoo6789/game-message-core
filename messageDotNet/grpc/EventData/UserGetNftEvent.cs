using System;
using UnityEngine;

/// <summary>
/// use nft event
/// </summary>
[Serializable]
public class UserGetNFTEvent
{ // 消息版本号 值为毫秒时间戳
    public long MsgVersion;
    public long UserId;
    public GrpcItemBaseInfo[] Items; 
}