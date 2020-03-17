using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class Player : MonoBehaviour {

    public Camera hitcam;
    // Use this for initialization
    void Start()
    {
        
    }


    // Update is called once per frame
    void Update () {
        ProcessPlayerInput();

    }

    private void ProcessPlayerInput()
    {
        // Exit if we haven't received any input
        if (!Input.anyKey)
        {
            return;
        }

        //if (Input.inputString.Contains("t"))
        //{
        //    RaycastHit hit;
        //    Ray ray = Camera.main.ScreenPointToRay(new Vector3(Screen.width / 2.0f, Screen.height / 2.0f, 0));
        //    if (Physics.Raycast(ray, out hit, 4.0f))
        //    {
        //        WorldData.SetBlockLightWithRegeneration((int)hit.point.x, (int)hit.point.z, (int)hit.point.y, 255);
        //        m_World.RegenerateChunks();
        //    }
        //}
        int textspeed = 5;
        if (Input.GetKey(KeyCode.W))
        {
            RaycastHit hit;
            Ray ray = hitcam.ScreenPointToRay(new Vector3(Screen.width / 2.0f, Screen.height / 2.0f, 0));
            if (Physics.Raycast(ray, out hit, 4.0f))
            {
                Debug.Log("yyyy");
            }
            else {
                Debug.Log("xxxx");
            }
            this.transform.Translate(Vector3.forward * Time.deltaTime * textspeed);
        }
        else if (Input.GetKey(KeyCode.S))
        {
            this.transform.Translate(Vector3.back * Time.deltaTime * textspeed);
        }
        else if (Input.GetKey(KeyCode.A))
        {
            this.transform.Translate(Vector3.up * Time.deltaTime * textspeed);
        }
        else if (Input.GetKey(KeyCode.D))
        {
            this.transform.Translate(Vector3.down * Time.deltaTime * textspeed);
        }
        if (Input.GetKeyDown(KeyCode.Space))
        {
            this.GetComponent<Rigidbody>().AddForce(Vector3.up);
        }

        //if (Input.GetKey(KeyCode.Mouse0))
        //{
        //    //m_World.RemoveBlockAt(blockHitPoint);
        //    RaycastHit hit;
        //    Ray ray = Camera.main.ScreenPointToRay(new Vector3(Screen.width / 2.0f, Screen.height / 2.0f, 0));
        //    if (Physics.Raycast(ray, out hit, 4.0f))
        //    {
        //        // Get the hit position...plus a little more. 
        //        Vector3 worldHitPosition = hit.point + (ray.direction.normalized * 0.01f);
        //        // We have the 'global' block position, we need to convert that to the local map position.
        //        Vector3i blockMapPosition = m_World.GlobalToLocalMapBlockPosition(new Vector3i((int)worldHitPosition.x, (int)worldHitPosition.z, (int)worldHitPosition.y));
        //        m_World.Dig(blockMapPosition, worldHitPosition);
        //    }
        //}

    }
}
