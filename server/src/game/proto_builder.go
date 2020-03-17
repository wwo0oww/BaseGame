package game

import (
	"game/core"
	"proto/net"
)

func Player2PPlayer(rawPlayer interface{}) *net.PPlayer {
	player := rawPlayer.(struct{Player_id string;Pos core.Position})
	return &net.PPlayer{
		PlayerID: player.Player_id,
		Position: Postion2PPosition(player.Pos),
	}
}

func Postion2PPosition(pos core.Position) *net.PPosition {
	return &net.PPosition{
		X: pos.X,
		Y: pos.Y,
	}
}
