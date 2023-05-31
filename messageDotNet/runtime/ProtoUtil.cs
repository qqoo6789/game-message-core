using System;
using UnityEngine;
using Google.Protobuf;
using GameMessageCore;
using System.Collections.Generic;
using System.Reflection;
using Snappy.Sharp;

namespace ProtoBuf.Runtime
{
    public static class ProtoUtil
    {
        public static readonly SnappyCompressor Compressor = new();
        public static readonly SnappyDecompressor Decompressor = new();
        private static readonly byte[] s_compressBytes = new byte[1024 * 8]; // 压缩字节 缓冲区
        private static readonly Dictionary<string, MessageParser> s_protoClassParserDic = new(); // 协议类解析器parser字典，避免反复反射获取

        /// <summary>
        /// 编码成字节流 [type][bodyLen][body]
        /// </summary>
        /// <param name="message">消息体</param>
        /// <param name="packetType">消息类型</param>
        /// <param name="headerLen">整体头字节长度</param>
        /// <param name="typeHeaderLen">类型头字节长度</param>
        /// <param name="bodyHeaderLen">内容头字节长度</param>
        /// <param name="isNeedCompress">是否需要压缩数据</param>
        /// <returns></returns>
        public static byte[] EncodeToBytes(ProtoMessage message, int headerLen, int typeHeaderLen, int bodyHeaderLen, bool isNeedCompress)
        {
            // 内容长度
            short bodyLength = (short)message.MsgData.CalculateSize();

            // 分配字节数组 长度为[总头长+类型头长+内容头长+内容长度]
            byte[] destination = new byte[headerLen + typeHeaderLen + bodyHeaderLen + bodyLength];
            // 类型头开始索引
            int typeHeaderStartIndex = headerLen;
            // 内容头开始索引            
            int bodyHeaderStartIndex = headerLen + typeHeaderLen;
            // 内容开始索引     
            int bodyStartIndex = headerLen + typeHeaderLen + bodyHeaderLen;

            //（一般而言，socket 才会进入这个判断。给整个消息体加一个长度头）
            if (headerLen > 0)
            {
                byte[] headerBytes = BitConverter.GetBytes(typeHeaderLen + bodyHeaderLen + bodyLength);
                if (headerBytes.Length != headerLen)
                {
                    Debug.LogError($"protobuf Serialize headerLen size error,cur={headerBytes.Length} expected={headerLen}");
                    return null;
                }
                Buffer.BlockCopy(headerBytes, 0, destination, 0, headerLen);
            }

            if (typeHeaderLen > 0)
            {
                // 写入 包类型头
                byte[] typeHeaderBytes = BitConverter.GetBytes((int)message.MsgType);
                if (typeHeaderBytes.Length != typeHeaderLen)
                {
                    Debug.LogError($"protobuf Serialize typeHeaderLen size error,cur={typeHeaderBytes.Length} expected={typeHeaderLen}");
                    return null;
                }
                Buffer.BlockCopy(typeHeaderBytes, 0, destination, typeHeaderStartIndex, typeHeaderLen);
            }


            if (bodyHeaderLen > 0)
            {
                // 写入 内容头
                byte[] bodyHeaderBytes = BitConverter.GetBytes(bodyLength);
                if (bodyHeaderBytes.Length != bodyHeaderLen)
                {
                    Debug.LogError($"protobuf Serialize bodyHeaderLen size error,cur={bodyHeaderBytes.Length} expected={bodyHeaderLen}");
                    return null;
                }
                Buffer.BlockCopy(bodyHeaderBytes, 0, destination, bodyHeaderStartIndex, bodyHeaderLen);
            }

            // 写入内容
            Span<byte> bodyBytes = new(destination, bodyStartIndex, bodyLength);
            message.MsgData.WriteTo(bodyBytes);

            // 压缩
            if (isNeedCompress)
            {
                // byte[] compressBytes = new byte[Compressor.MaxCompressedLength(destination.Length)];
                int compressLen = Compressor.Compress(destination, 0, destination.Length, s_compressBytes);
                return new Span<byte>(s_compressBytes, 0, compressLen).ToArray();
            }
            return destination;
        }

