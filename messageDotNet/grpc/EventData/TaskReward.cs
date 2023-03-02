using System;
using UnityEngine;

/// <summary>
/// task reward event
/// </summary>
[Serializable]
public class UserTaskRewardEvent
{ // 消息版本号 值为毫秒时间戳
    public long MsgVersion;
    public long UserId;
    public GrpcTalentExp[] Exps;
    public GrpcItemBaseInfo[] Items;
    public bool TaskListReward;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }

}