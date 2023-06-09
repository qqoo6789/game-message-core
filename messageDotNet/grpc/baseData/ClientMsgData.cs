using System;
using UnityEngine;

// 对应 proto.Vector3
[Serializable]
public class MultiClientMsgData : ICustomReference
{
    public long[] UserList;
    public int MsgId;
    public byte[] MsgBody;

    public void Dispose()
    {
        CustomReferencePool.Release(this);
    }
    public static MultiClientMsgData Create(long[] userList, int msgId, byte[] msgBody)
    {
        MultiClientMsgData reference = CustomReferencePool.Acquire<MultiClientMsgData>();
        reference.UserList = userList;
        reference.MsgId = msgId;
        reference.MsgBody = msgBody;
        return reference;
    }
    public void Clear()
    {
        UserList = null;
        MsgId = 0;
        MsgBody = null;
    }
}