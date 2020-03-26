package player

import (
	"fmt"
	"game"
	"game/config"
	"game/core"
	"game/playerMgr"
	"strconv"
)

func NewPlayerID() string {
	return fmt.Sprintf("%.2d%.8d", config.GetServerID(), game.MaxPlayerNum)
}
func NewPlayer(playerID string) *Player {
	IPlayerID, err := strconv.Atoi(playerID)
	if err != nil {
		IPlayerID = 0
	}
	player := playerMgr.Instance().Players[int32(IPlayerID)]
	if player != nil {
		return playerMgr.Instance().Players[int32(IPlayerID)].(*Player)
	}

	player1 := &Player{Player_id: int32(IPlayerID), Pos: InitPlayerPosition()}
	playerMgr.Instance().AddPlayer(player1.Player_id, player1)
	return player1
}

func InitPlayerPosition() core.Position {
	w, h := config.GetMapWH()
	return core.Position{X: config.GetAreaSize() * w / 2, Y: config.GetAreaSize() * h / 2}
}
