using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using Google.Protobuf;     //引用DLL
using System.Threading;
using Net;

public class heartbeat  {

    public static void Handle(IMessage msg) {
        m_heartbeat_toc toc = msg as m_heartbeat_toc;
        var i = toc.Timestamp;
    }
    public static void startheart()
    {
        Thread th = new Thread(new ThreadStart(ThrHeartbeat)); //创建线程
        th.Start(); //启动线程

    }
    public static void ThrHeartbeat()
    {
        while (true)
        {
            m_heartbeat_tos tos = new m_heartbeat_tos();
            WSocket.Send(tos);
            Thread.Sleep(1000);
        }
    }
}
