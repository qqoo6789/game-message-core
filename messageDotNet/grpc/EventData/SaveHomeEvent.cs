using System;
using UnityEngine;

/// <summary>
/// 玩家数据保持到main server, event json 交互数据结构
/// </summary>
[Serializable]
public class SaveHomeEvent : EventDataBase
{
    public long MsgVersion;
    public GrpcHomeData Data;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
}