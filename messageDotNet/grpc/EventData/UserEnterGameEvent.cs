using System;
using UnityEngine;

/// <summary>
/// use nft event
/// </summary>
[Serializable]
public class UserEnterGameEvent
{ // 消息版本号 值为毫秒时间戳
    public long MsgVersion;
    public string SceneServiceAppId;
    public int MapId;
    public GrpcPlayerBaseData BaseData;
    public GrpcVector3 Position;
    public GrpcVector3 Dir;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
}