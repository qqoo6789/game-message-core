package grpc

type ManagerServiceAction string

const (
	ManagerServiceActionServiceTime ManagerServiceAction = "QueryServerTime"
	ManagerServiceActionRegister    ManagerServiceAction = "ManagerActionServiceRegister"
	ManagerServiceActionDestroy     ManagerServiceAction = "ManagerActionServiceDestroy"
)

type ProtoMessageAction string

const (
	ProtoMessageActionPullClientMessage         ProtoMessageAction = "PullClientMessage"
	ProtoMessageActionBroadCastToClient         ProtoMessageAction = "BroadCastToClient"
	ProtoMessageActionMultipleBroadCastToClient ProtoMessageAction = "MultipleBroadCastToClient"
)

type UserAction string

const (
	UserActionLeaveGame UserAction = "UserLeaveGame"
)
