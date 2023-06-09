using System;
using UnityEngine;

/// <summary>
/// event json 交互数据结构公共数据字段
/// </summary>
[Serializable]
public class EventDataBase
{
    public int MapId;
    public long UserId;
    public float PosX;
    public float PosY;
    public float PosZ; 
}