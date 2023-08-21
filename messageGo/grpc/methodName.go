package grpc

const (
	RpcInvokeBroadSystemMessage string = "RpcInvokeBroadSystemMessage"
	RpcInvokeCheckUserLogin     string = "RpcInvokeCheckUserLogin"
	RpcInvokeTickOutPlayer      string = "RpcInvokeTickOutPlayer"
	RpcInvokeGetTeamData        string = "RpcInvokeGetTeamData"
)

type ManagerServiceAction string

const (
	ManagerServiceActionServiceTime        ManagerServiceAction = "QueryServerTime"
	ManagerServiceActionRegister           ManagerServiceAction = "ManagerActionServiceRegister"
	ManagerServiceActionStartService       ManagerServiceAction = "ManagerActionStartService"
	ManagerServiceActionSelectService      ManagerServiceAction = "ManagerActionSelectService"
	ManagerServiceActionMultiSelectService ManagerServiceAction = "ManagerActionMultiSelectService"
)

type MainServiceAction string

const (
	MainServiceActionUseNFT             MainServiceAction = "MainServiceActionUseNFT"
	MainServiceActionTakeNFT            MainServiceAction = "MainServiceActionTakeNFT"
	MainServiceActionGetAllBuild        MainServiceAction = "MainServiceActionGetAllBuild"
	MainServiceActionGetHomeData        MainServiceAction = "MainServiceActionGetHomeData"
	MainServiceActionSaveHomeData       MainServiceAction = "MainServiceActionSaveHomeData"
	MainServiceActionUpHomeLastSaveTime MainServiceAction = "MainServiceActionUpHomeLastSaveTime"

	MultiUpdateAnimalBaseData MainServiceAction = "MainServiceActionMultiUpdateAnimalBaseData"
	GetUserAnimalList         MainServiceAction = "MainServiceActionGetUserAnimalList"
	CaptureAnimal             MainServiceAction = "MainServiceActionCaptureAnimal"
)

type ProtoMessageAction string

const (
	ProtoMessageActionPullClientMessage         ProtoMessageAction = "PullClientMessage"
	ProtoMessageActionBroadCastToClient         ProtoMessageAction = "BroadCastToClient"
	ProtoMessageActionMultipleBroadCastToClient ProtoMessageAction = "MultipleBroadCastToClient"
)

type UserAction string

const (
	UserActionGetUserData       UserAction = "GetUserData"
	UserActionUpdateUsedAvatar  UserAction = "UpdateUsedAvatar"
	UserActionUpdateUserProfile UserAction = "UpdateUserProfile"
)

type ChangeServiceAction string

const (
	// user请求进入目标scene service（预进入阶段）
	ChangeServiceActionApplyEnterService ChangeServiceAction = "ChangeServiceActionApplyEnterService"
	// user正式进入目标scene service（正式进入阶段）
	ChangeServiceActionJoinService ChangeServiceAction = "ChangeServiceActionJoinService"
)
