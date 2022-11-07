package grpc

type SubscriptionEvent string

const (
	SubscriptionEventServiceUnregister SubscriptionEvent = "EventTopicServiceUnregister"

	SubscriptionEventSavePlayerData     SubscriptionEvent = "EventTopicSavePlayerData"
	SubscriptionEventKillMonster        SubscriptionEvent = "EventTopicKillMonster"
	SubscriptionEventPlayerDeath        SubscriptionEvent = "EventTopicPlayerDeath"
	SubscriptionEventUseNFT             SubscriptionEvent = "EventTopicUseNFT"
	SubscriptionEventUserEnterGame      SubscriptionEvent = "EventTopicUserEnterGame"
	SubscriptionEventUserLeaveGame      SubscriptionEvent = "EventTopicUserLeaveGame"
	SubscriptionEventUserTaskReward     SubscriptionEvent = "EventTopicUserTaskReward"
	SubscriptionEventUserTaskListFinish SubscriptionEvent = "EventTopicTaskListFinish"

	// user nft build events
	SubscriptionEventNftBuildAdd    SubscriptionEvent = "EventTopicNftBuildAdd"
	SubscriptionEventNftBuildUpdate SubscriptionEvent = "EventTopicNftBuildUpdate"
	SubscriptionEventNftBuildRemove SubscriptionEvent = "EventTopicNftBuildRemove"
)
