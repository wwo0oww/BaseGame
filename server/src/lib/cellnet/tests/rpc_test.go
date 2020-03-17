package tests

import (
	"fmt"
	goproto "github.com/golang/protobuf/proto"
	"lib/cellnet"
	"lib/cellnet/peer"
	"lib/cellnet/proc"
	"lib/cellnet/proc/tcp"
	"lib/cellnet/rpc"
	"proto/net"
	prpc "proto/rpc"
	"testing"
	"time"
)

const syncRPC_Address = "127.0.0.1:3344"

var (
	syncRPC_Signal  *SignalTester
	asyncRPC_Signal *SignalTester
	typeRPC_Signal  *SignalTester

	rpc_Acceptor cellnet.Peer
)

//func test(){
//	pkg, err := importer.Default().Import("testing")
//	if err != nil {
//		fmt.Println("error:", err)
//		return
//	}
//	for _, declName := range pkg.Scope().Names() {
//		fmt.Println(declName)
//	}
//}

func rpc_StartServer() {
	//test()
	queue := cellnet.NewEventQueue()

	rpc_Acceptor = peer.NewGenericPeer("tcp.Acceptor", "server", syncRPC_Address, queue)

	proc.BindProcessorHandler(rpc_Acceptor, "tcp.ltv", func(ev cellnet.Event) {
		switch msg := ev.Message().(type) {
		case *TestEchoACK1:
			log.Debugln("server recv rpc ", *msg)
			Data := &prpc.TestEchoACK{}
			goproto.Unmarshal(msg.Value, Data)
			fmt.Println("recv ", Data.Value, " ++ ", Data.Msg)
			ev.(interface {
				Reply(interface{})
			}).Reply(&TestEchoACK1{
				Msg:   msg.Msg,
				Value: msg.Value,
			})
		case *cellnet.SessionConnected:
			fmt.Println("client connected")
		default:
			Msg := &prpc.TestEchoACK{
				Msg:   "type",
				Value: 1234,
			}
			MsgData,err := goproto.Marshal(Msg)
			fmt.Println(MsgData)
			if err != nil{
				fmt.Println("编码错误")
			}else{
				ev.Session().Send(&net.TestEchoACK{
					Msg:   "type",
					Value: 1234,
				})
			}
			fmt.Println("default", ev.Message())
		}

	})
	rpc_Acceptor.Start()

	queue.StartLoop()
}
func syncRPC_OnClientEvent(ev cellnet.Event) {

	switch ev.Message().(type) {
	case *cellnet.SessionConnected:
		for i := 0; i < 2; i++ {

			// 同步阻塞请求必须并发启动，否则客户端无法接收数据
			go func(id int) {

				//Data := &prpc.TestEchoACK{}
				//Data.Value = 110
				//Data.Msg = "hello 110"
				//Data1, err1 := goproto.Marshal(Data)

				result, err := rpc.CallSync(ev.Session(), &TestEchoACK{
					Msg:   "type",
					Value: 1234,
				}, time.Second*5)

				if err != nil {
					syncRPC_Signal.Log(err)
					syncRPC_Signal.FailNow()
					return
				}

				msg := result.(*TestEchoACK1)
				log.Debugln("client sync recv:", msg.Msg, (string)(msg.Value), id*100)

				syncRPC_Signal.Done(id * 100)

			}(i + 1)
		}
	}
}

func asyncRPC_OnClientEvent(ev cellnet.Event) {

	switch ev.Message().(type) {
	case *cellnet.SessionConnected:
		for i := 0; i < 2; i++ {

			copy := i + 1

			rpc.Call(ev.Session(), &TestEchoACK{
				Msg:   "async",
				Value: 1234,
			}, time.Second*5, func(feedback interface{}) {

				switch v := feedback.(type) {
				case error:
					asyncRPC_Signal.Log(v)
					asyncRPC_Signal.FailNow()
				case *TestEchoACK:
					log.Debugln("client sync recv:", v.Msg)
					asyncRPC_Signal.Done(copy)
				}

			})

		}
	}
}

func typeRPC_OnClientEvent(ev cellnet.Event) {

	switch ev.Message().(type) {
	case *cellnet.SessionConnected:
		for i := 0; i < 2; i++ {

			copy := i + 1

			// 注意, 这里不能使用CallType, 异步第一次回来后, 就将rpc上下文清楚,导致第二次之后的回调无法触发, 不属于bug
			rpc.CallSyncType(ev.Session(), &TestEchoACK{
				Msg:   "type",
				Value: 1234,
			}, time.Second*5, func(ack *TestEchoACK, err error) {

				if err != nil {
					panic(err)
				}

				log.Debugln("client type sync recv:", ack)
				typeRPC_Signal.Done(copy)

			})

		}
	}
}

func rpc_StartClient(eventFunc func(event cellnet.Event)) {

	queue := cellnet.NewEventQueue()

	p := peer.NewGenericPeer("tcp.Connector", "client", syncRPC_Address, queue)

	proc.BindProcessorHandler(p, "tcp.ltv.type", eventFunc)

	p.Start()

	queue.StartLoop()
}

func TestSyncRPC(t *testing.T) {
	syncRPC_Signal = NewSignalTester(t)

	rpc_StartServer()

	rpc_StartClient(syncRPC_OnClientEvent)
	syncRPC_Signal.WaitAndExpect("sync not recv data ", 100, 200)
	time.Sleep(time.Second * 100)
	rpc_Acceptor.Stop()
}

func TestASyncRPC(t *testing.T) {

	asyncRPC_Signal = NewSignalTester(t)

	rpc_StartServer()

	rpc_StartClient(asyncRPC_OnClientEvent)
	asyncRPC_Signal.WaitAndExpect("async not recv data ", 1, 2)

	rpc_Acceptor.Stop()
}

func TestTypeRPC(t *testing.T) {

	typeRPC_Signal = NewSignalTester(t)

	rpc_StartServer()

	rpc_StartClient(typeRPC_OnClientEvent)
	typeRPC_Signal.WaitAndExpect("type rpc not recv data ", 1, 2)

	rpc_Acceptor.Stop()
}

func init() {
	// 对TypeRPC增强
	proc.RegisterProcessor("tcp.ltv.type", func(bundle proc.ProcessorBundle, userCallback cellnet.EventCallback, args ...interface{}) {

		bundle.SetTransmitter(new(tcp.TCPMessageTransmitter))
		bundle.SetHooker(proc.NewMultiHooker(new(tcp.MsgHooker), new(rpc.TypeRPCHooker)))
		bundle.SetCallback(proc.NewQueuedEventCallback(userCallback))
	})

}
