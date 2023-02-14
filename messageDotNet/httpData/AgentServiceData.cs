using System;
using UnityEngine;

/// <summary>
/// http请求网关数据的response
/// </summary>
[Serializable]
public class AgentServiceResp
{
    public int ErrorCode;//TODO: NEED USE GLOBAL ERROR CODE
    public string ErrorMessage;
    public string Url;
    public int Online;
    public int MaxOnline;
    public long CreatedAt;
    public long UpdateAt;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }

}