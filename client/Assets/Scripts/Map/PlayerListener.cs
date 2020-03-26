using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using NMObj;

public class PlayerListener : MonoBehaviour
{

    public Obj PlayerObj;
    // Use this for initialization
    void Start()
    {

    }


    // Update is called once per frame
    void Update()
    {
        ProcessPlayerInput();

    }

    private void ProcessPlayerInput()
    {
        // Exit if we haven't received any input

        if (PlayerObj == null) return;
        bool bMove = false;
        if (Input.GetKey(KeyCode.W))
        {
            PlayerObj.Move(ObjDirection.FORWARD);
            bMove = true;
        }
        else if (Input.GetKey(KeyCode.S))
        {
            PlayerObj.Move(ObjDirection.BACK);
            bMove = true;
        }
        else if (Input.GetKey(KeyCode.A))
        {
            PlayerObj.Move(ObjDirection.LEFT);
            bMove = true;
        }
        else if (Input.GetKey(KeyCode.D))
        {
            PlayerObj.Move(ObjDirection.RIGHT);
            bMove = true;
        }
        if (!bMove && PlayerObj.Status == ObjStatus.MOVE)
        {
            PlayerObj.StopMove();
        }
        //if (Input.GetKeyDown(KeyCode.Space))
        //{
        //    this.GetComponent<Rigidbody>().AddForce(Vector3.up);
        //}

    }
}
