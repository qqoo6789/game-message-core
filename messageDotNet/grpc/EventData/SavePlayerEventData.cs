using System;

/// <summary>
/// 玩家数据保持到main server, event json 交互数据结构
/// </summary>
[Serializable]
internal class SavePlayerEventData : EventDataBase
{
    public int CurHP;
    public float DirX;
    public float DirY;
    public float DirZ;
}