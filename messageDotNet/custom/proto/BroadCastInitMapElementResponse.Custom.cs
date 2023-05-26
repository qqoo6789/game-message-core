/*
 * @Author: xiang huan
 * @Date: 2023-04-25 13:46:11
 * @Description: 协议使用引用池，注意清理干净
 * @FilePath: /meland-scene-server/Assets/Plugins/SharedCore/game-message-core/messageDotNet/custom/proto/BroadCastInitMapElementResponse.Custom.cs
 * 
 */

using System.Collections.Generic;

namespace GameMessageCore
{
    public sealed partial class BroadCastInitMapElementResp : ICustomReference
    {
        public static BroadCastInitMapElementResp Create(List<EntityWithLocation> addEntitys, List<DestructionElementData> destruction, bool final)
        {
            BroadCastInitMapElementResp reference = CustomReferencePool.Acquire<BroadCastInitMapElementResp>();
            if (addEntitys != null)
            {
                reference.Entity.AddRange(addEntitys);
            }
            if (destruction != null)
            {
                reference.DestructionElements.AddRange(destruction);
            }
            reference.Final = final;
            return reference;
        }
        public void Dispose()
        {
            CustomReferencePool.Release(this);
        }
        public void Clear()
        {
            for (int i = 0; i < Entity.Count; i++)
            {
                Entity[i].IsLock = false;
            }
            Entity.Clear();
            DestructionElements.Clear();
            Final = false;
        }
    }
}
