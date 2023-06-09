using System;
using UnityEngine;

/// <summary>
/// 服务信息用于服务 注册和释放 request
/// </summary>
[Serializable]
public class ServiceRegisterInput
{
    // 服务数据
    public ServiceData Service;
    // register service current time MS
    public long RegisterAt;
 
}


/// <summary>
/// 服务信息用于服务 注册和释放 response
/// </summary>
[Serializable]
public class ServiceRegisterOutput
{
    public bool Success;
    // register service current time MS
    public long RegisterAt;
    // service manager current time MS
    public long ManagerAt;
 
}