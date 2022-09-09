using System;
using UnityEngine;



/// <summary>
/// 拾取 event json 交互数据结构
/// </summary>
[Serializable]
internal class PackDropEventData
{
    public int ServerId;
    public int MapId;
    public int UserId;
    public int DropCid;
    public string DropName;
    public float PosX;
    public float PosY;
    public float PosZ;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }

}