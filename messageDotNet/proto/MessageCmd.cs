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
            "ChFtZXNzYWdlX2NtZC5wcm90bxIPZ2FtZU1lc3NhZ2VDb3JlKpIQCgxFbnZl",
            "bG9wZVR5cGUSCwoHVW5rbm93bhAAEhYKEEJyb2FkQ2FzdFRpY2tPdXQQmbMC",
            "Eg0KB0l0ZW1HZXQQgYAEEg0KB0l0ZW1Vc2UQg4AEEg4KCEl0ZW1Ecm9wEIWA",
            "BBISCgxVcGRhdGVBdmF0YXIQh4AEEhIKDFVubG9hZEF2YXRhchCJgAQSFgoQ",
            "QnJvYWRDYXN0SXRlbUFkZBCRgAQSGQoTQnJvYWRDYXN0SXRlbVVwZGF0ZRCS",
            "gAQSFgoQQnJvYWRDYXN0SXRlbURlbBCTgAQSHQoXQnJvYWRDYXN0VXBkYXRl",
            "SXRlbVNsb3QQlIAEEhEKC0dldEl0ZW1TbG90EJWABBIVCg9VcGdyYWRlSXRl",
            "bVNsb3QQl4AEEhIKDFNpZ25pblBsYXllchCZgAQSEwoNU2lnbk91dFBsYXll",
            "chChgAQSEAoKUXVlcnlMYW5kcxClgAQSFwoRQnJvYWRDYXN0SW5pdExhbmQQ",
            "p4AEEhcKEUJyb2FkQ2FzdEluaXRJdGVtEKmABBILCgVCdWlsZBCxgAQSDwoJ",
            "UmVjeWNsaW5nELOABBINCgdDaGFyZ2VkELWABBINCgdIYXJ2ZXN0ELeABBIQ",
            "CgpDb2xsZWN0aW9uELmABBIaChRCcm9hZENhc3RNdWx0aVVwTGFuZBDBgAQS",
            "EwoNU2VsZk5mdEJ1aWxkcxDDgAQSHgoYQnJvYWRDYXN0U2VsZkJ1aWxkVXBk",
            "YXRlEMWABBIhChtCcm9hZENhc3RTZWxmQnVpbGRSZWN5Y2xpbmcQx4AEEhEK",
            "C01pbnRCYXR0ZXJ5EMmABBISCgxRdWVyeUdyYW5hcnkQ0YAEEhQKDkdyYW5h",
            "cnlDb2xsZWN0ENOABBIcChZCcm9hZENhc3RHcmFuYXJ5VXBkYXRlENWABBIc",
            "ChZCcm9hZENhc3RVcEdyYW5hcnlJdGVtENeABBIXChFVcGdyYWRlVGFsZW50",
            "Tm9kZRDZgAQSFAoOUXVlcnlUYWxlbnRFeHAQ4YAEEhoKFEJyb2FkQ2FzdFVw",
            "VGFsZW50RXhwEOOABBIRCgtRdWVyeVBsYXllchCBgAgSEgoMQ3JlYXRlUGxh",
            "eWVyEIOACBIOCghFbnRlck1hcBCBgAwSGAoSVXBkYXRlU2VsZkxvY2F0aW9u",
            "EIOADBIOCghVc2VTa2lsbBCFgAwSEwoNUmVzcGF3blBsYXllchCJgAwSHQoX",
            "QnJvYWRDYXN0SW5pdE1hcEVsZW1lbnQQkYAMEh4KGEJyb2FkQ2FzdE1hcEVu",
            "dGl0eVVwZGF0ZRCSgAwSHAoWQnJvYWRDYXN0RW50aXR5RGVzdHJveRCTgAwS",
            "GQoTQnJvYWRDYXN0RW50aXR5TW92ZRCUgAwSGwoVQnJvYWRDYXN0RW50aXR5",
            "Q29tYmF0EJWADBIcChZCcm9hZENhc3RSZXNwYXduUGxheWVyEJaADBIOCghU",
            "ZWxlcG9ydBCXgAwSIgocQnJvYWRDYXN0RW50aXR5UHJvZmlsZVVwZGF0ZRCZ",
            "gAwSGwoVQnJvYWRDYXN0TW9uc3RlckRlYXRoEKGADBIhChtCcm9hZENhc3RF",
            "bnRpdHlBdmF0YXJVcGRhdGUQooAMEiUKH0Jyb2FkQ2FzdEVudGl0eUJhdHRs",
            "ZURhdGFVcGRhdGUQo4AMEiQKHkJyb2FkQ2FzdEVudGl0eUJ1aWxkRGF0YVVw",
            "ZGF0ZRClgAwSGQoTUGxheWVyQWN0aW9uQ29sbGVjdBCngAwSIgocQnJvYWRD",
            "YXN0UGxheWVyQWN0aW9uQ29sbGVjdBCpgAwSFgoQU2NlbmVEZXN0cnVjdGlv",
            "bhCzgAwSHwoZQnJvYWRDYXN0U2NlbmVEZXN0cnVjdGlvbhC1gAwSFgoQUGxh",
            "eWVyQWN0aW9uQ2hhdBC3gAwSHwoZQnJvYWRDYXN0UGxheWVyQWN0aW9uQ2hh",
            "dBC5gAwSEQoLR2V0SG9tZURhdGEQwYAMEhUKD1F1ZXJ5U2VydmVyVGltZRDD",
            "gAwSHAoWQnJvYWRDYXN0Q2hhbmdlU2VydmljZRDFgAwSIAoaQnJvYWRDYXN0",
            "SG9tZURhdGFJbml0QmF0Y2gQyYAMEh4KGEJyb2FkQ2FzdEhvbWVTa2lsbFJl",
            "c3VsdBDRgAwSFQoPUXVlcnlUYWxlbnRUcmVlENOADBIfChlCcm9hZENhc3RU",
            "YWxlbnRUcmVlVXBkYXRlENWADBIPCglTZWxmVGFza3MQgYAQEhkKE0Jyb2Fk",
            "Q2FzdFVwZGF0ZVRhc2sQg4AQEh0KF0Jyb2FkQ2FzdFVwZGF0ZVRhc2tMaXN0",
            "EISAEBIQCgpBY2NlcHRUYXNrEIWAEBIUCg5BY2NlcHRUYXNrTGlzdBCHgBAS",
            "EQoLQWJhbmRvblRhc2sQiYAQEhUKD0FiYW5kb25UYXNrTGlzdBCRgBASEAoK",
            "VGFza1Jld2FyZBCTgBASGQoTVXBncmFkZVRhc2tQcm9ncmVzcxCVgBASGQoT",
            "QnJvYWRDYXN0VGFza1Jld2FyZBCXgBASHQoXQnJvYWRDYXN0VGFza0xpc3RS",
            "ZXdhcmQQmYAQEhUKD1NlbmRDaGF0TWVzc2FnZRCBgBQSGwoVQnJvYWRDYXN0",
            "Q2hhdE1lc3NhZ2VzEIOAFBIgChpCcm9hZENhc3RSZW1vdmVDaGF0TWVzc2Fn",
            "ZRCEgBQSHgoYQnJvYWRDYXN0VXBkYXRlQ2hhdFN0YXRlEIWAFBIKCgRQaW5n",
            "EIGAGBIdChdCcm9hZENhc3RNc2dBZ2dyZWdhdGlvbhCBgBxiBnByb3RvMw=="));
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
    /// 解锁/升级 天赋节点
    /// </summary>
    [pbr::OriginalName("UpgradeTalentNode")] UpgradeTalentNode = 65625,
    /// <summary>
    /// 查询天赋经验
    /// </summary>
    [pbr::OriginalName("QueryTalentExp")] QueryTalentExp = 65633,
    /// <summary>
    /// 广播更新天赋经验
    /// </summary>
    [pbr::OriginalName("BroadCastUpTalentExp")] BroadCastUpTalentExp = 65635,
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
    [pbr::OriginalName("BroadCastHomeDataInitBatch")] BroadCastHomeDataInitBatch = 196681,
    [pbr::OriginalName("BroadCastHomeSkillResult")] BroadCastHomeSkillResult = 196689,
    /// <summary>
    /// 查询天赋数据
    /// </summary>
    [pbr::OriginalName("QueryTalentTree")] QueryTalentTree = 196691,
    /// <summary>
    /// 广播更新天赋树数据
    /// </summary>
    [pbr::OriginalName("BroadCastTalentTreeUpdate")] BroadCastTalentTreeUpdate = 196693,
    /// <summary>
    ///taskServer协议 : 0x04ZZZZ  任务服务 协议 ************************************
    /// </summary>
    [pbr::OriginalName("SelfTasks")] SelfTasks = 262145,
    /// <summary>
    /// 任务  进度更新(推送)
    /// </summary>
    [pbr::OriginalName("BroadCastUpdateTask")] BroadCastUpdateTask = 262147,
    /// <summary>
    /// 任务链进度更新(推送)
    /// </summary>
    [pbr::OriginalName("BroadCastUpdateTaskList")] BroadCastUpdateTaskList = 262148,
    /// <summary>
    /// 领取任务
    /// </summary>
    [pbr::OriginalName("AcceptTask")] AcceptTask = 262149,
    /// <summary>
    /// 领取任务链任务
    /// </summary>
    [pbr::OriginalName("AcceptTaskList")] AcceptTaskList = 262151,
    /// <summary>
    /// 放弃任务(任务有保护时间)
    /// </summary>
    [pbr::OriginalName("AbandonTask")] AbandonTask = 262153,
    /// <summary>
    /// 放弃任务链的任务(任务有保护时间)
    /// </summary>
    [pbr::OriginalName("AbandonTaskList")] AbandonTaskList = 262161,
    /// <summary>
    /// 获取任务奖励(附带提交任务功能)
    /// </summary>
    [pbr::OriginalName("TaskReward")] TaskReward = 262163,
    /// <summary>
    /// 上报更新任务进度
    /// </summary>
    [pbr::OriginalName("UpgradeTaskProgress")] UpgradeTaskProgress = 262165,
    /// <summary>
    /// 推送获取的任务奖励
    /// </summary>
    [pbr::OriginalName("BroadCastTaskReward")] BroadCastTaskReward = 262167,
    /// <summary>
    /// 推送获取的任务链奖励
    /// </summary>
    [pbr::OriginalName("BroadCastTaskListReward")] BroadCastTaskListReward = 262169,
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
    /// <summary>
    ///协议 : 0x07ZZZZ  广播合并 协议 ************************************
    /// </summary>
    [pbr::OriginalName("BroadCastMsgAggregation")] BroadCastMsgAggregation = 458753,
  }

  #endregion

}

#endregion Designer generated code
