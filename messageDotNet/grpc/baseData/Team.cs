using System;
using UnityEngine;


[Serializable]
public class TeamMember
{
    public long UserId;
    public string Name;
    public string Icon;
}

[Serializable]
public class Team
{
    public long Id;
    public string Name;
    public long Leader;
    public TeamMember[] Members;
}
