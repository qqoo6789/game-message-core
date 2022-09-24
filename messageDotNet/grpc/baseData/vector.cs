using System;
using UnityEngine;

// 对应 proto.Vector3
[Serializable]
public class GrpcVector3
{
    public float32 X;
    public float32 Y;
    public float32 Z;

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
        Y = v.Z;
        Z = v.Z;
    }
	
    public GameMessageCore.Vector3 ToProtoData()
    {
        return new GameMessageCore.Vector3()
        {
            X = p.X,
            Y = p.Y,
            Z = p.Z,
        };
    }

}