using UnityEngine;
using UnityEditor;
using Google.Protobuf;
using NMObj;
using System.Collections.Generic;

public class MapMsg
{
    enum ObjUpdateType
    {
        ADD = 1,
        UPT = 2,
        DEL = 3
    }
    public static void Handler(IMessage rawMsg)
    {
        if (rawMsg.GetType() == typeof(Net.m_obj_move_toc))
        {
            ObjMove(rawMsg);
            return;
        }
        if (rawMsg.GetType() == typeof(Net.m_map_player_toc))
        {
            InitPlayer(rawMsg);
            return;
        }
        if (rawMsg.GetType() == typeof(Net.m_obj_update_toc))
        {
            UpdateObj(rawMsg);
            return;

        }
        if (rawMsg.GetType() == typeof(Net.m_map_info_toc))
        {
            InitMapInfo(rawMsg);
        }
        Debug.Log("IMessage" + rawMsg);
        // Obj terrain = ObjMgr.NewObj(new Vector3(x * Config.TerrainMeshX + MapCenterPos.x, 0, z * Config.TerrainMeshX + MapCenterPos.z), T_Obj.T_TERRAIN.GenObjType(), 1);
        // terrain.Init(); ;
    }

    public static void InitMapInfo(IMessage rawMsg)
    {
        var msg = rawMsg as Net.m_map_info_toc;
        Map.Instance.ServerFrame = msg.FrameCount;
        foreach (var item in msg.ObjInfo)
        {
            TryAddObj(item);
        }
    }

    public static void ObjMove(IMessage rawMsg)
    {
        var msg = rawMsg as Net.m_obj_move_toc;
        Debug.Log("m_obj_move_toc" + msg);
        int objID = msg.ObjId;
        var obj = ObjMgr.Objs[objID];
        obj.GPose = Common.GetLocalPos(msg.Pos);
        if (obj.HasStatus(ObjStatus.MOVE) && (msg.Direction == (int)ObjDirection.STOP || msg.Direction == (int)ObjDirection.NONE))
        {
            obj.DoStopMove();
        }
        else
        {
            if (!(msg.Direction == (int)ObjDirection.STOP || msg.Direction == (int)ObjDirection.NONE))
            {
                obj.DoMove((ObjDirection)msg.Direction);
            }
        }
    }

    public static void InitPlayer(IMessage rawMsg)
    {
        var msg = rawMsg as Net.m_map_player_toc;
        Debug.Log("m_map_player_toc" + msg);
        Map.PObj = msg.ObjInfo;
    }

    public static void UpdateObj(IMessage rawMsg)
    {
        var msg = rawMsg as Net.m_obj_update_toc;
        Debug.Log("m_obj_update_toc" + msg);
        switch (msg.Type)
        {
            case (int)ObjUpdateType.ADD:
            case (int)ObjUpdateType.UPT:
                foreach (var item in msg.ObjInfo) { 
                    TryAddObj(item); 
                }
                break;
            case (int)ObjUpdateType.DEL:
                break;
            default:
                break;
        }

    }
    public static Obj TryAddObj(Net.p_obj ObjInfo)
    {
        var pos = Common.GetLocalPos(ObjInfo.Pos);
        Obj Obj;
        if (ObjMgr.Objs.ContainsKey(ObjInfo.Id))
        {
            Obj = ObjMgr.Objs[ObjInfo.Id];

        }
        else
        {
            Obj = ObjMgr.NewObj(pos, ObjInfo.Type, ObjInfo.Id);
            Obj.Init();
        }

        if (ObjInfo.Pos.X != 0 && ObjInfo.Pos.Y != 0)
        {
            Obj.GPose = pos;
        }
        if (!ObjInfo.Name.Equals(""))
        {
            Obj.SetName(ObjInfo.Name);
        }

        if (ObjInfo.Status == (int)(ObjStatus.NONE) && Obj.HasStatus(ObjStatus.MOVE))
        {
            Obj.DoStopMove();
        }
        else if (ObjInfo.Status == (int)ObjStatus.MOVE)
        {
            Obj.SetSpeed(ObjInfo.Speed);
            Obj.DoMove((ObjDirection)ObjInfo.Direction);
        }
        return Obj;
    }
}