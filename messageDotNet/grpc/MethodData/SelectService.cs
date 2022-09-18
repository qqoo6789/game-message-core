using System;
using UnityEngine;

/// <summary>
/// 向manager service 查询服务信息
/// </summary>
[Serializable]
internal class ManagerActionSelectServiceInput
{
    // 消息版本号 值为毫秒时间戳
    public long MsgVersion;
    public GameMessageCore.ServiceType ServiceType;
    public int MapId;
    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
}


/// <summary>
/// 服务信息用于服务 注册和释放 response
/// </summary>
[Serializable]
internal class ManagerActionSelectServiceOutput
{
    public int ErrorCode;
    public string ErrorMessage;
    public GameMessageCore.ServiceType ServiceType;
    public string ServiceAppId;
    public int MapId;
    public string Host;
    public int Port;
    public int Online;
    public int MaxOnline;
    public long CreateAt;
    public long UpdateAt;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }

}