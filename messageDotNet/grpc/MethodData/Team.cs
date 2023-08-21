using System;
using UnityEngine;

/// <summary>
/// call friend service, get team data for user id  
/// </summary>
[Serializable]
public class TeamDataInput
{
    public long UserId;
}
[Serializable]
public class TeamDataOutput
{
    public bool Success;
    public string ErrMsg;
    public bool exist;
    public Team TeamData;
}
