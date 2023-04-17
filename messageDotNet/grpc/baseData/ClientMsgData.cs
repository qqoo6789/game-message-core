using System;
using UnityEngine;

// 对应 proto.Vector3
[Serializable]
public class MultiClientMsgData
{
    public long[] UserList;
    public int MsgId;
    public byte[] MsgBody;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }

}