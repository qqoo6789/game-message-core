package grpc

type ManagerServiceAction string

const (
	ManagerServiceActionServiceTime   ManagerServiceAction = "QueryServerTime"
	ManagerServiceActionRegister      ManagerServiceAction = "ManagerActionServiceRegister"
	ManagerServiceActionDestroy       ManagerServiceAction = "ManagerActionServiceDestroy"
	ManagerServiceActionSelectService ManagerServiceAction = "ManagerActionSelectService"
)

type ProtoMessageAction string

const (
	ProtoMessageActionPullClientMessage         ProtoMessageAction = "PullClientMessage"
	ProtoMessageActionBroadCastToClient         ProtoMessageAction = "BroadCastToClient"
	ProtoMessageActionMultipleBroadCastToClient ProtoMessageAction = "MultipleBroadCastToClient"
)

type UserAction string

const (
	UserActionUpdateUsedAvatar UserAction = "UpdateUsedAvatar"
	UserActionLeaveGame        UserAction = "UserLeaveGame"
)
