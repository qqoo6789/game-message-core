package grpc

type AppId string

const (
	GAME_SERVICE_APPID_MANAGER    AppId = "game-service-manager"
	GAME_SERVICE_APPID_MAIN       AppId = "game-service-main"
	GAME_SERVICE_APPID_ACCOUNT    AppId = "game-service-account"
	GAME_SERVICE_APPID_TASK       AppId = "game-service-task"
	GAME_SERVICE_APPID_CHAT       AppId = "game-service-chat"
	GAME_SERVICE_APPID_CONTROLLER AppId = "game-service-controller"
	GAME_SERVICE_APPID_LOG        AppId = "game-service-log"
	GAME_SERVICE_APPID_SCENE_XX   AppId = "game-service-scene-xx" // xx为动态ID
	GAME_SERVICE_APPID_AGENT_XX   AppId = "game-service-agent-xx" // xx为动态ID
)
