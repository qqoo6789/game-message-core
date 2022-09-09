package grpc

type ManagerServiceAction string

const (
	ManagerServiceActionServiceTime ManagerServiceAction = "QueryServerTime"
	ManagerServiceActionRegister    ManagerServiceAction = "ManagerActionServiceRegister"
	ManagerServiceActionDestroy     ManagerServiceAction = "ManagerActionServiceDestroy"
)
