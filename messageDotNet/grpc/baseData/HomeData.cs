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
    public long LastSaveMs;
    public string SoilJson;
    public string ResourceJson;
    public string AnimalJson;
    public string AnimalSceneJson;
    public GrpcAnimalBaseData[] AnimalList;
    public bool UseDefaultData;

    private const int USE_HASH_COMPARE_JSON_LENGTH = 10000;//长度大于这个值时 用hash比较json字符串

    public bool EqualsSaveHomeData(GrpcHomeData targetData)
    {
        if (targetData == null)
        {
            return false;
        }

        //列举已知需要判断是否相等的字段 后续忘记添加当做不相等保护起来不至于不保存数据
        if (EqualsJson(SoilJson, targetData.SoilJson) &&
            EqualsJson(ResourceJson, targetData.ResourceJson) &&
            EqualsJson(AnimalJson, targetData.AnimalJson) &&
            EqualsJson(AnimalSceneJson, targetData.AnimalSceneJson))
        {
            return true;
        }

        return false;
    }

    private bool EqualsJson(string json1, string json2)
    {
        if (string.IsNullOrEmpty(json1) || string.IsNullOrEmpty(json2))
        {
            return false;
        }

        // 通过实测TestStringComparePerformance 只要字符串到达GB级别 用hash比较性能才能更高 肯定达不到GB级别
        // 另外在M1 mac上测试
        // 前一半相同 后一半不同字符串时 5M大小 string.Compare 开销在6ms左右  string.Equals在11ms hash在10ms
        // 完全相同字符串 string.Compare和hash相当 Equals两倍开销
        // 完全不同 都远远大于hash性能
        // 所以直接选定string.Compare
        // if (Mathf.Min(json1.Length, json2.Length) > USE_HASH_COMPARE_JSON_LENGTH)//长度都很大时 不能简单字符串比较 需要转hash比较 性能更高
        // {
        //     return json1.GetHashCode() == json2.GetHashCode();
        // }
        // else
        // {
        return string.Compare(json1, json2) == 0;
        // }
    }

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }

}
