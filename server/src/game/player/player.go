package player

import (
	"fmt"
	"game/core"
)

type Player struct {
	Player_id int32
	Pos       core.Position
}

func (self *Player) JoinMap() {

}

func (self *Player) ToString() string {
	return fmt.Sprintf("Player_%.10d", self.Player_id)
}

func (self *Player) GetPos() core.Position {
	return self.Pos
}
