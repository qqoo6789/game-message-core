using System;
using UnityEngine;

/// <summary>
/// call main service save user home data
/// </summary>
[Serializable]
public class MainServiceActionSaveHomeDataInput
{
    public long UserId;
    public GrpcHomeData Data;
}
[Serializable]
public class MainServiceActionSaveHomeDataOutput
{
    public bool Success;
    public string ErrMsg;
  
}

/// <summary>
/// call main service update  user home data last save time
/// </summary>
[Serializable]
public class MainServiceActionUpHomeLastSaveTimeInput
{
    public long UserId;
    public long LastSaveMs; 
}
[Serializable]
public class MainServiceActionUpHomeLastSaveTimeOutput
{
    public bool Success;
    public string ErrMsg; 
}

/// <summary>
/// select user home data for main service
/// </summary>
[Serializable]
public class MainServiceActionGetHomeDataInput
{
    public long UserId; 
}
[Serializable]
public class MainServiceActionGetHomeDataOutput
{
    public bool Success;
    public string ErrMsg;
    public long UserId;
    public GrpcHomeData Data; 

}

/// <summary>
/// call main service save animal list data
/// </summary>
[Serializable]
public class MultiUpdateAnimalBaseDataInput
{
    public long UserId;
    public GrpcAnimalBaseData[] Animals; 
}
[Serializable]
public class MultiUpdateAnimalBaseDataOutput
{
    public bool Success;
    public string ErrMsg; 

}

/// <summary>
///  GetUserAnimalList  world to main service
/// </summary>
[Serializable]
public class GetUserAnimalListInput
{
    public long UserId; 
}
[Serializable]
public class GetUserAnimalListOutput
{
    public bool Success;
    public string ErrMsg;
    public long UserId;
    public GrpcAnimalBaseData[] Animals;
 
}

/// <summary>
///  GetUserAnimalList  world to main service
/// </summary>
[Serializable]
public class CaptureAnimalInput
{
    public long UserId;
    public long FreedAnimalId;
    public GrpcAnimalBaseData Animal;
     
}
[Serializable]
public class CaptureAnimalOutput
{
    public bool Success;
    public string ErrMsg;
}



