
public enum EventExpire : uint
{
    EVENT_EXPIRE_DEFAULT = 0,
    EVENT_EXPIRE_SEC_5 = 1 * 5,
    EVENT_EXPIRE_SEC_30 = 1 * 30,
    EVENT_EXPIRE_MIN_1 = EVENT_EXPIRE_SEC_30 * 2,
    EVENT_EXPIRE_HOUR_1 = EVENT_EXPIRE_MIN_1 * 60,
    EVENT_EXPIRE_HOUR_12 = EVENT_EXPIRE_HOUR_1 * 12,
    EVENT_EXPIRE_DAY = EVENT_EXPIRE_HOUR_1 * 24,
    EVENT_EXPIRE_WEEK = EVENT_EXPIRE_DAY * 7,
}


/// <summary>
/// MELAND service pubsub event topic name 
/// </summary>
public static class SubscriptionEvent
{
    /// <summary>
    /// 服务销毁事件
    /// </summary>
    public const string SERVICE_UNREGISTER = "EventTopicServiceUnregister";
    /// <summary>
    /// 玩家完成服务切换事件
    /// </summary>
    public const string USER_CHANGE_SERVICE = "EventTopicUserChangeService";

    /// <summary>
    /// 添加nft
    /// </summary>
    public const string MINT_NFT = "EventTopicMintNFT";
    /// <summary>
    /// 批量添加nft
    /// </summary>
    public const string MULTI_MINT_NFT = "EventTopicMultiMintNFT";
    /// <summary>
    /// 存储玩家数据
    /// </summary>
    public const string SAVE_PLAYER_DATA = "EventTopicSavePlayerData";
    /// <summary>
    /// 击杀怪物
    /// </summary>
    public const string KILL_MONSTER = "EventTopicKillMonster";
    /// <summary>
    /// 玩家死亡
    /// </summary>
    public const string PLAYER_DEATH = "EventTopicPlayerDeath";
    /// <summary>
    /// use NFT event
    /// </summary>
    public const string USE_NFT = "EventTopicUseNFT";
    /// <summary>
    /// get NFT event
    /// </summary>
    public const string GET_NFT = "EventTopicGetNFT";
    /// <summary>
    /// user enter game event
    /// </summary>
    public const string USER_ENTER_GAME = "EventTopicUserEnterGame";
    /// <summary>
    /// 玩家离开游戏(断开网络链接时)
    /// </summary>
    public const string USER_LEAVE_GAME = "EventTopicUserLeaveGame";
    /// <summary>
    /// add nft build event
    /// </summary>
    public const string NFT_BUILD_ADD = "EventTopicNftBuildAdd";

    /// <summary>
    /// update nft build event
    /// </summary>
    public const string NFT_BUILD_UPDATE = "EventTopicNftBuildUpdate";
    /// <summary>
    /// remove nft build event
    /// </summary>
    public const string NFT_BUILD_REMOVE = "EventTopicNftBuildRemove";
    /// <summary>
    /// task finish event
    /// </summary>
    public const string TASK_FINISH = "EventTopicTaskFinish";
    /// <summary>
    /// task list finish event
    /// </summary>
    public const string TASK_LIST_FINISH = "EventTopicTaskListFinish";

    /// <summary>
    /// home Granary Stockpile Event
    /// </summary>
    public const string GRANARY_STOCKPILE = "EventTopicGranaryStockpile";

    /// <summary>
    /// update user talent data event
    /// </summary>
    public const string UPDATE_TALENT = "EventTopicUpdateTalent";

    /// <summary>
    /// add  user talent tree exp event
    /// </summary>
    public const string ADD_TALENT_EXP = "EventTopicAddTalentExp";


    /// <summary>
    /// freed animal event
    /// </summary>
    public const string FREED_ANIMAL = "EventTopicFreedAnimal";

    /// <summary>
    /// Capture Animal event
    /// </summary>
    public const string CAPTURE_ANIMAL = "EventTopicCaptureAnimal";


    /// <summary>
    /// 家园播种事件
    /// </summary>
    public const string HOME_SOWING = "EventTopicHomeSowing";

    /// <summary>
    /// 家园放置动物食物事件
    /// </summary>
    public const string HOME_PUT_ANIMAL_FOOD = "EventTopicHomePutAnimalFood";

    /// <summary>
    /// 家园收获土地产出事件
    /// </summary>
    public const string HOME_HARVEST = "EventTopicHomeHarvest";


    /// <summary>
    /// 开始采集物品 搭配状态(collect_start || collect_end)使用
    /// </summary>
    public const string COLLECT = "EventTopicCollect";

    /// <summary>
    ///  日志采集
    /// </summary>
    public const string GLOBAL_LOG = "EventTopicGlobalLog";

    // friend events
    public const string CREATE_TEAM = "EventTopicCreateTeam";
    public const string DISBAND_TEAM = "EventTopicDisbandTeam";
    public const string JOIN_TEAM = "EventTopicJoinTeam";
    public const string QUIT_TEAM = "EventTopicQuitTeam";

}

