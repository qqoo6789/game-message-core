using System;
using UnityEngine;

/// <summary>
/// use nft event
/// </summary>
[Serializable]
public class TickOutPlayerEvent
{ // 消息版本号 值为毫秒时间戳
    public long MsgVersion;
    public long UserId;
    public string AgentId;
    public string SceneServiceAppId;
    public string SocketId;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
}