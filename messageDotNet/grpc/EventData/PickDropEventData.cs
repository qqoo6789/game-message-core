using System;

/// <summary>
/// 拾取 event json 交互数据结构
/// </summary>
[Serializable]
internal class PickDropEventData : EventDataBase
{
    public int ServerId;
    public int DropCid;
    public string DropName;
}