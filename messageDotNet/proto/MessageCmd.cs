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
            "ChFtZXNzYWdlX2NtZC5wcm90bxIPZ2FtZU1lc3NhZ2VDb3JlKqwTCgxFbnZl",
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
            "VGFsZW50RXhwEOOABBIVCg9RdWVyeUFuaW1hbExpc3QQ5YAEEh8KGUJyb2Fk",
            "Q2FzdFVwZGF0ZUFuaW1hbExpc3QQ54AEEhEKC0ZyZWVkQW5pbWFsEOmABBIR",
            "CgtRdWVyeVBsYXllchCBgAgSEgoMQ3JlYXRlUGxheWVyEIOACBIOCghFbnRl",
            "ck1hcBCBgAwSGAoSVXBkYXRlU2VsZkxvY2F0aW9uEIOADBIOCghVc2VTa2ls",
            "bBCFgAwSEwoNUmVzcGF3blBsYXllchCJgAwSHQoXQnJvYWRDYXN0SW5pdE1h",
            "cEVsZW1lbnQQkYAMEh4KGEJyb2FkQ2FzdE1hcEVudGl0eVVwZGF0ZRCSgAwS",
            "HAoWQnJvYWRDYXN0RW50aXR5RGVzdHJveRCTgAwSGQoTQnJvYWRDYXN0RW50",
            "aXR5TW92ZRCUgAwSGwoVQnJvYWRDYXN0RW50aXR5Q29tYmF0EJWADBIcChZC",
            "cm9hZENhc3RSZXNwYXduUGxheWVyEJaADBIOCghUZWxlcG9ydBCXgAwSIgoc",
            "QnJvYWRDYXN0RW50aXR5UHJvZmlsZVVwZGF0ZRCZgAwSGwoVQnJvYWRDYXN0",
            "TW9uc3RlckRlYXRoEKGADBIhChtCcm9hZENhc3RFbnRpdHlBdmF0YXJVcGRh",
            "dGUQooAMEiUKH0Jyb2FkQ2FzdEVudGl0eUJhdHRsZURhdGFVcGRhdGUQo4AM",
            "EiQKHkJyb2FkQ2FzdEVudGl0eUJ1aWxkRGF0YVVwZGF0ZRClgAwSGQoTUGxh",
            "eWVyQWN0aW9uQ29sbGVjdBCngAwSIgocQnJvYWRDYXN0UGxheWVyQWN0aW9u",
            "Q29sbGVjdBCpgAwSFgoQU2NlbmVEZXN0cnVjdGlvbhCzgAwSHwoZQnJvYWRD",
            "YXN0U2NlbmVEZXN0cnVjdGlvbhC1gAwSFgoQUGxheWVyQWN0aW9uQ2hhdBC3",
            "gAwSHwoZQnJvYWRDYXN0UGxheWVyQWN0aW9uQ2hhdBC5gAwSEQoLR2V0SG9t",
            "ZURhdGEQwYAMEhUKD1F1ZXJ5U2VydmVyVGltZRDDgAwSHAoWQnJvYWRDYXN0",
            "Q2hhbmdlU2VydmljZRDFgAwSIAoaQnJvYWRDYXN0SG9tZURhdGFJbml0QmF0",
            "Y2gQyYAMEh4KGEJyb2FkQ2FzdEhvbWVTa2lsbFJlc3VsdBDRgAwSFQoPUXVl",
            "cnlUYWxlbnRUcmVlENOADBIfChlCcm9hZENhc3RUYWxlbnRUcmVlVXBkYXRl",
            "ENWADBIkCh5Ccm9hZENhc3RBbmltYWxTY2VuZURhdGFVcGRhdGUQ14AMEhoK",
            "FEJyb2FkQ2FzdEFuaW1hbERlYXRoENmADBIgChpCcm9hZENhc3RBbmltYWxB",
            "dXRvUHJvZHVjZRDhgAwSGQoTQnJvYWRDYXN0UGlja3VwRHJvcBDjgAwSGQoT",
            "QW5pbWFsU3BlY2lhbEFjdGlvbhDlgAwSEwoNQ2FwdHVyZUFuaW1hbBDngAwS",
            "HgoYQnJvYWRDYXN0RW50aXR5TW92ZVNwZWVkEOmADBImCiBCcm9hZENhc3RF",
            "bnRpdHlDYXB0dXJlRGF0YVVwZGF0ZRDwgAwSEQoLRXhpdENhcHR1cmUQ8YAM",
            "EhwKFkVudGl0eUFjY3VtdWxhdGVTdGF0dXMQ8oAMEiUKH0Jyb2FkQ2FzdEVu",
            "dGl0eUFjY3VtdWxhdGVTdGF0dXMQ84AMEg8KCVNlbGZUYXNrcxCBgBASGQoT",
            "QnJvYWRDYXN0VXBkYXRlVGFzaxCDgBASHQoXQnJvYWRDYXN0VXBkYXRlVGFz",
            "a0xpc3QQhIAQEhAKCkFjY2VwdFRhc2sQhYAQEhQKDkFjY2VwdFRhc2tMaXN0",
            "EIeAEBIRCgtBYmFuZG9uVGFzaxCJgBASFQoPQWJhbmRvblRhc2tMaXN0EJGA",
            "EBIQCgpUYXNrUmV3YXJkEJOAEBIZChNVcGdyYWRlVGFza1Byb2dyZXNzEJWA",
            "EBIZChNCcm9hZENhc3RUYXNrUmV3YXJkEJeAEBIdChdCcm9hZENhc3RUYXNr",
            "TGlzdFJld2FyZBCZgBASFQoPU2VuZENoYXRNZXNzYWdlEIGAFBIbChVCcm9h",
            "ZENhc3RDaGF0TWVzc2FnZXMQg4AUEiAKGkJyb2FkQ2FzdFJlbW92ZUNoYXRN",
            "ZXNzYWdlEISAFBIeChhCcm9hZENhc3RVcGRhdGVDaGF0U3RhdGUQhYAUEgoK",
            "BFBpbmcQgYAYEh0KF0Jyb2FkQ2FzdE1zZ0FnZ3JlZ2F0aW9uEIGAHGIGcHJv",
            "dG8z"));
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
    /// 请求Animal列表
    /// </summary>
    [pbr::OriginalName("QueryAnimalList")] QueryAnimalList = 65637,
    /// <summary>
    /// 广播批量更新Animal列表(add | update | remove)
    /// </summary>
    [pbr::OriginalName("BroadCastUpdateAnimalList")] BroadCastUpdateAnimalList = 65639,
    /// <summary>
    /// 释放Animal
    /// </summary>
    [pbr::OriginalName("FreedAnimal")] FreedAnimal = 65641,
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
    /// 广播更新动物场数据
    /// </summary>
    [pbr::OriginalName("BroadCastAnimalSceneDataUpdate")] BroadCastAnimalSceneDataUpdate = 196695,
    /// <summary>
    /// 广播动物死亡
    /// </summary>
    [pbr::OriginalName("BroadCastAnimalDeath")] BroadCastAnimalDeath = 196697,
    /// <summary>
    /// 广播动物自动生产
    /// </summary>
    [pbr::OriginalName("BroadCastAnimalAutoProduce")] BroadCastAnimalAutoProduce = 196705,
    /// <summary>
    /// 广播拾取掉落物
    /// </summary>
    [pbr::OriginalName("BroadCastPickupDrop")] BroadCastPickupDrop = 196707,
    /// <summary>
    /// 请求动物特殊行为
    /// </summary>
    [pbr::OriginalName("AnimalSpecialAction")] AnimalSpecialAction = 196709,
    /// <summary>
    /// 抓捕Animal(同步可能释放一个 Animal) to scene service
    /// </summary>
    [pbr::OriginalName("CaptureAnimal")] CaptureAnimal = 196711,
    /// <summary>
    /// 实体速度
    /// </summary>
    [pbr::OriginalName("BroadCastEntityMoveSpeed")] BroadCastEntityMoveSpeed = 196713,
    [pbr::OriginalName("BroadCastEntityCaptureDataUpdate")] BroadCastEntityCaptureDataUpdate = 196720,
    /// <summary>
    /// 客户端请求退出捕获状态
    /// </summary>
    [pbr::OriginalName("ExitCapture")] ExitCapture = 196721,
    /// <summary>
    /// 实体蓄力状态请求
    /// </summary>
    [pbr::OriginalName("EntityAccumulateStatus")] EntityAccumulateStatus = 196722,
    /// <summary>
    /// 实体蓄力状态广播
    /// </summary>
    [pbr::OriginalName("BroadCastEntityAccumulateStatus")] BroadCastEntityAccumulateStatus = 196723,
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
