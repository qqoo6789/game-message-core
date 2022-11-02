using System;
using UnityEngine;

/// <summary>
/// 服务信息用于服务 注册和释放 request
/// </summary>
[Serializable]
public class ServiceRegisterInput
{
    // 消息版本号 值为毫秒时间戳
    public long MsgVersion;

    // 服务数据
    public ServiceData Service;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }

}


/// <summary>
/// 服务信息用于服务 注册和释放 response
/// </summary>
[Serializable]
public class ServiceRegisterOutput
{
    // 消息版本号 值为毫秒时间戳
    public long MsgVersion;
    public bool Success;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }

}