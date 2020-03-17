package player

import (
	"game"
	"game/config"
	"game/map"
	"proto"
)

func (self *Player) SendMsgToMap(Msg interface{}) {
	if isLocalMsg() {
		emap.GetMapInstance().AddMessage(Msg)
		return
	}
	game.Node().SendMsgToMapNode(proto.EncodeMapMsg(self, Msg))
}

func isLocalMsg() bool {
	//判断是否是rpc消息
	return game.Node().GetName() == config.GetMapNodeName()
}
