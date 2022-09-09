package grpc

type SubscriptionEvent string

const SubscriptionEventServiceRegister SubscriptionEvent = "EventTopicRegisterService"
const SubscriptionEventServiceDestroy SubscriptionEvent = "EventTopicDestroyService"
const SubscriptionEventSavePlayerData SubscriptionEvent = "EventTopicSavePlayerData"
const SubscriptionEventPickDrop SubscriptionEvent = "EventTopicPickDrop"
const SubscriptionEventKillMonster SubscriptionEvent = "EventTopicKillMonster"
const SubscriptionEventPlayerDeath SubscriptionEvent = "EventTopicPlayerDeath"
