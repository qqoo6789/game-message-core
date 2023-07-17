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

}


[Serializable]
public class AddTalentExpInfo
{
    public long Owner;
    public GrpcTalentExp[] AddExps;
    public EntityInfo FromEntity;
}


/// <summary>
/// add  user talent tree exp event
/// </summary>
[Serializable]
public class AddTalentExpEvent
{ // 消息版本号 值为毫秒时间戳
    public long MsgVersion;
    public ServiceData FormService;
    public AddTalentExpInfo AddExp;
}
