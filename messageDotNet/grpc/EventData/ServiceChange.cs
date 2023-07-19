using System;
using UnityEngine;


/// <summary>
/// 玩家完成服务切换事件
/// </summary>
[Serializable]
public class UserChangeServiceEvent
{
    // 消息版本号 值为毫秒时间戳
    public long MsgVersion;
    public long UserId;
    public ServiceData FormService;
    public ServiceData ToService;
    public string State;  // ChangeMapStart | ChangeMapEnd | ChangeMapCancel
    public string UserAgentAppId;
    public string UserSocketId;
    public GrpcVector3 UserPosition;
    public GrpcVector3 UserDir;
}