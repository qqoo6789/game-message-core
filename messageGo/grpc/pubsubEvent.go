package grpc

type SubscriptionEvent string

const (
	SubscriptionEventSavePlayerData SubscriptionEvent = "EventTopicSavePlayerData"
	SubscriptionEventPickDrop       SubscriptionEvent = "EventTopicPickDrop"
	SubscriptionEventKillMonster    SubscriptionEvent = "EventTopicKillMonster"
	SubscriptionEventPlayerDeath    SubscriptionEvent = "EventTopicPlayerDeath"
)
