package router

import (
	"fmt"
	"game/client"
	"game/heartbeat"
	"game/login"
	"lib/cellnet"
	"proto/net"
)

func Router(ev cellnet.Event) {
	client := (ev.Session()).(interface {
		GetClient() *client.Client
	}).GetClient()
	switch msg := ev.Message().(type) {
	case *cellnet.SessionConnected:
		fmt.Println("client connected")
	case *net.MLoginTos:
		login.Handle(*(ev.Message()).(*net.MLoginTos), client)
	case *net.MHeartbeatTos:
		heartbeat.Handle(*(ev.Message()).(*net.MHeartbeatTos), client)
	default:
		fmt.Println("unknown msg", msg)
	}
}
