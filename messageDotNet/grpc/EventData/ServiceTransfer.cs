using System;
using UnityEngine;

/// <summary>
/// SPOT服务迁移事件
/// </summary>
[Serializable]
public class ServiceTransferEvent
{
    // 消息版本号 值为毫秒时间戳
    public long MsgVersion;
    public ServiceData Service;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
}