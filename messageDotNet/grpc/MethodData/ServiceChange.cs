/*
 * @Author: alex
 * @Date: 2022-12-9 10:00
 * @Description: user 切换 scene service 需要的grpc calls 相关交换协议
 * @FilePath: /meland-scene-server/Assets/Plugins/SharedCore/game-message-core/messageDotNet/grpc/MethodData/ServiceChange.cs
 */

using System;
using UnityEngine;

/// <summary>
/// 向manager service 请求启动服务 request
/// </summary>
[Serializable]
public class StartServiceInput
{
    public GameMessageCore.ServiceType ServiceType;
    public GameMessageCore.SceneServiceSubType SceneSerSubType;
    public long OwnerId;
    public int MapId;
}
[Serializable]
public class StartServiceOutput
{
    public bool Success;
    public string ErrMsg;
    public ServiceData SerInfo;

}


/// <summary>
/// user请求进入目标scene service request（预进入阶段）
/// </summary>
[Serializable]
public class ApplyEnterServiceInput
{
    public long ApplyUser;
    public ServiceData FromSer;
}
[Serializable]
public class ApplyEnterServiceOutput
{
    public long UserId;
    public bool Success;
    public string ErrMsg;
}


/// <summary>
/// user正式进入目标scene service request（正式进入阶段）
/// </summary>
[Serializable]
public class JoinServiceInput
{
    public long UserId;
    public string UserAgentAppId;
    public string UserSocketId;
    public int CurHp;
    public GrpcVector3 UserPosition;
    public GrpcVector3 UserDir;
    public ServiceData FromSer;
    public string SkillEffectData;
    public long UsedMs;

}
[Serializable]
public class JoinServiceOutput
{
    public long UserId;
    public bool Success;
    public string ErrMsg;

}


