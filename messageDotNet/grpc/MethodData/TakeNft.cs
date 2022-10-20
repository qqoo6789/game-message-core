using System;
using UnityEngine;

[Serializable]
public class TakeNftData
{
    public string NftId;
    public int ItemCid;
    public int Num;
    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
}

/// <summary>
/// 通知mainService 扣除User nft input
/// </summary>
[Serializable]
public class MainServiceActionTakeNftInput
{
    // 消息版本号 值为毫秒时间戳
    public long MsgVersion;
    public int UserId;
    public TakeNftData[] TakeNfts;
    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
}
[Serializable]
public class MainServiceActionTakeNftOutput
{
    public bool Success;
    public string FailedMsg;
    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }

}