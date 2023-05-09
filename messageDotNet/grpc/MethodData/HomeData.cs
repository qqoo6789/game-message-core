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
    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
}
[Serializable]
public class MainServiceActionSaveHomeDataOutput
{
    public bool Success;
    public string ErrMsg;
    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
}

/// <summary>
/// call main service update  user home data last save time
/// </summary>
[Serializable]
public class MainServiceActionUpHomeLastSaveTimeInput
{
    public long UserId;
    public long LastSaveMs;
    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
}
[Serializable]
public class MainServiceActionUpHomeLastSaveTimeOutput
{
    public bool Success;
    public string ErrMsg;
    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
}

/// <summary>
/// select user home data for main service
/// </summary>
[Serializable]
public class MainServiceActionGetHomeDataInput
{
    public long UserId;
    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
}
[Serializable]
public class MainServiceActionGetHomeDataOutput
{
    public bool Success;
    public string ErrMsg;
    public long UserId;
    public GrpcHomeData Data;
    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }

}

/// <summary>
/// call main service save animal list data
/// </summary>
[Serializable]
public class MultiUpdateAnimalBaseDataInput
{
    public long UserId;
    public GrpcAnimalBaseData[] Animals;
    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
}
[Serializable]
public class MultiUpdateAnimalBaseDataOutput
{
    public bool Success;
    public string ErrMsg;
    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }

}

/// <summary>
///  GetUserAnimalList  world to main service
/// </summary>
[Serializable]
public class GetUserAnimalListInput
{
    public long UserId;
    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
}
[Serializable]
public class GetUserAnimalListOutput
{
    public bool Success;
    public string ErrMsg;
    public long UserId;
    public GrpcAnimalBaseData[] Animals;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }

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
    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
}
[Serializable]
public class CaptureAnimalOutput
{
    public bool Success;
    public string ErrMsg;
    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
}



