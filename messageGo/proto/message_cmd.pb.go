// 应用交互消息定义.
//***************************************************************************
// 消息协议号 ： 服务端主动推送的以 BroadCast 开头   如 BroadCastxxx
// 消息结构   ： server -> client 以 Resp  结尾   如 BroadCastXXResp
// 消息结构   ： server -> client 以 Resp  结尾   如 xxxResp
// 消息结构   ： client -> server 以 Req   结尾   如 xxxReq
//***************************************************************************

//***************************************************************************
// 协议ID规则       : 0xXXZZZZ, 其中 XX 为服务ID头, ZZZZ为具体的协议ID, 示例如下
// 保留id协议       : 0x00ZZZZ
// mainServer协议   : 0x01ZZZZ  主数据服务协议
// accountSer协议   : 0x02ZZZZ  账号服务 协议
// screenSer 协议   : 0x03ZZZZ  战斗服务 协议
// taskServer协议   : 0x04ZZZZ  任务服务 协议
// chatServer协议   : 0x05ZZZZ  聊天服务 协议
// agentServer协议  : 0x06ZZZZ  网关服务 协议
// ServiceMgr协议   : 0x09ZZZZ  服务管理 协议(客户端不能调用)
// CTRlServer协议   : 0x0aZZZZ  控制服务 协议
//**************************************************************************

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: message_cmd.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EnvelopeType int32

