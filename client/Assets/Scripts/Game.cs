using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;
using WoW.Common.Threading;

public enum GameStatus {
    None,
    Login,
    Game
}

public class Game : MonoBehaviour {
    public static int nFrame = 80;
    public static int nServerFrame = 20;// 后端为1s 20帧
    public bool Debug = true;
    public static Game Instance;
    public static int nDurFrame = 6;
    public static int SinglePlayer = 0;// 0 联网 1 单机
    private GameStatus m_status;
    void Awake()
    {
        Instance = this;
        Application.targetFrameRate = nFrame;

    }
	// Use this for initialization
	void Start () {

        InitTestLog();

    }


	// Update is called once per frame
	void Update () {
		
	}

    void InitTestLog()
    {
        var LogInfo = GameObject.Find("ShowTab/LogInfo").gameObject;
        GameObjmgr.Register("LogInfo", LogInfo);
    }

    public GameStatus Status
    {
        get { return m_status; }
        set {m_status = value; }
    }

    public static void Log<T>(T log)
    {
        var data = new TProto();
        Dispatcher.RunMethmod(new Dispatcher.Action((x) => { DoLog(log.ToString()); }), data);
    }
    static void DoLog(string log)
    {
        if (Game.Instance.Debug)
        {
            Text text = GameObjmgr.GetObjByName("LogInfo").GetComponent<Text>();
            text.text += "\n"+System.DateTime.Now + ":"+ log;
        }
    }
}
