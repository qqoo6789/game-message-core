using System;
using UnityEngine;


// home animal base data
[Serializable]
public class GrpcAnimalBaseData
{
    public long AnimalId;
    public string Name;
    public int Cid;
    public int Favorability;
    public long CreateMs;
    public long UpdateMs;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }

}

// home data
[Serializable]
public class GrpcHomeData
{
    public string SoilJson;
    public string ResourceJson;
    public string AnimalJson;
    public string AnimalSceneJson;
    public GrpcAnimalBaseData[] AnimalList;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }

}