const (
	//保留id协议      : 0x00ZZZZ   ***********************************************
	EnvelopeType_Unknown          EnvelopeType = 0
	EnvelopeType_BroadCastTickOut EnvelopeType = 39321 // 踢玩家下线
	//mainServer协议 : 0x01ZZZZ  主数据服务协议************************************
	EnvelopeType_ItemGet                     EnvelopeType = 65537
	EnvelopeType_ItemUse                     EnvelopeType = 65539
	EnvelopeType_ItemDrop                    EnvelopeType = 65541
	EnvelopeType_UpdateAvatar                EnvelopeType = 65543
	EnvelopeType_UnloadAvatar                EnvelopeType = 65545
	EnvelopeType_BroadCastItemAdd            EnvelopeType = 65553 // 添加道具
	EnvelopeType_BroadCastItemUpdate         EnvelopeType = 65554 // 更新道具
	EnvelopeType_BroadCastItemDel            EnvelopeType = 65555 // del道具
	EnvelopeType_SigninPlayer                EnvelopeType = 65561 // 登录角色
	EnvelopeType_SignOutPlayer               EnvelopeType = 65569 // 角色退出游戏
	EnvelopeType_QueryLands                  EnvelopeType = 65573 // 请求所有占地地块信息
	EnvelopeType_BroadCastInitLand           EnvelopeType = 65575 // 请求QueryLands后地块数据分批次推送
	EnvelopeType_BroadCastInitItem           EnvelopeType = 65577 // 初始玩家道具数据
	EnvelopeType_Build                       EnvelopeType = 65585 // 建造
	EnvelopeType_Recycling                   EnvelopeType = 65587 // 拆除
	EnvelopeType_Charged                     EnvelopeType = 65589 // 充电charged
	EnvelopeType_Harvest                     EnvelopeType = 65591 // 收获(harvest)自己建造物的产出(有电量的build)
	EnvelopeType_Collection                  EnvelopeType = 65593 // 采集/偷取(collection) 他人的或者自己的没电量的建造物产出
	EnvelopeType_BroadCastMultiUpLand        EnvelopeType = 65601 // 广播批量更新地格信息
	EnvelopeType_SelfNftBuilds               EnvelopeType = 65603 // 请求自己所有的建造物
	EnvelopeType_BroadCastSelfBuildUpdate    EnvelopeType = 65605 // 广播更新建造物数据,只广播给owner
	EnvelopeType_BroadCastSelfBuildRecycling EnvelopeType = 65607 // 广播移除ui界面中的建筑物,只广播给owner
	EnvelopeType_MintBattery                 EnvelopeType = 65609 // 使用token 购买电池(透传给 theweb3)
	EnvelopeType_QueryGranary                EnvelopeType = 65617 // 查询仓库所有道具
	EnvelopeType_GranaryCollect              EnvelopeType = 65619 // 取出仓库中的所有道具
	EnvelopeType_BroadCastGranaryUpdate      EnvelopeType = 65621 // 全量更新仓库数据
	EnvelopeType_BroadCastUpGranaryItem      EnvelopeType = 65623 // 更新仓库道具
	EnvelopeType_UpgradeTalentNode           EnvelopeType = 65625 // 解锁/升级 天赋节点
	EnvelopeType_QueryTalentExp              EnvelopeType = 65633 // 查询天赋经验
	EnvelopeType_BroadCastUpTalentExp        EnvelopeType = 65635 // 广播更新天赋经验
	EnvelopeType_QueryAnimalList             EnvelopeType = 65637 // 请求Animal列表
	EnvelopeType_BroadCastUpdateAnimalList   EnvelopeType = 65639 // 广播批量更新Animal列表(add | update | remove)
	EnvelopeType_FreedAnimal                 EnvelopeType = 65641 // 释放Animal
	EnvelopeType_QueryPlayerSetting          EnvelopeType = 65649 // 请求客户端所有自定义存储数据
	EnvelopeType_SavePlayerSetting           EnvelopeType = 65651 // 存储客户端自定义存储数据
	//accountSer协议 : 0x02ZZZZ  账号服务 协议 ************************************
	EnvelopeType_QueryPlayer    EnvelopeType = 131073
	EnvelopeType_CreatePlayer   EnvelopeType = 131075
	EnvelopeType_Login          EnvelopeType = 131077
	EnvelopeType_BroadCastQueue EnvelopeType = 131079
	//sceneSer 协议 : 0x03ZZZZ  场景服务 协议 ************************************
	EnvelopeType_EnterMap                         EnvelopeType = 196609
	EnvelopeType_UpdateSelfLocation               EnvelopeType = 196611
	EnvelopeType_UseSkill                         EnvelopeType = 196613
	EnvelopeType_RespawnPlayer                    EnvelopeType = 196617
	EnvelopeType_BroadCastInitMapElement          EnvelopeType = 196625
	EnvelopeType_BroadCastMapEntityUpdate         EnvelopeType = 196626
	EnvelopeType_BroadCastEntityDestroy           EnvelopeType = 196627
	EnvelopeType_BroadCastEntityMove              EnvelopeType = 196628
	EnvelopeType_BroadCastEntityCombat            EnvelopeType = 196629
	EnvelopeType_BroadCastRespawnPlayer           EnvelopeType = 196630
	EnvelopeType_Teleport                         EnvelopeType = 196631
	EnvelopeType_BroadCastEntityProfileUpdate     EnvelopeType = 196633
	EnvelopeType_BroadCastMonsterDeath            EnvelopeType = 196641
	EnvelopeType_BroadCastEntityAvatarUpdate      EnvelopeType = 196642
	EnvelopeType_BroadCastEntityBattleDataUpdate  EnvelopeType = 196643
	EnvelopeType_BroadCastEntityBuildDataUpdate   EnvelopeType = 196645
	EnvelopeType_PlayerActionCollect              EnvelopeType = 196647
	EnvelopeType_BroadCastPlayerActionCollect     EnvelopeType = 196649
	EnvelopeType_SceneDestruction                 EnvelopeType = 196659
	EnvelopeType_BroadCastSceneDestruction        EnvelopeType = 196661
	EnvelopeType_PlayerActionChat                 EnvelopeType = 196663
	EnvelopeType_BroadCastPlayerActionChat        EnvelopeType = 196665
	EnvelopeType_GetHomeData                      EnvelopeType = 196673
	EnvelopeType_QueryServerTime                  EnvelopeType = 196675 // 客户端请求服务器时间
	EnvelopeType_BroadCastChangeService           EnvelopeType = 196677
	EnvelopeType_BroadCastHomeDataInitBatch       EnvelopeType = 196681
	EnvelopeType_BroadCastHomeSkillResult         EnvelopeType = 196689
	EnvelopeType_QueryTalentTree                  EnvelopeType = 196691 // 查询天赋数据
	EnvelopeType_BroadCastTalentTreeUpdate        EnvelopeType = 196693 // 广播更新天赋树数据
	EnvelopeType_BroadCastAnimalSceneDataUpdate   EnvelopeType = 196695 // 广播更新动物场数据
	EnvelopeType_BroadCastAnimalDeath             EnvelopeType = 196697 // 广播动物死亡
	EnvelopeType_BroadCastAnimalAutoProduce       EnvelopeType = 196705 // 广播动物自动生产
	EnvelopeType_BroadCastPickupDrop              EnvelopeType = 196707 // 广播拾取掉落物
	EnvelopeType_AnimalSpecialAction              EnvelopeType = 196709 // 请求动物特殊行为
	EnvelopeType_CaptureAnimal                    EnvelopeType = 196711 // 抓捕Animal(同步可能释放一个 Animal) to scene service
	EnvelopeType_BroadCastEntityMoveSpeed         EnvelopeType = 196713 // 实体速度
	EnvelopeType_BroadCastEntityCaptureDataUpdate EnvelopeType = 196720
	EnvelopeType_ExitCapture                      EnvelopeType = 196721 // 客户端请求退出捕获状态
	EnvelopeType_EntityAccumulateStatus           EnvelopeType = 196722 // 实体蓄力状态请求
	EnvelopeType_BroadCastEntityAccumulateStatus  EnvelopeType = 196723 // 实体蓄力状态广播
	EnvelopeType_EntityDialogue                   EnvelopeType = 196725 // 实体对话
	EnvelopeType_BroadCastEntityEquipmentSkill    EnvelopeType = 196727 // 广播实体装备技能
	EnvelopeType_BroadCastEntityCaptureTrapSkill  EnvelopeType = 196729 // 捕获释放技能结果
	EnvelopeType_BroadCastMapEntityBorn           EnvelopeType = 196737 // 广播地图实体出生
	//taskServer协议 : 0x04ZZZZ  任务服务 协议 ************************************
	EnvelopeType_SelfTasks               EnvelopeType = 262145 // 查询玩家已接的任务列表（进行中状态）
	EnvelopeType_BroadCastUpdateTask     EnvelopeType = 262147 // 任务  进度更新(推送)
	EnvelopeType_BroadCastUpdateTaskList EnvelopeType = 262148 // 任务链进度更新(推送)
	EnvelopeType_AcceptTask              EnvelopeType = 262149 // 领取任务
	EnvelopeType_AcceptTaskList          EnvelopeType = 262151 // 领取任务链任务
	EnvelopeType_AbandonTask             EnvelopeType = 262153 // 放弃任务(任务有保护时间)
	EnvelopeType_AbandonTaskList         EnvelopeType = 262161 // 放弃任务链的任务(任务有保护时间)
	EnvelopeType_TaskReward              EnvelopeType = 262163 // 获取任务奖励(附带提交任务功能)
	EnvelopeType_UpgradeTaskProgress     EnvelopeType = 262165 // 上报更新任务进度
	EnvelopeType_BroadCastTaskReward     EnvelopeType = 262167 // 推送获取的任务奖励
	EnvelopeType_BroadCastTaskListReward EnvelopeType = 262169 // 推送获取的任务链奖励
	//chatServer协议 : 0x05ZZZZ  聊天服务 协议 ************************************
	EnvelopeType_SendChatMessage            EnvelopeType = 327681 // 玩家发送聊天消息
	EnvelopeType_BroadCastChatMessages      EnvelopeType = 327683 // 批量推送的聊天消息
	EnvelopeType_BroadCastRemoveChatMessage EnvelopeType = 327684
	EnvelopeType_BroadCastUpdateChatState   EnvelopeType = 327685
	//agentService协议 : 0x06ZZZZ  网关服务 协议 ************************************
	EnvelopeType_Ping EnvelopeType = 393217
)

