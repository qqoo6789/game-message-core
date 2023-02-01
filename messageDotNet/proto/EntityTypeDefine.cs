// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: entityTypeDefine.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace GameMessageCore {

  /// <summary>Holder for reflection information generated from entityTypeDefine.proto</summary>
  public static partial class EntityTypeDefineReflection {

    #region Descriptor
    /// <summary>File descriptor for entityTypeDefine.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static EntityTypeDefineReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "ChZlbnRpdHlUeXBlRGVmaW5lLnByb3RvEg9nYW1lTWVzc2FnZUNvcmUqswEK",
            "CkVudGl0eVR5cGUSEQoNRW50aXR5VHlwZUFsbBAAEhQKEEVudGl0eVR5cGVQ",
            "bGF5ZXIQARIVChFFbnRpdHlUeXBlTW9uc3RlchACEhYKEkVudGl0eVR5cGVO",
            "ZnRCdWlsZBADEhEKDUVudGl0eVR5cGVOcGMQBBIWChJFbnRpdHlUeXBlUmVz",
            "b3VyY2UQBRIRCg1FbnRpdHlUeXBlTWF4EAYSDwoKTWFpblBsYXllchDnB2IG",
            "cHJvdG8z"));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { },
          new pbr::GeneratedClrTypeInfo(new[] {typeof(global::GameMessageCore.EntityType), }, null, null));
    }
    #endregion

  }
  #region Enums
  public enum EntityType {
    [pbr::OriginalName("EntityTypeAll")] All = 0,
    /// <summary>
    /// 玩家
    /// </summary>
    [pbr::OriginalName("EntityTypePlayer")] Player = 1,
    /// <summary>
    /// 怪物
    /// </summary>
    [pbr::OriginalName("EntityTypeMonster")] Monster = 2,
    /// <summary>
    /// 建造物
    /// </summary>
    [pbr::OriginalName("EntityTypeNftBuild")] NftBuild = 3,
    /// <summary>
    /// npc
    /// </summary>
    [pbr::OriginalName("EntityTypeNpc")] Npc = 4,
    /// <summary>
    /// 资源
    /// </summary>
    [pbr::OriginalName("EntityTypeResource")] Resource = 5,
    /// <summary>
    /// 公共类型最大值  
    /// </summary>
    [pbr::OriginalName("EntityTypeMax")] Max = 6,
    /// <summary>
    /// 客户端专用，主角
    /// </summary>
    [pbr::OriginalName("MainPlayer")] MainPlayer = 999,
  }

  #endregion

}

#endregion Designer generated code
