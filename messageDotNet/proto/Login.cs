// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: login.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace GameMessageCore {

  /// <summary>Holder for reflection information generated from login.proto</summary>
  public static partial class LoginReflection {

    #region Descriptor
    /// <summary>File descriptor for login.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static LoginReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "Cgtsb2dpbi5wcm90bxIPZ2FtZU1lc3NhZ2VDb3JlKi8KCkxvZ2luU3RhdGUS",
            "CQoFTGVhdmUQABIJCgVRdWV1ZRABEgsKB1BsYXlpbmcQAmIGcHJvdG8z"));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { },
          new pbr::GeneratedClrTypeInfo(new[] {typeof(global::GameMessageCore.LoginState), }, null, null));
    }
    #endregion

  }
  #region Enums
  public enum LoginState {
    [pbr::OriginalName("Leave")] Leave = 0,
    [pbr::OriginalName("Queue")] Queue = 1,
    [pbr::OriginalName("Playing")] Playing = 2,
  }

  #endregion

}

#endregion Designer generated code
