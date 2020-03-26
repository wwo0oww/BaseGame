package npc

import (
	"game/core"
)

//npc类型
type NPCType int32

const (
	NPC_PLAYER NPCType = 1
	NPC_TEST NPCType = 2
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

