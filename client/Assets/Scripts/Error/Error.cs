using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using WoW.Common.Threading;

public class Error : MonoBehaviour
{
    public static void logErrorMain(string error)
    {
        var data = new TProto();
        Dispatcher.RunMethmod(new Dispatcher.Action((x) => { logError(error); }), data);
    }
    public static void logError(string error)
    {
        switch (Game.Instance.Status)
        {
            case GameStatus.Login:
                Login.logError(error);
                break;
            default:
                break;
        }

    }



}
