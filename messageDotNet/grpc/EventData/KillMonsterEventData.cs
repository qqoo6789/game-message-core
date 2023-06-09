using System;
using UnityEngine;

/// <summary>
/// kill monster event json 交互数据结构
/// </summary>
[Serializable]
public class KillMonsterEventData
{
    public long MsgVersion;
    public string SceneServiceAppId;
    public int MapId;
    public long OwnerId;
    public long KillerId;
    public float PosX;
    public float PosY;
    public float PosZ;
    public int MonsterCid;
    public string MonsterName;
    public GrpcItemBaseInfo[] DropList; 
}