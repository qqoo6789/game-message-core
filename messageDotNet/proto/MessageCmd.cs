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
            "ChFtZXNzYWdlX2NtZC5wcm90bxIPZ2FtZU1lc3NhZ2VDb3JlKvwTCgxFbnZl",
            "bG9wZVR5cGUSCwoHVW5rbm93bhAAEhYKEEJyb2FkQ2FzdFRpY2tPdXQQmbMC",
            "Eg0KB0l0ZW1HZXQQgYAEEg0KB0l0ZW1Vc2UQg4AEEg4KCEl0ZW1Ecm9wEIWA",
            "BBISCgxVcGRhdGVBdmF0YXIQh4AEEhIKDFVubG9hZEF2YXRhchCJgAQSFgoQ",
            "QnJvYWRDYXN0SXRlbUFkZBCRgAQSGQoTQnJvYWRDYXN0SXRlbVVwZGF0ZRCS",
            "gAQSFgoQQnJvYWRDYXN0SXRlbURlbBCTgAQSEgoMU2lnbmluUGxheWVyEJmA",
            "BBITCg1TaWduT3V0UGxheWVyEKGABBIQCgpRdWVyeUxhbmRzEKWABBIXChFC",
            "cm9hZENhc3RJbml0TGFuZBCngAQSFwoRQnJvYWRDYXN0SW5pdEl0ZW0QqYAE",
            "EgsKBUJ1aWxkELGABBIPCglSZWN5Y2xpbmcQs4AEEg0KB0NoYXJnZWQQtYAE",
            "Eg0KB0hhcnZlc3QQt4AEEhAKCkNvbGxlY3Rpb24QuYAEEhoKFEJyb2FkQ2Fz",
            "dE11bHRpVXBMYW5kEMGABBITCg1TZWxmTmZ0QnVpbGRzEMOABBIeChhCcm9h",
            "ZENhc3RTZWxmQnVpbGRVcGRhdGUQxYAEEiEKG0Jyb2FkQ2FzdFNlbGZCdWls",
            "ZFJlY3ljbGluZxDHgAQSEQoLTWludEJhdHRlcnkQyYAEEhIKDFF1ZXJ5R3Jh",
            "bmFyeRDRgAQSFAoOR3JhbmFyeUNvbGxlY3QQ04AEEhwKFkJyb2FkQ2FzdEdy",
            "YW5hcnlVcGRhdGUQ1YAEEhwKFkJyb2FkQ2FzdFVwR3JhbmFyeUl0ZW0Q14AE",
            "EhcKEVVwZ3JhZGVUYWxlbnROb2RlENmABBIUCg5RdWVyeVRhbGVudEV4cBDh",
            "gAQSGgoUQnJvYWRDYXN0VXBUYWxlbnRFeHAQ44AEEhUKD1F1ZXJ5QW5pbWFs",
            "TGlzdBDlgAQSHwoZQnJvYWRDYXN0VXBkYXRlQW5pbWFsTGlzdBDngAQSEQoL",
            "RnJlZWRBbmltYWwQ6YAEEhgKElF1ZXJ5UGxheWVyU2V0dGluZxDxgAQSFwoR",
            "U2F2ZVBsYXllclNldHRpbmcQ84AEEhEKC1F1ZXJ5UGxheWVyEIGACBISCgxD",
            "cmVhdGVQbGF5ZXIQg4AIEgsKBUxvZ2luEIWACBIUCg5Ccm9hZENhc3RRdWV1",
            "ZRCHgAgSDgoIRW50ZXJNYXAQgYAMEhgKElVwZGF0ZVNlbGZMb2NhdGlvbhCD",
            "gAwSDgoIVXNlU2tpbGwQhYAMEhMKDVJlc3Bhd25QbGF5ZXIQiYAMEh0KF0Jy",
            "b2FkQ2FzdEluaXRNYXBFbGVtZW50EJGADBIeChhCcm9hZENhc3RNYXBFbnRp",
            "dHlVcGRhdGUQkoAMEhwKFkJyb2FkQ2FzdEVudGl0eURlc3Ryb3kQk4AMEhkK",
            "E0Jyb2FkQ2FzdEVudGl0eU1vdmUQlIAMEhsKFUJyb2FkQ2FzdEVudGl0eUNv",
            "bWJhdBCVgAwSHAoWQnJvYWRDYXN0UmVzcGF3blBsYXllchCWgAwSDgoIVGVs",
            "ZXBvcnQQl4AMEiIKHEJyb2FkQ2FzdEVudGl0eVByb2ZpbGVVcGRhdGUQmYAM",
            "EhsKFUJyb2FkQ2FzdE1vbnN0ZXJEZWF0aBChgAwSIQobQnJvYWRDYXN0RW50",
            "aXR5QXZhdGFyVXBkYXRlEKKADBIlCh9Ccm9hZENhc3RFbnRpdHlCYXR0bGVE",
            "YXRhVXBkYXRlEKOADBIkCh5Ccm9hZENhc3RFbnRpdHlCdWlsZERhdGFVcGRh",
            "dGUQpYAMEhkKE1BsYXllckFjdGlvbkNvbGxlY3QQp4AMEiIKHEJyb2FkQ2Fz",
            "dFBsYXllckFjdGlvbkNvbGxlY3QQqYAMEhYKEFNjZW5lRGVzdHJ1Y3Rpb24Q",
            "s4AMEh8KGUJyb2FkQ2FzdFNjZW5lRGVzdHJ1Y3Rpb24QtYAMEhYKEFBsYXll",
            "ckFjdGlvbkNoYXQQt4AMEh8KGUJyb2FkQ2FzdFBsYXllckFjdGlvbkNoYXQQ",
            "uYAMEhEKC0dldEhvbWVEYXRhEMGADBIVCg9RdWVyeVNlcnZlclRpbWUQw4AM",
            "EhwKFkJyb2FkQ2FzdENoYW5nZVNlcnZpY2UQxYAMEiAKGkJyb2FkQ2FzdEhv",
            "bWVEYXRhSW5pdEJhdGNoEMmADBIeChhCcm9hZENhc3RIb21lU2tpbGxSZXN1",
            "bHQQ0YAMEhUKD1F1ZXJ5VGFsZW50VHJlZRDTgAwSHwoZQnJvYWRDYXN0VGFs",
            "ZW50VHJlZVVwZGF0ZRDVgAwSJAoeQnJvYWRDYXN0QW5pbWFsU2NlbmVEYXRh",
            "VXBkYXRlENeADBIaChRCcm9hZENhc3RBbmltYWxEZWF0aBDZgAwSIAoaQnJv",
            "YWRDYXN0QW5pbWFsQXV0b1Byb2R1Y2UQ4YAMEhkKE0Jyb2FkQ2FzdFBpY2t1",
            "cERyb3AQ44AMEhkKE0FuaW1hbFNwZWNpYWxBY3Rpb24Q5YAMEhMKDUNhcHR1",
            "cmVBbmltYWwQ54AMEh4KGEJyb2FkQ2FzdEVudGl0eU1vdmVTcGVlZBDpgAwS",
            "JgogQnJvYWRDYXN0RW50aXR5Q2FwdHVyZURhdGFVcGRhdGUQ8IAMEhEKC0V4",
            "aXRDYXB0dXJlEPGADBIcChZFbnRpdHlBY2N1bXVsYXRlU3RhdHVzEPKADBIl",
            "Ch9Ccm9hZENhc3RFbnRpdHlBY2N1bXVsYXRlU3RhdHVzEPOADBIUCg5FbnRp",
            "dHlEaWFsb2d1ZRD1gAwSIwodQnJvYWRDYXN0RW50aXR5RXF1aXBtZW50U2tp",
            "bGwQ94AMEiUKH0Jyb2FkQ2FzdEVudGl0eUNhcHR1cmVUcmFwU2tpbGwQ+YAM",
            "Eg8KCVNlbGZUYXNrcxCBgBASGQoTQnJvYWRDYXN0VXBkYXRlVGFzaxCDgBAS",
            "HQoXQnJvYWRDYXN0VXBkYXRlVGFza0xpc3QQhIAQEhAKCkFjY2VwdFRhc2sQ",
            "hYAQEhQKDkFjY2VwdFRhc2tMaXN0EIeAEBIRCgtBYmFuZG9uVGFzaxCJgBAS",
            "FQoPQWJhbmRvblRhc2tMaXN0EJGAEBIQCgpUYXNrUmV3YXJkEJOAEBIZChNV",
            "cGdyYWRlVGFza1Byb2dyZXNzEJWAEBIZChNCcm9hZENhc3RUYXNrUmV3YXJk",
            "EJeAEBIdChdCcm9hZENhc3RUYXNrTGlzdFJld2FyZBCZgBASFQoPU2VuZENo",
            "YXRNZXNzYWdlEIGAFBIbChVCcm9hZENhc3RDaGF0TWVzc2FnZXMQg4AUEiAK",
            "GkJyb2FkQ2FzdFJlbW92ZUNoYXRNZXNzYWdlEISAFBIeChhCcm9hZENhc3RV",
            "cGRhdGVDaGF0U3RhdGUQhYAUEgoKBFBpbmcQgYAYYgZwcm90bzM="));
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
    /// 请求客户端所有自定义存储数据
    /// </summary>
    [pbr::OriginalName("QueryPlayerSetting")] QueryPlayerSetting = 65649,
    /// <summary>
    /// 存储客户端自定义存储数据
    /// </summary>
    [pbr::OriginalName("SavePlayerSetting")] SavePlayerSetting = 65651,
    /// <summary>
    ///accountSer协议 : 0x02ZZZZ  账号服务 协议 ************************************
    /// </summary>
    [pbr::OriginalName("QueryPlayer")] QueryPlayer = 131073,
    [pbr::OriginalName("CreatePlayer")] CreatePlayer = 131075,
    [pbr::OriginalName("Login")] Login = 131077,
    [pbr::OriginalName("BroadCastQueue")] BroadCastQueue = 131079,
    /// <summary>
    ///sceneSer 协议 : 0x03ZZZZ  场景服务 协议 ************************************
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
    /// 实体对话
    /// </summary>
    [pbr::OriginalName("EntityDialogue")] EntityDialogue = 196725,
    /// <summary>
    /// 广播实体装备技能
    /// </summary>
    [pbr::OriginalName("BroadCastEntityEquipmentSkill")] BroadCastEntityEquipmentSkill = 196727,
    /// <summary>
    /// 捕获释放技能结果
    /// </summary>
    [pbr::OriginalName("BroadCastEntityCaptureTrapSkill")] BroadCastEntityCaptureTrapSkill = 196729,
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
    ///agentService协议 : 0x06ZZZZ  网关服务 协议 ************************************
    /// </summary>
    [pbr::OriginalName("Ping")] Ping = 393217,
  }

  #endregion

}

#endregion Designer generated code
