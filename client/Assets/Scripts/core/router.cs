using Google.Protobuf;
using System.Collections;
using System.Collections.Generic;
using WoW.Common.Threading;
using UnityEngine;

public class rounter
{
    public delegate void Handler(IMessage msg);
    static Hashtable HandlerMap = new Hashtable();
    public static void Register(IMessage t, Dispatcher.Action handler)
    {
       
        HandlerMap[t.GetType()] = handler;
    }
    public static Dispatcher.Action GetHandler(IMessage t)
    {
        return HandlerMap[t.GetType()] as Dispatcher.Action;
    }
}
