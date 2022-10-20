using System;
using UnityEngine;


[Serializable]
public class NftBuildAddEvent
{ // 消息版本号 值为毫秒时间戳
    public long MsgVersion;
    public GrpcNftBuild Build;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
}


[Serializable]
public class NftBuildUpdateEvent
{ // 消息版本号 值为毫秒时间戳
    public long MsgVersion;
    public GrpcNftBuild Build;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
}


[Serializable]
public class NftBuildRemoveEvent
{ // 消息版本号 值为毫秒时间戳
    public long MsgVersion;
    public long EntityId;
    public string FromNft;
    public long Owner;
    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
}