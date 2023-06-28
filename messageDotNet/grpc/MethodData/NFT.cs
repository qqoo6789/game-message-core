using System;
using UnityEngine;

[Serializable]
public class GrpcNftBaseData
{
    public string NftId;
    public int ItemCid;
    public int Num;
}

/// <summary>
/// call mainService User use nft
/// </summary>
[Serializable]
public class MainServiceActionUseNftInput
{
    public long UserId;
    public GrpcNftBaseData Nft;
}
[Serializable]
public class MainServiceActionUseNftOutput
{
    public bool Success;
    public string FailedMsg;
}

/// <summary>
/// 通知mainService 扣除User nft input
/// </summary>
[Serializable]
public class MainServiceActionTakeNftInput
{
    public long UserId;
    public GrpcNftBaseData[] TakeNfts;
}
[Serializable]
public class MainServiceActionTakeNftOutput
{
    public bool Success;
    public string FailedMsg;
}

/// <summary>
/// dapr call mainService mint User nft input
/// </summary>
[Serializable]
public class MainServiceActionMintNftInput
{
    public long UserId;
    public GrpcItemBaseInfo Item;
}
[Serializable]
public class MainServiceActionMintNftOutput
{
    public bool Success;
    public string FailedMsg;
}

/// <summary>
/// dapr call mainService mint User nft input
/// </summary>
[Serializable]
public class MainServiceActionMultiMintNftInput
{
    public long UserId;
    public GrpcItemBaseInfo[] Items;
}
[Serializable]
public class MainServiceActionMultiMintNftOutput
{
    public bool Success;
    public string FailedMsg;

}