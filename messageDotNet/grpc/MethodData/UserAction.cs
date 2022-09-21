using System;
using UnityEngine;

/// <summary>
/// player exit game 
/// </summary>
[Serializable]
public class UserLeaveGameInput
{
    // 消息版本号 值为毫秒时间戳
    public long MsgVersion;
    public string AgentAppId;
    public long UserId;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
}
[Serializable]
public class UserLeaveGameOutput
{
    public bool Success;

    public string ErrMsg;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }

}

/// <summary>
///  更新玩家使用的装备
/// </summary>
[Serializable]
public class UpdateUsedAvatarInput
{
    // 消息版本号 值为毫秒时间戳
    public long MsgVersion;
    public long UserId;
    public GameMessageCore.PlayerAvatar[] UsingAvatars;
    public GameMessageCore.EntityProfile CurProfile;
    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
}
[Serializable]
public class UpdateUsedAvatarOutput
{
    public bool Success;
    public string ErrMsg;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }

}

/// <summary>
/// 更新玩家战斗属性(升级/装备槽等级变更)
/// </summary>
[Serializable]
public class UpdateUserProfileInput
{
    // 消息版本号 值为毫秒时间戳
    public long MsgVersion;
    public long UserId;

    public GameMessageCore.EntityProfile CurProfile;
    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
}
[Serializable]
public class UpdateUserProfileOutput
{
    public bool Success;
    public string ErrMsg;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }

}