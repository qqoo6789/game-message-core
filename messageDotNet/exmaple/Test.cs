using System;

namespace ProtoBuf.Runtime
{
    public class Test
    {
        public static void test()
        {
            // int headerTypeSize = 4;
            // int headerBodySize = 4;
            // GameMessageCore.Envelope fromEnvelope = new GameMessageCore.Envelope();
            // fromEnvelope.Type = GameMessageCore.EnvelopeType.ItemUse;
            // fromEnvelope.SeqId = 1;
            // fromEnvelope.ErrorCode = 1;
            // fromEnvelope.ErrorMessage = "";
            // fromEnvelope.ItemUseRequest = new GameMessageCore.ItemUseRequest();
            // fromEnvelope.ItemUseRequest.ItemId = "1";
            // byte[] bytes = ProtoUtil.EncodeToBytes(fromEnvelope, (int)GameMessageCore.EnvelopeType.ItemUse, 0, headerTypeSize, headerBodySize);

            // int header = ProtoUtil.DecodeHeader(new Span<byte>(bytes, 0, headerTypeSize).ToArray());
            // GameMessageCore.Envelope toEnvelope = ProtoUtil.Decode(new Span<byte>(bytes, headerTypeSize + headerBodySize, bytes.Length - headerTypeSize - headerBodySize));


        }
    }
}