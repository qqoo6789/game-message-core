using System;
using UnityEngine;

/// <summary>
/// 获取 所有建造物信息 for main service
/// </summary>
[Serializable]
public class MainServiceActionGetAllBuildInput
{
    public int MapId;
    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
}
[Serializable]
public class MainServiceActionGetAllBuildOutput
{
    public bool Success;
    public string ErrMsg;
    public GrpcNftBuild[] AllBuilds;
    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }

}