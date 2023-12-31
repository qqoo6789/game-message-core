﻿/*
 * @Author: xiang huan
 * @Date: 2023-04-23 20:38:52
 * @Description: 自定义引用集合
 * @FilePath: /meland-unity/Assets/Plugins/SharedCore/game-message-core/messageDotNet/custom/CustomReferenceCollection.cs
 * 
 */


using System;
using System.Collections.Generic;
public class CustomReferenceCollection
{
    private readonly Queue<object> _references;
    public Type ReferenceType { get; private set; }
    public int UsingReferenceCount { get; private set; }
    public int AcquireReferenceCount { get; private set; }
    public int ReleaseReferenceCount { get; private set; }
    public int AddReferenceCount { get; private set; }
    public int RemoveReferenceCount { get; private set; }

    public int UnusedReferenceCount => _references.Count;

    public CustomReferenceCollection(Type referenceType)
    {
        _references = new Queue<object>();
        ReferenceType = referenceType;
        UsingReferenceCount = 0;
        AcquireReferenceCount = 0;
        ReleaseReferenceCount = 0;
        AddReferenceCount = 0;
        RemoveReferenceCount = 0;
    }

    public T Acquire<T>() where T : class, new()
    {
        if (typeof(T) != ReferenceType)
        {
            throw new Exception("Custom Type is invalid.");
        }
        return Acquire() as T;
    }
    public object Acquire()
    {
        UsingReferenceCount++;
        AcquireReferenceCount++;
        lock (_references)
        {
            if (_references.Count > 0)
            {
                return _references.Dequeue();
            }
        }
        AddReferenceCount++;
        return Activator.CreateInstance(ReferenceType);
    }

    public void Release(object reference)
    {
        lock (reference)
        {
            if (reference is ICustomReference customReference)
            {
                customReference.Clear();
            }
        }

        lock (_references)
        {
            _references.Enqueue(reference);
        }
        ReleaseReferenceCount++;
        UsingReferenceCount--;
    }

    public void Add<T>(int count) where T : class, new()
    {
        if (typeof(T) != ReferenceType)
        {
            throw new Exception("Custom Type is invalid.");
        }
        Add(count);
    }

    public void Add(int count)
    {
        lock (_references)
        {
            AddReferenceCount += count;
            while (count-- > 0)
            {
                _references.Enqueue(Activator.CreateInstance(ReferenceType));
            }
        }
    }

    public void Remove(int count)
    {
        lock (_references)
        {
            if (count > _references.Count)
            {
                count = _references.Count;
            }

            RemoveReferenceCount += count;
            while (count-- > 0)
            {
                _ = _references.Dequeue();
            }
        }
    }

    public void RemoveAll()
    {
        lock (_references)
        {
            RemoveReferenceCount += _references.Count;
            _references.Clear();
        }
    }
}
