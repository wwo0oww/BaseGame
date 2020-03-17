using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class SensitiveWord{

    public static string strError_1 = "存在敏感词";

	public static bool CheckWord(string str,out string outStr) 
    {
        // todo 敏感词检测
        outStr = str;
        return true;
    }
}
