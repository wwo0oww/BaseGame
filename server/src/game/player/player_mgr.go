package player

import (
	"fmt"
	"game"
	"game/config"
	"game/core"
	"strconv"
)

func NewPlayerID() string {
	return fmt.Sprintf("%.2d%.8d", config.GetServerID(), game.MaxPlayerNum)
}

func NewPlayer(PlayerID string) Player {
	IPlayerID, err := strconv.Atoi(PlayerID)
	if err != nil {
		IPlayerID = 0
	}
	return Player{Player_id: int32(IPlayerID), Pos: InitPlayerPosition()}
}

func InitPlayerPosition() core.Position {
	return core.Position{X: 0, Y: 0}
}
