using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Linq;
using System.Text;
using System.Net.Sockets;
using System.Net;
using System.Threading;
using System.IO;
using Net;  
using Google.Protobuf;     //引用DLL

public class socket : MonoBehaviour {

	const Int16 bodysize = 2;
	const Int16 msgIDSize = 2;

	

    //创建连接的Socket
    Socket socketSend;
	//创建接收客户端发送消息的线程
	Thread threadReceive;
	void Start () {
		try
		{
            socketSend = new Socket(AddressFamily.InterNetwork, SocketType.Stream, ProtocolType.Tcp);
			IPAddress ip = IPAddress.Parse("127.0.0.1");
			socketSend.Connect(ip, Convert.ToInt32(3344));

			Debug.Log("连接成功");
			
			//开启一个新的线程不停的接收服务器发送消息的线程
			threadReceive = new Thread(new ThreadStart(Receive));
			//设置为后台线程
			threadReceive.IsBackground = true;
			threadReceive.Start();
		}
		catch (Exception ex)
		{
			Debug.Log("连接服务端出错:" + ex.ToString());
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
				byte[] buffer = new byte[2048];
				//实际接收到的字节数
				int r = socketSend.Receive(buffer,bodysize,SocketFlags.None);
				Debug.Log("收到服务端消息:"+buffer+" 长度："+r);

				if (r == 0)
				{
					Debug.Log("错误的包头长度:" +r);
					break;
				}
				else
				{
					for(int i = 0;i < msgIDSize;i++){
						Debug.Log(buffer[i]);
					}
					Int16 Size = BitConverter.ToInt16( buffer,0 );
					Debug.Log("包大小为"+Size);
					r = socketSend.Receive(buffer,Size,SocketFlags.None);

					if (r != Size){
						Debug.Log("错误的包体长度:" +r);
					}
					else{

						byte[] buffer1 = new byte[msgIDSize];
						for(int i = 0;i < msgIDSize;i++){
							buffer1[i] = buffer[i];
						}
						Int16 MsgID = BitConverter.ToInt16( buffer1,0 );
						Debug.Log("消息id为" +MsgID);
						//sendBytesCount += (UInt32)System.Text.Encoding.Default.GetBytes(buffer).Length;\
						byte[] buffer2 = new byte[Size - msgIDSize];
						for(int i = msgIDSize;i < r;i++){
							buffer2[i - msgIDSize] = buffer[i]; 
						}
                        Debug.Log(BitConverter.ToString(buffer2));
                        IMessage IMperson = new TestEchoACK();
                        //TestEchoACK t1 = new TestEchoACK();
                        //t1.Msg = "type";
                        //t1.Value = 1234;
                        //byte[] databytes = t1.ToByteArray();
                        //Debug.Log(BitConverter.ToString(databytes));
                        //Debug.Log((TestEchoACK)IMperson.Descriptor.Parser.ParseFrom(databytes));
                        TestEchoACK p1 = new TestEchoACK();
                        p1 = (TestEchoACK)IMperson.Descriptor.Parser.ParseFrom(buffer2);
						Debug.Log(p1.Msg);
					}

				}
				
				
			}
		}
		catch (Exception ex)
		{
			Debug.Log("接收服务端发送的消息出错:" + ex.ToString());
		}
	}
	
	
	// Update is called once per frame
	void Update () {
		
	}
}

	