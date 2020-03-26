package player

import (
	"fmt"
	"game/client"
	"game/core"
)

type Player struct {
	Player_id  int32
	Pos        core.Position
	client     *client.Client
	OnlineFlag int32
}

func (self*Player)GetPlayerID() int32{
	return self.Player_id
}

func (self *Player) JoinMap() {

}

func (self *Player) ToString() string {
	return fmt.Sprintf("Player_%.10d", self.Player_id)
}

func (self *Player) GetPos() core.Position {
	return self.Pos
}

func (self *Player) SetClient(client *client.Client) {
	client.PlayerPr = self
	self.client = client
}

func (self *Player) SendMsg(msg interface{}) {
	if self.client.Session() != nil {
		self.client.Session().Send(msg)
	}
}

func (self *Player) OnLine() {
	self.OnlineFlag = 1
}

func (self *Player) OffLine() {
	self.OnlineFlag = 0
}
