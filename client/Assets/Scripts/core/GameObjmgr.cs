using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class GameObjmgr : MonoBehaviour
{
    static Dictionary<string, GameObject> objMap = new Dictionary<string, GameObject>();
    // Use this for initialization
    void Start()
    {

    }

    // Update is called once per frame
    void Update()
    {

    }
    public static void Register(string name, GameObject obj)
    {
        objMap.Add(name, obj);
    }

    public static void Unregister(string name)
    {
        objMap.Remove(name);
    }

    public static GameObject GetObjByName(string name)
    {
        GameObject obj = null;
        objMap.TryGetValue(name, out obj);
        return obj;
    }

}
