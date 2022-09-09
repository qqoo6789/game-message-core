package grpc

type SubscriptionEvent string

const SubscriptionEventServiceRegister SubscriptionEvent = "RegisterService"
const SubscriptionEventServiceDestroy SubscriptionEvent = "DestroyService"
