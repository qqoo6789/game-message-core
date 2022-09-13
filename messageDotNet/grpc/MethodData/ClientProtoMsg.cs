using System;
using UnityEngine;

/// <summary>
/// agent service forward client message request
/// </summary>
[Serializable]
internal class ClientProtoMessageInput
{
    // 消息版本号 值为毫秒时间戳
    public long MsgVersion;
    public string AgentAppId;
    public long UserId;
    public byte[] ProtoBytesReq;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
}


/// <summary>
/// agent service forward client message response
/// </summary>
[Serializable]
internal class ClientProtoMessageOutput
{
    // 消息版本号 值为毫秒时间戳
    public long MsgVersion;
    public string AgentAppId;
    public long UserId;
    public byte[] ProtoBytesResp;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }

}