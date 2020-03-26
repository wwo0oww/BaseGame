using System.Collections;
using System.Collections.Generic;
using UnityEngine;
namespace NMObj
{
    public enum ObjType : byte
    {
        NPC = 1,
        BUILDING = 2,
        TERRAIN = 3
    }

    public class ObjMgr
    {
        public static Dictionary<int, Obj> Objs = new Dictionary<int, Obj>();
        private static Dictionary<int, List<Obj>> ObjsCache = new Dictionary<int, List<Obj>>();
        public static Obj NewObj(Vector3 pos, int objType, int id)
        {
            int NextobjType;
            int CurObjType = T_Obj.GetCurObjType(objType, out NextobjType);
            Obj newObj = GetOne(objType);
            if (newObj == null)
            {
                switch (CurObjType)
                {
                    case (int)ObjType.NPC:
                        newObj = NPCMrg.NewObj(pos, objType, NextobjType, id);
                        break;
                    case (int)ObjType.BUILDING:
                        newObj = NPCMrg.NewObj(pos, objType, NextobjType, id); // todo
                        break;
                    case (int)ObjType.TERRAIN:
                        newObj = new Obj(pos, objType, id); // todo
                        break;
                    default:
                        newObj = NPCMrg.NewObj(pos, objType, NextobjType, id);// todo
                        break;
                }
            }
            else
            {
                newObj.GPose = pos;
                newObj.ObjType = objType;
                newObj.SetID(id);
            }
            Objs[newObj.ID] = newObj;
            return newObj;
        }

        public static void RecycleObj(Obj obj)
        {
            obj.GPose = new Vector3(0, -100, 0);
            obj.DoStopMove();
            Objs.Remove(obj.ID);
            List<Obj> objs = new List<Obj>();
            if (ObjsCache.TryGetValue(obj.ObjType, out objs))
            {
                objs.Add(obj);
            }
            else
            {
                objs.Add(obj);
                ObjsCache.Add(obj.ObjType, objs);
            }

        }

        private static Obj GetOne(int objType)
        {
            List<Obj> objs = new List<Obj>();
            Obj obj = null;
            if (ObjsCache.TryGetValue(objType, out objs))
            {
                if (objs.Count > 0)
                {
                    obj = objs[0];
                    objs.RemoveAt(0);
                }
            }
            return obj;
        }
    }
}
