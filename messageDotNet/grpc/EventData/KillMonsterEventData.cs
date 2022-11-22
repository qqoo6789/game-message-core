using System;
using UnityEngine;

/// <summary>
/// kill monster event json 交互数据结构
/// </summary>
[Serializable]
public class KillMonsterEventData : EventDataBase
{
    public long MsgVersion;
    public string SceneServiceAppId;
    public int MonsterCid;
    public string MonsterName;
    public GrpcItemBaseInfo[] DropList;
    public int Exp;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
}