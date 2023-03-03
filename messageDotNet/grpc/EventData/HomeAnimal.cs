using System;
using UnityEngine;


[Serializable]
public class FreedAnimalEvent
{
    public long MsgVersion;
    public long UserId;
    public GrpcAnimalBaseData Animal;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
}

[Serializable]
public class CaptureAnimalEvent
{
    public long MsgVersion;
    public long UserId;
    public GrpcAnimalBaseData FreedAnimal;
    public GrpcAnimalBaseData CaptureAnimal;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
}