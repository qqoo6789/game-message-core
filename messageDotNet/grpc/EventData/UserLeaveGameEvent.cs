using System;
using UnityEngine;

/// <summary>
/// use leave game event
/// </summary>
[Serializable]
public class UserLeaveGameEvent
{ // 消息版本号 值为毫秒时间戳
    public long MsgVersion;
    public string SceneServiceAppId;
    public long UserId; 
}