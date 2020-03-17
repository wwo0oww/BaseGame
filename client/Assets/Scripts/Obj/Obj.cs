using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class Obj : MonoBehaviour
{
    private int m_Angular = 0;
    public RectTransform p1;
    public RectTransform p2;

    // Use this for initialization
    void Start()
    {

    }

    // Update is called once per frame
    void Update()
    {
        //gameObject.transform.position += Vector3.left / Game.nDurFrame;
        if (Time.frameCount % Game.nDurFrame == 0)
        {
            var rotation = gameObject.transform.rotation;
            Debug.Log(Common.GetUnityDirection(new Vector2(rotation.x, rotation.y), Vector2.zero));
        }

        //V3.Set(0, 0, GetUnityDirection(p1.anchoredPosition, p2.anchoredPosition) - 90);
        //p2.localEulerAngles = V3;
    }

}
