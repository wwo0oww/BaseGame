using System.Collections;
using System.Collections.Generic;
using System.Threading;
using UnityEngine;

public class Common {
    public  delegate void DelayEvent();
    public static void Delay(int time,DelayEvent even) {
        Thread.Sleep(time);
        even();
    }


    /// <summary>
    /// 两个点的角度
    /// 0°正右
    /// 90°正上
    /// -90°正下
    /// </summary>
    /// <param name="p1"></param>
    /// <param name="p2"></param>
    /// <returns></returns>
    /// 
    public static float PointToAngle(Vector2 p1, Vector2 p2)
    {
        Vector2 p;
        p.x = p2.x - p1.x;
        p.y = p2.y - p1.y;
        return Mathf.Atan2(p.y, p.x) * 180 / Mathf.PI;
    }
    /// <summary>
    /// -180-180 转为 0-360
    /// </summary>
    /// <param name="angle"></param>
    /// <returns></returns>
    public static float Angle180To360(float angle)
    {
        if (angle >= 0 && angle <= 180)
            return angle;
        else
            return 360 + angle;
    }
    /// <summary>
    /// 返回Unity的角度
    /// </summary>
    /// <param name="p1"></param>
    /// <param name="p2"></param>
    /// <returns></returns>
    public static float GetUnityDirection(Vector2 p1, Vector2 p2)
    {
        float angle = Angle180To360(PointToAngle(p1, p2));
        Debug.Log(angle);
        float temp = 360 * 0.125f;//分为8个方向
        float dir = 0;
        for (int i = 0; i < 8; i++)
        {
            if (angle >= (i * temp) - (temp * 0.5f) && angle < (i * temp) + (temp * 0.5f))
            {
                dir = i * temp;
                break;
            }
        }
        return dir;
    }

    /// <summary>
    /// 新建线程
    /// </summary>
    /// <param name="p1"></param>
    /// <param name="p2"></param>
    /// <returns></returns>
    public static void NewThread(ThreadStart func ) {
        Thread thr = new Thread(func);
        thr.Start();
    }
}
