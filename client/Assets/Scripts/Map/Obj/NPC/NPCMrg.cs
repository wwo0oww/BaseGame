using System.Collections;
using System.Collections.Generic;
using UnityEngine;
namespace NMObj
{
    public enum NPCType : byte
    {
        PlAYER = 1,
        TESTNPC = 2,
        
    }

    public class NPCMrg
    {
        /// <summary>
        /// 
        /// </summary>
        /// <param name="pos"></param>
        /// <param name="baseObjType">基础类型如1001001【Obj=>NPC=>TestNPC】</param>
        /// <param name="curObjType">去除了上一层类型的当前类型，例testObj类型值为1001001【Obj=>NPC=>TestNPC】，则CurObjType为1001【NPC=>TestNPC】</param>
        /// <param name="id"></param>
        /// <param name="pos"></param>
        /// <returns></returns>
        public static NPC NewObj(Vector3 pos, int baseObjType, int curObjType,int id)
        {
            int NextobjType;
            int CurobjType = T_Obj.GetCurObjType(baseObjType, out NextobjType);
            switch (CurobjType)
            {
                case (int)NPCType.TESTNPC:
                    return new TestNPC(pos, baseObjType, id);
                case (int)NPCType.PlAYER:
                    return new Player(pos, baseObjType, id);
                default:
                    return new NPC(pos, baseObjType, id);// todo
            }
        }
    }
}
