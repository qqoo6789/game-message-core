
/// <summary>
/// MELAND service call method API name
/// </summary>
public static class ManagerServiceAction
{
    /// <summary>
    /// 获取服务时间戳
    /// </summary>
    public const string SERVICE_TIMESTAMP = "QueryServerTime";
    /// <summary>
    /// 注册服务
    /// </summary>
    public const string REGISTER_SERVICE = "ManagerActionServiceRegister";
    /// <summary>
    /// 释放服务
    /// </summary>
    public const string DESTROY_SERVICE = "ManagerActionServiceDestroy";
    /// <summary>
    /// 查询服务信息
    /// </summary>
    public const string SELECT_SERVICE = "ManagerActionSelectService";


}

/// <summary>
/// MELAND proto message Action calls method name 
/// </summary>
public static class ProtoMessageAction
{
    /// <summary>
    /// 转发客户端proto Message
    /// </summary>
    public const string PULL_CLIENT_MESSAGE = "PullClientMessage";
    /// <summary>
    /// services 推送 proto message 到客户端 
    /// </summary>
    public const string BROAD_CAST_TO_CLIENT = "BroadCastToClient";
    /// <summary>
    /// services 批量推送 proto message 到客户端 
    /// </summary>
    public const string MULTIPLE_BROAD_CAST_TO_CLIENT = "MultipleBroadCastToClient";
    /// <summary>
    /// 玩家离开游戏(断开网络链接时)
    /// </summary>
    public const string PLAYER_LEAVE_GAME = "PlayerLeaveGame";
}

/// <summary>
/// MELAND user action method api
/// </summary>
public static class UserAction
{
    /// <summary>
    /// 更新玩家战斗属性(升级/装备槽等级变更)
    /// </summary>
    public const string UPDATE_USER_PROFILE = "UpdateUserProfile";
    /// <summary>
    /// 更新玩家使用的装备
    /// </summary>
    public const string UPDATE_USED_AVATAR = "UpdateUsedAvatar";
    /// <summary>
    /// 玩家离开游戏(断开网络链接时)
    /// </summary>
    public const string USER_LEAVE_GAME = "UserLeaveGame";

    /// <summary>
    ///  查询玩家详细数据 for main service
    /// </summary>
    public const string GET_USER_DATA = "GetUserData";

}