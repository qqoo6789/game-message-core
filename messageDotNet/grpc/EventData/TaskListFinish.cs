using System;
using UnityEngine;

/// <summary>
/// task list finish event
/// </summary>
[Serializable]
public class TaskListFinishEvent
{ // 消息版本号 值为毫秒时间戳
    public long MsgVersion;
    public long UserId;
    public GameMessageCore.TaskListType TaskListType;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }

}