using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;


public class Inputmgr : MonoBehaviour
{
    public delegate void dealErrorFun(string error);

    public delegate bool CheckFun();

    static Dictionary<GameObject, string> InputMap = new Dictionary<GameObject, string>();
    static Dictionary<string, GameObject> InputMap1 = new Dictionary<string, GameObject>();


    static Dictionary<string, dealErrorFun> InputErrMap = new Dictionary<string, dealErrorFun>();

    static Dictionary<string, CheckFun> EndMap = new Dictionary<string, CheckFun>();
    static Dictionary<string, CheckFun> ChangedMap = new Dictionary<string, CheckFun>();

    static Hashtable InputCache = new Hashtable();

    public struct ObjCheckFun
    {
        public CheckFun end;
        public CheckFun change;
    }

    // Use this for initialization
    void Start()
    {

    }


    public static void RegisterListener(GameObject obj, string Name)
    {
        var Input = obj.GetComponent<InputField>();
        Input.onValueChanged.AddListener(
                    delegate (string str)
                    {
                        string outStr;
                        Changed_Value(Name, str, out outStr);
                    });
        Input.onEndEdit.AddListener(
                    delegate (string str)
                    {
                        string outStr;
                        End_Value(Name, str, out outStr);
                    });
    }

    // Update is called once per frame
    void Update()
    {

    }

    public static string GetCache(string name)
    {
        string str = InputCache[name] as string;
        return str == null ? "" : str;
    }

    public static GameObject GetObj(string name)
    {
        return InputMap1[name] as GameObject;
    }

    public static bool Changed_Value(string name, string str, out string outStr)
    {
        InputCache[name] = str;
       // print("正在输入:" + str);
        CheckFun Checkfun;
        if (ChangedMap.TryGetValue(name, out Checkfun))
        {
            if (!Checkfun())
            {
                outStr = str;
                return false;
            }
        }
        return CheckWorld(name, str, out outStr);

    }
    public static bool End_Value(string name, string str, out string outStr)
    {
        InputCache[name] = str;
      //  print("最终文本:" + str);
        CheckFun Checkfun;
        if (EndMap.TryGetValue(name, out Checkfun))
        {
            if (!Checkfun())
            {
                outStr = str;
                return false;
            }
        }
        return CheckWorld(name, str, out outStr);
    }
    public static void Register(string name, GameObject obj)
    {
        Register(name, obj, new ObjCheckFun());
    }
    public static void Register(string name, GameObject obj, ObjCheckFun fun)
    {
        if (fun.end != null)
        {
            EndMap[name] = fun.end;
        }
        if (fun.change != null)
        {
            ChangedMap[name] = fun.change;
        }
        InputMap.Add(obj, name);
        InputMap1.Add(name, obj);
        RegisterListener(obj, name);

    }

    public static void RegisterError(string name, dealErrorFun Func)
    {
        InputErrMap.Add(name, Func);
    }

    public static bool CheckWorld(string name, string str, out string outStr)
    {
        // todo 敏感词
        if (SensitiveWord.CheckWord(str, out outStr))
        {
            return true;
        }
        else
        {
            dealErrorFun func;
            if (InputErrMap.TryGetValue(name, out func))
            {
                func(SensitiveWord.strError_1);
            }
            return false;
        }
    }

}
