/*
 * @Author: xiang huan
 * @Date: 2023-04-25 13:46:11
 * @Description: 协议使用引用池，注意清理干净
 * @FilePath: /meland-scene-server/Assets/Plugins/SharedCore/game-message-core/messageDotNet/custom/proto/Envelope.Custom.cs
 * 
 */
using System.Reflection;

namespace GameMessageCore
{
    public sealed partial class Envelope : ICustomReference
    {
        private PropertyInfo _payloadInfo;
        public static Envelope Create(EnvelopeType type, string suffix, int errCode, string errMsg, object payload)
        {
            Envelope reference = CustomReferencePool.Acquire<Envelope>();
            reference.Type = type;
            reference.ErrorCode = errCode;
            reference.ErrorMessage = errMsg;

            string name = type.ToString() + suffix.ToString();
            reference._payloadInfo = reference.GetType().GetProperty(name);
            reference._payloadInfo.SetValue(reference, payload);
            reference.SeqId = 0;

            return reference;
        }

        public void Dispose()
        {
            CustomReferencePool.Release(this);
        }

        public void Clear()
        {

            object payload = _payloadInfo.GetValue(this);
            if (payload is not null and ICustomReference)
            {
                CustomReferencePool.Release(payload);
            }
            _payloadInfo.SetValue(this, null);

            Type = EnvelopeType.Unknown;
            ErrorCode = 0;
            ErrorMessage = "";
            SeqId = 0;
        }
    }
}