// Enum value maps for EnvelopeType.
var (
	EnvelopeType_name = map[int32]string{
		0:      "Unknown",
		39321:  "BroadCastTickOut",
		65537:  "ItemGet",
		65539:  "ItemUse",
		65541:  "ItemDrop",
		65543:  "UpdateAvatar",
		65545:  "UnloadAvatar",
		65553:  "BroadCastItemAdd",
		65554:  "BroadCastItemUpdate",
		65555:  "BroadCastItemDel",
		65561:  "SigninPlayer",
		65569:  "SignOutPlayer",
		65573:  "QueryLands",
		65575:  "BroadCastInitLand",
		65577:  "BroadCastInitItem",
		65585:  "Build",
		65587:  "Recycling",
		65589:  "Charged",
		65591:  "Harvest",
		65593:  "Collection",
		65601:  "BroadCastMultiUpLand",
		65603:  "SelfNftBuilds",
		65605:  "BroadCastSelfBuildUpdate",
		65607:  "BroadCastSelfBuildRecycling",
		65609:  "MintBattery",
		65617:  "QueryGranary",
		65619:  "GranaryCollect",
		65621:  "BroadCastGranaryUpdate",
		65623:  "BroadCastUpGranaryItem",
		65625:  "UpgradeTalentNode",
		65633:  "QueryTalentExp",
		65635:  "BroadCastUpTalentExp",
		65637:  "QueryAnimalList",
		65639:  "BroadCastUpdateAnimalList",
		65641:  "FreedAnimal",
		65649:  "QueryPlayerSetting",
		65651:  "SavePlayerSetting",
		131073: "QueryPlayer",
		131075: "CreatePlayer",
		131077: "Login",
		131079: "BroadCastQueue",
		196609: "EnterMap",
		196611: "UpdateSelfLocation",
		196613: "UseSkill",
		196617: "RespawnPlayer",
		196625: "BroadCastInitMapElement",
		196626: "BroadCastMapEntityUpdate",
		196627: "BroadCastEntityDestroy",
		196628: "BroadCastEntityMove",
		196629: "BroadCastEntityCombat",
		196630: "BroadCastRespawnPlayer",
		196631: "Teleport",
		196633: "BroadCastEntityProfileUpdate",
		196641: "BroadCastMonsterDeath",
		196642: "BroadCastEntityAvatarUpdate",
		196643: "BroadCastEntityBattleDataUpdate",
		196645: "BroadCastEntityBuildDataUpdate",
		196647: "PlayerActionCollect",
		196649: "BroadCastPlayerActionCollect",
		196659: "SceneDestruction",
		196661: "BroadCastSceneDestruction",
		196663: "PlayerActionChat",
		196665: "BroadCastPlayerActionChat",
		196673: "GetHomeData",
		196675: "QueryServerTime",
		196677: "BroadCastChangeService",
		196681: "BroadCastHomeDataInitBatch",
		196689: "BroadCastHomeSkillResult",
		196691: "QueryTalentTree",
		196693: "BroadCastTalentTreeUpdate",
		196695: "BroadCastAnimalSceneDataUpdate",
		196697: "BroadCastAnimalDeath",
		196705: "BroadCastAnimalAutoProduce",
		196707: "BroadCastPickupDrop",
		196709: "AnimalSpecialAction",
		196711: "CaptureAnimal",
		196713: "BroadCastEntityMoveSpeed",
		196720: "BroadCastEntityCaptureDataUpdate",
		196721: "ExitCapture",
		196722: "EntityAccumulateStatus",
		196723: "BroadCastEntityAccumulateStatus",
		196725: "EntityDialogue",
		196727: "BroadCastEntityEquipmentSkill",
		196729: "BroadCastEntityCaptureTrapSkill",
		196737: "BroadCastMapEntityBorn",
		262145: "SelfTasks",
		262147: "BroadCastUpdateTask",
		262148: "BroadCastUpdateTaskList",
		262149: "AcceptTask",
		262151: "AcceptTaskList",
		262153: "AbandonTask",
		262161: "AbandonTaskList",
		262163: "TaskReward",
		262165: "UpgradeTaskProgress",
		262167: "BroadCastTaskReward",
		262169: "BroadCastTaskListReward",
		327681: "SendChatMessage",
		327683: "BroadCastChatMessages",
		327684: "BroadCastRemoveChatMessage",
		327685: "BroadCastUpdateChatState",
		393217: "Ping",
	}
	EnvelopeType_value = map[string]int32{
		"Unknown":                          0,
		"BroadCastTickOut":                 39321,
		"ItemGet":                          65537,
		"ItemUse":                          65539,
		"ItemDrop":                         65541,
		"UpdateAvatar":                     65543,
		"UnloadAvatar":                     65545,
		"BroadCastItemAdd":                 65553,
		"BroadCastItemUpdate":              65554,
		"BroadCastItemDel":                 65555,
		"SigninPlayer":                     65561,
		"SignOutPlayer":                    65569,
		"QueryLands":                       65573,
		"BroadCastInitLand":                65575,
		"BroadCastInitItem":                65577,
		"Build":                            65585,
		"Recycling":                        65587,
		"Charged":                          65589,
		"Harvest":                          65591,
		"Collection":                       65593,
		"BroadCastMultiUpLand":             65601,
		"SelfNftBuilds":                    65603,
		"BroadCastSelfBuildUpdate":         65605,
		"BroadCastSelfBuildRecycling":      65607,
		"MintBattery":                      65609,
		"QueryGranary":                     65617,
		"GranaryCollect":                   65619,
		"BroadCastGranaryUpdate":           65621,
		"BroadCastUpGranaryItem":           65623,
		"UpgradeTalentNode":                65625,
		"QueryTalentExp":                   65633,
		"BroadCastUpTalentExp":             65635,
		"QueryAnimalList":                  65637,
		"BroadCastUpdateAnimalList":        65639,
		"FreedAnimal":                      65641,
		"QueryPlayerSetting":               65649,
		"SavePlayerSetting":                65651,
		"QueryPlayer":                      131073,
		"CreatePlayer":                     131075,
		"Login":                            131077,
		"BroadCastQueue":                   131079,
		"EnterMap":                         196609,
		"UpdateSelfLocation":               196611,
		"UseSkill":                         196613,
		"RespawnPlayer":                    196617,
		"BroadCastInitMapElement":          196625,
		"BroadCastMapEntityUpdate":         196626,
		"BroadCastEntityDestroy":           196627,
		"BroadCastEntityMove":              196628,
		"BroadCastEntityCombat":            196629,
		"BroadCastRespawnPlayer":           196630,
		"Teleport":                         196631,
		"BroadCastEntityProfileUpdate":     196633,
		"BroadCastMonsterDeath":            196641,
		"BroadCastEntityAvatarUpdate":      196642,
		"BroadCastEntityBattleDataUpdate":  196643,
		"BroadCastEntityBuildDataUpdate":   196645,
		"PlayerActionCollect":              196647,
		"BroadCastPlayerActionCollect":     196649,
		"SceneDestruction":                 196659,
		"BroadCastSceneDestruction":        196661,
		"PlayerActionChat":                 196663,
		"BroadCastPlayerActionChat":        196665,
		"GetHomeData":                      196673,
		"QueryServerTime":                  196675,
		"BroadCastChangeService":           196677,
		"BroadCastHomeDataInitBatch":       196681,
		"BroadCastHomeSkillResult":         196689,
		"QueryTalentTree":                  196691,
		"BroadCastTalentTreeUpdate":        196693,
		"BroadCastAnimalSceneDataUpdate":   196695,
		"BroadCastAnimalDeath":             196697,
		"BroadCastAnimalAutoProduce":       196705,
		"BroadCastPickupDrop":              196707,
		"AnimalSpecialAction":              196709,
		"CaptureAnimal":                    196711,
		"BroadCastEntityMoveSpeed":         196713,
		"BroadCastEntityCaptureDataUpdate": 196720,
		"ExitCapture":                      196721,
		"EntityAccumulateStatus":           196722,
		"BroadCastEntityAccumulateStatus":  196723,
		"EntityDialogue":                   196725,
		"BroadCastEntityEquipmentSkill":    196727,
		"BroadCastEntityCaptureTrapSkill":  196729,
		"BroadCastMapEntityBorn":           196737,
		"SelfTasks":                        262145,
		"BroadCastUpdateTask":              262147,
		"BroadCastUpdateTaskList":          262148,
		"AcceptTask":                       262149,
		"AcceptTaskList":                   262151,
		"AbandonTask":                      262153,
		"AbandonTaskList":                  262161,
		"TaskReward":                       262163,
		"UpgradeTaskProgress":              262165,
		"BroadCastTaskReward":              262167,
		"BroadCastTaskListReward":          262169,
		"SendChatMessage":                  327681,
		"BroadCastChatMessages":            327683,
		"BroadCastRemoveChatMessage":       327684,
		"BroadCastUpdateChatState":         327685,
		"Ping":                             393217,
	}
)

