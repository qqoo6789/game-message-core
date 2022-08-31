using System;

namespace ProtoBuf.Runtime
{
    public class Test
    {
        public static void test()
        {
            int headerSize = 4;
            GameMessageCore.Envelope fromEnvelope = new GameMessageCore.Envelope();
            fromEnvelope.Type = GameMessageCore.EnvelopeType.ItemUse;
            fromEnvelope.SeqId = 1;
            fromEnvelope.ErrorCode = 1;
            fromEnvelope.ErrorMessage = "";
            fromEnvelope.ItemUseRequest = new GameMessageCore.ItemUseRequest();
            fromEnvelope.ItemUseRequest.ItemId = "1";
            byte[] bytes = ProtoUtil.EncodeToBytes(fromEnvelope, headerSize);

            int header = ProtoUtil.DecodeHeader(new Span<byte>(bytes, 0, headerSize).ToArray());
            GameMessageCore.Envelope toEnvelope = ProtoUtil.Decode(new Span<byte>(bytes, headerSize, bytes.Length - headerSize));
        }
    }
}