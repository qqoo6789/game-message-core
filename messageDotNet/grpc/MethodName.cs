

/// <summary>
/// service method API name
/// </summary>
public static class RpcInvoke
{
    public const string TICK_OUT_PLAYER = "RpcInvokeTickOutPlayer";
    public const string GET_TEAM_DATA = "RpcInvokeGetTeamData";
}


/// <summary>
/// MELAND service call method API name
/// </summary>
public static class ManagerServiceAction
{
    /// <summary>
    /// 启动服务
    /// </summary>
    public const string START_SERVICE = "ManagerActionStartService";
    /// <summary>
    /// 获取服务时间戳
    /// </summary>
    public const string SERVICE_TIMESTAMP = "QueryServerTime";
    /// <summary>
    /// 注册服务
    /// </summary>
    public const string REGISTER_SERVICE = "ManagerActionServiceRegister";
    /// <summary>
    /// 查询服务信息
    /// </summary>
    public const string SELECT_SERVICE = "ManagerActionSelectService";

    public const string MULTI_SELECT_SERVICE = "ManagerActionMultiSelectService";

}

/// <summary>
///  main service call method API name
/// </summary>
public static class MainServiceAction
{
    /// <summary>
    /// use nft
    /// </summary>
    public const string USE_NFT = "MainServiceActionUseNFT";
    /// <summary>
    /// 扣除nft
    /// </summary>
    public const string TAKE_NFT = "MainServiceActionTakeNFT";
    /// <summary>
    /// 获取所有建造物数据 for main service
    /// </summary>
    public const string GET_ALL_BUILD = "MainServiceActionGetAllBuild";

    /// <summary>
    /// 查询玩家家园数据
    /// </summary>
    public const string GET_HOME_DATA = "MainServiceActionGetHomeData";
    /// <summary>
    /// 保存玩家家园数据
    /// </summary>
    public const string SAVE_HOME_DATA = "MainServiceActionSaveHomeData";
    /// <summary>
    /// 更新家园数据最后一次存储时间
    /// </summary>
    public const string UP_HOME_LAST_SAVE_TIME = "MainServiceActionUpHomeLastSaveTime";

    /// <summary>
    /// 批量存储animal基础信息
    /// </summary>
    public const string UPDATE_ANIMAL_BASE_DATA = "MainServiceActionMultiUpdateAnimalBaseData";
    /// <summary>
    /// 查询玩家animal列表信息(主要是大世界抓捕时使用)
    /// </summary>
    public const string GET_USER_ANIMAL_LIST = "MainServiceActionGetUserAnimalList";
    /// <summary>
    /// 抓捕到动物(主要是大世界抓捕时使用)
    /// </summary>
    public const string CAPTURE_ANIMAL = "MainServiceActionCaptureAnimal";

}

/// <summary>
/// proto message Action calls method name 
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
}

/// <summary>
/// user action method api
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
    ///  查询玩家详细数据 for main service
    /// </summary>
    public const string GET_USER_DATA = "GetUserData";
}

/// <summary>
/// user change service action
/// </summary>
public static class ChangeServiceAction
{
    /// <summary>
    /// user请求进入目标scene service（预进入阶段）
    /// </summary>
    public const string APPLY_ENTER_SERVICE = "ChangeServiceActionApplyEnterService";
    /// <summary>
    /// user正式进入目标scene service（正式进入阶段）
    /// </summary>
    public const string JOIN_SERVICE = "ChangeServiceActionJoinService";
}


