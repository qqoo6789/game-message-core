package grpc

type EventExpire uint32

const (
	EVENT_EXPIRE_DEFAULT EventExpire = 0
	EVENT_EXPIRE_SEC_5   EventExpire = 1 * 5
	EVENT_EXPIRE_SEC_30  EventExpire = 1 * 30
	EVENT_EXPIRE_MIN_1   EventExpire = EVENT_EXPIRE_SEC_30 * 2
	EVENT_EXPIRE_HOUR_1  EventExpire = EVENT_EXPIRE_MIN_1 * 60
	EVENT_EXPIRE_HOUR_12 EventExpire = EVENT_EXPIRE_HOUR_1 * 12
	EVENT_EXPIRE_DAY     EventExpire = EVENT_EXPIRE_HOUR_1 * 24
	EVENT_EXPIRE_WEEK    EventExpire = EVENT_EXPIRE_DAY * 7
)

type SubscriptionEvent string

const (
	// 动态读取配置数据
	SubscriptionEventReloadConfigData SubscriptionEvent = "EventTopicReloadConfigData"

	// 服务注册｜注销｜迁移
	SubscriptionEventServiceUnregister SubscriptionEvent = "EventTopicServiceUnregister"

	// 无缝切换地图
	SubscriptionEventUserChangeService SubscriptionEvent = "EventTopicUserChangeService"

	// 清理指定userId 缓存数据
	SubscriptionEventClearUserCache SubscriptionEvent = "EventTopicClearUserCache"

	SubscriptionEventMintNFT        SubscriptionEvent = "EventTopicMintNFT"
	SubscriptionEventMultiMintNFT   SubscriptionEvent = "EventTopicMultiMintNFT"
	SubscriptionEventSavePlayerData SubscriptionEvent = "EventTopicSavePlayerData"
	SubscriptionEventKillMonster    SubscriptionEvent = "EventTopicKillMonster"
	SubscriptionEventPlayerDeath    SubscriptionEvent = "EventTopicPlayerDeath"
	SubscriptionEventUseNFT         SubscriptionEvent = "EventTopicUseNFT"
	SubscriptionEventGetNft         SubscriptionEvent = "EventTopicGetNFT"

	SubscriptionEventCreateUser     SubscriptionEvent = "EventTopicCreateUser"
	SubscriptionEventUserEnterGame  SubscriptionEvent = "EventTopicUserEnterGame"
	SubscriptionEventUserLeaveGame  SubscriptionEvent = "EventTopicUserLeaveGame"
	SubscriptionEventTaskFinish     SubscriptionEvent = "EventTopicTaskFinish"
	SubscriptionEventTaskListFinish SubscriptionEvent = "EventTopicTaskListFinish"
	SubscriptionEventAcceptTask     SubscriptionEvent = "EventTopicAcceptTask"
	SubscriptionEventAbandonTask    SubscriptionEvent = "EventTopicAbandonTask"

	//《invalid event》
	SubscriptionEventSlotLevelUpgrade SubscriptionEvent = "EventTopicSlotLevelUpgrade"

	// user nft build events
	SubscriptionEventNftBuildAdd    SubscriptionEvent = "EventTopicNftBuildAdd"
	SubscriptionEventNftBuildUpdate SubscriptionEvent = "EventTopicNftBuildUpdate"
	SubscriptionEventNftBuildRemove SubscriptionEvent = "EventTopicNftBuildRemove"

	// home Granary Stockpile Event
	SubscriptionEventGranaryStockpile SubscriptionEvent = "EventTopicGranaryStockpile"

	// update user talent data event
	SubscriptionEventUpdateTalent  SubscriptionEvent = "EventTopicUpdateTalent"
	SubscriptionEventAddTalentExp  SubscriptionEvent = "EventTopicAddTalentExp"
	SubscriptionEventTakeTalentExp SubscriptionEvent = "EventTopicTakeTalentExp"

	// Capture & FREED Animal Event
	SubscriptionEventFreedAnimal   SubscriptionEvent = "EventTopicFreedAnimal"
	SubscriptionEventCaptureAnimal SubscriptionEvent = "EventTopicCaptureAnimal"
)
