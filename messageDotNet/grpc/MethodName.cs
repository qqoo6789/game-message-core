
/// <summary>
/// MELAND service call method API name
/// </summary>
internal static class ManagerServiceAction
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
    public const string DESTROY_SERVICE = "ManagerActionSelectService";


}

/// <summary>
/// MELAND proto message Action calls method name 
/// </summary>
internal static class ProtoMessageAction
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
internal static class UserAction
{
    /// <summary>
    /// 玩家离开游戏(断开网络链接时)
    /// </summary>
    public const string USER_LEAVE_GAME = "UserLeaveGame";
}