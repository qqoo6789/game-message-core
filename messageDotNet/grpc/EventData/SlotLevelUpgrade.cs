using System;
using UnityEngine;

[Serializable]
public class SlotLevelUpgradeEvent
{ // 消息版本号 值为毫秒时间戳
    public long MsgVersion;
    public long UserId;
    public int SlotPos;
    public int Level;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }

}