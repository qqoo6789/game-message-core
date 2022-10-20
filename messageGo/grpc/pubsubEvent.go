package grpc

type SubscriptionEvent string

const (
	SubscriptionEventSavePlayerData SubscriptionEvent = "EventTopicSavePlayerData"
	SubscriptionEventKillMonster    SubscriptionEvent = "EventTopicKillMonster"
	SubscriptionEventPlayerDeath    SubscriptionEvent = "EventTopicPlayerDeath"
	SubscriptionEventUseNFT         SubscriptionEvent = "EventTopicUseNFT"
	SubscriptionEventUserEnterGame  SubscriptionEvent = "EventTopicUserEnterGame"
	SubscriptionEventUserTaskReward SubscriptionEvent = "EventTopicUserTaskReward"

	// user nft build events
	SubscriptionEventNftBuildAdd    SubscriptionEvent = "EventTopicNftBuildAdd"
	SubscriptionEventNftBuildUpdate SubscriptionEvent = "EventTopicNftBuildUpdate"
	SubscriptionEventNftBuildRemove SubscriptionEvent = "EventTopicNftBuildRemove"
)
