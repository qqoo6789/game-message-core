using System;
using UnityEngine;



[Serializable]
public class TickOutPlayerInput
{
    public long UserId;
    public string AgentId;
    public string SceneServiceAppId;
    public string SocketId;
    public GameMessageCore.TickOutType TickOutCode;
}

[Serializable]
public class TickOutPlayerOutput
{
    public bool Success;
    public string ErrMsg;
}