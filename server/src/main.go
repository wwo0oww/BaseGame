package main

import (
	"fmt"
	"game"
	"game/router"
	"game/sql"
	"lib/cellnet"
	"lib/cellnet/peer"
	_ "lib/cellnet/peer/tcp"
	"lib/cellnet/proc"
	_ "lib/cellnet/proc/tcp"
	_ "proto"
	"time"
)

const Address = "169.254.70.149:3344"

var (
	Acceptor cellnet.Peer
	status   bool
)


func main() {
	status = true
	StartServer()
	sql.Startsql()
	game.StartGame()
	//tos  := net.MLoginTos{Pwd:"22222222",Name:"2222",}
	go test()
	for {
		if !status {
			break
		}
		time.Sleep(time.Second)
	}
}

func StartServer() {
	queue := cellnet.NewEventQueue()


	Acceptor = peer.NewGenericPeer("tcp.Acceptor", "server", Address, queue)

	proc.BindProcessorHandler(Acceptor, "tcp.ltv", router.Router)

	Acceptor.Start()

	queue.StartLoop()

}

func test() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("xx", err)
		}
	}()
	ch := make(chan []byte, 2)
	ch <- []byte("mgs 1")
	ch <- []byte("mgs 2")
	fmt.Println("end")

}
