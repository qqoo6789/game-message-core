using System;
using UnityEngine;


/// <summary>
/// dapr call 采集物品
/// </summary>
[Serializable]
public class CollectEvent
{
    // 消息版本号 值为毫秒时间戳
    public long MsgVersion;
    public ServiceData FormService;
    public long UserId;
    public EntityInfo TargetEntity;
    public string State; // collect_start || collect_end
    public GrpcItemBaseInfo DropItem;
}
