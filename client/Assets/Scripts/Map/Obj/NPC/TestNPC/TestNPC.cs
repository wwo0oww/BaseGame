using System.Collections;
using System.Collections.Generic;
using UnityEngine;

namespace NMObj
{
    public partial class TestNPC: NPC
    {
        public TestNPC(Vector3 pos, int objType, int id) : base(pos, objType,id){

        }
        public override void Init()
        {
            base.Init();
            GPose = new Vector3(GPose.x, 1, GPose.z);
        }
    }
}
