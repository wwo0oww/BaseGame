using System.Collections;
using System.Collections.Generic;
using UnityEngine;

namespace NMObj {
    public partial class Player : NPC
    {
        public Player(Vector3 pos, int objType, int id) : base(pos, objType, id)
        {

        }

        public override void Init()
        {
            Size = new Vector3(5, 5, 5);
            base.Init();
            GPose = new Vector3(GPose.x,1,GPose.z);
        }
    }
}
