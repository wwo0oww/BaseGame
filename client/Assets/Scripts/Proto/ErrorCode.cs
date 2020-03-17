using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class ErrorCode : MonoBehaviour
{
    public static string NONE = "";
    public static string ERRORCODE_1 = "两次密码不一致";
    public static string ERRORCODE_2 = "密码必须为6~12个字符";
    public static string ERRORCODE_3 = "玩家名必须为3~12个字符";
    static Hashtable ErrorMap = new Hashtable();
    public static void InitCode()
    {
        int start_index = 1;
        ErrorMap.Add(start_index++, "系统错误");
        ErrorMap.Add(start_index++, "用户名或密码错误");
        ErrorMap.Add(start_index++, "用户名已被使用");
        ErrorMap.Add(start_index++, "用户名不存在");
    }
    public static string GetError(int ErrorCode)
    {
        string str = ErrorMap[ErrorCode] as string;
        return str == null ? "未知错误" : str;
    }
}
