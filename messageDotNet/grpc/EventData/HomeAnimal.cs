using System;
using UnityEngine;


[Serializable]
public class FreedAnimalEvent
{
    public long MsgVersion;
    public long UserId;
    public GrpcAnimalBaseData Animal; 
}

[Serializable]
public class CaptureAnimalEvent
{
    public long MsgVersion;
    public string ServiceAppId;
    public int MapId;
    public long UserId;
    public GrpcAnimalBaseData FreedAnimal;
    public GrpcAnimalBaseData CaptureAnimal; 
}