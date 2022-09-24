using System;
using UnityEngine;
/// <summary>
/// 玩家死亡 event json 交互数据结构
/// </summary>
[Serializable]
public class PlayerDeathEventData : EventDataBase
{
    public int KillerType;
    public long KillerId;
    public string KillerName;
    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
}