        /// <summary>
        /// 解码协议数据 该方法解码数据格式为 [type][bodyLen][body]
        /// </summary>
        /// <param name="protoMsgResults">装载结果</param>
        /// <param name="source">源字节数组</param>
        /// <param name="typeHeaderLen">类型头字节长度</param>
        /// <param name="bodyHeaderLen">内容头字节长度</param>
        /// <param name="suffix">消息后缀</param>
        /// <param name="isNeedDeCompress">是否需要解压</param>
        /// <returns>解析结果数组</returns>
        public static void Decode(List<ProtoMessage> protoMsgResults, byte[] source, int typeHeaderLen, int bodyHeaderLen, eProtoMsgSuffix suffix, bool isNeedDeCompress)
        {
            protoMsgResults.Clear();
            if (typeHeaderLen <= 0)
            {
                Debug.LogError($"protobuf Decode typeHeaderLen error,typeHeaderLen={typeHeaderLen} expected > 0");
                return;
            }
            if (bodyHeaderLen <= 0)
            {
                Debug.LogError($"protobuf Decode bodyHeaderLen error,bodyHeaderLen={bodyHeaderLen} expected > 0");
                return;
            }
            if (isNeedDeCompress)
            {
                source = Decompressor.Decompress(source, 0, source.Length);
            }
            for (int i = 0; i < source.Length;)
            {
                // 取的[类型]部分的Span,并转为 EnvelopeType
                Span<byte> typeHeaderSpan = new(source, i, typeHeaderLen);
                EnvelopeType msgType = (EnvelopeType)BitConverter.ToInt32(typeHeaderSpan);

                // 取的[内容长度]部分的Span，并转为 长度数值
                Span<byte> bodyHeaderSpan = new(source, i + typeHeaderLen, bodyHeaderLen);
                short bodyLen = BitConverter.ToInt16(bodyHeaderSpan);

                // 取的[内容]部分的Span
                Span<byte> bodySpan = new(source, i + typeHeaderLen + bodyHeaderLen, bodyLen);

                IMessage message = ParseMessage(bodySpan, msgType, suffix);
                if (message != null)
                {
                    protoMsgResults.Add(ProtoMessage.Create(msgType, suffix, message));
                }

                // 移动指针
                i += typeHeaderLen + bodyHeaderLen + bodyLen;
            }
        }


        /// <summary>
        /// 解码协议数据 该方法解码数据格式为 [body] ,type 消息类型由外部传入
        /// </summary>
        /// <param name="body"></param>
        /// <param name="msgType"></param>
        /// <param name="suffix"></param>
        /// <param name="isNeedDeCompress">是否需要解压</param>
        /// <returns></returns>
        public static ProtoMessage Decode(byte[] body, EnvelopeType msgType, eProtoMsgSuffix suffix, bool isNeedDeCompress)
        {
            if (isNeedDeCompress)
            {
                body = Decompressor.Decompress(body, 0, body.Length);
            }
            IMessage message = ParseMessage(body, msgType, suffix);
            return message != null ? ProtoMessage.Create(msgType, suffix, message) : null;
        }

        /// <summary>
        /// 获取解析器
        /// </summary>
        /// <param name="msgType"></param>
        /// <param name="suffix"></param>
        /// <returns></returns>
        private static MessageParser GetMessageParser(EnvelopeType msgType, eProtoMsgSuffix suffix)
        {
            // 消息体类名称
            string respClassName = $"{ProtoDefine.PROTO_NAMESPACE}.{msgType}{suffix}";

            if (s_protoClassParserDic.TryGetValue(respClassName, out MessageParser messageParser))
            {
                return messageParser;
            }

            Type respClassType = Type.GetType(respClassName);
            if (respClassType != null)
            {
                // 获取parser
                PropertyInfo parserProperty = respClassType.GetProperty("Parser", BindingFlags.Static | BindingFlags.Public);
                MessageParser parser = (MessageParser)parserProperty.GetValue(null);
                s_protoClassParserDic.Add(respClassName, parser);
                return parser;
            }
            else
            {
                // 类型未找到
                Debug.LogError($"GetMessageParser no find respClassName :{respClassName}");
            }
            return null;
        }

        /// <summary>
        /// 解析 Span message
        /// </summary>
        /// <param name="body"></param>
        /// <param name="msgType"></param>
        /// <param name="suffix"></param>
        /// <returns></returns>
        private static IMessage ParseMessage(Span<byte> body, EnvelopeType msgType, eProtoMsgSuffix suffix)
        {
            try
            {
                MessageParser messageParser = GetMessageParser(msgType, suffix);

                if (messageParser != null)
                {
                    IMessage message = messageParser.ParseFrom(body);
                    return message;
                }
            }
            catch (Exception e)
            {
                Debug.LogError($"ParseMessage parser error :{e}");
            }

            return null;
        }

        /// <summary>
        /// 头长字节数组 转 内容长度
        /// </summary>
        /// <param name="headerBytes"></param>
        /// <returns></returns>
        public static int DecodeHeader(byte[] headerBytes)
        {
            return BitConverter.ToInt32(headerBytes, 0);
        }

    }
}