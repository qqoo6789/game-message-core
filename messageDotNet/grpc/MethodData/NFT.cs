using System;
using UnityEngine;

[Serializable]
public class GrpcNftBaseData
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
/// call mainService User use nft
/// </summary>
[Serializable]
public class MainServiceActionUseNftInput
{
    public long UserId;
    public GrpcNftBaseData Nft;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
}
[Serializable]
public class MainServiceActionUseNftOutput
{
    public bool Success;
    public string FailedMsg;

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
    public long UserId;
    public GrpcNftBaseData[] TakeNfts;

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

/// <summary>
/// dapr call mainService mint User nft input
/// </summary>
[Serializable]
public class MainServiceActionMintNftInput
{
    public long UserId;
    public GrpcItemBaseInfo Item;
    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
}
[Serializable]
public class MainServiceActionMintNftOutput
{
    public bool Success;
    public string FailedMsg;
    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }

}