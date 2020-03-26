using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using System;
using Google.Protobuf;     //引用DLL
using Net;

public class gen_proto
{

    static Int16 start_index = 1000;
    public static Hashtable protoIDMap = new Hashtable();
    public static Hashtable protoTypeMap = new Hashtable();

    public static void InitProtoId()
    {
        add(new TestEchoACK());
        add(new m_login_tos());
        add(new m_login_toc()); rounter.Register(new m_login_toc(), Login.Handler);
        add(new m_heartbeat_tos());
        add(new m_heartbeat_toc()); rounter.Register(new m_login_toc(), Login.Handler);
        add(new m_obj_update_toc()); rounter.Register(new m_obj_update_toc(), MapMsg.Handler);
        add(new m_map_info_toc()); rounter.Register(new m_map_info_toc(), MapMsg.Handler);
        add(new m_map_player_toc()); rounter.Register(new m_map_player_toc(), MapMsg.Handler);
        add(new m_obj_move_toc()); rounter.Register(new m_obj_move_toc(), MapMsg.Handler);
        add(new m_obj_move_tos()); rounter.Register(new m_obj_move_tos(), MapMsg.Handler);
    }

    public static void add(IMessage t)
    {

        protoIDMap[t.GetType()] = ++start_index;
        protoTypeMap[start_index] = t;
    }
}
