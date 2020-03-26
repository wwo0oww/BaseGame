using System.Collections;
using System.Collections.Generic;
using UnityEngine;
namespace NMObj
{

    /// <summary>
    ///  用于方便获取obj类型值，和解析obj类型值。每个类需要有  方法GenObjType
    /// </summary>
    public static class T_Obj
    {
        public const int LevelNum = 1000;
        public static int GenObjType()
        {
            return 1;
        }
        public static int GetCurObjType(int objType, out int RemainType)
        {
            RemainType = objType / LevelNum;
            return objType % LevelNum;
        }
        /// <summary>
        /// obj=>npc
        /// </summary>
        public static class T_NPC
        {
            public static int GenObjType()
            {
                return T_Obj.GenObjType() * LevelNum + (int)(ObjType.NPC);
            }

            /// <summary>
            /// obj=>npc=>testobj
            /// </summary>
            public static class T_TestNPC
            {
                public static int GenObjType()
                {
                    return T_NPC.GenObjType() * LevelNum + (int)(NPCType.TESTNPC);
                }
            }

            /// <summary>
            /// obj=>npc=>player
            /// </summary>
            public static class T_PLAYER
            {
                public static int GenObjType()
                {
                    return T_NPC.GenObjType() * LevelNum + (int)(NPCType.PlAYER);
                }
            }
        }

        /// <summary>
        /// obj=>terrain
        /// </summary>
        public static class T_TERRAIN
        {
            /// <summary>
            /// 获取 地形 的类型值
            /// </summary>
            /// <returns></returns>
            public static int GenObjType()
            {
                return T_Obj.GenObjType() * LevelNum + (int)(ObjType.TERRAIN); ;
            }
        }
    }

}