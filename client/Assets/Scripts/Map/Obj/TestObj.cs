using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using NMObj;

public class TestObj : MonoBehaviour
{

    public Obj testobj;
    // Use this for initialization
    void Start()
    {
        for (int x = 0; x < 10; x++) {
            //Instantiate(Map.Instance.ObjPrefab,new Vector3(x,1,1), Quaternion.identity);
            testobj = ObjMgr.NewObj(new Vector3(x, 0, 0), T_Obj.T_NPC.T_TestNPC.GenObjType(), 1);
            testobj.Init();
        }
        //testobj = ObjMgr.NewObj(new Vector3(122, 117, 115), T_Obj.NPC().TestNPC().GenObjType(), 1);
        // testobj.Init();
        //testobj.Light();

        // var testobj1 = ObjMgr.NewObj(new Vector3(1, 1, 2), T_Obj.NPC().TestNPC().GenObjType(), 2);
        //testobj1.Init();
        // testobj1.Light();
    }

    // Update is called once per frame
    void Update()
    {

    }
}
