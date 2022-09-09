using System;
using UnityEngine;

/// <summary>
/// 服务信息用于服务注册
/// </summary>
[Serializable]
internal class ServiceData
{
    public long Id;
    public string Name;
    public string APPID;
    public GameMessageCore.ServiceType ServiceType;
    public string Host;
    public int Port;
    public int MapId;
    public int Online;
    public int MaxOnline;
    public long CreatedAt;
    public long UpdatedAt;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }

}