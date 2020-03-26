package emap

import (
	"game/client"
	"game/core"
	mMsg "game/map/msg"
	"game/map/obj"
	"proto/net"
)

func Handle(tos net.MObjMoveTos, client *client.Client) {
	if tos.Direction > core.DirectionNum || tos.Direction == (int32)(core.DIRECTION_NONE) {
		return // todo 统一错误处理函数
	}
	client.PlayerPr.(IPlayer).SendMsgToMap(&mMsg.ObjMove{Obj: &obj.Obj{
		ObjID: client.PlayerPr.(IPlayer).GetPlayerID()},
		Direction: (core.DIRECTION)(tos.Direction)})
}
