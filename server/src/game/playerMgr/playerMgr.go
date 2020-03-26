package playerMgr

import (
	"errors"
	"fmt"
)

type IPlayer interface {
	SendMsg(interface{})
}

var instance *PlayerMgr

type PlayerMgr struct{
	Players map[int32]interface{}
}

func (self*PlayerMgr) AddPlayer(playerID int32,player interface{}){
	self.Players[playerID] = player
}

func  (self*PlayerMgr) SendToPlayer(playerID int32, msg interface{}) error {
	if self.Players[playerID] == nil {
		return errors.New(fmt.Sprint("player: %d not exsit", playerID))
	}
	self.Players[playerID].(IPlayer).SendMsg(msg)
	return nil
}

func Instance()*PlayerMgr{
	if instance == nil{
		instance = &PlayerMgr{Players:make(map[int32]interface{})}
	}
	return instance
}