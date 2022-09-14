using System;
using UnityEngine;

/// <summary>
/// agent service forward client message request
/// </summary>
[Serializable]
internal class PullClientMessageInput
{
    // 消息版本号 值为毫秒时间戳
    public long MsgVersion;
    public string AgentAppId;
    public long UserId;
    public string SocketId;
    public int MsgId;
    public byte[] MsgBody;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
}

/// <summary>
/// agent service forward client message response
/// </summary>
[Serializable]
internal class PullClientMessageOutput
{
    public bool Success;
    public string ErrMsg;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
}

/// <summary>
/// other service broadcast proto message to client request
/// </summary>
[Serializable]
internal class BroadCastToClientInput
{
    // 消息版本号 值为毫秒时间戳
    public long MsgVersion;
    public string ServiceAppId;
    public long UserId;
    public string SocketId;
    public int MsgId;
    public byte[] MsgBody;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
}

/// <summary>
/// agent service forward client message response
/// </summary>
[Serializable]
internal class BroadCastToClientOutput
{
    public bool Success;
    public string ErrMsg;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
}


/// <summary>
/// other service multiple broadcast proto message to client request
/// </summary>
[Serializable]
internal class MultipleBroadCastToClientInput
{
    // 消息版本号 值为毫秒时间戳
    public long MsgVersion;
    public string ServiceAppId;
    public long[] UserList;
    public int MsgId;
    public byte[] MsgBody;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
}

/// <summary>
/// other service multiple broadcast proto message to client response
/// </summary>
[Serializable]
internal class MultipleBroadCastToClientOutput
{
    public bool Success;
    public string ErrMsg;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
}
