using System;
using UnityEngine;

/// <summary>
/// 获取 所有建造物信息 for main service
/// </summary>
[Serializable]
public class MainServiceActionGetAllBuildInput
{
    public int MapId;
    
}
[Serializable]
public class MainServiceActionGetAllBuildOutput
{
    public bool Success;
    public string ErrMsg;
    public GrpcNftBuild[] AllBuilds;
     

}