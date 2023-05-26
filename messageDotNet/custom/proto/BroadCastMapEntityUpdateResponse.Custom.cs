/*
 * @Author: xiang huan
 * @Date: 2023-04-25 13:46:11
 * @Description: 协议使用引用池
 * @FilePath: /meland-scene-server/Assets/Plugins/SharedCore/game-message-core/messageDotNet/custom/proto/BroadCastMapEntityUpdateResponse.Custom.cs
 * 
 */

using System.Collections.Generic;

namespace GameMessageCore
{
    public sealed partial class BroadCastMapEntityUpdateResp : ICustomReference
    {

        public static BroadCastMapEntityUpdateResp Create(List<EntityId> removeEntitys, List<EntityWithLocation> addEntitys, List<DestructionElementData> destructionData)
        {
            BroadCastMapEntityUpdateResp reference = CustomReferencePool.Acquire<BroadCastMapEntityUpdateResp>();
            if (removeEntitys != null)
            {
                reference.EntityRemoved.AddRange(removeEntitys);
            }
            if (addEntitys != null)
            {
                reference.EntityAdded.AddRange(addEntitys);
            }
            if (destructionData != null)
            {
                reference.DestructionElements.AddRange(destructionData);
            }
            return reference;
        }
        public void Dispose()
        {
            CustomReferencePool.Release(this);
        }
        public void Clear()
        {
            for (int i = 0; i < EntityAdded.Count; i++)
            {
                EntityAdded[i].IsLock = false;
            }
            EntityAdded.Clear();
            EntityRemoved.Clear();
            DestructionElements.Clear();
        }
    }
}
