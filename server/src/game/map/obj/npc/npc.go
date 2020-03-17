package npc

import (
	"game/core"
)

type IMNPC interface {
	ToString() string
	GetPos() core.Position
	GetSize() core.Position
}
type MNPC struct {
	npc_id string
	Pos    core.Position
}

