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
            "ChFtZXNzYWdlX2NtZC5wcm90bxIPZ2FtZU1lc3NhZ2VDb3JlKuQOCgxFbnZl",
            "bG9wZVR5cGUSCwoHVW5rbm93bhAAEhYKEEJyb2FkQ2FzdFRpY2tPdXQQmbMC",
            "Eg0KB0l0ZW1HZXQQgYAEEg0KB0l0ZW1Vc2UQg4AEEg4KCEl0ZW1Ecm9wEIWA",
            "BBISCgxVcGRhdGVBdmF0YXIQh4AEEhIKDFVubG9hZEF2YXRhchCJgAQSFgoQ",
            "QnJvYWRDYXN0SXRlbUFkZBCRgAQSGQoTQnJvYWRDYXN0SXRlbVVwZGF0ZRCS",
            "gAQSFgoQQnJvYWRDYXN0SXRlbURlbBCTgAQSHQoXQnJvYWRDYXN0VXBkYXRl",
            "SXRlbVNsb3QQlIAEEhEKC0dldEl0ZW1TbG90EJWABBIVCg9VcGdyYWRlSXRl",
            "bVNsb3QQl4AEEhIKDFNpZ25pblBsYXllchCZgAQSEwoNU2lnbk91dFBsYXll",
            "chChgAQSGAoSVXBncmFkZVBsYXllckxldmVsEKOABBIQCgpRdWVyeUxhbmRz",
            "EKWABBIXChFCcm9hZENhc3RJbml0TGFuZBCngAQSFwoRQnJvYWRDYXN0SW5p",
            "dEl0ZW0QqYAEEgsKBUJ1aWxkELGABBIPCglSZWN5Y2xpbmcQs4AEEg0KB0No",
            "YXJnZWQQtYAEEg0KB0hhcnZlc3QQt4AEEhAKCkNvbGxlY3Rpb24QuYAEEhoK",
            "FEJyb2FkQ2FzdE11bHRpVXBMYW5kEMGABBITCg1TZWxmTmZ0QnVpbGRzEMOA",
            "BBIeChhCcm9hZENhc3RTZWxmQnVpbGRVcGRhdGUQxYAEEiEKG0Jyb2FkQ2Fz",
            "dFNlbGZCdWlsZFJlY3ljbGluZxDHgAQSEQoLTWludEJhdHRlcnkQyYAEEhIK",
            "DFF1ZXJ5R3JhbmFyeRDRgAQSFAoOR3JhbmFyeUNvbGxlY3QQ04AEEhwKFkJy",
            "b2FkQ2FzdEdyYW5hcnlVcGRhdGUQ1YAEEhwKFkJyb2FkQ2FzdFVwR3JhbmFy",
            "eUl0ZW0Q14AEEhEKC1F1ZXJ5VGFsZW50ENmABBIXChFVcGdyYWRlVGFsZW50",
            "Tm9kZRDhgAQSEQoLUXVlcnlQbGF5ZXIQgYAIEhIKDENyZWF0ZVBsYXllchCD",
            "gAgSDgoIRW50ZXJNYXAQgYAMEhgKElVwZGF0ZVNlbGZMb2NhdGlvbhCDgAwS",
            "DgoIVXNlU2tpbGwQhYAMEhMKDVJlc3Bhd25QbGF5ZXIQiYAMEh0KF0Jyb2Fk",
            "Q2FzdEluaXRNYXBFbGVtZW50EJGADBIeChhCcm9hZENhc3RNYXBFbnRpdHlV",
            "cGRhdGUQkoAMEhwKFkJyb2FkQ2FzdEVudGl0eURlc3Ryb3kQk4AMEhkKE0Jy",
            "b2FkQ2FzdEVudGl0eU1vdmUQlIAMEhsKFUJyb2FkQ2FzdEVudGl0eUNvbWJh",
            "dBCVgAwSHAoWQnJvYWRDYXN0UmVzcGF3blBsYXllchCWgAwSDgoIVGVsZXBv",
            "cnQQl4AMEiIKHEJyb2FkQ2FzdEVudGl0eVByb2ZpbGVVcGRhdGUQmYAMEhsK",
            "FUJyb2FkQ2FzdE1vbnN0ZXJEZWF0aBChgAwSIQobQnJvYWRDYXN0RW50aXR5",
            "QXZhdGFyVXBkYXRlEKKADBIlCh9Ccm9hZENhc3RFbnRpdHlCYXR0bGVEYXRh",
            "VXBkYXRlEKOADBIkCh5Ccm9hZENhc3RFbnRpdHlCdWlsZERhdGFVcGRhdGUQ",
            "pYAMEhkKE1BsYXllckFjdGlvbkNvbGxlY3QQp4AMEiIKHEJyb2FkQ2FzdFBs",
            "YXllckFjdGlvbkNvbGxlY3QQqYAMEhYKEFNjZW5lRGVzdHJ1Y3Rpb24Qs4AM",
            "Eh8KGUJyb2FkQ2FzdFNjZW5lRGVzdHJ1Y3Rpb24QtYAMEhYKEFBsYXllckFj",
            "dGlvbkNoYXQQt4AMEh8KGUJyb2FkQ2FzdFBsYXllckFjdGlvbkNoYXQQuYAM",
            "EhEKC0dldEhvbWVEYXRhEMGADBIVCg9RdWVyeVNlcnZlclRpbWUQw4AMEhwK",
            "FkJyb2FkQ2FzdENoYW5nZVNlcnZpY2UQxYAMEhkKE0hvbWVSZXNvdXJjZU9w",
            "ZXJhdGUQx4AMEiAKGkJyb2FkQ2FzdEhvbWVEYXRhSW5pdEJhdGNoEMmADBIP",
            "CglTZWxmVGFza3MQgYAQEh0KF0Jyb2FkQ2FzdFVwZGF0ZVRhc2tMaXN0EIOA",
            "EBIQCgpBY2NlcHRUYXNrEIWAEBIVCg9BYmFuZG9ubWVudFRhc2sQh4AQEhAK",
            "ClRhc2tSZXdhcmQQiYAQEhQKDlRhc2tMaXN0UmV3YXJkEJGAEBIZChNVcGdy",
            "YWRlVGFza1Byb2dyZXNzEJOAEBIZChNCcm9hZENhc3RUYXNrUmV3YXJkEJWA",
            "EBIVCg9TZW5kQ2hhdE1lc3NhZ2UQgYAUEhsKFUJyb2FkQ2FzdENoYXRNZXNz",
            "YWdlcxCDgBQSIAoaQnJvYWRDYXN0UmVtb3ZlQ2hhdE1lc3NhZ2UQhIAUEh4K",
            "GEJyb2FkQ2FzdFVwZGF0ZUNoYXRTdGF0ZRCFgBQSCgoEUGluZxCBgBhiBnBy",
            "b3RvMw=="));
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
    /// 踢玩家下线
    /// </summary>
    [pbr::OriginalName("BroadCastTickOut")] BroadCastTickOut = 39321,
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
    /// 登录角色
    /// </summary>
    [pbr::OriginalName("SigninPlayer")] SigninPlayer = 65561,
    /// <summary>
    /// 角色退出游戏
    /// </summary>
    [pbr::OriginalName("SignOutPlayer")] SignOutPlayer = 65569,
    /// <summary>
    /// 手动点击升级
    /// </summary>
    [pbr::OriginalName("UpgradePlayerLevel")] UpgradePlayerLevel = 65571,
    /// <summary>
    /// 请求所有占地地块信息
    /// </summary>
    [pbr::OriginalName("QueryLands")] QueryLands = 65573,
    /// <summary>
    /// 请求QueryLands后地块数据分批次推送  
    /// </summary>
    [pbr::OriginalName("BroadCastInitLand")] BroadCastInitLand = 65575,
    /// <summary>
    /// 初始玩家道具数据     
    /// </summary>
    [pbr::OriginalName("BroadCastInitItem")] BroadCastInitItem = 65577,
    /// <summary>
    /// 建造
    /// </summary>
    [pbr::OriginalName("Build")] Build = 65585,
    /// <summary>
    /// 拆除
    /// </summary>
    [pbr::OriginalName("Recycling")] Recycling = 65587,
    /// <summary>
    /// 充电charged
    /// </summary>
    [pbr::OriginalName("Charged")] Charged = 65589,
    /// <summary>
    /// 收获(harvest)自己建造物的产出(有电量的build) 
    /// </summary>
    [pbr::OriginalName("Harvest")] Harvest = 65591,
    /// <summary>
    /// 采集/偷取(collection) 他人的或者自己的没电量的建造物产出 
    /// </summary>
    [pbr::OriginalName("Collection")] Collection = 65593,
    /// <summary>
    /// 广播批量更新地格信息
    /// </summary>
    [pbr::OriginalName("BroadCastMultiUpLand")] BroadCastMultiUpLand = 65601,
    /// <summary>
    /// 请求自己所有的建造物
    /// </summary>
    [pbr::OriginalName("SelfNftBuilds")] SelfNftBuilds = 65603,
    /// <summary>
    /// 广播更新建造物数据,只广播给owner
    /// </summary>
    [pbr::OriginalName("BroadCastSelfBuildUpdate")] BroadCastSelfBuildUpdate = 65605,
    /// <summary>
    /// 广播移除ui界面中的建筑物,只广播给owner
    /// </summary>
    [pbr::OriginalName("BroadCastSelfBuildRecycling")] BroadCastSelfBuildRecycling = 65607,
    /// <summary>
    /// 使用token 购买电池(透传给 theweb3)
    /// </summary>
    [pbr::OriginalName("MintBattery")] MintBattery = 65609,
    /// <summary>
    /// 查询仓库所有道具
    /// </summary>
    [pbr::OriginalName("QueryGranary")] QueryGranary = 65617,
    /// <summary>
    /// 取出仓库中的所有道具
    /// </summary>
    [pbr::OriginalName("GranaryCollect")] GranaryCollect = 65619,
    /// <summary>
    /// 全量更新仓库数据
    /// </summary>
    [pbr::OriginalName("BroadCastGranaryUpdate")] BroadCastGranaryUpdate = 65621,
    /// <summary>
    /// 更新仓库道具
    /// </summary>
    [pbr::OriginalName("BroadCastUpGranaryItem")] BroadCastUpGranaryItem = 65623,
    /// <summary>
    /// 查询天赋数据
    /// </summary>
    [pbr::OriginalName("QueryTalent")] QueryTalent = 65625,
    /// <summary>
    /// 解锁/升级 天赋节点
    /// </summary>
    [pbr::OriginalName("UpgradeTalentNode")] UpgradeTalentNode = 65633,
    /// <summary>
    ///accountSer协议 : 0x02ZZZZ  账号服务 协议 ************************************
    /// </summary>
    [pbr::OriginalName("QueryPlayer")] QueryPlayer = 131073,
    [pbr::OriginalName("CreatePlayer")] CreatePlayer = 131075,
    /// <summary>
    ///sceneSer 协议 : 0x03ZZZZ  战斗服务 协议 ************************************
    /// </summary>
    [pbr::OriginalName("EnterMap")] EnterMap = 196609,
    [pbr::OriginalName("UpdateSelfLocation")] UpdateSelfLocation = 196611,
    [pbr::OriginalName("UseSkill")] UseSkill = 196613,
    [pbr::OriginalName("RespawnPlayer")] RespawnPlayer = 196617,
    [pbr::OriginalName("BroadCastInitMapElement")] BroadCastInitMapElement = 196625,
    [pbr::OriginalName("BroadCastMapEntityUpdate")] BroadCastMapEntityUpdate = 196626,
    [pbr::OriginalName("BroadCastEntityDestroy")] BroadCastEntityDestroy = 196627,
    [pbr::OriginalName("BroadCastEntityMove")] BroadCastEntityMove = 196628,
    [pbr::OriginalName("BroadCastEntityCombat")] BroadCastEntityCombat = 196629,
    [pbr::OriginalName("BroadCastRespawnPlayer")] BroadCastRespawnPlayer = 196630,
    [pbr::OriginalName("Teleport")] Teleport = 196631,
    [pbr::OriginalName("BroadCastEntityProfileUpdate")] BroadCastEntityProfileUpdate = 196633,
    [pbr::OriginalName("BroadCastMonsterDeath")] BroadCastMonsterDeath = 196641,
    [pbr::OriginalName("BroadCastEntityAvatarUpdate")] BroadCastEntityAvatarUpdate = 196642,
    [pbr::OriginalName("BroadCastEntityBattleDataUpdate")] BroadCastEntityBattleDataUpdate = 196643,
    [pbr::OriginalName("BroadCastEntityBuildDataUpdate")] BroadCastEntityBuildDataUpdate = 196645,
    [pbr::OriginalName("PlayerActionCollect")] PlayerActionCollect = 196647,
    [pbr::OriginalName("BroadCastPlayerActionCollect")] BroadCastPlayerActionCollect = 196649,
    [pbr::OriginalName("SceneDestruction")] SceneDestruction = 196659,
    [pbr::OriginalName("BroadCastSceneDestruction")] BroadCastSceneDestruction = 196661,
    [pbr::OriginalName("PlayerActionChat")] PlayerActionChat = 196663,
    [pbr::OriginalName("BroadCastPlayerActionChat")] BroadCastPlayerActionChat = 196665,
    [pbr::OriginalName("GetHomeData")] GetHomeData = 196673,
    /// <summary>
    /// 客户端请求服务器时间
    /// </summary>
    [pbr::OriginalName("QueryServerTime")] QueryServerTime = 196675,
    [pbr::OriginalName("BroadCastChangeService")] BroadCastChangeService = 196677,
    [pbr::OriginalName("HomeResourceOperate")] HomeResourceOperate = 196679,
    [pbr::OriginalName("BroadCastHomeDataInitBatch")] BroadCastHomeDataInitBatch = 196681,
    /// <summary>
    ///taskServer协议 : 0x04ZZZZ  任务服务 协议 ************************************
    /// </summary>
    [pbr::OriginalName("SelfTasks")] SelfTasks = 262145,
    /// <summary>
    /// 任务链进度更新(推送)
    /// </summary>
    [pbr::OriginalName("BroadCastUpdateTaskList")] BroadCastUpdateTaskList = 262147,
    /// <summary>
    /// 领取任务
    /// </summary>
    [pbr::OriginalName("AcceptTask")] AcceptTask = 262149,
    /// <summary>
    /// 放弃任务(任务有保护时间)
    /// </summary>
    [pbr::OriginalName("AbandonmentTask")] AbandonmentTask = 262151,
    /// <summary>
    /// 获取任务奖励(附带提交任务功能)
    /// </summary>
    [pbr::OriginalName("TaskReward")] TaskReward = 262153,
    /// <summary>
    /// 获取任务链奖励
    /// </summary>
    [pbr::OriginalName("TaskListReward")] TaskListReward = 262161,
    /// <summary>
    /// 上报更新任务进度
    /// </summary>
    [pbr::OriginalName("UpgradeTaskProgress")] UpgradeTaskProgress = 262163,
    /// <summary>
    /// 推送获取的任务奖励
    /// </summary>
    [pbr::OriginalName("BroadCastTaskReward")] BroadCastTaskReward = 262165,
    /// <summary>
    ///chatServer协议 : 0x05ZZZZ  聊天服务 协议 ************************************
    /// </summary>
    [pbr::OriginalName("SendChatMessage")] SendChatMessage = 327681,
    /// <summary>
    /// 批量推送的聊天消息
    /// </summary>
    [pbr::OriginalName("BroadCastChatMessages")] BroadCastChatMessages = 327683,
    [pbr::OriginalName("BroadCastRemoveChatMessage")] BroadCastRemoveChatMessage = 327684,
    [pbr::OriginalName("BroadCastUpdateChatState")] BroadCastUpdateChatState = 327685,
    /// <summary>
    ///getaway   协议 : 0x06ZZZZ  网关服务 协议 ************************************
    /// </summary>
    [pbr::OriginalName("Ping")] Ping = 393217,
  }

  #endregion

}

#endregion Designer generated code
