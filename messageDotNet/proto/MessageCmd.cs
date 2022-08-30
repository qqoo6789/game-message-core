// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: message_cmd.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace GameMessageCore {

  /// <summary>Holder for reflection information generated from message_cmd.proto</summary>
  public static partial class MessageCmdReflection {

    #region Descriptor
    /// <summary>File descriptor for message_cmd.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static MessageCmdReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "ChFtZXNzYWdlX2NtZC5wcm90bxIPZ2FtZU1lc3NhZ2VDb3JlKrgCCgxFbnZl",
            "bG9wZVR5cGUSCwoHVW5rbm93bhAAEg0KB0l0ZW1HZXQQgYAEEg0KB0l0ZW1V",
            "c2UQg4AEEg4KCEl0ZW1Ecm9wEIWABBISCgxVcGRhdGVBdmF0YXIQh4AEEhIK",
            "DFVubG9hZEF2YXRhchCJgAQSFgoQQnJvYWRDYXN0SXRlbUFkZBCRgAQSGQoT",
            "QnJvYWRDYXN0SXRlbVVwZGF0ZRCSgAQSFgoQQnJvYWRDYXN0SXRlbURlbBCT",
            "gAQSHQoXQnJvYWRDYXN0VXBkYXRlSXRlbVNsb3QQlIAEEhEKC0dldEl0ZW1T",
            "bG90EJWABBIVCg9VcGdyYWRlSXRlbVNsb3QQl4AEEhEKC1F1ZXJ5UGxheWVy",
            "EIGACBISCgxDcmVhdGVQbGF5ZXIQg4AIEgoKBFBpbmcQgYAYYgZwcm90bzM="));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { },
          new pbr::GeneratedClrTypeInfo(new[] {typeof(global::GameMessageCore.EnvelopeType), }, null, null));
    }
    #endregion

  }
  #region Enums
  public enum EnvelopeType {
    /// <summary>
    ///保留id协议      : 0x00ZZZZ   *********************************************** 
    /// </summary>
    [pbr::OriginalName("Unknown")] Unknown = 0,
    /// <summary>
    ///mainServer协议 : 0x01ZZZZ  主数据服务协议************************************
    /// </summary>
    [pbr::OriginalName("ItemGet")] ItemGet = 65537,
    [pbr::OriginalName("ItemUse")] ItemUse = 65539,
    [pbr::OriginalName("ItemDrop")] ItemDrop = 65541,
    [pbr::OriginalName("UpdateAvatar")] UpdateAvatar = 65543,
    [pbr::OriginalName("UnloadAvatar")] UnloadAvatar = 65545,
    /// <summary>
    /// 添加道具
    /// </summary>
    [pbr::OriginalName("BroadCastItemAdd")] BroadCastItemAdd = 65553,
    /// <summary>
    /// 更新道具
    /// </summary>
    [pbr::OriginalName("BroadCastItemUpdate")] BroadCastItemUpdate = 65554,
    /// <summary>
    /// del道具
    /// </summary>
    [pbr::OriginalName("BroadCastItemDel")] BroadCastItemDel = 65555,
    /// <summary>
    /// 推送 玩家道具槽
    /// </summary>
    [pbr::OriginalName("BroadCastUpdateItemSlot")] BroadCastUpdateItemSlot = 65556,
    /// <summary>
    /// 查询 玩家道具槽信息
    /// </summary>
    [pbr::OriginalName("GetItemSlot")] GetItemSlot = 65557,
    /// <summary>
    /// 升级 玩家道具槽
    /// </summary>
    [pbr::OriginalName("UpgradeItemSlot")] UpgradeItemSlot = 65559,
    /// <summary>
    ///accountSer协议 : 0x02ZZZZ  账号服务 协议 ************************************
    /// </summary>
    [pbr::OriginalName("QueryPlayer")] QueryPlayer = 131073,
    [pbr::OriginalName("CreatePlayer")] CreatePlayer = 131075,
    /// <summary>
    ///getaway   协议 : 0x06ZZZZ  网关服务 协议 ************************************
    /// </summary>
    [pbr::OriginalName("Ping")] Ping = 393217,
  }

  #endregion

}

#endregion Designer generated code
