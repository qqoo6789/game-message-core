/// <summary>
/// dapr event topic 与其他服务监听的topic 一一对应 
/// </summary>
public class EventTopicDefine
{
    public const string EVENT_TOPIC_SAVE_PLAYER_DATA = "EventTopicSavePlayerData";
    public const string EVENT_TOPIC_PICK_DROP = "EventTopicPickDrop";
    public const string EVENT_TOPIC_KILL_MONSTER = "EventTopicKillMonster";
    public const string EVENT_TOPIC_PLAYER_DEATH = "EventTopicPlayerDeath";
}