using System;
using UnityEngine;

/// <summary>
/// update user talent data event
/// </summary>
[Serializable]
public class UpdateTalentEvent
{ // 消息版本号 值为毫秒时间戳
    public long MsgVersion;
    public long UserId;
    public GrpcTalentLevel[] Levels;
    public GrpcTalentTreeUpdate[] Trees;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }

}