func (x EnvelopeType) Enum() *EnvelopeType {
	p := new(EnvelopeType)
	*p = x
	return p
}

func (x EnvelopeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnvelopeType) Descriptor() protoreflect.EnumDescriptor {
	return file_message_cmd_proto_enumTypes[0].Descriptor()
}

func (EnvelopeType) Type() protoreflect.EnumType {
	return &file_message_cmd_proto_enumTypes[0]
}

func (x EnvelopeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EnvelopeType.Descriptor instead.
func (EnvelopeType) EnumDescriptor() ([]byte, []int) {
	return file_message_cmd_proto_rawDescGZIP(), []int{0}
}

var File_message_cmd_proto protoreflect.FileDescriptor

var file_message_cmd_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6d, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x67, 0x61, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x43, 0x6f, 0x72, 0x65, 0x2a, 0x9a, 0x14, 0x0a, 0x0c, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e,
	0x10, 0x00, 0x12, 0x16, 0x0a, 0x10, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73, 0x74, 0x54,
	0x69, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x10, 0x99, 0xb3, 0x02, 0x12, 0x0d, 0x0a, 0x07, 0x49, 0x74,
	0x65, 0x6d, 0x47, 0x65, 0x74, 0x10, 0x81, 0x80, 0x04, 0x12, 0x0d, 0x0a, 0x07, 0x49, 0x74, 0x65,
	0x6d, 0x55, 0x73, 0x65, 0x10, 0x83, 0x80, 0x04, 0x12, 0x0e, 0x0a, 0x08, 0x49, 0x74, 0x65, 0x6d,
	0x44, 0x72, 0x6f, 0x70, 0x10, 0x85, 0x80, 0x04, 0x12, 0x12, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x10, 0x87, 0x80, 0x04, 0x12, 0x12, 0x0a, 0x0c,
	0x55, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x10, 0x89, 0x80, 0x04,
	0x12, 0x16, 0x0a, 0x10, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73, 0x74, 0x49, 0x74, 0x65,
	0x6d, 0x41, 0x64, 0x64, 0x10, 0x91, 0x80, 0x04, 0x12, 0x19, 0x0a, 0x13, 0x42, 0x72, 0x6f, 0x61,
	0x64, 0x43, 0x61, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x10,
	0x92, 0x80, 0x04, 0x12, 0x16, 0x0a, 0x10, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x44, 0x65, 0x6c, 0x10, 0x93, 0x80, 0x04, 0x12, 0x12, 0x0a, 0x0c, 0x53,
	0x69, 0x67, 0x6e, 0x69, 0x6e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x10, 0x99, 0x80, 0x04, 0x12,
	0x13, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x75, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x10, 0xa1, 0x80, 0x04, 0x12, 0x10, 0x0a, 0x0a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4c, 0x61, 0x6e,
	0x64, 0x73, 0x10, 0xa5, 0x80, 0x04, 0x12, 0x17, 0x0a, 0x11, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43,
	0x61, 0x73, 0x74, 0x49, 0x6e, 0x69, 0x74, 0x4c, 0x61, 0x6e, 0x64, 0x10, 0xa7, 0x80, 0x04, 0x12,
	0x17, 0x0a, 0x11, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73, 0x74, 0x49, 0x6e, 0x69, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x10, 0xa9, 0x80, 0x04, 0x12, 0x0b, 0x0a, 0x05, 0x42, 0x75, 0x69, 0x6c,
	0x64, 0x10, 0xb1, 0x80, 0x04, 0x12, 0x0f, 0x0a, 0x09, 0x52, 0x65, 0x63, 0x79, 0x63, 0x6c, 0x69,
	0x6e, 0x67, 0x10, 0xb3, 0x80, 0x04, 0x12, 0x0d, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65,
	0x64, 0x10, 0xb5, 0x80, 0x04, 0x12, 0x0d, 0x0a, 0x07, 0x48, 0x61, 0x72, 0x76, 0x65, 0x73, 0x74,
	0x10, 0xb7, 0x80, 0x04, 0x12, 0x10, 0x0a, 0x0a, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x10, 0xb9, 0x80, 0x04, 0x12, 0x1a, 0x0a, 0x14, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43,
	0x61, 0x73, 0x74, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x55, 0x70, 0x4c, 0x61, 0x6e, 0x64, 0x10, 0xc1,
	0x80, 0x04, 0x12, 0x13, 0x0a, 0x0d, 0x53, 0x65, 0x6c, 0x66, 0x4e, 0x66, 0x74, 0x42, 0x75, 0x69,
	0x6c, 0x64, 0x73, 0x10, 0xc3, 0x80, 0x04, 0x12, 0x1e, 0x0a, 0x18, 0x42, 0x72, 0x6f, 0x61, 0x64,
	0x43, 0x61, 0x73, 0x74, 0x53, 0x65, 0x6c, 0x66, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x10, 0xc5, 0x80, 0x04, 0x12, 0x21, 0x0a, 0x1b, 0x42, 0x72, 0x6f, 0x61, 0x64,
	0x43, 0x61, 0x73, 0x74, 0x53, 0x65, 0x6c, 0x66, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x63,
	0x79, 0x63, 0x6c, 0x69, 0x6e, 0x67, 0x10, 0xc7, 0x80, 0x04, 0x12, 0x11, 0x0a, 0x0b, 0x4d, 0x69,
	0x6e, 0x74, 0x42, 0x61, 0x74, 0x74, 0x65, 0x72, 0x79, 0x10, 0xc9, 0x80, 0x04, 0x12, 0x12, 0x0a,
	0x0c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x47, 0x72, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x10, 0xd1, 0x80,
	0x04, 0x12, 0x14, 0x0a, 0x0e, 0x47, 0x72, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x10, 0xd3, 0x80, 0x04, 0x12, 0x1c, 0x0a, 0x16, 0x42, 0x72, 0x6f, 0x61, 0x64,
	0x43, 0x61, 0x73, 0x74, 0x47, 0x72, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x10, 0xd5, 0x80, 0x04, 0x12, 0x1c, 0x0a, 0x16, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61,
	0x73, 0x74, 0x55, 0x70, 0x47, 0x72, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x10,
	0xd7, 0x80, 0x04, 0x12, 0x17, 0x0a, 0x11, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x54, 0x61,
	0x6c, 0x65, 0x6e, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x10, 0xd9, 0x80, 0x04, 0x12, 0x14, 0x0a, 0x0e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x45, 0x78, 0x70, 0x10, 0xe1,
	0x80, 0x04, 0x12, 0x1a, 0x0a, 0x14, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73, 0x74, 0x55,
	0x70, 0x54, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x45, 0x78, 0x70, 0x10, 0xe3, 0x80, 0x04, 0x12, 0x15,
	0x0a, 0x0f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x4c, 0x69, 0x73,
	0x74, 0x10, 0xe5, 0x80, 0x04, 0x12, 0x1f, 0x0a, 0x19, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61,
	0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x4c, 0x69,
	0x73, 0x74, 0x10, 0xe7, 0x80, 0x04, 0x12, 0x11, 0x0a, 0x0b, 0x46, 0x72, 0x65, 0x65, 0x64, 0x41,
	0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x10, 0xe9, 0x80, 0x04, 0x12, 0x18, 0x0a, 0x12, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x10,
	0xf1, 0x80, 0x04, 0x12, 0x17, 0x0a, 0x11, 0x53, 0x61, 0x76, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x10, 0xf3, 0x80, 0x04, 0x12, 0x11, 0x0a, 0x0b,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x10, 0x81, 0x80, 0x08, 0x12,
	0x12, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x10,
	0x83, 0x80, 0x08, 0x12, 0x0b, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x10, 0x85, 0x80, 0x08,
	0x12, 0x14, 0x0a, 0x0e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73, 0x74, 0x51, 0x75, 0x65,
	0x75, 0x65, 0x10, 0x87, 0x80, 0x08, 0x12, 0x0e, 0x0a, 0x08, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x4d,
	0x61, 0x70, 0x10, 0x81, 0x80, 0x0c, 0x12, 0x18, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x65, 0x6c, 0x66, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x83, 0x80, 0x0c,
	0x12, 0x0e, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x10, 0x85, 0x80, 0x0c,
	0x12, 0x13, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x70, 0x61, 0x77, 0x6e, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x10, 0x89, 0x80, 0x0c, 0x12, 0x1d, 0x0a, 0x17, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61,
	0x73, 0x74, 0x49, 0x6e, 0x69, 0x74, 0x4d, 0x61, 0x70, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x10, 0x91, 0x80, 0x0c, 0x12, 0x1e, 0x0a, 0x18, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73,
	0x74, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x10, 0x92, 0x80, 0x0c, 0x12, 0x1c, 0x0a, 0x16, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73,
	0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x44, 0x65, 0x73, 0x74, 0x72, 0x6f, 0x79, 0x10, 0x93,
	0x80, 0x0c, 0x12, 0x19, 0x0a, 0x13, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73, 0x74, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x4d, 0x6f, 0x76, 0x65, 0x10, 0x94, 0x80, 0x0c, 0x12, 0x1b, 0x0a,
	0x15, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x43, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x10, 0x95, 0x80, 0x0c, 0x12, 0x1c, 0x0a, 0x16, 0x42, 0x72,
	0x6f, 0x61, 0x64, 0x43, 0x61, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x61, 0x77, 0x6e, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x10, 0x96, 0x80, 0x0c, 0x12, 0x0e, 0x0a, 0x08, 0x54, 0x65, 0x6c, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x10, 0x97, 0x80, 0x0c, 0x12, 0x22, 0x0a, 0x1c, 0x42, 0x72, 0x6f, 0x61,
	0x64, 0x43, 0x61, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x10, 0x99, 0x80, 0x0c, 0x12, 0x1b, 0x0a, 0x15,
	0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73, 0x74, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72,
	0x44, 0x65, 0x61, 0x74, 0x68, 0x10, 0xa1, 0x80, 0x0c, 0x12, 0x21, 0x0a, 0x1b, 0x42, 0x72, 0x6f,
	0x61, 0x64, 0x43, 0x61, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x41, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x10, 0xa2, 0x80, 0x0c, 0x12, 0x25, 0x0a, 0x1f,
	0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x42,
	0x61, 0x74, 0x74, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x10,
	0xa3, 0x80, 0x0c, 0x12, 0x24, 0x0a, 0x1e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73, 0x74,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x44, 0x61, 0x74, 0x61, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x10, 0xa5, 0x80, 0x0c, 0x12, 0x19, 0x0a, 0x13, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x10, 0xa7, 0x80, 0x0c, 0x12, 0x22, 0x0a, 0x1c, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73,
	0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x10, 0xa9, 0x80, 0x0c, 0x12, 0x16, 0x0a, 0x10, 0x53, 0x63, 0x65, 0x6e,
	0x65, 0x44, 0x65, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0xb3, 0x80, 0x0c,
	0x12, 0x1f, 0x0a, 0x19, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73, 0x74, 0x53, 0x63, 0x65,
	0x6e, 0x65, 0x44, 0x65, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0xb5, 0x80,
	0x0c, 0x12, 0x16, 0x0a, 0x10, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x68, 0x61, 0x74, 0x10, 0xb7, 0x80, 0x0c, 0x12, 0x1f, 0x0a, 0x19, 0x42, 0x72, 0x6f,
	0x61, 0x64, 0x43, 0x61, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x68, 0x61, 0x74, 0x10, 0xb9, 0x80, 0x0c, 0x12, 0x11, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x48, 0x6f, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x61, 0x10, 0xc1, 0x80, 0x0c, 0x12, 0x15, 0x0a,
	0x0f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65,
	0x10, 0xc3, 0x80, 0x0c, 0x12, 0x1c, 0x0a, 0x16, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73,
	0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x10, 0xc5,
	0x80, 0x0c, 0x12, 0x20, 0x0a, 0x1a, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73, 0x74, 0x48,
	0x6f, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x61, 0x49, 0x6e, 0x69, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x10, 0xc9, 0x80, 0x0c, 0x12, 0x1e, 0x0a, 0x18, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73,
	0x74, 0x48, 0x6f, 0x6d, 0x65, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x10, 0xd1, 0x80, 0x0c, 0x12, 0x15, 0x0a, 0x0f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x61, 0x6c,
	0x65, 0x6e, 0x74, 0x54, 0x72, 0x65, 0x65, 0x10, 0xd3, 0x80, 0x0c, 0x12, 0x1f, 0x0a, 0x19, 0x42,
	0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73, 0x74, 0x54, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x54, 0x72,
	0x65, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x10, 0xd5, 0x80, 0x0c, 0x12, 0x24, 0x0a, 0x1e,
	0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73, 0x74, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x53,
	0x63, 0x65, 0x6e, 0x65, 0x44, 0x61, 0x74, 0x61, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x10, 0xd7,
	0x80, 0x0c, 0x12, 0x1a, 0x0a, 0x14, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73, 0x74, 0x41,
	0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x44, 0x65, 0x61, 0x74, 0x68, 0x10, 0xd9, 0x80, 0x0c, 0x12, 0x20,
	0x0a, 0x1a, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73, 0x74, 0x41, 0x6e, 0x69, 0x6d, 0x61,
	0x6c, 0x41, 0x75, 0x74, 0x6f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x10, 0xe1, 0x80, 0x0c,
	0x12, 0x19, 0x0a, 0x13, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73, 0x74, 0x50, 0x69, 0x63,
	0x6b, 0x75, 0x70, 0x44, 0x72, 0x6f, 0x70, 0x10, 0xe3, 0x80, 0x0c, 0x12, 0x19, 0x0a, 0x13, 0x41,
	0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x53, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x10, 0xe5, 0x80, 0x0c, 0x12, 0x13, 0x0a, 0x0d, 0x43, 0x61, 0x70, 0x74, 0x75, 0x72,
	0x65, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x10, 0xe7, 0x80, 0x0c, 0x12, 0x1e, 0x0a, 0x18, 0x42,
	0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4d, 0x6f,
	0x76, 0x65, 0x53, 0x70, 0x65, 0x65, 0x64, 0x10, 0xe9, 0x80, 0x0c, 0x12, 0x26, 0x0a, 0x20, 0x42,
	0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x43, 0x61,
	0x70, 0x74, 0x75, 0x72, 0x65, 0x44, 0x61, 0x74, 0x61, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x10,
	0xf0, 0x80, 0x0c, 0x12, 0x11, 0x0a, 0x0b, 0x45, 0x78, 0x69, 0x74, 0x43, 0x61, 0x70, 0x74, 0x75,
	0x72, 0x65, 0x10, 0xf1, 0x80, 0x0c, 0x12, 0x1c, 0x0a, 0x16, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x41, 0x63, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x10, 0xf2, 0x80, 0x0c, 0x12, 0x25, 0x0a, 0x1f, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73,
	0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x41, 0x63, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x10, 0xf3, 0x80, 0x0c, 0x12, 0x14, 0x0a, 0x0e, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x75, 0x65, 0x10, 0xf5, 0x80,
	0x0c, 0x12, 0x23, 0x0a, 0x1d, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73, 0x74, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x6b, 0x69,
	0x6c, 0x6c, 0x10, 0xf7, 0x80, 0x0c, 0x12, 0x25, 0x0a, 0x1f, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43,
	0x61, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x43, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65,
	0x54, 0x72, 0x61, 0x70, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x10, 0xf9, 0x80, 0x0c, 0x12, 0x1c, 0x0a,
	0x16, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73, 0x74, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x42, 0x6f, 0x72, 0x6e, 0x10, 0x81, 0x81, 0x0c, 0x12, 0x0f, 0x0a, 0x09, 0x53,
	0x65, 0x6c, 0x66, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x10, 0x81, 0x80, 0x10, 0x12, 0x19, 0x0a, 0x13,
	0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x61, 0x73, 0x6b, 0x10, 0x83, 0x80, 0x10, 0x12, 0x1d, 0x0a, 0x17, 0x42, 0x72, 0x6f, 0x61, 0x64,
	0x43, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69,
	0x73, 0x74, 0x10, 0x84, 0x80, 0x10, 0x12, 0x10, 0x0a, 0x0a, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74,
	0x54, 0x61, 0x73, 0x6b, 0x10, 0x85, 0x80, 0x10, 0x12, 0x14, 0x0a, 0x0e, 0x41, 0x63, 0x63, 0x65,
	0x70, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x10, 0x87, 0x80, 0x10, 0x12, 0x11,
	0x0a, 0x0b, 0x41, 0x62, 0x61, 0x6e, 0x64, 0x6f, 0x6e, 0x54, 0x61, 0x73, 0x6b, 0x10, 0x89, 0x80,
	0x10, 0x12, 0x15, 0x0a, 0x0f, 0x41, 0x62, 0x61, 0x6e, 0x64, 0x6f, 0x6e, 0x54, 0x61, 0x73, 0x6b,
	0x4c, 0x69, 0x73, 0x74, 0x10, 0x91, 0x80, 0x10, 0x12, 0x10, 0x0a, 0x0a, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x10, 0x93, 0x80, 0x10, 0x12, 0x19, 0x0a, 0x13, 0x55, 0x70,
	0x67, 0x72, 0x61, 0x64, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x10, 0x95, 0x80, 0x10, 0x12, 0x19, 0x0a, 0x13, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61,
	0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x10, 0x97, 0x80, 0x10,
	0x12, 0x1d, 0x0a, 0x17, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73, 0x74, 0x54, 0x61, 0x73,
	0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x10, 0x99, 0x80, 0x10, 0x12,
	0x15, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x10, 0x81, 0x80, 0x14, 0x12, 0x1b, 0x0a, 0x15, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43,
	0x61, 0x73, 0x74, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x10,
	0x83, 0x80, 0x14, 0x12, 0x20, 0x0a, 0x1a, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73, 0x74,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x10, 0x84, 0x80, 0x14, 0x12, 0x1e, 0x0a, 0x18, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61,
	0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x10, 0x85, 0x80, 0x14, 0x12, 0x0a, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x10, 0x81, 0x80,
	0x18, 0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x47, 0x6f,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_cmd_proto_rawDescOnce sync.Once
	file_message_cmd_proto_rawDescData = file_message_cmd_proto_rawDesc
)

func file_message_cmd_proto_rawDescGZIP() []byte {
	file_message_cmd_proto_rawDescOnce.Do(func() {
		file_message_cmd_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_cmd_proto_rawDescData)
	})
	return file_message_cmd_proto_rawDescData
}

var file_message_cmd_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_message_cmd_proto_goTypes = []interface{}{
	(EnvelopeType)(0), // 0: gameMessageCore.EnvelopeType
}
var file_message_cmd_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_message_cmd_proto_init() }
func file_message_cmd_proto_init() {
	if File_message_cmd_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_message_cmd_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_cmd_proto_goTypes,
		DependencyIndexes: file_message_cmd_proto_depIdxs,
		EnumInfos:         file_message_cmd_proto_enumTypes,
	}.Build()
	File_message_cmd_proto = out.File
	file_message_cmd_proto_rawDesc = nil
	file_message_cmd_proto_goTypes = nil
	file_message_cmd_proto_depIdxs = nil
}
