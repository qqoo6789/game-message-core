using System;
using UnityEngine;


[Serializable]
public class GrpcTalentExp
{
    public GameMessageCore.TalentType TalentType;
    public long Exp;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
}


[Serializable]
public class GrpcTalentLevel
{
    public GameMessageCore.TalentType TalentType;
    public int Level;
    public int MasterLevel;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
    public GameMessageCore.TalentLevel ToProtoData()
    {
        return new GameMessageCore.TalentLevel()
        {
            Type = TalentType,
            Level = Level,
            MasterLevel = MasterLevel,
        };
    }
}


[Serializable]
public class GrpcTalentNodeData
{
    public int NodeId;
    public int Level;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }
    public GameMessageCore.TalentNode ToProtoData()
    {
        return new GameMessageCore.TalentNode()
        {
            NodeId = NodeId,
            Level = Level,
        };
    }
}



[Serializable]
public class GrpcTalentTree
{
    public GameMessageCore.TalentType TalentType;
    public GrpcTalentNodeData[] Nodes;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }

    public GameMessageCore.TalentTree ToProtoData()
    {
        GameMessageCore.TalentTree tree = new() { Type = TalentType };
        if (Nodes != null)
        {
            for (int i = 0; i < Nodes.Length; i++)
            {
                tree.UnlockNodes.Add(Nodes[i].ToProtoData());
            }
        }
        return tree;
    }
}


[Serializable]
public class GrpcTalentData
{
    public GrpcTalentLevel[] Levels;
    public GrpcTalentTree[] Trees;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }

    public GameMessageCore.TalentData ToProtoData()
    {
        GameMessageCore.TalentData data = new();
        if (Levels != null)
        {
            int lvLength = Levels.Length;
            for (int i = 0; i < lvLength; i++)
            {
                data.LevelData.Add(Levels[i].ToProtoData());
            }
        }
        if (Trees != null)
        {
            int treeLength = Trees.Length;
            for (int i = 0; i < treeLength; i++)
            {
                data.Tree.Add(Trees[i].ToProtoData());
            }
        }
        return data;
    }
}


[Serializable]
public class GrpcTalentTreeUpdate
{
    public GameMessageCore.TalentType TalentType;
    public GrpcTalentNodeData[] AddNodes;
    public GrpcTalentNodeData[] UpdateNodes;
    public int[] RemoveNodes;

    public string ToJson()
    {
        return JsonUtility.ToJson(this);
    }

    public GameMessageCore.TalentTreeUpdate ToProtoData()
    {
        GameMessageCore.TalentTreeUpdate data = new() { Type = TalentType };
        if (AddNodes != null)
        {
            int length = AddNodes.Length;
            for (int i = 0; i < length; i++)
            {
                data.AddNodes.Add(AddNodes[i].ToProtoData());
            }
        }
        if (UpdateNodes != null)
        {
            int length = UpdateNodes.Length;
            for (int i = 0; i < length; i++)
            {
                data.UpdateNodes.Add(UpdateNodes[i].ToProtoData());
            }
        }
        if (RemoveNodes != null)
        {
            int length = RemoveNodes.Length;
            for (int i = 0; i < length; i++)
            {
                data.RemoveNodes.Add(RemoveNodes[i]);
            }
        }
        return data;
    }
}

