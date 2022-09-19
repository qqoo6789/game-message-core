package grpc

type SubscriptionEvent string

const (
	SubscriptionEventSavePlayerData SubscriptionEvent = "EventTopicSavePlayerData"
	SubscriptionEventKillMonster    SubscriptionEvent = "EventTopicKillMonster"
	SubscriptionEventPlayerDeath    SubscriptionEvent = "EventTopicPlayerDeath"
	SubscriptionEventUseNFT         SubscriptionEvent = "EventTopicUseNFT"
)
