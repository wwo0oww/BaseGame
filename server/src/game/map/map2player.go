package emap

import (
	"game"
	"game/config"
	"game/map/area"
	"game/playerMgr"
	"proto"
)

func (self *Map) SendMsgToMap(msg interface{}) {
	if isLocalMsg() {
		playerMgr.Instance().SendToPlayer(msg.(*area.AreaMsg).PlayerID, msg.(*area.AreaMsg).Msg)
		return
	}
	game.Node().SendMsgToGameNode(proto.EncodeMapMsg(self, msg))
}

func isLocalMsg() bool {
	//判断是否是rpc消息
	return game.Node().GetName() == config.GetMapNodeName()
}
