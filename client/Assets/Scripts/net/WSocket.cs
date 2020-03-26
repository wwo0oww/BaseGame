using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using System;
using System.ComponentModel;
using System.Linq;
using System.Text;
using System.Net.Sockets;
using System.Net;
using System.Threading;
using System.IO;
using Net;
using Google.Protobuf;     //引用DLL
using WoW.Common.Threading;

public class WSocket : MonoBehaviour
{

    const Int16 bodySize = 2;
    const Int16 msgIDSize = 2;

    int MaxConSec = 10;
    int NowConSec = 0;
    int NowConCount = 1;
    //创建连接的Socket
    static Socket socketSend;
    //创建接收客户端发送消息的线程
    Thread threadReceive;
    void Start()
    {
        NowConCount = 1;
        ErrorCode.InitCode();
        gen_proto.InitProtoId();
        Thread th = new Thread(new ThreadStart(trycon)); //创建线程
        th.Start(); //启动线程

    }
    void trycon()
    {
       // if (Game.Instance.Status != GameStatus.Login) return;

        try
        {
            socketSend = new Socket(AddressFamily.InterNetwork, SocketType.Stream, ProtocolType.Tcp);
            IPAddress ip = IPAddress.Parse("169.254.70.149");
            socketSend.Connect(ip, Convert.ToInt32(3344));

            Debug.Log("连接成功");

            //开启一个新的线程不停的接收服务器发送消息的线程
            threadReceive = new Thread(new ThreadStart(Receive));
            //设置为后台线程
            threadReceive.IsBackground = true;
            threadReceive.Start();

            heartbeat.startheart(); // 开启心跳

            Error.logErrorMain(String.Format("请输入用户名和密码"));

            Thread.Sleep(3000);
            Error.logErrorMain(String.Format(""));
        }
        catch (Exception ex)
        {
            Debug.Log("连接服务端出错:" + ex.ToString());
            NowConSec = 0;
            MaxConSec = MaxConSec * NowConCount;
            NowConCount++;
            while (NowConSec < MaxConSec)
            {
                Error.logErrorMain(String.Format("无法连接到服务器，{0:D2}秒后重新尝试连接", MaxConSec - NowConSec));
                NowConSec++;
                Thread.Sleep(1000);
            }
            Error.logErrorMain(String.Format("正在尝试重新连接服务器"));
            trycon();
        }
    }
    /// <summary>
    /// 接口服务器发送的消息
    /// </summary>
    private void Receive()
    {
        try
        {
            while (true)
            {
                int nSize = 2048;
                byte[] buffer = new byte[nSize];
                //实际接收到的字节数
                int r = socketSend.Receive(buffer, bodySize, SocketFlags.None);
                //socketSend.Close();
                if (r == 0)
                {
                    Debug.Log("错误的包头长度:" + r);
                    continue;
                }
                else
                {

                    Int16 Size = BitConverter.ToInt16(buffer, 0);
                    if (Size > nSize) {
                        buffer = new byte[Size];
                    }
                    r = socketSend.Receive(buffer, Size, SocketFlags.None);
                    if (r != Size)
                    {
                        Debug.Log("错误的包体长度:" + r);
                    }
                    else
                    {

                        byte[] buffer1 = new byte[msgIDSize];
                        for (int i = 0; i < msgIDSize; i++)
                        {
                            buffer1[i] = buffer[i];
                        }
                        Int16 MsgID = BitConverter.ToInt16(buffer1, 0);


                        //sendBytesCount += (UInt32)System.Text.Encoding.Default.GetBytes(buffer).Length;\
                        byte[] buffer2 = new byte[Size - msgIDSize];
                        for (int i = msgIDSize; i < r; i++)
                        {
                            buffer2[i - msgIDSize] = buffer[i];
                        }
                        IMessage IMsg = gen_proto.protoTypeMap[MsgID] as IMessage;
                        TestEchoACK p1 = new TestEchoACK();
                        var t = IMsg.Descriptor.Parser.ParseFrom(buffer2);
                        var maction = new Dispatcher.MAction();
                        maction.action = rounter.GetHandler(t);
                        maction.msg = t;
                        Dispatcher.Run(maction);
                    }

                }


            }
        }
        catch (Exception ex)
        {
            Debug.Log("接收服务端发送的消息出错:" + ex.ToString());
        }
    }

    public static void Send(IMessage raw)
    {
        byte[] data = raw.ToByteArray();
        Int16 msgID = (Int16)gen_proto.protoIDMap[raw.GetType()];

        var msgLen = data.Length + msgIDSize;
        byte[] sendData = new byte[bodySize + msgIDSize + data.Length];
        byte[] BMsgLen = new byte[bodySize];
        byte[] BProtoID = new byte[msgIDSize];

        BMsgLen = BitConverter.GetBytes(msgLen);
        BProtoID = BitConverter.GetBytes(msgID);
        Array.Copy(BMsgLen, sendData, bodySize);

        Buffer.BlockCopy(BProtoID, 0, sendData, bodySize, msgIDSize);

        Buffer.BlockCopy(data, 0, sendData, bodySize + msgIDSize, data.Length);

        Array.Copy(BMsgLen, sendData, bodySize);

        socketSend.Send(sendData, SocketFlags.None);
    }



    // Update is called once per frame
    void Update()
    {

    }



}

