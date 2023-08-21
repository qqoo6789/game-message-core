using System;
using UnityEngine;


[Serializable]
public class CreateTeamEvent
{ // 消息版本号 值为毫秒时间戳
    public long MsgVersion;
    public Team TeamData;
}


[Serializable]
public class DisbandTeamEvent
{ // 消息版本号 值为毫秒时间戳
    public long MsgVersion;
    public long TeamId;
}


[Serializable]
public class JoinTeamEvent
{ // 消息版本号 值为毫秒时间戳
    public long MsgVersion;
    public Team TeamData;
    public TeamMember JoinMember;
}


[Serializable]
public class QuitTeamEvent
{ // 消息版本号 值为毫秒时间戳
    public long MsgVersion;
    public long TeamId;
    public long Member;
}

