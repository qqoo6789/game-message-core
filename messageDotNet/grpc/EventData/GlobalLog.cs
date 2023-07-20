using System;
using UnityEngine;


/// <summary>
/// dapr call 日志采集
/// </summary>
[Serializable]
public class GlobalLogEvent
{
    // 消息版本号 值为毫秒时间戳
    public long MsgVersion;
    public ServiceData FormService;
    public GameMessageCore.GameLogType Action;
    public long UserId;
    public string ObjectType;
    public long ObjectId;
    public string ObjectName;
    public int64 ObjectOwner;
    public uint ObjectCid;
    public int ObjectNumber;
    public GrpcVector3 Pos;
    public string Desc;
}
