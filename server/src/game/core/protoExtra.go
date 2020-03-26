package core

import (
	"game/config"
	"proto/net"
)

type ObjUpdateType int32

const (
	ObjUpdateType_Add ObjUpdateType = 1 // 增加
	ObjUpdateType_Upt                   // 更新
	ObjUpdateType_Del                   // 删除
)

func Pos2PPos(area interface{}, position Position) *net.PPosition {
	areaPos := area.(interface{ GetPos() Position }).GetPos()
	size := config.GetAreaSize()
	return &net.PPosition{X: areaPos.X*size + position.X, Y: areaPos.Y*size + position.Y}
}
