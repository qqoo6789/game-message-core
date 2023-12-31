// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: service.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace GameMessageCore {

  /// <summary>Holder for reflection information generated from service.proto</summary>
  public static partial class ServiceReflection {

    #region Descriptor
    /// <summary>File descriptor for service.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static ServiceReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "Cg1zZXJ2aWNlLnByb3RvEg9nYW1lTWVzc2FnZUNvcmUqhQIKC1NlcnZpY2VU",
            "eXBlEhYKElNlcnZpY2VUeXBlVW5rbm93bhAAEhMKD1NlcnZpY2VUeXBlTWFp",
            "bhABEhYKElNlcnZpY2VUeXBlQWNjb3VudBACEhQKEFNlcnZpY2VUeXBlU2Nl",
            "bmUQAxITCg9TZXJ2aWNlVHlwZVRhc2sQBBITCg9TZXJ2aWNlVHlwZUNoYXQQ",
            "BRIUChBTZXJ2aWNlVHlwZUFnZW50EAYSFgoSU2VydmljZVR5cGVNYW5hZ2Vy",
            "EAcSGQoVU2VydmljZVR5cGVDb250cm9sbGVyEAgSEgoOU2VydmljZVR5cGVM",
            "b2cQCRIUChBTZXJ2aWNlVHlwZUxpbWl0EAoqSwoTU2NlbmVTZXJ2aWNlU3Vi",
            "VHlwZRISCg5Vbmtub3duU3ViVHlwZRAAEgkKBVdvcmxkEAESCAoESG9tZRAC",
            "EgsKB0R1bmdlb24QA2IGcHJvdG8z"));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { },
          new pbr::GeneratedClrTypeInfo(new[] {typeof(global::GameMessageCore.ServiceType), typeof(global::GameMessageCore.SceneServiceSubType), }, null, null));
    }
    #endregion

  }
  #region Enums
  public enum ServiceType {
    [pbr::OriginalName("ServiceTypeUnknown")] Unknown = 0,
    /// <summary>
    /// 主服务
    /// </summary>
    [pbr::OriginalName("ServiceTypeMain")] Main = 1,
    /// <summary>
    /// 账号服务
    /// </summary>
    [pbr::OriginalName("ServiceTypeAccount")] Account = 2,
    /// <summary>
    /// 场景(战斗)服务
    /// </summary>
    [pbr::OriginalName("ServiceTypeScene")] Scene = 3,
    /// <summary>
    /// 任务服务
    /// </summary>
    [pbr::OriginalName("ServiceTypeTask")] Task = 4,
    /// <summary>
    /// 聊天服务
    /// </summary>
    [pbr::OriginalName("ServiceTypeChat")] Chat = 5,
    /// <summary>
    /// 网关服务
    /// </summary>
    [pbr::OriginalName("ServiceTypeAgent")] Agent = 6,
    /// <summary>
    /// 服务管理
    /// </summary>
    [pbr::OriginalName("ServiceTypeManager")] Manager = 7,
    /// <summary>
    /// 后台控制服务
    /// </summary>
    [pbr::OriginalName("ServiceTypeController")] Controller = 8,
    /// <summary>
    /// 日志服务
    /// </summary>
    [pbr::OriginalName("ServiceTypeLog")] Log = 9,
    /// <summary>
    /// 服务类型上限标记必须放在最后
    /// </summary>
    [pbr::OriginalName("ServiceTypeLimit")] Limit = 10,
  }

  public enum SceneServiceSubType {
    [pbr::OriginalName("UnknownSubType")] UnknownSubType = 0,
    [pbr::OriginalName("World")] World = 1,
    [pbr::OriginalName("Home")] Home = 2,
    [pbr::OriginalName("Dungeon")] Dungeon = 3,
  }

  #endregion

}

#endregion Designer generated code
