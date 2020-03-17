package heartbeat

import (
	"errors"
	"fmt"
	"game/client"
	"game/config"
	"proto/net"
	"time"
)
func Handle(_ net.MHeartbeatTos, client *client.Client){
	client.Heartindex++
}

func Start(client *client.Client){
	client.Heartindex = 0
	report_ticker := time.NewTicker(time.Second * time.Duration(config.HeartBeat))
	go func() {
		for range report_ticker.C {
			if client.Session() == nil {
				break
			}
			go checkheart(client)
		}
	}()
}

func checkheart(client *client.Client){
	if client.Heartindex == 0{
		client.Stop( errors.New(fmt.Sprintf("no receive heartbeat more than %d",config.HeartBeat)))
	}
	client.Heartindex = 0
}