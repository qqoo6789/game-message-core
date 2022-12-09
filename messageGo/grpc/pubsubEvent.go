package grpc

type SubscriptionEvent string

const (
	SubscriptionEventServiceStarted    SubscriptionEvent = "EventTopicServiceStarted"
	SubscriptionEventServiceUnregister SubscriptionEvent = "EventTopicServiceUnregister"

	SubscriptionEventUserApplyEnterServiceRes SubscriptionEvent = "EventTopicUserApplyEnterServiceRes"
	SubscriptionEventUserJoinServiceRes       SubscriptionEvent = "EventTopicUserJoinServiceRes"
	SubscriptionEventUserChangeService        SubscriptionEvent = "EventTopicUserChangeService"

	SubscriptionEventSavePlayerData   SubscriptionEvent = "EventTopicSavePlayerData"
	SubscriptionEventKillMonster      SubscriptionEvent = "EventTopicKillMonster"
	SubscriptionEventPlayerDeath      SubscriptionEvent = "EventTopicPlayerDeath"
	SubscriptionEventUseNFT           SubscriptionEvent = "EventTopicUseNFT"
	SubscriptionEventUserEnterGame    SubscriptionEvent = "EventTopicUserEnterGame"
	SubscriptionEventUserLeaveGame    SubscriptionEvent = "EventTopicUserLeaveGame"
	SubscriptionEventUserTaskReward   SubscriptionEvent = "EventTopicUserTaskReward"
	SubscriptionEventTaskFinish       SubscriptionEvent = "EventTopicTaskFinish"
	SubscriptionEventTaskListFinish   SubscriptionEvent = "EventTopicTaskListFinish"
	SubscriptionEventSlotLevelUpgrade SubscriptionEvent = "EventTopicSlotLevelUpgrade"
	SubscriptionEventUserLevelUpgrade SubscriptionEvent = "EventTopicUserLevelUpgrade"

	// user nft build events
	SubscriptionEventNftBuildAdd    SubscriptionEvent = "EventTopicNftBuildAdd"
	SubscriptionEventNftBuildUpdate SubscriptionEvent = "EventTopicNftBuildUpdate"
	SubscriptionEventNftBuildRemove SubscriptionEvent = "EventTopicNftBuildRemove"

	SubscriptionEventTickOutPlayer SubscriptionEvent = "EventTopicTickOutPlayer"

	SubscriptionEventSaveHomeData SubscriptionEvent = "EventTopicSaveHomeData"
)
