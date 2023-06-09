using System;
using UnityEngine;

/// <summary>
/// 向manager service 查询服务信息
/// </summary>
[Serializable]
public class ManagerActionSelectServiceInput
{
    public GameMessageCore.ServiceType ServiceType;
    public GameMessageCore.SceneServiceSubType SceneSerSubType;
    public long OwnerId;
    public int MapId;
}


/// <summary>
/// 服务信息用于服务 注册和释放 response
/// </summary>
[Serializable]
public class ManagerActionSelectServiceOutput
{
    public int ErrorCode;
    public string ErrorMessage;
    public ServiceData Service;
 

}

/// <summary>
/// 向manager service 批量查询服务信息
/// </summary>
[Serializable]
public class MultiSelectServiceInput
{
    public GameMessageCore.ServiceType ServiceType;
    public GameMessageCore.SceneServiceSubType SceneSerSubType;
    public long OwnerId;
    public int MapId;
    
}
[Serializable]
public class MultiSelectServiceOutput
{
    public int ErrorCode;
    public string ErrorMessage;
    public ServiceData[] Services;

}
