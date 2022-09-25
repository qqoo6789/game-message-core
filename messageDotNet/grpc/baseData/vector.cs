using System;
using UnityEngine;

// 对应 proto.Vector3
[Serializable]
public class GrpcVector3
{
    public float X;
    public float Y;
    public float Z;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }

    public void Set(GameMessageCore.Vector3 v)
    {
        if (v == null)
        {
            return;
        }
        X = v.X;
        Y = v.Y;
        Z = v.Z;
    }

    public GameMessageCore.Vector3 ToProtoData()
    {
        return new GameMessageCore.Vector3()
        {
            X = X,
            Y = Y,
            Z = Z,
        };
    }

}