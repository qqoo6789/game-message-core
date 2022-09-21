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
    public GameMessageCore.PlayerBaseData BaseData;
    public GameMessageCore.Vector3 Position;
    public GameMessageCore.Vector3 Dir;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
}