
/// <summary>
/// MELAND service pubsub event topic name 
/// </summary>
internal static class SubscriptionEvent
{
    /// <summary>
    /// 注册服务
    /// </summary>
    public const string REGISTER_SERVICE = "EventTopicRegisterService";
    /// <summary>
    /// 释放服务
    /// </summary>
    public const string DESTROY_SERVICE = "EventTopicDestroyService";

    /// <summary>
    /// 存储玩家数据
    /// </summary>
    public const string SAVE_PLAYER_DATA = "EventTopicSavePlayerData";
    /// <summary>
    /// 玩家拾取道具
    /// </summary>
    public const string PICK_DROP = "EventTopicPickDrop";
    /// <summary>
    /// 击杀怪物
    /// </summary>
    public const string KILL_MONSTER = "EventTopicKillMonster";
    /// <summary>
    /// 玩家死亡
    /// </summary>
    public const string PLAYER_DEATH = "EventTopicPlayerDeath";
}