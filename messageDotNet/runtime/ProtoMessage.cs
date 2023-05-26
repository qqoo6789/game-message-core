

using System.Reflection;
using GameMessageCore;
using Google.Protobuf;

public class ProtoMessage : ICustomReference
{
    // 消息体
    public IMessage MsgData { get; private set; }

    // 消息类型
    public EnvelopeType MsgType { get; private set; }
    // 消息后缀类型
    public eProtoMsgSuffix Suffix;

    public static ProtoMessage Create(EnvelopeType type, eProtoMsgSuffix suffix, IMessage payload)
    {
        ProtoMessage reference = CustomReferencePool.Acquire<ProtoMessage>();
        reference.SetData(payload, type, suffix);
        return reference;
    }

    public void SetData(IMessage msgData, EnvelopeType msgType, eProtoMsgSuffix suffix)
    {
        MsgData = msgData;
        MsgType = msgType;
        Suffix = suffix;
    }

    public int SeqId
    {
        get
        {
            if (Suffix == eProtoMsgSuffix.Req)
            {
                ReqTitle reqTitle = GetAddReqTitle(false);
                if (reqTitle != null)
                {
                    return reqTitle.SeqId;
                }
                return reqTitle.SeqId;
            }
            else if (Suffix == eProtoMsgSuffix.Resp)
            {
                RespTitle respTitle = GetRespTitle();
                if (respTitle != null)
                {
                    return respTitle.SeqId;
                }
            }
            return 0;
        }
        set
        {
            if (Suffix == eProtoMsgSuffix.Req)
            {
                ReqTitle reqTitle = GetAddReqTitle(true);
                reqTitle.SeqId = value;
            }
            else if (Suffix == eProtoMsgSuffix.Resp)
            {
                RespTitle respTitle = GetRespTitle();
                if (respTitle != null)
                {
                    respTitle.SeqId = value;
                }
            }
        }
    }

    public int ErrorCode
    {
        get
        {
            RespTitle respTitle = GetRespTitle();
            if (respTitle != null)
            {
                return respTitle.ErrorCode;
            }
            return 0;
        }
    }

    public string ErrorMessage
    {
        get
        {
            RespTitle respTitle = GetRespTitle();
            if (respTitle != null)
            {
                return respTitle.ErrorMessage;
            }
            return "";
        }
    }

    // 设置ResTitle
    public void SetResTitle(int seqId, int errCode, string errMsg)
    {
        PropertyInfo respProperty = GetMsgDataRespTitleProperty();
        object respObj = respProperty.GetValue(MsgData);
        if (respObj == null)
        {
            respProperty.SetValue(MsgData, new RespTitle() { SeqId = seqId, ErrorCode = errCode, ErrorMessage = errMsg });
        }
        else
        {
            RespTitle respTitle = respObj as RespTitle;
            respTitle.SeqId = seqId;
            respTitle.ErrorCode = errCode;
            respTitle.ErrorMessage = errMsg;
        }
    }

    // 设置ReqTitle
    public void SetReqTitle(int seqId)
    {
        PropertyInfo reqProperty = GetMsgDataReqTitleProperty();
        object reqObj = reqProperty.GetValue(MsgData);
        if (reqObj == null)
        {
            reqProperty.SetValue(MsgData, new ReqTitle() { SeqId = seqId });
        }
        else
        {
            ReqTitle reqTitle = reqObj as ReqTitle;
            reqTitle.SeqId = seqId;
        }
    }

    /// <summary>
    /// 获取与添加 ReqTitle
    /// </summary>
    /// <param name="isAdd">是否需要添加</param>
    /// <returns></returns>
    public ReqTitle GetAddReqTitle(bool isAdd)
    {
        PropertyInfo reqProperty = GetMsgDataReqTitleProperty();
        if (reqProperty != null)
        {
            ReqTitle reqTitle = (ReqTitle)reqProperty.GetValue(MsgData, null);
            if (reqTitle == null && isAdd)
            {
                ReqTitle newReqTitle = new ReqTitle();
                reqProperty.SetValue(MsgData, newReqTitle);
                return newReqTitle;
            }
            return reqTitle;
        }
        return null;
    }

    /// <summary>
    /// 获取RespTitle 属性
    /// </summary>
    /// <returns></returns>
    public RespTitle GetRespTitle()
    {
        PropertyInfo respProperty = GetMsgDataRespTitleProperty();
        if (respProperty != null)
        {
            RespTitle respTitle = (RespTitle)respProperty.GetValue(MsgData, null);
            return respTitle;
        }
        return null;
    }


    // 获取ReqTitle属性
    private PropertyInfo GetMsgDataReqTitleProperty()
    {
        return MsgData.GetType().GetProperty("ReqTitle");
    }

    // 获取ResTitle属性
    private PropertyInfo GetMsgDataRespTitleProperty()
    {
        return MsgData.GetType().GetProperty("ResTitle");
    }


    public void Dispose()
    {
        CustomReferencePool.Release(this);
    }

    public void Clear()
    {
        MsgData = null;
        MsgType = EnvelopeType.Unknown;
        Suffix = eProtoMsgSuffix.None;
    }
}