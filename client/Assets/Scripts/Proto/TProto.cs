using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using Google.Protobuf;
using Google.Protobuf.Reflection;

/// <summary>
/// 空结构，只是为了传值
/// </summary>
public class TProto : IMessage {


   
    public MessageDescriptor Descriptor;

    MessageDescriptor IMessage.Descriptor
    {
        get
        {
            throw new System.NotImplementedException();
        }
    }


    public int CalculateSize()
    {
        throw new System.NotImplementedException();
    }

    public void MergeFrom(CodedInputStream input)
    {
        throw new System.NotImplementedException();
    }

    public void WriteTo(CodedOutputStream output)
    {
        throw new System.NotImplementedException();
    }

    
}
