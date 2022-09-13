
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
}



/// <summary>
/// MELAND Game Service Action call method API name
/// </summary>
internal static class GameServiceAction
{
    /// <summary>
    /// 转发客户端proto Message
    /// </summary>
    public const string CLIENT_PROTO_MESSAGE = "ClientProtoMessage";
    /// <summary>
    /// 玩家离开游戏(断开网络链接时)
    /// </summary>
    public const string PLAYER_LEAVE_GAME = "PlayerLeaveGame";
}