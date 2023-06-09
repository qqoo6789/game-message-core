using System;
using UnityEngine;

/// <summary>
/// service 销毁事件
/// </summary>
[Serializable]
public class ServiceUnregisterEvent
{
    // 消息版本号 值为毫秒时间戳
    public long MsgVersion;
    public ServiceData Service; 
}