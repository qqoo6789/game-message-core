using System;

/// <summary>
/// 玩家死亡 event json 交互数据结构
/// </summary>
[Serializable]
internal class PlayerDeathEventData : EventDataBase
{
    public int KillerType;
    public long KillerId;
    public string KillerName;
}