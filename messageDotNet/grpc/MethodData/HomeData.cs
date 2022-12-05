using System;
using UnityEngine;

/// <summary>
/// 获取 所有建造物信息 for main service
/// </summary>
[Serializable]
public class MainServiceActionGetHomeDataInput
{
    public long UserId;
    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
}
[Serializable]
public class MainServiceActionGetHomeDataOutput
{
    public bool Success;
    public string ErrMsg;
    public GrpcHomeData Data;
    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }

}