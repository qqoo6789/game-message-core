using System;


/// <summary>
/// kill monster event json 交互数据结构
/// </summary>
[Serializable]
internal class KillMonsterEventData : EventDataBase
{
    public int MonsterCid;
    public string MonsterName;
    public GameMessageCore.ItemBaseInfo[] DropList;
    public int Exp;
}