using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.UI;
using Net;
using Google.Protobuf;
using System.Threading;
using WoW.Common.Threading;

public class Login : MonoBehaviour
{
    // Use this for initialization
    enum STATUS
    {
        LOGIN = 1,// 登录
        REG = 2// 注册
    }
    static STATUS status = STATUS.LOGIN;
    static int errorCode = 0;
    public GameObject mapSc;
    public GameObject errorObj;
    public static GameObject MapSc; //todo 使用gameobject.find方法一直找不到main，先暂时通过ui界面获取

    void Start()
    {
        MapSc = mapSc;
        Game.Instance.Status = GameStatus.Login;
        // 注册obj
        GameObject root = GameObject.Find("Login");
        var loginError = root.transform.Find("Login/Error").gameObject;
        GameObjmgr.Register("loginError", loginError);
        var regError = root.transform.Find("Register/Error").gameObject;
        GameObjmgr.Register("regError", regError);

        //注册按钮
        var btn = root.transform.Find("Login/Button").GetComponent<Button>();
        btn.onClick.AddListener(BtnLogin);
        var btn1 = root.transform.Find("Login/BtnReg").GetComponent<Button>();
        btn1.onClick.AddListener(BtnToReg);
        var Rbtn = root.transform.Find("Register/Button").GetComponent<Button>();
        Rbtn.onClick.AddListener(BtnReg);
        var btnReturn = root.transform.Find("Register/BtnReturn").GetComponent<Button>();
        btnReturn.onClick.AddListener(BtnToLogin);

        // 注册 输入文本
        var account = root.transform.Find("Login/Account").gameObject;
        var pwd = root.transform.Find("Login/PWD").gameObject;
        pwd.GetComponent<InputField>().contentType = InputField.ContentType.Password;
        Inputmgr.Register("account", account);
        Inputmgr.Register("pwd", pwd);

        // 注册 输入文本
        var raccount = root.transform.Find("Register/Account").gameObject;
        var rpwd = root.transform.Find("Register/PWD").gameObject;
        rpwd.GetComponent<InputField>().contentType = InputField.ContentType.Password;
        var rpwd1 = root.transform.Find("Register/PWD1").gameObject;
        rpwd1.GetComponent<InputField>().contentType = InputField.ContentType.Password;
        Inputmgr.Register("raccount", raccount, new Inputmgr.ObjCheckFun() { end = CheckRegAcc, change = CheckRegAcc });
        Inputmgr.Register("rpwd", rpwd, new Inputmgr.ObjCheckFun() { end = CheckRegPwd1, change = CheckRegPwd1 });
        Inputmgr.Register("rpwd1", rpwd1, new Inputmgr.ObjCheckFun() { end = CheckRegPwd, change = CheckRegPwd });

    }


    public static void logError(string error)
    {
        switch (status)
        {
            case STATUS.LOGIN:
                GameObjmgr.GetObjByName("loginError").GetComponent<Text>().text = error;
                break;
            case STATUS.REG:
                GameObjmgr.GetObjByName("regError").GetComponent<Text>().text = error;
                break;
        }
    }

    // Update is called once per frame
    void Update()
    {

    }

    public static bool CheckRegAcc()
    {
        string acc = Inputmgr.GetCache("raccount");
        if (acc.Length < 3 || acc.Length > 12)
        {
            logError(ErrorCode.ERRORCODE_3);
            Inputmgr.GetObj("raccount").GetComponent<InputField>().ActivateInputField();
            return false;
        }
        logError(ErrorCode.NONE);
        return true;
    }

    public static bool CheckRegPwd1()
    {
        string pwd = Inputmgr.GetCache("rpwd");
        if (pwd.Length < 6 || pwd.Length > 12)
        {
            logError(ErrorCode.ERRORCODE_2);
            Inputmgr.GetObj("rpwd").GetComponent<InputField>().ActivateInputField();
            return false;
        }
        logError(ErrorCode.NONE);
        return true;
    }

    public static bool CheckRegPwd()
    {
        string pwd = Inputmgr.GetCache("rpwd");
        string pwd1 = Inputmgr.GetCache("rpwd1");
        Debug.Log(pwd + " == " + pwd1);
        if (!pwd.Equals(pwd1))
        {
            logError(ErrorCode.ERRORCODE_1);
            return false;
        }
        logError(ErrorCode.NONE);
        return true;
    }

    public static void BtnReg()
    {

        string account = Inputmgr.GetCache("raccount");
        string pwd = Inputmgr.GetCache("rpwd");

        m_login_tos t1 = new m_login_tos();
        t1.Name = account;
        t1.Pwd = MD5Helper.Md5(pwd);
        t1.Op = (int)STATUS.REG;
        WSocket.Send(t1);
    }

    public static void BtnLogin()
    {
        string account = Inputmgr.GetCache("account");
        string pwd = Inputmgr.GetCache("pwd");

        m_login_tos t1 = new m_login_tos();
        t1.Name = account;
        t1.Pwd = MD5Helper.Md5(pwd);
        t1.Op = (int)STATUS.LOGIN;
        WSocket.Send(t1);
    }

    public static void BtnToReg()
    {
        Transform root = GameObject.Find("Login").transform;
        GameObject Login = root.Find("Login").gameObject;
        Login.SetActive(false);
        GameObject Reg = root.Find("Register").gameObject;
        Reg.SetActive(true);
        logError(ErrorCode.NONE);
        status = STATUS.REG;
    }
    public static void BtnToLogin()
    {
        Transform rlogin = GameObject.Find("Login").transform;
        GameObject Login = rlogin.Find("Login").gameObject;
        Login.SetActive(true);
        GameObject Reg = rlogin.Find("Register").gameObject;
        Reg.SetActive(false);
        logError(ErrorCode.NONE);
        status = STATUS.LOGIN;
    }

    public static void Handler(IMessage msg)
    {
        Debug.Log(msg);
        m_login_toc toc = msg as m_login_toc;
        Debug.Log(toc);
        if (toc.Errcode != 0)
        {
            logError(ErrorCode.GetError(toc.Errcode));
            return;
        }
        if (toc.Op == (int)STATUS.REG)
        {
            BtnToLogin();
        }
        else
        {
            GameObject.Find("Root/Login").gameObject.SetActive(false);
            MapSc.SetActive(true);
            Game.Instance.Status = GameStatus.Game;

        }
    }

}
