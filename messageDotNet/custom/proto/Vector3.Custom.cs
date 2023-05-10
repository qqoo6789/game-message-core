/*
 * @Author: xiang huan
 * @Date: 2023-04-25 13:46:11
 * @Description: 协议使用引用池，注意清理干净
 * @FilePath: /meland-scene-server/Assets/Plugins/SharedCore/game-message-core/messageDotNet/custom/proto/Vector3.Custom.cs
 * 
 */

namespace GameMessageCore
{
    public sealed partial class Vector3
    {
        public void Set(float x, float y, float z)
        {
            X = x;
            Y = y;
            Z = z;
        }
    }
}
